import { BadRequestException, HttpException, Injectable } from '@nestjs/common';
import { UserLoginDto, UserRegisterDto } from './interfaces/user.dto';
import { UsersService } from '../users/users.service';
import { JwtService } from '@nestjs/jwt';
import { UserPayload } from './interfaces/auth.interface';
import { BcryptService } from './bcrypt.service';

@Injectable()
export class AuthService {
  constructor(
    private readonly userService: UsersService,
    private readonly jwtService: JwtService,
    private readonly bcryptService: BcryptService,
  ) {}
  async login(userLogin: UserLoginDto): Promise<Record<string, any>> {
    try {
      if (!userLogin || !userLogin.email || !userLogin.password) {
        throw new HttpException('Unauthenticated', 401);
      }

      const user = await this.userService.getUserByEmail(userLogin.email);

      if (!user) throw new BadRequestException('User Not Found');

      const valid = await this.bcryptService.verify(
        userLogin.password,
        user.password,
      );

      if (!valid) throw new BadRequestException('Email/Password is wrong!');

      const token = await this.jwtService.signAsync({
        id: user.id,
        email: user.email,
      } as UserPayload);

      return {
        token,
        name: user.name,
      };
    } catch (err) {
      throw err;
    }
  }

  async register(userRegister: UserRegisterDto) {
    const user = await this.userService.getUserByEmail(userRegister.email);

    if (user) {
      throw new BadRequestException('You are already registered');
    }

    userRegister.password = await this.bcryptService.hashing(
      userRegister.password,
    );

    const generatedUser = await this.userService.saveUser(userRegister);

    const token = await this.jwtService.signAsync({
      id: generatedUser.id,
      email: generatedUser.email,
    } as UserPayload);

    return { token, name: userRegister.name };
  }

  async authenticate(token: string): Promise<boolean> {
    const userPayload: UserPayload = await this.jwtService.verifyAsync(token);

    await this.userService.getUserById(userPayload.id);

    return true;
  }
}

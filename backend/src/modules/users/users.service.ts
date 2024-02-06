import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './entities/user.entity';
import { Repository } from 'typeorm';
import { UserRegisterDto } from '../auth/interfaces/user.dto';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User) private readonly userRepo: Repository<User>,
  ) {}

  getAllUsers() {
    return this.userRepo.createQueryBuilder('user').getMany();
  }

  getUserById(id: number) {
    return this.userRepo
      .createQueryBuilder('user')
      .where('user.id = :id', { id })
      .getOneOrFail();
  }

  getUserByEmail(email: string) {
    return this.userRepo
      .createQueryBuilder('user')
      .where('user.email=:email', { email })
      .getOne();
  }

  async saveUser(user: UserRegisterDto) {
    const builtUser = this.userRepo.create(user);

    await this.userRepo
      .createQueryBuilder()
      .insert()
      .into(User)
      .values(builtUser)
      .execute();

    return builtUser;
  }
}

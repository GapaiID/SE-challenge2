import { Body, Controller, Post } from '@nestjs/common';
import { UserLoginDto, UserRegisterDto } from './interfaces/user.dto';
import { AuthService } from './auth.service';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('login')
  login(@Body() body: UserLoginDto) {
    return this.authService.login(body);
  }

  @Post('register')
  register(@Body() body: UserRegisterDto) {
    return this.authService.register(body);
  }
}

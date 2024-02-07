import { Global, Module, OnModuleInit } from '@nestjs/common';
import { AuthService } from './auth.service';
import { UsersModule } from '../users/users.module';
import { AuthGuard } from './auth.guard';
import { JwtModule } from '@nestjs/jwt';
import { AuthController } from './auth.controller';
import { BcryptService } from './bcrypt.service';

@Global()
@Module({
  imports: [
    UsersModule,
    JwtModule.registerAsync({
      useFactory: () => ({
        secret: process.env.JWT_SECRET,
        signOptions: {
          expiresIn: '6h',
        },
      }),
    }),
  ],
  providers: [AuthService, AuthGuard, BcryptService],
  exports: [AuthService],
  controllers: [AuthController],
})
export class AuthModule implements OnModuleInit {
  constructor(private readonly bcryptService: BcryptService) {}

  async onModuleInit() {
    await this.bcryptService.generateSalt();
  }
}

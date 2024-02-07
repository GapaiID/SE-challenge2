import {
  CanActivate,
  ExecutionContext,
  HttpException,
  Injectable,
} from '@nestjs/common';
import { AuthService } from './auth.service';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private readonly authService: AuthService) {}

  async canActivate(ctx: ExecutionContext): Promise<boolean> {
    const req = ctx.switchToHttp().getRequest();

    const { authorization } = req.headers;

    if (!authorization) throw new HttpException('Unauthenticated', 401);

    const [type, token] = authorization.split(' ');

    if (!/bearer/is.test(type)) {
      throw new HttpException('Unauthenticated', 401);
    }

    return await this.authService.authenticate(token);
  }
}

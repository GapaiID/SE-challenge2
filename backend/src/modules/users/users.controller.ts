import {
  BadRequestException,
  Controller,
  Get,
  Param,
  UseGuards,
} from '@nestjs/common';
import { UsersService } from './users.service';
import { AuthGuard } from '../auth/auth.guard';

@UseGuards(AuthGuard)
@Controller('users')
export class UsersController {
  constructor(private readonly userService: UsersService) {}

  @Get('/')
  getUsers() {
    return this.userService.getAllUsers();
  }

  @Get('/:id')
  getUserById(@Param('id') userId: number) {
    userId = +userId;

    if (isNaN(userId)) {
      throw new BadRequestException('Invalid User ID');
    }

    if (userId < 1) {
      throw new BadRequestException('Invalid User ID');
    }

    return this.userService.getUserById(userId);
  }
}

import { Controller, Get, Param, UseGuards } from '@nestjs/common';
import { PostsService } from './posts.service';
import { AuthGuard } from '../auth/auth.guard';

@UseGuards(AuthGuard)
@Controller('posts')
export class PostsController {
  constructor(private readonly postsService: PostsService) {}

  @Get('/')
  getAllPost() {
    return this.postsService.getAllPost();
  }

  @Get('/:id')
  getPostById(@Param('id') postId: string) {
    return this.postsService.getPostById(postId);
  }
}

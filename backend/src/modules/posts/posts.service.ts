import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Post } from './entities/post.entity';

@Injectable()
export class PostsService {
  constructor(
    @InjectRepository(Post) private readonly postRepo: Repository<Post>,
  ) {}

  async getAllPost() {
    const data = await this.postRepo.query(`
select 
  p.id, 
  p.title, 
  p."content", 
  p."image_src",
  string_agg(t.tag, '|') tags, 
  c.id as creator_id, 
  c."name" as creator_name, 
  p.created_at, 
  p.updated_at 
from posts p 
left join post_tags pt on pt.post_id = p.id 
left join  users c on c.id = p.creator_id 
left join  tags t on t.id = pt.tag_id
group by 
  p.id,
  p.title,
  p."content",
  c.id,
  c."name",
  p.created_at,
  p.updated_at 
order by p.id ASC
;`);

    return data.map((el) => {
      const tagSet = new Set(el.tags.split('|'));
      el.tags = [...tagSet];

      return el;
    });
  }
}

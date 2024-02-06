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

  async getPostById(id: string) {
    const data = await this.postRepo.query(
      `
SELECT 
  p.id,
  p.title,
  p.content, 
  c.name as creator_name,
  c.id as creator_id,
  (
    string_agg(t.id || ':' || t.tag, ', ')
  ) as tags,
  p.created_at,
  p.updated_at
FROM posts p
LEFT JOIN users c ON c.id = p.creator_id 
LEFT JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN tags t ON pt.tag_id = t.id
WHERE p.id = $1
GROUP BY 
  p.id,
  p.title,
  p.content,
  c.name,
  c.id,
  p.created_at,
  p.updated_at
  LIMIT 1;`,
      [id],
    );

    const splittedTags = data[0].tags.split(', ').map((elem) => {
      const [id, tag] = elem.split(':');

      return { id, tag };
    });

    data[0].tags = splittedTags;

    return data[0];
  }
}

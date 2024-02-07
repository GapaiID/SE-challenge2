import {
  Column,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { Post } from './post.entity';
import { Tag } from 'src/modules/tags/entities/tag.entity';

@Entity({ name: 'post_tags' })
export class PostTag {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'int', nullable: false })
  post_id: number;
  @JoinColumn({ name: 'post_id', referencedColumnName: 'id' })
  @ManyToOne(() => Post, (post) => post.tags)
  posts: Post;

  @Column({ type: 'int', nullable: false })
  tag_id: number;
  @JoinColumn({ name: 'tag_id', referencedColumnName: 'id' })
  @ManyToOne(() => Tag, (tag) => tag.posts)
  tags: Tag;

  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  created_at: Date;
  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  updated_at: Date;
}

import { PostTag } from 'src/modules/posts/entities/post-tags.entity';
import { Post } from 'src/modules/posts/entities/post.entity';
import {
  PrimaryGeneratedColumn,
  Column,
  Entity,
  JoinColumn,
  OneToMany,
} from 'typeorm';

@Entity({ name: 'tags' })
export class Tag {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'varchar', length: 255, nullable: false })
  tag: string;

  @JoinColumn({ name: 'posts', referencedColumnName: 'tag_id' })
  @OneToMany(() => PostTag, (postTag) => postTag.post_id, {
    onDelete: 'CASCADE',
    onUpdate: 'CASCADE',
  })
  posts: Post[];

  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  created_at: Date;
  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  updated_at: Date;
}

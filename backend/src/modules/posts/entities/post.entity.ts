import { Tag } from 'src/modules/tags/entities/tag.entity';
import { User } from 'src/modules/users/entities/user.entity';
import {
  PrimaryGeneratedColumn,
  Column,
  ManyToOne,
  JoinColumn,
  Entity,
  OneToMany,
} from 'typeorm';
import { PostTag } from './post-tags.entity';

@Entity({ name: 'posts' })
export class Post {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'varchar', length: 255, nullable: false })
  title: string;

  @Column({ type: 'text', nullable: false })
  content: string;

  @Column({ type: 'text', nullable: true })
  image_src: string;

  @Column({ type: 'int', nullable: false })
  creator_id: number;

  @JoinColumn({ name: 'creator_id', referencedColumnName: 'id' })
  @ManyToOne(() => User, (user) => user.posts, {
    onUpdate: 'CASCADE',
    onDelete: 'CASCADE',
    orphanedRowAction: 'delete',
  })
  creator: User;

  @JoinColumn({ name: 'id', referencedColumnName: 'post_id' })
  @OneToMany(() => PostTag, (postTag) => postTag.post_id, {
    onUpdate: 'CASCADE',
    onDelete: 'CASCADE',
    orphanedRowAction: 'delete',
  })
  tags: Tag[];

  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  created_at: Date;
  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  updated_at: Date;
}

import { Post } from 'src/modules/posts/entities/post.entity';
import {
  PrimaryGeneratedColumn,
  Column,
  OneToMany,
  Entity,
  JoinColumn,
} from 'typeorm';

@Entity({ name: 'users' })
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ type: 'varchar', length: 255, nullable: false })
  name: string;

  @Column({ type: 'varchar', length: 255, nullable: false, unique: true })
  email: string;

  @Column({ type: 'varchar', length: 255, nullable: false })
  password: string;

  @Column({ type: 'varchar', nullable: false, length: 255 })
  description: string;

  @Column({ type: 'text', nullable: true })
  profile_pic: string;

  @JoinColumn({ name: 'posts', referencedColumnName: 'creator_id' })
  @OneToMany(() => Post, (post) => post.creator, {
    onDelete: 'CASCADE',
    onUpdate: 'CASCADE',
  })
  posts: Post[];

  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  created_at: Date;
  @Column({ type: 'timestamp', nullable: false, default: 'NOW()' })
  updated_at: Date;
}

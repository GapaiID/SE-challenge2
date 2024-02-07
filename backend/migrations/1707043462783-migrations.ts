import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migrations1707043462783 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    const postTableQuery = `
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    creator_id INTEGER NOT NULL REFERENCES users(id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW() 
)`;

    await queryRunner.query(postTableQuery);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    const dropPostsQuery = `DROP TABLE posts IF EXISTS`;
    await queryRunner.query(dropPostsQuery);
  }
}

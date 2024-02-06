import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migrations1707043778561 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    const userTableQuery = `
CREATE TABLE IF NOT EXISTS post_tags (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    post_id INT NOT NULL REFERENCES posts(id),
    tag_id INT NOT NULL REFERENCES tags(id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW() 
)`;

    await queryRunner.query(userTableQuery);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    const dropUserQuery = `DROP TABLE post_tags IF EXISTS`;
    await queryRunner.query(dropUserQuery);
  }
}

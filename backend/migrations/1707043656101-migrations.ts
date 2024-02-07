import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migrations1707043656101 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    const tagsTableQuery = `
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    tag VARCHAR(255) UNIQUE NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW() 
)`;
    await queryRunner.query(tagsTableQuery);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    const dropUserQuery = `DROP TABLE tags IF EXISTS `;
    await queryRunner.query(dropUserQuery);
  }
}

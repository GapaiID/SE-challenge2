import { MigrationInterface, QueryRunner } from 'typeorm';

export class Migrations1707042990336 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    const userTableQuery = `
CREATE TABLE IF NOT EXISTS users  (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW() 
)`;

    await queryRunner.query(userTableQuery);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    const dropUserQuery = `DROP TABLE users IF EXISTS`;
    await queryRunner.query(dropUserQuery);
  }
}

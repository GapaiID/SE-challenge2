import { Module } from '@nestjs/common';
import { UsersModule } from './modules/users/users.module';
import { PostsModule } from './modules/posts/posts.module';
import { TagsModule } from './modules/tags/tags.module';
import { TypeOrmModule, TypeOrmModuleOptions } from '@nestjs/typeorm';
import { ConfigModule } from '@nestjs/config';
import { AuthModule } from './modules/auth/auth.module';

const configs = [
  ConfigModule.forRoot({
    isGlobal: true,
  }),
  TypeOrmModule.forRootAsync({
    imports: [ConfigModule],
    useFactory: async (): Promise<TypeOrmModuleOptions> => {
      const type = 'postgres';
      const host: string = process.env.DB_HOST ?? 'postgres';
      let port = +process.env.DB_PORT;
      port = isNaN(port) ? +port : 5432;
      const username: string = process.env.DB_USERN;
      const password: string = process.env.DB_PASS;
      const database: string = process.env.DB_NAME;
      const env = process.env.NODE_ENV;
      const synchronize: boolean = env && env === 'production' ? false : true;

      return {
        type,
        host,
        port,
        username,
        password,
        database,
        migrationsRun: true,
        migrations: [__dirname + '/../migrations/*.{js,ts}'],
        entities: [__dirname + '/modules/*/entities/*.entity.{js,ts}'],
        synchronize,
      };
    },
  }),
];

@Module({
  imports: [UsersModule, PostsModule, TagsModule, ...configs, AuthModule],
})
export class AppModule {}

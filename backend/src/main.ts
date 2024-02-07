import { HttpAdapterHost, NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import helmet from 'helmet';
import { GlobalExceptionFilter } from './exceptions/global.exception';
import { ValidationPipe } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.use(helmet());

  app.useGlobalPipes(new ValidationPipe());

  const httpAdapterHost = app.get(HttpAdapterHost);
  app.useGlobalFilters(new GlobalExceptionFilter(httpAdapterHost));

  app.enableCors({
    origin: process.env.ORIGIN,
  });

  await app.listen(3000);
}
bootstrap();

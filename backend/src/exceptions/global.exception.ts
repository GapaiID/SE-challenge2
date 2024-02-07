import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { HttpAdapterHost } from '@nestjs/core';
import { JsonWebTokenError } from '@nestjs/jwt';
import { TypeORMError } from 'typeorm';

@Catch()
export class GlobalExceptionFilter implements ExceptionFilter {
  constructor(private readonly httpAdapterHost: HttpAdapterHost) {}

  catch(exception: any, host: ArgumentsHost) {
    console.log(exception);
    const { httpAdapter } = this.httpAdapterHost;

    const ctx = host.switchToHttp();

    const httpStatus = this.getStatus(exception);
    const message = this.getMessage(exception);

    const responseBody = {
      statusCode: httpStatus,
      message: Array.isArray(message) ? message : [message],
      timestamp: new Date().toISOString(),
    };

    httpAdapter.reply(ctx.getResponse(), responseBody, httpStatus);
  }

  getMessage(err: any) {
    switch (true) {
      case err instanceof HttpException:
        const response: string | Record<string, any> = err.getResponse();
        return typeof response === 'object' ? response.message : err.message;

      case err instanceof JsonWebTokenError:
        return 'Unauthenticated';

      case err instanceof TypeORMError:
        return err?.message ?? 'Database Error';

      default:
        return 'Internal Server Error';
    }
  }

  getStatus(err: any): number {
    switch (true) {
      case err instanceof HttpException:
        return err.getStatus();

      case err instanceof JsonWebTokenError:
        return 401;

      case err instanceof TypeORMError:
        return HttpStatus.SERVICE_UNAVAILABLE;

      default:
        return 500;
    }
  }
}

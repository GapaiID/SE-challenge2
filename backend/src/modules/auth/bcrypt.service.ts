import { Injectable } from '@nestjs/common';
import * as bcrypt from 'bcrypt';

@Injectable()
export class BcryptService {
  private secret = +process.env.BCRYPT_SECRET;
  private salt: string;

  async hashing(password: string) {
    return bcrypt.hash(password, this.salt);
  }

  async generateSalt() {
    this.salt = await bcrypt.genSalt(this.secret);
  }

  async verify(password: string, dbPass: string) {
    return bcrypt.compare(password, dbPass);
  }
}

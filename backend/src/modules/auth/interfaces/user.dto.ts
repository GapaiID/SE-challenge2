import { IsNotEmpty, IsString, IsUrl, ValidateIf } from 'class-validator';

export class UserRegisterDto {
  @IsNotEmpty()
  @IsString()
  public name: string;

  @IsNotEmpty()
  @IsString()
  public email: string;

  @IsNotEmpty()
  @IsString()
  public password: string;

  @IsNotEmpty()
  @IsString()
  public description: string;

  @ValidateIf((data) => !!data.profile_pic)
  @IsString()
  @IsUrl()
  public profile_pic?: string;
}

export class UserLoginDto {
  @IsNotEmpty()
  @IsString()
  public readonly email: string;

  @IsNotEmpty()
  @IsString()
  public readonly password: string;
}

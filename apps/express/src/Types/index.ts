import { Request } from 'express'

export interface UserReqBody {
  email: string
  password: string
  phoneNumber: string
  fullName: string
}

export interface LoginRequestBody {
  email: string
  password: string
}

export interface UserPayload {
  id?: string
  role: string
}

export interface UserRequest extends Request {
  decoded?: UserPayload
}

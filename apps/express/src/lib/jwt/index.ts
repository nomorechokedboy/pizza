import jwt from 'jsonwebtoken'
import { refreshTokenKey, SECRET_KEY } from '../../config/env'
import { UserPayload } from '../../Types'

export default function genToken(payload: UserPayload) {
  if (!SECRET_KEY) throw Error(`Secret key is ${SECRET_KEY}`)

  return jwt.sign(payload, SECRET_KEY, { expiresIn: '1m' })
}

export const genRefreshToken = (payload: UserPayload) => {
  if (!refreshTokenKey) throw Error(`Refresh token key is ${refreshTokenKey}`)

  return jwt.sign(payload, refreshTokenKey, { expiresIn: '2h' })
}

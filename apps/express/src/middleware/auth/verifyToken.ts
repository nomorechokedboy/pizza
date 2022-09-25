import { Response, NextFunction } from 'express'
import jwt from 'jsonwebtoken'
import { refreshTokenKey } from '../../config/env'
import { UserPayload, UserRequest } from '../../Types'

export default function verifyToken(
  req: UserRequest,
  res: Response,
  next: NextFunction
) {
  const refreshToken = req.headers.authorization?.split(' ')[1]
  if (!refreshToken) throw Error('Refresh token is undefined!')
  if (!refreshTokenKey) throw Error('Refresh token key is undefined!')

  try {
    const decoded = jwt.verify(refreshToken, refreshTokenKey) as UserPayload
    if (decoded) {
      req.decoded = decoded
      return next()
    }

    throw Error('Invalid token!')
  } catch (e) {
    next(e)
    return res.status(401).json({ error: 'Invalid token!' })
  }
}

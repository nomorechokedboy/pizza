import { NextFunction, Response } from 'express'
import jwt from 'jsonwebtoken'
import { SECRET_KEY } from '../../config/env'
import { authFailed } from '../../constants'
import { UserPayload, UserRequest } from '../../Types'

export default function isAdmin(
  req: UserRequest,
  res: Response,
  next: NextFunction
) {
  const token = req.headers.authorization?.split(' ')[1]
  if (!token)
    return res.status(403).json({ error: 'Invalid request (without token)' })
  if (!SECRET_KEY) throw Error(`Secret key is ${SECRET_KEY}`)

  try {
    const decoded = jwt.verify(token, SECRET_KEY) as UserPayload
    req.decoded = decoded

    if (decoded.role === 'admin') next()
    throw Error(authFailed)
  } catch (e) {
    console.error(e)
    return res.status(401).json({ error: authFailed })
  }
}

import bcrypt from 'bcrypt'
import { NextFunction, Request, Response } from 'express'
import { serverError, wrongUser } from '../constants'
import genToken, { genRefreshToken } from '../lib/jwt'
import { LoginRequestBody, UserReqBody, UserRequest } from '../Types'
import User from './model'

export const register = async (
  req: Request<unknown, unknown, UserReqBody, unknown>,
  res: Response,
  next: NextFunction
) => {
  const { email, password, phoneNumber, fullName } = req.body
  const data = { email, password, fullName, phoneNumber }

  try {
    const find = await User.findOne({ email })
    if (find)
      return res.status(401).json({
        errors: [
          {
            email: [
              `The email ${email} has been used, please used another email!`
            ]
          }
        ]
      })

    const user = new User({
      ...data,
      role: 'user'
    })
    const saveUser = await user.save()

    return res
      .status(201)
      .json({ message: `Register success for email: ${saveUser.email}` })
  } catch (e) {
    next(e)
    res.status(500).json({ error: serverError + new Date().toString() })
  }
}

export const login = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const { email, password } = req.body as LoginRequestBody

  try {
    const user = await User.findOne({ email })
    if (!user) return res.status(404).json({ error: wrongUser })

    const match = await bcrypt.compare(password, user.password)

    if (!match) return res.status(404).json({ error: wrongUser })

    const { _id: id, role } = user
    const payload = { id, role }

    return res.json({
      token: genToken(payload),
      refreshToken: genRefreshToken(payload)
    })
  } catch (e) {
    next(e)
    res.status(500).json({ error: serverError + new Date().toString() })
  }
}

export const getUser = async (
  req: UserRequest,
  res: Response,
  next: NextFunction
) => {
  const { id } = req.params
  const confirmId = req.decoded?.id

  try {
    if (id !== confirmId)
      return res.status(404).json({ error: 'You have no authorized!' })

    const userInfo = await User.findById(id).select('fullName')
    if (!userInfo) return res.status(404).json({ error: 'Invalid user!' })

    return res.json({ info: userInfo })
  } catch (e) {
    next(e)
    return res.status(500).json({ error: serverError + new Date().toString() })
  }
}

export const refreshToken = (
  req: UserRequest,
  res: Response,
  next: NextFunction
) => {
  try {
    if (req.decoded) {
      const { id, role } = req.decoded
      const payload = { id, role }

      res.json({
        token: genToken(payload),
        refreshToken: genRefreshToken(payload)
      })
    }
  } catch (e) {
    next(e)

    return res.status(500).json({ error: serverError + new Date().toString() })
  }
}

// export const logout = (req: Request,)

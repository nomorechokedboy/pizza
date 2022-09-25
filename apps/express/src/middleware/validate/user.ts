import { NextFunction, Request, Response } from 'express'
import Validator from 'validatorjs'
import { loginRules as rules } from '../../constants'

export const register = (req: Request, res: Response, next: NextFunction) => {
  const registerRules = {
    ...rules,
    fullName: 'required|min:8|max:30',
    phoneNumber: 'required|min:10|max:12',
    cPassword: 'same:password'
  }

  const validation = new Validator(req.body, registerRules)
  if (validation.fails()) return res.status(400).json(validation.errors.all())

  next()
}

export const login = (req: Request, res: Response, next: NextFunction) => {
  const validation = new Validator(req.body, rules)
  if (validation.fails()) return res.status(400).json(validation.errors.all())

  next()
}

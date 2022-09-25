import express from 'express'
import * as controller from './controller'
import * as validate from '../middleware/validate/user'
import isUser from '../middleware/auth/isUser'
import verifyToken from '../middleware/auth/verifyToken'
const userRouter = express.Router()

userRouter.post('/register', validate.register, controller.register)
userRouter.post('/login', validate.login, controller.login)
userRouter.get('/refreshToken', verifyToken, controller.refreshToken)
userRouter.get('/:id/detail', isUser, controller.getUser)

export default userRouter

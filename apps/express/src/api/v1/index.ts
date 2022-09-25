import productRouter from '../../product'
import userRouter from '../../user'
import { Router } from 'express'

const routerV1 = Router()

routerV1.use('/product', productRouter)
routerV1.use('/user', userRouter)

export default routerV1

import express from 'express'
import * as controller from './controller'
const productRouter = express.Router()

productRouter.post('/cookie', controller.postTest)
productRouter.get('/paging', controller.paging)
productRouter.get('/:slug/detail', controller.getDetail)
// productRouter.get('/faker', controller.genData);
productRouter.get('/', controller.getAll)

export default productRouter

import { NextFunction, Request, Response } from 'express'
import { RedisClient } from '../config'
import { serverError } from '../constants/messages'
import { productModel } from './model'
// import { Product } from '@pizza/core';
// import { faker } from '@faker-js/faker';

// const ProductType = ['pizza', 'starter', 'drink', 'combo'];

// function randomNumber(min: number, max: number) {
//   return Math.floor(Math.random() * (max - min)) + min;
// }

export const paging = async (
  req: Request<
    unknown,
    unknown,
    unknown,
    { page: string; pageSize: string; search: string }
  >,
  res: Response
) => {
  const page = parseInt(req.query.page) || 0
  const pageSize = parseInt(req.query.pageSize) || 20
  const { search } = req.query
  const cacheKey = `products:paging:page=${page}:pageSize=${pageSize}:search=${search}`
  const client: RedisClient = req.app.get('redisClient')

  try {
    const hit = await client.get(cacheKey)
    if (hit) return res.json({ data: JSON.parse(hit) })

    const products = await productModel
      .find(search ? { $text: { $search: search } } : {})
      .skip(page * pageSize)
      .limit(pageSize)
      .lean()
      .exec()

    await client.setEx(cacheKey, 3600, JSON.stringify(products))

    res.json({ data: products })
  } catch (e) {
    console.error(e)
    res.status(500).json({ error: serverError + new Date().toString() })
  }
}

export const getDetail = async (req: Request, res: Response) => {
  const { slug } = req.params

  try {
    const hitRepo = await productModel.findOne({ slug }).lean().exec()
    if (!hitRepo)
      return res
        .status(404)
        .json({ error: `Cannot found product with slug ${slug}` })

    res.json({ data: hitRepo })
  } catch (e) {
    console.error(e)

    res.status(500).json({ error: serverError + new Date().toString() })
  }
}

// export const genData = async (
//   _: Request,
//   res: Response,
//   next: NextFunction,
// ) => {
//   const products: Product[] = Array.from(Array(100)).map((_1) => {
//     const type = ProductType[randomNumber(0, 3)];
//     if (!type) throw Error(`type is ${type}`);

//     return {
//       description: faker.commerce.productDescription(),
//       img: faker.image.food(1600, 1100, true),
//       name: faker.commerce.productName(),
//       price: parseInt(faker.commerce.price(10000, 500000, 0)),
//       type,
//     };
//   });

//   try {
//     const created = await productModel.create(products);

//     res.json({ data: created });
//   } catch (e) {
//     console.error(e);

//     next(e);
//   }
// };

export const getAll = async (req: Request, res: Response) => {
  const cacheKey = 'products:getAll'
  const client: RedisClient = req.app.get('redisClient')

  try {
    const hit = await client.get(cacheKey)
    res.cookie('refreshToken', 'chanasdasdasdged', { maxAge: 60 * 1000 * 5 })

    if (hit) return res.json({ data: JSON.parse(hit) })

    const products = await productModel.find({}).lean().exec()
    await client.setEx(cacheKey, 60, JSON.stringify(products))

    res.json({ data: products })
  } catch (e) {
    console.log(e)
    res.status(500).json({ error: 'Server Internal Error' })
  }
}

export const postTest = (req: Request, res: Response, _: NextFunction) => {
  const { cookies } = req
  console.log({ cookies })
  return res.json({ cookies })
}

import mongoose from 'mongoose'
import slug from 'mongoose-slug-generator'

export interface Product {
  _id?: string
  name: string
  description: string
  img: string
  price: number
  type: string
  slug?: string
  createdAt?: Date
  updatedAt?: Date
  __v?: number
}

mongoose.plugin(slug)

const schema = new mongoose.Schema<Product>(
  {
    name: {
      type: String,
      require: true,
      unique: true
    },
    description: {
      type: String,
      required: true
    },
    img: {
      type: String,
      required: true,
      unique: true
    },
    price: {
      type: Number,
      required: true
    },
    type: {
      type: String,
      required: true,
      unique: true
    },
    slug: {
      type: String,
      slug: 'name'
    }
  },
  { timestamps: true }
)

schema.index({ name: 'text', description: 'text' })

export const productModel = mongoose.model<Product>('product', schema)

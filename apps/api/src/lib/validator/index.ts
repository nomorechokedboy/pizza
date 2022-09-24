import Validator from 'validatorjs'
import mongoose from 'mongoose'

interface ObjectIdLike {
  id: string | Buffer
  __id?: string
  toHexString(): string
}

type IsValidObjectIdInput =
  | string
  | number
  | mongoose.Types.ObjectId
  | ObjectIdLike
  | Buffer
  | Uint8Array

Validator.register(
  'ObjectId',
  (value) => mongoose.Types.ObjectId.isValid(value as IsValidObjectIdInput),
  'The :attribute is not ObjectId'
)

export default Validator

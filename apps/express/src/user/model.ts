import mongoose from 'mongoose'
import bcrypt from 'bcrypt'

interface User extends mongoose.Document {
  fullName: string
  email: string
  phoneNumber: string
  password: string
  role: string
}

const schema = new mongoose.Schema(
  {
    fullName: {
      type: String,
      required: true
    },
    email: {
      type: String,
      required: true,
      unique: true
    },
    phoneNumber: {
      type: String,
      required: true
    },
    password: {
      type: String,
      required: true
    },
    role: {
      type: String,
      required: true
    }
  },
  { timestamps: true }
)

schema.pre<User>('save', function (next) {
  if (this.isNew) {
    const salt = bcrypt.genSaltSync(10)
    this.password = bcrypt.hashSync(this.password, salt)
  }

  next()
})

export default mongoose.model<User>('user', schema)

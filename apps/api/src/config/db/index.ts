import mongoose from 'mongoose'
import { MONGODB } from '../env'

export function connectDb() {
  if (!MONGODB) throw Error(`MONGODB is ${MONGODB}`)

  mongoose
    .connect(MONGODB, {
      useNewUrlParser: true,
      useUnifiedTopology: true
    } as mongoose.ConnectOptions)
    .then(async () => {
      console.log('Connect Mongodb Success!')
    })
    .catch((e) => console.error(e))
}

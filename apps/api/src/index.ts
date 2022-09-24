import cookieParser from 'cookie-parser'
import cors, { CorsOptions } from 'cors'
import express, { Handler } from 'express'
import { createServer, Server } from 'http'
import morgan from 'morgan'
import routerV1 from './api'
import { connectDb, MORGAN, PORT } from './config'

const CORS_WHITELIST = [
  'http://localhost:5001/',
  'http://localhost:5000/',
  'https://pizza-api-nomorechokedboy.cloud.okteto.net'
]

const app = express()

connectDb()

app.use(cors(CORS_WHITELIST as CorsOptions))
app.use(express.json({}))
app.use(cookieParser())
app.use(
  express.urlencoded({
    extended: true
  })
)

morgan.format(
  'myformat',
  '[:date[clf]] ":method :url" :status :res[content-length] - :response-time ms'
)

if (MORGAN === '1') {
  app.use('/api/*', morgan('myformat') as Handler)
}

app.use('/api/v1', routerV1)
app.use('/healthcheck', (_, res) => {
  const healthcheck = {
    uptime: process.uptime(),
    message: 'I am fine',
    timestamp: Date.now()
  }

  try {
    res.json(healthcheck)
  } catch (e) {
    // if (e instanceof Error) next(new HttpException(e.message, 503));
    console.log(e)
  }
})

if (import.meta.env.PROD) {
  const server: Server = createServer(app)
  server.on('error', (e) => {
    if (e) throw e
  })

  server.listen(PORT || 5000, () => {
    console.log(`Stikinote api on http://localhost:${PORT}`)
  })
}

export const viteNodeApp = app

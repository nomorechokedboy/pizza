import { createClient } from 'redis'
import { REDIS_HOST, REDIS_PORT } from '../env'

export const client = createClient({
  socket: {
    host: REDIS_HOST || 'localhost',
    port: parseInt(REDIS_PORT) || 6379
  }
})

client.on('error', (e) => console.log(`Redis client error: ${e}`))

export type RedisClient = typeof client

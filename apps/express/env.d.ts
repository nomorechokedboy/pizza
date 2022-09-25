declare global {
  namespace NodeJS {
    interface ProcessEnv {
      PIZZA_SERVER_PORT: number
      MONGODB: string
      MORGAN: string
      SECRET_KEY: string
      refreshTokenKey: string
      UID_GID: string
      REDIS_HOST: string
      REDIS_PORT: number
    }
  }
}

export {}

/// <reference types="vite/client" />
interface ImportMetaEnv {
  readonly VITE_APP_TITLE: string
  PIZZA_SERVER_PORT: string
  MONGODB: string
  MORGAN: string
  SECRET_KEY: string
  refreshTokenKey: string
  UID_GID: string
  REDIS_HOST: string
  REDIS_PORT: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}

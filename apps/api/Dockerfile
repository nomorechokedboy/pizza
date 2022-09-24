ARG env=prod

FROM node:14-alpine as deps
WORKDIR /usr/src/app
RUN yarn global add pnpm
COPY package.json .npmrc /usr/src/app/
RUN pnpm i

FROM node:14-alpine AS builder
WORKDIR /usr/src/app

RUN yarn global add pnpm

COPY --from=deps /usr/src/app/node_modules ./node_modules
COPY --from=deps /usr/src/app/pnpm-lock.yaml .
COPY --from=deps /usr/src/app/package.json .
COPY . .
RUN if [ "$env" = "prod" ] ; then pnpm build && pnpm i --prod --ignore-scripts --prefer-offline --frozen-lockfile ; else echo 'adu' ; fi

FROM node:14-alpine

WORKDIR /usr/src/app

COPY --chown=node:node --from=builder /usr/src/app/build ./build
COPY --chown=node:node --from=builder /usr/src/app/node_modules ./node_modules
COPY --chown=node:node --from=builder /usr/src/app/package.json ./package.json

CMD ["yarn", "start"]

# docker build -t nomorechokedboy/web-project:pizza-api-prod .


# FROM node:14-alpine as deps
# WORKDIR /usr/src/app
# RUN yarn global add pnpm
# COPY package.json .npmrc /usr/src/app/
# RUN pnpm i

# FROM node:14-alpine as dev
# COPY --from=deps /usr/src/app/node_modules ./node_modules
# COPY --from=deps /usr/src/app/pnpm-lock.yaml .
# COPY --from=deps /usr/src/app/package.json .
# COPY . .

# FROM node:14-alpine AS builder
# WORKDIR /usr/src/app

# RUN yarn global add pnpm

# COPY --from=deps /usr/src/app/node_modules ./node_modules
# COPY --from=deps /usr/src/app/pnpm-lock.yaml .
# COPY --from=deps /usr/src/app/package.json .
# COPY . .
# RUN pnpm build && pnpm i --prod --ignore-scripts --prefer-offline --frozen-lockfile

# FROM node:14-alpine

# WORKDIR /usr/src/app

# COPY --chown=node:node --from=builder /usr/src/app/build ./build
# COPY --chown=node:node --from=builder /usr/src/app/node_modules ./node_modules
# COPY --chown=node:node --from=builder /usr/src/app/package.json ./package.json

# CMD ["yarn", "start"]

# # docker build -t nomorechokedboy/web-project:pizza-api-prod .


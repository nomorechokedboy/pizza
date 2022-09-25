FROM node:14-alpine as starter
WORKDIR /usr/src/app

RUN yarn global add pnpm turbo
COPY . .
RUN turbo prune --scope=express --docker

FROM node:14-alpine AS installer
ARG env
WORKDIR /usr/src/app
RUN yarn global add pnpm

COPY pnpm-workspace.yaml .
COPY --from=starter /usr/src/app/out/json/ .
COPY --from=starter /usr/src/app/out/pnpm-lock.yaml .
RUN pnpm i

COPY --from=starter /usr/src/app/out/full/ .
COPY turbo.json .
RUN if [ "$env" = "prod" ] ; then pnpm build --filter=express... && pnpm i --prod --frozen-lockfile --offline ; fi
# RUN  pnpm build --filter=api... && pnpm i --prod --frozen-lockfile --offline --filter=api

FROM node:14-alpine as runner

WORKDIR /usr/src/app

USER node
COPY --from=installer /usr/src/app .

CMD node apps/api/build/index.js

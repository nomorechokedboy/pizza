FROM rust:1.68 AS builder

RUN mkdir repo repo/apps && cargo new app/apps/noisy-boi
WORKDIR /app

COPY ./Cargo.lock ./Cargo.toml ./
# COPY ./apps/noisy-boi/migrations ./apps/noisy-boi/migrations
COPY ./apps/noisy-boi/Cargo.toml ./apps/noisy-boi/

RUN cargo b -r

COPY ./apps/noisy-boi/src ./apps/noisy-boi/src
RUN touch /app/apps/noisy-boi/src/main.rs \ 
    && cargo b -r

FROM busybox:1.36 AS deps

ARG BUSYBOX_VERSION=1.31.0-i686-uclibc
ADD https://busybox.net/downloads/binaries/$BUSYBOX_VERSION/busybox_WGET /wget
RUN chmod a+x /wget

FROM gcr.io/distroless/cc:nonroot AS runtime

COPY --from=deps /wget /usr/bin/wget
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /

COPY --from=builder /app/target/release/noisy-boi .
ENTRYPOINT ["./noisy-boi"]


FROM rust:1.68 AS builder

RUN cargo new app
WORKDIR /app

COPY Cargo.lock Cargo.toml ./
COPY migrations ./migrations

RUN cargo b -r

COPY src ./src
RUN touch /app/src/main.rs \ 
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

COPY --from=builder /app/target/release/tags-keeper .
ENTRYPOINT ["./tags-keeper"]


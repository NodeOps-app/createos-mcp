# syntax=docker/dockerfile:1
FROM debian:stable-slim
COPY --link --from=gcr.io/distroless/base /etc/ssl/certs/ /etc/ssl/certs/
RUN --mount=type=cache,target=/var/cache/apt \
  --mount=type=cache,target=/var/lib/apt/lists \
  apt-get update && apt-get install -y --no-install-recommends openssl curl
WORKDIR /app
COPY --link ./nodeops ./
ENTRYPOINT [ "/bin/sh", "-c" ]

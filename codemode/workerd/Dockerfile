FROM oven/bun:1 AS deps
WORKDIR /app
COPY package.json bun.lock* ./
RUN bun install --frozen-lockfile

FROM oven/bun:1 AS build
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN bun run build

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
    libssl3 ca-certificates \
 && rm -rf /var/lib/apt/lists/* \
 && groupadd --system codemode --gid 65532 \
 && useradd  --system --uid 65532 --gid 65532 --no-create-home --shell /sbin/nologin codemode
WORKDIR /app
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/dist ./dist
COPY config.capnp dispatcher.js host.js api-proxy.js api-sdk-builder.js ./
USER codemode
EXPOSE 8787
ENTRYPOINT ["./node_modules/.bin/workerd", "serve", "--experimental", "config.capnp"]

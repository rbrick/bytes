FROM node:18-alpine as node_builder

WORKDIR /bytes
COPY . .

RUN npm ci && npm run build

FROM node:18-alpine as node_runner
WORKDIR /bytes

# copy NPM files
COPY --from=node_builder /bytes/package.json .
COPY --from=node_builder /bytes/package-lock.json .

# copy next js files
COPY --from=node_builder /bytes/next.config.js .
COPY --from=node_builder /bytes/public ./public
COPY --from=node_builder /bytes/.next ./.next
COPY --from=node_builder /bytes/.next/standalone .


EXPOSE 3000
ENTRYPOINT [ "node", "server.js"]

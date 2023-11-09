FROM node:18-alpine

WORKDIR /bytes
COPY . .

RUN npm ci && npm run build

EXPOSE 3000
ENTRYPOINT [ "npx", "next", "start" ]

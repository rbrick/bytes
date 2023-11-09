FROM node:18

WORKDIR /bytes
COPY . .

RUN npm ci && npm run build

ENTRYPOINT [ "npx", "next", "start" ]

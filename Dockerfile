FROM golang:1.21-alpine as go_builder

WORKDIR /bytes
COPY . .
RUN go build 

FROM alpine

COPY --from=go_builder /bytes/bytes .

EXPOSE 1337

RUN stat bytes

CMD ["./bytes", "http" ]


# FROM node:18 as node_builder

# COPY --from=go_builder . .
# WORKDIR "/frontend/"

# RUN ls -a && npm ci && npm run build
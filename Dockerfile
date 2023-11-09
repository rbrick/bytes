FROM golang:1.21-alpine as go_builder

WORKDIR /bytes
COPY . .
RUN go build 

FROM alpine

COPY --from=go_builder /bytes/bytes .

EXPOSE 1337

CMD ["./bytes", "http" ]
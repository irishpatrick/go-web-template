FROM golang:1.23-bookworm AS builder

ARG app=go-web-template

WORKDIR /src
COPY . .
RUN go build . ; strip $app

FROM alpine:latest AS runner

RUN apk add libc6-compat
COPY --from=builder /src/$app /usr/local/bin/app

ARG PORT=80
ENV HOST="0.0.0.0"
ENV PORT=$PORT
EXPOSE $PORT
CMD ["app"]
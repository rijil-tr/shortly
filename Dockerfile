FROM golang:1.11.9-alpine3.9 as BUILD
RUN apk add --no-cache git gcc musl-dev
WORKDIR /app

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
COPY ./ ./

WORKDIR /app/cmd
RUN go build -o shortly

FROM alpine:3.7
WORKDIR /app
COPY --from=BUILD /app/cmd/shortly /app
COPY --from=BUILD /app/public /app/public
EXPOSE 8080
CMD ["/app/shortly"]

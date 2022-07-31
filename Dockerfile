#BUILD
FROM golang:1.18-alpine AS build_base

#RUN apk add --no-cache git

WORKDIR /tmp/bot

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/bot .

#DEPLOY
FROM alpine:3.16 
RUN apk add ca-certificates

COPY --from=build_base /tmp/bot/out/bot /app/bot
COPY --from=build_base /tmp/bot/img/* /app/img/

EXPOSE 3000

WORKDIR /app
ENTRYPOINT ["/app/bot"]

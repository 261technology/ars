FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /app/guin
COPY . .
COPY ./frontend /app/guin/frontend
COPY .env /app/guin/.env
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
RUN GOOS=linux go build -ldflags="-s -w" -o ./guin-app ./main.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /go/bin
COPY --from=build /app/guin /go/bin
EXPOSE 4000
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
ENTRYPOINT /go/bin/guin-app --port 4000
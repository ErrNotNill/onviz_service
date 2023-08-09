FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /root/project/bitrix
COPY . .
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go
FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /root/project/bitrix
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 9090
ENTRYPOINT ./onviz --port 9090

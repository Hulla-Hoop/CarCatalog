FROM golang:latest AS build
WORKDIR /app
COPY . /app
RUN GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app main.go

FROM scratch AS runtime
COPY --from=build /app/app .
EXPOSE 8080/tcp
ENTRYPOINT ["./app"]

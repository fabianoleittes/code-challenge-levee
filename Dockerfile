FROM golang:alpine as builder

ENV GO111MODULE=on \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	MONGODB_HOST=mongodb \
	MONGODB_DATABASE=job \
	APP_NAME=code-challenge-levee-api \
	APP_PORT=3001

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -a --installsuffix cgo --ldflags="-s" -o main

FROM scratch

COPY --from=builder /build .

ENTRYPOINT ["./main"]

EXPOSE 3001

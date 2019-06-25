FROM golang:1.12.6 as builder
LABEL maintainer="AlistairFink <alistairfink@gmail.com>"

WORKDIR /go/src/github.com/alistairfink/AlistairFink-Website-BackEnd
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/AlistairFink-Website-BackEnd .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/AlistairFink-Website-BackEnd .
COPY --from=builder /go/src/github.com/alistairfink/AlistairFink-Website-BackEnd/Config.json .

EXPOSE 41691

CMD ["./AlistairFink-Website-BackEnd"] 
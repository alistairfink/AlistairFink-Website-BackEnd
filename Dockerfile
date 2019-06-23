FROM golang:1.12.6

LABEL maintainer="AlistairFink <alistairfink@gmail.com>"

WORKDIR $GOPATH/src/github.com/alistairfink/AlistairFink-Website-BackEnd

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 41691

CMD ["AlistairFink-Website-BackEnd"]
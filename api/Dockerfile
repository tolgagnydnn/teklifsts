FROM golang:1.13.6

RUN go get github.com/beego/bee

WORKDIR /go/src/api

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8090

CMD ["bee", "run", "-downdoc=true", "-gendoc=true"]
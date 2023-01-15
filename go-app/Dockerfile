FROM golang:1.17-alpine
ENV ROOT=/go/src/app
RUN mkdir ${ROOT}
WORKDIR ${ROOT}

COPY go.mod go.sum ./
COPY . .
RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary
EXPOSE 8080
CMD ["/go/src/app/binary"]
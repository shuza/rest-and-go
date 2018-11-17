FROM golang
COPY . /go/src/rest-and-go
WORKDIR /go/src/rest-and-go
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 8000
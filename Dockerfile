FROM golang
COPY . /rest-and-go
WORKDIR /rest-and-go
RUN go mod vendor
ENTRYPOINT go run main.go
EXPOSE 8000

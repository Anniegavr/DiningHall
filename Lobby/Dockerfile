FROM golang:1.16-alpine
RUN mkdir /new
WORKDIR /new
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /lobby
EXPOSE 8081
CMD ["/lobby"]

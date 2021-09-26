FROM docker.io/library/golang:latest
RUN mkdir /build
WORKDIR /build
RUN export GO111MODULE=on
RUN go get github.com/Anniegavr/Lobby/Lobby/main
RUN cd /build && git clone https://github.com/Anniegavr/Lobby
RUN cd /build/Lobby/Lobby/main && go build
EXPOSE 8081
ENTRYPOINT "/build/Lobby/Lobby/main/main"

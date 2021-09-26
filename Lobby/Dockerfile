FROM docker.io/library/golang:latest
RUN mkdir /new
WORKDIR /new
RUN export GO111MODULE=on
RUN go get github.com/Anniegavr/Lobby/Lobby/main
RUN cd /new && git clone https://github.com/Anniegavr/Lobby
RUN cd /new/Lobby/Lobby/main && go build
EXPOSE 8081
ENTRYPOINT "/new/Lobby/Lobby/main/main"

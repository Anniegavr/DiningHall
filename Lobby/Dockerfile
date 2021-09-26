FROM docker.io/library/golang:latest
RUN mkdir /build
WORKDIR /build
RUN export GO111MODULE=on
RUN go get github.com/Anniegavr/Lobby/main
RUN cd /build && git clone https://github.com/Anniegavr/Lobby
RUN cd /build/RESTaurant/main && go build
EXPOSE 8080
ENTRYPOINT "/build/RESTaurant/main/main"

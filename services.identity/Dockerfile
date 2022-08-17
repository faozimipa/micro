FROM golang:alpine
WORKDIR /build
COPY services.identity services.identity
COPY shared shared
WORKDIR /build/services.identity

RUN go mod download
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.identity/main .
RUN cp /build/services.identity/src/app.env .
CMD ["/dist/main"]
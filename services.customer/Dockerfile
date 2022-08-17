FROM golang:alpine
WORKDIR /build
COPY services.customer services.customer
COPY shared shared
WORKDIR /build/services.customer

# RUN go mod download
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.customer/main .
RUN cp /build/services.customer/src/app.env .
CMD ["/dist/main"]

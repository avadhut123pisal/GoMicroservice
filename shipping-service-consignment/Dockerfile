FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app

WORKDIR app

ENV GO111MODULE=on

COPY . .

RUN go mod download

#build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shipping-service-consignment

#RUN CONTAINER
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app

WORKDIR app

COPY --from=builder /app/shipping-service-consignment .

CMD ["./shipping-service-consignment"]
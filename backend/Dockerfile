FROM golang:1.20-alpine as builder

WORKDIR /tmp/app

RUN apk update

RUN apk add --update gcc musl-dev

COPY . .

ENV CGO_ENABLED=1 

ENV CGO_CFLAGS=-D_LARGEFILE64_SOURCE

RUN go get -d ./...

RUN CGO_ENABLED=1 

RUN GOOS=linux go install ./...

# RUN GOOS=linux go build -tags=linux,darwin,window

RUN GOOS=linux go build


# Unit tests
# RUN CGO_ENABLED=0 go test -v


FROM alpine:3.9 

COPY --from=builder /tmp/app ./

EXPOSE 1323

CMD [ "./process-logs" ]
FROM golang:1.19-alpine

WORKDIR /go/src
COPY ./src .
RUN apk upgrade --update && \
    apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest  
# RUN go mod download
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]

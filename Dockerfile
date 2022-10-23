FROM golang:alpine AS build

RUN apk add --update git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/hoffme/ddd-backend

COPY . .

RUN go build -o /go/bin/app cmd/backend/main.go

FROM scratch

COPY --from=build /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
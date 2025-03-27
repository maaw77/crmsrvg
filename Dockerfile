FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

COPY . $GOPATH/src/crmsrvg/

WORKDIR $GOPATH/src/crmsrvg
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN  $GOPATH/bin/migrate -verbose -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST_DB}:${PORT_DB}/${POSTGRES_DB}?sslmode=disable' -path ./migrations up 
RUN go mod tidy
# RUN go mod download
RUN mkdir /app
RUN go build -o /app/crm .

FROM alpine:latest

RUN mkdir /app
COPY --from=builder /app/crm /app/crm
COPY --from=builder go//src/crmsrvg/docs /app/
COPY --from=builder go//src/crmsrvg/config/config.yaml /app/config/

WORKDIR /app 
# CMD ["/app/zbot", "run"]
CMD ["/app/crm"]
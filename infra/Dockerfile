FROM golang:1.19 as build
ENV CGO_ENABLED 0
ENV SERVICE user-mng-service

ARG WORKDIR=/app
WORKDIR $WORKDIR

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o "$SERVICE" ./cmd/service && \
    wget https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz &&  \
    tar -xvf migrate.linux-amd64.tar.gz

#FROM scratch
FROM alpine
COPY --from=build app/user-mng-service .
COPY --from=build app/migrate .
COPY --from=build app/migrations/sql ./migrations
COPY --from=build app/scripts/run.sh .

RUN chmod +x ./run.sh

CMD ./run.sh

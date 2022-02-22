FROM golang:1.17.7-alpine3.15 AS build
COPY . /app/
WORKDIR /app
RUN go build -o ./bin/task-mgmt ./cmd/task-mgmt

FROM alpine:3.15
WORKDIR /
COPY --from=build /app/bin /
CMD ["/task-mgmt"]

# syntax=docker/dockerfile:1

FROM  golang:1.21-alpine as builder
ENV CGO_ENABLED=0
WORKDIR /
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN go build -o /gorinha .

FROM gcr.io/distroless/base-debian11
LABEL maintainer="Erick Amorim <ericklima.ca@yahoo.com>"
COPY --from=builder /gorinha /gorinha
COPY --from=builder /files /files
ENTRYPOINT ["/gorinha"]

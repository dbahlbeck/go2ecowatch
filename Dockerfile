FROM golang:1.25.0-alpine as build-stage
WORKDIR /app
RUN pwd
COPY go.mod go.sum ./
COPY internal/ ./internal
COPY cmd/ ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/go2ecowatch.go

FROM build-stage
WORKDIR /
COPY --from=build-stage /app/go2ecowatch /go2ecowatch

ENTRYPOINT ["/go2ecowatch"]
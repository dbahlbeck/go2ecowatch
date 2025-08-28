FROM golang:1.25.0-alpine as build-stage
WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go2ecowatch

FROM build-stage
WORKDIR /
COPY --from=build-stage /go2ecowatch /go2ecowatch

ENTRYPOINT ["/go2ecowatch"]
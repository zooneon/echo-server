FROM --platform=$BUILDPLATFORM golang:1.20-alpine3.17 AS builder
ADD . /app
WORKDIR /app
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o echo-server ./cmd/echo-server/main.go

FROM scratch
COPY --from=builder /app/echo-server /
ENV PORT=8080
ENTRYPOINT [ "/echo-server" ]
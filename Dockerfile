# Сборка: Go кросс-компилирует под целевую платформу на платформе сборщика,
# поэтому для arm64 эмуляция (QEMU) не нужна.
FROM --platform=$BUILDPLATFORM golang:1.27 AS build
ARG TARGETOS TARGETARCH
# Реальная версия передаётся снаружи: --build-arg VERSION=v0.1.1 (CI, make docker).
# "dev" — значение по умолчанию для голого `docker build .`, как и в cmd/cmd.go.
ARG VERSION=dev
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath \
    -ldflags="-s -w -X github.com/hoggerty/mcp-yandex-tracker/cmd.VERSION=${VERSION}" \
    -o /out/yandex-tracker-mcp .

# distroless/static содержит CA-сертификаты (нужны для HTTPS к API Трекера) и
# пользователя nonroot; shell и пакетного менеджера нет.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/yandex-tracker-mcp /usr/local/bin/yandex-tracker-mcp
# Используется только при `server --transport http --http-addr :8080`.
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/yandex-tracker-mcp"]
CMD ["server"]

# Сборка

```sh
make build
```

Версия бинарника (`--version`) берётся из `git describe --tags --always --dirty`
и прокидывается через `-ldflags` (см. `Makefile`). Обычный `go build -o yandex-tracker-mcp .`
тоже работает, но версия в этом случае будет `dev`.

## Ручная настройка `.env`

Публичный путь настройки — wizard (`./yandex-tracker-mcp init`, см. README). Вручную
`.env` нужен только при локальной разработке без wizard'а:

```sh
cp .env.example .env
# затем отредактировать .env своими значениями
```

## Релизы

Релизы собираются [GoReleaser](https://goreleaser.com) в GitHub Actions
(`.github/workflows/release.yml`) при пуше тега вида `vX.Y.Z`:

```sh
git tag v0.1.0
git push origin v0.1.0
```

Workflow прогоняет тесты, собирает бинарники для linux/darwin/windows × amd64/arm64,
упаковывает их в архивы (`.tar.gz`, для Windows — `.zip`) вместе с `README.md`,
`LICENSE` и `.env.example`, считает `checksums.txt` и публикует GitHub Release с changelog.

Версия прокидывается в `cmd.VERSION` через `-ldflags` (как и в `Makefile`), поэтому
`--version` и `serverInfo` показывают имя тега, например `v0.1.0`.

Теги с суффиксом (`v0.2.0-rc.1`) публикуются как pre-release и не становятся Latest.

Тот же workflow дополнительно собирает multi-arch Docker-образ (`linux/amd64`,
`linux/arm64`) и публикует его в GHCR как `ghcr.io/hoggerty/mcp-yandex-tracker`
с тегами `X.Y.Z` и `X.Y`; для тегов без суффикса (не pre-release) добавляется
`latest`. Локальная сборка образа: `make docker` (результат — `ghcr.io/hoggerty/mcp-yandex-tracker:dev`).

GoReleaser для локальной проверки не нужно устанавливать глобально — версия закреплена
в `tools/go.mod` (Go tool dependency). `make release-check`/`make release-snapshot` сами
соберут его в `tools/bin` при первом запуске.

Проверить сборку локально без публикации (результат — в `dist/`):

```sh
make release-check
make release-snapshot
```

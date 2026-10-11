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

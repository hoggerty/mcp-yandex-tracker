# ARCHITECTURE.md

Архитектура **yandex-tracker-mcp**: слои, зависимости между ними и где живёт сборка приложения.

---

## Слои

```
cmd  →  wire  →  inbound  →  logic  →  outbound
```

- `cmd` — CLI на cobra: парсит аргументы, запускает сервер или интерактивный wizard. Не знает ни про Tracker API,
  ни про MCP.
- `internal/wire` — composition root: собирает конфигурацию (`internal/wire/config`) и связывает `outbound`,
  `logic` и `inbound` в готовый к обслуживанию MCP-сервер (`wire.NewServer`).
- `internal/inbound` — адаптер MCP: схемы инструментов, разбор аргументов, форматирование результата
  (`inbound/result`). Знает про `logic`, но не про `trackerapi` напрямую.
- `internal/logic` — use-case сервисы (`issue.Service`, `queue.Service`, `user.Service`, `board.Service`). Знают про
  `trackerapi`, но ничего не знают про протокол MCP.
- `internal/outbound/trackerapi` — HTTP-клиент Yandex Tracker API. Ничего не знает ни про MCP, ни про use-case'ы.

Правило: тип или функция снизу вверх не видит то, что выше. `logic` не импортирует `inbound`, `outbound` не
импортирует `logic`, и ни один из трёх слоёв не импортирует `wire`.

## Сборка конфигурации

Сборка всех конфигов (чтение переменных окружения, их валидация, маппинг в `Config`-типы нижних слоёв) — задача
`internal/wire` и живёт в его подпакетах (`internal/wire/config`), а не в `outbound`/`logic`/`inbound`. Эти три слоя
принимают уже готовый `Config` и ничего не знают о том, что он пришёл из переменных окружения, а не из флагов или
файла. `outbound/trackerapi.Config` — обычная структура данных без тегов `env` и без функции `FromEnv`/`ConfigFromEnv`.

`cmd` обращается к сборке конфига только через `internal/wire/config`, не напрямую в `outbound`.

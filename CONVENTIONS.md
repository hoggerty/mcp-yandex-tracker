# CONVENTIONS.md

Соглашения, специфичные для **yandex-tracker-mcp**. Проверка: `make lint`, `make vet`, `make test`.

Слои, зависимости между ними и сборка конфигурации — в [ARCHITECTURE.md](ARCHITECTURE.md). Здесь — то, что из этой
архитектуры вытекает на уровне организации файлов и кода.

---

## Один файл — один метод

Пакет `logic/<domain>` и пакет `inbound/<domain>` описывают один набор операций дважды — как use-case и как
MCP-инструмент. Оба организованы одинаково:

- `service.go` (в `logic`) или `<domain>.go` (в `inbound`) содержит только package doc, объявление типа
  (`Service` / список `tool.Tool`) и конструктор (`New` / `Tools`). Новый метод сюда не дописывают.
- Каждый экспортируемый метод сервиса и каждый MCP-инструмент — отдельный файл, имя файла = имя операции в
  snake_case, без суффикса `_test`. Имя файла в `logic/<domain>` совпадает с именем файла в `inbound/<domain>`,
  который его вызывает:

  | Операция             | `logic/issue/*.go`       | `inbound/issue/*.go`      |
  |-----------------------|---------------------------|----------------------------|
  | получить задачу       | `get.go`                  | `get.go`                   |
  | список переходов      | `transitions_list.go`     | `transitions_list.go`      |
  | выполнить переход     | `transition_execute.go`   | `transition_execute.go`    |

  Это пара, а не совпадение: имя тянется от MCP tool (`issue_transitions_list`) вниз до use-case метода. Найти, что
  вызывает `getTool.Handle`, можно не читая импорты — просто открыть файл с тем же именем в `logic`.

Опускать этот разнос ради «это же одна строчка» не нужно — даже однострочный метод (`Service.Get`) живёт в своём
файле. Иначе пара имя-файла ↔ имя-метода ломается для всех остальных операций в том же сервисе.

Тот же принцип в `internal/wire/config`: `Config` только группирует секции, каждая секция — свой тип в своём файле
(`TrackerClientConfig` → `tracker_client.go`), `config.go` — только сам `Config` и `Load()`.

## Комментарии

Одна строка doc-комментария на экспортируемый тип/функцию, с именем в начале (`// Get retrieves...`). Комментарии,
объясняющие *что* делает код, а не *почему*, не пишутся — имя функции и так это говорит.

## Ошибки

`outbound/trackerapi` оборачивает транспортные ошибки через `fmt.Errorf("...: %w", err)` с глаголом-описанием шага
(`"encode request body"`, `"build request"`, `"request tracker API"`). Ответ API с кодом `>= 300` оборачивается в
`*APIError`, а не в `fmt.Errorf` — у него предсказуемая структура (`StatusCode`, `Body`), которую `inbound/result`
может использовать напрямую.

`logic` и `inbound` ошибку от `trackerapi` не оборачивают и не проверяют `errors.As` — она долетает до
`result.JSON`, который сам решает, как показать её MCP-клиенту. Добавлять контекст повторно (`fmt.Errorf("get issue:
%w", err)`) на этом пути не нужно: шаг, на котором ошибка возникла, уже виден из сообщения `trackerapi`.

## Новая операция

Чтобы добавить операцию в существующий домен (например, `issue`):

1. `outbound/trackerapi/<domain>.go` — метод клиента, вызывающий Tracker API.
2. `logic/<domain>/<operation>.go` — метод `Service`, вызывающий клиент.
3. `inbound/<domain>/<operation>.go` — тип `<operation>Tool` с `Description()`/`Handle()`.
4. Зарегистрировать тул в `Tools()` (`<domain>.go` в `inbound/<domain>`) и строку в таблице `README.md`.

## Тесты

- `make test` — юниты (`go test ./...`).
- `make vet`, `make lint` — статический анализ (`go vet`, `golangci-lint`).

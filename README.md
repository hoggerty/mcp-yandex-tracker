# yandex-tracker-mcp

MCP сервер на Go, который предоставляет операции
[Яндекс Трекера](https://tracker.yandex.ru/) как для
MCP-клиентов (Claude Code, Codex и т.д.).

## Инструменты (tools)

| Tool | Описание |
| --- | --- |
| `issue_get` | Получить задачу по ключу |
| `issue_search` | Поиск задач по языку запросов или структурированному фильтру |
| `issue_count` | Посчитать задачи, подходящие под запрос или фильтр |
| `issue_create` | Создать новую задачу |
| `issue_update` | Обновить поля задачи |
| `issue_transitions_list` | Список доступных переходов workflow |
| `issue_transition_execute` | Выполнить переход задачи |
| `issue_comments_list` | Список комментариев к задаче |
| `issue_comment_add` | Добавить комментарий (с возможностью призвать пользователей) |
| `issue_links_list` | Список связей задачи |
| `issue_link_add` | Связать две задачи |
| `issue_changelog_get` | Получить историю изменений полей задачи |
| `queues_list` | Список очередей |
| `queue_get` | Получить очередь по ключу |
| `user_current` | Профиль владельца токена |
| `users_list` | Список пользователей организации |
| `boards_list` | Список agile-досок |
| `board_sprints_list` | Список спринтов доски |

## Сборка

```sh
make build
```

Версия бинарника (`--version`) берётся из `git describe --tags --always --dirty`
и прокидывается через `-ldflags` (см. `Makefile`). Обычный `go build -o yandex-tracker-mcp .`
тоже работает, но версия в этом случае будет `dev`.

## Конфигурация

Сервер общается с клиентом через stdio и настраивается переменными
окружения. Их можно задать напрямую в окружении процесса либо положить в
файл `.env` в корне `yandex-tracker-mcp` — он подхватывается автоматически
при старте (если файла нет, это не ошибка). Пример со всеми переменными —
в [`.env.example`](.env.example):

```sh
cp .env.example .env
# затем отредактировать .env своими значениями
```

Переменные:

- `YANDEX_TRACKER_TOKEN` (обязательна) — OAuth-токен с доступом к Tracker
  API. См. [доступ к API](https://yandex.ru/support/tracker/ru/concepts/access).
- `YANDEX_TRACKER_ORG_ID` — ID организации, отправляется в заголовке
  `X-Org-ID` (организации Яндекс 360).
- `YANDEX_TRACKER_CLOUD_ORG_ID` — ID организации, отправляется в заголовке
  `X-Cloud-Org-ID` (организации Yandex Cloud). Должна быть задана ровно одна
  из `YANDEX_TRACKER_ORG_ID` / `YANDEX_TRACKER_CLOUD_ORG_ID`.
- `YANDEX_TRACKER_BASE_URL` — переопределяет базовый URL API (по умолчанию
  `https://api.tracker.yandex.net`).

## Запуск

Бинарник — это CLI на [cobra](https://github.com/spf13/cobra) с подкомандой
`server`. Без аргументов запускается `server` по умолчанию:

```sh
./yandex-tracker-mcp          # то же самое, что ./yandex-tracker-mcp server
./yandex-tracker-mcp --help
```

## Пример конфигурации клиента

```json
{
  "mcpServers": {
    "yandex-tracker": {
      "command": "/absolute/path/to/yandex-tracker-mcp",
      "env": {
        "YANDEX_TRACKER_TOKEN": "y0_...",
        "YANDEX_TRACKER_ORG_ID": "1234567"
      }
    }
  }
}
```

### Claude Code

```sh
claude mcp add -s user \
  -e YANDEX_TRACKER_TOKEN=y0_... \
  -e YANDEX_TRACKER_ORG_ID=1234567 \
  --transport stdio yandex-tracker /absolute/path/to/yandex-tracker-mcp
```

## Справочники по API Яндекс Трекера

- [Обзор API](https://yandex.ru/support/tracker/ru/about-api)
- [Доступ (OAuth / IAM)](https://yandex.ru/support/tracker/ru/concepts/access)
- [Общий формат запросов](https://yandex.ru/support/tracker/ru/common-format)
- [Коды ошибок](https://yandex.ru/support/tracker/ru/error-codes)

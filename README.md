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
| `issue_move` | Перенести задачу в другую очередь |
| `issue_suggest` | Подсказки задач по тексту (автодополнение) |
| `issue_search_release` | Освободить ресурсы scroll-поиска задач |
| `issue_checklist_add` | Добавить пункт в чеклист задачи |
| `issue_checklist_list` | Список пунктов чеклиста задачи |
| `issue_checklist_edit` | Массово изменить пункты чеклиста задачи |
| `issue_checklist_delete` | Удалить чеклист задачи целиком |
| `issue_checklist_item_delete` | Удалить один пункт чеклиста задачи |
| `issue_comment_get` | Получить комментарий задачи по id |
| `issue_comment_edit` | Изменить текст комментария задачи |
| `issue_comment_delete` | Удалить комментарий задачи |
| `issue_comment_reaction_add` | Поставить реакцию на комментарий |
| `issue_links_list_paginate` | Список связей задачи постранично |
| `issue_link_delete` | Удалить связь между задачами |
| `issue_attachment_list` | Список файлов, прикреплённых к задаче |
| `issue_attachment_get` | Метаданные прикреплённого файла |
| `report_create` | Сформировать отчёт по задачам (xlsx/csv/xml) |
| `report_search` | Найти ранее сформированные отчёты |
| `worklog_add` | Добавить запись о затраченном времени к задаче |
| `worklog_edit` | Изменить запись о затраченном времени |
| `worklog_delete` | Удалить запись о затраченном времени |
| `worklog_list` | Список записей о затраченном времени по задаче |
| `worklog_search` | Поиск записей о затраченном времени по всем задачам |
| `externallink_applications_list` | Список внешних приложений (источников внешних связей) |
| `externallink_list` | Список внешних связей задачи |
| `externallink_add` | Добавить внешнюю связь к задаче |
| `externallink_delete` | Удалить внешнюю связь задачи |
| `bulkchange_move` | Массово перенести задачи в другую очередь |
| `bulkchange_update` | Массово изменить поля задач |
| `bulkchange_transition` | Массово выполнить переход workflow для задач |
| `bulkchange_status` | Статус массовой операции |
| `bulkchange_errors` | Задачи с ошибками в массовой операции |
| `field_list` | Список глобальных полей организации |
| `field_create` | Создать глобальное поле |
| `field_get` | Получить параметры глобального поля |
| `field_edit_name` | Переименовать глобальное поле |
| `field_edit_values` | Изменить параметры поля (категория, видимость, значения) |
| `field_category_create` | Создать категорию полей |
| `field_category_edit` | Изменить категорию полей |
| `filter_create` | Создать сохранённый фильтр задач |
| `filter_get` | Получить сохранённый фильтр по id |
| `filter_update` | Изменить сохранённый фильтр |
| `filter_delete` | Удалить сохранённый фильтр |
| `workflow_create` | Создать workflow |
| `workflow_list` | Список всех workflow организации |
| `workflow_queue_list` | Workflow, назначенные типам задач очереди |
| `workflow_get` | Получить workflow по id |
| `workflow_edit` | Изменить workflow |
| `workflow_action_edit` | Изменить действие шага workflow |
| `workflow_delete` | Удалить workflow |
| `autoaction_create` | Создать автодействие на очереди |
| `autoaction_get` | Получить параметры автодействия |
| `autoaction_logs_list` | История запусков автодействия |
| `autoaction_log_get` | Один запуск автодействия с затронутыми задачами |
| `trigger_create` | Создать триггер на очереди |
| `trigger_list` | Список триггеров очереди |
| `trigger_get` | Получить параметры триггера |
| `trigger_edit` | Изменить триггер |
| `trigger_logs_list` | Журнал вызовов вебхука триггера |
| `macro_list` | Список макросов очереди |
| `macro_get` | Получить макрос по id |
| `macro_create` | Создать макрос на очереди |
| `macro_edit` | Изменить макрос |
| `macro_delete` | Удалить макрос |
| `component_list` | Список всех компонентов организации |
| `component_queue_list` | Список компонентов очереди |
| `component_get` | Получить компонент по id |
| `component_create` | Создать компонент на очереди |
| `component_edit` | Изменить компонент |
| `component_delete` | Удалить компонент |
| `component_access_user_get` | Права пользователя на компонент |
| `component_access_group_get` | Права группы на компонент |
| `version_create` | Создать версию на очереди |
| `version_queue_list` | Список версий очереди |
| `version_get` | Получить версию по id |
| `version_edit` | Изменить версию |
| `version_delete` | Удалить версию |
| `queues_list` | Список очередей |
| `queue_get` | Получить очередь по ключу |
| `queue_create` | Создать очередь |
| `queue_delete` | Удалить очередь |
| `queue_restore` | Восстановить удалённую очередь |
| `queue_fields_list` | Список обязательных полей очереди |
| `queue_tags_list` | Список тегов очереди |
| `queue_tag_delete` | Удалить тег очереди |
| `queue_access_manage` | Изменить права доступа к очереди |
| `queue_access_user_get` | Права пользователя в очереди |
| `queue_access_group_get` | Права группы в очереди |
| `queue_local_field_create` | Создать локальное поле очереди |
| `queue_local_fields_list` | Список локальных полей очереди |
| `queue_local_field_get` | Получить локальное поле очереди |
| `queue_local_field_edit` | Изменить локальное поле очереди |
| `user_current` | Профиль владельца токена |
| `users_list` | Список пользователей организации |
| `user_get` | Получить пользователя по логину или id |
| `user_list_relative` | Список пользователей с постраничной навигацией относительно id |
| `boards_list` | Список agile-досок |
| `board_sprints_list` | Список спринтов доски |
| `board_get` | Получить agile-доску по id |
| `board_create` | Создать agile-доску |
| `board_edit` | Изменить agile-доску |
| `board_delete` | Удалить agile-доску |
| `column_list` | Список колонок доски |
| `column_get` | Получить колонку доски по id |
| `column_create` | Создать колонку на доске |
| `column_edit` | Изменить колонку доски |
| `column_delete` | Удалить колонку доски |
| `sprint_get` | Получить спринт по id |
| `sprint_create` | Создать спринт на доске |
| `sprint_edit` | Изменить спринт |
| `sprint_start` | Запустить спринт |
| `sprint_archive` | Архивировать спринт |
| `sprint_delete` | Удалить спринт |
| `dashboard_create` | Создать дашборд |
| `dashboard_cycle_time_widget_create` | Создать виджет Cycle time на дашборде |
| `issuetype_list` | Список типов задач организации |
| `issuetype_create` | Создать тип задачи |
| `issuetype_edit` | Изменить тип задачи |
| `status_list` | Список статусов задач организации |
| `status_create` | Создать статус задачи |
| `status_edit` | Изменить статус задачи |
| `resolution_list` | Список резолюций задач организации |
| `resolution_create` | Создать резолюцию |
| `resolution_edit` | Изменить резолюцию |
| `priority_list` | Список приоритетов задач организации |
| `priority_create` | Создать приоритет |
| `priority_edit` | Изменить приоритет |
| `entity_create` | Создать сущность (проект, портфель или цель) |
| `entity_get` | Получить сущность по идентификатору |
| `entity_update` | Изменить сущность |
| `entity_delete` | Удалить сущность |
| `entity_search` | Список сущностей заданного типа по фильтру |
| `entity_bulk_update` | Массово изменить несколько сущностей одного типа |
| `entity_events_list` | История изменений сущности |
| `entitycomment_add` | Добавить комментарий к сущности |
| `entitycomment_edit` | Изменить комментарий сущности |
| `entitycomment_list` | Список комментариев сущности |
| `entitycomment_get` | Получить комментарий сущности |
| `entitycomment_delete` | Удалить комментарий сущности |
| `entitychecklist_add` | Добавить пункт чек-листа сущности |
| `entitychecklist_edit` | Массово изменить чек-лист сущности |
| `entitychecklist_item_edit` | Изменить пункт чек-листа сущности |
| `entitychecklist_item_move` | Переместить пункт чек-листа сущности |
| `entitychecklist_delete` | Удалить все пункты чек-листа сущности |
| `entitychecklist_item_delete` | Удалить пункт чек-листа сущности |
| `entityattachment_list` | Список вложений сущности |
| `entityattachment_get` | Получить вложение сущности |
| `entityattachment_delete` | Удалить вложение сущности |
| `entitylink_add` | Добавить связь сущности |
| `entitylink_list` | Список связей сущности |
| `entitylink_delete` | Удалить связь сущности |
| `entityaccess_access_get` | Получить расширенные права доступа к сущности |
| `entityaccess_permissions_get` | Получить права доступа к сущности |
| `entityaccess_access_edit` | Изменить расширенные права доступа к сущности |
| `entityaccess_permissions_edit` | Изменить права доступа к сущности |
| `project_create` | Создать проект |
| `project_get` | Получить проект по идентификатору |
| `project_list` | Список всех проектов |
| `project_queues_list` | Список очередей проекта |
| `project_update` | Изменить проект |
| `project_delete` | Удалить проект |
| `import_ticket` | Импортировать задачу с сохранением исторических метаданных (только администратор) |
| `import_comment` | Импортировать комментарий задачи (только администратор) |
| `import_link` | Импортировать связь задачи (только администратор) |
| `import_worklog` | Импортировать запись времени задачи (только администратор) |
| `gap_create` | Создать отсутствия сотрудников (только администратор) |
| `gap_search` | Найти отсутствия сотрудников (только администратор) |
| `gap_delete` | Удалить отсутствия сотрудников (только администратор) |

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

Либо запустить интерактивный wizard — он спросит токен и ID организации,
запишет `.env`, а затем предложит сразу зарегистрировать сервер в Claude Code
или Codex (или покажет JSON-сниппет для других клиентов):

```sh
./yandex-tracker-mcp init
```

Тот же выбор доступен и без wizard'а — отдельными командами. Каждая из
них, если обязательные переменные ещё не заполнены, сама предложит
запустить `init`:

```sh
./yandex-tracker-mcp install-claude   # зарегистрировать в Claude Code
./yandex-tracker-mcp install-codex    # зарегистрировать в Codex
./yandex-tracker-mcp snippet          # вывести JSON-сниппет для других клиентов
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

# Контракт бэкенд ↔ judge

Контракт между бэкендом и judge. Бэкенд и judge общаются только через очередь (Redis Streams) и типы из этого пакета.
Благодаря этому оба пишут код параллельно: бэкенд можно запускать с заглушкой judge,
а judge — тестировать без бэкенда.

**Правило:** изменения в `api/` — только через pull request, который одобрили оба.

## Поток

1. Бэкенд кладёт `Submission` в поток `judge:submissions` (поле `payload`, JSON).
2. Свободный judge-воркер забирает его (группа `judges`), проверяет посылку.
3. Judge кладёт `Result` в поток `judge:results`.
4. Бэкенд читает результат, считает урон и рассылает событие игрокам.

## Пример

Посылка:

```json
{
  "id": "sub-42",
  "match_id": "m-7",
  "round": 2,
  "player_id": "u-1",
  "problem_id": "sort-numbers",
  "language": "cpp",
  "source": "#include <bits/stdc++.h>\n..."
}
```

Результат:

```json
{
  "submission_id": "sub-42",
  "verdict": "AC",
  "rungs": [
    {"size_exp": 10, "passed": true, "instructions": 1830211, "memory_bytes": 3145728},
    {"size_exp": 11, "passed": true, "instructions": 3922450, "memory_bytes": 3145728}
  ],
  "rungs_passed": 11,
  "rungs_total": 13,
  "score": 0.846,
  "complexity": "O(n log n)"
}
```

## Подключение

```go
import "github.com/algoduel/algoduel-api"

var s api.Submission
```

При разработке из родительского репозитория `algoduel` ничего подключать не нужно:
`go.work` подставляет локальную папку `api/`. Для сборки репозитория отдельно
(например, в CI) один раз выполните в нём `go get github.com/algoduel/algoduel-api@main`.

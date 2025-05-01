# GitFame

Утилита для анализа вклада разработчиков в Git-репозиторий

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/gitfame)](https://goreportcard.com/report/github.com/yourusername/gitfame)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Основные возможности

- **3 ключевые метрики**:
  - Количество изменённых строк (LOC)
  - Количество уникальных коммитов
  - Затронутые файлы

- **Гибкая фильтрация**:
  - По расширениям файлов (`.go,.md`)
  - По языкам программирования (`go,markdown`)
  - Glob-паттерны включения/исключения

- **Форматы вывода**:
  - Табличный (с авто-форматированием)
  - CSV
  - JSON
  - JSON Lines

- **Производительность**:
  - Оптимизированная обработка через `git blame --porcelain`
  - Прогресс-бар для длительных операций (TODO)

## Установка ⚙️

### Через Go:
```bash
go install gitlab.com/justnurik/gitfame/cmd/gitfame@latest
```

### Сборка из исходников:
```bash
git clone https://github.com/justnurik/gitfame.git
cd gitfame/cmd/gitfame && go build -o gitfame .
```

## Использование

Базовый пример:
```bash
gitfame --repository=. --order-by=lines
```

Расширенный сценарий:
```bash
gitfame \
  --revision=HEAD~10 \
  --extensions=".go,.md" \
  --languages="go,markdown" \
  --exclude="vendor/*,testdata/*" \
  --format=json \
  --use-committer
```

### Все флаги
| Флаг             | Описание                                  | По умолчанию    |
|-------------------|------------------------------------------|-----------------|
| `--repository`    | Путь к репозиторию                       | Текущая директория |
| `--revision`      | Анализируемая ревизия                    | HEAD            |
| `--order-by`      | Критерий сортировки (lines/commits/files)| lines           |
| `--format`        | Формат вывода (tabular/csv/json/json-lines)| tabular       |
| `--extensions`    | Фильтр по расширениям файлов             | Все             |
| `--languages`     | Фильтр по языкам программирования        | Все             |
| `--exclude`       | Исключающие Glob-паттерны                | Нет             |
| `--restrict-to`   | Включающие Glob-паттерны                 | Все             |
| `--use-committer` | Использовать коммиттера вместо автора    | false           |

## Пример вывода

```text
Name                   Lines    Commits  Files
Alice Anderson         15028    92       47
Bob Brown              2843     15       12
Charlie Clark          762      8        5
```

## Технологический стек

- **Язык**: Go 1.20+
- **CLI Framework**: Cobra + pflags
- **Тестирование**: Go test + интеграционные тесты
- **Обработка Git**: Нативный вызов git-команд через os/exec
- **Форматирование**: text/tabwriter, encoding/csv, encoding/json

## Разработка

### Запуск тестов
```bash
go test -v ./test/integration/... -coverprofile=coverage.out
```

## Лицензия 📄

MIT License. Подробности в файле [LICENSE](LICENSE).

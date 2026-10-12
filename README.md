# Разработка на Go

[![test](https://github.com/masesmirnov/x5-go-course/actions/workflows/test.yml/badge.svg)](https://github.com/masesmirnov/x5-go-course/actions/workflows/test.yml)
[![go](https://img.shields.io/github/go-mod/go-version/masesmirnov/x5-go-course)](go.mod)

Мои домашки по курсу X5 Tech в ВШЭ, осень 2026.

| Урок | Тема | Задачи |
| --- | --- | --- |
| 2 | [Типы данных, управляющие конструкции и функции](lesson02) | FizzBuzz, статистика по суммам, категория заказа |

Каждый пуш GitHub сам прогоняет через gofmt, golangci-lint и тесты на Linux, macOS и Windows, а вывод всех программ видно прямо в прогоне. Локально всё запускается из корня:

```sh
git clone https://github.com/masesmirnov/x5-go-course
cd x5-go-course
go test ./...
```

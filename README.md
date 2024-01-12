# SPBU-Architecture-and-Design
Репозиторий с материлами для проетка по курсу СПбГУ МКН РПО Арихитектура и Проектирование Информационных систем

## Структура папок

```
/pkg/                полноценные библиотеки, предназначенные для использования проектами как внутри проекта
  /lib/              код библиотеки `lib`
/projects/           весь остальной код *без внешних зависимостей*, разбитый на "проекты" — независимые друг от друга наборы связанных бинарников
  /proj/             код проета `proj`
    /cmd/            все бинарники проекта `proj`
      /proj-engine/  proj.go (`package main`), Dockerfile, пример конфига... 
      /proj-client/  proj.go (`package main`), Dockerfile, пример конфига... 
    /internal/       все не-main Go пакеты проекта `proj`
    /Makefile        Makefile для сборки проекта `proj`
```

## Зависимости

- go1.20
- Docker
- [`goimports`](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [`staticcheck`](https://staticcheck.io/)

# Ссылка на гугл док с решением
https://docs.google.com/document/d/1UXTUbpHfLhNIqYzQq8N0F6ba0GpAgeGG2fByj9eWxhY/edit?usp=sharing

# Go Project

export DATA_SOURCE_NAME="user=postgres dbname=postgres sslmode=disable password=123"

sudo docker run -e POSTGRES_PASSWORD=123 --rm -p 5432:5432 postgres

https://hub.docker.com/_/postgres

sudo docker exec -it -u postgres peaceful_zhukovsky psql

## Next

- Подробнее про функции: передача аргументов, "formal" и "actual" function parameters
- Парсинг параметров командной строки: пакет flag https://pkg.go.dev/flag или фреймворк viper
- Попробовать запускать метод createSchema только если проброшен флаг --create-schema

* createTestPersons - реализовать тело функции: создает несколько структур Persons и возвращает их в слайсе
* createTestData - изменить поведение: убрать строчки с захардкоженными Persons. Вместо этого, в цикле проходит по полученному слайсу persons и выполняет INSERT sql запрос по подобию того как это написано сейчас

Темы: циклы, структуры, слайсы


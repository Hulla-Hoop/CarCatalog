# CarCatalog 
 
## Описание

Простой каталог автомобилей

## Запуск 

Запустить проект можно командой:
make run 

## Rest методы

swagger спецификация лежит в ./docs

### Добавление 

- curl

``` 
curl --location 'localhost:8090/insert' \
--header 'Content-Type: application/json' \
--data '{
    "regNums":["x120xx150","x123xx132"]
}'

```
- swagger

```
/insert:
    post:
      tags:
        - catalog
      summary: Добавление машин
      description: получает на вход  слайс гос номеров запрашивает по ним информацию и возвращает добавленные объекты
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                regNum:
                  $ref: '#/components/schemas/regNum'
      responses:
        '200':
          description: Возвращает слайс добавленых объектов
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/cars'          
        '500':
          description: Выводит ошибку
 ```


### Удаление 
- curl

```curl --location --request POST 'localhost:8090/delete?id=2' \
--data ''
 ```

- swagger

```
 /delete:
    post:
      tags: 
      - catalog
      summary: удаляет машину
      description: Возращает удаленную машину
      parameters:
        - name: id
          in: query
          description: id категории
          required: true
          schema:
            type: string
            default: 
      responses:
        '200':
          description: Возращает удаленную машину
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Car'
        '400':
          description: Пустой параметр id
        '500':
          description: Выводит ошибку
 ```
### Обновление

- curl 

```
curl --location 'localhost:8090/update?id=3' \
--header 'Content-Type: application/json' \
--data '{
    "regNum": "x120xx150",
    "mark": "Honda",
    "model": "Camry",
    "year": 2009,
    "owner": {
        "name": "Alie",
        "surname": "Jones",
        "patronymic": "Liam"
    }
}'

```

- swagger

```
 /update:
    post:
      tags: 
      - catalog
      summary: обновляет машину
      description: Возращает обновленную машину
      parameters:
        - name: id
          in: query
          description: id машины
          required: true
          schema:
            type: string
            default:
      requestBody:
        required: true
        description: айди не нужно заполнять и не обязательно указывать все поля обновятся только те поля что указаны
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Car'
      responses:
        '200':
          description: Возвращает обновленую запись
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Car'
        '400':
          description:  Пустой параметр id
        '500':
          description: Выводит ошибку
 ```

### Фильтрация и пагинация
- curl 

```
curl --location 'localhost:8090/filter?limit=&offset=&operator=&value=Ford&field=mark'
```

- swagger
```
  /filter:
    post:
      tags: 
      - catalog
      summary: Фильтрация и пагинация
      description: Возращает список машин удовлетворяющих фильтру, если не указывать параметры то выведет весь список машин
      parameters:
        - name: limit
          in: query
          description: количество записей на странице(не работает без offset)
          required: false
          schema:
            type: string
            default:
        - name: offset
          in: query
          description: c какой записи начинается страница(не работает без limit)
          required: false
          schema:
            type: string
            default:
        - name: operator
          in: query
          description: оператор сравнения field+operator+value(не имеет смысла вводить без field,value)
          required: false
          schema:
            type: string
            enum:
              - "eq"
              - "ne"
              - "gt"
              - "ge"
              - "lt"
              - "le"
        - name: value
          in: query
          description: значение которому должно соответствовать поле(не имеет смысла вводить без operator,field)
          required: false
          schema:
            type: string
            default:
        - name: field
          in: query
          description: поле по которому фильтруют записи(не имеет смысла вводить без operator,value)
          required: false
          schema:
            type: string
            enum:
              - "regNum"
              - "mark"
              - "model"
              - "year"
              - "name"
              - "surname"
              - "patronymic"
      responses:
        '200':
          description: Возвращает обновленую запись
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/cars'
        '500':
          description: Выводит ошибку
 ```
## Api для получения рандомных данных об автомобилях по гос номеру

- Лежит в папке ./CarApi и запускается командой make run.
- Сгенирировал библиотекой предоставляемой swagger.io по спецификации лежащей в папке ./CarApi/api/swagger.yaml и заполнил простыми функциями генирации не применяя потернов проектирования и не покрывая логами.
- Сервис нужен только для получения рандомных данных.

## База данных

- Используется Postgres размещеный в докер композе.
- Таблица создается путем миграций.
- Миграции применяются автоматически при запуске приложения также можно использовать команды: migrate-up/migrate-down
- Структура таблицы [!Структура Таблицы](/docs/carcatalog.png)
- Поместил все в одну таблицу так как небыло условий что у чеовека может быть несколько машин или наоборот

## Логи 
- Для логирования спользуется библиотека Логрус
- Добавлены реквест айди для сквозного логирования 
```
func ReqID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), "reqID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

 ```
- Добавлен уровень Debug логов

## Конфигурационные данные 
- Лежат в .env файле
- Подгружаются в пакете ./internal/config
``` 
#DB config
DB_HOST=localhost
DB_NAME=test
DB_USER=postgres
DB_PASSWORD=12345678
DB_PORT=5432
DB_SSL=disable

# Server config

SERVER_HOST = localhost
SERVER_PORT = 8090

# Remote Api 

LINK = http://localhost:8080/info?regNum=

```
# 4. Парсер: OID

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [GET] /api/v1/catalog/oid - Поиск OID по точной dotter notation

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

### Возможные коды ошибок

400: Bad Request - Параметр notation отсутствует или имеет некорректный формат
500: Internal Server Error - Внутренняя ошибка парсера или СУБД при поиске OID

---

## [GET] /api/v1/catalog/oid/exact - Поиск OID по dotter notation, названию MIB и производителю

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

### Возможные коды ошибок

400: Bad Request - Пропущены обязательные параметры запроса notation или mib
500: Internal Server Error - Ошибка СУБД при точном поиске записи OID

---

## [GET] /api/v1/catalog/oid/mib - Получить OID по названию MIB

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

### Возможные коды ошибок

400: Bad Request - Обязательный параметр названия MIB (name) не передан
500: Internal Server Error - Ошибка базы данных при выборке объектов MIB-файла

---

## [GET] /api/v1/catalog/oid/prefix - Поиск OID по префиксу с пагинацией

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

### Возможные коды ошибок

400: Bad Request - Отсутствует обязательный параметр префикса (prefix)
500: Internal Server Error - Ошибка пагинации или выполнения префиксного поиска в СУБД

---

## [GET] /api/v1/catalog/oid/vendor - Получить OID по производителю с пагинацией

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

### Возможные коды ошибок

500: Internal Server Error - Ошибка сервера или СУБД при получении объектов по вендору

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
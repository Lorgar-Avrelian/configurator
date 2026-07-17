# 7. Конфигурация: Сопоставления параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/mapping - Создать сопоставление

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

400: Bad Request - Ошибка валидации структуры или обязательных параметров сопоставления  
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении сопоставления

---

## [GET] /api/v1/catalog/mapping/{id} - Получить сопоставление по ID

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

400: Bad Request - Неверный числовой формат идентификатора сопоставления в пути  
404: Not Found - Сопоставление с указанным ID не найдено в системе  
500: Internal Server Error - Системная ошибка базы данных при извлечении сопоставления

---

## [PUT] /api/v1/catalog/mapping/{id} - Обновить сопоставление по ID

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

400: Bad Request - Некорректный ID в пути или ошибка валидации JSON-структуры запроса  
404: Not Found - Сопоставление для обновления с указанным ID не найдено  
500: Internal Server Error - Ошибка обновления записи на стороне базы данных

---

## [DELETE] /api/v1/catalog/mapping/{id} - Удалить сопоставление по ID

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

400: Bad Request - Некорректный формат ID сопоставления в пути запроса  
404: Not Found - Удаляемое сопоставление не обнаружено в базе данных  
500: Internal Server Error - Ошибка целостности СУБД при удалении сопоставления

---

## [GET] /api/v1/catalog/mapping/{id}/own - Получить изолированное сопоставление по ID

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

400: Bad Request - Неверный числовой формат идентификатора сопоставления  
404: Not Found - Изолированное сопоставление с данным ID не найдено  
500: Internal Server Error - Внутренняя ошибка сервера при чтении записи

---

## [PATCH] /api/v1/catalog/mapping/{prevId}/{newId} - Изменить ID сопоставления

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

400: Bad Request - Неверный формат одного из ID (prevId или newId) в пути запроса  
404: Not Found - Сопоставление с исходным ID не существует в системе  
500: Internal Server Error - Ошибка СУБД при транзакции изменения идентификатора

---

## [GET] /api/v1/catalog/mappings - Получить все сопоставления

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

500: Internal Server Error - Системная ошибка сервера при извлечении полного списка сопоставлений

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
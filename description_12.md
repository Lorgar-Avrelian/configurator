# 12. Конфигурация: Пороги

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/threshold - Создать порог

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

400: Bad Request - Ошибка валидации переданной JSON-структуры порога или пропущены обязательные поля  
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении структуры порога

## [GET] /api/v1/catalog/threshold/{id} - Получить порог по ID

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

400: Bad Request - Неверный числовой формат идентификатора порога в пути  
404: Not Found - Порог с указанным ID не найден в системе  
500: Internal Server Error - Системная ошибка базы данных при извлечении структуры порога

---

## [PUT] /api/v1/catalog/threshold/{id} - Обновить порог по ID

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

400: Bad Request - Ошибка валидации структуры тела или неверный числовой формат ID в пути  
404: Not Found - Обновляемый порог с указанным ID не найден  
500: Internal Server Error - Внутренняя ошибка СУБД при обновлении полей порога

---

## [DELETE] /api/v1/catalog/threshold/{id} - Удалить порог по ID

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

400: Bad Request - Некорректный числовой формат ID порога в пути запроса  
404: Not Found - Удаляемый порог не обнаружен в базе данных  
500: Internal Server Error - Ошибка целостности СУБД при каскадном удалении порога

---

## [GET] /api/v1/catalog/threshold/{id}/from-string - Получить эквивалентную строку выражения порога по ID

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

400: Bad Request - Неверный формат ID в пути запроса  
404: Not Found - Порог с указанным ID не существует в системе  
500: Internal Server Error - Ошибка десериализации дерева условий JSONB в текстовую строку

---

## [PUT] /api/v1/catalog/threshold/{id}/from-string - Обновить порог по ID из строки

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

400: Bad Request - Синтаксическая ошибка при парсинге строкового выражения порога  
404: Not Found - Обновляемый порог не найден  
500: Internal Server Error - Внутренняя ошибка сервера при конвертации строки в AST-дерево СУБД

---

## [PATCH] /api/v1/catalog/threshold/{prevId}/{newId} - Изменить ID порога

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

400: Bad Request - Неверный числовой формат одного из ID в пути запроса  
404: Not Found - Исходный порог с данным ID не найден  
500: Internal Server Error - Ошибка обновления первичного или внешних ключей в базе данных

---

## [POST] /api/v1/catalog/threshold/from-string - Создать порог из эквивалентной строки

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

400: Bad Request - Ошибка валидации структуры или синтаксическая ошибка в текстовом выражении query  
500: Internal Server Error - Ошибка построения AST-структуры условий и её сохранения в JSONB

---

## [GET] /api/v1/catalog/thresholds - Получить все пороги

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

500: Internal Server Error - Критическая ошибка сервера при чтении полного списка порогов из БД

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
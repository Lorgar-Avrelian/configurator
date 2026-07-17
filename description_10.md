# 10. Конфигурация: Конфигурации по-умолчанию

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/default-configuration - Создать конфигурацию по умолчанию

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации или отсутствие обязательного query-параметра indicator  
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении записи

---

## [GET] /api/v1/catalog/default-configuration/{id} - Получить конфигурацию по умолчанию по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный числовой формат идентификатора конфигурации в пути  
404: Not Found - Конфигурация по умолчанию с указанным ID не найдена  
500: Internal Server Error - Системная ошибка СУБД при чтении конфигурации

---

## [PUT] /api/v1/catalog/default-configuration/{id} - Обновить конфигурацию по умолчанию по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации query-параметров или неверный формат ID в пути  
404: Not Found - Конфигурация по умолчанию с указанным ID не обнаружена  
500: Internal Server Error - Внутренняя ошибка СУБД при обновлении связей конфигурации

---

## [DELETE] /api/v1/catalog/default-configuration/{id} - Удалить конфигурацию по умолчанию по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат идентификатора конфигурации  
404: Not Found - Конфигурация по умолчанию с указанным ID не найдена  
500: Internal Server Error - Ошибка целостности СУБД при удалении записи

---

## [PATCH] /api/v1/catalog/default-configuration/{prevId}/{newId} - Изменить ID конфигурации по умолчанию

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат prevId или newId в пути запроса  
404: Not Found - Конфигурация по умолчанию с исходным ID не существует  
500: Internal Server Error - Ошибка транзакции изменения идентификатора на уровне СУБД

---

## [GET] /api/v1/catalog/default-configurations - Получить все конфигурации по умолчанию

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1
```

Ответ 1:

```json

```

</details>

</details>

### Возможные коды ошибок

500: Internal Server Error - Критическая ошибка сервера при чтении списка конфигураций по умолчанию

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
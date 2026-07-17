# 6. Конфигурация: Индикаторы параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/indicator/param - Создать индикатор параметров

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

400: Bad Request - Ошибка валидации структуры (не переданы oid_id или dotter_notation)  
500: Internal Server Error - Ошибка вставки записи индикатора параметра в СУБД

---

## [GET] /api/v1/catalog/indicator/param/{id} - Получить индикатор параметров по ID

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

400: Bad Request - Неверный числовой формат идентификатора в пути URL  
404: Not Found - Индикатор параметров с указанным ID не найден  
500: Internal Server Error - Внутренняя ошибка базы данных

---

## [PUT] /api/v1/catalog/indicator/param/{id} - Обновить индикатор параметров по ID

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

400: Bad Request - Ошибка валидации JSON-тела или некорректный ID в пути  
404: Not Found - Обновляемый индикатор параметров не обнаружен в базе данных  
500: Internal Server Error - Системная ошибка СУБД при сохранении данных

---

## [DELETE] /api/v1/catalog/indicator/param/{id} - Удалить индикатор параметров по ID

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

400: Bad Request - Некорректный числовой идентификатор в пути запроса  
404: Not Found - Индикатор параметров не существует  
500: Internal Server Error - Не удалось удалить индикатор параметров из-за ошибки СУБД

---

## [GET] /api/v1/catalog/indicator/params - Получить все индикаторы параметров

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

500: Internal Server Error - Внутренняя ошибка сервера при чтении индикаторов параметров

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
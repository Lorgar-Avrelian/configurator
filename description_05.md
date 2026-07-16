# 5. Конфигурация: Индикаторы устройств

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/indicator/device - Создать индикатор устройства

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

400: Bad Request - Нарушена валидация структуры индикатора (например, пустой object_id)
500: Internal Server Error - Не удалось записать индикатор устройства в базу данных

---

## [GET] /api/v1/catalog/indicator/device/{id} - Получить индикатор по ID

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

400: Bad Request - Идентификатор в пути запроса имеет неверный формат
404: Not Found - Индикатор устройства с указанным ID не обнаружен
500: Internal Server Error - Внутренняя ошибка СУБД при чтении индикатора

---

## [PUT] /api/v1/catalog/indicator/device/{id} - Обновить индикатор по ID

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

400: Bad Request - Некорректный ID в пути или ошибка валидации JSON-структуры тела
404: Not Found - Индикатор устройства для обновления не найден в системе
500: Internal Server Error - Внутренняя ошибка сервера при изменении индикатора

## [DELETE] /api/v1/catalog/indicator/device/{id} - Удалить индикатор по ID

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

400: Bad Request - Неверный числовой формат идентификатора в пути
404: Not Found - Индикатор устройства с указанным ID не найден
500: Internal Server Error - Системная ошибка каскадного удаления на стороне БД

---

## [GET] /api/v1/catalog/indicator/devices - Получить все индикаторы устройств

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

500: Internal Server Error - Критическая ошибка сервера при чтении списка индикаторов

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
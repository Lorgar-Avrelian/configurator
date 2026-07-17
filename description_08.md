# 8. Конфигурация: Структура компонентов устройства

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/device-component - Создать составную часть устройства

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

400: Bad Request - Нарушена валидация структуры (например, отсутствует обязательный model_id)  
500: Internal Server Error - Не удалось записать составную часть устройства в базу данных

---

## [GET] /api/v1/catalog/device-component/{id} - Получить составную часть устройства по ID

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

400: Bad Request - Идентификатор составной части в пути имеет неверный формат  
404: Not Found - Составная часть устройства с указанным ID не найдена  
500: Internal Server Error - Системная ошибка СУБД при чтении составной части

---

## [PUT] /api/v1/catalog/device-component/{id} - Обновить составную часть устройства по ID

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

400: Bad Request - Некорректный ID в пути или ошибка валидации JSON-тела запроса  
404: Not Found - Составная часть устройства для обновления не найдена  
500: Internal Server Error - Внутренняя ошибка сервера при изменении составной части

---

## [DELETE] /api/v1/catalog/device-component/{id} - Удалить составную часть устройства по ID

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

400: Bad Request - Неверный формат идентификатора составной части в пути  
404: Not Found - Составная часть устройства с указанным ID не существует  
500: Internal Server Error - Системная ошибка каскадного удаления на стороне БД

---

## [GET] /api/v1/catalog/device-component/{id}/own - Получить изолированную составную часть устройства по ID

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

400: Bad Request - Некорректный числовой формат идентификатора составной части  
404: Not Found - Составная часть устройства с данным ID не найдена  
500: Internal Server Error - Внутренняя ошибка сервера при чтении изолированной записи

---

## [PATCH] /api/v1/catalog/device-component/{prevId}/{newId} - Изменить ID составной части устройства

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
404: Not Found - Составная часть устройства с исходным ID не существует  
500: Internal Server Error - Ошибка СУБД при изменении первичного или внешнего ключа

---

## [GET] /api/v1/catalog/device-components - Получить всю структуру составных частей устройств

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

500: Internal Server Error - Критическая ошибка сервера при формировании дерева составных частей

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
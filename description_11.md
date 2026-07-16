# 11. Конфигурация: Конфигурации устройств

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/configuration - Создать рабочую конфигурацию

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

400: Bad Request - Ошибка валидации или отсутствие обязательного query-параметра indicator
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении рабочей конфигурации

---

## [GET] /api/v1/catalog/configuration/{id} - Получить рабочую конфигурацию по ID

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

400: Bad Request - Некорректный числовой формат идентификатора конфигурации в пути
404: Not Found - Рабочая конфигурация с указанным ID не найдена
500: Internal Server Error - Системная ошибка СУБД при извлечении рабочей конфигурации

---

## [PUT] /api/v1/catalog/configuration/{id} - Обновить рабочую конфигурацию по ID

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

400: Bad Request - Ошибка валидации query-параметров или неверный формат ID в пути
404: Not Found - Рабочая конфигурация с указанным ID не обнаружена
500: Internal Server Error - Внутренняя ошибка СУБД при обновлении связей рабочей конфигурации

---

## [DELETE] /api/v1/catalog/configuration/{id} - Удалить рабочую конфигурацию по ID

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

400: Bad Request - Неверный числовой формат идентификатора рабочей конфигурации
404: Not Found - Рабочая конфигурация с указанным ID не найдена
500: Internal Server Error - Ошибка целостности СУБД при удалении записи рабочей конфигурации

---

## [GET] /api/v1/catalog/configurations - Получить все рабочие конфигурации

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

500: Internal Server Error - Критическая ошибка сервера при чтении списка всех рабочих конфигураций

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
# 2. Модельный каталог: Параметры

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/param - Создать параметр

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

400: Bad Request - Ошибка валидации структуры системного параметра
500: Internal Server Error - Ошибка записи нового параметра в СУБД

---

## [GET] /api/v1/catalog/param/{id} - Получить параметр по ID

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
404: Not Found - Системный параметр с указанным ID не найден
500: Internal Server Error - Внутренняя ошибка выполнения запроса в БД

## [PUT] /api/v1/catalog/param/{id} - Обновить параметр

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

400: Bad Request - Ошибка валидации или неверный числовой формат ID в пути
404: Not Found - Обновляемый системный параметр не найден в системе
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении параметров

---

## [DELETE] /api/v1/catalog/param/{id} - Удалить параметр

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

400: Bad Request - Неверный формат ID параметра в пути
404: Not Found - Удаляемый параметр не найден в базе данных
500: Internal Server Error - Ошибка целостности СУБД при удалении параметра

---

## [PATCH] /api/v1/catalog/param/{prevId}/{newId} - Изменить ID параметра

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

400: Bad Request - Неверный формат одного из идентификаторов (prevId или newId)
404: Not Found - Системный параметр с исходным ID не существует
500: Internal Server Error - Ошибка выполнения транзакции переименования ID в БД

---

## [GET] /api/v1/catalog/param/search - Поиск параметров по строке

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

400: Bad Request - Пустой поисковый запрос (параметр query отсутствует)
500: Internal Server Error - Внутренняя ошибка СУБД при обработке текстового запроса

---

## [GET] /api/v1/catalog/param/search/{id} - Получить компоненты по ID параметра

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

400: Bad Request - Неверный формат числового идентификатора параметра
500: Internal Server Error - Ошибка выборки компонентов, владеющих данным параметром

---

## [GET] /api/v1/catalog/param/unattached - Получить непривязанные параметры

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

500: Internal Server Error - Системная ошибка СУБД при получении изолированных параметров

---

## [GET] /api/v1/catalog/params - Получить все параметры

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

500: Internal Server Error - Критическая ошибка сервера при чтении списка параметров

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
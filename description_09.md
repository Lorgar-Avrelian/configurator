# 9. Конфигурация: Связь компонентов устройства и сопоставлений параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/link/device-component-mapping/{deviceComponentId}/{mappingId} - Связать составную часть устройства с сопоставлением параметра

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

400: Bad Request - Некорректные форматы переданных идентификаторов в пути запроса
404: Not Found - Составная часть устройства или сопоставление параметра не найдены
500: Internal Server Error - Ошибка СУБД при записи связи в ассоциативную таблицу

---

## [DELETE] /api/v1/catalog/link/device-component-mapping/{deviceComponentId}/{mappingId} - Удалить связь составной части устройства с сопоставлением параметра

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

400: Bad Request - Некорректный числовой формат переданных идентификаторов в пути
404: Not Found - Запись о связи составной части с сопоставлением не обнаружена
500: Internal Server Error - Системная ошибка базы данных при удалении связи

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
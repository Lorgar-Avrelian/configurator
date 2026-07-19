# 15. Просмотр: Значения параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!TIP]
> API из данного блока предназначены для просмотра карты сохранённых значений параметров устройств, необходимых для 
> расчёта результатов пороговых выражений в других устройствах.

---

## [POST] /api/v1/catalog/param-result - Получить отфильтрованные сохранённые значения параметров

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/param-result

{
  "component_title": "component",
  "host": "127.0.0.1",
  "internal_order": 1,
  "param_title": "name",
  "port": 161
}
```

Ответ 1:

```json
[
  {
    "component_title": "component",
    "host": "127.0.0.1",
    "internal_order": 1,
    "param_title": "name",
    "port": 161,
    "value": "Main Node"
  }
]
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации структуры параметров фильтрации или некорректный формат полей  
500: Internal Server Error - Ошибка базы данных при выполнении условной выборки результатов

---

## [GET] /api/v1/catalog/param-results - Получить все сохранённые значения параметров

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/param-results
```

Ответ 1:

```json
[
  {
    "component_title": "component",
    "host": "127.0.0.1",
    "internal_order": 1,
    "param_title": "name",
    "port": 161,
    "value": "Main Node"
  }
]
```

</details>

</details>

### Возможные коды ошибок

500: Internal Server Error - Критическая ошибка сервера при чтении полной таблицы результатов

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
# 13. Просмотр: В процессе конфигурирования

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!TIP]
> API из данного блока предназначены для поиска и просмотра списка устройств, находящихся в процессе конфигурирования,
> то есть находящихся на первичном опросе, а потому ещё не имеющих конфигурации опроса.

---

## [GET] /api/v1/catalog/config/in-progress - Получить все данные устройств, находящихся в процессе конфигурирования

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/config/in-progress

{}
```

Ответ 1:

```json
[
  {
    "host": "127.0.0.1",
    "port": 161,
    "protocol": "SNMP"
  }
]
```

</details>

</details>

### Возможные коды ошибок

404: Not Found - Устройства в процессе конфигурирования отсутствуют  
500: Internal Server Error - Системная ошибка сервера при получении списка устройств

---

## [POST] /api/v1/catalog/config/in-progress/search - Поиск данных устройств, находящихся в процессе конфигурирования

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/config/in-progress/search?host=127.0.0.1&port=161

{}
```

Ответ 1:

```json
[
  {
    "host": "127.0.0.1",
    "port": 161,
    "protocol": "SNMP"
  }
]
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный формат параметров фильтрации (host или port)  
404: Not Found - Искомое устройство в процессе конфигурации не найдено  
500: Internal Server Error - Ошибка СУБД при выполнении поиска записи

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
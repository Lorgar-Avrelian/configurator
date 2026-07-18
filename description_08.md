# 8. Конфигурация: Структура компонентов устройства

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!TIP]
> Структура устройства состоит из компонентов модельного (составных частей) устройства и связей, определяющих их
> взаимную вложенность.  
> API из этого блока предназначены для сборки из компонентов модельного каталога структуры, соответствующей реальному
> составу устройства.

---

## [POST] /api/v1/catalog/device-component - Создать составную часть устройства

> [!TIP]
> Поле `internal_order` внутри структуры является не обязательным и служит для уточнения порядкового номера составной
> части устройства (компонента) внутри родительской составной части или родительского узла группировки составных частей
> (компонентов) для случаев, когда составные части, лежащие на одном уровне в Схеме деления устройства имеют различный
> состав или обладают определёнными функциональными особенностями, требующими отдельного уточнения.  
> Для базовой составной части (компонента, соответствующего устройству в целом) поле `parent` должно быть равно `null`
> или не передано вовсе.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component

{
  "model": 2
}
```

Ответ 1:

```json
{
  "id": 8,
  "component": {
    "id": 2,
    "title": "physical",
    "name_en": "Physical",
    "name_ru": "Физический"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": []
}
```

</details>

<details><summary>Пример 2</summary>

Запрос 2:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component

{
  "model": 4,
  "parent": 8
}
```

Ответ 2:

```json
{
  "id": 8,
  "component": {
    "id": 2,
    "title": "physical",
    "name_en": "Physical",
    "name_ru": "Физический"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": [
    {
      "id": 9,
      "component": {
        "id": 4,
        "title": "interface",
        "name_en": "Interface",
        "name_ru": "Интерфейс"
      },
      "internal_order": null,
      "parent": 8,
      "mappings": [],
      "children": []
    }
  ]
}
```

</details>

<details><summary>Пример 3</summary>

Запрос 3:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component

{
  "model": 7,
  "internal_order": 2,
  "parent": 8
}
```

Ответ 3:

```json
{
  "id": 8,
  "component": {
    "id": 2,
    "title": "physical",
    "name_en": "Physical",
    "name_ru": "Физический"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": [
    {
      "id": 9,
      "component": {
        "id": 4,
        "title": "interface",
        "name_en": "Interface",
        "name_ru": "Интерфейс"
      },
      "internal_order": null,
      "parent": 8,
      "mappings": [],
      "children": []
    },
    {
      "id": 10,
      "component": {
        "id": 7,
        "title": "storage",
        "name_en": "Data Storage",
        "name_ru": "Хранилище данных"
      },
      "internal_order": 2,
      "parent": 8,
      "mappings": [],
      "children": []
    }
  ]
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Нарушена валидация структуры (например, отсутствует обязательный model_id)  
500: Internal Server Error - Не удалось записать составную часть устройства в базу данных

---

## [GET] /api/v1/catalog/device-component/{id} - Получить составную часть устройства по ID

> [!TIP]
> При получении дочерней составной части устройства по id в результате запроса будет возвращена вся структура устройства
> вплоть до запрошенной составной части (до уровня вложенности, соответствующего запрошенной составной части).  
> Это сделано в целях корректного отображения конфигурируемого узла пользователю.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component/8

{}
```

Ответ 1:

```json
{
  "id": 8,
  "component": {
    "id": 2,
    "title": "physical",
    "name_en": "Physical",
    "name_ru": "Физический"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": [
    {
      "id": 9,
      "component": {
        "id": 4,
        "title": "interface",
        "name_en": "Interface",
        "name_ru": "Интерфейс"
      },
      "internal_order": null,
      "parent": 8,
      "mappings": [],
      "children": []
    },
    {
      "id": 10,
      "component": {
        "id": 7,
        "title": "storage",
        "name_en": "Data Storage",
        "name_ru": "Хранилище данных"
      },
      "internal_order": 2,
      "parent": 8,
      "mappings": [],
      "children": []
    }
  ]
}
```

</details>

<details><summary>Пример 2</summary>

Запрос 2:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component/9

{}
```

Ответ 2:

```json
{
  "id": 8,
  "component": {
    "id": 2,
    "title": "physical",
    "name_en": "Physical",
    "name_ru": "Физический"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": [
    {
      "id": 9,
      "component": {
        "id": 4,
        "title": "interface",
        "name_en": "Interface",
        "name_ru": "Интерфейс"
      },
      "internal_order": null,
      "parent": 8,
      "mappings": [],
      "children": []
    },
    {
      "id": 10,
      "component": {
        "id": 7,
        "title": "storage",
        "name_en": "Data Storage",
        "name_ru": "Хранилище данных"
      },
      "internal_order": 2,
      "parent": 8,
      "mappings": [],
      "children": []
    }
  ]
}
```

</details>

<details><summary>Пример 3</summary>

Запрос 3:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component/11

{}
```

Ответ 3:

```json
{
  "id": 11,
  "component": {
    "id": 12,
    "title": "fdd",
    "name_en": "Floppy Disk",
    "name_ru": "Гибкий диск"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": []
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Идентификатор составной части в пути имеет неверный формат  
404: Not Found - Составная часть устройства с указанным ID не найдена  
500: Internal Server Error - Системная ошибка СУБД при чтении составной части

---

## [PUT] /api/v1/catalog/device-component/{id} - Обновить составную часть устройства по ID

> [!TIP]
> Данный API позволяет менять положение и тип составной части внутри структуры устройства и должен применяться для этих
> целей.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component/11

{
  "model": 12,
  "internal_order": 10,
  "parent": 9
}
```

Ответ 1:

```json
{
  "id": 8,
  "component": {
    "id": 2,
    "title": "physical",
    "name_en": "Physical",
    "name_ru": "Физический"
  },
  "internal_order": null,
  "parent": null,
  "mappings": [],
  "children": [
    {
      "id": 9,
      "component": {
        "id": 4,
        "title": "interface",
        "name_en": "Interface",
        "name_ru": "Интерфейс"
      },
      "internal_order": null,
      "parent": 8,
      "mappings": [],
      "children": [
        {
          "id": 11,
          "component": {
            "id": 12,
            "title": "fdd",
            "name_en": "Floppy Disk",
            "name_ru": "Гибкий диск"
          },
          "internal_order": 1,
          "parent": 9,
          "mappings": [],
          "children": []
        }
      ]
    },
    {
      "id": 10,
      "component": {
        "id": 7,
        "title": "storage",
        "name_en": "Data Storage",
        "name_ru": "Хранилище данных"
      },
      "internal_order": 2,
      "parent": 8,
      "mappings": [],
      "children": []
    }
  ]
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный ID в пути или ошибка валидации JSON-тела запроса  
404: Not Found - Составная часть устройства для обновления не найдена  
500: Internal Server Error - Внутренняя ошибка сервера при изменении составной части

---

## [DELETE] /api/v1/catalog/device-component/{id} - Удалить составную часть устройства по ID

> [!CAUTION]
> При удалении составной части из структуры устройства в БД также удаляются все сложенные в неё составные части.  
> При удалении базовой составной части (компонента, описывающего устройство в целом) происходит полное удаление
> структуры устройства из БД, что приводит также к удалению всех конфигураций опроса, опирающихся на удаляемую
> структуру.

> [!WARNING]
> В связи с вышесказанным желательно добавить либо кнопку, либо дополнительное всплывающее сообщение подтверждения
> проведения операции.

> [!IMPORTANT]
> При удалении составной части происходит сдвиг (уменьшение значения id на 1 и более) всего списка составных частей,
> следовавших за удалёнными компонентами, что не приводит к сбоям в работе системы.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component/11

{}
```

Ответ 1:

```json
{}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат идентификатора составной части в пути  
404: Not Found - Составная часть устройства с указанным ID не существует  
500: Internal Server Error - Системная ошибка каскадного удаления на стороне БД

---

## [GET] /api/v1/catalog/device-component/{id}/own - Получить изолированную составную часть устройства по ID

> [!TIP]
> Данный API предназначен для получения дочерней составной части без привязки к структуре устройства.  
> Удобен для целей редактирования отдельного узла.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/device-component/6/own

{}
```

Ответ 1:

```json
{
  "id": 6,
  "component": {
    "id": 7,
    "title": "storage",
    "name_en": "Data Storage",
    "name_ru": "Хранилище данных"
  },
  "internal_order": 1,
  "parent": 1,
  "mappings": [],
  "children": []
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный числовой формат идентификатора составной части  
404: Not Found - Составная часть устройства с данным ID не найдена  
500: Internal Server Error - Внутренняя ошибка сервера при чтении изолированной записи

---

## [PATCH] /api/v1/catalog/device-component/{prevId}/{newId} - Изменить ID составной части устройства

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

400: Bad Request - Неверный формат одного из ID (prevId или newId) в пути запроса  
404: Not Found - Составная часть устройства с исходным ID не существует  
500: Internal Server Error - Ошибка СУБД при изменении первичного или внешнего ключа

---

## [GET] /api/v1/catalog/device-components - Получить всю структуру составных частей устройств

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

500: Internal Server Error - Критическая ошибка сервера при формировании дерева составных частей

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
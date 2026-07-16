# 2. Модельный каталог: Параметры

> [!CAUTION]
> **_!!! ВСЕ ПЕРЕЧИСЛЕННЫЕ НИЖЕ API ДОСТУПНЫ ТОЛЬКО ДЛЯ ПРОФИЛЯ "dev" ПРИЛОЖЕНИЯ !!!_**

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/param - Создать параметр

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/param

{
  "title": "id",
  "name_en": "ID",
  "name_ru": "ID",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component identifier in the component group",
  "description_ru": "Идентификатор компонента в группе компонентов",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
```

Ответ 1:

```json
{
  "id": 3,
  "title": "id",
  "name_en": "ID",
  "name_ru": "ID",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component identifier in the component group",
  "description_ru": "Идентификатор компонента в группе компонентов",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
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
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/param/4

{}
```

Ответ 1:

```json
{
  "id": 4,
  "title": "name",
  "name_en": "Name",
  "name_ru": "Имя",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component name",
  "description_ru": "Название компонента",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
```

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат идентификатора в пути
404: Not Found - Системный параметр с указанным ID не найден
500: Internal Server Error - Внутренняя ошибка выполнения запроса в БД

## [PUT] /api/v1/catalog/param/{id} - Обновить параметр

> [!TIP]
> Поле "id", содержащееся в теле запроса, при обновлении данных параметра полностью игнорируется.  
> Запрос обновляет данные только самого параметра, имеющего id, равный переданному значению в пути запроса.  
> Если параметр был привязан к каким-либо компонентам, то привязки не исчезают.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/param/4

{
  "title": "not_name",
  "name_en": "NOT Name",
  "name_ru": "НЕ Имя",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component name",
  "description_ru": "Название компонента",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
```

Ответ 1:

```json
{
  "id": 4,
  "title": "not_name",
  "name_en": "NOT Name",
  "name_ru": "НЕ Имя",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component name",
  "description_ru": "Название компонента",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
```

Запрос 2:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/param/4

{
  "id": 4,
  "title": "name",
  "name_en": "Name",
  "name_ru": "Имя",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component name",
  "description_ru": "Название компонента",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
```

Ответ 2:

```json
{
  "id": 4,
  "title": "name",
  "name_en": "Name",
  "name_ru": "Имя",
  "type": "VARCHAR",
  "value": null,
  "description_en": "Component name",
  "description_ru": "Название компонента",
  "access": "USER",
  "saved": true,
  "visible": true,
  "diagram": false
}
```

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации или неверный числовой формат ID в пути
404: Not Found - Обновляемый системный параметр не найден в системе
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении параметров

---

## [DELETE] /api/v1/catalog/param/{id} - Удалить параметр

> [!CAUTION]
> При удалении параметра **_он удаляется из всех компонентов, к которым был привязан_**.

> [!WARNING]  
> В связи с вышесказанным желательно добавить либо кнопку, либо дополнительное всплывающее сообщение подтверждения
> проведения операции.

> [!IMPORTANT]  
> При удалении параметра происходит сдвиг (уменьшение значения id на 1) всего списка параметров, следовавших за
> удалённым, что не приводит к сбоям в работе системы.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/param/4

{}
```

Ответ 1:

```json
{}
```

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат ID параметра в пути
404: Not Found - Удаляемый параметр не найден в базе данных
500: Internal Server Error - Ошибка целостности СУБД при удалении параметра

---

## [PATCH] /api/v1/catalog/param/{prevId}/{newId} - Изменить ID параметра

> [!TIP]
> API предназначен для упорядочивания перечня параметров и не имеет никаких ограничений по новому значению id, кроме
> правила: id должен быть больше 0.  
> При изменении id параметра возможен сдвиг всех последующих параметров в списке параметров (увеличение их id на 1),
> который не приводит к возникновению ошибки в работе системы.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
PATCH https://nms-dev.opk-bulat.ru/api/v1/catalog/param/62/1

{}
```

Ответ 1:

```json
{}
```

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат одного из идентификаторов (prevId или newId)
404: Not Found - Системный параметр с исходным ID не существует
500: Internal Server Error - Ошибка выполнения транзакции переименования ID в БД

---

## [GET] /api/v1/catalog/param/search - Поиск параметров по строке

> [!TIP]
> API для поиска параметра по строке, содержащей часть или полное значение полей "title", "name_en", "name_ru",
> "description_en", "description_ru".

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/param/search?query=status

{}
```

Ответ 1:

```json
[
  {
    "id": 12,
    "title": "admin_status",
    "name_en": "Admin status",
    "name_ru": "Заданный режим работы",
    "type": "VARCHAR",
    "value": null,
    "description_en": "User-defined (or System-defined) operating mode of the component",
    "description_ru": "Заданный пользователем (или системой) режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 13,
    "title": "oper_status",
    "name_en": "Oper status",
    "name_ru": "Текущий режим работы",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Current operating mode of the component",
    "description_ru": "Текущий режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 16,
    "title": "admin_status_flag",
    "name_en": "Admin status",
    "name_ru": "Заданный режим работы",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "User-defined (or System-defined) operating mode of the component",
    "description_ru": "Заданный пользователем (или системой) режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 17,
    "title": "oper_status_flag",
    "name_en": "Oper status",
    "name_ru": "Текущий режим работы",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "Current operating mode of the component",
    "description_ru": "Текущий режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 33,
    "title": "status",
    "name_en": "Status",
    "name_ru": "Статус",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Status of the component",
    "description_ru": "Статус компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  }
]
```

Запрос 2 - Пример для ввода значения `количество` на кириллице:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/param/search?query=%D0%BA%D0%BE%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D1%82%D0%B2%D0%BE

{}
```

Ответ 1:

```json
[
  {
    "id": 46,
    "title": "count",
    "name_en": "Count",
    "name_ru": "Количество",
    "type": "INTEGER",
    "value": null,
    "description_en": "Number of components of this type in the parent component",
    "description_ru": "Количество компонентов данного типа в составе родительского компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 50,
    "title": "input_bytes_total",
    "name_en": "Input bytes (total)",
    "name_ru": "Принято байт (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of bytes received",
    "description_ru": "Общее количество полученных байт",
    "units_en": "Byte",
    "units_ru": "Байт",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 51,
    "title": "input_unicast_total",
    "name_en": "Input Unicast packets (total)",
    "name_ru": "Принято Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of subnetwork-unicast packets delivered to a higher-layer protocol",
    "description_ru": "Количество Unicast-пакетов в подсети, доставленных по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 52,
    "title": "input_nucast_total",
    "name_en": "Input Non-Unicast packets (total)",
    "name_ru": "Принято Non-Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of Non-Unicast packets delivered to a higher-layer protocol",
    "description_ru": "Количество Non-Unicast-пакетов в подсети, доставленных по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 53,
    "title": "input_discards_total",
    "name_en": "Input packets discards (total)",
    "name_ru": "Сброшено входящих пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of input packets discards",
    "description_ru": "Общее количество отброшенных входных пакетов",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 54,
    "title": "input_pkts_errors_total",
    "name_en": "Inbound errors (total)",
    "name_ru": "Входящие ошибки (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of inbound packets that contained errors preventing them from being deliverable to a higher-layer protocol",
    "description_ru": "Количество входящих пакетов, содержащих ошибки, препятствующие их передаче по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 55,
    "title": "input_pkts_unknown_total",
    "name_en": "Packets with unknown protocol (total)",
    "name_ru": "Пакеты с неизвестным протоколом (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of packets received via the interface which were discarded because of an unknown or unsupported protocol",
    "description_ru": "Количество входящих пакетов, содержащих ошибки, препятствующие их передаче по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 56,
    "title": "output_bytes_total",
    "name_en": "Output bytes (total)",
    "name_ru": "Передано байт (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of bytes transmitted out",
    "description_ru": "Общее количество переданных байт",
    "units_en": "Byte",
    "units_ru": "Байт",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 57,
    "title": "output_unicast_total",
    "name_en": "Output Unicast packets (total)",
    "name_ru": "Передано Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of packets that higher-level protocols requested be transmitted to a subnetwork-unicast address",
    "description_ru": "Общее количество пакетов, запрошенных протоколами более высокого уровня для передачи на Unicast-адрес в подсети",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 58,
    "title": "output_nucast_total",
    "name_en": "Output Non-Unicast packets (total)",
    "name_ru": "Передано Non-Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of packets that higher-level protocols requested be transmitted to a non-unicast address",
    "description_ru": "Общее количество пакетов, которые запрошены протоколами более высокого уровня для передачи на Non-Unicast-адрес",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 59,
    "title": "output_discards_total",
    "name_en": "Output packets discards (total)",
    "name_ru": "Сброшено исходящих пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of outbound packets which were chosen to be discarded",
    "description_ru": "Количество исходящих пакетов, которые были выбраны для сброса",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 60,
    "title": "output_pkts_errors_total",
    "name_en": "Outbound errors (total)",
    "name_ru": "Исходящие ошибки (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of outbound packets that could not be transmitted because of errors",
    "description_ru": "Количество исходящих пакетов, которые не удалось передать из-за ошибок",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  }
]
```

</details>

### Возможные коды ошибок

400: Bad Request - Пустой поисковый запрос (параметр query отсутствует)
500: Internal Server Error - Внутренняя ошибка СУБД при обработке текстового запроса

---

## [GET] /api/v1/catalog/param/search/{id} - Получить компоненты по ID параметра

> [!NOTE]  
> Если наименование параметра и его единицы измерения совпадают для нескольких компонентов, относящихся к различным
> компонентам-родителям, то данный параметр может быть прикреплён к каждому из них.  
> Например, параметр `rotation_speed` ('скорость вращения') прикреплён как к компоненту `cd_rom` ('CD-ROM'), так и к
> компоненту `flywheel` ('Маховик').

> [!TIP]
> В ответе на запрос отображаются только те компоненты, к которым непосредственно прикреплён параметр. То есть
> **_компоненты-наследники_**, также наследующие данный параметр, в ответе на запрос **_отражены не будут_**.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/param/search/4

{}
```

Ответ 1:

```json
[
  {
    "id": 1,
    "title": "component",
    "name_en": "Component",
    "name_ru": "Компонент",
    "plural_name_en": "Components",
    "plural_name_ru": "Компоненты",
    "base_component": null,
    "description_en": "Base component",
    "description_ru": "Базовый компонент",
    "access": "USER",
    "params": []
  }
]
```

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат числового идентификатора параметра
500: Internal Server Error - Ошибка выборки компонентов, владеющих данным параметром

---

## [GET] /api/v1/catalog/param/unattached - Получить непривязанные параметры

> [!TIP]
> API предназначен для получения списка параметров, не привязанных ни к одному компоненту, и выполняет справочную
> функцию на случай массового создания новых параметров в модельном каталоге и их последующего распределения по
> компонентам.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/param/unattached

{}
```

Ответ 1:

```json
[
  {
    "id": 2,
    "title": "name",
    "name_en": "Test",
    "name_ru": "Test",
    "description_en": "Test",
    "description_ru": "Test",
    "saved": true,
    "type": "VARCHAR",
    "units_en": "test",
    "units_ru": "test",
    "value": null,
    "access": "USER",
    "visible": true,
    "diagram": false
  }
]
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
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/params

{}
```

Ответ 1:

```json
[
  {
    "id": 1,
    "title": "component",
    "name_en": "Component ID",
    "name_ru": "ID компонента",
    "type": "BIGINT",
    "value": null,
    "description_en": "Service parameter for changing the component type according to the received value",
    "description_ru": "Служебный параметр для изменения типа компонента в соответствии с полученным значением",
    "access": "GUEST",
    "saved": false,
    "visible": false,
    "diagram": false
  },
  {
    "id": 2,
    "title": "icon",
    "name_en": "Icon",
    "name_ru": "Иконка",
    "type": "TEXT",
    "value": null,
    "description_en": "System path of the icon for displaying the component in the user interface",
    "description_ru": "Системный путь иконки для отображения компонента в интерфейсе пользователя",
    "access": "ADMIN",
    "saved": false,
    "visible": true,
    "diagram": false
  },
  {
    "id": 3,
    "title": "id",
    "name_en": "ID",
    "name_ru": "ID",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Component identifier in the component group",
    "description_ru": "Идентификатор компонента в группе компонентов",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 4,
    "title": "name",
    "name_en": "Name",
    "name_ru": "Имя",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Component name",
    "description_ru": "Название компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 5,
    "title": "tag",
    "name_en": "Tag",
    "name_ru": "Метка",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Label of the component group containing this component",
    "description_ru": "Метка группы компонентов, содержащей данный компонент",
    "access": "ADMIN",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 6,
    "title": "uuid",
    "name_en": "UUID",
    "name_ru": "UUID",
    "type": "UUID",
    "value": null,
    "description_en": "Universal Unique Identifier",
    "description_ru": "Универсальный уникальный идентификатор",
    "access": "ADMIN",
    "saved": true,
    "visible": false,
    "diagram": false
  },
  {
    "id": 7,
    "title": "base_uuid",
    "name_en": "Base UUID",
    "name_ru": "Базовый UUID",
    "type": "UUID",
    "value": null,
    "description_en": "Base component Universal Unique Identifier",
    "description_ru": "Универсальный уникальный идентификатор базового компонента",
    "access": "ADMIN",
    "saved": true,
    "visible": false,
    "diagram": false
  },
  {
    "id": 8,
    "title": "ip",
    "name_en": "IP",
    "name_ru": "IP",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Component IP Address",
    "description_ru": "IP-адрес компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 9,
    "title": "description",
    "name_en": "Description",
    "name_ru": "Описание",
    "type": "TEXT",
    "value": null,
    "description_en": "Component description",
    "description_ru": "Описание компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 10,
    "title": "condition",
    "name_en": "Condition",
    "name_ru": "Состояние",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Component condition",
    "description_ru": "Состояние компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 11,
    "title": "connection_state",
    "name_en": "Connected",
    "name_ru": "Подключено",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Connection state of the component",
    "description_ru": "Состояние подключения компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 12,
    "title": "admin_status",
    "name_en": "Admin status",
    "name_ru": "Заданный режим работы",
    "type": "VARCHAR",
    "value": null,
    "description_en": "User-defined (or System-defined) operating mode of the component",
    "description_ru": "Заданный пользователем (или системой) режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 13,
    "title": "oper_status",
    "name_en": "Oper status",
    "name_ru": "Текущий режим работы",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Current operating mode of the component",
    "description_ru": "Текущий режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 14,
    "title": "creation_time",
    "name_en": "Creation time",
    "name_ru": "Время создания",
    "type": "TIMESTAMPTZ",
    "value": null,
    "description_en": "Time of the creation component in the system",
    "description_ru": "Время создания компонента в системе",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 15,
    "title": "activation_time",
    "name_en": "Activation time",
    "name_ru": "Время активации",
    "type": "TIMESTAMPTZ",
    "value": null,
    "description_en": "Time when the component was submitted for the polling",
    "description_ru": "Время постановки компонента на опрос",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 16,
    "title": "admin_status_flag",
    "name_en": "Admin status",
    "name_ru": "Заданный режим работы",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "User-defined (or System-defined) operating mode of the component",
    "description_ru": "Заданный пользователем (или системой) режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 17,
    "title": "oper_status_flag",
    "name_en": "Oper status",
    "name_ru": "Текущий режим работы",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "Current operating mode of the component",
    "description_ru": "Текущий режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 18,
    "title": "connection_state_flag",
    "name_en": "Connected",
    "name_ru": "Подключено",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "Connection state of the component",
    "description_ru": "Состояние подключения компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 19,
    "title": "condition_flag",
    "name_en": "Condition",
    "name_ru": "Состояние",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "Component condition",
    "description_ru": "Состояние компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 20,
    "title": "active",
    "name_en": "Active",
    "name_ru": "Активно",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "Component Polling indicator",
    "description_ru": "Индикатор опроса компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 21,
    "title": "primary",
    "name_en": "Primary",
    "name_ru": "Основной",
    "type": "BOOLEAN",
    "value": null,
    "description_en": "Indicator that the component is the primary",
    "description_ru": "Индикатор того, что компонент является основным",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 22,
    "title": "last_updated",
    "name_en": "Last updated",
    "name_ru": "Последнее обновление",
    "type": "TIMESTAMPTZ",
    "value": null,
    "description_en": "Last updated time",
    "description_ru": "Время последнего обновления",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 23,
    "title": "last_alarm_time",
    "name_en": "Last alarm",
    "name_ru": "Последняя тревога",
    "type": "TIMESTAMPTZ",
    "value": null,
    "description_en": "Time of the last alarm",
    "description_ru": "Время последнего сигнала тревоги",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 24,
    "title": "last_alarm_level",
    "name_en": "Last alarm level",
    "name_ru": "Уровень последней тревоги",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Level of the last alarm",
    "description_ru": "Уровень последнего сигнала тревоги",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 25,
    "title": "contact",
    "name_en": "Contact",
    "name_ru": "Контакт",
    "type": "TEXT",
    "value": null,
    "description_en": "Contact information",
    "description_ru": "Контактные данные",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 26,
    "title": "address",
    "name_en": "Address",
    "name_ru": "Адрес",
    "type": "TEXT",
    "value": null,
    "description_en": "Address of the component location",
    "description_ru": "Адрес расположения компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 27,
    "title": "vendor",
    "name_en": "Vendor",
    "name_ru": "Производитель",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Vendor of the component",
    "description_ru": "Производитель компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 28,
    "title": "protocol",
    "name_en": "Protocol",
    "name_ru": "Протокол",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Polling protocol",
    "description_ru": "Протокол опроса",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 29,
    "title": "uptime",
    "name_en": "Uptime",
    "name_ru": "Время работы",
    "type": "TIME",
    "value": null,
    "description_en": "Uptime",
    "description_ru": "Время работы",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 30,
    "title": "system_time",
    "name_en": "System time",
    "name_ru": "Системное время",
    "type": "TIME",
    "value": null,
    "description_en": "System time",
    "description_ru": "Системное время",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 31,
    "title": "os",
    "name_en": "OS",
    "name_ru": "ОС",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Operating system",
    "description_ru": "Операционная система",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 32,
    "title": "mac",
    "name_en": "MAC",
    "name_ru": "MAC",
    "type": "VARCHAR",
    "value": null,
    "description_en": "MAC address",
    "description_ru": "MAC-адрес",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 33,
    "title": "status",
    "name_en": "Status",
    "name_ru": "Статус",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Status of the component",
    "description_ru": "Статус компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 34,
    "title": "type",
    "name_en": "Type",
    "name_ru": "Тип",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Type of the component",
    "description_ru": "Тип компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 35,
    "title": "mode",
    "name_en": "Mode",
    "name_ru": "Режим работы",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Mode of the component",
    "description_ru": "Режим работы компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 36,
    "title": "location",
    "name_en": "Location",
    "name_ru": "Местоположение",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Location of the component",
    "description_ru": "Местоположение компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 37,
    "title": "slug",
    "name_en": "Slug",
    "name_ru": "Слаг",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Unique alphanumeric identifier that is understandable to humans",
    "description_ru": "Уникальный и понятный человеку буквенно-цифровой идентификатор",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 38,
    "title": "depth",
    "name_en": "Depth",
    "name_ru": "Глубина",
    "type": "NUMERIC",
    "value": null,
    "description_en": "Component depth",
    "description_ru": "Глубина компонента",
    "units_en": "m",
    "units_ru": "м",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 39,
    "title": "width",
    "name_en": "Width",
    "name_ru": "Ширина",
    "type": "NUMERIC",
    "value": null,
    "description_en": "Component width",
    "description_ru": "Ширина компонента",
    "units_en": "m",
    "units_ru": "м",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 40,
    "title": "height",
    "name_en": "Height",
    "name_ru": "Высота",
    "type": "NUMERIC",
    "value": null,
    "description_en": "Component height",
    "description_ru": "Высота компонента",
    "units_en": "m",
    "units_ru": "м",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 41,
    "title": "role",
    "name_en": "Role",
    "name_ru": "Роль",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Component role",
    "description_ru": "Роль компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 42,
    "title": "front_image",
    "name_en": "Front image",
    "name_ru": "Изображение спереди",
    "type": "TEXT",
    "value": null,
    "description_en": "System path of the front image",
    "description_ru": "Системный путь изображения спереди",
    "access": "USER",
    "saved": true,
    "visible": false,
    "diagram": false
  },
  {
    "id": 43,
    "title": "rear_image",
    "name_en": "Rear image",
    "name_ru": "Изображение сзади",
    "type": "TEXT",
    "value": null,
    "description_en": "System path of the rear image",
    "description_ru": "Системный путь изображения сзади",
    "access": "USER",
    "saved": true,
    "visible": false,
    "diagram": false
  },
  {
    "id": 44,
    "title": "weight",
    "name_en": "Weight",
    "name_ru": "Масса",
    "type": "VARCHAR",
    "value": null,
    "description_en": "Component weight",
    "description_ru": "Масса компонента",
    "units_en": "g",
    "units_ru": "г",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 45,
    "title": "osi",
    "name_en": "OSI",
    "name_ru": "OSI",
    "type": "VARCHAR",
    "value": null,
    "description_en": "OSI services list",
    "description_ru": "Перечень сервисов OSI",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 46,
    "title": "count",
    "name_en": "Count",
    "name_ru": "Количество",
    "type": "INTEGER",
    "value": null,
    "description_en": "Number of components of this type in the parent component",
    "description_ru": "Количество компонентов данного типа в составе родительского компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 47,
    "title": "mtu",
    "name_en": "MTU",
    "name_ru": "MTU",
    "type": "INTEGER",
    "value": null,
    "description_en": "Maximum Transmission Unit",
    "description_ru": "Максимальный размер кадра",
    "units_en": "Byte",
    "units_ru": "Байт",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 48,
    "title": "bandwidth",
    "name_en": "Bandwidth",
    "name_ru": "Пропускная способность",
    "type": "INTEGER",
    "value": null,
    "description_en": "Current bandwidth",
    "description_ru": "Текущая пропускная способность",
    "units_en": "bps",
    "units_ru": "бит/с",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 49,
    "title": "last_changed",
    "name_en": "Last changed",
    "name_ru": "Последнее изменение",
    "type": "TIME",
    "value": null,
    "description_en": "Time the component entered its current operational state from the last component initialization",
    "description_ru": "Время перехода компонента в текущее рабочее состояние с момента последней инициализации компонента",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": false
  },
  {
    "id": 50,
    "title": "input_bytes_total",
    "name_en": "Input bytes (total)",
    "name_ru": "Принято байт (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of bytes received",
    "description_ru": "Общее количество полученных байт",
    "units_en": "Byte",
    "units_ru": "Байт",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 51,
    "title": "input_unicast_total",
    "name_en": "Input Unicast packets (total)",
    "name_ru": "Принято Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of subnetwork-unicast packets delivered to a higher-layer protocol",
    "description_ru": "Количество Unicast-пакетов в подсети, доставленных по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 52,
    "title": "input_nucast_total",
    "name_en": "Input Non-Unicast packets (total)",
    "name_ru": "Принято Non-Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of Non-Unicast packets delivered to a higher-layer protocol",
    "description_ru": "Количество Non-Unicast-пакетов в подсети, доставленных по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 53,
    "title": "input_discards_total",
    "name_en": "Input packets discards (total)",
    "name_ru": "Сброшено входящих пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of input packets discards",
    "description_ru": "Общее количество отброшенных входных пакетов",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 54,
    "title": "input_pkts_errors_total",
    "name_en": "Inbound errors (total)",
    "name_ru": "Входящие ошибки (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of inbound packets that contained errors preventing them from being deliverable to a higher-layer protocol",
    "description_ru": "Количество входящих пакетов, содержащих ошибки, препятствующие их передаче по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 55,
    "title": "input_pkts_unknown_total",
    "name_en": "Packets with unknown protocol (total)",
    "name_ru": "Пакеты с неизвестным протоколом (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of packets received via the interface which were discarded because of an unknown or unsupported protocol",
    "description_ru": "Количество входящих пакетов, содержащих ошибки, препятствующие их передаче по протоколу более высокого уровня",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 56,
    "title": "output_bytes_total",
    "name_en": "Output bytes (total)",
    "name_ru": "Передано байт (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of bytes transmitted out",
    "description_ru": "Общее количество переданных байт",
    "units_en": "Byte",
    "units_ru": "Байт",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 57,
    "title": "output_unicast_total",
    "name_en": "Output Unicast packets (total)",
    "name_ru": "Передано Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of packets that higher-level protocols requested be transmitted to a subnetwork-unicast address",
    "description_ru": "Общее количество пакетов, запрошенных протоколами более высокого уровня для передачи на Unicast-адрес в подсети",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 58,
    "title": "output_nucast_total",
    "name_en": "Output Non-Unicast packets (total)",
    "name_ru": "Передано Non-Unicast-пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Total number of packets that higher-level protocols requested be transmitted to a non-unicast address",
    "description_ru": "Общее количество пакетов, которые запрошены протоколами более высокого уровня для передачи на Non-Unicast-адрес",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 59,
    "title": "output_discards_total",
    "name_en": "Output packets discards (total)",
    "name_ru": "Сброшено исходящих пакетов (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of outbound packets which were chosen to be discarded",
    "description_ru": "Количество исходящих пакетов, которые были выбраны для сброса",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 60,
    "title": "output_pkts_errors_total",
    "name_en": "Outbound errors (total)",
    "name_ru": "Исходящие ошибки (всего)",
    "type": "BIGINT",
    "value": null,
    "description_en": "Number of outbound packets that could not be transmitted because of errors",
    "description_ru": "Количество исходящих пакетов, которые не удалось передать из-за ошибок",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 61,
    "title": "output_queue_length",
    "name_en": "Output queue length",
    "name_ru": "Длина исходящей очереди",
    "type": "INTEGER",
    "value": null,
    "description_en": "Length of the output packet queue",
    "description_ru": "Длина очереди исходящих пакетов",
    "units_en": "pkts",
    "units_ru": "пк",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 62,
    "title": "load_percent",
    "name_en": "Load",
    "name_ru": "Нагрузка",
    "type": "INTEGER",
    "value": null,
    "description_en": "Load",
    "description_ru": "Нагрузка",
    "units_en": "%",
    "units_ru": "%",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 63,
    "title": "cluster_size",
    "name_en": "Cluster size",
    "name_ru": "Размер кластера",
    "type": "INTEGER",
    "value": null,
    "description_en": "Physical size of one allocation unit (cluster/block)",
    "description_ru": "Физический размер одной единицы выделения (кластера/блока)",
    "units_en": "Byte",
    "units_ru": "Байт",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 64,
    "title": "storage_size",
    "name_en": "Storage size",
    "name_ru": "Размер хранилища",
    "type": "INTEGER",
    "value": null,
    "description_en": "Total amount of storage, expressed in the number of memory allocation units",
    "description_ru": "Общий объём хранилища, выраженный в количестве единиц выделения памяти",
    "units_en": "Allocation unit",
    "units_ru": "Единиц выделения",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  },
  {
    "id": 65,
    "title": "storage_used",
    "name_en": "Storage used",
    "name_ru": "Занято хранилища",
    "type": "INTEGER",
    "value": null,
    "description_en": "Current amount of occupied storage memory, expressed in the number of allocation units",
    "description_ru": "Текущий объём занятой памяти хранилища, выраженный в количестве единиц выделения",
    "units_en": "Allocation unit",
    "units_ru": "Единиц выделения",
    "access": "USER",
    "saved": true,
    "visible": true,
    "diagram": true
  }
]
```

</details>

### Возможные коды ошибок

500: Internal Server Error - Критическая ошибка сервера при чтении списка параметров

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
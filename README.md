# ОПИСАНИЕ API

<details><summary>1. Модельный каталог: Компоненты</summary>

# 1. Модельный каталог: Компоненты

---

## [POST] /api/v1/catalog/component - Создать компонент устройства

> [!TIP]
> В поле "base_component" указывается значение id родительского компонента. После создания все параметры наследуются от
> родительского компонента.

> [!CAUTION]
> Поле "base_component" является обязательным в создаваемом Конфигураторе, хотя и не является обязательным для этого
> API.  
> Это вызвано тем, что **_базовый компонент_** для всех остальных компонентов должен быть только один, и он
> **_уже создан_**.  
> Компонент с id, равным 1, является основой для создания всех остальных компонентов модельного каталога.
<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/api/v1/catalog/component

{
  "title": "phisical",
  "base_component": 1,
  "name_en": "Physical",
  "name_ru": "Физический",
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "access": "USER"
}
```

Ответ 1:

```json
{
  "id": 2,
  "title": "physical",
  "name_en": "Physical",
  "name_ru": "Физический",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "base_component": 1,
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "access": "USER",
  "params": [
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
    }
  ]
}
```
</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат тела запроса или ошибка валидации данных
500: Internal Server Error - Ошибка на стороне базы данных или сервера

---

## [GET] /api/v1/catalog/component/{id} - Получить компонент по ID
<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/component/2

{}
```

Ответ 1:

```json
{
  "id": 2,
  "title": "physical",
  "name_en": "Physical",
  "name_ru": "Физический",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "base_component": 1,
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "access": "USER",
  "params": [
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
    }
  ]
}
```
</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат идентификатора компонента в пути
404: Not Found - Компонент с указанным ID не найден в системе
500: Internal Server Error - Внутренняя ошибка базы данных

---

## [PUT] /api/v1/catalog/component/{id} - Обновить компонент по ID

> [!TIP]
> Поля "id" и "params", содержащиеся в теле запроса, при обновлении данных компонента полностью игнорируется.  
> Запрос обновляет данные только самого компонента, имеющего id, равный переданному значению в пути запроса.  
> Параметры компонента остаются теми же, что и были прикреплены к нему или его родительскому компоненту ранее.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/component/2

{
  "id": 3,
  "title": "phisical",
  "base_component": 1,
  "name_en": "NON Physical",
  "name_ru": "НЕ Физический",
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "access": "USER",
  "params": []
}
```

Ответ 1:

```json
{
  "id": 2,
  "title": "physical",
  "name_en": "NON Physical",
  "name_ru": "НЕ Физический",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "base_component": 1,
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "access": "USER",
  "params": [
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
    }
  ]
}
```

Запрос 2:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/component/2

{
  "id": 2,
  "title": "physical",
  "name_en": "Physical",
  "name_ru": "Физический",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "base_component": 1,
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "access": "USER",
  "params": [
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
    }
  ]
}
```

Ответ 2:

```json
{
  "id": 2,
  "title": "physical",
  "name_en": "Physical",
  "name_ru": "Физический",
  "plural_name_en": "Physicals",
  "plural_name_ru": "Физические",
  "base_component": 1,
  "description_en": "Physical component - Component with dimensional characteristics",
  "description_ru": "Физический компонент - Компонент с габаритными характеристиками",
  "access": "USER",
  "params": [
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
    }
  ]
}
```
</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат ID в пути или ошибка валидации JSON-тела
500: Internal Server Error - Ошибка обновления данных на стороне СУБД

---

## [DELETE] /api/v1/catalog/component/{id} - Удалить компонент по ID

> [!CAUTION]
> При удалении компонента **_также удаляются все его компоненты-потомки_**, а также связи всех удалённых компонентов с
> параметрами.

> [!WARNING]  
> В связи с вышесказанным желательно добавить либо кнопку, либо дополнительное всплывающее сообщение подтверждения
> проведения операции.
<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/component/2

{}
```

Ответ 1:

```json
{}
```
</details>

### Возможные коды ошибок

404: Not Found - Удаляемый компонент с данным ID не найден в системе

---

## [PATCH] /api/v1/catalog/component/{prevId}/{newId} - Изменить ID компонента

> [!TIP]
> API предназначен для упорядочивания перечня компонентов и имеет одно ограничение: id компонента-потомка должно
> следовать строго после (быть больше) id компонента-родителя.  
> Таким образом, базовый компонент `component` (id которого равен 1) всегда будет первым в списке компонентов.  
> При изменении id компонента возможен сдвиг всех последующих компонентов в списке компонентов (увеличение их id на 1),
> который не приводит к возникновению ошибки в работе системы.
<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/component/3/2

{}
```

Ответ 1:

```json
{}
```
</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат prevId или newId, либо операция отклонена сервисом
404: Not Found - Исходный компонент для изменения ID не найден
500: Internal Server Error - Ошибка изменения первичного или внешнего ключа в СУБД

---

## [GET] /api/v1/catalog/component/search - Поиск компонентов по строке
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

400: Bad Request - Строка поиска query отсутствует или является пустой
500: Internal Server Error - Ошибка выполнения полнотекстового запроса к БД

---

## [GET] /api/v1/catalog/components - Получить всю структуру подчиненности устройств

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

500: Internal Server Error - Ошибка формирования дерева зависимостей компонентов на сервере

---

</details>

# 2. Модельный каталог: Параметры

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

# 3. Модельный каталог: Связи

## [POST] /api/v1/catalog/link/component-param/{componentId}/{paramId} - Связать компонент с параметром

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

400: Bad Request - Некорректные форматы переданных идентификаторов в пути
404: Not Found - Переданный компонент или параметр не обнаружены в моделях
500: Internal Server Error - Ошибка СУБД при вставке связи в component_param

---

## [DELETE] /api/v1/catalog/link/component-param/{componentId}/{paramId} - Удалить связь компонента с параметром

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

400: Bad Request - Некорректные типы переданных параметров или компонентов в URL
404: Not Found - Запись о связи или сущность каталога не найдены
500: Internal Server Error - Системная ошибка БД при выполнении каскадного удаления связи

# 4. Парсер: OID

## [GET] /api/v1/catalog/oid - Поиск OID по точной dotter notation

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

400: Bad Request - Параметр notation отсутствует или имеет некорректный формат
500: Internal Server Error - Внутренняя ошибка парсера или СУБД при поиске OID

---

## [GET] /api/v1/catalog/oid/exact - Поиск OID по dotter notation, названию MIB и производителю

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

400: Bad Request - Пропущены обязательные параметры запроса notation или mib
500: Internal Server Error - Ошибка СУБД при точном поиске записи OID

---

## [GET] /api/v1/catalog/oid/mib - Получить OID по названию MIB

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

400: Bad Request - Обязательный параметр названия MIB (name) не передан
500: Internal Server Error - Ошибка базы данных при выборке объектов MIB-файла

---

## [GET] /api/v1/catalog/oid/prefix - Поиск OID по префиксу с пагинацией

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

400: Bad Request - Отсутствует обязательный параметр префикса (prefix)
500: Internal Server Error - Ошибка пагинации или выполнения префиксного поиска в СУБД

---

## [GET] /api/v1/catalog/oid/vendor - Получить OID по производителю с пагинацией

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

500: Internal Server Error - Ошибка сервера или СУБД при получении объектов по вендору

---

# 5. Конфигурация: Индикаторы устройств

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

# 6. Конфигурация: Индикаторы параметров

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

# 7. Конфигурация: Сопоставления параметров

## [POST] /api/v1/catalog/mapping - Создать сопоставление

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

400: Bad Request - Ошибка валидации структуры или обязательных параметров сопоставления
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении сопоставления

---

## [GET] /api/v1/catalog/mapping/{id} - Получить сопоставление по ID

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

400: Bad Request - Неверный числовой формат идентификатора сопоставления в пути
404: Not Found - Сопоставление с указанным ID не найдено в системе
500: Internal Server Error - Системная ошибка базы данных при извлечении сопоставления

---

## [PUT] /api/v1/catalog/mapping/{id} - Обновить сопоставление по ID

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

400: Bad Request - Некорректный ID в пути или ошибка валидации JSON-структуры запроса
404: Not Found - Сопоставление для обновления с указанным ID не найдено
500: Internal Server Error - Ошибка обновления записи на стороне базы данных

---

## [DELETE] /api/v1/catalog/mapping/{id} - Удалить сопоставление по ID

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

400: Bad Request - Некорректный формат ID сопоставления в пути запроса
404: Not Found - Удаляемое сопоставление не обнаружено в базе данных
500: Internal Server Error - Ошибка целостности СУБД при удалении сопоставления

---

## [GET] /api/v1/catalog/mapping/{id}/own - Получить изолированное сопоставление по ID

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

400: Bad Request - Неверный числовой формат идентификатора сопоставления
404: Not Found - Изолированное сопоставление с данным ID не найдено
500: Internal Server Error - Внутренняя ошибка сервера при чтении записи

---

## [PATCH] /api/v1/catalog/mapping/{prevId}/{newId} - Изменить ID сопоставления

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
404: Not Found - Сопоставление с исходным ID не существует в системе
500: Internal Server Error - Ошибка СУБД при транзакции изменения идентификатора

---

## [GET] /api/v1/catalog/mappings - Получить все сопоставления

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

500: Internal Server Error - Системная ошибка сервера при извлечении полного списка сопоставлений

# 8. Конфигурация: Структура компонентов устройства

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

# 9. Конфигурация: Связь компонентов устройства и сопоставлений параметров

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

# 10. Конфигурация: Конфигурации по-умолчанию

## [POST] /api/v1/catalog/default-configuration - Создать конфигурацию по умолчанию

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
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении записи

---

## [GET] /api/v1/catalog/default-configuration/{id} - Получить конфигурацию по умолчанию по ID

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
404: Not Found - Конфигурация по умолчанию с указанным ID не найдена
500: Internal Server Error - Системная ошибка СУБД при чтении конфигурации

---

## [PUT] /api/v1/catalog/default-configuration/{id} - Обновить конфигурацию по умолчанию по ID

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
404: Not Found - Конфигурация по умолчанию с указанным ID не обнаружена
500: Internal Server Error - Внутренняя ошибка СУБД при обновлении связей конфигурации

---

## [DELETE] /api/v1/catalog/default-configuration/{id} - Удалить конфигурацию по умолчанию по ID

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

400: Bad Request - Неверный числовой формат идентификатора конфигурации
404: Not Found - Конфигурация по умолчанию с указанным ID не найдена
500: Internal Server Error - Ошибка целостности СУБД при удалении записи

---

## [PATCH] /api/v1/catalog/default-configuration/{prevId}/{newId} - Изменить ID конфигурации по умолчанию

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

400: Bad Request - Неверный формат prevId или newId в пути запроса
404: Not Found - Конфигурация по умолчанию с исходным ID не существует
500: Internal Server Error - Ошибка транзакции изменения идентификатора на уровне СУБД

---

## [GET] /api/v1/catalog/default-configurations - Получить все конфигурации по умолчанию

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

500: Internal Server Error - Критическая ошибка сервера при чтении списка конфигураций по умолчанию

# 11. Конфигурация: Конфигурации устройств

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

# 12. Конфигурация: Пороги

## [POST] /api/v1/catalog/threshold - Создать порог

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

400: Bad Request - Ошибка валидации переданной JSON-структуры порога или пропущены обязательные поля
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении структуры порога

## [GET] /api/v1/catalog/threshold/{id} - Получить порог по ID

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

400: Bad Request - Неверный числовой формат идентификатора порога в пути
404: Not Found - Порог с указанным ID не найден в системе
500: Internal Server Error - Системная ошибка базы данных при извлечении структуры порога

---

## [PUT] /api/v1/catalog/threshold/{id} - Обновить порог по ID

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

400: Bad Request - Ошибка валидации структуры тела или неверный числовой формат ID в пути
404: Not Found - Обновляемый порог с указанным ID не найден
500: Internal Server Error - Внутренняя ошибка СУБД при обновлении полей порога

---

## [DELETE] /api/v1/catalog/threshold/{id} - Удалить порог по ID

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

400: Bad Request - Некорректный числовой формат ID порога в пути запроса
404: Not Found - Удаляемый порог не обнаружен в базе данных
500: Internal Server Error - Ошибка целостности СУБД при каскадном удалении порога

---

## [GET] /api/v1/catalog/threshold/{id}/from-string - Получить эквивалентную строку выражения порога по ID

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

400: Bad Request - Неверный формат ID в пути запроса
404: Not Found - Порог с указанным ID не существует в системе
500: Internal Server Error - Ошибка десериализации дерева условий JSONB в текстовую строку

---

## [PUT] /api/v1/catalog/threshold/{id}/from-string - Обновить порог по ID из строки

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

400: Bad Request - Синтаксическая ошибка при парсинге строкового выражения порога
404: Not Found - Обновляемый порог не найден
500: Internal Server Error - Внутренняя ошибка сервера при конвертации строки в AST-дерево СУБД

---

## [PATCH] /api/v1/catalog/threshold/{prevId}/{newId} - Изменить ID порога

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

400: Bad Request - Неверный числовой формат одного из ID в пути запроса
404: Not Found - Исходный порог с данным ID не найден
500: Internal Server Error - Ошибка обновления первичного или внешних ключей в базе данных

---

## [POST] /api/v1/catalog/threshold/from-string - Создать порог из эквивалентной строки

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

400: Bad Request - Ошибка валидации структуры или синтаксическая ошибка в текстовом выражении query
500: Internal Server Error - Ошибка построения AST-структуры условий и её сохранения в JSONB

---

## [GET] /api/v1/catalog/thresholds - Получить все пороги

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

500: Internal Server Error - Критическая ошибка сервера при чтении полного списка порогов из БД

# 13. Просмотр: В процессе конфигурирования

## [GET] /api/v1/catalog/config/in-progress - Получить все данные устройств, находящихся в процессе конфигурирования

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

404: Not Found - Устройства в процессе конфигурирования отсутствуют
500: Internal Server Error - Системная ошибка сервера при получении списка устройств

---

## [GET] /api/v1/catalog/config/in-progress/search - Поиск данных устройств, находящихся в процессе конфигурирования

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

400: Bad Request - Некорректный формат параметров фильтрации (host или port)
404: Not Found - Искомое устройство в процессе конфигурации не найдено
500: Internal Server Error - Ошибка СУБД при выполнении поиска записи

---

# 14. Просмотр: Рабочая конфигурация

## [GET] /api/v1/catalog/config/working - Получить рабочую конфигурацию устройства

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

400: Bad Request - Отсутствуют или неверно указаны обязательные query-параметры host и port
404: Not Found - Рабочая конфигурация для указанного сетевого узла не найдена
500: Internal Server Error - Внутренняя ошибка сервера при сборке полного графа конфигурации SNMP

---

# 15. Просмотр: Значения параметров

## [GET] /api/v1/catalog/param-result - Получить отфильтрованные сохранённые значения параметров

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

400: Bad Request - Ошибка валидации структуры параметров фильтрации или некорректный формат полей
500: Internal Server Error - Ошибка базы данных при выполнении условной выборки результатов

---

## [GET] /api/v1/catalog/param-results - Получить все сохранённые значения параметров

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

500: Internal Server Error - Критическая ошибка сервера при чтении полной таблицы результатов

---

# 16. Результат: Экспортировать БД в SQL скрипт

## [POST] /api/v1/catalog/save-result - Экспортировать результаты в SQL скрипт

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

400: Bad Request - Ошибка валидации конфигурационных данных для экспорта
500: Internal Server Error - Внутренняя ошибка сервера при записи файлов на диск или дампе таблиц СУБД

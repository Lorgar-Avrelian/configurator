# 14. Просмотр: Рабочая конфигурация

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!TIP]
> API из данного блока предназначен для поиска конфигурации опроса конкретного устройства по его данным.

---

## [GET] /api/v1/catalog/config/working - Получить рабочую конфигурацию устройства

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/config/working/catalog/config/working?host=127.0.0.1&port=161

{}
```

Ответ 1:

```json
{
  "host": "127.0.0.1",
  "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
  "login": "string",
  "password": "string",
  "port": 161,
  "privacy": "AES",
  "authentication": "MD5",
  "community": "public",
  "version": "v2c",
  "components": [
    {
      "access": "USER",
      "base_component": 0,
      "description_en": "string",
      "description_ru": "string",
      "id": 1,
      "name_en": "Component",
      "name_ru": "Компонент",
      "params": [
        {
          "access": "USER",
          "description_en": "string",
          "description_ru": "string",
          "diagram": true,
          "id": 2,
          "name_en": "Name",
          "name_ru": "Имя",
          "saved": true,
          "title": "name",
          "type": "VARCHAR",
          "units_en": "string",
          "units_ru": "string",
          "value": "string",
          "visible": true
        }
      ],
      "plural_name_en": "Components",
      "plural_name_ru": "Компоненты",
      "title": "component"
    }
  ],
  "configuration": {
    "device_component": {
      "children": [
        "string"
      ],
      "component": {
        "id": 1,
        "name_en": "chassis",
        "name_ru": "Шасси",
        "title": "Chassis"
      },
      "id": 1,
      "internal_order": 1,
      "mappings": [
        {
          "children": [
            "string"
          ],
          "coefficient": 0,
          "enum": {},
          "frequency": "MEDIUM",
          "from": 0,
          "id": 1,
          "indicator": {
            "dotter_notation": ".1.3.6.1.2.1.1.2",
            "id": 1,
            "oid": {
              "access": "read-only",
              "category": "system",
              "description": "The vendor's authoritative identification of the network management subsystem contained in the entity.",
              "dotter_notation": ".1.3.6.1.2.1.1.2",
              "enum": {},
              "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
              "mib": {
                "id": 4912,
                "name": "SNMPv2-MIB",
                "path": "SNMPv2-MIB.mib",
                "vendor": "BASE"
              },
              "name": "sysObjectID",
              "number": 2,
              "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
              "status": "current",
              "syntax": "OBJECT IDENTIFIER",
              "type": "OBJECT IDENTIFIER",
              "units": "string"
            }
          },
          "param": {
            "access": "USER",
            "description_en": "string",
            "description_ru": "string",
            "diagram": true,
            "id": 2,
            "name_en": "Name",
            "name_ru": "Имя",
            "saved": true,
            "title": "name",
            "type": "VARCHAR",
            "units_en": "string",
            "units_ru": "string",
            "value": "string",
            "visible": true
          },
          "position": 1,
          "position_type": "INTEGER",
          "value": "string"
        }
      ],
      "parent": 0
    },
    "id": 1,
    "indicator": {
      "contact": "sysadmin@company.com",
      "description": "Linux server-node-01 5.4.0-74-generic",
      "id": 1,
      "location": "Rack 04, Room 202",
      "name": "node-01.local",
      "object_id": ".1.3.6.1.4.1.8072.3.2.10",
      "services": 72
    }
  },,
  "thresholds": [
    {
      "author": "admin",
      "created": "string",
      "description": "string",
      "id": 1,
      "name": "Высокая загрузка CPU",
      "query": {
        "root": {
          "element": {
            "comparison": {
              "operator": ">",
              "target": {
                "host": "string",
                "port": 0,
                "protocol": "string",
                "target": {
                  "component": "string",
                  "field": "string",
                  "internal_order": 0,
                  "next": {},
                  "param": "string"
                }
              },
              "value": "90"
            },
            "expression": {}
          },
          "next": {},
          "operator": "AND"
        }
      },
      "target": {
        "host": "string",
        "port": 0,
        "protocol": "string",
        "target": {
          "component": "string",
          "field": "string",
          "internal_order": 0,
          "next": {},
          "param": "string"
        }
      },
      "value": "ALARM"
    }
  ]
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Отсутствуют или неверно указаны обязательные query-параметры host и port  
404: Not Found - Рабочая конфигурация для указанного сетевого узла не найдена  
500: Internal Server Error - Внутренняя ошибка сервера при сборке полного графа конфигурации SNMP

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
# 7. Конфигурация: Сопоставления параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!NOTE]
> **Сопоставления параметров** - это функция приведения полученного в результате опроса значения (значения параметра на
> устройстве) к истинному значению параметра.

> [!TIP]
> Сопоставление значений может производиться на основании коэффициента, являющегося множителем (например, устройство
> выдаёт значение времени в миллисекундах, а значение параметра отображается в секундах: тогда коэффициент будет равен
> 0.001), либо на основании перечисления (enum) (например, роль устройства в сети, возвращаемая устройством в виде
> целого значения в диапазоне от 1 до 128, но отображаемая в виде строкового описания функциональных возможностей), либо
> иметь какое-либо предопределённое (сконфигурированное) заранее значение, либо быть рассчитано на основании точечной
> нотации OID для точечных нотаций переменной длины.  
> Сопоставления параметров сводят воедино ссылку на параметр, ссылку на индикатор, перечисленные выше правила
> преобразования значения и частоту, с которой следует обновлять значение данного параметра (частоту опроса точечной
> нотации OID).

---

## [POST] /api/v1/catalog/mapping - Создать сопоставление

> [!TIP]
> Для данного API поля 'param', 'indicator' и 'frequency' являются обязательными.  
> Как было сказано выше, приведение значения выполняется по принципу "или-или". Это означает, что для сопоставления
> должны задаваться значения полей или `coefficient`, или `enum`, или `value`. В дополнение к перечисленным полям также
> может задаваться сочетание полей `from`, `position_type` и `position` (последнее сочетание позволяет однозначно
> извлечь значение из точечной нотации OID переменной длины). Это сочетание может быть и самостоятельным значением. 

В поле `position_type` указывается стандартный тип значений согласно спецификации Asn.1/MIB:

- Unspecified;
- Unknown;
- INTEGER;
- BIT STRING;
- OCTET STRING;
- OBJECT IDENTIFIER;
- IpAddress;
- Counter32;
- Gauge32;
- Timeticks;
- Opaque;
- Counter64.

В поле `from` указывается id предыдущего сопоставления, от значения точечной нотации индикатора которого должен
начинаться отсчёт позиции значения внутри точечной нотации OID переменной длины.

В поле `position` указывается порядковый номер позиции значения параметра внутри точечной нотации OID переменной длины.

> [!IMPORTANT]
> Значения полей `from`, `position_type` и `position` должны либо задаваться все, либо не задаваться вовсе.

Поле `frequency` может принимать одно из трёх значений, определяющих частоту опроса параметра:

- LOW - для ежедневного обновления значения;
- MEDIUM - для обновления значения каждые 15 минут (при каждом 3-м опросе);
- HIGH - для обновления значения каждые 5 минут (при каждом опросе).

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping

{
  "param": 21,
  "indicator": 36,
  "frequency": "MEDIUM",
  "coefficient": 1.1
}
```

Ответ 1:

```json
{
  "id": 36,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "MEDIUM",
  "value": null,
  "coefficient": 1.1,
  "enum": null,
  "position": null,
  "from": null,
  "position_type": null,
  "children": []
}
```

</details>

<details><summary>Пример 2</summary>

Запрос 2:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping

{
  "param": 21,
  "indicator": 36,
  "frequency": "LOW",
  "enum": {
    "1": "UP",
    "2": "DOWN"
  }
}
```

Ответ 2:

```json
{
  "id": 37,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "LOW",
  "value": null,
  "coefficient": null,
  "enum": {
    "1": "UP",
    "2": "DOWN"
  },
  "position": null,
  "from": null,
  "position_type": null,
  "children": []
}
```

</details>

<details><summary>Пример 3</summary>

Запрос 3:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping

{
  "param": 21,
  "indicator": 36,
  "frequency": "LOW",
  "value": "TEST"
}
```

Ответ 3:

```json
{
  "id": 38,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "LOW",
  "value": "TEST",
  "coefficient": null,
  "enum": null,
  "position": null,
  "from": null,
  "position_type": null,
  "children": []
}
```

</details>

<details><summary>Пример 4</summary>

Запрос 4:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping

{
  "param": 23,
  "indicator": 36,
  "frequency": "MEDIUM",
  "coefficient": 1.1,
  "from": 38,
  "position": 1,
  "position_type": "OCTET STRING"
}
```

Ответ 4:

```json
{
  "id": 38,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "LOW",
  "value": "TEST",
  "coefficient": null,
  "enum": null,
  "position": null,
  "from": null,
  "position_type": null,
  "children": [
    {
      "id": 39,
      "indicator": {
        "id": 36,
        "oid": {
          "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
          "mib": {
            "id": 4912,
            "path": "SNMPv2-MIB.mib",
            "name": "SNMPv2-MIB",
            "vendor": null
          },
          "type": "OBJECT-TYPE",
          "name": "sysObjectID",
          "number": 2,
          "dotter_notation": ".1.3.6.1.2.1.1.2",
          "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
          "syntax": "OBJECT IDENTIFIER",
          "status": "current",
          "access": "read-only",
          "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
          "category": "system"
        },
        "dotter_notation": ".1.3.6.1.2.1.1.2"
      },
      "param": {
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
      "frequency": "MEDIUM",
      "value": null,
      "coefficient": 1.1,
      "enum": null,
      "position": 1,
      "from": 38,
      "position_type": "OCTET STRING",
      "children": []
    }
  ]
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации структуры или обязательных параметров сопоставления  
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении сопоставления

---

## [GET] /api/v1/catalog/mapping/{id} - Получить сопоставление по ID

> [!TIP]
> При получении сопоставления по id нужно отобразить всю цепочку точечной нотации составного OID. То есть родительскую
> точечную нотацию, позицию следования после неё и содержание дочернего сопоставления.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping/35

{}
```

Ответ 1:

```json
{
  "id": 35,
  "indicator": {
    "id": 35,
    "oid": {
      "id": "0b816f5a-d2f6-3488-8003-ae528ab08ece",
      "mib": {
        "id": 4705,
        "path": "HOST-RESOURCES-MIB.mib",
        "name": "HOST-RESOURCES-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "hrStorageUsed",
      "number": 6,
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.6",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageUsed",
      "syntax": "Integer32 (0... 2147483647)",
      "status": "current",
      "access": "read-only",
      "description": "The amount of the storage represented by this entry that is allocated, in units of hrStorageAllocationUnits.",
      "category": "hrStorageEntry"
    },
    "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.6"
  },
  "param": {
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
  },
  "frequency": "HIGH",
  "value": null,
  "coefficient": null,
  "enum": null,
  "position": null,
  "from": null,
  "position_type": null,
  "children": []
}
```

</details>

<details><summary>Пример 2</summary>

Запрос 2:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping/39

{}
```

Ответ 2:

```json
{
  "id": 38,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "LOW",
  "value": "TEST",
  "coefficient": null,
  "enum": null,
  "position": null,
  "from": null,
  "position_type": null,
  "children": [
    {
      "id": 39,
      "indicator": {
        "id": 36,
        "oid": {
          "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
          "mib": {
            "id": 4912,
            "path": "SNMPv2-MIB.mib",
            "name": "SNMPv2-MIB",
            "vendor": null
          },
          "type": "OBJECT-TYPE",
          "name": "sysObjectID",
          "number": 2,
          "dotter_notation": ".1.3.6.1.2.1.1.2",
          "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
          "syntax": "OBJECT IDENTIFIER",
          "status": "current",
          "access": "read-only",
          "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
          "category": "system"
        },
        "dotter_notation": ".1.3.6.1.2.1.1.2"
      },
      "param": {
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
      "frequency": "MEDIUM",
      "value": null,
      "coefficient": 1.1,
      "enum": null,
      "position": 1,
      "from": 38,
      "position_type": "OCTET STRING",
      "children": []
    }
  ]
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат идентификатора сопоставления в пути  
404: Not Found - Сопоставление с указанным ID не найдено в системе  
500: Internal Server Error - Системная ошибка базы данных при извлечении сопоставления

---

## [PUT] /api/v1/catalog/mapping/{id} - Обновить сопоставление по ID

> [!TIP]
> На операцию обновления данных сопоставления распространяются те же требования, что и на операцию создания
> сопоставления.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping/39

{
  "param": 23,
  "indicator": 36,
  "frequency": "MEDIUM",
  "coefficient": 1.1,
  "from": 38,
  "position": 1,
  "position_type": "OCTET STRING"
}
```

Ответ 1:

```json
{
  "id": 38,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "LOW",
  "value": "TEST",
  "coefficient": null,
  "enum": null,
  "position": null,
  "from": null,
  "position_type": null,
  "children": [
    {
      "id": 39,
      "indicator": {
        "id": 36,
        "oid": {
          "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
          "mib": {
            "id": 4912,
            "path": "SNMPv2-MIB.mib",
            "name": "SNMPv2-MIB",
            "vendor": null
          },
          "type": "OBJECT-TYPE",
          "name": "sysObjectID",
          "number": 2,
          "dotter_notation": ".1.3.6.1.2.1.1.2",
          "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
          "syntax": "OBJECT IDENTIFIER",
          "status": "current",
          "access": "read-only",
          "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
          "category": "system"
        },
        "dotter_notation": ".1.3.6.1.2.1.1.2"
      },
      "param": {
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
      "frequency": "MEDIUM",
      "value": null,
      "coefficient": 1.1,
      "enum": null,
      "position": 1,
      "from": 38,
      "position_type": "OCTET STRING",
      "children": []
    }
  ]
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный ID в пути или ошибка валидации JSON-структуры запроса  
404: Not Found - Сопоставление для обновления с указанным ID не найдено  
500: Internal Server Error - Ошибка обновления записи на стороне базы данных

---

## [DELETE] /api/v1/catalog/mapping/{id} - Удалить сопоставление по ID

> [!CAUTION]
> При удалении сопоставления параметра из БД также удаляются все дочерние сопоставления параметров.

> [!IMPORTANT]
> При удалении сопоставления параметра происходит сдвиг (уменьшение значения id на 1 и более) всего списка
> сопоставлений, следовавших за удалённым, что не приводит к сбоям в работе системы.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping/38

{}
```

Ответ 1:

```json
{}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный формат ID сопоставления в пути запроса  
404: Not Found - Удаляемое сопоставление не обнаружено в базе данных  
500: Internal Server Error - Ошибка целостности СУБД при удалении сопоставления

---

## [GET] /api/v1/catalog/mapping/{id}/own - Получить изолированное сопоставление по ID

> [!TIP]
> Данный API предназначен для получения дочернего сопоставления без привязки к родительскому сопоставлению.
> Удобен для целей редактирования отдельного узла.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping/37/own

{}
```

Ответ 1:

```json
{
  "id": 37,
  "indicator": {
    "id": 36,
    "oid": {
      "id": "14e8713a-2f3c-3af1-8e6f-449d7a612227",
      "mib": {
        "id": 4912,
        "path": "SNMPv2-MIB.mib",
        "name": "SNMPv2-MIB",
        "vendor": null
      },
      "type": "OBJECT-TYPE",
      "name": "sysObjectID",
      "number": 2,
      "dotter_notation": ".1.3.6.1.2.1.1.2",
      "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
      "syntax": "OBJECT IDENTIFIER",
      "status": "current",
      "access": "read-only",
      "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.424242, it could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.",
      "category": "system"
    },
    "dotter_notation": ".1.3.6.1.2.1.1.2"
  },
  "param": {
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
  "frequency": "LOW",
  "value": null,
  "coefficient": null,
  "enum": {
    "1": "UP",
    "2": "DOWN"
  },
  "position": null,
  "from": null,
  "position_type": null,
  "children": []
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат идентификатора сопоставления  
404: Not Found - Изолированное сопоставление с данным ID не найдено  
500: Internal Server Error - Внутренняя ошибка сервера при чтении записи

---

## [PATCH] /api/v1/catalog/mapping/{prevId}/{newId} - Изменить ID сопоставления

> [!TIP]
> API предназначен для упорядочивания перечня сопоставлений параметров и имеет одно ограничение: id
> сопоставления-потомка должно следовать строго после (быть больше) id сопоставления-родителя.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PATCH https://nms-dev.opk-bulat.ru/api/v1/catalog/mapping/36/37

{}
```

Ответ 1:

```json
{}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат одного из ID (prevId или newId) в пути запроса  
404: Not Found - Сопоставление с исходным ID не существует в системе  
500: Internal Server Error - Ошибка СУБД при транзакции изменения идентификатора

---

## [GET] /api/v1/catalog/mappings - Получить все сопоставления

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/mappings

{}
```

Ответ 1:

```json
[
  {
    "id": 1,
    "indicator": {
      "id": 1,
      "oid": {
        "id": "7ad45d0c-670b-35af-a76f-a6cf59387c6e",
        "mib": {
          "id": 4912,
          "path": "SNMPv2-MIB.mib",
          "name": "SNMPv2-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "sysDescr",
        "number": 1,
        "dotter_notation": ".1.3.6.1.2.1.1.1",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysDescr",
        "syntax": "DisplayString (SIZE (0... 255))",
        "status": "current",
        "access": "read-only",
        "description": "A textual description of the entity. This value should include the full name and version identification of the system's hardware type, software operating-system, and networking software.",
        "category": "system"
      },
      "dotter_notation": ".1.3.6.1.2.1.1.1"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 2,
    "indicator": {
      "id": 2,
      "oid": {
        "id": "a52b392f-e4f1-3c8e-9915-402ad93ce41e",
        "mib": {
          "id": 4912,
          "path": "SNMPv2-MIB.mib",
          "name": "SNMPv2-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "sysUpTime",
        "number": 3,
        "dotter_notation": ".1.3.6.1.2.1.1.3",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysUpTime",
        "syntax": "TimeTicks",
        "status": "current",
        "access": "read-only",
        "description": "The time (in hundredths of a second) since the network management portion of the system was last re-initialized.",
        "category": "system"
      },
      "dotter_notation": ".1.3.6.1.2.1.1.3"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 3,
    "indicator": {
      "id": 3,
      "oid": {
        "id": "9a19a17b-17d7-313d-af89-0072bdc20ac3",
        "mib": {
          "id": 4912,
          "path": "SNMPv2-MIB.mib",
          "name": "SNMPv2-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "sysContact",
        "number": 4,
        "dotter_notation": ".1.3.6.1.2.1.1.4",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysContact",
        "syntax": "DisplayString (SIZE (0... 255))",
        "status": "current",
        "access": "read-write",
        "description": "The textual identification of the contact person for this managed node, together with information on how to contact this person. If no contact information is known, the value is the zero-length string.",
        "category": "system"
      },
      "dotter_notation": ".1.3.6.1.2.1.1.4"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 4,
    "indicator": {
      "id": 4,
      "oid": {
        "id": "57916063-e436-3dfd-9c0e-3baa10b92a06",
        "mib": {
          "id": 4912,
          "path": "SNMPv2-MIB.mib",
          "name": "SNMPv2-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "sysName",
        "number": 5,
        "dotter_notation": ".1.3.6.1.2.1.1.5",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysName",
        "syntax": "DisplayString (SIZE (0... 255))",
        "status": "current",
        "access": "read-write",
        "description": "An administratively-assigned name for this managed node. By convention, this is the node's fully-qualified domain name. If the name is unknown, the value is the zero-length string.",
        "category": "system"
      },
      "dotter_notation": ".1.3.6.1.2.1.1.5"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 5,
    "indicator": {
      "id": 5,
      "oid": {
        "id": "3e3a0474-1208-3b78-9e9f-5cb1df2acd17",
        "mib": {
          "id": 4912,
          "path": "SNMPv2-MIB.mib",
          "name": "SNMPv2-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "sysLocation",
        "number": 6,
        "dotter_notation": ".1.3.6.1.2.1.1.6",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysLocation",
        "syntax": "DisplayString (SIZE (0... 255))",
        "status": "current",
        "access": "read-write",
        "description": "The physical location of this node (e.g., 'telephone closet, 3rd floor'). If the location is unknown, the value is the zero-length string.",
        "category": "system"
      },
      "dotter_notation": ".1.3.6.1.2.1.1.6"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 6,
    "indicator": {
      "id": 6,
      "oid": {
        "id": "90059161-6896-382d-bd01-caeb9e3f0929",
        "mib": {
          "id": 4912,
          "path": "SNMPv2-MIB.mib",
          "name": "SNMPv2-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "sysServices",
        "number": 7,
        "dotter_notation": ".1.3.6.1.2.1.1.7",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysServices",
        "syntax": "INTEGER (0... 127)",
        "status": "current",
        "access": "read-only",
        "description": "A value which indicates the set of services that this entity may potentially offer. The value is a sum. This sum initially takes the value zero. Then, for each layer, L, in the range 1 through 7, that this node performs transactions for, 2 raised to (L - 1) is added to the sum. For example, a node which performs only routing functions would have a value of 4 (2^ (3-1)). In contrast, a node which is a host offering application services would have a value of 72 (2^ (4-1) + 2^ (7-1)). Note that in the context of the Internet suite of protocols, values should be calculated accordingly: layer functionality 1 physical (e.g., repeaters) 2 datalink/subnetwork (e.g., bridges) 3 internet (e.g., supports the IP) 4 end-to-end (e.g., supports the TCP) 7 applications (e.g., supports the SMTP) For systems including OSI protocols, layers 5 and 6 may also be counted.",
        "category": "system"
      },
      "dotter_notation": ".1.3.6.1.2.1.1.7"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": {
      "0": "Unknown or No Services",
      "1": "L1: Physical Layer Device (Hub/Repeater)",
      "2": "L2: Data Link Layer Device (L2 Switch/Bridge)",
      "3": "L1+L2: Managed L2 Switch",
      "4": "L3: Network Layer Device (Router)",
      "5": "L1+L3: Hardware Router",
      "6": "L2+L3: L3 Switch (No Host Services)",
      "7": "L1+L2+L3: Enterprise L3 Switch / Router",
      "8": "L4: Transport Layer Device",
      "9": "L1+L4: Transport Gateway",
      "10": "L2+L4: L2/L4 Traffic Shaper",
      "11": "L1+L2+L4: L2/L4 Managed Device",
      "12": "L3+L4: Core Router / Filtering Router",
      "13": "L1+L3+L4: Core Router Hardware",
      "14": "L2+L3+L4: L3/L4 Advanced Switch",
      "15": "L1+L2+L3+L4: Firewall / Packet Filter",
      "16": "L5: Session Layer Device",
      "17": "L1+L5: Session Gateway",
      "18": "L2+L5: Session Device",
      "19": "L1+L2+L5: Session Device",
      "20": "L3+L5: Session Router",
      "21": "L1+L3+L5: Session Router",
      "22": "L2+L3+L5: Session Router",
      "23": "L1+L2+L3+L5: Session Router",
      "24": "L4+L5: Session/Transport Device",
      "25": "L1+L4+L5: Session Device",
      "26": "L2+L4+L5: Session Device",
      "27": "L1+L2+L4+L5: Session Device",
      "28": "L3+L4+L5: Session Router",
      "29": "L1+L3+L4+L5: Session Router",
      "30": "L2+L3+L4+L5: Session Router",
      "31": "L1+L2+L3+L4+L5: Session Router",
      "32": "L6: Presentation Layer Device",
      "33": "L1+L6: Presentation Device",
      "34": "L2+L6: Presentation Device",
      "35": "L1+L2+L6: Presentation Device",
      "36": "L3+L6: Presentation Router",
      "37": "L1+L3+L6: Presentation Router",
      "38": "L2+L3+L6: Presentation Router",
      "39": "L1+L2+L3+L6: Presentation Router",
      "40": "L4+L6: Presentation Device",
      "41": "L1+L4+L6: Presentation Device",
      "42": "L2+L4+L6: Presentation Device",
      "43": "L1+L2+L4+L6: Presentation Device",
      "44": "L3+L4+L6: Presentation Router",
      "45": "L1+L3+L4+L6: Presentation Router",
      "46": "L2+L3+L4+L6: Presentation Router",
      "47": "L1+L2+L3+L4+L6: Presentation Router",
      "48": "L5+L6: Session/Presentation Device",
      "49": "L1+L5+L6: Presentation Device",
      "50": "L2+L5+L6: Presentation Device",
      "51": "L1+L2+L5+L6: Presentation Device",
      "52": "L3+L5+L6: Presentation Router",
      "53": "L1+L3+L5+L6: Presentation Router",
      "54": "L2+L3+L5+L6: Presentation Router",
      "55": "L1+L2+L3+L5+L6: Presentation Router",
      "56": "L4+L5+L6: Presentation Device",
      "57": "L1+L4+L5+L6: Presentation Device",
      "58": "L2+L4+L5+L6: Presentation Device",
      "59": "L1+L2+L4+L5+L6: Presentation Device",
      "60": "L3+L4+L5+L6: Presentation Router",
      "61": "L1+L3+L4+L5+L6: Presentation Router",
      "62": "L2+L3+L4+L5+L6: Presentation Router",
      "63": "L1+L2+L3+L4+L5+L6: Presentation Router",
      "64": "L7: Isolated Application / Service",
      "65": "L1+L7: Application End-Device",
      "66": "L2+L7: Network End-Device (Printer/IP-Camera/IoT)",
      "67": "L1+L2+L7: Managed Network End-Device",
      "68": "L3+L7: Application Router",
      "69": "L1+L3+L7: Application Router Hardware",
      "70": "L2+L3+L7: Multi-Layer Endpoint",
      "71": "L1+L2+L3+L7: Gateway Endpoint",
      "72": "L4+L7: Standard Host / Server / PC (Linux/Windows/macOS)",
      "73": "L1+L4+L7: Standard Host Hardware",
      "74": "L2+L4+L7: Local Server / NAS Storage",
      "75": "L1+L2+L4+L7: Managed Local Server",
      "76": "L3+L4+L7: Routing Host / Software Router",
      "77": "L1+L3+L4+L7: Routing Host Hardware",
      "78": "L2+L3+L4+L7: Multi-Layer Router",
      "79": "L1+L2+L3+L4+L7: Enterprise Gateway / NGFW / UTM / Proxy",
      "80": "L5+L7: App/Session Host",
      "81": "L1+L5+L7: App Host",
      "82": "L2+L5+L7: App Host",
      "83": "L1+L2+L5+L7: App Host",
      "84": "L3+L5+L7: App Router",
      "85": "L1+L3+L5+L7: App Router",
      "86": "L2+L3+L5+L7: App Router",
      "87": "L1+L2+L3+L5+L7: App Router",
      "88": "L4+L5+L7: App Server",
      "89": "L1+L4+L5+L7: App Server",
      "90": "L2+L4+L5+L7: App Server",
      "91": "L1+L2+L4+L5+L7: App Server",
      "92": "L3+L4+L5+L7: App Gateway",
      "93": "L1+L3+L4+L5+L7: App Gateway",
      "94": "L2+L3+L4+L5+L7: App Gateway",
      "95": "L1+L2+L3+L4+L5+L7: App Gateway",
      "96": "L6+L7: App/Presentation Host",
      "97": "L1+L6+L7: App Host",
      "98": "L2+L6+L7: App Host",
      "99": "L1+L2+L6+L7: App Host",
      "100": "L3+L6+L7: App Router",
      "101": "L1+L3+L6+L7: App Router",
      "102": "L2+L3+L6+L7: App Router",
      "103": "L1+L2+L3+L6+L7: App Router",
      "104": "L4+L6+L7: App Server",
      "105": "L1+L4+L6+L7: App Server",
      "106": "L2+L4+L6+L7: App Server",
      "107": "L1+L2+L4+L6+L7: App Server",
      "108": "L3+L4+L6+L7: App Gateway",
      "109": "L1+L3+L4+L6+L7: App Gateway",
      "110": "L2+L3+L4+L6+L7: App Gateway",
      "111": "L1+L2+L3+L4+L6+L7: App Gateway",
      "112": "L5+L6+L7: Multi-Layer App Host",
      "113": "L1+L5+L6+L7: App Host",
      "114": "L2+L5+L6+L7: App Host",
      "115": "L1+L2+L5+L6+L7: App Host",
      "116": "L3+L5+L6+L7: App Router",
      "117": "L1+L3+L5+L6+L7: App Router",
      "118": "L2+L3+L5+L6+L7: App Router",
      "119": "L1+L2+L3+L5+L6+L7: App Router",
      "120": "L4+L5+L6+L7: Full-Stack Server",
      "121": "L1+L4+L5+L6+L7: Full-Stack Server",
      "122": "L2+L4+L5+L6+L7: Full-Stack Server",
      "123": "L1+L2+L4+L5+L6+L7: Full-Stack Server",
      "124": "L3+L4+L5+L6+L7: Full-Stack Gateway",
      "125": "L1+L3+L4+L5+L6+L7: Full-Stack Gateway",
      "126": "L2+L3+L4+L5+L6+L7: Full-Stack Gateway",
      "127": "L1+L2+L3+L4+L5+L6+L7: Full 7-Layer OSI Device"
    },
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 7,
    "indicator": {
      "id": 7,
      "oid": {
        "id": "c4757f71-1042-3c2d-ad93-f70bcf1873d1",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifNumber",
        "number": 1,
        "dotter_notation": ".1.3.6.1.2.1.2.1",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifNumber",
        "syntax": "INTEGER",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of network interfaces (regardless of their current state) present on this system.",
        "category": "interfaces"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.1"
    },
    "param": {
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
    "frequency": "MEDIUM",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 8,
    "indicator": {
      "id": 8,
      "oid": {
        "id": "56323d01-cb9b-34bd-8986-68496c2a5302",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifIndex",
        "number": 1,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.1",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifIndex",
        "syntax": "INTEGER",
        "status": "mandatory",
        "access": "read-only",
        "description": "A unique value for each interface. Its value ranges between 1 and the value of ifNumber. The value for each interface must remain constant at least from one re-initialization of the entity's network management system to the next re - initialization.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.1"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 9,
    "indicator": {
      "id": 9,
      "oid": {
        "id": "7940375b-985f-3198-86e3-e29cdaf3221a",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifDescr",
        "number": 2,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.2",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifDescr",
        "syntax": "DisplayString (SIZE (0... 255))",
        "status": "mandatory",
        "access": "read-only",
        "description": "A textual string containing information about the interface. This string should include the name of the manufacturer, the product name and the version of the hardware interface.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.2"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 10,
    "indicator": {
      "id": 10,
      "oid": {
        "id": "4ba029a5-fc64-37af-90f9-95326bb4084e",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifType",
        "number": 3,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.3",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifType",
        "syntax": "INTEGER",
        "status": "mandatory",
        "access": "read-only",
        "description": "The type of interface, distinguished according to the physical/link protocol (s) immediately `below' the network layer in the protocol stack.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.3"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 11,
    "indicator": {
      "id": 11,
      "oid": {
        "id": "1760f0e6-1921-3fed-8f12-101bbea5320e",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifMtu",
        "number": 4,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.4",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifMtu",
        "syntax": "INTEGER",
        "status": "mandatory",
        "access": "read-only",
        "description": "The size of the largest datagram which can be sent/received on the interface, specified in octets. For interfaces that are used for transmitting network datagrams, this is the size of the largest network datagram that can be sent on the interface.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.4"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 12,
    "indicator": {
      "id": 12,
      "oid": {
        "id": "bbd1e80c-9abb-3c45-8b34-8c6279e71627",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifSpeed",
        "number": 5,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.5",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifSpeed",
        "syntax": "Gauge",
        "status": "mandatory",
        "access": "read-only",
        "description": "An estimate of the interface's current bandwidth in bits per second. For interfaces which do not vary in bandwidth or for those where no accurate estimation can be made, this object should contain the nominal bandwidth.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.5"
    },
    "param": {
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
    "frequency": "MEDIUM",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 13,
    "indicator": {
      "id": 13,
      "oid": {
        "id": "6199b813-1667-37f9-9c3c-59a25f724bd9",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifPhysAddress",
        "number": 6,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.6",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifPhysAddress",
        "syntax": "OCTET STRING",
        "status": "mandatory",
        "access": "read-only",
        "description": "The interface's address at the protocol layer immediately `below' the network layer in the protocol stack. For interfaces which do not have such an address (e.g., a serial line), this object should contain an octet string of zero length.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.6"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 14,
    "indicator": {
      "id": 14,
      "oid": {
        "id": "6e09ed60-7c87-302f-ba7a-9301e73c4235",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifAdminStatus",
        "number": 7,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.7",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifAdminStatus",
        "syntax": "INTEGER",
        "status": "mandatory",
        "access": "read-write",
        "description": "The desired state of the interface. The testing (3) state indicates that no operational packets can be passed.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.7"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 15,
    "indicator": {
      "id": 15,
      "oid": {
        "id": "e188d301-fbf2-38d8-8da9-4b697cc84f09",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOperStatus",
        "number": 8,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.8",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOperStatus",
        "syntax": "INTEGER",
        "status": "mandatory",
        "access": "read-only",
        "description": "The current operational state of the interface. The testing (3) state indicates that no operational packets can be passed.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.8"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 16,
    "indicator": {
      "id": 16,
      "oid": {
        "id": "99dc3487-3fed-3a0b-b956-d25a182072aa",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifLastChange",
        "number": 9,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.9",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifLastChange",
        "syntax": "TimeTicks",
        "status": "mandatory",
        "access": "read-only",
        "description": "The value of sysUpTime at the time the interface entered its current operational state. If the current state was entered prior to the last re - initialization of the local network management subsystem, then this object contains a zero value.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.9"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 17,
    "indicator": {
      "id": 17,
      "oid": {
        "id": "0023e8dd-7874-3006-a7f7-85c299be0b4b",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifInOctets",
        "number": 10,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.10",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifInOctets",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The total number of octets received on the interface, including framing characters.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.10"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 18,
    "indicator": {
      "id": 18,
      "oid": {
        "id": "3fa79c22-59e5-3305-87e9-452c83b6472b",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifInUcastPkts",
        "number": 11,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.11",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifInUcastPkts",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of subnetwork-unicast packets delivered to a higher-layer protocol.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.11"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 19,
    "indicator": {
      "id": 19,
      "oid": {
        "id": "b868cab1-7bcb-31b8-9ded-d48d5d4db081",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifInNUcastPkts",
        "number": 12,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.12",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifInNUcastPkts",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of non-unicast (i.e., subnetwork - broadcast or subnetwork-multicast) packets delivered to a higher-layer protocol.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.12"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 20,
    "indicator": {
      "id": 20,
      "oid": {
        "id": "d240b5ff-c566-32d0-980a-6316300527e3",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifInDiscards",
        "number": 13,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.13",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifInDiscards",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of inbound packets which were chosen to be discarded even though no errors had been detected to prevent their being deliverable to a higher-layer protocol. One possible reason for discarding such a packet could be to free up buffer space.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.13"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 21,
    "indicator": {
      "id": 21,
      "oid": {
        "id": "97be9bd6-e3f5-3bcb-9d3e-b0d6f670fb7e",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifInErrors",
        "number": 14,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.14",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifInErrors",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of inbound packets that contained errors preventing them from being deliverable to a higher-layer protocol.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.14"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 22,
    "indicator": {
      "id": 22,
      "oid": {
        "id": "770c3d04-c0b8-376c-be9c-32fa006c7bdb",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifInUnknownProtos",
        "number": 15,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.15",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifInUnknownProtos",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of packets received via the interface which were discarded because of an unknown or unsupported protocol.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.15"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 23,
    "indicator": {
      "id": 23,
      "oid": {
        "id": "c0a5bf88-69f6-380c-b2ea-1def8176e7ac",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOutOctets",
        "number": 16,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.16",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOutOctets",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The total number of octets transmitted out of the interface, including framing characters.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.16"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 24,
    "indicator": {
      "id": 24,
      "oid": {
        "id": "26bcdad2-5a4e-3d5c-96d2-fe879aef7d80",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOutUcastPkts",
        "number": 17,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.17",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOutUcastPkts",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The total number of packets that higher-level protocols requested be transmitted to a subnetwork-unicast address, including those that were discarded or not sent.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.17"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 25,
    "indicator": {
      "id": 25,
      "oid": {
        "id": "6efd22d7-9000-3adc-a28d-f6a16273776a",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOutNUcastPkts",
        "number": 18,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.18",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOutNUcastPkts",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The total number of packets that higher-level protocols requested be transmitted to a non - unicast (i.e., a subnetwork-broadcast or subnetwork-multicast) address, including those that were discarded or not sent.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.18"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 26,
    "indicator": {
      "id": 26,
      "oid": {
        "id": "b6bcb074-2367-3582-a664-5ac84a120c30",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOutDiscards",
        "number": 19,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.19",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOutDiscards",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of outbound packets which were chosen to be discarded even though no errors had been detected to prevent their being transmitted. One possible reason for discarding such a packet could be to free up buffer space.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.19"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 27,
    "indicator": {
      "id": 27,
      "oid": {
        "id": "4e06d5b8-003b-3acf-9c3c-4184749a891b",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOutErrors",
        "number": 20,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.20",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOutErrors",
        "syntax": "Counter",
        "status": "mandatory",
        "access": "read-only",
        "description": "The number of outbound packets that could not be transmitted because of errors.",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.20"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 28,
    "indicator": {
      "id": 28,
      "oid": {
        "id": "047d8706-a3d3-390b-893c-33998b907fc8",
        "mib": {
          "id": 4888,
          "path": "RFC1213-MIB.mib",
          "name": "RFC1213-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "ifOutQLen",
        "number": 21,
        "dotter_notation": ".1.3.6.1.2.1.2.2.1.21",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.interfaces.ifTable.ifEntry.ifOutQLen",
        "syntax": "Gauge",
        "status": "mandatory",
        "access": "read-only",
        "description": "The length of the output packet queue (in packets).",
        "category": "ifEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.2.2.1.21"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 29,
    "indicator": {
      "id": 29,
      "oid": {
        "id": "78791a7d-ba52-30ee-a86c-d4f4330a8336",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrProcessorLoad",
        "number": 2,
        "dotter_notation": ".1.3.6.1.2.1.25.3.3.1.2",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrDevice.hrProcessorTable.hrProcessorEntry.hrProcessorLoad",
        "syntax": "Integer32 (0... 100)",
        "status": "current",
        "access": "read-only",
        "description": "The average, over the last minute, of the percentage of time that this processor was not idle. Implementations may approximate this one minute smoothing period if necessary.",
        "category": "hrProcessorEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.3.3.1.2"
    },
    "param": {
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
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 30,
    "indicator": {
      "id": 30,
      "oid": {
        "id": "9d3f834f-70cc-33f7-aca4-92e066a4c56e",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrStorageIndex",
        "number": 1,
        "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.1",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageIndex",
        "syntax": "Integer32 (1... 2147483647)",
        "status": "current",
        "access": "read-only",
        "description": "A unique value for each logical storage area contained by the host.",
        "category": "hrStorageEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.1"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 31,
    "indicator": {
      "id": 31,
      "oid": {
        "id": "ebb6c104-b7aa-32e7-9ebf-571326bd98c8",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrStorageType",
        "number": 2,
        "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.2",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageType",
        "syntax": "OBJECT IDENTIFIER",
        "status": "current",
        "access": "read-only",
        "description": "The type of storage represented by this entry.",
        "category": "hrStorageEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.2"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": {
      ".1.3.6.1.2.1.25.2.1.1": "7",
      ".1.3.6.1.2.1.25.2.1.10": "16",
      ".1.3.6.1.2.1.25.2.1.2": "8",
      ".1.3.6.1.2.1.25.2.1.3": "10",
      ".1.3.6.1.2.1.25.2.1.4": "9",
      ".1.3.6.1.2.1.25.2.1.5": "11",
      ".1.3.6.1.2.1.25.2.1.6": "12",
      ".1.3.6.1.2.1.25.2.1.7": "13",
      ".1.3.6.1.2.1.25.2.1.8": "14",
      ".1.3.6.1.2.1.25.2.1.9": "15"
    },
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 32,
    "indicator": {
      "id": 32,
      "oid": {
        "id": "2b4616e3-66b3-3219-ba93-3c2909653510",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrStorageDescr",
        "number": 3,
        "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.3",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageDescr",
        "syntax": "OCTET STRING (SIZE (0... 255))",
        "status": "current",
        "access": "read-only",
        "description": "A description of the type and instance of the storage described by this entry.",
        "category": "hrStorageEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.3"
    },
    "param": {
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
    "frequency": "LOW",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 33,
    "indicator": {
      "id": 33,
      "oid": {
        "id": "8ed064ce-e18e-3efa-8aad-8f8ff1c74074",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrStorageAllocationUnits",
        "number": 4,
        "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.4",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageAllocationUnits",
        "syntax": "Integer32 (1... 2147483647)",
        "status": "current",
        "access": "read-only",
        "units": "Bytes",
        "description": "The size, in bytes, of the data objects allocated from this pool. If this entry is monitoring sectors, blocks, buffers, or packets, for example, this number will commonly be greater than one. Otherwise this number will typically be one.",
        "category": "hrStorageEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.4"
    },
    "param": {
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
    "frequency": "MEDIUM",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 34,
    "indicator": {
      "id": 34,
      "oid": {
        "id": "02bb3b81-1efe-3249-8515-08f2527fdb14",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrStorageSize",
        "number": 5,
        "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.5",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageSize",
        "syntax": "Integer32 (0... 2147483647)",
        "status": "current",
        "access": "read-write",
        "description": "The size of the storage represented by this entry, in units of hrStorageAllocationUnits. This object is writable to allow remote configuration of the size of the storage area in those cases where such an operation makes sense and is possible on the underlying system. For example, the amount of main memory allocated to a buffer pool might be modified or the amount of disk space allocated to virtual memory might be modified.",
        "category": "hrStorageEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.5"
    },
    "param": {
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
    "frequency": "MEDIUM",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  },
  {
    "id": 35,
    "indicator": {
      "id": 35,
      "oid": {
        "id": "0b816f5a-d2f6-3488-8003-ae528ab08ece",
        "mib": {
          "id": 4705,
          "path": "HOST-RESOURCES-MIB.mib",
          "name": "HOST-RESOURCES-MIB",
          "vendor": null
        },
        "type": "OBJECT-TYPE",
        "name": "hrStorageUsed",
        "number": 6,
        "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.6",
        "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.host.hrStorage.hrStorageTable.hrStorageEntry.hrStorageUsed",
        "syntax": "Integer32 (0... 2147483647)",
        "status": "current",
        "access": "read-only",
        "description": "The amount of the storage represented by this entry that is allocated, in units of hrStorageAllocationUnits.",
        "category": "hrStorageEntry"
      },
      "dotter_notation": ".1.3.6.1.2.1.25.2.3.1.6"
    },
    "param": {
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
    },
    "frequency": "HIGH",
    "value": null,
    "coefficient": null,
    "enum": null,
    "position": null,
    "from": null,
    "position_type": null,
    "children": []
  }
]
```

</details>

</details>

### Возможные коды ошибок

500: Internal Server Error - Системная ошибка сервера при извлечении полного списка сопоставлений

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
# 6. Конфигурация: Индикаторы параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!NOTE]
> **Индикаторы параметров** - это точечные нотации OID, привязанные к конкретному описанию в системе.

> [!TIP]
> Каждое описание OID в системе, загруженное из того или иного MIB, имеет свой уникальный идентификатор UUID.

> [!IMPORTANT]
> Использование API из данного блока должно производиться во взаимодействии с API
> [блока SNMP MIB Parser](./description_04.md) для поиска значения конкретного описания OID, к точечной нотации которого
> будет производиться привязка параметра в последующих шагах.  
> У пользователя также должна быть возможность ввести вручную необходимую точечную нотацию в случае, если искомого им
> описания ещё нет в системе (необходимый MIB ещё не был загружен в систему). В этом случае индикатор параметра
> создаётся без указания конкретного id (имеющего тип UUID) описания OID, но пользователя нужно будет предупредить о том
> что описание указанной им точечной нотации будет автоматически обновлено при последующей загрузке в систему нового
> файла MIB, имеющего подходящее описание.

---

## [POST] /api/v1/catalog/indicator/param - Создать индикатор параметров

> [!TIP]
> Индикаторы параметров являются уникальными. Повторная отправка запроса на создание нового индикатора, совпадающая по
> содержанию с уже существующим индикатором, вернёт данные оригинала (новый индикатор создан не будет).

> [!WARNING]
> При создании индикатора параметра может быть отправлено значение только `oid_id`. В этом случае значение
> точечной нотации будет подставлено автоматически.  
> При создании индикатора путём отправки только `dotter_notation` индикатор параметров будет создан без привязки к
> описанию OID в любом случае. Этот вариант должен использоваться для ручного ввода точечной нотации.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param

{
  "dotter_notation": ".1.3.6.1.2.1.1.2",
  "oid_id": "14e8713a-2f3c-3af1-8e6f-449d7a612227"
}
```

Ответ 1:

```json
{
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
}
```

</details>

<details><summary>Пример 2</summary>

Запрос 2:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param

{
  "dotter_notation": ".1.3.6.1.2.1.1.2"
}
```

Ответ 2:

```json
{
  "id": 37,
  "oid": null,
  "dotter_notation": ".1.3.6.1.2.1.1.2"
}
```

</details>

<details><summary>Пример 3</summary>

Запрос 3:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param

{
  "oid_id": "14e8713a-2f3c-3af1-8e6f-449d7a612227"
}
```

Ответ 3:

```json
{
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
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации структуры (не переданы oid_id или dotter_notation)  
500: Internal Server Error - Ошибка вставки записи индикатора параметра в СУБД

---

## [GET] /api/v1/catalog/indicator/param/{id} - Получить индикатор параметров по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param/36

{}
```

Ответ 1:

```json
{
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
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат идентификатора в пути URL  
404: Not Found - Индикатор параметров с указанным ID не найден  
500: Internal Server Error - Внутренняя ошибка базы данных

---

## [PUT] /api/v1/catalog/indicator/param/{id} - Обновить индикатор параметров по ID

> [!IMPORTANT]
> При редактировании индикатора параметра учитывается условие уникальности сочетания `dotter_notation` и `oid_id`. Таким
> образом, в случае редактирования индикатора параметра до состояния уже созданного ранее индикатора редактируемый
> индикатор будет удалён, а возвращено будет значение оригинального индикатора.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param/37

{
  "oid_id": "14e8713a-2f3c-3af1-8e6f-449d7a612227"
}
```

Ответ 1:

```json
{
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
}
```

</details>

<details><summary>Пример 2</summary>

Запрос 2:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param/36

{
  "dotter_notation": ".1.3.6.1.2.1.1.2",
  "oid_id": "14e8713a-2f3c-3af1-8e6f-449d7a612227"
}
```

Ответ 2:

```json
{
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
}
```

</details>

<details><summary>Пример 3</summary>

Запрос 3:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param/36

{
  "dotter_notation": ".1.3.6.1.2.1.1.2"
}
```

Ответ 3:

```json
{
  "id": 36,
  "oid": null,
  "dotter_notation": ".1.3.6.1.2.1.1.2"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации JSON-тела или некорректный ID в пути  
404: Not Found - Обновляемый индикатор параметров не обнаружен в базе данных  
500: Internal Server Error - Системная ошибка СУБД при сохранении данных

---

## [DELETE] /api/v1/catalog/indicator/param/{id} - Удалить индикатор параметров по ID

> [!CAUTION]
> При удалении индикатора параметров из БД также удаляются все связанные с ним сопоставления параметров.

> [!WARNING]
> В связи с вышесказанным желательно добавить либо кнопку, либо дополнительное всплывающее сообщение подтверждения
> проведения операции.

> [!IMPORTANT]
> При удалении индикатора параметра происходит сдвиг (уменьшение значения id на 1) всего списка индикаторов, следовавших
> за удалённым, что не приводит к сбоям в работе системы.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/param/36

{}
```

Ответ 1:

```json
{}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный числовой идентификатор в пути запроса  
404: Not Found - Индикатор параметров не существует  
500: Internal Server Error - Не удалось удалить индикатор параметров из-за ошибки СУБД

---

## [GET] /api/v1/catalog/indicator/params - Получить все индикаторы параметров

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/indicator/params

{}
```

Ответ 1:

```json
[
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  }
]
```

</details>

</details>

### Возможные коды ошибок

500: Internal Server Error - Внутренняя ошибка сервера при чтении индикаторов параметров

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
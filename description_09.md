# 9. Конфигурация: Связь компонентов устройства и сопоставлений параметров

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/catalog/link/device-component-mapping/{deviceComponentId}/{mappingId} - Связать составную часть устройства с сопоставлением параметра

> [!TIP]
> API предназначен для связывания составных частей и сопоставлений параметров по их id: первым указывается id составной 
> части, вторым - id сопоставления параметра.  
> В результате выполнения запроса будет получена обновлённая структура устройства.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/link/device-component-mapping/1/1

{}
```

Ответ 1:

```json
{
  "id": 1,
  "component": {
    "id": 1,
    "title": "component",
    "name_en": "Component",
    "name_ru": "Компонент"
  },
  "internal_order": 1,
  "parent": null,
  "mappings": [
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
    }
  ],
  "children": []
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректные форматы переданных идентификаторов в пути запроса  
404: Not Found - Составная часть устройства или сопоставление параметра не найдены  
500: Internal Server Error - Ошибка СУБД при записи связи в ассоциативную таблицу

---

## [DELETE] /api/v1/catalog/link/device-component-mapping/{deviceComponentId}/{mappingId} - Удалить связь составной части устройства с сопоставлением параметра

> [!TIP]
> API предназначен для отвязывания сопоставлений параметров от составной части по их id: первым указывается id составной 
> части, вторым - id сопоставления параметра.  
> В результате выполнения запроса будет получена обновлённая структура устройства.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/link/device-component-mapping/1/1

{}
```

Ответ 1:

```json
{
  "id": 1,
  "component": {
    "id": 1,
    "title": "component",
    "name_en": "Component",
    "name_ru": "Компонент"
  },
  "internal_order": 1,
  "parent": null,
  "mappings": [
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
    }
  ],
  "children": []
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный числовой формат переданных идентификаторов в пути  
404: Not Found - Запись о связи составной части с сопоставлением не обнаружена  
500: Internal Server Error - Системная ошибка базы данных при удалении связи

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
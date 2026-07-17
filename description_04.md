# 4. Парсер: OID

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

## [POST] /api/v1/mib-parser/oid - Поиск OID по точной dotter notation

> [!TIP]
> API предназначен для поиска OID по точечной нотации, части названия производителя, точному названию MIB и их
> комбинациям и поддерживает постраничный вывод данных.  
> По умолчанию (если значения соответствующих полей не заданы) выводится первая страница с результатами, а размер
> страницы устанавливается равным 100 записям на страницу.

> [!NOTE]
> Точечная нотация OID может быть описана как в стандартных MIB, относящихся к той или иной версии протокола SNMP, так и
> в MIB самого производителя.  
> API при заданном значении фильтра названия производителя учитывает данный факт, поэтому в результатах поиска могут
> присутствовать OID, описание которыъ приведено как в MIB производителя, так и в стандартных MIB.

> [!IMPORTANT]
> Ни одно из полей данного API не является обязательным.

Передаваемый в API JSON-объект может содержать следующие поля:

- `dotter_notation` - точечная нотация OID, по которой необходимо осуществить поиск (если не будет передана, то поиск
  будет осуществлён по всем точечным нотациям);
- `mib` - точное название MIB (не путать с названием файла, название MIB указано в его содержании);
- `page` - номер страницы с результатами (результаты строго разбиты по страницам, поэтому при запросе страницы 2 и
  последующих страниц возможно получение пустого массива данных по причине недостатка в количестве полученных записей
  для их отображения на этих страницах) - по умолчанию равно 1;
- `page_size` - размер страницы - по умолчанию равно 100;
- `prefix` - флаг префикса: если true - в результатах поиска будут выведены все OID, точечная нотация которых начинается
  с переданного значения `dotter_notation`, если false - точечная нотация которых строго равна переданному значению
  `dotter_notation`;
- `vendor` - часть названия производителя.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "dotter_notation": ".1.3.6.1.2.1.1.2",
  "mib": "SNMPv2-MIB",
  "page": 1,
  "page_size": 100,
  "prefix": true,
  "vendor": "Raisecom"
}
```

Ответ 1:

```json
[
  {
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
  }
]
```

Запрос 2:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "dotter_notation": ".1.3.6.1.2.1.1.2"
}
```

Ответ 2:

```json
[
  {
    "id": "9401065c-6d89-30f5-920f-b276b4a6a0ca",
    "mib": {
      "id": 4888,
      "path": "RFC1213-MIB.mib",
      "name": "RFC1213-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysObjectID",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "mandatory",
    "access": "read-only",
    "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.4242, it could assign the identifier 1.3.6.1.4.1.4242.1.1 to its `Fred Router'.",
    "category": "system"
  },
  {
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
  {
    "id": "488814ee-190f-3877-b0d1-e4717699b8e0",
    "mib": {
      "id": 2099,
      "path": "tplink/RFC1213-MIB.mib",
      "name": "RFC1213-MIB",
      "vendor": "TP-Link Systems Inc."
    },
    "type": "OBJECT-TYPE",
    "name": "sysObjectID",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "mandatory",
    "access": "read-only",
    "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.4242, it could assign the identifier 1.3.6.1.4.1.4242.1.1 to its `Fred Router'.",
    "category": "system"
  },
  {
    "id": "6fdbbeff-3922-304a-806c-75107305e8aa",
    "mib": {
      "id": 1963,
      "path": "junos/JNX-SNMPv2-CAPABILITY.mib",
      "name": "JNX-SNMPv2-CAPABILITY",
      "vendor": "ESO Consortium"
    },
    "type": "OBJECT-TYPE",
    "name": "sysObjectID",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "mandatory",
    "access": "read-only",
    "description": "It is the assigned identifier to represent platform name. For example enterprises. 2636.1.1.1.2.21 to jnxProductNameMX960",
    "category": "system"
  },
  {
    "id": "696d4e2f-3622-3e54-b3f9-6a30a54e9849",
    "mib": {
      "id": 6,
      "path": "janitza/JANITZA-MIB.mib",
      "name": "JANITZA-MIB",
      "vendor": "Janitza electronics GmbH"
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
  {
    "id": "9f924940-8c6b-326d-b197-801d09aca7db",
    "mib": {
      "id": 7,
      "path": "janitza/JANITZA-MIB-UMG96.mib",
      "name": "JANITZA-MIB-UMG96",
      "vendor": "Janitza electronics GmbH"
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
  {
    "id": "a77918df-7331-36ba-b534-62317cd98edd",
    "mib": {
      "id": 980,
      "path": "cdata/RFC1213-MIB.mib",
      "name": "RFC1213-MIB",
      "vendor": "Shenzhen C-Data Technology Co.,Ltd."
    },
    "type": "OBJECT-TYPE",
    "name": "sysObjectID",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "mandatory",
    "access": "read-only",
    "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.4242, it could assign the identifier 1.3.6.1.4.1.4242.1.1 to its `Fred Router'.",
    "category": "system"
  }
]
```

Запрос 3:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "mib": "SNMPv2-MIB"
}
```

Ответ 3:

```json
[
  {
    "id": "45786ee6-c7be-3c1b-bb48-3e2ae81d265c",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "system",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system",
    "status": null,
    "access": null,
    "category": "mib-2"
  },
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
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
  {
    "id": "2f1687bd-b832-3fc9-8246-bfb689fd818b",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysORLastChange",
    "number": 8,
    "dotter_notation": ".1.3.6.1.2.1.1.8",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORLastChange",
    "syntax": "TimeTicks",
    "status": "current",
    "access": "read-only",
    "description": "The value of sysUpTime at the time of the most recent change in state or value of any instance of sysORID.",
    "category": "system"
  },
  {
    "id": "0bd3558c-98bc-3365-b827-45cdc9643da9",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysORTable",
    "number": 9,
    "dotter_notation": ".1.3.6.1.2.1.1.9",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The (conceptual) table listing the capabilities of the local SNMP application acting as a command responder with respect to various MIB modules. SNMP entities having dynamically-configurable support of MIB modules will have a dynamically-varying number of conceptual rows.",
    "category": "system"
  },
  {
    "id": "adc65055-9c7a-31de-ae0c-82aa9eefe040",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysOREntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.1.9.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORTable.sysOREntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry (conceptual row) in the sysORTable.",
    "category": "sysORTable"
  },
  {
    "id": "e879334e-d45a-370d-82b3-4c867938138c",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysORIndex",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.1.9.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORTable.sysOREntry.sysORIndex",
    "syntax": "INTEGER (1... 2147483647)",
    "status": "current",
    "access": "not-accessible",
    "description": "The auxiliary variable used for identifying instances of the columnar objects in the sysORTable.",
    "category": "sysOREntry"
  },
  {
    "id": "5907eff0-3319-3039-b723-2cae65c6bdda",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysORID",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.1.9.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORTable.sysOREntry.sysORID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-only",
    "description": "An authoritative identification of a capabilities statement with respect to various MIB modules supported by the local SNMP application acting as a command responder.",
    "category": "sysOREntry"
  },
  {
    "id": "205c3c50-3762-38f8-ad2b-65169b301ecc",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysORDescr",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.1.9.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORTable.sysOREntry.sysORDescr",
    "syntax": "OCTET STRING (SIZE (0... 255))",
    "status": "current",
    "access": "read-only",
    "description": "A textual description of the capabilities identified by the corresponding instance of sysORID.",
    "category": "sysOREntry"
  },
  {
    "id": "9e2fc807-d76b-38a1-bdaa-96c1aa29d62b",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysORUpTime",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.1.9.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysORTable.sysOREntry.sysORUpTime",
    "syntax": "TimeTicks",
    "status": "current",
    "access": "read-only",
    "description": "The value of sysUpTime at the time this conceptual row was last instantiated.",
    "category": "sysOREntry"
  },
  {
    "id": "fb62cc16-91e1-39ad-9f59-d0156a11526e",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmp",
    "number": 11,
    "dotter_notation": ".1.3.6.1.2.1.11",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp",
    "status": null,
    "access": null,
    "category": "mib-2"
  },
  {
    "id": "8fe67dda-98e7-3f57-a9d0-36f013f0b501",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInPkts",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.11.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInPkts",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of messages delivered to the SNMP entity from the transport service.",
    "category": "snmp"
  },
  {
    "id": "c3768931-5244-3e51-b568-dca127209303",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutPkts",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.11.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutPkts",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Messages which were passed from the SNMP protocol entity to the transport service.",
    "category": "snmp"
  },
  {
    "id": "c59826e4-e7c3-31c6-8c50-7e62feb61e8a",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInBadVersions",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.11.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInBadVersions",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of SNMP messages which were delivered to the SNMP entity and were for an unsupported SNMP version.",
    "category": "snmp"
  },
  {
    "id": "525e0b99-72dd-3e91-974f-3aca1f6f338e",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInBadCommunityNames",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.11.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInBadCommunityNames",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of community-based SNMP messages (for example, SNMPv1) delivered to the SNMP entity which used an SNMP community name not known to said entity. Also, implementations which authenticate community-based SNMP messages using check (s) in addition to matching the community name (for example, by also checking whether the message originated from a transport address allowed to use a specified community name) MAY include in this value the number of messages which failed the additional check (s). It is strongly RECOMMENDED that the documentation for any security model which is used to authenticate community-based SNMP messages specify the precise conditions that contribute to this value.",
    "category": "snmp"
  },
  {
    "id": "3b6a1de3-c8c6-3536-b31e-3e1038609c2e",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInBadCommunityUses",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.11.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInBadCommunityUses",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of community-based SNMP messages (for example, SNMPv1) delivered to the SNMP entity which represented an SNMP operation that was not allowed for the SNMP community named in the message. The precise conditions under which this counter is incremented (if at all) depend on how the SNMP entity implements its access control mechanism and how its applications interact with that access control mechanism. It is strongly RECOMMENDED that the documentation for any access control mechanism which is used to control access to and visibility of MIB instrumentation specify the precise conditions that contribute to this value.",
    "category": "snmp"
  },
  {
    "id": "e696a664-7887-3e69-acea-375b511c35eb",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInASNParseErrs",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.11.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInASNParseErrs",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of ASN. 1 or BER errors encountered by the SNMP entity when decoding received SNMP messages.",
    "category": "snmp"
  },
  {
    "id": "19dd4b5a-b58c-352b-b162-51834b8ce5e2",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInTooBigs",
    "number": 8,
    "dotter_notation": ".1.3.6.1.2.1.11.8",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInTooBigs",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were delivered to the SNMP protocol entity and for which the value of the error-status field was `tooBig'.",
    "category": "snmp"
  },
  {
    "id": "b6ff6c1f-033f-3868-a0f9-4a20efe0f60e",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInNoSuchNames",
    "number": 9,
    "dotter_notation": ".1.3.6.1.2.1.11.9",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInNoSuchNames",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were delivered to the SNMP protocol entity and for which the value of the error-status field was `noSuchName'.",
    "category": "snmp"
  },
  {
    "id": "0e0ae856-9d97-3e75-abcf-2a1359ecb3d9",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInBadValues",
    "number": 10,
    "dotter_notation": ".1.3.6.1.2.1.11.10",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInBadValues",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were delivered to the SNMP protocol entity and for which the value of the error-status field was `badValue'.",
    "category": "snmp"
  },
  {
    "id": "3f04421c-777e-32ac-94b3-2b0861b01c24",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInReadOnlys",
    "number": 11,
    "dotter_notation": ".1.3.6.1.2.1.11.11",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInReadOnlys",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number valid SNMP PDUs which were delivered to the SNMP protocol entity and for which the value of the error-status field was `readOnly'. It should be noted that it is a protocol error to generate an SNMP PDU which contains the value `readOnly' in the error-status field, as such this object is provided as a means of detecting incorrect implementations of the SNMP.",
    "category": "snmp"
  },
  {
    "id": "0cbafb0d-8f7e-3ae7-9130-0b12d7fbe31d",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInGenErrs",
    "number": 12,
    "dotter_notation": ".1.3.6.1.2.1.11.12",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInGenErrs",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were delivered to the SNMP protocol entity and for which the value of the error-status field was `genErr'.",
    "category": "snmp"
  },
  {
    "id": "6149c2ac-254b-35e3-a479-3cb776bc2901",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInTotalReqVars",
    "number": 13,
    "dotter_notation": ".1.3.6.1.2.1.11.13",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInTotalReqVars",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of MIB objects which have been retrieved successfully by the SNMP protocol entity as the result of receiving valid SNMP Get-Request and Get-Next PDUs.",
    "category": "snmp"
  },
  {
    "id": "967f0e76-0e2f-3f68-b9aa-d7d0584dc8ba",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInTotalSetVars",
    "number": 14,
    "dotter_notation": ".1.3.6.1.2.1.11.14",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInTotalSetVars",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of MIB objects which have been altered successfully by the SNMP protocol entity as the result of receiving valid SNMP Set-Request PDUs.",
    "category": "snmp"
  },
  {
    "id": "b148011d-a8ff-3773-b404-1f5d0763997b",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInGetRequests",
    "number": 15,
    "dotter_notation": ".1.3.6.1.2.1.11.15",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInGetRequests",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Get-Request PDUs which have been accepted and processed by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "3a835510-1ba3-3214-8fe1-7500a4d73c62",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInGetNexts",
    "number": 16,
    "dotter_notation": ".1.3.6.1.2.1.11.16",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInGetNexts",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Get-Next PDUs which have been accepted and processed by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "e792cee4-9be2-347b-a7bc-d51f2355bb28",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInSetRequests",
    "number": 17,
    "dotter_notation": ".1.3.6.1.2.1.11.17",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInSetRequests",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Set-Request PDUs which have been accepted and processed by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "8245ddd4-710b-3874-9f30-d562a9c94a17",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInGetResponses",
    "number": 18,
    "dotter_notation": ".1.3.6.1.2.1.11.18",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInGetResponses",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Get-Response PDUs which have been accepted and processed by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "b85a8752-edef-379d-95f0-d67899130e45",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpInTraps",
    "number": 19,
    "dotter_notation": ".1.3.6.1.2.1.11.19",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpInTraps",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Trap PDUs which have been accepted and processed by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "000948b2-b66c-3754-9de4-ee25a96a0225",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutTooBigs",
    "number": 20,
    "dotter_notation": ".1.3.6.1.2.1.11.20",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutTooBigs",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were generated by the SNMP protocol entity and for which the value of the error-status field was `tooBig. '",
    "category": "snmp"
  },
  {
    "id": "85af6abf-6f41-3d10-856e-7eef68d9d0c1",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutNoSuchNames",
    "number": 21,
    "dotter_notation": ".1.3.6.1.2.1.11.21",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutNoSuchNames",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were generated by the SNMP protocol entity and for which the value of the error-status was `noSuchName'.",
    "category": "snmp"
  },
  {
    "id": "53988f80-1421-3c27-ba04-440a36c1a1db",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutBadValues",
    "number": 22,
    "dotter_notation": ".1.3.6.1.2.1.11.22",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutBadValues",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were generated by the SNMP protocol entity and for which the value of the error-status field was `badValue'.",
    "category": "snmp"
  },
  {
    "id": "a2559a42-671d-3986-8ed2-0864b19a537d",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutGenErrs",
    "number": 24,
    "dotter_notation": ".1.3.6.1.2.1.11.24",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutGenErrs",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP PDUs which were generated by the SNMP protocol entity and for which the value of the error-status field was `genErr'.",
    "category": "snmp"
  },
  {
    "id": "edbb7ea9-199d-3de1-b44c-3a810fa26f1f",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutGetRequests",
    "number": 25,
    "dotter_notation": ".1.3.6.1.2.1.11.25",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutGetRequests",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Get-Request PDUs which have been generated by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "c834c0cd-d811-3dd3-a206-076755bfc3e9",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutGetNexts",
    "number": 26,
    "dotter_notation": ".1.3.6.1.2.1.11.26",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutGetNexts",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Get-Next PDUs which have been generated by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "94005eed-db11-38db-8d9d-dbb325cbe14a",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutSetRequests",
    "number": 27,
    "dotter_notation": ".1.3.6.1.2.1.11.27",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutSetRequests",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Set-Request PDUs which have been generated by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "bd36815b-375f-38ea-aece-ccafab581c08",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutGetResponses",
    "number": 28,
    "dotter_notation": ".1.3.6.1.2.1.11.28",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutGetResponses",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Get-Response PDUs which have been generated by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "4017bbd8-ba89-34eb-987c-6945e6eaa209",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpOutTraps",
    "number": 29,
    "dotter_notation": ".1.3.6.1.2.1.11.29",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpOutTraps",
    "syntax": "Counter32",
    "status": "obsolete",
    "access": "read-only",
    "description": "The total number of SNMP Trap PDUs which have been generated by the SNMP protocol entity.",
    "category": "snmp"
  },
  {
    "id": "a190817d-6258-3bb2-8801-7e07fc5446a2",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpEnableAuthenTraps",
    "number": 30,
    "dotter_notation": ".1.3.6.1.2.1.11.30",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpEnableAuthenTraps",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-write",
    "description": "Indicates whether the SNMP entity is permitted to generate authenticationFailure traps. The value of this object overrides any configuration information; as such, it provides a means whereby all authenticationFailure traps may be disabled. Note that it is strongly recommended that this object be stored in non-volatile memory so that it remains constant across re-initializations of the network management system.",
    "category": "snmp"
  },
  {
    "id": "d368bdcf-978e-3b38-8d12-96030a0f3227",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpSilentDrops",
    "number": 31,
    "dotter_notation": ".1.3.6.1.2.1.11.31",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpSilentDrops",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of Confirmed Class PDUs (such as GetRequest-PDUs, GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and InformRequest-PDUs) delivered to the SNMP entity which were silently dropped because the size of a reply containing an alternate Response Class PDU (such as a Response-PDU) with an empty variable-bindings field was greater than either a local constraint or the maximum message size associated with the originator of the request.",
    "category": "snmp"
  },
  {
    "id": "e0a45efe-c484-35f0-bcc4-a015717d8f35",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpProxyDrops",
    "number": 32,
    "dotter_notation": ".1.3.6.1.2.1.11.32",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.snmp.snmpProxyDrops",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "description": "The total number of Confirmed Class PDUs (such as GetRequest-PDUs, GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and InformRequest-PDUs) delivered to the SNMP entity which were silently dropped because the transmission of the (possibly translated) message to a proxy target failed in a manner (other than a time-out) such that no Response Class PDU (such as a Response-PDU) could be returned.",
    "category": "snmp"
  },
  {
    "id": "599edfff-d1d9-38ec-8350-2b553ec6d770",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "MODULE-IDENTITY",
    "name": "snmpMIB",
    "number": 1,
    "dotter_notation": ".1.3.6.1.6.3.1",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB",
    "status": null,
    "access": null,
    "description": "The MIB module for SNMP entities. Copyright (C) The Internet Society (2002). This version of this MIB module is part of RFC 3418; see the RFC itself for full legal notices.",
    "category": "snmpModules"
  },
  {
    "id": "b0fe8c8f-5ed3-3d3f-bdf2-8efeb8d7d2d8",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpMIBObjects",
    "number": 1,
    "dotter_notation": ".1.3.6.1.6.3.1.1",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects",
    "status": null,
    "access": null,
    "category": "snmpMIB"
  },
  {
    "id": "e6dd8268-ddf1-31dc-aeba-266634602d0d",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpTrap",
    "number": 4,
    "dotter_notation": ".1.3.6.1.6.3.1.1.4",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTrap",
    "status": null,
    "access": null,
    "category": "snmpMIBObjects"
  },
  {
    "id": "9375fd8b-3cd0-3c7f-886b-6208e948b237",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpTrapOID",
    "number": 1,
    "dotter_notation": ".1.3.6.1.6.3.1.1.4.1",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTrap.snmpTrapOID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "accessible-for-notify",
    "description": "The authoritative identification of the notification currently being sent. This variable occurs as the second varbind in every SNMPv2-Trap-PDU and InformRequest-PDU.",
    "category": "snmpTrap"
  },
  {
    "id": "4919a05c-2807-32e8-a8e8-cd1e8f044be6",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpTrapEnterprise",
    "number": 3,
    "dotter_notation": ".1.3.6.1.6.3.1.1.4.3",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTrap.snmpTrapEnterprise",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "accessible-for-notify",
    "description": "The authoritative identification of the enterprise associated with the trap currently being sent. When an SNMP proxy agent is mapping an RFC1157 Trap-PDU into a SNMPv2-Trap-PDU, this variable occurs as the last varbind.",
    "category": "snmpTrap"
  },
  {
    "id": "09e17035-ed97-3b35-b5a9-e167130f5018",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpTraps",
    "number": 5,
    "dotter_notation": ".1.3.6.1.6.3.1.1.5",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTraps",
    "status": null,
    "access": null,
    "category": "snmpMIBObjects"
  },
  {
    "id": "0f100fb1-b1ee-3d9b-bc06-3d52700a2e29",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "NOTIFICATION-TYPE",
    "name": "coldStart",
    "number": 1,
    "dotter_notation": ".1.3.6.1.6.3.1.1.5.1",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTraps.coldStart",
    "status": "current",
    "access": null,
    "description": "A coldStart trap signifies that the SNMP entity, supporting a notification originator application, is reinitializing itself and that its configuration may have been altered.",
    "category": "snmpTraps"
  },
  {
    "id": "75754f4a-b342-3987-8ffd-b7216823463e",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "NOTIFICATION-TYPE",
    "name": "warmStart",
    "number": 2,
    "dotter_notation": ".1.3.6.1.6.3.1.1.5.2",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTraps.warmStart",
    "status": "current",
    "access": null,
    "description": "A warmStart trap signifies that the SNMP entity, supporting a notification originator application, is reinitializing itself such that its configuration is unaltered.",
    "category": "snmpTraps"
  },
  {
    "id": "45c61af3-a31f-333b-a397-2afedc6d5365",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "NOTIFICATION-TYPE",
    "name": "authenticationFailure",
    "number": 5,
    "dotter_notation": ".1.3.6.1.6.3.1.1.5.5",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpTraps.authenticationFailure",
    "status": "current",
    "access": null,
    "description": "An authenticationFailure trap signifies that the SNMP entity has received a protocol message that is not properly authenticated. While all implementations of SNMP entities MAY be capable of generating this trap, the snmpEnableAuthenTraps object indicates whether this trap will be generated.",
    "category": "snmpTraps"
  },
  {
    "id": "e19b448f-b874-3ac9-b527-0440eb7c139c",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpSet",
    "number": 6,
    "dotter_notation": ".1.3.6.1.6.3.1.1.6",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpSet",
    "status": null,
    "access": null,
    "category": "snmpMIBObjects"
  },
  {
    "id": "fb83e4fb-8fd0-3656-83b0-d7df41367104",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "snmpSetSerialNo",
    "number": 1,
    "dotter_notation": ".1.3.6.1.6.3.1.1.6.1",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBObjects.snmpSet.snmpSetSerialNo",
    "syntax": "INTEGER (0... 2147483647)",
    "status": "current",
    "access": "read-write",
    "description": "An advisory lock used to allow several cooperating command generator applications to coordinate their use of the SNMP set operation. This object is used for coarse-grain coordination. To achieve fine-grain coordination, one or more similar objects might be defined within each MIB group, as appropriate.",
    "category": "snmpSet"
  },
  {
    "id": "09f2c892-c6b2-3239-8190-7462485fbd7d",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpMIBConformance",
    "number": 2,
    "dotter_notation": ".1.3.6.1.6.3.1.2",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance",
    "status": null,
    "access": null,
    "category": "snmpMIB"
  },
  {
    "id": "31c00025-43eb-390d-8928-0092a7156884",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpMIBCompliances",
    "number": 1,
    "dotter_notation": ".1.3.6.1.6.3.1.2.1",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBCompliances",
    "status": null,
    "access": null,
    "category": "snmpMIBConformance"
  },
  {
    "id": "ffa1153d-70bd-38bc-9dd3-6b9e4f4ba919",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "MODULE-COMPLIANCE",
    "name": "snmpBasicCompliance",
    "number": 2,
    "dotter_notation": ".1.3.6.1.6.3.1.2.1.2",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBCompliances.snmpBasicCompliance",
    "status": "deprecated",
    "access": null,
    "description": "The compliance statement for SNMPv2 entities which implement the SNMPv2 MIB. This compliance statement is replaced by snmpBasicComplianceRev2.",
    "category": "snmpMIBCompliances"
  },
  {
    "id": "ddc82926-625b-3b04-9118-f728538094d9",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "MODULE-COMPLIANCE",
    "name": "snmpBasicComplianceRev2",
    "number": 3,
    "dotter_notation": ".1.3.6.1.6.3.1.2.1.3",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBCompliances.snmpBasicComplianceRev2",
    "status": "current",
    "access": null,
    "description": "The compliance statement for SNMP entities which implement this MIB module.",
    "category": "snmpMIBCompliances"
  },
  {
    "id": "4e55ad0e-76d8-37fd-ac30-87039dca7f44",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "snmpMIBGroups",
    "number": 2,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups",
    "status": null,
    "access": null,
    "category": "snmpMIBConformance"
  },
  {
    "id": "013fbdbb-20e4-3501-82fc-851d3d4b8bb1",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-GROUP",
    "name": "snmpSetGroup",
    "number": 5,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.5",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpSetGroup",
    "status": "current",
    "access": null,
    "description": "A collection of objects which allow several cooperating command generator applications to coordinate their use of the set operation.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "bf6101f1-6c9f-30be-9fd7-38a1f703b36f",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-GROUP",
    "name": "systemGroup",
    "number": 6,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.6",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.systemGroup",
    "status": "current",
    "access": null,
    "description": "The system group defines objects which are common to all managed systems.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "e2cf157e-0414-3830-ab7c-d42c7983da47",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "NOTIFICATION-GROUP",
    "name": "snmpBasicNotificationsGroup",
    "number": 7,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.7",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpBasicNotificationsGroup",
    "status": "current",
    "access": null,
    "description": "The basic notifications implemented by an SNMP entity supporting command responder applications.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "580d9d19-ce29-31c2-8a1c-3af3a289eafd",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-GROUP",
    "name": "snmpGroup",
    "number": 8,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.8",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpGroup",
    "status": "current",
    "access": null,
    "description": "A collection of objects providing basic instrumentation and control of an SNMP entity.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "09a86340-d50a-3222-ac4f-517d6c5f76b5",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-GROUP",
    "name": "snmpCommunityGroup1",
    "number": 9,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.9",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpCommunityGroup1",
    "status": "current",
    "access": null,
    "description": "A collection of objects providing basic instrumentation of a SNMP entity which supports community-based authentication.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "f3c0c225-370e-394f-bdee-e9f9a5b2a80c",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-GROUP",
    "name": "snmpObsoleteGroup",
    "number": 10,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.10",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpObsoleteGroup",
    "status": "obsolete",
    "access": null,
    "description": "A collection of objects from RFC 1213 made obsolete by this MIB module.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "151b1514-a112-36ee-a1c0-06fbdc59658a",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "NOTIFICATION-GROUP",
    "name": "snmpWarmStartNotificationGroup",
    "number": 11,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.11",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpWarmStartNotificationGroup",
    "status": "current",
    "access": null,
    "description": "An additional notification for an SNMP entity supporting command responder applications, if it is able to reinitialize itself such that its configuration is unaltered.",
    "category": "snmpMIBGroups"
  },
  {
    "id": "d10a71a8-b0b2-32a5-8ec7-8712102b882f",
    "mib": {
      "id": 4912,
      "path": "SNMPv2-MIB.mib",
      "name": "SNMPv2-MIB",
      "vendor": null
    },
    "type": "OBJECT-GROUP",
    "name": "snmpNotificationGroup",
    "number": 12,
    "dotter_notation": ".1.3.6.1.6.3.1.2.2.12",
    "object_descriptor": ".iso.org.dod.internet.snmpV2.snmpModules.snmpMIB.snmpMIBConformance.snmpMIBGroups.snmpNotificationGroup",
    "status": "current",
    "access": null,
    "description": "These objects are required for entities which support notification originator applications.",
    "category": "snmpMIBGroups"
  }
]
```

Запрос 4:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "page": 1,
  "page_size": 10
}
```

Ответ 4:

```json
[
  {
    "id": "625f76be-96b0-349d-a14b-9e9be2d4c51e",
    "mib": {
      "id": 4273,
      "path": "ADSL2-LINE-TC-MIB.mib",
      "name": "ADSL2-LINE-TC-MIB",
      "vendor": null
    },
    "type": "MODULE-IDENTITY",
    "name": "adsl2TCMIB",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.10.238.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.238.adsl2TCMIB",
    "status": null,
    "access": null,
    "description": "This MIB Module provides Textual Conventions to be used by the ADSL2-LINE-MIB module for the purpose of managing ADSL, ADSL2, and ADSL2+ lines. Copyright (C) The Internet Society (2006). This version of this MIB module is part of RFC 4706: see the RFC itself for full legal notices.",
    "category": "transmission"
  },
  {
    "id": "aac38345-fb98-305d-8237-99289b4629a8",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "MODULE-IDENTITY",
    "name": "adslExtMIB",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB",
    "status": null,
    "access": null,
    "description": "Copyright (C) The Internet Society (2002). This version of this MIB module is part of RFC 3440; see the RFC itself for full legal notices. This MIB Module is a supplement to the ADSL-LINE-MIB [RFC2662].",
    "category": "adslMIB"
  },
  {
    "id": "2f6124a3-51d7-3463-a84f-3a446e21bfe1",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "adslExtMibObjects",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects",
    "status": null,
    "access": null,
    "category": "adslExtMIB"
  },
  {
    "id": "f0703500-f4ed-3288-b30b-903f880f9fb5",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineExtTable",
    "number": 17,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "This table is an extension of RFC 2662. It contains ADSL line configuration and monitoring information. This includes the ADSL line's capabilities and actual ADSL transmission system.",
    "category": "adslExtMibObjects"
  },
  {
    "id": "235e26f9-eaf0-3eb4-9ea1-394a692c1a0c",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineExtEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable.adslLineExtEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry extends the adslLineEntry defined in [RFC2662]. Each entry corresponds to an ADSL line.",
    "category": "adslLineExtTable"
  },
  {
    "id": "0af68949-143e-301d-95c2-fa9853c3a4c1",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineTransAtucCap",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable.adslLineExtEntry.adslLineTransAtucCap",
    "syntax": "BITS",
    "status": "current",
    "access": "read-only",
    "description": "The transmission modes, represented by a bitmask that the ATU-C is capable of supporting. The modes available are limited by the design of the equipment.",
    "category": "adslLineExtEntry"
  },
  {
    "id": "b56d7437-061c-3fc3-b74c-240b56079b5d",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineTransAtucConfig",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable.adslLineExtEntry.adslLineTransAtucConfig",
    "syntax": "BITS",
    "status": "current",
    "access": "read-write",
    "description": "The transmission modes, represented by a bitmask, currently enabled by the ATU-C. The manager can only set those modes that are supported by the ATU-C. An ATU-C's supported modes are provided by AdslLineTransAtucCap.",
    "category": "adslLineExtEntry"
  },
  {
    "id": "da4e972b-65e9-3f2b-aa15-c9891ef41844",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineTransAtucActual",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable.adslLineExtEntry.adslLineTransAtucActual",
    "syntax": "BITS",
    "status": "current",
    "access": "read-only",
    "description": "The actual transmission mode of the ATU-C. During ADSL line initialization, the ADSL Transceiver Unit - Remote terminal end (ATU-R) will determine the mode used for the link. This value will be limited a single transmission mode that is a subset of those modes enabled by the ATU-C and denoted by adslLineTransAtucConfig. After an initialization has occurred, its mode is saved as the 'Current' mode and is persistence should the link go down. This object returns 0 (i.e. BITS with no mode bit set) if the mode is not known.",
    "category": "adslLineExtEntry"
  },
  {
    "id": "988c2628-db73-39d4-8a62-0b7a76e44dac",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineGlitePowerState",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable.adslLineExtEntry.adslLineGlitePowerState",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-only",
    "description": "The value of this object specifies the power state of this interface. L0 is power on, L1 is power on but reduced and L3 is power off. Power state cannot be configured by an operator but it can be viewed via the ifOperStatus object for the managed ADSL interface. The value of the object ifOperStatus is set to down (2) if the ADSL interface is in power state L3 and is set to up (1) if the ADSL line interface is in power state L0 or L1. If the object adslLineTransAtucActual is set to a G. 992.2 (G.Lite) -type transmission mode, the value of this object will be one of the valid power states: L0 (2), L1 (3), or L3 (4). Otherwise, its value will be none (1).",
    "category": "adslLineExtEntry"
  },
  {
    "id": "89a2938f-b462-3486-a98d-ce08b1711653",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslLineConfProfileDualLite",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.17.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslLineExtTable.adslLineExtEntry.adslLineConfProfileDualLite",
    "syntax": "OCTET STRING (SIZE (0... 255))",
    "status": "current",
    "access": "read-write",
    "description": "This object extends the definition an ADSL line and associated channels (when applicable) for cases when it is configured in dual mode, and operating in a G.Lite-type mode as denoted by adslLineTransAtucActual. Dual mode exists when the object, adslLineTransAtucConfig, is configured with one or more full-rate modes and one or more G.Lite modes simultaneously. When 'dynamic' profiles are implemented, the value of object is equal to the index of the applicable row in the ADSL Line Configuration Profile Table, AdslLineConfProfileTable defined in ADSL-MIB [RFC2662]. In the case when dual-mode has not been enabled, the value of the object will be equal to the value of the object adslLineConfProfile [RFC2662]. When `static' profiles are implemented, in much like the case of the object, adslLineConfProfileName [RFC2662], this object's value will need to algorithmically represent the characteristics of the line. In this case, the value of the line's ifIndex plus a value indicating the line mode type (e.g., G.Lite, Full-rate) will be used. Therefore, the profile's name is a string concatenating the ifIndex and one of the follow values: Full or Lite. This string will be fixed-length (i.e., 14) with leading zero (s). For example, the profile name for ifIndex that equals '15' and is a full rate line, it will be '0000000015Full'.",
    "category": "adslLineExtEntry"
  }
]
```

Запрос 5:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "page": 2,
  "page_size": 20,
  "vendor": "Raisecom"
}
```

Ответ 5:

```json
[
  {
    "id": "ff18138d-40b7-3a99-990d-07428ec3f998",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfCurr1DayFastR",
    "number": 9,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.9",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfCurr1DayFastR",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current day as measured by adslAtucPerfCurr1DayTimeElapsed [RFC2662], adslAtucPerfCurr1DayFastR reports the number of seconds during which there have been fast retrains.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "57bd313d-5f63-3239-a1a1-9d8d85de0316",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfCurr1DayFailedFastR",
    "number": 10,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.10",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfCurr1DayFailedFastR",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current day as measured by adslAtucPerfCurr1DayTimeElapsed [RFC2662], adslAtucPerfCurr1DayFailedFastR reports the number of seconds during which there have been failed fast retrains.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "2044a92b-ffd3-3d1e-b859-66c18f6f9c35",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfCurr1DaySesL",
    "number": 11,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.11",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfCurr1DaySesL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current day as measured by adslAtucPerfCurr1DayTimeElapsed [RFC2662], adslAtucPerfCurr1DaySesL reports the number of seconds during which there have been severely errored seconds-line.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "d957ce98-9e43-3718-8893-f2888956dda4",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfCurr1DayUasL",
    "number": 12,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.12",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfCurr1DayUasL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current day as measured by adslAtucPerfCurr1DayTimeElapsed [RFC2662], adslAtucPerfCurr1DayUasL reports the number of seconds during which there have been unavailable seconds-line.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "db223e4a-8e57-33a0-b3e0-c4663c94ad8b",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfPrev1DayFastR",
    "number": 13,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.13",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfPrev1DayFastR",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the previous day, adslAtucPerfPrev1DayFastR reports the number of seconds during which there were fast retrains.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "fe77b7e7-f66b-3223-ab1c-104c897fd473",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfPrev1DayFailedFastR",
    "number": 14,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.14",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfPrev1DayFailedFastR",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the previous day, adslAtucPerfPrev1DayFailedFastR reports the number of seconds during which there were failed fast retrains.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "6d847879-1a81-3679-92ac-a477f90862d6",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfPrev1DaySesL",
    "number": 15,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.15",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfPrev1DaySesL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the previous day, adslAtucPerfPrev1DaySesL reports the number of seconds during which there were severely errored seconds-line.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "74e43e7d-0306-30bd-8e3d-6d95d1ed5072",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucPerfPrev1DayUasL",
    "number": 16,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.18.1.16",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucPerfDataExtTable.adslAtucPerfDataExtEntry.adslAtucPerfPrev1DayUasL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the previous day, adslAtucPerfPrev1DayUasL reports the number of seconds during which there were unavailable seconds-line.",
    "category": "adslAtucPerfDataExtEntry"
  },
  {
    "id": "20f45274-2bd1-3e30-be84-4547943a5939",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucIntervalExtTable",
    "number": 19,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.19",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucIntervalExtTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "This table provides one row for each ATU-C performance data collection interval for ADSL physical interfaces whose IfEntries' ifType is equal to adsl (94).",
    "category": "adslExtMibObjects"
  },
  {
    "id": "48894de6-a238-3849-a396-f16c3fc385c4",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucIntervalExtEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.19.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucIntervalExtTable.adslAtucIntervalExtEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the adslAtucIntervalExtTable.",
    "category": "adslAtucIntervalExtTable"
  },
  {
    "id": "78e0d263-2188-3f2b-aca9-28fe26db413b",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucIntervalFastR",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.19.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucIntervalExtTable.adslAtucIntervalExtEntry.adslAtucIntervalFastR",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current interval, adslAtucIntervalFastR reports the current number of seconds during which there have been fast retrains.",
    "category": "adslAtucIntervalExtEntry"
  },
  {
    "id": "51b87d45-916f-3fb2-8d05-0655e21ce0f8",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucIntervalFailedFastR",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.19.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucIntervalExtTable.adslAtucIntervalExtEntry.adslAtucIntervalFailedFastR",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the each interval, adslAtucIntervalFailedFastR reports the number of seconds during which there have been failed fast retrains.",
    "category": "adslAtucIntervalExtEntry"
  },
  {
    "id": "db06c70b-4e69-3caa-bd68-fc617b9de7c8",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucIntervalSesL",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.19.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucIntervalExtTable.adslAtucIntervalExtEntry.adslAtucIntervalSesL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the each interval, adslAtucIntervalSesL reports the number of seconds during which there have been severely errored seconds-line.",
    "category": "adslAtucIntervalExtEntry"
  },
  {
    "id": "a512cf6f-0513-3e34-a713-cf762b753285",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAtucIntervalUasL",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.19.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAtucIntervalExtTable.adslAtucIntervalExtEntry.adslAtucIntervalUasL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the each interval, adslAtucIntervalUasL reports the number of seconds during which there have been unavailable seconds-line.",
    "category": "adslAtucIntervalExtEntry"
  },
  {
    "id": "94b681e4-8827-3c8f-b36f-a233c15f8fee",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAturPerfDataExtTable",
    "number": 20,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.20",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAturPerfDataExtTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "This table contains ADSL physical line counters not defined in the adslAturPerfDataTable from the ADSL-LINE-MIB [RFC2662].",
    "category": "adslExtMibObjects"
  },
  {
    "id": "17251530-b132-38f8-a141-98083e8d321c",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAturPerfDataExtEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.20.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAturPerfDataExtTable.adslAturPerfDataExtEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry extends the adslAturPerfDataEntry defined in [RFC2662]. Each entry corresponds to an ADSL line.",
    "category": "adslAturPerfDataExtTable"
  },
  {
    "id": "70b8a5fe-38b8-3137-bd04-16160c263941",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAturPerfStatSesL",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.20.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAturPerfDataExtTable.adslAturPerfDataExtEntry.adslAturPerfStatSesL",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "The value of this object reports the count of severely errored second-line since the last agent reset.",
    "category": "adslAturPerfDataExtEntry"
  },
  {
    "id": "2621faec-6ebc-368f-b7c8-3c7d0e2c5180",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAturPerfStatUasL",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.20.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAturPerfDataExtTable.adslAturPerfDataExtEntry.adslAturPerfStatUasL",
    "syntax": "Counter32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "The value of this object reports the count of unavailable seconds-line since the last agent reset.",
    "category": "adslAturPerfDataExtEntry"
  },
  {
    "id": "f351d652-c9c6-371f-8466-850290c45a50",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAturPerfCurr15MinSesL",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.20.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAturPerfDataExtTable.adslAturPerfDataExtEntry.adslAturPerfCurr15MinSesL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current 15-minute interval, adslAturPerfCurr15MinSesL reports the current number of seconds during which there have been severely errored seconds-line.",
    "category": "adslAturPerfDataExtEntry"
  },
  {
    "id": "ec96809a-a5fd-3401-915c-7bf278dc1975",
    "mib": {
      "id": 4211,
      "path": "ADSL-LINE-EXT-MIB.mib",
      "name": "ADSL-LINE-EXT-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "adslAturPerfCurr15MinUasL",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.10.94.3.1.20.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.transmission.adslMIB.adslExtMIB.adslExtMibObjects.adslAturPerfDataExtTable.adslAturPerfDataExtEntry.adslAturPerfCurr15MinUasL",
    "syntax": "Gauge32",
    "status": "current",
    "access": "read-only",
    "units": "seconds",
    "description": "For the current 15-minute interval, adslAturPerfCurr15MinUasL reports the current number of seconds during which there have been available seconds-line.",
    "category": "adslAturPerfDataExtEntry"
  }
]
```

Запрос 6:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "dotter_notation": ".1.3.6.1.2.1.1.2",
  "vendor": "Raisecom"
}
```

Ответ 6:

```json
[
  {
    "id": "9401065c-6d89-30f5-920f-b276b4a6a0ca",
    "mib": {
      "id": 4888,
      "path": "RFC1213-MIB.mib",
      "name": "RFC1213-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "sysObjectID",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.system.sysObjectID",
    "syntax": "OBJECT IDENTIFIER",
    "status": "mandatory",
    "access": "read-only",
    "description": "The vendor's authoritative identification of the network management subsystem contained in the entity. This value is allocated within the SMI enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous means for determining `what kind of box' is being managed. For example, if vendor `Flintstones, Inc. ' was assigned the subtree 1.3.6.1.4.1.4242, it could assign the identifier 1.3.6.1.4.1.4242.1.1 to its `Fred Router'.",
    "category": "system"
  },
  {
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
  }
]
```

Запрос 7:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "page": 10,
  "page_size": 100,
  "prefix": true,
  "vendor": "Raisecom"
}
```

Ответ 7:

```json
[
  {
    "id": "341202c9-ba9e-37fb-8db1-7bdbcf99e5c8",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamNextFree",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServTBParamId, or a zero to indicate that none exist.",
    "category": "diffServTBParam"
  },
  {
    "id": "ff7d86dc-171b-3ee1-823d-3cdebcb8001b",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamTable",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "This table enumerates a single set of token bucket meter parameters that a system may use to police a stream of traffic. Such meters are modeled here as having a single rate and a single burst size. Multiple entries are used when multiple rates/burst sizes are needed.",
    "category": "diffServTBParam"
  },
  {
    "id": "8a802fea-02e8-363c-8f04-dd0476555e2e",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry that describes a single set of token bucket parameters.",
    "category": "diffServTBParamTable"
  },
  {
    "id": "0ed56de4-f486-32fb-90a0-13b1fcf2229c",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Token Bucket Parameter entries. Managers obtain new values for row creation in this table by reading diffServTBParamNextFree.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "8e74faa4-d37b-382f-87b1-250aeaa8b0c1",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamType",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamType",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "The Metering algorithm associated with the Token Bucket parameters. zeroDotZero indicates this is unknown. Standard values for generic algorithms: diffServTBParamSimpleTokenBucket, diffServTBParamAvgRate, diffServTBParamSrTCMBlind, diffServTBParamSrTCMAware, diffServTBParamTrTCMBlind, diffServTBParamTrTCMAware, and diffServTBParamTswTCM are specified in this MIB as OBJECT - IDENTITYs; additional values may be further specified in other MIBs.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "1b631da1-35f6-312a-8426-620f0e349b2f",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamRate",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamRate",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "kilobits per second",
    "description": "The token-bucket rate, in kilobits per second (kbps). This attribute is used for: 1. CIR in RFC 2697 for srTCM 2. CIR and PIR in RFC 2698 for trTCM 3. CTR and PTR in RFC 2859 for TSWTCM 4. AverageRate in RFC 3290.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "6c752e91-f86d-3f5b-9d88-03580abdc407",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamBurstSize",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamBurstSize",
    "syntax": "INTEGER (0... '7FFFFFFF'h)",
    "status": "current",
    "access": "read-create",
    "units": "Bytes",
    "description": "The maximum number of bytes in a single transmission burst. This attribute is used for: 1. CBS and EBS in RFC 2697 for srTCM 2. CBS and PBS in RFC 2698 for trTCM 3. Burst Size in RFC 3290.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "b9c12bbf-0b31-3906-a89a-f64f92fc42b6",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamInterval",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamInterval",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "microseconds",
    "description": "The time interval used with the token bucket. For: 1. Average Rate Meter, the Informal Differentiated Services Model section 5.2.1, - Delta. 2. Simple Token Bucket Meter, the Informal Differentiated Services Model section 5.1, - time interval t. 3. RFC 2859 TSWTCM, - AVG_INTERVAL. 4. RFC 2697 srTCM, RFC 2698 trTCM, - token bucket update time interval.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "65edf03c-87e8-3fde-9aa3-b51e292d4143",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamStorage",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "e78d46de-718c-3d4f-ba0e-b2a749f8d839",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServTBParamStatus",
    "number": 7,
    "dotter_notation": ".1.3.6.1.2.1.97.1.4.2.1.7",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServTBParam.diffServTBParamTable.diffServTBParamEntry.diffServTBParamStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServTBParamEntry"
  },
  {
    "id": "5f049d26-be8e-36cd-88a2-ba169e045729",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "diffServAction",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction",
    "status": null,
    "access": null,
    "category": "diffServMIBObjects"
  },
  {
    "id": "e48c1523-2bc2-3980-9bd1-14169fd4319e",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionNextFree",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServActionId, or a zero to indicate that none exist.",
    "category": "diffServAction"
  },
  {
    "id": "1c160c7b-866a-3950-bfb6-2fd7f5682293",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionTable",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The Action Table enumerates actions that can be performed to a stream of traffic. Multiple actions can be concatenated. For example, traffic exiting from a meter may be counted, marked, and potentially dropped before entering a queue. Specific actions are indicated by diffServActionSpecific which points to an entry of a specific action type parameterizing the action in detail.",
    "category": "diffServAction"
  },
  {
    "id": "a6ca6055-db1a-3346-8ed3-a1c0a69b21e5",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "Each entry in the action table allows description of one specific action to be applied to traffic.",
    "category": "diffServActionTable"
  },
  {
    "id": "f88d6603-febb-3ef7-98e8-46c2fe1784d1",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry.diffServActionId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Action entries. Managers obtain new values for row creation in this table by reading diffServActionNextFree.",
    "category": "diffServActionEntry"
  },
  {
    "id": "0cfcb2a9-6215-3728-ab2b-ab6dfe63d3a1",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionInterface",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry.diffServActionInterface",
    "syntax": "Integer32 (0... 2147483647)",
    "status": "current",
    "access": "read-create",
    "description": "The interface index (value of ifIndex) that this action occurs on. This may be derived from the diffServDataPathStartEntry's index by extension through the various RowPointers. However, as this may be difficult for a network management station, it is placed here as well. If this is indeterminate, the value is zero. This is of especial relevance when reporting the counters which may apply to traffic crossing an interface: diffServCountActOctets, diffServCountActPkts, diffServAlgDropOctets, diffServAlgDropPkts, diffServAlgRandomDropOctets, and diffServAlgRandomDropPkts. It is also especially relevant to the queue and scheduler which may be subsequently applied.",
    "category": "diffServActionEntry"
  },
  {
    "id": "e19c111e-fd02-3029-a794-fe4210bb0749",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionNext",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry.diffServActionNext",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This selects the next Differentiated Services Functional Data Path Element to handle traffic for this data path. This RowPointer should point to an instance of one of: diffServClfrEntry diffServMeterEntry diffServActionEntry diffServAlgDropEntry diffServQEntry A value of zeroDotZero in this attribute indicates no further Differentiated Services treatment is performed on traffic of this data path. The use of zeroDotZero is the normal usage for the last functional data path element of the current data path. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServActionEntry"
  },
  {
    "id": "ee8a3b69-5af2-3540-83d3-a765a7f87df5",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionSpecific",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry.diffServActionSpecific",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "A pointer to an object instance providing additional information for the type of action indicated by this action table entry. For the standard actions defined by this MIB module, this should point to either a diffServDscpMarkActEntry or a diffServCountActEntry. For other actions, it may point to an object instance defined in some other MIB. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the Meter should be treated as if it were not present. This may lead to incorrect policy behavior.",
    "category": "diffServActionEntry"
  },
  {
    "id": "e17fd3dd-c054-3494-a1ee-a16a8f7853c5",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionStorage",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry.diffServActionStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServActionEntry"
  },
  {
    "id": "e252222f-a41b-3a5f-87aa-75a4f67e09c3",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServActionStatus",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.2.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServActionTable.diffServActionEntry.diffServActionStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServActionEntry"
  },
  {
    "id": "4f61f016-4571-343f-8422-fc4ff141431a",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServDscpMarkActTable",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServDscpMarkActTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "This table enumerates specific DSCPs used for marking or remarking the DSCP field of IP packets. The entries of this table may be referenced by a diffServActionSpecific attribute.",
    "category": "diffServAction"
  },
  {
    "id": "232e0b10-67b3-3adf-b29e-d85f08da501e",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServDscpMarkActEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.3.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServDscpMarkActTable.diffServDscpMarkActEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the DSCP mark action table that describes a single DSCP used for marking.",
    "category": "diffServDscpMarkActTable"
  },
  {
    "id": "9e176796-d9c4-38ff-a699-0d791de50591",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServDscpMarkActDscp",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.3.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServDscpMarkActTable.diffServDscpMarkActEntry.diffServDscpMarkActDscp",
    "syntax": "Integer32 (0... 63)",
    "status": "current",
    "access": "read-only",
    "description": "The DSCP that this Action will store into the DSCP field of the subject. It is quite possible that the only packets subject to this Action are already marked with this DSCP. Note also that Differentiated Services processing may result in packet being marked on both ingress to a network and on egress from it, and that ingress and egress can occur in the same router.",
    "category": "diffServDscpMarkActEntry"
  },
  {
    "id": "b2a34a40-5f91-33c6-95fc-ae688738f37c",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActNextFree",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServCountActId, or a zero to indicate that none exist.",
    "category": "diffServAction"
  },
  {
    "id": "23120177-ac79-3972-996e-d5be69dd1c27",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActTable",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "This table contains counters for all the traffic passing through an action element.",
    "category": "diffServAction"
  },
  {
    "id": "384088d4-6bfd-3eef-a0be-36e2c8c5d249",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable.diffServCountActEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the count action table describes a single set of traffic counters.",
    "category": "diffServCountActTable"
  },
  {
    "id": "a98b9bf7-a4e7-3f44-a397-83759eefdb86",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable.diffServCountActEntry.diffServCountActId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Count Action entries. Managers obtain new values for row creation in this table by reading diffServCountActNextFree.",
    "category": "diffServCountActEntry"
  },
  {
    "id": "6e523312-61eb-3302-a086-900b4bf9c39b",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActOctets",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable.diffServCountActEntry.diffServCountActOctets",
    "syntax": "Counter64",
    "status": "current",
    "access": "read-only",
    "description": "The number of octets at the Action data path element. Discontinuities in the value of this counter can occur at re - initialization of the management system and at other times as indicated by the value of ifCounterDiscontinuityTime on the relevant interface.",
    "category": "diffServCountActEntry"
  },
  {
    "id": "3279c436-4af6-35a6-ae09-764a3e6883b2",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActPkts",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable.diffServCountActEntry.diffServCountActPkts",
    "syntax": "Counter64",
    "status": "current",
    "access": "read-only",
    "description": "The number of packets at the Action data path element. Discontinuities in the value of this counter can occur at re - initialization of the management system and at other times as indicated by the value of ifCounterDiscontinuityTime on the relevant interface.",
    "category": "diffServCountActEntry"
  },
  {
    "id": "60bb941e-a7c2-3ae2-87d4-4ac4ff5eb65f",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActStorage",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable.diffServCountActEntry.diffServCountActStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServCountActEntry"
  },
  {
    "id": "f588a64e-8b76-3514-90f0-486dd094a85c",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServCountActStatus",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.5.5.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAction.diffServCountActTable.diffServCountActEntry.diffServCountActStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServCountActEntry"
  },
  {
    "id": "896a02b0-ac83-317a-85a4-e41c601a282c",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "diffServAlgDrop",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop",
    "status": null,
    "access": null,
    "category": "diffServMIBObjects"
  },
  {
    "id": "54de9da5-da0d-300b-8302-181c7be1b1a9",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropNextFree",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServAlgDropId, or a zero to indicate that none exist.",
    "category": "diffServAlgDrop"
  },
  {
    "id": "73706902-7ce3-3865-b5b9-8c47573ee6f9",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropTable",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The algorithmic drop table contains entries describing an element that drops packets according to some algorithm.",
    "category": "diffServAlgDrop"
  },
  {
    "id": "ab9b6d2c-3d06-3ac7-870d-1c9637f5a460",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry describes a process that drops packets according to some algorithm. Further details of the algorithm type are to be found in diffServAlgDropType and with more detail parameter entry pointed to by diffServAlgDropSpecific when necessary.",
    "category": "diffServAlgDropTable"
  },
  {
    "id": "d4e7ec60-f5b3-3c49-af18-6a0cd7121e0a",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Algorithmic Dropper entries. Managers obtain new values for row creation in this table by reading diffServAlgDropNextFree.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "66e4132b-d91d-30ad-9c9f-b2e9b73219a6",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropType",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropType",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The type of algorithm used by this dropper. The value other (1) requires further specification in some other MIB module. In the tailDrop (2) algorithm, diffServAlgDropQThreshold represents the maximum depth of the queue, pointed to by diffServAlgDropQMeasure, beyond which all newly arriving packets will be dropped. In the headDrop (3) algorithm, if a packet arrives when the current depth of the queue, pointed to by diffServAlgDropQMeasure, is at diffServAlgDropQThreshold, packets currently at the head of the queue are dropped to make room for the new packet to be enqueued at the tail of the queue. In the randomDrop (4) algorithm, on packet arrival, an Active Queue Management algorithm is executed which may randomly drop a packet. This algorithm may be proprietary, and it may drop either the arriving packet or another packet in the queue. diffServAlgDropSpecific points to a diffServRandomDropEntry that describes the algorithm. For this algorithm, diffServAlgDropQThreshold is understood to be the absolute maximum size of the queue and additional parameters are described in diffServRandomDropTable. The alwaysDrop (5) algorithm is as its name specifies; always drop. In this case, the other configuration values in this Entry are not meaningful; There is no useful 'next' processing step, there is no queue, and parameters describing the queue are not useful. Therefore, diffServAlgDropNext, diffServAlgDropMeasure, and diffServAlgDropSpecific are all zeroDotZero.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "1feaaa0a-1edc-3527-bf4b-03966f075e8e",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropNext",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropNext",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This selects the next Differentiated Services Functional Data Path Element to handle traffic for this data path. This RowPointer should point to an instance of one of: diffServClfrEntry diffServMeterEntry diffServActionEntry diffServQEntry A value of zeroDotZero in this attribute indicates no further Differentiated Services treatment is performed on traffic of this data path. The use of zeroDotZero is the normal usage for the last functional data path element of the current data path. When diffServAlgDropType is alwaysDrop (5), this object is ignored. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "a9a87b8d-cbcc-365d-8628-caf53a3d8d39",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropQMeasure",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropQMeasure",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "Points to an entry in the diffServQTable to indicate the queue that a drop algorithm is to monitor when deciding whether to drop a packet. If the row pointed to does not exist, the algorithmic dropper element is considered inactive. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "fab995cf-b9c9-38c2-b58b-23acb340f80b",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropQThreshold",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropQThreshold",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "Bytes",
    "description": "A threshold on the depth in bytes of the queue being measured at which a trigger is generated to the dropping algorithm, unless diffServAlgDropType is alwaysDrop (5) where this object is ignored. For the tailDrop (2) or headDrop (3) algorithms, this represents the depth of the queue, pointed to by diffServAlgDropQMeasure, at which the drop action will take place. Other algorithms will need to define their own semantics for this threshold.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "1ebab5d9-75f2-324b-96b8-536ec9f054c4",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropSpecific",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropSpecific",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "Points to a table entry that provides further detail regarding a drop algorithm. Entries with diffServAlgDropType equal to other (1) may have this point to a table defined in another MIB module. Entries with diffServAlgDropType equal to randomDrop (4) must have this point to an entry in diffServRandomDropTable. For all other algorithms specified in this MIB, this should take the value zeroDotZero. The diffServAlgDropType is authoritative for the type of the drop algorithm and the specific parameters for the drop algorithm needs to be evaluated based on the diffServAlgDropType. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "aa501c5b-5c76-3e55-b061-a9b0c7685375",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropOctets",
    "number": 7,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.7",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropOctets",
    "syntax": "Counter64",
    "status": "current",
    "access": "read-only",
    "description": "The number of octets that have been deterministically dropped by this drop process. Discontinuities in the value of this counter can occur at re - initialization of the management system and at other times as indicated by the value of ifCounterDiscontinuityTime on the relevant interface.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "20685ade-f280-38ce-80d2-ac3ae0a26a30",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropPkts",
    "number": 8,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.8",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropPkts",
    "syntax": "Counter64",
    "status": "current",
    "access": "read-only",
    "description": "The number of packets that have been deterministically dropped by this drop process. Discontinuities in the value of this counter can occur at re - initialization of the management system and at other times as indicated by the value of ifCounterDiscontinuityTime on the relevant interface.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "abbce2d2-b39b-3f27-99ea-86fc5b94a3cd",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgRandomDropOctets",
    "number": 9,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.9",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgRandomDropOctets",
    "syntax": "Counter64",
    "status": "current",
    "access": "read-only",
    "description": "The number of octets that have been randomly dropped by this drop process. This counter applies, therefore, only to random droppers. Discontinuities in the value of this counter can occur at re - initialization of the management system and at other times as indicated by the value of ifCounterDiscontinuityTime on the relevant interface.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "a33fa1cf-12d6-3b83-a6ce-9717215f2791",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgRandomDropPkts",
    "number": 10,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.10",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgRandomDropPkts",
    "syntax": "Counter64",
    "status": "current",
    "access": "read-only",
    "description": "The number of packets that have been randomly dropped by this drop process. This counter applies, therefore, only to random droppers. Discontinuities in the value of this counter can occur at re - initialization of the management system and at other times as indicated by the value of ifCounterDiscontinuityTime on the relevant interface.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "bfc62dd8-9e3b-3cfb-88e7-d1efae866dec",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropStorage",
    "number": 11,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.11",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "d5863c65-869b-3c34-976e-dfdb0b5cc9d3",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServAlgDropStatus",
    "number": 12,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.2.1.12",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServAlgDropTable.diffServAlgDropEntry.diffServAlgDropStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServAlgDropEntry"
  },
  {
    "id": "705164f6-903e-3712-8425-009c19092312",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropNextFree",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServRandomDropId, or a zero to indicate that none exist.",
    "category": "diffServAlgDrop"
  },
  {
    "id": "d926741c-3b29-37d0-9c50-9c774f75cd77",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropTable",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The random drop table contains entries describing a process that drops packets randomly. Entries in this table are pointed to by diffServAlgDropSpecific.",
    "category": "diffServAlgDrop"
  },
  {
    "id": "63bb7019-9ea4-3a9b-ac60-01447b30ac5d",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry describes a process that drops packets according to a random algorithm.",
    "category": "diffServRandomDropTable"
  },
  {
    "id": "d969e3cd-d905-31d6-8f9b-e7bc10895caa",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Random Drop entries. Managers obtain new values for row creation in this table by reading diffServRandomDropNextFree.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "402227e1-2a65-396d-8fd0-2b191e3fcdee",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropMinThreshBytes",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropMinThreshBytes",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "bytes",
    "description": "The average queue depth in bytes, beyond which traffic has a non-zero probability of being dropped. Changes in this variable may or may not be reflected in the reported value of diffServRandomDropMinThreshPkts.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "dae542a1-f7cf-3f86-bce5-abf1bb66bc1d",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropMinThreshPkts",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropMinThreshPkts",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "packets",
    "description": "The average queue depth in packets, beyond which traffic has a non-zero probability of being dropped. Changes in this variable may or may not be reflected in the reported value of diffServRandomDropMinThreshBytes.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "30254ac6-e484-35b7-a8aa-4092c94b6533",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropMaxThreshBytes",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropMaxThreshBytes",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "bytes",
    "description": "The average queue depth beyond which traffic has a probability indicated by diffServRandomDropProbMax of being dropped or marked. Note that this differs from the physical queue limit, which is stored in diffServAlgDropQThreshold. Changes in this variable may or may not be reflected in the reported value of diffServRandomDropMaxThreshPkts.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "cd41bd57-4f93-354f-bad1-b05e77c2a750",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropMaxThreshPkts",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropMaxThreshPkts",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "packets",
    "description": "The average queue depth beyond which traffic has a probability indicated by diffServRandomDropProbMax of being dropped or marked. Note that this differs from the physical queue limit, which is stored in diffServAlgDropQThreshold. Changes in this variable may or may not be reflected in the reported value of diffServRandomDropMaxThreshBytes.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "7f96c580-daf0-354d-bb36-4f553865ad0f",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropProbMax",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropProbMax",
    "syntax": "Unsigned32 (0... 1000)",
    "status": "current",
    "access": "read-create",
    "description": "The worst case random drop probability, expressed in drops per thousand packets. For example, if in the worst case every arriving packet may be dropped (100%) for a period, this has the value 1000. Alternatively, if in the worst case only one percent (1%) of traffic may be dropped, it has the value 10.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "b3ae1978-0c15-3125-8233-e395382022bb",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropWeight",
    "number": 7,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.7",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropWeight",
    "syntax": "Unsigned32 (0... 65536)",
    "status": "current",
    "access": "read-create",
    "description": "The weighting of past history in affecting the Exponentially Weighted Moving Average function that calculates the current average queue depth. The equation uses diffServRandomDropWeight/65536 as the coefficient for the new sample in the equation, and (65536 - diffServRandomDropWeight) /65536 as the coefficient of the old value. Implementations may limit the values of diffServRandomDropWeight to a subset of the possible range of values, such as powers of two. Doing this would facilitate implementation of the Exponentially Weighted Moving Average using shift instructions or registers.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "d9272f6a-e32b-3261-a6ff-dbf579532308",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropSamplingRate",
    "number": 8,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.8",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropSamplingRate",
    "syntax": "Unsigned32 (0... 1000000)",
    "status": "current",
    "access": "read-create",
    "description": "The number of times per second the queue is sampled for queue average calculation. A value of zero is used to mean that the queue is sampled approximately each time a packet is enqueued (or dequeued).",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "c9182684-c6a5-3f07-9a8d-a09162f6670a",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropStorage",
    "number": 9,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.9",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "cb06e86b-2e50-3af6-a922-efa7c6a67de3",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServRandomDropStatus",
    "number": 10,
    "dotter_notation": ".1.3.6.1.2.1.97.1.6.4.1.10",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServAlgDrop.diffServRandomDropTable.diffServRandomDropEntry.diffServRandomDropStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServRandomDropEntry"
  },
  {
    "id": "0af4058c-9fcc-3e8e-8c70-e6b505f7d8e2",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "diffServQueue",
    "number": 7,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue",
    "status": null,
    "access": null,
    "category": "diffServMIBObjects"
  },
  {
    "id": "e6a1e5e8-cfcb-3487-8af5-32af9805e2b1",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQNextFree",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServQId, or a zero to indicate that none exist.",
    "category": "diffServQueue"
  },
  {
    "id": "6474916a-3fbb-3854-8108-b947bc84bf08",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQTable",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The Queue Table enumerates the individual queues. Note that the MIB models queuing systems as composed of individual queues, one per class of traffic, even though they may in fact be structured as classes of traffic scheduled using a common calendar queue, or in other ways.",
    "category": "diffServQueue"
  },
  {
    "id": "7ed80dfa-f64f-37e4-b45f-ec6e92abd643",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the Queue Table describes a single queue or class of traffic.",
    "category": "diffServQTable"
  },
  {
    "id": "88b1553c-734c-3175-a3c1-d4d259ab75f3",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry.diffServQId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Queue entries. Managers obtain new values for row creation in this table by reading diffServQNextFree.",
    "category": "diffServQEntry"
  },
  {
    "id": "c62dc9d8-7c6f-3155-b3c4-f00fbc62dccc",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQNext",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry.diffServQNext",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This selects the next Differentiated Services Scheduler. The RowPointer must point to a diffServSchedulerEntry. A value of zeroDotZero in this attribute indicates an incomplete diffServQEntry instance. In such a case, the entry has no operational effect, since it has no parameters to give it meaning. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServQEntry"
  },
  {
    "id": "9f5eec2e-3e4f-30b2-ba6d-f2c416ca801e",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQMinRate",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry.diffServQMinRate",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This RowPointer indicates the diffServMinRateEntry that the scheduler, pointed to by diffServQNext, should use to service this queue. If the row pointed to is zeroDotZero, the minimum rate and priority is unspecified. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServQEntry"
  },
  {
    "id": "69fc0358-5845-3e90-aa84-fa24a1efc626",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQMaxRate",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry.diffServQMaxRate",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This RowPointer indicates the diffServMaxRateEntry that the scheduler, pointed to by diffServQNext, should use to service this queue. If the row pointed to is zeroDotZero, the maximum rate is the line speed of the interface. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServQEntry"
  },
  {
    "id": "ea9384b6-b37d-3ba2-a55f-9f9c2b7205e3",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQStorage",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry.diffServQStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServQEntry"
  },
  {
    "id": "f3af4466-4996-3968-9624-83f1b20c0abd",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServQStatus",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.7.2.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServQueue.diffServQTable.diffServQEntry.diffServQStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServQEntry"
  },
  {
    "id": "a4137fe1-da5d-3353-9b06-dee830e15845",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT IDENTIFIER",
    "name": "diffServScheduler",
    "number": 8,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler",
    "status": null,
    "access": null,
    "category": "diffServMIBObjects"
  },
  {
    "id": "c8e88fec-f125-3a34-b586-82e26eaf53b8",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerNextFree",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServSchedulerId, or a zero to indicate that none exist.",
    "category": "diffServScheduler"
  },
  {
    "id": "8c20022e-3f87-3dd2-a8de-63d603db0da7",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerTable",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The Scheduler Table enumerates packet schedulers. Multiple scheduling algorithms can be used on a given data path, with each algorithm described by one diffServSchedulerEntry.",
    "category": "diffServScheduler"
  },
  {
    "id": "4acdf29a-a608-3db0-9716-b1e6dd923339",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the Scheduler Table describing a single instance of a scheduling algorithm.",
    "category": "diffServSchedulerTable"
  },
  {
    "id": "b9dcca4e-4f7f-3952-a375-e89350831169",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Scheduler entries. Managers obtain new values for row creation in this table by reading diffServSchedulerNextFree.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "0642cc18-7cde-392e-b90a-77783298cfd6",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerNext",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerNext",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This selects the next Differentiated Services Functional Data Path Element to handle traffic for this data path. This normally is null (zeroDotZero), or points to a diffServSchedulerEntry or a diffServQEntry. However, this RowPointer may also point to an instance of: diffServClfrEntry, diffServMeterEntry, diffServActionEntry, diffServAlgDropEntry. It would point another diffServSchedulerEntry when implementing multiple scheduler methods for the same data path, such as having one set of queues scheduled by WRR and that group participating in a priority scheduling system in which other queues compete with it in that way. It might also point to a second scheduler in a hierarchical scheduling system. If the row pointed to is zeroDotZero, no further Differentiated Services treatment is performed on traffic of this data path. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "c1f23317-76fc-314e-adf6-faed5d0af779",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerMethod",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerMethod",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "The scheduling algorithm used by this Scheduler. zeroDotZero indicates that this is unknown. Standard values for generic algorithms: diffServSchedulerPriority, diffServSchedulerWRR, and diffServSchedulerWFQ are specified in this MIB; additional values may be further specified in other MIBs.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "7925c512-1881-3989-b887-beefb568bf57",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerMinRate",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerMinRate",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This RowPointer indicates the entry in diffServMinRateTable which indicates the priority or minimum output rate from this scheduler. This attribute is used only when there is more than one level of scheduler. When it has the value zeroDotZero, it indicates that no minimum rate or priority is imposed. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "e4ad0dde-db82-3813-b51a-5e5fd99a075f",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerMaxRate",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerMaxRate",
    "syntax": "OBJECT IDENTIFIER",
    "status": "current",
    "access": "read-create",
    "description": "This RowPointer indicates the entry in diffServMaxRateTable which indicates the maximum output rate from this scheduler. When more than one maximum rate applies (eg, when a multi-rate shaper is in view), it points to the first of those rate entries. This attribute is used only when there is more than one level of scheduler. When it has the value zeroDotZero, it indicates that no maximum rate is imposed. Setting this to point to a target that does not exist results in an inconsistentValue error. If the row pointed to is removed or becomes inactive by other means, the treatment is as if this attribute contains a value of zeroDotZero.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "a93480f2-6272-36f6-a485-d0abbfdb2563",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerStorage",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "6f046ff5-3d45-3a2f-89f3-693cfc6937b4",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServSchedulerStatus",
    "number": 7,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.2.1.7",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServSchedulerTable.diffServSchedulerEntry.diffServSchedulerStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServSchedulerEntry"
  },
  {
    "id": "f0f23427-6f33-3161-9d92-95b92fc80052",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateNextFree",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServMinRateId, or a zero to indicate that none exist.",
    "category": "diffServScheduler"
  },
  {
    "id": "a4f2a9b9-367a-3175-bacc-2824aeb6a8e5",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateTable",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The Minimum Rate Parameters Table enumerates individual sets of scheduling parameter that can be used/reused by Queues and Schedulers.",
    "category": "diffServScheduler"
  },
  {
    "id": "922e8d04-20a7-38d3-8259-cf8636142992",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the Minimum Rate Parameters Table describes a single set of scheduling parameters for use by one or more queues or schedulers.",
    "category": "diffServMinRateTable"
  },
  {
    "id": "a135443c-2ac8-35d3-b3f6-8b63b24423b0",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry.diffServMinRateId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Scheduler Parameter entries. Managers obtain new values for row creation in this table by reading diffServMinRateNextFree.",
    "category": "diffServMinRateEntry"
  },
  {
    "id": "80bd100f-770f-3ad1-8697-7f62f7db4452",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRatePriority",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry.diffServMinRatePriority",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "description": "The priority of this input to the associated scheduler, relative to the scheduler's other inputs. A queue or scheduler with a larger numeric value will be served before another with a smaller numeric value.",
    "category": "diffServMinRateEntry"
  },
  {
    "id": "3863d161-1612-3748-b0bc-584a40fa9dd8",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateAbsolute",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry.diffServMinRateAbsolute",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "kilobits per second",
    "description": "The minimum absolute rate, in kilobits/sec, that a downstream scheduler element should allocate to this queue. If the value is zero, then there is effectively no minimum rate guarantee. If the value is non-zero, the scheduler will assure the servicing of this queue to at least this rate. Note that this attribute value and that of diffServMinRateRelative are coupled: changes to one will affect the value of the other. They are linked by the following equation, in that setting one will change the other: diffServMinRateRelative = (diffServMinRateAbsolute*1000000) /ifSpeed or, if appropriate: diffServMinRateRelative = diffServMinRateAbsolute/ifHighSpeed",
    "category": "diffServMinRateEntry"
  },
  {
    "id": "4fe5081b-b7fa-3828-ab8a-b3fe39b1c992",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateRelative",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry.diffServMinRateRelative",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "description": "The minimum rate that a downstream scheduler element should allocate to this queue, relative to the maximum rate of the interface as reported by ifSpeed or ifHighSpeed, in units of 1/1000 of 1. If the value is zero, then there is effectively no minimum rate guarantee. If the value is non-zero, the scheduler will assure the servicing of this queue to at least this rate. Note that this attribute value and that of diffServMinRateAbsolute are coupled: changes to one will affect the value of the other. They are linked by the following equation, in that setting one will change the other: diffServMinRateRelative = (diffServMinRateAbsolute*1000000) /ifSpeed or, if appropriate: diffServMinRateRelative = diffServMinRateAbsolute/ifHighSpeed",
    "category": "diffServMinRateEntry"
  },
  {
    "id": "e8c120b2-23b3-3842-9158-0b92fa38ec70",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateStorage",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry.diffServMinRateStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServMinRateEntry"
  },
  {
    "id": "d60650b9-a962-356c-abb1-2b4f53cc04b8",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMinRateStatus",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.4.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMinRateTable.diffServMinRateEntry.diffServMinRateStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServMinRateEntry"
  },
  {
    "id": "22d1c991-48c9-342a-91ed-39be7ac73e7e",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateNextFree",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateNextFree",
    "syntax": "Unsigned32 (0... 4294967295)",
    "status": "current",
    "access": "read-only",
    "description": "This object contains an unused value for diffServMaxRateId, or a zero to indicate that none exist.",
    "category": "diffServScheduler"
  },
  {
    "id": "bf1c15b6-bda7-3b28-ad19-54a680be737a",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateTable",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "The Maximum Rate Parameter Table enumerates individual sets of scheduling parameter that can be used/reused by Queues and Schedulers.",
    "category": "diffServScheduler"
  },
  {
    "id": "dd4fa568-0747-3e5e-badf-02fc209709f6",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateEntry",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry",
    "syntax": "TABLE",
    "status": "current",
    "access": "not-accessible",
    "description": "An entry in the Maximum Rate Parameter Table describes a single set of scheduling parameters for use by one or more queues or schedulers.",
    "category": "diffServMaxRateTable"
  },
  {
    "id": "3684d0e0-def1-34f3-95ac-619597501d70",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateId",
    "number": 1,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.1",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateId",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that enumerates the Maximum Rate Parameter entries. Managers obtain new values for row creation in this table by reading diffServMaxRateNextFree.",
    "category": "diffServMaxRateEntry"
  },
  {
    "id": "f19d1db1-c252-3c81-bf25-a96665548024",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateLevel",
    "number": 2,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.2",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateLevel",
    "syntax": "Unsigned32 (1... 32)",
    "status": "current",
    "access": "not-accessible",
    "description": "An index that indicates which level of a multi-rate shaper is being given its parameters. A multi-rate shaper has some number of rate levels. Frame Relay's dual rate specification refers to a 'committed' and an 'excess' rate; ATM's dual rate specification refers to a 'mean' and a 'peak' rate. This table is generalized to support an arbitrary number of rates. The committed or mean rate is level 1, the peak rate (if any) is the highest level rate configured, and if there are other rates they are distributed in monotonically increasing order between them.",
    "category": "diffServMaxRateEntry"
  },
  {
    "id": "b349ae87-55f2-351e-a551-42737e2f053a",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateAbsolute",
    "number": 3,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.3",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateAbsolute",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "units": "kilobits per second",
    "description": "The maximum rate in kilobits/sec that a downstream scheduler element should allocate to this queue. If the value is zero, then there is effectively no maximum rate limit and that the scheduler should attempt to be work conserving for this queue. If the value is non-zero, the scheduler will limit the servicing of this queue to, at most, this rate in a non-work-conserving manner. Note that this attribute value and that of diffServMaxRateRelative are coupled: changes to one will affect the value of the other. They are linked by the following equation, in that setting one will change the other: diffServMaxRateRelative = (diffServMaxRateAbsolute*1000000) /ifSpeed or, if appropriate: diffServMaxRateRelative = diffServMaxRateAbsolute/ifHighSpeed",
    "category": "diffServMaxRateEntry"
  },
  {
    "id": "29343cc3-9199-3b4b-a84a-e78b28effd26",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateRelative",
    "number": 4,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.4",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateRelative",
    "syntax": "Unsigned32 (1... 4294967295)",
    "status": "current",
    "access": "read-create",
    "description": "The maximum rate that a downstream scheduler element should allocate to this queue, relative to the maximum rate of the interface as reported by ifSpeed or ifHighSpeed, in units of 1/1000 of 1. If the value is zero, then there is effectively no maximum rate limit and the scheduler should attempt to be work conserving for this queue. If the value is non-zero, the scheduler will limit the servicing of this queue to, at most, this rate in a non-work-conserving manner. Note that this attribute value and that of diffServMaxRateAbsolute are coupled: changes to one will affect the value of the other. They are linked by the following equation, in that setting one will change the other: diffServMaxRateRelative = (diffServMaxRateAbsolute*1000000) /ifSpeed or, if appropriate: diffServMaxRateRelative = diffServMaxRateAbsolute/ifHighSpeed",
    "category": "diffServMaxRateEntry"
  },
  {
    "id": "4562445f-64a9-3fc5-8354-92b753ab9d18",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateThreshold",
    "number": 5,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.5",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateThreshold",
    "syntax": "INTEGER (0... '7FFFFFFF'h)",
    "status": "current",
    "access": "read-create",
    "units": "Bytes",
    "description": "The number of bytes of queue depth at which the rate of a multi-rate scheduler will increase to the next output rate. In the last conceptual row for such a shaper, this threshold is ignored and by convention is zero.",
    "category": "diffServMaxRateEntry"
  },
  {
    "id": "eeb6e20e-c55c-3aba-acb9-96e70f525b6d",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateStorage",
    "number": 6,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.6",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateStorage",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The storage type for this conceptual row. Conceptual rows having the value 'permanent' need not allow write-access to any columnar objects in the row.",
    "category": "diffServMaxRateEntry"
  },
  {
    "id": "da068330-d28c-3e14-8fd1-83577bd48e83",
    "mib": {
      "id": 4394,
      "path": "DIFFSERV-MIB.mib",
      "name": "DIFFSERV-MIB",
      "vendor": null
    },
    "type": "OBJECT-TYPE",
    "name": "diffServMaxRateStatus",
    "number": 7,
    "dotter_notation": ".1.3.6.1.2.1.97.1.8.6.1.7",
    "object_descriptor": ".iso.org.dod.internet.mgmt.mib-2.diffServMib.diffServMIBObjects.diffServScheduler.diffServMaxRateTable.diffServMaxRateEntry.diffServMaxRateStatus",
    "syntax": "INTEGER",
    "status": "current",
    "access": "read-create",
    "description": "The status of this conceptual row. All writable objects in this row may be modified at any time. Setting this variable to 'destroy' when the MIB contains one or more RowPointers pointing to it results in destruction being delayed until the row is no longer used.",
    "category": "diffServMaxRateEntry"
  }
]
```

Запрос 8:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/oid

{
  "dotter_notation": ".1.3.6.1.2.1.1.2",
  "mib": "SNMPv2-MIB"
}
```

Ответ 8:

```json
[
  {
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
  }
]
```

</details>

### Возможные коды ошибок

400: Bad Request - Параметр notation отсутствует или имеет некорректный формат
500: Internal Server Error - Внутренняя ошибка парсера или СУБД при поиске OID

---

## [POST] /api/v1/mib-parser/parse - Парсинг файлов MIB

Передаваемый в API JSON-объект может содержать следующие поля:

- 

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/parse

{
  "vendor": "Raisecom",
  "path": [
    "iscom", "2128ea"
  ],
  "mibs": [
    {
      "data": "LS1NaWJOYW1lPXJhaXNlY29tDQotLSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PQ0KLS0gVmVyc2lvbiBpbmZvIA0KLS0NCi0tIFZlcnNpb24gMC4xIENyZWF0ZWQgMjAwMy41LjI0IGJ5IEdPTkdHVUlET05HDQotLSBUaGlzIHZlcnNpb24gb2YgTUlCIGlzIGNyZWF0ZWQganVzdCBmb3IgbWFuYWdlbWVudCBvZiBiYXNlIGFuZCB1c2VyLiANCi0tIA0KLS0gQ29weXJpZ2h0KGMpIDIwMDAtMjAwMiBieSBSYWlzZWNvbSAgTHRkLiANCg0KLS0gYWRkIGlzY29tMjExMEEtTUEtUFdSNCBieSBsamsgMjAxMDA3MTQNCi0tIGFkZCBpc2NvbTI4MjRHIGJ5IHpsdyAyMDA5MDkxNw0KLS0gYWRkIGlzY29tMjAwOWFtYSBieSBsamsgMjAwOTA5MDENCi0tIGFkZCBpc2NvbTMwMjYgIGlzY29tMjgyNiAgCQkJYnkgZ29uZ2d1aWRvbmcgMjAwMzA4MTkNCi0tIGFkZCBpc2NvbTQxMjQgIGlzY29tMjEyNiAgbWV0cm9jb20zMDAwICAgICAJYnkgZ29uZ2d1aWRvbmcgMjAwNDAyMDUNCi0tIHVwZGF0ZSByY2Z0IHRvIHJjMDAyICAsIGZpYmVyIHRvIHJjMDAzICwgdHJhbnNtaXRlciB0byByYzAwNCBieSBnb25nZ3VpZG9uZyAyMDA0MDIwNQ0KLS0gbWVyZ2UgaXNjb20zNTI2LGlzY29tMzAyNixpc2NvbTI4MjYsaXNjb20yMTI2LGlzY29tMjAxNixpc2NvbTIwMDgsaXNjb20yMDI2IHRvIElzY29tU3dpdGNoIGJ5IGxpcWlvbmcgMjAwNDEyMTgNCi0tIG1vZGlmeSByb3MtbGl0ZSBPSUQgZnJvbSAxMiB0byAxNiAgICAgICAgICAgICANCi0tIGFkZCBpc2NvbU1lZGlhQ29udmVydG9yLCBtb2RpZnkgUkM1ODFGRSBmcm9tIDEgdG8gMiwgbW9kaWZ5IFJDNTgxR0UgZnJvbSAyIHRvIDMsYnkgbGlxaW9uZyAyMDA2MDQxMQ0KLS0gYWRkIHJjNzAyYyBieSBzdW56aGFuZmVuZyAyMDA2MDQxMw0KLS0gYWRkIGlzY29tMjg1MiBieSB6aGFvaG9uZ2NlIDIwMDYwNTE3DQotLSBhZGQgIG9wY29tMTAwLTJjIGJ5IGxpdWp1bmZ1ICAyMDA2MDkxNCAgICAgIA0KLS0gYWRkICByYzk1My1nZXN0bTEgYnkgbGl1anVuZnUgIDIwMDYwOTI4ICAgIHs5fSANCi0tIGFkZCAgb3Bjb20zNTAwZSBsaXVqdW5mdSAgMjAwNjEwMjUgICAgezh9ICAgICANCi0tIGFkZCBpc2NvbTI5MjYsIGlzY29tMjkyNkYsIGlzY29tMjAxN0EsIGlzY29tMzAxMiBsaXFpb25nIDIwMDYxMjI4ICAgICAgDQotLSBhZGQgaXNjb20yMDE2QywgaXNjb20zMDI2RSwgaXNjb20zMDI4RiAsaXNjb20zMDUyICBsaXFpb25nIDIwMDcwMjA4DQotLSBtb2RpZnkgaXNjb20zMDEyIHRvIGlzY29tMzAxMkdGICB6aGFvaGMgMjAwNzAyMTYgICANCi0tIGFkZCBpc2NvbTUxMjQsIHJjMzAwMC0xNSxsaXFpb25nLCAyMDA3MDcxOCANCi0tIGFkZCByYzk1M2UtZ2VzdG0xICxsaXFpb25nLCAyMDA3MTAxNSAgICANCi0tIGFkZCByYzk1OS00ZmUxNmUxLCBsaXFpb25nLCAyMDA3MTIwNCAgIA0KLS0gYWRkIG9wY29tMzUwMGUtNiwgbGlxaW9uZywgMjAwNzEyMTcNCi0tIGFkZCByYzcwMi1nZXN0bTQsIGxpcWlvbmcsIDIwMDgwMTI1ICAgICAgICAgDQotLSBhZGQgb3Bjb20zMTA1ICwgbGlxaW9uZywgMjAwODAzMDMgICANCi0tIGFkZCBpc2NvbTIyNTAsbGl6aGltaW4sIDIwMDgwMzMxICAgIA0KLS0gbW9kaWZ5IGlzY29tMjI1MCB0byBpc2NvbTIxNTAtTUEsbGl6aGltaW4sIDIwMDgwNDE2DQotLSBhZGQgcmM3MDJnZXN0bTQsIGxpcWlvbmcsIDIwMDgwNTA4ICAgDQotLSBhZGQgaXNjb20yMTA5LU1BLCBpc2NvbTIxMDlBLU1BLGlzY29tMjExOC1NQSxpc2NvbTIxMjZTLU1BLGNoZW5qdW55b25nLCAyMDA4MDUxNSAgICANCi0tIGFkZCByYzcwMmQsIGxpcWlvbmcsIDIwMDgwNjEzICAgIA0KLS0gYWRkIG9wY29tMzEwNywgbGlxaW9uZywgMjAwODA3MDEgICANCi0tIGFkZCByYzAwNi02LCBsaXFpb25nLCAyMDA4MDgwNyAgICAgICAgIA0KLS0gYWRkIHJjOTU5LWdlc3RtMSwgbGlxaW9uZywgMjAwODA4MTEgICAgICAgDQotLSBhZGQgb3B0VWRTeXNNZ210LCBvcHRVZFN5c01vZHVsZXMsIGxpcWlvbmcsIDIwMDgwODE4DQotLSBhZGQgaXNjb20yMTI2RS1NQSwgemhhbnh1ZWNoYTAsIDIwMDgwODI3ICAgIA0KLS0gYWRkIGlzY29tMjEyNkYtTUEsIHpoYW54dWVjaGFvLCAyMDA4MDkxNg0KLS0gYWRkIGlzY29tMjEyNkZMLU1BLCB6aGFueHVlY2hhbywgMjAwODEwMDcNCi0tIGFkZCByYzU1MUItR0U0RkUsIHpoYW54dWVjaGFvLCAyMDA5MDIwNA0KLS0gYWRkIHJjMzAwMC02LCBsaXFpb25nLCAyMDA4MDkyMiAgICAgICAgDQotLSBhZGQgcmMxMjAxLTRmZTRlMXQxLCBsaXFpb25nLCAyMDA4MTExOCAgICAgDQotLSBhZGQgaXNjb20yMTI2RUEtTUEsIHpoYW54dWVjaGFvLCAyMDA4MTEyOA0KLS0gYWRkIGlzY29tMjExMGEtbWEsIGxqaywgMjAwODEyMTcNCi0tIGFkZCBvcGNvbTMxMDksIGxpcWlvbmcsIDIwMDgxMTI1ICAgIA0KLS0gYWRkIHJjMDA2LTNtLXMsIGxpcWlvbmcsIDIwMDgxMjI2ICAgICANCi0tIGFkZCB0ZG1vcFNlcmllcywgbGlxaW9uZywgMjAwODEyMjkgICAgICAgICAgICAgDQotLSBkZWxldGUgIHJjMTIwMS00ZmU0ZTF0MSwgbHEsIDIwMDgxMjI5ICAgICAgDQotLSBhZGQgb3Bjb20tMTAwLW9hdSwgbHEsIDIwMDkwMjIwICAgICAgDQotLSBhZGQgZGxjb21TZXJpZXMsIGxpcWlvbmcsIDIwMDkwNDIxICAgICAgICAgICANCi0tIGFkZCByYzU1Mi1nZSwgbHEsIDIwMDkwNjAxICAgIA0KLS0gbW9kaWZ5IFJDNTUxQi1HRTRGRSAxNiB0byAxMCwgenEsIDIwMDkwNzI0ICANCi0tIGFkZCByYzMwMDBlICwgbHEsIDIwMDkwODE4ICAgICAgICAgICANCi0tIGFkZCByYzk1My00ZmV4ZTF0MSwgbHEsIDIwMDkwOTAxICAgIA0KLS0gYWRkIHJjOTA1Zy00ZmUxNmUxLCByYzkwNWctZ2VzdG0xLCBscSwgMjAwOTA5MDQgICAgDQotLS0gYWRkIHJjOTUzZWItZ2VzdG0xLCBscSwgMjAxMDAxMDQgICAgICAgIA0KLS0gYWRkICByYzk1My00ZmU4ZTF0MWJsLHJjOTUzLTRmZTRlMXQxYmwscmM5NTMtNGZlOGUxICxyYzk1My00ZmU0ZTEgLHJjOTUxZS00ZmVlMSxscSwgMjAxMDAzMDgNCi0tIGFkZCAgcmMxMTA2ZS1mZS0yd3g0ICxyYzExMDZlLWZlLTJ3eDgsIGh4eiwyMDEwMDYwNA0KLS0gYWRkICBvcGNvbTM1MDBlLWMgLCBoeHosMjAxMDA4MjQNCi0tID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PQ0KDQoNClJBSVNFQ09NLUJBU0UtTUlCICAgIERFRklOSVRJT05TIDo6PSBCRUdJTg0KDQpJTVBPUlRTDQogICAgICAgIGVudGVycHJpc2VzCQlGUk9NIFJGQzExNTUtU01JOw0KDQoNCi0tID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09DQotLQ0KLS0gIE9yZ2FuaXphdGlvbiAgYnJhbmNoZXMNCi0tID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09DQoNCglyYWlzZWNvbQkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGVudGVycHJpc2VzIDg4ODYgfQ0KDQoNCg0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgcmFpc2Vjb20gIG5ldHdvcmtzIHByb2R1Y3RzICBicmFuY2hlcw0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCg0KLS0gIFJhaXNlQ29tICBNYW5hZ2VyIA0KCXJhaXNlY29tQWdlbnQJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbSAxIH0NCg0KLS0gIFRyYW5zbWl0Q29udmVydG9yU2VyaWVzIFNlcmllcw0KCXJjMDAyCQkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDIgfQ0KDQotLSAgVHJhbnNtaXRQREggU2VyaWVzICAxVQ0KCXJjMDAzCQkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDMgfQ0KDQotLSAgVHJhbnNtaXRQREggU2VyaWVzICAxMFUNCglyYzAwNAkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDQgfQ0KDQotLSAgRU9TQWNjZXNzIFNlcmllcyANCglyYzcwMUZFCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gNSB9DQoNCi0tICBJU0NPTSBTZXJpZXMgDQoJaXNjb21TZXJpZXMJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbSA2IH0NCg0KLS0gIE9QQ09NIFNlcmllcyANCglvcGNvbVNlcmllcwkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDcgfQ0KDQotLSAgUkFJU0VDT00gTWFuYWdlciANCglyYWlzZWNvbU1hbmFnZXIJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbSA4IH0NCg0KLS0gIFRyYW5zbWl0UERIIFNlcmllcyAgUEMgDQoJcGNBZ2VudAkJCSAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gOSB9DQoNCi0tICBUcmFuc21pdFBDTSBTZXJpZXMgDQoJcGNjb21TZXJpZXMJCSAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMTAgfQ0KDQotLSAgT0VNIFNlcmllcyANCglvZW1TZXJpZXMJCSAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMTEgfSAgIA0KCQ0KLS0gIHJjIFNlcmllcyANCglyY1NlcmllcwkJICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbSAxMiB9DQoNCi0tCUNvbW1vbiBNSUIgZm9yIE9wdGljYWwgU3lzdGVtIEdyb3VwDQoJcmFpc2Vjb21PcHRTeXNDb21tb24JT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMTUgfQ0KDQotLSAgcm9zLWxpZ2h0IFNlcmllcyANCglyb3NsaXRlU2VyaWVzCQkgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDE2IH0gICAgIA0KCQ0KLS0JZHJhZnQJICANCglkcmFmdCAgICAgICAgICAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMTcgfQ0KCQ0KLS0gIFBPTiBTZXJpZXMgDQoJcG9uU2VyaWVzCQkgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDE4IH0NCg0KLS0gCVRETW9QIFNlcmllcyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIA0KCXRkbW9wU2VyaWVzCSAgICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDE5IH0gICAgDQoJDQotLSAJZGxjb20gU2VyaWVzICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgDQoJZGxjb21TZXJpZXMJICAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMjAgfQ0KDQotLSAgcmFpc2Vjb21UZXJtaW5hbCBTZXJpZXMNCiAgICByYWlzZWNvbVRlcm1pbmFsTWdtdCAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbSAyMSB9DQoNCi0tICBtc2cgU2VyaWVzDQogICAgbXNnU2VyaWVzICAgICAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMjIgfQ0KDQotLSAgaVROIFNlcmllcw0KICAgIGlUTlNlcmllcyAgICAgICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDIzIH0NCg0KLS0gIGlQTiBTZXJpZXMgYWRkZWQgYnkgY2hlbmp1bnlvbmcNCiAgICBpUE5TZXJpZXMJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDIzIH0NCg0KLS0gIEdhemVsbGUgU3dpdGNoIFNlcmllcw0KICAgIGdhemVsbGVTd2l0Y2hTZXJpZXMgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDI2IH0NCg0KLS0gIEdhemVsbGUgVHJhbnNtaXQgU2VyaWVzDQogICAgZ2F6ZWxsZVRyYW5zbWl0U2VyaWVzICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMjcgfQ0KDQotLSAgR2F6ZWxsZSBSb3V0ZXIgU2VyaWVzDQogICAgZ2F6ZWxsZVJvdXRlclNlcmllcyAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb20gMjggfQ0KDQotLSAgU0hFTkxBTlhVTlRPTkcgU2VyaWVzDQoJc2hlbmxhbnh1bnRvbmdTZXJpZXMgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbSAyOSB9DQoNCi0tICBTSEVOTEFOWFVOVE9ORyBUcmFuc21pdCBTZXJpZXMNCglzbHRTZXJpZXMgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBzaGVubGFueHVudG9uZ1NlcmllcyAgIDEgfQ0KDQotLSBPVE4gUHJvZHVjdCBTZXJpZXMNCglPVE5TZXJpZXMgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDMwIH0NCg0KLS1ST1MgTWdtdA0KICAgIHJvc01nbXQgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tIDYwIH0NCg0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgcmFpc2Vjb20gQWdlbnQgDQotLSANCg0KCXJhaXNlY29tQ2x1c3RlcgkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tQWdlbnQgNiB9CQ0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCg0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgSVNDT00gU2VyaWVzIA0KLS0gDQoNCglpc2NvbVN3aXRjaAkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDEgfQkgIA0KCWlzY29tMzAyNgkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDIgfQ0KCWlzY29tMjgyNgkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDMgfQ0KCWlzY29tNDEyNAkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDQgfQ0KCWlzY29tMjEyNgkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDUgfQ0KICAgIGlzY29tMjAxNgkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDYgfQ0KICAgIGlzY29tMjAwOAkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDcgfQkNCiAgICBpc2NvbTQzMDAJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA4IH0JICAgDQoJaXNjb20yMDI2QiAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA5IH0gDQoJaXNjb20yODI2RSAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAxMCB9ICAgIA0KCWlzY29tMjgyOEYgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMTEgfSANCglpc2NvbTI4MTJHRiAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDEyIH0gDQoJaXNjb20yMTA5RiAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAxMyB9IA0KCWlzY29tMjAyNiAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMTQgfSANCglpc2NvbTIwMjUgICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDE1IH0gDQoJaXNjb20yMDE3ICAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAxNiB9IA0KCWlzY29tMjAwOSAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMTcgfSAgDQoJaXNjb20yMTI1ICAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAxOCB9IA0KCWlzY29tMjExNyAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMTkgfSANCglpc2NvbTIxMDkgICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDIwIH0gICANCglpc2NvbTIxMjZlICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDIxIH0JDQoJaXNjb20yODUyICAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAyMiB9CSANCglpc2NvbTIxMjZGICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDIzIH0gICAgICAJIA0KCWlzY29tRXBvbiAgICAgICAgICAJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMjQgfSAgICAgICAgICAgICAJIA0KCWlzY29tMjkyNEdGICAgICAgICAJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMjUgfSAgICAgCSANCglpc2NvbTIxMjZTICAgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDI2IH0gICAgICAgIAkgDQoJaXNjb201NTA0ICAgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDI3IH0NCglpc2NvbTIwMDlBICAgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDI4IH0gICAgIAkgDQoJaXNjb20yMTA5QSAgICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAyOSB9ICAgICAgIAkgDQoJaXNjb20yOTI2ICAgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDMwIH0gICAJIA0KCWlzY29tMjkyNkYgICAgICAgICAJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMzEgfSAgIAkgDQoJaXNjb20yMDE3QSAgICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAzMiB9ICAgIAkgDQoJaXNjb20zMDEyR0YgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAzMyB9ICAgICAgCSANCglpc2NvbTIwMTZDICAgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDM0IH0gIAkgDQoJaXNjb20zMDI2RSAgICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAzNSB9ICAJIA0KCWlzY29tMzAyOEYgICAgICAgICAJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgMzYgfSAgICAJIA0KCWlzY29tMzA1MiAgICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyAzNyB9DQoJaXNjb201MTI0ICAgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDM4IH0NCglpc2NvbTIxNTAtTUEgICAgICAgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDM5IH0NCglpc2NvbTIxMTggICAgICAgICAJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgNDAgfQ0KCWlzY29tMjgyOCAgICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA0NCB9DQoJaXNjb20yMTA5LU1BICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA0NSB9DQoJaXNjb20yMTA5QS1NQSAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA0NiB9DQoJaXNjb20yMTE4LU1BICAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA0NyB9DQoJaXNjb20yMTI2Uy1NQSAgICAgIAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA0OCB9DQoJaXNjb20yMTI2RS1NQSAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA1MCB9DQogICAgaXNjb20yMTI2Ri1NQSAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA1MSB9DQogICAgaXNjb20yMTI2RkwtTUEgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA1MiB9DQoJaXNjb20yMDE3UyAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA1MyB9DQoJaXNjb20yMTI2RUEtTUEgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA1NCB9DQoJaXNjb20yMTEwQS1NQQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgNTUgfQ0KCWlzY29tMjAwOUEtTUEJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDU2IH0NCglpc2NvbTI4MjRHCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpc2NvbVNlcmllcyA1NyB9DQoJaXNjb20yMTEwQS1NQS1QV1I0CU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDU4IH0NCglpc2NvbTI4MjhGLUMJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDYyIH0NCglpc2NvbTI4MjgtTUEJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDYzIH0NCglpU0NPTTI5MjRHRi00QwkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgOTYgfQ0KICAgIGlTQ09NMjkyNEdGLTRHRQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgaXNjb21TZXJpZXMgOTcgfQ0KICAgIHJBWDcxMSAgICAgICAJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGlzY29tU2VyaWVzIDEwMyB9DQoNCi0tID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09DQotLQ0KLS0gIE9QQ09NIFNlcmllcyANCi0tDQoNCglvcGNvbTMxMDAJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBvcGNvbVNlcmllcyAxIH0NCglvcGNvbTEwMC00CQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgMiB9DQoJb3Bjb20zNTAwCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgMyB9DQoJb3Bjb20zMTAxCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgNCB9DQoJb3Bjb20zMTAyCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgNSB9DQoJb3Bjb20zMTAzCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgNiB9DQoJb3Bjb20xMDAtMmMJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBvcGNvbVNlcmllcyA3IH0NCiAgICBvcGNvbTM1MDBlICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IG9wY29tU2VyaWVzIDggfSANCglvcGNvbTM1MDBlLTYgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgOSB9ICAgICANCglvcGNvbTMxMDUgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgMTAgfSAgICANCglvcGNvbTMxMDcgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgMTEgfSAgICAgDQoJb3Bjb20zMTA5ICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IG9wY29tU2VyaWVzIDEyIH0NCglvcGNvbS0xMDAtb2F1ICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IG9wY29tU2VyaWVzIDE1IH0NCglvcGNvbTM1MDBlLWMgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb3Bjb21TZXJpZXMgMTYgfQ0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgUkFJU0VDT00gTWFuYWdlciANCi0tDQoNCglpc2NvbVBNCQkJICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbU1hbmFnZXIgMSB9DQoNCg0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgT0VNIFNlcmllcyANCi0tIA0KDQoJaXNjb20zNDA4CQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgb2VtU2VyaWVzIDEgfQ0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgcmMgU2VyaWVzIA0KLS0gDQoNCglyYzk1MQkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDEgfQ0KICAgIHJjOTU3CQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMiB9DQoJcmM5NTIJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAzIH0NCglvcHRpY2FsdHJhbnNjZWl2ZXIJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgNCB9DQoJcmMwMDYJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyA1IH0NCglyYzcwMiAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgNiB9DQoJcmM3MDJjICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDcgfSAgDQoJcmMwMDYtMSAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDggfSAgIA0KCXJjOTUzLWdlc3RtMSAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyA5IH0gICANCglyYzk1M2UtM2ZlMTZlMSAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMTAgfQ0KCXJjMzAwMC0xNSAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAxMSB9DQoJcmM5NTNlLWdlc3RtMQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAxMiB9DQoJcmM5NTktNGZlMTZlMQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAxMyB9IA0KCXJjNzAyLWdlc3RtNAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAxNCB9CSAgIA0KCXJjNzAyZ2VzdG00ICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAxNSB9ICAgIA0KCXJjNzAyZCAgICAJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDE2IH0gIA0KCXJjMDA2LTYgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAxNyB9ICAgIAkNCglyYzk1OS1nZXN0bTEJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMTggfSAJDQoJcmMzMDAwLTYgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDE5IH0NCglyYzU1Mi1nZSAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMjAgfSANCglyYzAwNi0zbS1zICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMjEgfQ0KCXJjMzAwMGUJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAyMiB9DQoJcmM5NTMtNGZleGUxdDEgCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDIzIH0gDQoJcmM5MDVnLTRmZTE2ZTEJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMjQgfSAgIA0KCXJjOTA1Zy1nZXN0bTEJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMjUgfSANCiAJcmM5NTNlYi1nZXN0bTEJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMjYgfSANCiAJcmM5NTMtNGZlOGUxdDFibAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAyNyB9IA0KIAlyYzk1My00ZmU0ZTF0MWJsCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjU2VyaWVzIDI4IH0gDQogCXJjOTUzLTRmZThlMQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAyOSB9IA0KIAlyYzk1My00ZmU0ZTEJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNTZXJpZXMgMzAgfSAgDQogCXJjOTUxZS00ZmVlMQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY1NlcmllcyAzMSB9IA0KCXJjMTEwNmUtZmUtMnd4NAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0gIHsgcmNTZXJpZXMgMzIgfQ0KCXJjMTEwNmUtZmUtMnd4OAlPQkpFQ1QgSURFTlRJRklFUiA6Oj0gIHsgcmNTZXJpZXMgMzMgfQ0KDQotLSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PQ0KLS0NCi0tICBDb21tb24gTUlCIGZvciBPcHRpY2FsIFN5c3RlbSBHcm91cA0KLS0gDQoJb3B0U3lzTWdtdAkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb21PcHRTeXNDb21tb24gMSB9DQoJb3B0U3lzTW9kdWxlcwlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbU9wdFN5c0NvbW1vbiAyIH0NCglvcHRBZ2VudENhcGFiaWxpdHkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmFpc2Vjb21PcHRTeXNDb21tb24gMyB9DQoJb3B0VWRTeXNNZ210CQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByYWlzZWNvbU9wdFN5c0NvbW1vbiA0IH0NCglvcHRVZFN5c01vZHVsZXMJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJhaXNlY29tT3B0U3lzQ29tbW9uIDUgfQ0KLS0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0NCi0tDQotLSAgcm9zbGl0ZSBTZXJpZXMgDQotLSANCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIA0KCWlzY29tTWVkaWFDb252ZXJ0b3IJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByb3NsaXRlU2VyaWVzIDEgfQ0KCXJjNTgxRkUJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByb3NsaXRlU2VyaWVzIDIgfQkgDQoJcmM1ODFHRQkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJvc2xpdGVTZXJpZXMgMyB9ICAgICAgICAgICAgIA0KCXJjNTUxLUZFCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcm9zbGl0ZVNlcmllcyA0IH0JIA0KCXJjNTUxLUdFCQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcm9zbGl0ZVNlcmllcyA1IH0gICAgICAgICAgICANCglyYzU1MS00RkUJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByb3NsaXRlU2VyaWVzIDYgfQ0KCXJjNTUxQi1GRQkJCU9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJvc2xpdGVTZXJpZXMgNyB9CSANCglyYzU1MUItR0UJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByb3NsaXRlU2VyaWVzIDggfSAgICAgICAgICAgIA0KCXJjNTUxQi00RkUJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByb3NsaXRlU2VyaWVzIDkgfQ0KCXJjNTUxQi1HRTRGRQkJT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcm9zbGl0ZVNlcmllcyAxMCB9DQoJcmM1NTFFLTRHRQkJICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByb3NsaXRlU2VyaWVzIDExIH0NCglyYzU1MUUtR0UJCSAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcm9zbGl0ZVNlcmllcyAxMiB9DQoJcmM1NTFFLTRHRUYJCSAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcm9zbGl0ZVNlcmllcyAxMyB9DQoJDQotLSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PQ0KLS0NCi0tICBkcmFmdCANCi0tIA0KICANCglvYW0gICAgICAgICAgICAgICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IGRyYWZ0IDEgfSAgIA0KCWVwb24gICAgICAgICAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgZHJhZnQgMiB9DQoNCi0tID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09DQotLQ0KLS0gIGlQTg0KLS0NCglpUE4yMDEJCQlPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBpUE5TZXJpZXMgMSB9DQoNCkVORA0K",
      "name": "RAISECOM-BASE-MIB"
    },
    {
      "data": "LS1NaWJOYW1lPXJjRGhjcFNub29waW5nDQotLSAqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKg0KLS0gc3dpdGNoLWRoY3Bzbm9vcGluZy1taWIubWliOiAgUmFpc2Vjb20gREhDUCBTbm9vcGluZyBNSUIgZmlsZQ0KLS0NCi0tIERlYyAyMDA2LCBkb25neGlhb2dhbmcNCi0tDQotLSBDb3B5cmlnaHQgKGMpIDE5OTQtMjAwNiBieSBSYWlzZWNvbSBUZWNobm9sb2d5IENvLiwgTHRkLg0KLS0gQWxsIHJpZ2h0cyByZXNlcnZlZC4NCi0tDQotLSAqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKg0KDQotLSAqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKg0KLS0gTW9kaSBSZXBvcnQ/P0Zvcm1hdDogPG51bWJlcj4sIDx0aW1lPiwgPGF1dGhvcj4sIDxkZXNjPiANCi0tIDAyPz8yMDEwMDMyNj8/aHVvY2hhbywgIGNoYW5nZWluZyBhdDpyY0RoY3BTbm9vcGluZ0JpbmRMZWFzZSBPQkpFQ1QtVFlQRSAgU1lOVEFYICAgVW5zaWduZWQzMiAgICAoY29tcGxpZXIgZXJyb3IpICAgICAgIA0KLS0NCi0tIDAxLCAyMDA4MDcwMSwgd3VtaW5neXUsIGFkZCBub2RlIHJjRGhjcFNub29waW5nQmluZEN1cnJlbnRSb3dzLA0KLS0gICAgICAgICAgICAgICAgICAgICAgICAgcmNEaGNwU25vb3BpbmdCaW5kSGlzdG9yeU1heFJvd3MgYW5kIHRhYmxlDQotLSAgICAgICAgICAgICAgICAgICAgICAgICByY0RoY3BTbm9vcGluZ0JpbmRUYWJsZQ0KLS0NCi0tICoqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqDQoNCg0KUkFJU0VDT00tREhDUC1TTk9PUElORy1NSUIgIERFRklOSVRJT05TIDo6PSBCRUdJTg0KDQpJTVBPUlRTDQogICAgTU9EVUxFLUlERU5USVRZLA0KICAgIE5PVElGSUNBVElPTi1UWVBFLA0KICAgIE9CSkVDVC1UWVBFICAgICAgICAgICAgICAgICAgICAgICAgIEZST00gU05NUHYyLVNNSQ0KDQogICAgTU9EVUxFLUNPTVBMSUFOQ0UsDQogICAgTk9USUZJQ0FUSU9OLUdST1VQLA0KICAgIE9CSkVDVC1HUk9VUCAgICAgICAgICAgICAgICAgICAgICAgIEZST00gU05NUHYyLUNPTkYNCiAgICBNYWNBZGRyZXNzICAgICAgICAgICAgICAgICAgICAgICAgICBGUk9NIFNOTVB2Mi1UQw0KDQogICAgRW5hYmxlVmFyICAgICAgICAgICAgICAgICAgICAgICAgICAgRlJPTSBTV0lUQ0gtVEMNCg0KICAgIGlzY29tU3dpdGNoICAgICAgICAgICAgICAgICAgICAgICAgIEZST00gUkFJU0VDT00tQkFTRS1NSUIgICAgDQogICAgDQogICAgSW5ldEFkZHJlc3NJUHY0LA0KICAgIEluZXRBZGRyZXNzSVB2NiAgICAgICAgICAgICAgICAgICAgIEZST00gSU5FVC1BRERSRVNTLU1JQjsNCg0KDQogICAgcmNEaGNwU25vb3BpbmcgTU9EVUxFLUlERU5USVRZDQogICAgICAgIExBU1QtVVBEQVRFRCAgICAiMjAxMDEyMTAwMDAwWiINCiAgICAgICAgT1JHQU5JWkFUSU9OICAgICJSYWlzZWNvbSBUZWNobm9sb2d5IENvLiwgTHRkLiINCiAgICAgICAgQ09OVEFDVC1JTkZPDQogICAgICAgICAgICAgICAgIlJhaXNlY29tIFN5c3RlbXMNCg0KICAgICAgICAgICAgICAgIFBvc3RhbDogTm8uNjE3LzYxOSBIYWl0YWkgVG93ZXIsDQogICAgICAgICAgICAgICAgMjI5IEZvdXJ0aCBOb3J0aCBMb29wIE1pZGRsZSBSb2FkLA0KICAgICAgICAgICAgICAgIEhhaWRpYW4gRGlzdHJpY3QsIEJlaWppbmcsIFBSQw0KDQogICAgICAgICAgICAgICAgVGVsOiArODYtMTAtODI4ODQ0OTkNCg0KICAgICAgICAgICAgICAgIEUtbWFpbDogZG9uZ3hpYW9nYW5nQHJhaXNlY29tLmNvbSINCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICJUaGUgTUlCIG1vZHVsZSBpcyBmb3IgY29uZmlndXJhdGlvbiBvZiBESENQIFNub29waW5nDQogICAgICAgICBmZWF0dXJlLiBESENQIFNub29waW5nIGlzIGEgc2VjdXJpdHkgbWVjaGFuaXNtIHdoaWNoDQogICAgICAgICB1c2VzIGluZm9ybWF0aW9uIGdsZWFuZWQgZnJvbSBESENQIHBhY2tldHMgdG8gcHJvdmlkZQ0KICAgICAgICAgcGVyLXBvcnQgc2VjdXJpdHkgY2FwYWJpbGl0aWVzLiINCiAgICAgICAgUkVWSVNJT04gICAiMjAxMDEyMTAwMDAwWiINCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJUaGUgaW5pdGlhbCByZXZpc2lvbiBvZiB0aGlzIE1JQiBtb2R1bGUuIg0KICAgICAgICA6Oj0geyBpc2NvbVN3aXRjaCAyMyB9DQoNCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KLS0gZGVmaW5lIGdyb3VwcyBpbiByY0RoY3BTbm9vcGluZw0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tDQogICAgcmNEaGNwU25vb3BpbmdNaWJPYmplY3RzDQogICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjRGhjcFNub29waW5nIDEgfQ0KDQogICAgcmNEaGNwU25vb3BpbmdHcm91cA0KICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY0RoY3BTbm9vcGluZ01pYk9iamVjdHMgMSB9DQogICAgcmNEaGNwNlNub29waW5nR3JvdXANCiAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgcmNEaGNwU25vb3BpbmdNaWJPYmplY3RzIDIgfQ0KICAgIHJjRGhjcDRTbm9vcGluZ09wdGlvbkdyb3VwDQogICAgICAgIE9CSkVDVCBJREVOVElGSUVSIDo6PSB7IHJjRGhjcFNub29waW5nTWliT2JqZWN0cyAzIH0gICANCiAgICByY0RoY3A2U25vb3BpbmdPcHRpb25Hcm91cA0KICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyByY0RoY3BTbm9vcGluZ01pYk9iamVjdHMgNCB9DQoNCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KLS0gYmVnaW4tLXJjRGhjcFNub29waW5nR3JvdXAtLS0tLQ0KLS0gcmNEaGNwU25vb3BpbmdHcm91cCBzY2FsZSBNSUItLQ0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KcmNEaGNwU25vb3BpbmdFbmFibGUgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIEVuYWJsZVZhcg0KICAgICAgICBNQVgtQUNDRVNTIHJlYWQtd3JpdGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJUaGUgc3RhdGUgb2YgZ2xvYmFsIGRoY3Agc25vb3BpbmcuIEl0IGhhcyB0d28gdmFsdWVzLA0KICAgICAgICAgICAgb25lIGlzIGVuYWJsZSgxKSx3aGljaCBpbmRpY2F0ZXMgdGhhdCB0aGUgc3lzdGVtIHN0YXJ0IGRoY3Agc25vb3Bpbmc7DQogICAgICAgICAgICB0aGUgb3RoZXIgaXMgZGlzYWJsZSgyKSB0aGF0IG1lYW5zIGRoY3Agc25vb3BpbmcgaXMgaW52YWxpZCBpbiB0aGlzIHN5c3RlbS4iDQogICAgICAgIERFRlZBTCB7IGRpc2FibGUgfQ0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ0dyb3VwIDEgfQ0KDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCi0tIHJjRGhjcFNub29waW5nR3JvdXAgcmNEaGNwU25vb3BpbmdQb3J0VGFibGUtLQ0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tDQpyY0RoY3BTbm9vcGluZ1BvcnRUYWJsZSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggU0VRVUVOQ0UgT0YgUmNEaGNwU25vb3BpbmdQb3J0RW50cnkNCiAgICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgICBTVEFUVVMgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIkEgdGFibGUgcHJvdmlkZXMgdGhlIG1lY2hhbmlzbSB0byBjb250cm9sIERIQ1AgU25vb3BpbmcgcGVyIHBvcnQuIg0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ0dyb3VwIDIgfQ0KDQpyY0RoY3BTbm9vcGluZ1BvcnRFbnRyeSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggUmNEaGNwU25vb3BpbmdQb3J0RW50cnkNCiAgICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgICBTVEFUVVMgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIlRhYmxlIGVudHJ5IGZvciBzd2l0Y2ggZGhjcCBzbm9vcGluZyBjb25maWcgYmFzZWQgb24gcG9ydC4iDQogICAgICAgIElOREVYIHsgcmNEaGNwU25vb3BpbmdQb3J0SW5kZXggfQ0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ1BvcnRUYWJsZSAxIH0NCg0KUmNEaGNwU25vb3BpbmdQb3J0RW50cnkgOjo9IFNFUVVFTkNFDQogICAgew0KICAgICAgICByY0RoY3BTbm9vcGluZ1BvcnRJbmRleCAgICAgSU5URUdFUiwNCiAgICAgICAgcmNEaGNwU25vb3BpbmdQb3J0RW5hYmxlICAgIEVuYWJsZVZhciwNCiAgICAgICAgcmNEaGNwU25vb3BpbmdQb3J0VHJ1c3QgICAgIElOVEVHRVINCiAgICB9DQoNCnJjRGhjcFNub29waW5nUG9ydEluZGV4IE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBJTlRFR0VSDQogICAgICAgIE1BWC1BQ0NFU1Mgbm90LWFjY2Vzc2libGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJBbiBpbmRleCB0aGF0IHVuaXF1ZWx5IGlkZW50aWZpZXMgYSBjb25maWd1cmF0aW9uIGFib3V0IGRoY3Agc25vb3BpbmcuIg0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ1BvcnRFbnRyeSAxfQ0KDQpyY0RoY3BTbm9vcGluZ1BvcnRFbmFibGUgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIEVuYWJsZVZhcg0KICAgICAgICBNQVgtQUNDRVNTIHJlYWQtd3JpdGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJUaGUgc3RhdGUgb2YgZGhjcCBzbm9vcGluZyBhYm91dCBzcGVjaWZpZWQgcG9ydC4NCiAgICAgICAgICAgIEl0IGFsc28gaGFzIHR3byB2YWx1ZXMsb25lIGlzIGVuYWJsZWQoMSksd2hpY2ggaW5kaWNhdGVzIHRoYXQgdGhlIHBvcnQNCiAgICAgICAgICAgIHN0YXJ0IGRoY3Agc25vb3Bpbmc7IHRoZSBvdGhlciBpcyBkaXNhYmxlKDIpIHRoYXQgbWVhbnMgZGhjcCBzbm9vcGluZw0KICAgICAgICAgICAgaXMgaW52YWxpZCBvbiB0aGUgcG9ydC5UaGUgZGVmYXVsdCB2YWx1ZSBpcyBlbmFibGUoMSkuIg0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ1BvcnRFbnRyeSAyIH0NCg0KcmNEaGNwU25vb3BpbmdQb3J0VHJ1c3QgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYICAgICAgSU5URUdFUnt0cnVzdGVkKDEpLCB1bnRydXN0ZWQoMil9DQogICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtd3JpdGUNCiAgICAgICAgU1RBVFVTICAgICAgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIlRydXN0IHN0YXRlIG9mIHBvcnQuSXQgaGFzIHR3byB2YWx1ZXMsb25lIGlzIHRydXN0ZWQoMSksdGhlIG90aGVyIGlzIHVudHJ1c3RlZCgyKS4NCiAgICAgICAgICAgIFRoZSBkZWZhdWx0IHZhbHVlIGlzIHVudHJ1c3RlZCgyKS4iDQogICAgICAgIDo6PSB7IHJjRGhjcFNub29waW5nUG9ydEVudHJ5IDN9DQoNCnJjRGhjcFNub29waW5nQmluZEN1cnJlbnRSb3dzIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBJTlRFR0VSDQogICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiQ3VycmVudCByb3dzIG9mIGJpbmRpbmcgdGFibGUuIg0KICAgICAgICBERUZWQUwgeyAwIH0NCiAgICAgICAgOjo9IHsgcmNEaGNwU25vb3BpbmdHcm91cCAzIH0NCg0KcmNEaGNwU25vb3BpbmdCaW5kSGlzdG9yeU1heFJvd3MgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIElOVEVHRVINCiAgICAgICAgTUFYLUFDQ0VTUyByZWFkLW9ubHkNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJIaXN0b3J5IG1heCByb3dzIG9mIGJpbmRpbmcgdGFibGUuIg0KICAgICAgICBERUZWQUwgeyAwIH0NCiAgICAgICAgOjo9IHsgcmNEaGNwU25vb3BpbmdHcm91cCA0IH0gICANCiAgICAgICAgDQoNCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KLS0gcmNEaGNwU25vb3BpbmdHcm91cCByY0RoY3BTbm9vcGluZ0JpbmRUYWJsZS0tDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCnJjRGhjcFNub29waW5nQmluZFRhYmxlIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBTRVFVRU5DRSBPRiBSY0RoY3BTbm9vcGluZ0JpbmRFbnRyeQ0KICAgICAgICBNQVgtQUNDRVNTIG5vdC1hY2Nlc3NpYmxlDQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiQSB0YWJsZSByZWNvcmRzIGRoY3Agc25vb3BpbmcgYmluZGluZ3MuIg0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ0dyb3VwIDUgfQ0KDQpyY0RoY3BTbm9vcGluZ0JpbmRFbnRyeSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggUmNEaGNwU25vb3BpbmdCaW5kRW50cnkNCiAgICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgICBTVEFUVVMgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIlRhYmxlIGVudHJ5IGZvciBkaGNwIHNub29waW5nIGJpbmRpbmcuIg0KICAgICAgICBJTkRFWCB7IHJjRGhjcFNub29waW5nQmluZElwIH0NCiAgICAgICAgOjo9IHsgcmNEaGNwU25vb3BpbmdCaW5kVGFibGUgMSB9DQoNClJjRGhjcFNub29waW5nQmluZEVudHJ5IDo6PSBTRVFVRU5DRQ0KICAgIHsNCiAgICAgICAgcmNEaGNwU25vb3BpbmdCaW5kSXAgICAgSW5ldEFkZHJlc3NJUHY0LA0KICAgICAgICByY0RoY3BTbm9vcGluZ0JpbmRNYWMgICBNYWNBZGRyZXNzLA0KICAgICAgICByY0RoY3BTbm9vcGluZ0JpbmRMZWFzZSBVbnNpZ25lZDMyLA0KICAgICAgICByY0RoY3BTbm9vcGluZ0JpbmRWbGFuICBJTlRFR0VSKDEuLjQwOTQpLA0KICAgICAgICByY0RoY3BTbm9vcGluZ0JpbmRQb3J0ICBJTlRFR0VSDQogICAgfQ0KDQpyY0RoY3BTbm9vcGluZ0JpbmRJcCBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggSW5ldEFkZHJlc3NJUHY0DQogICAgICAgIE1BWC1BQ0NFU1Mgbm90LWFjY2Vzc2libGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJCaW5kaW5nIElQIGFkZHJlc3MsIHRoZSBpbmRleCBvZiB0aGlzIHRhYmxlIg0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ0JpbmRFbnRyeSAxfQ0KDQpyY0RoY3BTbm9vcGluZ0JpbmRNYWMgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIE1hY0FkZHJlc3MNCiAgICAgICAgTUFYLUFDQ0VTUyByZWFkLW9ubHkNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJCaW5kaW5nIE1BQyBhZGRyZXNzLiINCiAgICAgICAgOjo9IHsgcmNEaGNwU25vb3BpbmdCaW5kRW50cnkgMiB9DQoNCnJjRGhjcFNub29waW5nQmluZExlYXNlIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCAgICAgIFVuc2lnbmVkMzINCiAgICAgICAgTUFYLUFDQ0VTUyAgcmVhZC1vbmx5DQogICAgICAgIFNUQVRVUyAgICAgIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJCaW5kaW5nIGxlYXNlIHJlbWFpbmluZyB0aW1lKHNlYykuIg0KICAgICAgICA6Oj0geyByY0RoY3BTbm9vcGluZ0JpbmRFbnRyeSAzfQ0KDQpyY0RoY3BTbm9vcGluZ0JpbmRWbGFuIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCAgICAgIElOVEVHRVIgKDEuLjQwOTQpDQogICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiQmluZGluZyB2bGFuLiINCiAgICAgICAgOjo9IHsgcmNEaGNwU25vb3BpbmdCaW5kRW50cnkgNH0NCg0KcmNEaGNwU25vb3BpbmdCaW5kUG9ydCBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggICAgICBJTlRFR0VSDQogICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiQmluZGluZyBwb3J0LiINCiAgICAgICAgOjo9IHsgcmNEaGNwU25vb3BpbmdCaW5kRW50cnkgNX0NCg0KLS0gZW5kLS1yY0RoY3BTbm9vcGluZ0dyb3VwLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tICAgICAgDQoNCg0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tDQotLSBiZWdpbi0tcmNEaGNwNlNub29waW5nR3JvdXAtLS0tLQ0KLS0gcmNEaGNwNlNub29waW5nR3JvdXAgc2NhbGUgTUlCLS0NCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCnJjRGhjcDZTbm9vcGluZ0VuYWJsZSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggRW5hYmxlVmFyDQogICAgICAgIE1BWC1BQ0NFU1MgcmVhZC13cml0ZQ0KICAgICAgICBTVEFUVVMgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIlRoZSBzdGF0ZSBvZiBnbG9iYWwgZGhjcCBzbm9vcGluZy4gSXQgaGFzIHR3byB2YWx1ZXMsDQogICAgICAgICAgICBvbmUgaXMgZW5hYmxlKDEpLHdoaWNoIGluZGljYXRlcyB0aGF0IHRoZSBzeXN0ZW0gc3RhcnQgZGhjcCBzbm9vcGluZzsNCiAgICAgICAgICAgIHRoZSBvdGhlciBpcyBkaXNhYmxlKDIpIHRoYXQgbWVhbnMgZGhjcCBzbm9vcGluZyBpcyBpbnZhbGlkIGluIHRoaXMgc3lzdGVtLiINCiAgICAgICAgREVGVkFMIHsgZGlzYWJsZSB9DQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ0dyb3VwIDEgfQ0KDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCi0tIHJjRGhjcDZTbm9vcGluZ0dyb3VwIHJjRGhjcDZTbm9vcGluZ1BvcnRUYWJsZS0tDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCnJjRGhjcDZTbm9vcGluZ1BvcnRUYWJsZSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggU0VRVUVOQ0UgT0YgUmNEaGNwNlNub29waW5nUG9ydEVudHJ5DQogICAgICAgIE1BWC1BQ0NFU1Mgbm90LWFjY2Vzc2libGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJBIHRhYmxlIHByb3ZpZGVzIHRoZSBtZWNoYW5pc20gdG8gY29udHJvbCBESENQIFNub29waW5nIHBlciBwb3J0LiINCiAgICAgICAgOjo9IHsgcmNEaGNwNlNub29waW5nR3JvdXAgMiB9DQoNCnJjRGhjcDZTbm9vcGluZ1BvcnRFbnRyeSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggUmNEaGNwNlNub29waW5nUG9ydEVudHJ5DQogICAgICAgIE1BWC1BQ0NFU1Mgbm90LWFjY2Vzc2libGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJUYWJsZSBlbnRyeSBmb3Igc3dpdGNoIGRoY3Agc25vb3BpbmcgY29uZmlnIGJhc2VkIG9uIHBvcnQuIg0KICAgICAgICBJTkRFWCB7IHJjRGhjcDZTbm9vcGluZ1BvcnRJbmRleCB9DQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ1BvcnRUYWJsZSAxIH0NCg0KUmNEaGNwNlNub29waW5nUG9ydEVudHJ5IDo6PSBTRVFVRU5DRQ0KICAgIHsNCiAgICAgICAgcmNEaGNwNlNub29waW5nUG9ydEluZGV4ICAgICBJTlRFR0VSLA0KICAgICAgICByY0RoY3A2U25vb3BpbmdQb3J0RW5hYmxlICAgIEVuYWJsZVZhciwNCiAgICAgICAgcmNEaGNwNlNub29waW5nUG9ydFRydXN0ICAgICBJTlRFR0VSDQogICAgfQ0KDQpyY0RoY3A2U25vb3BpbmdQb3J0SW5kZXggT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIElOVEVHRVINCiAgICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgICBTVEFUVVMgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIkFuIGluZGV4IHRoYXQgdW5pcXVlbHkgaWRlbnRpZmllcyBhIGNvbmZpZ3VyYXRpb24gYWJvdXQgZGhjcCBzbm9vcGluZy4iDQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ1BvcnRFbnRyeSAxfQ0KDQpyY0RoY3A2U25vb3BpbmdQb3J0RW5hYmxlIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBFbmFibGVWYXINCiAgICAgICAgTUFYLUFDQ0VTUyByZWFkLXdyaXRlDQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiVGhlIHN0YXRlIG9mIGRoY3Agc25vb3BpbmcgYWJvdXQgc3BlY2lmaWVkIHBvcnQuDQogICAgICAgICAgICBJdCBhbHNvIGhhcyB0d28gdmFsdWVzLG9uZSBpcyBlbmFibGVkKDEpLHdoaWNoIGluZGljYXRlcyB0aGF0IHRoZSBwb3J0DQogICAgICAgICAgICBzdGFydCBkaGNwIHNub29waW5nOyB0aGUgb3RoZXIgaXMgZGlzYWJsZSgyKSB0aGF0IG1lYW5zIGRoY3Agc25vb3BpbmcNCiAgICAgICAgICAgIGlzIGludmFsaWQgb24gdGhlIHBvcnQuVGhlIGRlZmF1bHQgdmFsdWUgaXMgZW5hYmxlKDEpLiINCiAgICAgICAgOjo9IHsgcmNEaGNwNlNub29waW5nUG9ydEVudHJ5IDIgfQ0KDQpyY0RoY3A2U25vb3BpbmdQb3J0VHJ1c3QgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYICAgICAgSU5URUdFUnt0cnVzdGVkKDEpLCB1bnRydXN0ZWQoMil9DQogICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtd3JpdGUNCiAgICAgICAgU1RBVFVTICAgICAgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIlRydXN0IHN0YXRlIG9mIHBvcnQuSXQgaGFzIHR3byB2YWx1ZXMsb25lIGlzIHRydXN0ZWQoMSksdGhlIG90aGVyIGlzIHVudHJ1c3RlZCgyKS4NCiAgICAgICAgICAgIFRoZSBkZWZhdWx0IHZhbHVlIGlzIHVudHJ1c3RlZCgyKS4iDQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ1BvcnRFbnRyeSAzfQ0KICAgDQogICANCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KLS0gcmNEaGNwNlNub29waW5nR3JvdXAgc2NhbGUgTUlCLS0NCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KICAgDQpyY0RoY3A2U25vb3BpbmdCaW5kQ3VycmVudFJvd3MgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIElOVEVHRVINCiAgICAgICAgTUFYLUFDQ0VTUyByZWFkLW9ubHkNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJDdXJyZW50IHJvd3Mgb2YgYmluZGluZyB0YWJsZS4iDQogICAgICAgIERFRlZBTCB7IDAgfQ0KICAgICAgICA6Oj0geyByY0RoY3A2U25vb3BpbmdHcm91cCAzIH0NCg0KcmNEaGNwNlNub29waW5nQmluZEhpc3RvcnlNYXhSb3dzIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBJTlRFR0VSDQogICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiSGlzdG9yeSBtYXggcm93cyBvZiBiaW5kaW5nIHRhYmxlLiINCiAgICAgICAgREVGVkFMIHsgMCB9DQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ0dyb3VwIDQgfQ0KDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCi0tIHJjRGhjcDZTbm9vcGluZ0dyb3VwIHJjRGhjcDZTbm9vcGluZ0JpbmRUYWJsZS0tDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCnJjRGhjcDZTbm9vcGluZ0JpbmRUYWJsZSBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggU0VRVUVOQ0UgT0YgUmNEaGNwNlNub29waW5nQmluZEVudHJ5DQogICAgICAgIE1BWC1BQ0NFU1Mgbm90LWFjY2Vzc2libGUNCiAgICAgICAgU1RBVFVTIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJBIHRhYmxlIHJlY29yZHMgZGhjcCBzbm9vcGluZyBiaW5kaW5ncy4iDQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ0dyb3VwIDUgfQ0KDQpyY0RoY3A2U25vb3BpbmdCaW5kRW50cnkgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIFJjRGhjcDZTbm9vcGluZ0JpbmRFbnRyeQ0KICAgICAgICBNQVgtQUNDRVNTIG5vdC1hY2Nlc3NpYmxlDQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiVGFibGUgZW50cnkgZm9yIGRoY3Agc25vb3BpbmcgYmluZGluZy4iDQogICAgICAgIElOREVYIHsgcmNEaGNwNlNub29waW5nQmluZElwIH0NCiAgICAgICAgOjo9IHsgcmNEaGNwNlNub29waW5nQmluZFRhYmxlIDEgfQ0KDQpSY0RoY3A2U25vb3BpbmdCaW5kRW50cnkgOjo9IFNFUVVFTkNFDQogICAgew0KICAgICAgICByY0RoY3A2U25vb3BpbmdCaW5kSXAgICAgSW5ldEFkZHJlc3NJUHY2LA0KICAgICAgICByY0RoY3A2U25vb3BpbmdCaW5kTWFjICAgTWFjQWRkcmVzcywNCiAgICAgICAgcmNEaGNwNlNub29waW5nQmluZExlYXNlIFVuc2lnbmVkMzIsDQogICAgICAgIHJjRGhjcDZTbm9vcGluZ0JpbmRWbGFuICBJTlRFR0VSLA0KICAgICAgICByY0RoY3A2U25vb3BpbmdCaW5kUG9ydCAgSU5URUdFUg0KICAgIH0NCg0KcmNEaGNwNlNub29waW5nQmluZElwIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBJbmV0QWRkcmVzc0lQdjYNCiAgICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgICBTVEFUVVMgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIkJpbmRpbmcgSVAgYWRkcmVzcywgdGhlIGluZGV4IG9mIHRoaXMgdGFibGUiDQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ0JpbmRFbnRyeSAxfQ0KDQpyY0RoY3A2U25vb3BpbmdCaW5kTWFjIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCBNYWNBZGRyZXNzDQogICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiQmluZGluZyBNQUMgYWRkcmVzcy4iDQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ0JpbmRFbnRyeSAyIH0NCg0KcmNEaGNwNlNub29waW5nQmluZExlYXNlIE9CSkVDVC1UWVBFDQogICAgICAgIFNZTlRBWCAgICAgIFVuc2lnbmVkMzINCiAgICAgICAgTUFYLUFDQ0VTUyAgcmVhZC1vbmx5DQogICAgICAgIFNUQVRVUyAgICAgIGN1cnJlbnQNCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJCaW5kaW5nIGxlYXNlIHJlbWFpbmluZyB0aW1lKHNlYykuIg0KICAgICAgICA6Oj0geyByY0RoY3A2U25vb3BpbmdCaW5kRW50cnkgM30NCg0KcmNEaGNwNlNub29waW5nQmluZFZsYW4gT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYICAgICAgSU5URUdFUg0KICAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICAgU1RBVFVTICAgICAgY3VycmVudA0KICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgIkJpbmRpbmcgdmxhbi4iDQogICAgICAgIDo6PSB7IHJjRGhjcDZTbm9vcGluZ0JpbmRFbnRyeSA0fQ0KDQpyY0RoY3A2U25vb3BpbmdCaW5kUG9ydCBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggICAgICBJTlRFR0VSDQogICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiQmluZGluZyBwb3J0LiINCiAgICAgICAgOjo9IHsgcmNEaGNwNlNub29waW5nQmluZEVudHJ5IDV9DQoNCi0tIGVuZC0tcmNEaGNwNlNub29waW5nR3JvdXAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCg0KDQoNCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KLS0gYmVnaW4tLXJjRGhjcDRTbm9vcGluZ09wdGlvbkdyb3VwLS0tLS0NCi0tIHJjRGhjcDRTbm9vcGluZ09wdGlvbkdyb3VwIHNjYWxlIE1JQi0tDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCg0KDQpyY0RoY3BTbm9vcGluZ09wdGlvbkxpc3QgT0JKRUNULVRZUEUNCiAgICAgICAgU1lOVEFYIE9DVEVUIFNUUklORyAoU0laRSgzMikpDQogICAgICAgIE1BWC1BQ0NFU1MgcmVhZC13cml0ZQ0KICAgICAgICBTVEFUVVMgY3VycmVudCANCiAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICJUaGUgZW5hYmxlIHN0YXRlIGxpc3Qgb2YgZGhjcCBzbm9vcGluZyBWNCBjdW1zdG9tIG9wdGlvbiBzdXBwb3J0aW5nLiANCiAgICAgICAgICAgIEl0IGhhcyBpbiB0b3RhbCAyNTYgYml0LCBlYWNoIG9uZSBpbmRpY2F0ZXMgYSBzdGF0ZSBvZiBvbmUgb3B0aW9uIg0KICAgICAgICA6Oj0geyByY0RoY3A0U25vb3BpbmdPcHRpb25Hcm91cCAxIH0gICANCiANCi0tIGVuZC0tcmNEaGNwNFNub29waW5nT3B0aW9uR3JvdXAtLS0tLS0tDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCiAgIA0KICAgDQogICANCi0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQ0KLS0gYmVnaW4tLXJjRGhjcDZTbm9vcGluZ09wdGlvbkdyb3VwLS0tLS0gICAgICAgICAgICAgICANCi0tIHJjRGhjcDZTbm9vcGluZ09wdGlvbkdyb3VwIHNjYWxlIE1JQi0tDQotLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0NCg0KcmNEaGNwNlNub29waW5nT3B0aW9uTGlzdCBPQkpFQ1QtVFlQRQ0KICAgICAgICBTWU5UQVggT0NURVQgU1RSSU5HIChTSVpFKDMyKSkNCiAgICAgICAgTUFYLUFDQ0VTUyByZWFkLXdyaXRlDQogICAgICAgIFNUQVRVUyBjdXJyZW50DQogICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAiVGhlIGVuYWJsZSBzdGF0ZSBsaXN0IG9mIGRoY3Agc25vb3BpbmcgVjYgY3Vtc3RvbSBvcHRpb24gc3VwcG9ydGluZy4gDQogICAgICAgICAgICBJdCBoYXMgaW4gdG90YWwgMjU2IGJpdCwgZWFjaCBvbmUgaW5kaWNhdGVzIGEgc3RhdGUgb2Ygb25lIG9wdGlvbiINCiAgICAgICAgOjo9IHsgcmNEaGNwNlNub29waW5nT3B0aW9uR3JvdXAgMSB9DQogDQotLSBlbmQtLXJjRGhjcDZTbm9vcGluZ09wdGlvbkdyb3VwLS0tLS0tLQ0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tDQogICANCkVORA0K",
      "name": "RAISECOM-DHCP-SNOOPING-MIB"
    }
  ],
  "rewrite": true
}
```

Ответ 1:

```json
{
  "result": [
    { "mib": "RAISECOM-BASE-MIB", "path": "raisecom/iscom/2128ea/RAISECOM-BASE-MIB.mib" },
    { "mib": "RAISECOM-DHCP-SNOOPING-MIB", "path": "raisecom/iscom/2128ea/RAISECOM-DHCP-SNOOPING-MIB.mib" }
  ]
}
```

</details>

### Возможные коды ошибок

400: Bad Request - Получен некорректный JSON-объект
500: Internal Server Error - Ошибка СУБД при поиске записей OID

---

## [GET] /api/v1/mib-parser/oid/mib - Получить OID по названию MIB

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

## [GET] /api/v1/mib-parser/oid/prefix - Поиск OID по префиксу с пагинацией

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

## [GET] /api/v1/mib-parser/oid/vendor - Получить OID по производителю с пагинацией

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

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
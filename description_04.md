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
> присутствовать OID, описание которых приведено как в MIB производителя, так и в стандартных MIB.

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

> [!TIP]
> Данный API поддерживает массовую запись/перезапись файлов MIB, содержание которых должно передаваться в виде
> ByteString.

Передаваемый в API JSON-объект может содержать следующие поля:

- `vendor` - часть названия производителя, для которого должен быть загружен MIB (в случае, если не передано, то файлы
  MIB будут сохранены в базовом каталоге файлов MIB системы, что автоматически сделает их стандартными MIB - то есть
  MIB, содержащими описание OID, распространяющиеся на все устройства всех производителей);
- `path` - массив строк, содержащий путь директории внутри директории производителя, по которому должны быть сохранены
  файлы MIB (если поле `vendor` не заполнено, то данное поле будет игнорироваться и заполняться не должно) - для
  сохранения файлов в корневую директорию производителя необходимо передать пустой массив, либо не передавать значение
  поля вовсе;
- `mibs` - обязательное поле, содержащее данные записываемых/перезаписываемых MIB в виде перечисления:
    * `data` - текстовое содержание файла MIB в виде ByteString;
    * `name` - наименование создаваемого/перезаписываемого файла MIB в системе (расширение `.mib` не указывается);
- `rewrite` - флаг опции записи: если true, то в случае обнаружения по указанному пути существующего файла этот файл
  будет перезаписан, если false - будет создан новый файл, название которого будет иметь индекс записи, указанное в
  скобках (например, `RAISECOM-BASE-MIB(1).mib`).

В случае успешного выполнения запроса будет получена карта, содержащая указанное пользователем желаемое название файла
на путь созданного/перезаписанного файла в системе.

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
500: Internal Server Error - Ошибка системы при парсинге MIB

---

## [POST] /api/v1/mib-parser/mibs - Получить список всех MIB или всех MIB производителя

> [!TIP]
> API позволяет получить из системы список всех MIB производителя (если часть его названия указана в поле `vendor`),
> список имеющихся базовых MIB (если поле `vendor` передано, но его значение - пустая строка), а также список всех MIB
> (если поле `vendor` не передано).

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/mibs

{
  "vendor": "Raisecom"
}
```

Ответ 1:

```json
{
    "mibs": [
        {
            "mib": "AUTO-CONFIGURATION-MIB",
            "path": "raisecom/AUTO-CONFIGURATION-MIB.mib"
        },
        {
            "mib": "CONVERTOR-SYSTEM-MIB",
            "path": "raisecom/CONVERTOR-SYSTEM-MIB.mib"
        },
        {
            "mib": "CONVERTOR-VLAN-MIB",
            "path": "raisecom/CONVERTOR-VLAN-MIB.mib"
        },
        {
            "mib": "DHCP-CLIENT-MIB",
            "path": "raisecom/DHCP-CLIENT-MIB.mib"
        },
        {
            "mib": "DHCP-OPTION-MIB",
            "path": "raisecom/DHCP-OPTION-MIB.mib"
        },
        {
            "mib": "DHCP-RELAY-MIB",
            "path": "raisecom/DHCP-RELAY-MIB.mib"
        },
        {
            "mib": "DHCP-SERVER-MIB",
            "path": "raisecom/DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "DHCP-SNOOPING-MIB",
            "path": "raisecom/DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "IPDHCP-RELAY-MIB",
            "path": "raisecom/IPDHCP-RELAY-MIB.mib"
        },
        {
            "mib": "IPDHCP-SERVER-MIB",
            "path": "raisecom/IPDHCP-SERVER-MIB.mib"
        },
        {
            "mib": "LLDP-PRI-MIB",
            "path": "raisecom/LLDP-PRI-MIB.mib"
        },
        {
            "mib": "LLDP-STD-MIB",
            "path": "raisecom/LLDP-STD-MIB.mib"
        },
        {
            "mib": "OUTBAND-MGMT-MIB",
            "path": "raisecom/OUTBAND-MGMT-MIB.mib"
        },
        {
            "mib": "RAISECOM-ACL-MIB",
            "path": "raisecom/RAISECOM-ACL-MIB.mib"
        },
        {
            "mib": "RAISECOM-ALARM-MGMT-MIB",
            "path": "raisecom/RAISECOM-ALARM-MGMT-MIB.mib"
        },
        {
            "mib": "RAISECOM-APS-MIB",
            "path": "raisecom/RAISECOM-APS-MIB.mib"
        },
        {
            "mib": "RAISECOM-ARP-MIB",
            "path": "raisecom/RAISECOM-ARP-MIB.mib"
        },
        {
            "mib": "RAISECOM-AUTOPROVISIONMDEV-MIB",
            "path": "raisecom/RAISECOM-AUTOPROVISIONMDEV-MIB.mib"
        },
        {
            "mib": "RAISECOM-AUTOPROVISIONRDEV-MIB",
            "path": "raisecom/RAISECOM-AUTOPROVISIONRDEV-MIB.mib"
        },
        {
            "mib": "RAISECOM-BANNER-MIB",
            "path": "raisecom/RAISECOM-BANNER-MIB.mib"
        },
        {
            "mib": "RAISECOM-BASE-MIB",
            "path": "raisecom/RAISECOM-BASE-MIB.mib"
        },
        {
            "mib": "RAISECOM-BFD-MIB",
            "path": "raisecom/RAISECOM-BFD-MIB.mib"
        },
        {
            "mib": "RAISECOM-CFM-MIB",
            "path": "raisecom/RAISECOM-CFM-MIB.mib"
        },
        {
            "mib": "RAISECOM-COMMON-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-COMMON-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP-CLIENT-MIB",
            "path": "raisecom/RAISECOM-DHCP-CLIENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP-OPTION-MIB",
            "path": "raisecom/RAISECOM-DHCP-OPTION-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP-SNOOPING-MIB",
            "path": "raisecom/RAISECOM-DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP6-CLIENT-MIB",
            "path": "raisecom/RAISECOM-DHCP6-CLIENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP6-RELAY-MIB",
            "path": "raisecom/RAISECOM-DHCP6-RELAY-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP6-SERVER-MIB",
            "path": "raisecom/RAISECOM-DHCP6-SERVER-MIB.mib"
        },
        {
            "mib": "RAISECOM-DOT1AG-MIB",
            "path": "raisecom/RAISECOM-DOT1AG-MIB.mib"
        },
        {
            "mib": "RAISECOM-DOT1X-MIB",
            "path": "raisecom/RAISECOM-DOT1X-MIB.mib"
        },
        {
            "mib": "RAISECOM-ELMI-MIB",
            "path": "raisecom/RAISECOM-ELMI-MIB.mib"
        },
        {
            "mib": "RAISECOM-ELPS-MIB",
            "path": "raisecom/RAISECOM-ELPS-MIB.mib"
        },
        {
            "mib": "RAISECOM-ERPS-MIB",
            "path": "raisecom/RAISECOM-ERPS-MIB.mib"
        },
        {
            "mib": "RAISECOM-ETHERSAM-MIB",
            "path": "raisecom/RAISECOM-ETHERSAM-MIB.mib"
        },
        {
            "mib": "RAISECOM-EXTEND-OAM-UPGRADE-MIB",
            "path": "raisecom/RAISECOM-EXTEND-OAM-UPGRADE-MIB.mib"
        },
        {
            "mib": "RAISECOM-EXTLOOPBACK-MIB",
            "path": "raisecom/RAISECOM-EXTLOOPBACK-MIB.mib"
        },
        {
            "mib": "RAISECOM-EXTOAM-MIB",
            "path": "raisecom/RAISECOM-EXTOAM-MIB.mib"
        },
        {
            "mib": "RAISECOM-FANMONITOR-MIB",
            "path": "raisecom/RAISECOM-FANMONITOR-MIB.mib"
        },
        {
            "mib": "RAISECOM-GARP-MIB",
            "path": "raisecom/RAISECOM-GARP-MIB.mib"
        },
        {
            "mib": "RAISECOM-IGMPL2-MIB",
            "path": "raisecom/RAISECOM-IGMPL2-MIB.mib"
        },
        {
            "mib": "RAISECOM-IP-BASE-MIB",
            "path": "raisecom/RAISECOM-IP-BASE-MIB.mib"
        },
        {
            "mib": "RAISECOM-IPMCAST-MIB",
            "path": "raisecom/RAISECOM-IPMCAST-MIB.mib"
        },
        {
            "mib": "RAISECOM-IPSOURCEGUARD-MIB",
            "path": "raisecom/RAISECOM-IPSOURCEGUARD-MIB.mib"
        },
        {
            "mib": "RAISECOM-KEEPALIVE-MIB",
            "path": "raisecom/RAISECOM-KEEPALIVE-MIB.mib"
        },
        {
            "mib": "RAISECOM-KEYCHAIN-MIB",
            "path": "raisecom/RAISECOM-KEYCHAIN-MIB.mib"
        },
        {
            "mib": "RAISECOM-L2CP-MIB",
            "path": "raisecom/RAISECOM-L2CP-MIB.mib"
        },
        {
            "mib": "RAISECOM-LBDETECT-MIB",
            "path": "raisecom/RAISECOM-LBDETECT-MIB.mib"
        },
        {
            "mib": "RAISECOM-LINKAGGREGATION-MIB",
            "path": "raisecom/RAISECOM-LINKAGGREGATION-MIB.mib"
        },
        {
            "mib": "RAISECOM-LLDP-STD-MIB",
            "path": "raisecom/RAISECOM-LLDP-STD-MIB.mib"
        },
        {
            "mib": "RAISECOM-LOOPBACK-MIB",
            "path": "raisecom/RAISECOM-LOOPBACK-MIB.mib"
        },
        {
            "mib": "RAISECOM-MCAST-MIB",
            "path": "raisecom/RAISECOM-MCAST-MIB.mib"
        },
        {
            "mib": "RAISECOM-MGMD-MIB",
            "path": "raisecom/RAISECOM-MGMD-MIB.mib"
        },
        {
            "mib": "RAISECOM-MLACP-MIB",
            "path": "raisecom/RAISECOM-MLACP-MIB.mib"
        },
        {
            "mib": "RAISECOM-MODULE-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-MODULE-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-MPLS-LSPV-MIB",
            "path": "raisecom/RAISECOM-MPLS-LSPV-MIB.mib"
        },
        {
            "mib": "RAISECOM-MPLS-MIB",
            "path": "raisecom/RAISECOM-MPLS-MIB.mib"
        },
        {
            "mib": "RAISECOM-MPLS-QOS-MIB",
            "path": "raisecom/RAISECOM-MPLS-QOS-MIB.mib"
        },
        {
            "mib": "RAISECOM-NAT-MIB",
            "path": "raisecom/RAISECOM-NAT-MIB.mib"
        },
        {
            "mib": "RAISECOM-NDP-MIB",
            "path": "raisecom/RAISECOM-NDP-MIB.mib"
        },
        {
            "mib": "RAISECOM-NMS-ACC-MIB",
            "path": "raisecom/RAISECOM-NMS-ACC-MIB.mib"
        },
        {
            "mib": "RAISECOM-NOTIFICATION-MIB",
            "path": "raisecom/RAISECOM-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "RAISECOM-NTP-MIB",
            "path": "raisecom/RAISECOM-NTP-MIB.mib"
        },
        {
            "mib": "RAISECOM-OAM-MIB",
            "path": "raisecom/RAISECOM-OAM-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPT-REMOTE-RMC-MIB",
            "path": "raisecom/RAISECOM-OPT-REMOTE-RMC-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-DEVICE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-DEVICE-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-ENTITY-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-ENTITY-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-MODULE-TYPE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-MODULE-TYPE-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-MONITOR-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-MONITOR-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-TRANSCEIVER-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-UDETH-INTERFACE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-UDETH-INTERFACE-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-UDSFP-INTERFACE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-UDSFP-INTERFACE-MIB.mib"
        },
        {
            "mib": "RAISECOM-OSPF-MIB",
            "path": "raisecom/RAISECOM-OSPF-MIB.mib"
        },
        {
            "mib": "RAISECOM-OSPFV2-MIB",
            "path": "raisecom/RAISECOM-OSPFV2-MIB.mib"
        },
        {
            "mib": "RAISECOM-PAE-MIB",
            "path": "raisecom/RAISECOM-PAE-MIB.mib"
        },
        {
            "mib": "RAISECOM-PERF-MIB",
            "path": "raisecom/RAISECOM-PERF-MIB.mib"
        },
        {
            "mib": "RAISECOM-PIM-MIB",
            "path": "raisecom/RAISECOM-PIM-MIB.mib"
        },
        {
            "mib": "RAISECOM-POE-MIB",
            "path": "raisecom/RAISECOM-POE-MIB.mib"
        },
        {
            "mib": "RAISECOM-PON-DEVICE-MIB",
            "path": "raisecom/RAISECOM-PON-DEVICE-MIB.mib"
        },
        {
            "mib": "RAISECOM-PONSERIES-BASE-MIB",
            "path": "raisecom/RAISECOM-PONSERIES-BASE-MIB.mib"
        },
        {
            "mib": "RAISECOM-PONSERIES-TC",
            "path": "raisecom/RAISECOM-PONSERIES-TC.mib"
        },
        {
            "mib": "RAISECOM-PORTSTATISTIC-MIB",
            "path": "raisecom/RAISECOM-PORTSTATISTIC-MIB.mib"
        },
        {
            "mib": "RAISECOM-POWERMONITOR-MIB",
            "path": "raisecom/RAISECOM-POWERMONITOR-MIB.mib"
        },
        {
            "mib": "RAISECOM-PPPOEAGENT-MIB",
            "path": "raisecom/RAISECOM-PPPOEAGENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-PTP-MIB",
            "path": "raisecom/RAISECOM-PTP-MIB.mib"
        },
        {
            "mib": "RAISECOM-QINQ-MIB",
            "path": "raisecom/RAISECOM-QINQ-MIB.mib"
        },
        {
            "mib": "RAISECOM-QOS-MIB",
            "path": "raisecom/RAISECOM-QOS-MIB.mib"
        },
        {
            "mib": "RAISECOM-RCFT-MIB",
            "path": "raisecom/RAISECOM-RCFT-MIB.mib"
        },
        {
            "mib": "RAISECOM-RCMP-MIB",
            "path": "raisecom/RAISECOM-RCMP-MIB.mib"
        },
        {
            "mib": "RAISECOM-RELAY-MIB",
            "path": "raisecom/RAISECOM-RELAY-MIB.mib"
        },
        {
            "mib": "RAISECOM-REMOTE-MANAGEMENT-LOCAL-MIB",
            "path": "raisecom/RAISECOM-REMOTE-MANAGEMENT-LOCAL-MIB.mib"
        },
        {
            "mib": "RAISECOM-REMOTE-MANAGEMENT-REMOTE-MIB",
            "path": "raisecom/RAISECOM-REMOTE-MANAGEMENT-REMOTE-MIB.mib"
        },
        {
            "mib": "RAISECOM-RIP-MIB",
            "path": "raisecom/RAISECOM-RIP-MIB.mib"
        },
        {
            "mib": "RAISECOM-RIP2-MIB",
            "path": "raisecom/RAISECOM-RIP2-MIB.mib"
        },
        {
            "mib": "RAISECOM-RNDP-MIB",
            "path": "raisecom/RAISECOM-RNDP-MIB.mib"
        },
        {
            "mib": "RAISECOM-ROUTEMANAGE-MIB",
            "path": "raisecom/RAISECOM-ROUTEMANAGE-MIB.mib"
        },
        {
            "mib": "RAISECOM-ROUTEPOLICY-MIB",
            "path": "raisecom/RAISECOM-ROUTEPOLICY-MIB.mib"
        },
        {
            "mib": "RAISECOM-RRCP-MIB",
            "path": "raisecom/RAISECOM-RRCP-MIB.mib"
        },
        {
            "mib": "RAISECOM-RRCP-VLAN-MIB",
            "path": "raisecom/RAISECOM-RRCP-VLAN-MIB.mib"
        },
        {
            "mib": "RAISECOM-RTDP-MIB",
            "path": "raisecom/RAISECOM-RTDP-MIB.mib"
        },
        {
            "mib": "RAISECOM-SCHEDULE-MIB",
            "path": "raisecom/RAISECOM-SCHEDULE-MIB.mib"
        },
        {
            "mib": "RAISECOM-SLA-MIB",
            "path": "raisecom/RAISECOM-SLA-MIB.mib"
        },
        {
            "mib": "RAISECOM-SROUTE-MIB",
            "path": "raisecom/RAISECOM-SROUTE-MIB.mib"
        },
        {
            "mib": "RAISECOM-SSH-MIB",
            "path": "raisecom/RAISECOM-SSH-MIB.mib"
        },
        {
            "mib": "RAISECOM-SYSLOG-SERVICE-MIB",
            "path": "raisecom/RAISECOM-SYSLOG-SERVICE-MIB.mib"
        },
        {
            "mib": "RAISECOM-SYSTEM-MIB",
            "path": "raisecom/RAISECOM-SYSTEM-MIB.mib"
        },
        {
            "mib": "RAISECOM-UPGRADE-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-UPGRADE-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-USER-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-USER-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-VCT-MIB",
            "path": "raisecom/RAISECOM-VCT-MIB.mib"
        },
        {
            "mib": "RAISECOM-VLANGROUP-MIB",
            "path": "raisecom/RAISECOM-VLANGROUP-MIB.mib"
        },
        {
            "mib": "RAISECOM-VLANMACCOPY-MIB",
            "path": "raisecom/RAISECOM-VLANMACCOPY-MIB.mib"
        },
        {
            "mib": "RAISECOM-VLANPROTECT-MIB",
            "path": "raisecom/RAISECOM-VLANPROTECT-MIB.mib"
        },
        {
            "mib": "RAISECOM-VRRP-MIB",
            "path": "raisecom/RAISECOM-VRRP-MIB.mib"
        },
        {
            "mib": "RAISECOM-WEBSERVER-MIB",
            "path": "raisecom/RAISECOM-WEBSERVER-MIB.mib"
        },
        {
            "mib": "RAISECOM-WRED-MIB",
            "path": "raisecom/RAISECOM-WRED-MIB.mib"
        },
        {
            "mib": "RC002-INTERVAL-PERFORMANCE-STAT-MIB",
            "path": "raisecom/RC002-INTERVAL-PERFORMANCE-STAT-MIB.mib"
        },
        {
            "mib": "RC002-LOCAL-DEVICE-PORT-MIB",
            "path": "raisecom/RC002-LOCAL-DEVICE-PORT-MIB.mib"
        },
        {
            "mib": "RC002-REMOTE-DEVICE-MIB",
            "path": "raisecom/RC002-REMOTE-DEVICE-MIB.mib"
        },
        {
            "mib": "RC002-REMOTEII-DEVICE-MIB",
            "path": "raisecom/RC002-REMOTEII-DEVICE-MIB.mib"
        },
        {
            "mib": "ROSMGMT-ALARM-MGMT-MIB",
            "path": "raisecom/ROSMGMT-ALARM-MGMT-MIB.mib"
        },
        {
            "mib": "ROSMGMT-COMMON-MANAGEMENT-MIB",
            "path": "raisecom/ROSMGMT-COMMON-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "ROSMGMT-MEMORY-MIB",
            "path": "raisecom/ROSMGMT-MEMORY-MIB.mib"
        },
        {
            "mib": "ROSMGMT-OPTICAL-TRANSCEIVER-MIB",
            "path": "raisecom/ROSMGMT-OPTICAL-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "ROSMGMT-OSPFV2-MIB",
            "path": "raisecom/ROSMGMT-OSPFv2-MIB.mib"
        },
        {
            "mib": "ROSMGMT-SYSTEM-MIB",
            "path": "raisecom/ROSMGMT-SYSTEM-MIB.mib"
        },
        {
            "mib": "ROSMGMT-VERSION-MIB",
            "path": "raisecom/ROSMGMT-VERSION-MIB.mib"
        },
        {
            "mib": "SWITCH-AUTO-CONFIGURATION-MIB",
            "path": "raisecom/SWITCH-AUTO-CONFIGURATION-MIB.mib"
        },
        {
            "mib": "SWITCH-CCP-MIB",
            "path": "raisecom/SWITCH-CCP-MIB.mib"
        },
        {
            "mib": "SWITCH-CLKMGMT-MIB",
            "path": "raisecom/SWITCH-CLKMGMT-MIB.mib"
        },
        {
            "mib": "SWITCH-CpuLimit-MIB",
            "path": "raisecom/SWITCH-CpuLimit-MIB.mib"
        },
        {
            "mib": "SWITCH-CPUPRO-MIB",
            "path": "raisecom/SWITCH-CPUPRO-MIB.mib"
        },
        {
            "mib": "SWITCH-DAI-MIB",
            "path": "raisecom/SWITCH-DAI-MIB.mib"
        },
        {
            "mib": "SWITCH-ERING-MIB",
            "path": "raisecom/SWITCH-ERING-MIB.mib"
        },
        {
            "mib": "SWITCH-FILTER-MIB",
            "path": "raisecom/SWITCH-FILTER-MIB.mib"
        },
        {
            "mib": "SWITCH-IFEXTEND-MIB",
            "path": "raisecom/SWITCH-IFEXTEND-MIB.mib"
        },
        {
            "mib": "SWITCH-IGMPSNOOP-MIB",
            "path": "raisecom/SWITCH-IGMPSNOOP-MIB.mib"
        },
        {
            "mib": "SWITCH-INTERFACE-PORT-MIB",
            "path": "raisecom/SWITCH-INTERFACE-PORT-MIB.mib"
        },
        {
            "mib": "SWITCH-L3-MIB",
            "path": "raisecom/SWITCH-L3-MIB.mib"
        },
        {
            "mib": "SWITCH-L3FILTER-MIB",
            "path": "raisecom/SWITCH-L3FILTER-MIB.mib"
        },
        {
            "mib": "SWITCH-LINKSTATETRACK-MIB",
            "path": "raisecom/SWITCH-LINKSTATETRACK-MIB.mib"
        },
        {
            "mib": "SWITCH-MACCONFIG-MIB",
            "path": "raisecom/SWITCH-MACCONFIG-MIB.mib"
        },
        {
            "mib": "SWITCH-MEMORYMANGMENT-MIB",
            "path": "raisecom/SWITCH-MEMORYMANGMENT-MIB.mib"
        },
        {
            "mib": "SWITCH-MSTP-MIB",
            "path": "raisecom/SWITCH-MSTP-MIB.mib"
        },
        {
            "mib": "SWITCH-MULTISYS-MIB",
            "path": "raisecom/SWITCH-MULTISYS-MIB.mib"
        },
        {
            "mib": "SWITCH-MVR-MIB",
            "path": "raisecom/SWITCH-MVR-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTBACKUP-MIB",
            "path": "raisecom/SWITCH-PORTBACKUP-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTMIRROR-MIB",
            "path": "raisecom/SWITCH-PORTMIRROR-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTPEERBACKUP-MIB",
            "path": "raisecom/SWITCH-PORTPEERBACKUP-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTSECURITY-MIB",
            "path": "raisecom/SWITCH-PORTSECURITY-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTSTATISTIC-MIB",
            "path": "raisecom/SWITCH-PORTSTATISTIC-MIB.mib"
        },
        {
            "mib": "SWITCH-RATELIMIT-MIB",
            "path": "raisecom/SWITCH-RATELIMIT-MIB.mib"
        },
        {
            "mib": "SWITCH-RMON-MIB",
            "path": "raisecom/SWITCH-RMON-MIB.mib"
        },
        {
            "mib": "SWITCH-RSTP-MIB",
            "path": "raisecom/SWITCH-RSTP-MIB.mib"
        },
        {
            "mib": "SWITCH-SLOTCARDMGMT-MIB",
            "path": "raisecom/SWITCH-SLOTCARDMGMT-MIB.mib"
        },
        {
            "mib": "SWITCH-SNMP-MIB",
            "path": "raisecom/SWITCH-SNMP-MIB.mib"
        },
        {
            "mib": "SWITCH-SNTP-MIB",
            "path": "raisecom/SWITCH-SNTP-MIB.mib"
        },
        {
            "mib": "SWITCH-SYNCE-MIB",
            "path": "raisecom/SWITCH-SYNCE-MIB.mib"
        },
        {
            "mib": "SWITCH-SYSTEM-MIB",
            "path": "raisecom/SWITCH-SYSTEM-MIB.mib"
        },
        {
            "mib": "SWITCH-TC",
            "path": "raisecom/SWITCH-TC.mib"
        },
        {
            "mib": "SWITCH-TRUNK-MIB",
            "path": "raisecom/SWITCH-TRUNK-MIB.mib"
        },
        {
            "mib": "SWITCH-VLAN-MIB",
            "path": "raisecom/SWITCH-VLAN-MIB.mib"
        },
        {
            "mib": "SWITCH-VLANCFG-MIB",
            "path": "raisecom/SWITCH-VLANCFG-MIB.mib"
        },
        {
            "mib": "SWITCH-VLANPORT-RATELIMIT-MIB",
            "path": "raisecom/SWITCH-VLANPORT-RATELIMIT-MIB.mib"
        },
        {
            "mib": "SWTICH-SERVICE-MIB",
            "path": "raisecom/SWTICH-SERVICE-MIB.mib"
        },
        {
            "mib": "SWTICH-VLANXC-MIB",
            "path": "raisecom/SWTICH-VLANXC-MIB.mib"
        }
    ]
}
```

Запрос 2:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/mibs

{
  "vendor": ""
}
```

Ответ 2:

```json
{
    "mibs": [
        {
            "mib": "ADSL-LINE-EXT-MIB",
            "path": "ADSL-LINE-EXT-MIB.mib"
        },
        {
            "mib": "ADSL-LINE-MIB",
            "path": "ADSL-LINE-MIB.mib"
        },
        {
            "mib": "ADSL-TC-MIB",
            "path": "ADSL-TC-MIB.mib"
        },
        {
            "mib": "ADSL2-LINE-TC-MIB",
            "path": "ADSL2-LINE-TC-MIB.mib"
        },
        {
            "mib": "ALARM-MIB",
            "path": "ALARM-MIB.mib"
        },
        {
            "mib": "BGP4-MIB",
            "path": "BGP4-MIB.mib"
        },
        {
            "mib": "BGP4V2-TC-MIB",
            "path": "BGP4V2-TC-MIB.mib"
        },
        {
            "mib": "BRIDGE-MIB",
            "path": "BRIDGE-MIB.mib"
        },
        {
            "mib": "CAPWAP-BASE-MIB",
            "path": "CAPWAP-BASE-MIB-draft06.mib"
        },
        {
            "mib": "DIAL-CONTROL-MIB",
            "path": "DIAL-CONTROL-MIB.mib"
        },
        {
            "mib": "DIFFSERV-DSCP-TC",
            "path": "DIFFSERV-DSCP-TC.mib"
        },
        {
            "mib": "DIFFSERV-MIB",
            "path": "DIFFSERV-MIB.mib"
        },
        {
            "mib": "DISMAN-EVENT-MIB",
            "path": "DISMAN-EVENT-MIB.mib"
        },
        {
            "mib": "DISMAN-NSLOOKUP-MIB",
            "path": "DISMAN-NSLOOKUP-MIB.mib"
        },
        {
            "mib": "DISMAN-PING-MIB",
            "path": "DISMAN-PING-MIB.mib"
        },
        {
            "mib": "DISMAN-SCHEDULE-MIB",
            "path": "DISMAN-SCHEDULE-MIB.mib"
        },
        {
            "mib": "DISMAN-SCRIPT-MIB",
            "path": "DISMAN-SCRIPT-MIB.mib"
        },
        {
            "mib": "DISMAN-TRACEROUTE-MIB",
            "path": "DISMAN-TRACEROUTE-MIB.mib"
        },
        {
            "mib": "DLSW-MIB",
            "path": "DLSW-MIB.mib"
        },
        {
            "mib": "DNS-RESOLVER-MIB",
            "path": "DNS-RESOLVER-MIB.mib"
        },
        {
            "mib": "DNS-SERVER-MIB",
            "path": "DNS-SERVER-MIB.mib"
        },
        {
            "mib": "DOCS-CABLE-DEVICE-MIB",
            "path": "DOCS-CABLE-DEVICE-MIB.mib"
        },
        {
            "mib": "DOCS-IF-MIB",
            "path": "DOCS-IF-MIB.mib"
        },
        {
            "mib": "DOT3-OAM-MIB",
            "path": "DOT3-OAM-MIB.mib"
        },
        {
            "mib": "DS1-MIB",
            "path": "DS1-MIB.mib"
        },
        {
            "mib": "DS3-MIB",
            "path": "DS3-MIB.mib"
        },
        {
            "mib": "DVMRP-MIB",
            "path": "DVMRP-MIB.mib"
        },
        {
            "mib": "DVMRP-STD-MIB",
            "path": "DVMRP-STD-MIB.mib"
        },
        {
            "mib": "ENTITY-MIB",
            "path": "ENTITY-MIB.mib"
        },
        {
            "mib": "ENTITY-SENSOR-MIB",
            "path": "ENTITY-SENSOR-MIB.mib"
        },
        {
            "mib": "ENTITY-STATE-MIB",
            "path": "ENTITY-STATE-MIB.mib"
        },
        {
            "mib": "ENTITY-STATE-TC-MIB",
            "path": "ENTITY-STATE-TC-MIB.mib"
        },
        {
            "mib": "EtherLike-MIB",
            "path": "EtherLike-MIB.mib"
        },
        {
            "mib": "FCMGMT-MIB",
            "path": "FCMGMT-MIB.mib"
        },
        {
            "mib": "FDDI-SMT73-MIB",
            "path": "FDDI-SMT73-MIB.mib"
        },
        {
            "mib": "FLOAT-TC-MIB",
            "path": "FLOAT-TC-MIB.mib"
        },
        {
            "mib": "FRAME-RELAY-DTE-MIB",
            "path": "FRAME-RELAY-DTE-MIB.mib"
        },
        {
            "mib": "GBOND-MIB",
            "path": "GBOND-MIB.mib"
        },
        {
            "mib": "HC-ALARM-MIB",
            "path": "HC-ALARM-MIB.mib"
        },
        {
            "mib": "HC-PerfHist-TC-MIB",
            "path": "HC-PerfHist-TC-MIB.mib"
        },
        {
            "mib": "HC-RMON-MIB",
            "path": "HC-RMON-MIB.mib"
        },
        {
            "mib": "HCNUM-TC",
            "path": "HCNUM-TC.mib"
        },
        {
            "mib": "HDSL2-SHDSL-LINE-MIB",
            "path": "HDSL2-SHDSL-LINE-MIB.mib"
        },
        {
            "mib": "HOST-RESOURCES-MIB",
            "path": "HOST-RESOURCES-MIB.mib"
        },
        {
            "mib": "HOST-RESOURCES-TYPES",
            "path": "HOST-RESOURCES-TYPES.mib"
        },
        {
            "mib": "IANA-ADDRESS-FAMILY-NUMBERS-MIB",
            "path": "IANA-ADDRESS-FAMILY-NUMBERS-MIB.mib"
        },
        {
            "mib": "IANA-CHARSET-MIB",
            "path": "IANA-CHARSET-MIB.mib"
        },
        {
            "mib": "IANA-ENTITY-MIB",
            "path": "IANA-ENTITY-MIB.mib"
        },
        {
            "mib": "IANA-GMPLS-TC-MIB",
            "path": "IANA-GMPLS-TC-MIB.mib"
        },
        {
            "mib": "IANA-ITU-ALARM-TC-MIB",
            "path": "IANA-ITU-ALARM-TC-MIB.mib"
        },
        {
            "mib": "IANA-LANGUAGE-MIB",
            "path": "IANA-LANGUAGE-MIB.mib"
        },
        {
            "mib": "IANA-MAU-MIB",
            "path": "IANA-MAU-MIB.mib"
        },
        {
            "mib": "IANA-PRINTER-MIB",
            "path": "IANA-PRINTER-MIB.mib"
        },
        {
            "mib": "IANA-PWE3-MIB",
            "path": "IANA-PWE3-MIB.mib"
        },
        {
            "mib": "IANA-RTPROTO-MIB",
            "path": "IANA-RTPROTO-MIB.mib"
        },
        {
            "mib": "IANAifType-MIB",
            "path": "IANAifType-MIB.mib"
        },
        {
            "mib": "IEEE-802DOT17-RPR-MIB",
            "path": "IEEE-802DOT17-RPR-MIB.mib"
        },
        {
            "mib": "IEEE8021-BRIDGE-MIB",
            "path": "IEEE8021-BRIDGE-MIB.mib"
        },
        {
            "mib": "IEEE8021-CFM-MIB",
            "path": "IEEE8021-CFM-MIB.mib"
        },
        {
            "mib": "IEEE8021-CFMD8-MIB",
            "path": "IEEE8021-CFMD8-MIB.mib"
        },
        {
            "mib": "IEEE8021-PAE-MIB",
            "path": "IEEE8021-PAE-MIB.mib"
        },
        {
            "mib": "IEEE8021-Q-BRIDGE-MIB",
            "path": "IEEE8021-Q-BRIDGE-MIB.mib"
        },
        {
            "mib": "IEEE8021-SECY-MIB",
            "path": "IEEE8021-SECY-MIB.mib"
        },
        {
            "mib": "IEEE8021-TC-MIB",
            "path": "IEEE8021-TC-MIB.mib"
        },
        {
            "mib": "IEEE802171-CFM-MIB",
            "path": "IEEE802171-CFM-MIB.mib"
        },
        {
            "mib": "IEEE8023-LAG-MIB",
            "path": "IEEE8023-LAG-MIB.mib"
        },
        {
            "mib": "IEEE802dot11-MIB",
            "path": "IEEE802dot11-MIB.mib"
        },
        {
            "mib": "IGMP-MIB",
            "path": "IGMP-MIB.mib"
        },
        {
            "mib": "IF-MIB",
            "path": "IF-MIB.mib"
        },
        {
            "mib": "IGMP-STD-MIB",
            "path": "IGMP-STD-MIB.mib"
        },
        {
            "mib": "INET-ADDRESS-MIB",
            "path": "INET-ADDRESS-MIB.mib"
        },
        {
            "mib": "INT-SERV-MIB",
            "path": "INT-SERV-MIB.mib"
        },
        {
            "mib": "INTEGRATED-SERVICES-MIB",
            "path": "INTEGRATED-SERVICES-MIB.mib"
        },
        {
            "mib": "IP-FORWARD-MIB",
            "path": "IP-FORWARD-MIB.mib"
        },
        {
            "mib": "IP-MIB",
            "path": "IP-MIB.mib"
        },
        {
            "mib": "IPV6-FLOW-LABEL-MIB",
            "path": "IPV6-FLOW-LABEL-MIB.mib"
        },
        {
            "mib": "IPMROUTE-STD-MIB",
            "path": "IPMROUTE-STD-MIB.mib"
        },
        {
            "mib": "IPMROUTE-MIB",
            "path": "IPMROUTE-MIB.mib"
        },
        {
            "mib": "IPV6-ICMP-MIB",
            "path": "IPV6-ICMP-MIB.mib"
        },
        {
            "mib": "IPV6-MIB",
            "path": "IPV6-MIB.mib"
        },
        {
            "mib": "IPV6-MLD-MIB",
            "path": "IPV6-MLD-MIB.mib"
        },
        {
            "mib": "IPV6-TC",
            "path": "IPV6-TC.mib"
        },
        {
            "mib": "IPV6-TCP-MIB",
            "path": "IPV6-TCP-MIB.mib"
        },
        {
            "mib": "IPV6-UDP-MIB",
            "path": "IPV6-UDP-MIB.mib"
        },
        {
            "mib": "ISDN-MIB",
            "path": "ISDN-MIB.mib"
        },
        {
            "mib": "ITU-ALARM-TC-MIB",
            "path": "ITU-ALARM-TC-MIB.mib"
        },
        {
            "mib": "ISIS-MIB",
            "path": "ISIS-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-DCBX-MIB",
            "path": "LLDP-EXT-DCBX-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-DOT1-MIB",
            "path": "LLDP-EXT-DOT1-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-DOT3-MIB",
            "path": "LLDP-EXT-DOT3-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-MED-MIB",
            "path": "LLDP-EXT-MED-MIB.mib"
        },
        {
            "mib": "LLDP-MIB",
            "path": "LLDP-MIB.mib"
        },
        {
            "mib": "LLDP-V2-MIB",
            "path": "LLDP-V2-MIB.mib"
        },
        {
            "mib": "LLDP-V2-TC-MIB",
            "path": "LLDP-V2-TC-MIB.mib"
        },
        {
            "mib": "MGMD-STD-MIB",
            "path": "MGMD-STD-MIB.mib"
        },
        {
            "mib": "MAU-MIB",
            "path": "MAU-MIB.mib"
        },
        {
            "mib": "MPLS-L3VPN-STD-MIB",
            "path": "MPLS-L3VPN-STD-MIB.mib"
        },
        {
            "mib": "MPLS-LDP-STD-MIB",
            "path": "MPLS-LDP-STD-MIB.mib"
        },
        {
            "mib": "MPLS-LSR-MIB",
            "path": "MPLS-LSR-MIB.mib"
        },
        {
            "mib": "MPLS-LSR-STD-MIB",
            "path": "MPLS-LSR-STD-MIB.mib"
        },
        {
            "mib": "MPLS-TC-STD-MIB",
            "path": "MPLS-TC-STD-MIB.mib"
        },
        {
            "mib": "MPLS-TE-MIB",
            "path": "MPLS-TE-MIB.mib"
        },
        {
            "mib": "MPLS-TE-STD-MIB",
            "path": "MPLS-TE-STD-MIB.mib"
        },
        {
            "mib": "MPLS-VPN-MIB",
            "path": "MPLS-VPN-MIB.mib"
        },
        {
            "mib": "MSTP-MIB",
            "path": "MSTP-MIB.mib"
        },
        {
            "mib": "MTA-MIB",
            "path": "MTA-MIB.mib"
        },
        {
            "mib": "NETWORK-SERVICES-MIB",
            "path": "NETWORK-SERVICES-MIB.mib"
        },
        {
            "mib": "NOTIFICATION-LOG-MIB",
            "path": "NOTIFICATION-LOG-MIB.mib"
        },
        {
            "mib": "OSPF-MIB",
            "path": "OSPF-MIB.mib"
        },
        {
            "mib": "OSPF-TRAP-MIB",
            "path": "OSPF-TRAP-MIB.mib"
        },
        {
            "mib": "OSPFV3-MIB",
            "path": "OSPFV3-MIB.mib"
        },
        {
            "mib": "PerfHist-TC-MIB",
            "path": "PerfHist-TC-MIB.mib"
        },
        {
            "mib": "P-BRIDGE-MIB",
            "path": "P-BRIDGE-MIB.mib"
        },
        {
            "mib": "PIM-MIB",
            "path": "PIM-MIB.mib"
        },
        {
            "mib": "POWER-ETHERNET-MIB",
            "path": "POWER-ETHERNET-MIB.mib"
        },
        {
            "mib": "PPVPN-TC-MIB",
            "path": "PPVPN-TC-MIB.mib"
        },
        {
            "mib": "PTOPO-MIB",
            "path": "PTOPO-MIB.mib"
        },
        {
            "mib": "Printer-MIB",
            "path": "Printer-MIB.mib"
        },
        {
            "mib": "PW-STD-MIB",
            "path": "PW-STD-MIB.mib"
        },
        {
            "mib": "PW-TC-STD-MIB",
            "path": "PW-TC-STD-MIB.mib"
        },
        {
            "mib": "RFC-1215",
            "path": "RFC-1215.mib"
        },
        {
            "mib": "Q-BRIDGE-MIB",
            "path": "Q-BRIDGE-MIB.mib"
        },
        {
            "mib": "RFC-1212",
            "path": "RFC-1212.mib"
        },
        {
            "mib": "RFC1155-SMI",
            "path": "RFC1155-SMI.mib"
        },
        {
            "mib": "RFC1213-MIB",
            "path": "RFC1213-MIB.mib"
        },
        {
            "mib": "RFC1284-MIB",
            "path": "RFC1284-MIB.mib"
        },
        {
            "mib": "RFC1271-MIB",
            "path": "RFC1271-MIB.mib"
        },
        {
            "mib": "RFC1389-MIB",
            "path": "RFC1389-MIB.mib"
        },
        {
            "mib": "RIPv2-MIB",
            "path": "RIPv2-MIB.mib"
        },
        {
            "mib": "RMON-MIB",
            "path": "RMON-MIB.mib"
        },
        {
            "mib": "RMON2-MIB",
            "path": "RMON2-MIB.mib"
        },
        {
            "mib": "RSTP-MIB",
            "path": "RSTP-MIB.mib"
        },
        {
            "mib": "SCTP-MIB",
            "path": "SCTP-MIB.mib"
        },
        {
            "mib": "SMON-MIB",
            "path": "SMON-MIB.mib"
        },
        {
            "mib": "SNA-SDLC-MIB",
            "path": "SNA-SDLC-MIB.mib"
        },
        {
            "mib": "SNMP-COMMUNITY-MIB",
            "path": "SNMP-COMMUNITY-MIB.mib"
        },
        {
            "mib": "SNMP-MPD-MIB",
            "path": "SNMP-MPD-MIB.mib"
        },
        {
            "mib": "SNMP-FRAMEWORK-MIB",
            "path": "SNMP-FRAMEWORK-MIB.mib"
        },
        {
            "mib": "SNMP-NOTIFICATION-MIB",
            "path": "SNMP-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "SNMP-PROXY-MIB",
            "path": "SNMP-PROXY-MIB.mib"
        },
        {
            "mib": "SNMP-REPEATER-MIB",
            "path": "SNMP-REPEATER-MIB.mib"
        },
        {
            "mib": "SNMP-TARGET-MIB",
            "path": "SNMP-TARGET-MIB.mib"
        },
        {
            "mib": "SNMP-USER-BASED-SM-MIB",
            "path": "SNMP-USER-BASED-SM-MIB.mib"
        },
        {
            "mib": "SNMP-USM-AES-MIB",
            "path": "SNMP-USM-AES-MIB.mib"
        },
        {
            "mib": "SNMP-USM-DH-OBJECTS-MIB",
            "path": "SNMP-USM-DH-OBJECTS-MIB.mib"
        },
        {
            "mib": "SNMP-VIEW-BASED-ACM-MIB",
            "path": "SNMP-VIEW-BASED-ACM-MIB.mib"
        },
        {
            "mib": "SNMPv2-CONF",
            "path": "SNMPv2-CONF.mib"
        },
        {
            "mib": "SNMPv2-SMI-v1",
            "path": "SNMPv2-SMI-v1.mib"
        },
        {
            "mib": "SNMPv2-MIB",
            "path": "SNMPv2-MIB.mib"
        },
        {
            "mib": "SNMPv2-SMI",
            "path": "SNMPv2-SMI.mib"
        },
        {
            "mib": "SNMPv2-TC-v1",
            "path": "SNMPv2-TC-v1.mib"
        },
        {
            "mib": "SNMPv2-TM",
            "path": "SNMPv2-TM.mib"
        },
        {
            "mib": "SNMPv2-TC",
            "path": "SNMPv2-TC.mib"
        },
        {
            "mib": "SWITCH-TC",
            "path": "SWITCH-TC.mib"
        },
        {
            "mib": "SONET-MIB",
            "path": "SONET-MIB.mib"
        },
        {
            "mib": "TCP-MIB",
            "path": "TCP-MIB.mib"
        },
        {
            "mib": "SYSAPPL-MIB",
            "path": "SYSAPPL-MIB.mib"
        },
        {
            "mib": "TRANSPORT-ADDRESS-MIB",
            "path": "TRANSPORT-ADDRESS-MIB.mib"
        },
        {
            "mib": "TOKEN-RING-RMON-MIB",
            "path": "TOKEN-RING-RMON-MIB.mib"
        },
        {
            "mib": "UUID-TC-MIB",
            "path": "UUID-TC-MIB.mib"
        },
        {
            "mib": "UDP-MIB",
            "path": "UDP-MIB.mib"
        },
        {
            "mib": "TUNNEL-MIB",
            "path": "TUNNEL-MIB.mib"
        },
        {
            "mib": "UPS-MIB",
            "path": "UPS-MIB.mib"
        },
        {
            "mib": "VDSL-LINE-MIB",
            "path": "VDSL-LINE-MIB.mib"
        },
        {
            "mib": "VPN-TC-STD-MIB",
            "path": "VPN-TC-STD-MIB.mib"
        },
        {
            "mib": "VDSL2-LINE-TC-MIB",
            "path": "VDSL2-LINE-TC-MIB.mib"
        },
        {
            "mib": "VDSL2-LINE-MIB",
            "path": "VDSL2-LINE-MIB.mib"
        },
        {
            "mib": "VRRPV3-MIB",
            "path": "VRRPV3-MIB.mib"
        },
        {
            "mib": "VRRP-MIB",
            "path": "VRRP-MIB.mib"
        }
    ]
}
```

Запрос 3:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/mib-parser/mibs

{}
```

Ответ 3:

```json
{
    "mibs": [
        {
            "mib": "MAIPU-SMI",
            "path": "maipu/MAIPU-SMI.mib"
        },
        {
            "mib": "VELOCLOUD-MIB",
            "path": "vmware/VELOCLOUD-MIB.mib"
        },
        {
            "mib": "CT-BROADCAST-MIB",
            "path": "enterasys/CT-BROADCAST-MIB.mib"
        },
        {
            "mib": "AE-ALARM-TABLE-MIB",
            "path": "calix/AE-ALARM-TABLE-MIB.mib"
        },
        {
            "mib": "VMWARE-CIMOM-MIB",
            "path": "vmware/VMWARE-CIMOM-MIB.mib"
        },
        {
            "mib": "JANITZA-MIB",
            "path": "janitza/JANITZA-MIB.mib"
        },
        {
            "mib": "JANITZA-MIB-UMG96",
            "path": "janitza/JANITZA-MIB-UMG96.mib"
        },
        {
            "mib": "EXAGRID-MIB",
            "path": "exagrid/EXAGRID-MIB.mib"
        },
        {
            "mib": "EdgeSwitch-BOXSERVICES-PRIVATE-MIB",
            "path": "edgeswitch/EdgeSwitch-BOXSERVICES-PRIVATE-MIB.mib"
        },
        {
            "mib": "VELOCLOUD-EDGE-MIB",
            "path": "vmware/VELOCLOUD-EDGE-MIB.mib"
        },
        {
            "mib": "Vega-MIB",
            "path": "sangoma/VEGA-GATEWAY-MIB.mib"
        },
        {
            "mib": "UHP-MIB",
            "path": "uhp/UHP-MIB.mib"
        },
        {
            "mib": "VMWARE-AGENTCAP-MIB",
            "path": "vmware/VMWARE-AGENTCAP-MIB.mib"
        },
        {
            "mib": "AE-PM-TABLE-MIB",
            "path": "calix/AE-PM-TABLE-MIB.mib"
        },
        {
            "mib": "SPSv1-MIB",
            "path": "bladeshelter/SPSv1-MIB.mib"
        },
        {
            "mib": "ALCOMA-MIB",
            "path": "alcoma/ALCOMA-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-AAA-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-AAA-MIB.mib"
        },
        {
            "mib": "F5-BIGIP-APM-MIB",
            "path": "f5/F5-BIGIP-APM-MIB.mib"
        },
        {
            "mib": "GAMATRONIC-MIB",
            "path": "gamatronic/GAMATRONIC-MIB.mib"
        },
        {
            "mib": "CT-CMMPHYS-MIB",
            "path": "enterasys/CT-CMMPHYS-MIB.mib"
        },
        {
            "mib": "CUMULUS-BGPUN-MIB",
            "path": "cumulus/CUMULUS-BGPUN-MIB.mib"
        },
        {
            "mib": "CET-TSI-MIB",
            "path": "cet/CET-TSI-MIB.mib"
        },
        {
            "mib": "VMWARE-ENV-MIB",
            "path": "vmware/VMWARE-ENV-MIB.mib"
        },
        {
            "mib": "HHMSAGENT-MIB",
            "path": "akcp/HHMSAGENT-MIB.mib"
        },
        {
            "mib": "EdgeSwitch-INVENTORY-MIB",
            "path": "edgeswitch/EdgeSwitch-INVENTORY-MIB.mib"
        },
        {
            "mib": "ADVA-FSP3000ALM-MIB",
            "path": "adva/ADVA-FSP3000ALM-MIB.mib"
        },
        {
            "mib": "MPIOS-MIB",
            "path": "maipu/MPIOS-MIB.mib"
        },
        {
            "mib": "ZTE-AN-CHASSIS-MIB",
            "path": "zte/ZTE-AN-CHASSIS-MIB.mib"
        },
        {
            "mib": "proware-SNMP-MIB",
            "path": "proware/proware-SNMP-MIB.mib"
        },
        {
            "mib": "ARICENT-CFA-MIB",
            "path": "cambium/cnmatrix/ARICENT-CFA-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-APP-FINGERPRINT-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-APP-FINGERPRINT-MIB.mib"
        },
        {
            "mib": "AE-TC",
            "path": "calix/AE-TC.mib"
        },
        {
            "mib": "COMTROL-ES8510-MIB",
            "path": "comtrol/ES8510-MIB.mib"
        },
        {
            "mib": "SEH-PSRV-MIB",
            "path": "seh/SEH-PSRV-MIB.mib"
        },
        {
            "mib": "F5-BIGIP-COMMON-MIB",
            "path": "f5/F5-BIGIP-COMMON-MIB.mib"
        },
        {
            "mib": "ARICENT-ISS-MIB",
            "path": "eltexmes24xx/ARICENT-ISS-MIB.mib"
        },
        {
            "mib": "EdgeSwitch-LOGGING-MIB",
            "path": "edgeswitch/EdgeSwitch-LOGGING-MIB.mib"
        },
        {
            "mib": "CT-CONTAINER-MIB",
            "path": "enterasys/CT-CONTAINER-MIB.mib"
        },
        {
            "mib": "CET-TSI-SMI",
            "path": "cet/CET-TSI-SMI.mib"
        },
        {
            "mib": "CUMULUS-BGPVRF-MIB",
            "path": "cumulus/CUMULUS-BGPVRF-MIB.mib"
        },
        {
            "mib": "ZTE-AN-ENVMON-MIB",
            "path": "zte/ZTE-AN-ENVMON-MIB.mib"
        },
        {
            "mib": "ELTEX-MES-ISS-CPU-UTIL-MIB",
            "path": "eltexmes24xx/ELTEX-MES-ISS-CPU-UTIL-MIB.mib"
        },
        {
            "mib": "EdgeSwitch-REF-MIB",
            "path": "edgeswitch/EdgeSwitch-REF-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-AUTO-FABRIC-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-AUTO-FABRIC-MIB.mib"
        },
        {
            "mib": "ARICENT-ISS-MIB",
            "path": "cambium/cnmatrix/ARICENT-ISS-MIB.mib"
        },
        {
            "mib": "AE-VOICE-STATS-MIB",
            "path": "calix/AE-VOICE-STATS-MIB.mib"
        },
        {
            "mib": "CUMULUS-SNMP-MIB",
            "path": "cumulus/CUMULUS-SNMP-MIB.mib"
        },
        {
            "mib": "ADVA-FSPR7-CAP-MIB",
            "path": "adva/ADVA-FSPR7-CAP-MIB.mib"
        },
        {
            "mib": "VMWARE-ESX-AGENTCAP-MIB",
            "path": "vmware/VMWARE-ESX-AGENTCAP-MIB.mib"
        },
        {
            "mib": "NETTRACK-E3METER-CTR-SNMP-MIB",
            "path": "riedo/NETTRACK-E3METER-CTR-SNMP-MIB.mib"
        },
        {
            "mib": "AT-BOARDS-MIB",
            "path": "allied/AT-BOARDS-MIB.mib"
        },
        {
            "mib": "ELTEX-MES-ISS-ENV-MIB",
            "path": "eltexmes24xx/ELTEX-MES-ISS-ENV-MIB.mib"
        },
        {
            "mib": "ARICENT-POE-MIB",
            "path": "cambium/cnmatrix/ARICENT-POE-MIB.mib"
        },
        {
            "mib": "ZTE-AN-OPTICAL-MODULE-MIB",
            "path": "zte/ZTE-AN-OPTICAL-MODULE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-BASE",
            "path": "nokia/aos7/ALCATEL-IND1-BASE.mib"
        },
        {
            "mib": "ANUE-MIB",
            "path": "ixia/ANUE-MIB.mib"
        },
        {
            "mib": "AXOS-ALARM-MIB",
            "path": "calix/AXOS-ALARM-MIB.mib"
        },
        {
            "mib": "ADVA-FSPR7-CFM-EXTENSION-MIB",
            "path": "adva/ADVA-FSPR7-CFM-EXTENSION-MIB.mib"
        },
        {
            "mib": "T3610-MIB",
            "path": "comet/T3610-MIB.mib"
        },
        {
            "mib": "IFT-SNMP-MIB",
            "path": "infortrend/IFT-SNMP-MIB.mib"
        },
        {
            "mib": "Sentry3-MIB",
            "path": "sentry/Sentry3-MIB.mib"
        },
        {
            "mib": "CT-ELS10-MIB",
            "path": "enterasys/CT-ELS10-MIB.mib"
        },
        {
            "mib": "VMWARE-HEARTBEAT-MIB",
            "path": "vmware/VMWARE-HEARTBEAT-MIB.mib"
        },
        {
            "mib": "ELTEX-MES-ISS-MIB",
            "path": "eltexmes24xx/ELTEX-MES-ISS-MIB.mib"
        },
        {
            "mib": "F5-BIGIP-GLOBAL-MIB",
            "path": "f5/F5-BIGIP-GLOBAL-MIB.mib"
        },
        {
            "mib": "ABBMODULARUPS-MIB",
            "path": "abb/ABB-MODULARUPS-MIB.mib"
        },
        {
            "mib": "ZTE-AN-SMI",
            "path": "zte/ZTE-AN-SMI.mib"
        },
        {
            "mib": "ALCATEL-IND1-BFD-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-BFD-MIB.mib"
        },
        {
            "mib": "AT-BRI-MIB",
            "path": "allied/AT-BRI-MIB.mib"
        },
        {
            "mib": "COLUBRIS-802DOT1X-ACCESS-MIB",
            "path": "hpmsm/COLUBRIS-802DOT1X-MIB.my.mib"
        },
        {
            "mib": "Stulz-WIB8000-MIB",
            "path": "stulz/Stulz-WIB8000-MIB.mib"
        },
        {
            "mib": "Axos-Card-MIB",
            "path": "calix/Axos-Card-MIB.mib"
        },
        {
            "mib": "EdgeSwitch-SWITCHING-MIB",
            "path": "edgeswitch/EdgeSwitch-SWITCHING-MIB.mib"
        },
        {
            "mib": "ONEACCESS-CELLULAR-MIB",
            "path": "oneaccess/ONEACCESS-CELLULAR-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-CHANNEL-MIB",
            "path": "procera/PACKETLOGIC-CHANNEL-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-MED-CAMBIUM-MIB",
            "path": "cambium/cnmatrix/LLDP-EXT-MED-CAMBIUM-MIB.mib"
        },
        {
            "mib": "NETTRACK-E3METER-SNMP-MIB",
            "path": "riedo/NETTRACK-E3METER-SNMP-MIB.mib"
        },
        {
            "mib": "Sentry4-MIB",
            "path": "sentry/Sentry4-MIB.mib"
        },
        {
            "mib": "ELTEX-PHY-MIB",
            "path": "eltexmes24xx/ELTEX-PHY-MIB.mib"
        },
        {
            "mib": "FIREBRICK-BGP-MIB",
            "path": "firebrick/FIREBRICK-BGP-MIB.mib"
        },
        {
            "mib": "CT-FASTPATH-DHCPSERVER-MIB",
            "path": "enterasys/CT-FASTPATH-DHCPSERVER-MIB.mib"
        },
        {
            "mib": "ADVA-FSPR7-DEF-MIB",
            "path": "adva/ADVA-FSPR7-DEF-MIB.mib"
        },
        {
            "mib": "ZTE-AN-SOFTWARE-MIB",
            "path": "zte/ZTE-AN-SOFTWARE-MIB.mib"
        },
        {
            "mib": "VMWARE-HZECC-AGENTCAP-MIB",
            "path": "vmware/VMWARE-HZECC-AGENTCAP-MIB.mib"
        },
        {
            "mib": "SPAGENT-MIB",
            "path": "akcp/SPAGENT-MIB.mib"
        },
        {
            "mib": "AT-CAPABILITIES-MIB",
            "path": "allied/AT-CAPABILITIES-MIB.mib"
        },
        {
            "mib": "COLUBRIS-AAA-CLIENT-MIB",
            "path": "hpmsm/COLUBRIS-AAA-CLIENT-MIB.my.mib"
        },
        {
            "mib": "SNR-ERD-4",
            "path": "snrerd/SNR-ERD-4.mib"
        },
        {
            "mib": "Axos-Ont-MIB",
            "path": "calix/Axos-Ont-MIB.mib"
        },
        {
            "mib": "ELTEX-MES-HWENVIROMENT-MIB",
            "path": "eltexmes23xx/ELTEX-MES-HWENVIROMENT-MIB.mib"
        },
        {
            "mib": "FIREBRICK-CPU-MIB",
            "path": "firebrick/FIREBRICK-CPU-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-BGP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-BGP-MIB.mib"
        },
        {
            "mib": "VMWARE-HZECC-EVENT-MIB",
            "path": "vmware/VMWARE-HZECC-EVENT-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-HW-MIB",
            "path": "procera/PACKETLOGIC-HW-MIB.mib"
        },
        {
            "mib": "CT-FASTPATH-PROTECTED-PORT-MIB",
            "path": "enterasys/CT-FASTPATH-PROTECTED-PORT-MIB.mib"
        },
        {
            "mib": "COLUBRIS-BANDWIDTH-CONTROL-MIB",
            "path": "hpmsm/COLUBRIS-BANDWIDTH-CONTROL-MIB.my.mib"
        },
        {
            "mib": "HUAWEI-AAA-MIB",
            "path": "huawei/HUAWEI-AAA-MIB.mib"
        },
        {
            "mib": "BROADCOM-POWER-ETHERNET-MIB",
            "path": "dell/BROADCOM-POWER-ETHERNET-MIB.mib"
        },
        {
            "mib": "INNO-MIB",
            "path": "innovaphone/INNO-MIB.mib"
        },
        {
            "mib": "ONEACCESS-GLOBAL-REG",
            "path": "oneaccess/ONEACCESS-GLOBAL-REG.mib"
        },
        {
            "mib": "ZXR10-MIB",
            "path": "zte/ZXR10-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-HW-SENSORS-MIB",
            "path": "procera/PACKETLOGIC-HW-SENSORS-MIB.mib"
        },
        {
            "mib": "AT-DHCP-MIB",
            "path": "allied/AT-DHCP-MIB.mib"
        },
        {
            "mib": "RNX-UPDU-MIB",
            "path": "riedo/RNX-UPDU-MIB.mib"
        },
        {
            "mib": "CT-FLASH-MIB",
            "path": "enterasys/CT-FLASH-MIB.mib"
        },
        {
            "mib": "ELTEX-SMI-ACTUAL",
            "path": "eltexmes24xx/ELTEX-SMI-ACTUAL.mib"
        },
        {
            "mib": "NETAPP-MIB",
            "path": "netapp/NETAPP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-CDP-MIB",
            "path": "hpmsm/COLUBRIS-CDP-MIB.my.mib"
        },
        {
            "mib": "Axos-System-MIB",
            "path": "calix/Axos-System-MIB.mib"
        },
        {
            "mib": "FIREBRICK-GLOBAL",
            "path": "firebrick/FIREBRICK-GLOBAL-MIB.mib"
        },
        {
            "mib": "ELTEX-MES-PHYSICAL-DESCRIPTION-MIB",
            "path": "eltexmes23xx/ELTEX-MES-PHYSICAL-DESCRIPTION-MIB.mib"
        },
        {
            "mib": "ONEACCESS-SHDSL-MIB",
            "path": "oneaccess/ONEACCESS-SHDSL-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-HW-SFP-MIB",
            "path": "procera/PACKETLOGIC-HW-SFP-MIB.mib"
        },
        {
            "mib": "VIPTELA-APP-ROUTE",
            "path": "viptela/VIPTELA-APP-ROUTE.mib"
        },
        {
            "mib": "ALCATEL-IND1-CAPMAN-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-CAPMAN-MIB.mib"
        },
        {
            "mib": "F5-BIGIP-LOCAL-MIB",
            "path": "f5/F5-BIGIP-LOCAL-MIB.mib"
        },
        {
            "mib": "VMWARE-NSX-AGENTCAP-MIB",
            "path": "vmware/VMWARE-NSX-AGENTCAP-MIB.mib"
        },
        {
            "mib": "DKT-CATV-MIB",
            "path": "dkt/DKT-CATV-MIB.mib"
        },
        {
            "mib": "GREENBONE-PRODUCT-MIB",
            "path": "greenbone/GREENBONE-PRODUCT-MIB.mib"
        },
        {
            "mib": "BROADCOM-REF-MIB",
            "path": "dell/BROADCOM-REF-MIB.mib"
        },
        {
            "mib": "FIREBRICK-IPSEC-MIB",
            "path": "firebrick/FIREBRICK-IPSEC-MIB.mib"
        },
        {
            "mib": "HUAWEI-ACL-MIB",
            "path": "huawei/HUAWEI-ACL-MIB.mib"
        },
        {
            "mib": "ONEACCESS-SYS-MIB",
            "path": "oneaccess/ONEACCESS-SYS-MIB.mib"
        },
        {
            "mib": "ADVA-FSPR7-MIB",
            "path": "adva/ADVA-FSPR7-MIB.mib"
        },
        {
            "mib": "COLUBRIS-CLIENT-TRACKING-MIB",
            "path": "hpmsm/COLUBRIS-CLIENT-TRACKING-MIB.my.mib"
        },
        {
            "mib": "VIPTELA-BFD",
            "path": "viptela/VIPTELA-BFD.mib"
        },
        {
            "mib": "AT-DOS-MIB",
            "path": "allied/AT-DOS-MIB.mib"
        },
        {
            "mib": "TERRA-DEFINITIONS-MIB",
            "path": "terra/TERRA-DEFINITIONS-MIB.mib"
        },
        {
            "mib": "Axos-Trap-MIB",
            "path": "calix/Axos-Trap-MIB.mib"
        },
        {
            "mib": "ELTEX-MES",
            "path": "eltexmes23xx/ELTEX-MES.mib"
        },
        {
            "mib": "DKT-FE-MIB",
            "path": "dkt/DKT-FE-MIB.mib"
        },
        {
            "mib": "CT-FPS-SERVICES-MIB",
            "path": "enterasys/CT-FPS-SERVICES-MIB.mib"
        },
        {
            "mib": "ADVA-FSPR7-MODULE-ENCRYPTION-MIB",
            "path": "adva/ADVA-FSPR7-MODULE-ENCRYPTION-MIB.mib"
        },
        {
            "mib": "NETWORK-APPLIANCE-MIB",
            "path": "netapp/NETWORK-APPLIANCE-MIB.mib"
        },
        {
            "mib": "FIREBRICK-L2TP-MIB",
            "path": "firebrick/FIREBRICK-L2TP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-CONNECTION-LIMITING-MIB",
            "path": "hpmsm/COLUBRIS-CONNECTION-LIMITING-MIB.my.mib"
        },
        {
            "mib": "NMS-IF-MIB",
            "path": "pbn/NMS-IF-MIB.MIB.mib"
        },
        {
            "mib": "ISILON-MIB",
            "path": "emc/ISILON-MIB.mib"
        },
        {
            "mib": "TERRA-PRODUCTS-MIB",
            "path": "terra/TERRA-PRODUCTS-MIB.mib"
        },
        {
            "mib": "HUAWEI-ALARM-MIB",
            "path": "huawei/HUAWEI-ALARM-MIB.mib"
        },
        {
            "mib": "CALIX-PRODUCT-MIB",
            "path": "calix/CALIX-PRODUCT-MIB.mib"
        },
        {
            "mib": "AT-DS3-MIB",
            "path": "allied/AT-DS3-MIB.mib"
        },
        {
            "mib": "VIPTELA-BRIDGE",
            "path": "viptela/VIPTELA-BRIDGE.mib"
        },
        {
            "mib": "DELL-MM-MIB-SMIv2",
            "path": "dell/DELL-MM-MIB-SMIv2.mib"
        },
        {
            "mib": "VMWARE-NSX-MANAGER-AGENTCAP-MIB",
            "path": "vmware/VMWARE-NSX-MANAGER-AGENTCAP-MIB.mib"
        },
        {
            "mib": "ELTEX-SMI-ACTUAL",
            "path": "eltexmes23xx/ELTEX-SMI-ACTUAL.mib"
        },
        {
            "mib": "F5-BIGIP-SYSTEM-MIB",
            "path": "f5/F5-BIGIP-SYSTEM-MIB.mib"
        },
        {
            "mib": "COLUBRIS-DEVICE-DOT1X-MIB",
            "path": "hpmsm/COLUBRIS-DEVICE-DOT1X-MIB.my.mib"
        },
        {
            "mib": "FIREBRICK-MIB",
            "path": "firebrick/FIREBRICK-MIB.mib"
        },
        {
            "mib": "CT-HSIMPHYS-MIB",
            "path": "enterasys/CT-HSIMPHYS-MIB.mib"
        },
        {
            "mib": "PBI-4000P-5000P-MIB",
            "path": "pbi/PBI-4000P-5000P-MIB.mib"
        },
        {
            "mib": "DKT-GE-MIB",
            "path": "dkt/DKT-GE-MIB.mib"
        },
        {
            "mib": "CALIX-SMI",
            "path": "calix/CALIX-SMI.mib"
        },
        {
            "mib": "NMS-LLDP-MIB",
            "path": "pbn/NMS-LLDP.mib"
        },
        {
            "mib": "ALCATEL-IND1-CHASSIS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-CHASSIS-MIB.mib"
        },
        {
            "mib": "AT-ENVMON-MIB",
            "path": "allied/AT-ENVMON-MIB.mib"
        },
        {
            "mib": "VIPTELA-DOT1X",
            "path": "viptela/VIPTELA-DOT1X.mib"
        },
        {
            "mib": "F5-BIGIP-WAM-MIB",
            "path": "f5/F5-BIGIP-WAM-MIB.mib"
        },
        {
            "mib": "FIREBRICK-MONITORING",
            "path": "firebrick/FIREBRICK-MONITORING.mib"
        },
        {
            "mib": "TERRA-sdi410C-MIB",
            "path": "terra/TERRA-sdi410C-MIB.mib"
        },
        {
            "mib": "COLUBRIS-DEVICE-EVENT-MIB",
            "path": "hpmsm/COLUBRIS-DEVICE-EVENT-MIB.my.mib"
        },
        {
            "mib": "Juniper-Accounting-CONF",
            "path": "junose/broken/Juniper-Accounting-CONF.mib"
        },
        {
            "mib": "CT-PIC-MIB",
            "path": "enterasys/CT-PIC-MIB.mib"
        },
        {
            "mib": "PBI-MAIN-MIB",
            "path": "pbi/PBI-MAIN-MIB.mib"
        },
        {
            "mib": "DKT-GENERIC-MIB",
            "path": "dkt/DKT-GENERIC-MIB.mib"
        },
        {
            "mib": "VMWARE-NSX-MANAGER-MIB",
            "path": "vmware/VMWARE-NSX-MANAGER-MIB.mib"
        },
        {
            "mib": "MARVELL-POE-MIB",
            "path": "eltexmes23xx/MARVELL-POE-MIB.mib"
        },
        {
            "mib": "HUAWEI-ALARM-RELIABILITY-MIB",
            "path": "huawei/HUAWEI-ALARM-RELIABILITY-MIB.mib"
        },
        {
            "mib": "F5-COMMON-SMI-MIB",
            "path": "f5/F5-COMMON-SMI-MIB.mib"
        },
        {
            "mib": "NMS-MEMORY-POOL-MIB",
            "path": "pbn/NMS-MEMORY-POOL-MIB.MIB.mib"
        },
        {
            "mib": "LAN",
            "path": "peplink/LAN.mib"
        },
        {
            "mib": "ALCATEL-IND1-CONFIG-MGR-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-CONFIG-MGR-MIB.mib"
        },
        {
            "mib": "VIPTELA-GLOBAL",
            "path": "viptela/VIPTELA-GLOBAL.mib"
        },
        {
            "mib": "FIREBRICK-RUNSTATS-MIB",
            "path": "firebrick/FIREBRICK-RUNSTATS-MIB.mib"
        },
        {
            "mib": "DKT-MIB",
            "path": "dkt/DKT-MIB.mib"
        },
        {
            "mib": "AT-EPSR-MIB",
            "path": "allied/AT-EPSR-MIB.mib"
        },
        {
            "mib": "CALIX-EOAM-EXT-MIB",
            "path": "calix/CALIX-SOAM-EXT-MIB.mib"
        },
        {
            "mib": "PBI-MGSIGNALCHARACTERISTICS-MIB",
            "path": "pbi/PBI-MGSIGNALCHARACTERISTICS-MIB.mib"
        },
        {
            "mib": "DELL-MM-MIB",
            "path": "dell/DELL-MM-MIB.mib"
        },
        {
            "mib": "HUAWEI-APS-MIB",
            "path": "huawei/HUAWEI-APS-MIB.mib"
        },
        {
            "mib": "RADLAN-DEVICEPARAMS-MIB",
            "path": "eltexmes23xx/RADLAN-DEVICEPARAMS-MIB.mib"
        },
        {
            "mib": "COLUBRIS-DEVICE-IF-MIB",
            "path": "hpmsm/COLUBRIS-DEVICE-IF-MIB.my.mib"
        },
        {
            "mib": "ADVA-FSPR7-PM-MIB",
            "path": "adva/ADVA-FSPR7-PM-MIB.mib"
        },
        {
            "mib": "TERRA-sdi480-MIB",
            "path": "terra/TERRA-sdi480-MIB.mib"
        },
        {
            "mib": "NMS-PROCESS-MIB",
            "path": "pbn/NMS-PROCESS-MIB.MIB.mib"
        },
        {
            "mib": "Juniper-ATM-1483-Profile-CONF",
            "path": "junose/broken/Juniper-ATM-1483-Profile-CONF.mib"
        },
        {
            "mib": "F5-EM-MIB",
            "path": "f5/F5-EM-MIB.mib"
        },
        {
            "mib": "VIPTELA-HARDWARE",
            "path": "viptela/VIPTELA-HARDWARE.mib"
        },
        {
            "mib": "AT-ETH-MIB",
            "path": "allied/AT-ETH-MIB.mib"
        },
        {
            "mib": "CT-PRIORITY-CLASSIFY-MIB",
            "path": "enterasys/CT-PRIORITY-CLASSIFY-MIB.mib"
        },
        {
            "mib": "FIREBRICK-VOIP-MIB",
            "path": "firebrick/FIREBRICK-VOIP-MIB.mib"
        },
        {
            "mib": "TERRA-SMI",
            "path": "terra/TERRA-SMI.mib"
        },
        {
            "mib": "HUAWEI-ASPF-EUDM-MIB",
            "path": "huawei/HUAWEI-ASPF-EUDM-MIB.mib"
        },
        {
            "mib": "DKT-RMON-MIB",
            "path": "dkt/DKT-RMON-MIB.mib"
        },
        {
            "mib": "E5-110-AESCOMMON-MIB",
            "path": "calix/E5-110-AESCOMMON-MIB.mib"
        },
        {
            "mib": "PBI-MGSYSTEM-MIB",
            "path": "pbi/PBI-MGSYSTEM-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-BGP4-V2-MIB",
            "path": "dell/DELL-NETWORKING-BGP4-V2-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DA-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-DA-MIB.mib"
        },
        {
            "mib": "Juniper-Dos-Protection-Platform-CONF",
            "path": "junose/broken/Juniper-Dos-Protection-Platform-CONF.mib"
        },
        {
            "mib": "RADLAN-File",
            "path": "eltexmes23xx/RADLAN-File.mib"
        },
        {
            "mib": "VMWARE-NSX-MIB",
            "path": "vmware/VMWARE-NSX-MIB.mib"
        },
        {
            "mib": "COLUBRIS-DEVICE-MIB",
            "path": "hpmsm/COLUBRIS-DEVICE-MIB.my.mib"
        },
        {
            "mib": "PK-SOFTWARE-APPLIANCE-V2",
            "path": "primekey/PK-SOFTWARE-APPLIANCE-V2.mib"
        },
        {
            "mib": "E5-110-AS-ATM-MIB",
            "path": "calix/E5-110-AS-ATM-MIB.mib"
        },
        {
            "mib": "NMS-QOS-PIB-MIB",
            "path": "pbn/NMS-QOS-PIB-MIB.MIB.mib"
        },
        {
            "mib": "VIPTELA-OMP",
            "path": "viptela/VIPTELA-OMP.mib"
        },
        {
            "mib": "DELL-NETWORKING-BMP-MIB",
            "path": "dell/DELL-NETWORKING-BMP-MIB.mib"
        },
        {
            "mib": "IDKT-F2-MIB",
            "path": "dkt/IDKT-F2-MIB.mib"
        },
        {
            "mib": "VMWARE-OBSOLETE-MIB",
            "path": "vmware/VMWARE-OBSOLETE-MIB.mib"
        },
        {
            "mib": "CTELS100-NG-MIB",
            "path": "enterasys/CTELS100-NG-MIB.mib"
        },
        {
            "mib": "DigiPower-PDU-MIB",
            "path": "digipower/DigiPower-PDU-MIB.mib"
        },
        {
            "mib": "Juniper-DOS-PROTECTION-PLATFORM-MIB",
            "path": "junose/broken/Juniper-DOS-PROTECTION-PLATFORM-MIB.mib"
        },
        {
            "mib": "AT-FILE-MIB",
            "path": "allied/AT-FILE-MIB.mib"
        },
        {
            "mib": "F5-PLATFORM-STATS-MIB",
            "path": "f5/F5-PLATFORM-STATS-MIB.mib"
        },
        {
            "mib": "CPI-UNITY-MIB",
            "path": "chatsworth/CPI-UNITY-MIB.mib"
        },
        {
            "mib": "HUAWEI-ATK-EUDM-MIB",
            "path": "huawei/HUAWEI-ATK-EUDM-MIB.mib"
        },
        {
            "mib": "PRIMEKEY-APPLIANCE-MIB",
            "path": "primekey/PRIMEKEY-APPLIANCE-MIB.mib"
        },
        {
            "mib": "E5-110-IESCOMMON-MIB",
            "path": "calix/E5-110-IESCOMMON-MIB.mib"
        },
        {
            "mib": "DISMUNTELv00-MIB",
            "path": "himoinsa/DISMUNTELv00-MIB.mib"
        },
        {
            "mib": "CTFPS-MIB",
            "path": "enterasys/CTFPS-MIB.mib"
        },
        {
            "mib": "ADVA-FSPR7-TC-MIB",
            "path": "adva/ADVA-FSPR7-TC-MIB.mib"
        },
        {
            "mib": "RADLAN-HWENVIROMENT",
            "path": "eltexmes23xx/RADLAN-HWENVIROMENT.mib"
        },
        {
            "mib": "ALCATEL-IND1-DCBX-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-DCBX-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-MAC-AUTHENTICATION-MIB",
            "path": "foundry/FOUNDRY-SN-MAC-AUTHENTICATION-MIB.mib"
        },
        {
            "mib": "COLUBRIS-DEVICE-WDS-MIB",
            "path": "hpmsm/COLUBRIS-DEVICE-WDS-MIB.my.mib"
        },
        {
            "mib": "VIPTELA-OPER-BGP",
            "path": "viptela/VIPTELA-OPER-BGP.mib"
        },
        {
            "mib": "VMWARE-PRODUCTS-MIB",
            "path": "vmware/VMWARE-PRODUCTS-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-NIC-MIB",
            "path": "procera/PACKETLOGIC-NIC-MIB.mib"
        },
        {
            "mib": "ADVA-MIB",
            "path": "adva/ADVA-MIB.mib"
        },
        {
            "mib": "SKYHIGHSECURITY-SMI",
            "path": "skyhigh/SKYHIGHSECURITY-SMI.mib"
        },
        {
            "mib": "CTFRAMER-CONFIG-MIB",
            "path": "enterasys/CTFRAMER-CONFIG-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-MIB",
            "path": "procera/PACKETLOGIC-MIB.mib"
        },
        {
            "mib": "HIMOINSAv14-MIB",
            "path": "himoinsa/HIMOINSAv14-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-MAC-VLAN-MIB",
            "path": "foundry/FOUNDRY-SN-MAC-VLAN-MIB.mib"
        },
        {
            "mib": "NSCRTV-ROOT",
            "path": "edfa/NSCRTV-ROOT.mib"
        },
        {
            "mib": "RADLAN-MIB",
            "path": "eltexmes23xx/RADLAN-MIB.mib"
        },
        {
            "mib": "AT-FIREWALL-MIB",
            "path": "allied/AT-FIREWALL-MIB.mib"
        },
        {
            "mib": "COLUBRIS-DEVICE-WIRELESS-MIB",
            "path": "hpmsm/COLUBRIS-DEVICE-WIRELESS-MIB.my.mib"
        },
        {
            "mib": "AOS-CORE-CONDITION-MIB",
            "path": "adva/AOS-CORE-CONDITION-MIB.mib"
        },
        {
            "mib": "SKYHIGHSECURITY-SWG-MIB",
            "path": "skyhigh/SKYHIGHSECURITY-SWG-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-BPSTATS-MIB",
            "path": "dell/DELL-NETWORKING-BPSTATS-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-MRP-MIB",
            "path": "foundry/FOUNDRY-SN-MRP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DEVICES",
            "path": "nokia/aos7/ALCATEL-IND1-DEVICES.mib"
        },
        {
            "mib": "BLADETYPE2-ACL-MIB",
            "path": "hp/BLADETYPE2-ACL-MIB.mib"
        },
        {
            "mib": "NMS-SMI",
            "path": "pbn/NMS-SMI.mib"
        },
        {
            "mib": "HUAWEI-ATK-MIB",
            "path": "huawei/HUAWEI-ATK-MIB.mib"
        },
        {
            "mib": "SNWL-COMMON-MIB",
            "path": "sonicwall/SNWL-COMMON-MIB.mib"
        },
        {
            "mib": "Juniper-HTTP-MIB",
            "path": "junose/broken/Juniper-HTTP-MIB.mib"
        },
        {
            "mib": "VIPTELA-OPER-MULTICAST",
            "path": "viptela/VIPTELA-OPER-MULTICAST.mib"
        },
        {
            "mib": "AOS-CORE-FACILITY-MIB",
            "path": "adva/AOS-CORE-FACILITY-MIB.mib"
        },
        {
            "mib": "VMWARE-RESOURCES-MIB",
            "path": "vmware/VMWARE-RESOURCES-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-OVERVIEW-MIB",
            "path": "procera/PACKETLOGIC-OVERVIEW-MIB.mib"
        },
        {
            "mib": "HUAWEI-ATM-MIB",
            "path": "huawei/HUAWEI-ATM-MIB.mib"
        },
        {
            "mib": "EQL-DCB-MIB",
            "path": "equallogic/EQL-DCB-MIB.mib"
        },
        {
            "mib": "E5-110-MIB",
            "path": "calix/E5-110-MIB.mib"
        },
        {
            "mib": "BLUECOAT-LICENSE-MIB",
            "path": "bluecoat/BLUECOAT-LICENSE-MIB.mib"
        },
        {
            "mib": "AT-FLASH-MIB",
            "path": "allied/AT-FLASH-MIB.mib"
        },
        {
            "mib": "SNWL-SSLVPN-MIB",
            "path": "sonicwall/SNWL-SSLVPN-MIB.mib"
        },
        {
            "mib": "RADLAN-PHY-MIB",
            "path": "eltexmes23xx/RADLAN-PHY-MIB.mib"
        },
        {
            "mib": "CTIF-EXT-MIB",
            "path": "enterasys/CTIF-EXT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DHCPV6-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-DHCPV6-MIB.mib"
        },
        {
            "mib": "COLUBRIS-802DOT11-MIB",
            "path": "hpmsm/COLUBRIS-IEEE802DOT11.my.mib"
        },
        {
            "mib": "DELL-NETWORKING-CHASSIS-MIB",
            "path": "dell/DELL-NETWORKING-CHASSIS-MIB.mib"
        },
        {
            "mib": "EQL-LLDP-MIB",
            "path": "equallogic/EQL-LLDP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-LICENSE-MIB",
            "path": "hpmsm/COLUBRIS-LICENSE-MIB.my.mib"
        },
        {
            "mib": "BLADETYPE2-NETWORK-MIB",
            "path": "hp/BLADETYPE2-NETWORK-MIB.mib"
        },
        {
            "mib": "GUDEADS-ESB7213-MIB",
            "path": "gude/GUDEADS-ESB7213-MIB.mib"
        },
        {
            "mib": "VIPTELA-OPER-OSPF",
            "path": "viptela/VIPTELA-OPER-OSPF.mib"
        },
        {
            "mib": "VMWARE-ROOT-MIB",
            "path": "vmware/VMWARE-ROOT-MIB.mib"
        },
        {
            "mib": "E5-110-TRAPS-MIB",
            "path": "calix/E5-110-TRAPS-MIB.mib"
        },
        {
            "mib": "AT-IGMP-MIB",
            "path": "allied/AT-IGMP-MIB.mib"
        },
        {
            "mib": "NMS-TC",
            "path": "pbn/NMS-TC.mib"
        },
        {
            "mib": "SONICWALL-FIREWALL-IP-STATISTICS-MIB",
            "path": "sonicwall/SONICWALL-FIREWALL-IP-STATISTICS-MIB.mib"
        },
        {
            "mib": "Juniper-IGMP-CONF",
            "path": "junose/broken/Juniper-IGMP-CONF.mib"
        },
        {
            "mib": "PACKETLOGIC-RAID-MIB",
            "path": "procera/PACKETLOGIC-RAID-MIB.mib"
        },
        {
            "mib": "AOS-DOMAIN-OTN-PM-MIB",
            "path": "adva/AOS-DOMAIN-OTN-PM-MIB.mib"
        },
        {
            "mib": "BLUECOAT-MIB",
            "path": "bluecoat/BLUECOAT-MIB.mib"
        },
        {
            "mib": "HUAWEI-BASE-TRAP-MIB",
            "path": "huawei/HUAWEI-BASE-TRAP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-MAINTENANCE-MIB",
            "path": "hpmsm/COLUBRIS-MAINTENANCE-MIB.my.mib"
        },
        {
            "mib": "E5-111-AESCOMMON-MIB",
            "path": "calix/E5-111-AESCOMMON-MIB.mib"
        },
        {
            "mib": "GUDEADS-ESB7214-MIB",
            "path": "gude/GUDEADS-ESB7214-MIB.mib"
        },
        {
            "mib": "AT-INSTALL-MIB",
            "path": "allied/AT-INSTALL-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DHL-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-DHL-MIB.mib"
        },
        {
            "mib": "COLUBRIS-PRODUCTS-MIB",
            "path": "hpmsm/COLUBRIS-PRODUCTS-MIB.my.mib"
        },
        {
            "mib": "PACKETLOGIC-SNOOPER-DHCP-MIB",
            "path": "procera/PACKETLOGIC-SNOOPER-DHCP-MIB.mib"
        },
        {
            "mib": "CTINB-MIB",
            "path": "enterasys/CTINB-MIB.mib"
        },
        {
            "mib": "EQLACCESS-MIB",
            "path": "equallogic/EQLACCESS-MIB.mib"
        },
        {
            "mib": "BLADETYPE2-PHYSICAL-MIB",
            "path": "hp/BLADETYPE2-PHYSICAL-MIB.mib"
        },
        {
            "mib": "VMWARE-SRM-EVENT-MIB",
            "path": "vmware/VMWARE-SRM-EVENT-MIB.mib"
        },
        {
            "mib": "RADLAN-Physicaldescription-MIB",
            "path": "eltexmes23xx/RADLAN-Physicaldescription-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-DCB-MIB",
            "path": "dell/DELL-NETWORKING-DCB-MIB.mib"
        },
        {
            "mib": "CM-ALARM-MIB",
            "path": "adva/CM-ALARM-MIB.mib"
        },
        {
            "mib": "SONICWALL-SMI",
            "path": "sonicwall/SONICWALL-SMI.mib"
        },
        {
            "mib": "CTINB2-MIB",
            "path": "enterasys/CTINB2-MIB.mib"
        },
        {
            "mib": "PBN-MIB",
            "path": "pbn/PBN-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DOT3-OAM-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-DOT3-OAM-MIB.mib"
        },
        {
            "mib": "HUAWEI-BFD-MIB",
            "path": "huawei/HUAWEI-BFD-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SEGMENT-MIB",
            "path": "bluecoat/BLUECOAT-SEGMENT-MIB.mib"
        },
        {
            "mib": "GUDEADS-ETS-MIB",
            "path": "gude/GUDEADS-ETS-MIB.mib"
        },
        {
            "mib": "VMWARE-SYSTEM-MIB",
            "path": "vmware/VMWARE-SYSTEM-MIB.mib"
        },
        {
            "mib": "VIPTELA-OPER-SYSTEM",
            "path": "viptela/VIPTELA-OPER-SYSTEM.mib"
        },
        {
            "mib": "Juniper-Internet-CONF",
            "path": "junose/broken/Juniper-Internet-CONF.mib"
        },
        {
            "mib": "AT-INTERFACES-MIB",
            "path": "allied/AT-INTERFACES-MIB.mib"
        },
        {
            "mib": "RADLAN-rndMng",
            "path": "eltexmes23xx/RADLAN-rndMng.mib"
        },
        {
            "mib": "CM-COMMON-MIB",
            "path": "adva/CM-COMMON-MIB.mib"
        },
        {
            "mib": "COLUBRIS-PUBLIC-ACCESS-MIB",
            "path": "hpmsm/COLUBRIS-PUBLIC-ACCESS-MIB.my.mib"
        },
        {
            "mib": "ALCATEL-IND1-DVMRP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-DVMRP-MIB.mib"
        },
        {
            "mib": "CTRMONXT-MIB",
            "path": "enterasys/CTRMONXT-MIB.mib"
        },
        {
            "mib": "EQLAGENT-MIB",
            "path": "equallogic/EQLAGENT-MIB.mib"
        },
        {
            "mib": "PACKETLOGIC-TRAP-MIB",
            "path": "procera/PACKETLOGIC-TRAP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-PUBLIC-ACCESS-RETENTION-MIB",
            "path": "hpmsm/COLUBRIS-PUBLIC-ACCESS-RETENTION-MIB.my.mib"
        },
        {
            "mib": "BLUECOAT-SG-ATTACK-MIB",
            "path": "bluecoat/BLUECOAT-SG-ATTACK-MIB.mib"
        },
        {
            "mib": "VMWARE-TC-MIB",
            "path": "vmware/VMWARE-TC-MIB.mib"
        },
        {
            "mib": "VIPTELA-OPER-VPN",
            "path": "viptela/VIPTELA-OPER-VPN.mib"
        },
        {
            "mib": "E5-111-AS-ATM-MIB",
            "path": "calix/E5-111-AS-ATM-MIB.mib"
        },
        {
            "mib": "CONTROLBOX-TH332-MIB",
            "path": "controlbox/CONTROLBOX-TH332-MIB.mib"
        },
        {
            "mib": "PBN-ROOT",
            "path": "pbn/PBN-ROOT.mib"
        },
        {
            "mib": "AT-ISDN-MIB",
            "path": "allied/AT-ISDN-MIB.mib"
        },
        {
            "mib": "FE-FIREEYE-MIB",
            "path": "trellix/FE-FIREEYE-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-FIB-MIB",
            "path": "dell/DELL-NETWORKING-FIB-MIB.mib"
        },
        {
            "mib": "BLADETYPE2-QOS-MIB",
            "path": "hp/BLADETYPE2-QOS-MIB.mib"
        },
        {
            "mib": "Juniper-IP-MIB",
            "path": "junose/broken/Juniper-IP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-QOS-MIB",
            "path": "hpmsm/COLUBRIS-QOS-MIB.my.mib"
        },
        {
            "mib": "VMW-TUNNEL-SERVER-AGENTCAP-MIB",
            "path": "vmware/VMWARE-TUNNEL-SERVER-AGENTCAP-MIB.mib"
        },
        {
            "mib": "CTRON-ALIAS-MIB",
            "path": "enterasys/CTRON-ALIAS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-E-SERVICE-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-E-SERVICE-MIB.mib"
        },
        {
            "mib": "HUAWEI-BGP-ACCOUNTING-MIB",
            "path": "huawei/HUAWEI-BGP-ACCOUNTING-MIB.mib"
        },
        {
            "mib": "AT-LB-MIB",
            "path": "allied/AT-LB-MIB.mib"
        },
        {
            "mib": "EQLAPPLIANCE-MIB",
            "path": "equallogic/EQLAPPLIANCE-MIB.mib"
        },
        {
            "mib": "CM-ENTITY-MIB",
            "path": "adva/CM-ENTITY-MIB.mib"
        },
        {
            "mib": "ACCEDIAN-SMI",
            "path": "accedian/ACCEDIAN-SMI.mib"
        },
        {
            "mib": "COLUBRIS-SATELLITE-MANAGEMENT-MIB",
            "path": "hpmsm/COLUBRIS-SATELLITE-MANAGEMENT-MIB.my.mib"
        },
        {
            "mib": "MERIDIAN2-MIB",
            "path": "endrun/MERIDIAN2-MIB.mib"
        },
        {
            "mib": "VEEAM-MIB",
            "path": "veeam/VEEAM-MIB.mib"
        },
        {
            "mib": "VIPTELA-POLICY",
            "path": "viptela/VIPTELA-POLICY.mib"
        },
        {
            "mib": "TELTONIKA-MIB",
            "path": "teltonika/TELTONIKA-MIB.mib"
        },
        {
            "mib": "CTRON-AP3000-MIB",
            "path": "enterasys/CTRON-AP3000-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-AUTHENTICATION-MIB",
            "path": "bluecoat/BLUECOAT-SG-AUTHENTICATION-MIB.mib"
        },
        {
            "mib": "E5-111-IESCOMMON-MIB",
            "path": "calix/E5-111-IESCOMMON-MIB.mib"
        },
        {
            "mib": "EQLCONTROLLER-MIB",
            "path": "equallogic/EQLCONTROLLER-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-ERP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-ERP-MIB.mib"
        },
        {
            "mib": "MCAFEE-ATD-MIB",
            "path": "trellix/MCAFEE-ATD-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-FIPSNOOPING-MIB",
            "path": "dell/DELL-NETWORKING-FIPSNOOPING-MIB.mib"
        },
        {
            "mib": "VMWARE-TUNNEL-SERVER-MIB",
            "path": "vmware/VMWARE-TUNNEL-SERVER-MIB.mib"
        },
        {
            "mib": "Juniper-IPv6-Profile-CONF",
            "path": "junose/broken/Juniper-IPv6-Profile-CONF.mib"
        },
        {
            "mib": "HUAWEI-BGP-GR-MIB",
            "path": "huawei/HUAWEI-BGP-GR-MIB.mib"
        },
        {
            "mib": "ACD-ALARM-MIB",
            "path": "accedian/ACD-ALARM-MIB.mib"
        },
        {
            "mib": "SONOMA-MIB",
            "path": "endrun/SONOMA-MIB.mib"
        },
        {
            "mib": "WTI-CONSOLE-MIB",
            "path": "wti/WTI-CONSOLE-MIB.mib"
        },
        {
            "mib": "BLADETYPE2-SWITCH-MIB",
            "path": "hp/BLADETYPE2-SWITCH-MIB.mib"
        },
        {
            "mib": "AT-LOADER-MIB",
            "path": "allied/AT-LOADER-MIB.mib"
        },
        {
            "mib": "CTRON-AppleTalk-ROUTER-MIB",
            "path": "enterasys/CTRON-AppleTalk-ROUTER-MIB.mib"
        },
        {
            "mib": "VMWARE-VA-AGENTCAP-MIB",
            "path": "vmware/VMWARE-VA-AGENTCAP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-SENSOR-MIB",
            "path": "hpmsm/COLUBRIS-SENSOR-MIB.my.mib"
        },
        {
            "mib": "EQLDISK-MIB",
            "path": "equallogic/EQLDISK-MIB.mib"
        },
        {
            "mib": "VIPTELA-SECURITY",
            "path": "viptela/VIPTELA-SECURITY.mib"
        },
        {
            "mib": "BLUECOAT-SG-DISK-MIB",
            "path": "bluecoat/BLUECOAT-SG-DISK-MIB.mib"
        },
        {
            "mib": "WTI-MPC-MIB",
            "path": "wti/WTI-MPC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-EVB-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-EVB-MIB.mib"
        },
        {
            "mib": "BLADETYPE2-TRAP-MIB",
            "path": "hp/BLADETYPE2-TRAP-MIB.mib"
        },
        {
            "mib": "E5-111-MIB",
            "path": "calix/E5-111-MIB.mib"
        },
        {
            "mib": "AT-PAE-MIB",
            "path": "allied/AT-PAE-MIB.mib"
        },
        {
            "mib": "ACD-CFM-MIB",
            "path": "accedian/ACD-CFM-MIB.mib"
        },
        {
            "mib": "MCAFEE-INTRUVERT-SMI",
            "path": "trellix/MCAFEE-INTRUVERT-SMI.mib"
        },
        {
            "mib": "BLUECOAT-SG-FAILOVER-MIB",
            "path": "bluecoat/BLUECOAT-SG-FAILOVER-MIB.mib"
        },
        {
            "mib": "CTRON-APPN-MIB",
            "path": "enterasys/CTRON-APPN-MIB.mib"
        },
        {
            "mib": "HUAWEI-BGP-VPN-MIB",
            "path": "huawei/HUAWEI-BGP-VPN-MIB.mib"
        },
        {
            "mib": "Juniper-MPLS-CONF",
            "path": "junose/broken/Juniper-MPLS-CONF.mib"
        },
        {
            "mib": "TELTONIKA-RUTM-MIB",
            "path": "teltonika/TELTONIKA-RUTM-MIB.mib"
        },
        {
            "mib": "TEMPUSLXUNISON-MIB",
            "path": "endrun/TEMPUSLXUNISON-MIB.mib"
        },
        {
            "mib": "E5-111-TRAPS-MIB",
            "path": "calix/E5-111-TRAPS-MIB.mib"
        },
        {
            "mib": "WTI-POWER-MIB",
            "path": "wti/WTI-POWER-MIB.mib"
        },
        {
            "mib": "EQLGROUP-MIB",
            "path": "equallogic/EQLGROUP-MIB.mib"
        },
        {
            "mib": "ACD-DESC-MIB",
            "path": "accedian/ACD-DESC-MIB.mib"
        },
        {
            "mib": "VMWARE-VC-EVENT-MIB",
            "path": "vmware/VMWARE-VC-EVENT-MIB.mib"
        },
        {
            "mib": "TRELLIX-INTRUVERT-SMI",
            "path": "trellix/TRELLIX-INTRUVERT-SMI.mib"
        },
        {
            "mib": "CTRON-BDG-MIB",
            "path": "enterasys/CTRON-BDG-MIB.mib"
        },
        {
            "mib": "Juniper-Multicast-Router-CONF",
            "path": "junose/broken/Juniper-Multicast-Router-CONF.mib"
        },
        {
            "mib": "AT-PIM-MIB",
            "path": "allied/AT-PIM-MIB.mib"
        },
        {
            "mib": "TELTONIKA-RUTX-MIB",
            "path": "teltonika/TELTONIKA-RUTX-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-FIPS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-FIPS-MIB.mib"
        },
        {
            "mib": "EQLINTERNAL-MIB",
            "path": "equallogic/EQLINTERNAL-MIB.mib"
        },
        {
            "mib": "E5-120-AS-ATM-MIB",
            "path": "calix/E5-120-AS-ATM-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-HEALTHCHECK-MIB",
            "path": "bluecoat/BLUECOAT-SG-HEALTHCHECK-MIB.mib"
        },
        {
            "mib": "TRELLIX-INTRUVERT-TC",
            "path": "trellix/TRELLIX-INTRUVERT-TC.mib"
        },
        {
            "mib": "DELL-NETWORKING-FPSTATS-MIB",
            "path": "dell/DELL-NETWORKING-FPSTATS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BLS-MIB",
            "path": "huawei/HUAWEI-BLS-MIB.mib"
        },
        {
            "mib": "GAM-PRODUCTS-MIB",
            "path": "positron/GAM-PRODUCTS-MIB.mib"
        },
        {
            "mib": "BLADETYPE4-NETWORK-MIB",
            "path": "hp/BLADETYPE4-NETWORK-MIB.mib"
        },
        {
            "mib": "CM-IP-MIB",
            "path": "adva/CM-IP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-SMI",
            "path": "hpmsm/COLUBRIS-SMI.my.mib"
        },
        {
            "mib": "CTRON-BRIDGE-MIB",
            "path": "enterasys/CTRON-BRIDGE-MIB.mib"
        },
        {
            "mib": "AT-PING-MIB",
            "path": "allied/AT-PING-MIB.mib"
        },
        {
            "mib": "VIPTELA-TRAPS",
            "path": "viptela/VIPTELA-TRAPS.mib"
        },
        {
            "mib": "BLUECOAT-SG-HEALTHMONITOR-MIB",
            "path": "bluecoat/BLUECOAT-SG-HEALTHMONITOR-MIB.mib"
        },
        {
            "mib": "E5-120-IESCOMMON-MIB",
            "path": "calix/E5-120-IESCOMMON-MIB.mib"
        },
        {
            "mib": "COLUBRIS-SYSLOG-MIB",
            "path": "hpmsm/COLUBRIS-SYSLOG-MIB.my.mib"
        },
        {
            "mib": "VMWARE-VCHA-MIB",
            "path": "vmware/VMWARE-VCHA-MIB.mib"
        },
        {
            "mib": "VIPTELA-WLAN",
            "path": "viptela/VIPTELA-WLAN.mib"
        },
        {
            "mib": "DELL-NETWORKING-IF-EXTENSION-MIB",
            "path": "dell/DELL-NETWORKING-IF-EXTENSION-MIB.mib"
        },
        {
            "mib": "CM-FACILITY-MIB",
            "path": "adva/CM-FACILITY-MIB.mib"
        },
        {
            "mib": "ALPHA-RESOURCE-MIB",
            "path": "alpha/ALPHA-CONVERTER-SYS-MIB.mib"
        },
        {
            "mib": "TREDESS-FS-MIB",
            "path": "tredess/TREDESS-FS-MIB.mib"
        },
        {
            "mib": "EQLIPADDR-MIB",
            "path": "equallogic/EQLIPADDR-MIB.mib"
        },
        {
            "mib": "CTRON-BUS-MIB",
            "path": "enterasys/CTRON-BUS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-COPS-MIB",
            "path": "huawei/HUAWEI-BRAS-COPS-MIB.mib"
        },
        {
            "mib": "GAM-SYSUTIL-MIB",
            "path": "positron/GAM-SYSUTIL-MIB.mib"
        },
        {
            "mib": "AT-PRI-MIB",
            "path": "allied/AT-PRI-MIB.mib"
        },
        {
            "mib": "FIBROLAN-COMMON-MIB",
            "path": "fibrolan/FIBROLAN-COMMON-MIB.mib"
        },
        {
            "mib": "ACD-DISCOVERY-MIB",
            "path": "accedian/ACD-DISCOVERY-MIB.mib"
        },
        {
            "mib": "VIPTELA-WWAN",
            "path": "viptela/VIPTELA-WWAN.mib"
        },
        {
            "mib": "VMWARE-VCOPS-EVENT-MIB",
            "path": "vmware/VMWARE-VCOPS-EVENT-MIB.mib"
        },
        {
            "mib": "EQLIPSEC-MIB",
            "path": "equallogic/EQLIPSEC-MIB.mib"
        },
        {
            "mib": "COLUBRIS-SYSTEM-MIB",
            "path": "hpmsm/COLUBRIS-SYSTEM-MIB.my.mib"
        },
        {
            "mib": "DELL-NETWORKING-ISIS-MIB",
            "path": "dell/DELL-NETWORKING-ISIS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-DPI-MIB",
            "path": "huawei/HUAWEI-BRAS-DPI-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-ICAP-MIB",
            "path": "bluecoat/BLUECOAT-SG-ICAP-MIB.mib"
        },
        {
            "mib": "CM-PERFORMANCE-MIB",
            "path": "adva/CM-PERFORMANCE-MIB.mib"
        },
        {
            "mib": "CTRON-CDP-MIB",
            "path": "enterasys/CTRON-CDP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-GLOBALROUTETABLE-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-GLOBALROUTETABLE-MIB.mib"
        },
        {
            "mib": "GAM-TC",
            "path": "positron/GAM-TC.mib"
        },
        {
            "mib": "ALPHA-NOTIFICATION-MIB",
            "path": "alpha/ALPHA-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "E5-120-MIB",
            "path": "calix/E5-120-MIB.mib"
        },
        {
            "mib": "AT-PRODUCT-MIB",
            "path": "allied/AT-PRODUCT-MIB.mib"
        },
        {
            "mib": "PEAKFLOW-DOS-MIB",
            "path": "arbornet/ARBORNET-PEAKFLOW-DOS-MIB.mib"
        },
        {
            "mib": "CITRIX-NetScaler-SD-WAN-MIB",
            "path": "citrix/CITRIX-NetScaler-SD-WAN-MIB.mib"
        },
        {
            "mib": "VMWARE-VMINFO-MIB",
            "path": "vmware/VMWARE-VMINFO-MIB.mib"
        },
        {
            "mib": "TRELLIX-SENSOR-CONF-MIB",
            "path": "trellix/TRELLIX-SENSOR-CONF-MIB.mib"
        },
        {
            "mib": "FIBROLAN-DEVICE-MIB",
            "path": "fibrolan/FIBROLAN-DEVICE-MIB.mib"
        },
        {
            "mib": "EQLISCSI-MIB",
            "path": "equallogic/EQLISCSI-MIB.mib"
        },
        {
            "mib": "CM-PROTECTION-MIB",
            "path": "adva/CM-PROTECTION-MIB.mib"
        },
        {
            "mib": "COLUBRIS-TC",
            "path": "hpmsm/COLUBRIS-TC.my.mib"
        },
        {
            "mib": "CTRON-CHASSIS-MIB",
            "path": "enterasys/CTRON-CHASSIS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-GRE-MIB",
            "path": "huawei/HUAWEI-BRAS-GRE-MIB.mib"
        },
        {
            "mib": "E5-120-TRAPS-MIB",
            "path": "calix/E5-120-TRAPS-MIB.mib"
        },
        {
            "mib": "ACD-FILTER-MIB",
            "path": "accedian/ACD-FILTER-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-GVRP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-GVRP-MIB.mib"
        },
        {
            "mib": "POSITRON-SMI",
            "path": "positron/POSITRON-SMI.mib"
        },
        {
            "mib": "AT-PVSTPM-MIB",
            "path": "allied/AT-PVSTPM-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-LINK-AGGREGATION-MIB",
            "path": "dell/DELL-NETWORKING-LINK-AGGREGATION-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-POLICY-MIB",
            "path": "bluecoat/BLUECOAT-SG-POLICY-MIB.mib"
        },
        {
            "mib": "PEAKFLOW-SP-MIB",
            "path": "arbornet/ARBORNET-PEAKFLOW-SP-MIB.mib"
        },
        {
            "mib": "MAS-MIB-SMIV2-MIB",
            "path": "citrix/MAS-MIB-SMIV2-MIB.mib"
        },
        {
            "mib": "CM-REDUNDANCY-MIB",
            "path": "adva/CM-REDUNDANCY-MIB.mib"
        },
        {
            "mib": "BLADETYPE5-NETWORK-MIB",
            "path": "hp/BLADETYPE5-NETWORK-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-IFNET-MIB",
            "path": "huawei/HUAWEI-BRAS-IFNET-MIB.mib"
        },
        {
            "mib": "E5-121-AS-ATM-MIB",
            "path": "calix/E5-121-AS-ATM-MIB.mib"
        },
        {
            "mib": "ALPHA-RECTIFIER-SYS-MIB",
            "path": "alpha/ALPHA-RECTIFIER-SYS-MIB.mib"
        },
        {
            "mib": "TRELLIX-SENSOR-PERF-MIB",
            "path": "trellix/TRELLIX-SENSOR-PERF-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-ACM110",
            "path": "fibrolan/FIBROLAN-MIB-ACM110.mib"
        },
        {
            "mib": "VMWARE-VRNI-AGENTCAP-MIB",
            "path": "vmware/VMWARE-VRNI-AGENTCAP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-HA-VLAN-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-HA-VLAN-MIB.mib"
        },
        {
            "mib": "PEAKFLOW-TMS-MIB",
            "path": "arbornet/ARBORNET-PEAKFLOW-TMS-MIB.mib"
        },
        {
            "mib": "AT-QOS-MIB",
            "path": "allied/AT-QOS-MIB.mib"
        },
        {
            "mib": "COLUBRIS-TOOLS-MIB",
            "path": "hpmsm/COLUBRIS-TOOLS-MIB.my.mib"
        },
        {
            "mib": "ACD-PAA-MIB",
            "path": "accedian/ACD-PAA-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-MAC-NOTIFICATION-MIB",
            "path": "dell/DELL-NETWORKING-MAC-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-PROXY-MIB",
            "path": "bluecoat/BLUECOAT-SG-PROXY-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-HEALTH-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-HEALTH-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-IPTN-MIB",
            "path": "huawei/HUAWEI-BRAS-IPTN-MIB.mib"
        },
        {
            "mib": "EQLMEMBER-MIB",
            "path": "equallogic/EQLMEMBER-MIB.mib"
        },
        {
            "mib": "COLUBRIS-USAGE-INFORMATION-MIB",
            "path": "hpmsm/COLUBRIS-USAGE-INFORMATION-MIB.my.mib"
        },
        {
            "mib": "TRELLIX-SENSOR-SMI",
            "path": "trellix/TRELLIX-SENSOR-SMI.mib"
        },
        {
            "mib": "CTRON-COMMON-MIB",
            "path": "enterasys/CTRON-COMMON-MIB.mib"
        },
        {
            "mib": "AT-SMI-MIB",
            "path": "allied/AT-SMI-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-DEVICECONVERTER",
            "path": "fibrolan/FIBROLAN-MIB-DEVICECONVERTERS.mib"
        },
        {
            "mib": "ALPHA-RESOURCE-MIB",
            "path": "alpha/ALPHA-RESOURCE-MIB.mib"
        },
        {
            "mib": "ACD-POLICY-MIB",
            "path": "accedian/ACD-POLICY-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-OPENFLOW-MIB",
            "path": "dell/DELL-NETWORKING-OPENFLOW-MIB.mib"
        },
        {
            "mib": "ARBOR-SMI",
            "path": "arbornet/ARBORNET-SMI.mib"
        },
        {
            "mib": "QTECH-MIB",
            "path": "qtech/QTECH-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IGMP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IGMP-MIB.mib"
        },
        {
            "mib": "E5-121-IESCOMMON-MIB",
            "path": "calix/E5-121-IESCOMMON-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-SENSOR-MIB",
            "path": "bluecoat/BLUECOAT-SG-SENSOR-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-L2TP-MIB",
            "path": "huawei/HUAWEI-BRAS-L2TP-MIB.mib"
        },
        {
            "mib": "ADMIN-MASTER-MIB",
            "path": "fs/ADMIN-MASTER-MIB.mib"
        },
        {
            "mib": "AT-STACK-MIB",
            "path": "allied/AT-STACK-MIB.mib"
        },
        {
            "mib": "SMARTNODE-MIB",
            "path": "patton/SMARTNODE-MIB.mib"
        },
        {
            "mib": "BLADETYPE6-NETWORK-MIB",
            "path": "hp/BLADETYPE6-NETWORK-MIB.mib"
        },
        {
            "mib": "CM-SA-MIB",
            "path": "adva/CM-SA-MIB.mib"
        },
        {
            "mib": "COLUBRIS-USER-ACCOUNT-MIB",
            "path": "hpmsm/COLUBRIS-USER-ACCOUNT-MIB.my.mib"
        },
        {
            "mib": "ACD-PORT-MIB",
            "path": "accedian/ACD-PORT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-INTERSWITCH-PROTOCOL-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-INTERSWITCH-PROTOCOL-MIB.mib"
        },
        {
            "mib": "CDATA-COMMON-SMI",
            "path": "cdata/CDATA-COMMON-SMI.mib"
        },
        {
            "mib": "CTRON-CSMACD-MIB",
            "path": "enterasys/CTRON-CSMACD-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-PRODUCTS-MIB",
            "path": "dell/DELL-NETWORKING-PRODUCTS-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-GBE-MCM1000",
            "path": "fibrolan/FIBROLAN-MIB-GBE-MCM1000.mib"
        },
        {
            "mib": "NS-ROOT-MIB",
            "path": "citrix/NS-MIB-SMIV2-MIB.mib"
        },
        {
            "mib": "AT-SWITCH-MIB",
            "path": "allied/AT-SWITCH-MIB.mib"
        },
        {
            "mib": "EQLNAS-MIB",
            "path": "equallogic/EQLNAS-MIB.mib"
        },
        {
            "mib": "PACKETFLUX-SMI",
            "path": "packetflux/PACKETFLUX-SMI.mib"
        },
        {
            "mib": "COLUBRIS-USER-SESSION-MIB",
            "path": "hpmsm/COLUBRIS-USER-SESSION-MIB.my.mib"
        },
        {
            "mib": "VMWARE-VRNI-MIB",
            "path": "vmware/VMWARE-VRNI-MIB.mib"
        },
        {
            "mib": "Argus-MIB",
            "path": "alpha/Argus-MIB.mib"
        },
        {
            "mib": "MIB",
            "path": "citrix/SDX-MIB-SMIV2-MIB.mib"
        },
        {
            "mib": "ACD-REGULATOR-MIB",
            "path": "accedian/ACD-REGULATOR-MIB.mib"
        },
        {
            "mib": "CTRON-DCM-MIB",
            "path": "enterasys/CTRON-DCM-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-USAGE-MIB",
            "path": "bluecoat/BLUECOAT-SG-USAGE-MIB.mib"
        },
        {
            "mib": "CONFIG-MIB",
            "path": "hp/CONFIG-MIB.mib"
        },
        {
            "mib": "E5-121-MIB",
            "path": "calix/E5-121-MIB.mib"
        },
        {
            "mib": "CM-SAT-MIB",
            "path": "adva/CM-SAT-MIB.mib"
        },
        {
            "mib": "ERRP-MIB",
            "path": "fs/ERRP-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-GBE-MCM1000X",
            "path": "fibrolan/FIBROLAN-MIB-GBE-MCM1000X.mib"
        },
        {
            "mib": "DELL-NETWORKING-SMI",
            "path": "dell/DELL-NETWORKING-SMI.mib"
        },
        {
            "mib": "EQLRAID-MIB",
            "path": "equallogic/EQLRAID-MIB.mib"
        },
        {
            "mib": "CDATA-EPON-MIB",
            "path": "cdata/CDATA-EPON-MIB.mib"
        },
        {
            "mib": "VMWARE-VROPS-AGENTCAP-MIB",
            "path": "vmware/VMWARE-VROPS-AGENTCAP-MIB.mib"
        },
        {
            "mib": "PACKETFLUX-STANDBYPOWER",
            "path": "packetflux/PACKETFLUX-STANDBYPOWER.mib"
        },
        {
            "mib": "ALCATEL-IND1-IP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IP-MIB.mib"
        },
        {
            "mib": "AT-SYSINFO-MIB",
            "path": "allied/AT-SYSINFO-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-GBE-MCM1000XRL",
            "path": "fibrolan/FIBROLAN-MIB-GBE-MCM1000XRL.mib"
        },
        {
            "mib": "HUAWEI-BRAS-MULTICAST-MIB",
            "path": "huawei/HUAWEI-BRAS-MULTICAST-MIB.mib"
        },
        {
            "mib": "BLUECOAT-SG-WCCP-MIB",
            "path": "bluecoat/BLUECOAT-SG-WCCP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-VIRTUAL-AP-MIB",
            "path": "hpmsm/COLUBRIS-VIRTUAL-AP-MIB.my.mib"
        },
        {
            "mib": "ACD-SA-MIB",
            "path": "accedian/ACD-SA-MIB.mib"
        },
        {
            "mib": "CTRON-DECIV-ROUTER-MIB",
            "path": "enterasys/CTRON-DECIV-ROUTER-MIB.mib"
        },
        {
            "mib": "FS-AC-MGMT-MIB",
            "path": "fs/FS-AC-MGMT-MIB.mib"
        },
        {
            "mib": "E5-121-TRAPS-MIB",
            "path": "calix/E5-121-TRAPS-MIB.mib"
        },
        {
            "mib": "Argus-Power-System-MIB",
            "path": "alpha/Argus-Power-System-MIB.mib"
        },
        {
            "mib": "PROFLINE-MIB",
            "path": "profline/PROFLINE.mib"
        },
        {
            "mib": "EQLREPLPARTNER-MIB",
            "path": "equallogic/EQLREPLPARTNER-MIB.mib"
        },
        {
            "mib": "MIMOSA-MIB-TC",
            "path": "mimosa/MIMOSA-MIB-TC.mib"
        },
        {
            "mib": "DELL-NETWORKING-SYSLOG-MIB",
            "path": "dell/DELL-NETWORKING-SYSLOG-MIB.mib"
        },
        {
            "mib": "CPQHLTH-MIB",
            "path": "hp/CPQHLTH-MIB.mib"
        },
        {
            "mib": "SCHLEIFENBAUER-DATABUS-MIB",
            "path": "schleifenbauer/SCHLEIFENBAUER-DATABUS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-MVLAN-MIB",
            "path": "huawei/HUAWEI-BRAS-MVLAN-MIB.mib"
        },
        {
            "mib": "CDATA-GPON-MIB",
            "path": "cdata/CDATA-GPON-MIB.mib"
        },
        {
            "mib": "VMWARE-VROPS-MIB",
            "path": "vmware/VMWARE-VROPS-MIB.mib"
        },
        {
            "mib": "COLUBRIS-VSC-MIB",
            "path": "hpmsm/COLUBRIS-VSC-MIB.my.mib"
        },
        {
            "mib": "CTRON-DEVICE-MIB",
            "path": "enterasys/CTRON-DEVICE-MIB.mib"
        },
        {
            "mib": "ACD-SFP-MIB",
            "path": "accedian/ACD-SFP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPMRM-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IPMRM-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-GBECONVERTERS",
            "path": "fibrolan/FIBROLAN-MIB-GBECONVERTERS.mib"
        },
        {
            "mib": "EQLSTORAGECONTAINER-MIB",
            "path": "equallogic/EQLSTORAGECONTAINER-MIB.mib"
        },
        {
            "mib": "AT-TRIGGER-MIB",
            "path": "allied/AT-TRIGGER-MIB.mib"
        },
        {
            "mib": "CM-SECURITY-MIB",
            "path": "adva/CM-SECURITY-MIB.mib"
        },
        {
            "mib": "TERACOM-TCW210-TH-MIB",
            "path": "teracom/TERACOM-TCW210-TH-MIB.mib"
        },
        {
            "mib": "PROFLINE-SFD-MIB",
            "path": "profline/PROFLINESFD.mib"
        },
        {
            "mib": "E7-Calix-MIB",
            "path": "calix/E7-Calix-MIB.mib"
        },
        {
            "mib": "COLUBRIS-WDS-MIB",
            "path": "hpmsm/COLUBRIS-WDS-MIB.my.mib"
        },
        {
            "mib": "SCHLEIFENBAUER-SMI",
            "path": "schleifenbauer/SCHLEIFENBAUER-SMI.mib"
        },
        {
            "mib": "MIMOSA-NETWORKS-BASE-MIB",
            "path": "mimosa/MIMOSA-NETWORKS-BASE-MIB.mib"
        },
        {
            "mib": "FS-ENTITY-MIB",
            "path": "fs/FS-ENTITY-MIB.mib"
        },
        {
            "mib": "CTRON-DHCP-MIB",
            "path": "enterasys/CTRON-DHCP-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-PPPoX-MIB",
            "path": "huawei/HUAWEI-BRAS-PPPoX-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-SYSTEM-COMPONENT-MIB",
            "path": "dell/DELL-NETWORKING-SYSTEM-COMPONENT-MIB.mib"
        },
        {
            "mib": "CPQHOST-MIB",
            "path": "hp/CPQHOST-MIB.mib"
        },
        {
            "mib": "AT-TTY-MIB",
            "path": "allied/AT-TTY-MIB.mib"
        },
        {
            "mib": "PNETMOD-MIB",
            "path": "bke/PNETMOD-MIB.mib"
        },
        {
            "mib": "ACD-SHAPER-MIB",
            "path": "accedian/ACD-SHAPER-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-LTA41MA-V2",
            "path": "fibrolan/FIBROLAN-MIB-LTA41MA.mib"
        },
        {
            "mib": "E7-TC",
            "path": "calix/E7-TC.mib"
        },
        {
            "mib": "SITE-MONITORING-MIB",
            "path": "alpha/SITE-MONITORING-MIB.mib"
        },
        {
            "mib": "TERACOM-TCW220-MIB",
            "path": "teracom/TERACOM-TCW220-MIB.mib"
        },
        {
            "mib": "CM-SYSTEM-MIB",
            "path": "adva/CM-SYSTEM-MIB.mib"
        },
        {
            "mib": "EQLSTORAGEPOOL-MIB",
            "path": "equallogic/EQLSTORAGEPOOL-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPRM-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IPRM-MIB.mib"
        },
        {
            "mib": "MIMOSA-NETWORKS-BFIVE-MIB",
            "path": "mimosa/MIMOSA-NETWORKS-BFIVE-MIB.mib"
        },
        {
            "mib": "CDATA-GPON-MIB2",
            "path": "cdata/CDATA-GPON-MIB2.mib"
        },
        {
            "mib": "ACD-SMAP-MIB",
            "path": "accedian/ACD-SMAP-MIB.mib"
        },
        {
            "mib": "COLUBRIS-WIRELESS-CLIENT-MIB",
            "path": "hpmsm/COLUBRIS-WIRELESS-CLIENT-MIB.my.mib"
        },
        {
            "mib": "CTRON-DLSW-MIB",
            "path": "enterasys/CTRON-DLSW-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-LTA41xE1-V2",
            "path": "fibrolan/FIBROLAN-MIB-LTA41XE1.mib"
        },
        {
            "mib": "FS-FIBER-MIB",
            "path": "fs/FS-FIBER-MIB.mib"
        },
        {
            "mib": "AtiEdgeSwitch-MIB",
            "path": "allied/AtiEdgeSwitch-MIB.mib"
        },
        {
            "mib": "CPQIDA-MIB",
            "path": "hp/CPQIDA-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-QOS-MIB",
            "path": "huawei/HUAWEI-BRAS-QOS-MIB.mib"
        },
        {
            "mib": "CM-TOPOLOGY-MIB",
            "path": "adva/CM-TOPOLOGY-MIB.mib"
        },
        {
            "mib": "EQLTAG-MIB",
            "path": "equallogic/EQLTAG-MIB.mib"
        },
        {
            "mib": "TERACOM-TCW241-MIB",
            "path": "teracom/TERACOM-TCW241-MIB.mib"
        },
        {
            "mib": "ACD-TID-MIB",
            "path": "accedian/ACD-TID-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPRMV6-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IPRMV6-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-TC",
            "path": "dell/DELL-NETWORKING-TC.mib"
        },
        {
            "mib": "FS-MEMORY-MIB",
            "path": "fs/FS-MEMORY-MIB.mib"
        },
        {
            "mib": "OCCAM-ENTITY-MIB",
            "path": "calix/OCCAM-ENTITY-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MCM100-xE1",
            "path": "fibrolan/FIBROLAN-MIB-MCM100-xE1.mib"
        },
        {
            "mib": "GW-EPON-DEV-MIB",
            "path": "gwd/GW-EPON-DEV-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-RADIUS-MIB",
            "path": "huawei/HUAWEI-BRAS-RADIUS-MIB.mib"
        },
        {
            "mib": "ENDACE-EDA-MIB",
            "path": "endace/ENDACE-EDA-MIB.mib"
        },
        {
            "mib": "CTRON-DOWNLOAD-MIB",
            "path": "enterasys/CTRON-DOWNLOAD-MIB.mib"
        },
        {
            "mib": "DEV-CFG-MIB",
            "path": "adva/DEV-CFG-MIB.mib"
        },
        {
            "mib": "MIMOSA-NETWORKS-PTMP-MIB",
            "path": "mimosa/MIMOSA-NETWORKS-PTMP-MIB.mib"
        },
        {
            "mib": "AtiL2-MIB",
            "path": "allied/AtiL2-MIB.mib"
        },
        {
            "mib": "EQLVOLBALANCER-MIB",
            "path": "equallogic/EQLVOLBALANCER-MIB.mib"
        },
        {
            "mib": "EDFA-oa-MIB",
            "path": "cdata/EDFA-oa-MIB.mib"
        },
        {
            "mib": "HWg-WLD-MIB",
            "path": "hwg/HWg-WLD-MIB.mib"
        },
        {
            "mib": "OCCAM-ETHERLIKE-MIB",
            "path": "calix/OCCAM-ETHERLIKE-MIB.mib"
        },
        {
            "mib": "DELL-NETWORKING-TRAP-EVENT-MIB",
            "path": "dell/DELL-NETWORKING-TRAP-EVENT-MIB.mib"
        },
        {
            "mib": "NUTANIX-MIB",
            "path": "nutanix-aos/NUTANIX-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-RUI-MIB",
            "path": "huawei/HUAWEI-BRAS-RUI-MIB.mib"
        },
        {
            "mib": "ENDACE-ERFSTREAM-MIB",
            "path": "endace/ENDACE-ERFSTREAM-MIB.mib"
        },
        {
            "mib": "MITEL-APPCMN-MIB",
            "path": "mitel/MITEL-APPCMN-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPSEC-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IPSEC-MIB.mib"
        },
        {
            "mib": "DEV-ID-MIB",
            "path": "adva/DEV-ID-MIB.mib"
        },
        {
            "mib": "GW-EPON-MIB",
            "path": "gwd/GW-EPON-MIB.mib"
        },
        {
            "mib": "CPQPOWER-MIB",
            "path": "hp/CPQPOWER-MIB.mib"
        },
        {
            "mib": "FS-MIB",
            "path": "fs/FS-MIB.mib"
        },
        {
            "mib": "TERACOM-TCW242-MIB",
            "path": "teracom/TERACOM-TCW242-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MCM1xx",
            "path": "fibrolan/FIBROLAN-MIB-MCM1xx.mib"
        },
        {
            "mib": "CTRON-ELAN-MIB",
            "path": "enterasys/CTRON-ELAN-MIB.mib"
        },
        {
            "mib": "FS-PROCESS-MIB",
            "path": "fs/FS-PROCESS-MIB.mib"
        },
        {
            "mib": "GEIST-MIB-V3",
            "path": "geist/GEIST-MIB-V3.mib"
        },
        {
            "mib": "ENDACE-HBA-MIB",
            "path": "endace/ENDACE-HBA-MIB.mib"
        },
        {
            "mib": "OCCAM-MLT-MIB",
            "path": "calix/OCCAM-MLT-MIB.mib"
        },
        {
            "mib": "MITEL-APPLICATION-PLATFORM-LIST-MIB",
            "path": "mitel/MITEL-APPLICATION-PLATFORM-LIST-MIB.mib"
        },
        {
            "mib": "AtiStackInfo-MIB",
            "path": "allied/AtiStackInfo-MIB.mib"
        },
        {
            "mib": "GEIST-V4-MIB",
            "path": "geist/GEIST-V4-MIB.mib"
        },
        {
            "mib": "TERACOM-TCW260-MIB",
            "path": "teracom/TERACOM-TCW260-MIB.mib"
        },
        {
            "mib": "EPON-EOC-MIB",
            "path": "cdata/EPON-EOC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPV6-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-IPV6-MIB.mib"
        },
        {
            "mib": "GWTT-SMI",
            "path": "gwd/GWTT-SMI.mib"
        },
        {
            "mib": "DELL-NETWORKING-VIRTUAL-LINK-TRUNK-MIB",
            "path": "dell/DELL-NETWORKING-VIRTUAL-LINK-TRUNK-MIB.mib"
        },
        {
            "mib": "EQLVOLUME-MIB",
            "path": "equallogic/EQLVOLUME-MIB.mib"
        },
        {
            "mib": "CIENA-6500R-INVENTORY-AMPS-MIB",
            "path": "ciena/CIENA-6500R-INVENTORY-AMPS-MIB.mib"
        },
        {
            "mib": "CPQRACK-MIB",
            "path": "hp/CPQRACK-MIB.mib"
        },
        {
            "mib": "F3-AMP-MIB",
            "path": "adva/F3-AMP-MIB.mib"
        },
        {
            "mib": "CTRON-ENTITY-STATE-MIB",
            "path": "enterasys/CTRON-ENTITY-STATE-MIB.mib"
        },
        {
            "mib": "EQUALLOGIC-SMI",
            "path": "equallogic/EQUALLOGIC-SMI.mib"
        },
        {
            "mib": "FIBROLAN-MIB-METRO-STAR-MV",
            "path": "fibrolan/FIBROLAN-MIB-METRO-STAR-MV.mib"
        },
        {
            "mib": "AtiSwitch-MIB",
            "path": "allied/AtiSwitch-MIB.mib"
        },
        {
            "mib": "ENDACE-HWMON-MIB",
            "path": "endace/ENDACE-HWMON-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-ISIS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-ISIS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-SBC-MIB",
            "path": "huawei/HUAWEI-BRAS-SBC-MIB.mib"
        },
        {
            "mib": "LM-SENSORS-MIB",
            "path": "davis/LM-SENSORS-MIB.mib"
        },
        {
            "mib": "CIENA-6500R-INVENTORY-MIB",
            "path": "ciena/CIENA-6500R-INVENTORY-MIB.mib"
        },
        {
            "mib": "OCCAM-REG-MODULE",
            "path": "calix/OCCAM-REG-MODULE.mib"
        },
        {
            "mib": "RAD-MIB",
            "path": "packetlight/RAD-MIB.mib"
        },
        {
            "mib": "FD-OLT-MIB",
            "path": "cdata/FD-OLT-MIB.mib"
        },
        {
            "mib": "DELL-RAC-MIB",
            "path": "dell/DELL-RAC-MIB.mib"
        },
        {
            "mib": "CTRON-ENTITY-STATE-TC-MIB",
            "path": "enterasys/CTRON-ENTITY-STATE-TC-MIB.mib"
        },
        {
            "mib": "MITEL-APPLIST-MIB",
            "path": "mitel/MITEL-APPLIST-MIB.mib"
        },
        {
            "mib": "CPQSINFO-MIB",
            "path": "hp/CPQSINFO-MIB.mib"
        },
        {
            "mib": "F3-BFD-MIB",
            "path": "adva/F3-BFD-MIB.mib"
        },
        {
            "mib": "ENDACE-INVMGR-MIB",
            "path": "endace/ENDACE-INVMGR-MIB.mib"
        },
        {
            "mib": "Perseus-MIB",
            "path": "hwg/Perseus-MIB.mib"
        },
        {
            "mib": "IPS-AUTH-MIB",
            "path": "equallogic/IPS-AUTH-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-ISIS-SPB-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-ISIS-SPB-MIB.mib"
        },
        {
            "mib": "TELESYN-ATI-TC",
            "path": "allied/TELESYN-ATI-TC.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MSM100U",
            "path": "fibrolan/FIBROLAN-MIB-MSM100U.mib"
        },
        {
            "mib": "MITEL-CMNALM-MIB",
            "path": "mitel/MITEL-CMNALM-MIB.mib"
        },
        {
            "mib": "UCD-DEMO-MIB",
            "path": "davis/UCD-DEMO-MIB.mib"
        },
        {
            "mib": "CTRON-ENVIRONMENT-MIB",
            "path": "enterasys/CTRON-ENVIRONMENT-MIB.mib"
        },
        {
            "mib": "F3-BRIDGE-MIB",
            "path": "adva/F3-BRIDGE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LAG-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-LAG-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MSM2500U",
            "path": "fibrolan/FIBROLAN-MIB-MSM2500U.mib"
        },
        {
            "mib": "FS-SMI",
            "path": "fs/FS-SMI.mib"
        },
        {
            "mib": "ISCSI-MIB",
            "path": "equallogic/ISCSI-MIB.mib"
        },
        {
            "mib": "CTRON-ETHERNET-PARAMETERS-MIB",
            "path": "enterasys/CTRON-ETHERNET-PARAMETERS-MIB.mib"
        },
        {
            "mib": "CIENA-6500R-SLOTS-MIB",
            "path": "ciena/CIENA-6500R-SLOTS-MIB.mib"
        },
        {
            "mib": "ALPHACOM-MIB",
            "path": "zenitel/ALPHACOM-MIB.mib"
        },
        {
            "mib": "DELL-REF-MIB",
            "path": "dell/DELL-REF-MIB.mib"
        },
        {
            "mib": "MITEL-MIB",
            "path": "mitel/MITEL-MIB.mib"
        },
        {
            "mib": "CPQSTDEQ-MIB",
            "path": "hp/CPQSTDEQ-MIB.mib"
        },
        {
            "mib": "OCCAM-SENSOR-MIB",
            "path": "calix/OCCAM-SENSOR-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MSM622U",
            "path": "fibrolan/FIBROLAN-MIB-MSM622U.mib"
        },
        {
            "mib": "IPFIX-MIB",
            "path": "endace/ENDACE-IPFIX-MIB.mib"
        },
        {
            "mib": "UCD-DISKIO-MIB",
            "path": "davis/UCD-DISKIO-MIB.mib"
        },
        {
            "mib": "F3-CAPABILITIES-MIB",
            "path": "adva/F3-CAPABILITIES-MIB.mib"
        },
        {
            "mib": "FS-SYSTEM-MIB",
            "path": "fs/FS-SYSTEM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LPS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-LPS-MIB.mib"
        },
        {
            "mib": "FD-ONU-MIB",
            "path": "cdata/FD-ONU-MIB.mib"
        },
        {
            "mib": "CTRON-ETWMIM-MIB",
            "path": "enterasys/CTRON-ETWMIM-MIB.mib"
        },
        {
            "mib": "SL-ALARM-MIB",
            "path": "packetlight/SL-ALARM-MIB.mib"
        },
        {
            "mib": "CIENA-6500R-TYPES-MIB",
            "path": "ciena/CIENA-6500R-TYPES-MIB.mib"
        },
        {
            "mib": "SCSI-MIB",
            "path": "equallogic/SCSI-MIB.mib"
        },
        {
            "mib": "ES4552BH2-MIB",
            "path": "ignitenet/ES4552BH2-MIB.mib"
        },
        {
            "mib": "FS-TC",
            "path": "fs/FS-TC.mib"
        },
        {
            "mib": "FAN-MIB",
            "path": "hp/FAN-MIB.mib"
        },
        {
            "mib": "VS-DEVICE-MIB",
            "path": "zenitel/VS-Device-MIB.mib"
        },
        {
            "mib": "UBNT-AFLTU-MIB",
            "path": "ubnt/UBNT-AFLTU-MIB.mib"
        },
        {
            "mib": "FS-SWITCH-V2-MIB",
            "path": "fs/FS-SIWTCH-V2-MIB.mib"
        },
        {
            "mib": "DELL-SHADOW-MIB",
            "path": "dell/DELL-SHADOW-MIB.mib"
        },
        {
            "mib": "ENDACE-MIB",
            "path": "endace/ENDACE-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-SRVCFG-DEVICE-MIB",
            "path": "huawei/HUAWEI-BRAS-SRVCFG-DEVICE-MIB.mib"
        },
        {
            "mib": "POSEIDON-MIB",
            "path": "hwg/POSEIDON-MIB.mib"
        },
        {
            "mib": "UCD-DLMOD-MIB",
            "path": "davis/UCD-DLMOD-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MSMODULE",
            "path": "fibrolan/FIBROLAN-MIB-MSMODULE.mib"
        },
        {
            "mib": "F3-CFM-MIB",
            "path": "adva/F3-CFM-MIB.mib"
        },
        {
            "mib": "OCCAM-SHELF-MIB",
            "path": "calix/OCCAM-SHELF-MIB.mib"
        },
        {
            "mib": "FD-PERFORMANCE-MIB",
            "path": "cdata/FD-PERFORMANCE-MIB.mib"
        },
        {
            "mib": "CTRON-FDDI-FNB-MIB",
            "path": "enterasys/CTRON-FDDI-FNB-MIB.mib"
        },
        {
            "mib": "UBNT-AirFIBER-MIB",
            "path": "ubnt/UBNT-AirFIBER-MIB.mib"
        },
        {
            "mib": "SL-ALS-MIB",
            "path": "packetlight/SL-ALS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-8021X-MIB",
            "path": "ciena/CIENA-CES-8021X-MIB.mib"
        },
        {
            "mib": "DELL-SNMP-UPS-MIB",
            "path": "dell/DELL-SNMP-UPS-MIB.mib"
        },
        {
            "mib": "HP-AUTZ-MIB",
            "path": "hp/HP-AUTZ-MIB.mib"
        },
        {
            "mib": "GARP-MIB",
            "path": "fs/GARP-MIB.mib"
        },
        {
            "mib": "ALPINE-GEN-CARD-EDFA-MIB",
            "path": "alpineoe/ALPINE-GEN-CARD-EDFA-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MAC-ADDRESS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-MAC-ADDRESS-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-REDUNDANCYDEVICES",
            "path": "fibrolan/FIBROLAN-MIB-REDUNDANCYDEVICES.mib"
        },
        {
            "mib": "IOMEGANAS-MIB",
            "path": "lenovo/IOMEGANAS-MIB.mib"
        },
        {
            "mib": "ENDACE-MODULE-MIB",
            "path": "endace/ENDACE-MODULE-MIB.mib"
        },
        {
            "mib": "GBNDeviceOEM-MIB",
            "path": "fs/GBNDeviceOEM-MIB.mib"
        },
        {
            "mib": "CTRON-FDDI-STAT-MIB",
            "path": "enterasys/CTRON-FDDI-STAT-MIB.mib"
        },
        {
            "mib": "UBNT-AirMAX-MIB",
            "path": "ubnt/UBNT-AirMAX-MIB.mib"
        },
        {
            "mib": "F3-CONNECTGUARD-MIB",
            "path": "adva/F3-CONNECTGUARD-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MAC-SERVER-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-MAC-SERVER-MIB.mib"
        },
        {
            "mib": "GBNDevicePoe-MIB",
            "path": "fs/GBNDevicePoe-MIB.mib"
        },
        {
            "mib": "FD-SWITCH-MIB",
            "path": "cdata/FD-SWITCH-MIB.mib"
        },
        {
            "mib": "DELL-STORAGE-SC-MIB",
            "path": "dell/DELL-STORAGE-SC-MIB.mib"
        },
        {
            "mib": "DATUM-MIB",
            "path": "microsemi/DATUM-MIB.mib"
        },
        {
            "mib": "UCD-IPFWACC-MIB",
            "path": "davis/UCD-IPFWACC-MIB.mib"
        },
        {
            "mib": "SL-CHASSIS-MIB",
            "path": "packetlight/SL-CHASSIS-MIB.mib"
        },
        {
            "mib": "ENDACE-STREAM-MIB",
            "path": "endace/ENDACE-STREAM-MIB.mib"
        },
        {
            "mib": "HP-BASE-MIB",
            "path": "hp/HP-BASE-MIB.mib"
        },
        {
            "mib": "F3-DATAEXPORT-MIB",
            "path": "adva/F3-DATAEXPORT-MIB.mib"
        },
        {
            "mib": "UBNT-EdgeMAX-MIB",
            "path": "ubnt/UBNT-EdgeMAX-MIB.mib"
        },
        {
            "mib": "ALPINE-GEN-CARD-TDCM-MIB",
            "path": "alpineoe/ALPINE-GEN-CARD-TDCM-MIB.mib"
        },
        {
            "mib": "CIENA-CES-AAA-MIB",
            "path": "ciena/CIENA-CES-AAA-MIB.mib"
        },
        {
            "mib": "SL-DRY-CON-MIB",
            "path": "packetlight/SL-DRY-CON-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MLD-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-MLD-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-SFP",
            "path": "fibrolan/FIBROLAN-MIB-SFP.mib"
        },
        {
            "mib": "UBNT-MIB",
            "path": "ubnt/UBNT-MIB.mib"
        },
        {
            "mib": "HP-CAR-MIB",
            "path": "hp/HP-CAR-MIB.mib"
        },
        {
            "mib": "ENDACE-SYSTEM-MIB",
            "path": "endace/ENDACE-SYSTEM-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-SRVCFG-EAP-MIB",
            "path": "huawei/HUAWEI-BRAS-SRVCFG-EAP-MIB.mib"
        },
        {
            "mib": "GBNDeviceStack-MIB",
            "path": "fs/GBNDeviceStack-MIB.mib"
        },
        {
            "mib": "Dell-Vendor-MIB",
            "path": "dell/Dell-Vendor-MIB.mib"
        },
        {
            "mib": "F3-ELMI-MIB",
            "path": "adva/F3-ELMI-MIB.mib"
        },
        {
            "mib": "FIBROLAN-SFP-MIB",
            "path": "fibrolan/FIBROLAN-SFP-MIB.mib"
        },
        {
            "mib": "MICROSEMI-PDSINE-MIB",
            "path": "microsemi/MICROSEMI-PDSINE-MIB.mib"
        },
        {
            "mib": "CTRON-FNBTR-MIB",
            "path": "enterasys/CTRON-FNBTR-MIB.mib"
        },
        {
            "mib": "LENOVO-ENV-MIB",
            "path": "lenovo/LENOVO-ENV-MIB.mib"
        },
        {
            "mib": "STE-MIB",
            "path": "hwg/STE-MIB.mib"
        },
        {
            "mib": "IGNITENET-MIB",
            "path": "ignitenet/IGNITENET-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MULTI-CHASSIS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-MULTI-CHASSIS-MIB.mib"
        },
        {
            "mib": "FD-SYSTEM-MIB",
            "path": "cdata/FD-SYSTEM-MIB.mib"
        },
        {
            "mib": "UCD-SNMP-MIB",
            "path": "davis/UCD-SNMP-MIB.mib"
        },
        {
            "mib": "UBNT-UFIBER-MIB",
            "path": "ubnt/UBNT-UFIBER-MIB.mib"
        },
        {
            "mib": "HP-DOT1X-EXTENSIONS-MIB",
            "path": "hp/HP-DOT1X-EXTENSIONS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-ACCESS-LIST-MIB",
            "path": "ciena/CIENA-CES-ACCESS-LIST-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-METRO-STAR",
            "path": "fibrolan/METROSTAR-MIB.mib"
        },
        {
            "mib": "F3-EOMPLS-MIB",
            "path": "adva/F3-EOMPLS-MIB.mib"
        },
        {
            "mib": "CTRON-FRONTPANEL-MIB",
            "path": "enterasys/CTRON-FRONTPANEL-MIB.mib"
        },
        {
            "mib": "ENDACEModuleType-MIB",
            "path": "endace/ENDACEModuleType-MIB.mib"
        },
        {
            "mib": "GBNDeviceSWAPI-MIB",
            "path": "fs/GBNDeviceSWAPI-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MVRP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-MVRP-MIB.mib"
        },
        {
            "mib": "ALPINE-ROOT",
            "path": "alpineoe/ALPINE-ROOT.mib"
        },
        {
            "mib": "SL-EDFA-MIB",
            "path": "packetlight/SL-EDFA-MIB.mib"
        },
        {
            "mib": "UBNT-UniFi-MIB",
            "path": "ubnt/UBNT-UniFi-MIB.mib"
        },
        {
            "mib": "LENOVO-PRODUCTS-MIB",
            "path": "lenovo/LENOVO-PRODUCTS-MIB.mib"
        },
        {
            "mib": "FIBROLAN-MIB-MS-TRAPS",
            "path": "fibrolan/MS-TRAPS-MIB.mib"
        },
        {
            "mib": "F3-EOTDM-MIB",
            "path": "adva/F3-EOTDM-MIB.mib"
        },
        {
            "mib": "HP-ENTITY-MIB",
            "path": "hp/HP-ENTITY-MIB.mib"
        },
        {
            "mib": "GBNDeviceSwitch-MIB",
            "path": "fs/GBNDeviceSwitch-MIB.mib"
        },
        {
            "mib": "WebGraph-8xThermometer-US-MIB",
            "path": "wut/WebGraph-8xThermometer-US-MIB.mib"
        },
        {
            "mib": "FD-TRAP-MIB",
            "path": "cdata/FD-TRAP-MIB.mib"
        },
        {
            "mib": "UI-AF60-MIB",
            "path": "ubnt/UI-AF60-MIB.mib"
        },
        {
            "mib": "CIENA-CES-ACL-MIB",
            "path": "ciena/CIENA-CES-ACL-MIB.mib"
        },
        {
            "mib": "SSU2000-MIB",
            "path": "microsemi/SSU2000-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-NETSEC-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-NETSEC-MIB.mib"
        },
        {
            "mib": "HMPRIV-MGMT-SNMP-MIB",
            "path": "hirschmann/HMPRIV-MGMT-SNMP-MIB.mib"
        },
        {
            "mib": "GBNL2Dhcp6Snooping-MIB",
            "path": "fs/GBNL2Dhcp6Snooping-MIB.mib"
        },
        {
            "mib": "KMIB",
            "path": "cirpack/KMIB-MIB.mib"
        },
        {
            "mib": "F3-ERP-MIB",
            "path": "adva/F3-ERP-MIB.mib"
        },
        {
            "mib": "CTRON-IF-REMAP-2-MIB",
            "path": "enterasys/CTRON-IF-REMAP-2-MIB.mib"
        },
        {
            "mib": "WebGraph-Thermo-Hygro-Barometer-MIB",
            "path": "wut/WebGraph-Thermo-Hygro-Barometer-MIB.mib"
        },
        {
            "mib": "LENOVO-SMI-MIB",
            "path": "lenovo/LENOVO-SMI-MIB.mib"
        },
        {
            "mib": "DELLEMC-OS10-BGP4V2-MIB",
            "path": "dell/DELLEMC-OS10-BGP4V2-MIB.mib"
        },
        {
            "mib": "GBNL2DhcpSnooping-MIB",
            "path": "fs/GBNL2DhcpSnooping-MIB.mib"
        },
        {
            "mib": "SYMM-SMI",
            "path": "microsemi/SYMM-SMI.mib"
        },
        {
            "mib": "STE2-MIB",
            "path": "hwg/STE2-MIB.mib"
        },
        {
            "mib": "MBG-SNMP-FDMXPT-MIB",
            "path": "meinberg/MBG-SNMP-FDMXPT-MIB.mib"
        },
        {
            "mib": "CIENA-CES-ALARM-MIB",
            "path": "ciena/CIENA-CES-ALARM-MIB.mib"
        },
        {
            "mib": "LTNET-COMMONINFO-MIB",
            "path": "cdata/LTNET-COMMONINFO-MIB.mib"
        },
        {
            "mib": "SL-ENTITY-MIB",
            "path": "packetlight/SL-ENTITY-MIB.mib"
        },
        {
            "mib": "BARRACUDA-REF",
            "path": "barracuda/BARRACUDA-REF-MIB.mib"
        },
        {
            "mib": "ALPINE-TDCM-EDFA-MIB",
            "path": "alpineoe/ALPINE-TDCM-EDFA-MIB.mib"
        },
        {
            "mib": "CCPOWER-MIB",
            "path": "ccpower/CCPOWER-MIB.mib"
        },
        {
            "mib": "HP-ICF-8023-RPTR",
            "path": "hp/HP-ICF-8023-RPTR.mib"
        },
        {
            "mib": "HUAWEI-BRAS-SRVCFG-STATICUSER-MIB",
            "path": "huawei/HUAWEI-BRAS-SRVCFG-STATICUSER-MIB.mib"
        },
        {
            "mib": "F3-ESM-MIB",
            "path": "adva/F3-ESM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-NTP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-NTP-MIB.mib"
        },
        {
            "mib": "CTRON-IF-REMAP-MIB",
            "path": "enterasys/CTRON-IF-REMAP-MIB.mib"
        },
        {
            "mib": "LTNET-ROOT",
            "path": "cdata/LTNET-ROOT.mib"
        },
        {
            "mib": "WAYSTREAM-COPY-MIB",
            "path": "waystream/WAYSTREAM-COPY-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-AAA-MIB",
            "path": "nokia/ALCATEL-IND1-AAA-MIB.mib"
        },
        {
            "mib": "APRISAXE-EVENTS-4RF",
            "path": "4rf/4RF-APRISAXE-EVENTS.mib"
        },
        {
            "mib": "GBNL2PortSecurity-MIB",
            "path": "fs/GBNL2PortSecurity-MIB.mib"
        },
        {
            "mib": "WebGraph-Thermo-Hygrometer-MIB",
            "path": "wut/WebGraph-Thermo-Hygrometer-MIB.mib"
        },
        {
            "mib": "SL-ETH-MIB",
            "path": "packetlight/SL-ETH-MIB.mib"
        },
        {
            "mib": "LENOVO-XCC-ALERT-MIB",
            "path": "lenovo/LENOVO-XCC-ALERT-MIB.mib"
        },
        {
            "mib": "NETELASTIC-FLEXBNG-IPPOOL",
            "path": "netelastic/NETELASTIC-FLEXBNG-IPPOOL.mib"
        },
        {
            "mib": "ALCATEL-IND1-BASE",
            "path": "nokia/ALCATEL-IND1-BASE.mib"
        },
        {
            "mib": "MBG-SNMP-LT-MIB",
            "path": "meinberg/MBG-SNMP-LT-MIB.mib"
        },
        {
            "mib": "HP-ICF-ARP-PROTECT",
            "path": "hp/HP-ICF-ARP-PROTECT.mib"
        },
        {
            "mib": "DELLEMC-OS10-CHASSIS-MIB",
            "path": "dell/DELLEMC-OS10-CHASSIS-MIB.mib"
        },
        {
            "mib": "BWS-MIB",
            "path": "barracuda/BWS-MIB.mib"
        },
        {
            "mib": "AV-SME-PLATFORM-MIB",
            "path": "avaya/AV-SME-PLATFORM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-OPENFLOW-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-OPENFLOW-MIB.mib"
        },
        {
            "mib": "CIENA-CES-BENCHMARK-MIB",
            "path": "ciena/CIENA-CES-BENCHMARK-MIB.mib"
        },
        {
            "mib": "WAYSTREAM-IGMP-CACHE-MIB",
            "path": "waystream/WAYSTREAM-IGMP-CACHE-MIB.mib"
        },
        {
            "mib": "SL-EVENT-MIB",
            "path": "packetlight/SL-EVENT-MIB.mib"
        },
        {
            "mib": "GBNL2PppoePlus-MIB",
            "path": "fs/GBNL2PppoePlus-MIB.mib"
        },
        {
            "mib": "NE-ALARM-MIB",
            "path": "cdata/NE-ALARM-MIB.mib"
        },
        {
            "mib": "HP-ICF-AUTORUN",
            "path": "hp/HP-ICF-AUTORUN.mib"
        },
        {
            "mib": "NETELASTIC-FLEXBNG-IPPOOLV6",
            "path": "netelastic/NETELASTIC-FLEXBNG-IPPOOLV6.mib"
        },
        {
            "mib": "CTRON-IGMP-MIB",
            "path": "enterasys/CTRON-IGMP-MIB.mib"
        },
        {
            "mib": "PHION-MIB",
            "path": "barracuda/PHION-MIB.mib"
        },
        {
            "mib": "WebGraph-Thermo-Hygrometer-US-MIB",
            "path": "wut/WebGraph-Thermo-Hygrometer-US-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-OSPF-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-OSPF-MIB.mib"
        },
        {
            "mib": "F3-FPM-MIB",
            "path": "adva/F3-FPM-MIB.mib"
        },
        {
            "mib": "DELLEMC-OS10-PRODUCTS-MIB",
            "path": "dell/DELLEMC-OS10-PRODUCTS-MIB.mib"
        },
        {
            "mib": "AV-SME-PLATFORM-PROD-MIB",
            "path": "avaya/AV-SME-PLATFORM-PROD-MIB.mib"
        },
        {
            "mib": "GBNL2QACL-MIB",
            "path": "fs/GBNL2QACL-MIB.mib"
        },
        {
            "mib": "SL-FT-MIB",
            "path": "packetlight/SL-FT-MIB.mib"
        },
        {
            "mib": "WAYSTREAM-MIB",
            "path": "waystream/WAYSTREAM-MIB.mib"
        },
        {
            "mib": "CIENA-CES-BFD-MIB",
            "path": "ciena/CIENA-CES-BFD-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-BFD-MIB",
            "path": "nokia/ALCATEL-IND1-BFD-MIB.mib"
        },
        {
            "mib": "APRISAXE-MIB-4RF",
            "path": "4rf/4RF-APRISAXE-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPONEOC-MOD-EOC-MIB",
            "path": "cdata/NSCRTV-EPONEOC-MOD-EOC-MIB.mib"
        },
        {
            "mib": "HP-ICF-BASIC",
            "path": "hp/HP-ICF-BASIC.mib"
        },
        {
            "mib": "MBG-SNMP-LTNG-MIB",
            "path": "meinberg/MBG-SNMP-LTNG-MIB.mib"
        },
        {
            "mib": "DELLEMC-OS10-SMI-MIB",
            "path": "dell/DELLEMC-OS10-SMI-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-OSPF3-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-OSPF3-MIB.mib"
        },
        {
            "mib": "GBNL2Switch-MIB",
            "path": "fs/GBNL2Switch-MIB.mib"
        },
        {
            "mib": "APRISAXE-TC-4RF",
            "path": "4rf/4RF-APRISAXE-TC.mib"
        },
        {
            "mib": "WAYSTREAM-PRODUCTS-MIB",
            "path": "waystream/WAYSTREAM-PRODUCTS-MIB.mib"
        },
        {
            "mib": "NETELASTIC-FLEXBNG-MIB",
            "path": "netelastic/NETELASTIC-FLEXBNG-MIB.mib"
        },
        {
            "mib": "CTRON-IP-ROUTER-MIB",
            "path": "enterasys/CTRON-IP-ROUTER-MIB.mib"
        },
        {
            "mib": "AVAYAGEN-MIB",
            "path": "avaya/AVAYAGEN-MIB.mib"
        },
        {
            "mib": "SL-L2TOPOLOGY-MIB",
            "path": "packetlight/SL-L2TOPOLOGY-MIB.mib"
        },
        {
            "mib": "HP-ICF-BRIDGE",
            "path": "hp/HP-ICF-BRIDGE.mib"
        },
        {
            "mib": "ALCATEL-IND1-BGP-MIB",
            "path": "nokia/ALCATEL-IND1-BGP-MIB.mib"
        },
        {
            "mib": "MBG-SNMP-ROOT-MIB",
            "path": "meinberg/MBG-SNMP-ROOT-MIB.mib"
        },
        {
            "mib": "GBNL3-MIB",
            "path": "fs/GBNL3-MIB.mib"
        },
        {
            "mib": "COMMON-4RF",
            "path": "4rf/4RF-COMMON-MIB.mib"
        },
        {
            "mib": "F3-JDSU-MIB",
            "path": "adva/F3-JDSU-MIB.mib"
        },
        {
            "mib": "NETELASTIC-FLEXBNG-PPPOE",
            "path": "netelastic/NETELASTIC-FLEXBNG-PPPOE.mib"
        },
        {
            "mib": "DELLEMC-OS10-TC-MIB",
            "path": "dell/DELLEMC-OS10-TC-MIB.mib"
        },
        {
            "mib": "LENOVO-XCC-MIB",
            "path": "lenovo/LENOVO-XCC-MIB.mib"
        },
        {
            "mib": "GBNL3DhcpRelay-MIB",
            "path": "fs/GBNL3DhcpRelay-MIB.mib"
        },
        {
            "mib": "WebGraph-Thermometer-MIB",
            "path": "wut/WebGraph-Thermometer-MIB.mib"
        },
        {
            "mib": "SL-MAIN-MIB",
            "path": "packetlight/SL-MAIN-MIB.mib"
        },
        {
            "mib": "NSCRTV-FTTX-EPON-MIB",
            "path": "cdata/NSCRTV-FTTX-EPON-MIB.mib"
        },
        {
            "mib": "NAG-MIB",
            "path": "snr/NAG-MIB.mib"
        },
        {
            "mib": "WAYSTREAM-RPM-MIB",
            "path": "waystream/WAYSTREAM-RPM-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-SRVCFGINTERFACE-MIB",
            "path": "huawei/HUAWEI-BRAS-SRVCFGINTERFACE-MIB.mib"
        },
        {
            "mib": "CIENA-CES-CFM-MIB",
            "path": "ciena/CIENA-CES-CFM-MIB.mib"
        },
        {
            "mib": "IPO-MIB",
            "path": "avaya/IPO-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-PIM-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-PIM-MIB.mib"
        },
        {
            "mib": "MBG-SNMP-XPT-MIB",
            "path": "meinberg/MBG-SNMP-XPT-MIB.mib"
        },
        {
            "mib": "CTRON-IPX-ROUTER-MIB",
            "path": "enterasys/CTRON-IPX-ROUTER-MIB.mib"
        },
        {
            "mib": "DellMDStorageArray",
            "path": "dell/DellMDStorageArray-MIB.mib"
        },
        {
            "mib": "NETELASTIC-FLEXBNG-SYSINFO",
            "path": "netelastic/NETELASTIC-FLEXBNG-SYSINFO.mib"
        },
        {
            "mib": "HP-ICF-CHAIN",
            "path": "hp/HP-ICF-CHAIN.mib"
        },
        {
            "mib": "IPO-PHONES-MIB",
            "path": "avaya/IPO-PHONES-MIB.mib"
        },
        {
            "mib": "WEBMON-EDGE-MATRIX-MIB",
            "path": "dantel/WEBMON-EDGE-MATRIX-MIB.mib"
        },
        {
            "mib": "COMMON-TC-4RF",
            "path": "4rf/4RF-COMMON-TC.mib"
        },
        {
            "mib": "GBNL3If-MIB",
            "path": "fs/GBNL3If-MIB.mib"
        },
        {
            "mib": "CIENA-CES-CHASSIS-MIB",
            "path": "ciena/CIENA-CES-CHASSIS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-POLICY-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-POLICY-MIB.mib"
        },
        {
            "mib": "NSCRTV-FTTX-GPON-MIB",
            "path": "cdata/NSCRTV-FTTX-GPON-MIB.mib"
        },
        {
            "mib": "SL-MUX-MIB",
            "path": "packetlight/SL-MUX-MIB.mib"
        },
        {
            "mib": "CTRON-MIB-NAMES",
            "path": "enterasys/CTRON-MIB-NAMES.mib"
        },
        {
            "mib": "WebGraph-Thermometer-NTC-MIB",
            "path": "wut/WebGraph-Thermometer-NTC-MIB.mib"
        },
        {
            "mib": "DellrPDU-MIB",
            "path": "dell/DellrPDU-MIB.mib"
        },
        {
            "mib": "MIB-4RF",
            "path": "4rf/4RF-MIB.mib"
        },
        {
            "mib": "WAYSTREAM-SMI",
            "path": "waystream/WAYSTREAM-SMI.mib"
        },
        {
            "mib": "EMC-1",
            "path": "morningstar/EMC-1.mib"
        },
        {
            "mib": "ALCATEL-IND1-PORT-MAPPING",
            "path": "nokia/aos7/ALCATEL-IND1-PORT-MAPPING.mib"
        },
        {
            "mib": "ALCATEL-IND1-CHASSIS-MIB",
            "path": "nokia/ALCATEL-IND1-CHASSIS-MIB.mib"
        },
        {
            "mib": "ADONIS-DNS-MIB",
            "path": "bluecatnetworks/ADONIS-DNS-MIB.mib"
        },
        {
            "mib": "HP-ICF-CHASSIS",
            "path": "hp/HP-ICF-CHASSIS.mib"
        },
        {
            "mib": "CTRON-NAT-MIB",
            "path": "enterasys/CTRON-NAT-MIB.mib"
        },
        {
            "mib": "LCOS-LX-MIB",
            "path": "lancom/LCOS-LX-MIB.mib"
        },
        {
            "mib": "IPO-PROD-MIB",
            "path": "avaya/IPO-PROD-MIB.mib"
        },
        {
            "mib": "BEGEMOT-ATM-FREEBSD-MIB",
            "path": "pfsense/BEGEMOT-ATM-FREEBSD-MIB.mib"
        },
        {
            "mib": "GBNL3Igmp-MIB",
            "path": "fs/GBNL3Igmp-MIB.mib"
        },
        {
            "mib": "CIENA-CES-CONFIG-MGMT-MIB",
            "path": "ciena/CIENA-CES-CONFIG-MGMT-MIB.mib"
        },
        {
            "mib": "GENSTAR",
            "path": "morningstar/GENSTAR.mib"
        },
        {
            "mib": "MEINBERG-OS-MIB",
            "path": "meinberg/MEINBERG-OS-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-ALARMS-MIB",
            "path": "cdata/NSCRTV-HFCEMS-ALARMS-MIB.mib"
        },
        {
            "mib": "PRODUCTS-MIB-4RF",
            "path": "4rf/4RF-PRODUCTS-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-USERVLAN-MIB",
            "path": "huawei/HUAWEI-BRAS-USERVLAN-MIB.mib"
        },
        {
            "mib": "SL-NE-MIB",
            "path": "packetlight/SL-NE-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-ALARM-MIB-DEPRECATED",
            "path": "stormshield/STORMSHIELD-ALARM-MIB-DEPRECATED.mib"
        },
        {
            "mib": "ALCATEL-IND1-CONFIG-MGR-MIB",
            "path": "nokia/ALCATEL-IND1-CONFIG-MGR-MIB.mib"
        },
        {
            "mib": "MORNINGSTAR",
            "path": "morningstar/MORNINGSTAR.mib"
        },
        {
            "mib": "BEGEMOT-ATM-MIB",
            "path": "pfsense/BEGEMOT-ATM.mib"
        },
        {
            "mib": "ALCATEL-IND1-PORT-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-PORT-MIB.mib"
        },
        {
            "mib": "DNOS-AUTHENTICATION-MANAGER-MIB",
            "path": "dell/DNOS-AUTHENTICATION-MANAGER-MIB.mib"
        },
        {
            "mib": "HP-ICF-CONNECTION-RATE-FILTER",
            "path": "hp/HP-ICF-CONNECTION-RATE-FILTER.mib"
        },
        {
            "mib": "BCN-COMMANDSERVER-MIB",
            "path": "bluecatnetworks/BCN-COMMANDSERVER-MIB.mib"
        },
        {
            "mib": "GRANDSTREAM-GWN-PRODUCTS-AP-MIB",
            "path": "grandstream/GRANDSTREAM-GWN-PRODUCTS-AP-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-COMMON-MIB",
            "path": "cdata/NSCRTV-HFCEMS-COMMON-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-VSM-MIB",
            "path": "huawei/HUAWEI-BRAS-VSM-MIB.mib"
        },
        {
            "mib": "F3-L3-MIB",
            "path": "adva/F3-L3-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP800-MIB",
            "path": "cambium/800/CAMBIUM-PTP800-MIB.mib"
        },
        {
            "mib": "CIENA-CES-DATAPLANE-MIB",
            "path": "ciena/CIENA-CES-DATAPLANE-MIB.mib"
        },
        {
            "mib": "GBNL3IPPool-MIB",
            "path": "fs/GBNL3IPPool-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-PORT-MIRRORING-MONITORING-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-PORT-MIRRORING-MONITORING-MIB.mib"
        },
        {
            "mib": "CTRON-OIDS",
            "path": "enterasys/CTRON-OIDS.mib"
        },
        {
            "mib": "STORMSHIELD-ALARM-MIB",
            "path": "stormshield/STORMSHIELD-ALARM-MIB.mib"
        },
        {
            "mib": "GRANDSTREAM-GWN-PRODUCTS-MIB",
            "path": "grandstream/GRANDSTREAM-GWN-PRODUCTS-MIB.mib"
        },
        {
            "mib": "GBNL3Ospf-MIB",
            "path": "fs/GBNL3Ospf-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DEVICES",
            "path": "nokia/ALCATEL-IND1-DEVICES.mib"
        },
        {
            "mib": "BEGEMOT-BRIDGE-MIB",
            "path": "pfsense/BEGEMOT-BRIDGE-MIB.mib"
        },
        {
            "mib": "CIENA-CES-DHCP-RELAY-MIB",
            "path": "ciena/CIENA-CES-DHCP-RELAY-MIB.mib"
        },
        {
            "mib": "F10-C-SERIES-CHASSIS-MIB",
            "path": "dell/F10-C-SERIES-CHASSIS-MIB.mib"
        },
        {
            "mib": "SL-OPT-APS-MIB",
            "path": "packetlight/SL-OPT-APS-MIB.mib"
        },
        {
            "mib": "F3-LAG-MIB",
            "path": "adva/F3-LAG-MIB.mib"
        },
        {
            "mib": "HP-ICF-DOWNLOAD",
            "path": "hp/HP-ICF-DOWNLOAD.mib"
        },
        {
            "mib": "InterSeptor-MIB",
            "path": "jacarta/InterSeptor-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-ASQ-STATS-MIB",
            "path": "stormshield/STORMSHIELD-ASQ-STATS-MIB.mib"
        },
        {
            "mib": "GBNL3PIM-MIB",
            "path": "fs/GBNL3PIM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-QCN-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-QCN-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP800-MIB",
            "path": "cambium/800/CAMBIUM-PTP800-V1-MIB.mib"
        },
        {
            "mib": "CTRON-ORP-HSIM-MIB",
            "path": "enterasys/CTRON-ORP-HSIM-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-OPTICALAMPLIFIER-MIB",
            "path": "cdata/NSCRTV-HFCEMS-OPTICALAMPLIFIER-MIB.mib"
        },
        {
            "mib": "HILLSTONE-DHCP-MIB",
            "path": "hillstone/HILLSTONE-DHCP-MIB.mib"
        },
        {
            "mib": "CIENA-CES-DHCPV6-CLIENT-MIB",
            "path": "ciena/CIENA-CES-DHCPV6-CLIENT-MIB.mib"
        },
        {
            "mib": "HUAWEI-BRAS-VT-MIB",
            "path": "huawei/HUAWEI-BRAS-VT-MIB.mib"
        },
        {
            "mib": "PROSTAR-MPPT",
            "path": "morningstar/PROSTAR-MPPT.mib"
        },
        {
            "mib": "GRANDSTREAM-GWN-PRODUCTS-SWITCH-MIB",
            "path": "grandstream/GRANDSTREAM-GWN-PRODUCTS-SWITCH-MIB.mib"
        },
        {
            "mib": "BCN-DHCPV4-MIB",
            "path": "bluecatnetworks/BCN-DHCPV4-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DHCP-SRV-MIB",
            "path": "nokia/ALCATEL-IND1-DHCP-SRV-MIB.mib"
        },
        {
            "mib": "F10-CHASSIS-MIB",
            "path": "dell/F10-CHASSIS-MIB.mib"
        },
        {
            "mib": "F3-NTP-MIB",
            "path": "adva/F3-NTP-MIB.mib"
        },
        {
            "mib": "CTRON-PORTMAP-MIB",
            "path": "enterasys/CTRON-PORTMAP-MIB.mib"
        },
        {
            "mib": "HUAWEI-BULKSTAT-MIB",
            "path": "huawei/HUAWEI-BULKSTAT-MIB.mib"
        },
        {
            "mib": "BEGEMOT-HAST-MIB",
            "path": "pfsense/BEGEMOT-HAST-MIB.mib"
        },
        {
            "mib": "SL-OTN-MIB",
            "path": "packetlight/SL-OTN-MIB.mib"
        },
        {
            "mib": "CIENA-CES-EXT-LAG-MIB",
            "path": "ciena/CIENA-CES-EXT-LAG-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-AUTHUSERS-MIB",
            "path": "stormshield/STORMSHIELD-AUTHUSERS-MIB.mib"
        },
        {
            "mib": "GBNL3Rip-MIB",
            "path": "fs/GBNL3Rip-MIB.mib"
        },
        {
            "mib": "GRANDSTREAM-GWN-ROOT-MIB",
            "path": "grandstream/GRANDSTREAM-GWN-ROOT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DOT1Q-MIB",
            "path": "nokia/ALCATEL-IND1-DOT1Q-MIB.mib"
        },
        {
            "mib": "HP-ICF-FAULT-FINDER-MIB",
            "path": "hp/HP-ICF-FAULT-FINDER-MIB.mib"
        },
        {
            "mib": "HILLSTONE-DNS-MIB",
            "path": "hillstone/HILLSTONE-DNS-MIB.mib"
        },
        {
            "mib": "F3-OSPF-MIB",
            "path": "adva/F3-OSPF-MIB.mib"
        },
        {
            "mib": "XAVI-XG6846-MIB",
            "path": "inteno/XAVI-XG6846-MIB.mib"
        },
        {
            "mib": "PROSTAR-PWM",
            "path": "morningstar/PROSTAR-PWM.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-PROPERTY-MIB",
            "path": "cdata/NSCRTV-HFCEMS-PROPERTY-MIB.mib"
        },
        {
            "mib": "BCN-DNS-MIB",
            "path": "bluecatnetworks/BCN-DNS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-FEATURE-LICENSE-MIB",
            "path": "ciena/CIENA-CES-FEATURE-LICENSE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-QOS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-QOS-MIB.mib"
        },
        {
            "mib": "HILLSTONE-FAN-MIB",
            "path": "hillstone/HILLSTONE-FAN-MIB.mib"
        },
        {
            "mib": "GS-GXW42XX-MIB",
            "path": "grandstream/GS-GXW42XX-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-RDP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-RDP-MIB.mib"
        },
        {
            "mib": "ALTEON-CHEETAH-BWM-MIB",
            "path": "alteonos/ALTEON-CHEETAH-BWM-MIB.mib"
        },
        {
            "mib": "SL-PM-MIB",
            "path": "packetlight/SL-PM-MIB.mib"
        },
        {
            "mib": "BEGEMOT-HOSTRES-MIB",
            "path": "pfsense/BEGEMOT-HOSTRES-MIB.mib"
        },
        {
            "mib": "CIENA-CES-FILE-TRANSFER-MIB",
            "path": "ciena/CIENA-CES-FILE-TRANSFER-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-AUTOUPDATE-MIB",
            "path": "stormshield/STORMSHIELD-AUTOUPDATE-MIB.mib"
        },
        {
            "mib": "BCN-HA-MIB",
            "path": "bluecatnetworks/BCN-HA-MIB.mib"
        },
        {
            "mib": "F10-FIB-MIB",
            "path": "dell/F10-FIB-MIB.mib"
        },
        {
            "mib": "CTRON-POWER-SUPPLY-MIB",
            "path": "enterasys/CTRON-POWER-SUPPLY-MIB.mib"
        },
        {
            "mib": "NSCRTV-PON-TREE-EXT-MIB",
            "path": "cdata/NSCRTV-PON-TREE-EXT-MIB.mib"
        },
        {
            "mib": "F3-OTN-MIB",
            "path": "adva/F3-OTN-MIB.mib"
        },
        {
            "mib": "GBNL3RouteCommon-MIB",
            "path": "fs/GBNL3RouteCommon-MIB.mib"
        },
        {
            "mib": "HILLSTONE-IF-MIB",
            "path": "hillstone/HILLSTONE-IF-MIB.mib"
        },
        {
            "mib": "GS-HT8XX-MIB",
            "path": "grandstream/GS-HT8XX-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DOT1X-MIB",
            "path": "nokia/ALCATEL-IND1-DOT1X-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-RIP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-RIP-MIB.mib"
        },
        {
            "mib": "CTRON-PPC-BAD-PACKETS-MIB",
            "path": "enterasys/CTRON-PPC-BAD-PACKETS-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-HA-MIB",
            "path": "stormshield/STORMSHIELD-HA-MIB.mib"
        },
        {
            "mib": "HP-ICF-GENERIC-RPTR",
            "path": "hp/HP-ICF-GENERIC-RPTR.mib"
        },
        {
            "mib": "HILLSTONE-IP-MIB",
            "path": "hillstone/HILLSTONE-IP-MIB.mib"
        },
        {
            "mib": "BEGEMOT-IP-MIB",
            "path": "pfsense/BEGEMOT-IP-MIB.mib"
        },
        {
            "mib": "SL-PORT-MIB",
            "path": "packetlight/SL-PORT-MIB.mib"
        },
        {
            "mib": "BCN-LICENSE-MIB",
            "path": "bluecatnetworks/BCN-LICENSE-MIB.mib"
        },
        {
            "mib": "F10-IF-EXTENSION-MIB",
            "path": "dell/F10-IF-EXTENSION-MIB.mib"
        },
        {
            "mib": "NSCRTV-ROOT",
            "path": "cdata/NSCRTV-ROOT.mib"
        },
        {
            "mib": "SUNSAVER-MPPT",
            "path": "morningstar/SUNSAVER-MPPT.mib"
        },
        {
            "mib": "LCOS-SX-GENERAL-MIB",
            "path": "lancom/LCOS-SX-GENERAL-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-RIPNG-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-RIPNG-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-HEALTH-MONITOR-MIB",
            "path": "stormshield/STORMSHIELD-HEALTH-MONITOR-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DOT3-OAM-MIB",
            "path": "nokia/ALCATEL-IND1-DOT3-OAM-MIB.mib"
        },
        {
            "mib": "BEGEMOT-MIB",
            "path": "pfsense/BEGEMOT-MIB.mib"
        },
        {
            "mib": "HILLSTONE-IPSEC-MIB",
            "path": "hillstone/HILLSTONE-IPSEC-MIB.mib"
        },
        {
            "mib": "HUAWEI-CBQOS-MIB",
            "path": "huawei/HUAWEI-CBQOS-MIB.mib"
        },
        {
            "mib": "ALTEON-CHEETAH-LAYER4-MIB",
            "path": "alteonos/ALTEON-CHEETAH-LAYER4-MIB.mib"
        },
        {
            "mib": "HUAWEI-CCC-MIB",
            "path": "huawei/HUAWEI-CCC-MIB.mib"
        },
        {
            "mib": "BCN-NTP-MIB",
            "path": "bluecatnetworks/BCN-NTP-MIB.mib"
        },
        {
            "mib": "CIENA-CES-ICL-MIB",
            "path": "ciena/CIENA-CES-ICL-MIB.mib"
        },
        {
            "mib": "SURESINE",
            "path": "morningstar/SURESINE.mib"
        },
        {
            "mib": "GBNPlatformChassis-MIB",
            "path": "fs/GBNPlatformChassis-MIB.mib"
        },
        {
            "mib": "F3-PBB-MIB",
            "path": "adva/F3-PBB-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-HOSTS-MIB",
            "path": "stormshield/STORMSHIELD-HOSTS-MIB.mib"
        },
        {
            "mib": "RFC1213-MIB",
            "path": "cdata/RFC1213-MIB.mib"
        },
        {
            "mib": "CTRON-PRIORITY-CLASSIFY-MIB",
            "path": "enterasys/CTRON-PRIORITY-CLASSIFY-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-ROUTEMAP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-ROUTEMAP-MIB.mib"
        },
        {
            "mib": "GBNPlatformGNLink-MIB",
            "path": "fs/GBNPlatformGNLink-MIB.mib"
        },
        {
            "mib": "HILLSTONE-MODULE-MIB",
            "path": "hillstone/HILLSTONE-MODULE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DRCTM-MIB",
            "path": "nokia/ALCATEL-IND1-DRCTM-MIB.mib"
        },
        {
            "mib": "RADWIN-MIB-WINLINK1000",
            "path": "radwin/RADWIN-MIB-WINLINK1000.mib"
        },
        {
            "mib": "F10-LINK-AGGREGATION-MIB",
            "path": "dell/F10-LINK-AGGREGATION-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-IF-MIB",
            "path": "stormshield/STORMSHIELD-IF-MIB.mib"
        },
        {
            "mib": "F3-POPM-MIB",
            "path": "adva/F3-POPM-MIB.mib"
        },
        {
            "mib": "LCOS-MIB",
            "path": "lancom/LCOS-MIB.mib"
        },
        {
            "mib": "BEGEMOT-MIB2-MIB",
            "path": "pfsense/BEGEMOT-MIB2-MIB.mib"
        },
        {
            "mib": "HP-ICF-GPPC-MIB",
            "path": "hp/HP-ICF-GPPC-MIB.mib"
        },
        {
            "mib": "CIENA-CES-IP-INTERFACE-MIB",
            "path": "ciena/CIENA-CES-IP-INTERFACE-MIB.mib"
        },
        {
            "mib": "CTRON-PRIORITY-EXTENSIONS-MIB",
            "path": "enterasys/CTRON-PRIORITY-EXTENSIONS-MIB.mib"
        },
        {
            "mib": "ALTEON-CHEETAH-LAYER7-MIB",
            "path": "alteonos/ALTEON-CHEETAH-LAYER7-MIB.mib"
        },
        {
            "mib": "GBNPlatformOAM-MIB",
            "path": "fs/GBNPlatformOAM-MIB.mib"
        },
        {
            "mib": "BCN-PRODUCTS-MIB",
            "path": "bluecatnetworks/BCN-PRODUCTS-MIB.mib"
        },
        {
            "mib": "LCOS-SX-MIB",
            "path": "lancom/LCOS-SX-MIB.mib"
        },
        {
            "mib": "AVIAT-ALARM-REPORTING-MIB",
            "path": "aviat-wtm/AVIAT-ALARM-REPORTING-MIB.mib"
        },
        {
            "mib": "SL-RADIUS-MIB",
            "path": "packetlight/SL-RADIUS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-LDP-MIB",
            "path": "ciena/CIENA-CES-LDP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SAA-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-SAA-MIB.mib"
        },
        {
            "mib": "SPC200",
            "path": "cdata/SPC200.mib"
        },
        {
            "mib": "F10-M-SERIES-CHASSIS-MIB",
            "path": "dell/F10-M-SERIES-CHASSIS-MIB.mib"
        },
        {
            "mib": "RUIJIE-ENTITY-MIB",
            "path": "ruijie/RUIJIE-ENTITY-MIB.mib"
        },
        {
            "mib": "GBNPlatformOAMMailalarm-MIB",
            "path": "fs/GBNPlatformOAMMailalarm-MIB.mib"
        },
        {
            "mib": "TRISTAR-MPPT-600V",
            "path": "morningstar/TRISTAR-MPPT-600V.mib"
        },
        {
            "mib": "HILLSTONE-NTP-MIB",
            "path": "hillstone/HILLSTONE-NTP-MIB.mib"
        },
        {
            "mib": "HUAWEI-CDP-COMPLIANCE-MIB",
            "path": "huawei/HUAWEI-CDP-COMPLIANCE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-DVMRP-MIB",
            "path": "nokia/ALCATEL-IND1-DVMRP-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-IPSEC-STATS-MIB",
            "path": "stormshield/STORMSHIELD-IPSEC-STATS-MIB.mib"
        },
        {
            "mib": "F3-PORTMIRROR-MIB",
            "path": "adva/F3-PORTMIRROR-MIB.mib"
        },
        {
            "mib": "ALTEON-CHEETAH-NETWORK-MIB",
            "path": "alteonos/ALTEON-CHEETAH-NETWORK-MIB.mib"
        },
        {
            "mib": "CTRON-RATE-POLICING-MIB",
            "path": "enterasys/CTRON-RATE-POLICING-MIB.mib"
        },
        {
            "mib": "BCN-SMI-MIB",
            "path": "bluecatnetworks/BCN-SMI-MIB.mib"
        },
        {
            "mib": "GBNPlatformOAMSntpClient-MIB",
            "path": "fs/GBNPlatformOAMSntpClient-MIB.mib"
        },
        {
            "mib": "HILLSTONE-POWER-MIB",
            "path": "hillstone/HILLSTONE-POWER-MIB.mib"
        },
        {
            "mib": "F10-PRODUCTS-MIB",
            "path": "dell/F10-PRODUCTS-MIB.mib"
        },
        {
            "mib": "SL-RETIMER-MIB",
            "path": "packetlight/SL-RETIMER-MIB.mib"
        },
        {
            "mib": "TRISTAR-MPPT",
            "path": "morningstar/TRISTAR-MPPT.mib"
        },
        {
            "mib": "STORMSHIELD-MODEL-MIB",
            "path": "stormshield/STORMSHIELD-MODEL-MIB.mib"
        },
        {
            "mib": "HP-ICF-INST-MON",
            "path": "hp/HP-ICF-INST-MON.mib"
        },
        {
            "mib": "SPIDCOM-ALARM-MIB",
            "path": "cdata/SPIDCOM-ALARM-MIB.mib"
        },
        {
            "mib": "HUAWEI-CE-PING-MIB",
            "path": "huawei/HUAWEI-CE-PING-MIB.mib"
        },
        {
            "mib": "CTRON-REMOTE-ACCESS-MIB",
            "path": "enterasys/CTRON-REMOTE-ACCESS-MIB.mib"
        },
        {
            "mib": "RUIJIE-FIBER-MIB",
            "path": "ruijie/RUIJIE-FIBER-MIB.mib"
        },
        {
            "mib": "F10-S-SERIES-CHASSIS-MIB",
            "path": "dell/F10-S-SERIES-CHASSIS-MIB.mib"
        },
        {
            "mib": "SL-ROADM-MIB",
            "path": "packetlight/SL-ROADM-MIB.mib"
        },
        {
            "mib": "CIENA-CES-MAC-MIB",
            "path": "ciena/CIENA-CES-MAC-MIB.mib"
        },
        {
            "mib": "TRISTAR",
            "path": "morningstar/TRISTAR.mib"
        },
        {
            "mib": "ALCATEL-IND1-SERVICE-MGR-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-SERVICE-MGR-MIB.mib"
        },
        {
            "mib": "ALTEON-CHEETAH-SWITCH-MIB",
            "path": "alteonos/ALTEON-CHEETAH-SWITCH-MIB.mib"
        },
        {
            "mib": "SPIDCOM-MIB",
            "path": "cdata/SPIDCOM-MIB.mib"
        },
        {
            "mib": "AVIAT-G826-MIB",
            "path": "aviat-wtm/AVIAT-G826-MIB.mib"
        },
        {
            "mib": "HP-ICF-IP-ROUTING",
            "path": "hp/HP-ICF-IP-ROUTING.mib"
        },
        {
            "mib": "F3-PTP-MIB",
            "path": "adva/F3-PTP-MIB.mib"
        },
        {
            "mib": "F3-PWE3-MIB",
            "path": "adva/F3-PWE3-MIB.mib"
        },
        {
            "mib": "GBNPlatformOAMSsh-MIB",
            "path": "fs/GBNPlatformOAMSsh-MIB.mib"
        },
        {
            "mib": "ALTEON-CS-PHYSICAL-MIB",
            "path": "alteonos/ALTEON-CS-PHYSICAL-MIB.mib"
        },
        {
            "mib": "HUAWEI-CLOCK-MIB",
            "path": "huawei/HUAWEI-CLOCK-MIB.mib"
        },
        {
            "mib": "HILLSTONE-PRODUCTS-MIB",
            "path": "hillstone/HILLSTONE-PRODUCTS-MIB.mib"
        },
        {
            "mib": "F10-Z-SERIES-CHASSIS-MIB",
            "path": "dell/F10-Z-SERIES-CHASSIS-MIB.mib"
        },
        {
            "mib": "BCN-SYSTEM-MIB",
            "path": "bluecatnetworks/BCN-SYSTEM-MIB.mib"
        },
        {
            "mib": "SL-SECU-MIB",
            "path": "packetlight/SL-SECU-MIB.mib"
        },
        {
            "mib": "BEGEMOT-NETGRAPH-MIB",
            "path": "pfsense/BEGEMOT-NETGRAPH.mib"
        },
        {
            "mib": "STORMSHIELD-OVPN-MIB",
            "path": "stormshield/STORMSHIELD-OVPNTABLE-MIB.mib"
        },
        {
            "mib": "RUIJIE-INTERFACE-MIB",
            "path": "ruijie/RUIJIE-INTERFACE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-E-SERVICE-MIB",
            "path": "nokia/ALCATEL-IND1-E-SERVICE-MIB.mib"
        },
        {
            "mib": "SPIDCOM-NOTIFICATION-MIB",
            "path": "cdata/SPIDCOM-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-ALARM-TC",
            "path": "ericsson/ERICSSON-ROUTER-ALARM-TC.mib"
        },
        {
            "mib": "CTRON-ROUTERS-INTERNAL-MIB",
            "path": "enterasys/CTRON-ROUTERS-INTERNAL-MIB.mib"
        },
        {
            "mib": "GBNPlatformOAMSyslog-MIB",
            "path": "fs/GBNPlatformOAMSyslog-MIB.mib"
        },
        {
            "mib": "CIENA-CES-MCAST-FILTER-MIB",
            "path": "ciena/CIENA-CES-MCAST-FILTER-MIB.mib"
        },
        {
            "mib": "AVIAT-MODEM-MIB",
            "path": "aviat-wtm/AVIAT-MODEM-MIB.mib"
        },
        {
            "mib": "HP-ICF-IPCONFIG",
            "path": "hp/HP-ICF-IPCONFIG.mib"
        },
        {
            "mib": "F3-SHG-MIB",
            "path": "adva/F3-SHG-MIB.mib"
        },
        {
            "mib": "ALTEON-ROOT-MIB",
            "path": "alteonos/ALTEON-ROOT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SESSION-MGR-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-SESSION-MGR-MIB.mib"
        },
        {
            "mib": "BEGEMOT-PF-MIB",
            "path": "pfsense/BEGEMOT-PF-MIB.mib"
        },
        {
            "mib": "HILLSTONE-SMI",
            "path": "hillstone/HILLSTONE-SMI.mib"
        },
        {
            "mib": "SPIDCOM-TRAPS",
            "path": "cdata/SPIDCOM-TRAPS.mib"
        },
        {
            "mib": "HUAWEI-CONFIG-MAN-MIB",
            "path": "huawei/HUAWEI-CONFIG-MAN-MIB.mib"
        },
        {
            "mib": "AVIAT-RF-MIB",
            "path": "aviat-wtm/AVIAT-RF-MIB.mib"
        },
        {
            "mib": "SL-SFP-MIB",
            "path": "packetlight/SL-SFP-MIB.mib"
        },
        {
            "mib": "F3-SYNC-MIB",
            "path": "adva/F3-SYNC-MIB.mib"
        },
        {
            "mib": "BCN-TC-MIB",
            "path": "bluecatnetworks/BCN-TC-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-CPU-METER-MIB",
            "path": "ericsson/ERICSSON-ROUTER-CPU-METER-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-ERP-MIB",
            "path": "nokia/ALCATEL-IND1-ERP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SLB-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-SLB-MIB.mib"
        },
        {
            "mib": "FASTPATH-BOXSERVICES-PRIVATE-MIB",
            "path": "dell/FASTPATH-BOXSERVICES-PRIVATE-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-POLICY-MIB",
            "path": "stormshield/STORMSHIELD-POLICY-MIB.mib"
        },
        {
            "mib": "GBNPlatformOAMTelnet-MIB",
            "path": "fs/GBNPlatformOAMTelnet-MIB.mib"
        },
        {
            "mib": "HUAWEI-CPU-MIB",
            "path": "huawei/HUAWEI-CPU-MIB.mib"
        },
        {
            "mib": "AVIAT-RXPERFORMANCE-EX-MIB",
            "path": "aviat-wtm/AVIAT-RXPERFORMANCE-EX-MIB.mib"
        },
        {
            "mib": "BEGEMOT-SNMPD-MIB",
            "path": "pfsense/BEGEMOT-SNMPD.mib"
        },
        {
            "mib": "ALCATEL-IND1-ETHERNET-OAM-MIB",
            "path": "nokia/ALCATEL-IND1-ETHERNET-OAM-MIB.mib"
        },
        {
            "mib": "RUIJIE-MEMORY-MIB",
            "path": "ruijie/RUIJIE-MEMORY-MIB.mib"
        },
        {
            "mib": "CIENA-CES-MGMT-INTERFACE-MIB",
            "path": "ciena/CIENA-CES-MGMT-INTERFACE-MIB.mib"
        },
        {
            "mib": "ALTEON-TIGON-SWITCH-MIB",
            "path": "alteonos/ALTEON-TIGON-SWITCH-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SNMP-AGENT-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-SNMP-AGENT-MIB.mib"
        },
        {
            "mib": "HILLSTONE-STATISTICS-MIB",
            "path": "hillstone/HILLSTONE-STATISTICS-MIB.mib"
        },
        {
            "mib": "VENDOR-COMMON-MIB",
            "path": "cdata/VENDOR-COMMON-MIB.mib"
        },
        {
            "mib": "CTRON-ROUTERS-MIB",
            "path": "enterasys/CTRON-ROUTERS-MIB.mib"
        },
        {
            "mib": "SL-SNTP-MIB",
            "path": "packetlight/SL-SNTP-MIB.mib"
        },
        {
            "mib": "BCN-TFTP-MIB",
            "path": "bluecatnetworks/BCN-TFTP-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-ENVMON-CAP",
            "path": "ericsson/ERICSSON-ROUTER-ENVMON-CAP.mib"
        },
        {
            "mib": "AVIAT-RXPERFORMANCE-MIB",
            "path": "aviat-wtm/AVIAT-RXPERFORMANCE-MIB.mib"
        },
        {
            "mib": "HUAWEI-DAD-MIB",
            "path": "huawei/HUAWEI-DAD-MIB.mib"
        },
        {
            "mib": "HP-ICF-JUMBO-MIB",
            "path": "hp/HP-ICF-JUMBO-MIB.mib"
        },
        {
            "mib": "FORCE10-BGP4-V2-MIB",
            "path": "dell/FORCE10-BGP4-V2-MIB.mib"
        },
        {
            "mib": "GBNServiceCM-MIB",
            "path": "fs/GBNServiceCM-MIB.mib"
        },
        {
            "mib": "BEGEMOT-WIRELESS-MIB",
            "path": "pfsense/BEGEMOT-WIRELESS-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-PROPERTY-MIB-DEPRECATED",
            "path": "stormshield/STORMSHIELD-PROPERTY-MIB-DEPRECATED.mib"
        },
        {
            "mib": "RUIJIE-PROCESS-MIB",
            "path": "ruijie/RUIJIE-PROCESS-MIB.mib"
        },
        {
            "mib": "ALTEON-TS-NETWORK-MIB",
            "path": "alteonos/ALTEON-TS-NETWORK-MIB.mib"
        },
        {
            "mib": "F3-SYNCJACK-MIB",
            "path": "adva/F3-SYNCJACK-MIB.mib"
        },
        {
            "mib": "CIENA-CES-MODULE-MIB",
            "path": "ciena/CIENA-CES-MODULE-MIB.mib"
        },
        {
            "mib": "FORCE10-COPY-CONFIG-MIB",
            "path": "dell/FORCE10-COPY-CONFIG-MIB.mib"
        },
        {
            "mib": "SL-SONET-MIB",
            "path": "packetlight/SL-SONET-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-GROUP-MOBILITY-MIB",
            "path": "nokia/ALCATEL-IND1-GROUP-MOBILITY-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SYSTEM-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-SYSTEM-MIB.mib"
        },
        {
            "mib": "GBNServiceMACAUTHEN-MIB",
            "path": "fs/GBNServiceMACAUTHEN-MIB.mib"
        },
        {
            "mib": "HILLSTONE-SYSTEM-MIB",
            "path": "hillstone/HILLSTONE-SYSTEM-MIB.mib"
        },
        {
            "mib": "XXX-MIB",
            "path": "cdata/XXX-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-PROPERTY-MIB",
            "path": "stormshield/STORMSHIELD-PROPERTY-MIB.mib"
        },
        {
            "mib": "BLUECATNETWORKS-MIB",
            "path": "bluecatnetworks/BLUECATNETWORKS-MIB.mib"
        },
        {
            "mib": "HP-ICF-L3MAC-MIB",
            "path": "hp/HP-ICF-L3MAC-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-ENVMON-MIB",
            "path": "ericsson/ERICSSON-ROUTER-ENVMON-MIB.mib"
        },
        {
            "mib": "FOKUS-MIB",
            "path": "pfsense/FOKUS-MIB.mib"
        },
        {
            "mib": "HUAWEI-DATASYNC-MIB",
            "path": "huawei/HUAWEI-DATASYNC-MIB.mib"
        },
        {
            "mib": "AVIAT-SWMANAGEMENT-MIB",
            "path": "aviat-wtm/AVIAT-SWMANAGEMENT-MIB.mib"
        },
        {
            "mib": "F3-TIMEZONE-MIB",
            "path": "adva/F3-TIMEZONE-MIB.mib"
        },
        {
            "mib": "GBNServiceRADIUS-MIB",
            "path": "fs/GBNServiceRADIUS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-GVRP-MIB",
            "path": "nokia/ALCATEL-IND1-GVRP-MIB.mib"
        },
        {
            "mib": "CTRON-SFCS-MIB",
            "path": "enterasys/CTRON-SFCS-MIB.mib"
        },
        {
            "mib": "RUIJIE-SMI",
            "path": "ruijie/RUIJIE-SMI.mib"
        },
        {
            "mib": "HILLSTONE-TEMPERATURE-MIB",
            "path": "hillstone/HILLSTONE-TEMPERATURE-MIB.mib"
        },
        {
            "mib": "ALTEON-TS-PHYSICAL-MIB",
            "path": "alteonos/ALTEON-TS-PHYSICAL-MIB.mib"
        },
        {
            "mib": "CIENA-CES-MPLS-MIB",
            "path": "ciena/CIENA-CES-MPLS-MIB.mib"
        },
        {
            "mib": "FORCE10-MONITORING-MIB",
            "path": "dell/FORCE10-MONITORING-MIB.mib"
        },
        {
            "mib": "EXTRAHOP-MIB",
            "path": "extrahop/EXTRAHOP-MIB.mib"
        },
        {
            "mib": "HUAWEI-DC-TRAP-MIB",
            "path": "huawei/HUAWEI-DC-TRAP-MIB.mib"
        },
        {
            "mib": "GBNServiceRMON-MIB",
            "path": "fs/GBNServiceRMON-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-QOS-MIB",
            "path": "stormshield/STORMSHIELD-QOS-MIB.mib"
        },
        {
            "mib": "FREEBSD-MIB",
            "path": "pfsense/FREEBSD-MIB.mib"
        },
        {
            "mib": "HP-ICF-LINKTEST",
            "path": "hp/HP-ICF-LINKTEST.mib"
        },
        {
            "mib": "CIENA-CES-NOTIFICATIONS-CONTROL",
            "path": "ciena/CIENA-CES-NOTIFICATIONS-CONTROL-MIB.mib"
        },
        {
            "mib": "BDCOM-MEMORY-POOL-MIB",
            "path": "bdcom/BDCOM-MEMORY-POOL-MIB.mib"
        },
        {
            "mib": "RUIJIE-SYSTEM-MIB",
            "path": "ruijie/RUIJIE-SYSTEM-MIB.mib"
        },
        {
            "mib": "AVIAT-TEXTCONVENTION-MIB",
            "path": "aviat-wtm/AVIAT-TEXTCONVENTION-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-OPTICAL-TRANSCEIVER-MIB",
            "path": "ericsson/ERICSSON-ROUTER-OPTICAL-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "SL-SONET-SUP-PM-MIB",
            "path": "packetlight/SL-SONET-SUP-PM-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-BASE-MIB",
            "path": "enterasys/CTRON-SFPS-BASE-MIB.mib"
        },
        {
            "mib": "HP-ICF-MLD-MIB",
            "path": "hp/HP-ICF-MLD-MIB.mib"
        },
        {
            "mib": "HILLSTONE-ZONE-MIB",
            "path": "hillstone/HILLSTONE-ZONE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-CHASSIS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-CHASSIS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-HEALTH-MIB",
            "path": "nokia/ALCATEL-IND1-HEALTH-MIB.mib"
        },
        {
            "mib": "HUAWEI-DEVICE-EXT-MIB",
            "path": "huawei/HUAWEI-DEVICE-EXT-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-SFP-MIB",
            "path": "ericsson/ERICSSON-ROUTER-SFP-MIB.mib"
        },
        {
            "mib": "FORCE10-MSTP-MIB",
            "path": "dell/FORCE10-MSTP-MIB.mib"
        },
        {
            "mib": "F3-TWAMP-MIB",
            "path": "adva/F3-TWAMP-MIB.mib"
        },
        {
            "mib": "RUIJIE-TC",
            "path": "ruijie/RUIJIE-TC.mib"
        },
        {
            "mib": "OMNITRON-MIB",
            "path": "omnitron/OMNITRON-MIB.mib"
        },
        {
            "mib": "SL-TESTS-MIB",
            "path": "packetlight/SL-TESTS-MIB.mib"
        },
        {
            "mib": "PARKS-PK700",
            "path": "parks/PARKS-PK700.mib"
        },
        {
            "mib": "CTRON-SFPS-BINDERY-MIB",
            "path": "enterasys/CTRON-SFPS-BINDERY-MIB.mib"
        },
        {
            "mib": "GBNServiceTACACS-MIB",
            "path": "fs/GBNServiceTACACS-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-ROUTE-MIB",
            "path": "stormshield/STORMSHIELD-ROUTE-MIB.mib"
        },
        {
            "mib": "MEF-SOAM-PM-MIB",
            "path": "mef/MEF-SOAM-PM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IGMP-MIB",
            "path": "nokia/ALCATEL-IND1-IGMP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-FILTER-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-FILTER-MIB.mib"
        },
        {
            "mib": "NTRON714FX6-MIB",
            "path": "redlion/NTRON714FX6-MIB.mib"
        },
        {
            "mib": "FORCE10-SMI",
            "path": "dell/FORCE10-SMI.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-SMI",
            "path": "ericsson/ERICSSON-ROUTER-SMI.mib"
        },
        {
            "mib": "STXN-GLOBALREGISTER-MIB",
            "path": "aviat-wtm/STXN-GLOBALREGISTER-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-SERVICES-MIB",
            "path": "stormshield/STORMSHIELD-SERVICES-MIB.mib"
        },
        {
            "mib": "LAG-ARCH-MIB",
            "path": "fs/LAG-ARCH-MIB.mib"
        },
        {
            "mib": "HP-ICF-OID",
            "path": "hp/HP-ICF-OID.mib"
        },
        {
            "mib": "FSP150-MIB",
            "path": "adva/FSP150-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-CALL-MIB",
            "path": "enterasys/CTRON-SFPS-CALL-MIB.mib"
        },
        {
            "mib": "CIENA-CES-OAM-MIB",
            "path": "ciena/CIENA-CES-OAM-MIB.mib"
        },
        {
            "mib": "BDCOM-PROCESS-MIB",
            "path": "bdcom/BDCOM-PROCESS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-GLOBAL-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-GLOBAL-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-INLINE-POWER-MIB",
            "path": "nokia/ALCATEL-IND1-INLINE-POWER-MIB.mib"
        },
        {
            "mib": "ipPbxNs-MIB",
            "path": "panasonic/ipPbxNs-MIB.mib"
        },
        {
            "mib": "OMNITRON-POE-MIB",
            "path": "omnitron/OMNITRON-POE-MIB.mib"
        },
        {
            "mib": "HUAWEI-DEVICE-MIB",
            "path": "huawei/HUAWEI-DEVICE-MIB.mib"
        },
        {
            "mib": "OA-SFP-MIB",
            "path": "adva/OA-SFP-MIB.mib"
        },
        {
            "mib": "SL-TRAP-MIB",
            "path": "packetlight/SL-TRAP-MIB.mib"
        },
        {
            "mib": "FORCE10-SYSTEM-COMPONENT-MIB",
            "path": "dell/FORCE10-SYSTEM-COMPONENT-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-SYS-RESOURCES-MIB",
            "path": "ericsson/ERICSSON-ROUTER-SYS-RESOURCES-MIB.mib"
        },
        {
            "mib": "NETONIX-SWITCH-MIB",
            "path": "netonix/NETONIX-SWITCH-MIB.mib"
        },
        {
            "mib": "LLDPPRIVATE-MIB",
            "path": "fs/LLDPPRIVATE-MIB.mib"
        },
        {
            "mib": "SIXNET-MIB",
            "path": "redlion/SIXNET-MIB.mib"
        },
        {
            "mib": "DANTHERM-COOLING-MIB",
            "path": "dantherm/DANTHERM-COOLING-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-CHASSIS-MIB",
            "path": "enterasys/CTRON-SFPS-CHASSIS-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-SMI-MIB",
            "path": "stormshield/STORMSHIELD-SMI-MIB.mib"
        },
        {
            "mib": "MEF-SOAM-TC-MIB",
            "path": "mef/MEF-SOAM-TC-MIB.mib"
        },
        {
            "mib": "CIENA-CES-OSPF-MIB",
            "path": "ciena/CIENA-CES-OSPF-MIB.mib"
        },
        {
            "mib": "OS-COMMON-TC-MIB",
            "path": "adva/OS-COMMON-TC-MIB.mib"
        },
        {
            "mib": "HP-ICF-OSPF",
            "path": "hp/HP-ICF-OSPF.mib"
        },
        {
            "mib": "ALCATEL-IND1-INTERSWITCH-PROTOCOL-MIB",
            "path": "nokia/ALCATEL-IND1-INTERSWITCH-PROTOCOL-MIB.mib"
        },
        {
            "mib": "OMNITRON-TC-MIB",
            "path": "omnitron/OMNITRON-TC-MIB.mib"
        },
        {
            "mib": "EES-POWER-FERRO-MIB",
            "path": "emerson/EES-POWER-FERRO-MIB.mib"
        },
        {
            "mib": "SL-XPDR-MIB",
            "path": "packetlight/SL-XPDR-MIB.mib"
        },
        {
            "mib": "BDCOM-QOS-PIB-MIB",
            "path": "bdcom/BDCOM-QOS-PIB-MIB.mib"
        },
        {
            "mib": "PAN-COMMON-MIB",
            "path": "paloaltonetworks/PAN-COMMON-MIB.mib"
        },
        {
            "mib": "ERICSSON-ROUTER-TC",
            "path": "ericsson/ERICSSON-ROUTER-TC.mib"
        },
        {
            "mib": "CTRON-SFPS-COMMON-MIB",
            "path": "enterasys/CTRON-SFPS-COMMON-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-LDP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-LDP-MIB.mib"
        },
        {
            "mib": "AIRESPACE-REF-MIB",
            "path": "cisco/AIRESPACE-REF-MIB.mib"
        },
        {
            "mib": "FREENAS-MIB",
            "path": "ixsystems/FREENAS-MIB.mib"
        },
        {
            "mib": "MAC-NOTIFICATION-MIB",
            "path": "fs/MAC-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "EES-POWER-MIB",
            "path": "emerson/EES-POWER-MIB.mib"
        },
        {
            "mib": "MEF-UNI-EVC-MIB",
            "path": "mef/MEF-UNI-EVC-MIB.mib"
        },
        {
            "mib": "ELTEX-LTP8X-STANDALONE",
            "path": "eltex/ELTEX-LTP8X-STANDALONE.mib"
        },
        {
            "mib": "STORMSHIELD-SYSTEM-MONITOR-MIB",
            "path": "stormshield/STORMSHIELD-SYSTEM-MONITOR-MIB.mib"
        },
        {
            "mib": "AT-ALMMON-MIB",
            "path": "awplus/AT-ALMMON-MIB.mib"
        },
        {
            "mib": "FORCE10-TC",
            "path": "dell/FORCE10-TC.mib"
        },
        {
            "mib": "NSCRTV-EPON-ALARM-MGM-MIB",
            "path": "fs/NSCRTV-EPON-ALARM-MGM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IP-MIB",
            "path": "nokia/ALCATEL-IND1-IP-MIB.mib"
        },
        {
            "mib": "HUAWEI-DHCP-SNOOPING-MIB",
            "path": "huawei/HUAWEI-DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "AIRESPACE-SWITCHING-MIB",
            "path": "cisco/AIRESPACE-SWITCHING-MIB.mib"
        },
        {
            "mib": "ERICSSON-TC-MIB",
            "path": "ericsson/ERICSSON-TC-MIB.mib"
        },
        {
            "mib": "RIELLO-MIB",
            "path": "riello/RIELLO-MIB.mib"
        },
        {
            "mib": "HP-ICF-PIM",
            "path": "hp/HP-ICF-PIM.mib"
        },
        {
            "mib": "PAN-ENTITY-EXT-MIB",
            "path": "paloaltonetworks/PAN-ENTITY-EXT-MIB.mib"
        },
        {
            "mib": "LANTRONIX-MIB",
            "path": "lantronix/LANTRONIX-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-CONN-MIB",
            "path": "enterasys/CTRON-SFPS-CONN-MIB.mib"
        },
        {
            "mib": "BDCOM-SMI",
            "path": "bdcom/BDCOM-SMI.mib"
        },
        {
            "mib": "NSCRTV-EPON-IGMP-MGM-MIB",
            "path": "fs/NSCRTV-EPON-IGMP-MGM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-MPLS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-MPLS-MIB.mib"
        },
        {
            "mib": "NETSURE-MIB-004-A",
            "path": "emerson/NETSURE-MIB-004-A.mib"
        },
        {
            "mib": "ERICSSON-TOP-MIB",
            "path": "ericsson/ERICSSON-TOP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPM-VLAN-MIB",
            "path": "nokia/ALCATEL-IND1-IPM-VLAN-MIB.mib"
        },
        {
            "mib": "ELTEX-LTP8X",
            "path": "eltex/ELTEX-LTP8X.mib"
        },
        {
            "mib": "STORMSHIELD-VPN-MIB",
            "path": "stormshield/STORMSHIELD-VPN-MIB.mib"
        },
        {
            "mib": "TRUENAS-MIB",
            "path": "ixsystems/TRUENAS-MIB.mib"
        },
        {
            "mib": "HUAWEI-DHCPR-MIB",
            "path": "huawei/HUAWEI-DHCPR-MIB.mib"
        },
        {
            "mib": "RIELLOMDU-MIB",
            "path": "riello/RIELLOMDU-MIB.mib"
        },
        {
            "mib": "FORCE10-TRAP-EVENT-MIB",
            "path": "dell/FORCE10-TRAP-EVENT-MIB.mib"
        },
        {
            "mib": "BDCOM-TC",
            "path": "bdcom/BDCOM-TC.mib"
        },
        {
            "mib": "AT-ATMF-MIB",
            "path": "awplus/AT-ATMF-MIB.mib"
        },
        {
            "mib": "HP-ICF-POE-MIB",
            "path": "hp/HP-ICF-POE-MIB.mib"
        },
        {
            "mib": "AIRESPACE-WIRELESS-MIB",
            "path": "cisco/AIRESPACE-WIRELESS-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPON-PERFORMANCE-STAT-MIB",
            "path": "fs/NSCRTV-EPON-PERFORMANCE-STAT-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-VPNIKESA-MIB",
            "path": "stormshield/STORMSHIELD-VPNIKESA-MIB.mib"
        },
        {
            "mib": "CIENA-CES-PM",
            "path": "ciena/CIENA-CES-PM-MIB.mib"
        },
        {
            "mib": "HP-ICF-PROVIDER-BRIDGE",
            "path": "hp/HP-ICF-PROVIDER-BRIDGE.mib"
        },
        {
            "mib": "LANTRONIX-SLC-MIB",
            "path": "lantronix/LANTRONIX-SLC-MIB.mib"
        },
        {
            "mib": "PAN-GLOBAL-REG",
            "path": "paloaltonetworks/PAN-GLOBAL-REG-MIB.mib"
        },
        {
            "mib": "HUAWEI-DHCPS-MIB",
            "path": "huawei/HUAWEI-DHCPS-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-CONNECTION-MIB",
            "path": "enterasys/CTRON-SFPS-CONNECTION-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-OAM-TEST-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-OAM-TEST-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPMRM-MIB",
            "path": "nokia/ALCATEL-IND1-IPMRM-MIB.mib"
        },
        {
            "mib": "ARHANGELSK-GLOBAL-REG",
            "path": "majorpower/ARHANGELSK-GLOBAL-REG.mib"
        },
        {
            "mib": "RIELLOMTS-MIB",
            "path": "riello/RIELLOMTS-MIB.mib"
        },
        {
            "mib": "ELTEX-PP4",
            "path": "eltex/ELTEX-PP4.mib"
        },
        {
            "mib": "CIENA-CES-PORT-MIB",
            "path": "ciena/CIENA-CES-PORT-MIB.mib"
        },
        {
            "mib": "AT-BOARDS-MIB",
            "path": "awplus/AT-BOARDS-MIB.mib"
        },
        {
            "mib": "VEC-MIBv5-9",
            "path": "emerson/VEC-MIBv5-9.mib"
        },
        {
            "mib": "STORMSHIELD-VPNSA-MIB",
            "path": "stormshield/STORMSHIELD-VPNSA-MIB.mib"
        },
        {
            "mib": "ASYNCOS-MAIL-MIB",
            "path": "cisco/ASYNCOS-MAIL-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPON-QOS-MGM-MIB",
            "path": "fs/NSCRTV-EPON-QOS-MGM-MIB.mib"
        },
        {
            "mib": "MINI-LINK-MIB",
            "path": "ericsson/MINI-LINK-MIB.mib"
        },
        {
            "mib": "HIK-DEVICE-MIB",
            "path": "hikvision/HIK-DEVICE-MIB.mib"
        },
        {
            "mib": "NMS-CARD-SYS-MIB",
            "path": "bdcom/NMS-CARD-SYS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPMS-MIB",
            "path": "nokia/ALCATEL-IND1-IPMS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-PORT-XCVR-MIB",
            "path": "ciena/CIENA-CES-PORT-XCVR-MIB.mib"
        },
        {
            "mib": "IDRAC-MIB-SMIv2",
            "path": "dell/IDRAC-MIB-SMIv2.mib"
        },
        {
            "mib": "HP-ICF-RATE-LIMIT-MIB",
            "path": "hp/HP-ICF-RATE-LIMIT-MIB.mib"
        },
        {
            "mib": "PAN-GLOBAL-TC",
            "path": "paloaltonetworks/PAN-GLOBAL-TC-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-DIAGSTATS-MIB",
            "path": "enterasys/CTRON-SFPS-DIAGSTATS-MIB.mib"
        },
        {
            "mib": "ASYNCOSWEBSECURITYAPPLIANCE-MIB",
            "path": "cisco/ASYNCOSWEBSECURITYAPPLIANCE-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPON-SNI-MIB",
            "path": "fs/NSCRTV-EPON-SNI-MIB.mib"
        },
        {
            "mib": "ELTEX-SMI-ACTUAL",
            "path": "eltex/ELTEX-SMI-ACTUAL.mib"
        },
        {
            "mib": "ELTEK-BC2000-DC-POWER-MIB",
            "path": "eltek/ELTEK-BC2000-DC-POWER-MIB.mib"
        },
        {
            "mib": "AT-CHASSIS-MIB",
            "path": "awplus/AT-CHASSIS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-PTP-MIB",
            "path": "ciena/CIENA-CES-PTP-MIB.mib"
        },
        {
            "mib": "PAN-LC-MIB",
            "path": "paloaltonetworks/PAN-LC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-PORT-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-PORT-MIB.mib"
        },
        {
            "mib": "HP-ICF-RIP",
            "path": "hp/HP-ICF-RIP.mib"
        },
        {
            "mib": "NMS-CHASSIS",
            "path": "bdcom/NMS-CHASSIS.mib"
        },
        {
            "mib": "BASIS-MIB",
            "path": "cisco/BASIS-MIB.mib"
        },
        {
            "mib": "HUAWEI-DHCPV6-SERVER-MIB",
            "path": "huawei/HUAWEI-DHCPV6-SERVER-MIB.mib"
        },
        {
            "mib": "RIELLOUPS-MIB",
            "path": "riello/RIELLOUPS-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPON-STP-MGM-MIB",
            "path": "fs/NSCRTV-EPON-STP-MGM-MIB.mib"
        },
        {
            "mib": "PT-FM-MIB",
            "path": "ericsson/PT-FM-MIB.mib"
        },
        {
            "mib": "HIKVISION-MIB",
            "path": "hikvision/HIKVISION-MIB.mib"
        },
        {
            "mib": "STORMSHIELD-VPNSP-MIB",
            "path": "stormshield/STORMSHIELD-VPNSP-MIB.mib"
        },
        {
            "mib": "SO-MUX-MIB",
            "path": "solidoptics/SO-MUX-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPRM-MIB",
            "path": "nokia/ALCATEL-IND1-IPRM-MIB.mib"
        },
        {
            "mib": "AT-DHCPSN-MIB",
            "path": "awplus/AT-DHCPSN-MIB.mib"
        },
        {
            "mib": "IDRAC-MIB",
            "path": "dell/IDRAC-MIB.mib"
        },
        {
            "mib": "HIPATH-WIRELESS-DOT11-EXTNS-MIB",
            "path": "ewc/HIPATH-WIRELESS-DOT11-EXTNS-MIB.mib"
        },
        {
            "mib": "HP-ICF-SECURITY",
            "path": "hp/HP-ICF-SECURITY.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-QOS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-QOS-MIB.mib"
        },
        {
            "mib": "HUAWEI-DISMAN-PING-MIB",
            "path": "huawei/HUAWEI-DISMAN-PING-MIB.mib"
        },
        {
            "mib": "NMS-EPON-OLT-PON",
            "path": "bdcom/NMS-EPON-OLT-PON.mib"
        },
        {
            "mib": "ELTEK-COMMON-MIB",
            "path": "eltek/ELTEK-COMMON-MIB.mib"
        },
        {
            "mib": "SENSORTRAP-MIB",
            "path": "riello/SENSORTRAP-MIB.mib"
        },
        {
            "mib": "CIENA-CES-RADIUS-CLIENT-MIB",
            "path": "ciena/CIENA-CES-RADIUS-CLIENT-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-DIRECTORY-MIB",
            "path": "enterasys/CTRON-SFPS-DIRECTORY-MIB.mib"
        },
        {
            "mib": "PT-MIB",
            "path": "ericsson/PT-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPON-SYSTEM-MIB",
            "path": "fs/NSCRTV-EPON-SYSTEM-MIB.mib"
        },
        {
            "mib": "NSCRTV-EXTENSION-GY",
            "path": "glassway/NSCRTV-EXTENSION-GY.mib"
        },
        {
            "mib": "RADIO-BRIDGE-MIB",
            "path": "siklu/RADIO-BRIDGE-MIB.mib"
        },
        {
            "mib": "PAN-PRODUCTS-MIB",
            "path": "paloaltonetworks/PAN-PRODUCT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPRMV6-MIB",
            "path": "nokia/ALCATEL-IND1-IPRMV6-MIB.mib"
        },
        {
            "mib": "AT-DNS-CLIENT-MIB",
            "path": "awplus/AT-DNS-CLIENT.mib"
        },
        {
            "mib": "WISI-GTMODULES-MIB",
            "path": "wisi/WISI-GTMODULES-MIB.mib"
        },
        {
            "mib": "ELTEK-DISTRIBUTED-MIB",
            "path": "eltek/ELTEK-DISTRIBUTED-MIB.mib"
        },
        {
            "mib": "HP-ICF-SNMP-MIB",
            "path": "hp/HP-ICF-SNMP-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-ESYS-MIB",
            "path": "enterasys/CTRON-SFPS-ESYS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-SAP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-SAP-MIB.mib"
        },
        {
            "mib": "CERENT-454-MIB",
            "path": "cisco/CERENT-454-MIB.mib"
        },
        {
            "mib": "NMS-FAN-TRAP",
            "path": "bdcom/NMS-FAN-TRAP.mib"
        },
        {
            "mib": "HUAWEI-DLDP-MIB",
            "path": "huawei/HUAWEI-DLDP-MIB.mib"
        },
        {
            "mib": "PT-MONITOR-MIB",
            "path": "ericsson/PT-MONITOR-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-ALARMS-MIB",
            "path": "glassway/NSCRTV-HFCEMS-ALARMS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-RAPS-MIB",
            "path": "ciena/CIENA-CES-RAPS-MIB.mib"
        },
        {
            "mib": "AT-ENVMONv2-MIB",
            "path": "awplus/AT-ENVMONv2-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPON-VLAN-MGM-MIB",
            "path": "fs/NSCRTV-EPON-VLAN-MGM-MIB.mib"
        },
        {
            "mib": "WISI-GTSENSORS-MIB",
            "path": "wisi/WISI-GTSENSORS-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-EVENTLOG-MIB",
            "path": "enterasys/CTRON-SFPS-EVENTLOG-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPSEC-MIB",
            "path": "nokia/ALCATEL-IND1-IPSEC-MIB.mib"
        },
        {
            "mib": "HP-ICF-STACK",
            "path": "hp/HP-ICF-STACK.mib"
        },
        {
            "mib": "HIPATH-WIRELESS-HWC-MIB",
            "path": "ewc/HIPATH-WIRELESS-HWC-MIB.mib"
        },
        {
            "mib": "ADVANTECH-COMMON-MIB",
            "path": "advantech/ADVANTECH-COMMON-MIB.mib"
        },
        {
            "mib": "NSCRTV-EPONEOC-EPON-MIB",
            "path": "fs/NSCRTV-EPONEOC-EPON-MIB.mib"
        },
        {
            "mib": "AT-EPSRv2-MIB",
            "path": "awplus/AT-EPSRv2-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-COMMON-MIB",
            "path": "glassway/NSCRTV-HFCEMS-COMMON-MIB.mib"
        },
        {
            "mib": "PT-PM-MIB",
            "path": "ericsson/PT-PM-MIB.mib"
        },
        {
            "mib": "CERENT-ENVMON-MIB",
            "path": "cisco/CERENT-ENVMON-MIB.mib"
        },
        {
            "mib": "CIENA-CES-RMON-MIB",
            "path": "ciena/CIENA-CES-RMON-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPV6-MIB",
            "path": "nokia/ALCATEL-IND1-IPV6-MIB.mib"
        },
        {
            "mib": "SP2-MIB",
            "path": "eltek/SP2-MIB.mib"
        },
        {
            "mib": "HUAWEI-E-TRUNK-MIB",
            "path": "huawei/HUAWEI-E-TRUNK-MIB.mib"
        },
        {
            "mib": "NMS-GPON-MIB",
            "path": "bdcom/NMS-GPON-MIB.mib"
        },
        {
            "mib": "HIPATH-WIRELESS-PRODUCTS-MIB",
            "path": "ewc/HIPATH-WIRELESS-PRODUCTS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-SDP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-SDP-MIB.mib"
        },
        {
            "mib": "IMCO-BIG-MIB",
            "path": "imco/IMCO-BIG-MIB.mib"
        },
        {
            "mib": "PT-RADIOLINK-MIB",
            "path": "ericsson/PT-RADIOLINK-MIB.mib"
        },
        {
            "mib": "HP-ICF-TC",
            "path": "hp/HP-ICF-TC.mib"
        },
        {
            "mib": "PAN-TRAPS",
            "path": "paloaltonetworks/PAN-TRAPS.mib"
        },
        {
            "mib": "OAP-C1-OEO",
            "path": "fs/OAP-C1-OEO.mib"
        },
        {
            "mib": "ADVANTECH-EKI-PRONEER-MIB",
            "path": "advantech/ADVANTECH-EKI-PRONEER-MIB.mib"
        },
        {
            "mib": "MIB-Dell-10892",
            "path": "dell/MIB-Dell-10892.mib"
        },
        {
            "mib": "CTRON-SFPS-FLOOD-MIB",
            "path": "enterasys/CTRON-SFPS-FLOOD-MIB.mib"
        },
        {
            "mib": "SanAppliance-MIB",
            "path": "dell/SanAppliance-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-IPX-MIB",
            "path": "nokia/ALCATEL-IND1-IPX-MIB.mib"
        },
        {
            "mib": "AT-FIBER-MONITORING-MIB",
            "path": "awplus/AT-FIBER-MONITORING-MIB.mib"
        },
        {
            "mib": "NetWare-Host-Ext-MIB",
            "path": "novell/NetWare-Host-Ext-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-OPTICALAMPLIFIER-MIB",
            "path": "glassway/NSCRTV-HFCEMS-OPTICALAMPLIFIER-MIB.mib"
        },
        {
            "mib": "RBT-MIB",
            "path": "riverbed/RBT-MIB.mib"
        },
        {
            "mib": "HIPATH-WIRELESS-SMI",
            "path": "ewc/HIPATH-WIRELESS-SMI.mib"
        },
        {
            "mib": "WISI-GTSETTINGS-MIB",
            "path": "wisi/WISI-GTSETTINGS-MIB.mib"
        },
        {
            "mib": "CERENT-FC-MIB",
            "path": "cisco/CERENT-FC-MIB.mib"
        },
        {
            "mib": "PT-SFP-MIB",
            "path": "ericsson/PT-SFP-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-INCLUDE-MIB",
            "path": "enterasys/CTRON-SFPS-INCLUDE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-SERV-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-SERV-MIB.mib"
        },
        {
            "mib": "GANDI-MIB",
            "path": "gandi/GANDI-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LAG-MIB",
            "path": "nokia/ALCATEL-IND1-LAG-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-OPTICALTRANSMITTERDIRECTLY-MIB",
            "path": "glassway/NSCRTV-HFCEMS-OPTICALTRANSMITTERDIRECTLY-MIB.mib"
        },
        {
            "mib": "CIENA-CES-RSVPTE-MIB",
            "path": "ciena/CIENA-CES-RSVPTE-MIB.mib"
        },
        {
            "mib": "NETSCREEN-ADDR-MIB",
            "path": "screenos/NETSCREEN-ADDR-MIB.mib"
        },
        {
            "mib": "NMS-IF-MIB",
            "path": "bdcom/NMS-IF-MIB.mib"
        },
        {
            "mib": "HP-ICF-TRANSCEIVER-MIB",
            "path": "hp/HP-ICF-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "OAP-NMU",
            "path": "fs/OAP-NMU.mib"
        },
        {
            "mib": "StorageManagement-MIB",
            "path": "dell/StorageManagement-MIB.mib"
        },
        {
            "mib": "HUAWEI-ENERGYMNGT-MIB",
            "path": "huawei/HUAWEI-ENERGYMNGT-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-L4SS-MIB",
            "path": "enterasys/CTRON-SFPS-L4SS-MIB.mib"
        },
        {
            "mib": "AT-FILEv2-MIB",
            "path": "awplus/AT-FILEv2-MIB.mib"
        },
        {
            "mib": "WISI-ROOT-MIB",
            "path": "wisi/WISI-ROOT-MIB.mib"
        },
        {
            "mib": "PIM-BSR-MIB",
            "path": "fs/PIM-BSR-MIB.mib"
        },
        {
            "mib": "IMCO-LSPS-MIB",
            "path": "imco/IMCO-LSPS-MIB.mib"
        },
        {
            "mib": "HH3C-3GMODEM-MIB",
            "path": "comware/HH3C-3GMODEM-MIB.mib"
        },
        {
            "mib": "STEELHEAD-MIB",
            "path": "riverbed/STEELHEAD-MIB.mib"
        },
        {
            "mib": "RUGGEDCOM-MIB",
            "path": "ros/RUGGEDCOM-MIB.mib"
        },
        {
            "mib": "PT-TRAP-MIB",
            "path": "ericsson/PT-TRAP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LBD-MIB",
            "path": "nokia/ALCATEL-IND1-LBD-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-TC-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-TC-MIB.mib"
        },
        {
            "mib": "HP-ICF-UDLD-MIB",
            "path": "hp/HP-ICF-UDLD-MIB.mib"
        },
        {
            "mib": "CERENT-GENERIC-MIB",
            "path": "cisco/CERENT-GENERIC-MIB.mib"
        },
        {
            "mib": "STP-MIB",
            "path": "fs/STP-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-MCAST-MIB",
            "path": "enterasys/CTRON-SFPS-MCAST-MIB.mib"
        },
        {
            "mib": "NSCRTV-HFCEMS-PROPERTY-MIB",
            "path": "glassway/NSCRTV-HFCEMS-PROPERTY-MIB.mib"
        },
        {
            "mib": "CIENA-CES-SECURITY-MIB",
            "path": "ciena/CIENA-CES-SECURITY-MIB.mib"
        },
        {
            "mib": "HH3C-8021PAE-MIB",
            "path": "comware/HH3C-8021X-EXT-MIB.mib"
        },
        {
            "mib": "AT-G8032v2-MIB",
            "path": "awplus/AT-G8032v2-MIB.mib"
        },
        {
            "mib": "NIMBLE-MIB",
            "path": "nimble/NIMBLE-MIB.mib"
        },
        {
            "mib": "CERENT-GENERIC-PM-MIB",
            "path": "cisco/CERENT-GENERIC-PM-MIB.mib"
        },
        {
            "mib": "APS-MIB",
            "path": "junos/APS-MIB.mib"
        },
        {
            "mib": "RUGGEDCOM-SYS-INFO-MIB",
            "path": "ros/RUGGEDCOM-SYS-INFO-MIB.mib"
        },
        {
            "mib": "NMS-LLDP-MIB",
            "path": "bdcom/NMS-LLDP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LICENSE-MANAGER-MIB",
            "path": "nokia/ALCATEL-IND1-LICENSE-MANAGER-MIB.mib"
        },
        {
            "mib": "NETSCREEN-BGP4-MIB",
            "path": "screenos/NETSCREEN-BGP4-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-PATH-MIB",
            "path": "enterasys/CTRON-SFPS-PATH-MIB.mib"
        },
        {
            "mib": "HUAWEI-ENTITY-EXTENT-MIB",
            "path": "huawei/HUAWEI-ENTITY-EXTENT-MIB.mib"
        },
        {
            "mib": "RBN-CPU-METER-CAP",
            "path": "ericsson/RBN-CPU-METER-CAP.mib"
        },
        {
            "mib": "NSCRTV-ROOT",
            "path": "glassway/NSCRTV-ROOT.mib"
        },
        {
            "mib": "WISI-TANGRAM-MIB",
            "path": "wisi/WISI-TANGRAM-MIB.mib"
        },
        {
            "mib": "HH3C-8021X-EXT2-MIB",
            "path": "comware/HH3C-8021X-EXT2-MIB.mib"
        },
        {
            "mib": "CIENA-CES-SSH-MIB",
            "path": "ciena/CIENA-CES-SSH-MIB.mib"
        },
        {
            "mib": "HP-ICF-UDP-FORWARD",
            "path": "hp/HP-ICF-UDP-FORWARD.mib"
        },
        {
            "mib": "AT-HHM-MIB",
            "path": "awplus/AT-HHM-MIB.mib"
        },
        {
            "mib": "CERENT-GLOBAL-REGISTRY",
            "path": "cisco/CERENT-GLOBAL-REGISTRY.mib"
        },
        {
            "mib": "NETSCREEN-CERTIFICATE-MIB",
            "path": "screenos/NETSCREEN-CERTIFICATE-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-PKTMGR-MIB",
            "path": "enterasys/CTRON-SFPS-PKTMGR-MIB.mib"
        },
        {
            "mib": "HH3C-AAA-MIB",
            "path": "comware/HH3C-AAA-MIB.mib"
        },
        {
            "mib": "CIENA-CES-SW-XGRADE-MIB",
            "path": "ciena/CIENA-CES-SW-XGRADE-MIB.mib"
        },
        {
            "mib": "NMS-OPTICAL-PORT-MIB",
            "path": "bdcom/NMS-OPTICAL-PORT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LLDP-MED-MIB",
            "path": "nokia/ALCATEL-IND1-LLDP-MED-MIB.mib"
        },
        {
            "mib": "AGENT-GENERAL-MIB",
            "path": "dlink/AGENT-GENERAL-MIB.mib"
        },
        {
            "mib": "companyMIB",
            "path": "orvaldi/companyMIB.mib"
        },
        {
            "mib": "ATM-MIB",
            "path": "junos/ATM-MIB.mib"
        },
        {
            "mib": "NETSCREEN-CHASSIS-MIB",
            "path": "screenos/NETSCREEN-CHASSIS-MIB.mib"
        },
        {
            "mib": "CERENT-HC-RMON-MIB",
            "path": "cisco/CERENT-HC-RMON-MIB.mib"
        },
        {
            "mib": "AT-IP-MIB",
            "path": "awplus/AT-IP-MIB.mib"
        },
        {
            "mib": "RBN-CPU-METER-MIB",
            "path": "ericsson/RBN-CPU-METER-MIB.mib"
        },
        {
            "mib": "HP-ICF-USER-PROFILE-MIB",
            "path": "hp/HP-ICF-USER-PROFILE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TIMETRA-VRTR-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TIMETRA-VRTR-MIB.mib"
        },
        {
            "mib": "NMS-POWER-MIB",
            "path": "bdcom/NMS-POWER-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-POLICY-MIB",
            "path": "enterasys/CTRON-SFPS-POLICY-MIB.mib"
        },
        {
            "mib": "HUAWEI-ENTITY-TRAP-MIB",
            "path": "huawei/HUAWEI-ENTITY-TRAP-MIB.mib"
        },
        {
            "mib": "HH3C-AAA-NASID-MIB",
            "path": "comware/HH3C-AAA-NASID-MIB.mib"
        },
        {
            "mib": "DLINK-ID-REC-MIB",
            "path": "dlink/DLINK-ID-REC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-LPS-MIB",
            "path": "nokia/ALCATEL-IND1-LPS-MIB.mib"
        },
        {
            "mib": "NETBOTZ410-MIB",
            "path": "netbotz/NETBOTZ410-MIB.mib"
        },
        {
            "mib": "CERENT-IF-EXT-MIB",
            "path": "cisco/CERENT-IF-EXT-MIB.mib"
        },
        {
            "mib": "CIENA-CES-SYSLOG-COLLECTOR-MIB",
            "path": "ciena/CIENA-CES-SYSLOG-COLLECTOR-MIB.mib"
        },
        {
            "mib": "NMS-SMI",
            "path": "bdcom/NMS-SMI.mib"
        },
        {
            "mib": "RBN-ENVMON-MIB",
            "path": "ericsson/RBN-ENVMON-MIB.mib"
        },
        {
            "mib": "AT-LICENSE-MIB",
            "path": "awplus/AT-LICENSE-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-PORT-MIB",
            "path": "enterasys/CTRON-SFPS-PORT-MIB.mib"
        },
        {
            "mib": "HUAWEI-ENVIRONMENT-MIB",
            "path": "huawei/HUAWEI-ENVIRONMENT-MIB.mib"
        },
        {
            "mib": "DLINKSW-AAA-ACCOUNTING-MIB",
            "path": "dlink/DLINKSW-AAA-ACCOUNTING-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MAC-ADDRESS-MIB",
            "path": "nokia/ALCATEL-IND1-MAC-ADDRESS-MIB.mib"
        },
        {
            "mib": "ATM-TC-MIB",
            "path": "junos/ATM-TC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TRAP-MGR-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-TRAP-MGR-MIB.mib"
        },
        {
            "mib": "RUGGEDCOM-TRAPS-MIB",
            "path": "ros/RUGGEDCOM-TRAPS-MIB.mib"
        },
        {
            "mib": "BDCOM-MEMORY-POOL-MIB",
            "path": "fs/bdcom/BDCOM-MEMORY-POOL-MIB.mib"
        },
        {
            "mib": "CIENA-CES-SYSTEM-CONFIG-MIB",
            "path": "ciena/CIENA-CES-SYSTEM-CONFIG-MIB.mib"
        },
        {
            "mib": "NETSCREEN-IDS-MIB",
            "path": "screenos/NETSCREEN-IDS-MIB.mib"
        },
        {
            "mib": "RBN-SMI",
            "path": "ericsson/RBN-SMI.mib"
        },
        {
            "mib": "TN-LAG-MIB",
            "path": "nokia/1830/TN-LAG-MIB.mib"
        },
        {
            "mib": "AT-LINKTRAP-MIB",
            "path": "awplus/AT-LINKTRAP-MIB.mib"
        },
        {
            "mib": "HH3C-ACFP-MIB",
            "path": "comware/HH3C-ACFP-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-RESOLVE-MIB",
            "path": "enterasys/CTRON-SFPS-RESOLVE-MIB.mib"
        },
        {
            "mib": "BFD-STD-MIB",
            "path": "junos/BFD-STD-MIB.mib"
        },
        {
            "mib": "HUAWEI-EPON-MIB",
            "path": "huawei/HUAWEI-EPON-MIB.mib"
        },
        {
            "mib": "DLINKSW-AAA-AUTH-MIB",
            "path": "dlink/DLINKSW-AAA-AUTH-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MAC-SERVER-MIB",
            "path": "nokia/ALCATEL-IND1-MAC-SERVER-MIB.mib"
        },
        {
            "mib": "SIAE-ACM-STATISTICS-MIB",
            "path": "siae/SIAE-ACM-STATISTICS-MIB.mib"
        },
        {
            "mib": "HP-ICF-VG-RPTR",
            "path": "hp/HP-ICF-VG-RPTR.mib"
        },
        {
            "mib": "BDCOM-PROCESS-MIB",
            "path": "fs/bdcom/BDCOM-PROCESS-MIB.mib"
        },
        {
            "mib": "AT-LOG-MIB",
            "path": "awplus/AT-LOG-MIB.mib"
        },
        {
            "mib": "NetBotz50-MIB",
            "path": "netbotz/NetBotz50-MIB.mib"
        },
        {
            "mib": "RBN-TC",
            "path": "ericsson/RBN-TC.mib"
        },
        {
            "mib": "CERENT-MSDWDM-MIB",
            "path": "cisco/CERENT-MSDWDM-MIB.mib"
        },
        {
            "mib": "NETSCREEN-INTERFACE-MIB",
            "path": "screenos/NETSCREEN-INTERFACE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-UDLD-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-UDLD-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-SFLSP-MIB",
            "path": "enterasys/CTRON-SFPS-SFLSP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MLD-MIB",
            "path": "nokia/ALCATEL-IND1-MLD-MIB.mib"
        },
        {
            "mib": "DLINKSW-AAA-COMMON-MIB",
            "path": "dlink/DLINKSW-AAA-COMMON-MIB.mib"
        },
        {
            "mib": "MIKROTIK-MIB",
            "path": "mikrotik/MIKROTIK-MIB.mib"
        },
        {
            "mib": "HH3C-ACL-MIB",
            "path": "comware/HH3C-ACL-MIB.mib"
        },
        {
            "mib": "HP-ICF-VRRP-MIB",
            "path": "hp/HP-ICF-VRRP-MIB.mib"
        },
        {
            "mib": "AT-LOOPPROTECT-MIB",
            "path": "awplus/AT-LOOPPROTECT-MIB.mib"
        },
        {
            "mib": "CERENT-OPTICAL-MONITOR-MIB",
            "path": "cisco/CERENT-OPTICAL-MONITOR-MIB.mib"
        },
        {
            "mib": "HUAWEI-ERPS-MIB",
            "path": "huawei/HUAWEI-ERPS-MIB.mib"
        },
        {
            "mib": "XF-RADIOLINK-PTP-MODEM-MIB",
            "path": "ericsson/XF-RADIOLINK-PTP-MODEM-MIB.mib"
        },
        {
            "mib": "CIENA-CES-TACACS-CLIENT-MIB",
            "path": "ciena/CIENA-CES-TACACS-CLIENT-MIB.mib"
        },
        {
            "mib": "NETSCREEN-IP-ARP-MIB",
            "path": "screenos/NETSCREEN-IP-ARP-MIB.mib"
        },
        {
            "mib": "BDCOM-QOS-PIB-MIB",
            "path": "fs/bdcom/BDCOM-QOS-PIB-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-MVRP-MIB",
            "path": "nokia/ALCATEL-IND1-MVRP-MIB.mib"
        },
        {
            "mib": "HH3C-AFC-MIB",
            "path": "comware/HH3C-AFC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-UDP-RELAY-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-UDP-RELAY-MIB.mib"
        },
        {
            "mib": "HP-ICF-XRRP",
            "path": "hp/HP-ICF-XRRP.mib"
        },
        {
            "mib": "SWITCH",
            "path": "fs/SWITCH.mib"
        },
        {
            "mib": "OPENBSD-BASE-MIB",
            "path": "openbsd/OPENBSD-BASE-MIB.mib"
        },
        {
            "mib": "TN-OAM-TEST-MIB",
            "path": "nokia/1830/TN-OAM-TEST-MIB.mib"
        },
        {
            "mib": "DLINKSW-AAA-SERVER-MIB",
            "path": "dlink/DLINKSW-AAA-SERVER-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-SIZE-MIB",
            "path": "enterasys/CTRON-SFPS-SIZE-MIB.mib"
        },
        {
            "mib": "NETBOTZV2-MIB",
            "path": "netbotz/NETBOTZV2-MIB.mib"
        },
        {
            "mib": "BGP4-V2-MIB-JUNIPER",
            "path": "junos/BGP4-V2-MIB-JUNIPER.mib"
        },
        {
            "mib": "CIENA-CES-TIME-SYNC-MIB",
            "path": "ciena/CIENA-CES-TIME-SYNC-MIB.mib"
        },
        {
            "mib": "SIAE-AGGRL1-MANAGEMENT-MIB",
            "path": "siae/SIAE-AGGRL1-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "CERENT-TC",
            "path": "cisco/CERENT-TC.mib"
        },
        {
            "mib": "AT-MIBVERSION-MIB",
            "path": "awplus/AT-MIBVERSION-MIB.mib"
        },
        {
            "mib": "XF-RADIOLINK-PTP-RADIO-MIB",
            "path": "ericsson/XF-RADIOLINK-PTP-RADIO-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-NETSEC-MIB",
            "path": "nokia/ALCATEL-IND1-NETSEC-MIB.mib"
        },
        {
            "mib": "HP-IF-EXT-MIB",
            "path": "hp/HP-IF-EXT-MIB.mib"
        },
        {
            "mib": "OPENBSD-CARP-MIB",
            "path": "openbsd/OPENBSD-CARP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VIRTUAL-CHASSIS-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VIRTUAL-CHASSIS-MIB.mib"
        },
        {
            "mib": "BDCOM-SMI",
            "path": "fs/bdcom/BDCOM-SMI.mib"
        },
        {
            "mib": "DLINKSW-ACL-MIB",
            "path": "dlink/DLINKSW-ACL-MIB.mib"
        },
        {
            "mib": "CISCO-AAA-SERVER-MIB",
            "path": "cisco/CISCO-AAA-SERVER-MIB.mib"
        },
        {
            "mib": "PowerNet-MIB",
            "path": "apc/PowerNet-MIB.mib"
        },
        {
            "mib": "NETSCREEN-IPPOOL-MIB",
            "path": "screenos/NETSCREEN-IPPOOL-MIB.mib"
        },
        {
            "mib": "CIENA-CES-TWAMP-MIB",
            "path": "ciena/CIENA-CES-TWAMP-MIB.mib"
        },
        {
            "mib": "SIAE-ALARM-MIB",
            "path": "siae/SIAE-ALARM-MIB.mib"
        },
        {
            "mib": "XPPC-MIB",
            "path": "logmaster/UPSMATE-MIB.mib"
        },
        {
            "mib": "TN-PMON-MIB",
            "path": "nokia/1830/TN-PMON-MIB.mib"
        },
        {
            "mib": "HH3C-ARP-ENTRY-MIB",
            "path": "comware/HH3C-ARP-ENTRY-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-SOFTLINK-MIB",
            "path": "enterasys/CTRON-SFPS-SOFTLINK-MIB.mib"
        },
        {
            "mib": "HUAWEI-ERRORDOWN-MIB",
            "path": "huawei/HUAWEI-ERRORDOWN-MIB.mib"
        },
        {
            "mib": "XF-RADIOLINK-PTP-TERMINAL-MIB",
            "path": "ericsson/XF-RADIOLINK-PTP-TERMINAL-MIB.mib"
        },
        {
            "mib": "BDCOM-TC",
            "path": "fs/bdcom/BDCOM-TC.mib"
        },
        {
            "mib": "DLINKSW-ASP-MIB",
            "path": "dlink/DLINKSW-ASP-MIB.mib"
        },
        {
            "mib": "NETSCREEN-NAT-MIB",
            "path": "screenos/NETSCREEN-NAT-MIB.mib"
        },
        {
            "mib": "AT-NTP-MIB",
            "path": "awplus/AT-NTP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-NTP-MIB",
            "path": "nokia/ALCATEL-IND1-NTP-MIB.mib"
        },
        {
            "mib": "OPENBSD-MEM-MIB",
            "path": "openbsd/OPENBSD-MEM-MIB.mib"
        },
        {
            "mib": "DRAFT-MSDP-MIB",
            "path": "junos/DRAFT-MSDP-MIB.mib"
        },
        {
            "mib": "HUAWEI-ETHARP-MIB",
            "path": "huawei/HUAWEI-ETHARP-MIB.mib"
        },
        {
            "mib": "CISCO-AAA-SESSION-MIB",
            "path": "cisco/CISCO-AAA-SESSION-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VIRTUAL-FLOW-CONTROL-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VIRTUAL-FLOW-CONTROL-MIB.mib"
        },
        {
            "mib": "SIAE-CARRIER-AGGRL1-MIB",
            "path": "siae/SIAE-CARRIER-AGGRL1-MIB.mib"
        },
        {
            "mib": "HH3C-ARP-RATELIMIT-MIB",
            "path": "comware/HH3C-ARP-RATELIMIT-MIB.mib"
        },
        {
            "mib": "TNMS-NBI-MIB",
            "path": "coriant/TNMS-NBI-MIB.mib"
        },
        {
            "mib": "DASAN-ACCESS-MIB",
            "path": "dasan/DASAN-ACCESS-MIB.mib"
        },
        {
            "mib": "NMS-CARD-SYS-MIB",
            "path": "fs/bdcom/NMS-CARD-SYS-MIB.mib"
        },
        {
            "mib": "CIENA-CES-VLLI-MIB",
            "path": "ciena/CIENA-CES-VLLI-MIB.mib"
        },
        {
            "mib": "HP-LASERJET-COMMON-MIB",
            "path": "hp/HP-LASERJET-COMMON-MIB.mib"
        },
        {
            "mib": "DLINKSW-BPDU-PROTECTION-MIB",
            "path": "dlink/DLINKSW-BPDU-PROTECTION-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-OSPF-MIB",
            "path": "nokia/ALCATEL-IND1-OSPF-MIB.mib"
        },
        {
            "mib": "CONTROLLER-MIB",
            "path": "ucopia/CONTROLLER-MIB.mib"
        },
        {
            "mib": "TN-PORT-MIB",
            "path": "nokia/1830/TN-PORT-MIB.mib"
        },
        {
            "mib": "HH3C-ARP-SOURCE-SUPPRESSION-MIB",
            "path": "comware/HH3C-ARP-SOURCE-SUPPRESSION-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-TAP-MIB",
            "path": "enterasys/CTRON-SFPS-TAP-MIB.mib"
        },
        {
            "mib": "OPENBSD-PF-MIB",
            "path": "openbsd/OPENBSD-PF-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VIRTUALROUTER-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VIRTUALROUTER-MIB.mib"
        },
        {
            "mib": "AT-PLUGGABLE-DIAGNOSTICS-MIB",
            "path": "awplus/AT-PLUGGABLE-DIAGNOSTICS-MIB.mib"
        },
        {
            "mib": "XF-RADIOLINK-RLT-MIB",
            "path": "ericsson/XF-RADIOLINK-RLT-MIB.mib"
        },
        {
            "mib": "CIENA-GLOBAL-MIB",
            "path": "ciena/CIENA-GLOBAL-MIB.mib"
        },
        {
            "mib": "SIAE-CFGM-MIB",
            "path": "siae/SIAE-CFGM-MIB.mib"
        },
        {
            "mib": "HP-MEMPROC-MIB",
            "path": "hp/HP-MEMPROC-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VLAN-MGR-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VLAN-MGR-MIB.mib"
        },
        {
            "mib": "NETSCREEN-NSRP-MIB",
            "path": "screenos/NETSCREEN-NSRP-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-TOPOLOGY-MIB",
            "path": "enterasys/CTRON-SFPS-TOPOLOGY-MIB.mib"
        },
        {
            "mib": "DS1-MIB",
            "path": "junos/DS1-MIB.mib"
        },
        {
            "mib": "CIENA-PRO-SOFTWARE-MIB",
            "path": "ciena/CIENA-PRO-SOFTWARE-MIB.mib"
        },
        {
            "mib": "DASAN-ACCESS-SLOT-H248-MIB",
            "path": "dasan/DASAN-ACCESS-SLOT-H248-MIB.mib"
        },
        {
            "mib": "CISCO-AUTH-FRAMEWORK-MIB",
            "path": "cisco/CISCO-AUTH-FRAMEWORK-MIB.mib"
        },
        {
            "mib": "NETSCREEN-OSPF-MIB",
            "path": "screenos/NETSCREEN-OSPF-MIB.mib"
        },
        {
            "mib": "NMS-CHASSIS",
            "path": "fs/bdcom/NMS-CHASSIS.mib"
        },
        {
            "mib": "TN-SERV-MIB",
            "path": "nokia/1830/TN-SERV-MIB.mib"
        },
        {
            "mib": "DLINKSW-CABLE-DIAG-MIB",
            "path": "dlink/DLINKSW-CABLE-DIAG-MIB.mib"
        },
        {
            "mib": "SIAE-CLOG-MIB",
            "path": "siae/SIAE-CLOG-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-OSPF3-MIB",
            "path": "nokia/ALCATEL-IND1-OSPF3-MIB.mib"
        },
        {
            "mib": "HP-PROCURVE-420-PRIVATE-MIB",
            "path": "hp/HP-PROCURVE-420-PRIVATE-MIB.mib"
        },
        {
            "mib": "OPENBSD-SENSORS-MIB",
            "path": "openbsd/OPENBSD-SENSORS-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-VLAN-MIB",
            "path": "enterasys/CTRON-SFPS-VLAN-MIB.mib"
        },
        {
            "mib": "XF-SOFTWARE-MIB",
            "path": "ericsson/XF-SOFTWARE-MIB.mib"
        },
        {
            "mib": "HH3C-ATM-DXI-MIB",
            "path": "comware/HH3C-ATM-DXI-MIB.mib"
        },
        {
            "mib": "CTS-cvt_wac_wpc_mac_mpc3112-MIB",
            "path": "cts/CTS-cvt_wac_wpc_mac_mpc3112-MIB.mib"
        },
        {
            "mib": "SIAE-ECFM-EXT-MIB",
            "path": "siae/SIAE-ECFM-EXT-MIB.mib"
        },
        {
            "mib": "AT-PRODUCT-MIB",
            "path": "awplus/AT-PRODUCT-MIB.mib"
        },
        {
            "mib": "HUAWEI-ETHOAM-MIB",
            "path": "huawei/HUAWEI-ETHOAM-MIB.mib"
        },
        {
            "mib": "CIENA-PRO-TYPES-MIB",
            "path": "ciena/CIENA-PRO-TYPES-MIB.mib"
        },
        {
            "mib": "NETSCREEN-OSPF-TRAP-MIB",
            "path": "screenos/NETSCREEN-OSPF-TRAP-MIB.mib"
        },
        {
            "mib": "DLINKSW-CPU-PROTECT-MIB",
            "path": "dlink/DLINKSW-CPU-PROTECT-MIB.mib"
        },
        {
            "mib": "CTRON-SFPS-VSTP-MIB",
            "path": "enterasys/CTRON-SFPS-VSTP-MIB.mib"
        },
        {
            "mib": "DASAN-ACCESS-SLOT-MGCP-MIB",
            "path": "dasan/DASAN-ACCESS-SLOT-MGCP-MIB.mib"
        },
        {
            "mib": "XF-TOP-MIB",
            "path": "ericsson/XF-TOP-MIB.mib"
        },
        {
            "mib": "HH3C-BFD-STD-MIB",
            "path": "comware/HH3C-BFD-STD-MIB.mib"
        },
        {
            "mib": "HES-3112-MIB",
            "path": "cts/HES-3112-MIB.mib"
        },
        {
            "mib": "OPENBSD-SNMPD-CONF",
            "path": "openbsd/OPENBSD-SNMPD-CONF.mib"
        },
        {
            "mib": "ALCATEL-IND1-PARTITIONED-MGR-MIB",
            "path": "nokia/ALCATEL-IND1-PARTITIONED-MGR-MIB.mib"
        },
        {
            "mib": "CISCO-BGP4-MIB",
            "path": "cisco/CISCO-BGP4-MIB.mib"
        },
        {
            "mib": "NMS-EPON-OLT-PON",
            "path": "fs/bdcom/NMS-EPON-OLT-PON.mib"
        },
        {
            "mib": "CIENA-PRODUCTS-MIB",
            "path": "ciena/CIENA-PRODUCTS-MIB.mib"
        },
        {
            "mib": "AT-PTP-MIB",
            "path": "awplus/AT-PTP-MIB.mib"
        },
        {
            "mib": "ESO-CONSORTIUM-MIB",
            "path": "junos/ESO-CONSORTIUM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VLAN-STP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VLAN-STP-MIB.mib"
        },
        {
            "mib": "SIAE-EQUIP-MIB",
            "path": "siae/SIAE-EQUIP-MIB.mib"
        },
        {
            "mib": "HP-SN-AGENT-MIB",
            "path": "hp/HP-SN-AGENT-MIB.mib"
        },
        {
            "mib": "CTRON-SMARTTRUNK-MIB",
            "path": "enterasys/CTRON-SMARTTRUNK-MIB.mib"
        },
        {
            "mib": "HH3C-BGP-EVPN-MIB",
            "path": "comware/HH3C-BGP-EVPN-MIB.mib"
        },
        {
            "mib": "HUAWEI-EVC-MIB",
            "path": "huawei/HUAWEI-EVC-MIB.mib"
        },
        {
            "mib": "CISCO-BRIDGE-DOMAIN-MIB",
            "path": "cisco/CISCO-BRIDGE-DOMAIN-MIB.mib"
        },
        {
            "mib": "AT-QOSv2-MIB",
            "path": "awplus/AT-QOSv2-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-PIM-MIB",
            "path": "nokia/ALCATEL-IND1-PIM-MIB.mib"
        },
        {
            "mib": "NETSCREEN-POLICY-MIB",
            "path": "screenos/NETSCREEN-POLICY-MIB.mib"
        },
        {
            "mib": "DLINKSW-DAI-MIB",
            "path": "dlink/DLINKSW-DAI-MIB.mib"
        },
        {
            "mib": "NMS-FAN-TRAP",
            "path": "fs/bdcom/NMS-FAN-TRAP.mib"
        },
        {
            "mib": "SIAE-EQUIPTYPE-MIB",
            "path": "siae/SIAE-EQUIPTYPE-MIB.mib"
        },
        {
            "mib": "HP-SN-APPLETALK-MIB",
            "path": "hp/HP-SN-APPLETALK-MIB.mib"
        },
        {
            "mib": "CTMMIBCUSTOM",
            "path": "ctm/CTMMIBCUSTOM.mib"
        },
        {
            "mib": "CIENA-SMI",
            "path": "ciena/CIENA-SMI.mib"
        },
        {
            "mib": "CISCO-BRIDGE-EXT-MIB",
            "path": "cisco/CISCO-BRIDGE-EXT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VRRP-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VRRP-MIB.mib"
        },
        {
            "mib": "AT-RESOURCE-MIB",
            "path": "awplus/AT-RESOURCE-MIB.mib"
        },
        {
            "mib": "TN-TC-MIB",
            "path": "nokia/1830/TN-TC-MIB.mib"
        },
        {
            "mib": "DLINKSW-DDM-MIB",
            "path": "dlink/DLINKSW-DDM-MIB.mib"
        },
        {
            "mib": "Juniper-IP-POLICY-MIB",
            "path": "juniper/Juniper-IP-POLICY-MIB.mib"
        },
        {
            "mib": "IES-3110-MIB",
            "path": "cts/IES-3110-MIB.mib"
        },
        {
            "mib": "ETHER-WIS",
            "path": "junos/ETHER-WIS.mib"
        },
        {
            "mib": "CIENA-TC",
            "path": "ciena/CIENA-TC.mib"
        },
        {
            "mib": "SIAE-FEATUREKEYS-MIB",
            "path": "siae/SIAE-FEATUREKEYS-MIB.mib"
        },
        {
            "mib": "HP-SN-BGP4-GROUP-MIB",
            "path": "hp/HP-SN-BGP4-GROUP-MIB.mib"
        },
        {
            "mib": "CTMMIB",
            "path": "ctm/CTMMIBV2.mib"
        },
        {
            "mib": "CTRON-SSR-CAPACITY-MIB",
            "path": "enterasys/CTRON-SSR-CAPACITY-MIB.mib"
        },
        {
            "mib": "HH3C-BGP-VPN-MIB",
            "path": "comware/HH3C-BGP-VPN-MIB.mib"
        },
        {
            "mib": "DASAN-ACCESS-SLOT-POTS-MIB",
            "path": "dasan/DASAN-ACCESS-SLOT-POTS-MIB.mib"
        },
        {
            "mib": "TROPIC-ACCESSPORT-MIB",
            "path": "nokia/1830/TROPIC-ACCESSPORT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VRRP3-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-VRRP3-MIB.mib"
        },
        {
            "mib": "AT-SETUP-MIB",
            "path": "awplus/AT-SETUP-MIB.mib"
        },
        {
            "mib": "DLINKSW-DDP-CLIENT-MIB",
            "path": "dlink/DLINKSW-DDP-CLIENT-MIB.mib"
        },
        {
            "mib": "HUAWEI-EVPN-MIB",
            "path": "huawei/HUAWEI-EVPN-MIB.mib"
        },
        {
            "mib": "NETSCREEN-PRODUCTS-MIB",
            "path": "screenos/NETSCREEN-PRODUCTS-MIB.mib"
        },
        {
            "mib": "SIAE-HC-MIB",
            "path": "siae/SIAE-HC-MIB.mib"
        },
        {
            "mib": "CISCO-CASA-MIB",
            "path": "cisco/CISCO-CASA-MIB.mib"
        },
        {
            "mib": "CIENA-WS-ALARM-MIB",
            "path": "ciena/CIENA-WS-ALARM-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-POLICY-MIB",
            "path": "nokia/ALCATEL-IND1-POLICY-MIB.mib"
        },
        {
            "mib": "FR-MFR-MIB",
            "path": "junos/FR-MFR-MIB.mib"
        },
        {
            "mib": "NMS-GPON-MIB",
            "path": "fs/bdcom/NMS-GPON-MIB.mib"
        },
        {
            "mib": "JUNIPER-MIB",
            "path": "juniper/JUNIPER-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-WEBMGT-MIB",
            "path": "nokia/aos7/ALCATEL-IND1-WEBMGT-MIB.mib"
        },
        {
            "mib": "CISCO-CAT6K-CROSSBAR-MIB",
            "path": "cisco/CISCO-CAT6K-CROSSBAR-MIB.mib"
        },
        {
            "mib": "NETSCREEN-QOS-MIB",
            "path": "screenos/NETSCREEN-QOS-MIB.mib"
        },
        {
            "mib": "TROPIC-ALARMPANEL-MIB",
            "path": "nokia/1830/TROPIC-ALARMPANEL-MIB.mib"
        },
        {
            "mib": "ADR155C-MIB",
            "path": "sagemcom/ADR155C-MIB.mib"
        },
        {
            "mib": "CIENA-WS-BLADE-MIB",
            "path": "ciena/CIENA-WS-BLADE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-PORT-MAPPING",
            "path": "nokia/ALCATEL-IND1-PORT-MAPPING.mib"
        },
        {
            "mib": "CTRON-SSR-CONFIG-MIB",
            "path": "enterasys/CTRON-SSR-CONFIG-MIB.mib"
        },
        {
            "mib": "DASAN-ACCESS-SLOT-SIP-MIB",
            "path": "dasan/DASAN-ACCESS-SLOT-SIP-MIB.mib"
        },
        {
            "mib": "IPS-3106-SE-PB-MIB",
            "path": "cts/IPS-3106-SE-PB-MIB.mib"
        },
        {
            "mib": "HH3C-BGP4V2-MIB",
            "path": "comware/HH3C-BGP4V2-MIB.mib"
        },
        {
            "mib": "NMS-IF-MIB",
            "path": "fs/bdcom/NMS-IF-MIB.mib"
        },
        {
            "mib": "Juniper-MIBs",
            "path": "juniper/Juniper-MIBs.mib"
        },
        {
            "mib": "TEGILE-MIB",
            "path": "tegile/TEGILE-MIB.mib"
        },
        {
            "mib": "DLINKSW-DHCP-FILTER-MIB",
            "path": "dlink/DLINKSW-DHCP-FILTER-MIB.mib"
        },
        {
            "mib": "CIENA-WS-CHASSIS-MIB",
            "path": "ciena/CIENA-WS-CHASSIS-MIB.mib"
        },
        {
            "mib": "AT-SMI-MIB",
            "path": "awplus/AT-SMI-MIB.mib"
        },
        {
            "mib": "TROPIC-AMPLIFIER-MIB",
            "path": "nokia/1830/TROPIC-AMPLIFIER-MIB.mib"
        },
        {
            "mib": "NETSCREEN-RESOURCE-MIB",
            "path": "screenos/NETSCREEN-RESOURCE-MIB.mib"
        },
        {
            "mib": "CISCO-CDP-MIB",
            "path": "cisco/CISCO-CDP-MIB.mib"
        },
        {
            "mib": "HH3C-BLG-MIB",
            "path": "comware/HH3C-BLG-MIB.mib"
        },
        {
            "mib": "ADR2500C-MIB",
            "path": "sagemcom/ADR2500C-MIB.mib"
        },
        {
            "mib": "HUAWEI-FCOE-MIB",
            "path": "huawei/HUAWEI-FCOE-MIB.mib"
        },
        {
            "mib": "NMS-LLDP-MIB",
            "path": "fs/bdcom/NMS-LLDP-MIB.mib"
        },
        {
            "mib": "SIAE-HITLESS-AGGRL1-MIB",
            "path": "siae/SIAE-HITLESS-AGGRL1-MIB.mib"
        },
        {
            "mib": "HP-SN-IGMP-MIB",
            "path": "hp/HP-SN-IGMP-MIB.mib"
        },
        {
            "mib": "CIENA-WS-CONFIGURATION-MIB",
            "path": "ciena/CIENA-WS-CONFIGURATION-MIB.mib"
        },
        {
            "mib": "IPS-3110-MIB",
            "path": "cts/IPS-3110-MIB.mib"
        },
        {
            "mib": "NETSCREEN-RIPv2-MIB",
            "path": "screenos/NETSCREEN-RIPv2-MIB.mib"
        },
        {
            "mib": "DASAN-ADSL-MIB",
            "path": "dasan/DASAN-ADSL-MIB.mib"
        },
        {
            "mib": "DLINKSW-DHCP-RELAY-MIB",
            "path": "dlink/DLINKSW-DHCP-RELAY-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-PORT-MIB",
            "path": "nokia/ALCATEL-IND1-PORT-MIB.mib"
        },
        {
            "mib": "GGSN-MIB",
            "path": "junos/GGSN-MIB.mib"
        },
        {
            "mib": "HH3C-BPA-MIB",
            "path": "comware/HH3C-BPA-MIB.mib"
        },
        {
            "mib": "SIAE-IFEXT-MIB",
            "path": "siae/SIAE-IFEXT-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-HARDWARE-MIB",
            "path": "enterasys/CTRON-SSR-HARDWARE-MIB.mib"
        },
        {
            "mib": "NMS-OPTICAL-PORT-MIB",
            "path": "fs/bdcom/NMS-OPTICAL-PORT-MIB.mib"
        },
        {
            "mib": "ADR63E1-MIB",
            "path": "sagemcom/ADR63E1-MIB.mib"
        },
        {
            "mib": "IPS-3110-PB-MIB",
            "path": "cts/IPS-3110-PB-MIB.mib"
        },
        {
            "mib": "TROPIC-CARD-MIB",
            "path": "nokia/1830/TROPIC-CARD-MIB.mib"
        },
        {
            "mib": "CIENA-WS-ENCRYPTION-MIB",
            "path": "ciena/CIENA-WS-ENCRYPTION-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-PORT-MIRRORING-MONITORING-MIB",
            "path": "nokia/ALCATEL-IND1-PORT-MIRRORING-MONITORING-MIB.mib"
        },
        {
            "mib": "Juniper-ROUTER-MIB",
            "path": "juniper/Juniper-ROUTER-MIB.mib"
        },
        {
            "mib": "HH3C-BRAS-ACCESS-MIB",
            "path": "comware/HH3C-BRAS-ACCESS-MIB.mib"
        },
        {
            "mib": "HP-SN-IP-ACL-MIB",
            "path": "hp/HP-SN-IP-ACL-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SCHEDULE-MIB",
            "path": "screenos/NETSCREEN-SCHEDULE-MIB.mib"
        },
        {
            "mib": "AT-SWITCH-MIB",
            "path": "awplus/AT-SWITCH-MIB.mib"
        },
        {
            "mib": "CISCO-CEF-MIB",
            "path": "cisco/CISCO-CEF-MIB.mib"
        },
        {
            "mib": "GMPLS-LSR-STD-MIB",
            "path": "junos/GMPLS-LSR-STD-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-L2-MIB",
            "path": "enterasys/CTRON-SSR-L2-MIB.mib"
        },
        {
            "mib": "HUAWEI-FLASH-MAN-MIB",
            "path": "huawei/HUAWEI-FLASH-MAN-MIB.mib"
        },
        {
            "mib": "APS-MIB-JUNI",
            "path": "junose/APS-MIB-JUNI.mib"
        },
        {
            "mib": "HH3C-CATV-TRANSCEIVER-MIB",
            "path": "comware/HH3C-CATV-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "SIAE-LLF-MIB",
            "path": "siae/SIAE-LLF-MIB.mib"
        },
        {
            "mib": "EQUIPMENT-MIB",
            "path": "sagemcom/EQUIPMENT-MIB.mib"
        },
        {
            "mib": "DASAN-AUTORESET-MIB",
            "path": "dasan/DASAN-AUTORESET-MIB.mib"
        },
        {
            "mib": "NMS-POWER-MIB",
            "path": "fs/bdcom/NMS-POWER-MIB.mib"
        },
        {
            "mib": "AT-SYSINFO-MIB",
            "path": "awplus/AT-SYSINFO-MIB.mib"
        },
        {
            "mib": "HUAWEI-FR-QOS-MIB",
            "path": "huawei/HUAWEI-FR-QOS-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-L3-MIB",
            "path": "enterasys/CTRON-SSR-L3-MIB.mib"
        },
        {
            "mib": "DLINKSW-DHCP-SERVER-MIB",
            "path": "dlink/DLINKSW-DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-QOS-MIB",
            "path": "nokia/ALCATEL-IND1-QOS-MIB.mib"
        },
        {
            "mib": "SIAE-MAB-MIB",
            "path": "siae/SIAE-MAB-MIB.mib"
        },
        {
            "mib": "CIENA-WS-ENCRYPTION-RPC-MIB",
            "path": "ciena/CIENA-WS-ENCRYPTION-RPC-MIB.mib"
        },
        {
            "mib": "PowerNet-MIB",
            "path": "apc-cpdu/PowerNet-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SERVICE-MIB",
            "path": "screenos/NETSCREEN-SERVICE-MIB.mib"
        },
        {
            "mib": "TROPIC-CONTROLCARD-MIB",
            "path": "nokia/1830/TROPIC-CONTROLCARD-MIB.mib"
        },
        {
            "mib": "CISCO-CEF-TC",
            "path": "cisco/CISCO-CEF-TC.mib"
        },
        {
            "mib": "GMPLS-TC-STD-MIB",
            "path": "junos/GMPLS-TC-STD-MIB.mib"
        },
        {
            "mib": "JUNIPER-SMI",
            "path": "juniper/JUNIPER-SMI.mib"
        },
        {
            "mib": "HP-SN-IP-MIB",
            "path": "hp/HP-SN-IP-MIB.mib"
        },
        {
            "mib": "NMS-SMI",
            "path": "fs/bdcom/NMS-SMI.mib"
        },
        {
            "mib": "TROPIC-FAN-MIB",
            "path": "nokia/1830/TROPIC-FAN-MIB.mib"
        },
        {
            "mib": "APS-MIB",
            "path": "junose/APS-MIB.mib"
        },
        {
            "mib": "GIGE-MIB",
            "path": "sagemcom/GIGE-MIB.mib"
        },
        {
            "mib": "DASAN-BRIDGE-MIB",
            "path": "dasan/DASAN-BRIDGE-MIB.mib"
        },
        {
            "mib": "GMPLS-TE-STD-MIB",
            "path": "junos/GMPLS-TE-STD-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-POLICY-MIB",
            "path": "enterasys/CTRON-SSR-POLICY-MIB.mib"
        },
        {
            "mib": "AT-TRIGGER-MIB",
            "path": "awplus/AT-TRIGGER-MIB.mib"
        },
        {
            "mib": "HH3C-CBQOS2-MIB",
            "path": "comware/HH3C-CBQOS2-MIB.mib"
        },
        {
            "mib": "HUAWEI-FTP-MIB",
            "path": "huawei/HUAWEI-FTP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-RDP-MIB",
            "path": "nokia/ALCATEL-IND1-RDP-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-ADMIN-USR-MIB",
            "path": "screenos/NETSCREEN-SET-ADMIN-USR-MIB.mib"
        },
        {
            "mib": "CIENA-WS-LICENSE-MIB",
            "path": "ciena/CIENA-WS-LICENSE-MIB.mib"
        },
        {
            "mib": "DLINKSW-DHCP-SNOOPING-MIB",
            "path": "dlink/DLINKSW-DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "SIAE-MANOP-MIB",
            "path": "siae/SIAE-MANOP-MIB.mib"
        },
        {
            "mib": "HAWK-I2-MIB",
            "path": "sinetica/HAWK-I2-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-SERVICE-STATUS-MIB",
            "path": "enterasys/CTRON-SSR-SERVICE-STATUS-MIB.mib"
        },
        {
            "mib": "SAMLEXAMERICA-MIB",
            "path": "samlex/SAMLEXAMERICA-MIB.mib"
        },
        {
            "mib": "IANA-GMPLS-TC-MIB",
            "path": "junos/IANA-GMPLS-TC-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-AUTH-MIB",
            "path": "screenos/NETSCREEN-SET-AUTH-MIB.mib"
        },
        {
            "mib": "HP-SN-IP-VRRP-MIB",
            "path": "hp/HP-SN-IP-VRRP-MIB.mib"
        },
        {
            "mib": "HH3C-CFCARD-MIB",
            "path": "comware/HH3C-CFCARD-MIB.mib"
        },
        {
            "mib": "LOG-MIB",
            "path": "sagemcom/LOG-MIB.mib"
        },
        {
            "mib": "AT-UDLD-MIB",
            "path": "awplus/AT-UDLD-MIB.mib"
        },
        {
            "mib": "CISCO-CHANNEL-MIB",
            "path": "cisco/CISCO-CHANNEL-MIB.mib"
        },
        {
            "mib": "Juniper-TC",
            "path": "juniper/Juniper-TC.mib"
        },
        {
            "mib": "ATM-MIB",
            "path": "junose/ATM-MIB.mib"
        },
        {
            "mib": "DASAN-DHCP-MIB",
            "path": "dasan/DASAN-DHCP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-RIP-MIB",
            "path": "nokia/ALCATEL-IND1-RIP-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-SMI-MIB",
            "path": "enterasys/CTRON-SSR-SMI-MIB.mib"
        },
        {
            "mib": "CIENA-WS-MIB",
            "path": "ciena/CIENA-WS-MIB.mib"
        },
        {
            "mib": "QUANTA-LB6M-REF-MIB",
            "path": "quanta/cheetahref.my.mib"
        },
        {
            "mib": "HH3C-COMMON-SYSTEM-MIB",
            "path": "comware/HH3C-COMMON-SYSTEM-MIB.mib"
        },
        {
            "mib": "IPMCAST-MIB-CAPABILITY",
            "path": "junos/IPMCAST-MIB-CAPABILITY.mib"
        },
        {
            "mib": "SIAE-PMFTP-MIB",
            "path": "siae/SIAE-PMFTP-MIB.mib"
        },
        {
            "mib": "TROPIC-GENERIC-NOTIFICATION-MIB",
            "path": "nokia/1830/TROPIC-GENERIC-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "CTRON-SSR-TRAP-MIB",
            "path": "enterasys/CTRON-SSR-TRAP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-RIPNG-MIB",
            "path": "nokia/ALCATEL-IND1-RIPNG-MIB.mib"
        },
        {
            "mib": "PERFORMANCE-MIB",
            "path": "sagemcom/PERFORMANCE-MIB.mib"
        },
        {
            "mib": "NETGEAR-DOT1X-ADVANCED-FEATURES-MIB",
            "path": "quanta/dot1xAdvanced.my.mib"
        },
        {
            "mib": "ROOMALERT11E-MIB",
            "path": "avtech/ROOMALERT11E-MIB.mib"
        },
        {
            "mib": "HH3C-CONFIG-MAN-MIB",
            "path": "comware/HH3C-CONFIG-MAN-MIB.mib"
        },
        {
            "mib": "CISCO-CLASS-BASED-QOS-MIB",
            "path": "cisco/CISCO-CLASS-BASED-QOS-MIB.mib"
        },
        {
            "mib": "HP-SN-IPX-MIB",
            "path": "hp/HP-SN-IPX-MIB.mib"
        },
        {
            "mib": "TROPIC-GLOBAL-REG",
            "path": "nokia/1830/TROPIC-GLOBAL-REG.mib"
        },
        {
            "mib": "DASAN-DHCP-R-MIB",
            "path": "dasan/DASAN-DHCP-R-MIB.mib"
        },
        {
            "mib": "CIENA-WS-NOTIFICATION-MIB",
            "path": "ciena/CIENA-WS-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "DLINKSW-DHCP6-CLIENT-MIB",
            "path": "dlink/DLINKSW-DHCP6-CLIENT-MIB.mib"
        },
        {
            "mib": "AT-USER-MIB",
            "path": "awplus/AT-USER-MIB.mib"
        },
        {
            "mib": "CISCO-CONFIG-MAN-MIB",
            "path": "cisco/CISCO-CONFIG-MAN-MIB.mib"
        },
        {
            "mib": "ATM-TC-MIB",
            "path": "junose/ATM-TC-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-DHCP-MIB",
            "path": "screenos/NETSCREEN-SET-DHCP-MIB.mib"
        },
        {
            "mib": "HH3C-CONTEXT-MIB",
            "path": "comware/HH3C-CONTEXT-MIB.mib"
        },
        {
            "mib": "IPMCAST-MIB",
            "path": "junos/IPMCAST-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-ROUTEMAP-MIB",
            "path": "nokia/ALCATEL-IND1-ROUTEMAP-MIB.mib"
        },
        {
            "mib": "NETGEAR-DOT1X-AUTHENTICATION-SERVER-MIB",
            "path": "quanta/dot1x_auth_serv.my.mib"
        },
        {
            "mib": "CIENA-WS-PLATFORM-ALARM-MIB",
            "path": "ciena/CIENA-WS-PLATFORM-ALARM-MIB.mib"
        },
        {
            "mib": "ROOMALERT12E-MIB",
            "path": "avtech/ROOMALERT12E-MIB.mib"
        },
        {
            "mib": "HP-SN-MPLS-LSR-MIB",
            "path": "hp/HP-SN-MPLS-LSR-MIB.mib"
        },
        {
            "mib": "HUAWEI-FWD-PAF-TRAP-MIB",
            "path": "huawei/HUAWEI-FWD-PAF-TRAP-MIB.mib"
        },
        {
            "mib": "Juniper-UNI-SMI",
            "path": "juniper/Juniper-UNI-SMI.mib"
        },
        {
            "mib": "HH3C-CUPM-CP-MIB",
            "path": "comware/HH3C-CUPM-CP-MIB.mib"
        },
        {
            "mib": "SIAE-PMG828-MIB",
            "path": "siae/SIAE-PMG828-MIB.mib"
        },
        {
            "mib": "CTRON-TIMED-RESET-MIB",
            "path": "enterasys/CTRON-TIMED-RESET-MIB.mib"
        },
        {
            "mib": "PPP-BRIDGE-NCP-MIB",
            "path": "sagemcom/PPP-BRIDGE-NCP-MIB.mib"
        },
        {
            "mib": "DASAN-DSL-MIB",
            "path": "dasan/DASAN-DSL-MIB.mib"
        },
        {
            "mib": "IPV6-FLOW-LABEL-MIB",
            "path": "junos/IPV6-FLOW-LABEL-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PLATFORM-ENCRYPTION-MIB",
            "path": "ciena/CIENA-WS-PLATFORM-ENCRYPTION-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SAA-MIB",
            "path": "nokia/ALCATEL-IND1-SAA-MIB.mib"
        },
        {
            "mib": "DLINKSW-DHCP6-GUARD-MIB",
            "path": "dlink/DLINKSW-DHCP6-GUARD-MIB.mib"
        },
        {
            "mib": "ROOMALERT24E-MIB",
            "path": "avtech/ROOMALERT24E-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-DNS-MIB",
            "path": "screenos/NETSCREEN-SET-DNS-MIB.mib"
        },
        {
            "mib": "TROPIC-L1SERVICE-MIB",
            "path": "nokia/1830/TROPIC-L1SERVICE-MIB.mib"
        },
        {
            "mib": "ATM2-MIB",
            "path": "junose/ATM2-MIB.mib"
        },
        {
            "mib": "HH3C-CUPM-UP-MIB",
            "path": "comware/HH3C-CUPM-UP-MIB.mib"
        },
        {
            "mib": "LAG-MIB",
            "path": "quanta/dot3ad.my.mib"
        },
        {
            "mib": "SIAE-PMRXPWR-MIB",
            "path": "siae/SIAE-PMRXPWR-MIB.mib"
        },
        {
            "mib": "CISCO-CONTEXT-MAPPING-MIB",
            "path": "cisco/CISCO-CONTEXT-MAPPING-MIB.mib"
        },
        {
            "mib": "HH3C-CUSP-MIB",
            "path": "comware/HH3C-CUSP-MIB.mib"
        },
        {
            "mib": "CTRON-TRANSLATION-MIB",
            "path": "enterasys/CTRON-TRANSLATION-MIB.mib"
        },
        {
            "mib": "TROPIC-OCH-MIB",
            "path": "nokia/1830/TROPIC-OCH-MIB.mib"
        },
        {
            "mib": "JNX-DOT3OAM-CAPABILITY",
            "path": "junos/JNX-DOT3OAM-CAPABILITY.mib"
        },
        {
            "mib": "DLINKSW-DHCP6-RELAY-MIB",
            "path": "dlink/DLINKSW-DHCP6-RELAY-MIB.mib"
        },
        {
            "mib": "HP-SN-MPLS-TC-MIB",
            "path": "hp/HP-SN-MPLS-TC-MIB.mib"
        },
        {
            "mib": "DASAN-EPON-MIB",
            "path": "dasan/DASAN-EPON-MIB.mib"
        },
        {
            "mib": "AT-UWC-WLAN-SWITCH-MIB",
            "path": "awplus/AT-UWC-WLAN-SWITCH-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PLATFORM-PM-MIB",
            "path": "ciena/CIENA-WS-PLATFORM-PM-MIB.mib"
        },
        {
            "mib": "SIAE-PMTXPWR-MIB",
            "path": "siae/SIAE-PMTXPWR-MIB.mib"
        },
        {
            "mib": "HUAWEI-FWD-RES-TRAP-MIB",
            "path": "huawei/HUAWEI-FWD-RES-TRAP-MIB.mib"
        },
        {
            "mib": "NETGEAR-BGP-MIB",
            "path": "quanta/fastpathbgp.my.mib"
        },
        {
            "mib": "NETSCREEN-SET-EMAIL-MIB",
            "path": "screenos/NETSCREEN-SET-EMAIL-MIB.mib"
        },
        {
            "mib": "CISCO-DIAL-CONTROL-MIB",
            "path": "cisco/CISCO-DIAL-CONTROL-MIB.mib"
        },
        {
            "mib": "PPP-LCP-MIB",
            "path": "sagemcom/PPP-LCP-MIB.mib"
        },
        {
            "mib": "CTRON-TX-QUEUE-ARBITRATION-MIB",
            "path": "enterasys/CTRON-TX-QUEUE-ARBITRATION-MIB.mib"
        },
        {
            "mib": "SIAE-POWER-SUPPLY-MIB",
            "path": "siae/SIAE-POWER-SUPPLY-MIB.mib"
        },
        {
            "mib": "ROOMALERT26W-MIB",
            "path": "avtech/ROOMALERT26W-MIB.mib"
        },
        {
            "mib": "HH3C-DAR-MIB",
            "path": "comware/HH3C-DAR-MIB.mib"
        },
        {
            "mib": "NETGEAR-DNS-RESOLVER-CONTROL-MIB",
            "path": "quanta/fastpathdnsclient_control.my.mib"
        },
        {
            "mib": "JUNIPER-WX-COMMON-MIB",
            "path": "juniper/JUNIPER-WX-COMMON-MIB.mib"
        },
        {
            "mib": "HUAWEI-GTL-MIB",
            "path": "huawei/HUAWEI-GTL-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SERVICE-MGR-MIB",
            "path": "nokia/ALCATEL-IND1-SERVICE-MGR-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PLATFORM-TYPEDEFS-MIB",
            "path": "ciena/CIENA-WS-PLATFORM-TYPEDEFS-MIB.mib"
        },
        {
            "mib": "HP-SN-MPLS-TE-MIB",
            "path": "hp/HP-SN-MPLS-TE-MIB.mib"
        },
        {
            "mib": "AT-VCSTACK-MIB",
            "path": "awplus/AT-VCSTACK-MIB.mib"
        },
        {
            "mib": "DVMRP-STD-MIB-JUNI",
            "path": "junose/DVMRP-STD-MIB-JUNI.mib"
        },
        {
            "mib": "DLINKSW-DHCP6-SERVER-MIB",
            "path": "dlink/DLINKSW-DHCP6-SERVER-MIB.mib"
        },
        {
            "mib": "ROOMALERT32E-MIB",
            "path": "avtech/ROOMALERT32E-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-GEN-MIB",
            "path": "screenos/NETSCREEN-SET-GEN-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SESSION-MGR-MIB",
            "path": "nokia/ALCATEL-IND1-SESSION-MGR-MIB.mib"
        },
        {
            "mib": "DASAN-GEPON-MIB",
            "path": "dasan/DASAN-GEPON-MIB.mib"
        },
        {
            "mib": "HH3C-DHCP-SERVER-MIB",
            "path": "comware/HH3C-DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "JNX-GDOI-MIB",
            "path": "junos/JNX-GDOI-MIB.mib"
        },
        {
            "mib": "SIAE-PTP-MIB",
            "path": "siae/SIAE-PTP-MIB.mib"
        },
        {
            "mib": "NETGEAR-INVENTORY-MIB",
            "path": "quanta/fastpathinventory.my.mib"
        },
        {
            "mib": "CISCO-DMN-DSG-DIAG-MIB",
            "path": "cisco/CISCO-DMN-DSG-DIAG-MIB.mib"
        },
        {
            "mib": "IF-INVERTED-STACK-MIB",
            "path": "junose/IF-INVERTED-STACK-MIB.mib"
        },
        {
            "mib": "PROTECTION-MIB",
            "path": "sagemcom/PROTECTION-MIB.mib"
        },
        {
            "mib": "AT-VLAN-MIB",
            "path": "awplus/AT-VLAN-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-GLOBAL-REG",
            "path": "juniper/JUNIPER-WX-GLOBAL-REG.mib"
        },
        {
            "mib": "DLINKSW-DNS-MIB",
            "path": "dlink/DLINKSW-DNS-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-GLB-MIB",
            "path": "screenos/NETSCREEN-SET-GLB-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SLB-MIB",
            "path": "nokia/ALCATEL-IND1-SLB-MIB.mib"
        },
        {
            "mib": "HUAWEI-GTSM-MIB",
            "path": "huawei/HUAWEI-GTSM-MIB.mib"
        },
        {
            "mib": "CTRON-UPS-MIB",
            "path": "enterasys/CTRON-UPS-MIB.mib"
        },
        {
            "mib": "JNX-IF-CAPABILITY",
            "path": "junos/JNX-IF-CAPABILITY.mib"
        },
        {
            "mib": "CISCO-DMN-DSG-ROOT-MIB",
            "path": "cisco/CISCO-DMN-DSG-ROOT-MIB.mib"
        },
        {
            "mib": "AT-XEM-MIB",
            "path": "awplus/AT-XEM-MIB.mib"
        },
        {
            "mib": "ROOMALERT32S-MIB",
            "path": "avtech/ROOMALERT32S-MIB.mib"
        },
        {
            "mib": "SIAE-QUEUE-DEPTH-MIB",
            "path": "siae/SIAE-QUEUE-DEPTH-MIB.mib"
        },
        {
            "mib": "HP-SN-OSPF-GROUP-MIB",
            "path": "hp/HP-SN-OSPF-GROUP-MIB.mib"
        },
        {
            "mib": "SAGEM-DR-MIB",
            "path": "sagemcom/SAGEM-DR-MIB.mib"
        },
        {
            "mib": "HH3C-DHCP-SNOOP2-MIB",
            "path": "comware/HH3C-DHCP-SNOOP2-MIB.mib"
        },
        {
            "mib": "DLINKSW-DOS-PREVENT-MIB",
            "path": "dlink/DLINKSW-DOS-PREVENT-MIB.mib"
        },
        {
            "mib": "CISCO-DMN-DSG-TUNING-MIB",
            "path": "cisco/CISCO-DMN-DSG-TUNING-MIB.mib"
        },
        {
            "mib": "DASAN-GFAST-MIB",
            "path": "dasan/DASAN-GFAST-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SNMP-AGENT-MIB",
            "path": "nokia/ALCATEL-IND1-SNMP-AGENT-MIB.mib"
        },
        {
            "mib": "CTRON-VLAN-CLASSIFY-MIB",
            "path": "enterasys/CTRON-VLAN-CLASSIFY-MIB.mib"
        },
        {
            "mib": "Juniper-AAA-MIB",
            "path": "junose/Juniper-AAA-MIB.mib"
        },
        {
            "mib": "NETGEAR-IPV6-LOOPBACK-MIB",
            "path": "quanta/fastpathipv6loopback.my.mib"
        },
        {
            "mib": "ROOMALERT3E-MIB",
            "path": "avtech/ROOMALERT3E-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PM-MIB",
            "path": "ciena/CIENA-WS-PM-MIB.mib"
        },
        {
            "mib": "HP-SN-POS-GROUP-MIB",
            "path": "hp/HP-SN-POS-GROUP-MIB.mib"
        },
        {
            "mib": "HUAWEI-HGMP-MIB",
            "path": "huawei/HUAWEI-HGMP-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-LOG-MIB",
            "path": "screenos/NETSCREEN-SET-LOG-MIB.mib"
        },
        {
            "mib": "CTRON-VLAN-EXTENSIONS-MIB",
            "path": "enterasys/CTRON-VLAN-EXTENSIONS-MIB.mib"
        },
        {
            "mib": "SDH-ETS-MIB",
            "path": "sagemcom/SDH-ETS-MIB.mib"
        },
        {
            "mib": "ATKK-WLAN-ACCESS",
            "path": "awplus/ATKK-WLAN-ACCESS-MIB.mib"
        },
        {
            "mib": "DASAN-GIGABIT-OPTIC-TRANSCEIVER-MIB",
            "path": "dasan/DASAN-GIGABIT-OPTIC-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-GLOBAL-TC",
            "path": "juniper/JUNIPER-WX-GLOBAL-TC.mib"
        },
        {
            "mib": "JNX-IP-CAPABILITY",
            "path": "junos/JNX-IP-CAPABILITY.mib"
        },
        {
            "mib": "DLINKSW-DOT1X-EXT-MIB",
            "path": "dlink/DLINKSW-DOT1X-EXT-MIB.mib"
        },
        {
            "mib": "HP-SN-ROOT-MIB",
            "path": "hp/HP-SN-ROOT-MIB.mib"
        },
        {
            "mib": "TROPIC-OPTICALPORT-MIB",
            "path": "nokia/1830/TROPIC-OPTICALPORT-MIB.mib"
        },
        {
            "mib": "SIAE-RADIO-ENCRYPTION-MIB",
            "path": "siae/SIAE-RADIO-ENCRYPTION-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PORT-MIB",
            "path": "ciena/CIENA-WS-PORT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SSH-MIB",
            "path": "nokia/ALCATEL-IND1-SSH-MIB.mib"
        },
        {
            "mib": "SESSION-MIB",
            "path": "sagemcom/SESSION-MIB.mib"
        },
        {
            "mib": "HUAWEI-HQOS-MIB",
            "path": "huawei/HUAWEI-HQOS-MIB.mib"
        },
        {
            "mib": "CISCO-DOCS-EXT-MIB",
            "path": "cisco/CISCO-DOCS-EXT-MIB.mib"
        },
        {
            "mib": "HH3C-DHCP4-MIB",
            "path": "comware/HH3C-DHCP4-MIB.mib"
        },
        {
            "mib": "CTRON-WAN-IMUX-MIB",
            "path": "enterasys/CTRON-WAN-IMUX-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-SNMP-MIB",
            "path": "screenos/NETSCREEN-SET-SNMP-MIB.mib"
        },
        {
            "mib": "TROPIC-OTH-MIB",
            "path": "nokia/1830/TROPIC-OTH-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PTP-MIB",
            "path": "ciena/CIENA-WS-PTP-MIB.mib"
        },
        {
            "mib": "DLINKSW-ENTITY-EXT-MIB",
            "path": "dlink/DLINKSW-ENTITY-EXT-MIB.mib"
        },
        {
            "mib": "NETGEAR-IPV6-TUNNEL-MIB",
            "path": "quanta/fastpathipv6tunnel.my.mib"
        },
        {
            "mib": "HUAWEI-HTTP-MIB",
            "path": "huawei/HUAWEI-HTTP-MIB.mib"
        },
        {
            "mib": "DASAN-MCAST-MIB",
            "path": "dasan/DASAN-MCAST-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-MIB",
            "path": "juniper/JUNIPER-WX-MIB.mib"
        },
        {
            "mib": "ROOMALERT4E-MIB",
            "path": "avtech/ROOMALERT4E-MIB.mib"
        },
        {
            "mib": "HP-SN-ROUTER-TRAP-MIB",
            "path": "hp/HP-SN-ROUTER-TRAP-MIB.mib"
        },
        {
            "mib": "HH3C-DHCP6-MIB",
            "path": "comware/HH3C-DHCP6-MIB.mib"
        },
        {
            "mib": "CTRON-WAN-MIB",
            "path": "enterasys/CTRON-WAN-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-SYSTIME-MIB",
            "path": "screenos/NETSCREEN-SET-SYSTIME-MIB.mib"
        },
        {
            "mib": "DLINKSW-ERPS-MIB",
            "path": "dlink/DLINKSW-ERPS-MIB.mib"
        },
        {
            "mib": "JNX-IPSEC-MONITOR-MIB",
            "path": "junos/JNX-IPSEC-MONITOR-MIB.mib"
        },
        {
            "mib": "CISCO-DOT11-ASSOCIATION-MIB",
            "path": "cisco/CISCO-DOT11-ASSOCIATION-MIB.mib"
        },
        {
            "mib": "NETGEAR-ISDP-MIB",
            "path": "quanta/fastpathisdp.my.mib"
        },
        {
            "mib": "Juniper-AAA-Server-CONF",
            "path": "junose/Juniper-AAA-Server-CONF.mib"
        },
        {
            "mib": "SIAE-RADIO-SYSTEM-MIB",
            "path": "siae/SIAE-RADIO-SYSTEM-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SET-URL-FILTER-MIB",
            "path": "screenos/NETSCREEN-SET-URL-FILTER-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-STACK-MANAGER-MIB",
            "path": "nokia/ALCATEL-IND1-STACK-MANAGER-MIB.mib"
        },
        {
            "mib": "HP-SN-SW-L4-SWITCH-GROUP-MIB",
            "path": "hp/HP-SN-SW-L4-SWITCH-GROUP-MIB.mib"
        },
        {
            "mib": "TROPIC-OTUODU-MIB",
            "path": "nokia/1830/TROPIC-OTUODU-MIB.mib"
        },
        {
            "mib": "HH3C-DHCP6-SERVER-MIB",
            "path": "comware/HH3C-DHCP6-SERVER-MIB.mib"
        },
        {
            "mib": "CIENA-WS-PTP-MODEM-MIB",
            "path": "ciena/CIENA-WS-PTP-MODEM-MIB.mib"
        },
        {
            "mib": "SHELF-MIB",
            "path": "sagemcom/SHELF-MIB.mib"
        },
        {
            "mib": "DASAN-NOTIFICATION-V1",
            "path": "dasan/DASAN-NOTIFICATION-V1.mib"
        },
        {
            "mib": "TRAPEZE-NETWORKS-ROOT-MIB",
            "path": "juniper/TRAPEZE-NETWORKS-ROOT-MIB.mib"
        },
        {
            "mib": "CISCO-DOT11-IF-MIB",
            "path": "cisco/CISCO-DOT11-IF-MIB.mib"
        },
        {
            "mib": "HH3C-DHCPR-MIB",
            "path": "comware/HH3C-DHCPR-MIB.mib"
        },
        {
            "mib": "UNITRENDS-SNMP",
            "path": "unitrends/UNITRENDS-SNMP.mib"
        },
        {
            "mib": "NETSCREEN-SET-WEB-MIB",
            "path": "screenos/NETSCREEN-SET-WEB-MIB.mib"
        },
        {
            "mib": "HUAWEI-HWTACACS-MIB",
            "path": "huawei/HUAWEI-HWTACACS-MIB.mib"
        },
        {
            "mib": "SIAE-RET-MIB",
            "path": "siae/SIAE-RET-MIB.mib"
        },
        {
            "mib": "DATA-DOMAIN-MIB",
            "path": "datadomain/DATA-DOMAIN-MIB.mib"
        },
        {
            "mib": "NETGEAR-LOGGING-MIB",
            "path": "quanta/fastpathlogging.my.mib"
        },
        {
            "mib": "CIENA-WS-PTP-PLUGGABLE-MIB",
            "path": "ciena/CIENA-WS-PTP-PLUGGABLE-MIB.mib"
        },
        {
            "mib": "Juniper-ACCOUNTING-MIB",
            "path": "junose/Juniper-ACCOUNTING-MIB.mib"
        },
        {
            "mib": "ROOMALERT7E-MIB",
            "path": "avtech/ROOMALERT7E-MIB.mib"
        },
        {
            "mib": "CTRON-WAN-MULTI-IMUX-MIB",
            "path": "enterasys/CTRON-WAN-MULTI-IMUX-MIB.mib"
        },
        {
            "mib": "JNX-L2TP-MIB",
            "path": "junos/JNX-L2TP-MIB.mib"
        },
        {
            "mib": "DLINKSW-ERROR-DISABLE-MIB",
            "path": "dlink/DLINKSW-ERROR-DISABLE-MIB.mib"
        },
        {
            "mib": "NETGEAR-LOOPBACK-MIB",
            "path": "quanta/fastpathloopback.my.mib"
        },
        {
            "mib": "CISCO-ENHANCED-IMAGE-MIB",
            "path": "cisco/CISCO-ENHANCED-IMAGE-MIB.mib"
        },
        {
            "mib": "HP-SN-SWITCH-GROUP-MIB",
            "path": "hp/HP-SN-SWITCH-GROUP-MIB.mib"
        },
        {
            "mib": "HH3C-DHCPRELAY-MIB",
            "path": "comware/HH3C-DHCPRELAY-MIB.mib"
        },
        {
            "mib": "CIENA-WS-SERVICE-DOMAIN-MIB",
            "path": "ciena/CIENA-WS-SERVICE-DOMAIN-MIB.mib"
        },
        {
            "mib": "SPRIF-MIB",
            "path": "sagemcom/SPRIF-MIB.mib"
        },
        {
            "mib": "DASAN-NOTIFICATION",
            "path": "dasan/DASAN-NOTIFICATION.mib"
        },
        {
            "mib": "SIAE-SAFE-MODE-MIB",
            "path": "siae/SIAE-SAFE-MODE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-SYSTEM-MIB",
            "path": "nokia/ALCATEL-IND1-SYSTEM-MIB.mib"
        },
        {
            "mib": "NETSCREEN-SMI",
            "path": "screenos/NETSCREEN-SMI.mib"
        },
        {
            "mib": "JNX-MPLS-TE-P2MP-STD-MIB",
            "path": "junos/JNX-MPLS-TE-P2MP-STD-MIB.mib"
        },
        {
            "mib": "TROPIC-PSD-MIB",
            "path": "nokia/1830/TROPIC-PSD-MIB.mib"
        },
        {
            "mib": "EKINOPS-MGNT2-MIB",
            "path": "ekinops/EKINOPS-MGNT2-MIB.mib"
        },
        {
            "mib": "TRAPEZE-NETWORKS-SYSTEM-MIB",
            "path": "juniper/TRAPEZE-NETWORKS-SYSTEM-MIB.mib"
        },
        {
            "mib": "NETGEAR-MULTICAST-MIB",
            "path": "quanta/fastpathmulticast.my.mib"
        },
        {
            "mib": "DLINKSW-FS-MIB",
            "path": "dlink/DLINKSW-FS-MIB.mib"
        },
        {
            "mib": "HP-SN-TRAP-MIB",
            "path": "hp/HP-SN-TRAP-MIB.mib"
        },
        {
            "mib": "SIAE-SECURITY-MANAGEMENT-MIB",
            "path": "siae/SIAE-SECURITY-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "CISCO-ENHANCED-MEMPOOL-MIB",
            "path": "cisco/CISCO-ENHANCED-MEMPOOL-MIB.mib"
        },
        {
            "mib": "Juniper-ADDRESS-POOL-MIB",
            "path": "junose/Juniper-ADDRESS-POOL-MIB.mib"
        },
        {
            "mib": "ROOMALERTST4E-MIB",
            "path": "avtech/ROOMALERTST4E-MIB.mib"
        },
        {
            "mib": "CIENA-WS-SERVICE-MIB",
            "path": "ciena/CIENA-WS-SERVICE-MIB.mib"
        },
        {
            "mib": "XCONNECTION-MIB",
            "path": "sagemcom/XCONNECTION-MIB.mib"
        },
        {
            "mib": "DASAN-PRODUCTS-MIB",
            "path": "dasan/DASAN-PRODUCTS-MIB.mib"
        },
        {
            "mib": "NETSCREEN-TRAP-MIB",
            "path": "screenos/NETSCREEN-TRAP-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TP-DEVICES",
            "path": "nokia/ALCATEL-IND1-TP-DEVICES.mib"
        },
        {
            "mib": "CTRON-WEBVIEW-MIB",
            "path": "enterasys/CTRON-WEBVIEW-MIB.mib"
        },
        {
            "mib": "TELTONIKA-MIB",
            "path": "witchos/TELTONIKA-MIB.mib"
        },
        {
            "mib": "NETGEAR-NSF-MIB",
            "path": "quanta/fastpathnsf.my.mib"
        },
        {
            "mib": "JNX-OPT-IF-EXT-MIB",
            "path": "junos/JNX-OPT-IF-EXT-MIB.mib"
        },
        {
            "mib": "HUAWEI-IF-EXT-MIB",
            "path": "huawei/HUAWEI-IF-EXT-MIB.mib"
        },
        {
            "mib": "EKINOPS-MGNT2-NMS-MIB",
            "path": "ekinops/EKINOPS-MGNT2-NMS-MIB.mib"
        },
        {
            "mib": "DLINKSW-GENMGMT-MIB",
            "path": "dlink/DLINKSW-GENMGMT-MIB.mib"
        },
        {
            "mib": "SIAE-SENSOR-TEMP-MIB",
            "path": "siae/SIAE-SENSOR-TEMP-MIB.mib"
        },
        {
            "mib": "CIENA-WS-SOFTWARE-MIB",
            "path": "ciena/CIENA-WS-SOFTWARE-MIB.mib"
        },
        {
            "mib": "CISCO-ENHANCED-SLB-MIB",
            "path": "cisco/CISCO-ENHANCED-SLB-MIB.mib"
        },
        {
            "mib": "EKINOPS-MIB",
            "path": "ekinops/EKINOPS-MIB.mib"
        },
        {
            "mib": "TROPIC-PTP-MIB",
            "path": "nokia/1830/TROPIC-PTP-MIB.mib"
        },
        {
            "mib": "HH3C-DHCPS-MIB",
            "path": "comware/HH3C-DHCPS-MIB.mib"
        },
        {
            "mib": "NETGEAR-ROUTE-POLICY-MIB",
            "path": "quanta/fastpathroutepolicy.my.mib"
        },
        {
            "mib": "CISCO-ENHANCED-WRED-MIB",
            "path": "cisco/CISCO-ENHANCED-WRED-MIB.mib"
        },
        {
            "mib": "HUAWEI-IF-QOS-MIB",
            "path": "huawei/HUAWEI-IF-QOS-MIB.mib"
        },
        {
            "mib": "CTSMTMIB-MIB",
            "path": "enterasys/CTSMTMIB-MIB.mib"
        },
        {
            "mib": "TEMPAGER-MIB",
            "path": "avtech/TEMPAGER-MIB.mib"
        },
        {
            "mib": "HP-SN-VSRP-MIB",
            "path": "hp/HP-SN-VSRP-MIB.mib"
        },
        {
            "mib": "NETSCREEN-UAC-MIB",
            "path": "screenos/NETSCREEN-UAC-MIB.mib"
        },
        {
            "mib": "SIAE-SFP-MIB",
            "path": "siae/SIAE-SFP-MIB.mib"
        },
        {
            "mib": "TROPIC-SHELF-MIB",
            "path": "nokia/1830/TROPIC-SHELF-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-TRAP-MGR-MIB",
            "path": "nokia/ALCATEL-IND1-TRAP-MGR-MIB.mib"
        },
        {
            "mib": "A10-AX-MIB",
            "path": "a10/A10-AX-MIB.mib"
        },
        {
            "mib": "DELIBERANT-MIB",
            "path": "deliberant/DELIBERANT-MIB.mib"
        },
        {
            "mib": "DLINKSW-GVRP-MIB",
            "path": "dlink/DLINKSW-GVRP-MIB.mib"
        },
        {
            "mib": "DASAN-QOS-MIB",
            "path": "dasan/DASAN-QOS-MIB.mib"
        },
        {
            "mib": "Juniper-Agents",
            "path": "junose/Juniper-Agents.mib"
        },
        {
            "mib": "NETGEAR-ROUTING-MIB",
            "path": "quanta/fastpathrouting.my.mib"
        },
        {
            "mib": "HH3C-DHCPSNOOP-MIB",
            "path": "comware/HH3C-DHCPSNOOP-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-ALARM-MIB",
            "path": "cisco/CISCO-ENTITY-ALARM-MIB.mib"
        },
        {
            "mib": "BAICELLS-MIB",
            "path": "baicells/cpe/BAICELLS-MIB.mib"
        },
        {
            "mib": "TROPIC-SLOT-MIB",
            "path": "nokia/1830/TROPIC-SLOT-MIB.mib"
        },
        {
            "mib": "HUAWEI-IFIT-MIB",
            "path": "huawei/HUAWEI-IFIT-MIB.mib"
        },
        {
            "mib": "CIENA-WS-SYSTEM-MIB",
            "path": "ciena/CIENA-WS-SYSTEM-MIB.mib"
        },
        {
            "mib": "JNX-OPT-IF-MIB",
            "path": "junos/JNX-OPT-IF-MIB.mib"
        },
        {
            "mib": "TEMPAGER3E-MIB",
            "path": "avtech/TEMPAGER3E-MIB.mib"
        },
        {
            "mib": "A10-AX-NOTIFICATIONS",
            "path": "a10/A10-AX-NOTIFICATIONS.mib"
        },
        {
            "mib": "CTTRAPLOG-MIB",
            "path": "enterasys/CTTRAPLOG-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-GATEWAY-MIB",
            "path": "screenos/NETSCREEN-VPN-GATEWAY-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-ASSET-MIB",
            "path": "cisco/CISCO-ENTITY-ASSET-MIB.mib"
        },
        {
            "mib": "HH3C-DISK-MIB",
            "path": "comware/HH3C-DISK-MIB.mib"
        },
        {
            "mib": "JNX-PPP-MIB",
            "path": "junos/JNX-PPP-MIB.mib"
        },
        {
            "mib": "HUAWEI-IMA-MIB",
            "path": "huawei/HUAWEI-IMA-MIB.mib"
        },
        {
            "mib": "SIAE-SOFT-MIB",
            "path": "siae/SIAE-SOFT-MIB.mib"
        },
        {
            "mib": "TN-ACL-MIB",
            "path": "transition/TN-ACL-MIB.mib"
        },
        {
            "mib": "A10-COMMON-MIB",
            "path": "a10/A10-COMMON-MIB.mib"
        },
        {
            "mib": "NETGEAR-ROUTING6-MIB",
            "path": "quanta/fastpathrouting6.my.mib"
        },
        {
            "mib": "TROPIC-SOFTWARE-MIB",
            "path": "nokia/1830/TROPIC-SOFTWARE-MIB.mib"
        },
        {
            "mib": "Juniper-ATM-1483-Profile-MIB",
            "path": "junose/Juniper-ATM-1483-Profile-MIB.mib"
        },
        {
            "mib": "DLB-802DOT11-EXT-MIB",
            "path": "deliberant/DLB-802DOT11-EXT-MIB.mib"
        },
        {
            "mib": "HP-SNTPclientConfiguration-MIB",
            "path": "hp/HP-SNTPclientConfiguration-MIB.mib"
        },
        {
            "mib": "CIENA-WS-TYPEDEFS-MIB",
            "path": "ciena/CIENA-WS-TYPEDEFS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-UDLD-MIB",
            "path": "nokia/ALCATEL-IND1-UDLD-MIB.mib"
        },
        {
            "mib": "ENTERASYS-MIB-NAMES",
            "path": "enterasys/ENTERASYS-MIB-NAMES.mib"
        },
        {
            "mib": "DASAN-ROUTER-MIB",
            "path": "dasan/DASAN-ROUTER-MIB.mib"
        },
        {
            "mib": "JNX-PPPOE-MIB",
            "path": "junos/JNX-PPPOE-MIB.mib"
        },
        {
            "mib": "HUAWEI-INFOCENTER-MIB",
            "path": "huawei/HUAWEI-INFOCENTER-MIB.mib"
        },
        {
            "mib": "CAREL-RITTAL-LCP-3311-MIB",
            "path": "carel/CAREL-RITTAL-LCP-3311-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-DIAG-MIB",
            "path": "cisco/CISCO-ENTITY-DIAG-MIB.mib"
        },
        {
            "mib": "DLINKSW-IF-COUNTER-MIB",
            "path": "dlink/DLINKSW-IF-COUNTER-MIB.mib"
        },
        {
            "mib": "NETGEAR-SFLOW-MIB",
            "path": "quanta/fastpathsflow.my.mib"
        },
        {
            "mib": "SIAE-SYNC-MIB",
            "path": "siae/SIAE-SYNC-MIB.mib"
        },
        {
            "mib": "JNX-SNMPv2-CAPABILITY",
            "path": "junos/JNX-SNMPv2-CAPABILITY.mib"
        },
        {
            "mib": "ENTERASYS-RESOURCE-UTILIZATION-MIB",
            "path": "enterasys/ENTERASYS-RESOURCE-UTILIZATION-MIB.mib"
        },
        {
            "mib": "HH3C-DLDP-MIB",
            "path": "comware/HH3C-DLDP-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-IKE-MIB",
            "path": "screenos/NETSCREEN-VPN-IKE-MIB.mib"
        },
        {
            "mib": "DLB-ATHDRV-STATS-MIB",
            "path": "deliberant/DLB-ATHDRV-STATS-MIB.mib"
        },
        {
            "mib": "Juniper-ATM-CONF",
            "path": "junose/Juniper-ATM-CONF.mib"
        },
        {
            "mib": "CISCO-ENTITY-DIAG-TC-MIB",
            "path": "cisco/CISCO-ENTITY-DIAG-TC-MIB.mib"
        },
        {
            "mib": "DASAN-SHDSL-MIB",
            "path": "dasan/DASAN-SHDSL-MIB.mib"
        },
        {
            "mib": "DLINKSW-IMPB-MIB",
            "path": "dlink/DLINKSW-IMPB-MIB.mib"
        },
        {
            "mib": "SIAE-TREE-MIB",
            "path": "siae/SIAE-TREE-MIB.mib"
        },
        {
            "mib": "BGP4V2-MIB",
            "path": "brocade/BGP4V2-MIB.mib"
        },
        {
            "mib": "JUNIPER-ALARM-EXT-MIB",
            "path": "junos/JUNIPER-ALARM-EXT-MIB.mib"
        },
        {
            "mib": "TN-ARP-INSPECTION-MIB",
            "path": "transition/TN-ARP-INSPECTION-MIB.mib"
        },
        {
            "mib": "DASAN-SMI",
            "path": "dasan/DASAN-SMI.mib"
        },
        {
            "mib": "CAREL-ug40cdz-MIB",
            "path": "carel/CAREL-ug40cdz-MIB.mib"
        },
        {
            "mib": "CIENA-WS-XCVR-MIB",
            "path": "ciena/CIENA-WS-XCVR-MIB.mib"
        },
        {
            "mib": "HP-STACK-MIB",
            "path": "hp/HP-STACK-MIB.mib"
        },
        {
            "mib": "SIAE-UNIT-MIB",
            "path": "siae/SIAE-UNIT-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-UDP-RELAY-MIB",
            "path": "nokia/ALCATEL-IND1-UDP-RELAY-MIB.mib"
        },
        {
            "mib": "NETGEAR-SNTP-CLIENT-MIB",
            "path": "quanta/fastpathsntp.my.mib"
        },
        {
            "mib": "HH3C-DLDP2-MIB",
            "path": "comware/HH3C-DLDP2-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-DISPLAY-MIB",
            "path": "cisco/CISCO-ENTITY-DISPLAY-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-L2TP-MIB",
            "path": "screenos/NETSCREEN-VPN-L2TP-MIB.mib"
        },
        {
            "mib": "IRM-OIDS",
            "path": "enterasys/IRM-OIDS.mib"
        },
        {
            "mib": "DLB-GENERIC-MIB",
            "path": "deliberant/DLB-GENERIC-MIB.mib"
        },
        {
            "mib": "HUAWEI-INNER-LINK-MIB",
            "path": "huawei/HUAWEI-INNER-LINK-MIB.mib"
        },
        {
            "mib": "KELVIN-pCOWeb-Chiller-MIB",
            "path": "carel/KELVIN-pCOWeb-Chiller-MIB.mib"
        },
        {
            "mib": "SIAE-UNITYPE-MIB",
            "path": "siae/SIAE-UNITYPE-MIB.mib"
        },
        {
            "mib": "TN-CES-MIB",
            "path": "transition/TN-CES-MIB.mib"
        },
        {
            "mib": "HH3C-DNS-MIB",
            "path": "comware/HH3C-DNS-MIB.mib"
        },
        {
            "mib": "Juniper-Autoconfigure-CONF",
            "path": "junose/Juniper-Autoconfigure-CONF.mib"
        },
        {
            "mib": "DLINKSW-IP-EXT-MIB",
            "path": "dlink/DLINKSW-IP-EXT-MIB.mib"
        },
        {
            "mib": "BROCADE-IEEE8021-PAE-CAPABILITY-MIB",
            "path": "brocade/BROCADE-IEEE8021-PAE-CAPABILITY-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VIRTUALROUTER-MIB",
            "path": "nokia/ALCATEL-IND1-VIRTUALROUTER-MIB.mib"
        },
        {
            "mib": "CIENA-WS-XCVR-MODEM-MIB",
            "path": "ciena/CIENA-WS-XCVR-MODEM-MIB.mib"
        },
        {
            "mib": "JUNIPER-ALARM-MIB",
            "path": "junos/JUNIPER-ALARM-MIB.mib"
        },
        {
            "mib": "TROPIC-STATISTICS-MIB",
            "path": "nokia/1830/TROPIC-STATISTICS-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-EXT-MIB",
            "path": "cisco/CISCO-ENTITY-EXT-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-MANUAL-MIB",
            "path": "screenos/NETSCREEN-VPN-MANUAL-MIB.mib"
        },
        {
            "mib": "HP-SWITCH-PL-MIB",
            "path": "hp/HP-SWITCH-PL-MIB.mib"
        },
        {
            "mib": "DASAN-SNMP-MIB",
            "path": "dasan/DASAN-SNMP-MIB.mib"
        },
        {
            "mib": "DLB-RADIO3-DRV-MIB",
            "path": "deliberant/DLB-RADIO3-DRV-MIB.mib"
        },
        {
            "mib": "TN-CES-ROUTING-MIB",
            "path": "transition/TN-CES-ROUTING-MIB.mib"
        },
        {
            "mib": "SCS-ks-MIB",
            "path": "carel/SCS-ks-MIB.mib"
        },
        {
            "mib": "DLINKSW-IP-SOURCE-GUARD-MIB",
            "path": "dlink/DLINKSW-IP-SOURCE-GUARD-MIB.mib"
        },
        {
            "mib": "TROPIC-SYNCE-MIB",
            "path": "nokia/1830/TROPIC-SYNCE-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VLAN-MGR-MIB",
            "path": "nokia/ALCATEL-IND1-VLAN-MGR-MIB.mib"
        },
        {
            "mib": "NETWORK-DIAGS",
            "path": "enterasys/NETWORK-DIAGS.mib"
        },
        {
            "mib": "Juniper-AUTOCONFIGURE-MIB",
            "path": "junose/Juniper-AUTOCONFIGURE-MIB.mib"
        },
        {
            "mib": "HH3C-DOMAIN-MIB",
            "path": "comware/HH3C-DOMAIN-MIB.mib"
        },
        {
            "mib": "NETGEAR-SWITCHING-MIB",
            "path": "quanta/fastpathswitching.my.mib"
        },
        {
            "mib": "BROCADE-IEEE8023-LAG-CAPABILITY-MIB",
            "path": "brocade/BROCADE-IEEE8023-LAG-CAPABILITY-MIB.mib"
        },
        {
            "mib": "JUNIPER-ANALYZER-MIB",
            "path": "junos/JUNIPER-ANALYZER-MIB.mib"
        },
        {
            "mib": "CIENA-WS-XCVR-PLUGGABLE-MIB",
            "path": "ciena/CIENA-WS-XCVR-PLUGGABLE-MIB.mib"
        },
        {
            "mib": "SIAE-USER-MIB",
            "path": "siae/SIAE-USER-MIB.mib"
        },
        {
            "mib": "HUAWEI-IPFPM-MIB",
            "path": "huawei/HUAWEI-IPFPM-MIB.mib"
        },
        {
            "mib": "ATEN-PE-CFG-STR",
            "path": "aten/ATEN-PE-CFG-STR.mib"
        },
        {
            "mib": "ROUTER-OIDS",
            "path": "enterasys/ROUTER-OIDS.mib"
        },
        {
            "mib": "CISCO-ENTITY-FRU-CONTROL-MIB",
            "path": "cisco/CISCO-ENTITY-FRU-CONTROL-MIB.mib"
        },
        {
            "mib": "NETGEAR-AUTHENTICATION-MANAGER-MIB",
            "path": "quanta/fastpath_auth_mgr.my.mib"
        },
        {
            "mib": "TROPIC-SYSTEM-MIB",
            "path": "nokia/1830/TROPIC-SYSTEM-MIB.mib"
        },
        {
            "mib": "HP-SwitchStack-MIB",
            "path": "hp/HP-SwitchStack-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-QFP-MIB",
            "path": "cisco/CISCO-ENTITY-QFP-MIB.mib"
        },
        {
            "mib": "Juniper-BGP-CONF",
            "path": "junose/Juniper-BGP-CONF.mib"
        },
        {
            "mib": "NETSCREEN-VPN-MON-MIB",
            "path": "screenos/NETSCREEN-VPN-MON-MIB.mib"
        },
        {
            "mib": "TN-CONFIG-MIB",
            "path": "transition/TN-CONFIG-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VLAN-STACKING-MIB",
            "path": "nokia/ALCATEL-IND1-VLAN-STACKING-MIB.mib"
        },
        {
            "mib": "HUAWEI-IPHC-MIB",
            "path": "huawei/HUAWEI-IPHC-MIB.mib"
        },
        {
            "mib": "JUNIPER-ATM-COS-MIB",
            "path": "junos/JUNIPER-ATM-COS-MIB.mib"
        },
        {
            "mib": "NETGEAR-BOXSERVICES-PRIVATE-MIB",
            "path": "quanta/fastpath_boxservices.my.mib"
        },
        {
            "mib": "DLINKSW-IPMCAST-EXT-MIB",
            "path": "dlink/DLINKSW-IPMCAST-EXT-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-ACMT-MIB",
            "path": "comware/HH3C-DOT11-ACMT-MIB.mib"
        },
        {
            "mib": "DASAN-SWITCH-MIB",
            "path": "dasan/DASAN-SWITCH-MIB.mib"
        },
        {
            "mib": "UPS2-MIB",
            "path": "enterasys/UPS2-MIB.mib"
        },
        {
            "mib": "BROCADE-LLDP-CAPABILITY-MIB",
            "path": "brocade/BROCADE-LLDP-CAPABILITY-MIB.mib"
        },
        {
            "mib": "WWP-LEOS-PORT-MIB",
            "path": "ciena/CIENA-WWP-LEOS-PORT-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-PHASEONE-MIB",
            "path": "screenos/NETSCREEN-VPN-PHASEONE-MIB.mib"
        },
        {
            "mib": "HP-SYSTEM-MIB",
            "path": "hp/HP-SYSTEM-MIB.mib"
        },
        {
            "mib": "ATEN-PE-CFG",
            "path": "aten/ATEN-PE-CFG.mib"
        },
        {
            "mib": "TROPIC-TC",
            "path": "nokia/1830/TROPIC-TC.mib"
        },
        {
            "mib": "JUNIPER-ATM-MIB",
            "path": "junos/JUNIPER-ATM-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-SENSOR-MIB",
            "path": "cisco/CISCO-ENTITY-SENSOR-MIB.mib"
        },
        {
            "mib": "TN-DDMI-MIB",
            "path": "transition/TN-DDMI-MIB.mib"
        },
        {
            "mib": "NETGEAR-CAPTIVE-PORTAL-MIB",
            "path": "quanta/fastpath_captive_portal.my.mib"
        },
        {
            "mib": "DASAN-TC",
            "path": "dasan/DASAN-TC.mib"
        },
        {
            "mib": "HUAWEI-IPMCAST-MIB",
            "path": "huawei/HUAWEI-IPMCAST-MIB.mib"
        },
        {
            "mib": "JUNIPER-BFD-MIB",
            "path": "junos/JUNIPER-BFD-MIB.mib"
        },
        {
            "mib": "DLINKSW-IPV6-SNOOPING-MIB",
            "path": "dlink/DLINKSW-IPV6-SNOOPING-MIB.mib"
        },
        {
            "mib": "SPECTRACOM-GLOBAL-REG-MIB",
            "path": "spectracom/SPECTRACOM-GLOBAL-MIB.mib"
        },
        {
            "mib": "DASAN-THRESHOLD-MIB",
            "path": "dasan/DASAN-THRESHOLD-MIB.mib"
        },
        {
            "mib": "BROCADE-LLDP-EXT-DOT3-CAPABILITY-MIB",
            "path": "brocade/BROCADE-LLDP-EXT-DOT3-CAPABILITY-MIB.mib"
        },
        {
            "mib": "TROPIC-WAVEKEY-MIB",
            "path": "nokia/1830/TROPIC-WAVEKEY-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-APMT-MIB",
            "path": "comware/HH3C-DOT11-APMT-MIB.mib"
        },
        {
            "mib": "HP-USER-AUTH",
            "path": "hp/HP-USER-AUTH.mib"
        },
        {
            "mib": "TN-ACCESS-MGMT-MIB",
            "path": "transition/TN-DEV-ACCESS-MGMT-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-PHASETWO-MIB",
            "path": "screenos/NETSCREEN-VPN-PHASETWO-MIB.mib"
        },
        {
            "mib": "JUNIPER-CFGMGMT-MIB",
            "path": "junos/JUNIPER-CFGMGMT-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP500-MIB",
            "path": "cambium/300/CAMBIUM-PTP500-MIB.mib"
        },
        {
            "mib": "DLINKSW-IPV6-SRC-GUARD-MIB",
            "path": "dlink/DLINKSW-IPV6-SRC-GUARD-MIB.mib"
        },
        {
            "mib": "DLINK-ID-REC-MIB",
            "path": "dlink_dgs1250/DLINK-ID-REC-MIB.mib"
        },
        {
            "mib": "DASAN-TS-1000-MIB",
            "path": "dasan/DASAN-TS-1000-MIB.mib"
        },
        {
            "mib": "HP-VLAN-CAR-MIB",
            "path": "hp/HP-VLAN-CAR-MIB.mib"
        },
        {
            "mib": "SPECTRACOM-NTP-V4-MIB",
            "path": "spectracom/SPECTRACOM-NTP-V4-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VLAN-STP-MIB",
            "path": "nokia/ALCATEL-IND1-VLAN-STP-MIB.mib"
        },
        {
            "mib": "Juniper-BGP-MIB",
            "path": "junose/Juniper-BGP-MIB.mib"
        },
        {
            "mib": "TN-DEV-AGGREGATION-MIB",
            "path": "transition/TN-DEV-AGGREGATION-MIB.mib"
        },
        {
            "mib": "HUAWEI-IPPOOL-MIB",
            "path": "huawei/HUAWEI-IPPOOL-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VPN-USER-MIB",
            "path": "screenos/NETSCREEN-VPN-USER-MIB.mib"
        },
        {
            "mib": "DLINKSW-JWAC-MIB",
            "path": "dlink/DLINKSW-JWAC-MIB.mib"
        },
        {
            "mib": "DASAN-USER-MANAGEMENT-MIB",
            "path": "dasan/DASAN-USER-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "WWP-LEOS-PORT-XCVR-MIB",
            "path": "ciena/CIENA-WWP-LEOS-PORT-XCVR-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VRRP-MIB",
            "path": "nokia/ALCATEL-IND1-VRRP-MIB.mib"
        },
        {
            "mib": "BROCADE-MAPS-MIB",
            "path": "brocade/BROCADE-MAPS-MIB.mib"
        },
        {
            "mib": "Juniper-Bridge-CONF",
            "path": "junose/Juniper-Bridge-CONF.mib"
        },
        {
            "mib": "SPECTRACOM-PTP-MIB",
            "path": "spectracom/SPECTRACOM-PTP-MIB.mib"
        },
        {
            "mib": "TELTONIKA-MIB",
            "path": "rutos/TELTONIKA-MIB.mib"
        },
        {
            "mib": "HP-VLAN",
            "path": "hp/HP-VLAN.mib"
        },
        {
            "mib": "DLINKSW-DDM-MIB",
            "path": "dlink_dgs1250/DLINKSW-DDM-MIB.mib"
        },
        {
            "mib": "DLINKSW-L2FDB-MIB",
            "path": "dlink/DLINKSW-L2FDB-MIB.mib"
        },
        {
            "mib": "NETGEAR-DCBX-MIB",
            "path": "quanta/fastpath_dcbx.my.mib"
        },
        {
            "mib": "MOXA-AWK4131A-MIB",
            "path": "moxa/MOXA-AWK4131A-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VR-BGP4-MIB",
            "path": "screenos/NETSCREEN-VR-BGP4-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-CFG-MIB",
            "path": "comware/HH3C-DOT11-CFG-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-VRRP3-MIB",
            "path": "nokia/ALCATEL-IND1-VRRP3-MIB.mib"
        },
        {
            "mib": "SWITCH",
            "path": "fs/centec/SWITCH.mib"
        },
        {
            "mib": "HUAWEI-IPSESSION-MIB",
            "path": "huawei/HUAWEI-IPSESSION-MIB.mib"
        },
        {
            "mib": "DPW-ATM-MIB",
            "path": "dasan/DPW-ATM-MIB.mib"
        },
        {
            "mib": "JUNIPER-CHASSIS-CLUSTER-MIB",
            "path": "junos/JUNIPER-CHASSIS-CLUSTER-MIB.mib"
        },
        {
            "mib": "WWP-LEOS-SYSTEM-CONFIG-MIB",
            "path": "ciena/CIENA-WWP-LEOS-SYSTEM-CONFIG.mib"
        },
        {
            "mib": "Juniper-BRIDGE-ETHERNET-MIB",
            "path": "junose/Juniper-BRIDGE-ETHERNET-MIB.mib"
        },
        {
            "mib": "CISCO-ENTITY-VENDORTYPE-OID-MIB",
            "path": "cisco/CISCO-ENTITY-VENDORTYPE-OID-MIB.mib"
        },
        {
            "mib": "BLADE-MIB",
            "path": "ibm/BLADE-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VR-MIB",
            "path": "screenos/NETSCREEN-VR-MIB.mib"
        },
        {
            "mib": "SPECTRACOM-PTPBASE-MIB",
            "path": "spectracom/SPECTRACOM-PTPBASE-MIB.mib"
        },
        {
            "mib": "ICF-VG-RPTR",
            "path": "hp/ICF-VG-RPTR.mib"
        },
        {
            "mib": "DLINKSW-LACP-EXT-MIB",
            "path": "dlink/DLINKSW-LACP-EXT-MIB.mib"
        },
        {
            "mib": "RFC1213-MIB",
            "path": "tplink/RFC1213-MIB.mib"
        },
        {
            "mib": "HUAWEI-IPV6-MIB",
            "path": "huawei/HUAWEI-IPV6-MIB.mib"
        },
        {
            "mib": "MOXA-DEVICE-IO-MIB",
            "path": "moxa/MOXA-DEVICE-IO-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-CFGEXT-MIB",
            "path": "comware/HH3C-DOT11-CFGEXT-MIB.mib"
        },
        {
            "mib": "BROCADE-NP-TM-STATS-MIB",
            "path": "brocade/BROCADE-NP-TM-STATS-MIB.mib"
        },
        {
            "mib": "TN-DEV-RFC2544-MIB",
            "path": "transition/TN-DEV-RFC2544-MIB.mib"
        },
        {
            "mib": "SLE-AM-MIB",
            "path": "dasan/SLE-AM-MIB.mib"
        },
        {
            "mib": "NETSWITCH-DMA-MIB",
            "path": "hp/NETSWITCH-DMA-MIB.mib"
        },
        {
            "mib": "Juniper-BRIDGE-MIB",
            "path": "junose/Juniper-BRIDGE-MIB.mib"
        },
        {
            "mib": "NETGEAR-DHCPSERVER-PRIVATE-MIB",
            "path": "quanta/fastpath_dhcp.my.mib"
        },
        {
            "mib": "DLINKSW-ENTITY-EXT-MIB",
            "path": "dlink_dgs1250/DLINKSW-ENTITY-EXT-MIB.mib"
        },
        {
            "mib": "TPLINK-DDMBIASCURTHRESHOLD-MIB",
            "path": "tplink/TPLINK-DDMBIASCURTHRESHOLD-MIB.mib"
        },
        {
            "mib": "WWP-PRODUCTS-MIB",
            "path": "ciena/CIENA-WWP-PRODUCTS-MIB.mib"
        },
        {
            "mib": "DLINKSW-LBD-MIB",
            "path": "dlink/DLINKSW-LBD-MIB.mib"
        },
        {
            "mib": "NETGEAR-DHCPCLIENT-PRIVATE-MIB",
            "path": "quanta/fastpath_dhcpclient.my.mib"
        },
        {
            "mib": "NETSWITCH-DRIVERS-MIB",
            "path": "hp/NETSWITCH-DRIVERS-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-WCCP-MIB",
            "path": "nokia/ALCATEL-IND1-WCCP-MIB.mib"
        },
        {
            "mib": "JUNIPER-CHASSIS-DEFINES-MIB",
            "path": "junos/JUNIPER-CHASSIS-DEFINES-MIB.mib"
        },
        {
            "mib": "MOXA-DUALHOMING-MIB",
            "path": "moxa/MOXA-DUALHOMING-MIB.mib"
        },
        {
            "mib": "SPECTRACOM-XSYNC-MIB",
            "path": "spectracom/SPECTRACOM-XSYNC-MIB.mib"
        },
        {
            "mib": "TN-DEV-SYS-IP2-MIB",
            "path": "transition/TN-DEV-SYS-IP2-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VR-OSPF-MIB",
            "path": "screenos/NETSCREEN-VR-OSPF-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-LIC-MIB",
            "path": "comware/HH3C-DOT11-LIC-MIB.mib"
        },
        {
            "mib": "TPLINK-DDMMANAGE-MIB",
            "path": "tplink/TPLINK-DDMMANAGE-MIB.mib"
        },
        {
            "mib": "BLADESPPALT-MIB",
            "path": "ibm/BLADESPPALT-MIB.mib"
        },
        {
            "mib": "DLINKSW-TC-MIB",
            "path": "dlink_dgs1250/DLINKSW-TC-MIB.mib"
        },
        {
            "mib": "SLE-BGP-MIB",
            "path": "dasan/SLE-BGP-MIB.mib"
        },
        {
            "mib": "BROCADE-PRODUCTS-MIB",
            "path": "brocade/BROCADE-PRODUCTS-MIB.mib"
        },
        {
            "mib": "DLINKSW-LED-MIB",
            "path": "dlink/DLINKSW-LED-MIB.mib"
        },
        {
            "mib": "Juniper-Bridged-Ethernet-CONF",
            "path": "junose/Juniper-Bridged-Ethernet-CONF.mib"
        },
        {
            "mib": "MOXA-EDS510E-MIB",
            "path": "moxa/MOXA-EDS510E-MIB.mib"
        },
        {
            "mib": "CISCO-ENVMON-MIB",
            "path": "cisco/CISCO-ENVMON-MIB.mib"
        },
        {
            "mib": "WWP-SMI",
            "path": "ciena/CIENA-WWP-SMI.mib"
        },
        {
            "mib": "PANDUIT-MIB",
            "path": "panduit/PANDUIT-MIB.mib"
        },
        {
            "mib": "HUAWEI-ISIS-CONF-MIB",
            "path": "huawei/HUAWEI-ISIS-CONF-MIB.mib"
        },
        {
            "mib": "NETSWITCH-MIB",
            "path": "hp/NETSWITCH-MIB.mib"
        },
        {
            "mib": "BTI8xx-INTERFACE-MIB",
            "path": "bti/BTI8xx-INTERFACE-MIB.mib"
        },
        {
            "mib": "NETGEAR-DENIALOFSERVICE-PRIVATE-MIB",
            "path": "quanta/fastpath_dos.my.mib"
        },
        {
            "mib": "DLINKSW-LLDP-EXT-MIB",
            "path": "dlink/DLINKSW-LLDP-EXT-MIB.mib"
        },
        {
            "mib": "CISCO-ERR-DISABLE-MIB",
            "path": "cisco/CISCO-ERR-DISABLE-MIB.mib"
        },
        {
            "mib": "BTI8xx-MIB",
            "path": "bti/BTI8xx-MIB.mib"
        },
        {
            "mib": "Brocade-REG-MIB",
            "path": "brocade/Brocade-REG-MIB.mib"
        },
        {
            "mib": "NETGEAR-KEYING-PRIVATE-MIB",
            "path": "quanta/fastpath_keying.my.mib"
        },
        {
            "mib": "JUNIPER-CHASSIS-FWDD-MIB",
            "path": "junos/JUNIPER-CHASSIS-FWDD-MIB.mib"
        },
        {
            "mib": "ALCATEL-IND1-WEBMGT-MIB",
            "path": "nokia/ALCATEL-IND1-WEBMGT-MIB.mib"
        },
        {
            "mib": "TPLINK-DDMRXPOWTHRESHOLD-MIB",
            "path": "tplink/TPLINK-DDMRXPOWTHRESHOLD-MIB.mib"
        },
        {
            "mib": "HUAWEI-KEYCHAIN-MIB",
            "path": "huawei/HUAWEI-KEYCHAIN-MIB.mib"
        },
        {
            "mib": "TN-DEV-SYS-IPMGMT-MIB",
            "path": "transition/TN-DEV-SYS-IPMGMT-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VR-RIPv2-MIB",
            "path": "screenos/NETSCREEN-VR-RIPv2-MIB.mib"
        },
        {
            "mib": "MOXA-EDS528E-MIB",
            "path": "moxa/MOXA-EDS528E-MIB.mib"
        },
        {
            "mib": "CISCO-ETHER-CFM-MIB",
            "path": "cisco/CISCO-ETHER-CFM-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-PROBE-MIB",
            "path": "comware/HH3C-DOT11-PROBE-MIB.mib"
        },
        {
            "mib": "ADIC-INTELLIGENT-STORAGE-MIB",
            "path": "adic/ADIC-INTELLIGENT-STORAGE-MIB.mib"
        },
        {
            "mib": "BROCADE-SYSLOG-MIB",
            "path": "brocade/BROCADE-SYSLOG-MIB.mib"
        },
        {
            "mib": "NETGEAR-LLPF-PRIVATE-MIB",
            "path": "quanta/fastpath_llpf.my.mib"
        },
        {
            "mib": "Juniper-Bridging-Manager-CONF",
            "path": "junose/Juniper-Bridging-Manager-CONF.mib"
        },
        {
            "mib": "TPLINK-DDMSTATUS-MIB",
            "path": "tplink/TPLINK-DDMSTATUS-MIB.mib"
        },
        {
            "mib": "Brocade-TC",
            "path": "brocade/Brocade-TC-MIB.mib"
        },
        {
            "mib": "BTI8xx-SFP-MIB",
            "path": "bti/BTI8xx-SFP-MIB.mib"
        },
        {
            "mib": "TN-DEV-SYS-SNMPMGMT-MIB",
            "path": "transition/TN-DEV-SYS-SNMPMGMT-MIB.mib"
        },
        {
            "mib": "HUAWEI-KOMPELLA-MIB",
            "path": "huawei/HUAWEI-KOMPELLA-MIB.mib"
        },
        {
            "mib": "CISCO-FIREWALL-MIB",
            "path": "cisco/CISCO-FIREWALL-MIB.mib"
        },
        {
            "mib": "DLINKSW-MAC-AUTH-MIB",
            "path": "dlink/DLINKSW-MAC-AUTH-MIB.mib"
        },
        {
            "mib": "ALCATEL-ISIS-MIB",
            "path": "nokia/ALCATEL-ISIS-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-QOS-MIB",
            "path": "comware/HH3C-DOT11-QOS-MIB.mib"
        },
        {
            "mib": "JUNIPER-COLLECTOR-MIB",
            "path": "junos/JUNIPER-COLLECTOR-MIB.mib"
        },
        {
            "mib": "ECS2100-MIB",
            "path": "edgecos/ECS2100-MIB.mib"
        },
        {
            "mib": "Juniper-Bridging-Manager-MIB",
            "path": "junose/Juniper-Bridging-Manager-MIB.mib"
        },
        {
            "mib": "NETSCREEN-VSYS-MIB",
            "path": "screenos/NETSCREEN-VSYS-MIB.mib"
        },
        {
            "mib": "MOXA-EDSG508E-MIB",
            "path": "moxa/MOXA-EDSG508E-MIB.mib"
        },
        {
            "mib": "BROCADE-VCS-MIB",
            "path": "brocade/BROCADE-VCS-MIB.mib"
        },
        {
            "mib": "DLINKSW-MGMD-SNOOPING-MIB",
            "path": "dlink/DLINKSW-MGMD-SNOOPING-MIB.mib"
        },
        {
            "mib": "ALCATEL-STATIC-FRR-MIB",
            "path": "nokia/ALCATEL-STATIC-FRR-MIB.mib"
        },
        {
            "mib": "SLE-BRIDGE-MIB",
            "path": "dasan/SLE-BRIDGE-MIB.mib"
        },
        {
            "mib": "HUAWEI-L2IF-MIB",
            "path": "huawei/HUAWEI-L2IF-MIB.mib"
        },
        {
            "mib": "MOXA-EDSG512E-MIB",
            "path": "moxa/MOXA-EDSG512E-MIB.mib"
        },
        {
            "mib": "CISCO-FLASH-MIB",
            "path": "cisco/CISCO-FLASH-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-REF-MIB",
            "path": "comware/HH3C-DOT11-REF-MIB.mib"
        },
        {
            "mib": "SLE-CLOCK-MIB",
            "path": "dasan/SLE-CLOCK-MIB.mib"
        },
        {
            "mib": "POWERSUPPLY-MIB",
            "path": "hp/POWERSUPPLY-MI.mib"
        },
        {
            "mib": "NETGEAR-MGMT-SECURITY-MIB",
            "path": "quanta/fastpath_mgmt_security.my.mib"
        },
        {
            "mib": "DLINKSW-ND-INSPECT-MIB",
            "path": "dlink/DLINKSW-ND-INSPECT-MIB.mib"
        },
        {
            "mib": "TN-DEV-SYS-UPGRADER-MIB",
            "path": "transition/TN-DEV-SYS-UPGRADER-MIB.mib"
        },
        {
            "mib": "FA-EXT-MIB",
            "path": "brocade/FA-EXT-MIB.mib"
        },
        {
            "mib": "NETSCREEN-ZONE-MIB",
            "path": "screenos/NETSCREEN-ZONE-MIB.mib"
        },
        {
            "mib": "Juniper-CLI-CONF",
            "path": "junose/Juniper-CLI-CONF.mib"
        },
        {
            "mib": "BTI8xx-SYSTEM-MIB",
            "path": "bti/BTI8xx-SYSTEM-MIB.mib"
        },
        {
            "mib": "ADIC-MANAGEMENT-MIB",
            "path": "adic/ADIC-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "SEMI-MIB",
            "path": "hp/SEMI-MIB.mib"
        },
        {
            "mib": "JUNIPER-COS-MIB",
            "path": "junos/JUNIPER-COS-MIB.mib"
        },
        {
            "mib": "DLINKSW-NETWORK-ACCESS-MIB",
            "path": "dlink/DLINKSW-NETWORK-ACCESS-MIB.mib"
        },
        {
            "mib": "ECS3510-MIB",
            "path": "edgecos/ECS3510-MIB.mib"
        },
        {
            "mib": "CISCO-FRAME-RELAY-MIB",
            "path": "cisco/CISCO-FRAME-RELAY-MIB.mib"
        },
        {
            "mib": "NETGEAR-VPC-MIB",
            "path": "quanta/fastpath_mlag.my.mib"
        },
        {
            "mib": "HH3C-DOT11-ROAM-MIB",
            "path": "comware/HH3C-DOT11-ROAM-MIB.mib"
        },
        {
            "mib": "TPLINK-DDMTEMPTHRESHOLD-MIB",
            "path": "tplink/TPLINK-DDMTEMPTHRESHOLD-MIB.mib"
        },
        {
            "mib": "HUAWEI-L2MAM-MIB",
            "path": "huawei/HUAWEI-L2MAM-MIB.mib"
        },
        {
            "mib": "SLE-CONFIG-MIB",
            "path": "dasan/SLE-CONFIG-MIB.mib"
        },
        {
            "mib": "JUNIPER-DCU-MIB",
            "path": "junos/JUNIPER-DCU-MIB.mib"
        },
        {
            "mib": "ADIC-SANMGR-PROXY-MIB",
            "path": "adic/ADIC-SANMGR-PROXY-MIB.mib"
        },
        {
            "mib": "FCMGMT-MIB",
            "path": "brocade/FCMGMT-MIB.mib"
        },
        {
            "mib": "CME-MIB",
            "path": "ibm/CME-MIB.mib"
        },
        {
            "mib": "ALU-MICROWAVE-MIB",
            "path": "nokia/ALU-MICROWAVE-MIB.mib"
        },
        {
            "mib": "Juniper-CLI-MIB",
            "path": "junose/Juniper-CLI-MIB.mib"
        },
        {
            "mib": "B100-MIB",
            "path": "kemp/B100-MIB.mib"
        },
        {
            "mib": "STATISTICS-MIB",
            "path": "hp/STATISTICS-MIB.mib"
        },
        {
            "mib": "BTI8xx-TC-MIB",
            "path": "bti/BTI8xx-TC-MIB.mib"
        },
        {
            "mib": "TN-DEV-SYS-USER-MIB",
            "path": "transition/TN-DEV-SYS-USER-MIB.mib"
        },
        {
            "mib": "CISCO-HSRP-EXT-MIB",
            "path": "cisco/CISCO-HSRP-EXT-MIB.mib"
        },
        {
            "mib": "Juniper-COPS-CONF",
            "path": "junose/Juniper-COPS-CONF.mib"
        },
        {
            "mib": "HUAWEI-L2MULTICAST-MIB",
            "path": "huawei/HUAWEI-L2MULTICAST-MIB.mib"
        },
        {
            "mib": "FIBRE-CHANNEL-FE-MIB",
            "path": "brocade/FIBRE-CHANNEL-FE-MIB.mib"
        },
        {
            "mib": "MOXA-EDSG512E8POE-MIB",
            "path": "moxa/MOXA-EDSG512E8POE-MIB.mib"
        },
        {
            "mib": "TPLINK-DDMTXPOWTHRESHOLD-MIB",
            "path": "tplink/TPLINK-DDMTXPOWTHRESHOLD-MIB.mib"
        },
        {
            "mib": "NETGEAR-MVR-PRIVATE-MIB",
            "path": "quanta/fastpath_mvr.my.mib"
        },
        {
            "mib": "HH3C-DOT11-RRM-MIB",
            "path": "comware/HH3C-DOT11-RRM-MIB.mib"
        },
        {
            "mib": "DLINKSW-NETWORK-PROTOCOL-PORT-PROTECT-MIB",
            "path": "dlink/DLINKSW-NETWORK-PROTOCOL-PORT-PROTECT-MIB.mib"
        },
        {
            "mib": "TN-DEV-SYS-XNTP-MIB",
            "path": "transition/TN-DEV-SYS-xNTP-MIB.mib"
        },
        {
            "mib": "ECS4100-52T-MIB",
            "path": "edgecos/ECS4100-52T-MIB.mib"
        },
        {
            "mib": "SLE-DCN-MIB",
            "path": "dasan/SLE-DCN-MIB.mib"
        },
        {
            "mib": "CISCO-HSRP-MIB",
            "path": "cisco/CISCO-HSRP-MIB.mib"
        },
        {
            "mib": "BroadworksMaintenance",
            "path": "broadworks/BroadworksMaintenance.mib"
        },
        {
            "mib": "BARCO-CLICKSHARE-MIB",
            "path": "barco/BARCO-CLICKSHARE-MIB.mib"
        },
        {
            "mib": "GPFS-MIB",
            "path": "ibm/GPFS-MIB.mib"
        },
        {
            "mib": "JUNIPER-DFC-MIB",
            "path": "junos/JUNIPER-DFC-MIB.mib"
        },
        {
            "mib": "DIGIPDU-MIB",
            "path": "planet/DIGIPDU-MIB.mib"
        },
        {
            "mib": "HUAWEI-L2TP-EXT-MIB",
            "path": "huawei/HUAWEI-L2TP-EXT-MIB.mib"
        },
        {
            "mib": "ALU-SAR-GLOBAL-MIB",
            "path": "nokia/ALU-SAR-GLOBAL-MIB.mib"
        },
        {
            "mib": "IPVS-MIB",
            "path": "kemp/IPVS-MIB.mib"
        },
        {
            "mib": "NETGEAR-PFC-MIB",
            "path": "quanta/fastpath_pfc.my.mib"
        },
        {
            "mib": "HH3C-DOT11-SA-MIB",
            "path": "comware/HH3C-DOT11-SA-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-DHCP-SERVER-MIB",
            "path": "cisco/CISCO-IETF-DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "ASAM-ALARM-MIB",
            "path": "nokia/ASAM-ALARM-MIB.mib"
        },
        {
            "mib": "MOXA-EDSG512EPoE-MIB",
            "path": "moxa/MOXA-EDSG512EPoE-MIB.mib"
        },
        {
            "mib": "FOUNDRY-BFD-STD-MIB",
            "path": "brocade/FOUNDRY-BFD-STD-MIB.mib"
        },
        {
            "mib": "TPLINK-DDMVOLTHRESHOLD-MIB",
            "path": "tplink/TPLINK-DDMVOLTHRESHOLD-MIB.mib"
        },
        {
            "mib": "TN-DEV-VLAN-TRANSLATION-MIB",
            "path": "transition/TN-DEV-VLAN-TRANSLATION-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-SAVI-MIB",
            "path": "comware/HH3C-DOT11-SAVI-MIB.mib"
        },
        {
            "mib": "SLE-DEBUG-MIB",
            "path": "dasan/SLE-DEBUG-MIB.mib"
        },
        {
            "mib": "JUNIPER-DOM-MIB",
            "path": "junos/JUNIPER-DOM-MIB.mib"
        },
        {
            "mib": "ONE4NET-MIB",
            "path": "kemp/ONE4NET-MIB.mib"
        },
        {
            "mib": "VUTLAN-SYSTEM-MIB",
            "path": "vutlan/VUTLAN-SYSTEM-MIB.mib"
        },
        {
            "mib": "SYNOLOGY-DISK-MIB",
            "path": "synology/SYNOLOGY-DISK-MIB.mib"
        },
        {
            "mib": "IBM-3200-MIB",
            "path": "ibm/IBM-3200-MIB.mib"
        },
        {
            "mib": "DLINKSW-NTP-MIB",
            "path": "dlink/DLINKSW-NTP-MIB.mib"
        },
        {
            "mib": "ECS4110-MIB",
            "path": "edgecos/ECS4110-MIB.mib"
        },
        {
            "mib": "HUAWEI-L2VLAN-MIB",
            "path": "huawei/HUAWEI-L2VLAN-MIB.mib"
        },
        {
            "mib": "NETGEAR-PORTSECURITY-PRIVATE-MIB",
            "path": "quanta/fastpath_portsecurity.my.mib"
        },
        {
            "mib": "ASAM-EQUIP-MIB",
            "path": "nokia/ASAM-EQUIP-MIB.mib"
        },
        {
            "mib": "TPLINK-DOT1Q-VLAN-MIB",
            "path": "tplink/TPLINK-DOT1Q-VLAN-MIB.mib"
        },
        {
            "mib": "MOXA-EDSG516E-MIB",
            "path": "moxa/MOXA-EDSG516E-MIB.mib"
        },
        {
            "mib": "DLINKSW-PACKET-MONITOR-MIB",
            "path": "dlink/DLINKSW-PACKET-MONITOR-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-ISIS-MIB",
            "path": "cisco/CISCO-IETF-ISIS-MIB.mib"
        },
        {
            "mib": "BW-BroadworksApplicationServer",
            "path": "broadworks/BW-BroadworksApplicationServer.mib"
        },
        {
            "mib": "FOUNDRY-CAR-MIB",
            "path": "brocade/FOUNDRY-CAR-MIB.mib"
        },
        {
            "mib": "HALON-SP-MIB",
            "path": "halon/HALON-SP-MIB.mib"
        },
        {
            "mib": "TN-DHCP-MIB",
            "path": "transition/TN-DHCP-MIB.mib"
        },
        {
            "mib": "SLE-DEVICE-MIB",
            "path": "dasan/SLE-DEVICE-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-STATION-MIB",
            "path": "comware/HH3C-DOT11-STATION-MIB.mib"
        },
        {
            "mib": "SYNOLOGY-RAID-MIB",
            "path": "synology/SYNOLOGY-RAID-MIB.mib"
        },
        {
            "mib": "HUAWEI-L2VPN-MIB",
            "path": "huawei/HUAWEI-L2VPN-MIB.mib"
        },
        {
            "mib": "TPLINK-IPV6ADDR-MIB",
            "path": "tplink/TPLINK-IPV6ADDR-MIB.mib"
        },
        {
            "mib": "ASAM-SYSTEM-MIB",
            "path": "nokia/ASAM-SYSTEM-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-ALARM-MIB",
            "path": "arris/d5/ARRIS-D5-ALARM-MIB.mib"
        },
        {
            "mib": "JUNIPER-EVENT-MIB",
            "path": "junos/JUNIPER-EVENT-MIB.mib"
        },
        {
            "mib": "NETGEAR-POWER-ETHERNET-MIB",
            "path": "quanta/fastpath_power_ethernet.my.mib"
        },
        {
            "mib": "ECS4120-MIB",
            "path": "edgecos/ECS4120-MIB.mib"
        },
        {
            "mib": "SLE-DHCP-MIB",
            "path": "dasan/SLE-DHCP-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-PW-ENET-MIB",
            "path": "cisco/CISCO-IETF-PW-ENET-MIB.mib"
        },
        {
            "mib": "BW-BroadworksMediaServer",
            "path": "broadworks/BW-BroadworksMediaServer.mib"
        },
        {
            "mib": "TN-DHCP-RELAY-MIB",
            "path": "transition/TN-DHCP-RELAY-MIB.mib"
        },
        {
            "mib": "SYNOLOGY-SYSTEM-MIB",
            "path": "synology/SYNOLOGY-SYSTEM-MIB.mib"
        },
        {
            "mib": "HUAWEI-L3VLAN-MIB",
            "path": "huawei/HUAWEI-L3VLAN-MIB.mib"
        },
        {
            "mib": "MOXA-EDSP506E-MIB",
            "path": "moxa/MOXA-EDSP506E-MIB.mib"
        },
        {
            "mib": "TPLINK-IPV6STATICROUTE-MIB",
            "path": "tplink/TPLINK-IPV6STATICROUTE-MIB.mib"
        },
        {
            "mib": "FOUNDRY-LAG-MIB",
            "path": "brocade/FOUNDRY-LAG-MIB.mib"
        },
        {
            "mib": "ASAM-TC-MIB",
            "path": "nokia/ASAM-TC-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-CHANNEL-MODE-LOG-MIB",
            "path": "arris/d5/ARRIS-D5-CHANNEL-MODE-LOG-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-PW-FR-MIB",
            "path": "cisco/CISCO-IETF-PW-FR-MIB.mib"
        },
        {
            "mib": "DLINKSW-POE-MIB",
            "path": "dlink/DLINKSW-POE-MIB.mib"
        },
        {
            "mib": "Juniper-COPS-MIB",
            "path": "junose/Juniper-COPS-MIB.mib"
        },
        {
            "mib": "TN-DHCP-SERVER-MIB",
            "path": "transition/TN-DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "INSYDE-IPMI-MIB",
            "path": "supervyse-openbmc/INSYDE-IPMI-MIB.mib"
        },
        {
            "mib": "TPLINK-LLDP-MIB",
            "path": "tplink/TPLINK-LLDP-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-WIDS-MIB",
            "path": "comware/HH3C-DOT11-WIDS-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-DTI-EXT-MIB",
            "path": "arris/d5/ARRIS-D5-DTI-EXT-MIB.mib"
        },
        {
            "mib": "SLE-DHCP-SNOOPING-MIB",
            "path": "dasan/SLE-DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "BW-BroadworksNetworkServer",
            "path": "broadworks/BW-BroadworksNetworkServer.mib"
        },
        {
            "mib": "NETGEAR-QOS-COS-MIB",
            "path": "quanta/fastpath_qos_cos.my.mib"
        },
        {
            "mib": "MOXA-EDSP510A8POE-MIB",
            "path": "moxa/MOXA-EDSP510A8POE-MIB.mib"
        },
        {
            "mib": "TN-DHCP-SNOOPING-MIB",
            "path": "transition/TN-DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "TPLINK-LLDPINFO-MIB",
            "path": "tplink/TPLINK-LLDPINFO-MIB.mib"
        },
        {
            "mib": "HUAWEI-L3VPN-EXT-MIB",
            "path": "huawei/HUAWEI-L3VPN-EXT-MIB.mib"
        },
        {
            "mib": "IBM-6611-APPN-MIB",
            "path": "ibm/IBM-6611-APPN-MIB.mib"
        },
        {
            "mib": "DLINKSW-PORT-SECURITY-MIB",
            "path": "dlink/DLINKSW-PORT-SECURITY-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-PW-MIB",
            "path": "cisco/CISCO-IETF-PW-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-WIPS-MIB",
            "path": "comware/HH3C-DOT11-WIPS-MIB.mib"
        },
        {
            "mib": "ECS4210-MIB",
            "path": "edgecos/ECS4210-MIB.mib"
        },
        {
            "mib": "IMP-HF528-MIB",
            "path": "tailyn/IMP-HF528-MIB.mib"
        },
        {
            "mib": "IPMCAST-MIB",
            "path": "nokia/IPMCAST-MIB.mib"
        },
        {
            "mib": "FOUNDRY-MAC-VLAN-MIB",
            "path": "brocade/FOUNDRY-MAC-VLAN-MIB.mib"
        },
        {
            "mib": "Juniper-DHCP-CONF",
            "path": "junose/Juniper-DHCP-CONF.mib"
        },
        {
            "mib": "TN-ELPS-MIB",
            "path": "transition/TN-ELPS-MIB.mib"
        },
        {
            "mib": "BIANCA-BRICK-MIB",
            "path": "bintec/BIANCA-BRICK-MIB.mib"
        },
        {
            "mib": "ZHONE-CARD-RESOURCES-MIB",
            "path": "zhone/ZHONE-CARD-RESOURCES-MIB.mib"
        },
        {
            "mib": "HUAWEI-LDT-MIB",
            "path": "huawei/HUAWEI-LDT-MIB.mib"
        },
        {
            "mib": "TPLINK-MIB",
            "path": "tplink/TPLINK-MIB.mib"
        },
        {
            "mib": "NETGEAR-QOS-ISCSI-MIB",
            "path": "quanta/fastpath_qos_iscsi.my.mib"
        },
        {
            "mib": "JUNIPER-EX-MAC-NOTIFICATION-MIB",
            "path": "junos/JUNIPER-EX-MAC-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-PW-MPLS-MIB",
            "path": "cisco/CISCO-IETF-PW-MPLS-MIB.mib"
        },
        {
            "mib": "IPX",
            "path": "nokia/IPX.mib"
        },
        {
            "mib": "MOXA-FIBER-CHECK-MIB",
            "path": "moxa/MOXA-FIBER-CHECK-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-DVB-EIS-MIB",
            "path": "arris/d5/ARRIS-D5-DVB-EIS-MIB.mib"
        },
        {
            "mib": "NETGEAR-QOS-AUTOVOIP-MIB",
            "path": "quanta/fastpath_qos_voip.my.mib"
        },
        {
            "mib": "BIANCA-BRICK-MIBRES-MIB",
            "path": "bintec/BIANCA-BRICK-MIBRES-MIB.mib"
        },
        {
            "mib": "SLE-DHCPV6-MIB",
            "path": "dasan/SLE-DHCPV6-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-WLANEXT-MIB",
            "path": "comware/HH3C-DOT11-WLANEXT-MIB.mib"
        },
        {
            "mib": "TPLINK-POWER-OVER-ETHERNET-MIB",
            "path": "tplink/TPLINK-POWER-OVER-ETHERNET-MIB.mib"
        },
        {
            "mib": "TN-ENTITY-SENSOR-MIB",
            "path": "transition/TN-ENTITY-SENSOR-MIB.mib"
        },
        {
            "mib": "FOUNDRY-POE-MIB",
            "path": "brocade/FOUNDRY-POE-MIB.mib"
        },
        {
            "mib": "JUNIPER-EX-SMI",
            "path": "junos/JUNIPER-EX-SMI.mib"
        },
        {
            "mib": "ZHONE-INTERFACE-TRANSLATION-MIB",
            "path": "zhone/ZHONE-INTERFACE-TRANSLATION-MIB.mib"
        },
        {
            "mib": "PDU-MIB",
            "path": "raritan/PDU-MIB.mib"
        },
        {
            "mib": "DLINKSW-POWER-SAVING-MIB",
            "path": "dlink/DLINKSW-POWER-SAVING-MIB.mib"
        },
        {
            "mib": "MOXA-GENERAL-MIB",
            "path": "moxa/MOXA-GENERAL-MIB.mib"
        },
        {
            "mib": "CISCO-IETF-PW-TC-MIB",
            "path": "cisco/CISCO-IETF-PW-TC-MIB.mib"
        },
        {
            "mib": "HUAWEI-LI-MIB",
            "path": "huawei/HUAWEI-LI-MIB.mib"
        },
        {
            "mib": "NETGEAR-OUTBOUNDTELNET-PRIVATE-MIB",
            "path": "quanta/fastpath_telnet.my.mib"
        },
        {
            "mib": "BINTEC-MIB",
            "path": "bintec/BINTEC-MIB.mib"
        },
        {
            "mib": "Juniper-DHCP-MIB",
            "path": "junose/Juniper-DHCP-MIB.mib"
        },
        {
            "mib": "ECS4510-MIB",
            "path": "edgecos/ECS4510-MIB.mib"
        },
        {
            "mib": "JUNIPER-EXPERIMENT-MIB",
            "path": "junos/JUNIPER-EXPERIMENT-MIB.mib"
        },
        {
            "mib": "IBM-AIX-MIB",
            "path": "ibm/IBM-AIX-MIB.mib"
        },
        {
            "mib": "TPLINK-PRODUCTS-MIB",
            "path": "tplink/TPLINK-PRODUCTS-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11-WM2U-MIB",
            "path": "comware/HH3C-DOT11-WM2U-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-DVB-PSIG-MIB",
            "path": "arris/d5/ARRIS-D5-DVB-PSIG-MIB.mib"
        },
        {
            "mib": "CISCO-IF-EXTENSION-MIB",
            "path": "cisco/CISCO-IF-EXTENSION-MIB.mib"
        },
        {
            "mib": "RFC1158-MIB",
            "path": "bintec/RFC1158-MIB.mib"
        },
        {
            "mib": "Juniper-DHCPv6-CONF",
            "path": "junose/Juniper-DHCPv6-CONF.mib"
        },
        {
            "mib": "TN-ERPS-MIB",
            "path": "transition/TN-ERPS-MIB.mib"
        },
        {
            "mib": "HUAWEI-LINE-MIB",
            "path": "huawei/HUAWEI-LINE-MIB.mib"
        },
        {
            "mib": "SLE-EPON-MIB",
            "path": "dasan/SLE-EPON-MIB.mib"
        },
        {
            "mib": "MOXA-IKS6726A-MIB",
            "path": "moxa/MOXA-IKS6726A-MIB.mib"
        },
        {
            "mib": "ISIS-MIB",
            "path": "nokia/ISIS-MIB.mib"
        },
        {
            "mib": "ECS4610-24F-MIB",
            "path": "edgecos/ECS4610-24F-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-AGENT-MIB",
            "path": "brocade/FOUNDRY-SN-AGENT-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-ENT-CHASSIS-MIB",
            "path": "arris/d5/ARRIS-D5-ENT-CHASSIS-MIB.mib"
        },
        {
            "mib": "PDU2-MIB",
            "path": "raritan/PDU2-MIB.mib"
        },
        {
            "mib": "Zhone",
            "path": "zhone/ZHONE-MIB.mib"
        },
        {
            "mib": "CISCO-IF-THRESHOLD-MIB",
            "path": "cisco/CISCO-IF-THRESHOLD-MIB.mib"
        },
        {
            "mib": "HH3C-DOT11S-MESH-MIB",
            "path": "comware/HH3C-DOT11S-MESH-MIB.mib"
        },
        {
            "mib": "DLINKSW-QOS-MIB",
            "path": "dlink/DLINKSW-QOS-MIB.mib"
        },
        {
            "mib": "RemoteKVMDevice-MIB",
            "path": "raritan/RemoteKVMDevice-MIB.mib"
        },
        {
            "mib": "JUNIPER-FABRIC-CHASSIS",
            "path": "junos/JUNIPER-FABRIC-CHASSIS.mib"
        },
        {
            "mib": "TPLINK-STATICROUTE-MIB",
            "path": "tplink/TPLINK-STATICROUTE-MIB.mib"
        },
        {
            "mib": "NETGEAR-TIMERANGE-MIB",
            "path": "quanta/fastpath_timerange.my.mib"
        },
        {
            "mib": "IBM-BCM-MIB",
            "path": "ibm/IBM-BCM-MIB.mib"
        },
        {
            "mib": "CISCO-IGMP-SNOOPING-MIB",
            "path": "cisco/CISCO-IGMP-SNOOPING-MIB.mib"
        },
        {
            "mib": "Zhone-TC",
            "path": "zhone/ZHONE-TC-MIB.mib"
        },
        {
            "mib": "HH3C-DOT3-EFM-EPON-MIB",
            "path": "comware/HH3C-DOT3-EFM-EPON-MIB.mib"
        },
        {
            "mib": "ITF-MIB-EXT",
            "path": "nokia/ITF-MIB-EXT.mib"
        },
        {
            "mib": "TPLINK-SYSINFO-MIB",
            "path": "tplink/TPLINK-SYSINFO-MIB.mib"
        },
        {
            "mib": "TN-ETHSOAM-MIB",
            "path": "transition/TN-ETHSOAM-MIB.mib"
        },
        {
            "mib": "NETGEAR-TIMEZONE-PRIVATE-MIB",
            "path": "quanta/fastpath_timezone.my.mib"
        },
        {
            "mib": "ES3510MA-MIB",
            "path": "edgecos/ES3510MA-MIB.mib"
        },
        {
            "mib": "A3COM-HUAWEI-DEVICE-MIB",
            "path": "3com/A3COM-HUAWEI-DEVICE-MIB.mib"
        },
        {
            "mib": "ALCATEL-NGOAW-BASE-MIB",
            "path": "nokia/stellar/ALCATEL-NGOAW-BASE-MIB.mib"
        },
        {
            "mib": "MOXA-IKS6728A-8POE-MIB",
            "path": "moxa/MOXA-IKS6728A-8POE-MIB.mib"
        },
        {
            "mib": "HUAWEI-LLDP-MIB",
            "path": "huawei/HUAWEI-LLDP-MIB.mib"
        },
        {
            "mib": "SLE-FAULTMGMT-MIB",
            "path": "dasan/SLE-FAULTMGMT-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-ENT-SENSOR-THRESH-MIB",
            "path": "arris/d5/ARRIS-D5-ENT-SENSOR-THRESH-MIB.mib"
        },
        {
            "mib": "Juniper-DHCPv6-MIB",
            "path": "junose/Juniper-DHCPv6-MIB.mib"
        },
        {
            "mib": "BISON-ROUTER-MIB",
            "path": "bison/BISON-ROUTER-MIB.mib"
        },
        {
            "mib": "JUNIPER-FABRIC-MIB",
            "path": "junos/JUNIPER-FABRIC-MIB.mib"
        },
        {
            "mib": "DLINKSW-RA-GUARD-MIB",
            "path": "dlink/DLINKSW-RA-GUARD-MIB.mib"
        },
        {
            "mib": "CISCO-IMAGE-MIB",
            "path": "cisco/CISCO-IMAGE-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-APPLETALK-MIB",
            "path": "brocade/FOUNDRY-SN-APPLETALK-MIB.mib"
        },
        {
            "mib": "NETGEAR-UDLD-MIB",
            "path": "quanta/fastpath_udld.my.mib"
        },
        {
            "mib": "HH3C-DRNI-MIB",
            "path": "comware/HH3C-DRNI-MIB.mib"
        },
        {
            "mib": "TPLINK-SYSMONITOR-MIB",
            "path": "tplink/TPLINK-SYSMONITOR-MIB.mib"
        },
        {
            "mib": "IBM-CPS-MIB",
            "path": "ibm/IBM-CPS-MIB.mib"
        },
        {
            "mib": "TN-ETHSOAM-PM-MIB",
            "path": "transition/TN-ETHSOAM-PM-MIB.mib"
        },
        {
            "mib": "OA-SMI",
            "path": "openaccess/OA-SMI.mib"
        },
        {
            "mib": "Juniper-DISMAN-EVENT-MIB",
            "path": "junose/Juniper-DISMAN-EVENT-MIB.mib"
        },
        {
            "mib": "ITF-MIB",
            "path": "nokia/ITF-MIB.mib"
        },
        {
            "mib": "MOXA-IKS6728A-MIB",
            "path": "moxa/MOXA-IKS6728A-MIB.mib"
        },
        {
            "mib": "ALCATEL-NGOAW-DEVICES-MIB",
            "path": "nokia/stellar/ALCATEL-NGOAW-DEVICES-MIB.mib"
        },
        {
            "mib": "DLINKSW-SAFEGUARD-ENGINE-MIB",
            "path": "dlink/DLINKSW-SAFEGUARD-ENGINE-MIB.mib"
        },
        {
            "mib": "CISCO-IMAGE-TC",
            "path": "cisco/CISCO-IMAGE-TC.mib"
        },
        {
            "mib": "JUNIPER-FIREWALL-MIB",
            "path": "junos/JUNIPER-FIREWALL-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-IP-MIB",
            "path": "arris/d5/ARRIS-D5-IP-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-ARP-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-ARP-GROUP-MIB.mib"
        },
        {
            "mib": "A3COM-HUAWEI-LswDEVM-MIB",
            "path": "3com/A3COM-HUAWEI-LswDEVM-MIB.mib"
        },
        {
            "mib": "TPLINK-TC-MIB",
            "path": "tplink/TPLINK-TC-MIB.mib"
        },
        {
            "mib": "HH3C-DSP-MIB",
            "path": "comware/HH3C-DSP-MIB.mib"
        },
        {
            "mib": "HUAWEI-LOAD-BACKUP-MIB",
            "path": "huawei/HUAWEI-LOAD-BACKUP-MIB.mib"
        },
        {
            "mib": "JUNIPER-FRU-MIB",
            "path": "junos/JUNIPER-FRU-MIB.mib"
        },
        {
            "mib": "ES3528MO-MIB",
            "path": "edgecos/ES3528MO-MIB.mib"
        },
        {
            "mib": "A3COM-HUAWEI-OID-MIB",
            "path": "3com/A3COM-HUAWEI-OID-MIB.mib"
        },
        {
            "mib": "NETGEAR-FIPSNOOPING-MIB",
            "path": "quanta/fip_snooping.my.mib"
        },
        {
            "mib": "OACOMMON-MIB",
            "path": "openaccess/OACOMMON-MIB.mib"
        },
        {
            "mib": "LANGTAG-TC-MIB",
            "path": "nokia/LANGTAG-TC-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-LICENSE-MIB",
            "path": "arris/d5/ARRIS-D5-LICENSE-MIB.mib"
        },
        {
            "mib": "TN-EVC-MIB",
            "path": "transition/TN-EVC-MIB.mib"
        },
        {
            "mib": "MOXA-POE-BT-MIB",
            "path": "moxa/MOXA-POE-BT-MIB.mib"
        },
        {
            "mib": "DLINKSW-SFLOW-MIB",
            "path": "dlink/DLINKSW-SFLOW-MIB.mib"
        },
        {
            "mib": "HH3C-DVPN-MIB",
            "path": "comware/HH3C-DVPN-MIB.mib"
        },
        {
            "mib": "HUAWEI-LOOPDETECT-MIB",
            "path": "huawei/HUAWEI-LOOPDETECT-MIB.mib"
        },
        {
            "mib": "CISCO-IP-STAT-MIB",
            "path": "cisco/CISCO-IP-STAT-MIB.mib"
        },
        {
            "mib": "DeltaUPS-MIB",
            "path": "delta/DeltaUPS-MIB.mib"
        },
        {
            "mib": "OAW-AP1101",
            "path": "nokia/stellar/OAW-AP1101.mib"
        },
        {
            "mib": "TN-FRA-MIB",
            "path": "transition/TN-FRA-MIB.mib"
        },
        {
            "mib": "TEL2N-MIB",
            "path": "2n/TEL2N-MIB.mib"
        },
        {
            "mib": "MPLS-LDP-MIB",
            "path": "nokia/MPLS-LDP-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-BGP4-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-BGP4-GROUP-MIB.mib"
        },
        {
            "mib": "ES3528MV2-MIB",
            "path": "edgecos/ES3528MV2-MIB.mib"
        },
        {
            "mib": "NETGEAR-GREENETHERNET-PRIVATE-MIB",
            "path": "quanta/green_ethernet.my.mib"
        },
        {
            "mib": "ARRIS-D5-QAM-EXT-MIB",
            "path": "arris/d5/ARRIS-D5-QAM-EXT-MIB.mib"
        },
        {
            "mib": "Juniper-DNS-CONF",
            "path": "junose/Juniper-DNS-CONF.mib"
        },
        {
            "mib": "SLE-GPON-MIB",
            "path": "dasan/SLE-GPON-MIB.mib"
        },
        {
            "mib": "JUNIPER-HOSTRESOURCES-MIB",
            "path": "junos/JUNIPER-HOSTRESOURCES-MIB.mib"
        },
        {
            "mib": "HH3C-E1-MIB",
            "path": "comware/HH3C-E1-MIB.mib"
        },
        {
            "mib": "A3Com-products-MIB",
            "path": "3com/A3Com-products-MIB.mib"
        },
        {
            "mib": "OAW-AP1201",
            "path": "nokia/stellar/OAW-AP1201.mib"
        },
        {
            "mib": "IBM-Director-Alert-MIB",
            "path": "ibm/IBM-Director-Alert-MIB.mib"
        },
        {
            "mib": "MOXA-PORT-MIB",
            "path": "moxa/MOXA-PORT-MIB.mib"
        },
        {
            "mib": "PIM-BSR-MIB",
            "path": "nokia/PIM-BSR-MIB.mib"
        },
        {
            "mib": "HUAWEI-M-LAG-MIB",
            "path": "huawei/HUAWEI-M-LAG-MIB.mib"
        },
        {
            "mib": "SLE-ISIS-MIB",
            "path": "dasan/SLE-ISIS-MIB.mib"
        },
        {
            "mib": "GLOBAL-REG",
            "path": "delta/GLOBAL-REG.mib"
        },
        {
            "mib": "DLINKSW-SNMP-MIB",
            "path": "dlink/DLINKSW-SNMP-MIB.mib"
        },
        {
            "mib": "OAW-AP1201BG",
            "path": "nokia/stellar/OAW-AP1201BG.mib"
        },
        {
            "mib": "A3COM0004-GENERIC",
            "path": "3com/A3COM0004-GENERIC.mib"
        },
        {
            "mib": "ATS-MIB",
            "path": "ats/ATS-MIB.mib"
        },
        {
            "mib": "TN-HQOS-MIB",
            "path": "transition/TN-HQOS-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-QAM-MIB",
            "path": "arris/d5/ARRIS-D5-QAM-MIB.mib"
        },
        {
            "mib": "TAIT-COMMON-MIB",
            "path": "tait/TAIT-COMMON-MIB.mib"
        },
        {
            "mib": "IEEE8021-PFC-MIB",
            "path": "quanta/ieee8021_pfc.my.mib"
        },
        {
            "mib": "JUNIPER-IF-ACCOUNTING-MIB",
            "path": "junos/JUNIPER-IF-ACCOUNTING-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-CAM-MIB",
            "path": "brocade/FOUNDRY-SN-CAM-MIB.mib"
        },
        {
            "mib": "Juniper-DNS-MIB",
            "path": "junose/Juniper-DNS-MIB.mib"
        },
        {
            "mib": "CISCO-IP-URPF-MIB",
            "path": "cisco/CISCO-IP-URPF-MIB.mib"
        },
        {
            "mib": "DLINKSW-SSH-MIB",
            "path": "dlink/DLINKSW-SSH-MIB.mib"
        },
        {
            "mib": "SLE-MLSQOS-MIB",
            "path": "dasan/SLE-MLSQOS-MIB.mib"
        },
        {
            "mib": "OAW-AP1201H",
            "path": "nokia/stellar/OAW-AP1201H.mib"
        },
        {
            "mib": "HH3C-E1T1VI-MIB",
            "path": "comware/HH3C-E1T1VI-MIB.mib"
        },
        {
            "mib": "PIM-STD-MIB",
            "path": "nokia/PIM-STD-MIB.mib"
        },
        {
            "mib": "TN-HTTPS-MIB",
            "path": "transition/TN-HTTPS-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-QAM-POST-MIB",
            "path": "arris/d5/ARRIS-D5-QAM-POST-MIB.mib"
        },
        {
            "mib": "ORION-BASE-MIB",
            "path": "delta/ORION-BASE-MIB.mib"
        },
        {
            "mib": "MOXA-PT7528V2-MIB",
            "path": "moxa/MOXA-PT7528V2-MIB.mib"
        },
        {
            "mib": "A3COM0352-STACK-CONFIG",
            "path": "3com/A3COM0352-STACK-CONFIG.mib"
        },
        {
            "mib": "MGMD-STD-MIB",
            "path": "quanta/mgmd.my.mib"
        },
        {
            "mib": "HH3C-EFM-COMMON-MIB",
            "path": "comware/HH3C-EFM-COMMON-MIB.mib"
        },
        {
            "mib": "CISCO-IPSEC-FLOW-MONITOR-MIB",
            "path": "cisco/CISCO-IPSEC-FLOW-MONITOR-MIB.mib"
        },
        {
            "mib": "TPDIN2-MIB",
            "path": "tycon/TPDIN2-MIB.mib"
        },
        {
            "mib": "MOXA-SWITCHING-MIB",
            "path": "moxa/MOXA-SWITCHING-MIB.mib"
        },
        {
            "mib": "IBM-ELAN-MIB",
            "path": "ibm/IBM-ELAN-MIB.mib"
        },
        {
            "mib": "TN-IP-SOURCE-GUARD-MIB",
            "path": "transition/TN-IP-SOURCE-GUARD-MIB.mib"
        },
        {
            "mib": "OAW-AP1201HL",
            "path": "nokia/stellar/OAW-AP1201HL.mib"
        },
        {
            "mib": "FOUNDRY-SN-IGMP-MIB",
            "path": "brocade/FOUNDRY-SN-IGMP-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-BFD-MIB",
            "path": "dasan/SLE-MPLS-TP-BFD-MIB.mib"
        },
        {
            "mib": "NETGEAR-MMRP-MIB",
            "path": "quanta/mmrp.my.mib"
        },
        {
            "mib": "RIPSAP",
            "path": "nokia/RIPSAP.mib"
        },
        {
            "mib": "TPDIN3-MIB",
            "path": "tycon/TPDIN3-MIB.mib"
        },
        {
            "mib": "HUAWEI-MA5200-DEVICE-MIB",
            "path": "huawei/HUAWEI-MA5200-DEVICE-MIB.mib"
        },
        {
            "mib": "Juniper-Dos-Protection-CONF",
            "path": "junose/Juniper-Dos-Protection-CONF.mib"
        },
        {
            "mib": "OAW-AP1201L",
            "path": "nokia/stellar/OAW-AP1201L.mib"
        },
        {
            "mib": "DLINKSW-SSL-MIB",
            "path": "dlink/DLINKSW-SSL-MIB.mib"
        },
        {
            "mib": "TAIT-INFRA93-94SERIES-COMMON-MIB",
            "path": "tait/TAIT-INFRA93-94SERIES-COMMON-MIB.mib"
        },
        {
            "mib": "JUNIPER-IF-MIB",
            "path": "junos/JUNIPER-IF-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-SFP-MIB",
            "path": "arris/d5/ARRIS-D5-SFP-MIB.mib"
        },
        {
            "mib": "ATM-DXI-MIB",
            "path": "atm/ATM-DXI-MIB.mib"
        },
        {
            "mib": "CISCO-ISDN-MIB",
            "path": "cisco/CISCO-ISDN-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-IP-ACL-MIB",
            "path": "brocade/FOUNDRY-SN-IP-ACL-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-LPS-MIB",
            "path": "dasan/SLE-MPLS-TP-LPS-MIB.mib"
        },
        {
            "mib": "MOXA-SYSTEM-INFO-MIB",
            "path": "moxa/MOXA-SYSTEM-INFO-MIB.mib"
        },
        {
            "mib": "OAW-AP1221",
            "path": "nokia/stellar/OAW-AP1221.mib"
        },
        {
            "mib": "RS-232-MIB",
            "path": "nokia/RS-232-MIB.mib"
        },
        {
            "mib": "HUAWEI-MA5200-MIB",
            "path": "huawei/HUAWEI-MA5200-MIB.mib"
        },
        {
            "mib": "RAJANT-CORPORATION-MIB",
            "path": "rajant/RAJANT-CORPORATION-MIB.mib"
        },
        {
            "mib": "NETGEAR-MRP-MIB",
            "path": "quanta/mrp.my.mib"
        },
        {
            "mib": "CISCO-L2L3-INTERFACE-CONFIG-MIB",
            "path": "cisco/CISCO-L2L3-INTERFACE-CONFIG-MIB.mib"
        },
        {
            "mib": "IBM-ENETDISPATCHER-MIB",
            "path": "ibm/IBM-ENETDISPATCHER-MIB.mib"
        },
        {
            "mib": "TN-IPMC-SNOOPING-MIB",
            "path": "transition/TN-IPMC-SNOOPING-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-SLOT-EXT-MIB",
            "path": "arris/d5/ARRIS-D5-SLOT-EXT-MIB.mib"
        },
        {
            "mib": "HH3C-ENTITY-EXT-MIB",
            "path": "comware/HH3C-ENTITY-EXT-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-NODE-MIB",
            "path": "dasan/SLE-MPLS-TP-NODE-MIB.mib"
        },
        {
            "mib": "OAW-AP1222",
            "path": "nokia/stellar/OAW-AP1222.mib"
        },
        {
            "mib": "TAIT-INFRA93-94SERIES-TC-MIB",
            "path": "tait/TAIT-INFRA93-94SERIES-TC-MIB.mib"
        },
        {
            "mib": "MOXA-SYSTEM-UTILIZATION-MIB",
            "path": "moxa/MOXA-SYSTEM-UTILIZATION-MIB.mib"
        },
        {
            "mib": "NETGEAR-MVRP-MIB",
            "path": "quanta/mvrp.my.mib"
        },
        {
            "mib": "JUNIPER-IFOPTICS-MIB",
            "path": "junos/JUNIPER-IFOPTICS-MIB.mib"
        },
        {
            "mib": "ICT-INVERTER-MIB",
            "path": "ict/ICT-INVERTER-MIB.mib"
        },
        {
            "mib": "SFLOW-MIB",
            "path": "nokia/SFLOW-MIB.mib"
        },
        {
            "mib": "DLINKSW-STACK-MIB",
            "path": "dlink/DLINKSW-STACK-MIB.mib"
        },
        {
            "mib": "ATM-FORUM-ADDR-REG",
            "path": "atm/ATM-FORUM-ADDR-REG.mib"
        },
        {
            "mib": "Juniper-DOS-PROTECTION-MIB",
            "path": "junose/Juniper-DOS-PROTECTION-MIB.mib"
        },
        {
            "mib": "CISCO-LAG-MIB",
            "path": "cisco/CISCO-LAG-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-IP-MIB",
            "path": "brocade/FOUNDRY-SN-IP-MIB.mib"
        },
        {
            "mib": "JUNIPER-IFOTN-MIB",
            "path": "junos/JUNIPER-IFOTN-MIB.mib"
        },
        {
            "mib": "HUAWEI-MAC-AUTHEN-MIB",
            "path": "huawei/HUAWEI-MAC-AUTHEN-MIB.mib"
        },
        {
            "mib": "TAIT-INFRA93SERIES-MIB",
            "path": "tait/TAIT-INFRA93SERIES-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-SOFTWARE-MGR-MIB",
            "path": "arris/d5/ARRIS-D5-SOFTWARE-MGR-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-OAM-MIB",
            "path": "dasan/SLE-MPLS-TP-OAM-MIB.mib"
        },
        {
            "mib": "TN-LACP-MIB",
            "path": "transition/TN-LACP.mib"
        },
        {
            "mib": "PIM-BSR-MIB",
            "path": "quanta/pimbsrrfc5240.my.mib"
        },
        {
            "mib": "ALVARION-DOT11-WLAN-MIB",
            "path": "alvarion/ALVARION-DOT11-WLAN-MIB.mib"
        },
        {
            "mib": "OAW-AP1231",
            "path": "nokia/stellar/OAW-AP1231.mib"
        },
        {
            "mib": "HH3C-ENTITY-VENDORTYPE-OID-MIB",
            "path": "comware/HH3C-ENTITY-VENDORTYPE-OID-MIB.mib"
        },
        {
            "mib": "DLINKSW-STORM-CTRL-MIB",
            "path": "dlink/DLINKSW-STORM-CTRL-MIB.mib"
        },
        {
            "mib": "MOXA-TCST-MIB",
            "path": "moxa/MOXA-TCST-MIB.mib"
        },
        {
            "mib": "IBM-FRAMERELAY-MIB",
            "path": "ibm/IBM-FRAMERELAY-MIB.mib"
        },
        {
            "mib": "ATM-FORUM-ILMI40-MIB",
            "path": "atm/ATM-FORUM-ILMI40-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-PRO-IF-MIB",
            "path": "dasan/SLE-MPLS-TP-PRO-IF-MIB.mib"
        },
        {
            "mib": "TAIT-INFRA93SERIES-TC-MIB",
            "path": "tait/TAIT-INFRA93SERIES-TC-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-AP-MIB",
            "path": "cisco/CISCO-LWAPP-AP-MIB.mib"
        },
        {
            "mib": "SFP-MIB",
            "path": "nokia/SFP-MIB.mib"
        },
        {
            "mib": "ALVARION-DOT11-WLAN-TST-MIB",
            "path": "alvarion/ALVARION-DOT11-WLAN-TST-MIB.mib"
        },
        {
            "mib": "HUAWEI-MACBIND-MIB",
            "path": "huawei/HUAWEI-MACBIND-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-TEST-MIB",
            "path": "arris/d5/ARRIS-D5-TEST-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-IP-VRRP-MIB",
            "path": "brocade/FOUNDRY-SN-IP-VRRP-MIB.mib"
        },
        {
            "mib": "Juniper-IP-POLICY-MIB",
            "path": "junos/Juniper-IP-POLICY-MIB.mib"
        },
        {
            "mib": "NETGEAR-QOS-MIB",
            "path": "quanta/qos.my.mib"
        },
        {
            "mib": "OAW-AP1232",
            "path": "nokia/stellar/OAW-AP1232.mib"
        },
        {
            "mib": "TN-LINK-OAM-MIB",
            "path": "transition/TN-LINK-OAM-MIB.mib"
        },
        {
            "mib": "HH3C-ENTRELATION-MIB",
            "path": "comware/HH3C-ENTRELATION-MIB.mib"
        },
        {
            "mib": "SYSTEM-MIB",
            "path": "nokia/SYSTEM-MIB.mib"
        },
        {
            "mib": "DLINKSW-STP-EXT-MIB",
            "path": "dlink/DLINKSW-STP-EXT-MIB.mib"
        },
        {
            "mib": "ATM-FORUM-M4-MIB",
            "path": "atm/ATM-FORUM-M4-MIB.mib"
        },
        {
            "mib": "MOXA-TURBOCHAIN-MIB",
            "path": "moxa/MOXA-TURBOCHAIN-MIB.mib"
        },
        {
            "mib": "JUNIPER-IPFORWARD-MIB",
            "path": "junos/JUNIPER-IPFORWARD-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-DOT11-CLIENT-MIB",
            "path": "cisco/CISCO-LWAPP-DOT11-CLIENT-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-PW-MIB",
            "path": "dasan/SLE-MPLS-TP-PW-MIB.mib"
        },
        {
            "mib": "TN-LLDP-EXT-MIB",
            "path": "transition/TN-LLDP-EXT-MIB.mib"
        },
        {
            "mib": "Juniper-DS1-CONF",
            "path": "junose/Juniper-DS1-CONF.mib"
        },
        {
            "mib": "NETGEAR-QOS-ACL-MIB",
            "path": "quanta/qos_acl.my.mib"
        },
        {
            "mib": "HUAWEI-MACSEC-MIB",
            "path": "huawei/HUAWEI-MACSEC-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-DOT11-MIB",
            "path": "cisco/CISCO-LWAPP-DOT11-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-VIDEO-ERM-MIB",
            "path": "arris/d5/ARRIS-D5-VIDEO-ERM-MIB.mib"
        },
        {
            "mib": "TAIT-TN9300-MIB",
            "path": "tait/TAIT-TN9300-MIB.mib"
        },
        {
            "mib": "MOXA-TURBORINGV2-MIB",
            "path": "moxa/MOXA-TURBORINGV2-MIB.mib"
        },
        {
            "mib": "OAW-AP1251",
            "path": "nokia/stellar/OAW-AP1251.mib"
        },
        {
            "mib": "ICT-PDU-MIB",
            "path": "ict/ICT-PDU-MIB.mib"
        },
        {
            "mib": "HH3C-EOC-COMMON-MIB",
            "path": "comware/HH3C-EOC-COMMON-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-PW-STATISTICS-MIB",
            "path": "dasan/SLE-MPLS-TP-PW-STATISTICS-MIB.mib"
        },
        {
            "mib": "HUAWEI-MASTERKEY-MIB",
            "path": "huawei/HUAWEI-MASTERKEY-MIB.mib"
        },
        {
            "mib": "NETGEAR-QOS-DIFFSERV-EXTENSIONS-MIB",
            "path": "quanta/qos_diffserv_extensions.my.mib"
        },
        {
            "mib": "WIPIPE-MIB",
            "path": "cradlepoint/WIPIPE-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-MOBILITY-EXT-MIB",
            "path": "cisco/CISCO-LWAPP-MOBILITY-EXT-MIB.mib"
        },
        {
            "mib": "JUNIPER-IPSEC-FLOW-MON-MIB",
            "path": "junos/JUNIPER-IPSEC-FLOW-MON-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-IPX-MIB",
            "path": "brocade/FOUNDRY-SN-IPX-MIB.mib"
        },
        {
            "mib": "TAIT-TN9300-TC",
            "path": "tait/TAIT-TN9300-TC.mib"
        },
        {
            "mib": "DLINKSW-SURVEILLANCE-VLAN-MIB",
            "path": "dlink/DLINKSW-SURVEILLANCE-VLAN-MIB.mib"
        },
        {
            "mib": "OAW-AP1251D",
            "path": "nokia/stellar/OAW-AP1251D.mib"
        },
        {
            "mib": "ATM-FORUM-MIB",
            "path": "atm/ATM-FORUM-MIB.mib"
        },
        {
            "mib": "ICT-PLATINUM-MIB",
            "path": "ict/ICT-PLATINUM-MIB.mib"
        },
        {
            "mib": "BWA-DOT11-WLAN-MIB",
            "path": "alvarion/BWA-DOT11-WLAN-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-RF-MIB",
            "path": "cisco/CISCO-LWAPP-RF-MIB.mib"
        },
        {
            "mib": "JUNIPER-IPv4-MIB",
            "path": "junos/JUNIPER-IPv4-MIB.mib"
        },
        {
            "mib": "HUAWEI-MC-TRUNK-MIB",
            "path": "huawei/HUAWEI-MC-TRUNK-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-TUNNEL-MIB",
            "path": "dasan/SLE-MPLS-TP-TUNNEL-MIB.mib"
        },
        {
            "mib": "NETGEAR-QOS-DIFFSERV-PRIVATE-MIB",
            "path": "quanta/qos_diffserv_private.my.mib"
        },
        {
            "mib": "ARRIS-D5-VIDEO-IP-BUNDLE-MIB",
            "path": "arris/d5/ARRIS-D5-VIDEO-IP-BUNDLE-MIB.mib"
        },
        {
            "mib": "OAW-AP1321",
            "path": "nokia/stellar/OAW-AP1321.mib"
        },
        {
            "mib": "TN-LLDP-MIB",
            "path": "transition/TN-LLDP-MIB.mib"
        },
        {
            "mib": "HH3C-EPON-DEVICE-MIB",
            "path": "comware/HH3C-EPON-DEVICE-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-MAC-AUTHENTICATION-MIB",
            "path": "brocade/FOUNDRY-SN-MAC-AUTHENTICATION-MIB.mib"
        },
        {
            "mib": "Juniper-DS1-MIB",
            "path": "junose/Juniper-DS1-MIB.mib"
        },
        {
            "mib": "NETGEAR-RADIUS-AUTH-CLIENT-MIB",
            "path": "quanta/radius.my.mib"
        },
        {
            "mib": "TIMETRA-BGP-MIB",
            "path": "nokia/TIMETRA-BGP-MIB.mib"
        },
        {
            "mib": "DLINKSW-SWITCHPORT-MIB",
            "path": "dlink/DLINKSW-SWITCHPORT-MIB.mib"
        },
        {
            "mib": "TAIT-TNADMIN-MIB",
            "path": "tait/TAIT-TNADMIN-MIB.mib"
        },
        {
            "mib": "DATACOM-REG",
            "path": "datacom/DATACOM-REG.mib"
        },
        {
            "mib": "CPS-MIB",
            "path": "cyberpower/CPS-MIB.mib"
        },
        {
            "mib": "TIMETRA-CELLULAR-MIB",
            "path": "nokia/TIMETRA-CELLULAR-MIB.mib"
        },
        {
            "mib": "ATM-FORUM-SNMP-M4-MIB",
            "path": "atm/ATM-FORUM-SNMP-M4-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-VIDEO-MIB",
            "path": "arris/d5/ARRIS-D5-VIDEO-MIB.mib"
        },
        {
            "mib": "HUAWEI-MEMORY-MIB",
            "path": "huawei/HUAWEI-MEMORY-MIB.mib"
        },
        {
            "mib": "RADIUS-ACC-CLIENT-MIB",
            "path": "quanta/radius_acc_client.my.mib"
        },
        {
            "mib": "IBM-GbTOR-10G-L2L3-MIB",
            "path": "ibm/IBM-GbTOR-10G-L2L3-MIB.mib"
        },
        {
            "mib": "TN-LOAM-EXT-MIB",
            "path": "transition/TN-LOAM-EXT-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-SYS-MIB",
            "path": "cisco/CISCO-LWAPP-SYS-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-TUNNEL-STATISTICS-MIB",
            "path": "dasan/SLE-MPLS-TP-TUNNEL-STATISTICS-MIB.mib"
        },
        {
            "mib": "JUNIPER-IPv6-MIB",
            "path": "junos/JUNIPER-IPv6-MIB.mib"
        },
        {
            "mib": "TAIT-TNADMIN-MODULE-MIB",
            "path": "tait/TAIT-TNADMIN-MODULE-MIB.mib"
        },
        {
            "mib": "DATACOM-SMI",
            "path": "datacom/DATACOM-SMI.mib"
        },
        {
            "mib": "ICT-POWERSYSTEM-MIB",
            "path": "ict/ICT-POWERSYSTEM-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-MRP-MIB",
            "path": "brocade/FOUNDRY-SN-MRP-MIB.mib"
        },
        {
            "mib": "HH3C-EPON-FB-MIB",
            "path": "comware/HH3C-EPON-FB-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-TC-MIB",
            "path": "cisco/CISCO-LWAPP-TC-MIB.mib"
        },
        {
            "mib": "OAW-AP1322",
            "path": "nokia/stellar/OAW-AP1322.mib"
        },
        {
            "mib": "ATM-FORUM-SRVC-REG",
            "path": "atm/ATM-FORUM-SRVC-REG.mib"
        },
        {
            "mib": "ARRIS-D5-VIDEO-SESSION-MIB",
            "path": "arris/d5/ARRIS-D5-VIDEO-SESSION-MIB.mib"
        },
        {
            "mib": "DLINKSW-SYSLOG-MIB",
            "path": "dlink/DLINKSW-SYSLOG-MIB.mib"
        },
        {
            "mib": "HUAWEI-MFF-MIB",
            "path": "huawei/HUAWEI-MFF-MIB.mib"
        },
        {
            "mib": "SLE-MPLS-TP-VPLS-MIB",
            "path": "dasan/SLE-MPLS-TP-VPLS-MIB.mib"
        },
        {
            "mib": "Juniper-DS3-CONF",
            "path": "junose/Juniper-DS3-CONF.mib"
        },
        {
            "mib": "RADIUS-AUTH-CLIENT-MIB",
            "path": "quanta/radius_auth_client.my.mib"
        },
        {
            "mib": "TAIT-TNADMIN-TC",
            "path": "tait/TAIT-TNADMIN-TC.mib"
        },
        {
            "mib": "TN-LOOP-PROTECT-MIB",
            "path": "transition/TN-LOOP-PROTECT-MIB.mib"
        },
        {
            "mib": "JUNIPER-JDHCP-MIB",
            "path": "junos/JUNIPER-JDHCP-MIB.mib"
        },
        {
            "mib": "OAW-AP1361",
            "path": "nokia/stellar/OAW-AP1361.mib"
        },
        {
            "mib": "RUCKUS-CTRL-MIB",
            "path": "ruckus/RUCKUS-CTRL-MIB.mib"
        },
        {
            "mib": "DMOS-SYSMON-MIB",
            "path": "datacom/DMOS-SYSMON-MIB.mib"
        },
        {
            "mib": "DLINKSW-SYSTEM-FILE-MIB",
            "path": "dlink/DLINKSW-SYSTEM-FILE-MIB.mib"
        },
        {
            "mib": "CISCO-LWAPP-WLAN-MIB",
            "path": "cisco/CISCO-LWAPP-WLAN-MIB.mib"
        },
        {
            "mib": "ATM-FORUM-TC-MIB",
            "path": "atm/ATM-FORUM-TC-MIB.mib"
        },
        {
            "mib": "Juniper-DS3-MIB",
            "path": "junose/Juniper-DS3-MIB.mib"
        },
        {
            "mib": "ARRIS-D5-VIDEO-VIF-MIB",
            "path": "arris/d5/ARRIS-D5-VIDEO-VIF-MIB.mib"
        },
        {
            "mib": "HUAWEI-MFLP-MIB",
            "path": "huawei/HUAWEI-MFLP-MIB.mib"
        },
        {
            "mib": "PT-MIB",
            "path": "protelevision/PT-MIB.mib"
        },
        {
            "mib": "HH3C-EPON-MIB",
            "path": "comware/HH3C-EPON-MIB.mib"
        },
        {
            "mib": "OAW-AP1361D",
            "path": "nokia/stellar/OAW-AP1361D.mib"
        },
        {
            "mib": "FOUNDRY-SN-NOTIFICATION-MIB",
            "path": "brocade/FOUNDRY-SN-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "JUNIPER-JDHCPV6-MIB",
            "path": "junos/JUNIPER-JDHCPV6-MIB.mib"
        },
        {
            "mib": "SOCOMECPDU-MIB",
            "path": "socomec/SOCOMECPDU-MIB.mib"
        },
        {
            "mib": "DMOS-TC-MIB",
            "path": "datacom/DMOS-TC-MIB.mib"
        },
        {
            "mib": "SLE-MVQOS-MIB",
            "path": "dasan/SLE-MVQOS-MIB.mib"
        },
        {
            "mib": "DLINKSW-TC-MIB",
            "path": "dlink/DLINKSW-TC-MIB.mib"
        },
        {
            "mib": "RUCKUS-DEVICE-MIB",
            "path": "ruckus/RUCKUS-DEVICE-MIB.mib"
        },
        {
            "mib": "ATM-MIB",
            "path": "atm/ATM-MIB.mib"
        },
        {
            "mib": "OAW-AP1362",
            "path": "nokia/stellar/OAW-AP1362.mib"
        },
        {
            "mib": "TN-MAC-MIB",
            "path": "transition/TN-MAC-MIB.mib"
        },
        {
            "mib": "Juniper-DVMRP-CONF",
            "path": "junose/Juniper-DVMRP-CONF.mib"
        },
        {
            "mib": "IANA-RTPROTO-MIB",
            "path": "quanta/rtproto.my.mib"
        },
        {
            "mib": "ARRIS-D5-WAN-POST-MIB",
            "path": "arris/d5/ARRIS-D5-WAN-POST-MIB.mib"
        },
        {
            "mib": "DLINKSW-TELNET-MIB",
            "path": "dlink/DLINKSW-TELNET-MIB.mib"
        },
        {
            "mib": "CISCO-MAC-NOTIFICATION-MIB",
            "path": "cisco/CISCO-MAC-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "HH3C-EPON-UNI-MIB",
            "path": "comware/HH3C-EPON-UNI-MIB.mib"
        },
        {
            "mib": "SOCOMECUPS-MIB",
            "path": "socomec/SOCOMECUPS-MIB.mib"
        },
        {
            "mib": "RUCKUS-HWINFO-MIB",
            "path": "ruckus/RUCKUS-HWINFO-MIB.mib"
        },
        {
            "mib": "HUAWEI-MGMD-STD-MIB",
            "path": "huawei/HUAWEI-MGMD-STD-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-OSPF-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-OSPF-GROUP-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-AUTH-MIB",
            "path": "junos/JUNIPER-JS-AUTH-MIB.mib"
        },
        {
            "mib": "DMswitch-MIB",
            "path": "datacom/DMswitch-MIB.mib"
        },
        {
            "mib": "PT3080-MIB",
            "path": "protelevision/PT3080-MIB.mib"
        },
        {
            "mib": "ARRIS-MIB",
            "path": "arris/d5/ARRIS-MIB.mib"
        },
        {
            "mib": "SFLOW-MIB",
            "path": "quanta/sflow.my.mib"
        },
        {
            "mib": "SLE-NETWORK-MIB",
            "path": "dasan/SLE-NETWORK-MIB.mib"
        },
        {
            "mib": "DLINKSW-TIME-MIB",
            "path": "dlink/DLINKSW-TIME-MIB.mib"
        },
        {
            "mib": "AXIS-ROOT-MIB",
            "path": "axis/AXIS-ROOT-MIB.mib"
        },
        {
            "mib": "CXR-TS-MIB",
            "path": "cxr-networks/CXR-TS-MIB.mib"
        },
        {
            "mib": "TN-MGMT-MIB",
            "path": "transition/TN-MGMT-MIB.mib"
        },
        {
            "mib": "RUCKUS-PRODUCTS-MIB",
            "path": "ruckus/RUCKUS-PRODUCTS-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-POS-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-POS-GROUP-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP500-MIB",
            "path": "cambium/500/CAMBIUM-PTP500-MIB.mib"
        },
        {
            "mib": "ATM-TC-MIB",
            "path": "atm/ATM-TC-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-CERT-MIB",
            "path": "junos/JUNIPER-JS-CERT-MIB.mib"
        },
        {
            "mib": "Juniper-DVMRP-MIB",
            "path": "junose/Juniper-DVMRP-MIB.mib"
        },
        {
            "mib": "TELESTE-COMMON-MIB",
            "path": "teleste/TELESTE-COMMON-MIB.mib"
        },
        {
            "mib": "SNMP-RESEARCH-MIB",
            "path": "quanta/snmp-res.my.mib"
        },
        {
            "mib": "DLINKSW-TIME-RANGE-MIB",
            "path": "dlink/DLINKSW-TIME-RANGE-MIB.mib"
        },
        {
            "mib": "HUAWEI-MIB",
            "path": "huawei/HUAWEI-MIB.mib"
        },
        {
            "mib": "TIMETRA-CHASSIS-MIB",
            "path": "nokia/TIMETRA-CHASSIS-MIB.mib"
        },
        {
            "mib": "TIMETRA-DHCP-SERVER-MIB",
            "path": "nokia/TIMETRA-DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "HH3C-EVB-MIB",
            "path": "comware/HH3C-EVB-MIB.mib"
        },
        {
            "mib": "AXIS-VIDEO-MIB",
            "path": "axis/AXIS-VIDEO-MIB.mib"
        },
        {
            "mib": "SLE-Network6-MIB",
            "path": "dasan/SLE-Network6-MIB.mib"
        },
        {
            "mib": "SIM-MIB",
            "path": "arris/d5/SIM-MIB.mib"
        },
        {
            "mib": "TN-MGMT-TDM-MIB",
            "path": "transition/TN-MGMT-TDM-MIB.mib"
        },
        {
            "mib": "VIPRINET-MIB",
            "path": "viprinet/VIPRINET-MIB.mib"
        },
        {
            "mib": "CISCO-MEDIA-GATEWAY-MIB",
            "path": "cisco/CISCO-MEDIA-GATEWAY-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-DNS-MIB",
            "path": "junos/JUNIPER-JS-DNS-MIB.mib"
        },
        {
            "mib": "BKTEL-HFC862-BASE-MIB",
            "path": "bktel/BKTEL-HFC862-BASE-MIB.mib"
        },
        {
            "mib": "DLINKSW-TRAFFIC-SEGMENT-MIB",
            "path": "dlink/DLINKSW-TRAFFIC-SEGMENT-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-ROOT-MIB",
            "path": "brocade/FOUNDRY-SN-ROOT-MIB.mib"
        },
        {
            "mib": "SLE-OSPF-MIB",
            "path": "dasan/SLE-OSPF-MIB.mib"
        },
        {
            "mib": "IBM-GbTOR-G8052-MIB",
            "path": "ibm/IBM-GbTOR-G8052-MIB.mib"
        },
        {
            "mib": "CISCO-MEMORY-POOL-MIB",
            "path": "cisco/CISCO-MEMORY-POOL-MIB.mib"
        },
        {
            "mib": "SR-AGENT-INFO-MIB",
            "path": "quanta/srAgentInfo.my.mib"
        },
        {
            "mib": "TN-MIRRORING-MIB",
            "path": "transition/TN-MIRRORING-MIB.mib"
        },
        {
            "mib": "MNI-PROTEUS-AMT-MIB",
            "path": "mni/MNI-PROTEUS-AMT-MIB.mib"
        },
        {
            "mib": "PICA-PRIVATE-MIB",
            "path": "picos/PICA-PRIVATE-MIB.mib"
        },
        {
            "mib": "CONEL-GPS-MIB",
            "path": "icr-os/CONEL-GPS-MIB.mib"
        },
        {
            "mib": "RUCKUS-ROOT-MIB",
            "path": "ruckus/RUCKUS-ROOT-MIB.mib"
        },
        {
            "mib": "TACACS-CLIENT-MIB",
            "path": "quanta/tacacsclient.my.mib"
        },
        {
            "mib": "Juniper-Entity-CONF",
            "path": "junose/Juniper-Entity-CONF.mib"
        },
        {
            "mib": "TIMETRA-FILTER-MIB",
            "path": "nokia/TIMETRA-FILTER-MIB.mib"
        },
        {
            "mib": "TIMETRA-GLOBAL-MIB",
            "path": "nokia/TIMETRA-GLOBAL-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-FLOW-MIB",
            "path": "junos/JUNIPER-JS-FLOW-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-ROUTER-TRAP-MIB",
            "path": "brocade/FOUNDRY-SN-ROUTER-TRAP-MIB.mib"
        },
        {
            "mib": "HH3C-EVC-MIB",
            "path": "comware/HH3C-EVC-MIB.mib"
        },
        {
            "mib": "HUAWEI-MINM-MIB",
            "path": "huawei/HUAWEI-MINM-MIB.mib"
        },
        {
            "mib": "CISCO-MVPN-MIB",
            "path": "cisco/CISCO-MVPN-MIB.mib"
        },
        {
            "mib": "TN-MPLS-MIB",
            "path": "transition/TN-MPLS-MIB.mib"
        },
        {
            "mib": "SLE-OSPFv3-MIB",
            "path": "dasan/SLE-OSPFv3-MIB.mib"
        },
        {
            "mib": "TELESTE-LUMINATO-MIB",
            "path": "teleste/TELESTE-LUMINATO-MIB.mib"
        },
        {
            "mib": "BKTEL-HFC862-BES-V01-MIB",
            "path": "bktel/BKTEL-HFC862-BES-V01-MIB.mib"
        },
        {
            "mib": "RUCKUS-SWINFO-MIB",
            "path": "ruckus/RUCKUS-SWINFO-MIB.mib"
        },
        {
            "mib": "ENLOGIC-PDU-MIB",
            "path": "enlogic/ENLOGIC-PDU-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-STACKING-MIB",
            "path": "brocade/FOUNDRY-SN-STACKING-MIB.mib"
        },
        {
            "mib": "VOLIUS-OA-MIB",
            "path": "volius/VOLIUS-OA-MIB.mib"
        },
        {
            "mib": "CISCO-NAC-TC-MIB",
            "path": "cisco/CISCO-NAC-TC-MIB.mib"
        },
        {
            "mib": "DLINKSW-VLAN-MIB",
            "path": "dlink/DLINKSW-VLAN-MIB.mib"
        },
        {
            "mib": "TN-MRP-MIB",
            "path": "transition/TN-MRP-MIB.mib"
        },
        {
            "mib": "USM-TARGET-TAG-MIB",
            "path": "quanta/usm-tran.my.mib"
        },
        {
            "mib": "RUCKUS-SYSTEM-MIB",
            "path": "ruckus/RUCKUS-SYSTEM-MIB.mib"
        },
        {
            "mib": "AUTO-CONFIGURATION-MIB",
            "path": "raisecom/AUTO-CONFIGURATION-MIB.mib"
        },
        {
            "mib": "HH3C-EVI-MIB",
            "path": "comware/HH3C-EVI-MIB.mib"
        },
        {
            "mib": "TELESTE-LUMINATO-MIB2",
            "path": "teleste/TELESTE-LUMINATO-MIB2.mib"
        },
        {
            "mib": "Juniper-ERX-Registry",
            "path": "junose/Juniper-ERX-Registry.mib"
        },
        {
            "mib": "JUNIPER-JS-IDP-MIB",
            "path": "junos/JUNIPER-JS-IDP-MIB.mib"
        },
        {
            "mib": "SLE-PERFORMANCEMGMT-MIB",
            "path": "dasan/SLE-PERFORMANCEMGMT-MIB.mib"
        },
        {
            "mib": "HUAWEI-MIRROR-MIB",
            "path": "huawei/HUAWEI-MIRROR-MIB.mib"
        },
        {
            "mib": "CISCO-NS-MIB",
            "path": "cisco/CISCO-NS-MIB.mib"
        },
        {
            "mib": "BKTEL-HFC862-HMSNE-MIB",
            "path": "bktel/BKTEL-HFC862-HMSNE-MIB.mib"
        },
        {
            "mib": "CONEL-INFO-MIB",
            "path": "icr-os/CONEL-INFO-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-SW-L4-SWITCH-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-SW-L4-SWITCH-GROUP-MIB.mib"
        },
        {
            "mib": "CONVERTOR-SYSTEM-MIB",
            "path": "raisecom/CONVERTOR-SYSTEM-MIB.mib"
        },
        {
            "mib": "HUAWEI-MODULE-INFO-MIB",
            "path": "huawei/HUAWEI-MODULE-INFO-MIB.mib"
        },
        {
            "mib": "Juniper-ERX-System-CONF",
            "path": "junose/Juniper-ERX-System-CONF.mib"
        },
        {
            "mib": "TIMETRA-LDP-MIB",
            "path": "nokia/TIMETRA-LDP-MIB.mib"
        },
        {
            "mib": "TN-MVR-MIB",
            "path": "transition/TN-MVR-MIB.mib"
        },
        {
            "mib": "IBOOTPDU-MIB",
            "path": "dataprobe/IBOOTPDU-MIB.mib"
        },
        {
            "mib": "HH3C-EVPN-MIB",
            "path": "comware/HH3C-EVPN-MIB.mib"
        },
        {
            "mib": "RUCKUS-SZ-CONFIG-WLAN-MIB",
            "path": "ruckus/RUCKUS-SZ-CONFIG-WLAN-MIB.mib"
        },
        {
            "mib": "CONVERTOR-VLAN-MIB",
            "path": "raisecom/CONVERTOR-VLAN-MIB.mib"
        },
        {
            "mib": "ENLOGIC-PDU2-MIB",
            "path": "enlogic/ENLOGIC-PDU2-MIB.mib"
        },
        {
            "mib": "SLE-PM-MIB",
            "path": "dasan/SLE-PM-MIB.mib"
        },
        {
            "mib": "CISCO-OPTICAL-MONITOR-MIB",
            "path": "cisco/CISCO-OPTICAL-MONITOR-MIB.mib"
        },
        {
            "mib": "DLINKSW-VOICE-VLAN-MIB",
            "path": "dlink/DLINKSW-VOICE-VLAN-MIB.mib"
        },
        {
            "mib": "TELESTE-ROOT-MIB",
            "path": "teleste/TELESTE-ROOT-MIB.mib"
        },
        {
            "mib": "ARISTA-BGP4V2-MIB",
            "path": "arista/ARISTA-BGP4V2-MIB.mib"
        },
        {
            "mib": "DHCP-CLIENT-MIB",
            "path": "raisecom/DHCP-CLIENT-MIB.mib"
        },
        {
            "mib": "SFOS-FIREWALL-MIB",
            "path": "sophos/SFOS-FIREWALL-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-SWITCH-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-SWITCH-GROUP-MIB.mib"
        },
        {
            "mib": "TIMETRA-LLDP-MIB",
            "path": "nokia/TIMETRA-LLDP-MIB.mib"
        },
        {
            "mib": "CONEL-IO-MIB",
            "path": "icr-os/CONEL-IO-MIB.mib"
        },
        {
            "mib": "HUAWEI-MP-MIB",
            "path": "huawei/HUAWEI-MP-MIB.mib"
        },
        {
            "mib": "BKTEL-HFC862-NECE-MIB",
            "path": "bktel/BKTEL-HFC862-NECE-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-IF-EXT-MIB",
            "path": "junos/JUNIPER-JS-IF-EXT-MIB.mib"
        },
        {
            "mib": "HH3C-FAILOVER-MIB",
            "path": "comware/HH3C-FAILOVER-MIB.mib"
        },
        {
            "mib": "VOLIUS-OR-MIB",
            "path": "volius/VOLIUS-OR-MIB.mib"
        },
        {
            "mib": "Juniper-ERX-System-MIB",
            "path": "junose/Juniper-ERX-System-MIB.mib"
        },
        {
            "mib": "DHCP-OPTION-MIB",
            "path": "raisecom/DHCP-OPTION-MIB.mib"
        },
        {
            "mib": "TN-NAS-MIB",
            "path": "transition/TN-NAS-MIB.mib"
        },
        {
            "mib": "CYBEROAM-MIB",
            "path": "cyberoam/CYBEROAM-MIB.mib"
        },
        {
            "mib": "SLE-PPPOE-MIB",
            "path": "dasan/SLE-PPPOE-MIB.mib"
        },
        {
            "mib": "IBM-GbTOR-G8264-MIB",
            "path": "ibm/IBM-GbTOR-G8264-MIB.mib"
        },
        {
            "mib": "CISCO-OPTICAL-PATCH-MIB",
            "path": "cisco/CISCO-OPTICAL-PATCH-MIB.mib"
        },
        {
            "mib": "BKTEL-HFC862-OA-V01-MIB",
            "path": "bktel/BKTEL-HFC862-OA-V01-MIB.mib"
        },
        {
            "mib": "HUAWEI-MPLS-EXTEND-MIB",
            "path": "huawei/HUAWEI-MPLS-EXTEND-MIB.mib"
        },
        {
            "mib": "ARISTA-BGP4V2-TC-MIB",
            "path": "arista/ARISTA-BGP4V2-TC-MIB.mib"
        },
        {
            "mib": "HH3C-FC-FLOGIN-MIB",
            "path": "comware/HH3C-FC-FLOGIN-MIB.mib"
        },
        {
            "mib": "VOLIUS-OS-MIB",
            "path": "volius/VOLIUS-OS-MIB.mib"
        },
        {
            "mib": "TIMETRA-LOG-MIB",
            "path": "nokia/TIMETRA-LOG-MIB.mib"
        },
        {
            "mib": "DLINKSW-WEB-AUTH-MIB",
            "path": "dlink/DLINKSW-WEB-AUTH-MIB.mib"
        },
        {
            "mib": "IPSEC-ISAKMP-IKE-DOI-TC",
            "path": "watchguard/IPSEC-ISAKMP-IKE-DOI-TC.mib"
        },
        {
            "mib": "FOUNDRY-SN-TRAP-MIB",
            "path": "brocade/FOUNDRY-SN-TRAP-MIB.mib"
        },
        {
            "mib": "TN-OTDR-MIB",
            "path": "transition/TN-OTDR-MIB.mib"
        },
        {
            "mib": "DAHUA-SNMP-MIB",
            "path": "dahua/DAHUA-SNMP-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-IPSEC-VPN-MIB",
            "path": "junos/JUNIPER-JS-IPSEC-VPN-MIB.mib"
        },
        {
            "mib": "RUCKUS-SZ-EVENT-MIB",
            "path": "ruckus/RUCKUS-SZ-EVENT-MIB.mib"
        },
        {
            "mib": "HH3C-FC-NAME-SERVER-MIB",
            "path": "comware/HH3C-FC-NAME-SERVER-MIB.mib"
        },
        {
            "mib": "DHCP-RELAY-MIB",
            "path": "raisecom/DHCP-RELAY-MIB.mib"
        },
        {
            "mib": "SLE-QOS-MIB",
            "path": "dasan/SLE-QOS-MIB.mib"
        },
        {
            "mib": "Juniper-ES2-Registry",
            "path": "junose/Juniper-ES2-Registry.mib"
        },
        {
            "mib": "ARISTA-ENTITY-SENSOR-MIB",
            "path": "arista/ARISTA-ENTITY-SENSOR-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-VSRP-MIB",
            "path": "brocade/FOUNDRY-SN-VSRP-MIB.mib"
        },
        {
            "mib": "DLINKSW-WEB-COMMON-MIB",
            "path": "dlink/DLINKSW-WEB-COMMON-MIB.mib"
        },
        {
            "mib": "BKTEL-HFC862-OVTX-V11-MIB",
            "path": "bktel/BKTEL-HFC862-OVTX-V11-MIB.mib"
        },
        {
            "mib": "HH3C-FC-PING-MIB",
            "path": "comware/HH3C-FC-PING-MIB.mib"
        },
        {
            "mib": "VOLIUS-OT-MIB",
            "path": "volius/VOLIUS-OT-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-NAT-MIB",
            "path": "junos/JUNIPER-JS-NAT-MIB.mib"
        },
        {
            "mib": "HUAWEI-MPLSLDP-MIB",
            "path": "huawei/HUAWEI-MPLSLDP-MIB.mib"
        },
        {
            "mib": "CISCO-PAGP-MIB",
            "path": "cisco/CISCO-PAGP-MIB.mib"
        },
        {
            "mib": "RUCKUS-SZ-SYSTEM-MIB",
            "path": "ruckus/RUCKUS-SZ-SYSTEM-MIB.mib"
        },
        {
            "mib": "FOUNDRY-SN-WIRELESS-GROUP-MIB",
            "path": "brocade/FOUNDRY-SN-WIRELESS-GROUP-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-CLIENT-MIB",
            "path": "watchguard/WATCHGUARD-CLIENT-MIB.mib"
        },
        {
            "mib": "AI-AP-MIB",
            "path": "arubaos/AI-AP-MIB.mib"
        },
        {
            "mib": "TN-PORT-MIB",
            "path": "transition/TN-PORT-MIB.mib"
        },
        {
            "mib": "DHCP-SERVER-MIB",
            "path": "raisecom/DHCP-SERVER-MIB.mib"
        },
        {
            "mib": "ARISTA-GENERAL-MIB",
            "path": "arista/ARISTA-GENERAL-MIB.mib"
        },
        {
            "mib": "DataAire-dap4-al-MIB",
            "path": "dataaire/DataAire-dap4-al-MIB.mib"
        },
        {
            "mib": "Juniper-Ethernet-CONF",
            "path": "junose/Juniper-Ethernet-CONF.mib"
        },
        {
            "mib": "HH3C-FC-PSM-MIB",
            "path": "comware/HH3C-FC-PSM-MIB.mib"
        },
        {
            "mib": "EQUIPMENT-MIB",
            "path": "dlink/EQUIPMENT-MIB.mib"
        },
        {
            "mib": "CISCO-PORT-SECURITY-MIB",
            "path": "cisco/CISCO-PORT-SECURITY-MIB.mib"
        },
        {
            "mib": "ARUBA-MGMT-MIB",
            "path": "arubaos/ARUBA-MGMT-MIB.mib"
        },
        {
            "mib": "SLE-RED-MIB",
            "path": "dasan/SLE-RED-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-PACKET-MIRROR-MIB",
            "path": "junos/JUNIPER-JS-PACKET-MIRROR-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-HA-MIB",
            "path": "watchguard/WATCHGUARD-HA-MIB.mib"
        },
        {
            "mib": "FOUNDRY-VLAN-CAR-MIB",
            "path": "brocade/FOUNDRY-VLAN-CAR-MIB.mib"
        },
        {
            "mib": "TN-POWER-SUPPLY-MIB",
            "path": "transition/TN-POWER-SUPPLY-MIB.mib"
        },
        {
            "mib": "DHCP-SNOOPING-MIB",
            "path": "raisecom/DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "RUCKUS-SZ-WLAN-MIB",
            "path": "ruckus/RUCKUS-SZ-WLAN-MIB.mib"
        },
        {
            "mib": "TIMETRA-MPLS-MIB",
            "path": "nokia/TIMETRA-MPLS-MIB.mib"
        },
        {
            "mib": "ARISTA-IF-MIB",
            "path": "arista/ARISTA-IF-MIB.mib"
        },
        {
            "mib": "HH3C-FC-TC-MIB",
            "path": "comware/HH3C-FC-TC-MIB.mib"
        },
        {
            "mib": "HUAWEI-MPLSLSR-EXT-MIB",
            "path": "huawei/HUAWEI-MPLSLSR-EXT-MIB.mib"
        },
        {
            "mib": "ARUBA-MIB",
            "path": "arubaos/ARUBA-MIB.mib"
        },
        {
            "mib": "CISCO-POWER-ETHERNET-EXT-MIB",
            "path": "cisco/CISCO-POWER-ETHERNET-EXT-MIB.mib"
        },
        {
            "mib": "SINGLE-IP-MIB",
            "path": "dlink/SINGLE-IP-MIB.mib"
        },
        {
            "mib": "IPDHCP-RELAY-MIB",
            "path": "raisecom/IPDHCP-RELAY-MIB.mib"
        },
        {
            "mib": "SLE-RIP-MIB",
            "path": "dasan/SLE-RIP-MIB.mib"
        },
        {
            "mib": "IANA-PWE3-MIB",
            "path": "brocade/IANA-PWE3-MIB.mib"
        },
        {
            "mib": "ARISTA-NEXTHOP-GROUP-MIB",
            "path": "arista/ARISTA-NEXTHOP-GROUP-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-POLICY-MIB",
            "path": "junos/JUNIPER-JS-POLICY-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-INFO-SYSTEM-MIB",
            "path": "watchguard/WATCHGUARD-INFO-SYSTEM-MIB.mib"
        },
        {
            "mib": "Juniper-ETHERNET-MIB",
            "path": "junose/Juniper-ETHERNET-MIB.mib"
        },
        {
            "mib": "TN-PROTECTION-MIB",
            "path": "transition/TN-Protection-MIB.mib"
        },
        {
            "mib": "CHECKPOINT-MIB",
            "path": "checkpoint/CHECKPOINT-MIB.mib"
        },
        {
            "mib": "APEX-MIB",
            "path": "arris/APEX-MIB.mib"
        },
        {
            "mib": "HH3C-FC-TRACE-ROUTE-MIB",
            "path": "comware/HH3C-FC-TRACE-ROUTE-MIB.mib"
        },
        {
            "mib": "RUCKUS-TC-MIB",
            "path": "ruckus/RUCKUS-TC-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-CMM-MIB",
            "path": "arris/ARRIS-C3-CMM-MIB.mib"
        },
        {
            "mib": "IBM-GbTOR-G8264CS-MIB",
            "path": "ibm/IBM-GbTOR-G8264CS-MIB.mib"
        },
        {
            "mib": "ARUBA-TC",
            "path": "arubaos/ARUBA-TC.mib"
        },
        {
            "mib": "HUAWEI-MPLSOAM-MIB",
            "path": "huawei/HUAWEI-MPLSOAM-MIB.mib"
        },
        {
            "mib": "FAN-MIB",
            "path": "fiberhome/FAN-MIB.mib"
        },
        {
            "mib": "TIMETRA-NAT-MIB",
            "path": "nokia/TIMETRA-NAT-MIB.mib"
        },
        {
            "mib": "CISCO-PRIVATE-VLAN-MIB",
            "path": "cisco/CISCO-PRIVATE-VLAN-MIB.mib"
        },
        {
            "mib": "ARISTA-SMI-MIB",
            "path": "arista/ARISTA-SMI-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-IPSEC-ENDPOINT-PAIR-MIB",
            "path": "watchguard/WATCHGUARD-IPSEC-ENDPOINT-PAIR-MIB.mib"
        },
        {
            "mib": "RUCKUS-UNLEASHED-EVENT-MIB",
            "path": "ruckus/RUCKUS-UNLEASHED-EVENT-MIB.mib"
        },
        {
            "mib": "CPPM-MIB",
            "path": "arubaos/CPPM-MIB.mib"
        },
        {
            "mib": "SLE-RIPng-MIB",
            "path": "dasan/SLE-RIPng-MIB.mib"
        },
        {
            "mib": "SWDGS1510PRIMGMT-MIB",
            "path": "dlink/SWDGS1510PRIMGMT-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-FPD-MIB",
            "path": "arris/ARRIS-C3-FPD-MIB.mib"
        },
        {
            "mib": "Juniper-Event-Manager-CONF",
            "path": "junose/Juniper-Event-Manager-CONF.mib"
        },
        {
            "mib": "IPDHCP-SERVER-MIB",
            "path": "raisecom/IPDHCP-SERVER-MIB.mib"
        },
        {
            "mib": "CONEL-MIB",
            "path": "icr-os/CONEL-MIB.mib"
        },
        {
            "mib": "TN-PROV-MIB",
            "path": "transition/TN-PROV-MIB.mib"
        },
        {
            "mib": "HH3C-FC-ZONE-SERVER-MIB",
            "path": "comware/HH3C-FC-ZONE-SERVER-MIB.mib"
        },
        {
            "mib": "SW-MIB",
            "path": "brocade/SW-MIB.mib"
        },
        {
            "mib": "CONEL-MBUS-MIB",
            "path": "icr-os/CONEL-MBUS-MIB.mib"
        },
        {
            "mib": "HUAWEI-MPLSOAM-PS-MIB",
            "path": "huawei/HUAWEI-MPLSOAM-PS-MIB.mib"
        },
        {
            "mib": "TIMERANGE-MIB",
            "path": "dlink/TIMERANGE-MIB.mib"
        },
        {
            "mib": "WLSR-AP-MIB",
            "path": "arubaos/WLSR-AP-MIB.mib"
        },
        {
            "mib": "JUNIPER-EXPERIMENT-MIB",
            "path": "junose/JUNIPER-EXPERIMENT-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-SCREENING-MIB",
            "path": "junos/JUNIPER-JS-SCREENING-MIB.mib"
        },
        {
            "mib": "GEPON-OLT-COMMON-MIB",
            "path": "fiberhome/GEPON-OLT-COMMON-MIB.mib"
        },
        {
            "mib": "ARISTA-VRF-MIB",
            "path": "arista/ARISTA-VRF-MIB.mib"
        },
        {
            "mib": "LLDP-PRI-MIB",
            "path": "raisecom/LLDP-PRI-MIB.mib"
        },
        {
            "mib": "SWBASE-MIB",
            "path": "brocade/SWBASE-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-IF-MIB",
            "path": "arris/ARRIS-C3-IF-MIB.mib"
        },
        {
            "mib": "HUAWEI-MSDP-MIB",
            "path": "huawei/HUAWEI-MSDP-MIB.mib"
        },
        {
            "mib": "CISCO-PROCESS-MIB",
            "path": "cisco/CISCO-PROCESS-MIB.mib"
        },
        {
            "mib": "IBM-GbTOR-G8264T-MIB",
            "path": "ibm/IBM-GbTOR-G8264T-MIB.mib"
        },
        {
            "mib": "WLSX-AUTH-MIB",
            "path": "arubaos/WLSX-AUTH-MIB.mib"
        },
        {
            "mib": "WRI-CPU-MIB",
            "path": "fiberhome/WRI-CPU-MIB.mib"
        },
        {
            "mib": "RUCKUS-UNLEASHED-SYSTEM-MIB",
            "path": "ruckus/RUCKUS-UNLEASHED-SYSTEM-MIB.mib"
        },
        {
            "mib": "TIMETRA-OAM-TEST-MIB",
            "path": "nokia/TIMETRA-OAM-TEST-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-IPSEC-SA-MON-MIB-EXT",
            "path": "watchguard/WATCHGUARD-IPSEC-SA-MON-MIB-EXT.mib"
        },
        {
            "mib": "CONEL-MOBILE-MIB",
            "path": "icr-os/CONEL-MOBILE-MIB.mib"
        },
        {
            "mib": "SLE-RMON-MIB",
            "path": "dasan/SLE-RMON-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-MAC-MIB",
            "path": "arris/ARRIS-C3-MAC-MIB.mib"
        },
        {
            "mib": "HH3C-FCOE-MIB",
            "path": "comware/HH3C-FCOE-MIB.mib"
        },
        {
            "mib": "SYSTEM-MIB",
            "path": "brocade/SYSTEM-MIB.mib"
        },
        {
            "mib": "LLDP-STD-MIB",
            "path": "raisecom/LLDP-STD-MIB.mib"
        },
        {
            "mib": "ZONE-DEFENSE-MGMT-MIB",
            "path": "dlink/ZONE-DEFENSE-MGMT-MIB.mib"
        },
        {
            "mib": "TN-PTP-MIB",
            "path": "transition/TN-PTP-MIB.mib"
        },
        {
            "mib": "HUAWEI-MSTP-MIB",
            "path": "huawei/HUAWEI-MSTP-MIB.mib"
        },
        {
            "mib": "WLSX-CTS-MIB",
            "path": "arubaos/WLSX-CTS-MIB.mib"
        },
        {
            "mib": "Juniper-Experiment",
            "path": "junose/Juniper-Experiment.mib"
        },
        {
            "mib": "WRI-DEVICE-MIB",
            "path": "fiberhome/WRI-DEVICE-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-SMI",
            "path": "junos/JUNIPER-JS-SMI.mib"
        },
        {
            "mib": "DRAGONWAVE-HCP-MIB",
            "path": "dragonwave/DRAGONWAVE-HCP-MIB.mib"
        },
        {
            "mib": "WRI-MEMORY-MIB",
            "path": "fiberhome/WRI-MEMORY-MIB.mib"
        },
        {
            "mib": "CISCO-PRODUCTS-MIB",
            "path": "cisco/CISCO-PRODUCTS-MIB.mib"
        },
        {
            "mib": "HH3C-FCOE-MODE-MIB",
            "path": "comware/HH3C-FCOE-MODE-MIB.mib"
        },
        {
            "mib": "TN-PRIVATE-VLAN-MIB",
            "path": "transition/TN-PVLAN-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-IPSEC-TUNNEL-MIB",
            "path": "watchguard/WATCHGUARD-IPSEC-TUNNEL-MIB.mib"
        },
        {
            "mib": "RUCKUS-UNLEASHED-WLAN-MIB",
            "path": "ruckus/RUCKUS-UNLEASHED-WLAN-MIB.mib"
        },
        {
            "mib": "HUAWEI-MULTICAST-MIB",
            "path": "huawei/HUAWEI-MULTICAST-MIB.mib"
        },
        {
            "mib": "JUNIPER-JS-UTM-AV-MIB",
            "path": "junos/JUNIPER-JS-UTM-AV-MIB.mib"
        },
        {
            "mib": "IBM-INTERFACE-MIB",
            "path": "ibm/IBM-INTERFACE-MIB.mib"
        },
        {
            "mib": "CONEL-STATUS-MIB",
            "path": "icr-os/CONEL-STATUS-MIB.mib"
        },
        {
            "mib": "WLSX-ESI-MIB",
            "path": "arubaos/WLSX-ESI-MIB.mib"
        },
        {
            "mib": "PEGASUS-LEAN-TRAP-MIB",
            "path": "pegasus/PEGASUS-LEAN-TRAP-MIB.mib"
        },
        {
            "mib": "CISCO-QOS-PIB-MIB",
            "path": "cisco/CISCO-QOS-PIB-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-NTP-MIB",
            "path": "arris/ARRIS-C3-NTP-MIB.mib"
        },
        {
            "mib": "SLE-SECURITY-MIB",
            "path": "dasan/SLE-SECURITY-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-MIB",
            "path": "watchguard/WATCHGUARD-MIB.mib"
        },
        {
            "mib": "HH3C-FDMI-MIB",
            "path": "comware/HH3C-FDMI-MIB.mib"
        },
        {
            "mib": "Juniper-File-Transfer-CONF",
            "path": "junose/Juniper-File-Transfer-CONF.mib"
        },
        {
            "mib": "CAMBIUM-PTP600-MIB",
            "path": "cambium/600/CAMBIUM-PTP600-MIB.mib"
        },
        {
            "mib": "HUAWEI-NAT-EUDM-MIB",
            "path": "huawei/HUAWEI-NAT-EUDM-MIB.mib"
        },
        {
            "mib": "OUTBAND-MGMT-MIB",
            "path": "raisecom/OUTBAND-MGMT-MIB.mib"
        },
        {
            "mib": "DRAGONWAVE-HORIZON-IDU-MIB",
            "path": "dragonwave/DRAGONWAVE-HORIZON-IDU-MIB.mib"
        },
        {
            "mib": "JUNIPER-JVAE-INFRA-MIB",
            "path": "junos/JUNIPER-JVAE-INFRA-MIB.mib"
        },
        {
            "mib": "RUCKUS-ZD-AAA-MIB",
            "path": "ruckus/RUCKUS-ZD-AAA-MIB.mib"
        },
        {
            "mib": "WRI-POWER-MIB",
            "path": "fiberhome/WRI-POWER-MIB.mib"
        },
        {
            "mib": "HH3C-FIREWALL-MIB",
            "path": "comware/HH3C-FIREWALL-MIB.mib"
        },
        {
            "mib": "SLE-SFLOW-MIB",
            "path": "dasan/SLE-SFLOW-MIB.mib"
        },
        {
            "mib": "TN-QOS-EXT-MIB",
            "path": "transition/TN-QOS-EXT.mib"
        },
        {
            "mib": "Juniper-FILE-XFER-MIB",
            "path": "junose/Juniper-FILE-XFER-MIB.mib"
        },
        {
            "mib": "TIMETRA-PORT-MIB",
            "path": "nokia/TIMETRA-PORT-MIB.mib"
        },
        {
            "mib": "WLSX-HA-MIB",
            "path": "arubaos/WLSX-HA-MIB.mib"
        },
        {
            "mib": "CISCO-QOS-POLICY-CONFIG-MIB",
            "path": "cisco/CISCO-QOS-POLICY-CONFIG-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-POST-MIB",
            "path": "arris/ARRIS-C3-POST-MIB.mib"
        },
        {
            "mib": "IBM-LAN-EMULATION-EXTENSION-MIB",
            "path": "ibm/IBM-LAN-EMULATION-EXTENSION-MIB.mib"
        },
        {
            "mib": "JUNIPER-JVAE-NODE-MIB",
            "path": "junos/JUNIPER-JVAE-NODE-MIB.mib"
        },
        {
            "mib": "WRI-SMI",
            "path": "fiberhome/WRI-SMI.mib"
        },
        {
            "mib": "PEGASUS-MIB",
            "path": "pegasus/PEGASUS-MIB.mib"
        },
        {
            "mib": "HH3C-FLASH-MAN-MIB",
            "path": "comware/HH3C-FLASH-MAN-MIB.mib"
        },
        {
            "mib": "SLE-SNMP-MIB",
            "path": "dasan/SLE-SNMP-MIB.mib"
        },
        {
            "mib": "RAISECOM-ACL-MIB",
            "path": "raisecom/RAISECOM-ACL-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-POLICY-MIB",
            "path": "watchguard/WATCHGUARD-POLICY-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-RANGING-MIB",
            "path": "arris/ARRIS-C3-RANGING-MIB.mib"
        },
        {
            "mib": "CONEL-XCCNT-MIB",
            "path": "icr-os/CONEL-XCCNT-MIB.mib"
        },
        {
            "mib": "CISCO-QOS-TC-MIB",
            "path": "cisco/CISCO-QOS-TC-MIB.mib"
        },
        {
            "mib": "Juniper-Fractional-T1-CONF",
            "path": "junose/Juniper-Fractional-T1-CONF.mib"
        },
        {
            "mib": "RUCKUS-ZD-AP-MIB",
            "path": "ruckus/RUCKUS-ZD-AP-MIB.mib"
        },
        {
            "mib": "MWRM-NETWORK-MIB",
            "path": "ceraos/MWRM-NETWORK-MIB.mib"
        },
        {
            "mib": "WLSX-IFEXT-MIB",
            "path": "arubaos/WLSX-IFEXT-MIB.mib"
        },
        {
            "mib": "JUNIPER-L2ALD-MIB",
            "path": "junos/JUNIPER-L2ALD-MIB.mib"
        },
        {
            "mib": "TN-RFC2544-MIB",
            "path": "transition/TN-RFC2544-MIB.mib"
        },
        {
            "mib": "HUAWEI-NAT-MIB",
            "path": "huawei/HUAWEI-NAT-MIB.mib"
        },
        {
            "mib": "RAISECOM-ALARM-MGMT-MIB",
            "path": "raisecom/RAISECOM-ALARM-MGMT-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-PRODUCTS-MIB",
            "path": "watchguard/WATCHGUARD-PRODUCTS-MIB.mib"
        },
        {
            "mib": "DRAGONWAVE-HORIZON-QUANTUM-MIB",
            "path": "dragonwave/DRAGONWAVE-HORIZON-QUANTUM-MIB.mib"
        },
        {
            "mib": "JUNIPER-L2CP-FEATURES-MIB",
            "path": "junos/JUNIPER-L2CP-FEATURES-MIB.mib"
        },
        {
            "mib": "HH3C-FLEXE-MIB",
            "path": "comware/HH3C-FLEXE-MIB.mib"
        },
        {
            "mib": "PEGASUS-SDH-MIB",
            "path": "pegasus/PEGASUS-SDH-MIB.mib"
        },
        {
            "mib": "HUAWEI-ND-MIB",
            "path": "huawei/HUAWEI-ND-MIB.mib"
        },
        {
            "mib": "DWI-HARMONY-PRIVATE-MIB",
            "path": "dragonwave/DWI-HARMONY-PRIVATE-MIB.mib"
        },
        {
            "mib": "MWRM-PM-MIB",
            "path": "ceraos/MWRM-PM-MIB.mib"
        },
        {
            "mib": "WLSX-MESH-MIB",
            "path": "arubaos/WLSX-MESH-MIB.mib"
        },
        {
            "mib": "SLE-SYNCE-MIB",
            "path": "dasan/SLE-SYNCE-MIB.mib"
        },
        {
            "mib": "WRI-TEMPERATURE-MIB",
            "path": "fiberhome/WRI-TEMPERATURE-MIB.mib"
        },
        {
            "mib": "TIMETRA-QOS-MIB",
            "path": "nokia/TIMETRA-QOS-MIB.mib"
        },
        {
            "mib": "Juniper-FRACTIONAL-T1-MIB",
            "path": "junose/Juniper-FRACTIONAL-T1-MIB.mib"
        },
        {
            "mib": "IBM-LES-BUS-MIB",
            "path": "ibm/IBM-LES-BUS-MIB.mib"
        },
        {
            "mib": "CISCO-QUEUE-MIB",
            "path": "cisco/CISCO-QUEUE-MIB.mib"
        },
        {
            "mib": "JUNIPER-LDP-MIB",
            "path": "junos/JUNIPER-LDP-MIB.mib"
        },
        {
            "mib": "RUCKUS-ZD-EVENT-MIB",
            "path": "ruckus/RUCKUS-ZD-EVENT-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-SM-MIB",
            "path": "arris/ARRIS-C3-SM-MIB.mib"
        },
        {
            "mib": "HUAWEI-NDB-MIB",
            "path": "huawei/HUAWEI-NDB-MIB.mib"
        },
        {
            "mib": "EQUIPMENT-COMMON-MIB",
            "path": "dragonwave/EQUIPMENT-COMMON-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-SMI",
            "path": "watchguard/WATCHGUARD-SMI.mib"
        },
        {
            "mib": "HH3C-FLOWTEMPLATE-MIB",
            "path": "comware/HH3C-FLOWTEMPLATE-MIB.mib"
        },
        {
            "mib": "IONODES-IONSERIES-MIB",
            "path": "ionodes/IONODES-IONSERIES-MIB.mib"
        },
        {
            "mib": "RAISECOM-APS-MIB",
            "path": "raisecom/RAISECOM-APS-MIB.mib"
        },
        {
            "mib": "WLSX-MOBILITY-MIB",
            "path": "arubaos/WLSX-MOBILITY-MIB.mib"
        },
        {
            "mib": "Juniper-Frame-Relay-CONF",
            "path": "junose/Juniper-Frame-Relay-CONF.mib"
        },
        {
            "mib": "WRI-VOLTAGE-MIB",
            "path": "fiberhome/WRI-VOLTAGE-MIB.mib"
        },
        {
            "mib": "TN-S-FLOW-MIB",
            "path": "transition/TN-S-FLOW-MIB.mib"
        },
        {
            "mib": "MWRM-RADIO-MIB",
            "path": "ceraos/MWRM-RADIO-MIB.mib"
        },
        {
            "mib": "HORIZON-EQUIPMENT-LOG-MIB",
            "path": "dragonwave/HORIZON-EQUIPMENT-LOG-MIB.mib"
        },
        {
            "mib": "SLE-SYSTEMMAINTENANCE-MIB",
            "path": "dasan/SLE-SYSTEMMAINTENANCE-MIB.mib"
        },
        {
            "mib": "CISCO-REMOTE-ACCESS-MONITOR-MIB",
            "path": "cisco/CISCO-REMOTE-ACCESS-MONITOR-MIB.mib"
        },
        {
            "mib": "SUB10SYSTEMS-MIB",
            "path": "sub10/SUB10SYSTEMS-MIB.mib"
        },
        {
            "mib": "JUNIPER-LICENSE-MIB",
            "path": "junos/JUNIPER-LICENSE-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-SYSTEM-CONFIG-MIB",
            "path": "watchguard/WATCHGUARD-SYSTEM-CONFIG-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-STATS-MIB",
            "path": "arris/ARRIS-C3-STATS-MIB.mib"
        },
        {
            "mib": "HH3C-FR-QOS-MIB",
            "path": "comware/HH3C-FR-QOS-MIB.mib"
        },
        {
            "mib": "RAISECOM-ARP-MIB",
            "path": "raisecom/RAISECOM-ARP-MIB.mib"
        },
        {
            "mib": "RUCKUS-ZD-SYSTEM-MIB",
            "path": "ruckus/RUCKUS-ZD-SYSTEM-MIB.mib"
        },
        {
            "mib": "MPBC-2RU-MIB",
            "path": "mpb/MPBC-2RU-MIB.mib"
        },
        {
            "mib": "HUAWEI-NETCONF-MIB",
            "path": "huawei/HUAWEI-NETCONF-MIB.mib"
        },
        {
            "mib": "WLSX-MON-MIB",
            "path": "arubaos/WLSX-MON-MIB.mib"
        },
        {
            "mib": "EATON-ATS2-MIB",
            "path": "eaton/EATON-ATS2-MIB.mib"
        },
        {
            "mib": "TIMETRA-SAP-MIB",
            "path": "nokia/TIMETRA-SAP-MIB.mib"
        },
        {
            "mib": "EIP-DNSBLAST-MIB",
            "path": "efficientip/EIP-DNSBLAST-MIB.mib"
        },
        {
            "mib": "SLE-TC-MIB",
            "path": "dasan/SLE-TC-MIB.mib"
        },
        {
            "mib": "HORIZON-MIB",
            "path": "dragonwave/HORIZON-MIB.mib"
        },
        {
            "mib": "TN-SA-MIB",
            "path": "transition/TN-SA-MIB.mib"
        },
        {
            "mib": "EATON-EMP-MIB",
            "path": "eaton/EATON-EMP-MIB.mib"
        },
        {
            "mib": "Juniper-FRAME-RELAY-MIB",
            "path": "junose/Juniper-FRAME-RELAY-MIB.mib"
        },
        {
            "mib": "WATCHGUARD-SYSTEM-STATISTICS-MIB",
            "path": "watchguard/WATCHGUARD-SYSTEM-STATISTICS-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-SUBINT-MIB",
            "path": "arris/ARRIS-C3-SUBINT-MIB.mib"
        },
        {
            "mib": "TIMETRA-SDP-MIB",
            "path": "nokia/TIMETRA-SDP-MIB.mib"
        },
        {
            "mib": "ZMTEL-ODU-MIB",
            "path": "zmtel/ZMTEL-ODU-MIB.mib"
        },
        {
            "mib": "RUCKUS-ZD-WLAN-CONFIG-MIB",
            "path": "ruckus/RUCKUS-ZD-WLAN-CONFIG-MIB.mib"
        },
        {
            "mib": "TN-SAT-LOOPBACK-MIB",
            "path": "transition/TN-SAT-LOOPBACK-MIB.mib"
        },
        {
            "mib": "HH3C-FTM-MIB",
            "path": "comware/HH3C-FTM-MIB.mib"
        },
        {
            "mib": "MWRM-UNIT-MIB",
            "path": "ceraos/MWRM-UNIT-MIB.mib"
        },
        {
            "mib": "CISCO-RESILIENT-ETHERNET-PROTOCOL-MIB",
            "path": "cisco/CISCO-RESILIENT-ETHERNET-PROTOCOL-MIB.mib"
        },
        {
            "mib": "WLSX-RS-MIB",
            "path": "arubaos/WLSX-RS-MIB.mib"
        },
        {
            "mib": "HORIZON-ODU-MIB",
            "path": "dragonwave/HORIZON-ODU-MIB.mib"
        },
        {
            "mib": "IBM-LES-LECS-MIB",
            "path": "ibm/IBM-LES-LECS-MIB.mib"
        },
        {
            "mib": "EIP-MON-MIB",
            "path": "efficientip/EIP-MON-MIB.mib"
        },
        {
            "mib": "RAISECOM-AUTOPROVISIONMDEV-MIB",
            "path": "raisecom/RAISECOM-AUTOPROVISIONMDEV-MIB.mib"
        },
        {
            "mib": "SLE-VOIP-MIB",
            "path": "dasan/SLE-VOIP-MIB.mib"
        },
        {
            "mib": "Juniper-Ha-Redundancy-CONF",
            "path": "junose/Juniper-Ha-Redundancy-CONF.mib"
        },
        {
            "mib": "HUAWEI-NETSTREAM-MIB",
            "path": "huawei/HUAWEI-NETSTREAM-MIB.mib"
        },
        {
            "mib": "WLSX-SNR-MIB",
            "path": "arubaos/WLSX-SNR-MIB.mib"
        },
        {
            "mib": "TN-SECURITY-AAA-MIB",
            "path": "transition/TN-SECURITY-AAA-MIB.mib"
        },
        {
            "mib": "RUCKUS-ZD-WLAN-MIB",
            "path": "ruckus/RUCKUS-ZD-WLAN-MIB.mib"
        },
        {
            "mib": "ARRIS-C3-TFTPD-MIB",
            "path": "arris/ARRIS-C3-TFTPD-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYS-SECURITYPROFILE-MIB",
            "path": "junos/JUNIPER-LSYS-SECURITYPROFILE-MIB.mib"
        },
        {
            "mib": "EATON-EPDU-MIB",
            "path": "eaton/EATON-EPDU-MIB.mib"
        },
        {
            "mib": "EIP-STATS-MIB",
            "path": "efficientip/EIP-STATS-MIB.mib"
        },
        {
            "mib": "HH3C-GOLD-MIB",
            "path": "comware/HH3C-GOLD-MIB.mib"
        },
        {
            "mib": "FOURELLE-SMI",
            "path": "venturi/FOURELLE-SMI.mib"
        },
        {
            "mib": "HUAWEI-NTP-MIB",
            "path": "huawei/HUAWEI-NTP-MIB.mib"
        },
        {
            "mib": "RAISECOM-AUTOPROVISIONRDEV-MIB",
            "path": "raisecom/RAISECOM-AUTOPROVISIONRDEV-MIB.mib"
        },
        {
            "mib": "MWR-ETHERNET-MIB",
            "path": "dragonwave/MWR-ETHERNET-MIB.mib"
        },
        {
            "mib": "EATON-OIDS",
            "path": "eaton/EATON-OIDS.mib"
        },
        {
            "mib": "TN-SECURITY-SWITCH-SSH-MIB",
            "path": "transition/TN-SECURITY-SWITCH-SSH-MIB.mib"
        },
        {
            "mib": "Juniper-HDLC-CONF",
            "path": "junose/Juniper-HDLC-CONF.mib"
        },
        {
            "mib": "VEL-HOST2-MIB",
            "path": "vigintos/VEL-HOST2-MIB.mib"
        },
        {
            "mib": "HH3C-GRE-MIB",
            "path": "comware/HH3C-GRE-MIB.mib"
        },
        {
            "mib": "CISCO-RF-MIB",
            "path": "cisco/CISCO-RF-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-CPU-MIB",
            "path": "junos/JUNIPER-LSYSSP-CPU-MIB.mib"
        },
        {
            "mib": "IBM-MIB",
            "path": "ibm/IBM-MIB.mib"
        },
        {
            "mib": "TIC-RMTI4-G9000-G2020-MIB",
            "path": "toshiba/TIC-RMTI4-G9000-G2020-MIB.mib"
        },
        {
            "mib": "DHCP-MIB",
            "path": "microsoft/DHCP-MIB.mib"
        },
        {
            "mib": "RAISECOM-BANNER-MIB",
            "path": "raisecom/RAISECOM-BANNER-MIB.mib"
        },
        {
            "mib": "MWR-RADIO-MC-MIB",
            "path": "dragonwave/MWR-RADIO-MC-MIB.mib"
        },
        {
            "mib": "EATON-SENSOR-MIB",
            "path": "eaton/EATON-SENSOR-MIB.mib"
        },
        {
            "mib": "TN-SIP-MIB",
            "path": "transition/TN-SIP-MIB.smi.mib"
        },
        {
            "mib": "CISCO-RTTMON-ICMP-MIB",
            "path": "cisco/CISCO-RTTMON-ICMP-MIB.mib"
        },
        {
            "mib": "ASENTRIA-ROOT-MIB",
            "path": "asentria/ASENTRIA-ROOT-MIB.mib"
        },
        {
            "mib": "FOURELLE-VENTURI2-MIB",
            "path": "venturi/VENTURI-REL2-MIB.mib"
        },
        {
            "mib": "VEL-MIB",
            "path": "vigintos/VEL-MIB.mib"
        },
        {
            "mib": "ARRIS-CM-CAPABILITY-MIB",
            "path": "arris/ARRIS-CM-CAPABILITY-MIB.mib"
        },
        {
            "mib": "WLSX-STATS-MIB",
            "path": "arubaos/WLSX-STATS-MIB.mib"
        },
        {
            "mib": "SLE-VRRP-MIB",
            "path": "dasan/SLE-VRRP-MIB.mib"
        },
        {
            "mib": "FROGFOOT-RESOURCES-MIB",
            "path": "frogfoot/FROGFOOT-RESOURCES-MIB.mib"
        },
        {
            "mib": "HH3C-HGMP-MIB",
            "path": "comware/HH3C-HGMP-MIB.mib"
        },
        {
            "mib": "HUAWEI-NTP-TRAP-MIB",
            "path": "huawei/HUAWEI-NTP-TRAP-MIB.mib"
        },
        {
            "mib": "PM8ECCMIB",
            "path": "schneider/PM8ECCMIB.mib"
        },
        {
            "mib": "RAISECOM-BASE-MIB",
            "path": "raisecom/RAISECOM-BASE-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-FLOWGATE-MIB",
            "path": "junos/JUNIPER-LSYSSP-FLOWGATE-MIB.mib"
        },
        {
            "mib": "MSFT-MIB",
            "path": "microsoft/MSFT-MIB.mib"
        },
        {
            "mib": "ARRIS-CM-DEVICE-MIB",
            "path": "arris/ARRIS-CM-DEVICE-MIB.mib"
        },
        {
            "mib": "NEXANS-BM-MIB",
            "path": "nexans/NEXANS-BM-MIB.mib"
        },
        {
            "mib": "Juniper-HDLC-MIB",
            "path": "junose/Juniper-HDLC-MIB.mib"
        },
        {
            "mib": "CISCO-RTTMON-MIB",
            "path": "cisco/CISCO-RTTMON-MIB.mib"
        },
        {
            "mib": "TN-SYNCE-MIB",
            "path": "transition/TN-SYNCE-MIB.mib"
        },
        {
            "mib": "WLSX-SWITCH-MIB",
            "path": "arubaos/WLSX-SWITCH-MIB.mib"
        },
        {
            "mib": "HUAWEI-NVO3-MIB",
            "path": "huawei/HUAWEI-NVO3-MIB.mib"
        },
        {
            "mib": "VENTURI-SERVER-CONFIG-MIB",
            "path": "venturi/VENTURI-SERVER-CONFIG-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-FLOWSESS-MIB",
            "path": "junos/JUNIPER-LSYSSP-FLOWSESS-MIB.mib"
        },
        {
            "mib": "SLEV2-DHCP-MIB",
            "path": "dasan/SLEV2-DHCP-MIB.mib"
        },
        {
            "mib": "MG-SNMP-UPS-MIB",
            "path": "eaton/MG-SNMP-UPS-MIB.mib"
        },
        {
            "mib": "HH3C-HPEOC-MIB",
            "path": "comware/HH3C-HPEOC-MIB.mib"
        },
        {
            "mib": "TRIPPLITE-12X",
            "path": "poweralert/TRIPPLITE-12X.mib"
        },
        {
            "mib": "ARUBAWIRED-AAA-MIB",
            "path": "arubaos-cx/ARUBAWIRED-AAA-MIB.mib"
        },
        {
            "mib": "ARRIS-CMTS-FFT-MIB",
            "path": "arris/ARRIS-CMTS-FFT-MIB.mib"
        },
        {
            "mib": "CISCO-RTTMON-TC-MIB",
            "path": "cisco/CISCO-RTTMON-TC-MIB.mib"
        },
        {
            "mib": "SITEBOSS-360-STD-MIB",
            "path": "asentria/SITEBOSS-360-STD-MIB.mib"
        },
        {
            "mib": "TIMETRA-SERV-MIB",
            "path": "nokia/TIMETRA-SERV-MIB.mib"
        },
        {
            "mib": "RAISECOM-BFD-MIB",
            "path": "raisecom/RAISECOM-BFD-MIB.mib"
        },
        {
            "mib": "VENTURI-SERVER-MIB",
            "path": "venturi/VENTURI-SERVER-MIB.mib"
        },
        {
            "mib": "Juniper-HOST-MIB",
            "path": "junose/Juniper-HOST-MIB.mib"
        },
        {
            "mib": "TN-SYS-LOG-MIB",
            "path": "transition/TN-SYS-LOG-MIB.mib"
        },
        {
            "mib": "HUAWEI-OPENFLOW-MIB",
            "path": "huawei/HUAWEI-OPENFLOW-MIB.mib"
        },
        {
            "mib": "ARRIS-MIB",
            "path": "arris/ARRIS-MIB.mib"
        },
        {
            "mib": "NEXANS-MIB",
            "path": "nexans/NEXANS-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-CHASSIS-MIB",
            "path": "arubaos-cx/ARUBAWIRED-CHASSIS-MIB.mib"
        },
        {
            "mib": "WLSX-SYSTEMEXT-MIB",
            "path": "arubaos/WLSX-SYSTEMEXT-MIB.mib"
        },
        {
            "mib": "IBM-NetFinity-Text-Alert-MIB",
            "path": "ibm/IBM-NetFinity-Text-Alert-MIB.mib"
        },
        {
            "mib": "TRIPPLITE-MIB",
            "path": "poweralert/TRIPPLITE-MIB.mib"
        },
        {
            "mib": "PDU-MIB",
            "path": "eaton/PDU-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATCONEBIND-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATCONEBIND-MIB.mib"
        },
        {
            "mib": "RAISECOM-CFM-MIB",
            "path": "raisecom/RAISECOM-CFM-MIB.mib"
        },
        {
            "mib": "Juniper-HTTP-Profile-MIB",
            "path": "junose/Juniper-HTTP-Profile-MIB.mib"
        },
        {
            "mib": "VENTURI-SERVER-STATS-MIB",
            "path": "venturi/VENTURI-SERVER-STATS-MIB.mib"
        },
        {
            "mib": "HH3C-IDS-MIB",
            "path": "comware/HH3C-IDS-MIB.mib"
        },
        {
            "mib": "SLEV2-EPON-IM-MIB",
            "path": "dasan/SLEV2-EPON-IM-MIB.mib"
        },
        {
            "mib": "TN-TC",
            "path": "transition/TN-TC.mib"
        },
        {
            "mib": "CISCO-SLB-EXT-MIB",
            "path": "cisco/CISCO-SLB-EXT-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-CIPT-MIB",
            "path": "arubaos-cx/ARUBAWIRED-CIPT-MIB.mib"
        },
        {
            "mib": "HUAWEI-OSPFV2-MIB",
            "path": "huawei/HUAWEI-OSPFV2-MIB.mib"
        },
        {
            "mib": "BCS-IDENT-MIB",
            "path": "arris/BCS-IDENT-MIB.mib"
        },
        {
            "mib": "RAISECOM-COMMON-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-COMMON-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "ELTEX-MES-HWENVIROMENT-MIB",
            "path": "eltexmes21xx/ELTEX-MES-HWENVIROMENT-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATDSTPOOL-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATDSTPOOL-MIB.mib"
        },
        {
            "mib": "SITEBOSS-530-STD-MIB",
            "path": "asentria/SITEBOSS-530-STD-MIB.mib"
        },
        {
            "mib": "CORIANT-GROOVE-MIB",
            "path": "infinera/CORIANT-GROOVE-MIB.mib"
        },
        {
            "mib": "WLSX-TRAP-MIB",
            "path": "arubaos/WLSX-TRAP-MIB.mib"
        },
        {
            "mib": "TN-THERMAL-PROTECTION-MIB",
            "path": "transition/TN-THERMAL-PROTECTION-MIB.mib"
        },
        {
            "mib": "CISCO-SLB-HEALTH-MON-MIB",
            "path": "cisco/CISCO-SLB-HEALTH-MON-MIB.mib"
        },
        {
            "mib": "RPS-SC200-MIB",
            "path": "eaton/RPS-SC200-MIB.mib"
        },
        {
            "mib": "Juniper-IGMP-MIB",
            "path": "junose/Juniper-IGMP-MIB.mib"
        },
        {
            "mib": "VENTURI-SERVER-SYSTEM-MIB",
            "path": "venturi/VENTURI-SERVER-SYSTEM-MIB.mib"
        },
        {
            "mib": "ELTEX-MES",
            "path": "eltexmes21xx/ELTEX-MES.mib"
        },
        {
            "mib": "RAISECOM-DHCP-CLIENT-MIB",
            "path": "raisecom/RAISECOM-DHCP-CLIENT-MIB.mib"
        },
        {
            "mib": "LUM-EQUIPMENT-MIB",
            "path": "infinera/LUM-EQUIPMENT-MIB.mib"
        },
        {
            "mib": "TRIPPLITE-PRODUCTS",
            "path": "poweralert/TRIPPLITE-PRODUCTS.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATDSTRULE-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATDSTRULE-MIB.mib"
        },
        {
            "mib": "HH3C-IF-EXT-MIB",
            "path": "comware/HH3C-IF-EXT-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-CONFIG-MIB",
            "path": "arubaos-cx/ARUBAWIRED-CONFIG-MIB.mib"
        },
        {
            "mib": "IBM-OSA-MIB",
            "path": "ibm/IBM-OSA-MIB.mib"
        },
        {
            "mib": "TN-TT-LOOP-MIB",
            "path": "transition/TN-TT-LOOP-MIB.mib"
        },
        {
            "mib": "WLSX-TUNNELEDNODE-MIB",
            "path": "arubaos/WLSX-TUNNELEDNODE-MIB.mib"
        },
        {
            "mib": "BCS-TRAPS-MIB",
            "path": "arris/BCS-TRAPS-MIB.mib"
        },
        {
            "mib": "VENTURI-SERVER-TRAP-MIB",
            "path": "venturi/VENTURI-SERVER-TRAP-MIB.mib"
        },
        {
            "mib": "SLEV2-MULTICAST-MIB",
            "path": "dasan/SLEV2-MULTICAST-MIB.mib"
        },
        {
            "mib": "RpsSc300Mib",
            "path": "eaton/RPS-SC300-MIB.mib"
        },
        {
            "mib": "HH3C-IFQOS2-MIB",
            "path": "comware/HH3C-IFQOS2-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-DIST-SERVICES-MIB",
            "path": "arubaos-cx/ARUBAWIRED-DIST-SERVICES-MIB.mib"
        },
        {
            "mib": "ELTEX-SMI-ACTUAL",
            "path": "eltexmes21xx/ELTEX-SMI-ACTUAL.mib"
        },
        {
            "mib": "SITEBOSS-550-STD-MIB",
            "path": "asentria/SITEBOSS-550-STD-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATPOIPNUM-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATPOIPNUM-MIB.mib"
        },
        {
            "mib": "TRIPPLITE",
            "path": "poweralert/TRIPPLITE.mib"
        },
        {
            "mib": "HUAWEI-OSPFV3-MIB",
            "path": "huawei/HUAWEI-OSPFV3-MIB.mib"
        },
        {
            "mib": "Juniper-IKE-CONF",
            "path": "junose/Juniper-IKE-CONF.mib"
        },
        {
            "mib": "VENTURI-TC",
            "path": "venturi/VENTURI-TC.mib"
        },
        {
            "mib": "LUM-IFBASIC-MIB",
            "path": "infinera/LUM-IFBASIC-MIB.mib"
        },
        {
            "mib": "CADANT-CMTS-EQUIPMENT-MIB",
            "path": "arris/CADANT-CMTS-EQUIPMENT-MIB.mib"
        },
        {
            "mib": "TN-UDLD-MIB",
            "path": "transition/TN-UDLD.mib"
        },
        {
            "mib": "TELECOM-MIB",
            "path": "eaton/TELECOM-MIB.mib"
        },
        {
            "mib": "HH3C-IKE-MONITOR-MIB",
            "path": "comware/HH3C-IKE-MONITOR-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-FAN-MIB",
            "path": "arubaos-cx/ARUBAWIRED-FAN-MIB.mib"
        },
        {
            "mib": "CADANT-PRODUCTS-MIB",
            "path": "arris/CADANT-PRODUCTS-MIB.mib"
        },
        {
            "mib": "TIMETRA-SYSTEM-MIB",
            "path": "nokia/TIMETRA-SYSTEM-MIB.mib"
        },
        {
            "mib": "CISCO-SLB-MIB",
            "path": "cisco/CISCO-SLB-MIB.mib"
        },
        {
            "mib": "TN-VLAN-MGMT-MIB",
            "path": "transition/TN-VLAN-MGMT-MIB.mib"
        },
        {
            "mib": "TIMETRA-SUBSCRIBER-MGMT-MIB",
            "path": "nokia/TIMETRA-SUBSCRIBER-MGMT-MIB.mib"
        },
        {
            "mib": "WLSX-USER-MIB",
            "path": "arubaos/WLSX-USER-MIB.mib"
        },
        {
            "mib": "VENTURI-WIRELESS-SMI",
            "path": "venturi/VENTURI-WIRELESS-SMI.mib"
        },
        {
            "mib": "SITEBOSS-571-STD-MIB",
            "path": "asentria/SITEBOSS-571-STD-MIB.mib"
        },
        {
            "mib": "SLEV2-PPPoE-MIB",
            "path": "dasan/SLEV2-PPPoE-MIB.mib"
        },
        {
            "mib": "NBS-CMMC-MIB",
            "path": "mrv/NBS-CMMC-MIB.mib"
        },
        {
            "mib": "IBM-SERVERAID-MIB",
            "path": "ibm/IBM-SERVERAID-MIB.mib"
        },
        {
            "mib": "HH3C-INFOCENTER-MIB",
            "path": "comware/HH3C-INFOCENTER-MIB.mib"
        },
        {
            "mib": "TIMETRA-TC-MG-MIB",
            "path": "nokia/TIMETRA-TC-MG-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP-OPTION-MIB",
            "path": "raisecom/RAISECOM-DHCP-OPTION-MIB.mib"
        },
        {
            "mib": "RADLAN-DEVICEPARAMS-MIB",
            "path": "eltexmes21xx/RADLAN-DEVICEPARAMS-MIB.mib"
        },
        {
            "mib": "LUM-IFOTN-MIB",
            "path": "infinera/LUM-IFOTN-MIB.mib"
        },
        {
            "mib": "Juniper-IKE-MIB",
            "path": "junose/Juniper-IKE-MIB.mib"
        },
        {
            "mib": "XUPS-MIB",
            "path": "eaton/XUPS-MIB.mib"
        },
        {
            "mib": "HUAWEI-PERFMGMT-MIB",
            "path": "huawei/HUAWEI-PERFMGMT-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATSRCNOPATAD-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATSRCNOPATAD-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-FANTRAY-MIB",
            "path": "arubaos-cx/ARUBAWIRED-FANTRAY-MIB.mib"
        },
        {
            "mib": "TN-XSTP-MIB",
            "path": "transition/TN-XSTP-MIB.mib"
        },
        {
            "mib": "CISCO-SMI",
            "path": "cisco/CISCO-SMI.mib"
        },
        {
            "mib": "CADANT-TC",
            "path": "arris/CADANT-TC-MIB.mib"
        },
        {
            "mib": "WLSX-USER6-MIB",
            "path": "arubaos/WLSX-USER6-MIB.mib"
        },
        {
            "mib": "HH3C-IP-ADDRESS-MIB",
            "path": "comware/HH3C-IP-ADDRESS-MIB.mib"
        },
        {
            "mib": "RADLAN-File",
            "path": "eltexmes21xx/RADLAN-File.mib"
        },
        {
            "mib": "RAISECOM-DHCP-SNOOPING-MIB",
            "path": "raisecom/RAISECOM-DHCP-SNOOPING-MIB.mib"
        },
        {
            "mib": "CISCO-SNAPSHOT-MIB",
            "path": "cisco/CISCO-SNAPSHOT-MIB.mib"
        },
        {
            "mib": "HH3C-IP-BROADCAST-MIB",
            "path": "comware/HH3C-IP-BROADCAST-MIB.mib"
        },
        {
            "mib": "XPPC-MIB",
            "path": "phoenixtec/XPPC-MIB.mib"
        },
        {
            "mib": "Juniper-Interfaces-CONF",
            "path": "junose/Juniper-Interfaces-CONF.mib"
        },
        {
            "mib": "RAISECOM-DHCP6-CLIENT-MIB",
            "path": "raisecom/RAISECOM-DHCP6-CLIENT-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-INTERFACE-MIB",
            "path": "arubaos-cx/ARUBAWIRED-INTERFACE-MIB.mib"
        },
        {
            "mib": "ECRESO-FM-TRANS-MIB",
            "path": "worldcastsystems/ECRESO-FM-TRANS-MIB.mib"
        },
        {
            "mib": "WLSX-VOICE-MIB",
            "path": "arubaos/WLSX-VOICE-MIB.mib"
        },
        {
            "mib": "TN-Y1564-MIB",
            "path": "transition/TN-Y1564-MIB.mib"
        },
        {
            "mib": "RADLAN-HWENVIROMENT",
            "path": "eltexmes21xx/RADLAN-HWENVIROMENT.mib"
        },
        {
            "mib": "SLEV2-QOS-MIB",
            "path": "dasan/SLEV2-QOS-MIB.mib"
        },
        {
            "mib": "POLYCOM-BASE-MIB",
            "path": "polycom/POLYCOM-BASE-MIB.mib"
        },
        {
            "mib": "DOCS-BPI2-MIB",
            "path": "arris/DOCS-BPI2-MIB.mib"
        },
        {
            "mib": "HH3C-IPA-MIB",
            "path": "comware/HH3C-IPA-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATSRCPATAD-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATSRCPATAD-MIB.mib"
        },
        {
            "mib": "NBS-CMMCENUM-MIB",
            "path": "mrv/NBS-CMMCENUM-MIB.mib"
        },
        {
            "mib": "LUM-IFOTNMON-MIB",
            "path": "infinera/LUM-IFOTNMON-MIB.mib"
        },
        {
            "mib": "Juniper-IP-Policy-CONF",
            "path": "junose/Juniper-IP-Policy-CONF.mib"
        },
        {
            "mib": "SITEBOSS-572-STD-MIB",
            "path": "asentria/SITEBOSS-572-STD-MIB.mib"
        },
        {
            "mib": "TN-ZERO-TOUCH-PROVISION-MIB",
            "path": "transition/TN-ZERO-TOUCH-PROVISION-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP6-RELAY-MIB",
            "path": "raisecom/RAISECOM-DHCP6-RELAY-MIB.mib"
        },
        {
            "mib": "CELLULAR",
            "path": "peplink/CELLULAR.mib"
        },
        {
            "mib": "RADLAN-MIB",
            "path": "eltexmes21xx/RADLAN-MIB.mib"
        },
        {
            "mib": "HUAWEI-PERFORMANCE-MIB",
            "path": "huawei/HUAWEI-PERFORMANCE-MIB.mib"
        },
        {
            "mib": "HH3C-IPRAN-DCN-MIB",
            "path": "comware/HH3C-IPRAN-DCN-MIB.mib"
        },
        {
            "mib": "SLEV2-Security-MIB",
            "path": "dasan/SLEV2-Security-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATSRCPOOL-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATSRCPOOL-MIB.mib"
        },
        {
            "mib": "DOCS-CABLE-DEVICE-TRAP-MIB",
            "path": "arris/DOCS-CABLE-DEVICE-TRAP-MIB.mib"
        },
        {
            "mib": "WLSX-WLAN-MIB",
            "path": "arubaos/WLSX-WLAN-MIB.mib"
        },
        {
            "mib": "TRANSITION-SMI",
            "path": "transition/TRANSITION-SMI.mib"
        },
        {
            "mib": "RADLAN-PHY-MIB",
            "path": "eltexmes21xx/RADLAN-PHY-MIB.mib"
        },
        {
            "mib": "IBM-SVC-MIB",
            "path": "ibm/IBM-SVC-MIB.mib"
        },
        {
            "mib": "RAISECOM-DHCP6-SERVER-MIB",
            "path": "raisecom/RAISECOM-DHCP6-SERVER-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-LED-LOCATOR-MIB",
            "path": "arubaos-cx/ARUBAWIRED-LED-LOCATOR-MIB.mib"
        },
        {
            "mib": "ENVIROMUX-1W-MIB",
            "path": "nti/ENVIROMUX-1W-MIB.mib"
        },
        {
            "mib": "POLYCOM-ENDPOINT-MIB",
            "path": "polycom/POLYCOM-ENDPOINT-MIB.mib"
        },
        {
            "mib": "SLEV2-SNMP-MIB",
            "path": "dasan/SLEV2-SNMP-MIB.mib"
        },
        {
            "mib": "LUM-IFPERF-MIB",
            "path": "infinera/LUM-IFPERF-MIB.mib"
        },
        {
            "mib": "DEVICE",
            "path": "peplink/DEVICE.mib"
        },
        {
            "mib": "NBS-COHERENT-MIB",
            "path": "mrv/NBS-COHERENT-MIB.mib"
        },
        {
            "mib": "RADLAN-Physicaldescription-MIB",
            "path": "eltexmes21xx/RADLAN-Physicaldescription-MIB.mib"
        },
        {
            "mib": "CISCO-STACK-MIB",
            "path": "cisco/CISCO-STACK-MIB.mib"
        },
        {
            "mib": "Juniper-IP-POLICY-MIB",
            "path": "junose/Juniper-IP-POLICY-MIB.mib"
        },
        {
            "mib": "RAISECOM-DOT1AG-MIB",
            "path": "raisecom/RAISECOM-DOT1AG-MIB.mib"
        },
        {
            "mib": "TIMETRA-TC-MIB",
            "path": "nokia/TIMETRA-TC-MIB.mib"
        },
        {
            "mib": "ENVIROMUXMICRO-MIB",
            "path": "nti/ENVIROMUXMICRO-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-LLDP-MIB",
            "path": "arubaos-cx/ARUBAWIRED-LLDP-MIB.mib"
        },
        {
            "mib": "LINKSYS-3SW2SWTABLES-MIB",
            "path": "linksys/LINKSYS-3SW2SWTABLES-MIB.mib"
        },
        {
            "mib": "DOCS-QOS-MIB",
            "path": "arris/DOCS-QOS-MIB.mib"
        },
        {
            "mib": "NBS-CONNECTIVITY-MIB",
            "path": "mrv/NBS-CONNECTIVITY-MIB.mib"
        },
        {
            "mib": "CISCO-STACKWISE-MIB",
            "path": "cisco/CISCO-STACKWISE-MIB.mib"
        },
        {
            "mib": "TRANSITION-TC",
            "path": "transition/TRANSITION-TC.mib"
        },
        {
            "mib": "NAS-MIB",
            "path": "qnap/NAS-MIB.mib"
        },
        {
            "mib": "HH3C-IPSEC-MONITOR-MIB",
            "path": "comware/HH3C-IPSEC-MONITOR-MIB.mib"
        },
        {
            "mib": "RADLAN-rndMng",
            "path": "eltexmes21xx/RADLAN-rndMng.mib"
        },
        {
            "mib": "IBM-TN3270E-MIB",
            "path": "ibm/IBM-TN3270E-MIB.mib"
        },
        {
            "mib": "HUAWEI-PFLT-EUDM-MIB",
            "path": "huawei/HUAWEI-PFLT-EUDM-MIB.mib"
        },
        {
            "mib": "CISCO-STP-EXTENSIONS-MIB",
            "path": "cisco/CISCO-STP-EXTENSIONS-MIB.mib"
        },
        {
            "mib": "NBS-EFM-MIB",
            "path": "mrv/NBS-EFM-MIB.mib"
        },
        {
            "mib": "HH3C-IPSEC-MONITOR-V2-MIB",
            "path": "comware/HH3C-IPSEC-MONITOR-V2-MIB.mib"
        },
        {
            "mib": "LINKSYS-AAA",
            "path": "linksys/LINKSYS-AAA.mib"
        },
        {
            "mib": "GRE",
            "path": "peplink/GRE.mib"
        },
        {
            "mib": "Juniper-IP-Profile-CONF",
            "path": "junose/Juniper-IP-Profile-CONF.mib"
        },
        {
            "mib": "DCS-MIB",
            "path": "vertiv/DCS-MIB.mib"
        },
        {
            "mib": "RAISECOM-DOT1X-MIB",
            "path": "raisecom/RAISECOM-DOT1X-MIB.mib"
        },
        {
            "mib": "TIMETRA-VRTR-MIB",
            "path": "nokia/TIMETRA-VRTR-MIB.mib"
        },
        {
            "mib": "RITTAL-CMC-III-CAPABILITY-MIB",
            "path": "rittal/RITTAL-CMC-III-CAPABILITY-MIB.mib"
        },
        {
            "mib": "HUAWEI-PGI-MIB",
            "path": "huawei/HUAWEI-PGI-MIB.mib"
        },
        {
            "mib": "CISCO-SWITCH-ENGINE-MIB",
            "path": "cisco/CISCO-SWITCH-ENGINE-MIB.mib"
        },
        {
            "mib": "PRVT-ALARM-MIB",
            "path": "telco-systems/binox/PRVT-ALARM-MIB.mib"
        },
        {
            "mib": "HH3C-IPV6-ADDRESS-MIB",
            "path": "comware/HH3C-IPV6-ADDRESS-MIB.mib"
        },
        {
            "mib": "VERTIV-ITA2-MIB",
            "path": "vertiv/VERTIV-ITA2-MIB.mib"
        },
        {
            "mib": "DOCS-SUBMGT-MIB",
            "path": "arris/DOCS-SUBMGT-MIB.mib"
        },
        {
            "mib": "LINKSYS-BANNER-MIB",
            "path": "linksys/LINKSYS-BANNER-MIB.mib"
        },
        {
            "mib": "NBS-EUSM-MIB",
            "path": "mrv/NBS-EUSM-MIB.mib"
        },
        {
            "mib": "RITTAL-CMC-III-MIB",
            "path": "rittal/RITTAL-CMC-III-MIB.mib"
        },
        {
            "mib": "IPSEC-VPN",
            "path": "peplink/IPSEC-VPN.mib"
        },
        {
            "mib": "PERLE-IOLAN-SDS-MIB",
            "path": "perle/PERLE-IOLAN-SDS-MIB.mib"
        },
        {
            "mib": "RAISECOM-ELMI-MIB",
            "path": "raisecom/RAISECOM-ELMI-MIB.mib"
        },
        {
            "mib": "HH3C-IPX-MIB",
            "path": "comware/HH3C-IPX-MIB.mib"
        },
        {
            "mib": "CISCO-SWITCH-QOS-MIB",
            "path": "cisco/CISCO-SWITCH-QOS-MIB.mib"
        },
        {
            "mib": "IBM-TS3500-MIBv1",
            "path": "ibm/IBM-TS3500-MIBv1.mib"
        },
        {
            "mib": "LUM-IFPHYSICAL-MIB",
            "path": "infinera/LUM-IFPHYSICAL-MIB.mib"
        },
        {
            "mib": "Juniper-IP-PROFILE-MIB",
            "path": "junose/Juniper-IP-PROFILE-MIB.mib"
        },
        {
            "mib": "INFINERA-ENTITY-CHASSIS-MIB",
            "path": "iqnos/INFINERA-ENTITY-CHASSIS.mib"
        },
        {
            "mib": "HUAWEI-PIM-BSR-MIB",
            "path": "huawei/HUAWEI-PIM-BSR-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-LOOPPROTECT-MIB",
            "path": "arubaos-cx/ARUBAWIRED-LOOPPROTECT-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATSRCRULE-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATSRCRULE-MIB.mib"
        },
        {
            "mib": "DOCS-TEST-MIB",
            "path": "arris/DOCS-TEST-MIB.mib"
        },
        {
            "mib": "VERTIV-V5-MIB",
            "path": "vertiv/VERTIV-V5-MIB.mib"
        },
        {
            "mib": "PRVT-CFM-MIB",
            "path": "telco-systems/binox/PRVT-CFM-MIB.mib"
        },
        {
            "mib": "MPLS-LSR-MIB",
            "path": "telco-systems/binos/MPLS-LSR-MIB.mib"
        },
        {
            "mib": "RITTAL-CMC-III-PRODUCTS-MIB",
            "path": "rittal/RITTAL-CMC-III-PRODUCTS-MIB.mib"
        },
        {
            "mib": "RAISECOM-ELPS-MIB",
            "path": "raisecom/RAISECOM-ELPS-MIB.mib"
        },
        {
            "mib": "LINKSYS-BaudRate-MIB",
            "path": "linksys/LINKSYS-BaudRate-MIB.mib"
        },
        {
            "mib": "HH3C-ISDN-MIB",
            "path": "comware/HH3C-ISDN-MIB.mib"
        },
        {
            "mib": "NBS-FAN-MIB",
            "path": "mrv/NBS-FAN-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-MACNOTIFY-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MACNOTIFY-MIB.mib"
        },
        {
            "mib": "CISCO-SYSLOG-MIB",
            "path": "cisco/CISCO-SYSLOG-MIB.mib"
        },
        {
            "mib": "SEAGATESYSTEMTRAP-MIB",
            "path": "seagate/SEAGATESYSTEMTRAP-MIB.mib"
        },
        {
            "mib": "PRVT-CONFIGCHANGE-MIB",
            "path": "telco-systems/binox/PRVT-CONFIGCHANGE-MIB.mib"
        },
        {
            "mib": "DSR4410MD-MIB",
            "path": "arris/DSR4410MD-MIB.mib"
        },
        {
            "mib": "Juniper-IP-Tunnel-CONF",
            "path": "junose/Juniper-IP-Tunnel-CONF.mib"
        },
        {
            "mib": "HUAWEI-PIM-STD-MIB",
            "path": "huawei/HUAWEI-PIM-STD-MIB.mib"
        },
        {
            "mib": "NBS-FEC-MIB",
            "path": "mrv/NBS-FEC-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-NATSTATICRULE-MIB",
            "path": "junos/JUNIPER-LSYSSP-NATSTATICRULE-MIB.mib"
        },
        {
            "mib": "LINKSYS-BONJOUR-MIB",
            "path": "linksys/LINKSYS-BONJOUR-MIB.mib"
        },
        {
            "mib": "PEPVPN-SPEEDFUSION",
            "path": "peplink/PEPVPN-SPEEDFUSION.mib"
        },
        {
            "mib": "MPLS-TC-PRIV-STDEXT-MIB",
            "path": "telco-systems/binos/MPLS-TC-PRIV-STDEXT-MIB.mib"
        },
        {
            "mib": "RITTAL-CMC-TC-MIB",
            "path": "rittal/RITTAL-CMC-TC-MIB.mib"
        },
        {
            "mib": "PERLE-MCR-MGT-MIB",
            "path": "perle/PERLE-MCR-MGT-MIB.mib"
        },
        {
            "mib": "HUAWEI-POE-MIB",
            "path": "huawei/HUAWEI-POE-MIB.mib"
        },
        {
            "mib": "NBS-FECPM-MIB",
            "path": "mrv/NBS-FECPM-MIB.mib"
        },
        {
            "mib": "PRVT-CR-LDP-MIB",
            "path": "telco-systems/binox/PRVT-CR-LDP-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-10GEPON-PM-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-10GEPON-PM-GROUP-MIB.mib"
        },
        {
            "mib": "HH3C-ISIS-MIB",
            "path": "comware/HH3C-ISIS-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-MCLAG-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MCLAG-MIB.mib"
        },
        {
            "mib": "RAISECOM-ERPS-MIB",
            "path": "raisecom/RAISECOM-ERPS-MIB.mib"
        },
        {
            "mib": "CISCO-TC-NO-U32",
            "path": "cisco/CISCO-TC-NO-U32.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-POLICY-MIB",
            "path": "junos/JUNIPER-LSYSSP-POLICY-MIB.mib"
        },
        {
            "mib": "Juniper-IP-TUNNEL-MIB",
            "path": "junose/Juniper-IP-TUNNEL-MIB.mib"
        },
        {
            "mib": "WAN",
            "path": "peplink/WAN.mib"
        },
        {
            "mib": "LINKSYS-BRGMACSWITCH-MIB",
            "path": "linksys/LINKSYS-BRGMACSWITCH-MIB.mib"
        },
        {
            "mib": "RITTAL-SMI-MIB",
            "path": "rittal/RITTAL-SMI.mib"
        },
        {
            "mib": "UBIQUOSS-10GEPON-PON-MAC-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-10GEPON-PON-MAC-GROUP-MIB.mib"
        },
        {
            "mib": "PRIV-LSL-MIB",
            "path": "telco-systems/binos/PRIV-LSL-MIB.mib"
        },
        {
            "mib": "ULAFPLUS-MIB",
            "path": "siemens/PLUS-MIB.mib"
        },
        {
            "mib": "LUM-INVENTORY-MIB",
            "path": "infinera/LUM-INVENTORY-MIB.mib"
        },
        {
            "mib": "HH3C-ISSU-MIB",
            "path": "comware/HH3C-ISSU-MIB.mib"
        },
        {
            "mib": "NBS-JUMPER-MIB",
            "path": "mrv/NBS-JUMPER-MIB.mib"
        },
        {
            "mib": "AIRPORT-BASESTATION-3-MIB",
            "path": "airport/AIRPORT-BASESTATION-3-MIB.mib"
        },
        {
            "mib": "LINKSYS-BRIDGE-SECURITY",
            "path": "linksys/LINKSYS-BRIDGE-SECURITY.mib"
        },
        {
            "mib": "PRVT-EFM-OAM-MIB",
            "path": "telco-systems/binox/PRVT-EFM-OAM-MIB.mib"
        },
        {
            "mib": "RAISECOM-ETHERSAM-MIB",
            "path": "raisecom/RAISECOM-ETHERSAM-MIB.mib"
        },
        {
            "mib": "INFINERA-PM-BANDCTP-MIB",
            "path": "iqnos/INFINERA-PM-BANDCTP.mib"
        },
        {
            "mib": "IBM-TS3500-MIBv2",
            "path": "ibm/IBM-TS3500-MIBv2.mib"
        },
        {
            "mib": "ARUBAWIRED-MDNS-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MDNS-MIB.mib"
        },
        {
            "mib": "QUANTASTOR-SYS-STATS",
            "path": "osnexus/QUANTASTOR-SYS-STATS.mib"
        },
        {
            "mib": "HH3C-L2ISOLATE-MIB",
            "path": "comware/HH3C-L2ISOLATE-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-10GEPON-PON-PROFILE-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-10GEPON-PON-PROFILE-GROUP-MIB.mib"
        },
        {
            "mib": "NBS-META-MIB",
            "path": "mrv/NBS-META-MIB.mib"
        },
        {
            "mib": "Juniper-IPsec-Tunnel-CONF",
            "path": "junose/Juniper-IPsec-Tunnel-CONF.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-POLICYWCNT-MIB",
            "path": "junos/JUNIPER-LSYSSP-POLICYWCNT-MIB.mib"
        },
        {
            "mib": "SN-MSPS-SCX200-MIB",
            "path": "siemens/SN-MSPS-SCX200-MIB.mib"
        },
        {
            "mib": "CISCO-TC",
            "path": "cisco/CISCO-TC.mib"
        },
        {
            "mib": "HUAWEI-PORT-MIB",
            "path": "huawei/HUAWEI-PORT-MIB.mib"
        },
        {
            "mib": "WLC",
            "path": "peplink/WLC.mib"
        },
        {
            "mib": "PRIV-VENDORDEF-MIB",
            "path": "telco-systems/binos/PRIV-VENDORDEF-MIB.mib"
        },
        {
            "mib": "TACHYON-MIB",
            "path": "tachyon/TACHYON-MIB.mib"
        },
        {
            "mib": "NBS-MIB",
            "path": "mrv/NBS-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-10GEPON-SERVICE-POLICY-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-10GEPON-SERVICE-POLICY-GROUP-MIB.mib"
        },
        {
            "mib": "BROTHER-MIB",
            "path": "brother/BROTHER-MIB.mib"
        },
        {
            "mib": "LINKSYS-BRIDGEMIBOBJECTS-MIB",
            "path": "linksys/LINKSYS-BRIDGEMIBOBJECTS-MIB.mib"
        },
        {
            "mib": "RAISECOM-EXTEND-OAM-UPGRADE-MIB",
            "path": "raisecom/RAISECOM-EXTEND-OAM-UPGRADE-MIB.mib"
        },
        {
            "mib": "TRAP-MIB",
            "path": "siemens/TRAP-MIB.mib"
        },
        {
            "mib": "HH3C-L2TP-MIB",
            "path": "comware/HH3C-L2TP-MIB.mib"
        },
        {
            "mib": "PRVT-BIST-MIB",
            "path": "telco-systems/binos/PRVT-BIST-MIB.mib"
        },
        {
            "mib": "PRVT-EPS-MIB",
            "path": "telco-systems/binox/PRVT-EPS-MIB.mib"
        },
        {
            "mib": "HUAWEI-PORTAL-MIB",
            "path": "huawei/HUAWEI-PORTAL-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-10GEPON-SOFTWARE-MANAGEMENT-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-10GEPON-SOFTWARE-MANAGEMENT-GROUP-MIB.mib"
        },
        {
            "mib": "CISCO-TCP-MIB",
            "path": "cisco/CISCO-TCP-MIB.mib"
        },
        {
            "mib": "RMCU",
            "path": "westmountainradio/RMCU.mib"
        },
        {
            "mib": "Juniper-IPsec-Tunnel-MIB",
            "path": "junose/Juniper-IPsec-Tunnel-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-SCHEDULER-MIB",
            "path": "junos/JUNIPER-LSYSSP-SCHEDULER-MIB.mib"
        },
        {
            "mib": "HH3C-L2VPN-MIB",
            "path": "comware/HH3C-L2VPN-MIB.mib"
        },
        {
            "mib": "NBS-OBA-MIB",
            "path": "mrv/NBS-OBA-MIB.mib"
        },
        {
            "mib": "EDS-MIB",
            "path": "eds/EDS-MIB.mib"
        },
        {
            "mib": "CISCO-TCPOFFLOAD-MIB",
            "path": "cisco/CISCO-TCPOFFLOAD-MIB.mib"
        },
        {
            "mib": "LUM-OA-MIB",
            "path": "infinera/LUM-OA-MIB.mib"
        },
        {
            "mib": "RAISECOM-EXTLOOPBACK-MIB",
            "path": "raisecom/RAISECOM-EXTLOOPBACK-MIB.mib"
        },
        {
            "mib": "ESPHOME-ESP32-SNMP-MIB",
            "path": "esphome/ESPHOME-ESP32-SNMP-MIB.mib"
        },
        {
            "mib": "IBM-TS4500-MIBv2",
            "path": "ibm/IBM-TS4500-MIBv2.mib"
        },
        {
            "mib": "LINKSYS-CDB-MIB",
            "path": "linksys/LINKSYS-CDB-MIB.mib"
        },
        {
            "mib": "ULAF2-MIB",
            "path": "siemens/ULAF2-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-MIB.mib"
        },
        {
            "mib": "PRVT-CES-MIB",
            "path": "telco-systems/binos/PRVT-CES-MIB.mib"
        },
        {
            "mib": "NBS-ODSYS-MIB",
            "path": "mrv/NBS-ODSYS-MIB.mib"
        },
        {
            "mib": "HH3C-L2VPN-PWE3-MIB",
            "path": "comware/HH3C-L2VPN-PWE3-MIB.mib"
        },
        {
            "mib": "IFOTEC-PRODUCTLIST-MIB",
            "path": "ifotec/IFOTEC-PRODUCTLIST-MIB.mib"
        },
        {
            "mib": "ESPHOME-ESP8266-SNMP-MIB",
            "path": "esphome/ESPHOME-ESP8266-SNMP-MIB.mib"
        },
        {
            "mib": "INFINERA-PM-OCHCTP-MIB",
            "path": "iqnos/INFINERA-PM-OCHCTP.mib"
        },
        {
            "mib": "HUAWEI-POWER-MIB",
            "path": "huawei/HUAWEI-POWER-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-ONTMANAGER-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-ONTMANAGER-GROUP-MIB.mib"
        },
        {
            "mib": "BROCADE-ACL-MIB",
            "path": "extreme/BROCADE-ACL-MIB.mib"
        },
        {
            "mib": "PRVT-EVENT-PROPAGATION-MIB",
            "path": "telco-systems/binox/PRVT-EVENT-PROPAGATION-MIB.mib"
        },
        {
            "mib": "LINKSYS-CLI-MIB",
            "path": "linksys/LINKSYS-CLI-MIB.mib"
        },
        {
            "mib": "HUAWEI-PPP-MIB",
            "path": "huawei/HUAWEI-PPP-MIB.mib"
        },
        {
            "mib": "BROCADE-CONTEXT-MAPPING-MIB",
            "path": "extreme/BROCADE-CONTEXT-MAPPING-MIB.mib"
        },
        {
            "mib": "HH3C-L4RDT-MIB",
            "path": "comware/HH3C-L4RDT-MIB.mib"
        },
        {
            "mib": "IES5206-MIB",
            "path": "zyxel/IES5206-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-MGMD-RMON-TRAP-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MGMD-RMON-TRAP-MIB.mib"
        },
        {
            "mib": "IES5206-TRAPS-MIB",
            "path": "zyxel/IES5206-TRAPS-MIB.mib"
        },
        {
            "mib": "Juniper-IPV6-PROFILE-MIB",
            "path": "junose/Juniper-IPV6-PROFILE-MIB.mib"
        },
        {
            "mib": "RAISECOM-EXTOAM-MIB",
            "path": "raisecom/RAISECOM-EXTOAM-MIB.mib"
        },
        {
            "mib": "PRVT-CFM-MIB",
            "path": "telco-systems/binos/PRVT-CFM-MIB.mib"
        },
        {
            "mib": "CISCO-UDLDP-MIB",
            "path": "cisco/CISCO-UDLDP-MIB.mib"
        },
        {
            "mib": "DPS-MIB-CG-V1",
            "path": "dpstelecom/DPS-MIB-CG-V1.mib"
        },
        {
            "mib": "LUM-OCM-MIB",
            "path": "infinera/LUM-OCM-MIB.mib"
        },
        {
            "mib": "IBM2210-MIB",
            "path": "ibm/IBM2210-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-PM-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-PM-GROUP-MIB.mib"
        },
        {
            "mib": "HUAWEI-PTP-MIB",
            "path": "huawei/HUAWEI-PTP-MIB.mib"
        },
        {
            "mib": "PRVT-HQOS-MIB",
            "path": "telco-systems/binox/PRVT-HQOS-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-SECUREWIRE-MIB",
            "path": "junos/JUNIPER-LSYSSP-SECUREWIRE-MIB.mib"
        },
        {
            "mib": "NBS-OPTIC-MIB",
            "path": "mrv/NBS-OPTIC-MIB.mib"
        },
        {
            "mib": "HH3C-LAG-MIB",
            "path": "comware/HH3C-LAG-MIB.mib"
        },
        {
            "mib": "LINKSYS-COPY-MIB",
            "path": "linksys/LINKSYS-COPY-MIB.mib"
        },
        {
            "mib": "BROCADE-INTERFACE-STATS-MIB",
            "path": "extreme/BROCADE-INTERFACE-STATS-MIB.mib"
        },
        {
            "mib": "RAISECOM-FANMONITOR-MIB",
            "path": "raisecom/RAISECOM-FANMONITOR-MIB.mib"
        },
        {
            "mib": "DPS-MIB-V38-V2",
            "path": "dpstelecom/DPS-MIB-V38-V2.mib"
        },
        {
            "mib": "ZYXEL-AESCOMMON-MIB",
            "path": "zyxel/ZYXEL-AESCOMMON-MIB.mib"
        },
        {
            "mib": "INFINERA-PM-OSCCTP-MIB",
            "path": "iqnos/INFINERA-PM-OSCCTP.mib"
        },
        {
            "mib": "Juniper-IS-IS-CONF",
            "path": "junose/Juniper-IS-IS-CONF.mib"
        },
        {
            "mib": "IFOTEC-SMI",
            "path": "ifotec/IFOTEC-SMI.mib"
        },
        {
            "mib": "ARUBAWIRED-MGMD-SNOOPING-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MGMD-SNOOPING-MIB.mib"
        },
        {
            "mib": "PRVT-IGMP-SNOOPING-MIB",
            "path": "telco-systems/binox/PRVT-IGMP-SNOOPING-MIB.mib"
        },
        {
            "mib": "PRVT-CONFIGCHANGE-MIB",
            "path": "telco-systems/binos/PRVT-CONFIGCHANGE-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-COMPUTE-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-COMPUTE-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-PM-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-PM-MIB.mib"
        },
        {
            "mib": "LUM-REG",
            "path": "infinera/LUM-REG.mib"
        },
        {
            "mib": "IBM2212-MIB",
            "path": "ibm/IBM2212-MIB.mib"
        },
        {
            "mib": "HUAWEI-PWE3-MIB",
            "path": "huawei/HUAWEI-PWE3-MIB.mib"
        },
        {
            "mib": "NBS-OSA-MIB",
            "path": "mrv/NBS-OSA-MIB.mib"
        },
        {
            "mib": "HH3C-LB-MIB",
            "path": "comware/HH3C-LB-MIB.mib"
        },
        {
            "mib": "BROCADE-MODULE-CPU-UTIL-MIB",
            "path": "extreme/BROCADE-MODULE-CPU-UTIL-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSP-ZONE-MIB",
            "path": "junos/JUNIPER-LSYSSP-ZONE-MIB.mib"
        },
        {
            "mib": "DPS-MIB-V38-V2EXT",
            "path": "dpstelecom/DPS-MIB-V38-V2EXT.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-PON-MAC-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-PON-MAC-GROUP-MIB.mib"
        },
        {
            "mib": "RAISECOM-GARP-MIB",
            "path": "raisecom/RAISECOM-GARP-MIB.mib"
        },
        {
            "mib": "INFINERA-PM-SCHCTP-MIB",
            "path": "iqnos/INFINERA-PM-SCHCTP.mib"
        },
        {
            "mib": "PRVT-INTERWORKING-OS-MIB",
            "path": "telco-systems/binox/PRVT-INTERWORKING-OS-MIB.mib"
        },
        {
            "mib": "ZYXEL-AS-ATM-MIB",
            "path": "zyxel/ZYXEL-AS-ATM-MIB.mib"
        },
        {
            "mib": "LINKSYS-CPU-COUNTERS-MIB",
            "path": "linksys/LINKSYS-CPU-COUNTERS-MIB.mib"
        },
        {
            "mib": "LIGO-WIRELESS-MIB",
            "path": "ligoos/LIGO-WIRELESS-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-MODULE-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MODULE-MIB.mib"
        },
        {
            "mib": "PRVT-CR-LDP-MIB",
            "path": "telco-systems/binos/PRVT-CR-LDP-MIB.mib"
        },
        {
            "mib": "HH3C-LBV2-MIB",
            "path": "comware/HH3C-LBV2-MIB.mib"
        },
        {
            "mib": "JUNIPER-LSYSSPAUTHENTRY-MIB",
            "path": "junos/JUNIPER-LSYSSPAUTHENTRY-MIB.mib"
        },
        {
            "mib": "NBS-OTNOH-MIB",
            "path": "mrv/NBS-OTNOH-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-PON-PROFILE-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-PON-PROFILE-GROUP-MIB.mib"
        },
        {
            "mib": "IBM2216-MIB",
            "path": "ibm/IBM2216-MIB.mib"
        },
        {
            "mib": "HUAWEI-PWE3-TNL-MIB",
            "path": "huawei/HUAWEI-PWE3-TNL-MIB.mib"
        },
        {
            "mib": "RAISECOM-IGMPL2-MIB",
            "path": "raisecom/RAISECOM-IGMPL2-MIB.mib"
        },
        {
            "mib": "DPS-MIB-V38",
            "path": "dpstelecom/DPS-MIB-V38.mib"
        },
        {
            "mib": "HH3C-LI-MIB",
            "path": "comware/HH3C-LI-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-EQUIPMENT-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-EQUIPMENT-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-SERVICE-POLICY-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-SERVICE-POLICY-GROUP-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-MEMORY-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-MEMORY-MIB.mib"
        },
        {
            "mib": "PRVT-DHCP-MIB",
            "path": "telco-systems/binos/PRVT-DHCP-MIB.mib"
        },
        {
            "mib": "BROCADE-MODULE-MEM-UTIL-MIB",
            "path": "extreme/BROCADE-MODULE-MEM-UTIL-MIB.mib"
        },
        {
            "mib": "Juniper-ISIS-MIB",
            "path": "junose/Juniper-ISIS-MIB.mib"
        },
        {
            "mib": "LUM-SYSINFO-MIB",
            "path": "infinera/LUM-SYSINFO-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-MSTP-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MSTP-MIB.mib"
        },
        {
            "mib": "DPS-TEXT-RTU-MIB",
            "path": "dpstelecom/DPS-TEXT-RTU-MIB.mib"
        },
        {
            "mib": "JUNIPER-MAC-MIB",
            "path": "junos/JUNIPER-MAC-MIB.mib"
        },
        {
            "mib": "PRVT-ISIS-MIB",
            "path": "telco-systems/binox/PRVT-ISIS-MIB.mib"
        },
        {
            "mib": "INFINERA-REG-MIB",
            "path": "iqnos/INFINERA-REG-MIB.mib"
        },
        {
            "mib": "RAISECOM-IP-BASE-MIB",
            "path": "raisecom/RAISECOM-IP-BASE-MIB.mib"
        },
        {
            "mib": "ZYXEL-AS-MIB",
            "path": "zyxel/ZYXEL-AS-MIB.mib"
        },
        {
            "mib": "HUAWEI-QINQ-MIB",
            "path": "huawei/HUAWEI-QINQ-MIB.mib"
        },
        {
            "mib": "LIGOWAVE-MIB",
            "path": "ligoos/LIGOWAVE-MIB.mib"
        },
        {
            "mib": "LINKSYS-DEBUGCAPABILITIES-MIB",
            "path": "linksys/LINKSYS-DEBUGCAPABILITIES-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-SOFTWARE-MANAGEMENT-GROUP-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-SOFTWARE-MANAGEMENT-GROUP-MIB.mib"
        },
        {
            "mib": "BROCADE-MPLS-MIB",
            "path": "extreme/BROCADE-MPLS-MIB.mib"
        },
        {
            "mib": "ALGPOWER-v2-MIB",
            "path": "algcom/ALGPOWER-v2.mib"
        },
        {
            "mib": "HH3C-LICENSE-MIB",
            "path": "comware/HH3C-LICENSE-MIB.mib"
        },
        {
            "mib": "ZYXEL-ES-CAPWAP",
            "path": "zyxel/ZYXEL-ES-CAPWAP.mib"
        },
        {
            "mib": "PRVT-L2TUNNELING-MIB",
            "path": "telco-systems/binox/PRVT-L2TUNNELING-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-MIB.mib"
        },
        {
            "mib": "PRVT-DRY-CONTACTS-MIB",
            "path": "telco-systems/binos/PRVT-DRY-CONTACTS-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-EPON-TC-MIB",
            "path": "ubiquoss/UBIQUOSS-EPON-TC-MIB.mib"
        },
        {
            "mib": "NBS-OTNPM-MIB",
            "path": "mrv/NBS-OTNPM-MIB.mib"
        },
        {
            "mib": "Juniper-L2TP-CONF",
            "path": "junose/Juniper-L2TP-CONF.mib"
        },
        {
            "mib": "HH3C-LLDP-EXT-MIB",
            "path": "comware/HH3C-LLDP-EXT-MIB.mib"
        },
        {
            "mib": "IBM3172-MIB",
            "path": "ibm/IBM3172-MIB.mib"
        },
        {
            "mib": "LINKSYS-DEVICEPARAMS-MIB",
            "path": "linksys/LINKSYS-DEVICEPARAMS-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-PROCESSOR-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-PROCESSOR-MIB.mib"
        },
        {
            "mib": "JUNIPER-MAG-MIB",
            "path": "junos/JUNIPER-MAG-MIB.mib"
        },
        {
            "mib": "ZYXEL-ES-COMMON",
            "path": "zyxel/ZYXEL-ES-COMMON.mib"
        },
        {
            "mib": "RAISECOM-IPMCAST-MIB",
            "path": "raisecom/RAISECOM-IPMCAST-MIB.mib"
        },
        {
            "mib": "HUAWEI-RBRP-MIB",
            "path": "huawei/HUAWEI-RBRP-MIB.mib"
        },
        {
            "mib": "LUM-SYSTEM-MIB",
            "path": "infinera/LUM-SYSTEM-MIB.mib"
        },
        {
            "mib": "NBS-PART-MIB",
            "path": "mrv/NBS-PART-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-INTERFACE-MIB",
            "path": "ubiquoss/UBIQUOSS-INTERFACE-MIB.mib"
        },
        {
            "mib": "JUNIPER-MBG-SMI",
            "path": "junos/JUNIPER-MBG-SMI.mib"
        },
        {
            "mib": "PRVT-EGRESS-COUNTERS-MIB",
            "path": "telco-systems/binos/PRVT-EGRESS-COUNTERS-MIB.mib"
        },
        {
            "mib": "ALGSMXS-v0-MIB",
            "path": "algcom/ALGSMXS-v0-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-MVRP-MIB",
            "path": "arubaos-cx/ARUBAWIRED-MVRP-MIB.mib"
        },
        {
            "mib": "Juniper-L2TP-Dialout-Generator-CONF",
            "path": "junose/Juniper-L2TP-Dialout-Generator-CONF.mib"
        },
        {
            "mib": "HH3C-LOCAL-AAA-SERVER-MIB",
            "path": "comware/HH3C-LOCAL-AAA-SERVER-MIB.mib"
        },
        {
            "mib": "RAISECOM-IPSOURCEGUARD-MIB",
            "path": "raisecom/RAISECOM-IPSOURCEGUARD-MIB.mib"
        },
        {
            "mib": "BROCADE-OPTICAL-MONITORING-MIB",
            "path": "extreme/BROCADE-OPTICAL-MONITORING-MIB.mib"
        },
        {
            "mib": "NBS-PRBS-MIB",
            "path": "mrv/NBS-PRBS-MIB.mib"
        },
        {
            "mib": "IB-DHCPONE-MIB",
            "path": "infoblox/IB-DHCPONE-MIB.mib"
        },
        {
            "mib": "LINKSYS-DHCP-MIB",
            "path": "linksys/LINKSYS-DHCP-MIB.mib"
        },
        {
            "mib": "ZYXEL-ES-ProWLAN",
            "path": "zyxel/ZYXEL-ES-ProWLAN.mib"
        },
        {
            "mib": "UBIQUOSS-MAC-MIB",
            "path": "ubiquoss/UBIQUOSS-MAC-MIB.mib"
        },
        {
            "mib": "HH3C-LOGIC-VOLUME-MIB",
            "path": "comware/HH3C-LOGIC-VOLUME-MIB.mib"
        },
        {
            "mib": "PRVT-ELMI-MIB",
            "path": "telco-systems/binos/PRVT-ELMI-MIB.mib"
        },
        {
            "mib": "PRVT-LLDP-MIB",
            "path": "telco-systems/binox/PRVT-LLDP-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-STORAGE-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-STORAGE-MIB.mib"
        },
        {
            "mib": "HUAWEI-RIPV2-EXT-MIB",
            "path": "huawei/HUAWEI-RIPV2-EXT-MIB.mib"
        },
        {
            "mib": "IBM6611-MIB",
            "path": "ibm/IBM6611-MIB.mib"
        },
        {
            "mib": "RAISECOM-KEEPALIVE-MIB",
            "path": "raisecom/RAISECOM-KEEPALIVE-MIB.mib"
        },
        {
            "mib": "JUNIPER-MIB",
            "path": "junos/JUNIPER-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-NETWORKING-OID",
            "path": "arubaos-cx/ARUBAWIRED-NETWORKING-OID.mib"
        },
        {
            "mib": "LINKSYS-DHCPCL-MIB",
            "path": "linksys/LINKSYS-DHCPCL-MIB.mib"
        },
        {
            "mib": "LUM-WDM-MIB",
            "path": "infinera/LUM-WDM-MIB.mib"
        },
        {
            "mib": "EXALINK-FUSION-MIB",
            "path": "exalink-fusion/EXALINK-FUSION-MIB.mib"
        },
        {
            "mib": "ZYXEL-ES-RF-MANAGEMENT",
            "path": "zyxel/ZYXEL-ES-RF-MANAGEMENT.mib"
        },
        {
            "mib": "BROCADE-PRODUCTS-MIB",
            "path": "extreme/BROCADE-PRODUCTS-MIB.mib"
        },
        {
            "mib": "NBS-REDUNDANCY-MIB",
            "path": "mrv/NBS-REDUNDANCY-MIB.mib"
        },
        {
            "mib": "PRVT-EPS-MIB",
            "path": "telco-systems/binos/PRVT-EPS-MIB.mib"
        },
        {
            "mib": "INFINERA-TC-MIB",
            "path": "iqnos/INFINERA-TC-MIB.mib"
        },
        {
            "mib": "PRVT-LMGR-MIB",
            "path": "telco-systems/binox/PRVT-LMGR-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-SMI",
            "path": "ubiquoss/UBIQUOSS-SMI.mib"
        },
        {
            "mib": "Brocade-REG-MIB",
            "path": "extreme/BROCADE-REG-MIB.mib"
        },
        {
            "mib": "HH3C-LPBKDT-MIB",
            "path": "comware/HH3C-LPBKDT-MIB.mib"
        },
        {
            "mib": "A4400-CPU-MIB",
            "path": "alcatel/A4400-CPU-MIB.mib"
        },
        {
            "mib": "RAISECOM-KEYCHAIN-MIB",
            "path": "raisecom/RAISECOM-KEYCHAIN-MIB.mib"
        },
        {
            "mib": "IB-DNSONE-MIB",
            "path": "infoblox/IB-DNSONE-MIB.mib"
        },
        {
            "mib": "LINKSYS-DHCPv6-CLIENT",
            "path": "linksys/LINKSYS-DHCPv6-CLIENT.mib"
        },
        {
            "mib": "Juniper-MIBs",
            "path": "junos/Juniper-MIBs.mib"
        },
        {
            "mib": "ARUBAWIRED-PM-MIB",
            "path": "arubaos-cx/ARUBAWIRED-PM-MIB.mib"
        },
        {
            "mib": "ZYXEL-ES-SMI",
            "path": "zyxel/ZYXEL-ES-SMI.mib"
        },
        {
            "mib": "IBMACCOUNTING-MIB",
            "path": "ibm/IBMACCOUNTING-MIB.mib"
        },
        {
            "mib": "Brocade-TC",
            "path": "extreme/BROCADE-TC.mib"
        },
        {
            "mib": "UBIQUOSS-STP-MIB",
            "path": "ubiquoss/UBIQUOSS-STP-MIB.mib"
        },
        {
            "mib": "Juniper-L2TP-Dialout-MIB",
            "path": "junose/Juniper-L2TP-Dialout-MIB.mib"
        },
        {
            "mib": "A4400-RTM-MIB",
            "path": "alcatel/A4400-RTM-MIB.mib"
        },
        {
            "mib": "HUAWEI-RM-EXT-MIB",
            "path": "huawei/HUAWEI-RM-EXT-MIB.mib"
        },
        {
            "mib": "HH3C-LSW-DEV-ADM-MIB",
            "path": "comware/HH3C-LSW-DEV-ADM-MIB.mib"
        },
        {
            "mib": "PRVT-LMM-MIB",
            "path": "telco-systems/binox/PRVT-LMM-MIB.mib"
        },
        {
            "mib": "CISCO-VISM-DSX0-MIB",
            "path": "cisco/CISCO-VISM-DSX0-MIB.mib"
        },
        {
            "mib": "PRVT-INTERWORKING-OS-MIB",
            "path": "telco-systems/binos/PRVT-INTERWORKING-OS-MIB.mib"
        },
        {
            "mib": "RADWARE-MIB",
            "path": "radware/RADWARE-MIB.mib"
        },
        {
            "mib": "NBS-ROADM-MIB",
            "path": "mrv/NBS-ROADM-MIB.mib"
        },
        {
            "mib": "LINKSYS-DHCPv6-RELAY",
            "path": "linksys/LINKSYS-DHCPv6-RELAY.mib"
        },
        {
            "mib": "ZYXEL-ES-WIRELESS",
            "path": "zyxel/ZYXEL-ES-WIRELESS.mib"
        },
        {
            "mib": "UBIQUOSS-SWITCH-INTERFACE-MIB",
            "path": "ubiquoss/UBIQUOSS-SWITCH-INTERFACE-MIB.mib"
        },
        {
            "mib": "BROCADE-TCAM-MIB",
            "path": "extreme/BROCADE-TCAM-MIB.mib"
        },
        {
            "mib": "CA-SNMP-MIB8",
            "path": "arraynetworks/CA-SNMP-MIB8.mib"
        },
        {
            "mib": "HH3C-LswDEVM-MIB",
            "path": "comware/HH3C-LswDEVM-MIB.mib"
        },
        {
            "mib": "IB-PLATFORMONE-MIB",
            "path": "infoblox/IB-PLATFORMONE-MIB.mib"
        },
        {
            "mib": "RAISECOM-L2CP-MIB",
            "path": "raisecom/RAISECOM-L2CP-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-POE-MIB",
            "path": "arubaos-cx/ARUBAWIRED-POE-MIB.mib"
        },
        {
            "mib": "HUAWEI-RPR-MIB",
            "path": "huawei/HUAWEI-RPR-MIB.mib"
        },
        {
            "mib": "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB",
            "path": "cisco/CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB.mib"
        },
        {
            "mib": "JUNIPER-MIMSTP-MIB",
            "path": "junos/JUNIPER-MIMSTP-MIB.mib"
        },
        {
            "mib": "NBS-SFF-MIB",
            "path": "mrv/NBS-SFF-MIB.mib"
        },
        {
            "mib": "ALCATEL-NMC-PROXY-AGENT-MIB",
            "path": "alcatel/ALCATEL-NMC-PROXY-AGENT-MIB.mib"
        },
        {
            "mib": "UBIQUOSS-SYSINFO-MIB",
            "path": "ubiquoss/UBIQUOSS-SYSINFO-MIB.mib"
        },
        {
            "mib": "PRVT-JDSU-MIB",
            "path": "telco-systems/binos/PRVT-JDSU-MIB.mib"
        },
        {
            "mib": "LINKSYS-DHCPv6",
            "path": "linksys/LINKSYS-DHCPv6.mib"
        },
        {
            "mib": "HH3C-LTE-MEC-MIB",
            "path": "comware/HH3C-LTE-MEC-MIB.mib"
        },
        {
            "mib": "BROCADE-TMSTATS-MIB",
            "path": "extreme/BROCADE-TMSTATS-MIB.mib"
        },
        {
            "mib": "PRVT-LOAD-BALANCE-MIB",
            "path": "telco-systems/binox/PRVT-LOAD-BALANCE-MIB.mib"
        },
        {
            "mib": "RAISECOM-LBDETECT-MIB",
            "path": "raisecom/RAISECOM-LBDETECT-MIB.mib"
        },
        {
            "mib": "NBS-SIGCOND-MIB",
            "path": "mrv/NBS-SIGCOND-MIB.mib"
        },
        {
            "mib": "CISCO-VLAN-MEMBERSHIP-MIB",
            "path": "cisco/CISCO-VLAN-MEMBERSHIP-MIB.mib"
        },
        {
            "mib": "UBQS-ACCESS-LIST-MIB",
            "path": "ubiquoss/UBQS-ACCESS-LIST-MIB.mib"
        },
        {
            "mib": "HYTERA-REPEATER-MIB",
            "path": "hytera/HYTERA-REPEATER-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-PORT-ACCESS-MIB",
            "path": "arubaos-cx/ARUBAWIRED-PORT-ACCESS-MIB.mib"
        },
        {
            "mib": "DCN-MIB",
            "path": "dcn/DCN-MIB.mib"
        },
        {
            "mib": "LINKSYS-Dlf-MIB",
            "path": "linksys/LINKSYS-Dlf-MIB.mib"
        },
        {
            "mib": "HPOV-NNM-MIB",
            "path": "alcatel/HPOV-NNM-MIB.mib"
        },
        {
            "mib": "HH3C-MAC-INFORMATION-MIB",
            "path": "comware/HH3C-MAC-INFORMATION-MIB.mib"
        },
        {
            "mib": "RAISECOM-LINKAGGREGATION-MIB",
            "path": "raisecom/RAISECOM-LINKAGGREGATION-MIB.mib"
        },
        {
            "mib": "NBS-SIGLANE-MIB",
            "path": "mrv/NBS-SIGLANE-MIB.mib"
        },
        {
            "mib": "BROCADE-UDLD-MIB",
            "path": "extreme/BROCADE-UDLD-MIB.mib"
        },
        {
            "mib": "ZYXEL-ES-ZyxelAPMgmt",
            "path": "zyxel/ZYXEL-ES-ZyxelAPMgmt.mib"
        },
        {
            "mib": "UBQS-AFSMGR-MIB",
            "path": "ubiquoss/UBQS-AFSMGR-MIB.mib"
        },
        {
            "mib": "HUAWEI-RRPP-MIB",
            "path": "huawei/HUAWEI-RRPP-MIB.mib"
        },
        {
            "mib": "PRVT-L2TUNNELING-MIB",
            "path": "telco-systems/binos/PRVT-L2TUNNELING-MIB.mib"
        },
        {
            "mib": "IB-SMI-MIB",
            "path": "infoblox/IB-SMI-MIB.mib"
        },
        {
            "mib": "CAMBIUM-MIB",
            "path": "cambium/cnpilotr/CAMBIUM-MIB.mib"
        },
        {
            "mib": "IBMAPPNMEMORY-MIB",
            "path": "ibm/IBMAPPNMEMORY-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-3G",
            "path": "adtran/ADTRAN-AOS-3G.mib"
        },
        {
            "mib": "Juniper-L2TP-MIB",
            "path": "junose/Juniper-L2TP-MIB.mib"
        },
        {
            "mib": "HH3C-MACSEC-MIB",
            "path": "comware/HH3C-MACSEC-MIB.mib"
        },
        {
            "mib": "PRVT-MAC-SECURITY-MIB",
            "path": "telco-systems/binox/PRVT-MAC-SECURITY-MIB.mib"
        },
        {
            "mib": "UBQS-ARP-MIB",
            "path": "ubiquoss/UBQS-ARP-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-PORTSECURITY-MIB",
            "path": "arubaos-cx/ARUBAWIRED-PORTSECURITY-MIB.mib"
        },
        {
            "mib": "RAISECOM-LLDP-STD-MIB",
            "path": "raisecom/RAISECOM-LLDP-STD-MIB.mib"
        },
        {
            "mib": "IRT-COMMONVARBINDS-MIB",
            "path": "rs/IRT-COMMONVARBINDS-MIB.mib"
        },
        {
            "mib": "CISCO-VOA-MIB",
            "path": "cisco/CISCO-VOA-MIB.mib"
        },
        {
            "mib": "NBS-SLA-MIB",
            "path": "mrv/NBS-SLA-MIB.mib"
        },
        {
            "mib": "EXTREME-BASE-MIB",
            "path": "extreme/EXTREME-BASE-MIB.mib"
        },
        {
            "mib": "NE843-MIB",
            "path": "gepower/NE843-MIB.mib"
        },
        {
            "mib": "PRVT-LLDP-MIB",
            "path": "telco-systems/binos/PRVT-LLDP-MIB.mib"
        },
        {
            "mib": "ZYXEL-GS2200-24-MIB",
            "path": "zyxel/ZYXEL-GS2200-24-MIB.mib"
        },
        {
            "mib": "LINKSYS-DNSCL-MIB",
            "path": "linksys/LINKSYS-DNSCL-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-DESKTOP-AUDITING",
            "path": "adtran/ADTRAN-AOS-DESKTOP-AUDITING.mib"
        },
        {
            "mib": "HH3C-MCDR-MIB",
            "path": "comware/HH3C-MCDR-MIB.mib"
        },
        {
            "mib": "UBQS-AUTO-RESET-MIB",
            "path": "ubiquoss/UBQS-AUTO-RESET-MIB.mib"
        },
        {
            "mib": "XMUX4-PLUS",
            "path": "fibernet/XMUX4-PLUS.mib"
        },
        {
            "mib": "RAISECOM-LOOPBACK-MIB",
            "path": "raisecom/RAISECOM-LOOPBACK-MIB.mib"
        },
        {
            "mib": "NORTEL-MIB",
            "path": "nortel/NORTEL-MIB.mib"
        },
        {
            "mib": "IRT-DVBT-SINGLETRANSMITTER-MIB",
            "path": "rs/IRT-DVBT-SINGLETRANSMITTER-MIB.mib"
        },
        {
            "mib": "NBS-STATS-MIB",
            "path": "mrv/NBS-STATS-MIB.mib"
        },
        {
            "mib": "HUAWEI-RSVPTE-MIB",
            "path": "huawei/HUAWEI-RSVPTE-MIB.mib"
        },
        {
            "mib": "Juniper-License-Mgr-CONF",
            "path": "junose/Juniper-License-Mgr-CONF.mib"
        },
        {
            "mib": "ARUBAWIRED-PORTVLAN-MIB",
            "path": "arubaos-cx/ARUBAWIRED-PORTVLAN-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-AAA-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-AAA-MIB.mib"
        },
        {
            "mib": "CISCO-VOICE-COMMON-DIAL-CONTROL-MIB",
            "path": "cisco/CISCO-VOICE-COMMON-DIAL-CONTROL-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-DNS-MIB",
            "path": "adtran/ADTRAN-AOS-DNS-MIB.mib"
        },
        {
            "mib": "HH3C-MDC-MIB",
            "path": "comware/HH3C-MDC-MIB.mib"
        },
        {
            "mib": "UBQS-CFM-MIB",
            "path": "ubiquoss/UBQS-CFM-MIB.mib"
        },
        {
            "mib": "EXTREME-CABLE-MIB",
            "path": "extreme/EXTREME-CABLE-MIB.mib"
        },
        {
            "mib": "RAISECOM-MCAST-MIB",
            "path": "raisecom/RAISECOM-MCAST-MIB.mib"
        },
        {
            "mib": "PRVT-MPLS-IF-MIB",
            "path": "telco-systems/binox/PRVT-MPLS-IF-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-DS1-EXT",
            "path": "adtran/ADTRAN-AOS-DS1-EXT.mib"
        },
        {
            "mib": "IRT-TRANSMITTER-SMI-MIB",
            "path": "rs/IRT-TRANSMITTER-SMI-MIB.mib"
        },
        {
            "mib": "PRVT-LMGR-MIB",
            "path": "telco-systems/binos/PRVT-LMGR-MIB.mib"
        },
        {
            "mib": "LINKSYS-DOT1X-MIB",
            "path": "linksys/LINKSYS-DOT1X-MIB.mib"
        },
        {
            "mib": "HH3C-MINM-MIB",
            "path": "comware/HH3C-MINM-MIB.mib"
        },
        {
            "mib": "NORTEL-OPTICAL-GENERIC-MIB",
            "path": "nortel/NORTEL-OPTICAL-GENERIC-MIB.mib"
        },
        {
            "mib": "IBMBNA-MIB",
            "path": "ibm/IBMBNA-MIB.mib"
        },
        {
            "mib": "NBS-SYSCOMM-MIB",
            "path": "mrv/NBS-SYSCOMM-MIB.mib"
        },
        {
            "mib": "Juniper-LICENSE-MIB",
            "path": "junose/Juniper-LICENSE-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-POWER-STAT-MIB",
            "path": "arubaos-cx/ARUBAWIRED-POWER-STAT-MIB.mib"
        },
        {
            "mib": "RAISECOM-MGMD-MIB",
            "path": "raisecom/RAISECOM-MGMD-MIB.mib"
        },
        {
            "mib": "RS-COMMON-MIB",
            "path": "rs/RS-COMMON-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP670-MIB",
            "path": "cambium/CAMBIUM-PTP670-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-DHCP-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-DHCP-MIB.mib"
        },
        {
            "mib": "ZYXEL-GS4012F-MIB",
            "path": "zyxel/ZYXEL-GS4012F-MIB.mib"
        },
        {
            "mib": "Montclair-MIB",
            "path": "montclair/Montclair-MIB.mib"
        },
        {
            "mib": "PRVT-MONITOR-SESSION-MIB",
            "path": "telco-systems/binos/PRVT-MONITOR-SESSION-MIB.mib"
        },
        {
            "mib": "PRVT-MPLS-LDP-MIB",
            "path": "telco-systems/binox/PRVT-MPLS-LDP-MIB.mib"
        },
        {
            "mib": "UBQS-CONFIG-MIB",
            "path": "ubiquoss/UBQS-CONFIG-MIB.mib"
        },
        {
            "mib": "CISCO-VOICE-DIAL-CONTROL-MIB",
            "path": "cisco/CISCO-VOICE-DIAL-CONTROL-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-DYING-GASP-MIB",
            "path": "adtran/ADTRAN-AOS-DYING-GASP-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-POWERSUPPLY-MIB",
            "path": "arubaos-cx/ARUBAWIRED-POWERSUPPLY-MIB.mib"
        },
        {
            "mib": "NORTEL-OPTICAL-PM-MIB",
            "path": "nortel/NORTEL-OPTICAL-PM-MIB.mib"
        },
        {
            "mib": "NBS-SYSLOG-SERVER-MIB",
            "path": "mrv/NBS-SYSLOG-SERVER-MIB.mib"
        },
        {
            "mib": "CISCO-VOICE-DNIS-MIB",
            "path": "cisco/CISCO-VOICE-DNIS-MIB.mib"
        },
        {
            "mib": "HUAWEI-SECSTAT-EUDM-MIB",
            "path": "huawei/HUAWEI-SECSTAT-EUDM-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-FAN-MIB",
            "path": "adtran/ADTRAN-AOS-FAN-MIB.mib"
        },
        {
            "mib": "CISCO-VPDN-MGMT-EXT-MIB",
            "path": "cisco/CISCO-VPDN-MGMT-EXT-MIB.mib"
        },
        {
            "mib": "RS-XX8000-COMMON-MIB",
            "path": "rs/RS-XX8000-COMMON-MIB.mib"
        },
        {
            "mib": "EXTREME-DLCS-MIB",
            "path": "extreme/EXTREME-DLCS-MIB.mib"
        },
        {
            "mib": "Juniper-Local-Address-Server-CONF",
            "path": "junose/Juniper-Local-Address-Server-CONF.mib"
        },
        {
            "mib": "RAISECOM-MLACP-MIB",
            "path": "raisecom/RAISECOM-MLACP-MIB.mib"
        },
        {
            "mib": "UBQS-CPU-MAC-FILTER-MIB",
            "path": "ubiquoss/UBQS-CPU-MAC-FILTER.mib"
        },
        {
            "mib": "HH3C-MIRRORGROUP-MIB",
            "path": "comware/HH3C-MIRRORGROUP-MIB.mib"
        },
        {
            "mib": "NBS-TRAPCONTROL-MIB",
            "path": "mrv/NBS-TRAPCONTROL-MIB.mib"
        },
        {
            "mib": "RAISECOM-MODULE-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-MODULE-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "CISCO-VPDN-MGMT-MIB",
            "path": "cisco/CISCO-VPDN-MGMT-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-PROVIDER-BRIDGE-MIB",
            "path": "arubaos-cx/ARUBAWIRED-PROVIDER-BRIDGE-MIB.mib"
        },
        {
            "mib": "PRVT-MPLS-IF-MIB",
            "path": "telco-systems/binos/PRVT-MPLS-IF-MIB.mib"
        },
        {
            "mib": "RAPID-CITY-BAY-STACK",
            "path": "nortel/RAPID-CITY-BAY-STACK.mib"
        },
        {
            "mib": "EXTREME-DOS-MIB",
            "path": "extreme/EXTREME-DOS-MIB.mib"
        },
        {
            "mib": "ACS-MIB",
            "path": "avocent/ACS-MIB.mib"
        },
        {
            "mib": "PRVT-MPLS-TE-MIB",
            "path": "telco-systems/binox/PRVT-MPLS-TE-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-MEDIAGATEWAY-MIB",
            "path": "adtran/ADTRAN-AOS-MEDIAGATEWAY-MIB.mib"
        },
        {
            "mib": "RS-XX8000-DVB-TX-MIB",
            "path": "rs/RS-XX8000-DVB-TX-MIB.mib"
        },
        {
            "mib": "LINKSYS-EEE-MIB",
            "path": "linksys/LINKSYS-EEE-MIB.mib"
        },
        {
            "mib": "HH3C-MP-MIB",
            "path": "comware/HH3C-MP-MIB.mib"
        },
        {
            "mib": "ZYXEL-HW-MONITOR-MIB",
            "path": "zyxel/ZYXEL-HW-MONITOR-MIB.mib"
        },
        {
            "mib": "UBQS-CPU-STATS-MIB",
            "path": "ubiquoss/UBQS-CPU-STATS-MIB.mib"
        },
        {
            "mib": "CISCO-VRF-MIB",
            "path": "cisco/CISCO-VRF-MIB.mib"
        },
        {
            "mib": "NBS-TUNABLE-MIB",
            "path": "mrv/NBS-TUNABLE-MIB.mib"
        },
        {
            "mib": "HH3C-MP-V2-MIB",
            "path": "comware/HH3C-MP-V2-MIB.mib"
        },
        {
            "mib": "RAISECOM-MPLS-LSPV-MIB",
            "path": "raisecom/RAISECOM-MPLS-LSPV-MIB.mib"
        },
        {
            "mib": "IBMCPU-MIB",
            "path": "ibm/IBMCPU-MIB.mib"
        },
        {
            "mib": "PRVT-MPLS-LDP-MIB",
            "path": "telco-systems/binos/PRVT-MPLS-LDP-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-MEF-MIB",
            "path": "adtran/ADTRAN-AOS-MEF-MIB.mib"
        },
        {
            "mib": "HUAWEI-SECSTAT-IP-MONITOR-MIB",
            "path": "huawei/HUAWEI-SECSTAT-IP-MONITOR-MIB.mib"
        },
        {
            "mib": "Juniper-Log-CONF",
            "path": "junose/Juniper-Log-CONF.mib"
        },
        {
            "mib": "ARUBAWIRED-RPVST-MIB",
            "path": "arubaos-cx/ARUBAWIRED-RPVST-MIB.mib"
        },
        {
            "mib": "EXTREME-EAPS-MIB",
            "path": "extreme/EXTREME-EAPS-MIB.mib"
        },
        {
            "mib": "PM-MIB",
            "path": "avocent/PM-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-EXAMPLE-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-EXAMPLE-MIB.mib"
        },
        {
            "mib": "PRVT-OSPF-MIB",
            "path": "telco-systems/binox/PRVT-OSPF-MIB.mib"
        },
        {
            "mib": "LSI-MegaRAID-SAS-MIB",
            "path": "lsi/LSI-MegaRAID-SAS-MIB.mib"
        },
        {
            "mib": "UBQS-DHCP-MIB",
            "path": "ubiquoss/UBQS-DHCP-MIB.mib"
        },
        {
            "mib": "RAISECOM-MPLS-MIB",
            "path": "raisecom/RAISECOM-MPLS-MIB.mib"
        },
        {
            "mib": "CISCO-VTP-MIB",
            "path": "cisco/CISCO-VTP-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-MUX-MIB",
            "path": "adtran/ADTRAN-AOS-MUX-MIB.mib"
        },
        {
            "mib": "HUAWEI-SECSTAT-MIB",
            "path": "huawei/HUAWEI-SECSTAT-MIB.mib"
        },
        {
            "mib": "NBS-USER-SESSION-MIB",
            "path": "mrv/NBS-USER-SESSION-MIB.mib"
        },
        {
            "mib": "LINKSYS-EMBWEB-MIB",
            "path": "linksys/LINKSYS-EMBWEB-MIB.mib"
        },
        {
            "mib": "HH3C-MPLS-LDP-MIB",
            "path": "comware/HH3C-MPLS-LDP-MIB.mib"
        },
        {
            "mib": "EXTREME-EDP-MIB",
            "path": "extreme/EXTREME-EDP-MIB.mib"
        },
        {
            "mib": "PRVT-PORTS-AGGREGATION-MIB",
            "path": "telco-systems/binox/PRVT-PORTS-AGGREGATION-MIB.mib"
        },
        {
            "mib": "Juniper-LOG-MIB",
            "path": "junose/Juniper-LOG-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-SWITCH-IMAGE-MIB",
            "path": "arubaos-cx/ARUBAWIRED-SWITCH-IMAGE-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-NETWORK-SYNC-MIB",
            "path": "adtran/ADTRAN-AOS-NETWORK-SYNC-MIB.mib"
        },
        {
            "mib": "LINKSYS-ENDOFMIB-MIB",
            "path": "linksys/LINKSYS-ENDOFMIB-MIB.mib"
        },
        {
            "mib": "EXTREME-ENH-DOS-MIB",
            "path": "extreme/EXTREME-ENH-DOS-MIB.mib"
        },
        {
            "mib": "PRVT-MPLS-TE-MIB",
            "path": "telco-systems/binos/PRVT-MPLS-TE-MIB.mib"
        },
        {
            "mib": "IPOMANII-MIB",
            "path": "ingrasys/IPOMANII-MIB.mib"
        },
        {
            "mib": "UBQS-DOT1BRIDGE-MIB",
            "path": "ubiquoss/UBQS-DOT1BRIDGE-MIB.mib"
        },
        {
            "mib": "NBS-VLAN-TAGS-MIB",
            "path": "mrv/NBS-VLAN-TAGS-MIB.mib"
        },
        {
            "mib": "RAISECOM-MPLS-QOS-MIB",
            "path": "raisecom/RAISECOM-MPLS-QOS-MIB.mib"
        },
        {
            "mib": "CISCO-WAN-3G-MIB",
            "path": "cisco/CISCO-WAN-3G-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-GTP-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-GTP-MIB.mib"
        },
        {
            "mib": "LINKSYS-ERRDISABLE-RECOVERY-MIB",
            "path": "linksys/LINKSYS-ERRDISABLE-RECOVERY-MIB.mib"
        },
        {
            "mib": "HH3C-MPLS-LSR-MIB",
            "path": "comware/HH3C-MPLS-LSR-MIB.mib"
        },
        {
            "mib": "PRVT-PROTO-PRIORITY-MIB",
            "path": "telco-systems/binox/PRVT-PROTO-PRIORITY-MIB.mib"
        },
        {
            "mib": "IBMESCON-MIB",
            "path": "ibm/IBMESCON-MIB.mib"
        },
        {
            "mib": "ZYXEL-IES5000-MIB",
            "path": "zyxel/ZYXEL-IES5000-MIB.mib"
        },
        {
            "mib": "ZYXEL-IESCOMMON-MIB",
            "path": "zyxel/ZYXEL-IESCOMMON-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-NETWORKMONITOR",
            "path": "adtran/ADTRAN-AOS-NETWORKMONITOR.mib"
        },
        {
            "mib": "SML-MIB",
            "path": "ibm/SML-MIB.mib"
        },
        {
            "mib": "UBQS-ENTITY-ALARM-MIB",
            "path": "ubiquoss/UBQS-ENTITY-ALARM-MIB.mib"
        },
        {
            "mib": "PRVT-MST-MIB",
            "path": "telco-systems/binos/PRVT-MST-MIB.mib"
        },
        {
            "mib": "EXTREME-ENTITY-MIB",
            "path": "extreme/EXTREME-ENTITY-MIB.mib"
        },
        {
            "mib": "HH3C-MPLS-VPN-BGP-MIB",
            "path": "comware/HH3C-MPLS-VPN-BGP-MIB.mib"
        },
        {
            "mib": "CISCO-WAN-CELL-EXT-MIB",
            "path": "cisco/CISCO-WAN-CELL-EXT-MIB.mib"
        },
        {
            "mib": "HUAWEI-SECURITY-MIB",
            "path": "huawei/HUAWEI-SECURITY-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-OVER-TEMP-PROTECTION-MIB",
            "path": "adtran/ADTRAN-AOS-OVER-TEMP-PROTECTION-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-SYSTEMINFO-MIB",
            "path": "arubaos-cx/ARUBAWIRED-SYSTEMINFO-MIB.mib"
        },
        {
            "mib": "RAISECOM-NAT-MIB",
            "path": "raisecom/RAISECOM-NAT-MIB.mib"
        },
        {
            "mib": "USHA-MIB",
            "path": "ingrasys/USHA-MIB.mib"
        },
        {
            "mib": "NET-SNMP-AGENT-MIB",
            "path": "netsnmp/NET-SNMP-AGENT-MIB.mib"
        },
        {
            "mib": "UBQS-ENTITY-MIB",
            "path": "ubiquoss/UBQS-ENTITY-MIB.mib"
        },
        {
            "mib": "PRVT-QOS-MIB",
            "path": "telco-systems/binox/PRVT-QOS-MIB.mib"
        },
        {
            "mib": "LINKSYS-File",
            "path": "linksys/LINKSYS-File.mib"
        },
        {
            "mib": "CISCO-WAN-OPTIMIZATION-MIB",
            "path": "cisco/CISCO-WAN-OPTIMIZATION-MIB.mib"
        },
        {
            "mib": "PRVT-NETWORK-LOOPBACK-TEST-MIB",
            "path": "telco-systems/binos/PRVT-NETWORK-LOOPBACK-TEST-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-PORT-SECURITY-MIB",
            "path": "adtran/ADTRAN-AOS-PORT-SECURITY-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-TEMPSENSOR-MIB",
            "path": "arubaos-cx/ARUBAWIRED-TEMPSENSOR-MIB.mib"
        },
        {
            "mib": "ServersCheck",
            "path": "serverscheck/ServersCheck.mib"
        },
        {
            "mib": "ZYXEL-MGS3712-MIB",
            "path": "zyxel/ZYXEL-MGS3712-MIB.mib"
        },
        {
            "mib": "NET-SNMP-EXAMPLES-MIB",
            "path": "netsnmp/NET-SNMP-EXAMPLES-MIB.mib"
        },
        {
            "mib": "EXTREME-ESRP-MIB",
            "path": "extreme/EXTREME-ESRP-MIB.mib"
        },
        {
            "mib": "RAISECOM-NDP-MIB",
            "path": "raisecom/RAISECOM-NDP-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-RMPS-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-RMPS-MIB.mib"
        },
        {
            "mib": "HH3C-MPLS-BGP-VPN-MIB",
            "path": "comware/HH3C-MPLS-VPN-MIB.mib"
        },
        {
            "mib": "ZYXEL-MIB",
            "path": "zyxel/ZYXEL-MIB.mib"
        },
        {
            "mib": "UBQS-ENVMON-MIB",
            "path": "ubiquoss/UBQS-ENVMON-MIB.mib"
        },
        {
            "mib": "PRVT-RAPS-MIB",
            "path": "telco-systems/binox/PRVT-RAPS-MIB.mib"
        },
        {
            "mib": "JUNIPER-MIB",
            "path": "junose/JUNIPER-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-POWER",
            "path": "adtran/ADTRAN-AOS-POWER.mib"
        },
        {
            "mib": "ARUBAWIRED-VSF-MIB",
            "path": "arubaos-cx/ARUBAWIRED-VSF-MIB.mib"
        },
        {
            "mib": "HH3C-MPLSEXT-MIB",
            "path": "comware/HH3C-MPLSEXT-MIB.mib"
        },
        {
            "mib": "CISCOSB-1-BONJOUR-SERVICE-MIB",
            "path": "cisco/CISCOSB-1-BONJOUR-SERVICE-MIB.mib"
        },
        {
            "mib": "PRVT-OEM1-PARAMS-MIB",
            "path": "telco-systems/binos/PRVT-OEM1-PARAMS-MIB.mib"
        },
        {
            "mib": "IBMESCONCUB-MIB",
            "path": "ibm/IBMESCONCUB-MIB.mib"
        },
        {
            "mib": "UBQS-ERPS-MIB",
            "path": "ubiquoss/UBQS-ERPS-MIB.mib"
        },
        {
            "mib": "HH3C-MPLSOAM-MIB",
            "path": "comware/HH3C-MPLSOAM-MIB.mib"
        },
        {
            "mib": "RAISECOM-NMS-ACC-MIB",
            "path": "raisecom/RAISECOM-NMS-ACC-MIB.mib"
        },
        {
            "mib": "CISCOSB-3SW2SWTABLES-MIB",
            "path": "cisco/CISCOSB-3SW2SWTABLES-MIB.mib"
        },
        {
            "mib": "NET-SNMP-EXTEND-MIB",
            "path": "netsnmp/NET-SNMP-EXTEND-MIB.mib"
        },
        {
            "mib": "LINKSYS-GREEN-MIB",
            "path": "linksys/LINKSYS-GREEN-MIB.mib"
        },
        {
            "mib": "EXTREME-FDB-MIB",
            "path": "extreme/EXTREME-FDB-MIB.mib"
        },
        {
            "mib": "NET-SNMP-MIB",
            "path": "netsnmp/NET-SNMP-MIB.mib"
        },
        {
            "mib": "RAISECOM-NOTIFICATION-MIB",
            "path": "raisecom/RAISECOM-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-SGW-GTP-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-SGW-GTP-MIB.mib"
        },
        {
            "mib": "HH3C-MPLSOAM-PS-MIB",
            "path": "comware/HH3C-MPLSOAM-PS-MIB.mib"
        },
        {
            "mib": "ZYXEL-NWA-SERIES_v1-4-2",
            "path": "zyxel/ZYXEL-NWA-SERIES_v1-4-2.mib"
        },
        {
            "mib": "UBQS-INTERFACE-MIB",
            "path": "ubiquoss/UBQS-INTERFACE-MIB.mib"
        },
        {
            "mib": "HUAWEI-SECURITY-PKI-MIB",
            "path": "huawei/HUAWEI-SECURITY-PKI-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-QOS",
            "path": "adtran/ADTRAN-AOS-QOS.mib"
        },
        {
            "mib": "PRVT-RAPS-SVC-MIB",
            "path": "telco-systems/binox/PRVT-RAPS-SVC-MIB.mib"
        },
        {
            "mib": "ASCOM-IPDECT-MIB",
            "path": "ascom/ASCOM-IPDECT-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-VSFv2-MIB",
            "path": "arubaos-cx/ARUBAWIRED-VSFv2-MIB.mib"
        },
        {
            "mib": "PRVT-OPR-LED-MANAGEMENT-MIB",
            "path": "telco-systems/binos/PRVT-OPR-LED-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "EXTREME-FILETRANSFER-MIB",
            "path": "extreme/EXTREME-FILETRANSFER-MIB.mib"
        },
        {
            "mib": "RAPID-CLIENT-MIB",
            "path": "nortel/RAPID-CLIENT-MIB.mib"
        },
        {
            "mib": "PRVT-REDUNDANCY-MIB",
            "path": "telco-systems/binox/PRVT-REDUNDANCY-MIB.mib"
        },
        {
            "mib": "CISCOSB-AAA",
            "path": "cisco/CISCOSB-AAA.mib"
        },
        {
            "mib": "RAISECOM-NTP-MIB",
            "path": "raisecom/RAISECOM-NTP-MIB.mib"
        },
        {
            "mib": "ARUBAWIRED-VSX-MIB",
            "path": "arubaos-cx/ARUBAWIRED-VSX-MIB.mib"
        },
        {
            "mib": "CISCOSB-BANNER-MIB",
            "path": "cisco/CISCOSB-BANNER-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-SIP-PROXY-MIB",
            "path": "adtran/ADTRAN-AOS-SIP-PROXY-MIB.mib"
        },
        {
            "mib": "UBQS-IP-MIB",
            "path": "ubiquoss/UBQS-IP-MIB.mib"
        },
        {
            "mib": "HH3C-MPLSTE-MIB",
            "path": "comware/HH3C-MPLSTE-MIB.mib"
        },
        {
            "mib": "PRVT-OSPF-EXTENSION-MIB",
            "path": "telco-systems/binos/PRVT-OSPF-EXTENSION-MIB.mib"
        },
        {
            "mib": "GIGAMON-SNMP-MIB",
            "path": "gigamon/GIGAMON-SNMP-MIB.mib"
        },
        {
            "mib": "RAPID-HA-MIB",
            "path": "nortel/RAPID-HA-MIB.mib"
        },
        {
            "mib": "NET-SNMP-PASS-MIB",
            "path": "netsnmp/NET-SNMP-PASS-MIB.mib"
        },
        {
            "mib": "ZYXEL-PRESTIGE-MIB",
            "path": "zyxel/ZYXEL-PRESTIGE-MIB.mib"
        },
        {
            "mib": "CISCOSB-BaudRate-MIB",
            "path": "cisco/CISCOSB-BaudRate-MIB.mib"
        },
        {
            "mib": "Juniper-MIBs",
            "path": "junose/Juniper-MIBs.mib"
        },
        {
            "mib": "UBQS-LAG-MIB",
            "path": "ubiquoss/UBQS-LAG-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-SIP-REGISTRATION",
            "path": "adtran/ADTRAN-AOS-SIP-REGISTRATION.mib"
        },
        {
            "mib": "LINKSYS-GVRP-MIB",
            "path": "linksys/LINKSYS-GVRP-MIB.mib"
        },
        {
            "mib": "CISCOSB-BONJOUR-MIB",
            "path": "cisco/CISCOSB-BONJOUR-MIB.mib"
        },
        {
            "mib": "PRVT-RESILIENT-LINK-MIB",
            "path": "telco-systems/binox/PRVT-RESILIENT-LINK-MIB.mib"
        },
        {
            "mib": "HH3C-MPM-MIB",
            "path": "comware/HH3C-MPM-MIB.mib"
        },
        {
            "mib": "IPI-CMM-CHASSIS-MIB",
            "path": "ipinfusion/IPI-CMM-CHASSIS-MIB.mib"
        },
        {
            "mib": "RAISECOM-OAM-MIB",
            "path": "raisecom/RAISECOM-OAM-MIB.mib"
        },
        {
            "mib": "UBQS-MISC-MIB",
            "path": "ubiquoss/UBQS-MISC-MIB.mib"
        },
        {
            "mib": "RAPID-INFO-SYSTEM-MIB",
            "path": "nortel/RAPID-INFO-SYSTEM-MIB.mib"
        },
        {
            "mib": "IbmFaultMgmt-MIB",
            "path": "ibm/IbmFaultMgmt-MIB.mib"
        },
        {
            "mib": "PRVT-PORT-SECURITY-MIB",
            "path": "telco-systems/binos/PRVT-PORT-SECURITY-MIB.mib"
        },
        {
            "mib": "EXTREME-NETFLOW-MIB",
            "path": "extreme/EXTREME-NETFLOW-MIB.mib"
        },
        {
            "mib": "NET-SNMP-TC",
            "path": "netsnmp/NET-SNMP-TC.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-SGW-SM-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-SGW-SM-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS-VQM",
            "path": "adtran/ADTRAN-AOS-VQM.mib"
        },
        {
            "mib": "RAPID-CITY",
            "path": "nortel/RAPID-CITY.mib"
        },
        {
            "mib": "CAMBIUM-PMP80211-MIB",
            "path": "cambium/CAMBIUM-PMP80211-MIB.txt.mib"
        },
        {
            "mib": "L-AM3440-A-Private",
            "path": "loop-telecom/L-AM3440-A-Private.mib"
        },
        {
            "mib": "CISCOSB-BRGMACSWITCH-MIB",
            "path": "cisco/CISCOSB-BRGMACSWITCH-MIB.mib"
        },
        {
            "mib": "HUAWEI-SECURITY-STAT-MIB",
            "path": "huawei/HUAWEI-SECURITY-STAT-MIB.mib"
        },
        {
            "mib": "LINKSYS-HWENVIROMENT",
            "path": "linksys/LINKSYS-HWENVIROMENT.mib"
        },
        {
            "mib": "NET-SNMP-VACM-MIB",
            "path": "netsnmp/NET-SNMP-VACM-MIB.mib"
        },
        {
            "mib": "ZYXEL-SAM1216",
            "path": "zyxel/ZYXEL-SAM1216.mib"
        },
        {
            "mib": "UBQS-MPLS-LDP-MIB",
            "path": "ubiquoss/UBQS-MPLS-LDP-MIB.mib"
        },
        {
            "mib": "AcAlarm",
            "path": "audiocodes/AC-ALARM-MIB.mib"
        },
        {
            "mib": "HH3C-MS-MAN-MIB",
            "path": "comware/HH3C-MS-MAN-MIB.mib"
        },
        {
            "mib": "CISCOSB-BRIDGE-SECURITY",
            "path": "cisco/CISCOSB-BRIDGE-SECURITY.mib"
        },
        {
            "mib": "ADTRAN-AOS-VRRP-MIB",
            "path": "adtran/ADTRAN-AOS-VRRP-MIB.mib"
        },
        {
            "mib": "EXTREME-NP-MIB",
            "path": "extreme/EXTREME-NP-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPT-REMOTE-RMC-MIB",
            "path": "raisecom/RAISECOM-OPT-REMOTE-RMC-MIB.mib"
        },
        {
            "mib": "Juniper-MROUTER-MIB",
            "path": "junose/Juniper-MROUTER-MIB.mib"
        },
        {
            "mib": "IBMFRBRS-MIB",
            "path": "ibm/IBMFRBRS-MIB.mib"
        },
        {
            "mib": "LINKSYS-IP",
            "path": "linksys/LINKSYS-IP.mib"
        },
        {
            "mib": "ZYXEL-STACKING-MIB",
            "path": "zyxel/ZYXEL-STACKING-MIB.mib"
        },
        {
            "mib": "PRVT-ROUTE-MIB",
            "path": "telco-systems/binox/PRVT-ROUTE-MIB.mib"
        },
        {
            "mib": "EXTREME-OSPF-MIB",
            "path": "extreme/EXTREME-OSPF-MIB.mib"
        },
        {
            "mib": "PRVT-PORTS-AGGREGATION-MIB",
            "path": "telco-systems/binos/PRVT-PORTS-AGGREGATION-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP250-MIB",
            "path": "cambium/CAMBIUM-PTP250-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-DEVICE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-DEVICE-MIB.mib"
        },
        {
            "mib": "CISCOSB-BRIDGEMIBOBJECTS-MIB",
            "path": "cisco/CISCOSB-BRIDGEMIBOBJECTS-MIB.mib"
        },
        {
            "mib": "HM800MIB",
            "path": "hitachi/HM800MIB.mib"
        },
        {
            "mib": "HH3C-MULTICAST-MIB",
            "path": "comware/HH3C-MULTICAST-MIB.mib"
        },
        {
            "mib": "UBQS-MPLS-MIB",
            "path": "ubiquoss/UBQS-MPLS-MIB.mib"
        },
        {
            "mib": "LINKSYS-ippreflist-MIB",
            "path": "linksys/LINKSYS-ippreflist-MIB.mib"
        },
        {
            "mib": "CISCOSB-CDB-MIB",
            "path": "cisco/CISCOSB-CDB-MIB.mib"
        },
        {
            "mib": "PRVT-PROXY-MANAGER-MIB",
            "path": "telco-systems/binos/PRVT-PROXY-MANAGER-MIB.mib"
        },
        {
            "mib": "RAPID-IPSEC-ENDPOINT-PAIR-MIB",
            "path": "nortel/RAPID-IPSEC-ENDPOINT-PAIR-MIB.mib"
        },
        {
            "mib": "AC-ANALOG-MIB",
            "path": "audiocodes/AC-ANALOG-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOS",
            "path": "adtran/ADTRAN-AOS.mib"
        },
        {
            "mib": "PRVT-RSVP-MIB",
            "path": "telco-systems/binox/PRVT-RSVP-MIB.mib"
        },
        {
            "mib": "IPI-CMM-IPMI-MIB",
            "path": "ipinfusion/IPI-CMM-IPMI-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-ENTITY-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-ENTITY-MIB.mib"
        },
        {
            "mib": "CAMBIUM-PTP650-MIB",
            "path": "cambium/CAMBIUM-PTP650-MIB.mib"
        },
        {
            "mib": "SFA-INFO",
            "path": "ddn/SFA-INFO.mib"
        },
        {
            "mib": "EXTREME-PBQOS-MIB",
            "path": "extreme/EXTREME-PBQOS-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-SM-IP-POOL-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-SM-IP-POOL-MIB.mib"
        },
        {
            "mib": "Juniper-Notification-Log-CONF",
            "path": "junose/Juniper-Notification-Log-CONF.mib"
        },
        {
            "mib": "CISCOSB-CDP-MIB",
            "path": "cisco/CISCOSB-CDP-MIB.mib"
        },
        {
            "mib": "UBQS-MPLS-PW-MIB",
            "path": "ubiquoss/UBQS-MPLS-PW-MIB.mib"
        },
        {
            "mib": "PRVT-PW-TDM-MIB",
            "path": "telco-systems/binos/PRVT-PW-TDM-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOSCPU",
            "path": "adtran/ADTRAN-AOSCPU.mib"
        },
        {
            "mib": "PRVT-SAA-MIB",
            "path": "telco-systems/binox/PRVT-SAA-MIB.mib"
        },
        {
            "mib": "LINKSYS-IpRouter",
            "path": "linksys/LINKSYS-IpRouter.mib"
        },
        {
            "mib": "AC-CONTROL-MIB",
            "path": "audiocodes/AC-CONTROL-MIB.mib"
        },
        {
            "mib": "EXTREME-POE-MIB",
            "path": "extreme/EXTREME-POE-MIB.mib"
        },
        {
            "mib": "HH3C-MULTICAST-SNOOPING-MIB",
            "path": "comware/HH3C-MULTICAST-SNOOPING-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-MODULE-TYPE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-MODULE-TYPE-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOSDOWNLOAD",
            "path": "adtran/ADTRAN-AOSDOWNLOAD.mib"
        },
        {
            "mib": "RAPID-IPSEC-SA-MON-MIB-EXT",
            "path": "nortel/RAPID-IPSEC-SA-MON-MIB-EXT.mib"
        },
        {
            "mib": "ZYXEL-TRANSCEIVER-MIB",
            "path": "zyxel/ZYXEL-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "IBMHPRNCL-MIB",
            "path": "ibm/IBMHPRNCL-MIB.mib"
        },
        {
            "mib": "UBQS-MPLS-RSVP-MIB",
            "path": "ubiquoss/UBQS-MPLS-RSVP-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAY-SM-MIB",
            "path": "junos/JUNIPER-MOBILE-GATEWAY-SM-MIB.mib"
        },
        {
            "mib": "AC-FAULT-TC",
            "path": "audiocodes/AC-FAULT-TC.mib"
        },
        {
            "mib": "LINKSYS-IPv6",
            "path": "linksys/LINKSYS-IPv6.mib"
        },
        {
            "mib": "EXTREME-PORT-MIB",
            "path": "extreme/EXTREME-PORT-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOSFILESYSTEM",
            "path": "adtran/ADTRAN-AOSFILESYSTEM.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-MONITOR-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-MONITOR-MIB.mib"
        },
        {
            "mib": "PRVT-QOS-MIB",
            "path": "telco-systems/binos/PRVT-QOS-MIB.mib"
        },
        {
            "mib": "Juniper-NS-Lookup-CONF",
            "path": "junose/Juniper-NS-Lookup-CONF.mib"
        },
        {
            "mib": "HH3C-NAT-MIB",
            "path": "comware/HH3C-NAT-MIB.mib"
        },
        {
            "mib": "ZYXEL-ZYWALL-MIB",
            "path": "zyxel/ZYXEL-ZYWALL-MIB.mib"
        },
        {
            "mib": "BENU-CGNAT-STATS-MIB",
            "path": "benuos/BENU-CGNAT-STATS-MIB.mib"
        },
        {
            "mib": "HUAWEI-SERVER-IBMC-MIB",
            "path": "huawei/HUAWEI-SERVER-IBMC-MIB.mib"
        },
        {
            "mib": "IPI-MODULE-MIB",
            "path": "ipinfusion/IPI-MODULE-MIB.mib"
        },
        {
            "mib": "PRVT-SERV-MIB",
            "path": "telco-systems/binox/PRVT-SERV-MIB.mib"
        },
        {
            "mib": "CISCOSB-CLI-MIB",
            "path": "cisco/CISCOSB-CLI-MIB.mib"
        },
        {
            "mib": "UBQS-MULTICAST-MIB",
            "path": "ubiquoss/UBQS-MULTICAST-MIB.mib"
        },
        {
            "mib": "LINKSYS-JUMBOFRAMES-MIB",
            "path": "linksys/LINKSYS-JUMBOFRAMES-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOSSNMP",
            "path": "adtran/ADTRAN-AOSSNMP.mib"
        },
        {
            "mib": "HH3C-NDEC-MIB",
            "path": "comware/HH3C-NDEC-MIB.mib"
        },
        {
            "mib": "UPS-MIB",
            "path": "mitsubishi/UPS-MIB.mib"
        },
        {
            "mib": "NMS",
            "path": "smartbyte/NMS.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GATEWAYS",
            "path": "junos/JUNIPER-MOBILE-GATEWAYS.mib"
        },
        {
            "mib": "EXTREME-POS-MIB",
            "path": "extreme/EXTREME-POS-MIB.mib"
        },
        {
            "mib": "RAPID-IPSEC-TUNNEL-MIB",
            "path": "nortel/RAPID-IPSEC-TUNNEL-MIB.mib"
        },
        {
            "mib": "ADTRAN-AOSUNIT",
            "path": "adtran/ADTRAN-AOSUNIT.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-TRANSCEIVER-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "UBQS-NTP-MIB",
            "path": "ubiquoss/UBQS-NTP-MIB.mib"
        },
        {
            "mib": "HH3C-NET-MAN-MIB",
            "path": "comware/HH3C-NET-MAN-MIB.mib"
        },
        {
            "mib": "ZYXEL-ZYWALL-ZLD-COMMON-MIB",
            "path": "zyxel/ZYXEL-ZYWALL-ZLD-COMMON-MIB.mib"
        },
        {
            "mib": "LINKSYS-LLDP-MIB",
            "path": "linksys/LINKSYS-LLDP-MIB.mib"
        },
        {
            "mib": "EXTREME-QOS-MIB",
            "path": "extreme/EXTREME-QOS-MIB.mib"
        },
        {
            "mib": "IBMHPRROUTETEST-MIB",
            "path": "ibm/IBMHPRROUTETEST-MIB.mib"
        },
        {
            "mib": "ADTRAN-COMMON-AOS",
            "path": "adtran/ADTRAN-COMMON-AOS.mib"
        },
        {
            "mib": "PRVT-RESILIENT-LINK-MIB",
            "path": "telco-systems/binos/PRVT-RESILIENT-LINK-MIB.mib"
        },
        {
            "mib": "BENU-CHASSIS-MIB",
            "path": "benuos/BENU-CHASSIS-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-UDETH-INTERFACE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-UDETH-INTERFACE-MIB.mib"
        },
        {
            "mib": "CISCOSB-COPY-MIB",
            "path": "cisco/CISCOSB-COPY-MIB.mib"
        },
        {
            "mib": "CMM3-MIB",
            "path": "cambium/CMM3-MIB.txt.mib"
        },
        {
            "mib": "PRVT-SPANNING-TREE-MIB",
            "path": "telco-systems/binox/PRVT-SPANNING-TREE-MIB.mib"
        },
        {
            "mib": "RAPID-MIB",
            "path": "nortel/RAPID-MIB.mib"
        },
        {
            "mib": "Juniper-OSPF-CONF",
            "path": "junose/Juniper-OSPF-CONF.mib"
        },
        {
            "mib": "UBQS-OSPF-MIB",
            "path": "ubiquoss/UBQS-OSPF-MIB.mib"
        },
        {
            "mib": "HH3C-NPV-MIB",
            "path": "comware/HH3C-NPV-MIB.mib"
        },
        {
            "mib": "SYNSO-UPSMIB",
            "path": "synso/SYNSO-UPSMIB.mib"
        },
        {
            "mib": "CISCOSB-CPU-COUNTERS-MIB",
            "path": "cisco/CISCOSB-CPU-COUNTERS-MIB.mib"
        },
        {
            "mib": "EXTREME-RTSTATS-MIB",
            "path": "extreme/EXTREME-RTSTATS-MIB.mib"
        },
        {
            "mib": "HUAWEI-SITE-MONITOR-MIB",
            "path": "huawei/HUAWEI-SITE-MONITOR-MIB.mib"
        },
        {
            "mib": "AC-KPI-MIB",
            "path": "audiocodes/AC-KPI-MIB.mib"
        },
        {
            "mib": "EXCELIANCE-MIB",
            "path": "haproxy/EXCELIANCE-MIB.mib"
        },
        {
            "mib": "ADTRAN-IF-PERF-HISTORY-MIB",
            "path": "adtran/ADTRAN-IF-PERF-HISTORY-MIB.mib"
        },
        {
            "mib": "LINKSYS-LOCALIZATION-MIB",
            "path": "linksys/LINKSYS-LOCALIZATION-MIB.mib"
        },
        {
            "mib": "HH3C-NQA-MIB",
            "path": "comware/HH3C-NQA-MIB.mib"
        },
        {
            "mib": "PRVT-RING-EPS-MIB",
            "path": "telco-systems/binos/PRVT-RING-EPS-MIB.mib"
        },
        {
            "mib": "BENU-DHCP-MIB",
            "path": "benuos/BENU-DHCP-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILE-GW-SGW-MFWD-MIB",
            "path": "junos/JUNIPER-MOBILE-GW-SGW-MFWD-MIB.mib"
        },
        {
            "mib": "CISCOSB-DEBUGCAPABILITIES-MIB",
            "path": "cisco/CISCOSB-DEBUGCAPABILITIES-MIB.mib"
        },
        {
            "mib": "RAPID-POLICY-MIB",
            "path": "nortel/RAPID-POLICY-MIB.mib"
        },
        {
            "mib": "Juniper-OSPF-MIB",
            "path": "junose/Juniper-OSPF-MIB.mib"
        },
        {
            "mib": "RAISECOM-OPTICAL-UDSFP-INTERFACE-MIB",
            "path": "raisecom/RAISECOM-OPTICAL-UDSFP-INTERFACE-MIB.mib"
        },
        {
            "mib": "TERRAGRAPH-RADIO-MIB",
            "path": "cambium/TERRAGRAPH-RADIO-MIB.mib"
        },
        {
            "mib": "AC-MEDIA-MIB",
            "path": "audiocodes/AC-MEDIA-MIB.mib"
        },
        {
            "mib": "UBQS-PB-MIB",
            "path": "ubiquoss/UBQS-PB-MIB.mib"
        },
        {
            "mib": "ADTRAN-MEF-PER-COS-PER-EVC-PERF-HISTORY-MIB",
            "path": "adtran/ADTRAN-MEF-PER-COS-PER-EVC-PERF-HISTORY-MIB.mib"
        },
        {
            "mib": "CISCOSB-DEVICEPARAMS-MIB",
            "path": "cisco/CISCOSB-DEVICEPARAMS-MIB.mib"
        },
        {
            "mib": "HH3C-NS-MIB",
            "path": "comware/HH3C-NS-MIB.mib"
        },
        {
            "mib": "EXTREME-SERVICES-MIB",
            "path": "extreme/EXTREME-SERVICES-MIB.mib"
        },
        {
            "mib": "BENU-ENTERPRISE-MIB",
            "path": "benuos/BENU-ENTERPRISE-MIB.mib"
        },
        {
            "mib": "WIS-BRIDGE-MIB",
            "path": "wis/WIS-BRIDGE-MIB.mib"
        },
        {
            "mib": "PRVT-STATISTICS-HISTORY-MIB",
            "path": "telco-systems/binox/PRVT-STATISTICS-HISTORY-MIB.mib"
        },
        {
            "mib": "IBMIROC-MIB",
            "path": "ibm/IBMIROC-MIB.mib"
        },
        {
            "mib": "LINKSYS-MAC-BASE-PRIO",
            "path": "linksys/LINKSYS-MAC-BASE-PRIO.mib"
        },
        {
            "mib": "RAPID-SYSTEM-CONFIG-MIB",
            "path": "nortel/RAPID-SYSTEM-CONFIG-MIB.mib"
        },
        {
            "mib": "PRVT-RIP-EXTENSION-MIB",
            "path": "telco-systems/binos/PRVT-RIP-EXTENSION-MIB.mib"
        },
        {
            "mib": "AC-ModularGW-MIB",
            "path": "audiocodes/AC-MODULARGATEWAY-MIB.mib"
        },
        {
            "mib": "RAISECOM-OSPF-MIB",
            "path": "raisecom/RAISECOM-OSPF-MIB.mib"
        },
        {
            "mib": "WIS-MIB",
            "path": "wis/WIS-MIB.mib"
        },
        {
            "mib": "HH3C-NTP-MIB",
            "path": "comware/HH3C-NTP-MIB.mib"
        },
        {
            "mib": "UBQS-PON-LAG-MIB",
            "path": "ubiquoss/UBQS-PON-LAG-MIB.mib"
        },
        {
            "mib": "Juniper-Packet-Mirror-CONF",
            "path": "junose/Juniper-Packet-Mirror-CONF.mib"
        },
        {
            "mib": "ADTRAN-MEF-PER-COS-PER-UNI-PERF-HISTORY-MIB",
            "path": "adtran/ADTRAN-MEF-PER-COS-PER-UNI-PERF-HISTORY-MIB.mib"
        },
        {
            "mib": "BENU-GENERAL-NOTIFICATION-MIB",
            "path": "benuos/BENU-GENERAL-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "RAPID-SYSTEM-STATISTICS-MIB",
            "path": "nortel/RAPID-SYSTEM-STATISTICS-MIB.mib"
        },
        {
            "mib": "WHISP-APS-MIB",
            "path": "cambium/WHISP-APS-MIB.mib"
        },
        {
            "mib": "PRVT-RSVP-MIB",
            "path": "telco-systems/binos/PRVT-RSVP-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILITY-CHARGING-MIB",
            "path": "junos/JUNIPER-MOBILITY-CHARGING-MIB.mib"
        },
        {
            "mib": "EXTREME-SLB-MIB",
            "path": "extreme/EXTREME-SLB-MIB.mib"
        },
        {
            "mib": "HH3C-NVGRE-MIB",
            "path": "comware/HH3C-NVGRE-MIB.mib"
        },
        {
            "mib": "UBQS-PRODUCTS-MIB",
            "path": "ubiquoss/UBQS-PRODUCTS-MIB.mib"
        },
        {
            "mib": "LINKSYS-MIB",
            "path": "linksys/LINKSYS-MIB.mib"
        },
        {
            "mib": "AC-PM-Analog-MIB",
            "path": "audiocodes/AC-PMANALOG-MIB.mib"
        },
        {
            "mib": "CISCOSB-DHCP-MIB",
            "path": "cisco/CISCOSB-DHCP-MIB.mib"
        },
        {
            "mib": "CAMBIUM-MIB",
            "path": "cambium/cnpilote/CAMBIUM-MIB.mib"
        },
        {
            "mib": "PRVT-STATISTICS-MIB",
            "path": "telco-systems/binox/PRVT-STATISTICS-MIB.mib"
        },
        {
            "mib": "OG-CONNECT-MIB",
            "path": "opengear/OG-CONNECT-MIB.mib"
        },
        {
            "mib": "HUAWEI-SLOG-EUDM-MIB",
            "path": "huawei/HUAWEI-SLOG-EUDM-MIB.mib"
        },
        {
            "mib": "Juniper-PACKET-MIRROR-MIB",
            "path": "junose/Juniper-PACKET-MIRROR-MIB.mib"
        },
        {
            "mib": "BENU-HOST-MIB",
            "path": "benuos/BENU-HOST-MIB.mib"
        },
        {
            "mib": "WHISP-BOX-MIBV2-MIB",
            "path": "cambium/WHISP-BOX-MIBV2-MIB.mib"
        },
        {
            "mib": "ADTRAN-MEF-PER-EVC-PERF-HISTORY-MIB",
            "path": "adtran/ADTRAN-MEF-PER-EVC-PERF-HISTORY-MIB.mib"
        },
        {
            "mib": "PRVT-SAA-MIB",
            "path": "telco-systems/binos/PRVT-SAA-MIB.mib"
        },
        {
            "mib": "CISCOSB-DHCPCL-MIB",
            "path": "cisco/CISCOSB-DHCPCL-MIB.mib"
        },
        {
            "mib": "IBMIROCAUTH-MIB",
            "path": "ibm/IBMIROCAUTH-MIB.mib"
        },
        {
            "mib": "HUAWEI-SMAP-MIB",
            "path": "huawei/HUAWEI-SMAP-MIB.mib"
        },
        {
            "mib": "RC-MLT-MIB",
            "path": "nortel/RC-MLT-MIB.mib"
        },
        {
            "mib": "AC-PM-ATM-MIB",
            "path": "audiocodes/AC-PMAtm-MIB.mib"
        },
        {
            "mib": "RAISECOM-OSPFV2-MIB",
            "path": "raisecom/RAISECOM-OSPFV2-MIB.mib"
        },
        {
            "mib": "OG-DATA-MIB",
            "path": "opengear/OG-DATA-MIB.mib"
        },
        {
            "mib": "PRVT-STORM-CTL-MIB",
            "path": "telco-systems/binox/PRVT-STORM-CTL-MIB.mib"
        },
        {
            "mib": "JUNIPER-MOBILITY-SGW-CHARGING-MIB",
            "path": "junos/JUNIPER-MOBILITY-SGW-CHARGING-MIB.mib"
        },
        {
            "mib": "ADTRAN-MEF-PER-UNI-PERF-HISTORY-MIB",
            "path": "adtran/ADTRAN-MEF-PER-UNI-PERF-HISTORY-MIB.mib"
        },
        {
            "mib": "CISCOSB-DHCPv6-CLIENT",
            "path": "cisco/CISCOSB-DHCPv6-CLIENT.mib"
        },
        {
            "mib": "Juniper-PIM-CONF",
            "path": "junose/Juniper-PIM-CONF.mib"
        },
        {
            "mib": "HH3C-OBJECT-INFO-MIB",
            "path": "comware/HH3C-OBJECT-INFO-MIB.mib"
        },
        {
            "mib": "UBQS-QOS-MIB",
            "path": "ubiquoss/UBQS-QOS-MIB.mib"
        },
        {
            "mib": "LINKSYS-MIR-MIB",
            "path": "linksys/LINKSYS-MIR-MIB.mib"
        },
        {
            "mib": "CISCOSB-DHCPv6-RELAY",
            "path": "cisco/CISCOSB-DHCPv6-RELAY.mib"
        },
        {
            "mib": "WHISP-GLOBAL-REG-MIB",
            "path": "cambium/WHISP-GLOBAL-REG-MIB.mib"
        },
        {
            "mib": "EXTREME-SNMPV3-MIB",
            "path": "extreme/EXTREME-SNMPV3-MIB.mib"
        },
        {
            "mib": "DCP-ALARM-MIB",
            "path": "smartoptics/DCP-ALARM-MIB.mib"
        },
        {
            "mib": "IBMIROCDIALOUT-MIB",
            "path": "ibm/IBMIROCDIALOUT-MIB.mib"
        },
        {
            "mib": "ADTRAN-MIB",
            "path": "adtran/ADTRAN-MIB.mib"
        },
        {
            "mib": "BENU-HTTP-CLIENT-MIB",
            "path": "benuos/BENU-HTTP-CLIENT-MIB.mib"
        },
        {
            "mib": "HUAWEI-SMARTLINK-MIB",
            "path": "huawei/HUAWEI-SMARTLINK-MIB.mib"
        },
        {
            "mib": "HH3C-OBJP-MIB",
            "path": "comware/HH3C-OBJP-MIB.mib"
        },
        {
            "mib": "PRVT-SERV-MIB",
            "path": "telco-systems/binos/PRVT-SERV-MIB.mib"
        },
        {
            "mib": "UBQS-REDUNDANCY-MIB",
            "path": "ubiquoss/UBQS-REDUNDANCY-MIB.mib"
        },
        {
            "mib": "RAISECOM-PAE-MIB",
            "path": "raisecom/RAISECOM-PAE-MIB.mib"
        },
        {
            "mib": "RC-VLAN-MIB",
            "path": "nortel/RC-VLAN-MIB.mib"
        },
        {
            "mib": "LINKSYS-MNGINF-MIB",
            "path": "linksys/LINKSYS-MNGINF-MIB.mib"
        },
        {
            "mib": "EXTREME-SOFTWARE-MONITOR-MIB",
            "path": "extreme/EXTREME-SOFTWARE-MONITOR-MIB.mib"
        },
        {
            "mib": "OG-FAILOVER-MIB",
            "path": "opengear/OG-FAILOVER-MIB.mib"
        },
        {
            "mib": "CISCOSB-DHCPv6",
            "path": "cisco/CISCOSB-DHCPv6.mib"
        },
        {
            "mib": "ADTRAN-TC",
            "path": "adtran/ADTRAN-TC.mib"
        },
        {
            "mib": "HUAWEI-SNMP-EXT-MIB",
            "path": "huawei/HUAWEI-SNMP-EXT-MIB.mib"
        },
        {
            "mib": "UBQS-SLD-MIB",
            "path": "ubiquoss/UBQS-SLD-MIB.mib"
        },
        {
            "mib": "HH3C-OFP-MIB",
            "path": "comware/HH3C-OFP-MIB.mib"
        },
        {
            "mib": "OG-HOST-MIB",
            "path": "opengear/OG-HOST-MIB.mib"
        },
        {
            "mib": "RAISECOM-PERF-MIB",
            "path": "raisecom/RAISECOM-PERF-MIB.mib"
        },
        {
            "mib": "PRVT-STATISTICS-CES-MIB",
            "path": "telco-systems/binos/PRVT-STATISTICS-CES-MIB.mib"
        },
        {
            "mib": "JUNIPER-MPLS-LDP-MIB",
            "path": "junos/JUNIPER-MPLS-LDP-MIB.mib"
        },
        {
            "mib": "PRVT-SUPER-VLAN-MIB",
            "path": "telco-systems/binox/PRVT-SUPER-VLAN-MIB.mib"
        },
        {
            "mib": "BENU-HTTP-SERVER-MIB",
            "path": "benuos/BENU-HTTP-SERVER-MIB.mib"
        },
        {
            "mib": "WHISP-SM-MIB",
            "path": "cambium/WHISP-SM-MIB.mib"
        },
        {
            "mib": "RC-VRF-MIB",
            "path": "nortel/RC-VRF-MIB.mib"
        },
        {
            "mib": "CISCOSB-DIGITALKEYMANAGE-MIB",
            "path": "cisco/CISCOSB-DIGITALKEYMANAGE-MIB.mib"
        },
        {
            "mib": "Juniper-PIM-MIB",
            "path": "junose/Juniper-PIM-MIB.mib"
        },
        {
            "mib": "DCP-INTERFACE-MIB",
            "path": "smartoptics/DCP-INTERFACE-MIB.mib"
        },
        {
            "mib": "SAF-ENTERPRISE",
            "path": "saf/SAF-ENTERPRISE.mib"
        },
        {
            "mib": "EXTREME-STACKING-MIB",
            "path": "extreme/EXTREME-STACKING-MIB.mib"
        },
        {
            "mib": "UBQS-SMI",
            "path": "ubiquoss/UBQS-SMI.mib"
        },
        {
            "mib": "OG-OMTELEM-MIB",
            "path": "opengear/OG-OMTELEM-MIB.mib"
        },
        {
            "mib": "HH3C-OID-MIB",
            "path": "comware/HH3C-OID-MIB.mib"
        },
        {
            "mib": "LINKSYS-MODEL-MIB",
            "path": "linksys/LINKSYS-MODEL-MIB.mib"
        },
        {
            "mib": "BENU-IP-MIB",
            "path": "benuos/BENU-IP-MIB.mib"
        },
        {
            "mib": "S5-CHASSIS-MIB",
            "path": "nortel/S5-CHASSIS-MIB.mib"
        },
        {
            "mib": "RAISECOM-PIM-MIB",
            "path": "raisecom/RAISECOM-PIM-MIB.mib"
        },
        {
            "mib": "PRVT-STATISTICSHISTORY-MIB",
            "path": "telco-systems/binos/PRVT-STATISTICSHISTORY-MIB.mib"
        },
        {
            "mib": "WHISP-TCV2-MIB",
            "path": "cambium/WHISP-TCV2-MIB.mib"
        },
        {
            "mib": "JUNIPER-NAT-MIB",
            "path": "junos/JUNIPER-NAT-MIB.mib"
        },
        {
            "mib": "AC-PM-Control-MIB",
            "path": "audiocodes/AC-PMCONTROL-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-ACCESS-LIST-MIB",
            "path": "telco-systems/binox/PRVT-SWITCH-ACCESS-LIST-MIB.mib"
        },
        {
            "mib": "Juniper-Ping-CONF",
            "path": "junose/Juniper-Ping-CONF.mib"
        },
        {
            "mib": "DCP-LINKVIEW-MIB",
            "path": "smartoptics/DCP-LINKVIEW-MIB.mib"
        },
        {
            "mib": "IBMIROCRLAN-MIB",
            "path": "ibm/IBMIROCRLAN-MIB.mib"
        },
        {
            "mib": "SAF-INTEGRA-MIB",
            "path": "saf/SAF-INTEGRA-MIB.mib"
        },
        {
            "mib": "OG-PATTERN-MIB",
            "path": "opengear/OG-PATTERN-MIB.mib"
        },
        {
            "mib": "RAISECOM-POE-MIB",
            "path": "raisecom/RAISECOM-POE-MIB.mib"
        },
        {
            "mib": "CISCOSB-Dlf-MIB",
            "path": "cisco/CISCOSB-Dlf-MIB.mib"
        },
        {
            "mib": "HH3C-OSPF-MIB",
            "path": "comware/HH3C-OSPF-MIB.mib"
        },
        {
            "mib": "EXTREME-STP-EXTENSIONS-MIB",
            "path": "extreme/EXTREME-STP-EXTENSIONS-MIB.mib"
        },
        {
            "mib": "UBQS-SNMP-MIB",
            "path": "ubiquoss/UBQS-SNMP-MIB.mib"
        },
        {
            "mib": "BENU-IPPOOL-MIB",
            "path": "benuos/BENU-IPPOOL-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-IPVLAN-MIB",
            "path": "telco-systems/binox/PRVT-SWITCH-IPVLAN-MIB.mib"
        },
        {
            "mib": "DCP-MIB",
            "path": "smartoptics/DCP-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-ACCESS-LIST-MIB",
            "path": "telco-systems/binos/PRVT-SWITCH-ACCESS-LIST-MIB.mib"
        },
        {
            "mib": "HUAWEI-SNMP-NOTIFICATION-MIB",
            "path": "huawei/HUAWEI-SNMP-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "S5-ROOT-MIB",
            "path": "nortel/S5-ROOT-MIB.mib"
        },
        {
            "mib": "AC-PM-Media-MIB",
            "path": "audiocodes/AC-PMMEDIA-MIB.mib"
        },
        {
            "mib": "RAISECOM-PON-DEVICE-MIB",
            "path": "raisecom/RAISECOM-PON-DEVICE-MIB.mib"
        },
        {
            "mib": "HH3C-PBR-MIB",
            "path": "comware/HH3C-PBR-MIB.mib"
        },
        {
            "mib": "OG-PRODUCTS-MIB",
            "path": "opengear/OG-PRODUCTS-MIB.mib"
        },
        {
            "mib": "CISCOSB-DNSCL-MIB",
            "path": "cisco/CISCOSB-DNSCL-MIB.mib"
        },
        {
            "mib": "SAF-INTEGRAB-MIB",
            "path": "saf/SAF-INTEGRAB-MIB.mib"
        },
        {
            "mib": "LINKSYS-MULTISESSIONTERMINAL-MIB",
            "path": "linksys/LINKSYS-MULTISESSIONTERMINAL-MIB.mib"
        },
        {
            "mib": "EXTREME-SYSTEM-MIB",
            "path": "extreme/EXTREME-SYSTEM-MIB.mib"
        },
        {
            "mib": "UBQS-SYSLOG-MIB",
            "path": "ubiquoss/UBQS-SYSLOG-MIB.mib"
        },
        {
            "mib": "HUAWEI-SSH-MIB",
            "path": "huawei/HUAWEI-SSH-MIB.mib"
        },
        {
            "mib": "IBMNETU-MIB",
            "path": "ibm/IBMNETU-MIB.mib"
        },
        {
            "mib": "BENU-KAFKA-CLIENT-MIB",
            "path": "benuos/BENU-KAFKA-CLIENT-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-EFM-OAM-MIB",
            "path": "telco-systems/binos/PRVT-SWITCH-EFM-OAM-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-MIB",
            "path": "telco-systems/binox/PRVT-SWITCH-MIB.mib"
        },
        {
            "mib": "Juniper-Policy-Manager-CONF",
            "path": "junose/Juniper-Policy-Manager-CONF.mib"
        },
        {
            "mib": "RAISECOM-PONSERIES-BASE-MIB",
            "path": "raisecom/RAISECOM-PONSERIES-BASE-MIB.mib"
        },
        {
            "mib": "AC-PM-MediaServer-MIB",
            "path": "audiocodes/AC-PMMEDIASERVER-MIB.mib"
        },
        {
            "mib": "DCP-OCH-MIB",
            "path": "smartoptics/DCP-OCH-MIB.mib"
        },
        {
            "mib": "S5-TCS-MIB",
            "path": "nortel/S5-TCS-MIB.mib"
        },
        {
            "mib": "HH3C-PEX-MIB",
            "path": "comware/HH3C-PEX-MIB.mib"
        },
        {
            "mib": "MCAFEE-MWG-MIB",
            "path": "mcafee/MCAFEE-MWG-MIB.mib"
        },
        {
            "mib": "UBQS-SYSRSC-MIB",
            "path": "ubiquoss/UBQS-SYSRSC-MIB.mib"
        },
        {
            "mib": "OG-SENSOR-MIB",
            "path": "opengear/OG-SENSOR-MIB.mib"
        },
        {
            "mib": "EXTREME-TRAP-MIB",
            "path": "extreme/EXTREME-TRAP-MIB.mib"
        },
        {
            "mib": "LINKSYS-openflow-MIB",
            "path": "linksys/LINKSYS-openflow-MIB.mib"
        },
        {
            "mib": "JUNIPER-OAM-MIB",
            "path": "junos/JUNIPER-OAM-MIB.mib"
        },
        {
            "mib": "RAISECOM-PONSERIES-TC",
            "path": "raisecom/RAISECOM-PONSERIES-TC.mib"
        },
        {
            "mib": "MSERIES-ALARM-MIB",
            "path": "smartoptics/MSERIES-ALARM-MIB.mib"
        },
        {
            "mib": "HUAWEI-SSL-MIB",
            "path": "huawei/HUAWEI-SSL-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-EV-PROP-MIB",
            "path": "telco-systems/binos/PRVT-SWITCH-EV-PROP-MIB.mib"
        },
        {
            "mib": "CISCOSB-DOT1X-MIB",
            "path": "cisco/CISCOSB-DOT1X-MIB.mib"
        },
        {
            "mib": "MCAFEE-SMI",
            "path": "mcafee/MCAFEE-SMI.mib"
        },
        {
            "mib": "AC-PM-PSTN-MIB",
            "path": "audiocodes/AC-PMPSTN-MIB.mib"
        },
        {
            "mib": "HH3C-PORT-SECURITY-MIB",
            "path": "comware/HH3C-PORT-SECURITY-MIB.mib"
        },
        {
            "mib": "PRVT-SYNC-ETHERNET-MIB",
            "path": "telco-systems/binox/PRVT-SYNC-ETHERNET-MIB.mib"
        },
        {
            "mib": "SWPRIMGMT-MIB",
            "path": "nortel/SWPRIMGMT-MIB.mib"
        },
        {
            "mib": "MSERIES-ENVMON-MIB",
            "path": "smartoptics/MSERIES-ENVMON-MIB.mib"
        },
        {
            "mib": "EXTREME-TRAPPOLL-MIB",
            "path": "extreme/EXTREME-TRAPPOLL-MIB.mib"
        },
        {
            "mib": "UBQS-SYSTEM-ACCESS-MIB",
            "path": "ubiquoss/UBQS-SYSTEM-ACCESS-MIB.mib"
        },
        {
            "mib": "Juniper-POLICY-MIB",
            "path": "junose/Juniper-POLICY-MIB.mib"
        },
        {
            "mib": "BENU-PLATFORM-DEFN-MIB",
            "path": "benuos/BENU-PLATFORM-DEFN-MIB.mib"
        },
        {
            "mib": "IBMTCPIPMVS-MIB",
            "path": "ibm/IBMTCPIPMVS-MIB.mib"
        },
        {
            "mib": "CISCOSB-EEE-MIB",
            "path": "cisco/CISCOSB-EEE-MIB.mib"
        },
        {
            "mib": "JUNIPER-OTN-MIB",
            "path": "junos/JUNIPER-OTN-MIB.mib"
        },
        {
            "mib": "SAF-INTEGRAE-MIB",
            "path": "saf/SAF-INTEGRAE-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-FIB-MIB",
            "path": "telco-systems/binos/PRVT-SWITCH-FIB-MIB.mib"
        },
        {
            "mib": "HH3C-PORTAL-MIB",
            "path": "comware/HH3C-PORTAL-MIB.mib"
        },
        {
            "mib": "RAISECOM-PORTSTATISTIC-MIB",
            "path": "raisecom/RAISECOM-PORTSTATISTIC-MIB.mib"
        },
        {
            "mib": "PRVT-SYS-MON-MIB",
            "path": "telco-systems/binox/PRVT-SYS-MON-MIB.mib"
        },
        {
            "mib": "HUAWEI-STACK-MIB",
            "path": "huawei/HUAWEI-STACK-MIB.mib"
        },
        {
            "mib": "OG-SIGNAL-MIB",
            "path": "opengear/OG-SIGNAL-MIB.mib"
        },
        {
            "mib": "CISCOSB-EMBWEB-MIB",
            "path": "cisco/CISCOSB-EMBWEB-MIB.mib"
        },
        {
            "mib": "AC-PM-SS7-MIB",
            "path": "audiocodes/AC-PMSS7-MIB.mib"
        },
        {
            "mib": "SAF-INTEGRAW-MIB",
            "path": "saf/SAF-INTEGRAW-MIB.mib"
        },
        {
            "mib": "MSERIES-MIB",
            "path": "smartoptics/MSERIES-MIB.mib"
        },
        {
            "mib": "SYNOPTICS-ROOT-MIB",
            "path": "nortel/SYNOPTICS-ROOT-MIB.mib"
        },
        {
            "mib": "LINKSYS-PHY-MIB",
            "path": "linksys/LINKSYS-PHY-MIB.mib"
        },
        {
            "mib": "PRVT-SWITCH-IPVLAN-MIB",
            "path": "telco-systems/binos/PRVT-SWITCH-IPVLAN-MIB.mib"
        },
        {
            "mib": "HH3C-POS-MIB",
            "path": "comware/HH3C-POS-MIB.mib"
        },
        {
            "mib": "NPT-EQUIPMENT-MIB",
            "path": "ribbon/NPT-EQUIPMENT-MIB.mib"
        },
        {
            "mib": "OG-SMI-MIB",
            "path": "opengear/OG-SMI-MIB.mib"
        },
        {
            "mib": "BENU-PLATFORM-MIB",
            "path": "benuos/BENU-PLATFORM-MIB.mib"
        },
        {
            "mib": "UBQS-SYSTEM-MIB",
            "path": "ubiquoss/UBQS-SYSTEM-MIB.mib"
        },
        {
            "mib": "JUNIPER-PAE-EXTENSION-MIB",
            "path": "junos/JUNIPER-PAE-EXTENSION-MIB.mib"
        },
        {
            "mib": "CISCOSB-ENDOFMIB-MIB",
            "path": "cisco/CISCOSB-ENDOFMIB-MIB.mib"
        },
        {
            "mib": "RAISECOM-POWERMONITOR-MIB",
            "path": "raisecom/RAISECOM-POWERMONITOR-MIB.mib"
        },
        {
            "mib": "EXTREME-V2TRAP-MIB",
            "path": "extreme/EXTREME-V2TRAP-MIB.mib"
        },
        {
            "mib": "AC-PM-System-MIB",
            "path": "audiocodes/AC-PMSYSTEM-MIB.mib"
        },
        {
            "mib": "HH3C-POSA-MIB",
            "path": "comware/HH3C-POSA-MIB.mib"
        },
        {
            "mib": "NPT-ROOT-MIB",
            "path": "ribbon/NPT-ROOT-MIB.mib"
        },
        {
            "mib": "UBQS-TC",
            "path": "ubiquoss/UBQS-TC.mib"
        },
        {
            "mib": "PRVT-SWITCH-MIB",
            "path": "telco-systems/binos/PRVT-SWITCH-MIB.mib"
        },
        {
            "mib": "IMM-MIB",
            "path": "ibm/IMM-MIB.mib"
        },
        {
            "mib": "MSERIES-PORT-MIB",
            "path": "smartoptics/MSERIES-PORT-MIB.mib"
        },
        {
            "mib": "EXTREME-VC-MIB",
            "path": "extreme/EXTREME-VC-MIB.mib"
        },
        {
            "mib": "Juniper-PPP-CONF",
            "path": "junose/Juniper-PPP-CONF.mib"
        },
        {
            "mib": "OG-STATUS-MIB",
            "path": "opengear/OG-STATUS-MIB.mib"
        },
        {
            "mib": "LINKSYS-Physicaldescription-MIB",
            "path": "linksys/LINKSYS-Physicaldescription-MIB.mib"
        },
        {
            "mib": "CISCOSB-ERRDISABLE-RECOVERY-MIB",
            "path": "cisco/CISCOSB-ERRDISABLE-RECOVERY-MIB.mib"
        },
        {
            "mib": "AATS-MIB",
            "path": "bats/AATS-MIB.mib"
        },
        {
            "mib": "SAF-INTEGRAX-MIB",
            "path": "saf/SAF-INTEGRAX-MIB.mib"
        },
        {
            "mib": "PRVT-TE-PARAM-MIB",
            "path": "telco-systems/binox/PRVT-TE-PARAM-MIB.mib"
        },
        {
            "mib": "HUAWEI-STORAGE-HARDWARE-MIB",
            "path": "huawei/HUAWEI-STORAGE-HARDWARE-MIB.mib"
        },
        {
            "mib": "HH3C-POWER-ETH-EXT-MIB",
            "path": "comware/HH3C-POWER-ETH-EXT-MIB.mib"
        },
        {
            "mib": "RADLAN-AAA",
            "path": "radlan/RADLAN-AAA.mib"
        },
        {
            "mib": "BENU-PLATFORM-SERVICE-MIB",
            "path": "benuos/BENU-PLATFORM-SERVICE-MIB.mib"
        },
        {
            "mib": "PRVT-SYS-INFO-MIB",
            "path": "telco-systems/binos/PRVT-SYS-INFO-MIB.mib"
        },
        {
            "mib": "OG-STATUSv2-MIB",
            "path": "opengear/OG-STATUSv2-MIB.mib"
        },
        {
            "mib": "CISCOSB-EVENTS-MIB",
            "path": "cisco/CISCOSB-EVENTS-MIB.mib"
        },
        {
            "mib": "JUNIPER-PFE-MIB",
            "path": "junos/JUNIPER-PFE-MIB.mib"
        },
        {
            "mib": "AC-PSTN-MIB",
            "path": "audiocodes/AC-PSTN-MIB.mib"
        },
        {
            "mib": "NPT-SYSTEM-MIB",
            "path": "ribbon/NPT-SYSTEM-MIB.mib"
        },
        {
            "mib": "RAISECOM-PPPOEAGENT-MIB",
            "path": "raisecom/RAISECOM-PPPOEAGENT-MIB.mib"
        },
        {
            "mib": "EXTREME-VLAN-MIB",
            "path": "extreme/EXTREME-VLAN-MIB.mib"
        },
        {
            "mib": "RADLAN-BaudRate-MIB",
            "path": "radlan/RADLAN-BaudRate-MIB.mib"
        },
        {
            "mib": "MSERIES-TC",
            "path": "smartoptics/MSERIES-TC.mib"
        },
        {
            "mib": "BENU-RADIUS-MIB",
            "path": "benuos/BENU-RADIUS-MIB.mib"
        },
        {
            "mib": "HH3C-REDUNDANCY-POWER-MIB",
            "path": "comware/HH3C-POWER-MIB.mib"
        },
        {
            "mib": "PRVT-SYS-MON-MIB",
            "path": "telco-systems/binos/PRVT-SYS-MON-MIB.mib"
        },
        {
            "mib": "SAF-IPRADIO",
            "path": "saf/SAF-IPRADIO.mib"
        },
        {
            "mib": "NPT-TC-MIB",
            "path": "ribbon/NPT-TC-MIB.mib"
        },
        {
            "mib": "LINKSYS-POE-MIB",
            "path": "linksys/LINKSYS-POE-MIB.mib"
        },
        {
            "mib": "PRVT-TEMIB-ENTITY-MIB",
            "path": "telco-systems/binox/PRVT-TEMIB-ENTITY-MIB.mib"
        },
        {
            "mib": "RADLAN-COPY-MIB",
            "path": "radlan/RADLAN-COPY-MIB.mib"
        },
        {
            "mib": "HH3C-PPP-MIB",
            "path": "comware/HH3C-PPP-MIB.mib"
        },
        {
            "mib": "AH-INTERFACE-MIB",
            "path": "aerohive/AH-INTERFACE-MIB.mib"
        },
        {
            "mib": "JUNIPER-PING-MIB",
            "path": "junos/JUNIPER-PING-MIB.mib"
        },
        {
            "mib": "CISCOSB-File",
            "path": "cisco/CISCOSB-File.mib"
        },
        {
            "mib": "BENU-SUB-TUNNEL-MIB",
            "path": "benuos/BENU-SUB-TUNNEL-MIB.mib"
        },
        {
            "mib": "OG-UPS-MIB",
            "path": "opengear/OG-UPS-MIB.mib"
        },
        {
            "mib": "SO-MIB",
            "path": "smartoptics/SO-MIB.mib"
        },
        {
            "mib": "IMMALERT-MIB",
            "path": "ibm/IMMALERT-MIB.mib"
        },
        {
            "mib": "QUIDOS-MIB",
            "path": "papouch/QUIDOS-MIB.mib"
        },
        {
            "mib": "RAISECOM-PTP-MIB",
            "path": "raisecom/RAISECOM-PTP-MIB.mib"
        },
        {
            "mib": "Juniper-PPP-MIB",
            "path": "junose/Juniper-PPP-MIB.mib"
        },
        {
            "mib": "PRVT-TWAMP-MIB",
            "path": "telco-systems/binox/PRVT-TWAMP-MIB.mib"
        },
        {
            "mib": "HUAWEI-STORAGE-NAS-MIB",
            "path": "huawei/HUAWEI-STORAGE-NAS-MIB.mib"
        },
        {
            "mib": "EXTREME-WIRELESS-MIB",
            "path": "extreme/EXTREME-WIRELESS-MIB.mib"
        },
        {
            "mib": "AC-QOS-MIB",
            "path": "audiocodes/AC-QOS-MIB.mib"
        },
        {
            "mib": "RADLAN-DNSCL-MIB",
            "path": "radlan/RADLAN-DNSCL-MIB.mib"
        },
        {
            "mib": "AH-MRP-MIB",
            "path": "aerohive/AH-MRP-MIB.mib"
        },
        {
            "mib": "PRVT-TE-PARAM-MIB",
            "path": "telco-systems/binos/PRVT-TE-PARAM-MIB.mib"
        },
        {
            "mib": "OGTRAP-MIB",
            "path": "opengear/OGTRAP-MIB.mib"
        },
        {
            "mib": "RAISECOM-QINQ-MIB",
            "path": "raisecom/RAISECOM-QINQ-MIB.mib"
        },
        {
            "mib": "EXTREMEdot11AP-MIB",
            "path": "extreme/EXTREMEDOT11AP-MIB.mib"
        },
        {
            "mib": "JUNIPER-PMon-MIB",
            "path": "junos/JUNIPER-PMon-MIB.mib"
        },
        {
            "mib": "SAF-MPMUX-MIB",
            "path": "saf/SAF-MPMUX-MIB.mib"
        },
        {
            "mib": "HH3C-PPP-OVER-SONET-MIB",
            "path": "comware/HH3C-PPP-OVER-SONET-MIB.mib"
        },
        {
            "mib": "AH-SMI-MIB",
            "path": "aerohive/AH-SMI-MIB.mib"
        },
        {
            "mib": "SO-TC-MIB",
            "path": "smartoptics/SO-TC-MIB.mib"
        },
        {
            "mib": "PRVT-VENDORDEF-MIB",
            "path": "telco-systems/binox/PRVT-VENDORDEF-MIB.mib"
        },
        {
            "mib": "RADLAN-File",
            "path": "radlan/RADLAN-File.mib"
        },
        {
            "mib": "HH3C-PPPOE-SERVER-MIB",
            "path": "comware/HH3C-PPPOE-SERVER-MIB.mib"
        },
        {
            "mib": "LINKSYS-POLICY-MIB",
            "path": "linksys/LINKSYS-POLICY-MIB.mib"
        },
        {
            "mib": "BENU-SYSLOG-MIB",
            "path": "benuos/BENU-SYSLOG-MIB.mib"
        },
        {
            "mib": "CISCOSB-FINDIT",
            "path": "cisco/CISCOSB-FINDIT.mib"
        },
        {
            "mib": "AH-SYSTEM-MIB",
            "path": "aerohive/AH-SYSTEM-MIB.mib"
        },
        {
            "mib": "RAISECOM-QOS-MIB",
            "path": "raisecom/RAISECOM-QOS-MIB.mib"
        },
        {
            "mib": "RADLAN-HWENVIROMENT",
            "path": "radlan/RADLAN-HWENVIROMENT.mib"
        },
        {
            "mib": "EXTREMEdot11f-MIB",
            "path": "extreme/EXTREMEDOT11F-MIB.mib"
        },
        {
            "mib": "AC-SS7-MIB",
            "path": "audiocodes/AC-SS7-MIB.mib"
        },
        {
            "mib": "NMS-OPTICAL-MIB",
            "path": "smartbyte/NMS-OPTICAL-MIB.mib"
        },
        {
            "mib": "PRVT-TEMIB-ENTITY-MIB",
            "path": "telco-systems/binos/PRVT-TEMIB-ENTITY-MIB.mib"
        },
        {
            "mib": "Juniper-PPP-Profile-CONF",
            "path": "junose/Juniper-PPP-Profile-CONF.mib"
        },
        {
            "mib": "OGTRAPv2-MIB",
            "path": "opengear/OGTRAPv2-MIB.mib"
        },
        {
            "mib": "NETGEAR-BOXSERVICES-PRIVATE-MIB",
            "path": "netgear/NETGEAR-BOXSERVICES-PRIVATE-MIB.mib"
        },
        {
            "mib": "PRVT-VRRP-MIB",
            "path": "telco-systems/binox/PRVT-VRRP-MIB.mib"
        },
        {
            "mib": "BENU-TWAG-STATS-MIB",
            "path": "benuos/BENU-TWAG-STATS-MIB.mib"
        },
        {
            "mib": "NWAYSMSS-MIB",
            "path": "ibm/NWAYSMSS-MIB.mib"
        },
        {
            "mib": "CISCOSB-GREEN-MIB",
            "path": "cisco/CISCOSB-GREEN-MIB.mib"
        },
        {
            "mib": "RAISECOM-RCFT-MIB",
            "path": "raisecom/RAISECOM-RCFT-MIB.mib"
        },
        {
            "mib": "RADLAN-IP",
            "path": "radlan/RADLAN-IP.mib"
        },
        {
            "mib": "HA-MIB",
            "path": "extreme/HA-MIB.mib"
        },
        {
            "mib": "PRVT-UPS-MIB",
            "path": "telco-systems/binos/PRVT-UPS-MIB.mib"
        },
        {
            "mib": "NETGEAR-REF-MIB",
            "path": "netgear/NETGEAR-REF-MIB.mib"
        },
        {
            "mib": "AC-SYSTEM-MIB",
            "path": "audiocodes/AC-SYSTEM-MIB.mib"
        },
        {
            "mib": "ICOTERA-I6400-SERIES",
            "path": "icotera/ICOTERA-I6400-SERIES-MIB.mib"
        },
        {
            "mib": "RADLAN-IPv6",
            "path": "radlan/RADLAN-IPv6.mib"
        },
        {
            "mib": "LINKSYS-ProtectedPorts-MIB",
            "path": "linksys/LINKSYS-ProtectedPorts-MIB.mib"
        },
        {
            "mib": "HH3C-PRODUCT-ID-MIB",
            "path": "comware/HH3C-PRODUCT-ID-MIB.mib"
        },
        {
            "mib": "NMS-SYS-MIB",
            "path": "smartbyte/NMS-SYS-MIB.mib"
        },
        {
            "mib": "TELDAT-MIB",
            "path": "teldat/TELDAT-MIB.mib"
        },
        {
            "mib": "RADLAN-LLDP-MIB",
            "path": "radlan/RADLAN-LLDP-MIB.mib"
        },
        {
            "mib": "AH_TRAP_MIB",
            "path": "aerohive/AH_TRAP_MIB.mib"
        },
        {
            "mib": "CISCOSB-GVRP-MIB",
            "path": "cisco/CISCOSB-GVRP-MIB.mib"
        },
        {
            "mib": "NETGEAR-SMART-SWITCHING-MIB",
            "path": "netgear/NETGEAR-SMART-SWITCHING-MIB.mib"
        },
        {
            "mib": "RAISECOM-RCMP-MIB",
            "path": "raisecom/RAISECOM-RCMP-MIB.mib"
        },
        {
            "mib": "DB7001-MIB",
            "path": "deva/DB7001-MIB.mib"
        },
        {
            "mib": "Juniper-PPP-PROFILE-MIB",
            "path": "junose/Juniper-PPP-PROFILE-MIB.mib"
        },
        {
            "mib": "HH3C-PROT-PRIORITY-MIB",
            "path": "comware/HH3C-PROT-PRIORITY-MIB.mib"
        },
        {
            "mib": "AUDIOCODES-TYPES-MIB",
            "path": "audiocodes/AC-TYPES.mib"
        },
        {
            "mib": "SWBASE-MIB",
            "path": "extreme/SWBASE-MIB.mib"
        },
        {
            "mib": "HUAWEI-STORAGE-SPACE-MIB",
            "path": "huawei/HUAWEI-STORAGE-SPACE-MIB.mib"
        },
        {
            "mib": "PRVT-Y1564-MIB",
            "path": "telco-systems/binos/PRVT-Y1564-MIB.mib"
        },
        {
            "mib": "CISCOSB-HWENVIROMENT",
            "path": "cisco/CISCOSB-HWENVIROMENT.mib"
        },
        {
            "mib": "BENU-VLAN-MIB",
            "path": "benuos/BENU-VLAN-MIB.mib"
        },
        {
            "mib": "JUNIPER-POWER-SUPPLY-UNIT-MIB",
            "path": "junos/JUNIPER-POWER-SUPPLY-UNIT-MIB.mib"
        },
        {
            "mib": "TELDAT-MON-CommonInfo-MIB",
            "path": "teldat/TELDAT-MON-CommonInfo-MIB.mib"
        },
        {
            "mib": "ICOTERA-I6800-SERIES",
            "path": "icotera/ICOTERA-I6800-SERIES-MIB.mib"
        },
        {
            "mib": "RAISECOM-RELAY-MIB",
            "path": "raisecom/RAISECOM-RELAY-MIB.mib"
        },
        {
            "mib": "CORERO-CMS-CLUSTERS-MIB",
            "path": "corero/CORERO-CMS-CLUSTERS-MIB.mib"
        },
        {
            "mib": "LINKSYS-QOS-CLI-MIB",
            "path": "linksys/LINKSYS-QOS-CLI-MIB.mib"
        },
        {
            "mib": "ETHERNET-MIB",
            "path": "exalt/ETHERNET-MIB.mib"
        },
        {
            "mib": "FJDARY-E150",
            "path": "fujitsu/FJDARY-E150.mib"
        },
        {
            "mib": "BENU-WAG-MIB",
            "path": "benuos/BENU-WAG-MIB.mib"
        },
        {
            "mib": "RADLAN-LOCALIZATION-MIB",
            "path": "radlan/RADLAN-LOCALIZATION-MIB.mib"
        },
        {
            "mib": "HH3C-PROTOCOL-VLAN-MIB",
            "path": "comware/HH3C-PROTOCOL-VLAN-MIB.mib"
        },
        {
            "mib": "AC-V5-MIB",
            "path": "audiocodes/AC-V5-MIB.mib"
        },
        {
            "mib": "Juniper-PPPoE-CONF",
            "path": "junose/Juniper-PPPoE-CONF.mib"
        },
        {
            "mib": "NETGEAR-SWITCHING-MIB",
            "path": "netgear/NETGEAR-SWITCHING-MIB.mib"
        },
        {
            "mib": "JUNIPER-PW-ATM-MIB",
            "path": "junos/JUNIPER-PW-ATM-MIB.mib"
        },
        {
            "mib": "LINKSYS-rlBrgMcMngr-MIB",
            "path": "linksys/LINKSYS-rlBrgMcMngr-MIB.mib"
        },
        {
            "mib": "PURESTORAGE-MIB",
            "path": "purestorage/PURESTORAGE-MIB.mib"
        },
        {
            "mib": "RADLAN-MAC-BASE-PRIO",
            "path": "radlan/RADLAN-MAC-BASE-PRIO.mib"
        },
        {
            "mib": "FSC-SERVERCONTROL2-MIB",
            "path": "fujitsu/FSC-SERVERCONTROL2-MIB.mib"
        },
        {
            "mib": "HH3C-PU-MAN-MIB",
            "path": "comware/HH3C-PU-MAN-MIB.mib"
        },
        {
            "mib": "ExaltComm-TRAPS-MIB",
            "path": "exalt/ExaltComm-TRAPS-MIB.mib"
        },
        {
            "mib": "RAISECOM-REMOTE-MANAGEMENT-LOCAL-MIB",
            "path": "raisecom/RAISECOM-REMOTE-MANAGEMENT-LOCAL-MIB.mib"
        },
        {
            "mib": "SYSTEM-MIB",
            "path": "extreme/SYSTEM-MIB.mib"
        },
        {
            "mib": "TTDP-MIB",
            "path": "westermo/TTDP-MIB.mib"
        },
        {
            "mib": "CORERO-CMS-DEVICES-MIB",
            "path": "corero/CORERO-CMS-DEVICES-MIB.mib"
        },
        {
            "mib": "DKSF-70-6-X-X-1",
            "path": "netping/DKSF-70-6-X-X-1.mib"
        },
        {
            "mib": "HH3C-PVST-MIB",
            "path": "comware/HH3C-PVST-MIB.mib"
        },
        {
            "mib": "LINKSYS-rlBrgMulticast-MIB",
            "path": "linksys/LINKSYS-rlBrgMulticast-MIB.mib"
        },
        {
            "mib": "ExaltComm",
            "path": "exalt/ExaltComm.mib"
        },
        {
            "mib": "RAISECOM-REMOTE-MANAGEMENT-REMOTE-MIB",
            "path": "raisecom/RAISECOM-REMOTE-MANAGEMENT-REMOTE-MIB.mib"
        },
        {
            "mib": "RADLAN-MIB",
            "path": "radlan/RADLAN-MIB.mib"
        },
        {
            "mib": "TELDAT-MON-CPU-MIB",
            "path": "teldat/TELDAT-MON-CPU-MIB.mib"
        },
        {
            "mib": "HUAWEI-SUPERLAG-MIB",
            "path": "huawei/HUAWEI-SUPERLAG-MIB.mib"
        },
        {
            "mib": "SNIA-SML-MIB",
            "path": "ibm/SNIA-SML-MIB.mib"
        },
        {
            "mib": "HH3C-QINQ-MIB",
            "path": "comware/HH3C-QINQ-MIB.mib"
        },
        {
            "mib": "JUNIPER-PW-TDM-MIB",
            "path": "junos/JUNIPER-PW-TDM-MIB.mib"
        },
        {
            "mib": "V1600D",
            "path": "vsolution/V1600D.mib"
        },
        {
            "mib": "BENU-WAG-STATS-MIB",
            "path": "benuos/BENU-WAG-STATS-MIB.mib"
        },
        {
            "mib": "CISCOSB-IP-SLA",
            "path": "cisco/CISCOSB-IP-SLA.mib"
        },
        {
            "mib": "WESTERMO-FRNT-MIB",
            "path": "westermo/WESTERMO-FRNT-MIB.mib"
        },
        {
            "mib": "SENSATRONICS-EM1",
            "path": "sensatronics/SENSATRONICS-EM1.mib"
        },
        {
            "mib": "AcAtm",
            "path": "audiocodes/AcAtm.mib"
        },
        {
            "mib": "RAY-MIB",
            "path": "ray/RAY-MIB.mib"
        },
        {
            "mib": "CMM4-MIB",
            "path": "cambium/cmm4/CMM4-MIB.mib"
        },
        {
            "mib": "RADLAN-MNGINF-MIB",
            "path": "radlan/RADLAN-MNGINF-MIB.mib"
        },
        {
            "mib": "ExaltComProducts",
            "path": "exalt/ExaltComProducts.mib"
        },
        {
            "mib": "RAISECOM-RIP-MIB",
            "path": "raisecom/RAISECOM-RIP-MIB.mib"
        },
        {
            "mib": "Juniper-PPPOE-MIB",
            "path": "junose/Juniper-PPPOE-MIB.mib"
        },
        {
            "mib": "TELDAT-MON-INTERF-CELLULAR-MIB",
            "path": "teldat/TELDAT-MON-INTERF-CELLULAR-MIB.mib"
        },
        {
            "mib": "Deltanet-MIB",
            "path": "deltanet/Deltanet-MIB.mib"
        },
        {
            "mib": "JUNIPER-RMON-MIB",
            "path": "junos/JUNIPER-RMON-MIB.mib"
        },
        {
            "mib": "LINKSYS-rlFft",
            "path": "linksys/LINKSYS-rlFft.mib"
        },
        {
            "mib": "RADLAN-PHY-MIB",
            "path": "radlan/RADLAN-PHY-MIB.mib"
        },
        {
            "mib": "SENSATRONICS-ITTM",
            "path": "sensatronics/SENSATRONICS-ITTM.mib"
        },
        {
            "mib": "HH3C-QINQV2-MIB",
            "path": "comware/HH3C-QINQV2-MIB.mib"
        },
        {
            "mib": "RAY3-MIB",
            "path": "ray/RAY3-MIB.mib"
        },
        {
            "mib": "Juniper-PPPoE-Profile-CONF",
            "path": "junose/Juniper-PPPoE-Profile-CONF.mib"
        },
        {
            "mib": "AcBoard",
            "path": "audiocodes/AcBoard.mib"
        },
        {
            "mib": "RAISECOM-RIP2-MIB",
            "path": "raisecom/RAISECOM-RIP2-MIB.mib"
        },
        {
            "mib": "SENSATRONICS-SMI",
            "path": "sensatronics/SENSATRONICS-SMI.mib"
        },
        {
            "mib": "WESTERMO-INTERFACE-MIB",
            "path": "westermo/WESTERMO-INTERFACE-MIB.mib"
        },
        {
            "mib": "INOVONICS-MODEL640-MIB",
            "path": "inovonics/INOVONICS-MODEL640-MIB.mib"
        },
        {
            "mib": "CISCOSB-IP",
            "path": "cisco/CISCOSB-IP.mib"
        },
        {
            "mib": "Juniper-ROUTER-MIB",
            "path": "junos/Juniper-ROUTER-MIB.mib"
        },
        {
            "mib": "HUAWEI-SWITCH-L2MAM-EXT-MIB",
            "path": "huawei/HUAWEI-SWITCH-L2MAM-EXT-MIB.mib"
        },
        {
            "mib": "RADLAN-Physicaldescription-MIB",
            "path": "radlan/RADLAN-Physicaldescription-MIB.mib"
        },
        {
            "mib": "CORERO-CMS-MIB",
            "path": "corero/CORERO-CMS-MIB.mib"
        },
        {
            "mib": "LINKSYS-rlInterfaces",
            "path": "linksys/LINKSYS-rlInterfaces.mib"
        },
        {
            "mib": "ExtendAirG2",
            "path": "exalt/ExtendAirG2.mib"
        },
        {
            "mib": "WESTERMO-OID-MIB",
            "path": "westermo/WESTERMO-OID-MIB.mib"
        },
        {
            "mib": "ATEN-IPMI-MIB",
            "path": "supermicro/ATEN-IPMI-MIB.mib"
        },
        {
            "mib": "CISCOSB-ippreflist-MIB",
            "path": "cisco/CISCOSB-ippreflist-MIB.mib"
        },
        {
            "mib": "ESI-MIB",
            "path": "zebra/ESI-MIB.mib"
        },
        {
            "mib": "HH3C-QOS-CAPABILITY-MIB",
            "path": "comware/HH3C-QOS-CAPABILITY-MIB.mib"
        },
        {
            "mib": "RAISECOM-RNDP-MIB",
            "path": "raisecom/RAISECOM-RNDP-MIB.mib"
        },
        {
            "mib": "Juniper-PPPOE-PROFILE-MIB",
            "path": "junose/Juniper-PPPOE-PROFILE-MIB.mib"
        },
        {
            "mib": "JUNIPER-RPF-MIB",
            "path": "junos/JUNIPER-RPF-MIB.mib"
        },
        {
            "mib": "HUAWEI-SYS-CLOCK-MIB",
            "path": "huawei/HUAWEI-SYS-CLOCK-MIB.mib"
        },
        {
            "mib": "RADLAN-QOS-CLI-MIB",
            "path": "radlan/RADLAN-QOS-CLI-MIB.mib"
        },
        {
            "mib": "TELDAT-MON-INTERF-WLAN-MIB",
            "path": "teldat/TELDAT-MON-INTERF-WLAN-MIB.mib"
        },
        {
            "mib": "CISCO-UNIFIED-COMPUTING-TC-MIB",
            "path": "cisco/CISCO-UNIFIED-COMPUTING-TC-MIB.mib"
        },
        {
            "mib": "QOS",
            "path": "exalt/QOS.mib"
        },
        {
            "mib": "CORERO-CMS-SEGMENTS-MIB",
            "path": "corero/CORERO-CMS-SEGMENTS-MIB.mib"
        },
        {
            "mib": "ADSL-LINE-EXT-MIB",
            "path": "ADSL-LINE-EXT-MIB.mib"
        },
        {
            "mib": "CISCOSB-IpRouter",
            "path": "cisco/CISCOSB-IpRouter.mib"
        },
        {
            "mib": "HH3C-QOS-PROFILE-MIB",
            "path": "comware/HH3C-QOS-PROFILE-MIB.mib"
        },
        {
            "mib": "Juniper-Products-MIB",
            "path": "junose/Juniper-Products-MIB.mib"
        },
        {
            "mib": "RAISECOM-ROUTEMANAGE-MIB",
            "path": "raisecom/RAISECOM-ROUTEMANAGE-MIB.mib"
        },
        {
            "mib": "AcGateway",
            "path": "audiocodes/AcGateway.mib"
        },
        {
            "mib": "RADLAN-QOS-SERV",
            "path": "radlan/RADLAN-QOS-SERV.mib"
        },
        {
            "mib": "INOVONICS-MODEL650-MIB",
            "path": "inovonics/INOVONICS-MODEL650-MIB.mib"
        },
        {
            "mib": "ZEBRA-QL-MIB",
            "path": "zebra/ZEBRA-QL-MIB.mib"
        },
        {
            "mib": "HUAWEI-SYS-MAN-MIB",
            "path": "huawei/HUAWEI-SYS-MAN-MIB.mib"
        },
        {
            "mib": "LAMBDATRAIL-MIB",
            "path": "deltanet/LAMBDATRAIL-MIB.mib"
        },
        {
            "mib": "TELDAT-SW-STRUCTURE-MIB",
            "path": "teldat/TELDAT-SW-STRUCTURE-MIB.mib"
        },
        {
            "mib": "SUPERMICRO-HEALTH-MIB",
            "path": "supermicro/SUPERMICRO-HEALTH-MIB.mib"
        },
        {
            "mib": "JUNIPER-RPM-MIB",
            "path": "junos/JUNIPER-RPM-MIB.mib"
        },
        {
            "mib": "SYSLOG",
            "path": "exalt/SYSLOG.mib"
        },
        {
            "mib": "LINKSYS-rlLcli-MIB",
            "path": "linksys/LINKSYS-rlLcli-MIB.mib"
        },
        {
            "mib": "RAISECOM-ROUTEPOLICY-MIB",
            "path": "raisecom/RAISECOM-ROUTEPOLICY-MIB.mib"
        },
        {
            "mib": "CORERO-CMS-STATISTICS-MIB",
            "path": "corero/CORERO-CMS-STATISTICS-MIB.mib"
        },
        {
            "mib": "AclV5",
            "path": "audiocodes/AclV5.mib"
        },
        {
            "mib": "CISCOSB-IPSTDACL-MIB",
            "path": "cisco/CISCOSB-IPSTDACL-MIB.mib"
        },
        {
            "mib": "RADLAN-rlInterfaces",
            "path": "radlan/RADLAN-rlInterfaces.mib"
        },
        {
            "mib": "WESTERMO-WEOS-MIB",
            "path": "westermo/WESTERMO-WEOS-MIB.mib"
        },
        {
            "mib": "Juniper-Profile-CONF",
            "path": "junose/Juniper-Profile-CONF.mib"
        },
        {
            "mib": "LIEBERT-GP-AGENT-MIB",
            "path": "liebert/LIEBERT-GP-AGENT-MIB.mib"
        },
        {
            "mib": "VLAN",
            "path": "exalt/VLAN.mib"
        },
        {
            "mib": "LINKSYS-RMON",
            "path": "linksys/LINKSYS-RMON.mib"
        },
        {
            "mib": "HH3C-RADIUS-MIB",
            "path": "comware/HH3C-RADIUS-MIB.mib"
        },
        {
            "mib": "RAISECOM-RRCP-MIB",
            "path": "raisecom/RAISECOM-RRCP-MIB.mib"
        },
        {
            "mib": "ADSL-LINE-MIB",
            "path": "ADSL-LINE-MIB.mib"
        },
        {
            "mib": "WESTERMO-WEOS-TECHPREVIEW-MIB",
            "path": "westermo/WESTERMO-WEOS-TECHPREVIEW-MIB.mib"
        },
        {
            "mib": "PanDacom-MIB",
            "path": "pandacom/PanDacom-MIB.mib"
        },
        {
            "mib": "Lambdatrail2s-MIB",
            "path": "deltanet/Lambdatrail2s-MIB.mib"
        },
        {
            "mib": "SUPERMICRO-SD5-MIB",
            "path": "supermicro/SUPERMICRO-SD5-MIB.mib"
        },
        {
            "mib": "HUAWEI-SYSLOG-MIB",
            "path": "huawei/HUAWEI-SYSLOG-MIB.mib"
        },
        {
            "mib": "ADSL-TC-MIB",
            "path": "ADSL-TC-MIB.mib"
        },
        {
            "mib": "HH3C-RAID-MIB",
            "path": "comware/HH3C-RAID-MIB.mib"
        },
        {
            "mib": "CORERO-CMS-SYSTEM-STATUS-MIB",
            "path": "corero/CORERO-CMS-SYSTEM-STATUS-MIB.mib"
        },
        {
            "mib": "Juniper-PROFILE-MIB",
            "path": "junose/Juniper-PROFILE-MIB.mib"
        },
        {
            "mib": "SUPERMICRO-SMI",
            "path": "supermicro/SUPERMICRO-SMI.mib"
        },
        {
            "mib": "RAISECOM-RRCP-VLAN-MIB",
            "path": "raisecom/RAISECOM-RRCP-VLAN-MIB.mib"
        },
        {
            "mib": "AcPerfH323SIPGateway",
            "path": "audiocodes/AcPerfH323SIPGateway.mib"
        },
        {
            "mib": "EPPC-MIB",
            "path": "powerwalker/EPPC-MIB.mib"
        },
        {
            "mib": "XIRRUS-MIB",
            "path": "xirrus_aos/XIRRUS-MIB.mib"
        },
        {
            "mib": "RADLAN-rlLcli-MIB",
            "path": "radlan/RADLAN-rlLcli-MIB.mib"
        },
        {
            "mib": "MDS-EVENT-MIB",
            "path": "gemds/MDS-EVENT-MIB.mib"
        },
        {
            "mib": "JUNIPER-RPS-MIB",
            "path": "junos/JUNIPER-RPS-MIB.mib"
        },
        {
            "mib": "SPEED-DUALLINE-FC",
            "path": "pandacom/SPEED-DUALLINE-FC.mib"
        },
        {
            "mib": "CISCOSB-IPv6",
            "path": "cisco/CISCOSB-IPv6.mib"
        },
        {
            "mib": "HUAWEI-SZONE-MIB",
            "path": "huawei/HUAWEI-SZONE-MIB.mib"
        },
        {
            "mib": "SILVERPEAK-MGMT-MIB",
            "path": "silverpeak/SILVERPEAK-MGMT-MIB.mib"
        },
        {
            "mib": "CORERO-MIB",
            "path": "corero/CORERO-MIB.mib"
        },
        {
            "mib": "FORCEPOINT-NGFW-ENGINE-MIB",
            "path": "forcepoint/FORCEPOINT-NGFW-ENGINE-MIB.mib"
        },
        {
            "mib": "HH3C-RBAC-MIB",
            "path": "comware/HH3C-RBAC-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-CONDITIONS-MIB",
            "path": "liebert/LIEBERT-GP-COND-MIB.mib"
        },
        {
            "mib": "MDS-IF-CELL-MIB",
            "path": "gemds/MDS-IF-CELL-MIB.mib"
        },
        {
            "mib": "PULSESECURE-PSG-MIB",
            "path": "pulse/PULSESECURE-PSG-MIB.mib"
        },
        {
            "mib": "LINKSYS-rndApplications",
            "path": "linksys/LINKSYS-rndApplications.mib"
        },
        {
            "mib": "Juniper-QoS-Manager-CONF",
            "path": "junose/Juniper-QoS-Manager-CONF.mib"
        },
        {
            "mib": "FORTINET-CORE-MIB",
            "path": "fortinet/FORTINET-CORE-MIB.mib"
        },
        {
            "mib": "RADLAN-RMON",
            "path": "radlan/RADLAN-RMON.mib"
        },
        {
            "mib": "JUNIPER-RSVP-MIB",
            "path": "junos/JUNIPER-RSVP-MIB.mib"
        },
        {
            "mib": "AcPerfMediaGateway",
            "path": "audiocodes/AcPerfMediaGateway.mib"
        },
        {
            "mib": "ADSL2-LINE-TC-MIB",
            "path": "ADSL2-LINE-TC-MIB.mib"
        },
        {
            "mib": "RAISECOM-RTDP-MIB",
            "path": "raisecom/RAISECOM-RTDP-MIB.mib"
        },
        {
            "mib": "MDS-IF-IEEE80211-MIB",
            "path": "gemds/MDS-IF-IEEE80211-MIB.mib"
        },
        {
            "mib": "STONESOFT-FIREWALL-MIB",
            "path": "forcepoint/STONESOFT-FIREWALL-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-CONTROLLER-MIB",
            "path": "liebert/LIEBERT-GP-CONTROLLER-MIB.mib"
        },
        {
            "mib": "RADLAN-rndApplications",
            "path": "radlan/RADLAN-rndApplications.mib"
        },
        {
            "mib": "HUAWEI-TAD-MIB",
            "path": "huawei/HUAWEI-TAD-MIB.mib"
        },
        {
            "mib": "SPEED-MULTILINE-MIB",
            "path": "pandacom/SPEED-MULTILINE-MIB.mib"
        },
        {
            "mib": "AcPerfMediaServices",
            "path": "audiocodes/AcPerfMediaServices.mib"
        },
        {
            "mib": "JUNIPER-RTM-MIB",
            "path": "junos/JUNIPER-RTM-MIB.mib"
        },
        {
            "mib": "MDS-IF-LN-MIB",
            "path": "gemds/MDS-IF-LN-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIADC-MIB",
            "path": "fortinet/FORTINET-FORTIADC-MIB.mib"
        },
        {
            "mib": "RAISECOM-SCHEDULE-MIB",
            "path": "raisecom/RAISECOM-SCHEDULE-MIB.mib"
        },
        {
            "mib": "HH3C-RBM-MIB",
            "path": "comware/HH3C-RBM-MIB.mib"
        },
        {
            "mib": "STONESOFT-NETNODE-MIB",
            "path": "forcepoint/STONESOFT-NETNODE-MIB.mib"
        },
        {
            "mib": "CISCOSB-IPV6FHS-MIB",
            "path": "cisco/CISCOSB-IPV6FHS-MIB.mib"
        },
        {
            "mib": "RADLAN-rndMng",
            "path": "radlan/RADLAN-rndMng.mib"
        },
        {
            "mib": "LINKSYS-rndMng",
            "path": "linksys/LINKSYS-rndMng.mib"
        },
        {
            "mib": "MDS-IF-LW-MIB",
            "path": "gemds/MDS-IF-LW-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-ENVIRONMENTAL-MIB",
            "path": "liebert/LIEBERT-GP-ENV-MIB.mib"
        },
        {
            "mib": "STONESOFT-SMI-MIB",
            "path": "forcepoint/STONESOFT-SMI-MIB.mib"
        },
        {
            "mib": "ALARM-MIB",
            "path": "ALARM-MIB.mib"
        },
        {
            "mib": "HUAWEI-TASK-MIB",
            "path": "huawei/HUAWEI-TASK-MIB.mib"
        },
        {
            "mib": "Juniper-QoS-MIB",
            "path": "junose/Juniper-QoS-MIB.mib"
        },
        {
            "mib": "JUNIPER-SCU-MIB",
            "path": "junos/JUNIPER-SCU-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIAP-MIB",
            "path": "fortinet/FORTINET-FORTIAP-MIB.mib"
        },
        {
            "mib": "SILVERPEAK-PRODUCTS-MIB",
            "path": "silverpeak/SILVERPEAK-PRODUCTS-MIB.mib"
        },
        {
            "mib": "CISCOSB-JUMBOFRAMES-MIB",
            "path": "cisco/CISCOSB-JUMBOFRAMES-MIB.mib"
        },
        {
            "mib": "SPEED-MUX-200G-MIB",
            "path": "pandacom/SPEED-MUX-200G-MIB.mib"
        },
        {
            "mib": "MDS-IF-NX-MIB",
            "path": "gemds/MDS-IF-NX-MIB.mib"
        },
        {
            "mib": "RADLAN-SENSORENTMIB",
            "path": "radlan/RADLAN-SENSORENTMIB.mib"
        },
        {
            "mib": "HH3C-RCP-MIB",
            "path": "comware/HH3C-RCP-MIB.mib"
        },
        {
            "mib": "RAISECOM-SLA-MIB",
            "path": "raisecom/RAISECOM-SLA-MIB.mib"
        },
        {
            "mib": "LINKSYS-SCT-MIB",
            "path": "linksys/LINKSYS-SCT-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIAUTHENTICATOR-MIB",
            "path": "fortinet/FORTINET-FORTIAUTHENTICATOR-MIB.mib"
        },
        {
            "mib": "SILVERPEAK-SMI",
            "path": "silverpeak/SILVERPEAK-SMI.mib"
        },
        {
            "mib": "CISCOSB-LBD-MIB",
            "path": "cisco/CISCOSB-LBD-MIB.mib"
        },
        {
            "mib": "RADLAN-SNMP-MIB",
            "path": "radlan/RADLAN-SNMP-MIB.mib"
        },
        {
            "mib": "MDS-ORBIT-SMI-MIB",
            "path": "gemds/MDS-ORBIT-SMI-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-FLEXIBLE-CONDITIONS-MIB",
            "path": "liebert/LIEBERT-GP-FLEXIBLE-COND-MIB.mib"
        },
        {
            "mib": "RAISECOM-SROUTE-MIB",
            "path": "raisecom/RAISECOM-SROUTE-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-CLIENT-CONF",
            "path": "junose/Juniper-RADIUS-CLIENT-CONF.mib"
        },
        {
            "mib": "BGP4-MIB",
            "path": "BGP4-MIB.mib"
        },
        {
            "mib": "CISCOSB-LLDP-MIB",
            "path": "cisco/CISCOSB-LLDP-MIB.mib"
        },
        {
            "mib": "LINKSYS-SECSD-MIB",
            "path": "linksys/LINKSYS-SECSD-MIB.mib"
        },
        {
            "mib": "SILVERPEAK-TC",
            "path": "silverpeak/SILVERPEAK-TC.mib"
        },
        {
            "mib": "HH3C-RCR-MIB",
            "path": "comware/HH3C-RCR-MIB.mib"
        },
        {
            "mib": "SPEED-AMP-MIB",
            "path": "pandacom/SPEEDAMP-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-FLEXIBLE-MIB",
            "path": "liebert/LIEBERT-GP-FLEXIBLE-MIB.mib"
        },
        {
            "mib": "MDS-REG-MIB",
            "path": "gemds/MDS-REG-MIB.mib"
        },
        {
            "mib": "JUNIPER-SECURE-ACCESS-PORT-MIB",
            "path": "junos/JUNIPER-SECURE-ACCESS-PORT-MIB.mib"
        },
        {
            "mib": "BGP4V2-TC-MIB",
            "path": "BGP4V2-TC-MIB.mib"
        },
        {
            "mib": "RAISECOM-SSH-MIB",
            "path": "raisecom/RAISECOM-SSH-MIB.mib"
        },
        {
            "mib": "CISCOSB-LOCALIZATION-MIB",
            "path": "cisco/CISCOSB-LOCALIZATION-MIB.mib"
        },
        {
            "mib": "MDS-SERIAL-MIB",
            "path": "gemds/MDS-SERIAL-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIEXTENDER-MIB",
            "path": "fortinet/FORTINET-FORTIEXTENDER-MIB.mib"
        },
        {
            "mib": "LINKSYS-SECURITY-SUITE",
            "path": "linksys/LINKSYS-SECURITY-SUITE.mib"
        },
        {
            "mib": "HH3C-RDDC-MIB",
            "path": "comware/HH3C-RDDC-MIB.mib"
        },
        {
            "mib": "RADLAN-SNMPv2",
            "path": "radlan/RADLAN-SNMPv2.mib"
        },
        {
            "mib": "LIEBERT-GP-NOTIFICATIONS-MIB",
            "path": "liebert/LIEBERT-GP-NOTIFY-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-CLIENT-MIB",
            "path": "junose/Juniper-RADIUS-CLIENT-MIB.mib"
        },
        {
            "mib": "CISCOSB-MAC-BASE-PRIO",
            "path": "cisco/CISCOSB-MAC-BASE-PRIO.mib"
        },
        {
            "mib": "MDS-SERVICE-GPS-MIB",
            "path": "gemds/MDS-SERVICE-GPS-MIB.mib"
        },
        {
            "mib": "RAISECOM-SYSLOG-SERVICE-MIB",
            "path": "raisecom/RAISECOM-SYSLOG-SERVICE-MIB.mib"
        },
        {
            "mib": "SPEEDCARRIER-MIB",
            "path": "pandacom/SPEEDCARRIER-MIB.mib"
        },
        {
            "mib": "BRIDGE-MIB",
            "path": "BRIDGE-MIB.mib"
        },
        {
            "mib": "LINKSYS-SMON-MIB",
            "path": "linksys/LINKSYS-SMON-MIB.mib"
        },
        {
            "mib": "JUNIPER-Services-MIB",
            "path": "junos/JUNIPER-Services-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-Disconnect-CONF",
            "path": "junose/Juniper-RADIUS-Disconnect-CONF.mib"
        },
        {
            "mib": "RADLAN-SSH-MIB",
            "path": "radlan/RADLAN-SSH-MIB.mib"
        },
        {
            "mib": "CISCOSB-MGMD-ROUTER-MIB",
            "path": "cisco/CISCOSB-MGMD-ROUTER-MIB.mib"
        },
        {
            "mib": "HH3C-RES-MON-MIB",
            "path": "comware/HH3C-RES-MON-MIB.mib"
        },
        {
            "mib": "MDS-SERVICES-MIB",
            "path": "gemds/MDS-SERVICES-MIB.mib"
        },
        {
            "mib": "SPEEDSINGLELINE-XFP-MIB",
            "path": "pandacom/SPEEDSINGLELINE-XFP-MIB.mib"
        },
        {
            "mib": "RAISECOM-SYSTEM-MIB",
            "path": "raisecom/RAISECOM-SYSTEM-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-PDU-MIB",
            "path": "liebert/LIEBERT-GP-PDU-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIGATE-MIB",
            "path": "fortinet/FORTINET-FORTIGATE-MIB.mib"
        },
        {
            "mib": "CAPWAP-BASE-MIB",
            "path": "CAPWAP-BASE-MIB-draft06.mib"
        },
        {
            "mib": "RADLAN-SSL",
            "path": "radlan/RADLAN-SSL.mib"
        },
        {
            "mib": "LINKSYS-SNMP-MIB",
            "path": "linksys/LINKSYS-SNMP-MIB.mib"
        },
        {
            "mib": "HUAWEI-TC-MIB",
            "path": "huawei/HUAWEI-TC-MIB.mib"
        },
        {
            "mib": "MDS-SYSTEM-MIB",
            "path": "gemds/MDS-SYSTEM-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-Disconnect-MIB",
            "path": "junose/Juniper-RADIUS-Disconnect-MIB.mib"
        },
        {
            "mib": "HH3C-RMON-EXT-MIB",
            "path": "comware/HH3C-RMON-EXT-MIB.mib"
        },
        {
            "mib": "CISCOSB-MIB",
            "path": "cisco/CISCOSB-MIB.mib"
        },
        {
            "mib": "JUNIPER-SIP-COMMON-MIB",
            "path": "junos/JUNIPER-SIP-COMMON-MIB.mib"
        },
        {
            "mib": "RAISECOM-UPGRADE-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-UPGRADE-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIMAIL-MIB",
            "path": "fortinet/FORTINET-FORTIMAIL-MIB.mib"
        },
        {
            "mib": "RADLAN-STACK-MIB",
            "path": "radlan/RADLAN-STACK-MIB.mib"
        },
        {
            "mib": "HUAWEI-TCP-MIB",
            "path": "huawei/HUAWEI-TCP-MIB.mib"
        },
        {
            "mib": "LINKSYS-SpecialBpdu-MIB",
            "path": "linksys/LINKSYS-SpecialBpdu-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-POWER-MIB",
            "path": "liebert/LIEBERT-GP-POWER-MIB.mib"
        },
        {
            "mib": "CISCOSB-MIR-MIB",
            "path": "cisco/CISCOSB-MIR-MIB.mib"
        },
        {
            "mib": "JUNIPER-SMI",
            "path": "junos/JUNIPER-SMI.mib"
        },
        {
            "mib": "RADLAN-SYSLOG-MIB",
            "path": "radlan/RADLAN-SYSLOG-MIB.mib"
        },
        {
            "mib": "DIAL-CONTROL-MIB",
            "path": "DIAL-CONTROL-MIB.mib"
        },
        {
            "mib": "HH3C-RMON-EXT2-MIB",
            "path": "comware/HH3C-RMON-EXT2-MIB.mib"
        },
        {
            "mib": "RAISECOM-USER-MANAGEMENT-MIB",
            "path": "raisecom/RAISECOM-USER-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIMANAGER-FORTIANALYZER-MIB",
            "path": "fortinet/FORTINET-FORTIMANAGER-FORTIANALYZER-MIB.mib"
        },
        {
            "mib": "CISCOSB-MNGINF-MIB",
            "path": "cisco/CISCOSB-MNGINF-MIB.mib"
        },
        {
            "mib": "HUAWEI-TDM-PSN-MIB",
            "path": "huawei/HUAWEI-TDM-PSN-MIB.mib"
        },
        {
            "mib": "JUNIPER-SNMP-SET-MIB",
            "path": "junos/JUNIPER-SNMP-SET-MIB.mib"
        },
        {
            "mib": "LINKSYS-SSL",
            "path": "linksys/LINKSYS-SSL.mib"
        },
        {
            "mib": "FORTINET-FORTISANDBOX-MIB",
            "path": "fortinet/FORTINET-FORTISANDBOX-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-Initiated-Request-MIB",
            "path": "junose/Juniper-RADIUS-Initiated-Request-MIB.mib"
        },
        {
            "mib": "HH3C-RPR-MIB",
            "path": "comware/HH3C-RPR-MIB.mib"
        },
        {
            "mib": "DIFFSERV-DSCP-TC",
            "path": "DIFFSERV-DSCP-TC.mib"
        },
        {
            "mib": "CISCOSB-MULTISESSIONTERMINAL-MIB",
            "path": "cisco/CISCOSB-MULTISESSIONTERMINAL-MIB.mib"
        },
        {
            "mib": "HUAWEI-TRILL-CONF-MIB",
            "path": "huawei/HUAWEI-TRILL-CONF-MIB.mib"
        },
        {
            "mib": "RADLAN-TIMESYNCHRONIZATION-MIB",
            "path": "radlan/RADLAN-TIMESYNCHRONIZATION-MIB.mib"
        },
        {
            "mib": "RAISECOM-VCT-MIB",
            "path": "raisecom/RAISECOM-VCT-MIB.mib"
        },
        {
            "mib": "CISCOSB-openflow-MIB",
            "path": "cisco/CISCOSB-openflow-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-REGISTRATION-MIB",
            "path": "liebert/LIEBERT-GP-REG-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTISWITCH-MIB",
            "path": "fortinet/FORTINET-FORTISWITCH-MIB.mib"
        },
        {
            "mib": "HH3C-RRPP-MIB",
            "path": "comware/HH3C-RRPP-MIB.mib"
        },
        {
            "mib": "LINKSYS-STACK-MIB",
            "path": "linksys/LINKSYS-STACK-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-Proxy-CONF",
            "path": "junose/Juniper-RADIUS-Proxy-CONF.mib"
        },
        {
            "mib": "HUAWEI-TRNG-MIB",
            "path": "huawei/HUAWEI-TRNG-MIB.mib"
        },
        {
            "mib": "RADLAN-Tuning",
            "path": "radlan/RADLAN-Tuning.mib"
        },
        {
            "mib": "JUNIPER-SOAM-PM-MIB",
            "path": "junos/JUNIPER-SOAM-PM-MIB.mib"
        },
        {
            "mib": "CISCOSB-PBR-MIB",
            "path": "cisco/CISCOSB-PBR-MIB.mib"
        },
        {
            "mib": "DIFFSERV-MIB",
            "path": "DIFFSERV-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIVOICE-MIB",
            "path": "fortinet/FORTINET-FORTIVOICE-MIB.mib"
        },
        {
            "mib": "RAISECOM-VLANGROUP-MIB",
            "path": "raisecom/RAISECOM-VLANGROUP-MIB.mib"
        },
        {
            "mib": "HH3C-RS485-MIB",
            "path": "comware/HH3C-RS485-MIB.mib"
        },
        {
            "mib": "JUNIPER-SONET-MIB",
            "path": "junos/JUNIPER-SONET-MIB.mib"
        },
        {
            "mib": "LINKSYS-STORMCTRL-MIB",
            "path": "linksys/LINKSYS-STORMCTRL-MIB.mib"
        },
        {
            "mib": "RADLAN-UDP",
            "path": "radlan/RADLAN-UDP.mib"
        },
        {
            "mib": "LIEBERT-GP-SRC-MIB",
            "path": "liebert/LIEBERT-GP-SRC-MIB.mib"
        },
        {
            "mib": "HUAWEI-TUNNEL-TE-MIB",
            "path": "huawei/HUAWEI-TUNNEL-TE-MIB.mib"
        },
        {
            "mib": "Juniper-RADIUS-Proxy-MIB",
            "path": "junose/Juniper-RADIUS-Proxy-MIB.mib"
        },
        {
            "mib": "RAISECOM-VLANMACCOPY-MIB",
            "path": "raisecom/RAISECOM-VLANMACCOPY-MIB.mib"
        },
        {
            "mib": "DISMAN-EVENT-MIB",
            "path": "DISMAN-EVENT-MIB.mib"
        },
        {
            "mib": "CISCOSB-PHY-MIB",
            "path": "cisco/CISCOSB-PHY-MIB.mib"
        },
        {
            "mib": "FORTINET-FORTIWEB-MIB",
            "path": "fortinet/FORTINET-FORTIWEB-MIB.mib"
        },
        {
            "mib": "JUNIPER-SP-MIB",
            "path": "junos/JUNIPER-SP-MIB.mib"
        },
        {
            "mib": "HH3C-RSA-MIB",
            "path": "comware/HH3C-RSA-MIB.mib"
        },
        {
            "mib": "LINKSYS-SYSLOG-MIB",
            "path": "linksys/LINKSYS-SYSLOG-MIB.mib"
        },
        {
            "mib": "RADLAN-vlan-MIB",
            "path": "radlan/RADLAN-vlan-MIB.mib"
        },
        {
            "mib": "Juniper-REDUNDANCY-MIB",
            "path": "junose/Juniper-REDUNDANCY-MIB.mib"
        },
        {
            "mib": "RAISECOM-VLANPROTECT-MIB",
            "path": "raisecom/RAISECOM-VLANPROTECT-MIB.mib"
        },
        {
            "mib": "HUAWEI-UNIMNG-MIB",
            "path": "huawei/HUAWEI-UNIMNG-MIB.mib"
        },
        {
            "mib": "JUNIPER-SRD-MIB",
            "path": "junos/JUNIPER-SRD-MIB.mib"
        },
        {
            "mib": "CISCOSB-Physicaldescription-MIB",
            "path": "cisco/CISCOSB-Physicaldescription-MIB.mib"
        },
        {
            "mib": "HH3C-SAN-AGG-MIB",
            "path": "comware/HH3C-SAN-AGG-MIB.mib"
        },
        {
            "mib": "LIEBERT-GP-SYSTEM-MIB",
            "path": "liebert/LIEBERT-GP-SYSTEM-MIB.mib"
        },
        {
            "mib": "FORTINET-MIB-280",
            "path": "fortinet/FORTINET-MIB-280.mib"
        },
        {
            "mib": "DISMAN-NSLOOKUP-MIB",
            "path": "DISMAN-NSLOOKUP-MIB.mib"
        },
        {
            "mib": "LINKSYS-SYSMNG-MIB",
            "path": "linksys/LINKSYS-SYSMNG-MIB.mib"
        },
        {
            "mib": "HUAWEI-USERLOG-MIB",
            "path": "huawei/HUAWEI-USERLOG-MIB.mib"
        },
        {
            "mib": "RAISECOM-VRRP-MIB",
            "path": "raisecom/RAISECOM-VRRP-MIB.mib"
        },
        {
            "mib": "CISCOSB-PIM-MIB",
            "path": "cisco/CISCOSB-PIM-MIB.mib"
        },
        {
            "mib": "Juniper-Registry",
            "path": "junose/Juniper-Registry.mib"
        },
        {
            "mib": "MERU-CONFIG-AP-MIB",
            "path": "fortinet/MERU-CONFIG-AP-MIB.mib"
        },
        {
            "mib": "HH3C-SAVA-MIB",
            "path": "comware/HH3C-SAVA-MIB.mib"
        },
        {
            "mib": "JUNIPER-SRX5000-SPU-MONITORING-MIB",
            "path": "junos/JUNIPER-SRX5000-SPU-MONITORING-MIB.mib"
        },
        {
            "mib": "LINKSYS-TELNET-MIB",
            "path": "linksys/LINKSYS-TELNET-MIB.mib"
        },
        {
            "mib": "RAISECOM-WEBSERVER-MIB",
            "path": "raisecom/RAISECOM-WEBSERVER-MIB.mib"
        },
        {
            "mib": "DISMAN-PING-MIB",
            "path": "DISMAN-PING-MIB.mib"
        },
        {
            "mib": "HUAWEI-VBST-MIB",
            "path": "huawei/HUAWEI-VBST-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-CONTROLLER-MIB",
            "path": "fortinet/MERU-CONFIG-CONTROLLER-MIB.mib"
        },
        {
            "mib": "HH3C-SECHIGH-MIB",
            "path": "comware/HH3C-SECHIGH-MIB.mib"
        },
        {
            "mib": "Juniper-RIP-CONF",
            "path": "junose/Juniper-RIP-CONF.mib"
        },
        {
            "mib": "CISCOSB-PNP",
            "path": "cisco/CISCOSB-PNP.mib"
        },
        {
            "mib": "JUNIPER-SUBSCRIBER-MIB",
            "path": "junos/JUNIPER-SUBSCRIBER-MIB.mib"
        },
        {
            "mib": "HUAWEI-VE-MIB",
            "path": "huawei/HUAWEI-VE-MIB.mib"
        },
        {
            "mib": "RAISECOM-WRED-MIB",
            "path": "raisecom/RAISECOM-WRED-MIB.mib"
        },
        {
            "mib": "LINKSYS-TIMESYNCHRONIZATION-MIB",
            "path": "linksys/LINKSYS-TIMESYNCHRONIZATION-MIB.mib"
        },
        {
            "mib": "HH3C-SECP-MIB",
            "path": "comware/HH3C-SECP-MIB.mib"
        },
        {
            "mib": "DISMAN-SCHEDULE-MIB",
            "path": "DISMAN-SCHEDULE-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-GRE-MIB",
            "path": "fortinet/MERU-CONFIG-GRE-MIB.mib"
        },
        {
            "mib": "Juniper-Router-CONF",
            "path": "junose/Juniper-Router-CONF.mib"
        },
        {
            "mib": "JUNIPER-SYSLOG-MIB",
            "path": "junos/JUNIPER-SYSLOG-MIB.mib"
        },
        {
            "mib": "HH3C-SESSION-MIB",
            "path": "comware/HH3C-SESSION-MIB.mib"
        },
        {
            "mib": "CISCOSB-POE-MIB",
            "path": "cisco/CISCOSB-POE-MIB.mib"
        },
        {
            "mib": "HUAWEI-VGMP-MIB",
            "path": "huawei/HUAWEI-VGMP-MIB.mib"
        },
        {
            "mib": "LINKSYS-TRACEROUTE-MIB",
            "path": "linksys/LINKSYS-TRACEROUTE-MIB.mib"
        },
        {
            "mib": "Juniper-TC",
            "path": "junos/Juniper-TC.mib"
        },
        {
            "mib": "MERU-CONFIG-ICR-MIB",
            "path": "fortinet/MERU-CONFIG-ICR-MIB.mib"
        },
        {
            "mib": "RC002-INTERVAL-PERFORMANCE-STAT-MIB",
            "path": "raisecom/RC002-INTERVAL-PERFORMANCE-STAT-MIB.mib"
        },
        {
            "mib": "DISMAN-SCRIPT-MIB",
            "path": "DISMAN-SCRIPT-MIB.mib"
        },
        {
            "mib": "CISCOSB-POLICY-MIB",
            "path": "cisco/CISCOSB-POLICY-MIB.mib"
        },
        {
            "mib": "Juniper-ROUTER-MIB",
            "path": "junose/Juniper-ROUTER-MIB.mib"
        },
        {
            "mib": "HH3C-SLBG-MIB",
            "path": "comware/HH3C-SLBG-MIB.mib"
        },
        {
            "mib": "LINKSYS-TRUNK-MIB",
            "path": "linksys/LINKSYS-TRUNK-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-LICENSE-MIB",
            "path": "fortinet/MERU-CONFIG-LICENSE-MIB.mib"
        },
        {
            "mib": "CISCOSB-PORT-STATISTICS-MIB",
            "path": "cisco/CISCOSB-PORT-STATISTICS-MIB.mib"
        },
        {
            "mib": "Juniper-SLEP-CONF",
            "path": "junose/Juniper-SLEP-CONF.mib"
        },
        {
            "mib": "JUNIPER-TIMING-NOTFNS-MIB",
            "path": "junos/JUNIPER-TIMING-NOTFNS-MIB.mib"
        },
        {
            "mib": "HH3C-SMLK-MIB",
            "path": "comware/HH3C-SMLK-MIB.mib"
        },
        {
            "mib": "RC002-LOCAL-DEVICE-PORT-MIB",
            "path": "raisecom/RC002-LOCAL-DEVICE-PORT-MIB.mib"
        },
        {
            "mib": "DISMAN-TRACEROUTE-MIB",
            "path": "DISMAN-TRACEROUTE-MIB.mib"
        },
        {
            "mib": "HUAWEI-VLL-STATISTIC-MIB",
            "path": "huawei/HUAWEI-VLL-STATISTIC-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-MACFILTERING-MIB",
            "path": "fortinet/MERU-CONFIG-MACFILTERING-MIB.mib"
        },
        {
            "mib": "CISCOSB-ProtectedPorts-MIB",
            "path": "cisco/CISCOSB-ProtectedPorts-MIB.mib"
        },
        {
            "mib": "LINKSYS-Tuning",
            "path": "linksys/LINKSYS-Tuning.mib"
        },
        {
            "mib": "JUNIPER-TLB-MIB",
            "path": "junos/JUNIPER-TLB-MIB.mib"
        },
        {
            "mib": "HUAWEI-VP-MIB",
            "path": "huawei/HUAWEI-VP-MIB.mib"
        },
        {
            "mib": "Juniper-SLEP-MIB",
            "path": "junose/Juniper-SLEP-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-PACKETCAPTURE-MIB",
            "path": "fortinet/MERU-CONFIG-PACKETCAPTURE-MIB.mib"
        },
        {
            "mib": "HH3C-SNA-DLSW-EXT-MIB",
            "path": "comware/HH3C-SNA-DLSW-EXT-MIB.mib"
        },
        {
            "mib": "RC002-REMOTE-DEVICE-MIB",
            "path": "raisecom/RC002-REMOTE-DEVICE-MIB.mib"
        },
        {
            "mib": "CISCOSB-QOS-APPS-MIB",
            "path": "cisco/CISCOSB-QOS-APPS-MIB.mib"
        },
        {
            "mib": "LINKSYS-TUNNEL-MIB",
            "path": "linksys/LINKSYS-TUNNEL-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-PORTPROFILE-MIB",
            "path": "fortinet/MERU-CONFIG-PORTPROFILE-MIB.mib"
        },
        {
            "mib": "DLSW-MIB",
            "path": "DLSW-MIB.mib"
        },
        {
            "mib": "HH3C-SNA-DLSW-MIB",
            "path": "comware/HH3C-SNA-DLSW-MIB.mib"
        },
        {
            "mib": "RC002-REMOTEII-DEVICE-MIB",
            "path": "raisecom/RC002-REMOTEII-DEVICE-MIB.mib"
        },
        {
            "mib": "JUNIPER-TRACEROUTE-MIB",
            "path": "junos/JUNIPER-TRACEROUTE-MIB.mib"
        },
        {
            "mib": "HUAWEI-VPLS-EXT-MIB",
            "path": "huawei/HUAWEI-VPLS-EXT-MIB.mib"
        },
        {
            "mib": "JUNIPER-SMI",
            "path": "junose/JUNIPER-SMI.mib"
        },
        {
            "mib": "CISCOSB-QOS-CLI-MIB",
            "path": "cisco/CISCOSB-QOS-CLI-MIB.mib"
        },
        {
            "mib": "LINKSYS-UDP",
            "path": "linksys/LINKSYS-UDP.mib"
        },
        {
            "mib": "DNS-RESOLVER-MIB",
            "path": "DNS-RESOLVER-MIB.mib"
        },
        {
            "mib": "HH3C-SNMP-EXT-MIB",
            "path": "comware/HH3C-SNMP-EXT-MIB.mib"
        },
        {
            "mib": "CISCOSB-QUEUE-STATISTICS-MIB",
            "path": "cisco/CISCOSB-QUEUE-STATISTICS-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-QOS-MIB",
            "path": "fortinet/MERU-CONFIG-QOS-MIB.mib"
        },
        {
            "mib": "ROSMGMT-ALARM-MGMT-MIB",
            "path": "raisecom/ROSMGMT-ALARM-MGMT-MIB.mib"
        },
        {
            "mib": "HUAWEI-VPLS-MIB",
            "path": "huawei/HUAWEI-VPLS-MIB.mib"
        },
        {
            "mib": "LINKSYS-vlan-MIB",
            "path": "linksys/LINKSYS-vlan-MIB.mib"
        },
        {
            "mib": "HH3C-SPB-MIB",
            "path": "comware/HH3C-SPB-MIB.mib"
        },
        {
            "mib": "JUNIPER-TUNNEL-STATS-MIB",
            "path": "junos/JUNIPER-TUNNEL-STATS-MIB.mib"
        },
        {
            "mib": "Juniper-SNMP-CONF",
            "path": "junose/Juniper-SNMP-CONF.mib"
        },
        {
            "mib": "DNS-SERVER-MIB",
            "path": "DNS-SERVER-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-ROGUEAP-MIB",
            "path": "fortinet/MERU-CONFIG-ROGUEAP-MIB.mib"
        },
        {
            "mib": "CISCOSB-RADIUSSRV",
            "path": "cisco/CISCOSB-RADIUSSRV.mib"
        },
        {
            "mib": "ROSMGMT-COMMON-MANAGEMENT-MIB",
            "path": "raisecom/ROSMGMT-COMMON-MANAGEMENT-MIB.mib"
        },
        {
            "mib": "HH3C-LswARP-MIB",
            "path": "comware/HH3C-SPLAT-ARP-MIB.mib"
        },
        {
            "mib": "LINKSYS-vlanVoice-MIB",
            "path": "linksys/LINKSYS-vlanVoice-MIB.mib"
        },
        {
            "mib": "HUAWEI-VPLS-TNL-MIB",
            "path": "huawei/HUAWEI-VPLS-TNL-MIB.mib"
        },
        {
            "mib": "JUNIPER-TWAMP-MIB",
            "path": "junos/JUNIPER-TWAMP-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-SECURITY-MIB",
            "path": "fortinet/MERU-CONFIG-SECURITY-MIB.mib"
        },
        {
            "mib": "Juniper-SNMP-MIB",
            "path": "junose/Juniper-SNMP-MIB.mib"
        },
        {
            "mib": "CISCOSB-Redistribute",
            "path": "cisco/CISCOSB-Redistribute.mib"
        },
        {
            "mib": "DOCS-CABLE-DEVICE-MIB",
            "path": "DOCS-CABLE-DEVICE-MIB.mib"
        },
        {
            "mib": "ROSMGMT-MEMORY-MIB",
            "path": "raisecom/ROSMGMT-MEMORY-MIB.mib"
        },
        {
            "mib": "HH3C-LswDHCP-MIB",
            "path": "comware/HH3C-SPLAT-DHCP-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-SECURITYCERT-MIB",
            "path": "fortinet/MERU-CONFIG-SECURITYCERT-MIB.mib"
        },
        {
            "mib": "LINKSYS-WBA-MIB",
            "path": "linksys/LINKSYS-WBA-MIB.mib"
        },
        {
            "mib": "CISCOSB-rlBrgMcMngr-MIB",
            "path": "cisco/CISCOSB-rlBrgMcMngr-MIB.mib"
        },
        {
            "mib": "Juniper-UNI-SMI",
            "path": "junos/Juniper-UNI-SMI.mib"
        },
        {
            "mib": "ROSMGMT-OPTICAL-TRANSCEIVER-MIB",
            "path": "raisecom/ROSMGMT-OPTICAL-TRANSCEIVER-MIB.mib"
        },
        {
            "mib": "Juniper-SONET-CONF",
            "path": "junose/Juniper-SONET-CONF.mib"
        },
        {
            "mib": "MERU-CONFIG-SNMP-MIB",
            "path": "fortinet/MERU-CONFIG-SNMP-MIB.mib"
        },
        {
            "mib": "HUAWEI-VPN-DIAGNOSTICS-MIB",
            "path": "huawei/HUAWEI-VPN-DIAGNOSTICS-MIB.mib"
        },
        {
            "mib": "LINKSYS-WeightedRandomTailDrop-MIB",
            "path": "linksys/LINKSYS-WeightedRandomTailDrop-MIB.mib"
        },
        {
            "mib": "HH3C-LswIGSP-MIB",
            "path": "comware/HH3C-SPLAT-IGSP-MIB.mib"
        },
        {
            "mib": "JUNIPER-URL-FILTER-MIB",
            "path": "junos/JUNIPER-URL-FILTER-MIB.mib"
        },
        {
            "mib": "CISCOSB-rlBrgMulticast-MIB",
            "path": "cisco/CISCOSB-rlBrgMulticast-MIB.mib"
        },
        {
            "mib": "Juniper-SSC-Client-CONF",
            "path": "junose/Juniper-SSC-Client-CONF.mib"
        },
        {
            "mib": "DOCS-IF-MIB",
            "path": "DOCS-IF-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-STATICSTATION-MIB",
            "path": "fortinet/MERU-CONFIG-STATICSTATION-MIB.mib"
        },
        {
            "mib": "CISCOSB-rldot1q-MIB",
            "path": "cisco/CISCOSB-rldot1q-MIB.mib"
        },
        {
            "mib": "HUAWEI-VRRP-EXT-MIB",
            "path": "huawei/HUAWEI-VRRP-EXT-MIB.mib"
        },
        {
            "mib": "ROSMGMT-OSPFV2-MIB",
            "path": "raisecom/ROSMGMT-OSPFv2-MIB.mib"
        },
        {
            "mib": "HUAWEI-VS-MIB",
            "path": "huawei/HUAWEI-VS-MIB.mib"
        },
        {
            "mib": "HH3C-LswINF-MIB",
            "path": "comware/HH3C-SPLAT-INF-MIB.mib"
        },
        {
            "mib": "JUNIPER-USER-AAA-MIB",
            "path": "junos/JUNIPER-USER-AAA-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-STATION-MIB",
            "path": "fortinet/MERU-CONFIG-STATION-MIB.mib"
        },
        {
            "mib": "Juniper-SSC-CLIENT-MIB",
            "path": "junose/Juniper-SSC-CLIENT-MIB.mib"
        },
        {
            "mib": "DOT3-OAM-MIB",
            "path": "DOT3-OAM-MIB.mib"
        },
        {
            "mib": "ROSMGMT-SYSTEM-MIB",
            "path": "raisecom/ROSMGMT-SYSTEM-MIB.mib"
        },
        {
            "mib": "HH3C-LswMAM-MIB",
            "path": "comware/HH3C-SPLAT-MAM-MIB.mib"
        },
        {
            "mib": "CISCOSB-rlFft",
            "path": "cisco/CISCOSB-rlFft.mib"
        },
        {
            "mib": "JUNIPER-USERFIREWALL-MIB",
            "path": "junos/JUNIPER-USERFIREWALL-MIB.mib"
        },
        {
            "mib": "HUAWEI-WAN-MIB",
            "path": "huawei/HUAWEI-WAN-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-VLAN-MIB",
            "path": "fortinet/MERU-CONFIG-VLAN-MIB.mib"
        },
        {
            "mib": "Juniper-Subscriber-CONF",
            "path": "junose/Juniper-Subscriber-CONF.mib"
        },
        {
            "mib": "ROSMGMT-VERSION-MIB",
            "path": "raisecom/ROSMGMT-VERSION-MIB.mib"
        },
        {
            "mib": "HH3C-LswMix-MIB",
            "path": "comware/HH3C-SPLAT-MIX-MIB.mib"
        },
        {
            "mib": "CISCOSB-rlInterfaces",
            "path": "cisco/CISCOSB-rlInterfaces.mib"
        },
        {
            "mib": "DS1-MIB",
            "path": "DS1-MIB.mib"
        },
        {
            "mib": "JUNIPER-UTIL-MIB",
            "path": "junos/JUNIPER-UTIL-MIB.mib"
        },
        {
            "mib": "HUAWEI-WARRANTY-MIB",
            "path": "huawei/HUAWEI-WARRANTY-MIB.mib"
        },
        {
            "mib": "Juniper-SUBSCRIBER-MIB",
            "path": "junose/Juniper-SUBSCRIBER-MIB.mib"
        },
        {
            "mib": "MERU-CONFIG-WLAN-MIB",
            "path": "fortinet/MERU-CONFIG-WLAN-MIB.mib"
        },
        {
            "mib": "CISCOSB-RLINVENTORYENT-MIB",
            "path": "cisco/CISCOSB-RLINVENTORYENT-MIB.mib"
        },
        {
            "mib": "SWITCH-AUTO-CONFIGURATION-MIB",
            "path": "raisecom/SWITCH-AUTO-CONFIGURATION-MIB.mib"
        },
        {
            "mib": "HH3C-LswMSTP-MIB",
            "path": "comware/HH3C-SPLAT-MSTP-MIB.mib"
        },
        {
            "mib": "DS3-MIB",
            "path": "DS3-MIB.mib"
        },
        {
            "mib": "Juniper-System-Clock-CONF",
            "path": "junose/Juniper-System-Clock-CONF.mib"
        },
        {
            "mib": "JUNIPER-VIRTUALCHASSIS-MIB",
            "path": "junos/JUNIPER-VIRTUALCHASSIS-MIB.mib"
        },
        {
            "mib": "SWITCH-CCP-MIB",
            "path": "raisecom/SWITCH-CCP-MIB.mib"
        },
        {
            "mib": "CISCOSB-rlIP-MIB",
            "path": "cisco/CISCOSB-rlIP-MIB.mib"
        },
        {
            "mib": "MERU-GLOBAL-STATISTICS-MIB",
            "path": "fortinet/MERU-GLOBAL-STATISTICS-MIB.mib"
        },
        {
            "mib": "HH3C-LswQos-MIB",
            "path": "comware/HH3C-SPLAT-QOS-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-AP-MIB",
            "path": "huawei/HUAWEI-WLAN-AP-MIB.mib"
        },
        {
            "mib": "DVMRP-MIB",
            "path": "DVMRP-MIB.mib"
        },
        {
            "mib": "SWITCH-CLKMGMT-MIB",
            "path": "raisecom/SWITCH-CLKMGMT-MIB.mib"
        },
        {
            "mib": "CISCOSB-rlLcli-MIB",
            "path": "cisco/CISCOSB-rlLcli-MIB.mib"
        },
        {
            "mib": "JUNIPER-VLAN-MIB",
            "path": "junos/JUNIPER-VLAN-MIB.mib"
        },
        {
            "mib": "Juniper-System-Clock-MIB",
            "path": "junose/Juniper-System-Clock-MIB.mib"
        },
        {
            "mib": "MERU-SMI",
            "path": "fortinet/MERU-SMI.mib"
        },
        {
            "mib": "HH3C-LswRSTP-MIB",
            "path": "comware/HH3C-SPLAT-RSTP-MIB.mib"
        },
        {
            "mib": "SWITCH-CpuLimit-MIB",
            "path": "raisecom/SWITCH-CpuLimit-MIB.mib"
        },
        {
            "mib": "JUNIPER-VMON-MIB",
            "path": "junos/JUNIPER-VMON-MIB.mib"
        },
        {
            "mib": "MERU-SUPPORTED-FEATURES-MIB",
            "path": "fortinet/MERU-SUPPORTED-FEATURES-MIB.mib"
        },
        {
            "mib": "DVMRP-STD-MIB",
            "path": "DVMRP-STD-MIB.mib"
        },
        {
            "mib": "CISCOSB-RMON",
            "path": "cisco/CISCOSB-RMON.mib"
        },
        {
            "mib": "HUAWEI-WLAN-AP-RADIO-MIB",
            "path": "huawei/HUAWEI-WLAN-AP-RADIO-MIB.mib"
        },
        {
            "mib": "HH3C-LswSMON-MIB",
            "path": "comware/HH3C-SPLAT-SMONEXT-MIB.mib"
        },
        {
            "mib": "SWITCH-CPUPRO-MIB",
            "path": "raisecom/SWITCH-CPUPRO-MIB.mib"
        },
        {
            "mib": "MERU-SYSLOG-MIB",
            "path": "fortinet/MERU-SYSLOG-MIB.mib"
        },
        {
            "mib": "JUNIPER-VPN-MIB",
            "path": "junos/JUNIPER-VPN-MIB.mib"
        },
        {
            "mib": "Juniper-System-MIB",
            "path": "junose/Juniper-System-MIB.mib"
        },
        {
            "mib": "CISCOSB-rndApplications",
            "path": "cisco/CISCOSB-rndApplications.mib"
        },
        {
            "mib": "ENTITY-MIB",
            "path": "ENTITY-MIB.mib"
        },
        {
            "mib": "HH3C-LswTRAP-MIB",
            "path": "comware/HH3C-SPLAT-TRAP-MIB.mib"
        },
        {
            "mib": "Juniper-TACACS-Plus-Client-CONF",
            "path": "junose/Juniper-TACACS-Plus-Client-CONF.mib"
        },
        {
            "mib": "JUNIPER-WIRELESS-WAN-MIB",
            "path": "junos/JUNIPER-WIRELESS-WAN-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-AP-SERVICE-MIB",
            "path": "huawei/HUAWEI-WLAN-AP-SERVICE-MIB.mib"
        },
        {
            "mib": "MERU-TOP10-STATISTICS-MIB",
            "path": "fortinet/MERU-TOP10-STATISTICS-MIB.mib"
        },
        {
            "mib": "HH3C-LswVLAN-MIB",
            "path": "comware/HH3C-SPLAT-VLAN-MIB.mib"
        },
        {
            "mib": "ENTITY-SENSOR-MIB",
            "path": "ENTITY-SENSOR-MIB.mib"
        },
        {
            "mib": "MERU-TC",
            "path": "fortinet/MERU-TC.mib"
        },
        {
            "mib": "CISCOSB-rndMng",
            "path": "cisco/CISCOSB-rndMng.mib"
        },
        {
            "mib": "SWITCH-DAI-MIB",
            "path": "raisecom/SWITCH-DAI-MIB.mib"
        },
        {
            "mib": "CISCOSB-ROUTEMAP-MIB",
            "path": "cisco/CISCOSB-ROUTEMAP-MIB.mib"
        },
        {
            "mib": "JUNIPER-WLAN-WAP-MIB",
            "path": "junos/JUNIPER-WLAN-WAP-MIB.mib"
        },
        {
            "mib": "HH3C-SRPOLICY-MIB",
            "path": "comware/HH3C-SRPOLICY-MIB.mib"
        },
        {
            "mib": "SWITCH-ERING-MIB",
            "path": "raisecom/SWITCH-ERING-MIB.mib"
        },
        {
            "mib": "Juniper-TACACS-Plus-Client-MIB",
            "path": "junose/Juniper-TACACS-Plus-Client-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-AP-UPDATE-MIB",
            "path": "huawei/HUAWEI-WLAN-AP-UPDATE-MIB.mib"
        },
        {
            "mib": "CISCOSB-SCT-MIB",
            "path": "cisco/CISCOSB-SCT-MIB.mib"
        },
        {
            "mib": "ENTITY-STATE-MIB",
            "path": "ENTITY-STATE-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-COMMON-MIB",
            "path": "junos/JUNIPER-WX-COMMON-MIB.mib"
        },
        {
            "mib": "MERU-TOPOLOGY-MIB",
            "path": "fortinet/MERU-TOPOLOGY-MIB.mib"
        },
        {
            "mib": "HH3C-SRV6-MIB",
            "path": "comware/HH3C-SRV6-MIB.mib"
        },
        {
            "mib": "SWITCH-FILTER-MIB",
            "path": "raisecom/SWITCH-FILTER-MIB.mib"
        },
        {
            "mib": "MERU-VOICE-STATISTICS-MIB",
            "path": "fortinet/MERU-VOICE-STATISTICS-MIB.mib"
        },
        {
            "mib": "SWITCH-IFEXTEND-MIB",
            "path": "raisecom/SWITCH-IFEXTEND-MIB.mib"
        },
        {
            "mib": "CISCOSB-SECSD-MIB",
            "path": "cisco/CISCOSB-SECSD-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-CAPWAP-MIB",
            "path": "huawei/HUAWEI-WLAN-CAPWAP-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-GLOBAL-REG",
            "path": "junos/JUNIPER-WX-GLOBAL-REG.mib"
        },
        {
            "mib": "CISCOSB-SECURITY-SUITE",
            "path": "cisco/CISCOSB-SECURITY-SUITE.mib"
        },
        {
            "mib": "Juniper-TC",
            "path": "junose/Juniper-TC.mib"
        },
        {
            "mib": "HH3C-SRV6POLICY-MIB",
            "path": "comware/HH3C-SRV6POLICY-MIB.mib"
        },
        {
            "mib": "MERU-WLAN-MIB",
            "path": "fortinet/MERU-WLAN-MIB.mib"
        },
        {
            "mib": "ENTITY-STATE-TC-MIB",
            "path": "ENTITY-STATE-TC-MIB.mib"
        },
        {
            "mib": "SWITCH-IGMPSNOOP-MIB",
            "path": "raisecom/SWITCH-IGMPSNOOP-MIB.mib"
        },
        {
            "mib": "CISCOSB-SENSORENTMIB",
            "path": "cisco/CISCOSB-SENSORENTMIB.mib"
        },
        {
            "mib": "JUNIPER-WX-GLOBAL-TC",
            "path": "junos/JUNIPER-WX-GLOBAL-TC.mib"
        },
        {
            "mib": "Juniper-Trace-Route-CONF",
            "path": "junose/Juniper-Trace-Route-CONF.mib"
        },
        {
            "mib": "HH3C-SSH-MIB",
            "path": "comware/HH3C-SSH-MIB.mib"
        },
        {
            "mib": "EtherLike-MIB",
            "path": "EtherLike-MIB.mib"
        },
        {
            "mib": "SWITCH-INTERFACE-PORT-MIB",
            "path": "raisecom/SWITCH-INTERFACE-PORT-MIB.mib"
        },
        {
            "mib": "CISCOSB-SMARTPORTS-MIB",
            "path": "cisco/CISCOSB-SMARTPORTS-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-MIB",
            "path": "junos/JUNIPER-WX-MIB.mib"
        },
        {
            "mib": "Juniper-TSM-CONF",
            "path": "junose/Juniper-TSM-CONF.mib"
        },
        {
            "mib": "HH3C-SSLVPN-MIB",
            "path": "comware/HH3C-SSLVPN-MIB.mib"
        },
        {
            "mib": "CISCOSB-SMON-MIB",
            "path": "cisco/CISCOSB-SMON-MIB.mib"
        },
        {
            "mib": "SWITCH-L3-MIB",
            "path": "raisecom/SWITCH-L3-MIB.mib"
        },
        {
            "mib": "FCMGMT-MIB",
            "path": "FCMGMT-MIB.mib"
        },
        {
            "mib": "L2L3-VPN-MCAST-MIB",
            "path": "junos/L2L3-VPN-MCAST-MIB.mib"
        },
        {
            "mib": "Juniper-TSM-MIB",
            "path": "junose/Juniper-TSM-MIB.mib"
        },
        {
            "mib": "SWITCH-L3FILTER-MIB",
            "path": "raisecom/SWITCH-L3FILTER-MIB.mib"
        },
        {
            "mib": "HH3C-STACK-MIB",
            "path": "comware/HH3C-STACK-MIB.mib"
        },
        {
            "mib": "CISCOSB-SNMP-MIB",
            "path": "cisco/CISCOSB-SNMP-MIB.mib"
        },
        {
            "mib": "FDDI-SMT73-MIB",
            "path": "FDDI-SMT73-MIB.mib"
        },
        {
            "mib": "LANGTAG-TC-MIB",
            "path": "junos/LANGTAG-TC-MIB.mib"
        },
        {
            "mib": "CISCOSB-SOCKET-MIB",
            "path": "cisco/CISCOSB-SOCKET-MIB.mib"
        },
        {
            "mib": "HH3C-STORAGE-MIB",
            "path": "comware/HH3C-STORAGE-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-GLOBAL-MIB",
            "path": "huawei/HUAWEI-WLAN-GLOBAL-MIB.mib"
        },
        {
            "mib": "FLOAT-TC-MIB",
            "path": "FLOAT-TC-MIB.mib"
        },
        {
            "mib": "SWITCH-LINKSTATETRACK-MIB",
            "path": "raisecom/SWITCH-LINKSTATETRACK-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-CONFIGURATION-MIB",
            "path": "huawei/HUAWEI-WLAN-CONFIGURATION-MIB.mib"
        },
        {
            "mib": "Juniper-UNI-ATM-MIB",
            "path": "junose/Juniper-UNI-ATM-MIB.mib"
        },
        {
            "mib": "HH3C-STORAGE-REF-MIB",
            "path": "comware/HH3C-STORAGE-REF-MIB.mib"
        },
        {
            "mib": "MCAST-VPN-MIB",
            "path": "junos/MCAST-VPN-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-MIB",
            "path": "huawei/HUAWEI-WLAN-MIB.mib"
        },
        {
            "mib": "CISCOSB-SPAN-MIB",
            "path": "cisco/CISCOSB-SPAN-MIB.mib"
        },
        {
            "mib": "SWITCH-MACCONFIG-MIB",
            "path": "raisecom/SWITCH-MACCONFIG-MIB.mib"
        },
        {
            "mib": "CISCOSB-SpecialBpdu-MIB",
            "path": "cisco/CISCOSB-SpecialBpdu-MIB.mib"
        },
        {
            "mib": "FRAME-RELAY-DTE-MIB",
            "path": "FRAME-RELAY-DTE-MIB.mib"
        },
        {
            "mib": "HH3C-STORAGE-SNAP-MIB",
            "path": "comware/HH3C-STORAGE-SNAP-MIB.mib"
        },
        {
            "mib": "Juniper-UNI-IF-MIB",
            "path": "junose/Juniper-UNI-IF-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-NPE-MIB",
            "path": "huawei/HUAWEI-WLAN-NPE-MIB.mib"
        },
        {
            "mib": "SWITCH-MEMORYMANGMENT-MIB",
            "path": "raisecom/SWITCH-MEMORYMANGMENT-MIB.mib"
        },
        {
            "mib": "MPLS-MIB",
            "path": "junos/MPLS-MIB.mib"
        },
        {
            "mib": "Juniper-UNI-SMI",
            "path": "junose/Juniper-UNI-SMI.mib"
        },
        {
            "mib": "HH3C-STORM-CONSTRAIN-MIB",
            "path": "comware/HH3C-STORM-CONSTRAIN-MIB.mib"
        },
        {
            "mib": "CISCOSB-SSH-MIB",
            "path": "cisco/CISCOSB-SSH-MIB.mib"
        },
        {
            "mib": "MPLS-MLDP-STD-MIB",
            "path": "junos/MPLS-MLDP-STD-MIB.mib"
        },
        {
            "mib": "GBOND-MIB",
            "path": "GBOND-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-SAC-MIB",
            "path": "huawei/HUAWEI-WLAN-SAC-MIB.mib"
        },
        {
            "mib": "CISCOSB-SSL",
            "path": "cisco/CISCOSB-SSL.mib"
        },
        {
            "mib": "SWITCH-MSTP-MIB",
            "path": "raisecom/SWITCH-MSTP-MIB.mib"
        },
        {
            "mib": "CISCOSB-STACK-MIB",
            "path": "cisco/CISCOSB-STACK-MIB.mib"
        },
        {
            "mib": "HH3C-SUBNET-VLAN-MIB",
            "path": "comware/HH3C-SUBNET-VLAN-MIB.mib"
        },
        {
            "mib": "SWITCH-MULTISYS-MIB",
            "path": "raisecom/SWITCH-MULTISYS-MIB.mib"
        },
        {
            "mib": "HC-ALARM-MIB",
            "path": "HC-ALARM-MIB.mib"
        },
        {
            "mib": "Juniper-UNI-SONET-MIB",
            "path": "junose/Juniper-UNI-SONET-MIB.mib"
        },
        {
            "mib": "HH3C-SYS-MAN-MIB",
            "path": "comware/HH3C-SYS-MAN-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-STATION-MIB",
            "path": "huawei/HUAWEI-WLAN-STATION-MIB.mib"
        },
        {
            "mib": "CISCOSB-STORMCTRL-MIB",
            "path": "cisco/CISCOSB-STORMCTRL-MIB.mib"
        },
        {
            "mib": "HC-PerfHist-TC-MIB",
            "path": "HC-PerfHist-TC-MIB.mib"
        },
        {
            "mib": "Juniper-V35-CONF",
            "path": "junose/Juniper-V35-CONF.mib"
        },
        {
            "mib": "OPT-IF-MIB",
            "path": "junos/OPT-IF-MIB.mib"
        },
        {
            "mib": "SWITCH-MVR-MIB",
            "path": "raisecom/SWITCH-MVR-MIB.mib"
        },
        {
            "mib": "HH3C-SYSLOG-MIB",
            "path": "comware/HH3C-SYSLOG-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-VAP-MIB",
            "path": "huawei/HUAWEI-WLAN-VAP-MIB.mib"
        },
        {
            "mib": "CISCOSB-SYSLOG-MIB",
            "path": "cisco/CISCOSB-SYSLOG-MIB.mib"
        },
        {
            "mib": "OSPFV3-MIB-JUNIPER",
            "path": "junos/OSPFV3-MIB-JUNIPER.mib"
        },
        {
            "mib": "SWITCH-PORTBACKUP-MIB",
            "path": "raisecom/SWITCH-PORTBACKUP-MIB.mib"
        },
        {
            "mib": "HH3C-T1-MIB",
            "path": "comware/HH3C-T1-MIB.mib"
        },
        {
            "mib": "HC-RMON-MIB",
            "path": "HC-RMON-MIB.mib"
        },
        {
            "mib": "CISCOSB-SYSMNG-MIB",
            "path": "cisco/CISCOSB-SYSMNG-MIB.mib"
        },
        {
            "mib": "Juniper-V35-MIB",
            "path": "junose/Juniper-V35-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTMIRROR-MIB",
            "path": "raisecom/SWITCH-PORTMIRROR-MIB.mib"
        },
        {
            "mib": "CISCOSB-TBI-MIB",
            "path": "cisco/CISCOSB-TBI-MIB.mib"
        },
        {
            "mib": "PCE-PCEP-MIB",
            "path": "junos/PCE-PCEP-MIB.mib"
        },
        {
            "mib": "HH3C-TE-TUNNEL-MIB",
            "path": "comware/HH3C-TE-TUNNEL-MIB.mib"
        },
        {
            "mib": "HUAWEI-WLAN-WIDS-SERVICE-MIB",
            "path": "huawei/HUAWEI-WLAN-WIDS-SERVICE-MIB.mib"
        },
        {
            "mib": "Juniper-VRRP-CONF",
            "path": "junose/Juniper-VRRP-CONF.mib"
        },
        {
            "mib": "HCNUM-TC",
            "path": "HCNUM-TC.mib"
        },
        {
            "mib": "SWITCH-PORTPEERBACKUP-MIB",
            "path": "raisecom/SWITCH-PORTPEERBACKUP-MIB.mib"
        },
        {
            "mib": "CISCOSB-TCPSESSIONS",
            "path": "cisco/CISCOSB-TCPSESSIONS.mib"
        },
        {
            "mib": "PerfHist-TC-MIB",
            "path": "junos/PerfHist-TC-MIB.mib"
        },
        {
            "mib": "HH3C-TRANSCEIVER-INFO-MIB",
            "path": "comware/HH3C-TRANSCEIVER-INFO-MIB.mib"
        },
        {
            "mib": "CISCOSB-TELNET-MIB",
            "path": "cisco/CISCOSB-TELNET-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-COMMON-MIB",
            "path": "junose/JUNIPER-WX-COMMON-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTSECURITY-MIB",
            "path": "raisecom/SWITCH-PORTSECURITY-MIB.mib"
        },
        {
            "mib": "PPP-LCP-MIB",
            "path": "junos/PPP-LCP-MIB.mib"
        },
        {
            "mib": "HDSL2-SHDSL-LINE-MIB",
            "path": "HDSL2-SHDSL-LINE-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-GLOBAL-REG",
            "path": "junose/JUNIPER-WX-GLOBAL-REG.mib"
        },
        {
            "mib": "HH3C-TRAP-MIB",
            "path": "comware/HH3C-TRAP-MIB.mib"
        },
        {
            "mib": "CISCOSB-TIMEBASED-PORT-SHUTDOWN-MIB",
            "path": "cisco/CISCOSB-TIMEBASED-PORT-SHUTDOWN-MIB.mib"
        },
        {
            "mib": "SWITCH-PORTSTATISTIC-MIB",
            "path": "raisecom/SWITCH-PORTSTATISTIC-MIB.mib"
        },
        {
            "mib": "HH3C-TRNG-MIB",
            "path": "comware/HH3C-TRNG-MIB.mib"
        },
        {
            "mib": "JUNIPER-WX-GLOBAL-TC",
            "path": "junose/JUNIPER-WX-GLOBAL-TC.mib"
        },
        {
            "mib": "HUAWEI-XPON-COMMON-MIB",
            "path": "huawei/HUAWEI-XPON-COMMON-MIB.mib"
        },
        {
            "mib": "RADIUS-ACC-CLIENT-MIB",
            "path": "junos/RADIUS-ACC-CLIENT-MIB.mib"
        },
        {
            "mib": "HOST-RESOURCES-MIB",
            "path": "HOST-RESOURCES-MIB.mib"
        },
        {
            "mib": "CISCOSB-TIMESYNCHRONIZATION-MIB",
            "path": "cisco/CISCOSB-TIMESYNCHRONIZATION-MIB.mib"
        },
        {
            "mib": "CISCOSB-TRACEROUTE-MIB",
            "path": "cisco/CISCOSB-TRACEROUTE-MIB.mib"
        },
        {
            "mib": "HH3C-TRNG2-MIB",
            "path": "comware/HH3C-TRNG2-MIB.mib"
        },
        {
            "mib": "SWITCH-RATELIMIT-MIB",
            "path": "raisecom/SWITCH-RATELIMIT-MIB.mib"
        },
        {
            "mib": "RADIUS-AUTH-CLIENT-MIB",
            "path": "junos/RADIUS-AUTH-CLIENT-MIB.mib"
        },
        {
            "mib": "HOST-RESOURCES-TYPES",
            "path": "HOST-RESOURCES-TYPES.mib"
        },
        {
            "mib": "JUNIPER-WX-MIB",
            "path": "junose/JUNIPER-WX-MIB.mib"
        },
        {
            "mib": "CISCOSB-TRAPS-MIB",
            "path": "cisco/CISCOSB-TRAPS-MIB.mib"
        },
        {
            "mib": "SWITCH-RMON-MIB",
            "path": "raisecom/SWITCH-RMON-MIB.mib"
        },
        {
            "mib": "HH3C-TUNNEL-MIB",
            "path": "comware/HH3C-TUNNEL-MIB.mib"
        },
        {
            "mib": "SNMP-COMMUNITY-MIB",
            "path": "junos/SNMP-COMMUNITY-MIB.mib"
        },
        {
            "mib": "PPP-IP-NCP-MIB",
            "path": "junose/PPP-IP-NCP-MIB.mib"
        },
        {
            "mib": "IANA-ADDRESS-FAMILY-NUMBERS-MIB",
            "path": "IANA-ADDRESS-FAMILY-NUMBERS-MIB.mib"
        },
        {
            "mib": "HUAWEI-XPON-MIB",
            "path": "huawei/HUAWEI-XPON-MIB.mib"
        },
        {
            "mib": "SWITCH-RSTP-MIB",
            "path": "raisecom/SWITCH-RSTP-MIB.mib"
        },
        {
            "mib": "CISCOSB-TRUNK-MIB",
            "path": "cisco/CISCOSB-TRUNK-MIB.mib"
        },
        {
            "mib": "HH3C-TWAMP-MIB",
            "path": "comware/HH3C-TWAMP-MIB.mib"
        },
        {
            "mib": "SNMP-FRAMEWORK-MIB",
            "path": "junos/SNMP-FRAMEWORK-MIB.mib"
        },
        {
            "mib": "HH3C-UI-MAN-MIB",
            "path": "comware/HH3C-UI-MAN-MIB.mib"
        },
        {
            "mib": "IANA-CHARSET-MIB",
            "path": "IANA-CHARSET-MIB.mib"
        },
        {
            "mib": "PPP-LCP-MIB",
            "path": "junose/PPP-LCP-MIB.mib"
        },
        {
            "mib": "CISCOSB-Tuning",
            "path": "cisco/CISCOSB-Tuning.mib"
        },
        {
            "mib": "SWITCH-SLOTCARDMGMT-MIB",
            "path": "raisecom/SWITCH-SLOTCARDMGMT-MIB.mib"
        },
        {
            "mib": "SNMP-MPD-MIB",
            "path": "junos/SNMP-MPD-MIB.mib"
        },
        {
            "mib": "HWMUSA-DEV-MIB",
            "path": "huawei/HWMUSA-DEV-MIB.mib"
        },
        {
            "mib": "HH3C-UNICAST-MIB",
            "path": "comware/HH3C-UNICAST-MIB.mib"
        },
        {
            "mib": "IANA-ENTITY-MIB",
            "path": "IANA-ENTITY-MIB.mib"
        },
        {
            "mib": "CISCOSB-TUNNEL-MIB",
            "path": "cisco/CISCOSB-TUNNEL-MIB.mib"
        },
        {
            "mib": "RADIUS-ACC-SERVER-MIB",
            "path": "junose/RADIUS-ACC-SERVER-MIB.mib"
        },
        {
            "mib": "HUAWEI-XQoS-MIB",
            "path": "huawei/HUAWEI-XQoS-MIB.mib"
        },
        {
            "mib": "SNMP-USER-BASED-SM-MIB",
            "path": "junos/SNMP-USER-BASED-SM-MIB.mib"
        },
        {
            "mib": "SWITCH-SNMP-MIB",
            "path": "raisecom/SWITCH-SNMP-MIB.mib"
        },
        {
            "mib": "ISM-HUAWEI-MIB",
            "path": "huawei/ISM-HUAWEI-MIB.mib"
        },
        {
            "mib": "HH3C-UPS-MIB",
            "path": "comware/HH3C-UPS-MIB.mib"
        },
        {
            "mib": "RADIUS-AUTH-SERVER-MIB",
            "path": "junose/RADIUS-AUTH-SERVER-MIB.mib"
        },
        {
            "mib": "CISCOSB-UDLD-MIB",
            "path": "cisco/CISCOSB-UDLD-MIB.mib"
        },
        {
            "mib": "SWITCH-SNTP-MIB",
            "path": "raisecom/SWITCH-SNTP-MIB.mib"
        },
        {
            "mib": "IANA-GMPLS-TC-MIB",
            "path": "IANA-GMPLS-TC-MIB.mib"
        },
        {
            "mib": "ISM-PERFORMANCE-MIB",
            "path": "huawei/ISM-PERFORMANCE-MIB.mib"
        },
        {
            "mib": "SNMP-VIEW-BASED-ACM-MIB",
            "path": "junos/SNMP-VIEW-BASED-ACM-MIB.mib"
        },
        {
            "mib": "HH3C-USER-MIB",
            "path": "comware/HH3C-USER-MIB.mib"
        },
        {
            "mib": "CISCOSB-UDP",
            "path": "cisco/CISCOSB-UDP.mib"
        },
        {
            "mib": "SWITCH-SYNCE-MIB",
            "path": "raisecom/SWITCH-SYNCE-MIB.mib"
        },
        {
            "mib": "IANA-ITU-ALARM-TC-MIB",
            "path": "IANA-ITU-ALARM-TC-MIB.mib"
        },
        {
            "mib": "ISM-STORAGE-SVC-MIB",
            "path": "huawei/ISM-STORAGE-SVC-MIB.mib"
        },
        {
            "mib": "SYSAPPL-MIB",
            "path": "junos/SYSAPPL-MIB.mib"
        },
        {
            "mib": "HH3C-USERLOG-MIB",
            "path": "comware/HH3C-USERLOG-MIB.mib"
        },
        {
            "mib": "CISCOSB-vlan-MIB",
            "path": "cisco/CISCOSB-vlan-MIB.mib"
        },
        {
            "mib": "SWITCH-SYSTEM-MIB",
            "path": "raisecom/SWITCH-SYSTEM-MIB.mib"
        },
        {
            "mib": "IANA-LANGUAGE-MIB",
            "path": "IANA-LANGUAGE-MIB.mib"
        },
        {
            "mib": "ISM-TRAP-MIB",
            "path": "huawei/ISM-TRAP-MIB.mib"
        },
        {
            "mib": "VPLS-BGP-DRAFT-01-MIB",
            "path": "junos/VPLS-BGP-DRAFT-01-MIB.mib"
        },
        {
            "mib": "CISCOSB-vlanVoice-MIB",
            "path": "cisco/CISCOSB-vlanVoice-MIB.mib"
        },
        {
            "mib": "SWITCH-TC",
            "path": "raisecom/SWITCH-TC.mib"
        },
        {
            "mib": "HH3C-VBR-MIB",
            "path": "comware/HH3C-VBR-MIB.mib"
        },
        {
            "mib": "IANA-MAU-MIB",
            "path": "IANA-MAU-MIB.mib"
        },
        {
            "mib": "CISCOSB-VRRP",
            "path": "cisco/CISCOSB-VRRP.mib"
        },
        {
            "mib": "VPLS-GENERIC-DRAFT-01-MIB",
            "path": "junos/VPLS-GENERIC-DRAFT-01-MIB.mib"
        },
        {
            "mib": "HH3C-VLANGROUP-MIB",
            "path": "comware/HH3C-VLANGROUP-MIB.mib"
        },
        {
            "mib": "SWITCH-TRUNK-MIB",
            "path": "raisecom/SWITCH-TRUNK-MIB.mib"
        },
        {
            "mib": "HH3C-VLANTERM-MIB",
            "path": "comware/HH3C-VLANTERM-MIB.mib"
        },
        {
            "mib": "NQA-MIB",
            "path": "huawei/NQA-MIB.mib"
        },
        {
            "mib": "CISCOSB-WBA-MIB",
            "path": "cisco/CISCOSB-WBA-MIB.mib"
        },
        {
            "mib": "VPLS-LDP-DRAFT-01-MIB",
            "path": "junos/VPLS-LDP-DRAFT-01-MIB.mib"
        },
        {
            "mib": "IANA-PRINTER-MIB",
            "path": "IANA-PRINTER-MIB.mib"
        },
        {
            "mib": "SWITCH-VLAN-MIB",
            "path": "raisecom/SWITCH-VLAN-MIB.mib"
        },
        {
            "mib": "CISCOSB-WeightedRandomTailDrop-MIB",
            "path": "cisco/CISCOSB-WeightedRandomTailDrop-MIB.mib"
        },
        {
            "mib": "OPTIX-BOARD-MANAGE-MIB",
            "path": "huawei/OPTIX-BOARD-MANAGE-MIB.mib"
        },
        {
            "mib": "HH3C-VM-MAN-MIB",
            "path": "comware/HH3C-VM-MAN-MIB.mib"
        },
        {
            "mib": "IANA-PWE3-MIB",
            "path": "IANA-PWE3-MIB.mib"
        },
        {
            "mib": "SWITCH-VLANCFG-MIB",
            "path": "raisecom/SWITCH-VLANCFG-MIB.mib"
        },
        {
            "mib": "CISCOSBLAN1-MIB",
            "path": "cisco/CISCOSBLAN1-MIB.mib"
        },
        {
            "mib": "HH3C-VMAP-MIB",
            "path": "comware/HH3C-VMAP-MIB.mib"
        },
        {
            "mib": "IANA-RTPROTO-MIB",
            "path": "IANA-RTPROTO-MIB.mib"
        },
        {
            "mib": "SWITCH-VLANPORT-RATELIMIT-MIB",
            "path": "raisecom/SWITCH-VLANPORT-RATELIMIT-MIB.mib"
        },
        {
            "mib": "OPTIX-MISC-MIB",
            "path": "huawei/OPTIX-MISC-MIB.mib"
        },
        {
            "mib": "CISCOSMB-MIB",
            "path": "cisco/CISCOSMB-MIB.mib"
        },
        {
            "mib": "HH3C-VNF-DEVICE-MIB",
            "path": "comware/HH3C-VNF-DEVICE-MIB.mib"
        },
        {
            "mib": "CISCOWAN-SMI",
            "path": "cisco/CISCOWAN-SMI.mib"
        },
        {
            "mib": "SWTICH-SERVICE-MIB",
            "path": "raisecom/SWTICH-SERVICE-MIB.mib"
        },
        {
            "mib": "IANAifType-MIB",
            "path": "IANAifType-MIB.mib"
        },
        {
            "mib": "OPTIX-OID-MIB",
            "path": "huawei/OPTIX-OID-MIB.mib"
        },
        {
            "mib": "OPTIX-NE-MIB",
            "path": "huawei/OPTIX-NE-MIB.mib"
        },
        {
            "mib": "HH3C-VOICE-CALL-ACTIVE-MIB",
            "path": "comware/HH3C-VOICE-CALL-ACTIVE-MIB.mib"
        },
        {
            "mib": "DIFF-SERV-MIB",
            "path": "cisco/DIFF-SERV-MIB.mib"
        },
        {
            "mib": "SWTICH-VLANXC-MIB",
            "path": "raisecom/SWTICH-VLANXC-MIB.mib"
        },
        {
            "mib": "HH3C-VOICE-CALL-HISTORY-MIB",
            "path": "comware/HH3C-VOICE-CALL-HISTORY-MIB.mib"
        },
        {
            "mib": "OPTIX-RTN-ODU-MGR-MIB",
            "path": "huawei/OPTIX-RTN-ODU-MGR-MIB.mib"
        },
        {
            "mib": "HH3C-VOICE-DIAL-CONTROL-MIB",
            "path": "comware/HH3C-VOICE-DIAL-CONTROL-MIB.mib"
        },
        {
            "mib": "IEEE-802DOT17-RPR-MIB",
            "path": "IEEE-802DOT17-RPR-MIB.mib"
        },
        {
            "mib": "DOCS-IF-MIB",
            "path": "cisco/DOCS-IF-MIB.mib"
        },
        {
            "mib": "IEEE8021-BRIDGE-MIB",
            "path": "IEEE8021-BRIDGE-MIB.mib"
        },
        {
            "mib": "IRONPORT-SMI",
            "path": "cisco/IRONPORT-SMI.mib"
        },
        {
            "mib": "HH3C-VOICE-IF-MIB",
            "path": "comware/HH3C-VOICE-IF-MIB.mib"
        },
        {
            "mib": "HH3C-VOICE-VLAN-MIB",
            "path": "comware/HH3C-VOICE-VLAN-MIB.mib"
        },
        {
            "mib": "MPLS-LDP-MIB",
            "path": "cisco/MPLS-LDP-MIB.mib"
        },
        {
            "mib": "IEEE8021-CFM-MIB",
            "path": "IEEE8021-CFM-MIB.mib"
        },
        {
            "mib": "OLD-CISCO-CHASSIS-MIB",
            "path": "cisco/OLD-CISCO-CHASSIS-MIB.mib"
        },
        {
            "mib": "HH3C-VPN-PEER-MIB",
            "path": "comware/HH3C-VPN-PEER-MIB.mib"
        },
        {
            "mib": "OLD-CISCO-CPU-MIB",
            "path": "cisco/OLD-CISCO-CPU-MIB.mib"
        },
        {
            "mib": "IEEE8021-CFMD8-MIB",
            "path": "IEEE8021-CFMD8-MIB.mib"
        },
        {
            "mib": "HH3C-VRRP-EXT-MIB",
            "path": "comware/HH3C-VRRP-EXT-MIB.mib"
        },
        {
            "mib": "IEEE8021-PAE-MIB",
            "path": "IEEE8021-PAE-MIB.mib"
        },
        {
            "mib": "OLD-CISCO-INTERFACES-MIB",
            "path": "cisco/OLD-CISCO-INTERFACES-MIB.mib"
        },
        {
            "mib": "HH3C-VSAN-MIB",
            "path": "comware/HH3C-VSAN-MIB.mib"
        },
        {
            "mib": "OLD-CISCO-MEMORY-MIB",
            "path": "cisco/OLD-CISCO-MEMORY-MIB.mib"
        },
        {
            "mib": "HH3C-VSI-MIB",
            "path": "comware/HH3C-VSI-MIB.mib"
        },
        {
            "mib": "IEEE8021-Q-BRIDGE-MIB",
            "path": "IEEE8021-Q-BRIDGE-MIB.mib"
        },
        {
            "mib": "OLD-CISCO-SYS-MIB",
            "path": "cisco/OLD-CISCO-SYS-MIB.mib"
        },
        {
            "mib": "IEEE8021-SECY-MIB",
            "path": "IEEE8021-SECY-MIB.mib"
        },
        {
            "mib": "HH3C-VXLAN-MIB",
            "path": "comware/HH3C-VXLAN-MIB.mib"
        },
        {
            "mib": "SA-CM-MIB",
            "path": "cisco/SA-CM-MIB.mib"
        },
        {
            "mib": "IEEE8021-TC-MIB",
            "path": "IEEE8021-TC-MIB.mib"
        },
        {
            "mib": "SA-CM-MTA-MIB",
            "path": "cisco/SA-CM-MTA-MIB.mib"
        },
        {
            "mib": "HH3C-WAPI-MIB",
            "path": "comware/HH3C-WAPI-MIB.mib"
        },
        {
            "mib": "HH3C-WEB-AUTHENTICATION-MIB",
            "path": "comware/HH3C-WEB-AUTHENTICATION-MIB.mib"
        },
        {
            "mib": "IEEE802171-CFM-MIB",
            "path": "IEEE802171-CFM-MIB.mib"
        },
        {
            "mib": "SA-HARDWARE-MIB",
            "path": "cisco/SA-HARDWARE-MIB.mib"
        },
        {
            "mib": "IEEE8023-LAG-MIB",
            "path": "IEEE8023-LAG-MIB.mib"
        },
        {
            "mib": "SA-MTA-MIB",
            "path": "cisco/SA-MTA-MIB.mib"
        },
        {
            "mib": "HH3C-WLAN-FLEXAPP-CFG-MIB",
            "path": "comware/HH3C-WLAN-FLEXAPP-CFG-MIB.mib"
        },
        {
            "mib": "HH3C-WIPS-MIB",
            "path": "comware/HH3C-WIPS-MIB.mib"
        },
        {
            "mib": "IEEE802dot11-MIB",
            "path": "IEEE802dot11-MIB.mib"
        },
        {
            "mib": "HH3C-WLANMT-MIB",
            "path": "comware/HH3C-WLANMT-MIB.mib"
        },
        {
            "mib": "SA-RG-MIB",
            "path": "cisco/SA-RG-MIB.mib"
        },
        {
            "mib": "IGMP-MIB",
            "path": "IGMP-MIB.mib"
        },
        {
            "mib": "IF-MIB",
            "path": "IF-MIB.mib"
        },
        {
            "mib": "IGMP-STD-MIB",
            "path": "IGMP-STD-MIB.mib"
        },
        {
            "mib": "INET-ADDRESS-MIB",
            "path": "INET-ADDRESS-MIB.mib"
        },
        {
            "mib": "INT-SERV-MIB",
            "path": "INT-SERV-MIB.mib"
        },
        {
            "mib": "INTEGRATED-SERVICES-MIB",
            "path": "INTEGRATED-SERVICES-MIB.mib"
        },
        {
            "mib": "IP-FORWARD-MIB",
            "path": "IP-FORWARD-MIB.mib"
        },
        {
            "mib": "IP-MIB",
            "path": "IP-MIB.mib"
        },
        {
            "mib": "IPV6-FLOW-LABEL-MIB",
            "path": "IPV6-FLOW-LABEL-MIB.mib"
        },
        {
            "mib": "IPMROUTE-STD-MIB",
            "path": "IPMROUTE-STD-MIB.mib"
        },
        {
            "mib": "IPMROUTE-MIB",
            "path": "IPMROUTE-MIB.mib"
        },
        {
            "mib": "IPV6-ICMP-MIB",
            "path": "IPV6-ICMP-MIB.mib"
        },
        {
            "mib": "IPV6-MIB",
            "path": "IPV6-MIB.mib"
        },
        {
            "mib": "IPV6-MLD-MIB",
            "path": "IPV6-MLD-MIB.mib"
        },
        {
            "mib": "IPV6-TC",
            "path": "IPV6-TC.mib"
        },
        {
            "mib": "IPV6-TCP-MIB",
            "path": "IPV6-TCP-MIB.mib"
        },
        {
            "mib": "IPV6-UDP-MIB",
            "path": "IPV6-UDP-MIB.mib"
        },
        {
            "mib": "ISDN-MIB",
            "path": "ISDN-MIB.mib"
        },
        {
            "mib": "ITU-ALARM-TC-MIB",
            "path": "ITU-ALARM-TC-MIB.mib"
        },
        {
            "mib": "ISIS-MIB",
            "path": "ISIS-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-DCBX-MIB",
            "path": "LLDP-EXT-DCBX-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-DOT1-MIB",
            "path": "LLDP-EXT-DOT1-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-DOT3-MIB",
            "path": "LLDP-EXT-DOT3-MIB.mib"
        },
        {
            "mib": "LLDP-EXT-MED-MIB",
            "path": "LLDP-EXT-MED-MIB.mib"
        },
        {
            "mib": "LLDP-MIB",
            "path": "LLDP-MIB.mib"
        },
        {
            "mib": "LLDP-V2-MIB",
            "path": "LLDP-V2-MIB.mib"
        },
        {
            "mib": "LLDP-V2-TC-MIB",
            "path": "LLDP-V2-TC-MIB.mib"
        },
        {
            "mib": "MGMD-STD-MIB",
            "path": "MGMD-STD-MIB.mib"
        },
        {
            "mib": "MAU-MIB",
            "path": "MAU-MIB.mib"
        },
        {
            "mib": "MPLS-L3VPN-STD-MIB",
            "path": "MPLS-L3VPN-STD-MIB.mib"
        },
        {
            "mib": "MPLS-LDP-STD-MIB",
            "path": "MPLS-LDP-STD-MIB.mib"
        },
        {
            "mib": "MPLS-LSR-MIB",
            "path": "MPLS-LSR-MIB.mib"
        },
        {
            "mib": "MPLS-LSR-STD-MIB",
            "path": "MPLS-LSR-STD-MIB.mib"
        },
        {
            "mib": "MPLS-TC-STD-MIB",
            "path": "MPLS-TC-STD-MIB.mib"
        },
        {
            "mib": "MPLS-TE-MIB",
            "path": "MPLS-TE-MIB.mib"
        },
        {
            "mib": "MPLS-TE-STD-MIB",
            "path": "MPLS-TE-STD-MIB.mib"
        },
        {
            "mib": "MPLS-VPN-MIB",
            "path": "MPLS-VPN-MIB.mib"
        },
        {
            "mib": "MSTP-MIB",
            "path": "MSTP-MIB.mib"
        },
        {
            "mib": "MTA-MIB",
            "path": "MTA-MIB.mib"
        },
        {
            "mib": "NETWORK-SERVICES-MIB",
            "path": "NETWORK-SERVICES-MIB.mib"
        },
        {
            "mib": "NOTIFICATION-LOG-MIB",
            "path": "NOTIFICATION-LOG-MIB.mib"
        },
        {
            "mib": "OSPF-MIB",
            "path": "OSPF-MIB.mib"
        },
        {
            "mib": "OSPF-TRAP-MIB",
            "path": "OSPF-TRAP-MIB.mib"
        },
        {
            "mib": "OSPFV3-MIB",
            "path": "OSPFV3-MIB.mib"
        },
        {
            "mib": "PerfHist-TC-MIB",
            "path": "PerfHist-TC-MIB.mib"
        },
        {
            "mib": "P-BRIDGE-MIB",
            "path": "P-BRIDGE-MIB.mib"
        },
        {
            "mib": "PIM-MIB",
            "path": "PIM-MIB.mib"
        },
        {
            "mib": "POWER-ETHERNET-MIB",
            "path": "POWER-ETHERNET-MIB.mib"
        },
        {
            "mib": "PPVPN-TC-MIB",
            "path": "PPVPN-TC-MIB.mib"
        },
        {
            "mib": "PTOPO-MIB",
            "path": "PTOPO-MIB.mib"
        },
        {
            "mib": "Printer-MIB",
            "path": "Printer-MIB.mib"
        },
        {
            "mib": "PW-STD-MIB",
            "path": "PW-STD-MIB.mib"
        },
        {
            "mib": "PW-TC-STD-MIB",
            "path": "PW-TC-STD-MIB.mib"
        },
        {
            "mib": "RFC-1215",
            "path": "RFC-1215.mib"
        },
        {
            "mib": "Q-BRIDGE-MIB",
            "path": "Q-BRIDGE-MIB.mib"
        },
        {
            "mib": "RFC-1212",
            "path": "RFC-1212.mib"
        },
        {
            "mib": "RFC1155-SMI",
            "path": "RFC1155-SMI.mib"
        },
        {
            "mib": "RFC1213-MIB",
            "path": "RFC1213-MIB.mib"
        },
        {
            "mib": "RFC1284-MIB",
            "path": "RFC1284-MIB.mib"
        },
        {
            "mib": "RFC1271-MIB",
            "path": "RFC1271-MIB.mib"
        },
        {
            "mib": "RFC1389-MIB",
            "path": "RFC1389-MIB.mib"
        },
        {
            "mib": "RIPv2-MIB",
            "path": "RIPv2-MIB.mib"
        },
        {
            "mib": "RMON-MIB",
            "path": "RMON-MIB.mib"
        },
        {
            "mib": "RMON2-MIB",
            "path": "RMON2-MIB.mib"
        },
        {
            "mib": "RSTP-MIB",
            "path": "RSTP-MIB.mib"
        },
        {
            "mib": "SCTP-MIB",
            "path": "SCTP-MIB.mib"
        },
        {
            "mib": "SMON-MIB",
            "path": "SMON-MIB.mib"
        },
        {
            "mib": "SNA-SDLC-MIB",
            "path": "SNA-SDLC-MIB.mib"
        },
        {
            "mib": "SNMP-COMMUNITY-MIB",
            "path": "SNMP-COMMUNITY-MIB.mib"
        },
        {
            "mib": "SNMP-MPD-MIB",
            "path": "SNMP-MPD-MIB.mib"
        },
        {
            "mib": "SNMP-FRAMEWORK-MIB",
            "path": "SNMP-FRAMEWORK-MIB.mib"
        },
        {
            "mib": "SNMP-NOTIFICATION-MIB",
            "path": "SNMP-NOTIFICATION-MIB.mib"
        },
        {
            "mib": "SNMP-PROXY-MIB",
            "path": "SNMP-PROXY-MIB.mib"
        },
        {
            "mib": "SNMP-REPEATER-MIB",
            "path": "SNMP-REPEATER-MIB.mib"
        },
        {
            "mib": "SNMP-TARGET-MIB",
            "path": "SNMP-TARGET-MIB.mib"
        },
        {
            "mib": "SNMP-USER-BASED-SM-MIB",
            "path": "SNMP-USER-BASED-SM-MIB.mib"
        },
        {
            "mib": "SNMP-USM-AES-MIB",
            "path": "SNMP-USM-AES-MIB.mib"
        },
        {
            "mib": "SNMP-USM-DH-OBJECTS-MIB",
            "path": "SNMP-USM-DH-OBJECTS-MIB.mib"
        },
        {
            "mib": "SNMP-VIEW-BASED-ACM-MIB",
            "path": "SNMP-VIEW-BASED-ACM-MIB.mib"
        },
        {
            "mib": "SNMPv2-CONF",
            "path": "SNMPv2-CONF.mib"
        },
        {
            "mib": "SNMPv2-SMI-v1",
            "path": "SNMPv2-SMI-v1.mib"
        },
        {
            "mib": "SNMPv2-MIB",
            "path": "SNMPv2-MIB.mib"
        },
        {
            "mib": "SNMPv2-SMI",
            "path": "SNMPv2-SMI.mib"
        },
        {
            "mib": "SNMPv2-TC-v1",
            "path": "SNMPv2-TC-v1.mib"
        },
        {
            "mib": "SNMPv2-TM",
            "path": "SNMPv2-TM.mib"
        },
        {
            "mib": "SNMPv2-TC",
            "path": "SNMPv2-TC.mib"
        },
        {
            "mib": "SWITCH-TC",
            "path": "SWITCH-TC.mib"
        },
        {
            "mib": "SONET-MIB",
            "path": "SONET-MIB.mib"
        },
        {
            "mib": "TCP-MIB",
            "path": "TCP-MIB.mib"
        },
        {
            "mib": "SYSAPPL-MIB",
            "path": "SYSAPPL-MIB.mib"
        },
        {
            "mib": "TRANSPORT-ADDRESS-MIB",
            "path": "TRANSPORT-ADDRESS-MIB.mib"
        },
        {
            "mib": "TOKEN-RING-RMON-MIB",
            "path": "TOKEN-RING-RMON-MIB.mib"
        },
        {
            "mib": "UUID-TC-MIB",
            "path": "UUID-TC-MIB.mib"
        },
        {
            "mib": "UDP-MIB",
            "path": "UDP-MIB.mib"
        },
        {
            "mib": "TUNNEL-MIB",
            "path": "TUNNEL-MIB.mib"
        },
        {
            "mib": "UPS-MIB",
            "path": "UPS-MIB.mib"
        },
        {
            "mib": "VDSL-LINE-MIB",
            "path": "VDSL-LINE-MIB.mib"
        },
        {
            "mib": "VPN-TC-STD-MIB",
            "path": "VPN-TC-STD-MIB.mib"
        },
        {
            "mib": "VDSL2-LINE-TC-MIB",
            "path": "VDSL2-LINE-TC-MIB.mib"
        },
        {
            "mib": "VDSL2-LINE-MIB",
            "path": "VDSL2-LINE-MIB.mib"
        },
        {
            "mib": "VRRPV3-MIB",
            "path": "VRRPV3-MIB.mib"
        },
        {
            "mib": "VRRP-MIB",
            "path": "VRRP-MIB.mib"
        }
    ]
}
```

</details>

### Возможные коды ошибок

400: Bad Request - Передан не валидный JSON  
500: Internal Server Error - Ошибка базы данных при получении данных файлов MIB

---

## [GET] /api/v1/mib-parser/mib/{path} - Чтение MIB, находящегося по указанному пути

> [!TIP]
> API позволяет получить текст файла MIB в виде ByteString.

<details><summary>Примеры запросов</summary>

### Примеры запросов

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/mib-parser/mib?path=SNMPv2-MIB.mib

{}
```

Ответ 1:

```json
{
    "mib": "LS0gKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioNCi0tIFNOTVB2Mi1NSUIubXk6ICBNYW5hZ2VtZW50IEluZm9ybWF0aW9uIEJhc2UgKE1JQikgZm9yIHRoZQ0KLS0gICAgICAgICAgICAgICAgIFNpbXBsZSBOZXR3b3JrIE1hbmFnZW1lbnQgUHJvdG9jb2wgKFNOTVApDQotLQ0KLS0gTWFyY2ggMjAwNiwgIFdlbiBYdQ0KLS0NCi0tIENvcHlyaWdodCAoYykgMjAwNiBieSBjaXNjbyBTeXN0ZW1zLCBJbmMuDQotLSBBbGwgcmlnaHRzIHJlc2VydmVkLg0KLS0gKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioNCi0tDQotLSBUaGlzIG1pYiB3YXMgZXh0cmFjdGVkIGZyb20gUkZDIDM0MTgNCg0KDQpTTk1QdjItTUlCIERFRklOSVRJT05TIDo6PSBCRUdJTg0KDQogICBJTVBPUlRTDQogICAgICAgTU9EVUxFLUlERU5USVRZLCBPQkpFQ1QtVFlQRSwgTk9USUZJQ0FUSU9OLVRZUEUsDQogICAgICAgVGltZVRpY2tzLCBDb3VudGVyMzIsIHNubXBNb2R1bGVzLCBtaWItMg0KICAgICAgICAgICBGUk9NIFNOTVB2Mi1TTUkNCiAgICAgICBEaXNwbGF5U3RyaW5nLCBUZXN0QW5kSW5jciwgVGltZVN0YW1wDQogICAgICAgICAgIEZST00gU05NUHYyLVRDDQogICAgICAgTU9EVUxFLUNPTVBMSUFOQ0UsIE9CSkVDVC1HUk9VUCwgTk9USUZJQ0FUSU9OLUdST1VQDQogICAgICAgICAgIEZST00gU05NUHYyLUNPTkY7DQoNCiAgIHNubXBNSUIgTU9EVUxFLUlERU5USVRZDQogICAgICAgTEFTVC1VUERBVEVEICIyMDAyMTAxNjAwMDBaIg0KICAgICAgIE9SR0FOSVpBVElPTiAiSUVURiBTTk1QdjMgV29ya2luZyBHcm91cCINCiAgICAgICBDT05UQUNULUlORk8NCiAgICAgICAgICAgICAgICJXRy1FTWFpbDogICBzbm1wdjNAbGlzdHMudGlzbGFicy5jb20NCiAgICAgICAgICAgICAgICBTdWJzY3JpYmU6ICBzbm1wdjMtcmVxdWVzdEBsaXN0cy50aXNsYWJzLmNvbQ0KDQogICAgICAgICAgICAgICAgQ28tQ2hhaXI6ICAgUnVzcyBNdW5keQ0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIE5ldHdvcmsgQXNzb2NpYXRlcyBMYWJvcmF0b3JpZXMNCiAgICAgICAgICAgICAgICBwb3N0YWw6ICAgICAxNTIwNCBPbWVnYSBEcml2ZSwgU3VpdGUgMzAwDQogICAgICAgICAgICAgICAgICAgICAgICAgICAgUm9ja3ZpbGxlLCBNRCAyMDg1MC00NjAxDQogICAgICAgICAgICAgICAgICAgICAgICAgICAgVVNBDQogICAgICAgICAgICAgICAgRU1haWw6ICAgICAgbXVuZHlAdGlzbGFicy5jb20NCiAgICAgICAgICAgICAgICBwaG9uZTogICAgICArMSAzMDEgOTQ3LTcxMDcNCg0KICAgICAgICAgICAgICAgIENvLUNoYWlyOiAgIERhdmlkIEhhcnJpbmd0b24NCiAgICAgICAgICAgICAgICAgICAgICAgICAgICBFbnRlcmFzeXMgTmV0d29ya3MNCiAgICAgICAgICAgICAgICBwb3N0YWw6ICAgICAzNSBJbmR1c3RyaWFsIFdheQ0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIFAuIE8uIEJveCA1MDA1DQogICAgICAgICAgICAgICAgICAgICAgICAgICAgUm9jaGVzdGVyLCBOSCAwMzg2Ni01MDA1DQogICAgICAgICAgICAgICAgICAgICAgICAgICAgVVNBDQogICAgICAgICAgICAgICAgRU1haWw6ICAgICAgZGJoQGVudGVyYXN5cy5jb20NCiAgICAgICAgICAgICAgICBwaG9uZTogICAgICArMSA2MDMgMzM3LTI2MTQNCg0KICAgICAgICAgICAgICAgIEVkaXRvcjogICAgIFJhbmR5IFByZXN1aG4NCiAgICAgICAgICAgICAgICAgICAgICAgICAgICBCTUMgU29mdHdhcmUsIEluYy4NCiAgICAgICAgICAgICAgICBwb3N0YWw6ICAgICAyMTQxIE5vcnRoIEZpcnN0IFN0cmVldA0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIFNhbiBKb3NlLCBDQSA5NTEzMQ0KICAgICAgICAgICAgICAgICAgICAgICAgICAgIFVTQQ0KICAgICAgICAgICAgICAgIEVNYWlsOiAgICAgIHJhbmR5X3ByZXN1aG5AYm1jLmNvbQ0KICAgICAgICAgICAgICAgIHBob25lOiAgICAgICsxIDQwOCA1NDYtMTAwNiINCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSBNSUIgbW9kdWxlIGZvciBTTk1QIGVudGl0aWVzLg0KDQogICAgICAgICAgICAgICAgQ29weXJpZ2h0IChDKSBUaGUgSW50ZXJuZXQgU29jaWV0eSAoMjAwMikuIFRoaXMNCiAgICAgICAgICAgICAgICB2ZXJzaW9uIG9mIHRoaXMgTUlCIG1vZHVsZSBpcyBwYXJ0IG9mIFJGQyAzNDE4Ow0KICAgICAgICAgICAgICAgIHNlZSB0aGUgUkZDIGl0c2VsZiBmb3IgZnVsbCBsZWdhbCBub3RpY2VzLg0KICAgICAgICAgICAgICAgIg0KICAgICAgIFJFVklTSU9OICAgICAgIjIwMDIxMDE2MDAwMFoiDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGlzIHJldmlzaW9uIG9mIHRoaXMgTUlCIG1vZHVsZSB3YXMgcHVibGlzaGVkIGFzDQogICAgICAgICAgICAgICAgUkZDIDM0MTguIg0KICAgICAgIFJFVklTSU9OICAgICAgIjE5OTUxMTA5MDAwMFoiDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGlzIHJldmlzaW9uIG9mIHRoaXMgTUlCIG1vZHVsZSB3YXMgcHVibGlzaGVkIGFzDQogICAgICAgICAgICAgICAgUkZDIDE5MDcuIg0KICAgICAgIFJFVklTSU9OICAgICAgIjE5OTMwNDAxMDAwMFoiDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgaW5pdGlhbCByZXZpc2lvbiBvZiB0aGlzIE1JQiBtb2R1bGUgd2FzIHB1Ymxpc2hlZA0KICAgICAgICAgICAgICAgYXMgUkZDIDE0NTAuIg0KICAgICAgIDo6PSB7IHNubXBNb2R1bGVzIDEgfQ0KDQogICBzbm1wTUlCT2JqZWN0cyBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBzbm1wTUlCIDEgfQ0KDQogICAtLSAgOjo9IHsgc25tcE1JQk9iamVjdHMgMSB9ICAgICAgICB0aGlzIE9JRCBpcyBvYnNvbGV0ZQ0KICAgLS0gIDo6PSB7IHNubXBNSUJPYmplY3RzIDIgfSAgICAgICAgdGhpcyBPSUQgaXMgb2Jzb2xldGUNCiAgIC0tICA6Oj0geyBzbm1wTUlCT2JqZWN0cyAzIH0gICAgICAgIHRoaXMgT0lEIGlzIG9ic29sZXRlDQoNCiAgIC0tIHRoZSBTeXN0ZW0gZ3JvdXANCiAgIC0tDQogICAtLSBhIGNvbGxlY3Rpb24gb2Ygb2JqZWN0cyBjb21tb24gdG8gYWxsIG1hbmFnZWQgc3lzdGVtcy4NCg0KICAgc3lzdGVtICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgbWliLTIgMSB9DQoNCiAgIHN5c0Rlc2NyIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICAgRGlzcGxheVN0cmluZyAoU0laRSAoMC4uMjU1KSkNCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJBIHRleHR1YWwgZGVzY3JpcHRpb24gb2YgdGhlIGVudGl0eS4gIFRoaXMgdmFsdWUgc2hvdWxkDQogICAgICAgICAgICAgICBpbmNsdWRlIHRoZSBmdWxsIG5hbWUgYW5kIHZlcnNpb24gaWRlbnRpZmljYXRpb24gb2YNCiAgICAgICAgICAgICAgIHRoZSBzeXN0ZW0ncyBoYXJkd2FyZSB0eXBlLCBzb2Z0d2FyZSBvcGVyYXRpbmctc3lzdGVtLA0KICAgICAgICAgICAgICAgYW5kIG5ldHdvcmtpbmcgc29mdHdhcmUuIg0KICAgICAgIDo6PSB7IHN5c3RlbSAxIH0NCg0KICAgc3lzT2JqZWN0SUQgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBPQkpFQ1QgSURFTlRJRklFUg0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB2ZW5kb3IncyBhdXRob3JpdGF0aXZlIGlkZW50aWZpY2F0aW9uIG9mIHRoZQ0KICAgICAgICAgICAgICAgbmV0d29yayBtYW5hZ2VtZW50IHN1YnN5c3RlbSBjb250YWluZWQgaW4gdGhlIGVudGl0eS4NCiAgICAgICAgICAgICAgIFRoaXMgdmFsdWUgaXMgYWxsb2NhdGVkIHdpdGhpbiB0aGUgU01JIGVudGVycHJpc2VzDQogICAgICAgICAgICAgICBzdWJ0cmVlICgxLjMuNi4xLjQuMSkgYW5kIHByb3ZpZGVzIGFuIGVhc3kgYW5kDQogICAgICAgICAgICAgICB1bmFtYmlndW91cyBtZWFucyBmb3IgZGV0ZXJtaW5pbmcgYHdoYXQga2luZCBvZiBib3gnIGlzDQogICAgICAgICAgICAgICBiZWluZyBtYW5hZ2VkLiAgRm9yIGV4YW1wbGUsIGlmIHZlbmRvciBgRmxpbnRzdG9uZXMsDQogICAgICAgICAgICAgICBJbmMuJyB3YXMgYXNzaWduZWQgdGhlIHN1YnRyZWUgMS4zLjYuMS40LjEuNDI0MjQyLA0KICAgICAgICAgICAgICAgaXQgY291bGQgYXNzaWduIHRoZSBpZGVudGlmaWVyIDEuMy42LjEuNC4xLjQyNDI0Mi4xLjENCiAgICAgICAgICAgICAgIHRvIGl0cyBgRnJlZCBSb3V0ZXInLiINCiAgICAgICA6Oj0geyBzeXN0ZW0gMiB9DQoNCiAgIHN5c1VwVGltZSBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIFRpbWVUaWNrcw0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB0aW1lIChpbiBodW5kcmVkdGhzIG9mIGEgc2Vjb25kKSBzaW5jZSB0aGUNCiAgICAgICAgICAgICAgIG5ldHdvcmsgbWFuYWdlbWVudCBwb3J0aW9uIG9mIHRoZSBzeXN0ZW0gd2FzIGxhc3QNCiAgICAgICAgICAgICAgIHJlLWluaXRpYWxpemVkLiINCiAgICAgICA6Oj0geyBzeXN0ZW0gMyB9DQoNCiAgIHN5c0NvbnRhY3QgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBEaXNwbGF5U3RyaW5nIChTSVpFICgwLi4yNTUpKQ0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtd3JpdGUNCiAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdGV4dHVhbCBpZGVudGlmaWNhdGlvbiBvZiB0aGUgY29udGFjdCBwZXJzb24gZm9yDQogICAgICAgICAgICAgICB0aGlzIG1hbmFnZWQgbm9kZSwgdG9nZXRoZXIgd2l0aCBpbmZvcm1hdGlvbiBvbiBob3cNCiAgICAgICAgICAgICAgIHRvIGNvbnRhY3QgdGhpcyBwZXJzb24uICBJZiBubyBjb250YWN0IGluZm9ybWF0aW9uIGlzDQogICAgICAgICAgICAgICBrbm93biwgdGhlIHZhbHVlIGlzIHRoZSB6ZXJvLWxlbmd0aCBzdHJpbmcuIg0KICAgICAgIDo6PSB7IHN5c3RlbSA0IH0NCg0KICAgc3lzTmFtZSBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIERpc3BsYXlTdHJpbmcgKFNJWkUgKDAuLjI1NSkpDQogICAgICAgTUFYLUFDQ0VTUyAgcmVhZC13cml0ZQ0KICAgICAgIFNUQVRVUyAgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIkFuIGFkbWluaXN0cmF0aXZlbHktYXNzaWduZWQgbmFtZSBmb3IgdGhpcyBtYW5hZ2VkDQogICAgICAgICAgICAgICBub2RlLiAgQnkgY29udmVudGlvbiwgdGhpcyBpcyB0aGUgbm9kZSdzIGZ1bGx5LXF1YWxpZmllZA0KICAgICAgICAgICAgICAgZG9tYWluIG5hbWUuICBJZiB0aGUgbmFtZSBpcyB1bmtub3duLCB0aGUgdmFsdWUgaXMNCiAgICAgICAgICAgICAgIHRoZSB6ZXJvLWxlbmd0aCBzdHJpbmcuIg0KICAgICAgIDo6PSB7IHN5c3RlbSA1IH0NCg0KICAgc3lzTG9jYXRpb24gT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBEaXNwbGF5U3RyaW5nIChTSVpFICgwLi4yNTUpKQ0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtd3JpdGUNCiAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgcGh5c2ljYWwgbG9jYXRpb24gb2YgdGhpcyBub2RlIChlLmcuLCAndGVsZXBob25lDQogICAgICAgICAgICAgICBjbG9zZXQsIDNyZCBmbG9vcicpLiAgSWYgdGhlIGxvY2F0aW9uIGlzIHVua25vd24sIHRoZQ0KICAgICAgICAgICAgICAgdmFsdWUgaXMgdGhlIHplcm8tbGVuZ3RoIHN0cmluZy4iDQogICAgICAgOjo9IHsgc3lzdGVtIDYgfQ0KDQogICBzeXNTZXJ2aWNlcyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIElOVEVHRVIgKDAuLjEyNykNCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJBIHZhbHVlIHdoaWNoIGluZGljYXRlcyB0aGUgc2V0IG9mIHNlcnZpY2VzIHRoYXQgdGhpcw0KICAgICAgICAgICAgICAgZW50aXR5IG1heSBwb3RlbnRpYWxseSBvZmZlci4gIFRoZSB2YWx1ZSBpcyBhIHN1bS4NCiAgICAgICAgICAgICAgIFRoaXMgc3VtIGluaXRpYWxseSB0YWtlcyB0aGUgdmFsdWUgemVyby4gVGhlbiwgZm9yDQogICAgICAgICAgICAgICBlYWNoIGxheWVyLCBMLCBpbiB0aGUgcmFuZ2UgMSB0aHJvdWdoIDcsIHRoYXQgdGhpcyBub2RlDQogICAgICAgICAgICAgICBwZXJmb3JtcyB0cmFuc2FjdGlvbnMgZm9yLCAyIHJhaXNlZCB0byAoTCAtIDEpIGlzIGFkZGVkDQogICAgICAgICAgICAgICB0byB0aGUgc3VtLiAgRm9yIGV4YW1wbGUsIGEgbm9kZSB3aGljaCBwZXJmb3JtcyBvbmx5DQogICAgICAgICAgICAgICByb3V0aW5nIGZ1bmN0aW9ucyB3b3VsZCBoYXZlIGEgdmFsdWUgb2YgNCAoMl4oMy0xKSkuDQogICAgICAgICAgICAgICBJbiBjb250cmFzdCwgYSBub2RlIHdoaWNoIGlzIGEgaG9zdCBvZmZlcmluZyBhcHBsaWNhdGlvbg0KICAgICAgICAgICAgICAgc2VydmljZXMgd291bGQgaGF2ZSBhIHZhbHVlIG9mIDcyICgyXig0LTEpICsgMl4oNy0xKSkuDQogICAgICAgICAgICAgICBOb3RlIHRoYXQgaW4gdGhlIGNvbnRleHQgb2YgdGhlIEludGVybmV0IHN1aXRlIG9mDQogICAgICAgICAgICAgICBwcm90b2NvbHMsIHZhbHVlcyBzaG91bGQgYmUgY2FsY3VsYXRlZCBhY2NvcmRpbmdseToNCg0KICAgICAgICAgICAgICAgICAgICBsYXllciAgICAgIGZ1bmN0aW9uYWxpdHkNCiAgICAgICAgICAgICAgICAgICAgICAxICAgICAgICBwaHlzaWNhbCAoZS5nLiwgcmVwZWF0ZXJzKQ0KICAgICAgICAgICAgICAgICAgICAgIDIgICAgICAgIGRhdGFsaW5rL3N1Ym5ldHdvcmsgKGUuZy4sIGJyaWRnZXMpDQogICAgICAgICAgICAgICAgICAgICAgMyAgICAgICAgaW50ZXJuZXQgKGUuZy4sIHN1cHBvcnRzIHRoZSBJUCkNCiAgICAgICAgICAgICAgICAgICAgICA0ICAgICAgICBlbmQtdG8tZW5kICAoZS5nLiwgc3VwcG9ydHMgdGhlIFRDUCkNCiAgICAgICAgICAgICAgICAgICAgICA3ICAgICAgICBhcHBsaWNhdGlvbnMgKGUuZy4sIHN1cHBvcnRzIHRoZSBTTVRQKQ0KDQogICAgICAgICAgICAgICBGb3Igc3lzdGVtcyBpbmNsdWRpbmcgT1NJIHByb3RvY29scywgbGF5ZXJzIDUgYW5kIDYNCiAgICAgICAgICAgICAgIG1heSBhbHNvIGJlIGNvdW50ZWQuIg0KICAgICAgIDo6PSB7IHN5c3RlbSA3IH0NCg0KICAgLS0gb2JqZWN0IHJlc291cmNlIGluZm9ybWF0aW9uDQogICAtLQ0KICAgLS0gYSBjb2xsZWN0aW9uIG9mIG9iamVjdHMgd2hpY2ggZGVzY3JpYmUgdGhlIFNOTVAgZW50aXR5J3MNCiAgIC0tIChzdGF0aWNhbGx5IGFuZCBkeW5hbWljYWxseSBjb25maWd1cmFibGUpIHN1cHBvcnQgb2YNCiAgIC0tIHZhcmlvdXMgTUlCIG1vZHVsZXMuDQoNCiAgIHN5c09STGFzdENoYW5nZSBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgVGltZVN0YW1wDQogICAgICAgTUFYLUFDQ0VTUyByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB2YWx1ZSBvZiBzeXNVcFRpbWUgYXQgdGhlIHRpbWUgb2YgdGhlIG1vc3QgcmVjZW50DQogICAgICAgICAgICAgICBjaGFuZ2UgaW4gc3RhdGUgb3IgdmFsdWUgb2YgYW55IGluc3RhbmNlIG9mIHN5c09SSUQuIg0KICAgICAgIDo6PSB7IHN5c3RlbSA4IH0NCg0KICAgc3lzT1JUYWJsZSBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgU0VRVUVOQ0UgT0YgU3lzT1JFbnRyeQ0KICAgICAgIE1BWC1BQ0NFU1Mgbm90LWFjY2Vzc2libGUNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSAoY29uY2VwdHVhbCkgdGFibGUgbGlzdGluZyB0aGUgY2FwYWJpbGl0aWVzIG9mDQogICAgICAgICAgICAgICB0aGUgbG9jYWwgU05NUCBhcHBsaWNhdGlvbiBhY3RpbmcgYXMgYSBjb21tYW5kDQogICAgICAgICAgICAgICByZXNwb25kZXIgd2l0aCByZXNwZWN0IHRvIHZhcmlvdXMgTUlCIG1vZHVsZXMuDQogICAgICAgICAgICAgICBTTk1QIGVudGl0aWVzIGhhdmluZyBkeW5hbWljYWxseS1jb25maWd1cmFibGUgc3VwcG9ydA0KICAgICAgICAgICAgICAgb2YgTUlCIG1vZHVsZXMgd2lsbCBoYXZlIGEgZHluYW1pY2FsbHktdmFyeWluZyBudW1iZXINCiAgICAgICAgICAgICAgIG9mIGNvbmNlcHR1YWwgcm93cy4iDQogICAgICAgOjo9IHsgc3lzdGVtIDkgfQ0KDQogICBzeXNPUkVudHJ5IE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICBTeXNPUkVudHJ5DQogICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgIFNUQVRVUyAgICAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiQW4gZW50cnkgKGNvbmNlcHR1YWwgcm93KSBpbiB0aGUgc3lzT1JUYWJsZS4iDQogICAgICAgSU5ERVggICAgICB7IHN5c09SSW5kZXggfQ0KICAgICAgIDo6PSB7IHN5c09SVGFibGUgMSB9DQoNCiAgIFN5c09SRW50cnkgOjo9IFNFUVVFTkNFIHsNCiAgICAgICBzeXNPUkluZGV4ICAgICBJTlRFR0VSLA0KICAgICAgIHN5c09SSUQgICAgICAgIE9CSkVDVCBJREVOVElGSUVSLA0KICAgICAgIHN5c09SRGVzY3IgICAgIERpc3BsYXlTdHJpbmcsDQogICAgICAgc3lzT1JVcFRpbWUgICAgVGltZVN0YW1wDQogICB9DQoNCiAgIHN5c09SSW5kZXggT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIElOVEVHRVIgKDEuLjIxNDc0ODM2NDcpDQogICAgICAgTUFYLUFDQ0VTUyBub3QtYWNjZXNzaWJsZQ0KICAgICAgIFNUQVRVUyAgICAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIGF1eGlsaWFyeSB2YXJpYWJsZSB1c2VkIGZvciBpZGVudGlmeWluZyBpbnN0YW5jZXMNCiAgICAgICAgICAgICAgIG9mIHRoZSBjb2x1bW5hciBvYmplY3RzIGluIHRoZSBzeXNPUlRhYmxlLiINCiAgICAgICA6Oj0geyBzeXNPUkVudHJ5IDEgfQ0KDQogICBzeXNPUklEIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICBPQkpFQ1QgSURFTlRJRklFUg0KICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJBbiBhdXRob3JpdGF0aXZlIGlkZW50aWZpY2F0aW9uIG9mIGEgY2FwYWJpbGl0aWVzDQogICAgICAgICAgICAgICBzdGF0ZW1lbnQgd2l0aCByZXNwZWN0IHRvIHZhcmlvdXMgTUlCIG1vZHVsZXMgc3VwcG9ydGVkDQogICAgICAgICAgICAgICBieSB0aGUgbG9jYWwgU05NUCBhcHBsaWNhdGlvbiBhY3RpbmcgYXMgYSBjb21tYW5kDQogICAgICAgICAgICAgICByZXNwb25kZXIuIg0KICAgICAgIDo6PSB7IHN5c09SRW50cnkgMiB9DQoNCiAgIHN5c09SRGVzY3IgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIERpc3BsYXlTdHJpbmcNCiAgICAgICBNQVgtQUNDRVNTIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiQSB0ZXh0dWFsIGRlc2NyaXB0aW9uIG9mIHRoZSBjYXBhYmlsaXRpZXMgaWRlbnRpZmllZA0KICAgICAgICAgICAgICAgYnkgdGhlIGNvcnJlc3BvbmRpbmcgaW5zdGFuY2Ugb2Ygc3lzT1JJRC4iDQogICAgICAgOjo9IHsgc3lzT1JFbnRyeSAzIH0NCg0KICAgc3lzT1JVcFRpbWUgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIFRpbWVTdGFtcA0KICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdmFsdWUgb2Ygc3lzVXBUaW1lIGF0IHRoZSB0aW1lIHRoaXMgY29uY2VwdHVhbA0KICAgICAgICAgICAgICAgcm93IHdhcyBsYXN0IGluc3RhbnRpYXRlZC4iDQogICAgICAgOjo9IHsgc3lzT1JFbnRyeSA0IH0NCg0KICAgLS0gdGhlIFNOTVAgZ3JvdXANCiAgIC0tDQogICAtLSBhIGNvbGxlY3Rpb24gb2Ygb2JqZWN0cyBwcm92aWRpbmcgYmFzaWMgaW5zdHJ1bWVudGF0aW9uIGFuZA0KICAgLS0gY29udHJvbCBvZiBhbiBTTk1QIGVudGl0eS4NCg0KICAgc25tcCAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgbWliLTIgMTEgfQ0KDQogICBzbm1wSW5Qa3RzIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBtZXNzYWdlcyBkZWxpdmVyZWQgdG8gdGhlIFNOTVANCiAgICAgICAgICAgICAgIGVudGl0eSBmcm9tIHRoZSB0cmFuc3BvcnQgc2VydmljZS4iDQogICAgICAgOjo9IHsgc25tcCAxIH0NCg0KICAgc25tcEluQmFkVmVyc2lvbnMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdG90YWwgbnVtYmVyIG9mIFNOTVAgbWVzc2FnZXMgd2hpY2ggd2VyZSBkZWxpdmVyZWQNCiAgICAgICAgICAgICAgIHRvIHRoZSBTTk1QIGVudGl0eSBhbmQgd2VyZSBmb3IgYW4gdW5zdXBwb3J0ZWQgU05NUA0KICAgICAgICAgICAgICAgdmVyc2lvbi4iDQogICAgICAgOjo9IHsgc25tcCAzIH0NCg0KICAgc25tcEluQmFkQ29tbXVuaXR5TmFtZXMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgY29tbXVuaXR5LWJhc2VkIFNOTVAgbWVzc2FnZXMgKGZvcg0KICAgICAgICAgICAgICBleGFtcGxlLCAgU05NUHYxKSBkZWxpdmVyZWQgdG8gdGhlIFNOTVAgZW50aXR5IHdoaWNoDQogICAgICAgICAgICAgIHVzZWQgYW4gU05NUCBjb21tdW5pdHkgbmFtZSBub3Qga25vd24gdG8gc2FpZCBlbnRpdHkuDQogICAgICAgICAgICAgIEFsc28sIGltcGxlbWVudGF0aW9ucyB3aGljaCBhdXRoZW50aWNhdGUgY29tbXVuaXR5LWJhc2VkDQogICAgICAgICAgICAgIFNOTVAgbWVzc2FnZXMgdXNpbmcgY2hlY2socykgaW4gYWRkaXRpb24gdG8gbWF0Y2hpbmcNCiAgICAgICAgICAgICAgdGhlIGNvbW11bml0eSBuYW1lIChmb3IgZXhhbXBsZSwgYnkgYWxzbyBjaGVja2luZw0KICAgICAgICAgICAgICB3aGV0aGVyIHRoZSBtZXNzYWdlIG9yaWdpbmF0ZWQgZnJvbSBhIHRyYW5zcG9ydCBhZGRyZXNzDQogICAgICAgICAgICAgIGFsbG93ZWQgdG8gdXNlIGEgc3BlY2lmaWVkIGNvbW11bml0eSBuYW1lKSBNQVkgaW5jbHVkZQ0KICAgICAgICAgICAgICBpbiB0aGlzIHZhbHVlIHRoZSBudW1iZXIgb2YgbWVzc2FnZXMgd2hpY2ggZmFpbGVkIHRoZQ0KICAgICAgICAgICAgICBhZGRpdGlvbmFsIGNoZWNrKHMpLiAgSXQgaXMgc3Ryb25nbHkgUkVDT01NRU5ERUQgdGhhdA0KICAgICAgICAgICAgICB0aGUgZG9jdW1lbnRhdGlvbiBmb3IgYW55IHNlY3VyaXR5IG1vZGVsIHdoaWNoIGlzIHVzZWQNCiAgICAgICAgICAgICAgdG8gYXV0aGVudGljYXRlIGNvbW11bml0eS1iYXNlZCBTTk1QIG1lc3NhZ2VzIHNwZWNpZnkNCiAgICAgICAgICAgICAgdGhlIHByZWNpc2UgY29uZGl0aW9ucyB0aGF0IGNvbnRyaWJ1dGUgdG8gdGhpcyB2YWx1ZS4iDQogICAgICAgOjo9IHsgc25tcCA0IH0NCg0KICAgc25tcEluQmFkQ29tbXVuaXR5VXNlcyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgQ291bnRlcjMyDQogICAgICAgTUFYLUFDQ0VTUyByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBjb21tdW5pdHktYmFzZWQgU05NUCBtZXNzYWdlcyAoZm9yDQogICAgICAgICAgICAgIGV4YW1wbGUsIFNOTVB2MSkgZGVsaXZlcmVkIHRvIHRoZSBTTk1QIGVudGl0eSB3aGljaA0KICAgICAgICAgICAgICByZXByZXNlbnRlZCBhbiBTTk1QIG9wZXJhdGlvbiB0aGF0IHdhcyBub3QgYWxsb3dlZCBmb3INCiAgICAgICAgICAgICAgdGhlIFNOTVAgY29tbXVuaXR5IG5hbWVkIGluIHRoZSBtZXNzYWdlLiAgVGhlIHByZWNpc2UNCiAgICAgICAgICAgICAgY29uZGl0aW9ucyB1bmRlciB3aGljaCB0aGlzIGNvdW50ZXIgaXMgaW5jcmVtZW50ZWQNCiAgICAgICAgICAgICAgKGlmIGF0IGFsbCkgZGVwZW5kIG9uIGhvdyB0aGUgU05NUCBlbnRpdHkgaW1wbGVtZW50cw0KICAgICAgICAgICAgICBpdHMgYWNjZXNzIGNvbnRyb2wgbWVjaGFuaXNtIGFuZCBob3cgaXRzIGFwcGxpY2F0aW9ucw0KICAgICAgICAgICAgICBpbnRlcmFjdCB3aXRoIHRoYXQgYWNjZXNzIGNvbnRyb2wgbWVjaGFuaXNtLiAgSXQgaXMNCiAgICAgICAgICAgICAgc3Ryb25nbHkgUkVDT01NRU5ERUQgdGhhdCB0aGUgZG9jdW1lbnRhdGlvbiBmb3IgYW55DQogICAgICAgICAgICAgIGFjY2VzcyBjb250cm9sIG1lY2hhbmlzbSB3aGljaCBpcyB1c2VkIHRvIGNvbnRyb2wgYWNjZXNzDQogICAgICAgICAgICAgIHRvIGFuZCB2aXNpYmlsaXR5IG9mIE1JQiBpbnN0cnVtZW50YXRpb24gc3BlY2lmeSB0aGUNCiAgICAgICAgICAgICAgcHJlY2lzZSBjb25kaXRpb25zIHRoYXQgY29udHJpYnV0ZSB0byB0aGlzIHZhbHVlLiINCiAgICAgICA6Oj0geyBzbm1wIDUgfQ0KDQogICBzbm1wSW5BU05QYXJzZUVycnMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdG90YWwgbnVtYmVyIG9mIEFTTi4xIG9yIEJFUiBlcnJvcnMgZW5jb3VudGVyZWQgYnkNCiAgICAgICAgICAgICAgIHRoZSBTTk1QIGVudGl0eSB3aGVuIGRlY29kaW5nIHJlY2VpdmVkIFNOTVAgbWVzc2FnZXMuIg0KICAgICAgIDo6PSB7IHNubXAgNiB9DQoNCiAgIHNubXBFbmFibGVBdXRoZW5UcmFwcyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIElOVEVHRVIgeyBlbmFibGVkKDEpLCBkaXNhYmxlZCgyKSB9DQogICAgICAgTUFYLUFDQ0VTUyAgcmVhZC13cml0ZQ0KICAgICAgIFNUQVRVUyAgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIkluZGljYXRlcyB3aGV0aGVyIHRoZSBTTk1QIGVudGl0eSBpcyBwZXJtaXR0ZWQgdG8NCiAgICAgICAgICAgICAgIGdlbmVyYXRlIGF1dGhlbnRpY2F0aW9uRmFpbHVyZSB0cmFwcy4gIFRoZSB2YWx1ZSBvZiB0aGlzDQogICAgICAgICAgICAgICBvYmplY3Qgb3ZlcnJpZGVzIGFueSBjb25maWd1cmF0aW9uIGluZm9ybWF0aW9uOyBhcyBzdWNoLA0KICAgICAgICAgICAgICAgaXQgcHJvdmlkZXMgYSBtZWFucyB3aGVyZWJ5IGFsbCBhdXRoZW50aWNhdGlvbkZhaWx1cmUNCiAgICAgICAgICAgICAgIHRyYXBzIG1heSBiZSBkaXNhYmxlZC4NCg0KICAgICAgICAgICAgICAgTm90ZSB0aGF0IGl0IGlzIHN0cm9uZ2x5IHJlY29tbWVuZGVkIHRoYXQgdGhpcyBvYmplY3QNCiAgICAgICAgICAgICAgIGJlIHN0b3JlZCBpbiBub24tdm9sYXRpbGUgbWVtb3J5IHNvIHRoYXQgaXQgcmVtYWlucw0KICAgICAgICAgICAgICAgY29uc3RhbnQgYWNyb3NzIHJlLWluaXRpYWxpemF0aW9ucyBvZiB0aGUgbmV0d29yaw0KICAgICAgICAgICAgICAgbWFuYWdlbWVudCBzeXN0ZW0uIg0KICAgICAgIDo6PSB7IHNubXAgMzAgfQ0KDQogICBzbm1wU2lsZW50RHJvcHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgQ29uZmlybWVkIENsYXNzIFBEVXMgKHN1Y2ggYXMNCiAgICAgICAgICAgICAgR2V0UmVxdWVzdC1QRFVzLCBHZXROZXh0UmVxdWVzdC1QRFVzLA0KICAgICAgICAgICAgICBHZXRCdWxrUmVxdWVzdC1QRFVzLCBTZXRSZXF1ZXN0LVBEVXMsIGFuZA0KICAgICAgICAgICAgICBJbmZvcm1SZXF1ZXN0LVBEVXMpIGRlbGl2ZXJlZCB0byB0aGUgU05NUCBlbnRpdHkgd2hpY2gNCiAgICAgICAgICAgICAgd2VyZSBzaWxlbnRseSBkcm9wcGVkIGJlY2F1c2UgdGhlIHNpemUgb2YgYSByZXBseQ0KICAgICAgICAgICAgICBjb250YWluaW5nIGFuIGFsdGVybmF0ZSBSZXNwb25zZSBDbGFzcyBQRFUgKHN1Y2ggYXMgYQ0KICAgICAgICAgICAgICBSZXNwb25zZS1QRFUpIHdpdGggYW4gZW1wdHkgdmFyaWFibGUtYmluZGluZ3MgZmllbGQNCiAgICAgICAgICAgICAgd2FzIGdyZWF0ZXIgdGhhbiBlaXRoZXIgYSBsb2NhbCBjb25zdHJhaW50IG9yIHRoZQ0KICAgICAgICAgICAgICBtYXhpbXVtIG1lc3NhZ2Ugc2l6ZSBhc3NvY2lhdGVkIHdpdGggdGhlIG9yaWdpbmF0b3Igb2YNCiAgICAgICAgICAgICAgdGhlIHJlcXVlc3QuIg0KICAgICAgIDo6PSB7IHNubXAgMzEgfQ0KDQogICBzbm1wUHJveHlEcm9wcyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgQ291bnRlcjMyDQogICAgICAgTUFYLUFDQ0VTUyByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgQ29uZmlybWVkIENsYXNzIFBEVXMNCiAgICAgICAgICAgICAgIChzdWNoIGFzIEdldFJlcXVlc3QtUERVcywgR2V0TmV4dFJlcXVlc3QtUERVcywNCiAgICAgICAgICAgICAgIEdldEJ1bGtSZXF1ZXN0LVBEVXMsIFNldFJlcXVlc3QtUERVcywgYW5kDQogICAgICAgICAgICAgICBJbmZvcm1SZXF1ZXN0LVBEVXMpIGRlbGl2ZXJlZCB0byB0aGUgU05NUCBlbnRpdHkgd2hpY2gNCiAgICAgICAgICAgICAgIHdlcmUgc2lsZW50bHkgZHJvcHBlZCBiZWNhdXNlIHRoZSB0cmFuc21pc3Npb24gb2YNCiAgICAgICAgICAgICAgIHRoZSAocG9zc2libHkgdHJhbnNsYXRlZCkgbWVzc2FnZSB0byBhIHByb3h5IHRhcmdldA0KICAgICAgICAgICAgICAgZmFpbGVkIGluIGEgbWFubmVyIChvdGhlciB0aGFuIGEgdGltZS1vdXQpIHN1Y2ggdGhhdA0KICAgICAgICAgICAgICAgbm8gUmVzcG9uc2UgQ2xhc3MgUERVIChzdWNoIGFzIGEgUmVzcG9uc2UtUERVKSBjb3VsZA0KICAgICAgICAgICAgICAgYmUgcmV0dXJuZWQuIg0KICAgICAgIDo6PSB7IHNubXAgMzIgfQ0KDQogICAtLSBpbmZvcm1hdGlvbiBmb3Igbm90aWZpY2F0aW9ucw0KICAgLS0NCiAgIC0tIGEgY29sbGVjdGlvbiBvZiBvYmplY3RzIHdoaWNoIGFsbG93IHRoZSBTTk1QIGVudGl0eSwgd2hlbg0KICAgLS0gc3VwcG9ydGluZyBhIG5vdGlmaWNhdGlvbiBvcmlnaW5hdG9yIGFwcGxpY2F0aW9uLA0KICAgLS0gdG8gYmUgY29uZmlndXJlZCB0byBnZW5lcmF0ZSBTTk1QdjItVHJhcC1QRFVzLg0KDQogICBzbm1wVHJhcCAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBzbm1wTUlCT2JqZWN0cyA0IH0NCg0KICAgc25tcFRyYXBPSUQgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIE9CSkVDVCBJREVOVElGSUVSDQogICAgICAgTUFYLUFDQ0VTUyBhY2Nlc3NpYmxlLWZvci1ub3RpZnkNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSBhdXRob3JpdGF0aXZlIGlkZW50aWZpY2F0aW9uIG9mIHRoZSBub3RpZmljYXRpb24NCiAgICAgICAgICAgICAgIGN1cnJlbnRseSBiZWluZyBzZW50LiAgVGhpcyB2YXJpYWJsZSBvY2N1cnMgYXMNCiAgICAgICAgICAgICAgIHRoZSBzZWNvbmQgdmFyYmluZCBpbiBldmVyeSBTTk1QdjItVHJhcC1QRFUgYW5kDQogICAgICAgICAgICAgICBJbmZvcm1SZXF1ZXN0LVBEVS4iDQogICAgICAgOjo9IHsgc25tcFRyYXAgMSB9DQoNCiAgIC0tICA6Oj0geyBzbm1wVHJhcCAyIH0gICB0aGlzIE9JRCBpcyBvYnNvbGV0ZQ0KDQogICBzbm1wVHJhcEVudGVycHJpc2UgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgIE9CSkVDVCBJREVOVElGSUVSDQogICAgICAgTUFYLUFDQ0VTUyBhY2Nlc3NpYmxlLWZvci1ub3RpZnkNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSBhdXRob3JpdGF0aXZlIGlkZW50aWZpY2F0aW9uIG9mIHRoZSBlbnRlcnByaXNlDQogICAgICAgICAgICAgICBhc3NvY2lhdGVkIHdpdGggdGhlIHRyYXAgY3VycmVudGx5IGJlaW5nIHNlbnQuICBXaGVuIGFuDQogICAgICAgICAgICAgICBTTk1QIHByb3h5IGFnZW50IGlzIG1hcHBpbmcgYW4gUkZDMTE1NyBUcmFwLVBEVQ0KICAgICAgICAgICAgICAgaW50byBhIFNOTVB2Mi1UcmFwLVBEVSwgdGhpcyB2YXJpYWJsZSBvY2N1cnMgYXMgdGhlDQogICAgICAgICAgICAgICBsYXN0IHZhcmJpbmQuIg0KICAgICAgIDo6PSB7IHNubXBUcmFwIDMgfQ0KDQogICAtLSAgOjo9IHsgc25tcFRyYXAgNCB9ICAgdGhpcyBPSUQgaXMgb2Jzb2xldGUNCg0KICAgLS0gd2VsbC1rbm93biB0cmFwcw0KDQogICBzbm1wVHJhcHMgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBzbm1wTUlCT2JqZWN0cyA1IH0NCg0KICAgY29sZFN0YXJ0IE5PVElGSUNBVElPTi1UWVBFDQogICAgICAgU1RBVFVTICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJBIGNvbGRTdGFydCB0cmFwIHNpZ25pZmllcyB0aGF0IHRoZSBTTk1QIGVudGl0eSwNCiAgICAgICAgICAgICAgIHN1cHBvcnRpbmcgYSBub3RpZmljYXRpb24gb3JpZ2luYXRvciBhcHBsaWNhdGlvbiwgaXMNCiAgICAgICAgICAgICAgIHJlaW5pdGlhbGl6aW5nIGl0c2VsZiBhbmQgdGhhdCBpdHMgY29uZmlndXJhdGlvbiBtYXkNCiAgICAgICAgICAgICAgIGhhdmUgYmVlbiBhbHRlcmVkLiINCiAgICAgICA6Oj0geyBzbm1wVHJhcHMgMSB9DQoNCiAgIHdhcm1TdGFydCBOT1RJRklDQVRJT04tVFlQRQ0KICAgICAgIFNUQVRVUyAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiQSB3YXJtU3RhcnQgdHJhcCBzaWduaWZpZXMgdGhhdCB0aGUgU05NUCBlbnRpdHksDQogICAgICAgICAgICAgICBzdXBwb3J0aW5nIGEgbm90aWZpY2F0aW9uIG9yaWdpbmF0b3IgYXBwbGljYXRpb24sDQogICAgICAgICAgICAgICBpcyByZWluaXRpYWxpemluZyBpdHNlbGYgc3VjaCB0aGF0IGl0cyBjb25maWd1cmF0aW9uDQogICAgICAgICAgICAgICBpcyB1bmFsdGVyZWQuIg0KICAgICAgIDo6PSB7IHNubXBUcmFwcyAyIH0NCg0KICAgLS0gTm90ZSB0aGUgbGlua0Rvd24gTk9USUZJQ0FUSU9OLVRZUEUgOjo9IHsgc25tcFRyYXBzIDMgfQ0KICAgLS0gYW5kIHRoZSBsaW5rVXAgTk9USUZJQ0FUSU9OLVRZUEUgOjo9IHsgc25tcFRyYXBzIDQgfQ0KICAgLS0gYXJlIGRlZmluZWQgaW4gUkZDIDI4NjMgW1JGQzI4NjNdDQoNCiAgIGF1dGhlbnRpY2F0aW9uRmFpbHVyZSBOT1RJRklDQVRJT04tVFlQRQ0KICAgICAgIFNUQVRVUyAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiQW4gYXV0aGVudGljYXRpb25GYWlsdXJlIHRyYXAgc2lnbmlmaWVzIHRoYXQgdGhlIFNOTVANCiAgICAgICAgICAgICAgICBlbnRpdHkgaGFzIHJlY2VpdmVkIGEgcHJvdG9jb2wgbWVzc2FnZSB0aGF0IGlzIG5vdA0KICAgICAgICAgICAgICAgIHByb3Blcmx5IGF1dGhlbnRpY2F0ZWQuICBXaGlsZSBhbGwgaW1wbGVtZW50YXRpb25zDQogICAgICAgICAgICAgICAgb2YgU05NUCBlbnRpdGllcyBNQVkgYmUgY2FwYWJsZSBvZiBnZW5lcmF0aW5nIHRoaXMNCiAgICAgICAgICAgICAgICB0cmFwLCB0aGUgc25tcEVuYWJsZUF1dGhlblRyYXBzIG9iamVjdCBpbmRpY2F0ZXMNCiAgICAgICAgICAgICAgICB3aGV0aGVyIHRoaXMgdHJhcCB3aWxsIGJlIGdlbmVyYXRlZC4iDQogICAgICAgOjo9IHsgc25tcFRyYXBzIDUgfQ0KDQogICAtLSBOb3RlIHRoZSBlZ3BOZWlnaGJvckxvc3Mgbm90aWZpY2F0aW9uIGlzIGRlZmluZWQNCiAgIC0tIGFzIHsgc25tcFRyYXBzIDYgfSBpbiBSRkMgMTIxMw0KDQogICAtLSB0aGUgc2V0IGdyb3VwDQogICAtLQ0KICAgLS0gYSBjb2xsZWN0aW9uIG9mIG9iamVjdHMgd2hpY2ggYWxsb3cgc2V2ZXJhbCBjb29wZXJhdGluZw0KICAgLS0gY29tbWFuZCBnZW5lcmF0b3IgYXBwbGljYXRpb25zIHRvIGNvb3JkaW5hdGUgdGhlaXIgdXNlIG9mIHRoZQ0KICAgLS0gc2V0IG9wZXJhdGlvbi4NCg0KICAgc25tcFNldCAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgc25tcE1JQk9iamVjdHMgNiB9DQoNCiAgIHNubXBTZXRTZXJpYWxObyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgVGVzdEFuZEluY3INCiAgICAgICBNQVgtQUNDRVNTIHJlYWQtd3JpdGUNCiAgICAgICBTVEFUVVMgICAgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIkFuIGFkdmlzb3J5IGxvY2sgdXNlZCB0byBhbGxvdyBzZXZlcmFsIGNvb3BlcmF0aW5nDQogICAgICAgICAgICAgICBjb21tYW5kIGdlbmVyYXRvciBhcHBsaWNhdGlvbnMgdG8gY29vcmRpbmF0ZSB0aGVpcg0KICAgICAgICAgICAgICAgdXNlIG9mIHRoZSBTTk1QIHNldCBvcGVyYXRpb24uDQoNCiAgICAgICAgICAgICAgIFRoaXMgb2JqZWN0IGlzIHVzZWQgZm9yIGNvYXJzZS1ncmFpbiBjb29yZGluYXRpb24uDQogICAgICAgICAgICAgICBUbyBhY2hpZXZlIGZpbmUtZ3JhaW4gY29vcmRpbmF0aW9uLCBvbmUgb3IgbW9yZSBzaW1pbGFyDQogICAgICAgICAgICAgICBvYmplY3RzIG1pZ2h0IGJlIGRlZmluZWQgd2l0aGluIGVhY2ggTUlCIGdyb3VwLCBhcw0KICAgICAgICAgICAgICAgYXBwcm9wcmlhdGUuIg0KICAgICAgIDo6PSB7IHNubXBTZXQgMSB9DQoNCiAgIC0tIGNvbmZvcm1hbmNlIGluZm9ybWF0aW9uDQoNCiAgIHNubXBNSUJDb25mb3JtYW5jZQ0KICAgICAgICAgICAgICAgICAgT0JKRUNUIElERU5USUZJRVIgOjo9IHsgc25tcE1JQiAyIH0NCg0KICAgc25tcE1JQkNvbXBsaWFuY2VzDQogICAgICAgICAgICAgICAgICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBzbm1wTUlCQ29uZm9ybWFuY2UgMSB9DQogICBzbm1wTUlCR3JvdXBzICBPQkpFQ1QgSURFTlRJRklFUiA6Oj0geyBzbm1wTUlCQ29uZm9ybWFuY2UgMiB9DQoNCiAgIC0tIGNvbXBsaWFuY2Ugc3RhdGVtZW50cw0KICAgLS0gICAgOjo9IHsgc25tcE1JQkNvbXBsaWFuY2VzIDEgfSAgICAgIHRoaXMgT0lEIGlzIG9ic29sZXRlDQogICBzbm1wQmFzaWNDb21wbGlhbmNlIE1PRFVMRS1DT01QTElBTkNFDQogICAgICAgU1RBVFVTICBkZXByZWNhdGVkDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgY29tcGxpYW5jZSBzdGF0ZW1lbnQgZm9yIFNOTVB2MiBlbnRpdGllcyB3aGljaA0KICAgICAgICAgICAgICAgaW1wbGVtZW50IHRoZSBTTk1QdjIgTUlCLg0KDQogICAgICAgICAgICAgICBUaGlzIGNvbXBsaWFuY2Ugc3RhdGVtZW50IGlzIHJlcGxhY2VkIGJ5DQogICAgICAgICAgICAgICBzbm1wQmFzaWNDb21wbGlhbmNlUmV2Mi4iDQogICAgICAgTU9EVUxFICAtLSB0aGlzIG1vZHVsZQ0KICAgICAgICAgICBNQU5EQVRPUlktR1JPVVBTIHsgc25tcEdyb3VwLCBzbm1wU2V0R3JvdXAsIHN5c3RlbUdyb3VwLA0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgc25tcEJhc2ljTm90aWZpY2F0aW9uc0dyb3VwIH0NCg0KICAgICAgICAgICBHUk9VUCAgIHNubXBDb21tdW5pdHlHcm91cDENCiAgICAgICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGlzIGdyb3VwIGlzIG1hbmRhdG9yeSBmb3IgU05NUHYyIGVudGl0aWVzIHdoaWNoDQogICAgICAgICAgICAgICBzdXBwb3J0IGNvbW11bml0eS1iYXNlZCBhdXRoZW50aWNhdGlvbi4iDQoNCiAgICAgICA6Oj0geyBzbm1wTUlCQ29tcGxpYW5jZXMgMiB9DQoNCiAgIHNubXBCYXNpY0NvbXBsaWFuY2VSZXYyIE1PRFVMRS1DT01QTElBTkNFDQogICAgICAgU1RBVFVTICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgY29tcGxpYW5jZSBzdGF0ZW1lbnQgZm9yIFNOTVAgZW50aXRpZXMgd2hpY2gNCiAgICAgICAgICAgICAgIGltcGxlbWVudCB0aGlzIE1JQiBtb2R1bGUuIg0KICAgICAgIE1PRFVMRSAgLS0gdGhpcyBtb2R1bGUNCiAgICAgICAgICAgTUFOREFUT1JZLUdST1VQUyB7IHNubXBHcm91cCwgc25tcFNldEdyb3VwLCBzeXN0ZW1Hcm91cCwNCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNubXBCYXNpY05vdGlmaWNhdGlvbnNHcm91cCB9DQoNCiAgICAgICAgICAgR1JPVVAgICBzbm1wQ29tbXVuaXR5R3JvdXAxDQogICAgICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhpcyBncm91cCBpcyBtYW5kYXRvcnkgZm9yIFNOTVAgZW50aXRpZXMgd2hpY2gNCiAgICAgICAgICAgICAgIHN1cHBvcnQgY29tbXVuaXR5LWJhc2VkIGF1dGhlbnRpY2F0aW9uLiINCg0KICAgICAgICAgICBHUk9VUCAgIHNubXBXYXJtU3RhcnROb3RpZmljYXRpb25Hcm91cA0KICAgICAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoaXMgZ3JvdXAgaXMgbWFuZGF0b3J5IGZvciBhbiBTTk1QIGVudGl0eSB3aGljaA0KICAgICAgICAgICAgICAgc3VwcG9ydHMgY29tbWFuZCByZXNwb25kZXIgYXBwbGljYXRpb25zLCBhbmQgaXMNCiAgICAgICAgICAgICAgIGFibGUgdG8gcmVpbml0aWFsaXplIGl0c2VsZiBzdWNoIHRoYXQgaXRzDQogICAgICAgICAgICAgICBjb25maWd1cmF0aW9uIGlzIHVuYWx0ZXJlZC4iDQoNCiAgICAgICA6Oj0geyBzbm1wTUlCQ29tcGxpYW5jZXMgMyB9DQoNCiAgIC0tIHVuaXRzIG9mIGNvbmZvcm1hbmNlDQoNCiAgIC0tICA6Oj0geyBzbm1wTUlCR3JvdXBzIDEgfSAgICAgICAgICAgdGhpcyBPSUQgaXMgb2Jzb2xldGUNCiAgIC0tICA6Oj0geyBzbm1wTUlCR3JvdXBzIDIgfSAgICAgICAgICAgdGhpcyBPSUQgaXMgb2Jzb2xldGUNCiAgIC0tICA6Oj0geyBzbm1wTUlCR3JvdXBzIDMgfSAgICAgICAgICAgdGhpcyBPSUQgaXMgb2Jzb2xldGUNCiAgIC0tICA6Oj0geyBzbm1wTUlCR3JvdXBzIDQgfSAgICAgICAgICAgdGhpcyBPSUQgaXMgb2Jzb2xldGUNCg0KICAgc25tcEdyb3VwIE9CSkVDVC1HUk9VUA0KICAgICAgIE9CSkVDVFMgeyBzbm1wSW5Qa3RzLA0KICAgICAgICAgICAgICAgICBzbm1wSW5CYWRWZXJzaW9ucywNCiAgICAgICAgICAgICAgICAgc25tcEluQVNOUGFyc2VFcnJzLA0KICAgICAgICAgICAgICAgICBzbm1wU2lsZW50RHJvcHMsDQogICAgICAgICAgICAgICAgIHNubXBQcm94eURyb3BzLA0KICAgICAgICAgICAgICAgICBzbm1wRW5hYmxlQXV0aGVuVHJhcHMgfQ0KICAgICAgIFNUQVRVUyAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiQSBjb2xsZWN0aW9uIG9mIG9iamVjdHMgcHJvdmlkaW5nIGJhc2ljIGluc3RydW1lbnRhdGlvbg0KICAgICAgICAgICAgICAgYW5kIGNvbnRyb2wgb2YgYW4gU05NUCBlbnRpdHkuIg0KICAgICAgIDo6PSB7IHNubXBNSUJHcm91cHMgOCB9DQoNCiAgIHNubXBDb21tdW5pdHlHcm91cDEgT0JKRUNULUdST1VQDQogICAgICAgT0JKRUNUUyB7IHNubXBJbkJhZENvbW11bml0eU5hbWVzLA0KICAgICAgICAgICAgICAgICBzbm1wSW5CYWRDb21tdW5pdHlVc2VzIH0NCiAgICAgICBTVEFUVVMgIGN1cnJlbnQNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIkEgY29sbGVjdGlvbiBvZiBvYmplY3RzIHByb3ZpZGluZyBiYXNpYyBpbnN0cnVtZW50YXRpb24NCiAgICAgICAgICAgICAgIG9mIGEgU05NUCBlbnRpdHkgd2hpY2ggc3VwcG9ydHMgY29tbXVuaXR5LWJhc2VkDQogICAgICAgICAgICAgICBhdXRoZW50aWNhdGlvbi4iDQogICAgICAgOjo9IHsgc25tcE1JQkdyb3VwcyA5IH0NCg0KICAgc25tcFNldEdyb3VwIE9CSkVDVC1HUk9VUA0KICAgICAgIE9CSkVDVFMgeyBzbm1wU2V0U2VyaWFsTm8gfQ0KICAgICAgIFNUQVRVUyAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiQSBjb2xsZWN0aW9uIG9mIG9iamVjdHMgd2hpY2ggYWxsb3cgc2V2ZXJhbCBjb29wZXJhdGluZw0KICAgICAgICAgICAgICAgY29tbWFuZCBnZW5lcmF0b3IgYXBwbGljYXRpb25zIHRvIGNvb3JkaW5hdGUgdGhlaXINCiAgICAgICAgICAgICAgIHVzZSBvZiB0aGUgc2V0IG9wZXJhdGlvbi4iDQogICAgICAgOjo9IHsgc25tcE1JQkdyb3VwcyA1IH0NCg0KICAgc3lzdGVtR3JvdXAgT0JKRUNULUdST1VQDQogICAgICAgT0JKRUNUUyB7IHN5c0Rlc2NyLCBzeXNPYmplY3RJRCwgc3lzVXBUaW1lLA0KICAgICAgICAgICAgICAgICBzeXNDb250YWN0LCBzeXNOYW1lLCBzeXNMb2NhdGlvbiwNCiAgICAgICAgICAgICAgICAgc3lzU2VydmljZXMsDQogICAgICAgICAgICAgICAgIHN5c09STGFzdENoYW5nZSwgc3lzT1JJRCwNCiAgICAgICAgICAgICAgICAgc3lzT1JVcFRpbWUsIHN5c09SRGVzY3IgfQ0KICAgICAgIFNUQVRVUyAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHN5c3RlbSBncm91cCBkZWZpbmVzIG9iamVjdHMgd2hpY2ggYXJlIGNvbW1vbiB0byBhbGwNCiAgICAgICAgICAgICAgIG1hbmFnZWQgc3lzdGVtcy4iDQogICAgICAgOjo9IHsgc25tcE1JQkdyb3VwcyA2IH0NCg0KICAgc25tcEJhc2ljTm90aWZpY2F0aW9uc0dyb3VwIE5PVElGSUNBVElPTi1HUk9VUA0KICAgICAgIE5PVElGSUNBVElPTlMgeyBjb2xkU3RhcnQsIGF1dGhlbnRpY2F0aW9uRmFpbHVyZSB9DQogICAgICAgU1RBVFVTICAgICAgICBjdXJyZW50DQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAiVGhlIGJhc2ljIG5vdGlmaWNhdGlvbnMgaW1wbGVtZW50ZWQgYnkgYW4gU05NUCBlbnRpdHkNCiAgICAgICAgICAgc3VwcG9ydGluZyBjb21tYW5kIHJlc3BvbmRlciBhcHBsaWNhdGlvbnMuIg0KICAgICAgIDo6PSB7IHNubXBNSUJHcm91cHMgNyB9DQoNCiAgIHNubXBXYXJtU3RhcnROb3RpZmljYXRpb25Hcm91cCBOT1RJRklDQVRJT04tR1JPVVANCiAgICAgIE5PVElGSUNBVElPTlMgeyB3YXJtU3RhcnQgfQ0KICAgICAgU1RBVFVTICAgICAgICBjdXJyZW50DQogICAgICBERVNDUklQVElPTg0KICAgICAgICAiQW4gYWRkaXRpb25hbCBub3RpZmljYXRpb24gZm9yIGFuIFNOTVAgZW50aXR5IHN1cHBvcnRpbmcNCiAgICAgICAgY29tbWFuZCByZXNwb25kZXIgYXBwbGljYXRpb25zLCBpZiBpdCBpcyBhYmxlIHRvIHJlaW5pdGlhbGl6ZQ0KICAgICAgICBpdHNlbGYgc3VjaCB0aGF0IGl0cyBjb25maWd1cmF0aW9uIGlzIHVuYWx0ZXJlZC4iDQogICAgIDo6PSB7IHNubXBNSUJHcm91cHMgMTEgfQ0KDQogICBzbm1wTm90aWZpY2F0aW9uR3JvdXAgT0JKRUNULUdST1VQDQogICAgICAgT0JKRUNUUyB7IHNubXBUcmFwT0lELCBzbm1wVHJhcEVudGVycHJpc2UgfQ0KICAgICAgIFNUQVRVUyAgY3VycmVudA0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlc2Ugb2JqZWN0cyBhcmUgcmVxdWlyZWQgZm9yIGVudGl0aWVzDQogICAgICAgICAgICAgICB3aGljaCBzdXBwb3J0IG5vdGlmaWNhdGlvbiBvcmlnaW5hdG9yIGFwcGxpY2F0aW9ucy4iDQogICAgICAgOjo9IHsgc25tcE1JQkdyb3VwcyAxMiB9DQoNCiAgIC0tIGRlZmluaXRpb25zIGluIFJGQyAxMjEzIG1hZGUgb2Jzb2xldGUgYnkgdGhlIGluY2x1c2lvbiBvZiBhDQogICAtLSBzdWJzZXQgb2YgdGhlIHNubXAgZ3JvdXAgaW4gdGhpcyBNSUINCg0KICAgc25tcE91dFBrdHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIE1lc3NhZ2VzIHdoaWNoIHdlcmUNCiAgICAgICAgICAgICAgIHBhc3NlZCBmcm9tIHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSB0byB0aGUNCiAgICAgICAgICAgICAgIHRyYW5zcG9ydCBzZXJ2aWNlLiINCiAgICAgICA6Oj0geyBzbm1wIDIgfQ0KDQogICAtLSB7IHNubXAgNyB9IGlzIG5vdCB1c2VkDQoNCiAgIHNubXBJblRvb0JpZ3MgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFBEVXMgd2hpY2ggd2VyZQ0KICAgICAgICAgICAgICAgZGVsaXZlcmVkIHRvIHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhbmQgZm9yDQogICAgICAgICAgICAgICB3aGljaCB0aGUgdmFsdWUgb2YgdGhlIGVycm9yLXN0YXR1cyBmaWVsZCB3YXMNCiAgICAgICAgICAgICAgIGB0b29CaWcnLiINCiAgICAgICA6Oj0geyBzbm1wIDggfQ0KDQogICBzbm1wSW5Ob1N1Y2hOYW1lcyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgIG9ic29sZXRlDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdG90YWwgbnVtYmVyIG9mIFNOTVAgUERVcyB3aGljaCB3ZXJlDQogICAgICAgICAgICAgICBkZWxpdmVyZWQgdG8gdGhlIFNOTVAgcHJvdG9jb2wgZW50aXR5IGFuZCBmb3INCiAgICAgICAgICAgICAgIHdoaWNoIHRoZSB2YWx1ZSBvZiB0aGUgZXJyb3Itc3RhdHVzIGZpZWxkIHdhcw0KICAgICAgICAgICAgICAgYG5vU3VjaE5hbWUnLiINCiAgICAgICA6Oj0geyBzbm1wIDkgfQ0KDQogICBzbm1wSW5CYWRWYWx1ZXMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFBEVXMgd2hpY2ggd2VyZQ0KICAgICAgICAgICAgICAgZGVsaXZlcmVkIHRvIHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhbmQgZm9yDQogICAgICAgICAgICAgICB3aGljaCB0aGUgdmFsdWUgb2YgdGhlIGVycm9yLXN0YXR1cyBmaWVsZCB3YXMNCiAgICAgICAgICAgICAgIGBiYWRWYWx1ZScuIg0KICAgICAgIDo6PSB7IHNubXAgMTAgfQ0KDQogICBzbm1wSW5SZWFkT25seXMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciB2YWxpZCBTTk1QIFBEVXMgd2hpY2ggd2VyZSBkZWxpdmVyZWQNCiAgICAgICAgICAgICAgIHRvIHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhbmQgZm9yIHdoaWNoIHRoZSB2YWx1ZQ0KICAgICAgICAgICAgICAgb2YgdGhlIGVycm9yLXN0YXR1cyBmaWVsZCB3YXMgYHJlYWRPbmx5Jy4gIEl0IHNob3VsZA0KICAgICAgICAgICAgICAgYmUgbm90ZWQgdGhhdCBpdCBpcyBhIHByb3RvY29sIGVycm9yIHRvIGdlbmVyYXRlIGFuDQogICAgICAgICAgICAgICBTTk1QIFBEVSB3aGljaCBjb250YWlucyB0aGUgdmFsdWUgYHJlYWRPbmx5JyBpbiB0aGUNCiAgICAgICAgICAgICAgIGVycm9yLXN0YXR1cyBmaWVsZCwgYXMgc3VjaCB0aGlzIG9iamVjdCBpcyBwcm92aWRlZA0KICAgICAgICAgICAgICAgYXMgYSBtZWFucyBvZiBkZXRlY3RpbmcgaW5jb3JyZWN0IGltcGxlbWVudGF0aW9ucyBvZg0KICAgICAgICAgICAgICAgdGhlIFNOTVAuIg0KICAgICAgIDo6PSB7IHNubXAgMTEgfQ0KDQogICBzbm1wSW5HZW5FcnJzIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICAgQ291bnRlcjMyDQogICAgICAgTUFYLUFDQ0VTUyAgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICAgb2Jzb2xldGUNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgU05NUCBQRFVzIHdoaWNoIHdlcmUgZGVsaXZlcmVkDQogICAgICAgICAgICAgICB0byB0aGUgU05NUCBwcm90b2NvbCBlbnRpdHkgYW5kIGZvciB3aGljaCB0aGUgdmFsdWUNCiAgICAgICAgICAgICAgIG9mIHRoZSBlcnJvci1zdGF0dXMgZmllbGQgd2FzIGBnZW5FcnInLiINCiAgICAgICA6Oj0geyBzbm1wIDEyIH0NCg0KICAgc25tcEluVG90YWxSZXFWYXJzIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICAgQ291bnRlcjMyDQogICAgICAgTUFYLUFDQ0VTUyAgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICAgb2Jzb2xldGUNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgTUlCIG9iamVjdHMgd2hpY2ggaGF2ZSBiZWVuDQogICAgICAgICAgICAgICByZXRyaWV2ZWQgc3VjY2Vzc2Z1bGx5IGJ5IHRoZSBTTk1QIHByb3RvY29sIGVudGl0eQ0KICAgICAgICAgICAgICAgYXMgdGhlIHJlc3VsdCBvZiByZWNlaXZpbmcgdmFsaWQgU05NUCBHZXQtUmVxdWVzdA0KICAgICAgICAgICAgICAgYW5kIEdldC1OZXh0IFBEVXMuIg0KICAgICAgIDo6PSB7IHNubXAgMTMgfQ0KDQogICBzbm1wSW5Ub3RhbFNldFZhcnMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBNSUIgb2JqZWN0cyB3aGljaCBoYXZlIGJlZW4NCiAgICAgICAgICAgICAgIGFsdGVyZWQgc3VjY2Vzc2Z1bGx5IGJ5IHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhcw0KICAgICAgICAgICAgICAgdGhlIHJlc3VsdCBvZiByZWNlaXZpbmcgdmFsaWQgU05NUCBTZXQtUmVxdWVzdCBQRFVzLiINCiAgICAgICA6Oj0geyBzbm1wIDE0IH0NCg0KICAgc25tcEluR2V0UmVxdWVzdHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIEdldC1SZXF1ZXN0IFBEVXMgd2hpY2gNCiAgICAgICAgICAgICAgIGhhdmUgYmVlbiBhY2NlcHRlZCBhbmQgcHJvY2Vzc2VkIGJ5IHRoZSBTTk1QDQogICAgICAgICAgICAgICBwcm90b2NvbCBlbnRpdHkuIg0KICAgICAgIDo6PSB7IHNubXAgMTUgfQ0KDQogICBzbm1wSW5HZXROZXh0cyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgIG9ic29sZXRlDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdG90YWwgbnVtYmVyIG9mIFNOTVAgR2V0LU5leHQgUERVcyB3aGljaCBoYXZlIGJlZW4NCiAgICAgICAgICAgICAgIGFjY2VwdGVkIGFuZCBwcm9jZXNzZWQgYnkgdGhlIFNOTVAgcHJvdG9jb2wgZW50aXR5LiINCiAgICAgICA6Oj0geyBzbm1wIDE2IH0NCg0KICAgc25tcEluU2V0UmVxdWVzdHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFNldC1SZXF1ZXN0IFBEVXMgd2hpY2gNCiAgICAgICAgICAgICAgIGhhdmUgYmVlbiBhY2NlcHRlZCBhbmQgcHJvY2Vzc2VkIGJ5IHRoZSBTTk1QIHByb3RvY29sDQogICAgICAgICAgICAgICBlbnRpdHkuIg0KICAgICAgIDo6PSB7IHNubXAgMTcgfQ0KDQogICBzbm1wSW5HZXRSZXNwb25zZXMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIEdldC1SZXNwb25zZSBQRFVzIHdoaWNoDQogICAgICAgICAgICAgICBoYXZlIGJlZW4gYWNjZXB0ZWQgYW5kIHByb2Nlc3NlZCBieSB0aGUgU05NUCBwcm90b2NvbA0KICAgICAgICAgICAgICAgZW50aXR5LiINCiAgICAgICA6Oj0geyBzbm1wIDE4IH0NCg0KICAgc25tcEluVHJhcHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFRyYXAgUERVcyB3aGljaCBoYXZlIGJlZW4NCiAgICAgICAgICAgICAgIGFjY2VwdGVkIGFuZCBwcm9jZXNzZWQgYnkgdGhlIFNOTVAgcHJvdG9jb2wgZW50aXR5LiINCiAgICAgICA6Oj0geyBzbm1wIDE5IH0NCg0KICAgc25tcE91dFRvb0JpZ3MgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFBEVXMgd2hpY2ggd2VyZSBnZW5lcmF0ZWQNCiAgICAgICAgICAgICAgIGJ5IHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhbmQgZm9yIHdoaWNoIHRoZSB2YWx1ZQ0KICAgICAgICAgICAgICAgb2YgdGhlIGVycm9yLXN0YXR1cyBmaWVsZCB3YXMgYHRvb0JpZy4nIg0KICAgICAgIDo6PSB7IHNubXAgMjAgfQ0KDQogICBzbm1wT3V0Tm9TdWNoTmFtZXMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFBEVXMgd2hpY2ggd2VyZSBnZW5lcmF0ZWQNCiAgICAgICAgICAgICAgIGJ5IHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhbmQgZm9yIHdoaWNoIHRoZSB2YWx1ZQ0KICAgICAgICAgICAgICAgb2YgdGhlIGVycm9yLXN0YXR1cyB3YXMgYG5vU3VjaE5hbWUnLiINCiAgICAgICA6Oj0geyBzbm1wIDIxIH0NCg0KICAgc25tcE91dEJhZFZhbHVlcyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgIG9ic29sZXRlDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdG90YWwgbnVtYmVyIG9mIFNOTVAgUERVcyB3aGljaCB3ZXJlIGdlbmVyYXRlZA0KICAgICAgICAgICAgICAgYnkgdGhlIFNOTVAgcHJvdG9jb2wgZW50aXR5IGFuZCBmb3Igd2hpY2ggdGhlIHZhbHVlDQogICAgICAgICAgICAgICBvZiB0aGUgZXJyb3Itc3RhdHVzIGZpZWxkIHdhcyBgYmFkVmFsdWUnLiINCiAgICAgICA6Oj0geyBzbm1wIDIyIH0NCg0KICAgLS0geyBzbm1wIDIzIH0gaXMgbm90IHVzZWQNCg0KICAgc25tcE91dEdlbkVycnMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFBEVXMgd2hpY2ggd2VyZSBnZW5lcmF0ZWQNCiAgICAgICAgICAgICAgIGJ5IHRoZSBTTk1QIHByb3RvY29sIGVudGl0eSBhbmQgZm9yIHdoaWNoIHRoZSB2YWx1ZQ0KICAgICAgICAgICAgICAgb2YgdGhlIGVycm9yLXN0YXR1cyBmaWVsZCB3YXMgYGdlbkVycicuIg0KICAgICAgIDo6PSB7IHNubXAgMjQgfQ0KDQogICBzbm1wT3V0R2V0UmVxdWVzdHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIEdldC1SZXF1ZXN0IFBEVXMgd2hpY2gNCiAgICAgICAgICAgICAgIGhhdmUgYmVlbiBnZW5lcmF0ZWQgYnkgdGhlIFNOTVAgcHJvdG9jb2wgZW50aXR5LiINCiAgICAgICA6Oj0geyBzbm1wIDI1IH0NCg0KICAgc25tcE91dEdldE5leHRzIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICAgQ291bnRlcjMyDQogICAgICAgTUFYLUFDQ0VTUyAgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICAgb2Jzb2xldGUNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgU05NUCBHZXQtTmV4dCBQRFVzIHdoaWNoIGhhdmUNCiAgICAgICAgICAgICAgIGJlZW4gZ2VuZXJhdGVkIGJ5IHRoZSBTTk1QIHByb3RvY29sIGVudGl0eS4iDQogICAgICAgOjo9IHsgc25tcCAyNiB9DQoNCiAgIHNubXBPdXRTZXRSZXF1ZXN0cyBPQkpFQ1QtVFlQRQ0KICAgICAgIFNZTlRBWCAgICAgIENvdW50ZXIzMg0KICAgICAgIE1BWC1BQ0NFU1MgIHJlYWQtb25seQ0KICAgICAgIFNUQVRVUyAgICAgIG9ic29sZXRlDQogICAgICAgREVTQ1JJUFRJT04NCiAgICAgICAgICAgICAgICJUaGUgdG90YWwgbnVtYmVyIG9mIFNOTVAgU2V0LVJlcXVlc3QgUERVcyB3aGljaA0KICAgICAgICAgICAgICAgaGF2ZSBiZWVuIGdlbmVyYXRlZCBieSB0aGUgU05NUCBwcm90b2NvbCBlbnRpdHkuIg0KICAgICAgIDo6PSB7IHNubXAgMjcgfQ0KDQogICBzbm1wT3V0R2V0UmVzcG9uc2VzIE9CSkVDVC1UWVBFDQogICAgICAgU1lOVEFYICAgICAgQ291bnRlcjMyDQogICAgICAgTUFYLUFDQ0VTUyAgcmVhZC1vbmx5DQogICAgICAgU1RBVFVTICAgICAgb2Jzb2xldGUNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIlRoZSB0b3RhbCBudW1iZXIgb2YgU05NUCBHZXQtUmVzcG9uc2UgUERVcyB3aGljaA0KICAgICAgICAgICAgICAgaGF2ZSBiZWVuIGdlbmVyYXRlZCBieSB0aGUgU05NUCBwcm90b2NvbCBlbnRpdHkuIg0KICAgICAgIDo6PSB7IHNubXAgMjggfQ0KDQogICBzbm1wT3V0VHJhcHMgT0JKRUNULVRZUEUNCiAgICAgICBTWU5UQVggICAgICBDb3VudGVyMzINCiAgICAgICBNQVgtQUNDRVNTICByZWFkLW9ubHkNCiAgICAgICBTVEFUVVMgICAgICBvYnNvbGV0ZQ0KICAgICAgIERFU0NSSVBUSU9ODQogICAgICAgICAgICAgICAiVGhlIHRvdGFsIG51bWJlciBvZiBTTk1QIFRyYXAgUERVcyB3aGljaCBoYXZlDQogICAgICAgICAgICAgICBiZWVuIGdlbmVyYXRlZCBieSB0aGUgU05NUCBwcm90b2NvbCBlbnRpdHkuIg0KICAgICAgIDo6PSB7IHNubXAgMjkgfQ0KDQogICBzbm1wT2Jzb2xldGVHcm91cCBPQkpFQ1QtR1JPVVANCiAgICAgICBPQkpFQ1RTIHsgc25tcE91dFBrdHMsIHNubXBJblRvb0JpZ3MsIHNubXBJbk5vU3VjaE5hbWVzLA0KICAgICAgICAgICAgICAgICBzbm1wSW5CYWRWYWx1ZXMsIHNubXBJblJlYWRPbmx5cywgc25tcEluR2VuRXJycywNCiAgICAgICAgICAgICAgICAgc25tcEluVG90YWxSZXFWYXJzLCBzbm1wSW5Ub3RhbFNldFZhcnMsDQogICAgICAgICAgICAgICAgIHNubXBJbkdldFJlcXVlc3RzLCBzbm1wSW5HZXROZXh0cywgc25tcEluU2V0UmVxdWVzdHMsDQogICAgICAgICAgICAgICAgIHNubXBJbkdldFJlc3BvbnNlcywgc25tcEluVHJhcHMsIHNubXBPdXRUb29CaWdzLA0KICAgICAgICAgICAgICAgICBzbm1wT3V0Tm9TdWNoTmFtZXMsIHNubXBPdXRCYWRWYWx1ZXMsDQogICAgICAgICAgICAgICAgIHNubXBPdXRHZW5FcnJzLCBzbm1wT3V0R2V0UmVxdWVzdHMsIHNubXBPdXRHZXROZXh0cywNCiAgICAgICAgICAgICAgICAgc25tcE91dFNldFJlcXVlc3RzLCBzbm1wT3V0R2V0UmVzcG9uc2VzLCBzbm1wT3V0VHJhcHMNCiAgICAgICAgICAgICAgICAgfQ0KICAgICAgIFNUQVRVUyAgb2Jzb2xldGUNCiAgICAgICBERVNDUklQVElPTg0KICAgICAgICAgICAgICAgIkEgY29sbGVjdGlvbiBvZiBvYmplY3RzIGZyb20gUkZDIDEyMTMgbWFkZSBvYnNvbGV0ZQ0KICAgICAgICAgICAgICAgYnkgdGhpcyBNSUIgbW9kdWxlLiINCiAgICAgICA6Oj0geyBzbm1wTUlCR3JvdXBzIDEwIH0NCg0KRU5EDQo="
}
```

</details>

### Возможные коды ошибок

400: Bad Request - Отсутствует обязательный параметр (path)  
404: Not Found - Отсутствует файла MIB по указанному пути в системе  
500: Internal Server Error - Ошибка чтения файла MIB

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
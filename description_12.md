# 12. Конфигурация: Пороги

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

> [!NOTE]
> **Пороги** - это система определения состояния аварии составных частей устройства и устройства в целом, а также 
> значений и состояний аварии их параметров по значениям определённого набора параметров и атрибутов, связанных между
> собой логическим выражением.

> [!TIP]
> Внутри системы каждая составная часть устройства и устройство в целом моделируются с использованием компонентов
> Модельного каталога, в котором каждый из компонентов является программной моделью какого-либо устройства, его 
> составной части, логической, функциональной или структурной сущности.  
> Каждый из указанных компонентов обладает определённым набором параметров, обусловленных назначением, составом и
> принципом работы компонента.  
> В системе в конфигурациях опроса устройств компоненты Модельного каталога выстраиваются в структуру, отражающую состав
> и схему деления устройства, а также задаются правила приведения значений, полученных в результате опроса устройства, к
> истинным значениям его параметров.  
> В процессе анализа результата опроса устройства на основании конфигурации опроса каждый его компонент приобретает свой 
> индекс (атрибут `internal_order`) внутри своего родительского компонента, а параметры - значения. Итоговый результат
> присваивается конкретному устройству с указанием протокола, имени хоста и номера порта, по которым производился 
> опрос.  
> Таким образом, результат опроса каждого отдельного устройства приводится к соответствующей ему структуре компонентов, 
> параметры которых приобретают те или иные значения. Эта структура компонентов внутри системы является моделью 
> устройства в целом и отображается пользователю в качестве результата опроса.  
> Каждый компонент и каждый параметр обладают атрибутом `alarm_level`, отражающим состояние тревоги на компоненте или по
> значению параметра.
> 
> Пороги являются инструментом пост-обработки информации и вычисляются на основании истинных значений параметров и 
> атрибутов. Определение целей и источников данных для порогов осуществляется с помощью узлов.

> [!NOTE]
> **Узлы порога** - это атрибуты компонентов, атрибуты и значения параметров, определяющие цели и источники данных
> порогов.

> [!TIP]
> К узлам порогов относятся:
> 
> - промежуточные узлы:
>   * протоколы опроса;
>   * имена хостов устройств;
>   * номера портов устройств;
>   * наименования (значения полей 'title') компонентов;
>   * наименования (значения полей 'title') параметров;
>   * атрибуты `internal_order` компонентов;
> - терминальные узлы:
>   * атрибуты `alarm_level` компонентов;
>   * атрибуты `alarm_level` параметров;
>   * атрибуты `value` параметров.
> 
> Узлы порогов объединяются в цепочки следования и должны начинаться с промежуточных узлов и заканчиваться терминальными
> узлами.  
> 
> Типовая структура узла должна соответствовать одному из следующих шаблонов:
> 
> ```text
> [PROTOCOL]:[host]:[port].['title' компонента].['internal_order' компонента].['alarm_level' компонента]
> [PROTOCOL]:[host]:[port].['title' компонента].['internal_order' компонента].['title' параметра].['alarm_level' параметра]
> [PROTOCOL]:[host]:[port].['title' компонента].['internal_order' компонента].['title' параметра].['value' параметра]
> [PROTOCOL]:[host]:[port].['title' компонента].['internal_order' компонента].['title' компонента].['internal_order' компонента]. ... ['alarm_level' компонента]
> [PROTOCOL]:[host]:[port].['title' компонента].['internal_order' компонента].['title' компонента].['internal_order' компонента]. ... ['title' параметра].['alarm_level' параметра]
> [PROTOCOL]:[host]:[port].['title' компонента].['internal_order' компонента].['title' компонента].['internal_order' компонента]. ... ['title' параметра].['value' параметра]
> ```
> 
> Как видно из приведённых шаблонов, фактически каждый узел является фильтром внутри итоговой структуры результата 
> опроса устройства, позволяющим однозначно назначить цель порога и определить, является ли текущее значение параметра 
> или атрибута источником данных.
> 
> Внутри системы каждый порог задаётся моделью, которая содержит:
> 
> - `id` - идентификатор порога;
> - `author` - имя автора порога;
> - `description` - текстовое описание порога;
> - `created` - дата и время создания порога;
> - `query` - логическое выражение, состоящее из узлов порогов, операторов сравнения и величин, с которыми производится 
>   сравнение, и других логических выражений, итоговым результатом вычисления которого должна быть булева величина;
> - `target` - целевой узел порога, в котором следует изменить значение при получении результата вычисления `query`,
>   равного true;
> - `value` - значение, присваиваемое целевому узлу порога.
> 
> Атрибут `alarm_level` компонентов может принимать следующие значения:
> 
> - NONE - состояние нормы;
> - ALARM - состояние аварии;
> - WARNING - состояние предупреждения.
> 
> Атрибут `alarm_level` параметров может принимать следующие значения:
> 
> - NONE - состояние нормы;
> - ALARM - состояние аварии;
> - WARNING - состояние предупреждения;
> - LOW ALARM - состояние аварии по нижнему пороговому значению;
> - LOW WARNING - состояние предупреждения по нижнему пороговому значению;
> - LOW ALARM - состояние аварии по верхнему пороговому значению;
> - LOW WARNING - состояние предупреждения по верхнему пороговому значению.
> 
> Атрибут `internal_order` может принимать только целочисленное значение, большее 0.
> 
> Атрибут `value` может принимать любое значение, но обязан соответствовать типу целевого параметра (тип параметра 
> содержится в результате опроса устройства, а также может быть получен с использованием API модельного каталога).
>
> В текущей модели порогов системы идентификация устройства по протоколу опроса, имени хоста и номеру порта сведена в
> единую структуру данных. Все оставшиеся узлы порогов являются самостоятельными универсальными единицами, в которых
> значение конкретного узла задаётся значением соответствующего этому узлу поля. Нужно отметить, что наименованию
> компонента всегда соответствует какое-либо значение `internal_order` (которое может не задаваться либо быть равным 
> null для случаев задания правила фильтрации "любой"), поэтому для задания такого сочетания используется единая
> универсальная единица. Наименование компонента или параметра тоже может быть "любым", тогда значение соответствующего 
> этому правилу фильтрации поля должно быть равно null, либо не задаваться вовсе. Для указания "этого" компонента или
> параметра в соответствующее поле ставится пустая строка, а для `internal_order` значение 0.

> [!WARNING]
> Определение того, подходит ли данный порог для анализа результата опроса текущего устройства или нет, осуществляется 
> автоматически. То есть действия всех порогов распространяются на все опрашиваемые устройства, независимо от протокола 
> опроса.  
> Это означает, что как для операций сохранения нового порога, так и для операций их удаления разумно было бы добавить
> кнопку или всплывающее окно подтверждения операции.

> [!CAUTION]
> API, содержащиеся в данном блоке можно разделить на 2 группы: структурные, которые доступны всем пользователям, и
> текстовые, которые будут доступны только в профиле "dev" системы.  
> 
> Текстовые API имеют синтаксис, соответствующий приведённым выше шаблонам порогов, и предназначены только для 
> обеспечения удобства разработки и начального конфигурирования. В дополнение к синтаксису приведённых шаблонов
> добавлены специальные символы:
> 
> - `#` - аналог ключевого слова this (этот) - соответствующее значение поля в структурированном json-объекте - пустая 
>   строка (`""`) или ноль (для `internal_order`);
> - `*` - аналог ключевого слова any (любой) - соответствует NULL-значениям полей.
>
> Целое значение для `internal_order` в текстовых API задаётся внутри квадратных скобок (например, `[1]`).
>
> Примеры текстовых выражений и соответствующих им json-объектов:
> 
> <details><summary>Пример 1</summary>
> 
> ```text
> IF *:*:*.cpu.*.cpu_core.*.load_percent > 90 THEN *:*:*.cpu.*.alarm_level = 'WARNING'
> ```
> 
> ```json
> {
>   "id": 1,
>   "name": "Высокая загрузка ядер CPU",
>   "description": "Проверка превышения порога загрузки процессора",
>   "author": "system_admin",
>   "created": "2026-07-13T19:04:37Z",
>   "query": {
>     "root": {
>       "element": {
>         "comparison": {
>           "target": {
>             "protocol": null,
>             "host": null,
>             "port": null,
>             "target": {
>               "component": "cpu",
>               "internal_order": null,
>               "param": null,
>               "field": null,
>               "next": {
>                 "component": "cpu_core",
>                 "internal_order": null,
>                 "param": null,
>                 "field": null,
>                 "next": {
>                   "component": null,
>                   "internal_order": null,
>                   "param": "load_percent",
>                   "field": null
>                 }
>               }
>             }
>           },
>           "operator": ">",
>           "value": "90"
>         }
>       }
>     }
>   },
>   "target": {
>     "protocol": null,
>     "host": null,
>     "port": null,
>     "target": {
>       "component": "cpu",
>       "internal_order": null,
>       "param": null,
>       "field": null,
>       "next": {
>         "component": null,
>         "internal_order": null,
>         "param": null,
>         "field": "alarm_level"
>       }
>     }
>   },
>   "value": "WARNING"
> }
> ```
> 
> </details>
> 
> <details><summary>Пример 2</summary>
> 
> ```text
> IF *:#:#.ups.*.battery.*.alarm_level == 'HIGH WARNING' 
>   OR ( *:#:#.ups.*.input.*.voltage.value > 240 OR *:#:#.ups.*.input.*.amperage.value > 1.2 ) 
>   OR ( 
>      ( *:#:#.ups.*.output.*.voltage.value > 240 OR *:#:#.ups.*.output.*.amperage.value > 1.2 ) 
>      AND ( *:#:#.ups.*.output.*.voltage.value > 240 OR *:#:#.ups.*.output.*.amperage.value > 1.2 ) 
>      AND ( *:#:#.ups.*.output.*.voltage.value > 240 OR *:#:#.ups.*.output.*.amperage.value > 1.2 ) 
>      ) 
>   OR *:#:#.ups.*.battery.*.charge_level.value <= 15 
> THEN #:#:#.ups.*.alarm_level = 'WARNING'
> ```
> 
> ```json
> {
>   "id": 2,
>   "name": "Высокая загрузка ядер CPU",
>   "description": "Проверка превышения порога",
>   "author": "system_admin",
>   "created": "2026-07-13T21:40:17Z",
>   "query": {
>     "root": {
>       "element": {
>         "comparison": {
>           "target": {
>             "protocol": null,
>             "host": "",
>             "port": 0,
>             "target": {
>               "component": "ups",
>               "internal_order": null,
>               "param": null,
>               "field": null,
>               "next": {
>                 "component": "battery",
>                 "internal_order": null,
>                 "param": null,
>                 "field": null,
>                 "next": {
>                   "component": null,
>                   "internal_order": null,
>                   "param": null,
>                   "field": "alarm_level"
>                 }
>               }
>             }
>           },
>           "operator": "==",
>           "value": "HIGH WARNING"
>         }
>       },
>       "operator": "OR",
>       "next": {
>         "element": {
>           "expression": {
>             "root": {
>               "element": {
>                 "comparison": {
>                   "target": {
>                     "protocol": null,
>                     "host": "",
>                     "port": 0,
>                     "target": {
>                       "component": "ups",
>                       "internal_order": null,
>                       "param": null,
>                       "field": null,
>                       "next": {
>                         "component": "input",
>                         "internal_order": null,
>                         "param": null,
>                         "field": null,
>                         "next": {
>                           "component": null,
>                           "internal_order": null,
>                           "param": "voltage",
>                           "field": null,
>                           "next": {
>                             "component": null,
>                             "internal_order": null,
>                             "param": null,
>                             "field": "value"
>                           }
>                         }
>                       }
>                     }
>                   },
>                   "operator": ">",
>                   "value": "240"
>                 }
>               },
>               "operator": "OR",
>               "next": {
>                 "element": {
>                   "comparison": {
>                     "target": {
>                       "protocol": null,
>                       "host": "",
>                       "port": 0,
>                       "target": {
>                         "component": "ups",
>                         "internal_order": null,
>                         "param": null,
>                         "field": null,
>                         "next": {
>                           "component": "input",
>                           "internal_order": null,
>                           "param": null,
>                           "field": null,
>                           "next": {
>                             "component": null,
>                             "internal_order": null,
>                             "param": "amperage",
>                             "field": null,
>                             "next": {
>                               "component": null,
>                               "internal_order": null,
>                               "param": null,
>                               "field": "value"
>                             }
>                           }
>                         }
>                       }
>                     },
>                     "operator": ">",
>                     "value": "1.2"
>                   }
>                 }
>               }
>             }
>           }
>         },
>         "operator": "OR",
>         "next": {
>           "element": {
>             "expression": {
>               "root": {
>                 "element": {
>                   "expression": {
>                     "root": {
>                       "element": {
>                         "comparison": {
>                           "target": {
>                             "protocol": null,
>                             "host": "",
>                             "port": 0,
>                             "target": {
>                               "component": "ups",
>                               "internal_order": null,
>                               "param": null,
>                               "field": null,
>                               "next": {
>                                 "component": "output",
>                                 "internal_order": null,
>                                 "param": null,
>                                 "field": null,
>                                 "next": {
>                                   "component": null,
>                                   "internal_order": null,
>                                   "param": "voltage",
>                                   "field": null,
>                                   "next": {
>                                     "component": null,
>                                     "internal_order": null,
>                                     "param": null,
>                                     "field": "value"
>                                   }
>                                 }
>                               }
>                             }
>                           },
>                           "operator": ">",
>                           "value": "240"
>                         }
>                       },
>                       "operator": "OR",
>                       "next": {
>                         "element": {
>                           "comparison": {
>                             "target": {
>                               "protocol": null,
>                               "host": "",
>                               "port": 0,
>                               "target": {
>                                 "component": "ups",
>                                 "internal_order": null,
>                                 "param": null,
>                                 "field": null,
>                                 "next": {
>                                   "component": "output",
>                                   "internal_order": null,
>                                   "param": null,
>                                   "field": null,
>                                   "next": {
>                                     "component": null,
>                                     "internal_order": null,
>                                     "param": "amperage",
>                                     "field": null,
>                                     "next": {
>                                       "component": null,
>                                       "internal_order": null,
>                                       "param": null,
>                                       "field": "value"
>                                     }
>                                   }
>                                 }
>                               }
>                             },
>                             "operator": ">",
>                             "value": "1.2"
>                           }
>                         }
>                       }
>                     }
>                   }
>                 },
>                 "operator": "AND",
>                 "next": {
>                   "element": {
>                     "expression": {
>                       "root": {
>                         "element": {
>                           "comparison": {
>                             "target": {
>                               "protocol": null,
>                               "host": "",
>                               "port": 0,
>                               "target": {
>                                 "component": "ups",
>                                 "internal_order": null,
>                                 "param": null,
>                                 "field": null,
>                                 "next": {
>                                   "component": "output",
>                                   "internal_order": null,
>                                   "param": null,
>                                   "field": null,
>                                   "next": {
>                                     "component": null,
>                                     "internal_order": null,
>                                     "param": "voltage",
>                                     "field": null,
>                                     "next": {
>                                       "component": null,
>                                       "internal_order": null,
>                                       "param": null,
>                                       "field": "value"
>                                     }
>                                   }
>                                 }
>                               }
>                             },
>                             "operator": ">",
>                             "value": "240"
>                           }
>                         },
>                         "operator": "OR",
>                         "next": {
>                           "element": {
>                             "comparison": {
>                               "target": {
>                                 "protocol": null,
>                                 "host": "",
>                                 "port": 0,
>                                 "target": {
>                                   "component": "ups",
>                                   "internal_order": null,
>                                   "param": null,
>                                   "field": null,
>                                   "next": {
>                                     "component": "output",
>                                     "internal_order": null,
>                                     "param": null,
>                                     "field": null,
>                                     "next": {
>                                       "component": null,
>                                       "internal_order": null,
>                                       "param": "amperage",
>                                       "field": null,
>                                       "next": {
>                                         "component": null,
>                                         "internal_order": null,
>                                         "param": null,
>                                         "field": "value"
>                                       }
>                                     }
>                                   }
>                                 }
>                               },
>                               "operator": ">",
>                               "value": "1.2"
>                             }
>                           }
>                         }
>                       }
>                     }
>                   },
>                   "operator": "AND",
>                   "next": {
>                     "element": {
>                       "expression": {
>                         "root": {
>                           "element": {
>                             "comparison": {
>                               "target": {
>                                 "protocol": null,
>                                 "host": "",
>                                 "port": 0,
>                                 "target": {
>                                   "component": "ups",
>                                   "internal_order": null,
>                                   "param": null,
>                                   "field": null,
>                                   "next": {
>                                     "component": "output",
>                                     "internal_order": null,
>                                     "param": null,
>                                     "field": null,
>                                     "next": {
>                                       "component": null,
>                                       "internal_order": null,
>                                       "param": "voltage",
>                                       "field": null,
>                                       "next": {
>                                         "component": null,
>                                         "internal_order": null,
>                                         "param": null,
>                                         "field": "value"
>                                       }
>                                     }
>                                   }
>                                 }
>                               },
>                               "operator": ">",
>                               "value": "240"
>                             }
>                           },
>                           "operator": "OR",
>                           "next": {
>                             "element": {
>                               "comparison": {
>                                 "target": {
>                                   "protocol": null,
>                                   "host": "",
>                                   "port": 0,
>                                   "target": {
>                                     "component": "ups",
>                                     "internal_order": null,
>                                     "param": null,
>                                     "field": null,
>                                     "next": {
>                                       "component": "output",
>                                       "internal_order": null,
>                                       "param": null,
>                                       "field": null,
>                                       "next": {
>                                         "component": null,
>                                         "internal_order": null,
>                                         "param": "amperage",
>                                         "field": null,
>                                         "next": {
>                                           "component": null,
>                                           "internal_order": null,
>                                           "param": null,
>                                           "field": "value"
>                                         }
>                                       }
>                                     }
>                                   }
>                                 },
>                                 "operator": ">",
>                                 "value": "1.2"
>                               }
>                             }
>                           }
>                         }
>                       }
>                     }
>                   }
>                 }
>               }
>             }
>           },
>           "operator": "OR",
>           "next": {
>             "element": {
>               "comparison": {
>                 "target": {
>                   "protocol": null,
>                   "host": "",
>                   "port": 0,
>                   "target": {
>                     "component": "ups",
>                     "internal_order": null,
>                     "param": null,
>                     "field": null,
>                     "next": {
>                       "component": "battery",
>                       "internal_order": null,
>                       "param": null,
>                       "field": null,
>                       "next": {
>                         "component": null,
>                         "internal_order": null,
>                         "param": null,
>                         "field": "charge_level",
>                         "next": {
>                           "component": null,
>                           "internal_order": null,
>                           "param": null,
>                           "field": "value"
>                         }
>                       }
>                     }
>                   }
>                 },
>                 "operator": "<=",
>                 "value": "15"
>               }
>             }
>           }
>         }
>       }
>     }
>   },
>   "target": {
>     "protocol": "",
>     "host": "",
>     "port": 0,
>     "target": {
>       "component": "ups",
>       "internal_order": null,
>       "param": null,
>       "field": null,
>       "next": {
>         "component": null,
>         "internal_order": null,
>         "param": null,
>         "field": "alarm_level"
>       }
>     }
>   },
>   "value": "WARNING"
> }
> ```
> 
> </details>

---

## [POST] /api/v1/catalog/threshold - Создать порог

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold

{
  "name": "Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

Ответ 1:

```json
{
  "id": 3,
  "name": "Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "created": "2026-07-19T07:12:09Z",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации переданной JSON-структуры порога или пропущены обязательные поля  
500: Internal Server Error - Внутренняя ошибка базы данных при сохранении структуры порога

## [GET] /api/v1/catalog/threshold/{id} - Получить порог по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/3

{}
```

Ответ 1:

```json
{
  "id": 3,
  "name": "Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "created": "2026-07-19T07:12:09Z",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат идентификатора порога в пути  
404: Not Found - Порог с указанным ID не найден в системе  
500: Internal Server Error - Системная ошибка базы данных при извлечении структуры порога

---

## [PUT] /api/v1/catalog/threshold/1 - Обновить порог по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/1

{
  "name": "НЕ Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

Ответ 1:

```json
{
  "id": 3,
  "name": "НЕ Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "created": "2026-07-19T07:12:09Z",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации структуры тела или неверный числовой формат ID в пути  
404: Not Found - Обновляемый порог с указанным ID не найден  
500: Internal Server Error - Внутренняя ошибка СУБД при обновлении полей порога

---

## [DELETE] /api/v1/catalog/threshold/{id} - Удалить порог по ID

> [!IMPORTANT]
> При удалении порога устройства происходит сдвиг (уменьшение значения id на 1) всего списка порогов, следовавших за
> удалённым, что не приводит к сбоям в работе системы.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
DELETE https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/3

{}
```

Ответ 1:

```json
{}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Некорректный числовой формат ID порога в пути запроса  
404: Not Found - Удаляемый порог не обнаружен в базе данных  
500: Internal Server Error - Ошибка целостности СУБД при каскадном удалении порога

---

## [GET] /api/v1/catalog/threshold/{id}/from-string - Получить эквивалентную строку выражения порога по ID

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/1/from-string

{}
```

Ответ 1:

```json
{
  "id": 1,
  "name": "Высокая загрузка ядер CPU",
  "description": "Проверка превышения порога загрузки процессора",
  "author": "system_admin",
  "expression": "IF *:*:*.cpu.*.cpu_core.*.load_percent > 90 THEN *:*:*.cpu.*.alarm_level = 'WARNING'",
  "created": "2026-07-13T19:04:37Z"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный формат ID в пути запроса  
404: Not Found - Порог с указанным ID не существует в системе  
500: Internal Server Error - Ошибка десериализации дерева условий JSONB в текстовую строку

---

## [PUT] /api/v1/catalog/threshold/{id}/from-string - Обновить порог по ID из строки

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PUT https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/2/from-string

{
  "author": "system_admin",
  "description": "Проверка ЦПУ",
  "expression": "IF *:*:*.cpu.[1].cpu_core.*.load_percent > 90 THEN *:*:*.cpu.[1].alarm_level = 'WARNING'",
  "name": "Высокая загрузка ядер CPU"
}
```

Ответ 1:

```json
{
  "id": 2,
  "name": "Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "created": "2026-07-19T07:17:17Z",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Синтаксическая ошибка при парсинге строкового выражения порога  
404: Not Found - Обновляемый порог не найден  
500: Internal Server Error - Внутренняя ошибка сервера при конвертации строки в AST-дерево СУБД

---

## [PATCH] /api/v1/catalog/threshold/{prevId}/{newId} - Изменить ID порога

> [!TIP]
> API предназначен для упорядочивания перечня порогов и производит смещение порога с одной указанной позиции на другую.

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
PATCH https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/1/2

{}
```

Ответ 1:

```json
{}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Неверный числовой формат одного из ID в пути запроса  
404: Not Found - Исходный порог с данным ID не найден  
500: Internal Server Error - Ошибка обновления первичного или внешних ключей в базе данных

---

## [POST] /api/v1/catalog/threshold/from-string - Создать порог из эквивалентной строки

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
POST https://nms-dev.opk-bulat.ru/api/v1/catalog/threshold/from-string

{
  "author": "system_admin",
  "description": "Проверка ЦПУ",
  "expression": "IF *:*:*.cpu.[1].cpu_core.*.load_percent > 90 THEN *:*:*.cpu.[1].alarm_level = 'WARNING'",
  "name": "Высокая загрузка ядер CPU"
}
```

Ответ 1:

```json
{
  "id": 3,
  "name": "Высокая загрузка ядер CPU",
  "description": "Проверка ЦПУ",
  "author": "system_admin",
  "created": "2026-07-19T07:17:17Z",
  "query": {
    "root": {
      "element": {
        "comparison": {
          "target": {
            "protocol": null,
            "host": null,
            "port": null,
            "target": {
              "component": "cpu",
              "internal_order": 1,
              "param": null,
              "field": null,
              "next": {
                "component": "cpu_core",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": null,
                  "internal_order": null,
                  "param": "load_percent",
                  "field": null
                }
              }
            }
          },
          "operator": ">",
          "value": "90"
        }
      }
    }
  },
  "target": {
    "protocol": null,
    "host": null,
    "port": null,
    "target": {
      "component": "cpu",
      "internal_order": 1,
      "param": null,
      "field": null,
      "next": {
        "component": null,
        "internal_order": null,
        "param": null,
        "field": "alarm_level"
      }
    }
  },
  "value": "WARNING"
}
```

</details>

</details>

### Возможные коды ошибок

400: Bad Request - Ошибка валидации структуры или синтаксическая ошибка в текстовом выражении query  
500: Internal Server Error - Ошибка построения AST-структуры условий и её сохранения в JSONB

---

## [GET] /api/v1/catalog/thresholds - Получить все пороги

<details><summary>Примеры запросов</summary>

### Примеры запросов

<details><summary>Пример 1</summary>

Запрос 1:

```http
GET https://nms-dev.opk-bulat.ru/api/v1/catalog/thresholds

{}
```

Ответ 1:

```json
[
  {
    "id": 1,
    "name": "Высокая загрузка ядер CPU",
    "description": "Проверка превышения порога загрузки процессора",
    "author": "system_admin",
    "created": "2026-07-13T19:04:37Z",
    "query": {
      "root": {
        "element": {
          "comparison": {
            "target": {
              "protocol": null,
              "host": null,
              "port": null,
              "target": {
                "component": "cpu",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": "cpu_core",
                  "internal_order": null,
                  "param": null,
                  "field": null,
                  "next": {
                    "component": null,
                    "internal_order": null,
                    "param": "load_percent",
                    "field": null
                  }
                }
              }
            },
            "operator": ">",
            "value": "90"
          }
        }
      }
    },
    "target": {
      "protocol": null,
      "host": null,
      "port": null,
      "target": {
        "component": "cpu",
        "internal_order": null,
        "param": null,
        "field": null,
        "next": {
          "component": null,
          "internal_order": null,
          "param": null,
          "field": "alarm_level"
        }
      }
    },
    "value": "WARNING"
  },
  {
    "id": 2,
    "name": "Высокая загрузка ядер CPU",
    "description": "Проверка превышения порога",
    "author": "system_admin",
    "created": "2026-07-13T21:40:17Z",
    "query": {
      "root": {
        "element": {
          "comparison": {
            "target": {
              "protocol": null,
              "host": "",
              "port": 0,
              "target": {
                "component": "ups",
                "internal_order": null,
                "param": null,
                "field": null,
                "next": {
                  "component": "battery",
                  "internal_order": null,
                  "param": null,
                  "field": null,
                  "next": {
                    "component": null,
                    "internal_order": null,
                    "param": null,
                    "field": "alarm_level"
                  }
                }
              }
            },
            "operator": "==",
            "value": "HIGH WARNING"
          }
        },
        "operator": "OR",
        "next": {
          "element": {
            "expression": {
              "root": {
                "element": {
                  "comparison": {
                    "target": {
                      "protocol": null,
                      "host": "",
                      "port": 0,
                      "target": {
                        "component": "ups",
                        "internal_order": null,
                        "param": null,
                        "field": null,
                        "next": {
                          "component": "input",
                          "internal_order": null,
                          "param": null,
                          "field": null,
                          "next": {
                            "component": null,
                            "internal_order": null,
                            "param": "voltage",
                            "field": null,
                            "next": {
                              "component": null,
                              "internal_order": null,
                              "param": null,
                              "field": "value"
                            }
                          }
                        }
                      }
                    },
                    "operator": ">",
                    "value": "240"
                  }
                },
                "operator": "OR",
                "next": {
                  "element": {
                    "comparison": {
                      "target": {
                        "protocol": null,
                        "host": "",
                        "port": 0,
                        "target": {
                          "component": "ups",
                          "internal_order": null,
                          "param": null,
                          "field": null,
                          "next": {
                            "component": "input",
                            "internal_order": null,
                            "param": null,
                            "field": null,
                            "next": {
                              "component": null,
                              "internal_order": null,
                              "param": "amperage",
                              "field": null,
                              "next": {
                                "component": null,
                                "internal_order": null,
                                "param": null,
                                "field": "value"
                              }
                            }
                          }
                        }
                      },
                      "operator": ">",
                      "value": "1.2"
                    }
                  }
                }
              }
            }
          },
          "operator": "OR",
          "next": {
            "element": {
              "expression": {
                "root": {
                  "element": {
                    "expression": {
                      "root": {
                        "element": {
                          "comparison": {
                            "target": {
                              "protocol": null,
                              "host": "",
                              "port": 0,
                              "target": {
                                "component": "ups",
                                "internal_order": null,
                                "param": null,
                                "field": null,
                                "next": {
                                  "component": "output",
                                  "internal_order": null,
                                  "param": null,
                                  "field": null,
                                  "next": {
                                    "component": null,
                                    "internal_order": null,
                                    "param": "voltage",
                                    "field": null,
                                    "next": {
                                      "component": null,
                                      "internal_order": null,
                                      "param": null,
                                      "field": "value"
                                    }
                                  }
                                }
                              }
                            },
                            "operator": ">",
                            "value": "240"
                          }
                        },
                        "operator": "OR",
                        "next": {
                          "element": {
                            "comparison": {
                              "target": {
                                "protocol": null,
                                "host": "",
                                "port": 0,
                                "target": {
                                  "component": "ups",
                                  "internal_order": null,
                                  "param": null,
                                  "field": null,
                                  "next": {
                                    "component": "output",
                                    "internal_order": null,
                                    "param": null,
                                    "field": null,
                                    "next": {
                                      "component": null,
                                      "internal_order": null,
                                      "param": "amperage",
                                      "field": null,
                                      "next": {
                                        "component": null,
                                        "internal_order": null,
                                        "param": null,
                                        "field": "value"
                                      }
                                    }
                                  }
                                }
                              },
                              "operator": ">",
                              "value": "1.2"
                            }
                          }
                        }
                      }
                    }
                  },
                  "operator": "AND",
                  "next": {
                    "element": {
                      "expression": {
                        "root": {
                          "element": {
                            "comparison": {
                              "target": {
                                "protocol": null,
                                "host": "",
                                "port": 0,
                                "target": {
                                  "component": "ups",
                                  "internal_order": null,
                                  "param": null,
                                  "field": null,
                                  "next": {
                                    "component": "output",
                                    "internal_order": null,
                                    "param": null,
                                    "field": null,
                                    "next": {
                                      "component": null,
                                      "internal_order": null,
                                      "param": "voltage",
                                      "field": null,
                                      "next": {
                                        "component": null,
                                        "internal_order": null,
                                        "param": null,
                                        "field": "value"
                                      }
                                    }
                                  }
                                }
                              },
                              "operator": ">",
                              "value": "240"
                            }
                          },
                          "operator": "OR",
                          "next": {
                            "element": {
                              "comparison": {
                                "target": {
                                  "protocol": null,
                                  "host": "",
                                  "port": 0,
                                  "target": {
                                    "component": "ups",
                                    "internal_order": null,
                                    "param": null,
                                    "field": null,
                                    "next": {
                                      "component": "output",
                                      "internal_order": null,
                                      "param": null,
                                      "field": null,
                                      "next": {
                                        "component": null,
                                        "internal_order": null,
                                        "param": "amperage",
                                        "field": null,
                                        "next": {
                                          "component": null,
                                          "internal_order": null,
                                          "param": null,
                                          "field": "value"
                                        }
                                      }
                                    }
                                  }
                                },
                                "operator": ">",
                                "value": "1.2"
                              }
                            }
                          }
                        }
                      }
                    },
                    "operator": "AND",
                    "next": {
                      "element": {
                        "expression": {
                          "root": {
                            "element": {
                              "comparison": {
                                "target": {
                                  "protocol": null,
                                  "host": "",
                                  "port": 0,
                                  "target": {
                                    "component": "ups",
                                    "internal_order": null,
                                    "param": null,
                                    "field": null,
                                    "next": {
                                      "component": "output",
                                      "internal_order": null,
                                      "param": null,
                                      "field": null,
                                      "next": {
                                        "component": null,
                                        "internal_order": null,
                                        "param": "voltage",
                                        "field": null,
                                        "next": {
                                          "component": null,
                                          "internal_order": null,
                                          "param": null,
                                          "field": "value"
                                        }
                                      }
                                    }
                                  }
                                },
                                "operator": ">",
                                "value": "240"
                              }
                            },
                            "operator": "OR",
                            "next": {
                              "element": {
                                "comparison": {
                                  "target": {
                                    "protocol": null,
                                    "host": "",
                                    "port": 0,
                                    "target": {
                                      "component": "ups",
                                      "internal_order": null,
                                      "param": null,
                                      "field": null,
                                      "next": {
                                        "component": "output",
                                        "internal_order": null,
                                        "param": null,
                                        "field": null,
                                        "next": {
                                          "component": null,
                                          "internal_order": null,
                                          "param": "amperage",
                                          "field": null,
                                          "next": {
                                            "component": null,
                                            "internal_order": null,
                                            "param": null,
                                            "field": "value"
                                          }
                                        }
                                      }
                                    }
                                  },
                                  "operator": ">",
                                  "value": "1.2"
                                }
                              }
                            }
                          }
                        }
                      }
                    }
                  }
                }
              }
            },
            "operator": "OR",
            "next": {
              "element": {
                "comparison": {
                  "target": {
                    "protocol": null,
                    "host": "",
                    "port": 0,
                    "target": {
                      "component": "ups",
                      "internal_order": null,
                      "param": null,
                      "field": null,
                      "next": {
                        "component": "battery",
                        "internal_order": null,
                        "param": null,
                        "field": null,
                        "next": {
                          "component": null,
                          "internal_order": null,
                          "param": null,
                          "field": "charge_level",
                          "next": {
                            "component": null,
                            "internal_order": null,
                            "param": null,
                            "field": "value"
                          }
                        }
                      }
                    }
                  },
                  "operator": "<=",
                  "value": "15"
                }
              }
            }
          }
        }
      }
    },
    "target": {
      "protocol": "",
      "host": "",
      "port": 0,
      "target": {
        "component": "ups",
        "internal_order": null,
        "param": null,
        "field": null,
        "next": {
          "component": null,
          "internal_order": null,
          "param": null,
          "field": "alarm_level"
        }
      }
    },
    "value": "WARNING"
  },
  {
    "id": 3,
    "name": "Высокая загрузка ядер CPU",
    "description": "Проверка ЦПУ",
    "author": "system_admin",
    "created": "2026-07-19T07:17:17Z",
    "query": {
      "root": {
        "element": {
          "comparison": {
            "target": {
              "protocol": null,
              "host": null,
              "port": null,
              "target": {
                "component": "cpu",
                "internal_order": 1,
                "param": null,
                "field": null,
                "next": {
                  "component": "cpu_core",
                  "internal_order": null,
                  "param": null,
                  "field": null,
                  "next": {
                    "component": null,
                    "internal_order": null,
                    "param": "load_percent",
                    "field": null
                  }
                }
              }
            },
            "operator": ">",
            "value": "90"
          }
        }
      }
    },
    "target": {
      "protocol": null,
      "host": null,
      "port": null,
      "target": {
        "component": "cpu",
        "internal_order": 1,
        "param": null,
        "field": null,
        "next": {
          "component": null,
          "internal_order": null,
          "param": null,
          "field": "alarm_level"
        }
      }
    },
    "value": "WARNING"
  }
]
```

</details>

</details>

### Возможные коды ошибок

500: Internal Server Error - Критическая ошибка сервера при чтении полного списка порогов из БД

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
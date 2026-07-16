# 3. Модельный каталог: Связи

> [!CAUTION]
> ВСЕ ПЕРЕЧИСЛЕННЫЕ НИЖЕ API ДОСТУПНЫ ТОЛЬКО ДЛЯ ПРОФИЛЯ "dev" ПРИЛОЖЕНИЯ

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)

---

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

---

> [[_**ОГЛАВЛЕНИЕ**_]](./README.md)
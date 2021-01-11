### Запуск

```
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```

#### Запуск тестов

```
docker-compose -f docker-compose.test.yml build
docker-compose -f docker-compose.test.yml up
```

### Использование

В приложении реализована асинхронная схема работы. Чтобы получит ответ на запрос, нужно получить ID задания
```
{
    "task_id": YOUR_TASK_ID
}
```
дождаться его выполнения и получить результат

```
curl --location --request GET 'http://localhost:8000/v1/tasks/?task_id=YOUR_TASK_ID'
``` 

##### Создание пользователя

```
curl --location --request POST 'http://localhost:8000/v1/users' --form 'id="user_id_1"'
```
```
{
    "result": {
        "id": "user_id_1",
        "status": "ok"
    }
}
```

#### Обновление продуктов
Параметры:
* user_id - ID пользователя
```
curl --location --request POST 'http://localhost:8000/v1/products/update/xlsx?user_id=user_id_1' --form 'file=@"./test_data/data_1.xlsx"'
```
```
{
    "result": {
        "count_deleted_products": 0,
        "count_new_products": 3,
        "count_updated_products": 0,
        "not_processed": []
    }
}
```
* * *
```
curl --location --request POST 'http://localhost:8000/v1/products/update/xlsx?user_id=user_id_1' --form 'file=@"./test_data/data_1_2.xlsx"'
```

```
{
    "result": {
        "count_deleted_products": 1,
        "count_new_products": 1,
        "count_updated_products": 1,
        "not_processed": []
    }
}
```
* * *
```
curl --location --request POST 'http://localhost:8000/v1/products/update/xlsx?user_id=user_id_1' --form 'file=@"./test_data/data_1.xlsx"'
```

```
{
    "result": {
        "count_deleted_products": 0,
        "count_new_products": 0,
        "count_updated_products": 0,
        "not_processed": [
            {
                "error": "available is not converted to bool",
                "indexes_in_xlsx": [
                    4
                ]
            },
            {
                "error": "name is required",
                "indexes_in_xlsx": [
                    1
                ]
            },
            {
                "error": "offer ID is required",
                "indexes_in_xlsx": [
                    0
                ]
            },
            {
                "error": "price is less than 0",
                "indexes_in_xlsx": [
                    2
                ]
            },
            {
                "error": "quality is less than 0",
                "indexes_in_xlsx": [
                    3
                ]
            }
        ]
    }
}
```
#### Получения продуктов по user_id, offer_id, названию
Параметры:
 * user_id - ID пользователя
 * offer_id - ID предмета в системе пользователя
 * search - строка для поиска по имени товара
```
curl --location --request GET 'http://localhost:8000/v1/products/?search=%D1%82%D0%B5%D0%BB%D0%B5'
```

```
{
    "result": {
        "count": 2,
        "data": [
            {
                "name": "телефон",
                "offer_id": "offer_id_2",
                "price": 20000,
                "quality": 20
            },
            {
                "name": "телевизор",
                "offer_id": "offer_id_1",
                "price": 10000,
                "quality": 1
            }
        ]
    }
}
```
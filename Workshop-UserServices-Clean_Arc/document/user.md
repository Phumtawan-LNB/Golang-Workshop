# User Service API Documentation
## Introduction

User Services power by kafka server. this services make for remember something or anything 

## Producer / Endpoints

### POST {url}/v1/user/create

Create user data in user services

**Request Body:**:

```json
{
    "name": "Pheem",
    "lat": 13.7,
    "long": 100.5,
    "age": 22,
    "first_name": "phumtawan",
    "last_name": "lunabut"
}
```

**Response**:

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "user create success",
    "status": "OK",
    "status_code": 200
}
```

**Event: UserAuthEvent**

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "name": "Pheem"
}
```

---
### Put {url}/v1/user/update

update user data by id

**Request Body:**

```json
{
  "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
  "name": "PheemZ",
  "lat": 15.7,
  "long": 105.5,
  "age": 23,
  "fist_name": "phumtawan",
  "last_name": "lunabut"
}
```

**Response**

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "user update success",
    "data": {
        "name": "PheemZ",
        "lat": 15.7,
        "long": 105.5,
        "age": 23,
        "fist_name": "phumtawan",
        "last_name": "lunabut"
    },
    "status": "OK",
    "status_code": 200
}
```
**Event: UserUpdateEvent**

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "name": "PheemZ"
}
```
---
### Get {url}/v1/user/readed

Get user readed history data by id

**Request Body:**

```json
{
  "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368"
}
```

**Response**

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "user read history success",
    "data": {
        "name": "PheemZ",
        "lat": 15.7,
        "long": 105.5,
        "age": 23,
        "fist_name": "phumtawan",
        "last_name": "lunabut"
    },
    "history": {
        "weather_search": {
            "weather_id": "xxx-xxx-xxx"{
                "name": "khon kaen",
                "quantity": 2
            },
            "weather_id": "xxx-xxx-xxx"{
                "name": "bangkok",
                "quantity": 3
            },
            "weather_id": "xxx-xxx-xxx"{
                "name": "LA (ROI-ET)",
                "quantity": 1
            }
        }
    },
    "status": "OK",
    "status_code": 200
}
```
---
### Delete {url}/v1/user/delete

Delete user data by id

**Request Body:**

```json
{
  "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368"
}
```

**Response**

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "user delete success",
    "status": "OK",
    "status_code": 200
}
```
**Event: UserDeleteEvent**

```json
{
    "id": "deeea3d7-0e8c-493d-b1ab-0538dec03368"
}
```
---
## Status Codes

<ul>
  <li>200 : OK. Request was successful.</li>
  <li>201 : Created. Resource was successfully created.</li>
  <li>400 : Bad request. The request was invalid or cannot be served.</li>
  <li>422 : Unprocessable Entity. may something broke on function and bad request body.</li>
</ul>

## Change Log

<ul>
  <li>2023-11-21 : build kafka and make user repository.</li>
  <li>2023-11-22 : make API {url}/user/....</li>
  <li>2023-11-22 : update documentation</li>
</ul>

## Support

If you have questions, you can solve it. But if you can't. You need to learn something more or contact me to say hello. This is my mail: Phumtawan.l@kkumail.com

---
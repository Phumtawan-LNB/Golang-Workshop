# Weather Service API Documentation
## Introduction

Weather Services is second Services power by kafka server. this services make for remember something or anything 

## Producer / Endpoints

### POST {url}/v1/weather/create

Generate weather data with lat,long. It can be done automatically.

**Request Body:**:

```json
{
    "lat": 13.7,
    "long": 100.5,
}
```

**Response**:

```json
{
    "id": "feeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "weather create success",
    "data": {
	    "name": "Khon Kaen",
	    "lat": 13.7,
	    "long": 100.5,
	    "Country": "Thailand",
	    "temp": 31,
	    "wind_dir": "ENE",
        "wind_kph": 13.0,
        "uv": 7.0
    },
    "status": "OK",
    "status_code": 200

}
```
---

### Get {url}/v1/weather/search

Get weather data by name and produce weather_id, name into kafka server.

**Request Body:**

```json
{
    "user_id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "name": "khon kaen"
}
```

**Response**:

```json
{
    "id": "keeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "weather read success",
    "data": {
	    "name": "Khon Kaen",
	    "lat": 13.7,
	    "long": 100.5,
	    "Country": "Thailand",
	    "temp": 31,
	    "Wind_dir": "ENE",
        "wind_kph": 13.0,
        "uv": 7.0
    },
    "status": "OK",
    "status_code": 200
}
```
**Event: WeatherSearchEvent**

```json
 {
    "weather_id": "deeea3d7-0e8c-493d-b1ab-0538dec03368",
    "weather_name": "khon kaen",
    "user_id": "feeea3d7-0e8c-493d-b1ab-0538dec03368",
    "user_name": "Pheem"
 }
```
---
### Put {url}/v1/weather/update

Update lat long data by id and current weather conditions will also be updated automatically.
Info: update all data with weather in database

**Request Body:**

```json
{
  "key": "9O1hb4rz7oPIvorzTd3N"
}
```

**Response**

```json
{
    "messge": "weather update success",
    "data": {
	    "Bangkok"{
            "lat": 15.7,
	        "long": 105.5,
	        "Country": "Thailand",
	        "temp": 33,
	        "Wind_dir": "NE",
            "wind_kph": 13.0,
            "uv": 7.0
        },
        "khon kaen"{
            "lat": 13.7,
	        "long": 100.5,
	        "Country": "Thailand",
	        "temp": 31,
	        "wind_dir": "ENE",
            "wind_kph": 13.0,
            "uv": 7.0
        }
    },
    "status": "OK",
    "status_code": 200
}
```
---
### Delete {url}/v1/weather/delete

Delete weather data by id

**Request Body:**

```json
{
  "id": "weeea3d7-0e8c-493d-b1ab-0538dec03368"
}
```

**Response**

```json
{
    "id": "weeea3d7-0e8c-493d-b1ab-0538dec03368",
    "messge": "weather delete success",
    "status": "OK",
    "status_code": 200
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
  <li>2023-11-21 : build kafka and make weather repository.</li>
  <li>2023-11-22 : make API {url}/weather/....</li>
  <li>2023-11-22 : update documentation</li>
</ul>

## Support

If you have questions, you can solve it. But if you can't. You need to learn something more or contact me to say hello. This is my mail: Phumtawan.l@kkumail.com

---
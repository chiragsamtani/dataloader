### Installation

* Go (1.19+)
* Docker

You can run the application using Go or Docker. Make scripts are also provided incase you would like to use them.

Run using Go using:
````
go run main.go
````

Run using Docker compose:
````
docker-compose up -d
````

The application runs on port 8080 (by default), the supported routes are mentioned on the API Contract
section below.

### API Contract


#### Fetches a list of hotels based on hotelIds or destinationI
<details>
<summary><code>POST</code> <code><b>/hotels</b></code> </summary>


##### Parameters

> | name            |  type                                                  | data type               | description                            |
> | --------------- | ------------------------------------------------------ | ----------------------- | -------------------------------------- |
> | hotelIds        |  either hotelIds or destinationId must be supplied     | []string                | List of unique hotelIds                |
> | destinationId   |  either hotelIds or destinationId must be supplied     | int                     | Single numeric destinationId           |


##### Responses

> | http code | content-type                      | response                                | description
> | --------- | --------------------------------- |-----------------------------------------|-------------------------------------------------------------------------------------|
> | `200`     | `application/json`                | `[]`                                    | DestinationId or hotelId not found or data doesn't exist                            |
> | `200`     | `application/json`                | `[<hotel_object>]]`                     | Valid hotel data returned                                                           |
> | `400`     | `application/json`                | `{"message": "Please specify at least one hotel ID(s) or a single destination ID"}` | HotelId or destinationId not supplied or data type of hotelId or destinationId is incorrect |
> | `400`     | `application/json`                | `{"message": "Request body must be in JSON format"}` | Make sure `Content-Type` headers are set to `application/json` and request body content is a valid JSON |
> | `405`     | `application/json`                | `{"message": "Method not allowed"}`      | Use POST as HTTP method, other methods are unsupported                              |

##### Example cURL

> ```javascript
>  curl --request POST --url http://localhost:8080/hotels --header 'Content-Type: application/json' --data '{ "hotel_ids": ["iJhz"] }'
> ```
</details>


**Response Object**

| Field Name  	    | Data Type   	    | Merge Strategy |
|---	            |---	            |--- |
|`id`   	        | String  	        | This is treated as the primary key of the data |
| `destinationId` | Numeric           | This can map to many hotels, that is one destinationId can span multiple hotels |
| `name`  	      | String 	        | Longest hotel name is chosen |
| `location` 	    | Object  	        | Country: will choose ISO-3601 compliant country codes and choose non-empty strings in that order <br />City: will choose longer length city between existing and new data <br /> Address: will choose longer address between existing and new data <br/>Lat: will choose first non-zero data <br/>Lng: will choose first non-zero data|
| `description`  	| String  	        | Longest hotel description is chosen
| `amenities`  	  | Array  	        | Union of existing and new data, filtering out any duplicate or similar data |
| `images`  	    | Array   	        | Union of existing and new data images, the URLs are first added to a Set to make sure we don't have any duplicate data, the returned object will be a unique Set of images with image link and captions |
| `booking_conditions` | Array   	        | Union between the existing and new data with no filter applied |

### Tests

Unit test can be run with:
```
go test ./...
```

Running tests with coverage:
```
go test -v -coverprofile cover.out ./...
```

### Project Structure


- **config**: contains all configuration and infra related files such as parsing environment variables
through .env files. We use `viper` as our configuration management tool
- **handler**: contains all the route or controller files
- **model**: contains all domain specific, business logic data and DTOs
- **repository**: contains all files for accessing and storing data. Supported
repository types currently are: `INMEMORY`. Future support for MySQL will
be added
- **service**: contains all business logic files
- **main.go**: server logic/entry point for Go programs

### Configuration

Data is purged after every app startup and data is loaded via the DataLoader
service with fresh dataset on app startup/initialization.

**SUPPLIER_CONFIG**: this is a comma-seperated key-value pair containing
supplier to URL relation.

**LOG_LEVEL**: Supported log levels are `debug`, `warn` and `error`


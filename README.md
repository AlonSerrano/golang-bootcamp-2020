# ZipCode of Mexico Golang for the Bootcamp  
  
## Introduction  
  
The following project is designed to obtain the postal data of the official post office of Mexico in csv format and then store the information in a database in MongoDB and proceed to query the colonies referring to a postal code as a result in json, in addition to that in order to update the database it is necessary to have a JWT token to perform the action  
  
## Requirements  

First [install](https://github.com/AlonSerrano/golang-bootcamp-2020/blob/Final_Deliverable/captures/swagger.png) go 
  
You need to install MongoDB either with the normal [installer](https://docs.mongodb.com/manual/installation/) or in  [docker](https://docs.docker.com/get-docker/) and run it on

> port 27017

Docker

```shell script
docker pull mongo
```

## Run the code

Run

```shell script
 go install
```

And finally

 ```shell script
  go run main.goo install
 ```

## About the project

### Config File

The config.env file has been added which contains the server configuration variables which it takes when starting the application through viper

 ```shell script
ServerAddress=localhost:8080
DBAddress=localhost
DBPort=27017
ZipCodesDB=https://www.correosdemexico.gob.mx/datosabiertos/cp/cpdescarga.txt
 ```


### Swagger

The swagger interface was also added to be able to analyze and document the end points in a simpler way with swaggo, which is updated with the command:
 ```shell script
swag init
 ```
This command automatically updates the interface of swagger with the comments added in each method

![Sawgger Interface](https://github.com/AlonSerrano/golang-bootcamp-2020/captures/swagger.png)

To be able to fill the data you need to register a user
```shell script
curl -X POST "http://localhost:8080/api/v1/user/signUp" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"birthdate\": \"string\", \"email\": \"string\", \"firstName\": \"string\", \"lastName\": \"string\", \"password\": \"string\", \"phone\": \"string\", \"secondLastName\": \"string\", \"secondName\": \"string\"}"
```

Example Response
```json
{
    "Message": "Registered user successfully",
    "Status": "200"
}
```

Then you have to get the token of the user that you previously create
```shell script
curl -X POST "http://localhost:8080/api/v1/user/signIn" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"email\": \"string\", \"password\": \"string\"}"
```
Example Response
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZW1haWwiOiJsdWlzX2Fsb25zb0BvdXRsb29rLmNvbSIsImV4cCI6MTYwNzkxMTYyMCwiZmlyc3ROYW1lIjoiTHVpcyIsImxhc3ROYW1lIjoiU2VycmFubyIsInN1YiI6MX0.ajRc7Jly3GVgsstryNbB8BuZcIDODIcUuOOk2Midjmo",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc5OTcxMjAsInN1YiI6MX0.sKcKidaW3BJ-JChRypy9H5PHQcgd1xCks4T2C3sknr0"
}
```

To be able to consult the data, you must first fill the tables, so you must execute a get-type request and sending the Token of the previous method on the header to the method:

```shell script
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" -XGET 'localhost:8080/api/v1/address/populate'
```


*What happens behind the scenes is that the code makes a request to a government endpoint which returns a csv with all the zip codes, a 15mb file _(145,110 approximate zip codes)_

Once you get the response of the inserted ids that will take approximately 15 seconds on complete, you can now check the colonies that belong to a postal code with the method

```shell script
  curl -XGET 'localhost:8080/api/v1/address/search/:zipCode'
```

Example Response
```json
{
    "Message": "Have been inserted 145110 postal codes updated by luis_alonso@outlook.com"
}
```

For example:

```shell script
  curl -XGET 'localhost:8080/api/v1/address/search/97306'
```

And an example of the previous request would be the following:

  ```json
    [
      {
        "Id": "56eb060711d84d67ab164e862c5eb0f1",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Chichi Suárez"
      },
      {
        "Id": "cd9c02f9983e4667ae3f730615263907",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Sitpach"
      },
      {
        "Id": "d24212569c7f42df9cb6b1b0f54f913d",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Villas de Oriente"
      },
      {
        "Id": "35467b546c21499e9870c4861fb8257d",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Los Héroes"
      },
      {
        "Id": "7e2386368a624f8c88e1ba827865ce7b",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Santa María Chí"
      },
      {
        "Id": "eee9a4cc733543d7a6d753b2ce40b5ab",
        "CodigoPostal": "97306",
        "Estado": "Yucatán",
        "EstadoISO": "MX-YUC",
        "Municipio": "Mérida",
        "Ciudad": "",
        "Barrio": "Chichi Díaz"
      }
    ]
```

## Unit test

Unit tests were created that test these end-point on the path

> /golang-bootcamp-2020/pkg/util/zipcodes_test.go

And

> /golang-bootcamp-2020/pkg/util/user_test.go

###For Zip Codes

First run the unit test

 ```go
  Test_getCSVCodes()
```

 ```go
  Test_searchZipCodes()
```

Another unit test was created to test the correct flushing of the zipcodes table, run it if you want

 ```go
  Test_dropZipCodes()
```

###For User
 ```go
  TestSignUp()
```


 ```go
  TestSignIn()
```


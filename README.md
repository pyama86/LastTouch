# LastTouch API
This is a simple API that returns the last update time of a given MySQL table. It requires the table's database and table name as URL query parameters and uses HTTP Basic Authentication for authentication.
## Requirements
- Go programming language (version 1.13 or later)
- MySQL Database
## Installation
1. Clone the repository.
2. Set the following environment variables based on your MySQL configuration:

   MYSQL_USER
   MYSQL_PASSWORD
   MYSQL_HOST
   MYSQL_PORT
   MYSQL_DB

Note: You can use the getEnv function with default values in the code if you don't want to use environment variables.
3. Set the environment variables for HTTP Basic Authentication:

   LASTTOUCH_USER
   LASTTOUCH_PASSWORD

4. In the code, replace the jwtSecret variable with your own secret key.
5. Install the necessary dependencies:

   go get -u github.com/go-sql-driver/mysql

6. Compile and run the application:

   go build
   ./main

## Usage
1. Send a GET request to the /getUpdateTime endpoint with URL query parameters db and table.
   For example:

   curl --location --request GET 'http://127.0.0.1:8080/getUpdateTime?db=my_database&table=my_table' --header 'Authorization: Basic <your_base64_encoded_credentials>'

2. If the request is successful, you will receive a JSON response with the update_time field.
   Example response:

json
   {
     "update_time": "2021-12-30 12:34:56"
   }

## License
This project is licensed under the MIT License.

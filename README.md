
# REST API project with Go, Docker and Postgres
This app has been written for the following purposes:
* creating a simple REST API in Golang,
* adding authentication middleware, 
* writing integration tests with a database in Go. 
I've used the Postgres database, and Docker to containerize the app. 
For database management, I've been using Soda CLI.  
 


## Available HTTP methods and endpoints

* GET "/players": This endpoint retrieves a list of all players in the app.

* GET "/players/1": This endpoint retrieves the details of the player with the ID of 1.
* POST "/players": This endpoint is used to create a new player in the app.

#### JSON payload for method POST and endpoint /players
```json
{
    "firstName":"Cristiano",
    "lastName":"Ronaldo",
    "age":37,
    "country":"Portugal",
    "club":"Manchester United",
    "position":"striker",
    "goals":24,
    "assists":6
}
```
* PATCH "/players/1": This endpoint is used to update the details of the player with the ID of 1.

#### JSON payload for method PATCH and endpoint /players/1
```json 
{
    "firstName":"Cristiano",
    "lastName":"Ronaldo",
    "age":37,
    "country":"Portugal",
    "club":"Real Madrid",
    "position":"striker",
    "goals": 27,
    "assists":12
}
```
* DELETE "/players/1": This endpoint is used to delete the player with the ID of 1 from the app.






## Installation

1. Download the app 

```bash
  git clone https://github.com/gavronski/rest-api
```
2. Add docker-compose.yml and database.yml files.

3. Build containers: 
```bash
  docker-compose up -d --build
```
4. Open the "app" container's cli (not db container) pasting the following command: 

```bash 
    docker exec app_container_name bash 
```
Then, run soda migrations and seed tables using command: 
```bash 
    soda migrate
```

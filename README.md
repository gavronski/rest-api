
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


## Running Tests

To run all unit tests, run the following command from the root directory.

```bash
  go test -v ./...
```

To run integration test, make sure that you are in the directory ./internal/repository/dbrepo .

### Important steps: 
#### 1. Creating docker connection 
```golang
// connect to docker
p, err := dockertest.NewPool("")
if err != nil {
  log.Fatalf("Could not connect to docker: %s", err)
}
```
#### 2. Setting up docker's options
```golang
// set up docker options
options := dockertest.RunOptions{
  Repository: "postgres",
  Env: []string{
    "POSTGRES_PASSWORD=" + password,
    "POSTGRES_USER=" + user,
    "POSTGRES_DB=" + dbName,
    "listen_addresses = '*'",
  },
}
```
#### 3. Start docker container 
```golang 
resource, err = pool.RunWithOptions(&options)
```

#### 4. Create tables structure 
```golang 
// populate the database with empty tables
err = createTables()
if err != nil {
  log.Fatalf("error while creating tables: %s", err)
}
```

#### 5. Run tests in the container
```golang 
// run tests
code := m.Run()
```

#### 6. Remove testing container
```golang 
// delete container after tests
_ = pool.Purge(resource)
```

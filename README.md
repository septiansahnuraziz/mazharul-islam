# API Mazharul Islam

## Getting Started

To start using the mazharul-islam console, follow the steps below:

1. Clone the repository:

    ```bash
    git clone git clone github.com/mizharul-islam.git
    ```

2. Jump into the project directory:

    ```bash
    cd mazharul-islam
    ```

3. Build the application:

    ```bash
    go build
    ```

4. Run the server from console:

    ```bash
    ./mazharul-islam server
    ```

## Available Commands

The following commands are available in the mazharul-islam console:

### server

- Description: Starts the server.
- Usage: 
    ```bash
    go run . server
    ```

### worker

- Description: Starts the worker.
- Usage:
    ```bash
    go run . worker
    ```

### migrate

- Description: Migrates the database.
- Usage:
    ```bash
    go run . migrate
    ```
- Optional flags:
    - --step: Sets the maximum migration steps.
    - --direction: Sets the migration direction. up is default value

### create-migration [filename]

- Description: Creates a new database migration file.
- Usage:
    ```bash
    go run . create-migration [migration name]
    ```
- Example:
    ```bash
    go run . create-migration create_customers_table 
    ```


## Requirement
- Go version 1.22.5 as minimum
- Postgres 12 as minimum

## Local Development
- **Please follow step by step:**
- `$ export GO111MODULE="on"`
- `$ go get github.com/mazharul-islam` or `git clone github.com/mazharul-islam.git``
- `$ export PATH=$GOPATH/bin:$PATH`
- `$ go mod tidy`
- run hot reloading:
  #### Server
    ```
    PORT=4000 make run
    ```
    #### Worker
    ```
    PORT=4001 make run-worker
    ```
- See http://localhost:4000/v1/ping

## Local Development using Docker
- ```$ docker build -t mazharul-islam .```
- ```$ docker run -p 4000:4000 mazharul-islam```

## Linter
- `$ make lint-prepare`
- `$ make lint`
- Please do resolve

## Docs
- Endpoint Dev: [http://localhost:4000/*](http://localhost:4000/ "Click")




# test-sms-2-pro
api about pokemon detail
### Setup and Run:

### Setup

#### install
    
 - [install Go](https://go.dev/doc/install)

- [install sqlite](https://www.sqlite.org/download.html)
#### config at path "config/confg.yaml"
1. config number port 
    ```bash
        port: { { app-port } }
    ```
2. config sqlite  database name 
    ```bash
        dbname: { { sqlite-dbname } }
    ```
3. config sqlite database path 
    ```bash
        dbpath: { { sqlite-dbpath } }
    ```
4. config json data file path and file name
    ```bash
        jsonDataFiles:
            pathFile: {{jsonDataFiles-pathFile}}
            nameFile: {{jsonDataFiles-nameFile}}
    ```
5. config jwt-key secret 
    ```bash
       secrets:
        jwt-key: {{secrets-jwt-key}}
    ```
### Run Go
1. run install all package.

    ```bash
    go mod tidy
    ```

    2. start go server.

    ```bash
    go run cmd/main.go
    ```
### test Go
1. run unit test all files and display coverage.

    ```bash
    go test ./... -cover
    ```

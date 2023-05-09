### Membuat migrations  
Pastikan sudah terdapat package migration di environment

1. Install package migration
    ```shell
    go install -tags ‘mysql,postgres’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```

2. membuat table migrations
    ```shell
    migrate create -ext sql -dir db/migrations <create_table_users>
    ```

3. Menjalankan migration up
    ```shell
     migrate -database "postgres://postgres:postgres@tcp(localhost:5432)/pmsapi" -path db/migrations up 
     #or
     migrate -database "postgres://postgres:postgres@localhost/authapi?sslmode=disable" -path db/migrations up
    ```
4. Menjalankan migration down / menurunkan versi migration
    ```shell
     migrate -database "<dirver database>://<username>:<password>@tcp(<host>:<port>)/<db name>" -path db/migrations down
    ```

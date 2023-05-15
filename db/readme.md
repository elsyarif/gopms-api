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
```sql
INSERT INTO groups
VALUES ('group-5cB1aj3CnSx4hBYU3PKDQ', 'Sinter', 'Sinter plant')

insert into servers
values ('server-R7wY5VHqD3B0NW0iuTHls', 'group-5cB1aj3CnSx4hBYU3PKDQ', 'SINTER1A', 'Data Center', 'online', 64, '172.21.83.31'),
('server-whqjPmXzWYrI3sIEZ0hOn', 'group-5cB1aj3CnSx4hBYU3PKDQ', 'SINTER1B', 'Data Center', 'Backup', 64, '172.21.83.31'),
('server-2t2jMjFtvFKPSXdPFlChV', 'group-5cB1aj3CnSx4hBYU3PKDQ', 'SINTERAI', 'Data Center', 'Online', 256, '172.21.83.31'),
('server-XUNbsTPkR22YETpn6C33Z', 'group-5cB1aj3CnSx4hBYU3PKDQ', 'IOSVR01', 'Sinter Plant ICT Computer Room', 'Standby', 8, '172.21.83.31'),
('server-kwdR9Ib7PtM06pdM8KzEY', 'group-5cB1aj3CnSx4hBYU3PKDQ', 'IOSVR02', 'Sinter Plant ICT Computer Room', 'Online', 8, '172.21.83.31'),
('server-Rf5ZTSL1Sls6Sn72cp3cd', 'group-5cB1aj3CnSx4hBYU3PKDQ', 'GATHERING PC', 'Sinter Plant ICT Computer Room', 'Online', 8, '172.21.83.31');

insert into disks
values ('disk-zpp962dSBU5A7OZiuDoUe', 'server-R7wY5VHqD3B0NW0iuTHls', 'C:', 278),
('disk-zFN7i7larVk1yWykfArlT', 'server-R7wY5VHqD3B0NW0iuTHls', 'D:', 119),
('disk-fNXvgHifgab3gWONpAQ0V', 'server-R7wY5VHqD3B0NW0iuTHls', 'E:', 239),
('disk-SMtLmsQmiUuXiq90xczl1', 'server-R7wY5VHqD3B0NW0iuTHls', 'F:', 196),

       ('disk-oiuxcc3Gl4eB4VDOsjht4', 'server-whqjPmXzWYrI3sIEZ0hOn', 'C:', 278),
       ('disk-o3iPCQ1cczB47BytcQhxS', 'server-whqjPmXzWYrI3sIEZ0hOn', 'D:', 119),
       ('disk-GEznDud0NbYBgTcLU1s7r', 'server-whqjPmXzWYrI3sIEZ0hOn', 'E:', 0),
       ('disk-jMw3Ea9BD4tT10r7b3HAY', 'server-whqjPmXzWYrI3sIEZ0hOn', 'F:', 196),

       ('disk-J9kXjeYnZYTDOwSh2XhU8', 'server-2t2jMjFtvFKPSXdPFlChV', 'C:', 1810),
       ('disk-YtkpgKK1ZuoHkTxpskcKS', 'server-2t2jMjFtvFKPSXdPFlChV', 'D:', 1810),

       ('disk-6YDIosQQbB4TgMZW969QT', 'server-XUNbsTPkR22YETpn6C33Z', 'C:', 440),
       ('disk-pXGBNKORlXwYUkaYmLzVI', 'server-XUNbsTPkR22YETpn6C33Z', 'D:', 458);
```
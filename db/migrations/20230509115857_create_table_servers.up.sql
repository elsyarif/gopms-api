CREATE TABLE servers (
    id          varchar(36) PRIMARY KEY NOT NULL,
    group_id    varchar(36)             NOT NULL,
    server_name varchar(50)             NOT NULL,
    location    varchar(50)             NOT NULL,
    status      varchar(20)             NOT NULL,
    memory      int                     NOT NULL,
    ip          varchar(20)             NOT NULL
);

ALTER TABLE servers ADD FOREIGN KEY ("group_id") REFERENCES groups ("id");
CREATE TABLE disks (
   id varchar(36) PRIMARY KEY NOT NULL,
   server_id varchar(36) NOT NULL,
   name varchar(10) NOT NULL,
   total int NOT NULL
);

ALTER TABLE disks ADD FOREIGN KEY ("server_id") REFERENCES servers ("id");
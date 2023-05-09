CREATE TABLE permissions (
     id varchar(36) PRIMARY KEY NOT NULL,
     name varchar(50) NOT NULL,
     description varchar(150)
);

CREATE TABLE permission_role (
     permission_id varchar(36) NOT NULL,
     role_id varchar(36) NOT NULL
);

ALTER TABLE permission_role ADD FOREIGN KEY ("permission_id") REFERENCES permissions("id");
ALTER TABLE permission_role ADD FOREIGN KEY ("role_id") REFERENCES roles("id");

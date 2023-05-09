CREATE TABLE roles (
   id varchar(36) PRIMARY KEY NOT NULL,
   name varchar(50) NOT NULL,
   description varchar(100)
);

CREATE TABLE user_role (
   user_id varchar(36),
   role_id varchar(36)
);

ALTER TABLE user_role ADD FOREIGN KEY ("user_id") REFERENCES users("id");
ALTER TABLE user_role ADD FOREIGN KEY ("role_id") REFERENCES roles("id");

CREATE TABLE groups (
    id varchar(36) PRIMARY KEY NOT NULL,
    name varchar(50) NOT NULL,
    description varchar(150)
);

CREATE TABLE group_user (
    user_id varchar(36) NOT NULL,
    group_id varchar(36) NOT NULL
);

ALTER TABLE group_user ADD FOREIGN KEY ("user_id") REFERENCES users("id");
ALTER TABLE group_user ADD FOREIGN KEY ("group_id") REFERENCES groups("id");
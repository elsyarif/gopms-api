CREATE TABLE users (
    id varchar(36) primary key not null,
    name varchar(50) not null,
    username varchar(36) not null,
    password varchar(200) not null,
    is_active BOOLEAN default true,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    unique (username)
);


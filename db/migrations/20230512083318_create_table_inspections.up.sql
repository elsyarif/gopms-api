CREATE TABLE inspections (
    id varchar(36) primary key not null,
    group_id varchar(36) not null ,
    group_name varchar(50) not null ,
    date date not null,
    user_by varchar(36) not null ,
    period_start date not null ,
    period_end date not null
);

CREATE TABLE inspection_server (
    id varchar(36) primary key not null,
    inspection_id varchar(36) not null ,
    server_id varchar(36) not null ,
    server_name varchar(50) not null ,
    cpu_usage int not null default 0,
    memory_usage int not null default 0
);

CREATE TABLE inspection_disk (
    id varchar(36) primary key not null,
    inspection_id varchar(36) not null ,
    inspection_server_id varchar(36) not null ,
    disk_id varchar(36) not null ,
    disk_name varchar(10) not null ,
    disk_usage int not null default 0
);

ALTER TABLE inspections ADD FOREIGN KEY ("group_id") REFERENCES groups("id");
ALTER TABLE inspection_server ADD FOREIGN KEY ("inspection_id") REFERENCES inspections("id");
ALTER TABLE inspection_server ADD FOREIGN KEY ("server_id") REFERENCES servers("id");
ALTER TABLE inspection_disk ADD FOREIGN KEY ("inspection_id") REFERENCES inspections("id");
ALTER TABLE inspection_disk ADD FOREIGN KEY ("inspection_server_id") REFERENCES inspection_server("id");
ALTER TABLE inspection_disk ADD FOREIGN KEY ("disk_id") REFERENCES disks("id");

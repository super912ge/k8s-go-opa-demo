ALTER TABLE team ADD COLUMN name varchar(255) NOT NULL;
ALTER TABLE team ADD COLUMN description varchar(255) NOT NULL;
ALTER TABLE team ADD COLUMN manager_id integer;
alter table team ADD constraint fk_manager_user FOREIGN KEY (manager_id) REFERENCES "user" (id);
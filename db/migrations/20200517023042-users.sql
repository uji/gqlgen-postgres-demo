
-- +migrate Up
create table if not exists users (
  id int not null primary key,
  name varchar(255) not null
);

-- +migrate Down
drop table if exists users;

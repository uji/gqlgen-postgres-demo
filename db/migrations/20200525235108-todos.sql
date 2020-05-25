
-- +migrate Up
create table if not exists todos (
  id int not null primary key,
  text text,
  done boolean not null,
  name varchar(255) not null
);

-- +migrate Down
drop table if exists todos;

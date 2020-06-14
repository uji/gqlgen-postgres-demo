
-- +migrate Up
alter table todos
add user_id int;

-- +migrate Down
alter table todos
drop column user_id;

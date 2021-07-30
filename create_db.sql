CREATE table users
(
    id         int generated always as identity,
    first_name varchar(40),
    last_name varchar(40),
    age int,
    creation_date timestamp with time zone default current_timestamp
)
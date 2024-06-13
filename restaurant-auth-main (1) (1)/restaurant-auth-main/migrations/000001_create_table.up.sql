create table users(
    id uuid primary key default gen_random_uuid(),
    user_name varchar(30) unique,
    password varchar(30),
    email varchar(50) unique,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delated_at bigint default 0
);
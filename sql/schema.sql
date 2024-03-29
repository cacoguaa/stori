create table if not exists transactions
(
    id BIGINT PRIMARY KEY,
    date text not null,
    value float8 not null
);
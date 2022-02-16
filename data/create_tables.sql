CREATE TABLE IF NOT EXISTS cookies (
    name  TEXT
        constraint cookies_pk
            primary key,
    value TEXT not null
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    chat_id integer
        constraint subscriptions_pk
            primary key
);

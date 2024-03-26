create table users
(
    id       serial       not null primary key,
    username varchar(50)  not null unique,
    password varchar(255) not null,
    status   int
);

create table sessions
(
    session_id varchar(255) not null primary key,
    user_id    bigint       not null,
    login_time timestamp    not null,
    status     int
)
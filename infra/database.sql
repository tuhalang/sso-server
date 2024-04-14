create table users
(
    id       serial       not null primary key,
    username varchar(50)  not null unique,
    password varchar(255) not null,
    status   int
);
-- $2a$10$B6804z8J5G/KesCaT6oLzeshHeLg.5Rsz5s1oF0fIcKtCZAPIZAdO

create table sessions
(
    session_id varchar(255) not null primary key,
    user_id    bigint       not null,
    login_time timestamp    not null,
    status     int
)
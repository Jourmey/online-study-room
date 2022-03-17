-- // 房间表
create table room
(
    id          integer
        primary key,
    created_at  datetime,
    updated_at  datetime,
    deleted_at  datetime,
    name        text,
    seat_number integer
);

create index idx_room_deleted_at
    on room (deleted_at);

INSERT INTO room (id, created_at, updated_at, deleted_at, name, seat_number) VALUES (1, null, null, null, 'room1', 10);


-- 座位表
create table seat
(
    id         integer
        primary key,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    room_id    integer,
    user_id    text,
    entered_at datetime,
    until      datetime,
    work_name  text,
    color_code text,
    x          integer,
    y          integer,
    wide       integer,
    high       integer
);

create index idx_seat_deleted_at
    on seat (deleted_at);

create index seat_room_id_index
    on seat (room_id);

create index seat_user_id_index
    on seat (user_id);

INSERT INTO seat (id, created_at, updated_at, deleted_at, room_id, user_id, entered_at, until, x, y, wide, high, work_name, color_code) VALUES (1, null, null, null, 1, '1', '2022-03-16 19:46:04.820127+08:00', '2022-03-16 17:46:04.820127+08:00', 0, 0, 10, 10, 'jourmey', 'blue');
INSERT INTO seat (id, created_at, updated_at, deleted_at, room_id, user_id, entered_at, until, x, y, wide, high, work_name, color_code) VALUES (2, null, null, null, 1, null, null, null, 10, 0, 10, 10, null, null);


-- 用户表
create table user
(
    id         integer
        primary key,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    user_name  text not null,
    nike_name  text not null,
    password   text not null
);

create index idx_user_deleted_at
    on user (deleted_at);

create unique index user_user_name_uindex
    on user (user_name);

INSERT INTO user (id, created_at, updated_at, deleted_at, user_name, nike_name, password) VALUES (1, null, null, null, '15735181677', 'jourmey', '123456');
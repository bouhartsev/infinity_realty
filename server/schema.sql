drop table if exists clients;
create table if not exists clients
(
    id         int(11) auto_increment primary key,
    name       text not null,
    surname    text not null,
    patronymic text not null,
    tel        text,
    email      text
);

drop table if exists realtors;
create table if not exists realtors
(
    id         int(11) auto_increment primary key,
    name       text not null,
    surname    text not null,
    patronymic text not null,
    commission float4
);

drop table if exists properties;
create table if not exists properties
(
    id               int(11) auto_increment primary key,
    type             int2 not null,
    address_city     text,
    address_street   text,
    address_building text,
    address_floor    text,
    cords_latitude   float4,
    cords_longitude  float4,
    floor            int2,
    floor_count      int2,
    room_count       int2,
    square           int2
);

drop table if exists offers;
create table if not exists offers
(
    id          int(11) auto_increment primary key,
    client_id   int4 not null references clients (id),
    realtor_id  int4 not null references realtors (id),
    property_id int4 not null references properties (id),
    price       int  not null,
    confirmed   bool default false
);

drop table if exists needs;
create table if not exists needs
(
    id                        int(11) auto_increment primary key,
    client_id                 int4 not null references clients (id),
    realtor_id                int4 not null references realtors (id),
    property_type             int2 not null,
    property_address_city     text,
    property_address_street   text,
    property_address_building text,
    property_address_floor    text,
    min_floor                 int2,
    max_floor                 int2,
    min_room_count            int2,
    max_room_count            int2,
    min_square                int2,
    max_square                int2,
    min_floor_count           int2,
    max_floor_count           int2,
    min_price                 int  not null,
    max_price                 int  not null,
    confirmed                 bool default false
);

drop table if exists deals;
create table if not exists deals
(
    id       int(11) auto_increment primary key,
    need_id  int4 not null references needs (id),
    offer_id int4 not null references offers (id)
);

drop table if exists events;
create table if not exists events
(
    id         int(11) auto_increment primary key,
    realtor_id int4      not null references realtors (id),
    datetime   timestamp not null,
    duration   time      not null,
    type       int2      not null,
    comment    text
);
create table if not exists users (
    id serial primary key,
    role smallint not null,
    name text,
    surname text,
    patronymic text,
    tel text,
    email text,
    commission float4,
    password text not null,

    constraint either_email check ((email is null) <> (tel is null))
);

create table if not exists properties (
    id serial primary key,
    type int2 not null,
    address_city text,
    address_street text,
    address_building text,
    address_floor text,
    cords_latitude float4,
    cords_longitude float4,
    floor int2,
    room_count int2,
    square int2,
    floor_count int2
);

create table if not exists offers (
    id serial primary key,
    client_id int4 not null references users(id),
    realtor_id int4 not null references users(id),
    property_id int4 not null references properties(id),
    price int not null
);

create table if not exists needs (
      id serial primary key,
      client_id int4 not null references users(id),
      property_type int2 not null,
      property_address_city text,
      property_address_street text,
      property_address_building text,
      property_address_floor text,
      min_floor int2,
      max_floor int2,
      min_room_count int2,
      max_room_count int2,
      min_square int2,
      max_square int2,
      min_floor_count int2,
      max_floor_count int2,
      min_price int not null,
      max_price int not null
);

create table if not exists deals (
    id serial primary key,
    need_id int4 not null references needs(id),
    offer_id int4 not null references offers(id),
    company_commission float4 not null,
    realtor_commission float4 not null,
    client_commission float4 not null
);

create table if not exists events (
     id serial primary key,
     realtor_id int4 not null references users(id),
     datetime timestamp not null,
     duration time not null,
     type int2 not null,
     comment text
);
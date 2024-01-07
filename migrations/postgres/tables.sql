create table driver(
    id uuid primary key ,
    full_name text,
    phone text
);

create table car(
    id uuid primary key ,
    model varchar(30),
    brand varchar(30),
    year integer,
    driver_id uuid references driver(id) ON DELETE CASCADE
);
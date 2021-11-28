CREATE TABLE product
(
    id             serial         not null unique ,
    title           varchar(255)   not null,
    description    varchar (255),
    created_at     int            not null
);
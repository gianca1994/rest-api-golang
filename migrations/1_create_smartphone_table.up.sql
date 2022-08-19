CREATE TABLE smartphone
(
    id             int          not null auto increment,
    name           varchar(150) not null,
    country_origin varchar(100) not null,
    os             varchar(100) not null,
    price          int(10) not null
)
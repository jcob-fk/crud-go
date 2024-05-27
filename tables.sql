CREATE TABLE usuario (
    id int auto_increment, 
    nombre varchar(200) not null,
    email varchar(200) not null,
    password varchar(200) not null,
    activo int DEFAULT(1),
    PRIMARY KEY(id)
);

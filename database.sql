

create table users(
    id INT GENERATED ALWAYS AS IDENTITY,
    email varchar(50) unique  not null,
    date_created timestamp default CURRENT_TIMESTAMP,
    first_name varchar(50) unique  not null,
    middle_name varchar(50) ,
    phone_number varchar(50),
    password text not null,
    firebase_id text,

    primary key (id)
)

create table distance(
    user_id int,
                         current_latitude double precision,
                         current_longitude double precision,
                         max_distance double precision,
                         current_latitude double precision,
                         current_longitude double precision,

    primary key (user_id),

    CONSTRAINT fk_users
        FOREIGN KEY(user_id)
            REFERENCES product(id)
)
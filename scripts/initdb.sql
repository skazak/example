DROP TABLE IF EXISTS "user";
CREATE TABLE "user"
(
    id    serial PRIMARY KEY,
    email varchar,
    password varchar
);

DROP TABLE IF EXISTS category;
CREATE TABLE category
(
    id    serial PRIMARY KEY,
    title varchar
);

DROP TABLE IF EXISTS item;
CREATE TABLE item
(
    id          serial PRIMARY KEY,
    title       varchar,
    category_id int REFERENCES category(id) ON DELETE CASCADE
);




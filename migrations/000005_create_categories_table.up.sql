create table "categories" (
    "id" serial primary key,
    "name" varchar(80)
);

insert into "categories"
("name")
values
('Music'),
('Arts'),
('Outdoors'),
('Worskshop'),
('Sport'),
('Festival'),
('Fashion');
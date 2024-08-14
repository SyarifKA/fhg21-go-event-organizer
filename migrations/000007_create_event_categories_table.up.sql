create table "event_categories" (
    "id" serial primary key,
    "event_id" int references "events"("id"),
    "category_id" int references "categories"("id")
);

insert into "event_categories"
("event_id", "category_id")
values
(1,5),(1,3),(2,5),(3,6),(4,6),(5,1),(6,6),(7,1),(7,6),(8,4),(8,6),(9,6),(9,3),(10,2);
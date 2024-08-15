create table "event_sections"(
    "id" serial primary key,
    "name" varchar(50),
    "price" int,
    "quantity" int,
    "event_id" int references "events"("id")
);
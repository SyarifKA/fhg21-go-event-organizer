create table "events" (
    "id" serial primary key,
    "image" varchar(255),
    "title" varchar(50),
    "date" varchar(30),
    "description" text,
    -- "location_id" int references "locations"("id"),
    "location_id" int,
    "created_by" int references "users"("id")
);
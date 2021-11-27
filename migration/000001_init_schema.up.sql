CREATE TABLE "product" (
    "id"            serial PRIMARY KEY,
    "title"         varchar(255) NOT NULL,
    "description"   varchar(500),
    "created_at"    int
)
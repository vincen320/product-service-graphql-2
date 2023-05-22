CREATE TABLE IF NOT EXISTS products(
    "id" bigserial PRIMARY KEY,
    "name" varchar,
    "description" varchar,
    "price" bigint,
    "created_by" bigint,
    "created_at" timestamp
);

CREATE TABLE IF NOT EXISTS product_attributes(
    "id" bigserial PRIMARY KEY,
    "product_id" bigint REFERENCES products(id),
    "attribute" varchar,
    "value" varchar
)

CREATE INDEX ON products(name)
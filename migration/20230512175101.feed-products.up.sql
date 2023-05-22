INSERT INTO products(
    "name"
    , "description"
    , "price"
    , "created_by"
    , "created_at"
) VALUES
('T-Shirt', 'nice tshirt for daily', 50000, 1, now()),
('Shirt', 'nice tshirt for night prom', 100000, 1, now()),
('Truck', 'up to 500kg load', 500000000, 1, now());

INSERT INTO product_attributes(
    "product_id"
    , "attribute"
    , "value"
) VALUES
(1, 'material', 'cotton'),
(1, 'size', 'XL'),
(2, 'material', 'cotton'),
(2, 'size', 'M'),
(3, 'engine', '5000cc'),
(3, 'wheel', '4')
ALTER TABLE products
    ADD COLUMN "type" varchar

UPDATE products p
    SET "type" = 1
    WHERE EXISTS(
        SELECT 1 FROM product_attributes pa
        WHERE p.id = pa.product_id
        AND pa.attribute IN ('material', 'size')
    )

UPDATE products p
    SET "type" = 2
    WHERE EXISTS(
        SELECT 1 FROM product_attributes pa
        WHERE p.id = pa.product_id
        AND pa.attribute IN ('engine', 'wheel')
    )

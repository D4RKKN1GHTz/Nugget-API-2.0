SET statement_timeout = 0;

CREATE TABLE cart_items (
    id VARCHAR DEFAULT gen_random_uuid() PRIMARY KEY,
    cart_id VARCHAR NOT NULL,
    product_id VARCHAR NOT NULL,
    amount      INT NOT NULL,
    total_price INT NOT NULL
);

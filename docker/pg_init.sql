
-- If table already exists, drop it
DROP TABLE IF EXISTS menu;
-- Create a new table
CREATE TABLE menu (
  item_id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  price MONEY NOT NULL
);

INSERT INTO menu (
  name, price
) 
VALUES (
  ('focaccia', 5.00),
  ('biancaneve', 5.50),
  ('margherita', 6.5)
)
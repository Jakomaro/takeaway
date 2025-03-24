
-- If table already exists, drop it
DROP TABLE IF EXISTS menu;
-- Create a new table
CREATE TABLE menu (
  item_id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  price MONEY NOT NULL
);

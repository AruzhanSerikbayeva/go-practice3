
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name text not null, 
    user_id INTEGER REFERENCES users(id) on delete CASCADE,
    CONSTRAINT unique_user_category unique (user_id, name)
);

CREATE INDEX idx_categories_user_id ON categories(user_id);

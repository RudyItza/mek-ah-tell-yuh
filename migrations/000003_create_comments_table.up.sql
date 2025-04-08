CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    story_id INT REFERENCES stories(id),
    user_id INT REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS videos (id VARCHAR(255), title
VARCHAR(255), description TEXT, author VARCHAR(255));
CREATE TABLE IF NOT EXISTS ratings (record_id VARCHAR(255),
record_type VARCHAR(255), user_id VARCHAR(255), value INT);

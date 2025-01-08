CREATE UNIQUE INDEX users_email_index ON users (email);
CREATE INDEX users_status_index ON users (status);
CREATE INDEX users_deleted_at_index ON users (deleted_at);
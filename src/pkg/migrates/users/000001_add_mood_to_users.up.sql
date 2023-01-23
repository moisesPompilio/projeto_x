CREATE TABLE IF NOT EXISTS users (
	id uuid,
	nick_name varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	PRIMARY KEY (id)
);
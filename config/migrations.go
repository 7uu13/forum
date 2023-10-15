package config

const UserTable = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	age INTEGER NOT NULL DEFAULT 0,
	gender TEXT NOT NULL,
	firstname TEXT NOT NULL,
	lastname TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
);
`

const PostTable = `
CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created DATETIME CURRENT_TIMESTAMP,
	user_id INTEGER,
	foreign key (user_id) REFERENCES users (id)
);		
`

const CategoryTable = `
CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	name_slug TEXT NOT NULL
);
`

const PostCategoryTable = `
CREATE TABLE IF NOT EXISTS posts_category (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	post_id INTEGER,
	category_id INTEGER,
	foreign key (post_id) REFERENCES posts (id),
	foreign key (category_id) REFERENCES categories (id)
);
`

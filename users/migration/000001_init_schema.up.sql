CREATE TABLE IF NOT EXISTS "users" (
    "id" serial PRIMARY KEY,
    "firstname" varchar NOT NULL,
    "secondname" varchar NOT NULL,
    "login" varchar UNIQUE NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL
);
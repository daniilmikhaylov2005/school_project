CREATE TABLE IF NOT EXISTS "magazine" (
    "magazine_code" int NOT NULL UNIQUE,
    "teacher_login" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "kid" (
    "magazine_code" int NOT NULL,
    "id" serial PRIMARY KEY,
    "fullname" varchar NOT NULL,
    "age" integer NOT NULL,
    "graduate" integer NOT NULL,
    FOREIGN KEY ("magazine_code")  REFERENCES "magazine" ("magazine_code") ON DELETE CASCADE
);
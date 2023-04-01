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

CREATE TABLE IF NOT EXISTS "grades" (
    "kid_id" int NOT NULL,
    "date" timestamp NOT NULL,
    "subject" varchar NOT NULL,
    "grade" int NOT NULL,
    FOREIGN KEY ("kid_id") REFERENCES "kid" ("id") ON DELETE CASCADE
);
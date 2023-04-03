CREATE TABLE IF NOT EXISTS "teacher_history" (
    "teacher_login" varchar NOT NULL,
    "action" text NOT NULL,
    FOREIGN KEY ("teacher_login") REFERENCES "users" ("login") ON DELETE CASCADE
);
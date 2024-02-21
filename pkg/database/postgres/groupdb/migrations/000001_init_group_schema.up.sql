CREATE TABLE "groups" (
  "group_id" serial PRIMARY KEY,
  "group_name" varchar UNIQUE NOT NULL,
  "group_creator_id" integer NOT NULL,
  "photo_url" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "group_members" (
  "group_id" serial REFERENCES "groups" ("group_id") ON DELETE CASCADE,
  "member_id" integer NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("group_id", "member_id")
);

CREATE TABLE "posts" (
  "post_id" serial PRIMARY KEY,
  "member_id" integer NOT NULL,
  "group_id" serial NOT NULL,
  "photo_url" varchar,
  "description" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "post_reactions" (
  "reaction_id" serial PRIMARY KEY,
  "post_id" serial NOT NULL,
  "reaction" varchar NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "post_comments" (
  "comment_id" serial PRIMARY KEY,
  "post_id" integer NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
)
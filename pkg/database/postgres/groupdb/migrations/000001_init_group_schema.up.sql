CREATE TABLE "groups" (
  "group_id" serial PRIMARY KEY,
  "group_name" varchar UNIQUE NOT NULL,
  "group_creator_id" integer NOT NULL,
  "photo_url" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "group_members" (
  "group_id" serial REFERENCES "groups" ("group_id") ON DELETE CASCADE,
  "member_id" integer NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY ("group_id", "member_id")
);

CREATE TABLE "posts" (
  "post_id" serial PRIMARY KEY,
  "member_id" integer NOT NULL,
  "group_id" serial NOT NULL REFERENCES "groups" ("group_id") ON DELETE CASCADE,
  "photo_url" varchar,
  "description" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "post_reactions" (
  "reaction_id" serial PRIMARY KEY,
  "post_id" serial NOT NULL REFERENCES "posts" ("post_id") ON DELETE CASCADE,
  "member_id" integer NOT NULL,
  "reaction" varchar NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "post_comments" (
  "comment_id" serial PRIMARY KEY,
  "post_id" integer NOT NULL REFERENCES "posts" ("post_id") ON DELETE CASCADE,
  "member_id" integer NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "tags" (
    "tag_id" serial PRIMARY KEY,
    "tag_name" varchar UNIQUE NOT NULL,
    "icon_url" varchar
);

CREATE TABLE "group_tags" (
    "group_id" integer NOT NULL,
    "tag_id" integer NOT NULL,
    PRIMARY KEY ("group_id", "tag_id"),
    FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id"),
    FOREIGN KEY ("tag_id") REFERENCES "tags" ("tag_id")
);
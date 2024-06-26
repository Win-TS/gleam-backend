CREATE TABLE "groups" (
  "group_id" SERIAL PRIMARY KEY,
  "group_name" varchar UNIQUE NOT NULL,
  "group_creator_id" integer NOT NULL,
  "description" varchar,
  "photo_url" varchar,
  "tag_id" integer NOT NULL,
  "frequency" integer NOT NULL,
  "max_members" integer NOT NULL DEFAULT 25,
  "group_type" varchar NOT NULL DEFAULT 'social',
  "visibility" boolean NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "group_members" (
  "group_id" integer NOT NULL,
  "member_id" integer NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  PRIMARY KEY ("group_id", "member_id")
);

CREATE TABLE "group_requests" (
  "group_id" integer NOT NULL,
  "member_id" integer NOT NULL,
  "description" varchar,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  PRIMARY KEY ("group_id", "member_id")
);

CREATE TABLE "posts" (
  "post_id" SERIAL PRIMARY KEY,
  "member_id" integer NOT NULL,
  "group_id" SERIAL NOT NULL,
  "photo_url" varchar,
  "description" varchar,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "post_reactions" (
  "reaction_id" SERIAL PRIMARY KEY,
  "post_id" SERIAL NOT NULL,
  "member_id" integer NOT NULL,
  "reaction" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "post_comments" (
  "comment_id" SERIAL PRIMARY KEY,
  "post_id" integer NOT NULL,
  "member_id" integer NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "tags" (
  "tag_id" SERIAL PRIMARY KEY,
  "tag_name" varchar UNIQUE NOT NULL,
  "icon_url" varchar,
  "category_id" integer
);

CREATE TABLE "streak_set" (
  "streak_set_id" SERIAL PRIMARY KEY,
  "group_id" INTEGER NOT NULL,
  "member_id" integer NOT NULL,
  "end_date" TIMESTAMP NOT NULL,
  "start_date" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "streaks" (
  "streak_id" SERIAL PRIMARY KEY,
  "streak_set_id" INTEGER NOT NULL,
  "max_streak_count" INTEGER NOT NULL DEFAULT 0,
  "total_streak_count" INTEGER NOT NULL DEFAULT 0,
  "weekly_streak_count" INTEGER NOT NULL DEFAULT 0,
  "completed" boolean NOT NULL DEFAULT false,
  "recent_date_added" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "tag_category" (
  "category_id" SERIAL PRIMARY KEY,
  "category_name" VARCHAR UNIQUE NOT NULL
);

ALTER TABLE "group_members" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id") ON DELETE CASCADE;

ALTER TABLE "group_requests" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id") ON DELETE CASCADE;

ALTER TABLE "posts" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id") ON DELETE CASCADE;

ALTER TABLE "post_reactions" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id") ON DELETE CASCADE;

ALTER TABLE "post_comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id") ON DELETE CASCADE;

ALTER TABLE "groups" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("tag_id");

ALTER TABLE "streaks" ADD FOREIGN KEY ("streak_set_id") REFERENCES "streak_set" ("streak_set_id") ON DELETE CASCADE;

ALTER TABLE "streak_set" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id") ON DELETE CASCADE;

ALTER TABLE "tags" ADD FOREIGN KEY ("category_id") REFERENCES "tag_category" ("category_id");
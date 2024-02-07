-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "comments" (
    "id" bigint NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "body" text NOT NULL,
    "user_id" bigint NOT NULL,
    "post_id" bigint NOT NULL,
    "created_at" TIMESTAMP NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL
);

ALTER TABLE "comments" ADD CONSTRAINT "comments_user_id_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "comments_user_id" ON "comments" ("user_id");

ALTER TABLE "comments" ADD CONSTRAINT "comments_post_id_fk_post_id" FOREIGN KEY ("post_id") REFERENCES "blog_posts" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "comments_post_id" ON "comments" ("post_id");
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "comments" CASCADE;
COMMIT;
-- +goose StatementEnd

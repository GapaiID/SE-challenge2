-- +goose Up
-- +goose StatementBegin
CREATE TABLE "followers" (
    "id" bigint NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "follower_id" integer NOT NULL,
    "user_id" integer NOT NULL,
    "created_at" TIMESTAMP NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL
);

ALTER TABLE "followers" ADD CONSTRAINT "followers_user_id_follower_id_uniq" UNIQUE ("user_id", "follower_id");
ALTER TABLE "followers" ADD CONSTRAINT "followers_follower_id_user_id" FOREIGN KEY ("follower_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "followers" ADD CONSTRAINT "followers_user_id_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "followers_follower_id" ON "followers" ("follower_id");
CREATE INDEX "followers_user_id" ON "followers" ("user_id");
COMMIT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "followers";
-- +goose StatementEnd

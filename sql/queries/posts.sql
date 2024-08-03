-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts
WHERE feed_id IN (
    SELECT feed_id FROM feed_follows
    WHERE user_id = $1
)
ORDER BY published_at DESC
LIMIT $2;

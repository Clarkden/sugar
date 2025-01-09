-- name: CreateCoupon :one
INSERT INTO coupons (code, domain) VALUES (?, ?) RETURNING code;

-- name: GetCouponsByDomain :many
SELECT code FROM coupons WHERE domain = ?;

-- name: DeleteCoupon :exec
DELETE FROM coupons WHERE code = ?;
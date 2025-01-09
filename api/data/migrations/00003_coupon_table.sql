-- +goose Up
-- +goose StatementBegin
CREATE TABLE coupons (
    code TEXT UNIQUE,
    domain TEXT

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coupons;
-- +goose StatementEnd

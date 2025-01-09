// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sugar

type Coupon struct {
	Code   *string `json:"code"`
	Domain *string `json:"domain"`
}

type Session struct {
	UserID    *int64  `json:"user_id"`
	SessionID *string `json:"session_id"`
	CreatedAt *int64  `json:"created_at"`
	ExpiresAt *int64  `json:"expires_at"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

package model

type Auth struct {
	ID       int    `db:"id"`
	Token    string `db:"token"`
	AuthType string `db:"auth_type"`
	UserID   int    `db:"user_id"`
}

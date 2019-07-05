package ports

type ProfileOutputPort struct {
	UserID       int    `json:"user_id"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
}

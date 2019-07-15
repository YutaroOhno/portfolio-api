package ports

import (
  "portfolio-api/entities"
)

type ProfileOutputPort struct {
	UserID       int    `json:"user_id"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	Link         []entities.Link `json:"links"`
}

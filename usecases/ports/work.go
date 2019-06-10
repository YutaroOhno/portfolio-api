package ports

import (
	"portfolio-api/entities"
)

type WorksOutputPort struct {
	Works []entities.Work `json:"works"`
}
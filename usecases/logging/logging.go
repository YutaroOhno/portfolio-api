package logging

import (
	"portfolio-api/usecases"
)

type Logging interface {
	Error(*usecases.UError)
	Warning(*usecases.UError)
	Info(string)
	Debug(string)
}

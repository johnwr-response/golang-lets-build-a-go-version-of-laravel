package middleware

import (
	"github.com/johnwr-response/celeritas"
	"myapp/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}

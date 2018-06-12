package controller

import (
	"github.com/alustau/issuer/database"
)

//Controller it's a handler
type Controller struct {
	DB *database.Connection
}

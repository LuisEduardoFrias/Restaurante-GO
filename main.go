package main

import (
	"github.com/LuisEduardoFrias/Restaurante/Route"
	"github.com/LuisEduardoFrias/Restaurante/Server"
)

func main() {
	Server.Localhost(Route.Routes())
}

package main

import (
	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Route"
	"github.com/LuisEduardoFrias/Restaurante/GoBackEnd/Server"
)

func main() {
	Server.Localhost(Route.Routes())
}

package main

import (
    "goApp/routes"
)

func main() {
    r := routes.InitRoutes()
    r.Run(":8080")
}

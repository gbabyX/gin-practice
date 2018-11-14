package main

import "./routes"

func main() {
	r := routes.InitRouter()
	r.Run()
}

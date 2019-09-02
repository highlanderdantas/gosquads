package main

import (
	"gopkg.in/macaron.v1"

	"log"

	conf "github.com/highlanderdantas/gosquads/backend/conf/app"
)

// application entrypoint
func main() {
	app := macaron.New()
	conf.SetupMiddlewares(app)
	conf.SetupRoutes(app)
	log.Println(`
	_____                                       _      
	|  __ \                                     | |     
	| |  \/  ___   ___   __ _  _   _   __ _   __| | ___ 
	| | __  / _ \ / __| / _  || | | | / _  | / _  |/ __|
	| |_\ \| (_) |\__ \| (_| || |_| || (_| || (_| |\__ \
	 \____/ \___/ |___/ \__, | \__,_| \__,_| \__,_||___/
						   | |                          
						   |_|  
				`)

	app.Run(8080)
}

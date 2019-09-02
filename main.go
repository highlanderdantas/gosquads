package main

import (
	"gopkg.in/macaron.v1"

	"log"

	conf "github.com/highlanderdantas/gosquads/conf/app"
)

// application entrypoint
func main() {
	app := macaron.New()
	conf.SetupMiddlewares(app)
	conf.SetupRoutes(app)
	/*
		Generated using http://www.kammerl.de/ascii/AsciiSignature.php - (Font: 'starwars')
		All signatures are made with FIGlet (c) 1991, 1993, 1994 Glenn Chappell and Ian Chai
		All fonts are taken from figlet.org and jave.de.
		Please check for Font Credits the figlet font database!
		Figlet Frontend - Written by Julius Kammerl - 2005
	*/
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

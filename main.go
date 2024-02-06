package main

import "github.com/GapaiID/SE-challenge2/cmd"

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	cmd.Execute()
}

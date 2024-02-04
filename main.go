package main

import "github.com/agung96tm/miblog/cmd"

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	cmd.Execute()
}

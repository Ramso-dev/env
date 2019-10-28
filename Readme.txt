package main

import (
	"veryverybadbot/badbot-envvars/env"
)

func main() {

	local := map[string]string{
		"ASD":        "asddsads",
		"TTT":        "xxxx"}

	cloudtest := map[string]string{
		"ASD":        "asddsads",
		"TTT":        "xxxx"}

	prod := map[string]string{
		"ASD":        "asddsads",
		"TTT":        "xxxx"}
	env.InitEnvVars(local, cloudtest, prod)
}
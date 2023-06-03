package utils

import "log"

/*
Check if err != nil
*/
func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

package main

import "os"

func AddRole(text string) {
	file, err := os.OpenFile("role.json", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	file.WriteString(text)
}

package main

import "fmt"

func getGun(gunTYpe string) (IGun, error) {
	if gunTYpe == "ak47" {
		return newAk47(), nil
	}
	if gunTYpe == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

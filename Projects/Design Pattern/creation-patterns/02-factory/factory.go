package main

import "fmt"

func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	if pubType == "newspaper" {
		return createNewsaper(name, pg, pub), nil
	}
	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}
	return nil, fmt.Errorf("Publication türü konusunda saçmaladın!")
}

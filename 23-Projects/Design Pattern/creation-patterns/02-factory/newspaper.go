package main

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("Bu %s gazetesidir.", n.name)
}

func createNewsaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}

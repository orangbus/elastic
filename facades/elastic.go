package facades

import (
	"log"

	"goravel/packages/elastic"
	"goravel/packages/elastic/contracts"
)

func Elastic() contracts.Elastic {
	instance, err := elastic.App.Make(elastic.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Elastic)
}

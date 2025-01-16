package facades

import (
	"log"

	"github.com/orangbus/elastic"
	"github.com/orangbus/elastic/contracts"
)

func Elastic() contracts.Elastic {
	instance, err := elastic.App.Make(elastic.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Elastic)
}

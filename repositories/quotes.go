package repositories

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/georgizhivankin/go-slacking-cow/models"
	"gopkg.in/mgo.v2"
)

func SaveQuote(quote models.Quote) bool {
	// Connect to Mongo
	config, err := config.LoadEnv()
	if err != nil {
		log.Panic(err)
	}
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s/%s", config.Database.DSN, config.Database.Port, config.Database.Name))

	// Exit with an error
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	if err != nil {
		log.Panic(err)
	}
	result := models.Quote{}

	session.DB(config.Database.Name).C("quotes").Insert(quote)
	err = session.DB(config.Database.Name).C("quotes").FindId(quote.Id).One(&result)

	if err != nil {
		return false
	}

	return true
}

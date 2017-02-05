package repositories

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/georgizhivankin/go-slacking-cow/models"
	"gopkg.in/mgo.v2"
)

func SaveQuote(quote models.Quote) models.Quote {
	// Connect to Mongo
	config, err := config.LoadEnv()
	if err != nil {
		log.Panic(err)
	}
	session, err := mgo.DialWithTimeout(fmt.Sprintf("mongodb://%s:%s/%s", config.Database.DSN, config.Database.Port, config.Database.Name), 5)

	// Exit with an error
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	fmt.Print("Connected")
	defer session.Close()

	session.DB(config.Database.Name).C("quotes").Insert(quote)

	return quote
}

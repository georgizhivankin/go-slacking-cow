package repositories

import (
	"fmt"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/georgizhivankin/go-slacking-cow/helpers"
	"github.com/georgizhivankin/go-slacking-cow/models"
	"gopkg.in/mgo.v2"
)

func SaveQuote(quote models.Quote) bool {
	// Connect to Mongo
	config, err := config.LoadEnv()
	helpers.CheckAndLogError(err, "panic")
	session, err := mgo.DialWithTimeout(fmt.Sprintf("%s:%s", config.Database.DSN, config.Database.Port), 3)

	// Exit with an error
	helpers.CheckAndLogError(err, "panic")
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	helpers.CheckAndLogError(err, "panic")

	result := models.Quote{}

	session.DB(config.Database.Name).C("quotes").Insert(quote)
	err = session.DB(config.Database.Name).C("quotes").FindId(quote.Id).One(&result)

	if err != nil {
		return false
	}

	return true
}

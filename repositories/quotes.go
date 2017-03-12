package repositories

import (
	"fmt"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/georgizhivankin/go-slacking-cow/helpers"
	"github.com/georgizhivankin/go-slacking-cow/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collectionName string = "quotes"

func GetSession() *mgo.Session {
	// Connect to Mongo
	config, err := config.LoadEnv()
	helpers.CheckAndLogError(err, "panic")
	session, err := mgo.Dial(fmt.Sprintf("%s:%s/%s", config.Database.DSN, config.Database.Port, config.Database.Name))
	helpers.CheckAndLogError(err, "panic")
	session.SetMode(mgo.Monotonic, true)
	return session
}

func FindById(quoteId bson.ObjectId, collectionName string) models.Quote {
	config, err := config.LoadEnv()
	helpers.CheckAndLogError(err, "panic")
	result := models.Quote{}
	session := GetSession()
	session.DB(config.Database.Name).C(collectionName).FindId(quoteId).One(&result)

	return result
}

func SaveQuote(quote *models.Quote) bool {
	config, err := config.LoadEnv()
	helpers.CheckAndLogError(err, "panic")
	session := GetSession()

	err = session.DB(config.Database.Name).C(collectionName).Insert(quote)
	helpers.CheckAndLogError(err, "panic")
	result := FindById(quote.Id, collectionName)

	if result.Quote == "" {
		return false
	}

	return true
}

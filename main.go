package main

import (
	"log"
	//"github.com/patjackson52/ticketmaster-discovery-go/disco"
	"./disco"
)

func main() {
	discoGateway := disco.NewBuilder().
		ApiKey("0yEuZlGBOMb2AAYs4XeaRDpmIaaDqLWB").
		BaseUrl("http://app.ticketmaster.com").
		Logging(true).
		Build()
// We can pass the parameters available in eventSearchParams to customise our search. Here I have used a few
// to keep the search narrowed to a given location
	params := map[string]string{disco.GEOPOINT: "9q8yyh8",disco.SIZE:"1",disco.RADIUS:"1",disco.UNIT:"miles"}
	eventResults, err := discoGateway.SearchEvents(params)

	logErr(err)

	log.Println(eventResults)
	log.Println(len(eventResults.EmbeddedEvents.Events))

	attrResults, err2 := discoGateway.SearchAttractions(params)

	logErr(err2)

	log.Println(attrResults)

	eventId := eventResults.EmbeddedEvents.Events[0].Id
	tapEventId := eventResults.EmbeddedEvents.Events[0].Source.Id

	eventDetails, err3 := discoGateway.GetEventDetails(eventId)
	logErr(err3)
	log.Println(eventDetails)

	invResults, err4 := discoGateway.GetInventoryStatusDetails([]string{eventId})
	logErr(err4)
	log.Println(invResults)

	topPicksResults, err5 := discoGateway.GetTopPicks(tapEventId, nil)

	logErr(err5)
	log.Println(topPicksResults)

}

func logErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

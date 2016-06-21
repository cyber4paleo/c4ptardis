package main

import (
	"bytes"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	// "encoding/json"
	"fmt"
	"os"
	// "gopkg.in/mgo.v2/bson"
	sparql "github.com/knakk/sparql"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: maxage
prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> 
prefix iodp: <http://data.oceandrilling.org/core/1/> 
prefix foaf: <http://xmlns.com/foaf/0.1/> 
prefix owl: <http://www.w3.org/2002/07/owl#> 
prefix xsd: <http://www.w3.org/2001/XMLSchema#> 
prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> 
SELECT (MAX(?age) AS ?maxage) (MIN(?age) AS ?minage)
WHERE {
   ?s iodp:leg "192" .
   ?s iodp:ma ?age   .
}


# tag: query2
PREFIX chronos: <http://www.chronos.org/loc-schema#>
PREFIX geo: <http://www.w3.org/2003/01/geo/wgs84_pos#>
SELECT  ?dataset ?ob  ?age ?depth ?long ?lat
FROM <http://chronos.org/janusamp#>
WHERE {
  ?ob chronos:age ?age . 
  ?ob chronos:depth ?depth . 
  ?ob <http://purl.org/linked-data/cube#dataSet>  ?dataset .
  ?dataset geo:long ?long .
  ?dataset geo:lat ?lat .
   ?dataset chronos:leg "{{.Leg}}" .  
   ?dataset chronos:site "{{.Site}}" .
   ?dataset chronos:hole "{{.Hole}}" . 
}
ORDER BY ?dataset DESC(?age)
LIMIT 1

`

type Expedition struct {
	Expedition    string `json:"expedition"`
	Site    string `json:"site"`
	Hole    string `json:"hole"`

}

func main() {
	// call mongo and lookup the redirection to use...
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	IndexCSVW(session)
	// IndexSchema(session)

}

func IndexCSVW(session *mgo.Session) {
	// Optional. Switch the session to a monotonic behavior.
	csvw := session.DB("expedire").C("features")

	// Find all Documents CSVW
	var csvwDocs []Expedition
	err := csvw.Find(nil).All(&csvwDocs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// index some data
	for _, item := range csvwDocs {
		// fmt.Printf("Indexed item %d with Leg: %s\n", index, item.Expedition)

		res := SPARQLCall(item.Expedition, item.Site, item.Hole)

		solutionsTest := res.Solutions() // map[string][]rdf.Term
		// fmt.Println("res.Solutions():")
		for _, i := range solutionsTest {
			fmt.Printf("Leg Site Hole Age: %s %s %s %s \n", item.Expedition, item.Site, item.Hole,  i["age"])
		}

	}

	// // Update Example for Mongo
	// 	colQuerier := bson.M{"name": "Ale"}
	// 	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	// 	err = c.Update(colQuerier, change)
	// 	if err != nil {
	// 		panic(err)
	// 	}

}

func SPARQLCall(leg string, site string, hole string) *sparql.Results {
	repo, err := sparql.NewRepo("http://opencoredata.org/sparql",
		sparql.Timeout(time.Millisecond*15000),
	)
	if err != nil {
		log.Printf("%s\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("query2", struct{ Leg, Site, Hole string }{leg, site, hole})
	if err != nil {
		log.Printf("%s\n", err)
	}

	res, err := repo.Query(q)
	if err != nil {
		log.Printf("%s\n", err)
	}

	return res
}

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}

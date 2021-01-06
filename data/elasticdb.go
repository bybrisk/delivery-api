package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	elastic "github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error) {
	client, err :=  elastic.NewClient(elastic.SetURL("https://elastic:w9XrZDRi0JZmxFV5vwk6tVCq@390142e4769147acb17debc402b8474b.ap-south-1.aws.elastic-cloud.com:9243"),elastic.SetSniff(false),elastic.SetHealthcheck(false))
	return client, err

}

func InsertDeilveryWithGeoCode(d *AddDeliveryRequestWithGeoCode) string {
	var res string
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		log.Error("ElasticSearch initialization ERROR : ")
		log.Error(err)
	}

	//get current date
	currentTime := time.Now()
	date:=currentTime.Format("01-02-2006")

	dataJSON, err := json.Marshal(&d)
	js := string(dataJSON)
	ind, err := esclient.Index().Index(date).BodyJson(js).Do(ctx)

	if err != nil {
		log.Error("insertDataToElastic inserting ERROR : ")
		log.Error(err)
	}else{
		res=ind.Id
		fmt.Println("Delivery added to ElasticSearch with ID : ",ind.Id)
	}
	return res
}

func InsertDeilveryWithoutGeoCode(d *AddDeliveryRequestWithoutGeoCode) string {
	var res string
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		log.Error("ElasticSearch initialization ERROR : ")
		log.Error(err)
	}

	//get current date
	currentTime := time.Now()
	date:=currentTime.Format("01-02-2006")

	dataJSON, err := json.Marshal(&d)
	js := string(dataJSON)
	ind, err := esclient.Index().Index(date).BodyJson(js).Do(ctx)

	if err != nil {
		log.Error("insertDataToElastic inserting ERROR : ")
		log.Error(err)
	}else{
		res=ind.Id
		fmt.Println("Delivery added to ElasticSearch with ID : ",ind.Id)
	}
	return res
}
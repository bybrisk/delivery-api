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

	client, err :=  elastic.NewClient(elastic.SetURL("https://elastic:YqZmasvS43ne95XNNC9ZzkWs@83d2cc8c5e8c4cfc95f00e145d0f9be3.ap-south-1.aws.elastic-cloud.com:9243"),elastic.SetSniff(false),elastic.SetHealthcheck(false))
	fmt.Println("ElasticSearch initialized...")
	return client, err

}

func insertDataToElastic(d *AddDeliveryRequest) {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		log.Error("ElasticSearch initialization ERROR : ")
		log.Error(err)
	}

	//get current date
	currentTime := time.Now()
	date:=currentTime.Format("01-02-2006")
	fmt.Println(date)

	dataJSON, err := json.Marshal(&d)
	js := string(dataJSON)
	ind, err := esclient.Index().Index(date).BodyJson(js).Do(ctx)

	if err != nil {
		log.Error("insertDataToElastic inserting ERROR : ")
		log.Error(err)
	}else{
		fmt.Println("[Elastic][InsertProduct]Insertion Successful")
		fmt.Println(ind)
	}
}
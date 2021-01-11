package data

import (
	"encoding/json"
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	"bytes"
	"context"
	"net/http"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/mitchellh/mapstructure"
	
)

var (
	clusterURLs = []string{"https://390142e4769147acb17debc402b8474b.ap-south-1.aws.elastic-cloud.com:9243"}
	username    = "elastic"
	password    = "w9XrZDRi0JZmxFV5vwk6tVCq"
  )

func InsertDeilveryWithGeoCode(d *AddDeliveryRequestWithGeoCode) string {
	
	var id string

	cfg := elasticsearch.Config{
		Addresses: clusterURLs,
		Username:  username,
		Password:  password,		
	  }
	  es, err := elasticsearch.NewClient(cfg)
	  if err != nil {
		log.Error("ElasticSearch initialization ERROR : ")
		log.Error(err)
	}

	currentTime := time.Now()
	date:=currentTime.Format("01-02-2006")

	res, _ := es.Index(date, esutil.NewJSONReader(&d))
	var r map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    	log.Printf("Error parsing the response body: %s", err)
    } else {
    	// Print the response status and indexed document version.
		id=fmt.Sprintf("%v", r["_id"])
    }	
	
	log.Print(id)
	return id
}

func InsertDeilveryWithoutGeoCode(d *AddDeliveryRequestWithoutGeoCode) string {
	
	var id string

	cfg := elasticsearch.Config{
		Addresses: clusterURLs,
		Username:  username,
		Password:  password,		
	  }
	  es, err := elasticsearch.NewClient(cfg)
	  if err != nil {
		log.Error("ElasticSearch initialization ERROR : ")
		log.Error(err)
	}

	currentTime := time.Now()
	date:=currentTime.Format("01-02-2006")

	res, _ := es.Index(date, esutil.NewJSONReader(&d))
	var r map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    	log.Printf("Error parsing the response body: %s", err)
    } else {
    	// Print the response status and indexed document version.
		id=fmt.Sprintf("%v", r["_id"])
    }	
	
	log.Print(id)
	return id
}

func FetchDeliveryByID(docID string)  SingleDeliveryDetail{
	var delivery SingleDeliveryDetail

	var r  map[string]interface{}
	cfg := elasticsearch.Config{
		Addresses: clusterURLs,
		Username:  username,
		Password:  password,		
	  }
	  es, err := elasticsearch.NewClient(cfg)
	  if err != nil {
		log.Error("ElasticSearch initialization ERROR : ")
		log.Error(err)
	}

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"ids": map[string]interface{}{
				"values": docID,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Error("Error encoding query : ")
		log.Error(err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Error("Error getting response : ")
		log.Error(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Error("Error parsing the response body : ")
			log.Error(err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Error("Error parsing the response body : ")
		log.Error(err)
	}

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		sourceMap := hit.(map[string]interface{})["_source"]
		//log.Printf("%s",sourceMap)
		mapstructure.Decode(sourceMap, &delivery)
	}

	return delivery
}

func UpdateDeilveryStatusES(d *UpdateDeliveryStatus) string {
	var id string
	//Encode the data
	postBody, _ := json.Marshal(map[string]map[string]string{
		"doc": map[string]string{
			"deliveryStatus": d.DeliveryStatus,
		 },
	 })
	 responseBody := bytes.NewBuffer(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post("https://elastic:w9XrZDRi0JZmxFV5vwk6tVCq@390142e4769147acb17debc402b8474b.ap-south-1.aws.elastic-cloud.com:9243/01-11-2021/_update/"+d.DeliveryID, "application/json", responseBody)
  
	 //Handle Error
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	var r map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
    	log.Printf("Error parsing the response body: %s", err)
    } else {
    	// Print the response status and indexed document version.
		id=fmt.Sprintf("%v", r["_id"])
    }

	fmt.Println(id)
	return id
}
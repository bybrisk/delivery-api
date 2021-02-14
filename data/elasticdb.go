package data

import (
	"encoding/json"
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	"bytes"
	"context"
	"net/http"
	"io/ioutil"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/mitchellh/mapstructure"
	
)

var Elasticurl string = "https://67b69d7c039a4780b5aa5abbecb09c1a.ap-south-1.aws.elastic-cloud.com:9243"
var UsernameElastic string = "elastic"
var Elasticpassword string = "OmantonZe8l1MaugY6ypelSE"
var urlAuthenticate string = "https://elastic:OmantonZe8l1MaugY6ypelSE@67b69d7c039a4780b5aa5abbecb09c1a.ap-south-1.aws.elastic-cloud.com:9243"

var (
	clusterURLs = []string{Elasticurl}
	username    = UsernameElastic
	password    = Elasticpassword
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
	if (d.DeliveryStatus=="Pending-Cancelled"){
		d.DeliveryStatus="Cancelled"
	}
	var id string
	//Encode the data
	postBody:=`{
		"script" : {
			"source": "ctx._source.deliveryStatus='`+d.DeliveryStatus+`';",
			"lang": "painless"  
		  },
		  "query": {
			  "ids" : {
			"values" : "`+d.DeliveryID+`"
			}
		  }
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_update_by_query?conflicts=proceed", "application/json", responseBody)
  
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
		id=fmt.Sprintf("%v", r["updated"])
    }

	fmt.Println(id)
	return id
}

func UpdateDeilveryAgentES(d *UpdateDeliveryAgent) string {
	var id string
	//Encode the data
	postBody:=`{
		"script" : {
			"source": "ctx._source.deliveryAgentID='`+d.DeliveryAgentID+`';",
			"lang": "painless"  
		  },
		  "query": {
			  "ids" : {
			"values" : "`+d.DeliveryID+`"
			}
		  }
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_update_by_query?conflicts=proceed", "application/json", responseBody)
  
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
		id=fmt.Sprintf("%v", r["updated"])
    }

	fmt.Println(id)
	return id
}

func FetchAllDeliveryES(key string,docID string) DeliveryResponseBulk {
	var deliveries DeliveryResponseBulk

	postBody:=`{
		"query": {
		  "bool": {
			"filter": [
			  {"term": {
				"`+key+`": "`+docID+`"
			  }}
			]
		  }
		},
		"sort" : [
    { "rankingTime" : "desc" }
  ]
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_search?size=1000", "application/json", responseBody)
  
	 //Handle Error
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	 body, err := ioutil.ReadAll(resp.Body)
	 if err != nil {
		log.Error("ReadAll ERROR : ")
		log.Error(err)
	 }
	 
	 err = json.Unmarshal(body, &deliveries)
	 if err != nil {
		log.Error("json.Unmarshal ERROR : ")
		log.Error(err)
    	} 
	return deliveries
}

func FetchPendingDeliveryByAgentIdES(key string,docID string) *DeliveryResponseBulk {
	var deliveries DeliveryResponseBulk

	postBody:=`{
		"query": {
		  "bool": {
			"filter": [
			  {"term": {"`+key+`": "`+docID+`"}},
			  {"term" : {"deliveryStatus.keyword" : "Pending" }}
			]
		  }
		}
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_search?size=1000", "application/json", responseBody)
  
	 //Handle Error
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	 body, err := ioutil.ReadAll(resp.Body)
	 if err != nil {
		log.Error("ReadAll ERROR : ")
		log.Error(err)
	 }
	 
	 err = json.Unmarshal(body, &deliveries)
	 if err != nil {
		log.Error("json.Unmarshal ERROR : ")
		log.Error(err)
    	}	
	return &deliveries
}

func FetchDeliveryHistoryByAgentIdES(key string,docID string) DeliveryResponseBulk {
	var deliveries DeliveryResponseBulk

	postBody:=`{
		"query": {
		  "bool": {
			"filter": [
			  {"term": {"`+key+`": "`+docID+`"}}
			],
			"must_not": [
			  {"term":{"deliveryStatus.keyword": "Pending"}}
			  ]
		  }
		}
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_search?size=1000", "application/json", responseBody)
  
	 //Handle Error
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	 body, err := ioutil.ReadAll(resp.Body)
	 if err != nil {
		log.Error("ReadAll ERROR : ")
		log.Error(err)
	 }
	 
	 err = json.Unmarshal(body, &deliveries)
	 if err != nil {
		log.Error("json.Unmarshal ERROR : ")
		log.Error(err)
    	} 
	return deliveries
}

func UpdateDeilveryDistanceES(d *UpdateDeliveryDistance) string{
	var id string
	//Encode the data
	postBody:=`{
		"script" : {
			"source": "ctx._source.distanceObserved=`+fmt.Sprintf("%f", d.Distance)+`;",
			"lang": "painless"  
		  },
		  "query": {
			  "ids" : {
			"values" : "`+d.DeliveryID+`"
			}
		  }
	  }`

	 responseBody := bytes.NewBufferString(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(urlAuthenticate+"/_all/_update_by_query?conflicts=proceed", "application/json", responseBody)
  
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
		id=fmt.Sprintf("%v", r["updated"])
    }

	fmt.Println(id)
	return id	
}
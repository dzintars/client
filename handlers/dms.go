package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/oswee/client/models"
)

func CreateDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	ref := r.FormValue("reference")
	da := r.FormValue("address")
	dz := r.FormValue("zip")
	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
	lng, _ := strconv.ParseFloat(r.FormValue("lng"), 64)
	tw, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	rs, _ := strconv.ParseInt(r.FormValue("sequence"), 10, 64)

	_, err = models.CreateDeliveryOrder(ref, da, dz, lat, lng, tw, rs)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

func UpdateDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	id, _ := strconv.ParseInt(r.FormValue("order_id"), 10, 64)
	ref := r.FormValue("reference")
	da := r.FormValue("address")
	dz := r.FormValue("zip")
	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
	lng, _ := strconv.ParseFloat(r.FormValue("lng"), 64)
	tw, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	rs, _ := strconv.ParseInt(r.FormValue("sequence"), 10, 64)

	_, err = models.UpdateDeliveryOrder(ref, da, dz, lat, lng, tw, rs, id)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

// DeleteDeliveryOrderPostHandler ...
func DeleteDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)

	_, err = models.DeleteDeliveryOrder(id)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

//models.GeoCode("Brīvības iela 55, Rīga")

func GoogleGeoCode(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	id, _ := strconv.ParseInt(r.FormValue("order_id"), 10, 64)
	// ToDO: Change hardcoded stakeholder ID to form value
	stakeholder, _ := strconv.ParseInt(r.FormValue("1"), 10, 64)
	da := r.FormValue("address")

	_, err = models.GeoCodeDeliveryOrder(id, stakeholder, da)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

func BatchGeoCode(w http.ResponseWriter, r *http.Request) {
	models.CreatePageView(r)
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	sID, _ := strconv.ParseInt(r.FormValue("stakeholder_id"), 10, 64)
	models.BatchGeocode(sID)
	http.Redirect(w, r, "/", 302)
}

type Route struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
}

// CreateRoutePost ...
func CreateRoutePost(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	// Retrieve body from the http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Body read: ", b)

	// Save route in Route struct
	var _route Route
	err = json.Unmarshal(b, &_route)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("Body to struct: ", _route)

	saveRouteToKafka(_route)
	fmt.Println("Saving to kafka...")
	// Convert job struct into json
	// jsonString, err := json.Marshal(_route)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	//fmt.Println("_route to JSON: ", jsonString)

	// Set content-type http header
	//w.Header().Set("content-type", "application/json")

	// Send back data as response
	//w.Write(jsonString)

	//http.Redirect(w, r, "/", 302)
}

func saveRouteToKafka(rt Route) {
	//fmt.Println("Save to Kafka")

	jsonString, err := json.Marshal(rt)

	routeString := string(jsonString)
	fmt.Println("routeString: ", routeString)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka1:9092,kafka2:9092,kafka3:9092"})
	if err != nil {
		panic(err)
	}

	// Produce messages to topic (asynchronously)
	topic := "go-test"
	for _, word := range []string{string(routeString)} {
		fmt.Println("Ranging: ", word)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(word),
		}, nil)
	}
	// Wait for message deliveries
	p.Flush(15 * 1000)
}

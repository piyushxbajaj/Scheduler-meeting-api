package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type meeting struct {
	ID                string    `json:"ID"`
	Title             string    `json:"Title"`
	Participants      int       `json:"Partcipants"`
	StartTime         time.Time `json:"StartTime"`
	EndTime           time.Time `json:"EndTime"`
	CreationTimestamp time.Time `json:"CreationTime"`
}

type part struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

type allMeet []meeting
type allPart []part

var meets = allMeet{
	{ID: "1", Title: "Learning GO", Participants: 10, StartTime: time.Date(2020, 11, 14, 10, 45, 16, 0, time.UTC), EndTime: time.Date(2020, 11, 15, 10, 45, 16, 0, time.UTC), CreationTimestamp: time.Date(2020, 8, 15, 10, 45, 16, 0, time.UTC)},
}

var parts = allPart{
	{
		Name:  "Piyush",
		Email: "piyushpiyush41@gmail.com",
		RSVP:  "Yes",
	},
}

func createSchedule(w http.ResponseWriter, r *http.Request) {
	var newMeet meeting
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter meeting title, participants, start_time and end_time to update")
	}

	json.Unmarshal(reqBody, &newMeet)
	meets = append(meets, newMeet)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newMeet)
}

func getOneMeet(w http.ResponseWriter, r *http.Request) {
	meetID := mux.Vars(r)["id"]

	for _, singleMeet := range meets {
		if singleMeet.ID == meetID {
			json.NewEncoder(w).Encode(singleMeet)
		}
	}
}

func getAllMeets(w http.ResponseWriter, r *http.Request) {
	meetStart := mux.Vars(r)["startTime"]
	meetEnd := mux.Vars(r)["endTime"]
	layout := "0001-01-01T00:00:00Z"
	tstart, err := time.Parse(layout, meetStart)
	tend, err := time.Parse(layout, meetEnd)
	if err != nil {
		fmt.Println(err)
	}
	for _, singleMeet := range meets {
		if singleMeet.StartTime == tstart && singleMeet.EndTime == tend {
			json.NewEncoder(w).Encode(singleMeet)
		}
	}
}

// func (ph *urlhandler) getlistofmeetings(response http.ResponseWriter, request *http.Request) { // A function to return all meeting which took place within the given time
// 	response.Header().Set("content-type", "application/json")
// 	var meetingsdb []Meeting
// 	collection := ph.db.Database("appointydb").Collection("meetingsdb")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	start1, ok := request.URL.Query()["start"]
// 	if !ok || len(start1) == 0 {
// 		// handle missing starttime
// 		return
// 	}
// 	startTimeStamp, err := time.Parse(time.RFC3339, start1[0])
// 	if err != nil {
// 		// invalid timestamp
// 		return
// 	}
// 	end1, ok := request.URL.Query()["end"]
// 	if !ok || len(end1) == 0 {
// 		// handle missing endtime
// 		return
// 	}
// 	endTimeStamp, err := time.Parse(time.RFC3339, end1[0])
// 	if err != nil {
// 		// invalid timestamp
// 		return
// 	}
// 	filter := bson.M{
// 		"starttime": bson.M{"$gte": startTimeStamp},
// 		"endtime":   bson.M{"$lte": endTimeStamp},
// 	}
// 	cursor, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var person Meeting
// 		cursor.Decode(&person)
// 		meetingsdb = append(meetingsdb, person)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(meetingsdb)
// }

func allMeetPart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllMeetings")
	json.NewEncoder(w).Encode(meets)

}

func returnAllParticipants(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllParticipants")
	json.NewEncoder(w).Encode(parts)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to meeting scheduler!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/meetings", createSchedule).Methods("POST")
	router.HandleFunc("/meeting/{id}", getOneMeet).Methods("GET")
	router.HandleFunc("/meetings?start={startTime}&end={endTime}", getAllMeets).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

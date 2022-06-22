package main

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Holds the data to display for each section
//This is udpated on event and routine callbacks
var sectionStatus []string

type callbackFunc func(*int, *int) (callbackFunc, error)

//Main configuration struct
var sectionCallbacks = []struct {
	eventCallback callbackFunc
	callback      func(int) error
	group         int
	interval      time.Duration
}{
	{dateEventCallback, dateCallback, 1, time.Second},
	{audioDownEventCallback, audioDownCallback, 2, time.Duration(0)},
	{audioEventCallback, audioCallback, 2, time.Second * 10},
	{audioUpEventCallback, audioUpCallback, 2, time.Duration(0)},
	{updateEventCallback, updateCallback, 3, time.Hour},
	{musicDownEventCallback, musicDownCallback, 4, time.Duration(0)},
	{musicEventCallback, musicCallback, 4, time.Second * 10},
	{musicUpEventCallback, musicUpCallback, 4, time.Duration(0)},
	{bitcoinEventCallback, bitcoinCallback, 5, time.Minute},
}

//Initialize each section with an empty
//string toprevent invalid memory access
func init() {
	for i := 0; i < len(sectionCallbacks); i++ {
		sectionStatus = append(sectionStatus, "")
	}
}

func main() {
	callbackInit()

	mux := mux.NewRouter()

	mux.HandleFunc("/eventSection/{section}/eventButton/{button}", eventHandler)

	err := http.ListenAndServe(":1058", mux)
	if err != nil {
		panic(err)
	}
}

//eventHandler validates requests sent to our
//http endpoint, and calls eventCaller after
func eventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	section, err := strconv.Atoi(vars["section"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	button, err := strconv.Atoi(vars["button"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = eventCaller(sectionCallbacks[section].eventCallback, section, button)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

//eventCaller calls a given callback function with the data
//from the eventHandler. Callback functions can return another
//callback function to update a seperate clickable section
func eventCaller(cbf callbackFunc, section int, button int) (err error) {
	cbf, err = cbf(&section, &button)
	if err != nil {
		return err
	}

	if cbf != nil {
		eventCaller(cbf, section, button)
	}

	//I'm leaving this to (potentially) run many times
	//in case functions called later on return an error
	err = printStatus()
	if err != nil {
		return err
	}

	return nil
}

//callbackInit starts the main update loops for each
//callback function. This should return almost instantly
func callbackInit() {
	for i := 0; i < len(sectionCallbacks); i++ {
		go func(i int) {
			for {
				sectionCallbacks[i].callback(i)

				err := printStatus()
				if err != nil {
					log.Println("Encountered Error: ", err)
				}

				if sectionCallbacks[i].interval == time.Duration(0) {
					return
				}

				time.Sleep(sectionCallbacks[i].interval)
			}
		}(i)
	}
}

//printStatus constructs and sets a string to xsetroot.
//It will error out if there are 0 values in sectionStatus
func printStatus() error {
	var status string = "\x01" + sectionStatus[0] + " -"

	for i := 1; i < len(sectionCallbacks); i++ {
		status += string(byte(i + 1))

		if sectionCallbacks[i-1].group != sectionCallbacks[i].group {
			status += "- "
		}

		status += sectionStatus[i]

		//Check if the current status section is the last one before checking the
		//group to avoid a segmentation fault when accessing the next section's group
		if i < len(sectionCallbacks)-1 && sectionCallbacks[i].group != sectionCallbacks[i+1].group {
			status += " -"
		}
	}

	return exec.Command("xsetroot", "-name", status).Run()
}

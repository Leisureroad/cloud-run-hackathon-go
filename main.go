package main

import (
	"encoding/json"
	"fmt"
	"log"
	// rand2 "math/rand"
	"net/http"
	"os"
	// "reflect"
)

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprint(w, "Let the battle begin!")
		return
	}

	var v ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := play(v)
	fmt.Fprint(w, resp)
}

func play(input ArenaUpdate) (response string) {
	// log.Printf("IN: %#v", input)

	// r := reflect.ValueOf(input)
	// f := reflect.Indirect(r).FieldByName("Arena")
	// log.Printf("value: ", input.Arena.State["PlayerState"].X)
	x :=input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].X
	y :=input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].Y
	d :=input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].Direction


	// direction :=input.Arena.State["PlayerState"].Direction
	log.Printf("%i", x)
	log.Printf("%i", y)
	log.Printf("%v", d)
  var xm int
  var ym int
	s :=input.Arena.State
	for playerUrl, playerState := range s {
	        fmt.Println("playerUrl:", playerUrl, "=>", "playerState:", playerState)
					px :=playerState.X
					py :=playerState.Y
					if (x-px == 1) {
					  xm = 1
						log.Printf("%v", xm)
				  }
					if (y-py == 1) {
						ym = 1
						log.Printf("%v", ym)
					}
	    }
	if (x != 0) {
		if (d == "S" || d == "E") {
	    return "R"
		} else {
			if (xm == 1) {
				return "R"
			}
			return "F"
		}
	}

	if (y != 0) {
		if (d == "N") {
	    return "F"
		} else {
			return "R"
		}
	}

	if (d == "E" ) {
	  return "T"
	} else {
		return "R"
	}

	// log.Printf("%v", f)

	// log.Printf("X: %#v", input.Arena.State.PlayerState["X"])

	// commands := []string{"T", "F"}
	// rand := rand2.Intn(2)


	// TODO add your implementation here to replace the random response
	// return commands[rand]
	// return "T"
}

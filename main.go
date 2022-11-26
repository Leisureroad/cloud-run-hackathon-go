package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
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
	log.Printf("IN: %#v", input)

	// r := reflect.ValueOf(input)
	// f := reflect.Indirect(r).FieldByName("Arena")
	// log.Printf("value: ", input.Arena.State["PlayerState"].X)
	x :=input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].X
	y :=input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].Y
	d :=input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].Direction
	wasHit := input.Arena.State["https://cloud-run-hackathon-go-pkqx6rxn2q-uc.a.run.app"].WasHit
	// direction :=input.Arena.State["PlayerState"].Direction
	// log.Printf("%i", x)
	// log.Printf("%i", y)
	// log.Printf("%v", d)
  var west int
	var east int
	var south int
	var north int

	s :=input.Arena.State
	x_max :=input.Arena.Dimensions[0]
	y_max :=input.Arena.Dimensions[1]
	for _, playerState := range s {
      // fmt.Println("playerUrl:", playerUrl, "=>", "playerState:", playerState)
			px :=playerState.X
			py :=playerState.Y
		  xw := x-px
			xe := px-x
			ys := y-py
			yn := py - y

			if (xw == 1 && ys == 0) {
			  west = 1
			}
			if (xe == 1 && ys == 0) {
			  east = 1
			}
			if (ys == 1 && xw == 0) {
			  south = 1
			}
			if (yn == 1 && xw == 0) {
			  north = 1
			}
	}

  if (wasHit == true) {
		commands := []string{"F", "R"}
		rand := rand2.Intn(2)
		log.Printf("wasHit: %v", commands[rand])
		return commands[rand]
	}
	
	if (y == y_max-1) {
		log.Printf("y_max: R")
		return "R"
	}
	if (x == x_max-1) {
		log.Printf("x_max: R")
		return "R"
	}

	if (west == 1 && d == "W") {
		log.Printf("west: T")
		return "T"
	} else if (east == 1 && d == "E") {
		log.Printf("east: T")
		return "T"
	} else if (south == 1 && d == "S") {
		log.Printf("south: T")
		return "T"
	} else if (north == 1 && d == "N") {
		log.Printf("north: T")
		return "T"
	} else {
		log.Printf("blank: F")
		return "F"
	}

	// if (x != 0) {
	// 	if (d == "S" || d == "E") {
	//     return "R"
	// 	} else if (d == "N" && y == 0) {
	// 		return "L"
	// 	} else if (xm == 1 && y!= 0) {
	// 		    return "R"
	// 	} else {
	// 		commands := []string{"T", "F"}
	// 		rand := rand2.Intn(2)
	// 		return commands[rand]
	// 	}
	// }
	//
	// if (y != 0) {
	// 	if (d == "N") {
	// 		commands := []string{"T", "F"}
	// 		rand := rand2.Intn(2)
	// 		return commands[rand]
	// 	} else {
	// 		return "R"
	// 	}
	// }
	//
	// if (d == "E" ) {
	//   return "T"
	// } else {
	// 	return "R"
	// }

	// log.Printf("%v", f)

	// log.Printf("X: %#v", input.Arena.State.PlayerState["X"])

	// commands := []string{"L", "R", "T", "F"}
	// rand := rand2.Intn(4)
	//
	//
	// // TODO add your implementation here to replace the random response
	// return commands[rand]
	// return "T"
}

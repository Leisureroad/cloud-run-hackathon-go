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

	// direction :=input.Arena.State["PlayerState"].Direction
	log.Printf("%i", x)
	log.Printf("%i", y)
	log.Printf("%v", d)
  var xw int
	var xe int
  var ys int
	var yn int

	s :=input.Arena.State
	x_max :=input.Arena.Dimensions[0]
	y_max :=input.Area.Dimensions[1]
	for playerUrl, playerState := range s {
      fmt.Println("playerUrl:", playerUrl, "=>", "playerState:", playerState)
			px :=playerState.X
			py :=playerState.Y
			if (x-px == 1) {
			  xw = 1
				log.Printf("%v", xw)
		  }
			if (px-x == 1) {
				xe = 1
				log.Printf("%v", xe)
			}
			if (y-py == 1) {
				ys = 1
				log.Printf("%v", ys)
			}
			if (py - y == 1) {
				yn = 1
				log.Printf("%v", yn)
			}
	}

	if (y == y_max) {
		return "R"
	}
	if (x == x_max) {
		return "R"
	}

	if (xw == 1 && d == "W") {
		return "T"
	} else if (xe == 1 && d == "E") {
		return "T"
	} else if (ys == 1 && d == "S") {
		return "T"
	} else if (yn == 1 && d == "N") {
		return "T"
	} else {
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

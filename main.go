package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func calculate(w http.ResponseWriter, r *http.Request) {

	data := [][2]string{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Fprintf(w, "invalid data %s", err)
		return
	}

	var starts, ends = make(map[string]struct{}), make(map[string]struct{})

	for _, d := range data {
		from := d[0]
		to := d[1]

		starts[from] = struct{}{}
		ends[to] = struct{}{}
	}
	for _, d := range data {
		from := d[0]
		to := d[1]

		delete(starts, to)
		delete(ends, from)
	}

	var start, end string
	if len(starts) != 1 {
		fmt.Fprintf(w, "no singular start")
		return
	} else {
		for k := range starts {
			start = k
		}
	}

	if len(ends) != 1 {
		fmt.Fprintf(w, "no singular end")
		return
	} else {
		for k := range ends {
			end = k
		}
	}

	json.NewEncoder(w).Encode([2]string{start, end})

}

func main() {
	http.HandleFunc("/calculate", calculate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

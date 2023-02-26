// Written by @agzg

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

const link = "https://projecteuler.net/project/resources/p022_names.txt"

func main() {
	r, err := http.Get(link)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	s := strings.ReplaceAll(string(b), `"`, "")

	names := strings.Split(s, ",")
	sort.Strings(names)

	var answer int

	for i, name := range names {
		var accum int
		for _, c := range name {
			accum += (int(c)-64)
		}
		answer += accum * (i+1)
	}
	fmt.Println(answer)
}

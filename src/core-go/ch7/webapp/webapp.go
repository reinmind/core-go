package webapp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type price float32
type database map[string]price

//Point a Point have x axis and y axis
type Point struct {
	X uint32 `json:"x"`
	Y uint32 `json:"y"`
}

func (p *Point) move() {
	p.X++
	p.Y++
}

func (p price) String() string {
	return fmt.Sprintf("%.2f", p)
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}
func (p *Point) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("%v\n", p)
	p.move()
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", b)
}
func main() {
	pt := Point{0, 0}
	db := database{"shore": 33, "water": 22}
	sm := http.NewServeMux()

	sm.HandleFunc("/db", http.HandlerFunc(db.ServeHTTP))
	sm.HandleFunc("/pt", http.HandlerFunc(pt.ServeHTTP))

	// db := database{"shoes": 32, "socks": 50}
	// go log.Fatal(http.ListenAndServe("localhost:8082", db))
	log.Fatal(http.ListenAndServe("localhost:8083", sm))
}

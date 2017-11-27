package main

import (
    "fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
    "encoding/json"
    "database/sql"
    "gopkg.in/doug-martin/goqu.v4"
    _ "gopkg.in/doug-martin/goqu.v4/adapters/postgres"
    _ "github.com/lib/pq"
)

var db *goqu.Database
func SetDB() {
    pgDb, err := sql.Open("postgres", "host=psql user=admin password=admin dbname=shakespeare sslmode=disable ")
    if err != nil {
        panic(err.Error())
    }
    db = goqu.New("postgres", pgDb)
}

type Work struct{
    Id int64 `db:"workid"`
    Title string `db:"title"`
    LongTitle string `db:"longtitle"`
    Year int `db:"year"`
    GenreType string `db:"genretype"`
    Notes string `db:"notes"`
    Source string `db:"source"`
    TotalWords int64 `db:"totalwords"`
    TotalParagraphs string `db:"totalparagraphs"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
    var works []Work
    if err := db.From("work").ScanStructs(&works); err != nil{
        fmt.Println(err.Error())
        return
    }
    works_json, err := json.Marshal(works)
    if err != nil{
       fmt.Println(err.Error())
       return
    }
    w.Write(works_json)
}

func main() {
    SetDB()
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", RootHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

package router

import (
	"core/config"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type HomeStruct struct {
	PageTitle string
}
type Ciudad struct {
	id     int    `json:"id"`
	nombre string `json:"nombre"`
}

func Home(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "dba:dba@tcp(127.0.0.1:3306)/accounts")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT id, nombre FROM ciudad")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var ciudad Ciudad
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ciudad.id, &ciudad.nombre)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(ciudad.nombre)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	data := HomeStruct{
		PageTitle: config.Get().AppTitle,
	}
	tmpl.Execute(w, data)
}

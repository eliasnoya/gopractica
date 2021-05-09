package router

import (
	"core/config"
	"core/model"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type HomeStruct struct {
	PageTitle string
}
type Ciudad struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

func Home(w http.ResponseWriter, r *http.Request) {

	results, err := model.GetConn().Query("SELECT id, nombre FROM ciudad")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var ciudad Ciudad
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ciudad.Id, &ciudad.Nombre)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(ciudad.Nombre)
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	data := HomeStruct{
		PageTitle: config.Get().AppTitle,
	}
	tmpl.Execute(w, data)
}

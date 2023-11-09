package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, nil)
}

func qrCode(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}

func main() {

	fmt.Println("Open App ...")
	http.HandleFunc("/", home)
	http.HandleFunc("/generator/", qrCode)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/mdp/qrterminal"
	"image/png"
	"net/http"
	"os"
)

func handlercurl(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	qrterminal.Generate(url, qrterminal.L, w)
}

func handlerpng(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	png.Encode(w, qrCode)
}

func main() {
	http.HandleFunc("/c", handlercurl)
	http.HandleFunc("/p", handlerpng)
	http.ListenAndServe(":8080", nil)
}

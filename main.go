package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func traerNotas(http.ResponseWriter, *http.Request) {

}

func traerNota(http.ResponseWriter, *http.Request) {

}

func crearNota(http.ResponseWriter, *http.Request) {

}

func editarNota(http.ResponseWriter, *http.Request) {

}

func borrarNota(http.ResponseWriter, *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(false)

	//Rutas
	router.HandleFunc("/notas", traerNotas).Methods("GET")
	router.HandleFunc("/notas/{}", traerNota).Methods("GET")
	router.HandleFunc("/notas", crearNota).Methods("POST")
	router.HandleFunc("/notas/{}", editarNota).Methods("PUT")
	router.HandleFunc("/notas/{}", borrarNota).Methods("DELETE")

	//Crear el servidor
	server := &http.Server{
		Addr:              ":8080",          //Puerto a usar
		Handler:           router,           //Rutas mux
		ReadHeaderTimeout: 10 * time.Second, //Tiempo maximo para leeer
		WriteTimeout:      10 * time.Second, //Tiempo maximo para escribir
		MaxHeaderBytes:    1 << 20,          //Tamaño de la cabecera
	}

	log.Println("Listening...")
	log.Println(server.ListenAndServe())
}

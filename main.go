package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/erikyvanov/Notas-API-REST/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func traerNota(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nota db.Nota

	//Obtener ID
	var p = mux.Vars(r)

	//convertir el ID string a ID BSON
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(p["id"])

	//Conectar con la base de datos
	coleccion := db.ConnectToMongoDB()

	//Se usa el id como flitro
	filtro := bson.M{"_id": id}

	err := coleccion.FindOne(context.TODO(), filtro).Decode(&nota)
	if err != nil {
		panic(err)
	}

	fmt.Println("Se envio una nota,", nota.ID)
	json.NewEncoder(w).Encode(nota)
}

func traerNotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var notas []db.Nota

	coleccion := db.ConnectToMongoDB()

	//Traer todas las notas
	cur, err := coleccion.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.TODO())

	//Decodificar cada tarea
	for cur.Next(context.TODO()) {
		var nota db.Nota

		err := cur.Decode(&nota)
		if err != nil {
			log.Fatal(err)
		}

		notas = append(notas, nota)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Se enviaron todas las notas")
	json.NewEncoder(w).Encode(notas)
}

func crearNota(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevaNota db.Nota
	err := json.NewDecoder(r.Body).Decode(&nuevaNota)

	if err != nil {
		panic(err)
	}

	//conectar con la base de datos
	collection := db.ConnectToMongoDB()

	insertResult, err := collection.InsertOne(context.TODO(), nuevaNota)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(insertResult)
	fmt.Println("Nota insertada, ID: ", insertResult.InsertedID)
}

func editarNota(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p = mux.Vars(r)

	//Convertir ID string en ID BSON
	id, _ := primitive.ObjectIDFromHex(p["id"])

	var nota db.Nota

	//Conectar con la base de datos
	coleccion := db.ConnectToMongoDB()

	filtro := bson.M{"_id": id}

	//Decodificar el objeto enviado
	json.NewDecoder(r.Body).Decode(&nota)

	//Pasar objeto a BSON Document
	update := bson.D{
		{"$set", bson.D{
			{"titulo", nota.Titulo},
			{"descripcion", nota.Descripcion},
			{"autor", nota.Autor},
		}},
	}

	err := coleccion.FindOneAndUpdate(context.TODO(), filtro, update).Decode(&nota)

	if err != nil {
		panic(err)
	}

	nota.ID = id

	fmt.Println("Se modifico un elemento", nota.ID)
	json.NewEncoder(w).Encode(nota)
}

func borrarNota(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(false)

	//Rutas
	router.HandleFunc("/notas", traerNotas).Methods("GET")
	router.HandleFunc("/notas/{id}", traerNota).Methods("GET")
	router.HandleFunc("/notas", crearNota).Methods("POST")
	router.HandleFunc("/notas/{id}", editarNota).Methods("PUT")
	router.HandleFunc("/notas/{id}", borrarNota).Methods("DELETE")

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

package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Nota es la estructura para guardar los datos
type Nota struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Titulo      string             `bson:"titulo"`
	Descripcion string             `bson:"descripcion"`
	Autor       string             `bson:"autor"`
}

//Constantes para manejar la base de datos
const (
	DataBaseName   string = "serverNotas"
	CollectionName string = "notas"
)

//Port es el puerto donde esta la base de datos
const Port string = "27017"

//ConnectToMongoDB es para conectarse a mongo DB
func ConnectToMongoDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%s", Port))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(DataBaseName).Collection(CollectionName)

	return collection
}

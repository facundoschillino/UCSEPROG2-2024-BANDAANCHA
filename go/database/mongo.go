package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client //el campo de este struct es un puntero a la instancia de mongo Client, que seria lo que representa la conexion a la base de datos.
} // aca usamos un puntero ya que necesitamos crear una unica instancia que sera utilizada en todo el programa, de modo que no creamos copias innecesarias. IMPORTANTE: Me brinda consistencia, al trabajar sobre una unica instancia me aseguro de que cuando inicio/apago se hace en todos lados

func NewMongoDB() *MongoDB {
	instancia := &MongoDB{}
	instancia.Connect()
	return instancia
}

func (mongoDB *MongoDB) GetClient() *mongo.Client {
	return mongoDB.Client
}

func (mongoDB *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mongoDB.Client = client

	return nil
}

func (mongoDB *MongoDB) Disconnect() error {
	return mongoDB.Client.Disconnect(context.Background())
}

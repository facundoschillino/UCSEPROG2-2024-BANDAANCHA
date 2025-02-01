package repositories

import (
	"Github/UCSEPROG2-2024-BANDAANCHA/go/database"
	"Github/UCSEPROG2-2024-BANDAANCHA/go/models"
	"Github/UCSEPROG2-2024-BANDAANCHA/go/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositoryInterface interface {
	ObtenerProductos() ([]models.Producto, error)
	ObtenerProductoPorID(id string) (models.Producto, error)
	CrearProducto(producto models.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(id primitive.ObjectID) (*mongo.DeleteOneResult, error)
	ModificarProducto(producto models.Producto) (*mongo.UpdateOneResult, error)
}

type ProductoRepository struct {
	db database.DB
}

func NewProductoRepository(db database.DB) *ProductoRepository {
	return &ProductoRepository{
		db: db,
	}
}

func (repo ProductoRepository) ObtenerProductos() {
	collection := repo.db.GetClient().Database("DBP").Collection("productos")
	filtro := bson.M{} // por ahora lo hacemos sin los filtros, eso lo agregamos mas adelante

	cursor, err := collection.Find(context.TODO(), filtro) //en cursor voy a guardar la lista completa sin filtrar
	defer cursor.Close(context.Background())
	var productos []models.Producto
	for cursor.Next(context.Background()) {
		var producto models.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		productos = append(productos, producto)
	}
	return productos, err
}
func (repository ProductoRepository) ObtenerProductoPorID(id string) (models.Producto, error) {
	collection := repository.db.GetClient().Database("DBP").Collection("productos")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	// Itera a través de los resultados
	var producto models.Producto

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&producto)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return producto, err
}
func (repo ProductoRepository) CrearProducto(producto models.Producto) (*mongo.InsertOneResult, error) {
	collection := repo.db.GetClient().Database("DBP").Collection("productos")
	resultado, err := collection.InsertOne(context.TODO, producto)
	return resultado, err
}

func (repository ProductoRepository) ModificarProducto(producto models.Producto) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("productos")

	filtro := bson.M{"_id": producto.ID}
	entidad := bson.M{"$set": bson.M{"nombre": producto.Nombre}}

	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)

	return resultado, err
}

func (repo ProductoRepository) EliminarProducto(id primitive.ObjectID) (*mongo.DeleteOneResult, error) {
	collection := repo.db.GetClient().Database("DBP").Collection("productos")
	filtro := bson.M{"_id": id}
	resultado, err := collection.DeleteOne(context.TODO(), filtro)
	return resultado, err
}

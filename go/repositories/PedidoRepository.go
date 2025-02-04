package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/dto"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PedidoRepositoryInterface interface {
	ObtenerPedidos() ([]models.Pedido, error)
	InsertarPedido(pedido models.Pedido) (*mongo.InsertOneResult, error)
	ObtenerPedidosAprobados() ([]models.Pedido, error)
	AceptarPedido(pedido models.Pedido) (*mongo.UpdateResult, error)
	CancelarPedido(pedido models.Pedido) (*mongo.UpdateResult, error)
	ParaEnviarPedido(pedido models.Pedido) (*mongo.UpdateResult, error)
	EnviadoPedido(pedido models.Pedido) (*mongo.UpdateResult, error)
	ObtenerPedidoPorID(pedidoConId models.Pedido) (models.Pedido, error)
	ObtenerPesoPedido(pedido models.Pedido) (int, error)
	ActualizarPedido(pedido models.Pedido) (*mongo.UpdateResult, error)
	ObtenerPedidosFiltrados(filtro dto.FiltroPedido) ([]models.Pedido, error)
}
type PedidoRepository struct {
	db DB
}

func NewPedidoRepository(db DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (repo PedidoRepository) ObtenerPesoPedido(pedido models.Pedido) (int, error) {
	//Obtener el el pedido por id, luego tomar su lista de productos. A cada producto multiplicarle su peso por la cantidad. A eso sumarlo y retornar ese total
	pedidoObtenido, err := repo.ObtenerPedidoPorID(pedido)
	if err != nil {
		return 0, err
	}
	var pesoTotal int
	for _, producto := range pedidoObtenido.Productos {
		pesoTotal += producto.PesoUnitario * producto.Cantidad
	}
	return pesoTotal, nil

}
func (repo PedidoRepository) ActualizarPedido(pedido models.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	pedido.FechaModificacion = time.Now()
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": pedido}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) ObtenerPedidosAprobados() ([]models.Pedido, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"Estado": "Aceptado"}
	cursor, err := lista.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	var Pedidos []models.Pedido
	for cursor.Next(context.Background()) {
		var Pedido models.Pedido
		err := cursor.Decode(&Pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		Pedidos = append(Pedidos, Pedido)
	}
	return Pedidos, err
}
func (repository PedidoRepository) ObtenerPedidoPorID(pedidoABuscar models.Pedido) (models.Pedido, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedidoABuscar.ID}
	cursor, err := collection.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	// Itera a través de los resultados
	var pedido models.Pedido
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return pedido, err
}
func (repo PedidoRepository) ObtenerPedidos() ([]models.Pedido, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{}
	cursor, err := lista.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())

	var Pedidos []models.Pedido
	for cursor.Next(context.Background()) {
		var Pedido models.Pedido
		err := cursor.Decode(&Pedido)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		Pedidos = append(Pedidos, Pedido)
	}
	return Pedidos, err
}
func (repo PedidoRepository) InsertarPedido(pedido models.Pedido) (*mongo.InsertOneResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	resultado, err := lista.InsertOne(context.TODO(), pedido)
	return resultado, err
}
func (repo PedidoRepository) AceptarPedido(pedido models.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Aceptado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) CancelarPedido(pedido models.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"estado": "Cancelado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) ParaEnviarPedido(pedido models.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Para enviar"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repo PedidoRepository) EnviadoPedido(pedido models.Pedido) (*mongo.UpdateResult, error) {
	lista := repo.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	filtro := bson.M{"_id": pedido.ID}
	entity := bson.M{"$set": bson.M{"Estado": "Enviado"}}
	resultado, err := lista.UpdateOne(context.TODO(), filtro, entity)
	return resultado, err
}
func (repository *PedidoRepository) obtenerPedidos(filtro bson.M) ([]models.Pedido, error) {
	collection := repository.db.GetClient().Database("BandaAncha").Collection("Pedidos")
	cursor, err := collection.Find(context.Background(), filtro)
	if err != nil {
		return nil, err
	}
	var pedidos []models.Pedido
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var pedido models.Pedido
		err := cursor.Decode(pedido)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return pedidos, nil
}
func (repo PedidoRepository) ObtenerPedidosFiltrados(filtro dto.FiltroPedido) ([]models.Pedido, error) {
	filtroGenerado := bson.M{}
	if filtro.Estado != "" {
		filtroGenerado["Estado"] = filtro.Estado
	}
	if !filtro.FechaMenor.IsZero() || !filtro.FechaMayor.IsZero() {
		filtroFecha := bson.M{}
		if !filtro.FechaMenor.IsZero() {
			filtroFecha["$gte"] = filtro.FechaMenor
		}
		if !filtro.FechaMayor.IsZero() {
			filtroFecha["$lte"] = filtro.FechaMayor
		}
		filtroGenerado["fecha_creacion"] = filtroFecha
	}
	return repo.obtenerPedidos(filtroGenerado)
}

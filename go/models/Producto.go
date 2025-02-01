package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Tipo                     TipoProducto       `bson:"tipo"`
	Nombre                   string             `bson:"nombre"`
	PesoUnitario             int64              `bson:"peso_unitario"`
	PrecioUnitario           int64              `bson:"precio_unitario"`
	StockMinimo              int64              `bson:"stock_minimo"`
	StockDisponible          int64              `bson:"stock_disponible"`
	FechaCreacion            time.Time          `bson:"fecha_creacion"`
	FechaUltimaActualizacion time.Time          `bson:"fecha_actualizacion"`
	Estado                   string             `bson:"estado"`
}

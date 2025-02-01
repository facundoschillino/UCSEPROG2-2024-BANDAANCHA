package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TipoProducto struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Nombre                   string             `bson:"nombre"`
	Descripcion              string             `bson:"descripcion"`
	FechaCreacion            primitive.DateTime `bson:"fecha_creacion"`
	FechaUltimaActualizacion primitive.DateTime `bson:"fecha_creacion"`
	Estado                   string             `bson:"estado"`
}

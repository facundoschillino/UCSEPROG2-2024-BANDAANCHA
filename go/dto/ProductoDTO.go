package dto

import (
	"time"

	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/models"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/utils"
)

type Producto struct {
	ID                       string              `json:"_id,omitempty"`
	Tipo                     models.TipoProducto `json:"tipo"`
	Nombre                   string              `json:"nombre"`
	PesoUnitario             int64               `json:"peso_unitario"`
	PrecioUnitario           int64               `json:"precio_unitario"`
	StockMinimo              int64               `json:"stock_minimo"`
	StockDisponible          int64               `json:"stock_disponible"`
	FechaCreacion            time.Time           `json:"fecha_creacion"`
	FechaUltimaActualizacion time.Time           `json:"fecha_actualizacion"`
	Estado                   string              `json:"estado"`
}

// Con esto transformo el model en un DTO
func NewProducto(producto models.Producto) *Producto {
	return &Producto{
		ID:                       utils.GetStringIDFromObjectID(producto.ID),
		Tipo:                     producto.Tipo,
		Nombre:                   producto.Nombre,
		PesoUnitario:             producto.PesoUnitario,
		PrecioUnitario:           producto.PrecioUnitario,
		StockMinimo:              producto.StockMinimo,
		StockDisponible:          producto.StockDisponible,
		FechaCreacion:            producto.FechaCreacion,
		FechaUltimaActualizacion: producto.FechaUltimaActualizacion,
		Estado:                   producto.Estado,
	}
}

// Con esto transformo un DTO en model
func (producto Producto) GetModel() models.Producto {
	return models.Producto{
		ID:                       utils.GetObjectIDFromStringID(producto.ID),
		Tipo:                     producto.Tipo,
		Nombre:                   producto.Nombre,
		PesoUnitario:             producto.PesoUnitario,
		PrecioUnitario:           producto.PrecioUnitario,
		StockMinimo:              producto.StockMinimo,
		StockDisponible:          producto.StockDisponible,
		FechaCreacion:            producto.FechaCreacion,
		FechaUltimaActualizacion: producto.FechaUltimaActualizacion,
		Estado:                   producto.Estado,
	}
}

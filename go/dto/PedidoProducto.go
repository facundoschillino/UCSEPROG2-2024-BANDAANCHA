package dto

import (
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/models"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/utils"
)

type PedidoProducto struct {
	ID             string `json:"id,omitempty"`
	CodigoProducto string `json:"codigo_producto"`
	PrecioUnitario int    `json:"precio_unitario"`
	PesoUnitario   int    `json:"peso_unitario"`
	Cantidad       int    `json:"cantidad"`
}

func NewPedidoProducto(pedidoProducto *models.PedidoProducto) *PedidoProducto {
	return &PedidoProducto{
		ID:             utils.GetStringIDFromObjectID(pedidoProducto.ID),
		CodigoProducto: pedidoProducto.CodigoProducto,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}

func (pedidoProducto PedidoProducto) GetModel() models.PedidoProducto {
	return models.PedidoProducto{
		ID:             utils.GetObjectIDFromStringID(pedidoProducto.ID),
		CodigoProducto: pedidoProducto.CodigoProducto,
		PesoUnitario:   pedidoProducto.PesoUnitario,
		PrecioUnitario: pedidoProducto.PrecioUnitario,
		Cantidad:       pedidoProducto.Cantidad,
	}
}

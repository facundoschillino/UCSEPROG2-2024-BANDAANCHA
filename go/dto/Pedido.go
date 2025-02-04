package dto

import (
	"time"

	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/models"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/utils"
)

type Pedido struct {
	ID                string           `json:"id,omitempty"`
	Productos         []PedidoProducto `json:"productos"`
	Destino           string           `json:"destino"`
	Estado            string           `json:"estado"`
	FechaCreacion     time.Time        `json:"fecha_creacion"`
	FechaModificacion time.Time        `json:"fecha_modificacion"`
}

func NewPedido(pedido models.Pedido) *Pedido {
	return &Pedido{
		ID:                utils.GetStringIDFromObjectID(pedido.ID),
		Productos:         []PedidoProducto{},
		Destino:           pedido.Destino,
		Estado:            "Pendiente",
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
}
func (pedido Pedido) GetModel() models.Pedido {
	return models.Pedido{
		ID:                utils.GetObjectIDFromStringID(pedido.ID),
		Productos:         pedido.getProductosElegidos(),
		Destino:           pedido.Destino,
		Estado:            pedido.Estado,
		FechaCreacion:     pedido.FechaCreacion,
		FechaModificacion: pedido.FechaModificacion,
	}
}

func (pedido Pedido) getProductosElegidos() []models.PedidoProducto {
	var productosElegidos []models.PedidoProducto
	for _, producto := range pedido.Productos {
		productosElegidos = append(productosElegidos, producto.GetModel())
	}
	return productosElegidos
}
func NewProductosPedido(productosElegidos []models.PedidoProducto) []PedidoProducto {
	var productosElegidosDto []PedidoProducto
	for _, producto := range productosElegidos {
		productosElegidosDto = append(productosElegidosDto, *NewPedidoProducto(&producto))
	}
	return productosElegidosDto
}

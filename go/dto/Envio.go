package dto

import (
	"time"

	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/models"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/utils"
)

type Envio struct {
	ID                string    `json:"id,omitempty"`
	PatenteCamion     string    `json:"patente_camion"`
	Pedidos           []string  `json:"pedidos"`
	Paradas           []Parada  `json:"paradas"`
	Estado            string    `json:"estado"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}

func NewEnvio(envio models.Envio) *Envio {
	return &Envio{
		ID:                utils.GetStringIDFromObjectID(envio.ID),
		PatenteCamion:     envio.PatenteCamion,
		Pedidos:           envio.Pedidos,
		Paradas:           []Parada{},
		Estado:            "A despachar",
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}
}
func (envio Envio) GetModel() models.Envio {
	return models.Envio{
		ID:                utils.GetObjectIDFromStringID(envio.ID),
		PatenteCamion:     envio.PatenteCamion,
		Pedidos:           envio.Pedidos,
		Paradas:           envio.getParadas(),
		Estado:            envio.Estado,
		FechaCreacion:     envio.FechaCreacion,
		FechaModificacion: envio.FechaModificacion,
	}
}
func (envio Envio) getParadas() []models.Parada {
	var paradasEnvio []models.Parada
	for _, parada := range envio.Paradas {
		paradasEnvio = append(paradasEnvio, parada.GetModel())
	}
	return paradasEnvio
}

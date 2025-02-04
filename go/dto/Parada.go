package dto

import (
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/models"
)

type Parada struct {
	Ciudad       string `json:"ciudad"`
	KmRecorridos int    `json:"km_recorridos"`
}

func NewParada(parada models.Parada) *Parada {
	return &Parada{
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}

func (parada Parada) GetModel() models.Parada {
	return models.Parada{
		Ciudad:       parada.Ciudad,
		KmRecorridos: parada.KmRecorridos,
	}
}

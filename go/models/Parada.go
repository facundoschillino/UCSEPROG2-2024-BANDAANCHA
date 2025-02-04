package models

type Parada struct {
	Ciudad       string `bson:"ciudad"`
	KmRecorridos int    `bson:"kmRecorridos"`
}

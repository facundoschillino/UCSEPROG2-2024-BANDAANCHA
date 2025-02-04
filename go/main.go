package main

import (
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/database"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/handlers"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/repositories"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/services"
	"github.com/gin-gonic/gin"
)

var (
	productoHandler *handlers.ProductoHandler
	router          *gin.Engine
)

func main() {
	router := gin.Default
	dependencies()
	mappingRoutes()
	router.Run(":8080")
}

func mappingRoutes() {
	//Listado de rutas
	router.GET("/productos", productoHandler.ObtenerProductos)
	router.GET("/productos/:id", productoHandler.ObtenerProductoPorID)
	router.POST("/productos", productoHandler.CrearProducto)
	router.PUT("/productos/:id", productoHandler.ModificarProducto)
	router.DELETE("/productos/:id", productoHandler.EliminarProducto)
}

// Generacion de los objetos que se van a usar en la api
func dependencies() {
	//Definicion de variables de interface
	var database database.DB
	var productoRepository repositories.ProductoRepositoryInterface
	var productoService services.ProductoInterface

	//Creamos los objetos reales y los pasamos como parametro
	database = database.NewMongoDB()
	productoRepository = repositories.NewProductoRepository(database)
	productoService = services.NewProductoService(productoRepository)
	productoHandler = handlers.NewProductoHandler(productoService)
}

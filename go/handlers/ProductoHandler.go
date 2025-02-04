package handlers

import (
	"net/http"

	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/dto"
	"github.com/facundoschillino/UCSEPROG2-2024-BANDAANCHA/go/services"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	productoService services.ProductoInterface
}

func NewProductoHandler(productoService services.ProductoInterface) *ProductoHandler {
	return &ProductoHandler{
		productoService: productoService,
	}
}
func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {
	productos := handler.productoService.ObtenerProductos()
	c.JSON(http.StatusOK, productos)
}

func (handler *ProductoHandler) ObtenerProductoPorID(c *gin.Context) {
	id := c.Param("id")
	producto := handler.productoService.ObtenerProductoPorID(id)

	c.JSON(http.StatusOK, producto)
}

func (handler *ProductoHandler) CrearProducto(c *gin.Context) {
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.productoService.CrearProducto(&producto)

	c.JSON(http.StatusCreated, resultado)
}

func (handler *ProductoHandler) ModificarProducto(c *gin.Context) {
	var producto dto.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	producto.ID = c.Param("id")
	resultado := handler.productoService.ModificarProducto(&producto)

	c.JSON(http.StatusCreated, resultado)
}

func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	productos := handler.productoService.EliminarProducto(id)

	c.JSON(http.StatusOK, productos)
}

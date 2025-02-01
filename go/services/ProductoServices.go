package services

import (
	"Github/UCSEPROG2-2024-BANDAANCHA/go/dto"
	"Github/UCSEPROG2-2024-BANDAANCHA/go/repositories"
	"Github/UCSEPROG2-2024-BANDAANCHA/go/utils"
)

type ProductoInterface interface {
	ObtenerProductos() []*dto.Producto
	ObtenerProductoPorID(id string) *dto.Producto
	EliminarProducto(id string) bool
	CrearProducto(producto *dto.Producto) bool
	ModificarProducto(prodcuto *dto.Producto) bool
}

type ProductoService struct {
	productoRepository repositories.ProductoRepositoryInterface
}

func NewProductoService(productoRepository repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		productoRepository: productoRepository,
	}

}

func (service *ProductoService) ObtenerProductos() []*dto.Producto {
	productosDB, _ := service.productoRepository.ObtenerProductos()
	var productos []*dto.Producto
	for _, productoDB := range productosDB {
		producto := dto.NewProducto(productoDB)
		productos = append(productos, producto)
	}
	return productos
}
func (service *ProductoService) ObtenerProductoPorID(id string) *dto.Producto {
	productoDB, err := service.productoRepository.ObtenerProductoPorID(id)

	var producto *dto.Producto
	if err == nil {
		producto = dto.NewProducto(productoDB)
	}

	return producto
}
func (service *ProductoService) CrearProducto(producto *dto.Producto) bool {
	service.productoRepository.CrearProducto(producto.GetModel())

	return true
}

func (service *ProductoService) ModificarProducto(producto *dto.Producto) bool {
	service.productoRepository.ModificarProducto(producto.GetModel())

	return true
}

func (service *ProductoService) EliminarProducto(id string) bool {
	service.productoRepository.EliminarProducto(utils.GetObjectIDFromStringID(id))

	return true
}

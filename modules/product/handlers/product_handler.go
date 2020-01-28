package handlers

import (
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/goproductapi/common/resolvers"
	"github.com/goproductapi/modules/product/dto"

	"github.com/gin-gonic/gin"
)

// ProductResolver struct
type ProductHandler struct {
	resolver *resolvers.Resolver
}

// NewProductResolver initialize new product handler object
func NewProductHandler(resolver *resolvers.Resolver) *ProductHandler {
	handler := &ProductHandler{
		resolver: resolver,
	}

	return handler
}

// Get product by UUID
func (h *ProductHandler) GetByUUID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.resolver.ProductService.GetByUUID(uuid.FromStringOrNil(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// Submit product
func (h *ProductHandler) Submit(ctx *gin.Context) {
	var request dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if request.Name == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}

	product, err := h.resolver.ProductService.Create(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// GetAll products
func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := h.resolver.ProductService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

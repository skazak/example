package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skazak/example/internal/app/delivery/http/response"
	"github.com/skazak/example/internal/app/model"
)

// ItemHandler represents the http handler for model.Item
type ItemHandler struct {
	Service model.ItemService
}

// InitItemHandler initializes the model.Item resource endpoints
func InitItemHandler(r *gin.RouterGroup, is model.ItemService) {
	handler := &ItemHandler{
		Service: is,
	}

	r.GET("/category/:id/items/", handler.GetByCategoryID)
	r.GET("/item/:id", handler.GetByID)
	r.POST("/item", handler.Store)
	r.PUT("/item/:id", handler.Update)
	r.DELETE("/item/:id", handler.Delete)
}

// Get gets all categories
func (ih *ItemHandler) Get(ctx *gin.Context) {
	categories, err := ih.Service.Get(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// GetByID gets item by given id
func (ih *ItemHandler) GetByCategoryID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	items, err := ih.Service.GetByCategoryID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// GetByID gets item by given id
func (ih *ItemHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	item, err := ih.Service.GetByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// Store stores the item by given request body
func (ih *ItemHandler) Store(ctx *gin.Context) {
	var item model.Item

	err := ctx.Bind(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = ih.Service.Store(ctx, &item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// Update updates the item by given request body
func (ih *ItemHandler) Update(ctx *gin.Context) {
	var item model.Item

	err := ctx.Bind(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = ih.Service.Update(ctx, &item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// Delete deletes item with provided ID.
func (ih *ItemHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = ih.Service.Delete(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
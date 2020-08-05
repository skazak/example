package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skazak/example/internal/app/delivery/http/response"
	"github.com/skazak/example/internal/app/model"
)

// CategoryHandler represents the http handler for model.Category
type CategoryHandler struct {
	Service model.CategoryService
}

// InitCategoryHandler initializes the model.Category resource endpoints
func InitCategoryHandler(r *gin.RouterGroup, cs model.CategoryService) {
	handler := &CategoryHandler{
		Service: cs,
	}

	r.GET("/category", handler.Get)
	r.GET("/category/:id", handler.GetByID)
	r.POST("/category", handler.Store)
	r.PUT("/category/:id", handler.Update)
	r.DELETE("/category/:id", handler.Delete)
}

// Get gets all categories
func (ch *CategoryHandler) Get(ctx *gin.Context) {

	if ctx.Query("full") == "true" {
		categories, err := ch.Service.GetWithItems(ctx)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, categories)
		return
	}

	categories, err := ch.Service.Get(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// GetByID gets category by given id
func (ch *CategoryHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	category, err := ch.Service.GetByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// Store stores the category by given request body
func (ch *CategoryHandler) Store(ctx *gin.Context) {
	var category model.Category

	err := ctx.Bind(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = ch.Service.Store(ctx, &category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// Update updates the category by given request body
func (ch *CategoryHandler) Update(ctx *gin.Context) {
	var category model.Category

	err := ctx.Bind(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	category.ID = id

	err = ch.Service.Update(ctx, &category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// Delete deletes category with provided ID.
func (ch *CategoryHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = ch.Service.Delete(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
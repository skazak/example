package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skazak/example/internal/app/delivery/http/response"
	"github.com/skazak/example/internal/app/model"
)

// UserHandler represents the http handler for model.User
type UserHandler struct {
	Service model.UserService
}

// InitUserHandler initializes the model.User resource endpoints
func InitUserHandler(r *gin.RouterGroup, is model.UserService) {
	handler := &UserHandler{
		Service: is,
	}

	r.GET("/user/:id", handler.GetByID)
}

// GetByID gets user by given id
func (uh *UserHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := uh.Service.GetByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

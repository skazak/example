package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/skazak/example/internal/app/delivery/http/auth"
	"github.com/skazak/example/internal/app/delivery/http/request"
	"github.com/skazak/example/internal/app/delivery/http/response"
	"github.com/skazak/example/internal/app/model"
)

// AuthHandler represents the http handler for authorization needs
type AuthHandler struct {
	Service model.UserService
}

// hashSHA256 is used for password hash calculation
func hashSHA256(str string) string {
	s := sha256.New()
	s.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(s.Sum(nil))
}

// Auth performs User authentication by Email & Pass
func (ah *AuthHandler) Auth(ctx *gin.Context) {
	var req *request.AuthRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	user, err := ah.Service.GetByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if user.Password != hashSHA256(req.Password) {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Message: "wrong password",
		})
		return
	}

	token, err := auth.NewClaims(user.ID, time.Now().Add(36*time.Hour).Unix()).Encode()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.AuthResponse{
		Token: token,
	})
}

// Register stores new User to DB if Email is not duplicated
func (ah *AuthHandler) Register(ctx *gin.Context) {
	var user *model.User
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	_, err = ah.Service.GetByEmail(ctx, user.Email)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "user already exists",
		})
		return
	}

	if err != pgx.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	user.Password = hashSHA256(user.Password)

	err = ah.Service.Store(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// InitUserHandler initializes the model.User resource endpoints
func InitAuthHandler(r *gin.RouterGroup, us model.UserService) {
	handler := &AuthHandler{
		Service: us,
	}

	r.POST("/auth", handler.Auth)
	r.POST("/register", handler.Register)
}
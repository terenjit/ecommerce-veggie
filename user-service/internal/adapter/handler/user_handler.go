package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/terenjit/ecommerce-veggie/user-service/config"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter/handler/request"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter/handler/response"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/domain/entity"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/service"
)

type UserHandlerInterface interface {
	SignIn(ctx echo.Context) error
}

type UserHandler struct {
	userService service.UserServiceInterface
	cfg         *config.Config
}

func NewUserHandler(e *echo.Echo, userService service.UserServiceInterface, cfg *config.Config) UserHandlerInterface {
	userHandler := &UserHandler{userService: userService, cfg: cfg}

	e.Use(middleware.Recover())
	e.POST("/signin", userHandler.SignIn)

	mid := adapter.NewMiddlewareAdapter(cfg)
	adminGroup := e.Group("/admin", mid.CheckToken())
	adminGroup.GET("/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	return userHandler
}

func (h *UserHandler) SignIn(c echo.Context) error {
	var (
		req        = request.SignInRequest{}
		resp       = response.Response{}
		respSignIn = response.SignInResponse{}
		ctx        = c.Request().Context()
	)

	err := c.Bind(&req)
	if err != nil {
		log.Error().Err(err).Msgf("SignIn : bind request failed: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	err = c.Validate(req)
	if err != nil {
		log.Error().Err(err).Msgf("SignIn : validate request failed: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	reqEntity := entity.UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}

	user, token, err := h.userService.SignIn(ctx, &reqEntity)
	if err != nil {
		if err.Error() == "404" {
			log.Error().Err(err).Msgf("SignIn : user not found for email: %s", req.Email)
			resp.Message = "user not found"
			resp.Data = nil
			return c.JSON(http.StatusNotFound, resp)
		}
		log.Error().Err(err).Msgf("SignIn : user sign in failed: %v", err)
		resp.Message = err.Error()
		resp.Data = nil
		return c.JSON(http.StatusInternalServerError, resp)
	}

	respSignIn.ID = user.ID
	respSignIn.Email = user.Email
	respSignIn.Name = user.Name
	respSignIn.Phone = user.Phone
	respSignIn.Role = user.RoleName
	respSignIn.Lat = user.Lat
	respSignIn.Lng = user.Lng
	respSignIn.AccessToken = token

	resp.Message = "User signed in successfully"
	resp.Data = respSignIn
	return c.JSON(http.StatusOK, resp)
}

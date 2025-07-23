package http

import (
	"net/http"

	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/application/dto"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/application/mapper"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", c.CreateUser)
		userRoutes.GET("/:id", c.FindById)
		userRoutes.GET("", c.FindAll)
		userRoutes.PUT("/:id", c.Update)
		userRoutes.DELETE("/:id", c.Delete)
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

	var user dto.UserRequestDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := c.userService.SaveUser(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := mapper.ToUserResponse(savedUser)

	ctx.JSON(http.StatusOK, res)

}

func (c *UserController) FindById(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User id is required"})
		return
	}

	user, err := c.userService.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := mapper.ToUserResponse(user)

	ctx.JSON(http.StatusOK, res)

}

func (c *UserController) FindAll(ctx *gin.Context) {

	users, err := c.userService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res []*dto.UserResponseDTO

	for _, user := range users {
		res = append(res, mapper.ToUserResponse(user))
	}

	if res == nil {
		res = []*dto.UserResponseDTO{}
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User id is required"})
		return
	}
	if err := c.userService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (c *UserController) Update(ctx *gin.Context) {
	var user dto.UserRequestDTO
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User id is required"})
		return
	}

	updatedUser, err := c.userService.UpdateUser(user, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := mapper.ToUserResponse(updatedUser)

	ctx.JSON(http.StatusOK, res)

}

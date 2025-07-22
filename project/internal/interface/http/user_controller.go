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
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

	var user dto.UserRequestDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := c.userService.SaveUser(user.Name, user.Email)

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

	response := mapper.ToUserResponse(user)

	ctx.JSON(http.StatusOK, response)

}

func (c *UserController) FindAll(ctx *gin.Context) {

	users, err := c.userService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var userResponse []*dto.UserResponseDTO

	for _, user := range users {
		userResponse = append(userResponse, mapper.ToUserResponse(user))
	}

	ctx.JSON(http.StatusOK, userResponse)
}

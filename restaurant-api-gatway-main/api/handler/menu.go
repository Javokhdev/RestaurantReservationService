package handler

import (
	"strconv"

	pb "github.com/dilshodforever/restaurant-submodule/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateMenu 		handles the creation of a new Menu
// @Summary 		Create Menu
// @Description 	Create page
// @Tags 			Menu
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Menu    true   "Create"
// @Success 		200   {string}  string    "Create Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/menu/create [post]
func (h *Handler) CreateMenu(ctx *gin.Context) {
	arr := pb.Menu{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Menu.CreateMenu(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateMenu 		handles the creation of a new Menu
// @Summary 		Update Menu
// @Description 	Update page
// @Tags 			Menu
// @Accept  		json
// @Produce  		json
// @Param     		id 		path    string   true    "Menu ID"
// @Param   		Update  body    pb.Menu  true    "Update"
// @Success 		200   {string}  string   "Update Successful"
// @Failure 		401   {string}  string   "Error while created"
// @Router			 /Menu/update/{id} [put]
func (h *Handler) UpdateMenu(ctx *gin.Context) {
	arr := pb.Menu{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Menu.UpdateMenu(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteMenu 		handles the creation of a new Menu
// @Summary 		Delete Menu
// @Description 	Delete page
// @Tags 			Menu
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true    "Menu ID"
// @Success 		200 {string}  string  "Delete Successful"
// @Failure 		401 {string}  string  "Error while Deleted"
// @Router 			/Menu/delete/{id} [delete]
func (h *Handler) DeleteMenu(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Menu.DeleteMenu(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllMenu 		handles the creation of a new Menu
// @Summary 		GetAll Menu
// @Description 	GetAll page
// @Tags			Menu
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Menu true    "Query parameter"
// @Success 		200  {object}  pb.GetAllMenus  "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAlld"
// @Router 			/Menu/getall [get]
func (h *Handler) GetAllMenu(ctx *gin.Context) {
	menu := &pb.Menu{}
	menu.Name = ctx.Param("email")
	p, _ := strconv.ParseFloat(ctx.Param("price"), 32)
	menu.Price = float32(p)
	menu.Description = ctx.Param("description")

	res, err := h.Menu.GetAllMenu(ctx, menu)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdMenu 		handles the creation of a new Menu
// @Summary 		GetById Menu
// @Description 	GetById page
// @Tags 			Menu
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string     true   "Menu ID"
// @Success 		200  {object}  pb.Menu   "GetById Successful"
// @Failure 		401  {string}  string    "Error while GetByIdd"
// @Router 			/Menu/getbyid/{id} [get]
func (h *Handler) GetByIdMenu(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Menu.GetByIdMenu(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

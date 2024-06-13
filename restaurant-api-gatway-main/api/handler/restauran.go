package handler

import (
	pb "github.com/dilshodforever/restaurant-submodule/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Restoran handles the creation of a new Public_Vote
// @Summary 		Create Public_Vote
// @Description 	Create page
// @Tags 			Restoran
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Restoran  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/restoran/create [post]
func (h *Handler) CreateRestoran(ctx *gin.Context) {
	arr := pb.Restoran{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Restoran.CreateRestoran(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateRestoran 	handles the creation of a new Restoran
// @Summary 		Update Restoran
// @Description 	Update page
// @Tags 			Restoran
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Restoran ID"
// @Param   		Update  body    pb.Restoran  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/restoran/update/{id} [put]
func (h *Handler) UpdateRestoran(ctx *gin.Context) {
	arr := pb.Restoran{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Restoran.UpdateRestoran(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteRestoran 	handles the creation of a new Restoran
// @Summary 		Delete Restoran
// @Description 	Delete page
// @Tags 			Restoran
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Restoran ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/restoran/delete/{id} [delete]
func (h *Handler) DeleteRestoran(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Restoran.DeleteRestoran(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllRestoran 	handles the creation of a new Restoran
// @Summary 		GetAll Restoran
// @Description 	GetAll page
// @Tags 			Restoran
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Restoran true  "Query parameter"
// @Success 		200  {object}  pb.GetAllRestorans  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/restoran/getall [get]
func (h *Handler) GetAllRestoran(ctx *gin.Context) {
	restoran := &pb.Restoran{}
	restoran.Address = ctx.Param("restoran")
	restoran.Name = ctx.Param("name")

	res, err := h.Restoran.GetAllRestoran(ctx, restoran)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdRestoran 	handles the creation of a new Restoran
// @Summary 		GetById Restoran
// @Description 	GetById page
// @Tags 			Restoran
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Restoran ID"
// @Success 		200 {object}  pb.Restoran   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/restoran/getbyid/{id} [get]
func (h *Handler) GetByIdRestoran(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Restoran.GetByIdRestoran(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

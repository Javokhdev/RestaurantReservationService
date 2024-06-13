package handler

import (
	pb "github.com/dilshodforever/restaurant-submodule/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			reservation_order handles the creation of a new Public_Vote
// @Summary 		Create Public_Vote
// @Description 	Create page
// @Tags 			ReservationOrder
// @Accept  		json
// @Produce  		json
// @Param   		Create   body    pb.ReservationOrder   true    "Create"
// @Success			200    {string}  string  			   "Create Successful"
// @Failure 		401    {string}  string  			   "Error while Created"
// @Router 			/public_vote/create [post]
func (h *Handler) CreateReservationOrder(ctx *gin.Context) {
	arr := pb.ReservationOrder{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.ReservationOrder.CreateReservationOrder(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateReservationOrder 	handles the creation of a new ReservationOrder
// @Summary 				Update ReservationOrder
// @Description 			Update page
// @Tags 					ReservationOrder
// @Accept  				json
// @Produce  				json
// @Param     				id path string true "ReservationOrder ID"
// @Param   				Update   body    pb.ReservationOrder  true   "Update"
// @Success 				200    {string}  string  		     "Update Successful"
// @Failure 				401    {string}  string  		     "Error while created"
// @Router 					/ReservationOrder/update/{id} [put]
func (h *Handler) UpdateReservationOrder(ctx *gin.Context) {
	arr := pb.ReservationOrder{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.ReservationOrder.UpdateReservationOrder(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteReservationOrder 	handles the creation of a new ReservationOrder
// @Summary 				Delete ReservationOrder
// @Description 			Delete page
// @Tags 					ReservationOrder
// @Accept  				json
// @Produce  				json
// @Param     				id     path    string   true "ReservationOrder ID"
// @Success 				200  {string}  string  "Delete Successful"
// @Failure 				401  {string}  string  "Error while Deleted"
// @Router 					/ReservationOrder/delete/{id} [delete]
func (h *Handler) DeleteReservationOrder(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.ReservationOrder.DeleteReservationOrder(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllReservationOrder 	handles the creation of a new ReservationOrder
// @Summary 				GetAll ReservationOrder
// @Description 			GetAll page
// @Tags 					ReservationOrder
// @Accept  				json
// @Produce 			 	json
// @Param 					query  query   pb.ReservationOrder true    "Query parameter"
// @Success 				200  {object}  pb.GetAllReservationOrders  "GetAll Successful"
// @Failure 				401  {string}  string  					   "Error while GetAlld"
// @Router 					/ReservationOrder/getall [get]
func (h *Handler) GetAllReservationOrder(ctx *gin.Context) {
	reservationOrder := &pb.ReservationOrder{}
	reservationOrder.Quantity = ctx.Param("quantity")

	res, err := h.ReservationOrder.GetAllReservationOrder(ctx, reservationOrder)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdReservationOrder 	handles the creation of a new ReservationOrder
// @Summary 				GetById ReservationOrder
// @Description 			GetById page
// @Tags 					ReservationOrder
// @Accept  				json
// @Produce  				json
// @Param     				id    path    string    true 		"ReservationOrder ID"
// @Success 				200 {object}  pb.ReservationOrder   "GetById Successful"
// @Failure 				401 {string}  string 				"Error while GetByIdd"
// @Router 					/reservationOrder/getbyid/{id} [get]
func (h *Handler) GetByIdReservationOrder(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.ReservationOrder.GetByIdReservationOrder(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

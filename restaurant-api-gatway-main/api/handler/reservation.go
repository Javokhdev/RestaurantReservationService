package handler

import (
	pb "github.com/dilshodforever/restaurant-submodule/genprotos"

	"github.com/gin-gonic/gin"
)



// CreateReservation 	handles the creation of a new Reservation
// @Summary 			Create Reservation
// @Description 		Create page
// @Tags 				Reservation
// @Accept  			json
// @Produce  			json
// @Param     			user_id path string true "Restoran ID"
// @Param     			restaurant_id  path string true "User id"
// @Param   			Create   body    pb.Reservation   true   "Create"
// @Success 			200    {string}  string          "Create Successful"
// @Failure 			401    {string}  string          "Error while Created"
// @Router       		/reservation/create/{user_id}/{restaurant_id} [post]
func (h *Handler) CreateReservation(ctx *gin.Context) {
	userid := ctx.Param("user_id")
	restoranid := ctx.Param("restaurant_id")
	arr := pb.Reservation{}
	arr.UserId = userid
	arr.RestaurantId = restoranid
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Reservation.CreateReservation(ctx, &arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, "Sucsess!!!")
}

// UpdateReservation 	handles the creation of a new Reservation
// @Summary 			Update Reservation
// @Description 		Update page
// @Tags 				Reservation
// @Accept  			json
// @Produce  			json
// @Param     			id path string true "Reservation ID"
// @Param   			Update   body    pb.Reservation  true   "Update"
// @Success 			200    {string}  string         "Update Successful"
// @Failure 			401    {string}  string         "Error while created"
// @Router 				/Reservation/update/{id} [put]
func (h *Handler) UpdateReservation(ctx *gin.Context) {
	arr := pb.Reservation{Id: ctx.Param("id")}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Reservation.UpdateReservation(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}

// DeleteReservation 	handles the creation of a new Reservation
// @Summary 			Delete Reservation
// @Description 		Delete page
// @Tags 				Reservation
// @Accept  			json
// @Produce 			json
// @Param     			id path string true "Reservation ID"
// @Success 			200 {string} string  "Delete Successful"
// @Failure 			401 {string} string  "Error while Deleted"
// @Router 				/Reservation/delete/{id} [delete]
func (h *Handler) DeleteReservation(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Reservation.DeleteReservation(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}

// GetAllReservation 	handles the creation of a new Reservation
// @Summary 			GetAll Reservation
// @Description 		GetAll page
// @Tags 				Reservation
// @Accept  			json
// @Produce 			json
// @Param 				query query  pb.Reservation true "Query parameter"
// @Success 			200 {object} pb.GetAllReservations    "GetAll Successful"
// @Failure 			401 {string} string          		  "Error while GetAlld"
// @Router 				/reservation/getall [get]
func (h *Handler) GetAllReservation(ctx *gin.Context) {
	reservation := &pb.Reservation{}
	reservation.ReservationTime = ctx.Param("reservation_time")
	reservation.Status = ctx.Param("status")

	res, err := h.Reservation.GetAllReservation(ctx, reservation)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdReservation 	handles the creation of a new Reservation
// @Summary 			GetById Reservation
// @Description 		GetById page
// @Tags 				Reservation
// @Accept  			json
// @Produce  			json
// @Param    			id    path    string    true   "Reservation ID"
// @Success 			200 {object}  pb.Reservation   "GetById Successful"
// @Failure 			401 {string}  string 		   "Error while GetByIdd"
// @Router 				/reservation/getbyid/{id} [get]
func (h *Handler) GetbyIdReservation(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Reservation.GetByIdReservation(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}



// CreateReservation 	handles the creation of a new Reservation
// @Summary 			Create Reservation
// @Description 		Create page
// @Tags 				Reservation
// @Accept  			json
// @Produce  			json
// @Param   			Create   body    pb.Reservation   true   "Create"
// @Success 			200    {string}  string          "Create Successful"
// @Failure 			401    {string}  string          "Error while Created"
// @Router       		/reservation/bron [post]
func (h *Handler) Reservations(ctx *gin.Context) {
	arr := pb.Reservation{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Reservation.Reservations(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Sucsess!!!")
}
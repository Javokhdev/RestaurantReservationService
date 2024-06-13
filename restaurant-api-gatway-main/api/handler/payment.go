package handler

import (
	pb "github.com/dilshodforever/restaurant-submodule/genprotos"
	"github.com/gin-gonic/gin"
)
type Payments struct{
	Paymen *pb.GetAllPayments
	Reservations *pb.GetAllReservations
}
// CreatePayment 	handles the creation of a new Payment
// @Summary 		Create Payment
// @Description 	Create page
// @Tags 			Payment
// @Accept  		json
// @Produce  		json
// @Param   		Create   body    pb.Payment    true   "Create"
// @Success 		200    {string}  string       "Create Successful"
// @Failure 		401    {string}  string       "Error while Created"
// @Router 			/Payment/create [post]
func (h *Handler) CreatePayment(ctx *gin.Context) {
	arr := pb.Payment{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Payment.CreatePayment(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdatePayment 	handles the creation of a new Payment
// @Summary 		Update Payment
// @Description 	Update page
// @Tags 			Payment
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Payment ID"
// @Param   		Update   body    pb.Payment   true   "Update"
// @Success 		200    {string}  string      "Update Successful"
// @Failure 		401    {string}  string      "Error while created"
// @Router 			/Payment/update/{id} [put]
func (h *Handler) UpdatePayment(ctx *gin.Context) {
	arr := pb.Payment{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Payment.UpdatePayment(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeletePayment 	handles the creation of a new Payment
// @Summary 		Delete Payment
// @Description 	Delete page
// @Tags 			Payment
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string   true   "Payment ID"
// @Success 		200 {string}  string  "Delete Successful"
// @Failure 		401 {string}  string  "Error  while Deleted"
// @Router 			/Payment/delete/{id} [delete]
func (h *Handler) DeletePayment(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Payment.DeletePayment(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllPayment 	handles the creation of a new Payment
// @Summary 		GetAll Payment
// @Description 	GetAll page
// @Tags 			Payment
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Payment   true   "Query parameter"
// @Success 		200  {object}  Payments{}  "GetAll Successful"
// @Failure 		401  {string}  string  			   "Error while GetAlld"
// @Router 			/payment/getall [get]
func (h *Handler) GetAllPayment(ctx *gin.Context) {
	payment := &pb.Payment{}
	payment.PaymentMethod = ctx.Param("payment_method")
	payment.PaymentStatus = ctx.Param("payment_status")
	pays:=Payments{}
	pay, err := h.Payment.GetAllPayment(ctx, payment)
	if err != nil {
		panic(err)
	}
	reservation := &pb.Reservation{}
	reservation.ReservationTime = ctx.Param("reservation_time")
	reservation.Status = ctx.Param("status")

	res, err:=h.Reservation.GetAllReservation(ctx, reservation)
	if err != nil {
		panic(err)
	}
	pays.Paymen=pay
	pays.Reservations=res

	ctx.JSON(200, pays)
}

// GetByIdPayment 	handles the creation of a new Payment
// @Summary 		GetById Payment
// @Description 	GetById page
// @Tags 			Payment
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string true  "Payment ID"
// @Success 		200  {object}  pb.Payment   "GetById Successful"
// @Failure 		401  {string}  string       "Error while GetByIdd"
// @Router 			/Payment/getbyid/{id} [get]
func (h *Handler) GetByIdPayment(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Payment.GetByIdPayment(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

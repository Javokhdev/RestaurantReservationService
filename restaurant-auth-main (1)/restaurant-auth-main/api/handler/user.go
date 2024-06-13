package handler

import (
	"github.com/dilshodforever/restaurant-auth/api/token"

	pb "github.com/dilshodforever/restaurant-auth/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateUser 		handles the creation of a new user
// @Summary 		Create User
// @Description 	Create page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.User    true   "Create"
// @Success 		200   {string}   pb.User   "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/user/registr [post]
func (h *Handler) RegisterUser(ctx *gin.Context){
		arr:=pb.User{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.User.CreateUser(ctx, &arr)
		if err!=nil{
			panic(err)
		}
	t:=token.GenereteJWTToken(&arr)
	ctx.JSON(200, t)
}

// UpdateUser 		handles the creation of a new user
// @Summary			Update User
// @Description 	Update page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id 		path   string     true   "User ID"
// @Param   		Update  body   pb.User    true   "Update"
// @Success 		200   {string} string    "Update Successful"
// @Failure 		401   {string} string    "Error while created"
// @Router 			/User/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context){
	arr:=pb.User{}
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.User.UpdateUser(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}


// DeleteUser 		handles the creation of a new User
// @Summary			Delete User
// @Description 	Delete page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id   path     string   true   "User ID"
// @Success 		200 {string}  string   "Delete Successful"
// @Failure 		401 {string}  string   "Error while Deleted"
// @Router 			/user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.User.DeleteUser(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllUser 		handles the creation of a new User
// @Summary 		GetAll User
// @Description 	GetAll page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param 			query  query  pb.User true    "Query parameter"
// @Success 		200 {object}  pb.GetAllUsers  "GetAll Successful"
// @Failure 		401 {string}  string  		  "Error while GetAll"
// @Router 			/user/getall  [get]
func (h *Handler) GetAllUser(ctx *gin.Context){
	user := &pb.User{}
	user.Email = ctx.Param("email")
	user.Password = ctx.Param("password")
	user.UserName = ctx.Param("user_name")

	res, err:=h.User.GetAllUser(ctx, user)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdUser 		handles the creation of a new User
// @Summary 		GetById User
// @Description 	GetById page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id   path      string   true    "User ID"
// @Success 		200 {object}   pb.User  "GetById Successful"
// @Failure 		401 {string}   string   "Error while GetByIdd"
// @Router 			/User/getbyid/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.User.GetByIdUser(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdUser 		handles the creation of a new User
// @Summary 		/LoginUser
// @Description 	LoginUser page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body  pb.User   true     "Create"
// @Success 		200 {object}  pb.User  "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context){
	user:=&pb.User{}
	err:=ctx.ShouldBindJSON(user)
	if err!=nil{
		panic(err)
	}
	res, err:=h.User.LoginUser(ctx, user)
	if err!=nil{
		panic(err)
	}
	t:=token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}


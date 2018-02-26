package router

import (
	"goji.io"
	"goji.io/pat"

	"github.com/mrizkip/backend-finding-dosen/handlers"
	"github.com/mrizkip/backend-finding-dosen/middlewares"
)

func NewRouter() *goji.Mux {
	rootRoute := goji.NewMux()
	rootRoute.HandleFunc(pat.Post("/login"), handlers.Login)
	rootRoute.HandleFunc(pat.Post("/register"), handlers.Register)

	userRoute := goji.SubMux()
	userRoute.Use(middlewares.VerifyToken)
	userRoute.HandleFunc(pat.Get("/my_profile"), handlers.FetchMyProfile)
	userRoute.HandleFunc(pat.Post("/change_profile"), handlers.ChangeProfile)
	userRoute.HandleFunc(pat.Post("/change_password"), handlers.ChangePassword)
	userRoute.HandleFunc(pat.Get("/fetch_dosens"), handlers.FetchAllDosenProfile)
	userRoute.HandleFunc(pat.Get("/:id/profile"), handlers.FetchUserProfileByID)

	dosenRoute := goji.SubMux()
	dosenRoute.Use(middlewares.VerifyToken)
	dosenRoute.Use(middlewares.VerifyRoleDosen)
	dosenRoute.HandleFunc(pat.Post("/update_status"), handlers.UpdateUserStatus)

	rootRoute.Handle(pat.New("/user/*"), userRoute)
	rootRoute.Handle(pat.New("/dosen/*"), dosenRoute)

	return rootRoute
}

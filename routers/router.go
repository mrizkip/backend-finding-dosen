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
	userRoute.HandleFunc(pat.Get("/:id/profile"), handlers.FetchUserProfile)

	dosenRoute := goji.SubMux()
	dosenRoute.Use(middlewares.VerifyToken)
	dosenRoute.Use(middlewares.VerifyRoleDosen)
	dosenRoute.HandleFunc(pat.Post("/update_status"), handlers.UpdateUserStatus)

	/*
		userRoute := goji.SubMux()
		userRoute.Use(middlewares.VerifyToken)
		userRoute.HandleFunc(pat.Post("/change_password"), handlers.ChangePassword)
		userRoute.HandleFunc(pat.Post("/change_profile"), handlers.ChangeProfile)
		userRoute.HandleFunc(pat.Post("/change_profile_photo"), handlers.ChangeProfilePhoto)
		userRoute.HandleFunc(pat.Get("/my_profile"), handlers.FetchMyProfile)
		userRoute.HandleFunc(pat.Post("/update_current_location"), handlers.UpdateUserCurrentLocation)
		userRoute.HandleFunc(pat.Get("/:username/profile"), handlers.FetchUserProfileByUsername)

		promoRoute := goji.SubMux()
		promoRoute.Use(middlewares.VerifyToken)
		promoRoute.HandleFunc(pat.Post("/new"), handlers.NewPromo)
		promoRoute.HandleFunc(pat.Get("/my_promos"), handlers.FetchMyPromos)
		promoRoute.HandleFunc(pat.Get("/my_promo/:id/attendees"), handlers.FetchMyPromoAttendeeByPromoId)
		promoRoute.HandleFunc(pat.Delete("/my_promo/:promo_id/attendee/:attendee_id"), handlers.DeletePromoAttendee)
		promoRoute.HandleFunc(pat.Get("/my_joined_promos"), handlers.FetchJoinedPromo)
		promoRoute.HandleFunc(pat.Post("/nearby_promos"), handlers.FetchNearbyPromosFromMyLocation)
		promoRoute.HandleFunc(pat.Get("/admin_promoted"), handlers.FetchAdminPromotedPromo)
		promoRoute.HandleFunc(pat.Post("/find"), handlers.FetchPromoByName)
		promoRoute.HandleFunc(pat.Get("/:id"), handlers.FetchPromoById)
		promoRoute.HandleFunc(pat.Delete("/:id"), handlers.DeleteMyPromoById)
		promoRoute.HandleFunc(pat.Post("/:id"), handlers.ChangeMyPromoById)
		promoRoute.HandleFunc(pat.Post("/:id/join"), handlers.JoinPromo)

		rootRoute.Handle(pat.New("/user/*"), userRoute)
		rootRoute.Handle(pat.New("/promo/*"), promoRoute)
	*/

	rootRoute.Handle(pat.New("/user/*"), userRoute)
	rootRoute.Handle(pat.New("/dosen/*"), dosenRoute)

	return rootRoute
}

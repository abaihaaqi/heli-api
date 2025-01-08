package api

import (
	"fmt"
	"net/http"

	"github.com/ijaybaihaqi/heli-api/service"
	"github.com/rs/cors"
)

type API struct {
	userService          service.UserService
	sessionService       service.SessionService
	applianceService     service.ApplianceService
	userApplianceService service.UserApplianceService
	consumptionService   service.ConsumptionService
	chatService          service.ChatService
	mux                  *http.ServeMux
}

func NewAPI(userService service.UserService, sessionService service.SessionService, applianceService service.ApplianceService, userApplianceService service.UserApplianceService, consumptionService service.ConsumptionService, chatService service.ChatService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		sessionService,
		applianceService,
		userApplianceService,
		consumptionService,
		chatService,
		mux,
	}

	// User
	mux.Handle("/auth/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/auth/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/auth/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	// Appliance
	mux.Handle("/appliance/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllAppliance))))

	// User Appliance
	mux.Handle("/user-appliance/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllUserAppliance))))
	mux.Handle("/user-appliance/get-all-rooms", api.Get(api.Auth(http.HandlerFunc(api.FetchUserApplianceRooms))))
	mux.Handle("/user-appliance/get", api.Get(api.Auth(http.HandlerFunc(api.FetchUserApplianceByID))))
	mux.Handle("/user-appliance/add", api.Post(api.Auth(http.HandlerFunc(api.StoreUserAppliance))))
	mux.Handle("/user-appliance/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateUserAppliance))))
	mux.Handle("/user-appliance/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteUserAppliance))))

	// Consumption
	mux.Handle("/consumption/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllConsumptions))))
	mux.Handle("/consumption/add", api.Post(api.Auth(http.HandlerFunc(api.StoreConsumption))))
	mux.Handle("/consumption/reset", api.Delete(api.Auth(http.HandlerFunc(api.ResetConsumption))))

	// Chat with AI
	mux.Handle("/chat/analyze-data", api.Post(api.Auth(http.HandlerFunc(api.AnalyzeData))))
	mux.Handle("/chat/new", api.Post(api.Auth(http.HandlerFunc(api.ChatWithAI))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(api.Handler())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", corsHandler)
}

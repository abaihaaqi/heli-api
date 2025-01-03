package main

import (
	"github.com/ijaybaihaqi/heli-api/api"
	"github.com/ijaybaihaqi/heli-api/db"
	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
	"github.com/ijaybaihaqi/heli-api/service"
)

func main() {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "heli_db",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	// err = conn.Migrator().DropTable("appliances")
	// if err != nil {
	// 	panic(err)
	// }

	conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Appliance{}, &model.UserAppliance{}, &model.Consumption{})

	// appliances := []model.Appliance{
	// 	{
	// 		Name:  "Refrigerator",
	// 		Image: "https://upload.wikimedia.org/wikipedia/commons/2/2e/US_Domestic_Refrigerator_-_Frigidaire.jpg",
	// 	},
	// 	{
	// 		Name:  "TV",
	// 		Image: "https://upload.wikimedia.org/wikipedia/commons/7/78/Early_portable_tv.jpg",
	// 	},
	// 	{
	// 		Name:  "EVCar",
	// 		Image: "https://upload.wikimedia.org/wikipedia/commons/8/8e/Ford_Explorer_EV_Auto_Zuerich_2023_1X7A1325.jpg",
	// 	},
	// 	{
	// 		Name:  "Lamp",
	// 		Image: "https://upload.wikimedia.org/wikipedia/commons/2/2f/Lamp_with_a_lampshade_illuminated_by_sunlight.jpg",
	// 	},
	// 	{
	// 		Name:  "Air Conditioner",
	// 		Image: "https://upload.wikimedia.org/wikipedia/commons/7/73/Room_air_conditioning_unit_above_a_green_curtain.jpg",
	// 	},
	// 	{
	// 		Name:  "Power Strip",
	// 		Image: "https://upload.wikimedia.org/wikipedia/commons/9/95/Pikendusjuhe.jpg",
	// 	},
	// }

	// if err := conn.Create(&appliances).Error; err != nil {
	// 	panic("failed to create default products")
	// }

	userRepo := repository.NewUserRepo(conn)
	sessionRepo := repository.NewSessionRepo(conn)
	applianceRepo := repository.NewApplianceRepo(conn)
	userApplianceRepo := repository.NewUserApplianceRepo(conn)
	consumptionRepo := repository.NewConsumptionRepo(conn)

	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	applianceService := service.NewApplianceService(applianceRepo)
	userApplianceService := service.NewUserApplianceService(userApplianceRepo)
	consumptionService := service.NewConsumptionService(consumptionRepo)

	mainAPI := api.NewAPI(userService, sessionService, applianceService, userApplianceService, consumptionService)
	mainAPI.Start()
}

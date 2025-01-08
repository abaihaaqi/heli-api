package main

import (
	"log"
	"os"

	"github.com/ijaybaihaqi/heli-api/api"
	"github.com/ijaybaihaqi/heli-api/db"
	"github.com/ijaybaihaqi/heli-api/model"
	"github.com/ijaybaihaqi/heli-api/repository"
	"github.com/ijaybaihaqi/heli-api/service"
)

func main() {
	// Load the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Retrieve the Hugging Face token from the environment variables
	token, ok := os.LookupEnv("HUGGINGFACE_TOKEN")
	if !ok {
		log.Fatal("HUGGINGFACE_TOKEN is must be set")
	}

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
	// 		Name:   "Refrigerator",
	// 		Image:  "https://upload.wikimedia.org/wikipedia/commons/2/2e/US_Domestic_Refrigerator_-_Frigidaire.jpg",
	// 		Energy: 1.2,
	// 	},
	// 	{
	// 		Name:   "TV",
	// 		Image:  "https://upload.wikimedia.org/wikipedia/commons/7/78/Early_portable_tv.jpg",
	// 		Energy: 0.8,
	// 	},
	// 	{
	// 		Name:   "EVCar",
	// 		Image:  "https://upload.wikimedia.org/wikipedia/commons/8/8e/Ford_Explorer_EV_Auto_Zuerich_2023_1X7A1325.jpg",
	// 		Energy: 99.9,
	// 	},
	// 	{
	// 		Name:   "Lamp",
	// 		Image:  "https://upload.wikimedia.org/wikipedia/commons/2/2f/Lamp_with_a_lampshade_illuminated_by_sunlight.jpg",
	// 		Energy: 0.5,
	// 	},
	// 	{
	// 		Name:   "Air Conditioner",
	// 		Image:  "https://upload.wikimedia.org/wikipedia/commons/7/73/Room_air_conditioning_unit_above_a_green_curtain.jpg",
	// 		Energy: 2.1,
	// 	},
	// 	{
	// 		Name:   "Power Strip",
	// 		Image:  "https://upload.wikimedia.org/wikipedia/commons/9/95/Pikendusjuhe.jpg",
	// 		Energy: 10,
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
	chatService := service.NewChatService(token)

	mainAPI := api.NewAPI(userService, sessionService, applianceService, userApplianceService, consumptionService, chatService)
	mainAPI.Start()
}

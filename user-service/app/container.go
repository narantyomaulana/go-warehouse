package app

import (
	"log"
	"micro-warehouse/user-service/configs"
	"micro-warehouse/user-service/controller"
	"micro-warehouse/user-service/database"
	"micro-warehouse/user-service/repository"
	"micro-warehouse/user-service/usecase"
)

type Container struct {
	RoleController controller.RoleControllerInterface
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.ConnectionPostgres(*config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	roleRepo := repository.NewRoleRepository(db.DB)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	roleController := controller.NewRoleController(roleUsecase)

	return &Container{
		RoleController: roleController,
	}
}

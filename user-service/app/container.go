package app

import (
	"log"
	"micro-warehouse/user-service/configs"
	"micro-warehouse/user-service/controller"
	"micro-warehouse/user-service/database"
	"micro-warehouse/user-service/repository"
	"micro-warehouse/user-service/service"
	"micro-warehouse/user-service/usecase"
)

type Container struct {
	RoleController controller.RoleControllerInterface
	UserController controller.UserControllerInterface
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.ConnectionPostgres(*config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	rabbitMQService, err := service.NewRabbitMQService(*config)
	if err != nil {
		log.Fatalf("Failed to connect to rabbitmq: %v", err)
	}

	roleRepo := repository.NewRoleRepository(db.DB)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	roleController := controller.NewRoleController(roleUsecase)

	userRepo := repository.NewUserRepository(db.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, rabbitMQService)
	userController := controller.NewUserController(userUsecase)

	return &Container{
		RoleController: roleController,
		UserController: userController,
	}
}

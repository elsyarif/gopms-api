package infrastructure

import (
	"github.com/elsyarif/pms-api/internal/applications/usecases"
	"github.com/elsyarif/pms-api/internal/domain/services"
	"github.com/elsyarif/pms-api/internal/infrastructure/persistence/postgresql/repositories"
	"github.com/elsyarif/pms-api/internal/interface/http/handler"
	"github.com/elsyarif/pms-api/pkg/encryption"
	"github.com/elsyarif/pms-api/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Container(db *sqlx.DB, app *gin.Engine) {
	idGenerator := uid.NewNanoId()
	hash := encryption.PasswordHash()

	// repository
	userRepository := repositories.NewUserRepositoryPostgres(db)
	authRepository := repositories.NewAuthRepositoryPostgres(db)
	groupRepository := repositories.NewGroupRepositoryPostgres(db)
	serverRepository := repositories.NewServerRepositoryPostgres(db)
	diskRepository := repositories.NewDiskRepositoryPostgres(db)
	inspectionRepository := repositories.NewInspectionRepositoryPostgres(db)

	// services
	userService := services.NewUserService(userRepository, idGenerator, hash)
	authService := services.NewAuthService(authRepository, userRepository, hash)
	groupService := services.NewGroupService(groupRepository, idGenerator)
	serverService := services.NewServerService(serverRepository, idGenerator)
	diskService := services.NewDiskService(diskRepository, idGenerator)
	inspectionService := services.NewInspectionService(inspectionRepository, idGenerator)

	// useCase
	userUseCase := usecases.NewUserUseCase(userService)
	authUseCae := usecases.NewAuthUseCase(authService, userService)
	groupUseCase := usecases.NewGroupUseCase(groupService)
	serverUseCase := usecases.NewServerUseCase(serverService, diskService, groupService)
	diskUseCase := usecases.NewDiskUseCase(diskService)
	inspectionUseCase := usecases.NewInspectionUseCase(inspectionService)
	reportUseCase := usecases.NewReportUseCase(inspectionService)

	// handler
	userHandler := handler.NewUserHandler(userUseCase)
	authHandler := handler.NewAuthHandler(authUseCae)
	groupHandler := handler.NewGroupHandler(groupUseCase, serverUseCase)
	serverHandler := handler.NewServerHandler(serverUseCase)
	diskHandler := handler.NewDiskHandler(diskUseCase)
	inspectionHandler := handler.NewInspectionHandler(inspectionUseCase)
	reportHandler := handler.NewReportHandler(reportUseCase)

	// routes
	userHandler.Routes(app)
	authHandler.Routes(app)
	groupHandler.Routes(app)
	serverHandler.Routes(app)
	diskHandler.Routes(app)
	inspectionHandler.Routes(app)
	reportHandler.Routes(app)
}

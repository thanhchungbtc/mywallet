package app

import (
	"net/http"

	"github.com/thanhchungbtc/mywallet/internal/app/handler"
	"github.com/thanhchungbtc/mywallet/internal/database"
	"github.com/thanhchungbtc/mywallet/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
)

type App struct {
	db     *database.DB
	router *gin.Engine
}

func New() (*App, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	a := &App{
		db: db,
	}
	a.setupRouter()

	a.Migrate()
	return a, nil
}

func (a *App) setupRouter() {
	router := gin.Default()
	a.router = router

	// setup router
	a.MountAdmin()

	api := router.Group("/api")
	handler.NewAccountHandler(a.db).RegisterRoutes(api.Group("/accounts"))
	handler.NewAuthHandler(a.db).RegisterRoutes(api.Group("/auth"))

}

func (a *App) Migrate() {
	a.db.AutoMigrate(model.Account{}, model.Category{}, model.Expense{})
}

func (a *App) MountAdmin() {
	Admin := admin.New(&admin.AdminConfig{DB: a.db.DB})
	Admin.AddResource(&model.Account{})
	Admin.AddResource(&model.Category{})
	Admin.AddResource(&model.Expense{})
	adminMux := http.NewServeMux()
	Admin.MountTo("/admin", adminMux)
	a.router.Any("/admin/*resources", gin.WrapH(adminMux))
}

func (a *App) Run(addr string) error {
	return a.router.Run(addr)
}

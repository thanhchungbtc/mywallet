package app

import (
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/thanhchungbtc/mywallet/server/app/api"
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/model"

	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
)

type App struct {
	db     *database.DB
	router *gin.Engine
}

func New(DB *database.DB) (*App, error) {

	a := &App{
		db: DB,
	}
	a.setupRouter()

	a.Migrate()
	return a, nil
}

func (a *App) setupRouter() {
	router := gin.Default()
	a.router = router

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		AllowMethods:    []string{"*"},
	}))

	// setup router
	a.MountAdmin()

	api.New(a.db).RegisterRoutes(router.Group("/api"))
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

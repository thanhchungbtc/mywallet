package app

import (
	"net/http"

	"github.com/thanhchungbtc/mywallet/internal/service"

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

type Router interface {
	RegisterRoutes(r gin.IRouter)
}

func (a *App) setupRouter() {
	router := gin.Default()
	a.router = router

	// setup router
	a.MountAdmin()

	s := service.New(a.db)
	apiRouters := []struct {
		router Router
		path   string
	}{
		{router: handler.NewAuthHandler(a.db), path: "/auth"},
		{router: handler.NewAccountHandler(a.db), path: "/accounts"},
		{router: handler.NewCategoryHandler(s), path: "/categories"},
	}

	api := router.Group("/api")
	api.Use(withDB(a.db))
	for _, router := range apiRouters {
		router.router.RegisterRoutes(api.Group(router.path))
	}
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

func withDB(db *database.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("MW_DB", db)
	}
}

func (a *App) Run(addr string) error {
	return a.router.Run(addr)
}

package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"neural_storage/cache/adapters/hotstorage"
	"neural_storage/config/core/services/config"
	"neural_storage/cube/core/roles"
	"neural_storage/cube/core/services/interactor/model"
	"neural_storage/cube/core/services/interactor/user"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/cube/handlers/http/v1/admin/adminmodels"
	"neural_storage/cube/handlers/http/v1/admin/adminusers"
	"neural_storage/cube/handlers/http/v1/admin/adminweights"
	"neural_storage/cube/handlers/http/v1/auth"
	"neural_storage/cube/handlers/http/v1/auth/registration"
	"neural_storage/cube/handlers/http/v1/common"
	"neural_storage/cube/handlers/http/v1/models"
	"neural_storage/cube/handlers/http/v1/models/modelsstat"
	"neural_storage/cube/handlers/http/v1/users"
	"neural_storage/cube/handlers/http/v1/users/userblock"
	"neural_storage/cube/handlers/http/v1/users/usersstat"
	"neural_storage/cube/handlers/http/v1/weights"
	"neural_storage/cube/handlers/http/v1/weights/weightsstat"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat/statmiddleware"

	_ "neural_storage/cube/docs"
)

// @title           Cube API
// @version         1.0
// @description     This is cube server.

// @license.name  MIT
// @license.url   https://mit-license.org/

// @host      localhost
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
type Server interface {
	Run(addr ...string) (err error)
}

func New(params config.Config, lg *logger.Logger) Server {
	engine := gin.New()

	engine.Use(gin.Recovery())

	engine.Use(logger.RequestIDSetter())
	engine.Use(logger.RequestLogger(lg))
	engine.Use(statmiddleware.MeasureResponseDuration())

	userInteractor := user.NewInteractor(lg, params.UserInfo())
	adminUserInteractor := user.NewInteractor(lg, params.AdminUserInfo())

	adminModelInterator := model.NewInteractor(lg, params.AdminModelInfo())

	statUserInterator := user.NewInteractor(lg, params.StatUserInfo())
	statModelInterator := model.NewInteractor(lg, params.StatModelInfo())

	initBasicRoutes(params, engine)

	initAuthRoutes(params, lg, engine, userInteractor)

	initModelsRoutes(params, lg, engine, userInteractor, adminModelInterator)

	initWeightsRoutes(params, lg, engine, userInteractor, adminModelInterator)

	initUsersRoutes(params, lg, engine, userInteractor, adminUserInteractor)

	initStatsRoutes(params, lg, engine, userInteractor, statUserInterator, statModelInterator)

	initFailure(lg, engine)

	return engine
}

func initBasicRoutes(params config.Config, engine *gin.Engine) {
	engine.GET("/prometheus", gin.WrapH(promhttp.Handler()))
	engine.GET("/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	if params.Debug() {
		engine.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.PersistAuthorization(true)))
	}
}

func initAuthRoutes(
	params config.Config,
	lg *logger.Logger,
	engine *gin.Engine,
	interactor *user.Interactor,
) jwt.JWTMiddleware {
	authManager := auth.NewHandler(lg, interactor)
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleUser, roles.RoleAdmin, roles.RoleStat}),
		authManager.Authorizator([]uint64{roles.RoleUser, roles.RoleAdmin, roles.RoleStat}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	login := func(c *gin.Context) { authMiddleware.LoginHandler(c) }

	regManager := registration.New(lg, interactor)

	v1 := engine.Group("/api/v1")
	{
		v1.POST("/registration", regManager.Registration)
		v1.POST("/login", login)
	}

	v1Authorized := engine.
		Group("/api/v1").
		Use(authMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		v1Authorized.GET("/refresh", authMiddleware.RefreshHandler)
		v1Authorized.GET("/logout", authMiddleware.LogoutHandler)
	}

	return authMiddleware
}

func initModelsRoutes(
	params config.Config,
	lg *logger.Logger,
	engine *gin.Engine,
	userInteractor *user.Interactor,
	adminInteractor *model.Interactor,
) {
	authManager := auth.NewHandler(lg, userInteractor)
	userAuthMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleUser}),
		authManager.Authorizator([]uint64{roles.RoleUser}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 := engine.
		Group("/api/v1").
		Use(userAuthMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		modelManager := models.New(
			lg,
			model.NewInteractor(lg, params.ModelInfo()),
			hotstorage.New(params.Cache()))

		v1.POST("/models", modelManager.Add)
		v1.PATCH("/models/:model_id", modelManager.Update)
	}

	commonAuthMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleUser, roles.RoleAdmin}),
		authManager.Authorizator([]uint64{roles.RoleUser, roles.RoleAdmin}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 = engine.
		Group("/api/v1").
		Use(commonAuthMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		modelManager := models.New(
			lg,
			model.NewInteractor(lg, params.ModelInfo()),
			hotstorage.New(params.Cache()))

		admModelManager := adminmodels.New(lg, adminInteractor)

		v1.GET("/models", modelManager.GetAll)

		v1.GET("/models/:model_id", modelManager.Get)
		v1.DELETE("/models/:model_id", common.New(admModelManager.Delete, modelManager.Delete).Handle)
	}
}

func initWeightsRoutes(
	params config.Config,
	lg *logger.Logger,
	engine *gin.Engine,
	userInterator *user.Interactor,
	adminInterator *model.Interactor,
) {
	authManager := auth.NewHandler(lg, userInterator)

	userAuthMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleUser}),
		authManager.Authorizator([]uint64{roles.RoleUser}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 := engine.
		Group("/api/v1").
		Use(userAuthMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		weightsManager := weights.New(
			lg,
			model.NewInteractor(lg, params.ModelInfo()),
			hotstorage.New(params.Cache()))

		v1.POST("/weights", weightsManager.Add)
		v1.PATCH("/weights/:weight_id", weightsManager.Update)
	}

	commonAuthMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleUser, roles.RoleAdmin}),
		authManager.Authorizator([]uint64{roles.RoleUser, roles.RoleAdmin}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 = engine.
		Group("/api/v1").
		Use(commonAuthMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		weightsManager := weights.New(
			lg,
			model.NewInteractor(lg, params.ModelInfo()),
			hotstorage.New(params.Cache()))

		admWeightManager := adminweights.New(lg, adminInterator)

		v1.GET("/weights", weightsManager.GetAll)

		v1.GET("/weights/:weight_id", weightsManager.Get)
		v1.DELETE("/weights/:weight_id", common.New(admWeightManager.Delete, weightsManager.Delete).Handle)
	}
}

func initUsersRoutes(
	params config.Config,
	lg *logger.Logger,
	engine *gin.Engine,
	userInteractor *user.Interactor,
	adminUserInteractor *user.Interactor,
) {
	authManager := auth.NewHandler(lg, userInteractor)
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleAdmin}),
		authManager.Authorizator([]uint64{roles.RoleAdmin}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath())

	v1 := engine.
		Group("/api/v1").
		Use(authMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		admUsrManager := adminusers.New(lg, adminUserInteractor)

		v1.DELETE("/users/:user_id", admUsrManager.Delete)

		usrBlockManager := userblock.New(lg, adminUserInteractor)

		v1.GET("/blocks/users/:user_id", usrBlockManager.Get)
		v1.DELETE("/blocks/users/:user_id", usrBlockManager.Delete)
		v1.PATCH("/blocks/users/:user_id", usrBlockManager.Update)
	}

	commonAuthMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleUser, roles.RoleAdmin}),
		authManager.Authorizator([]uint64{roles.RoleUser, roles.RoleAdmin}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 = engine.
		Group("/api/v1").
		Use(commonAuthMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		userManager := users.New(lg, userInteractor)

		v1.GET("/users", userManager.GetAll)
		v1.GET("/users/:user_id", userManager.Get)
	}
}

func initStatsRoutes(
	params config.Config,
	lg *logger.Logger,
	engine *gin.Engine,
	user *user.Interactor,
	statUserInteractor *user.Interactor,
	statModelInteractor *model.Interactor,
) {
	authManager := auth.NewHandler(lg, user)
	authMiddleware := jwt.NewJWTMiddleware(
		authManager.Authenticator([]uint64{roles.RoleStat}),
		authManager.Authorizator([]uint64{roles.RoleStat}),
		authManager.Payload,
		authManager.IdentityHandler,
		params.PrivKeyPath(),
		params.PubKeyPath(),
	)

	v1 := engine.
		Group("/api/v1").
		Use(authMiddleware.MiddlewareFunc()).
		Use(gzip.Gzip(gzip.BestCompression))
	{
		userManager := usersstat.New(lg, statUserInteractor)
		v1.GET("/users/stats", userManager.Get)

		modelManager := modelsstat.New(lg, statModelInteractor)
		v1.GET("/models/stats", modelManager.Get)

		weightsManager := weightsstat.New(lg, statModelInteractor)
		v1.GET("/weights/stats", weightsManager.Get)
	}
}

func initFailure(lg *logger.Logger, engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		lg.WithFields(logrus.Fields{"claims": claims}).Info("no route")
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}

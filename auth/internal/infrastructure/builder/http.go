// infrastructure/builder/http.go
package builder

import (
	httpcontroller "clirzy/auth/internal/infrastructure/adapter/controller/http"
	"clirzy/auth/internal/infrastructure/config"
	"clirzy/pkg/httpserver"

	"github.com/gin-gonic/gin"
)

func BuildHTTP(cfg *config.Config, uc *UseCases) *httpserver.Server {
	r := gin.Default()
	r.Use(gin.Recovery())

	handler := httpcontroller.NewAuthHandler(
		uc.Register,
		uc.Login,
		uc.Logout,
		uc.RefreshToken,
	)

	httpcontroller.RegisterRoutes(r, handler)

	return httpserver.NewServer(":"+cfg.HttpServer.Port, r)
}

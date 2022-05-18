package endpoints

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerApi struct {
	Server *server.Server
}

func (s *ServerApi) GetMaxPlayerCount(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, s.Server.MaxPlayerCount())
}

func (s *ServerApi) GetPlayerCount(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, s.Server.PlayerCount())
}

func (s *ServerApi) CloseServer(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, s.Server.Close())
}

func (s *ServerApi) OpenServer(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, s.Server.Start())
}

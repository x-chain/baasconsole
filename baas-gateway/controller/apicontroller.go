package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/x-chain/baasconsole/baas-gateway/service"
	"net/http"
	"time"
)

type ApiController struct {
	chainService     *service.ChainService
	channelService   *service.ChannelService
	chaincodeService *service.ChaincodeService
	dashboardService *service.DashboardService
	userService      *service.UserService
	roleService      *service.RoleService
}

func NewApiController(userService *service.UserService, roleService *service.RoleService, chainService *service.ChainService, channelService *service.ChannelService, chaincodeService *service.ChaincodeService, dashboardService *service.DashboardService) *ApiController {
	return &ApiController{
		userService:      userService,
		roleService:      roleService,
		chainService:     chainService,
		channelService:   channelService,
		chaincodeService: chaincodeService,
		dashboardService: dashboardService,
	}
}

func (a *ApiController) Upload(ctx *gin.Context) {
	// single file
	file, _ := ctx.FormFile("file")
	path := fmt.Sprintf("/tmp/%d", time.Now().UnixNano())
	ctx.SaveUploadedFile(file, path)
	ctx.String(http.StatusOK, path)

}

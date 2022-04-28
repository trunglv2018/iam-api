package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/trunglen/g/x/rest"
)

type AdminApi struct {
	rest.JsonRender
	*gin.RouterGroup
}

func NewAdminApi(routerGroup *gin.RouterGroup) *AdminApi {
	s := &AdminApi{
		RouterGroup: routerGroup,
	}
	s.GET("customer/list", s.getListCustomers)
	//customer api

	return s
}

func (s *AdminApi) getListCustomers(c *gin.Context) {
	s.SendData(c, []map[string]interface{}{
		{
			"name": "Trung 1",
		},
		{
			"name": "Trung 2",
		},
	})
}

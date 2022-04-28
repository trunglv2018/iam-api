package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/trunglen/g/x/rest"
)

type CustomerApi struct {
	rest.JsonRender
	*gin.RouterGroup
}

func NewCustomerApi(routerGroup *gin.RouterGroup) *CustomerApi {
	s := &CustomerApi{
		RouterGroup: routerGroup,
	}
	s.GET("product/list", s.getListProducts)
	//customer api

	return s
}

func (s *CustomerApi) getListProducts(c *gin.Context) {
	s.SendData(c, []map[string]interface{}{
		{
			"name": "Product 1",
		},
		{
			"name": "Product 2",
		},
	})
}

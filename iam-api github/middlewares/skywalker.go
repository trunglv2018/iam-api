package middlewares

import (
	"fmt"
	"strconv"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/gin-gonic/gin"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const componentIDGINHttpServer = 5006

//Middleware gin middleware return HandlerFunc  with tracing.
func SkywalkerMiddleware(engine *gin.Engine, tracer *go2sky.Tracer) gin.HandlerFunc {
	if engine == nil || tracer == nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		span, ctx, err := tracer.CreateEntrySpan(c.Request.Context(), getOperationName(c), func(key string) (string, error) {
			return c.Request.Header.Get(key), nil
		})
		if err != nil {
			c.Next()
			return
		}
		span.SetComponent(componentIDGINHttpServer)
		span.Tag(go2sky.TagHTTPMethod, c.Request.Method)
		span.Tag(go2sky.TagURL, c.Request.Host+c.Request.URL.Path)
		span.SetSpanLayer(agentv3.SpanLayer_Http)

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		if len(c.Errors) > 0 {
			span.Error(time.Now(), c.Errors.String())
		}
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(c.Writer.Status()))
		span.End()
	}
}

func getOperationName(c *gin.Context) string {
	return fmt.Sprintf("/%s%s", c.Request.Method, c.FullPath())
}

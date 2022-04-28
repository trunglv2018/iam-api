package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"gitlab.com/trunglen/iam-api/config"
)

func Proxy(c *gin.Context) {
	var remoteURL = config.GetConfig().GetString("reverse_proxy.remote_url")
	remote, err := url.Parse(remoteURL)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", remote.Host)
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

package workspace

import "github.com/gogf/gf/v2/net/ghttp"

func Deactivated(r *ghttp.Request) {
	r.Response.WriteTpl("deactivated.html")
}
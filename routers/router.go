// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bee/blog/controllers"
	"bee/blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// 过滤器
	beego.InsertFilter("/api/v1/comment/add", beego.BeforeRouter, utils.GetRequestTokenFilter)
	beego.InsertFilter("/api/v1/sub_comment/add", beego.BeforeRouter, utils.GetRequestTokenFilter)
	beego.InsertFilter("/api/v1/article/add", beego.BeforeRouter, utils.GetRequestTokenFilter)

	ns := beego.NewNamespace("/api/v1",

		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
		beego.NSNamespace("/sub_comment",
			beego.NSInclude(
				&controllers.SubCommentController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
		beego.NSNamespace("/sub_category",
			beego.NSInclude(
				&controllers.SubCategoryController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
				&controllers.ProfileController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
}

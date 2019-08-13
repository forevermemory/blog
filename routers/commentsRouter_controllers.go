package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "DeleteById",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetList",
            Router: `/getall`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetTotalCount",
            Router: `/getcount`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AcceptImage",
            Router: `/image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/getall`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CommentController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:CommentController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:CommentController"],
        beego.ControllerComments{
            Method: "DeleteComment",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "DeleteByContent",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetMessageCount",
            Router: `/getcount`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/profile/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/profile/get/:userId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/profile/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:ProfileController"],
        beego.ControllerComments{
            Method: "UpdateAvator",
            Router: `/profile/update/avator`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/getall`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCategoryController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCommentController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCommentController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:SubCommentController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:SubCommentController"],
        beego.ControllerComments{
            Method: "DeleteSubComment",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:TagController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:TagController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:TagController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetAllArticle",
            Router: `/getArticle`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:TagController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/getall`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "AllUser",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: `/getList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "CheckUsernameExist",
            Router: `/registe`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/registe`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["bee/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePassword",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

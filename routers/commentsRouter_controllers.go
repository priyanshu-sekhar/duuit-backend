package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["duuit-backend/controllers:GoalController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:GoalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:GoalController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:GoalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:GoalController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:GoalController"],
        beego.ControllerComments{
            Method: "AddGoalTracking",
            Router: "/:id/tracking",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:RootController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:RootController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["duuit-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["duuit-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Ping",
            Router: "/ping",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

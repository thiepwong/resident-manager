package routes

import (
	"os"

	"github.com/thiepwong/resident-manager/middlewares"
	"github.com/thiepwong/resident-manager/services"
	"github.com/thiepwong/smartid/pkg/logger"

	"github.com/thiepwong/resident-manager/repositories"

	"github.com/thiepwong/resident-manager/common"
	"github.com/thiepwong/resident-manager/datasources"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/controllers"
)

type SetHeader func(iris.Context)

func RegisterRoute(app *iris.Application, cors context.Handler, config *common.Config) {
	pg, err := datasources.GetPg(*config.Database.Postgre)
	if err != nil {
		logger.LogErr.Println(err.Error())
		os.Exit(2)
	}

	mvcResult := controllers.NewMvcResult(nil)

	//Register Employee Controller
	empRep := repositories.NewEmployeeRepository(pg)
	empSrv := services.NewEmployeeService(empRep)
	emp := mvc.New(app.Party("/employees", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	emp.Register(empSrv, mvcResult)
	emp.Handle(new(controllers.EmployeeController))

	//Register Department Controller
	deptRep := repositories.NewDepartmentRepository(pg)
	deptSrv := services.NewDepartmentService(deptRep)
	dept := mvc.New(app.Party("/department", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	dept.Register(deptSrv, mvcResult)
	dept.Handle(new(controllers.DepartmentController))

	//Register Notification Controller
	notiRep := repositories.NewNotificationRepository(pg)
	notiSrv := services.NewNotificationService(notiRep)
	noti := mvc.New(app.Party("/notification", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	noti.Register(notiSrv, mvcResult)
	noti.Handle(new(controllers.NotificationController))

	//Register Side Controller
	sideRep := repositories.NewSideRepository(pg)
	sideSrv := services.NewSideService(sideRep)
	side := mvc.New(app.Party("/side", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	side.Register(sideSrv, mvcResult)
	side.Handle(new(controllers.SideController))

}

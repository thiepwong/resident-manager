package routes

import (
	"os"

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
	emp := mvc.New(app.Party("/employees", cors).AllowMethods(iris.MethodOptions))
	emp.Register(empSrv, mvcResult)
	emp.Handle(new(controllers.EmployeeController))

}

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
	empSrv := services.NewEmployeeService(empRep, config)
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
	notiSrv := services.NewNotificationService(notiRep, config)
	noti := mvc.New(app.Party("/notification", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	noti.Register(notiSrv, mvcResult)
	noti.Handle(new(controllers.NotificationController))

	//Register Side Controller
	sideRep := repositories.NewSideRepository(pg)
	sideSrv := services.NewSideService(sideRep)
	side := mvc.New(app.Party("/side", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	side.Register(sideSrv, mvcResult)
	side.Handle(new(controllers.SideController))

	//Register Block Controller
	blockRep := repositories.NewBlockRepository(pg)
	blockSrv := services.NewBlockService(blockRep)
	block := mvc.New(app.Party("/block", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	block.Register(blockSrv, mvcResult)
	block.Handle(new(controllers.BlockController))

	//Register Area Controller
	areaRep := repositories.NewAreaRepository(pg)
	areaSrv := services.NewAreaService(areaRep)
	area := mvc.New(app.Party("/area", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	area.Register(areaSrv, mvcResult)
	area.Handle(new(controllers.AreaController))

	//Register Contact Controller
	contactRep := repositories.NewContactRepository(pg)
	contactSrv := services.NewContactService(contactRep)
	contact := mvc.New(app.Party("/contact", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	contact.Register(contactSrv, mvcResult)
	contact.Handle(new(controllers.ContactController))

	//Register Room Controller
	roomRep := repositories.NewRoomRepository(pg)
	roomSrv := services.NewRoomService(roomRep)
	room := mvc.New(app.Party("/room", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	room.Register(roomSrv, mvcResult)
	room.Handle(new(controllers.RoomController))

	//Register Feedback Controller
	feedbackrep := repositories.NewFeedbackRepository(pg)
	feedbackSrv := services.NewFeedbackService(feedbackrep)
	feedback := mvc.New(app.Party("/feedback", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	feedback.Register(feedbackSrv, mvcResult)
	feedback.Handle(new(controllers.FeedbackController))

	//Register Resident Controller
	residentRepo := repositories.NewResidentRepository(pg)
	residentSrv := services.NewResidentService(residentRepo)
	resident := mvc.New(app.Party("/resident", cors, middlewares.Authorization).AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost))
	resident.Register(residentSrv, mvcResult)
	resident.Handle(new(controllers.ResidentController))

}

package Routes

import (
	"Snap/Cards"
	"Snap/Controllers"
	"Snap/Dashboard"
	"Snap/FormBackend"
	"Snap/ModifyTables"
	"Snap/SendData"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func HTTPHandler(f http.HandlerFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		Controllers.User(c)
		h := http.HandlerFunc(f)
		handler := fasthttpadaptor.NewFastHTTPHandler(h)
		handler(c.Context())
		return nil
	}
}

func Setup(app *fiber.App) {
	app.Use(logger.New())
	app.Post("/api/Login", Controllers.Login)
	app.Use("/api/User", Controllers.User)
	app.Use("/api/GetViolations", Controllers.GetViolations)
	app.Use("/AddClass", HTTPHandler(FormBackend.AddClass))
	app.Use("/AddStudent", HTTPHandler(FormBackend.AddStudent))
	app.Use("/AddViolation", HTTPHandler(FormBackend.AddViolation))
	app.Use("/AddTeacher", HTTPHandler(FormBackend.AddTeacher))
	//app.Use("/AssignCapstone", HTTPHandler(FormBackend.RecordStdCapstone))
	app.Use("/AssignStudentClass", HTTPHandler(FormBackend.AssignStudentClass))
	app.Use("/AddVacation", HTTPHandler(FormBackend.AddVacation))
	app.Use("/AssignBussing", HTTPHandler(FormBackend.Bussing))
	app.Use("/AddBus", HTTPHandler(FormBackend.AddBus))
	app.Use("/RecordTeacherVacation", HTTPHandler(FormBackend.RecordVacation))
	app.Use("/AddTrip", HTTPHandler(FormBackend.AddTrip))
	app.Use("/SuperVisor", HTTPHandler(FormBackend.Supervisor))
	app.Use("/RecordStudentUniversity", HTTPHandler(FormBackend.RecordStudentUniversity))
	app.Use("/RecordStudentViolation", HTTPHandler(FormBackend.RecordStudentViolation))
	app.Use("/RecordTeacherViolation", HTTPHandler(FormBackend.RecordTeacherViolation))
	app.Use("/AssignRooming", HTTPHandler(FormBackend.Rooming))
	app.Use("/AddRoom", HTTPHandler(FormBackend.AddRoom))
	app.Use("/AddScholarShip", HTTPHandler(FormBackend.AddScholarShip))
	app.Use("/AddSchool", HTTPHandler(FormBackend.AddSchool))
	app.Use("/AddUniversity", HTTPHandler(FormBackend.AddUniversity))
	//app.Use("/EnterComp", HTTPHandler(FormBackend.Entercomp))
	app.Use("/RecordStudentScholarShip", HTTPHandler(FormBackend.AssignStudentScholarShip))
	app.Use("/Dashboard", HTTPHandler(Dashboard.Dashboard))
	app.Use("/AllStudents", HTTPHandler(SendData.ShowAllStudents))
	app.Use("/AllTeachers", HTTPHandler(SendData.ShowAllTeachers))
	app.Use("/AllVacations", HTTPHandler(SendData.ShowAllVacations))
	app.Use("/AllViolations", HTTPHandler(SendData.ShowAllViolations))
	app.Use("/AllUniversities", HTTPHandler(SendData.ShowAllUniversities))
	app.Use("/AllScholarShips", HTTPHandler(SendData.ShowAllScholarShips))
	app.Use("/TeacherCardView", HTTPHandler(Cards.TeacherCardView))
	app.Use("/StudentCardView", HTTPHandler(Cards.StudentCardView))
	app.Use("/ClassCardView", HTTPHandler(Cards.ClassCardView))
	app.Use("/DeleteFromTable", HTTPHandler(FormBackend.DeleteFromAnyTable))
	app.Use("/Grade/10", HTTPHandler(Cards.Grade10StudentCardView))
	app.Use("/Grade/11", HTTPHandler(Cards.Grade11StudentCardView))
	app.Use("/Grade/12", HTTPHandler(Cards.Grade12StudentCardView))
	app.Post("/api/RemoveData", ModifyTables.RemoveData)
	app.Post("/api/EditData", ModifyTables.EditData)
	app.Post("/api/WriteData", Controllers.WriteDataFromApp)
	//app.Static("/Static", "./Static")
	Cards.GetApp(*app)
	Cards.TeacherCardDetails()
	Cards.StudentCardDetails()
	Cards.ClassCardDetails()
	app.Get("/", HTTPHandler(Controllers.RedirectToDashboard))
	app.Get("/api/Logout", Controllers.Logout)
}

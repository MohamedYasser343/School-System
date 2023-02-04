package Cards

import (
	"Snap/Controllers"
	"Snap/Database"
	"Snap/FormBackend"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var (
	tmplTeacherView = template.Must(template.ParseFiles("./DashboardFiles/TeacherCardView.html"))
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

var app *fiber.App

func GetApp(passedApp fiber.App) {
	app = &passedApp
}

func TeacherCardDetails() {
	db := Database.DBConnect()
	app.Get("/Teachers/:TeacherNationalId", func(c *fiber.Ctx) error {
		Controllers.User(c)
		if Controllers.CurrentUser.Permission > 0 {
			// Get the id from the url
			id := c.Params("TeacherNationalId")
			// Get the user from the database
			var teacherData struct {
				TeacherId                   string
				TeacherCode                 string
				TeacherName                 string
				JobCadr                     string
				SupervisingPosOriginSchool  string
				SupervisingPosStemSchool    string
				EduLevel                    string
				HiringDate                  string
				OriginQualification         string
				ObtainedOriginQualification string
				QualificationDate           string
				Specialization              string
				OriginQualificationGrade    string
				HighQualification           string
				ObtainedHighQualification   string
				HighQualificationDate       string
				HighQualificationGrade      string
				FinancialGrade              string
				FinancialGradeDate          string
				OriginSchool                string
				OriginTeachingGov           string
				OriginCity                  string
				StemSchoolName              string
				CurrentStemSchoolEnterDate  string
				StemSchoolEnterDate         string
				MobileWhats                 string
				MobileCalls                 string
				Email                       string
				NationalId                  string
				Photo                       string
			}
			// Get the user from the database
			PopulateDataQuery, err := db.Query("SELECT * FROM `teachers` WHERE `NationalId` = ?", id)
			if err != nil {
				log.Println(err.Error())
			}
			for PopulateDataQuery.Next() {
				err = PopulateDataQuery.Scan(&teacherData.TeacherId, &teacherData.TeacherCode, &teacherData.TeacherName, &teacherData.JobCadr, &teacherData.SupervisingPosOriginSchool, &teacherData.SupervisingPosStemSchool, &teacherData.EduLevel, &teacherData.HiringDate, &teacherData.OriginQualification, &teacherData.ObtainedOriginQualification, &teacherData.QualificationDate, &teacherData.Specialization, &teacherData.OriginQualificationGrade, &teacherData.HighQualification, &teacherData.ObtainedHighQualification, &teacherData.HighQualificationDate, &teacherData.HighQualificationGrade, &teacherData.FinancialGrade, &teacherData.FinancialGradeDate, &teacherData.OriginSchool, &teacherData.OriginTeachingGov, &teacherData.OriginCity, &teacherData.StemSchoolName, &teacherData.CurrentStemSchoolEnterDate, &teacherData.StemSchoolEnterDate, &teacherData.MobileWhats, &teacherData.MobileCalls, &teacherData.Email, &teacherData.NationalId, &teacherData.Photo)
				if err != nil {
					log.Println(err.Error())
				}
			}
			defer PopulateDataQuery.Close()
			fmt.Println(teacherData.Photo)
			if teacherData.Photo == "" || teacherData.Photo == " " || teacherData.Photo == "0" {
				chars := []rune(teacherData.NationalId)
				if chars[12]%2 == 1 {
					teacherData.Photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
				} else {
					teacherData.Photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
				}
			}
			if teacherData.TeacherId == "" || teacherData.TeacherId == " " {
				return c.JSON(fiber.Map{
					"Message": "No Teacher Found.",
				})
			}
			// Send the user to the client
			return c.Render("TeacherCardDetails", fiber.Map{
				"TeacherName":                        teacherData.TeacherName,
				"TeacherSubject":                     teacherData.SupervisingPosOriginSchool,
				"TeacherCode":                        teacherData.TeacherCode,
				"TeacherPhoneWhatsapp":               teacherData.MobileWhats,
				"TeacherPhoneCalls":                  teacherData.MobileCalls,
				"TeacherCadrJob":                     teacherData.JobCadr,
				"TeacherPreviousJob":                 teacherData.SupervisingPosOriginSchool,
				"TeacherEducationStage":              teacherData.EduLevel,
				"TeacherHiringDate":                  teacherData.HiringDate,
				"TeacherOriginQualification":         teacherData.OriginQualification,
				"TeacherObtainedOriginQualification": teacherData.ObtainedOriginQualification,
				"TeacherOriginalQualificationDate":   teacherData.QualificationDate,
				"TeacherSpecialization":              teacherData.Specialization,
				"TeacherOriginQualificationGrade":    teacherData.OriginQualificationGrade,
				"TeacherHighQualification":           teacherData.HighQualification,
				"TeacherObtainedHighQualification":   teacherData.ObtainedHighQualification,
				"TeacherHighQualificationDate":       teacherData.HighQualificationDate,
				"TeacherHighQualificationGrade":      teacherData.HighQualificationGrade,
				"TeacherFinancialGrade":              teacherData.FinancialGrade,
				"TeacherFinancialGradeDate":          teacherData.FinancialGradeDate,
				"TeacherOriginalSchool":              teacherData.OriginSchool,
				"TeacherOriginalTeachingGovernment":  teacherData.OriginTeachingGov,
				"TeacherOriginalCity":                teacherData.OriginCity,
				"TeacherStemSchoolNames":             teacherData.StemSchoolName,
				"TeacherCurrentStemEntryDate":        teacherData.CurrentStemSchoolEnterDate,
				"TeacherStemEntryDate":               teacherData.StemSchoolEnterDate,
				"TeacherEmail":                       teacherData.Email,
				"TeacherNationalId":                  teacherData.NationalId,
				"Photo":                              teacherData.Photo,
			})
		} else if Controllers.CurrentUser.Permission == 0 {
			// Print No Permission
			return c.Render("NotLoggedIn", nil)
			// return FormBackend.TmplNotLoggedIn.Execute(c, nil)
		} else {
			return c.Render("NotEnoughPermission", nil)
			// return FormBackend.TmplNoPermission.Execute(c, nil)
		}
	})
}

func TeacherCardView(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		db := Database.DBConnect()
		var teacherCards []string
		TeacherNameList, err := db.Query("SELECT `TeacherName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var i int = 0
		var teacherCount int = 0
		for TeacherNameList.Next() {
			teacherCount++
		}
		TeacherCards, err := db.Query("SELECT `TeacherName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for TeacherCards.Next() {
			var teacher string
			err = TeacherCards.Scan(&teacher)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherCards = append(teacherCards, teacher+",")
			} else {
				teacherCards = append(teacherCards, teacher)
			}
		}
		defer TeacherCards.Close()
		var teacherSubjects []string
		i = 0
		TeacherSubjects, err := db.Query("SELECT `SupervisingPosStemSchool` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for TeacherSubjects.Next() {
			var teacherSubject string
			err = TeacherSubjects.Scan(&teacherSubject)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherSubjects = append(teacherSubjects, teacherSubject+",")
			} else {
				teacherSubjects = append(teacherSubjects, teacherSubject)
			}
		}
		defer TeacherSubjects.Close()

		var teacherGenders []string
		i = 0
		TeacherNationalIds, err := db.Query("SELECT `NationalId` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var nationalIds []string
		for TeacherNationalIds.Next() {
			var teacher string
			err = TeacherNationalIds.Scan(&teacher)
			if err != nil {
				fmt.Println(err)
			}
			i++
			var photo string
			chars := []rune(teacher)
			if chars[12]%2 == 1 {
				photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
			} else {
				photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
			}
			if i != teacherCount {
				teacherGenders = append(teacherGenders, photo+",")
				nationalIds = append(nationalIds, teacher+",")
			} else {
				teacherGenders = append(teacherGenders, photo)
				nationalIds = append(nationalIds, teacher)
			}
		}
		defer TeacherNationalIds.Close()

		tmplTeacherView.Execute(w, struct {
			Sidebar            template.HTML
			TeacherNationalIds []string
			TeacherList        []string
			TeacherSubjectList []string
			TeacherGenderList  []string
		}{
			Controllers.CurrentUser.SideBar,
			nationalIds,
			teacherCards,
			teacherSubjects,
			teacherGenders,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

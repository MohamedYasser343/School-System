package Cards

import (
	"Snap/Controllers"
	"Snap/Database"
	"Snap/FormBackend"
	"Snap/Sidebars"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	tmplStudentView = template.Must(template.ParseFiles("./DashboardFiles/StudentCardView.html"))
	tmplGradeView   = template.Must(template.ParseFiles("./DashboardFiles/GradeCardView.html"))
)

func StudentCardDetails() {
	db := Database.DBConnect()

	app.Get("/Students/:StudentNationalId", func(c *fiber.Ctx) error {
		Controllers.User(c)
		if Controllers.CurrentUser.Permission >= 6 {
			NationalId := c.Params("StudentNationalId")
			fmt.Println(NationalId)
			// Get the student's data
			StudentData, err := db.Query("SELECT * FROM `student` WHERE `National ID (14 digit)` = ?", NationalId)
			if err != nil {
				log.Println(err.Error())
			}
			var studentData struct {
				StdId          string
				Code           string
				EnglishName    string
				ArabicName     string
				Grade          string
				Specialization string
				SecondLang     string
				Religion       string
				Gender         string
				PortfolioNo    string
				NationalId     string
				Class          string
				PersonalEmail  string
				OfficialEmail  string
			}
			for StudentData.Next() {
				err = StudentData.Scan(&studentData.StdId, &studentData.Code, &studentData.EnglishName, &studentData.ArabicName, &studentData.Grade, &studentData.Specialization, &studentData.SecondLang, &studentData.Religion, &studentData.Gender, &studentData.PortfolioNo, &studentData.NationalId, &studentData.Class, &studentData.PersonalEmail, &studentData.OfficialEmail)
				if err != nil {
					fmt.Println(err)
				}
			}
			if studentData.NationalId == "" {
				return c.JSON(fiber.Map{
					"Message": "No Student Found",
				})
			}
			var photo string

			if studentData.Grade == "G10" {
				studentData.Grade = "الصف الأول"
			} else if studentData.Grade == "G11" {
				studentData.Grade = "الصف الثاني"
			} else if studentData.Grade == "G12" {
				studentData.Grade = "الصف الثالث"
			}

			if studentData.Religion == "Christian" {
				studentData.Religion = "مسيحي"
			} else if studentData.Religion == "Muslem" {
				studentData.Religion = "مسلم"
			}

			if studentData.SecondLang == "German" {
				studentData.SecondLang = "الماني"
			} else if studentData.SecondLang == "French" {
				studentData.SecondLang = "فرنساوي"
			}

			if studentData.Specialization == "Mathematic" {
				studentData.Specialization = "علمي علوم"
			} else if studentData.Specialization == "Science" {
				studentData.Specialization = "علمي رياضة"
			} else if studentData.Specialization == "" {
				studentData.Specialization = "..."
			}

			chars := []rune(studentData.NationalId)
			if chars[12]%2 == 1 {
				photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
				studentData.Gender = "ذكر"
			} else {
				photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
				studentData.Gender = "انثي"
			}
			//Return Student Data
			return c.Render("StudentCardDetails", fiber.Map{
				"StudentName":    studentData.ArabicName,
				"Photo":          photo,
				"Code":           studentData.Code,
				"EnglishName":    studentData.EnglishName,
				"Grade":          studentData.Grade,
				"Gender":         studentData.Gender,
				"NationalId":     studentData.NationalId,
				"SecondLanguage": studentData.SecondLang,
				"Specialization": studentData.Specialization,
				"Religion":       studentData.Religion,
				"CapstoneGroup":  studentData.PortfolioNo,
				"Class":          studentData.Class,
			})
		} else if Controllers.CurrentUser.Permission == 0 {
			// Print No Permission
			return FormBackend.TmplNotLoggedIn.Execute(c, nil)
		} else {
			return FormBackend.TmplNoPermission.Execute(c, nil)
		}
	})
}

func StudentCardView(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()
		var studentCards []string
		StudentNameList, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var i int = 0
		var studentCount int = 0
		for StudentNameList.Next() {
			studentCount++
		}
		StudentCards, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for StudentCards.Next() {
			var student string
			err = StudentCards.Scan(&student)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != studentCount {
				studentCards = append(studentCards, student+",")
			} else {
				studentCards = append(studentCards, student)
			}
		}
		defer StudentCards.Close()

		var studentGenders []string
		var studentLinks []string
		i = 0
		StudentNationalIds, err := db.Query("SELECT `National ID (14 digit)` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for StudentNationalIds.Next() {
			var student string
			err = StudentNationalIds.Scan(&student)
			if err != nil {
				fmt.Println(err)
			}
			i++
			var photo string
			chars := []rune(student)
			if chars[12]%2 == 1 {
				photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
			} else {
				photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
			}
			if i != studentCount {
				studentGenders = append(studentGenders, photo+",")
				studentLinks = append(studentLinks, student+",")
			} else {
				studentGenders = append(studentGenders, photo)
				studentLinks = append(studentLinks, student)
			}
		}
		defer StudentNationalIds.Close()

		tmplStudentView.Execute(w, struct {
			Sidebar           template.HTML
			StudentList       []string
			StudentGenderList []string
			StudentLinks      []string
		}{
			Sidebars.AdminSidebar,
			studentCards,
			studentGenders,
			studentLinks,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func Grade10StudentCardView(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()
		var studentCards []string
		StudentNameList := Database.SelectFromDB("English Name", "student", "Grade", "G10", db)

		var i int = 0
		var studentCount int = 0
		for StudentNameList.Next() {
			studentCount++
		}
		StudentCards := Database.SelectFromDB("English Name", "student", "Grade", "G10", db)

		for StudentCards.Next() {
			var student string
			StudentCards.Scan(&student)

			i++
			if i != studentCount {
				studentCards = append(studentCards, student+",")
			} else {
				studentCards = append(studentCards, student)
			}
		}
		defer StudentCards.Close()

		var studentGenders []string
		var studentLinks []string
		i = 0
		StudentNationalIds := Database.SelectFromDB("National ID (14 digit)", "student", "Grade", "G10", db)
		for StudentNationalIds.Next() {
			var student string
			StudentNationalIds.Scan(&student)
			i++
			var photo string
			chars := []rune(student)
			if chars[12]%2 == 1 {
				photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
			} else {
				photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
			}
			if i != studentCount {
				studentGenders = append(studentGenders, photo+",")
				studentLinks = append(studentLinks, student+",")
			} else {
				studentGenders = append(studentGenders, photo)
				studentLinks = append(studentLinks, student)
			}
		}
		defer StudentNationalIds.Close()

		tmplGradeView.Execute(w, struct {
			Sidebar           template.HTML
			StudentList       []string
			StudentGenderList []string
			StudentLinks      []string
			Grade             string
		}{
			Sidebars.AdminSidebar,
			studentCards,
			studentGenders,
			studentLinks,
			"Grade 10",
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func Grade11StudentCardView(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()
		var studentCards []string
		StudentNameList := Database.SelectFromDB("English Name", "student", "Grade", "G11", db)

		var i int = 0
		var studentCount int = 0
		for StudentNameList.Next() {
			studentCount++
		}
		StudentCards := Database.SelectFromDB("English Name", "student", "Grade", "G11", db)

		for StudentCards.Next() {
			var student string
			StudentCards.Scan(&student)

			i++
			if i != studentCount {
				studentCards = append(studentCards, student+",")
			} else {
				studentCards = append(studentCards, student)
			}
		}
		defer StudentCards.Close()

		var studentGenders []string
		var studentLinks []string
		i = 0
		StudentNationalIds := Database.SelectFromDB("National ID (14 digit)", "student", "Grade", "G11", db)
		for StudentNationalIds.Next() {
			var student string
			StudentNationalIds.Scan(&student)
			i++
			var photo string
			chars := []rune(student)
			if chars[12]%2 == 1 {
				photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
			} else {
				photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
			}
			if i != studentCount {
				studentGenders = append(studentGenders, photo+",")
				studentLinks = append(studentLinks, student+",")
			} else {
				studentGenders = append(studentGenders, photo)
				studentLinks = append(studentLinks, student)
			}
		}
		defer StudentNationalIds.Close()

		tmplGradeView.Execute(w, struct {
			Sidebar           template.HTML
			StudentList       []string
			StudentGenderList []string
			StudentLinks      []string
			Grade             string
		}{
			Sidebars.AdminSidebar,
			studentCards,
			studentGenders,
			studentLinks,
			"Grade 11",
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func Grade12StudentCardView(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()
		var studentCards []string
		StudentNameList := Database.SelectFromDB("English Name", "student", "Grade", "G12", db)

		var i int = 0
		var studentCount int = 0
		for StudentNameList.Next() {
			studentCount++
		}
		StudentCards := Database.SelectFromDB("English Name", "student", "Grade", "G12", db)

		for StudentCards.Next() {
			var student string
			StudentCards.Scan(&student)

			i++
			if i != studentCount {
				studentCards = append(studentCards, student+",")
			} else {
				studentCards = append(studentCards, student)
			}
		}
		defer StudentCards.Close()

		var studentGenders []string
		var studentLinks []string
		i = 0
		StudentNationalIds := Database.SelectFromDB("National ID (14 digit)", "student", "Grade", "G12", db)
		for StudentNationalIds.Next() {
			var student string
			StudentNationalIds.Scan(&student)
			i++
			var photo string
			chars := []rune(student)
			if chars[12]%2 == 1 {
				photo = "https://snap.hopto.org/AboutTeacher/images/avatarMale.jpg"
			} else {
				photo = "https://snap.hopto.org/Dashboard/img_avatar2.png"
			}
			if i != studentCount {
				studentGenders = append(studentGenders, photo+",")
				studentLinks = append(studentLinks, student+",")
			} else {
				studentGenders = append(studentGenders, photo)
				studentLinks = append(studentLinks, student)
			}
		}
		defer StudentNationalIds.Close()

		tmplGradeView.Execute(w, struct {
			Sidebar           template.HTML
			StudentList       []string
			StudentGenderList []string
			StudentLinks      []string
			Grade             string
		}{
			Sidebars.AdminSidebar,
			studentCards,
			studentGenders,
			studentLinks,
			"Grade 12",
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

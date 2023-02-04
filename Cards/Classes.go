package Cards

import (
	"Snap/Controllers"
	"Snap/Database"
	"Snap/FormBackend"
	"Snap/Sidebars"
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var tmplClass = template.Must(template.ParseFiles("./DashboardFiles/ClassCardView.html"))

func ClassCardView(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()
		defer db.Close()

		ClassCount := Database.SelectFromDB("*", "Class", "1", "1", db)
		defer ClassCount.Close()

		var classCount int = 0
		for ClassCount.Next() {
			classCount++
		}
		ClassList := Database.SelectFromDB("ClassName", "Class", "1", "1", db)

		defer ClassList.Close()
		var classList []string
		var i int = 0
		for ClassList.Next() {
			var class string
			ClassList.Scan(&class)
			i++
			if i != classCount {
				classList = append(classList, class+",")
			} else {
				classList = append(classList, class)
			}
		}
		tmplClass.Execute(w, struct {
			Sidebar   template.HTML
			ClassList []string
		}{
			Sidebars.AdminSidebar,
			classList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ClassCardDetails() {
	db := Database.DBConnect()
	var classCount int
	ClassCount := Database.SelectFromDB("*", "Class", "1", "1", db)

	for ClassCount.Next() {
		classCount++
	}
	defer ClassCount.Close()

	ClassCards := Database.SelectFromDB("ClassName", "Class", "1", "1", db)

	userGroup := app.Group("/Classes")

	var classCards []string

	for ClassCards.Next() {
		var class string
		ClassCards.Scan(&class)

		classCards = append(classCards, class)
	}
	defer ClassCards.Close()

	// ClassSend := Database.SelectFromDB("*", "Class", "1", "1", db)

	// var i = 0
	// Make a userGroup entry
	userGroup.Get("/:Class", func(c *fiber.Ctx) error {
		if true {
			db := Database.DBConnect()
			name := c.Params("Class")

			var studentCards []string
			StudentNameList := Database.SelectFromDB("English Name", "student", "Class", name, db)

			var i int = 0
			var studentCount int = 0
			for StudentNameList.Next() {
				studentCount++
			}
			StudentCards := Database.SelectFromDB("English Name", "student", "Class", name, db)

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
			StudentNationalIds := Database.SelectFromDB("National ID (14 digit)", "student", "Class", name, db)
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

			// Return Html page template
			return c.Render("ClassCardDetails", fiber.Map{
				"StudentList":       studentCards,
				"StudentGenderList": studentGenders,
				"StudentLinks":      studentLinks,
				"ClassName":         name,
			})
		} else if Controllers.CurrentUser.Permission == 0 {
			// Print No Permission
			return FormBackend.TmplNotLoggedIn.Execute(c, nil)
		} else {
			return FormBackend.TmplNoPermission.Execute(c, nil)
		}
	})
}

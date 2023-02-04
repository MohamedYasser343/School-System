package Dashboard

import (
	"Snap/Controllers"
	"Snap/FormBackend"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmplPrincipal = template.Must(template.ParseFiles("./DashboardFiles/index.html"))

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		var UserName = Controllers.CurrentUser.Name
		db, err := sql.Open("mysql", "snap:Snapsnap@2@tcp(92.205.60.182:3306)/snap")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()
		StdList, err := db.Query("SELECT * FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StdList.Close()
		var stdCount int = 0
		for StdList.Next() {
			stdCount++
		}
		TeacherList, err := db.Query("SELECT * FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer TeacherList.Close()
		var teacherCount int = 0
		for TeacherList.Next() {
			teacherCount++
		}
		OverLookerList, err := db.Query("SELECT * FROM `overlooking` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer OverLookerList.Close()
		var overLookerCount int
		for OverLookerList.Next() {
			overLookerCount++
		}
		Grade10, err := db.Query("SELECT * FROM `student` WHERE `Grade` = 'G10'")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer Grade10.Close()
		var grade10Count int
		for Grade10.Next() {
			grade10Count++
		}
		Grade11, err := db.Query("SELECT * FROM `student` WHERE `Grade` = 'G11'")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer Grade11.Close()
		var grade11Count int
		for Grade11.Next() {
			grade11Count++
		}

		Grade12, err := db.Query("SELECT * FROM `student` WHERE `Grade` = 'G12'")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer Grade12.Close()
		var grade12Count int
		for Grade12.Next() {
			grade12Count++
		}

		tmplPrincipal.Execute(w, struct {
			Sidebar         template.HTML
			StdCount        int
			TeacherCount    int
			OverLookerCount int
			User            string
			Grade10         int
			Grade11         int
			Grade12         int
		}{
			Controllers.CurrentUser.SideBar,
			stdCount,
			teacherCount,
			overLookerCount,
			UserName,
			grade10Count,
			grade11Count,
			grade12Count,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		fmt.Println("You are not allowed to access this page")
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

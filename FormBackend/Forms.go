package FormBackend

import (
	"Snap/AbstractFunctions"
	"Snap/Controllers"
	"Snap/Database"
	"Snap/Types"

	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mavihq/persian"
)

var (
	tmplAddStudent             = template.Must(template.ParseFiles("./DashboardFiles/AddStudent.html"))
	tmplAddViolation           = template.Must(template.ParseFiles("./DashboardFiles/AddViolation.html"))
	tmplAddTeacher             = template.Must(template.ParseFiles("./DashboardFiles/add-teacher.html"))
	tmplAddClass               = template.Must(template.ParseFiles("./DashboardFiles/AddClass.html"))
	tmplAddVacation            = template.Must(template.ParseFiles("./DashboardFiles/AddVacation.html"))
	tmplRecordVacation         = template.Must(template.ParseFiles("./DashboardFiles/RecordTeaVacation.html"))
	tmplAddBus                 = template.Must(template.ParseFiles("./DashboardFiles/AddBus.html"))
	tmplBussing                = template.Must(template.ParseFiles("./DashboardFiles/assign-bussing.html"))
	recordStdCapstone          = template.Must(template.ParseFiles("./Forms/RecordStudentCapstone.html"))
	tmplTrip                   = template.Must(template.ParseFiles("./Forms/Trip.html"))
	tmplSupervisor             = template.Must(template.ParseFiles("./Forms/Supervisor.html"))
	tmplAssignUniversity       = template.Must(template.ParseFiles("./DashboardFiles/AssignStudentUniversity.html"))
	tmplRecordStudentViolation = template.Must(template.ParseFiles("./DashboardFiles/RecordStdViolation.html"))
	tmplRecordTeacherViolation = template.Must(template.ParseFiles("./DashboardFiles/RecordTeaViolation.html"))
	tmplRomming                = template.Must(template.ParseFiles("./Forms/Romming.html"))
	tmplroom                   = template.Must(template.ParseFiles("./DashboardFiles/AddRoom.html"))
	tmplSchool                 = template.Must(template.ParseFiles("./Forms/School.html"))
	tmplUniversity             = template.Must(template.ParseFiles("./DashboardFiles/AddUniversity.html"))
	tmplEntercomp              = template.Must(template.ParseFiles("./Forms/EnterComp.html"))
	TmplNotLoggedIn            = template.Must(template.ParseFiles("./DashboardFiles/NotLoggedIn.html"))
	TmplNoPermission           = template.Must(template.ParseFiles("./DashboardFiles/NotEnoughPermission.html"))
	tmplAddScholarShip         = template.Must(template.ParseFiles("./DashboardFiles/AddScholarShip.html"))
	tmplAssignScholarShip      = template.Must(template.ParseFiles("./DashboardFiles/AssignStudentScholarShip.html"))
	tmplRemoveFromAnyTable     = template.Must(template.ParseFiles("./DashboardFiles/DeleteFromAnyTable.html"))
	tmplAssignStudentClass     = template.Must(template.ParseFiles("./DashboardFiles/AssignStudentToClass.html"))
)

var StudentDetails []Types.StudentDetail

//func HomePage(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		tmpl.Execute(w, nil)
//		return
//	}
//
//	var details = InputForm{
//		//Name: r.FormValue("Name"),
//		Radio: r.FormValue("Radio"),
//		Id:    r.FormValue("Id"),
//	}
//
//	_ = details
//	fmt.Println(details.Id, details.Radio)
//
//	db, err := sql.Open("mysql", Database.DBAddress)
//	defer db.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	violation := ""
//	switch details.Radio {
//	case "تاخر عن الطابور":
//		insert, err := db.Query("UPDATE StudentViolations SET LateForMorningLine = LateForMorningLine + 1 WHERE ID = ?", details.Id)
//		if err != nil {
//			panic(err.Error())
//		}
//		defer insert.Close()
//		violation = "Being Late For Morning Line."
//	case "اكل مخالف":
//		insert, err := db.Query("UPDATE StudentViolations SET BannedFood = BannedFood + 1 WHERE ID = ?", details.Id)
//		if err != nil {
//			panic(err.Error())
//		}
//		defer insert.Close()
//		violation = "Banned Food."
//	}
//
//	data := map[string]string{
//		"msg": "Hello World1",
//		"sum": "Happy Day",
//	}
//	c := fcm.NewFCM("AAAAaiyBXQk:APA91bEIkJdn7e1c5tEdee952MHDknzJ0kU2vLr4gPjORHGmddrkhSHAkyjBzU_8Ubycbl-BLSLGNGobOSaf9x3IedXgg-sqd4OJAB7z0TcosP-b3htZ9w5CZbXb15s9QdZqqiJiG_vN")
//	token, err := db.Query("SELECT Token FROM `users` WHERE id = ?;", details.Id)
//	if err != nil {
//		fmt.Print(err)
//	}
//	defer token.Close()
//	tokenStr := ""
//	for token.Next() {
//		err := token.Scan(&tokenStr)
//		if err != nil {
//			log.Fatal(err.Error())
//		}
//		fmt.Println(tokenStr)
//	}
//	if tokenStr != "" {
//		response, err := c.Send(fcm.Message{
//			Data:             data,
//			RegistrationIDs:  []string{tokenStr},
//			ContentAvailable: true,
//			// To:               "4oESYbdtfSQqvwO7gFXoo0stm3q1",
//			Priority: fcm.PriorityHigh,
//			Notification: fcm.Notification{
//				Title: "You Have A Report.",
//				Body:  "Mr. Ahmed Has Given You A Report For " + violation + ".",
//				Icon:  "@mipmap/ic_launcher",
//			},
//		})
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println("Status Code   :", response.StatusCode)
//		fmt.Println("Success       :", response.Success)
//		fmt.Println("Fail          :", response.Fail)
//		fmt.Println("Canonical_ids :", response.CanonicalIDs)
//		fmt.Println("Topic MsgId   :", response.MsgID)
//	}
//	//saveViolations()
//	//StudentDetails = append(StudentDetails, StudentDetail{details.Name, details.Id, details.Violations1})
//	tmpl.Execute(w, struct {
//	}{})
//}

func ExecuteError(w http.ResponseWriter, r *http.Request, isLoggedIn bool) {
	if !isLoggedIn {
		TmplNotLoggedIn.Execute(w, nil)
	} else {
		TmplNoPermission.Execute(w, nil)
	}
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		// Load HTML Template in string var

		db := Database.DBConnect()
		defer db.Close()
		// Get Class Count From Database
		ClassCount := Database.SelectFromDB("*", "Class", "1", "1", db)
		defer ClassCount.Close()

		var classCount int = 0
		for ClassCount.Next() {
			classCount++
		}
		// Get Class List From Database
		var classList []string
		rows, err := db.Query("SELECT `ClassName` FROM `Class`")
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		var i int = 0
		for rows.Next() {
			var className string
			err := rows.Scan(&className)
			if err != nil {
				log.Fatal(err.Error())
			}
			i++
			if i != classCount {
				classList = append(classList, className+",")
			} else {
				classList = append(classList, className)
			}
		}

		if r.Method != http.MethodPost {
			tmplAddStudent.Execute(w, struct {
				Sidebar   template.HTML
				ClassList []string
			}{
				Sidebar:   Controllers.CurrentUser.SideBar,
				ClassList: classList,
			})
			return
		}
		var details = Types.NewStudent{
			Name:           r.FormValue("Name"),
			NameE:          r.FormValue("NameE"),
			Grade:          r.FormValue("Grade"),
			Specialization: r.FormValue("Specialization"),
			SecondLang:     r.FormValue("SecondLang"),
			Religion:       r.FormValue("Religion"),
			PortfolioNo:    r.FormValue("PortfolioNo"),
			NationalId:     r.FormValue("NationalId"),
			Email:          r.FormValue("Email"),
			OfficialEmail:  r.FormValue("OfficialEmail"),
			Code:           r.FormValue("Code"),
			Class:          r.FormValue("Class"),
			UserName:       r.FormValue("UserName"),
			Password:       r.FormValue("Password"),
		}
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Get file uploaded via Form
		file, handler, err := r.FormFile("myfile")
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		defer file.Close()
		localFile, err := os.Create("/var/www/html/StudentImages/" + handler.Filename)
		studentPhotoLink := AbstractFunctions.ResolveHostIp() + "/StudentImages/" + handler.Filename
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		defer localFile.Close()

		// Copy the uploaded file data to the newly created file on the filesystem
		if _, err := io.Copy(localFile, file); err != nil {
			log.Fatal(err.Error())
			return
		}
		FatherName := strings.Split(details.NameE, " ")
		fmt.Println(FatherName)
		NationalId := persian.ToEnglishDigits(details.NationalId)
		chars := []rune(NationalId)
		//for i := 0; i < len(chars); i++ {
		//	char := string(chars[i])
		//	println(char)
		//}
		//30502130104231
		var Gender string
		if chars[12]%2 == 1 {
			Gender = "Male"
		} else {
			Gender = "Female"
		}
		birthCity := "N/A"
		birthCity = AbstractFunctions.ReturnBirthCity(chars)
		yearIndex := "0"
		switch string(chars[0]) {
		case "3":
			yearIndex = "20"
		case "2":
			yearIndex = "19"
			break
		}
		birthMonth := string(chars[3]) + string(chars[4])
		birthDay := string(chars[5]) + string(chars[6])
		yearIndex2 := string(chars[1]) + string(chars[2])
		birthYear := yearIndex + yearIndex2
		BirthDay := birthYear + "/" + birthMonth + "/" + birthDay
		fmt.Println(BirthDay)
		fmt.Println(Gender)
		fmt.Println(details.NationalId, details.Name)
		// StudentDetails = append(StudentDetails, StudentDetail{details.Name, details.NameE, details.Email, NationalId, details.Id, details.Gender, details.Religion, details.State, details.Address, details.SchoolType, details.SchoolCity, details.StudentNumber, details.StudentNumber2, details.FatherName, details.FatherNameE, details.FatherNationalId, details.FatherJob, details.FatherAddress, details.FatherMobile, details.FatherMobile2, details.MotherName, details.MotherNameE, details.MotherNationalId, details.MotherJob, details.MotherAddress, details.MotherMobile, details.MotherMobile2, 0, 0, 0, 0, BirthDay, birthCity, details.ParentName, details.ParentNameE, details.ParentNationalId, details.ParentJob, details.ParentAddress, details.ParentMobile, details.ParentMobile2})
		Null := "N.A"
		Permission, err := strconv.Atoi(details.Permission)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(studentPhotoLink, birthCity, Permission, Null)
		// insert, err := db.Query("INSERT INTO `student` (`stdname`, `stdnamee`, `stdpermesion`, `stdnationalid`, `stdgender`, `stdreligon`, `stdmail`, `stdstate`, `stdaddr`, `stdschool`, `stdschoolType`, `stdschoolProvince`, `stdschoolCity`, `stdmobile`, `stdmobile2`, `DadDeth`, `dadname`, `dadnationalid`, `dadjob`, `dadmobile`, `dadmobile2`, `dadaddr`, `MumDeth`, `mumname`, `mumnationalid`, `mumjob`, `mummobile`, `mummobile2`, `mumaddr`, `parentname`, `parentnationalid`, `parentjob`, `parentmobile`, `parentmobile2`, `parentaddr`, `stdbirthcity`, `stdbirthyear`, `stdday`, `stdmonth`, `stdyear`, `stdphotolink`, `UserName`, `Password`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", details.Name, details.NameE, Permission, details.NationalId, details.Gender, details.Gender, details.Religion, details.Email, details.State, details.Address, details.School, details.SchoolType, details.SchoolGovernment, details.SchoolCity, details.StudentNumber, Null, Null, details.FatherName, details.FatherNationalId, details.FatherJob, details.FatherMobile, details.FatherMobile2, details.FatherAddress, details.MotherLivingState, details.MotherName, details.MotherNationalId, details.MotherJob, details.MotherMobile, details.MotherMobile2, details.MotherAddress, details.ParentName, details.ParentNationalId, details.ParentJob, details.ParentMobile, details.ParentMobile2, details.ParentAddress, birthCity, birthYear, birthDay, birthMonth, Null, studentPhotoLink, details.UserName, details.Password)
		fmt.Println(details)
		insertStudent, err := db.Query("INSERT INTO `student` (`StdId`, `Code`, `English Name`, `Arabic Name`, `Grade`, `specilization`, `SecondLang`, `Religion`, `Gender`, `Portfolio No`, `National ID (14 digit)`, `Class`, `PersonalEmail`, `Email`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", details.Code, details.NameE, details.Name, details.Grade, details.Specialization, details.SecondLang, details.Religion, Gender, details.PortfolioNo, details.NationalId, details.Class, details.Email, details.OfficialEmail)
		insertUser, err := db.Query("INSERT INTO `users` VALUES(NULL, ?, ?, ?, ?)", details.NameE, details.UserName, details.Password, 1)
		if err != nil {
			panic(err.Error())
		}
		defer insertStudent.Close()
		defer insertUser.Close()
		// config := &firebase.Config{ProjectID: "snap-fcba6"}
		// app, err := firebase.NewApp(context.Background(), config)
		// if err != nil {
		// 	log.Fatalf("error initializing app: %v\n", err)
		// }
		// client, err := app.Auth(context.Background())
		// if err != nil {
		// 	log.Fatalf("error getting Auth client: %v\n", err)
		// }
		// params := (&auth.UserToCreate{}).
		// 	Email(details.Email).
		// 	EmailVerified(false).
		// 	Password("secretPassword").
		// 	Disabled(false)
		// u, err := client.CreateUser(context.Background(), params)
		// if err != nil {
		// 	log.Fatalf("error creating user: %v\n", err)
		// }
		// log.Printf("User created successfully : %v\n", u)

		tmplAddStudent.Execute(w, struct {
			Sidebar   template.HTML
			ClassList []string
		}{
			Sidebar:   Controllers.CurrentUser.SideBar,
			ClassList: classList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AssignStudentClass(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		db := Database.DBConnect()
		defer db.Close()
		// Get Class Count From Database
		ClassCount := Database.SelectFromDB("*", "Class", "1", "1", db)
		defer ClassCount.Close()

		var classCount int = 0
		for ClassCount.Next() {
			classCount++
		}
		// Get Class List From Database
		var classList []string
		rows, err := db.Query("SELECT `ClassName` FROM `Class`")
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		var i int = 0
		for rows.Next() {
			var className string
			err := rows.Scan(&className)
			if err != nil {
				log.Fatal(err.Error())
			}
			i++
			if i != classCount {
				classList = append(classList, className+",")
			} else {
				classList = append(classList, className)
			}
		}

		StudentCount, err := db.Query("SELECT * FROM `student` WHERE 1")
		defer StudentCount.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		var studentCount int

		for StudentCount.Next() {
			studentCount++
		}

		var studentList []string
		i = 0

		SelectStudents, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		defer SelectStudents.Close()
		if err != nil {
			fmt.Println(err.Error())
		}

		for SelectStudents.Next() {
			var student string
			err = SelectStudents.Scan(&student)
			if err != nil {
				fmt.Println(err.Error())
			}
			i++
			if i != studentCount {

				studentList = append(studentList, student+",")
			} else {
				studentList = append(studentList, student)
			}
		}
		if r.Method != http.MethodPost {
			tmplAssignStudentClass.Execute(w, struct {
				Sidebar     template.HTML
				ClassList   []string
				StudentList []string
			}{
				Sidebar:     Controllers.CurrentUser.SideBar,
				ClassList:   classList,
				StudentList: studentList,
			})
			return
		}
		var details = Types.AssignClass{
			StudentName: r.FormValue("StudentName"),
			ClassName:   r.FormValue("ClassName"),
		}
		// Get StudentId From Student Name
		var studentId int
		SelectStudentId, err := db.Query("SELECT `StdId` FROM `student` WHERE `English Name` = ?", details.StudentName)
		defer SelectStudentId.Close()
		if err != nil {
			fmt.Println(err.Error())
		}

		for SelectStudentId.Next() {
			err = SelectStudentId.Scan(&studentId)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		// Get ClassId From Class Name
		var classId int
		SelectClassId, err := db.Query("SELECT `ClassId` FROM `Class` WHERE `ClassName` = ?", details.ClassName)
		defer SelectClassId.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		for SelectClassId.Next() {
			err = SelectClassId.Scan(&classId)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		// Insert StudentId and ClassId Into StudentClass Table
		insertStudentClass, err := db.Query("INSERT INTO `ClassStd` VALUES(NULL, ?, ?)", classId, studentId)
		if err != nil {
			panic(err.Error())
		}
		defer insertStudentClass.Close()
		tmplAssignStudentClass.Execute(w, struct {
			Sidebar     template.HTML
			ClassList   []string
			StudentList []string
		}{
			Sidebar:     Controllers.CurrentUser.SideBar,
			ClassList:   classList,
			StudentList: studentList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddViolation(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplAddViolation.Execute(w, struct {
				Sidebar template.HTML
			}{
				Sidebar: Controllers.CurrentUser.SideBar,
			})
			return
		}

		var details = Types.NewViolation{
			ViolationKind:   r.FormValue("ViolationType"),
			ViolationPoints: r.FormValue("ViolationPoints"),
		}

		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `violations` (`violationid`, `violationName`, `violationPoint`) VALUES (NULL, ?, ?);", details.ViolationKind, details.ViolationPoints)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
		tmplAddViolation.Execute(w, struct {
			Sidebar template.HTML
		}{
			Sidebar: Controllers.CurrentUser.SideBar,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddTeacher(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplAddTeacher.Execute(w, struct {
				Sidebar template.HTML
			}{
				Controllers.CurrentUser.SideBar,
			})
			return
		}
		var details = Types.NewTeacher{
			Name:                     r.FormValue("Name"),
			Cadr:                     r.FormValue("Cadr"),
			Id:                       r.FormValue("Id"),
			LastJob:                  r.FormValue("LastJob"),
			STEMJob:                  r.FormValue("STEMJob"),
			Stage:                    r.FormValue("Stage"),
			HiringDate:               r.FormValue("HiringDate"),
			OrgQualification:         r.FormValue("OrgQualification"),
			PlaceOrgQualification:    r.FormValue("PlaceOrgQualification"),
			QualificationDate:        r.FormValue("QualificationDate"),
			Specialization:           r.FormValue("Specialization"),
			OrgQualificationGrade:    r.FormValue("OrgQualificationGrade"),
			HigherQualification:      r.FormValue("HigherQualification"),
			PlaceHigherQualification: r.FormValue("PlaceHigherQualification"),
			HighQualificationDate:    r.FormValue("HighQualificationDate"),
			HigherQualificationGrade: r.FormValue("HighQualificationGrade"),
			CurrentFinacialGrade:     r.FormValue("CurrentFinacialGrade"),
			ItsDate:                  r.FormValue("ItsDate"),
			OrgSchool:                r.FormValue("OrgSchool"),
			OrgAdministration:        r.FormValue("OrgAdministration"),
			OrgGovernment:            r.FormValue("OrgGovernment"),
			NameofSTEMSchool:         r.FormValue("NameofSTEMSchool"),
			DateJoinSTEMNow:          r.FormValue("DateJoinSTEMNow"),
			DateJoinSTEM:             r.FormValue("DateJoinSTEM"),
			NumberWhats:              r.FormValue("NumberWhats"),
			NumberPhone:              r.FormValue("NumberPhone"),
			EmailSTEM:                r.FormValue("EmailSTEM"),
			NationalID:               r.FormValue("NationalID"),
		}
		_ = details
		NationalId := persian.ToEnglishDigits(details.NationalID)
		chars := []rune(NationalId)
		for i := 0; i < len(chars); i++ {
			char := string(chars[i])
			println(char)
		}
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Get file uploaded via Form
		file, handler, err := r.FormFile("myfile")
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		defer file.Close()
		localFile, err := os.Create("/var/www/html/StudentImages/" + handler.Filename)
		studentPhotoLink := AbstractFunctions.ResolveHostIp() + "/StudentImages/" + handler.Filename
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		defer localFile.Close()

		// Copy the uploaded file data to the newly created file on the filesystem
		if _, err := io.Copy(localFile, file); err != nil {
			log.Fatal(err.Error())
			return
		}
		//30502130104231
		if chars[12]%2 == 1 {
			details.Gender = "Male"
		} else {
			details.Gender = "Female"
		}
		birthCity := "N/A"
		birthCity = AbstractFunctions.ReturnBirthCity(chars)
		_ = birthCity
		yearIndex := "0"
		switch string(chars[0]) {
		case "3":
			yearIndex = "20"
		case "2":
			yearIndex = "19"
			break
		}
		birthMonth := string(chars[3]) + string(chars[4])
		birthDay := string(chars[5]) + string(chars[6])
		yearIndex2 := string(chars[1]) + string(chars[2])
		birthYear := yearIndex + yearIndex2
		BirthDay := birthYear + "/" + birthMonth + "/" + birthDay
		fmt.Println(BirthDay)
		db := Database.DBConnect()
		//perm, err := strconv.Atoi(details.TeacherPermesion)
		//if err != nil {
		//	perm = 0
		//}
		insert, err := db.Query("INSERT INTO `teachers` (`teachersId`, `teacherCode`, `TeacherName`, `JobCadr`, `SupervisingPosOriginSchool`, `SupervisingPosStemSchool`, `EduLevel`, `HiringDate`, `OriginQualification`, `ObtainedOriginQualification`, `QualificationDate`, `Specialization`, `OriginQualificationGrade`, `HighQualification`, `ObtainedHighQualification`, `HighQualificationDate`, `HighQualificationGrade`, `FinancialGrade`, `FinancialGradeDate`, `OriginSchool`, `OriginTeachingGov`, `OriginCity`, `StemSchoolName`, `CurrentStemSchoolEnterDate`, `StemSchoolEnterDate`, `MobileWhats`, `MobileCalls`, `Email`, `NationalId`, `Photo`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", details.TeacherCode, details.Name, details.Cadr, details.LastJob, details.STEMJob, details.Stage, details.HiringDate, details.OrgQualification, details.PlaceOrgQualification, details.QualificationDate, details.Specialization, details.OrgQualificationGrade, details.HigherQualification, details.PlaceHigherQualification, details.HighQualificationDate, details.HigherQualificationGrade, details.CurrentFinacialGrade, details.ItsDate, details.OrgSchool, details.OrgAdministration, details.OrgGovernment, details.NameofSTEMSchool, details.DateJoinSTEMNow, details.DateJoinSTEM, details.NumberWhats, details.NumberPhone, details.EmailSTEM, details.NationalID, studentPhotoLink)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplAddTeacher.Execute(w, struct {
			Sidebar template.HTML
		}{
			Controllers.CurrentUser.SideBar,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddClass(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplAddClass.Execute(w, struct {
				Sidebar template.HTML
			}{
				Controllers.CurrentUser.SideBar,
			})
			return
		}
		var details = Types.NewClass{
			ClassName: r.FormValue("ClassName"),
			Grade:     r.FormValue("Grade"),
		}
		_ = details
		db := Database.DBConnect()
		defer db.Close()
		insert, err := db.Query("INSERT INTO `Class` (`ClassId`, `ClassName`, `ClassGrade`) VALUES (NULL, ?, ?);", details.ClassName, details.Grade)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplAddClass.Execute(w, struct {
			Sidebar template.HTML
		}{
			Controllers.CurrentUser.SideBar,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddVacation(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplAddVacation.Execute(w, struct {
				Sidebar template.HTML
			}{
				Controllers.CurrentUser.SideBar,
			})
			return
		}
		var Name = r.FormValue("VacationType")
		db := Database.DBConnect()
		defer db.Close()
		insert, err := db.Query("INSERT INTO `addvacation` (`addvacationid`, `vacationname`) VALUES (NULL, ?);", Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplAddVacation.Execute(w, struct {
			Sidebar template.HTML
		}{
			Controllers.CurrentUser.SideBar,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func Bussing(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 0 {
		var stdList []string
		db := Database.DBConnect()
		defer db.Close()
		StdList, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StdList.Close()
		var i int = 0
		var stdCount int = 0
		for StdList.Next() {
			stdCount++
		}
		fmt.Println(stdCount)
		StdList2, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for StdList2.Next() {
			var std string
			err = StdList2.Scan(&std)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != stdCount {
				stdList = append(stdList, std+",")
			} else {
				stdList = append(stdList, std)
			}
		}
		defer StdList2.Close()
		var busList []string
		BusList, err := db.Query("SELECT `endpoint` FROM `bus` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for BusList.Next() {
			var bus string
			err = BusList.Scan(&bus)
			if err != nil {
				fmt.Println(err)
			}
			busList = append(busList, bus)
		}
		defer BusList.Close()
		if r.Method != http.MethodPost {
			tmplBussing.Execute(w, struct {
				Sidebar template.HTML
				StdList []string
				BusList []string
			}{
				Controllers.CurrentUser.SideBar,
				stdList,
				busList,
			})
			return
		}
		var details = Types.AssignBussing{
			BusId: r.FormValue("BusId"),
			StdId: r.FormValue("StdName"),
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()
		StdId, err := db.Query("SELECT `StdId` FROM `student` WHERE `English Name` = ?;", details.StdId)
		if err != nil {
			fmt.Println(err.Error())
		}
		var stdid int
		for StdId.Next() {
			err = StdId.Scan(&stdid)
		}
		defer StdId.Close()
		BusId, err := db.Query("SELECT `busid` FROM `bus` WHERE `endpoint` = ?;", details.BusId)
		if err != nil {
			fmt.Println(err.Error())
		}
		var busid int
		for BusId.Next() {
			err = BusId.Scan(&busid)
		}
		defer BusId.Close()
		fmt.Println(stdid, busid)
		insert, err := db.Query("INSERT INTO `bussing` (`bussingid`, `busid`, `stdid`) VALUES (NULL, ?, ?);", busid, stdid)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplBussing.Execute(w, struct {
			Sidebar template.HTML
			StdList []string
			BusList []string
		}{
			Controllers.CurrentUser.SideBar,
			stdList,
			busList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddBus(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 0 {
		if r.Method != http.MethodPost {
			tmplAddBus.Execute(w, struct {
				Sidebar template.HTML
			}{
				Controllers.CurrentUser.SideBar,
			})
			return
		}
		var details = Types.NewBus{
			BusLine:     r.FormValue("BusLine"),
			DriverName:  r.FormValue("DriverName"),
			DriverPhone: r.FormValue("DriverPhone"),
			StartPoint:  r.FormValue("StartPoint"),
			StartTime:   r.FormValue("StartTime"),
			BusChairs:   r.FormValue("BusChairs"),
			EndPoint:    r.FormValue("EndPoint"),
			EndTime:     r.FormValue("EndTime"),
		}
		_ = details
		db := Database.DBConnect()
		defer db.Close()
		insert, err := db.Query("INSERT INTO `bus` (`busid`, `busline`, `drivername`, `driverNo`, `startpoint`, `starttime`, `busChairno`, `endpoint`, `endtime`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?);", details.BusLine, details.DriverName, details.DriverPhone, details.StartPoint, details.StartTime, details.BusChairs, details.EndPoint, details.EndTime)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplAddBus.Execute(w, struct {
			Sidebar template.HTML
		}{
			Controllers.CurrentUser.SideBar,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func RecordVacation(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		var teacherList []string
		var teacherCount int
		var i int = 0
		db := Database.DBConnect()
		defer db.Close()
		TeacherList, err := db.Query("SELECT * FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer TeacherList.Close()
		for TeacherList.Next() {
			teacherCount++
		}
		TeacherNames, err := db.Query("SELECT `TeacherName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for TeacherNames.Next() {
			var teacher string
			err = TeacherNames.Scan(&teacher)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherList = append(teacherList, teacher+",")
			} else {
				teacherList = append(teacherList, teacher)
			}
		}
		defer TeacherNames.Close()
		Vacations, err := db.Query("SELECT * FROM `addvacation` WHERE 1")

		var vacationTypes []string
		var vacationCount int
		if err != nil {
			log.Fatal(err.Error())
		}
		for Vacations.Next() {
			vacationCount++
		}
		defer Vacations.Close()
		i = 0
		VacationTypes, err := db.Query("SELECT `vacationname` FROM `addvacation` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for VacationTypes.Next() {
			var vacation string
			err = VacationTypes.Scan(&vacation)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != vacationCount {
				vacationTypes = append(vacationTypes, vacation+",")
			} else {
				vacationTypes = append(vacationTypes, vacation)
			}
		}
		defer VacationTypes.Close()
		if r.Method != http.MethodPost {
			tmplRecordVacation.Execute(w, struct {
				Sidebar       template.HTML
				VacationTypes []string
				TeacherNames  []string
			}{
				Controllers.CurrentUser.SideBar,
				vacationTypes,
				teacherList,
			})
			return
		}
		var details = Types.RecordVacationForm{
			TeacherName:  r.FormValue("TeacherName"),
			VacationType: r.FormValue("VacationType"),
			StartDate:    r.FormValue("StartDate"),
			EndDate:      r.FormValue("EndDate"),
		}
		_ = details
		TeacherId, err := db.Query("SELECT `teachersId` FROM `teachers` WHERE `TeacherName` = ?;", details.TeacherName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var teacherId int
		for TeacherId.Next() {
			err = TeacherId.Scan(&teacherId)
		}
		defer TeacherId.Close()

		VacationId, err := db.Query("SELECT `addvacationid` FROM `addvacation` WHERE `vacationname` = ?;", details.VacationType)
		if err != nil {
			fmt.Println(err.Error())
		}
		var vacationId int
		for VacationId.Next() {
			err = VacationId.Scan(&vacationId)
		}
		defer VacationId.Close()

		insert, err := db.Query("INSERT INTO `recordvacation` (`recordvacationid`, `teacherid`, `vacationtype`, `vacationfrom`, `vacationto`, `numdays`) VALUES (NULL, ?, ?, ?, ?, NULL);", teacherId, vacationId, details.StartDate, details.EndDate)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplRecordVacation.Execute(w, struct {
			Sidebar       template.HTML
			VacationTypes []string
			TeacherNames  []string
		}{
			Controllers.CurrentUser.SideBar,
			vacationTypes,
			teacherList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func RecordStdCapstone(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			recordStdCapstone.Execute(w, nil)
			return
		}

		var details = Types.AddStdToCapstone{
			CapstoneNumber: r.FormValue("CapstoneNumber"),
			Student1:       r.FormValue("Student1"),
			Student2:       r.FormValue("Student2"),
			Student3:       r.FormValue("Student3"),
			Student4:       r.FormValue("Student4"),
			Student5:       r.FormValue("Student5"),
			Student6:       r.FormValue("Student6"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `recordstdcapstone` (`recordstdcapstoneid`, `capstoneid`, `studentid`, `id2`, `id3`, `id4`, `id5`, `id6`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?);", details.CapstoneNumber, details.Student1, details.Student2, details.Student3, details.Student4, details.Student5, details.Student6)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddTrip(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplTrip.Execute(w, nil)
			return
		}

		var details = Types.NewTrip{
			TripName:          r.FormValue("TripName"),
			TripDate:          r.FormValue("TripDate"),
			TripPlace:         r.FormValue("TripPlace"),
			ChairsNumber:      r.FormValue("ChairsNumber"),
			TripKind:          r.FormValue("TripKind"),
			GeneralSupervisor: r.FormValue("GeneralSupervisor"),
			Supervisor1:       r.FormValue("Supervisor1"),
			Supervisor2:       r.FormValue("Supervisor2"),
			Supervisor3:       r.FormValue("Supervisor3"),
			Supervisor4:       r.FormValue("Supervisor4"),
			Supervisor5:       r.FormValue("Supervisor5"),
			LastDate:          r.FormValue("LastDate"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `trip` (`tripid`, `tripname`, `tripdate`, `tripplace`, `tripchairnum`, `triptype`, `publicoverlooking`, `overlooking`, `overlooking2`, `overlooking3`, `overlooking4`, `overlooking5`, `takeofftime`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", details.TripName, details.TripDate, details.TripPlace, details.ChairsNumber, details.TripKind, details.GeneralSupervisor, details.Supervisor1, details.Supervisor2, details.Supervisor3, details.Supervisor4, details.Supervisor5, details.LastDate)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func Supervisor(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplSupervisor.Execute(w, nil)
			return
		}

		var details = Types.NewSupervisor{
			SupervisorName:        r.FormValue("SupervisorName"),
			SupervisorNumber:      r.FormValue("SupervisorNumber"),
			SupervisorId:          r.FormValue("SupervisorId"),
			SupervisorAddress:     r.FormValue("SupervisorAddress"),
			GraduationCertificate: r.FormValue("GraduationCertificate"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `overlooking` (`overlookingid`, `overlookingname`, `overlookingmobile`, `overlookingnationalid`, `overlookingaddr`, `educationalqualification`) VALUES (NULL, ?, ?, ?, ?, ?);", details.SupervisorName, details.SupervisorNumber, details.SupervisorId, details.SupervisorAddress, details.GraduationCertificate)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func RecordStudentUniversity(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		var studentCount int
		db := Database.DBConnect()
		defer db.Close()
		StudentList, err := db.Query("SELECT * FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StudentList.Close()
		for StudentList.Next() {
			studentCount++
		}
		var studentList []string
		var i int = 0
		StudentNames, err := db.Query("SELECT `Arabic Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StudentNames.Close()
		for StudentNames.Next() {
			var student string
			err = StudentNames.Scan(&student)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != studentCount {
				studentList = append(studentList, student+",")
			} else {
				studentList = append(studentList, student)
			}
		}

		var universityCount int
		UniversityList, err := db.Query("SELECT * FROM `University` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniversityList.Close()
		for UniversityList.Next() {
			universityCount++
		}

		i = 0
		var universityList []string
		UniversityNames, err := db.Query("SELECT `UniversityName` FROM `University` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniversityNames.Close()
		for UniversityNames.Next() {
			var university string
			err = UniversityNames.Scan(&university)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != universityCount {
				universityList = append(universityList, university+",")
			} else {
				universityList = append(universityList, university)
			}
		}

		if r.Method != http.MethodPost {
			tmplAssignUniversity.Execute(w, struct {
				Sidebar        template.HTML
				StudentList    []string
				UniversityList []string
			}{
				Controllers.CurrentUser.SideBar,
				studentList,
				universityList,
			})
			return
		}

		var details = Types.AssignScholarShip{
			StudentName:     r.FormValue("StudentName"),
			UniversityName:  r.FormValue("UniversityName"),
			ScholarShipKind: r.FormValue("ScholarShipKind"),
		}

		StudentId, err := db.Query("SELECT `StdId` FROM `student` WHERE `Arabic Name` = ?;", details.StudentName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var studentId int
		for StudentId.Next() {
			err = StudentId.Scan(&studentId)
		}
		defer StudentId.Close()

		UniversityID, err := db.Query("SELECT `UniversityID` FROM `University` WHERE `UniversityName` = ?;", details.UniversityName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var universityID int
		for UniversityID.Next() {
			err = UniversityID.Scan(&universityID)
			if err != nil {
				log.Println(err.Error())
			}
		}
		defer UniversityID.Close()

		insert, err := db.Query("INSERT INTO `recordstduni` VALUES (NULL, ?, ?, ?)", studentId, universityID, details.ScholarShipKind)
		if err != nil {
			log.Println(err)
		}
		defer insert.Close()
		tmplAssignUniversity.Execute(w, struct {
			Sidebar        template.HTML
			StudentList    []string
			UniversityList []string
		}{
			Controllers.CurrentUser.SideBar,
			studentList,
			universityList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func RecordStudentViolation(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		db := Database.DBConnect()
		StudentCount, err := db.Query("SELECT * FROM `student` WHERE 1")
		defer StudentCount.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		var studentCount int

		for StudentCount.Next() {
			studentCount++
		}

		var studentList []string
		var i int

		SelectStudents, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		defer SelectStudents.Close()
		if err != nil {
			fmt.Println(err.Error())
		}

		for SelectStudents.Next() {
			var student string
			err = SelectStudents.Scan(&student)
			if err != nil {
				fmt.Println(err.Error())
			}
			i++
			if i != studentCount {

				studentList = append(studentList, student+",")
			} else {
				studentList = append(studentList, student)
			}
		}

		ViolationCount, err := db.Query("SELECT * FROM `violations` WHERE 1")
		defer ViolationCount.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		var violationCount int

		for ViolationCount.Next() {
			violationCount++
		}
		var violationList []string

		i = 0

		SelectViolations, err := db.Query("SELECT `violationName` FROM `violations` WHERE 1")

		for SelectViolations.Next() {
			var violation string
			err = SelectViolations.Scan(&violation)
			if err != nil {
				fmt.Println(err.Error())
			}
			i++
			if i != violationCount {
				violationList = append(violationList, violation+",")
			} else {
				violationList = append(violationList, violation)
			}
		}

		if r.Method != http.MethodPost {
			tmplRecordStudentViolation.Execute(w, struct {
				Sidebar        template.HTML
				StudentsList   []string
				ViolationsList []string
			}{
				Controllers.CurrentUser.SideBar,
				studentList,
				violationList,
			})
			return
		}

		var details = Types.NewRecordStudentViolation{
			StudentName:   r.FormValue("StudentName"),
			ViolationKind: r.FormValue("ViolationKind"),
			ViolationDate: r.FormValue("ViolationDate"),
		}

		SelectStudentId, err := db.Query("SELECT `StdId` FROM `student` WHERE `English Name` = \"" + details.StudentName + "\"")

		if err != nil {
			fmt.Println(err.Error())
		}
		var studentId string
		for SelectStudentId.Next() {
			err = SelectStudentId.Scan(&studentId)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer SelectStudentId.Close()
		SelectViolationId, err := db.Query("SELECT `violationid` FROM `violations` WHERE `violationName` = \"" + details.ViolationKind + "\"")

		if err != nil {
			fmt.Println(err.Error())
		}

		var violationId string
		for SelectViolationId.Next() {
			err = SelectViolationId.Scan(&violationId)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer SelectViolationId.Close()
		insert, err := db.Query("INSERT INTO `recordstdviolation` (`recordstdviolationid`, `studentid`, `violationid`, `violationdate`) VALUES (NULL, ?, ?, ?);", studentId, violationId, details.ViolationDate)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
		tmplRecordStudentViolation.Execute(w, struct {
			Sidebar        template.HTML
			StudentsList   []string
			ViolationsList []string
		}{
			Controllers.CurrentUser.SideBar,
			studentList,
			violationList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func RecordTeacherViolation(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		var teacherList []string
		var teacherCount int
		var i int = 0
		db := Database.DBConnect()
		defer db.Close()
		TeacherList, err := db.Query("SELECT * FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer TeacherList.Close()
		for TeacherList.Next() {
			teacherCount++
		}
		TeacherNames, err := db.Query("SELECT `TeacherName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for TeacherNames.Next() {
			var teacher string
			err = TeacherNames.Scan(&teacher)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherList = append(teacherList, teacher+",")
			} else {
				teacherList = append(teacherList, teacher)
			}
		}
		defer TeacherNames.Close()

		Violations, err := db.Query("SELECT * FROM `violations` WHERE 1")
		var violationTypes []string
		var violationCount int
		if err != nil {
			log.Fatal(err.Error())
		}
		for Violations.Next() {
			violationCount++
		}
		defer Violations.Close()
		i = 0
		ViolationTypes, err := db.Query("SELECT `violationName` FROM `violations` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for ViolationTypes.Next() {
			var violation string
			err = ViolationTypes.Scan(&violation)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != violationCount {
				violationTypes = append(violationTypes, violation+",")
			} else {
				violationTypes = append(violationTypes, violation)
			}
		}
		defer ViolationTypes.Close()

		if r.Method != http.MethodPost {
			tmplRecordTeacherViolation.Execute(w, struct {
				Sidebar        template.HTML
				TeacherNames   []string
				ViolationKinds []string
			}{
				Controllers.CurrentUser.SideBar,
				teacherList,
				violationTypes,
			})
			return
		}

		var details = Types.NewRecordTeacherViolation{
			TeacherName:   r.FormValue("TeacherName"),
			ViolationKind: r.FormValue("ViolationKind"),
			ViolationDate: r.FormValue("ViolationDate"),
		}

		TeacherId, err := db.Query("SELECT `teachersId` FROM `teachers` WHERE `TeacherName` = ?;", details.TeacherName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var teacherId int
		for TeacherId.Next() {
			err = TeacherId.Scan(&teacherId)
		}
		defer TeacherId.Close()

		ViolationId, err := db.Query("SELECT `violationid` FROM `violations` WHERE `violationName` = ?;", details.ViolationKind)
		if err != nil {
			fmt.Println(err.Error())
		}
		var violationId int
		for ViolationId.Next() {
			err = ViolationId.Scan(&violationId)
		}
		defer ViolationId.Close()

		insert, err := db.Query("INSERT INTO `recordteaviolation` (`recordteaviolationid`, `teacherid`, `teaviolationname`, `teaviolationdate`) VALUES (NULL, ?, ?, ?);", teacherId, violationId, details.ViolationDate)

		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		tmplRecordTeacherViolation.Execute(w, struct {
			Sidebar        template.HTML
			TeacherNames   []string
			ViolationKinds []string
		}{
			Controllers.CurrentUser.SideBar,
			teacherList,
			violationTypes,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func Rooming(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplRomming.Execute(w, nil)
			return
		}

		var details = Types.NewRooming{
			RoomNumber:  r.FormValue("RoomNumber"),
			StudentName: r.FormValue("StudentName"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `romming` (`rommingid`, `roomid`, `stdid`) VALUES (NULL, ?, ?);", details.RoomNumber, details.StudentName)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddRoom(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {

			db := Database.DBConnect()

			defer db.Close()
			// Fetch School FloorNo From Database
			FloorNo, err := db.Query("SELECT `FloorNo` FROM `school` WHERE 1")

			if err != nil {
				log.Println(err.Error())
			}
			var floorNo int
			for FloorNo.Next() {
				//Scan and Handle Error
				err = FloorNo.Scan(&floorNo)
				if err != nil {
					log.Println(err.Error())
				}
			}
			defer FloorNo.Close()
			tmplroom.Execute(w, struct {
				Sidebar template.HTML
				FloorNo int
			}{
				Controllers.CurrentUser.SideBar,
				floorNo,
			})
			return
		}

		var details = Types.NewRoom{
			RoomNumber:   r.FormValue("RoomNo"),
			RoomCapacity: r.FormValue("RoomCapacity"),
			FloorNo:      r.FormValue("FloorNo"),
		}

		db := Database.DBConnect()

		defer db.Close()
		// Fetch School FloorNo From Database
		FloorNo, err := db.Query("SELECT `FloorNo` FROM `school` WHERE 1")

		if err != nil {
			log.Println(err.Error())
		}
		var floorNo int
		for FloorNo.Next() {
			err = FloorNo.Scan(&floorNo)
			if err != nil {
				log.Println(err.Error())
			}
		}
		defer FloorNo.Close()

		insert, err := db.Query("INSERT INTO `room` (`roomid`, `roomno`, `maxnO`, `floorno`) VALUES (NULL, ?, ?, ?);", details.RoomNumber, details.RoomCapacity, details.FloorNo)

		if err != nil {
			log.Println(err.Error())
		}

		defer insert.Close()

		tmplroom.Execute(w, struct {
			Sidebar template.HTML
			FloorNo int
		}{
			Controllers.CurrentUser.SideBar,
			floorNo,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddScholarShip(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplAddScholarShip.Execute(w, struct {
				Sidebar template.HTML
			}{
				Controllers.CurrentUser.SideBar,
			})
			return
		}

		var details = Types.NewScholarShip{
			UniversityName: r.FormValue("UniversityName"),
			CollageName:    r.FormValue("CollageName"),
			Notes:          r.FormValue("Notes"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `scholarship` (`scolarshipid`, `universityname`, `collagename`, `notes`) VALUES (NULL, ?, ?, ?);", details.UniversityName, details.CollageName, details.Notes)

		if err != nil {
			log.Fatal(err.Error())
		}
		tmplAddScholarShip.Execute(w, struct {
			Sidebar template.HTML
		}{
			Controllers.CurrentUser.SideBar,
		})
		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddSchool(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplSchool.Execute(w, nil)
			return
		}

		var details = Types.NewSchool{
			SchoolName:    r.FormValue("SchoolName"),
			DateCreated:   r.FormValue("DateCreated"),
			PrincipalName: r.FormValue("PrincipalName"),
			DeputyName:    r.FormValue("DeputyName"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `school` (`schoolid`, `schoolname`, `year`, `princname`, `depname`) VALUES (NULL, ?, ?, ?, ?);", details.SchoolName, details.DateCreated, details.PrincipalName, details.DeputyName)

		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AddUniversity(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplUniversity.Execute(w, struct {
				Sidebar template.HTML
			}{
				Controllers.CurrentUser.SideBar,
			})
			return
		}

		var details = Types.NewUniversity{
			UniversityName: r.FormValue("UniversityName"),
			CollegeName:    r.FormValue("CollegeName"),
			CityName:       r.FormValue("CityName"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `University` (`UniversityID`, `UniversityName`, `CollageName`, `City`) VALUES (NULL, ?, ?, ?);", details.UniversityName, details.CollegeName, details.CityName)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
		tmplUniversity.Execute(w, struct {
			Sidebar template.HTML
		}{
			Controllers.CurrentUser.SideBar,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func Entercomp(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		if r.Method != http.MethodPost {
			tmplEntercomp.Execute(w, nil)
			return
		}

		var details = Types.NewEnterComp{
			CompetitionName: r.FormValue("CompetitionName"),
			StudentName:     r.FormValue("StudentName"),
		}

		_ = details
		db := Database.DBConnect()

		defer db.Close()
		insert, err := db.Query("INSERT INTO `entercompid` (`StudentName`, `competitionid`, `stdid`) VALUES (NULL, ?, ?);", details.CompetitionName, details.StudentName)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer insert.Close()
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func AssignStudentScholarShip(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		var studentCount int
		db := Database.DBConnect()
		defer db.Close()
		StudentList, err := db.Query("SELECT * FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StudentList.Close()
		for StudentList.Next() {
			studentCount++
		}
		var studentList []string
		var i int = 0
		StudentNames, err := db.Query("SELECT `Arabic Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StudentNames.Close()
		for StudentNames.Next() {
			var student string
			err = StudentNames.Scan(&student)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != studentCount {
				studentList = append(studentList, student+",")
			} else {
				studentList = append(studentList, student)
			}
		}

		var universityCount int
		UniversityList, err := db.Query("SELECT * FROM `scholarship` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniversityList.Close()
		for UniversityList.Next() {
			universityCount++
		}

		i = 0
		var universityList []string
		UniversityNames, err := db.Query("SELECT `universityname` FROM `scholarship` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniversityNames.Close()
		for UniversityNames.Next() {
			var university string
			err = UniversityNames.Scan(&university)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != universityCount {
				universityList = append(universityList, university+",")
			} else {
				universityList = append(universityList, university)
			}
		}

		if r.Method != http.MethodPost {
			tmplAssignScholarShip.Execute(w, struct {
				Sidebar        template.HTML
				StudentList    []string
				UniversityList []string
			}{
				Controllers.CurrentUser.SideBar,
				studentList,
				universityList,
			})
			return
		}

		var details = Types.AssignScholarShip{
			StudentName:     r.FormValue("StudentName"),
			UniversityName:  r.FormValue("UniversityName"),
			ScholarShipKind: r.FormValue("ScholarShipKind"),
		}

		StudentId, err := db.Query("SELECT `StdId` FROM `student` WHERE `Arabic Name` = ?;", details.StudentName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var studentId int
		for StudentId.Next() {
			err = StudentId.Scan(&studentId)
		}
		defer StudentId.Close()

		ScholarShipId, err := db.Query("SELECT `scolarshipid` FROM `scholarship` WHERE `universityname` = ?;", details.UniversityName)
		if err != nil {
			fmt.Println(err.Error())
		}
		var scholarShipId int
		for ScholarShipId.Next() {
			err = ScholarShipId.Scan(&scholarShipId)
		}
		defer ScholarShipId.Close()

		insert, err := db.Query("INSERT INTO `AssignScholar` VALUES (NULL, ?, ?, ?)", studentId, scholarShipId, details.ScholarShipKind)
		if err != nil {
			log.Println(err)
		}
		defer insert.Close()
		tmplAssignScholarShip.Execute(w, struct {
			Sidebar        template.HTML
			StudentList    []string
			UniversityList []string
		}{
			Controllers.CurrentUser.SideBar,
			studentList,
			universityList,
		})
	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

func DeleteFromAnyTable(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		db := Database.DBConnect()

		defer db.Close()
		GetCount, err := db.Query("SHOW TABLES")
		var count int
		for GetCount.Next() {
			count++
		}
		defer GetCount.Close()

		GetTables, err := db.Query("SHOW TABLES")

		var tables []string
		var i int = 0

		for GetTables.Next() {
			var table string
			err = GetTables.Scan(&table)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != count {
				tables = append(tables, table+",")
			} else {
				tables = append(tables, table)
			}
		}

		defer GetTables.Close()
		fmt.Println(tables)
		if r.Method != http.MethodPost {
			tmplRemoveFromAnyTable.Execute(w, struct {
				TableList []string
			}{
				tables,
			})
			return
		}
		var details = r.FormValue("TableName")
		GetColumns, err := db.Query("SELECT * FROM TABLE " + details)
		if err != nil {
			fmt.Println(err.Error())
		}
		var Columns []string
		for GetColumns.Next() {
			var column string
			err := GetColumns.Scan(&column)
			if err != nil {
				fmt.Println(err.Error())
			}
			Columns = append(Columns, column)
		}

	} else if Controllers.CurrentUser.Permission == 0 {
		ExecuteError(w, r, false)
	} else {
		ExecuteError(w, r, true)
	}
}

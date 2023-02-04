package SendData

import (
	"Snap/Controllers"
	"Snap/Database"
	"Snap/FormBackend"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	tmplAllStudent   = template.Must(template.ParseFiles("./DashboardFiles/all-student.html"))
	tmplAllTeacher   = template.Must(template.ParseFiles("./DashboardFiles/all-teacher.html"))
	tmplAllVacation  = template.Must(template.ParseFiles("./DashboardFiles/all-vacations.html"))
	tmplAllViolation = template.Must(template.ParseFiles("./DashboardFiles/all-violations.html"))
	// tmplAllClasses      = template.Must(template.ParseFiles("./DashboardFiles/all-classes.html"))
	tmplAllUniversities = template.Must(template.ParseFiles("./DashboardFiles/all-universities.html"))
	tmplAllScholarShips = template.Must(template.ParseFiles("./DashboardFiles/all-scholarships.html"))
)

func ShowAllStudents(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		var stdNameList []string
		db := Database.DBConnect()
		defer db.Close()
		StdNameList, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StdNameList.Close()
		var i int = 0
		var stdCount int = 0
		for StdNameList.Next() {
			stdCount++
		}
		StdNameList2, err := db.Query("SELECT `English Name` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for StdNameList2.Next() {
			var std string
			err = StdNameList2.Scan(&std)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != stdCount {
				stdNameList = append(stdNameList, std+",")
			} else {
				stdNameList = append(stdNameList, std)
			}
		}
		defer StdNameList2.Close()
		var stdGradeList []string
		StdGradeList, err := db.Query("SELECT `Grade` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		i = 0
		for StdGradeList.Next() {
			var stdGrade string
			err = StdGradeList.Scan(&stdGrade)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != stdCount {
				stdGradeList = append(stdGradeList, stdGrade+",")
			} else {
				stdGradeList = append(stdGradeList, stdGrade)
			}
		}
		defer StdGradeList.Close()
		var stdClassList []string
		StdClassList, err := db.Query("SELECT `Class` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		i = 0
		for StdClassList.Next() {
			var stdClass string
			err = StdClassList.Scan(&stdClass)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != stdCount {
				stdClassList = append(stdClassList, stdClass+",")
			} else {
				stdClassList = append(stdClassList, stdClass)
			}
		}
		var stdEmailList []string
		StdEmailList, err := db.Query("SELECT `Official mail` FROM `student` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		i = 0
		for StdEmailList.Next() {
			var stdEmail string
			err = StdEmailList.Scan(&stdEmail)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != stdCount {
				stdEmailList = append(stdEmailList, stdEmail+",")
			} else {
				stdEmailList = append(stdEmailList, stdEmail)
			}
		}
		defer StdEmailList.Close()
		tmplAllStudent.Execute(w, struct {
			StdNameList  []string
			StdGradeList []string
			StdClassList []string
			StdEmailList []string
		}{
			stdNameList,
			stdGradeList,
			stdClassList,
			stdEmailList,
		})
		return
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ShowAllTeachers(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		var teacherNameList []string
		db := Database.DBConnect()
		defer db.Close()
		TeacherNameList, err := db.Query("SELECT `TeacherName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var i int = 0
		var teacherCount int = 0
		for TeacherNameList.Next() {
			teacherCount++
		}
		defer TeacherNameList.Close()
		TeacherNameList2, err := db.Query("SELECT `TeacherName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		for TeacherNameList2.Next() {
			var teacherName string
			err = TeacherNameList2.Scan(&teacherName)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherNameList = append(teacherNameList, teacherName+",")
			} else {
				teacherNameList = append(teacherNameList, teacherName)
			}
		}
		defer TeacherNameList2.Close()
		i = 0
		TeacherNationalIds, err := db.Query("SELECT `NationalId` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherGenderList []string
		var teacherNationalIdList []string
		for TeacherNationalIds.Next() {
			var TeacherNationalId string
			err = TeacherNationalIds.Scan(&TeacherNationalId)
			if err != nil {
				fmt.Println(err)
			}
			chars := []rune(TeacherNationalId)
			var Gender string
			if chars[12]%2 == 1 {
				Gender = "Male"
			} else {
				Gender = "Female"
			}
			i++
			if i != teacherCount {
				teacherGenderList = append(teacherGenderList, Gender+",")
				teacherNationalIdList = append(teacherNationalIdList, TeacherNationalId+",")
			} else {
				teacherGenderList = append(teacherGenderList, Gender)
				teacherNationalIdList = append(teacherNationalIdList, TeacherNationalId)
			}
		}
		defer TeacherNationalIds.Close()
		i = 0
		TeacherJobList, err := db.Query("SELECT `SupervisingPosStemSchool` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherJobList []string
		for TeacherJobList.Next() {
			var TeacherJob string
			err = TeacherJobList.Scan(&TeacherJob)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherJobList = append(teacherJobList, TeacherJob+",")
			} else {
				teacherJobList = append(teacherJobList, TeacherJob)
			}
		}
		defer TeacherJobList.Close()
		i = 0
		TeacherPreviousJobList, err := db.Query("SELECT `SupervisingPosOriginSchool` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherPreviousJobList []string
		for TeacherPreviousJobList.Next() {
			var TeacherJob string
			err = TeacherPreviousJobList.Scan(&TeacherJob)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherPreviousJobList = append(teacherPreviousJobList, TeacherJob+",")
			} else {
				teacherPreviousJobList = append(teacherPreviousJobList, TeacherJob)
			}
		}
		defer TeacherPreviousJobList.Close()
		i = 0
		TeachersCadre, err := db.Query("SELECT `JobCadr` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherCadreList []string
		for TeachersCadre.Next() {
			var TeacherCadre string
			err = TeachersCadre.Scan(&TeacherCadre)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherCadreList = append(teacherCadreList, TeacherCadre+",")
			} else {
				teacherCadreList = append(teacherCadreList, TeacherCadre)
			}
		}
		defer TeachersCadre.Close()
		i = 0
		TeacherEducationLevels, err := db.Query("SELECT `EduLevel` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherEducationLevelsList []string
		for TeacherEducationLevels.Next() {
			var TeacherEducationLevel string
			err = TeacherEducationLevels.Scan(&TeacherEducationLevel)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherEducationLevelsList = append(teacherEducationLevelsList, TeacherEducationLevel+",")
			} else {
				teacherEducationLevelsList = append(teacherEducationLevelsList, TeacherEducationLevel)
			}
		}
		defer TeacherEducationLevels.Close()
		i = 0
		TeacherHiringDates, err := db.Query("SELECT `HiringDate` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherHiringDatesList []string
		for TeacherHiringDates.Next() {
			var TeacherHiringDate string
			err = TeacherHiringDates.Scan(&TeacherHiringDate)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherHiringDatesList = append(teacherHiringDatesList, TeacherHiringDate+",")
			} else {
				teacherHiringDatesList = append(teacherHiringDatesList, TeacherHiringDate)
			}
		}
		defer TeacherHiringDates.Close()
		i = 0
		Qualification, err := db.Query("SELECT `OriginQualification` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var qualificationList []string
		for Qualification.Next() {
			var Qualificatio string
			err = Qualification.Scan(&Qualificatio)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				qualificationList = append(qualificationList, Qualificatio+",")
			} else {
				qualificationList = append(qualificationList, Qualificatio)
			}
		}
		defer Qualification.Close()
		i = 0
		QualificationPlace, err := db.Query("SELECT `ObtainedOriginQualification` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var qualificationPlaceList []string
		for QualificationPlace.Next() {
			var QualificationPlac string
			err = QualificationPlace.Scan(&QualificationPlac)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				qualificationPlaceList = append(qualificationPlaceList, QualificationPlac+",")
			} else {
				qualificationPlaceList = append(qualificationPlaceList, QualificationPlac)
			}
		}
		defer QualificationPlace.Close()
		i = 0
		QualificationDate, err := db.Query("SELECT `QualificationDate` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var qualificationDateList []string
		for QualificationDate.Next() {
			var QualificationDat string
			err = QualificationDate.Scan(&QualificationDat)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				qualificationDateList = append(qualificationDateList, QualificationDat+",")
			} else {
				qualificationDateList = append(qualificationDateList, QualificationDat)
			}
		}
		defer QualificationDate.Close()

		i = 0
		TeacherSpecializations, err := db.Query("SELECT `Specialization` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherSpecializationsList []string
		for TeacherSpecializations.Next() {
			var TeacherSpecialization string
			err = TeacherSpecializations.Scan(&TeacherSpecialization)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherSpecializationsList = append(teacherSpecializationsList, TeacherSpecialization+",")
			} else {
				teacherSpecializationsList = append(teacherSpecializationsList, TeacherSpecialization)
			}
		}
		defer TeacherSpecializations.Close()
		i = 0
		TeacherOriginalQualificationGrades, err := db.Query("SELECT `OriginQualificationGrade` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherOriginalQualificationGradesList []string
		for TeacherOriginalQualificationGrades.Next() {
			var TeacherOriginalQualificationGrade string
			err = TeacherOriginalQualificationGrades.Scan(&TeacherOriginalQualificationGrade)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherOriginalQualificationGradesList = append(teacherOriginalQualificationGradesList, TeacherOriginalQualificationGrade+",")
			} else {
				teacherOriginalQualificationGradesList = append(teacherOriginalQualificationGradesList, TeacherOriginalQualificationGrade)
			}
		}
		defer TeacherOriginalQualificationGrades.Close()
		i = 0
		TeacherHighQualifications, err := db.Query("SELECT `HighQualification` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherHighQualificationsList []string
		for TeacherHighQualifications.Next() {
			var TeacherHighQualification string
			err = TeacherHighQualifications.Scan(&TeacherHighQualification)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherHighQualificationsList = append(teacherHighQualificationsList, TeacherHighQualification+",")
			} else {
				teacherHighQualificationsList = append(teacherHighQualificationsList, TeacherHighQualification)
			}
		}
		defer TeacherHighQualifications.Close()
		i = 0
		TeacherHighQualificationPlaces, err := db.Query("SELECT `ObtainedHighQualification` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherHighQualificationPlacesList []string
		for TeacherHighQualificationPlaces.Next() {
			var TeacherHighQualificationPlace string
			err = TeacherHighQualificationPlaces.Scan(&TeacherHighQualificationPlace)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherHighQualificationPlacesList = append(teacherHighQualificationPlacesList, TeacherHighQualificationPlace+",")
			} else {
				teacherHighQualificationPlacesList = append(teacherHighQualificationPlacesList, TeacherHighQualificationPlace)
			}
		}
		defer TeacherHighQualificationPlaces.Close()
		i = 0
		TeacherHighQualificationDates, err := db.Query("SELECT `HighQualificationDate` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherHighQualificationDatesList []string
		for TeacherHighQualificationDates.Next() {
			var TeacherHighQualificationDate string
			err = TeacherHighQualificationDates.Scan(&TeacherHighQualificationDate)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherHighQualificationDatesList = append(teacherHighQualificationDatesList, TeacherHighQualificationDate+",")
			} else {
				teacherHighQualificationDatesList = append(teacherHighQualificationDatesList, TeacherHighQualificationDate)
			}
		}
		defer TeacherHighQualificationDates.Close()
		i = 0
		TeacherHighQualificationGrades, err := db.Query("SELECT `HighQualificationGrade` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherHighQualificationGradesList []string
		for TeacherHighQualificationGrades.Next() {
			var TeacherHighQualificationGrade string
			err = TeacherHighQualificationGrades.Scan(&TeacherHighQualificationGrade)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherHighQualificationGradesList = append(teacherHighQualificationGradesList, TeacherHighQualificationGrade+",")
			} else {
				teacherHighQualificationGradesList = append(teacherHighQualificationGradesList, TeacherHighQualificationGrade)
			}
		}
		defer TeacherHighQualificationGrades.Close()

		i = 0
		TeacherFinancialGrades, err := db.Query("SELECT `FinancialGrade` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherFinancialGradesList []string
		for TeacherFinancialGrades.Next() {
			var TeacherFinancialGrade string
			err = TeacherFinancialGrades.Scan(&TeacherFinancialGrade)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherFinancialGradesList = append(teacherFinancialGradesList, TeacherFinancialGrade+",")
			} else {
				teacherFinancialGradesList = append(teacherFinancialGradesList, TeacherFinancialGrade)
			}
		}
		defer TeacherFinancialGrades.Close()
		i = 0
		TeacherFinancialGradeDates, err := db.Query("SELECT `FinancialGradeDate` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherFinancialGradeDatesList []string
		for TeacherFinancialGradeDates.Next() {
			var TeacherFinancialGradeDate string
			err = TeacherFinancialGradeDates.Scan(&TeacherFinancialGradeDate)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherFinancialGradeDatesList = append(teacherFinancialGradeDatesList, TeacherFinancialGradeDate+",")
			} else {
				teacherFinancialGradeDatesList = append(teacherFinancialGradeDatesList, TeacherFinancialGradeDate)
			}
		}
		defer TeacherFinancialGradeDates.Close()
		i = 0
		TeacherOriginSchools, err := db.Query("SELECT `OriginSchool` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherOriginSchoolsList []string
		for TeacherOriginSchools.Next() {
			var TeacherOriginSchool string
			err = TeacherOriginSchools.Scan(&TeacherOriginSchool)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherOriginSchoolsList = append(teacherOriginSchoolsList, TeacherOriginSchool+",")
			} else {
				teacherOriginSchoolsList = append(teacherOriginSchoolsList, TeacherOriginSchool)
			}
		}
		defer TeacherOriginSchools.Close()
		i = 0
		TeacherOriginTeachingGovs, err := db.Query("SELECT `OriginTeachingGov` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherOriginTeachingGovsList []string
		for TeacherOriginTeachingGovs.Next() {
			var TeacherOriginTeachingGov string
			err = TeacherOriginTeachingGovs.Scan(&TeacherOriginTeachingGov)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherOriginTeachingGovsList = append(teacherOriginTeachingGovsList, TeacherOriginTeachingGov+",")
			} else {
				teacherOriginTeachingGovsList = append(teacherOriginTeachingGovsList, TeacherOriginTeachingGov)
			}
		}
		defer TeacherOriginTeachingGovs.Close()
		i = 0
		TeacherOriginCitys, err := db.Query("SELECT `OriginCity` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherOriginCitysList []string
		for TeacherOriginCitys.Next() {
			var TeacherOriginCity string
			err = TeacherOriginCitys.Scan(&TeacherOriginCity)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherOriginCitysList = append(teacherOriginCitysList, TeacherOriginCity+",")
			} else {
				teacherOriginCitysList = append(teacherOriginCitysList, TeacherOriginCity)
			}
		}
		defer TeacherOriginCitys.Close()
		i = 0
		TeacherStemSchoolNames, err := db.Query("SELECT `StemSchoolName` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherStemSchoolNamesList []string
		for TeacherStemSchoolNames.Next() {
			var TeacherStemSchoolName string
			err = TeacherStemSchoolNames.Scan(&TeacherStemSchoolName)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherStemSchoolNamesList = append(teacherStemSchoolNamesList, TeacherStemSchoolName+",")
			} else {
				teacherStemSchoolNamesList = append(teacherStemSchoolNamesList, TeacherStemSchoolName)
			}
		}
		defer TeacherStemSchoolNames.Close()
		i = 0
		TeacherCurrentStemSchoolEnterDates, err := db.Query("SELECT `CurrentStemSchoolEnterDate` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherCurrentStemSchoolEnterDatesList []string
		for TeacherCurrentStemSchoolEnterDates.Next() {
			var TeacherCurrentStemSchoolEnterDate string
			err = TeacherCurrentStemSchoolEnterDates.Scan(&TeacherCurrentStemSchoolEnterDate)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherCurrentStemSchoolEnterDatesList = append(teacherCurrentStemSchoolEnterDatesList, TeacherCurrentStemSchoolEnterDate+",")
			} else {
				teacherCurrentStemSchoolEnterDatesList = append(teacherCurrentStemSchoolEnterDatesList, TeacherCurrentStemSchoolEnterDate)
			}
		}
		defer TeacherCurrentStemSchoolEnterDates.Close()
		i = 0
		TeacherStemSchoolEnterDates, err := db.Query("SELECT `StemSchoolEnterDate` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherStemSchoolEnterDatesList []string
		for TeacherStemSchoolEnterDates.Next() {
			var TeacherStemSchoolEnterDate string
			err = TeacherStemSchoolEnterDates.Scan(&TeacherStemSchoolEnterDate)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherStemSchoolEnterDatesList = append(teacherStemSchoolEnterDatesList, TeacherStemSchoolEnterDate+",")
			} else {
				teacherStemSchoolEnterDatesList = append(teacherStemSchoolEnterDatesList, TeacherStemSchoolEnterDate)
			}
		}
		defer TeacherStemSchoolEnterDates.Close()
		i = 0
		TeacherMobileWhatss, err := db.Query("SELECT `MobileWhats` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherMobileWhatssList []string
		for TeacherMobileWhatss.Next() {
			var TeacherMobileWhats string
			err = TeacherMobileWhatss.Scan(&TeacherMobileWhats)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherMobileWhatssList = append(teacherMobileWhatssList, TeacherMobileWhats+",")
			} else {
				teacherMobileWhatssList = append(teacherMobileWhatssList, TeacherMobileWhats)
			}
		}
		defer TeacherMobileWhatss.Close()
		i = 0
		TeacherMobileCallss, err := db.Query("SELECT `MobileCalls` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherMobileCallssList []string
		for TeacherMobileCallss.Next() {
			var TeacherMobileCalls string
			err = TeacherMobileCallss.Scan(&TeacherMobileCalls)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherMobileCallssList = append(teacherMobileCallssList, TeacherMobileCalls+",")
			} else {
				teacherMobileCallssList = append(teacherMobileCallssList, TeacherMobileCalls)
			}
		}
		defer TeacherMobileCallss.Close()
		i = 0
		TeacherEmails, err := db.Query("SELECT `Email` FROM `teachers` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		var teacherEmailsList []string
		for TeacherEmails.Next() {
			var TeacherEmail string
			err = TeacherEmails.Scan(&TeacherEmail)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherEmailsList = append(teacherEmailsList, TeacherEmail+",")
			} else {
				teacherEmailsList = append(teacherEmailsList, TeacherEmail)
			}
		}
		defer TeacherEmails.Close()
		// Get Teacher Id From Database
		i = 0
		TeacherIds, err := db.Query("SELECT `teachersId` FROM `teachers` WHERE 1")
		if err != nil {
			log.Println(err.Error())
		}
		var teacherIdsList []string
		for TeacherIds.Next() {
			var TeacherId string
			err = TeacherIds.Scan(&TeacherId)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != teacherCount {
				teacherIdsList = append(teacherIdsList, TeacherId+",")
			} else {
				teacherIdsList = append(teacherIdsList, TeacherId)
			}
		}
		defer TeacherIds.Close()

		tmplAllTeacher.Execute(w, struct {
			TeacherNameList                       []string
			TeacherGenderList                     []string
			TeacherJobList                        []string
			TeacherPreviousJobList                []string
			TeacherCadreList                      []string
			TeacherEducationLevelList             []string
			TeacherHiringDateList                 []string
			TeacherQualificationList              []string
			TeacherQualificationPlaceList         []string
			TeacherQualificationDateList          []string
			TeacherSpecializationList             []string
			TeacherOriginalQualificationGradeList []string
			TeacherHighQualificationList          []string
			TeacherHighQualificationPlaceList     []string
			TeacherHighQualificationDateList      []string
			TeacherHighQualificationGradeList     []string
			TeacherFinancialGradeList             []string
			TeacherFinancialGradeDateList         []string
			TeacherOriginSchoolList               []string
			TeacherOriginTeachingGovList          []string
			TeacherOriginCityList                 []string
			TeacherStemSchoolNameList             []string
			TeacherCurrentStemSchoolEnterDateList []string
			TeacherStemSchoolEnterDateList        []string
			TeacherMobileWhatsList                []string
			TeacherMobileCallsList                []string
			TeacherEmailList                      []string
			TeacherNationalIdList                 []string
			TeacherIds                            []string
		}{
			teacherNameList,
			teacherGenderList,
			teacherJobList,
			teacherCadreList,
			teacherPreviousJobList,
			teacherEducationLevelsList,
			teacherHiringDatesList,
			qualificationList,
			qualificationPlaceList,
			qualificationDateList,
			teacherSpecializationsList,
			teacherOriginalQualificationGradesList,
			teacherHighQualificationsList,
			teacherHighQualificationPlacesList,
			teacherHighQualificationDatesList,
			teacherHighQualificationGradesList,
			teacherFinancialGradesList,
			teacherFinancialGradeDatesList,
			teacherOriginSchoolsList,
			teacherOriginTeachingGovsList,
			teacherOriginCitysList,
			teacherStemSchoolNamesList,
			teacherCurrentStemSchoolEnterDatesList,
			teacherStemSchoolEnterDatesList,
			teacherMobileWhatssList,
			teacherMobileCallssList,
			teacherEmailsList,
			teacherNationalIdList,
			teacherIdsList,
		})
		return
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ShowAllVacations(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		db := Database.DBConnect()
		defer db.Close()
		VacationCount, err := db.Query("SELECT * FROM `addvacation` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer VacationCount.Close()
		var i int = 0
		var vacationCount int = 0
		for VacationCount.Next() {
			vacationCount++
		}
		Vacations, err := db.Query("SELECT `vacationname` FROM `addvacation` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer Vacations.Close()
		var vacationNames []string
		for Vacations.Next() {
			var vacation string
			err = Vacations.Scan(&vacation)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != vacationCount {
				vacationNames = append(vacationNames, vacation+",")
			} else {
				vacationNames = append(vacationNames, vacation)
			}
		}
		// Get VacationIds
		i = 0
		VacationIds, err := db.Query("SELECT `addvacationid` FROM `addvacation` WHERE 1")
		if err != nil {
			log.Println(err.Error())
		}
		var vacationIds []string
		for VacationIds.Next() {
			var vacationId string
			err = VacationIds.Scan(&vacationId)
			if err != nil {
				log.Println(err)
			}
			i++
			if i != vacationCount {
				vacationIds = append(vacationIds, vacationId+",")
			} else {
				vacationIds = append(vacationIds, vacationId)
			}
		}
		tmplAllVacation.Execute(w, struct {
			Sidebar       template.HTML
			VacationNames []string
			VacationIds   []string
		}{
			Controllers.CurrentUser.SideBar,
			vacationNames,
			vacationIds,
		})
		return
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ShowAllViolations(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission > 0 {
		db := Database.DBConnect()
		defer db.Close()
		ViolationCount, err := db.Query("SELECT * FROM `violations` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer ViolationCount.Close()
		var i int = 0
		var violationCount int = 0
		for ViolationCount.Next() {
			violationCount++
		}
		ViolationNames, err := db.Query("SELECT * FROM `violations` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer ViolationNames.Close()
		var violationNames []string
		var violationPoints []string
		for ViolationNames.Next() {
			var violationId int
			var violationName string
			var violationPoint int
			err = ViolationNames.Scan(&violationId, &violationName, &violationPoint)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(violationId, violationName, violationPoint)
			i++
			if i != violationCount {
				violationNames = append(violationNames, violationName+",")
				violationPoints = append(violationPoints, strconv.Itoa(violationPoint)+",")
			} else {
				violationNames = append(violationNames, violationName)
				violationPoints = append(violationPoints, strconv.Itoa(violationPoint))
			}
		}

		tmplAllViolation.Execute(w, struct {
			Sidebar         template.HTML
			ViolationNames  []string
			ViolationPoints []string
		}{
			Controllers.CurrentUser.SideBar,
			violationNames,
			violationPoints,
		})
		return
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ShowAllClasses(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()

		defer db.Close()
		ClassCount, err := db.Query("SELECT * FROM `Class` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer ClassCount.Close()

		var classCount int = 0
		for ClassCount.Next() {
			classCount++
		}
		ClassList, err := db.Query("SELECT `ClassName` FROM `Class` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer ClassList.Close()
		var classList []string
		var i int = 0
		for ClassList.Next() {
			var class string
			err = ClassList.Scan(&class)
			if err != nil {
				fmt.Println(err)
			}

			i++
			if i != classCount {
				classList = append(classList, class+",")
			} else {
				classList = append(classList, class)
			}
		}
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ShowAllUniversities(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()

		defer db.Close()
		UniCount, err := db.Query("SELECT * FROM `University` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniCount.Close()

		var uniCount int = 0
		for UniCount.Next() {
			uniCount++
		}
		UniList, err := db.Query("SELECT `UniversityName` FROM `University` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniList.Close()
		var uniList []string
		var i int = 0
		for UniList.Next() {
			var university string
			err = UniList.Scan(&university)
			if err != nil {
				fmt.Println(err)
			}

			i++
			if i != uniCount {
				uniList = append(uniList, university+",")
			} else {
				uniList = append(uniList, university)
			}
		}

		UniCities, err := db.Query("SELECT `CIty` FROM `University` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer UniCities.Close()
		var uniCities []string
		i = 0
		for UniCities.Next() {
			var university string
			err = UniCities.Scan(&university)
			if err != nil {
				fmt.Println(err)
			}

			i++
			if i != uniCount {
				uniCities = append(uniCities, university+",")
			} else {
				uniCities = append(uniCities, university)
			}
		}

		tmplAllUniversities.Execute(w, struct {
			UniList   []string
			UniCities []string
		}{
			uniList,
			uniCities,
		})
		return
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

func ShowAllScholarShips(w http.ResponseWriter, r *http.Request) {
	if Controllers.CurrentUser.Permission >= 6 {
		db := Database.DBConnect()
		defer db.Close()

		ScholarShipCount, err := db.Query("SELECT * FROM `AssignScholar` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer ScholarShipCount.Close()

		var scholarShipCount int = 0
		for ScholarShipCount.Next() {
			scholarShipCount++
		}

		StudentIds, err := db.Query("SELECT `StdId` FROM `AssignScholar` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer StudentIds.Close()

		var studentIds []string
		var i int = 0

		for StudentIds.Next() {
			var student string
			err = StudentIds.Scan(&student)
			if err != nil {
				fmt.Println(err)
			}

			i++
			if i != scholarShipCount {
				studentIds = append(studentIds, student+",")
			} else {
				studentIds = append(studentIds, student)
			}
		}

		// Iterate Over stdudentIds
		var studentNames []string
		i = 0
		for _, studentId := range studentIds {
			StudentName, err := db.Query("SELECT `English Name` FROM `student` WHERE `StdId` = ?", studentId)
			if err != nil {
				log.Println(err.Error())
			}
			i++
			if i != scholarShipCount {

				for StudentName.Next() {
					var student string
					err = StudentName.Scan(&student)
					if err != nil {
						fmt.Println(err)
					}
					studentNames = append(studentNames, student+",")
				}
			} else {
				for StudentName.Next() {
					var student string
					err = StudentName.Scan(&student)
					if err != nil {
						fmt.Println(err)
					}
					studentNames = append(studentNames, student)
				}
			}
		}

		ScholarIds, err := db.Query("SELECT `ScholarshipId` FROM `AssignScholar` WHERE 1")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer ScholarIds.Close()

		var scholarIds []string
		i = 0

		for ScholarIds.Next() {
			var scholar string
			err = ScholarIds.Scan(&scholar)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != scholarShipCount {
				scholarIds = append(scholarIds, scholar+",")
			} else {
				scholarIds = append(scholarIds, scholar)
			}
		}
		i = 0
		//Iterate Over Scholar Ids
		var scholarNames []string
		for _, scholarId := range scholarIds {
			ScholarName, err := db.Query("SELECT `universityname` FROM `scholarship` WHERE `scolarshipid` = ?", scholarId)
			if err != nil {
				log.Println(err.Error())
			}
			i++
			if i != scholarShipCount {
				for ScholarName.Next() {
					var scholarName string
					err = ScholarName.Scan(&scholarName)
					if err != nil {
						fmt.Println(err)
					}
					scholarNames = append(scholarNames, scholarName+",")
				}
			} else {
				for ScholarName.Next() {
					var scholarName string
					err = ScholarName.Scan(&scholarName)
					if err != nil {
						fmt.Println(err)
					}
					scholarNames = append(scholarNames, scholarName)
				}
			}
		}
		// Get ScholarShipKinds
		ScholarShipKinds, err := db.Query("SELECT `ScholarShipKind` FROM `AssignScholar` WHERE 1")
		if err != nil {
			log.Println(err.Error())
		}
		defer ScholarShipKinds.Close()
		var scholarShipKinds []string
		i = 0
		for ScholarShipKinds.Next() {
			var scholarShipKind string
			err = ScholarShipKinds.Scan(&scholarShipKind)
			if err != nil {
				fmt.Println(err)
			}
			i++
			if i != scholarShipCount {
				scholarShipKinds = append(scholarShipKinds, scholarShipKind+",")
			} else {
				scholarShipKinds = append(scholarShipKinds, scholarShipKind)
			}
		}
		tmplAllScholarShips.Execute(w, struct {
			StudentNames     []string
			UniNames         []string
			ScholarShipKinds []string
		}{
			studentNames,
			scholarNames,
			scholarShipKinds,
		})
		return
	} else if Controllers.CurrentUser.Permission == 0 {
		FormBackend.ExecuteError(w, r, false)
	} else {
		FormBackend.ExecuteError(w, r, true)
	}
}

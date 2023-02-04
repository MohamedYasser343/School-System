package Types

type StudentDetail struct {
	Name             string  `json:"Name"`
	NameE            string  `json:"Name_E"`
	Email            string  `json:"Email"`
	NationalId       string  `json:"National_Id"`
	Id               string  `json:"Id"`
	Gender           string  `json:"Gender"`
	Religion         string  `json:"Religion"`
	State            string  `json:"State"`
	Address          string  `json:"Address"`
	SchoolType       string  `json:"School_Type"`
	SchoolCity       string  `json:"School_City"`
	StudentNumber    string  `json:"Student_Number"`
	StudentNumber2   string  `json:"Student_Number_2"`
	FatherName       string  `json:"Father_Name"`
	FatherNameE      string  `json:"Father_Name_E"`
	FatherNationalId string  `json:"Father_National_Id"`
	FatherJob        string  `json:"Father_Job"`
	FatherAddress    string  `json:"Father_Address"`
	FatherMobile     string  `json:"Father_Mobile"`
	FatherMobile2    string  `json:"Father_Mobile2"`
	MotherName       string  `json:"Mother_Name"`
	MotherNameE      string  `json:"Mother_Name_E"`
	MotherNationalId string  `json:"Mother_National_Id"`
	MotherJob        string  `json:"Mother_Job"`
	MotherAddress    string  `json:"Mother_Address"`
	MotherMobile     string  `json:"Mother_Mobile"`
	MotherMobile2    string  `json:"Mother_Mobile2"`
	Violations1      float64 `json:"Violations1"`
	Violations2      float64 `json:"Violations2"`
	Violations3      float64 `json:"Violations3"`
	Violations4      float64 `json:"Violations4"`
	BirthDay         string  `json:"Birth_Day"`
	BirthCity        string  `json:"Birth_City"`
	ParentName       string  `json:"Parent_Name"`
	ParentNameE      string  `json:"Parent_Name_E"`
	ParentNationalId string  `json:"Parent_National_Id"`
	ParentJob        string  `json:"Parent_Job"`
	ParentAddress    string  `json:"Parent_Address"`
	ParentMobile     string  `json:"Parent_Mobile"`
	ParentMobile2    string  `json:"Parent_Mobile2"`
}

type NewStudent struct {
	Name           string `json:"Name"`
	NameE          string `json:"NameE"`
	Grade          string `json:"Grade"`
	Permission     string `json:"Permission"`
	Specialization string `json:"Specialization"`
	SecondLang     string `json:"SecondLang"`
	Religion       string `json:"Religion"`
	PortfolioNo    string `json:"PortfolioNo"`
	NationalId     string `json:"NationalId"`
	Email          string `json:"Email"`
	OfficialEmail  string `json:"OfficialEmail"`
	Code           string `json:"Code"`
	Class          string `json:"Class"`
	UserName       string `json:"UserName"`
	Password       string `json:"Password"`
}

type NewTeacher struct {
	TeacherCode              string `json:"TeacherCode"`
	Name                     string `json:"Name"`
	Cadr                     string `json:"Cadr"`
	Id                       string `json:"Id"`
	LastJob                  string `json:"LastJob"`
	STEMJob                  string `json:"STEMJob"`
	Stage                    string `json:"Stage"`
	HiringDate               string `json:"HiringDate"`
	OrgQualification         string `json:"OrgQualification"`
	PlaceOrgQualification    string `json:"PlaceOrgQualification"`
	QualificationDate        string `json:"QualificationDate"`
	Specialization           string `json:"Specialization"`
	OrgQualificationGrade    string `json:"OrgQualificationGrade"`
	HigherQualification      string `json:"HigherQualification"`
	PlaceHigherQualification string `json:"PlaceHigherQualification"`
	HighQualificationDate    string `json:"HighQualificationDate"`
	HigherQualificationGrade string `json:"HigherQualificationGrade"`
	CurrentFinacialGrade     string `json:"CurrentFinacialGrade"`
	ItsDate                  string `json:"ItsDate"`
	OrgSchool                string `json:"OrgSchool"`
	OrgAdministration        string `json:"OrgAdministration"`
	OrgGovernment            string `json:"Org Government"`
	NameofSTEMSchool         string `json:"NameofSTEMSchool"`
	DateJoinSTEMNow          string `json:"DateJoinSTEMNow"`
	DateJoinSTEM             string `json:"DateJoinSTEM"`
	NumberWhats              string `json:"NumberWhats"`
	NumberPhone              string `json:"NumberPhone"`
	EmailSTEM                string `json:"EmailSTEM"`
	NationalID               string `json:"NationalID"`
	Gender                   string `json:"Gender"`
}

type NewClass struct {
	ClassName string `json:"ClassName"`
	Grade     string `json:"Grade"`
}

type NewVacation struct {
	Name string `json:"Name"`
}

type InputForm struct {
	Radio string `json:"Radio"`
	Id    string `json:"Id"`
}

type AssignBussing struct {
	BusId string `json:"BusId"`
	StdId string `json:"StdId"`
}

type NewBus struct {
	BusLine     string `json:"BusLine"`
	DriverName  string `json:"DriverName"`
	DriverPhone string `json:"DriverPhone"`
	StartPoint  string `json:"StartPoint"`
	StartTime   string `json:"StartTime"`
	BusChairs   string `json:"BusChairs"`
	EndPoint    string `json:"EndPoint"`
	EndTime     string `json:"EndTime"`
}

type RecordVacationForm struct {
	TeacherName  string `json:"TeacherName"`
	VacationType string `json:"VacationType"`
	StartDate    string `json:"StartDate"`
	EndDate      string `json:"EndDate"`
}

type AddStdToCapstone struct {
	CapstoneNumber string `json:"CapstoneNumber"`
	Student1       string `json:"Student1"`
	Student2       string `json:"Student2"`
	Student3       string `json:"Student3"`
	Student4       string `json:"Student4"`
	Student5       string `json:"Student5"`
	Student6       string `json:"Student6"`
}

type NewViolation struct {
	ViolationKind   string `json:"ViolationKind"`
	ViolationPoints string `json:"ViolationPoints"`
}

type NewTrip struct {
	TripName          string `json:"TripName"`
	TripDate          string `json:"TripDate"`
	TripPlace         string `json:"TripPlace"`
	ChairsNumber      string `json:"ChairsNumber"`
	TripKind          string `json:"TripKind"`
	GeneralSupervisor string `json:"GeneralSupervisor"`
	Supervisor1       string `json:"Supervisor1"`
	Supervisor2       string `json:"Supervisor2"`
	Supervisor3       string `json:"Supervisor3"`
	Supervisor4       string `json:"Supervisor4"`
	Supervisor5       string `json:"Supervisor5"`
	LastDate          string `json:"LastDate"`
}

type NewSupervisor struct {
	SupervisorName        string `json:"SupervisorName"`
	SupervisorNumber      string `json:"SupervisorNumber"`
	SupervisorId          string `json:"SupervisorId"`
	SupervisorAddress     string `json:"SupervisorAddress"`
	GraduationCertificate string `json:"GraduationCertificate"`
}

type NewRecordStudentUniversity struct {
	StudentName      string `json:"StudentName"`
	UniversityName   string `json:"UniversityName"`
	SchoolarshipKind string `json:"SchoolarshipKind"`
}
type NewRecordStudentViolation struct {
	StudentName   string `json:"StudentName"`
	ViolationKind string `json:"ViolationKind"`
	ViolationDate string `json:"ViolationDate"`
}
type NewRecordTeacherViolation struct {
	TeacherName   string `json:"TeacherName"`
	ViolationKind string `json:"ViolationKind"`
	ViolationDate string `json:"ViolationDate"`
}

type NewRooming struct {
	RoomNumber  string `json:"RoomNumber"`
	StudentName string `json:"StudentName"`
}

type NewRoom struct {
	RoomNumber   string `json:"RoomNumber"`
	RoomCapacity string `json:"RoomCapacity"`
	FloorNo      string `json:"FloorNo"`
}

type NewScholarShip struct {
	UniversityName string `json:"UniversityName"`
	CollageName    string `json:"CollageName"`
	Notes          string `json:"Notes"`
}

type NewSchool struct {
	SchoolName    string `json:"SchoolName"`
	DateCreated   string `json:"DateCreated"`
	PrincipalName string `json:"PrincipalName"`
	DeputyName    string `json:"DeputyName"`
}

type NewUniversity struct {
	UniversityName string `json:"UniversityName"`
	CollegeName    string `json:"CollegeName"`
	CityName       string `json:"CityName"`
}

type NewEnterComp struct {
	CompetitionName string `json:"CompetitionName"`
	StudentName     string `json:"StudentName"`
}

type AssignScholarShip struct {
	StudentName     string `json:"StudentName"`
	UniversityName  string `json:"UniversityName"`
	ScholarShipKind string `json:"ScholarShipKind"`
}

type AssignClass struct {
	StudentName string `json:"StudentName"`
	ClassName   string `json:"ClassName"`
}

package Models

import "html/template"

type User struct {
	Id         uint
	Name       string
	Email      string
	Password   string
	Permission int
	SideBar    template.HTML
	Code       string
}

type Student struct {
	StdId          string
	Code           string
	EnglishName    string
	ArabicName     string
	Grade          string
	Specialization string
	SecondLang     string
	Religion       string
	Gender         string
	Capstone       string
	NationalId     string
	Class          string
	PersonalEmail  string
	OfficialEmail  string
}

type RemoveItem struct {
	Table      string
	ColumnName string
	RowData    string
}

type Data struct {
	TableName string
	Data      []string
}

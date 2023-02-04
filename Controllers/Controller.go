package Controllers

import (
	"Snap/AbstractFunctions"
	"Snap/Database"
	"Snap/Models"
	"Snap/Sidebars"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mavihq/persian"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = ""

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user Models.User
	fmt.Println(data)
	Database.DB.Where("email = ?", data["Email"]).First(&user)
	Database.DB.Where("password = ?", data["Password"]).First(&user)
	if user.Email != data["Email"] || data["Email"] == "" {
		c.Status(fiber.StatusNotFound)
		err := Logout(c)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "User Not Found",
		})
	} else if user.Password != data["Password"] {
		c.Status(fiber.StatusBadRequest)
		err := Logout(c)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 96).Unix(), // Valid For Four Day.
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Couldn't Log In",
		})
	}
	if user.Email == data["Email"] && user.Password == data["Password"] {
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 96), // Valid For 4 Days.
			HTTPOnly: true,
		}
		c.Cookie(&cookie)
		if &user != nil {
			return c.JSON(fiber.Map{
				"message":    "Success",
				"jwt":        token,
				"permission": user.Permission,
			})
		}
		if user.Permission == 1 {
			return c.JSON(fiber.Map{
				"message":    "Success",
				"permission": "1",
			})
		} else if user.Permission == 2 {
			return c.JSON(fiber.Map{
				"message":    "Success",
				"permission": "2",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Success",
		})
	}
	return c.JSON(fiber.Map{
		"message": "User Not Found",
	})
}

var CurrentUser *Models.User

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		var user Models.User
		user.Permission = 0
		CurrentUser = &user
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var user Models.User

	claims := token.Claims.(*jwt.StandardClaims)
	Database.DB.Where("email = ?", claims.Issuer).First(&user)
	fmt.Println(user.Permission)
	if user.Permission == 6 {
		user.SideBar = Sidebars.AdminSidebar
	} else if user.Permission == 4 {
		user.SideBar = Sidebars.TeacherAffairsSidebar
	}
	CurrentUser = &user
	fmt.Println(user.Code)
	db := Database.DBConnect()
	query := Database.SelectFromDB("*", "student", "Code", user.Code, db)
	// make student model
	var student Models.Student

	for query.Next() {
		err := query.Scan(&student.StdId, &student.Code, &student.EnglishName, &student.ArabicName, &student.Grade, &student.Specialization, &student.SecondLang, &student.Religion, &student.Gender, &student.Capstone, &student.NationalId, &student.Class, &student.PersonalEmail, &student.OfficialEmail)
		if err != nil {
			log.Println(err.Error())
		}
	}
	query.Scan(&student)
	fmt.Println(student)
	NationalId := persian.ToEnglishDigits(student.NationalId)
	chars := []rune(NationalId)
	fmt.Println(chars)
	if NationalId != "" {
		birthCity := "N/A"
		_ = birthCity
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
		BirthDay := birthDay + "/" + birthMonth + "/" + birthYear
		return c.JSON(fiber.Map{
			"user":           user.Name,
			"studentdetails": student,
			"birthday":       BirthDay,
			"city":           birthCity,
		})
	} else {
		return c.JSON(fiber.Map{
			"message": user.Permission,
		})
	}
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logged Out",
	})
}

func GetViolations(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	var user Models.User
	CurrentUser = &user
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	Database.DB.Where("email = ?", claims.Issuer).First(&user)
	db := Database.DBConnect()

	StudentId := Database.SelectFromDB("StdId", "student", "Code", user.Code, db)

	var id int
	for StudentId.Next() {
		err := StudentId.Scan(&id)
		if err != nil {
			log.Println(err.Error())
		}
	}

	query := Database.SelectFromDB("violationid", "recordstdviolation", "studentid", strconv.Itoa(id), db)
	var violations []string
	var violationNumbers []int
	for query.Next() {
		var ViolationId string
		query.Scan(&ViolationId)
		var ViolationName string
		GetViolationName := Database.SelectFromDB("violationName", "violations", "violationid", ViolationId+" LIMIT 1", db)
		for GetViolationName.Next() {
			err := GetViolationName.Scan(&ViolationName)
			if err != nil {
				log.Println(err.Error())
			}
			violations = append(violations, ViolationName)
		}
	}
	violations = RemoveDuplicates(violations)
	fmt.Println(violations)
	for _, i := range violations {
		fmt.Println(i)
		ViolationId := Database.SelectFromDB("violationid", "violations", "violationName", i, db)
		var violationId string
		for ViolationId.Next() {
			err := ViolationId.Scan(&violationId)
			if err != nil {
				log.Println(err.Error())
			}
		}
		//violation := Database.SelectFromDB("*", "recordstduni", "violationid", violationId, db)
		violation, err := db.Query("SELECT * FROM `recordstdviolation` WHERE `violationid` = ? && `studentid` = ?", violationId, id)

		if err != nil {
			log.Println(err.Error())
		}

		var index int
		for violation.Next() {
			index++
		}
		violationNumbers = append(violationNumbers, index)
	}
	return c.JSON(fiber.Map{
		"violations":       violations,
		"violationNumbers": violationNumbers,
	})
}

func RemoveDuplicates(Array []string) []string {

	check := make(map[string]int)
	d := append(Array, Array...)
	res := make([]string, 0)
	for _, val := range d {
		check[val] = 1
	}

	for Element, _ := range check {
		res = append(res, Element)
	}

	return res
}

func RedirectToDashboard(w http.ResponseWriter, r *http.Request) {
	if CurrentUser.Permission == 0 {
		http.Redirect(w, r, "https://snap.hopto.org", http.StatusSeeOther)
	} else if CurrentUser.Permission > 0 {
		http.Redirect(w, r, "https://snap.hopto.org/GoDashboard", http.StatusSeeOther)
	}
}

func WriteDataFromApp(c *fiber.Ctx) error {
	User(c)
	if CurrentUser.Permission >= 6 {
		var data map[string][]string
		if err := c.BodyParser(&data); err != nil {
			log.Println(err.Error())
			return err
		}
		fmt.Println(data)
		query := "INSERT INTO `" + data["Table"][0] + "` VALUES(" + "NULL" + ", "

		// loop over ColumnData
		for i := range data["ColumnData"] {
			if i != len(data["ColumnData"])-1 {
				query += "\"" + data["ColumnData"][i] + "\", "
			} else {
				query += "\"" + data["ColumnData"][i] + "\");"
			}
		}
		fmt.Println(query)
		db := Database.DBConnect()
		insert, err := db.Query(query)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		defer insert.Close()

		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

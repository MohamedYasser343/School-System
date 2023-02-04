package AbstractFunctions

import (
	"database/sql"
	"fmt"
	"github.com/mavihq/persian"
	"log"
	"net"
)

func ReplaceNationalId() {
	db, err := sql.Open("mysql", "snap:Snapsnap@2@tcp(92.205.60.182:3306)/snap")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	insert, err := db.Query("SELECT stdnationalid FROM `students`")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer insert.Close()
	var stdNationalId = "0"
	var Ids []string
	for insert.Next() {
		err := insert.Scan(&stdNationalId)
		if err != nil {
			log.Fatal(err.Error())
		}
		Ids = append(Ids, stdNationalId)
	}
	for _, Id := range Ids {
		finalId := persian.ToEnglishDigits(Id)
		fmt.Println(finalId)
		insert2, err := db.Query("UPDATE students SET stdNationalId = ? WHERE stdNationalId = ?", finalId, Id)
		defer insert2.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func ResolveHostIp() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()

			fmt.Println("Resolved Host IP: " + ip)
			return ip
		}
	}
	return ""
}

func ReturnBirthCity(chars []rune) string {
	switch string(chars[7]) + string(chars[8]) {
	case "01":
		return "Cairo"
	case "02":
		return "Alexandria"
	case "12":
		return "Dakhlia"
	case "13":
		return "Sharkia"
	case "14":
		return "Kaliobia"
	case "15":
		return "Kafr El Sheikh"
	case "16":
		return "Al 8arbya"
	case "17":
		return "Monofia"
	case "18":
		return "Behira"
	case "19":
		return "Ismalia"
	case "21":
		return "Giza"
	case "22":
		return "Beni Swaif"
	case "24":
		return "Al Menia"
	case "25":
		return "Asyuit"
	case "26":
		return "Sohag"
	case "27":
		return "Qna"
	case "28":
		return "Aswan"
	case "29":
		return "Luxor"
	case "33":
		return "Matro7"
	case "11":
		return "Domit"
	case "88":
		return "Main Central"
	case "03":
		return "Bor Saed"
	}
	return "Couldn't Resolve BirthCity."
}

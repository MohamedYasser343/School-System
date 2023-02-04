# School-System
Freelancing Project for my STEM SChool

This system was programmed by me and another friends help me in building and programming it.

---

# Table Of Content
This repo consists of:
1. snap.sql (database code)
2. backend files (written in go-lang)
3. frontend files (paid template was used and some js coding for fetch apis coming from backend)

---

# How To Install

install the project:

```
git clone https://github.com/MohamedYasser343/School-System.git
```

---

# How To Run

1. insert data in **snap.sql** using:

```sql
INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
```

2. config your database in **Database/Connection.go**:

```go
var DBAddress string = "<user>:<password>@tcp(<ip>:3306)/<database_name>"

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("<user>:<password>@tcp(<ip>:3306)/<database_name>"), &gorm.Config{})
	if err != nil {
		panic("Couldn't Connect To The Database.")
	}
	DB = connection
	connection.AutoMigrate(&Models.User{})
}
```

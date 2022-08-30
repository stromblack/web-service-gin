package database

import (
	"fmt"
	"synergy/web-service-gin/common"
	"synergy/web-service-gin/models"
	"time"

	_ "github.com/lib/pq"
)

func GetUsers() []models.User {
	db, _ := InitDb()
	defer db.Close()
	rows, err := db.Query("SELECT uid,username,password,email,created FROM T_USER_INFO")
	common.CheckErr(err)
	userList := make([]models.User, 0)
	// loop
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UserID, &user.UserName, &user.UserPassword, &user.Email, &user.Created)
		// append
		userList = append(userList, user)
	}
	fmt.Printf("# Query: %v \n", len(userList))
	return userList
}

func InsertUser(user models.User) models.User {
	db, _ := InitDb()
	// close db after return type
	defer db.Close()
	var LastInsertId int
	dt := time.Now()
	err := db.QueryRow("INSERT INTO T_USER_INFO(username,password,email,created) VALUES($1, $2, $3, $4) RETURNING uid;", user.UserName, user.UserPassword, user.Email, dt).Scan(&LastInsertId)
	common.CheckErr(err)
	fmt.Printf("# Inserting = %v", LastInsertId)
	user.UserID = int64(LastInsertId)
	return user
}

func UpdateUser(user models.User) bool {
	db, _ := InitDb()
	defer db.Close()
	fmt.Println("# Updating")
	stmt, err := db.Prepare("update T_USER_INFO set username=$1, email=$2 where uid=$3")
	common.CheckErr(err)
	res, err := stmt.Exec(user.UserName, user.Email, user.UserID)
	common.CheckErr(err)
	affect, err := res.RowsAffected()
	common.CheckErr(err)
	fmt.Println(affect, "rows changed")
	return affect > 0
}

func DeleteUser(userid int) bool {
	db, _ := InitDb()
	defer db.Close()
	fmt.Println("# Deleteing")
	stmt, err := db.Prepare("DELETE FROM T_USER_INFO WHERE uid = $1")
	common.CheckErr(err)
	res, err := stmt.Exec(userid)
	common.CheckErr(err)
	affect, err := res.RowsAffected()
	common.CheckErr(err)
	return affect > 0
}

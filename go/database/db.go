package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type logCreate struct {
	List []logCreateRow `json:"list"`
}
type logCreateRow struct {
	Createid int    `json:"create_id"`
	Owner    string `json:"owner"`
	Path     string `json:"path"`
	Date     string `json:"date"`
}

type logMove struct {
	List []logMoveRow `json:"list"`
}
type logMoveRow struct {
	Moveid  int    `json:"moveid"`
	Owner   string `json:"owner"`
	Origin  string `json:"origin"`
	Destiny string `json:"destiny"`
	Date    string `json:"date"`
}

type logDelete struct {
	List []logDeleteRow `json:"list"`
}

type logDeleteRow struct {
	Delid int    `json:"del_id"`
	Owner string `json:"owner"`
	Path  string `json:"path"`
	Date  string `json:"date"`
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "dev"
)

//Connect database test
func Connect() *sql.DB {
	/*dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWORD, DB_NAME)*/
	dbinfo := fmt.Sprintf("host=organization_db user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func Disconnect(db *sql.DB) string {
	err := db.Close()
	var res string
	if err != nil {
		res = "error al desconectar"
		panic(err)
	} else {
		res = "base de adtos deconectada"
	}
	return res
}

func Move(db *sql.DB, owner string, origen string, destino string) logMoveRow {
	var lastInsertId int
	var t logMoveRow
	t.Owner = owner
	t.Origin = origen
	t.Destiny = destino
	t.Date = time.Now().Format("01-02-2006 15:04:05")
	fmt.Println("# Inserting values")
	err = db.QueryRow("INSERT INTO logMovements(owner,origen,destiny,date) VALUES($1,$2,$3,$4) returning move_id;", owner, origen, destino, t.Date).Scan(&lastInsertId)
	checkErr(err)
	res := "last inserted id =" + strconv.Itoa(lastInsertId)
	fmt.Println(res)
	t.Moveid = lastInsertId
	return t
}

func LogMove(db *sql.DB) logMove {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM logMovements")
	checkErr(err)
	var aux logMove

	for rows.Next() {
		var temp logMoveRow
		err = rows.Scan(&temp.Moveid, &temp.Owner, &temp.Origin, &temp.Destiny, &temp.Date)
		checkErr(err)
		aux.List = append(aux.List, temp)
	}
	return aux
}

func Create(db *sql.DB, owner string, path string) logCreateRow {
	var t logCreateRow
	t.Date = time.Now().Format("01-02-2006 15:04:05")
	t.Owner = owner
	t.Path = path
	fmt.Println("# Inserting values")
	err = db.QueryRow("INSERT INTO logFolderCreations(owner,path,date) VALUES($1,$2,$3) returning create_id;", owner, path, t.Date).Scan(&t.Createid)
	checkErr(err)
	return t
}

func LogCreate(db *sql.DB) logCreate {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM logFolderCreations")
	checkErr(err)

	var aux logCreate

	for rows.Next() {
		var temp logCreateRow
		err = rows.Scan(&temp.Createid, &temp.Owner, &temp.Path, &temp.Date)
		checkErr(err)
		aux.List = append(aux.List, temp)
	}

	return aux
}

func Delete(db *sql.DB, owner string, path string) logDeleteRow {
	var t logDeleteRow
	var lastInsertId int
	t.Date = time.Now().Format("01-02-2006 15:04:05")
	t.Owner = owner
	t.Path = path
	fmt.Println("# Inserting values")
	err = db.QueryRow("INSERT INTO logDeletes(owner,path,date) VALUES($1,$2,$3) returning del_id;", owner, path, t.Date).Scan(&lastInsertId)
	checkErr(err)
	res := "last inserted id =" + strconv.Itoa(lastInsertId)
	fmt.Println(res)
	t.Delid = lastInsertId
	return t
}

func LogDelete(db *sql.DB) logDelete {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM logDeletes")
	checkErr(err)

	var aux logDelete

	for rows.Next() {
		var temp logDeleteRow
		err = rows.Scan(&temp.Delid, &temp.Owner, &temp.Path, &temp.Date)
		checkErr(err)
		aux.List = append(aux.List, temp)
	}
	return aux
}

/*func UpdateValues() string {
	fmt.Println("# Updating")
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("astaxieupdate", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	out := strconv.FormatInt(affect, 10) + "rows changed"

	fmt.Println(affect, "rows changed")
	return out

}

func Query() string {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	aux := fmt.Sprintf("uid | username | department | created \n")

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		aux += fmt.Sprintf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}
	return aux
}

func Delete() string {
	fmt.Println("# Deleting")
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	out := strconv.FormatInt(affect, 10) + "rows changed"
	return out
}*/

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERRRROROROROROROOR")
		panic(err)
	}
}

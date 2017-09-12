package main

import (
	"database/sql"
	"fmt"
)

func main() {

	db, err := sql.Open("mysql", "root:@/go?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("jimersylee", "dev", "2017-09-05")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("nameUpdate", id)
	checkErr(err)
	affect,err:=res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)


	//查询数据
	rows,err:=db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next(){
		var uid int
		var username string
		var department string
		var created string
		err=rows.Scan(&uid,&username,&department,&created)
		fmt.Println(uid,username,department,created)
	}

	//删除数据
	stmt, err=db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res,err=stmt.Exec(id)
	checkErr(err)

	affect,err=res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

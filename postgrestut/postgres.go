package postgrestut

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@172.17.0.2:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection secured")
	insert(db, "vivek")
	Update(db, "manav", "manav sinha")
	Delete(db, "vivek")
	readbyId(db, 3)
	read(db)
}

func read(db *sql.DB) {
	res, err := db.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		var name string
		var id int
		res.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func insert(db *sql.DB, name string) {
	// _, err := db.Query(`insert into student(name) values('?')`, name)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	stmt, err := db.Prepare(fmt.Sprintf(`insert into student(name) values('%s')`, name))
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

}

func Update(db *sql.DB, oldName string, newName string) {
	query := fmt.Sprintf(`update student set name = '%s' where name='%s';`, newName, oldName)
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

}

func Delete(db *sql.DB, name string) {

	// _,err := db.Query(`delete from student where name = '?';`, name)
	// if err!=nil{
	// 	log.Fatal()
	// }
	query := fmt.Sprintf(`delete from student where name = '%s';`, name)
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func readbyId(db *sql.DB, id int) {
	row, err := db.Query("select name from student where id =3")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var name string
		row.Scan(&name)
		fmt.Println(name)
	}
}

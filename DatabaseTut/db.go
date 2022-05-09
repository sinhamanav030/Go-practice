package databasetut

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db  *sql.DB
	ctx context.Context
	err error
)

func init() {
	// db, err = sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/company?charset=utf8")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Connection Success")
	// defer db.Close()
	// GetInfo()
}

func GetInfo() {
	rows, err := db.Query(`select * from employee;`)
	if err != nil {
		fmt.Println(err)
	}
	var s string
	var eid, cno int
	for rows.Next() {
		err = rows.Scan(&eid, &s, &cno)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(eid, s, cno)
	}

}

func Create() {
	stmt, err := db.Prepare(`create table customer(name varchar(20));`)

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

func Insert() {
	stmt, err := db.Prepare(`insert into customer values("manav")`)
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

func Update() {
	stmt, err := db.Prepare(`update customer set name = "manav sinha" where name="manav";`)
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

func Read() {
	rows, err := db.Query(`select * from customer;`)
	if err != nil {
		fmt.Println(err)
	}
	var s string
	for rows.Next() {
		err = rows.Scan(&s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}

}

func Delete() {
	stmt, err := db.Prepare(`delete from customer where name = "manav sinha";`)
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

//mongodb -username : manav03
//paswword : 7bRpcqdQtiKaUvZb

type User struct {
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
	Id     primitive.ObjectID `json:"id" bson:"_id"`
}

var collection *mongo.Collection

func ConnectMdb() {
	const conn = `mongodb+srv://manav03:7bRpcqdQtiKaUvZb@gopherslab.uyaxi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority`
	const dbname = "company"
	const colname = "user"

	clientOption := options.Client().ApplyURI(conn)
	cl, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatalln(err)
	}
	collection = cl.Database(dbname).Collection(colname)

	user := User{
		Name:   "Manav",
		Gender: "male",
		Age:    21,
		Id:     primitive.NewObjectID(),
	}
	in, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(in.InsertedID)
}

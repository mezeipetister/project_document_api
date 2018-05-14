package example

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	"github.com/soveran/redisurl"

	"gopkg.in/mgo.v2"
)

// User struct
type User struct {
	username string
	fname    string
	lname    string
	email    string
	password string
}

// Create new user
func Create(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("DEMO").C("users")

	err = collection.Insert(&User{
		"mezeipetister",
		"Peter",
		"Mezei",
		"mezeipetister@gmail.com",
		"HelloBello",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(rw, "Ok")
}

func Login(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

// Demo section

var counter int

// Create new user
func redisDemo(c echo.Context) {
	// Now we connect to the Redis server
	conn, err := redisurl.ConnectToURL("http://localhost:6379")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	counter++
	conn.Do("lpush", "numbers", counter)
}

type Person struct {
	Fname string
	Lname string
}

// Mongo Demo
func Mongo(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("DEMO").C("A")

	err = collection.Insert(&Person{"Peti", "Mezei"},
		&Person{"Krisztina", "Mezeiné Bégányi"})

	if err != nil {
		log.Fatal(err)
	}

	// return c.NoContent(http.StatusCreated)
}

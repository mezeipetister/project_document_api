package user

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/soveran/redisurl"

	"gopkg.in/mgo.v2"
)

var counter int

// Create new user
func Create(c echo.Context) error {
	// Now we connect to the Redis server
	conn, err := redisurl.ConnectToURL("http://localhost:6379")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	counter++
	conn.Do("lpush", "numbers", counter)

	return c.NoContent(http.StatusCreated)
}

type Person struct {
	Fname string
	Lname string
}

// Mongo Demo
func Mongo(c echo.Context) error {

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

	return c.NoContent(http.StatusCreated)
}

// Login controller
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

package main

import (
	"strconv"
	"sync"

	"api/database"
	"api/routes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gofiber/fiber/v2"
)

//xlsx file structure model
type Temp struct {
	ID        int    `json:id`
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
	Gender    string `json:gender`
	Country   string `json:country`
	Age       int    `json:age`
	Date      string `json:date`
}

//inserting 500 data into temps table
func go1(wg *sync.WaitGroup) {

	db := database.DBconnection()
	var data Temp

	defer wg.Done()
	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)

	}
	rows, err := f.GetRows("Sheet1")
	limit := len(rows)

	for i := 1; i < limit/4; i++ {
		for j := 0; j < 7; j++ {
			if j == 0 {
				i, err := strconv.Atoi(rows[i][0])
				if err != nil {
					fmt.Print(err)
				}
				data.ID = i

			}
			if j == 1 {
				data.Firstname = rows[i][1]
			}
			if j == 2 {
				data.Lastname = rows[i][2]
			}
			if j == 3 {
				data.Gender = rows[i][4]
			}
			if j == 4 {
				data.Country = rows[i][4]
			}
			if j == 5 {
				i, err := strconv.Atoi(rows[i][5])
				if err != nil {
					fmt.Print(err)
				}
				data.Age = i
			}
			if j == 6 {
				data.Date = rows[i][6]
			}

		}
		db.Create(data)
	}
}

//inserting 1000 data into temps table
func go2(wg *sync.WaitGroup) {
	var data Temp
	defer wg.Done()
	db := database.DBconnection()

	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)

	}
	rows, err := f.GetRows("Sheet1")
	limit := len(rows)

	for i := limit / 2; i < limit; i++ {
		for j := 0; j < 7; j++ {
			if j == 0 {
				i, err := strconv.Atoi(rows[i][0])
				if err != nil {
					fmt.Print(err)
				}
				data.ID = i

			}
			if j == 1 {
				data.Firstname = rows[i][1]
			}
			if j == 2 {
				data.Lastname = rows[i][2]
			}
			if j == 3 {
				data.Gender = rows[i][3]
			}
			if j == 4 {
				data.Country = rows[i][4]
			}
			if j == 5 {
				i, err := strconv.Atoi(rows[i][5])
				if err != nil {
					fmt.Print(err)
				}
				data.Age = i
			}

			if j == 6 {
				data.Date = rows[i][6]
			}

		}
		db.Create(data)

	}
}

//inserting 500 data into temps table
func go3(ch chan string) {
	db := database.DBconnection()
	var data Temp
	f, err := excelize.OpenFile(<-ch)
	if err != nil {
		fmt.Println(err)

	}
	rows, err := f.GetRows("Sheet1")
	limit := len(rows)

	for i := limit / 4; i < limit/2; i++ {
		for j := 0; j < 7; j++ {
			if j == 0 {
				i, err := strconv.Atoi(rows[i][0])
				if err != nil {
					fmt.Print(err)
				}
				data.ID = i

			}
			if j == 1 {
				data.Firstname = rows[i][1]
			}
			if j == 2 {
				data.Lastname = rows[i][2]
			}
			if j == 3 {
				data.Gender = rows[i][4]
			}
			if j == 4 {
				data.Country = rows[i][4]
			}
			if j == 5 {
				i, err := strconv.Atoi(rows[i][5])
				if err != nil {
					fmt.Print(err)
				}
				data.Age = i
			}
			if j == 6 {
				data.Date = rows[i][6]
			}

		}
		db.Create(data)
	}
}
func go4(ch chan string) {

	for i := 2; i < 7; i++ {
		fmt.Println(<-ch)
	}
}
func setroutes(app *fiber.App) {
	app.Get("/showuser/:id", routes.GetBook)
	app.Post("/create", routes.Create)
	app.Get("/showusers", routes.GetBooks)
	app.Delete("/delete/:id", routes.Delete)
	app.Put("/update/:id", routes.Update)
}

func main() {

	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		//channel
		ch := make(chan string, 10)
		ch <- "sample.xlsx"
		go go3(ch)
		go go4(ch)

		//sync.waitgroup
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go go1(wg)
		go go2(wg)
		wg.Wait()
		return c.SendString("data successfully uploaded")

	})

	setroutes(app)

	app.Listen(":3000")
}

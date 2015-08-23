package app

import (
	"./errors"
	"./models"
	"fmt"
	"os"
)

type app struct {
	inputFileName  string
	resultFilePath string
}

var (
	Errs   *myErrors.Error
	Client *models.Client
	File   *models.File
)

func New() app {
	return make(app)
}

func (a *app) init() {
	a.Errs.Init()
	a.Client.Init()
	a.File.Init()
	if len(os.Args) < 2 {
		a.Errs.Check(nil, 2, "")
	}
	a.inputFileName = os.Args[1]
	a.resultFilePath = os.Args[2]
}

func (a *app) Start() {
	a.init()
	File.Name = a.inputFileName
	fmt.Print("Open file " + File.Name + "... ")
	Errs.Check(File.Open())
	fmt.Println("OK")
	defer File.Basic.Close()
	fmt.Print("Read file " + File.Name + "... ")
	Errs.Check(File.Read())
	fmt.Println("OK")
	fmt.Println("Sending file " + File.Name + " as " + File.Name + "... ")
	Client.Connect()
	Client.SendFile(File)
	fmt.Println("OK")
}

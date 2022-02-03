package main

import (
	"fq_task_serivce/internal/controllers"
	dbinit "fq_task_serivce/internal/db"
	"fq_task_serivce/internal/routers"
	"fq_task_serivce/internal/services"
	"fq_task_serivce/internal/types"

	"github.com/StanDenisov/fq_utils/confclient"
	"github.com/StanDenisov/fq_utils/profile"
	"github.com/StanDenisov/fq_utils/users"
	"github.com/labstack/echo"
)

func main() {
	cnf := confclient.ParseFlagsAndGetConfig()

	e := echo.New()
	db := dbinit.ConnectDB(cnf)

	err := db.AutoMigrate(&types.Comment{}, &types.Issue{}, &users.User{}, &profile.Profile{})
	if err != nil {
		panic("Failed auto migrate")
	}

	ir := services.NewIssueService(db)
	ic := controllers.NewIssueController(ir)

	routers.RouteIssues(e, ic)

	e.Logger.Fatal(e.Start(":" + cnf.AppPort))
}

package main

import (
	"loan_tracker_api/deliveries/controllers"
	"loan_tracker_api/deliveries/router"
	_ "loan_tracker_api/docs"
	"loan_tracker_api/infrastructure"
	"loan_tracker_api/repository"
	"loan_tracker_api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

// @title           Loan Tracker API
// @version         1.0
// @description     Loan Tracker API for managing loans and user accounts.

// @contact.name   Dagim Mesfin
// @contact.email  dagmmesfin99@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

func main() {
	client := infrastructure.MongoDBInit() //mongodb initialization

	userrepo := repository.NewUserRepository(client)
	useruse := usecase.NewUserUsecase(userrepo, time.Second*300)
	usercont := controllers.NewUserController(useruse)

	loanrepo := repository.NewLoanRepository(client)
	loanuse := usecase.NewLoanUsecase(loanrepo, time.Second*300)
	loancont := controllers.NewLoanController(loanuse)

	r := gin.Default()
	router.SetRouter(r, usercont, client, loancont)
	r.Run()
}

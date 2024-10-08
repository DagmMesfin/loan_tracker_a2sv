package main

import (
	"loan_tracker_api/deliveries/controllers"
	"loan_tracker_api/deliveries/router"
	"loan_tracker_api/infrastructure"
	"loan_tracker_api/repository"
	"loan_tracker_api/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

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

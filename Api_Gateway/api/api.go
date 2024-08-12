package api

import (
	handlerC "api/api/handler/budgeting"

	_ "api/docs"
	_ "api/genprotos/auth"
	_ "api/genprotos/budgeting"

	"api/api/middlerware"
	"log"

	"github.com/casbin/casbin/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Budgeting SYSTEM API
// @version 1.0
// @description Developing a platform that helps users track their spending, set a budget and manage their financial goals
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin( /*AuthConn, */ budgetingConn *grpc.ClientConn) *gin.Engine {
	budgeting := handlerC.NewBudgetingHandler(budgetingConn)

	router := gin.Default()

	enforcer, err := casbin.NewEnforcer("/home/sobirov/go/src/gitlab.com/PersonalFinanceTracker/Api_Gateway/api/model.conf", "/home/sobirov/go/src/gitlab.com/PersonalFinanceTracker/Api_Gateway/api/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	sw := router.Group("/")
	sw.Use(middlerware.NewAuth(enforcer))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	account := router.Group("/account")
	{
		account.POST("/create", budgeting.CreateAccount)
		account.PUT("/update", budgeting.UpdateAccount)
		account.GET("/get/{id}", budgeting.GetAccount)
		account.DELETE("/delete/{id}",budgeting.DeleteAccount)
		account.GET("/get",budgeting.ListAccounts)
	}
	budget := router.Group("/budget")
	{
		budget.POST("/create", budgeting.CreateBudget)
		budget.PUT("/update", budgeting.UpdateBudget)
		budget.GET("/get/{id}", budgeting.GetBudget)
		budget.DELETE("/delete/{id}",budgeting.DeleteBudget)
		budget.GET("/get",budgeting.ListBudgets)
	}
	category := router.Group("/category")
	{
		category.POST("/create", budgeting.CreateCategory)
		category.PUT("/update", budgeting.UpdateCategory)
		category.GET("/get/{id}", budgeting.GetCategory)
		category.DELETE("/delete/{id}",budgeting.DeleteCategory)
		category.GET("/get",budgeting.ListCategories)
	}

	goal := router.Group("/goal")
	{
		goal.POST("/create", budgeting.CreateGoal)
		goal.PUT("/update", budgeting.UpdateGoal)
		goal.GET("/get/{id}", budgeting.GetGoal)
		goal.DELETE("/delete/{id}",budgeting.DeleteGoal)
		goal.GET("/get",budgeting.ListGoals)
	}

	transaction := router.Group("/transaction")
	{
		transaction.POST("/create", budgeting.CreateTransaction)
		transaction.PUT("/update", budgeting.UpdateTransaction)
		transaction.GET("/get/{id}", budgeting.GetTransaction)
		transaction.DELETE("/delete/{id}",budgeting.DeleteTransaction)
		transaction.GET("/get",budgeting.ListTransactions)
	}


	return router
}

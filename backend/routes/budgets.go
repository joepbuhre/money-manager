package routers

import (
	"context"
	"log"
	"money-manager/database"
	m "money-manager/middlewares"
	u "money-manager/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddBudgetRouter(rg *gin.RouterGroup, env *u.Env) {
	public := rg.Group("/budgets")

	public.GET("/", m.IsAllowedTo(env, "Budgets", m.CanRead), GetBudgets(env))

	public.POST("/", m.IsAllowedTo(env, "Budgets", m.CanCreate), CreateBudget(env))

	public.DELETE("/:budgetid", m.IsAllowedTo(env, "Budgets", m.CanDelete), DeleteBudget(env))

}

// Create a Budget.
//
//	@ID			create-Budget
//	@Produce	json
//	@Router		/Budgets [post]
func CreateBudget(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var newBudget database.CreateBudgetParams
		_, err = u.ReadRequestBody(c, &newBudget)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
			return
		}

		var budget database.Budget
		budget, err = env.Queries.CreateBudget(c, newBudget)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
			return
		}
		c.JSON(http.StatusOK, budget)

	}
}

// Get all Budgets.
//
//	@ID			get-Budget
//	@Produce	json
//	@Router		/Budgets [get]
func GetBudgets(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var Budgets, _ = env.Queries.ListBudgets(context.Background())
		c.JSON(http.StatusOK, Budgets)
	}
}

// Delete Budget
//
//	@ID		delete-Budget
//	@Router	/Budgets/:Budgetid [delete]
func DeleteBudget(env *u.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q = env.Queries
		var idStr string
		var id int
		var err error
		var errors u.JsonErrors
		idStr = c.Param("Budgetid")

		id, err = strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			errors.Add(u.ApiErrors[u.ErrInternal])
			c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
			return
		}

		// Delete Budget here
		err = q.DeleteBudget(context.Background(), int32(id))
		if err != nil {
			log.Println(err)
			errors.Add(u.ApiErrors[u.ErrInternal])
			c.AbortWithStatusJSON(errors.ErrorCode, errors.ToString())
			return
		}

		// Handle an accepted response
		c.Status(http.StatusAccepted)
	}
}

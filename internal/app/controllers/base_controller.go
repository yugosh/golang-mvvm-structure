package controllers

import (
	"BACKEND-GOLANG-MVVM/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	expressionService *services.ExpressionService
	formulaService    *services.FormulaService
}

func NewBaseController(es *services.ExpressionService, fs *services.FormulaService) *BaseController {
	return &BaseController{expressionService: es, formulaService: fs}
}

// Endpoint untuk mendapatkan daftar fungsi yang tersedia
func (bc *BaseController) GetAvailableFunctions(c *gin.Context) {
	functions := bc.formulaService.AvailableFunctions()

	c.JSON(http.StatusOK, functions)
}

// Endpoint untuk menghitung nilai dari ekspresi
func (bc *BaseController) CalculateExpression(c *gin.Context) {
	var req struct {
		Expression           string  `json:"expression"`
		BaseSalary           float64 `json:"base_salary"`
		AttendanceDays       int     `json:"attendance_days"`
		LateMinutes          int     `json:"late_minutes"`
		LatePenaltyPerMinute float64 `json:"late_penalty_per_minute"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parameters := map[string]interface{}{
		"BaseSalary":           req.BaseSalary,
		"AttendanceDays":       req.AttendanceDays,
		"LateMinutes":          req.LateMinutes,
		"LatePenaltyPerMinute": req.LatePenaltyPerMinute,
	}

	result, err := bc.expressionService.EvaluateExpression(req.Expression, parameters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

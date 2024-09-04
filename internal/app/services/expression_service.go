package services

import (
	"fmt"
	"log"
	"math"

	"github.com/Knetic/govaluate"
)

type ExpressionService struct {
	formulaService *FormulaService
}

func NewExpressionService(fs *FormulaService) *ExpressionService {
	return &ExpressionService{
		formulaService: fs,
	}
}

func (es *ExpressionService) EvaluateExpression(expr string, parameters map[string]interface{}) (string, error) {
	// Define functions that can be used in the expressions
	functions := map[string]govaluate.ExpressionFunction{
		//CATEGORY : PARAMETER
		"NilaiGaji": func(args ...interface{}) (interface{}, error) {
			log.Println("NilaiGaji:", es.formulaService.NilaiGaji(parameters))
			return es.formulaService.NilaiGaji(parameters), nil
		},
		"Hadir": func(args ...interface{}) (interface{}, error) {
			log.Println("Hadir:", es.formulaService.Hadir(parameters))
			return es.formulaService.Hadir(parameters), nil
		},
		"TelatMenit": func(args ...interface{}) (interface{}, error) {
			log.Println("TelatMenit:", es.formulaService.TelatMenit(parameters))
			return es.formulaService.TelatMenit(parameters), nil
		},
		"DendaPerMenit": func(args ...interface{}) (interface{}, error) {
			log.Println("DendaPerMenit:", es.formulaService.DendaPerMenit(parameters))
			return es.formulaService.DendaPerMenit(parameters), nil
		},

		//CATEGORY : LOGICAL
		"if": func(args ...interface{}) (interface{}, error) {
			condition := args[0].(bool)
			if condition {
				return args[1], nil
			}
			return args[2], nil
		},

		//CATEGORY : MATH
		"ABS": func(args ...interface{}) (interface{}, error) {
			return math.Abs(args[0].(float64)), nil
		},
		"MAX": func(args ...interface{}) (interface{}, error) {
			max := args[0].(float64)
			for _, arg := range args[1:] {
				value := arg.(float64)
				if value > max {
					max = value
				}
			}
			return max, nil
		},
		"MIN": func(args ...interface{}) (interface{}, error) {
			min := args[0].(float64)
			for _, arg := range args[1:] {
				value := arg.(float64)
				if value < min {
					min = value
				}
			}
			return min, nil
		},
		"SUM": func(args ...interface{}) (interface{}, error) {
			sum := 0.0
			for _, arg := range args {
				sum += arg.(float64)
			}
			return sum, nil
		},
		"AVG": func(args ...interface{}) (interface{}, error) {
			sum := 0.0
			for _, arg := range args {
				sum += arg.(float64)
			}
			return sum / float64(len(args)), nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
	if err != nil {
		log.Println("Error creating expression:", err)
		return "", err
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Println("Error evaluating expression:", err)
		return "", err
	}

	log.Println("Evaluation Result:", result)
	return fmt.Sprintf("%v", result), nil
}

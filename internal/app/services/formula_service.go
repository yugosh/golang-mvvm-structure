package services

type FormulaService struct{}

func NewFormulaService() *FormulaService {
	return &FormulaService{}
}

// Implementasi metode AvailableFunctions
func (fs *FormulaService) AvailableFunctions() []string {
	return []string{
		"NilaiGaji()",
		"Hadir()",
		"TelatMenit()",
		"DendaPerMenit()",

		"if(condition, trueResult, falseResult)",
		"ABS(n1)",
		"MAX(n1, n2)",
		"MIN(n1, n2)",
		"SUM(n1, n2, ..., nx)",
		"AVG(n1, n2, ..., nx)",
	}
}

func (fs *FormulaService) NilaiGaji(parameters map[string]interface{}) float64 {
	return parameters["BaseSalary"].(float64)
}

func (fs *FormulaService) Hadir(parameters map[string]interface{}) float64 {
	return float64(parameters["AttendanceDays"].(int)) // Konversi ke float64
}

func (fs *FormulaService) TelatMenit(parameters map[string]interface{}) float64 {
	return float64(parameters["LateMinutes"].(int)) // Konversi ke float64
}

func (fs *FormulaService) DendaPerMenit(parameters map[string]interface{}) float64 {
	return parameters["LatePenaltyPerMinute"].(float64)
}

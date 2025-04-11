package nutrition

// NutritionService описывает поведение сервиса для оценки сбалансированности рецептуры.
type NutritionService interface {
	// PredictSafety принимает состав нутриентов (в соответствующих единицах)
	// и массу порции (в граммах), возвращая риск несбалансированности (0-1),
	// сообщение с выявленными отклонениями и ошибку (если возникнет).
	PredictSafety(composition map[string]float64, portionWeight float64) (score float64, message string, err error)
}

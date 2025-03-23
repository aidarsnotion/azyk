package nutrition

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// NutritionModel — структура для нейросетевой модели оценки рецептуры.
type NutritionModel struct {
	g      *gorgonia.ExprGraph
	model  *gorgonia.Node
	input  *gorgonia.Node
	output *gorgonia.Node
	vm     gorgonia.VM
}

// NewNutritionModel создает и возвращает новый экземпляр NutritionModel.
func NewNutritionModel(inputSize, hiddenSize, outputSize int) (*NutritionModel, error) {
	g := gorgonia.NewGraph()

	// Создаем узел для входного вектора
	input := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, inputSize), gorgonia.WithName("input"))

	// Первый слой: входной -> скрытый
	w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(inputSize, hiddenSize), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	b1 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(hiddenSize), gorgonia.WithName("b1"), gorgonia.WithInit(gorgonia.Zeroes()))
	l1pre, err := gorgonia.Add(gorgonia.Must(gorgonia.Mul(input, w1)), b1)
	if err != nil {
		return nil, err
	}
	l1 := gorgonia.Must(gorgonia.Rectify(l1pre))

	// Второй слой: скрытый -> выходной
	w2 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(hiddenSize, outputSize), gorgonia.WithName("w2"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	b2 := gorgonia.NewVector(g, tensor.Float64, gorgonia.WithShape(outputSize), gorgonia.WithName("b2"), gorgonia.WithInit(gorgonia.Zeroes()))
	outPre, err := gorgonia.Add(gorgonia.Must(gorgonia.Mul(l1, w2)), b2)
	if err != nil {
		return nil, err
	}
	// Применяем SoftMax для получения вероятностей
	out := gorgonia.Must(gorgonia.SoftMax(outPre))

	// Создаем модель
	return &NutritionModel{
		g:      g,
		model:  out,
		input:  input,
		output: out,
	}, nil
}

// TrainModel — функция, которая обучает модель на синтетических данных.
// Здесь приведен упрощенный пример обучения с использованием случайных данных.
func (nm *NutritionModel) TrainModel(XTrain tensor.Tensor, yTrain tensor.Tensor, epochs int) error {
	// Здесь должна быть реализация обучения: определение функции потерь, оптимизатора и т.д.
	// Для демонстрации, мы имитируем процесс обучения.
	log.Println("Начало обучения модели...")
	time.Sleep(2 * time.Second)
	log.Println("Обучение завершено (демо-версия).")
	return nil
}

// Predict принимает входной вектор (содержания нутриентов) и возвращает предсказание модели.
func (nm *NutritionModel) Predict(inputData []float64) ([]float64, error) {
	// Обновляем входной узел
	err := gorgonia.Let(nm.input, tensor.New(tensor.WithShape(1, len(inputData)), tensor.WithBacking(inputData)))
	if err != nil {
		return nil, err
	}

	// Создаем виртуальную машину для выполнения графа
	nm.vm = gorgonia.NewTapeMachine(nm.g)
	defer nm.vm.Close()

	if err = nm.vm.RunAll(); err != nil {
		return nil, err
	}
	outputVal := nm.output.Value()
	// Преобразуем результат в []float64
	result, ok := outputVal.Data().([]float64)
	if !ok {
		return nil, fmt.Errorf("невозможно преобразовать результат")
	}
	return result, nil
}

// PredictSafety — функция, которая принимает состав нутриентов и массу порции,
// корректирует нормы и возвращает риск (вероятность небезопасности) и сообщение.
func PredictSafety(composition map[string]float64, portionWeight float64) (float64, string, error) {
	// Суточные нормы, скорректированные для порции
	rdaPortion := AdjustedRDA(portionWeight)
	keys := []string{"protein", "fat", "carbohydrates", "Na", "Ca", "Fe", "Zn"}
	inputVector := make([]float64, len(keys))
	message := ""

	for i, key := range keys {
		val, ok := composition[key]
		if !ok {
			val = 0
		}
		inputVector[i] = val
		percent := (val / rdaPortion[key]) * 100
		if percent > 150 {
			message += fmt.Sprintf("превышение по %s (%.0f%%); ", key, percent)
		} else if percent < 70 {
			message += fmt.Sprintf("дефицит по %s (%.0f%%); ", key, percent)
		}
	}

	// Создаем модель (например, вход размер 7, скрытый слой 16, выход 2)
	nModel, err := NewNutritionModel(len(keys), 16, 2)
	if err != nil {
		return 0, "", err
	}

	// Здесь предполагается, что модель уже обучена.
	// Для демонстрации можно пропустить этап обучения или вызвать TrainModel с синтетическими данными.
	XTrain, yTrain := generateTrainingData(1000, portionWeight)
	err = nModel.TrainModel(XTrain, yTrain, 1000)
	if err != nil {
		return 0, "", err
	}

	// Получаем предсказание
	pred, err := nModel.Predict(inputVector)
	if err != nil {
		return 0, "", err
	}

	// Предположим, индекс 1 соответствует классу "несбалансировано" (риск)
	riskScore := pred[1]
	if message == "" {
		message = "Нарушений не выявлено"
	}
	return riskScore, message, nil
}

// AdjustedRDA корректирует суточные нормы для порции, как описано ранее.
func AdjustedRDA(portionWeight float64) map[string]float64 {
	factor := portionWeight / 2000.0
	adjusted := make(map[string]float64)
	for key, value := range RDA {
		adjusted[key] = value * factor
	}
	return adjusted
}

// generateTrainingData генерирует синтетическую обучающую выборку для модели.
func generateTrainingData(numSamples int, portionWeight float64) (tensor.Tensor, tensor.Tensor) {
	rand.Seed(time.Now().UnixNano())
	keys := []string{"protein", "fat", "carbohydrates", "Na", "Ca", "Fe", "Zn"}
	rdaPortion := AdjustedRDA(portionWeight)
	data := make([][]float64, numSamples)
	labels := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		sample := make([]float64, len(keys))
		isBalanced := true
		for j, key := range keys {
			norm := rdaPortion[key]
			var val float64
			if rand.Float64() < 0.3 {
				// сбалансированное значение: 90-110%
				val = norm * (0.9 + 0.2*rand.Float64())
			} else {
				if rand.Float64() < 0.5 {
					val = norm * (0.5 + 0.2*rand.Float64())
				} else {
					val = norm * (1.5 + 0.5*rand.Float64())
				}
				isBalanced = false
			}
			sample[j] = val
		}
		data[i] = sample
		if isBalanced {
			labels[i] = 0 // сбалансированно
		} else {
			labels[i] = 1 // несбалансированно
		}
	}
	// Преобразуем данные в тензоры
	x := tensor.New(tensor.WithShape(numSamples, len(keys)), tensor.WithBacking(data))
	y := tensor.New(tensor.WithShape(numSamples), tensor.WithBacking(labels))
	return x, y
}

func main() {
	// Пример использования PredictSafety для 200 г порции.
	portionWeight := 200.0
	exampleComposition := map[string]float64{
		"protein":       10,  // г
		"fat":           12,  // г
		"carbohydrates": 20,  // г
		"Na":            300, // мг
		"Ca":            80,  // мг
		"Fe":            2,   // мг
		"Zn":            1,   // мг
	}
	risk, msg, err := PredictSafety(exampleComposition, portionWeight)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Риск несбалансированности: %.3f\n", risk)
	fmt.Printf("Сообщение: %s\n", msg)
}

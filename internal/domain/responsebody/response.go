package responsebody

type Response[T any] struct {
	Status      string `json:"STATUS"`                        // success, error, etc.
	Code        int    `json:"CODE"`                          // Внутренний код, нужно потом под коды прописать конфлюенс
	Message     string `json:"MESSAGE"`                       // сообщение понятное для человека
	RequestId   string `json:"REQUEST_ID"`                    // для логов
	Total       int    `json:"TOTAL_RECORDS_COUNT,omitempty"` // Общее количество записей
	TotalPages  int    `json:"TOTAL_PAGES"`
	CurrentPage int    `json:"CURRENT_PAGE"`
	Data        T      `json:"DATA"` // полезная нагрузка
}

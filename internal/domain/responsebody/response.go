package responsebody

type Response[T any] struct {
	Status      string `json:"status"`                        // success, error, etc.
	Code        int    `json:"code"`                          // Внутренний код, нужно потом под коды прописать конфлюенс
	Message     string `json:"message"`                       // сообщение понятное для человека
	RequestId   string `json:"request_id"`                    // для логов
	Total       int    `json:"total_records_count,omitempty"` // Общее количество записей
	TotalPages  int    `json:"total_pages"`
	CurrentPage int    `json:"current_page"`
	Data        T      `json:"data"` // полезная нагрузка
}

package responsebody

type Response[T any] struct {
	Cmd         string `json:"cmd"`                           // название метода
	Code        int    `json:"code"`                          // внутренний код, нужно потом под коды прописать конфлюенс
	Message     string `json:"message"`                       // сообщение понятное для человека
	RequestId   string `json:"request_id"`                    // для логов
	Total       int    `json:"total_records_count,omitempty"` // Общее количество записей
	TotalPages  int    `json:"total_pages"`
	CurrentPage int    `json:"current_page"`
	Data        T      `json:"data"` // полезная нагрузка
}

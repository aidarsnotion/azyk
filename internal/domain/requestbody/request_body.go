package requestbody

type BaseRequest struct {
	RequestID   string `json:"request_id" validate:"required"` // идентификатор запроса, для удобства дебага
	Cmd         string `json:"cmd" validate:"required"`        // Команда/название метода
	CurrentPage int    `json:"cuurent_page,omitempty"`         // Номер страницы
	PageSize    int    `json:"page_size,omitempty"`            // Кол-во элементов на странице
}

type RequestWithPayload[T any] struct {
	BaseRequest
	Data T `json:"DATA" validate:"required"`
}

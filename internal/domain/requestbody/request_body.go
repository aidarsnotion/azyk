package requestbody

type BaseRequest struct {
	RequestID   string `json:"REQUEST_ID" validate:"required"` // идентификатор запроса, для удобства дебага
	Cmd         string `json:"CMD" validate:"required"`        // Команда/название метода
	CurrentPage int    `json:"CURRENT_PAGE,omitempty"`         // Номер страницы
	PageSize    int    `json:"PAGE_SIZE,omitempty"`            // Кол-во элементов на странице
}

type RequestWithPayload[T any] struct {
	BaseRequest
	Data T `json:"DATA" validate:"required"`
}

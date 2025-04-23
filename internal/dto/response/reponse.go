package response

type APIResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Errors  any    `json:"errors,omitempty"`
}

func SuccessResponse[T any](msg string, data T) *APIResponse[T] {
	return &APIResponse[T]{
		Status:  "success",
		Message: msg,
		Data:    data,
	}
}

func ValidationResponse[T any](validationErrors map[string]string) *APIResponse[T] {
	var emptyObject = map[string]any{}
	return &APIResponse[T]{
		Status:  "validation error",
		Message: "Invalid request",
		Data:    any(emptyObject).(T),
		Errors:  validationErrors,
	}
}

func ErrorResponse[T any](msg string, errs any) *APIResponse[T] {
	var emptyObject = map[string]any{}
	return &APIResponse[T]{
		Status:  "error",
		Message: msg,
		Data:    any(emptyObject).(T),
		Errors:  errs,
	}
}

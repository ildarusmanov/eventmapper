package models

type JsonResponse struct {
	IsOk bool
	Status string
}

func CreateNewJsonResponse(isOk bool, status string) *JsonResponse {
	return &JsonResponse{isOk, status}
}
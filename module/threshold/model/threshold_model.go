package model

type ThresholdCreateRequest struct {
	Parameter string  `json:"parameter" validate:"required"`
	Threshold float64 `json:"threshold" validate:"required"`
}

type ThresholdUpdateRequest struct {
	Parameter string  `json:"parameter" validate:"required"`
	Threshold float64 `json:"threshold" validate:"required"`
}

type ThresholdResponse struct {
	Parameter string  `json:"parameter"`
	Threshold float64 `json:"threshold"`
}

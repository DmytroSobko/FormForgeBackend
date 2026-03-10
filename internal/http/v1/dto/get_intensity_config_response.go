package dto

type IntensityConfig struct {
	Type          string  `json:"type"`
	Multiplier    float64 `json:"multiplier"`
	FatigueFactor float64 `json:"fatigueFactor"`
}

type IntensityConfigsResponse struct {
	Intensities []IntensityConfig `json:"intensities"`
}

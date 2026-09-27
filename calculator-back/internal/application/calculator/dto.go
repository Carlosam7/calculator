package calculator

type EvaluateExpressionRequest struct {
	Expression string `json:"expression"`
}

type EvaluateExpressionResponse struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}

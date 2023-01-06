package sse

const (
	DebtCreated   = "debt_created"
	DebtCanceled  = "debt_canceled"
	DebtFulfilled = "debt_fulfilled"
)

var Events = []string{DebtCreated, DebtCanceled, DebtFulfilled}

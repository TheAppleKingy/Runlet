package entities

type Attempt struct {
	ID        int       `json:"id"`
	Amount    uint      `json:"amount"`
	Done      bool      `json:"done"`
	TestCases TestCases `json:"test_cases"`
}

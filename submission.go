// Package api описывает контракт между бэкендом и judge.
// Менять его можно только через pull request, который посмотрели оба.
package api

// Language — язык решения игрока.
type Language string

const (
	LangCPP    Language = "cpp"
	LangPython Language = "python"
)

// Submission — задание на проверку, которое бэкенд кладёт в очередь.
type Submission struct {
	ID        string   `json:"id"`
	MatchID   string   `json:"match_id"`
	Round     int      `json:"round"`
	PlayerID  string   `json:"player_id"`
	ProblemID string   `json:"problem_id"`
	Language  Language `json:"language"`
	Source    string   `json:"source"`
}

// Verdict — итог проверки посылки.
type Verdict string

const (
	VerdictAccepted      Verdict = "AC" // корректно; эффективность — в RungsPassed
	VerdictWrongAnswer   Verdict = "WA" // неверный ответ (на тестах корректности или на ступени)
	VerdictRuntimeError  Verdict = "RE" // упало
	VerdictCompileError  Verdict = "CE" // не скомпилировалось
	VerdictTimeLimit     Verdict = "TL" // не уложилось уже на тестах корректности
	VerdictMemoryLimit   Verdict = "ML" // превышена память
	VerdictInternalError Verdict = "IE" // ошибка judge, не игрока
)

// RungResult — результат одной ступени лестницы.
type RungResult struct {
	SizeExp      int    `json:"size_exp"` // n = 2^SizeExp
	Passed       bool   `json:"passed"`
	Instructions uint64 `json:"instructions"` // 0, пока замер инструкций не реализован
	MemoryBytes  int64  `json:"memory_bytes"`
}

// Result — ответ judge, который бэкенд читает из очереди.
type Result struct {
	SubmissionID string       `json:"submission_id"`
	Verdict      Verdict      `json:"verdict"`
	Rungs        []RungResult `json:"rungs,omitempty"`
	RungsPassed  int          `json:"rungs_passed"`
	RungsTotal   int          `json:"rungs_total"`
	Score        float64      `json:"score"`                // RungsPassed / RungsTotal
	Complexity   string       `json:"complexity,omitempty"` // оценка класса, например "O(n log n)"
	Message      string       `json:"message,omitempty"`    // текст ошибки компиляции и т. п.
}

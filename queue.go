package api

// Имена потоков Redis (Redis Streams), через которые общаются бэкенд и judge.
const (
	// Бэкенд пишет сюда Submission в поле "payload" (JSON).
	StreamSubmissions = "judge:submissions"
	// Judge пишет сюда Result в поле "payload" (JSON).
	StreamResults = "judge:results"
	// Группа потребителей judge-воркеров.
	JudgeConsumerGroup = "judges"
)

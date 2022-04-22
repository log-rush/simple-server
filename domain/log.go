package domain

import "context"

type Log struct {
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
	Stream    string `json:"stream"`
}

type LogUseCase interface {
	SendLog(ctx context.Context, log *Log) error
	SendLogBatch(ctx context.Context, logs *[]Log) error
}

type LogRepository interface {
	AddLogs(ctx context.Context, logs *[]Log) error
	FetchLogs(ctx context.Context, stream string) ([]Log, error)
}

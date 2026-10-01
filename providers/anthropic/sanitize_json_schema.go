package anthropic

import "github.com/nrbrd/ai-sdk-legacy/internal/anthropicschema"

func sanitizeJSONSchema(schema map[string]any) map[string]any {
	return anthropicschema.Sanitize(schema)
}

func formatConstraintName(key string) string {
	return anthropicschema.FormatConstraintName(key)
}

func formatConstraintValue(value any) string {
	return anthropicschema.FormatConstraintValue(value)
}

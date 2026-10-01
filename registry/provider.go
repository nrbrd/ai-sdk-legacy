package registry

import "github.com/nrbrd/ai-sdk-legacy/provider"

type Provider interface {
	LanguageModel(modelID string) (provider.LanguageModel, error)
}

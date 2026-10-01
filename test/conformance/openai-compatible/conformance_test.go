//go:build conformance

package openaicompatible_test

import (
	"testing"

	"github.com/nrbrd/ai-sdk-legacy/provider"
	openaicompatible "github.com/nrbrd/ai-sdk-legacy/providers/openai-compatible"
	"github.com/nrbrd/ai-sdk-legacy/test/conformance"
)

func TestConformance(t *testing.T) {
	providerDir := "."
	cases := conformance.DiscoverTestCases(t, providerDir)

	if len(cases) == 0 {
		t.Skip("no conformance test cases found")
	}

	factory := func(baseURL string, cfg *conformance.Config) (provider.LanguageModel, error) {
		model := openaicompatible.New(
			cfg.Model,
			openaicompatible.WithBaseURL(baseURL+"/v1"),
			openaicompatible.WithAPIKey("test-api-key"),
			openaicompatible.WithIncludeUsage(true),
			openaicompatible.WithStructuredOutputs(true),
		)
		return model, nil
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			conformance.RunTestCase(t, tc, factory)
		})
	}
}

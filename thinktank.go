package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/tmc/langchaingo/llms"
)

const DefaultConcurrency = 5

type ThinkTank struct {
	Prompt      string
	Model       llms.Model
	Experts     []*Expert
	Concurrency int
}

func NewThinkTank(model llms.Model, experts ...*Expert) *ThinkTank {
	return &ThinkTank{
		Prompt:      "Consider the points from all experts, then provide a comprehensive and definitive answer to the question: %s",
		Model:       model,
		Experts:     experts,
		Concurrency: DefaultConcurrency,
	}
}

type ThinkTankResult struct {
	ExpertResults []*ExpertResult
}

func (tt *ThinkTank) Evaluate(ctx context.Context, input string) *ThinkTankResult {

	var results []*ExpertResult
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, expert := range tt.Experts {
		wg.Add(1)

		go func(expert *Expert) {
			defer wg.Done()
			result, err := expert.Evaluate(ctx, input)
			result.Err = err

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(expert)
	}

	wg.Wait()

	return &ThinkTankResult{
		ExpertResults: results,
	}
}

func (tt *ThinkTank) Answer(ctx context.Context, input string) (string, error) {
	result := tt.Evaluate(ctx, input)

	var messages []llms.MessageContent

	// First we innocently pose the question verbatim
	originalQuestion := llms.MessageContent{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{llms.TextPart(input)},
	}

	messages = append(messages, originalQuestion)

	// Next we let each expert chime in
	for _, expertResult := range result.ExpertResults {

		if expertResult.Err == nil {
			opinion := fmt.Sprintf("%s: %s\n", expertResult.Expert.Name, expertResult.Text)

			expertOpinion := llms.MessageContent{
				Role:  llms.ChatMessageTypeSystem,
				Parts: []llms.ContentPart{llms.TextPart(opinion)},
			}

			messages = append(messages, expertOpinion)
		}
	}

	// Lastly we instruct the LLM to re-answer the question in light of
	// the new expert opinions
	thinkTankPrompt := fmt.Sprintf(tt.Prompt, input)
	systemMessage := llms.MessageContent{
		Role:  llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{llms.TextPart(thinkTankPrompt)},
	}

	messages = append(messages, systemMessage)

	content, err := tt.Model.GenerateContent(ctx, messages)

	if err != nil {
		return "", err
	}

	return content.Choices[0].Content, nil
}

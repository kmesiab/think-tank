package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var llm, _ = openai.New()

var callOptions = &llms.CallOptions{
	Model:          "gpt-4o-mini",
	CandidateCount: 0,
	MaxTokens:      1024,
	Temperature:    0.9,
	Tools:          nil,
	ToolChoice:     nil,
}

var economicsExpert = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Economics Expert",
	Description: "A highly skilled economist with expertise in international trade, economic growth, and policy analysis.",
	Prompt:      "Given the following context: %s\nCan you expand the question to include related economic concepts to enhance the answer?",
}

var politicsExpert = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Politics Expert",
	Description: "A highly knowledgeable political scientist with expertise in international relations, political science, and policy analysis.",
	Prompt:      "Given the following context: %s\nCan you expand the question to include related political concepts to enhance the answer?",
}

var ethicsExpert = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Ethics Expert",
	Description: "A highly knowledgeable ethics professor with expertise in international relations, ethical persuasion, and policy analysis.",
	Prompt:      "Given the following context: %s\nCan you expand the question to include related ethical concepts to enhance the answer?",
}

var devilsAdvocate = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Devil's Advocate",
	Description: "A pragmatist that reframes questions to elicit new insightful responses.",
	Prompt:      "What are the potential consequences of %s? Consider the potential benefits and drawbacks, as well as the potential for harm.",
}

func main() {

	ctx := context.Background()
	tk := NewThinkTank(llm, devilsAdvocate, economicsExpert, politicsExpert, ethicsExpert)

	input := "Will forgiving student loan debt have a long term positive or negative affect on the US economy?"

	finalAnswer, err := tk.Answer(ctx, input)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Final Answer: %s\n", finalAnswer)

}

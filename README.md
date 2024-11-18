# ThinkTank

ThinkTank is a modular, collaborative AI system designed to generate comprehensive
and nuanced answers to complex questions. By leveraging multiple AI-powered
`experts`, ThinkTank combines diverse perspectives to provide well-rounded
insights, mimicking the dynamics of human expert panels.

---

## Features

- **Collaborative AI Models**: Integrates multiple AI `experts`, each with
  unique prompts and areas of specialization.
- **Modular Design**: Easily add, remove, or customize experts to tailor
  ThinkTank to specific domains or use cases.
- **Emergent Intelligence**: Synthesizes insights from various experts,
  producing outputs that are greater than the sum of their parts.
- **Interdisciplinary Reasoning**: Handles complex, multifaceted queries
  requiring expertise in diverse fields.
- **Customizable Outputs**: Refines answers through iterative evaluation and
  context-aware synthesis.

---

## How It Works

ThinkTank orchestrates the following workflow:

1. **Input**: A user provides a question or topic.
2. **Expert Evaluation**: Each AI expert independently evaluates the input
   based on its unique expertise and prompt.
3. **Synthesis**: ThinkTank consolidates expert responses, incorporating their
   insights into a comprehensive final answer.

---

## Example Usage

Here’s how you can use ThinkTank in a real-world scenario:

### Input:
```plaintext
Will forgiving student loan debt have a long-term positive or negative effect
on the US economy?
```

### Output:
```plaintext
Final Answer: Forgiving student loan debt could provide immediate economic
relief to borrowers and stimulate consumer spending. However, it may lead to
inflationary pressures and a shift in fiscal priorities. Expert analyses
highlight potential benefits for low-income households but raise concerns
about long-term economic inequality and fiscal sustainability.
```

---

## Getting Started

### Prerequisites

- Go 1.23 or later
- OpenAI API credentials
- Homebrew (for dependency management)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kmesiab/think-tank.git
   cd think-tank
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up OpenAI API credentials:
   Ensure your OpenAI API key is set as an environment variable:
   ```bash
   export OPENAI_API_KEY="your-api-key-here"
   ```

4. Build the project:
   ```bash
   go build main.go
   ```

---

## Usage

Run the ThinkTank application with the following command:
```bash
./think-tank
```

### Customizing Experts

You can modify or add experts in the `main.go` file. Each expert is defined
as follows:
```go
var economicsExpert = &Expert{
CallOptions: callOptions,
LLM:         llm,
Name:        "Economics Expert",
Description: "A highly skilled economist with expertise in international
trade, economic growth, and policy analysis.",
Prompt:      "Given the following context: %s\\nCan you expand the question
to include related economic concepts to enhance the answer?",
}
```

To add a new expert, copy this structure and update the `Name`,
`Description`, and `Prompt` fields.

---

## Architecture

### Core Components

1. **ThinkTank**:
   - The main orchestrator that manages experts and synthesizes their outputs.

2. **Expert**:
   - Represents an AI model with a specific area of expertise, prompt, and
     evaluation logic.

3. **LLMs**:
   - Leverages the `langchaingo` library for seamless integration with
     OpenAI models.

### File Structure

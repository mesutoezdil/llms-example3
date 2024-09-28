# llms-example3

This project demonstrates how to interact with the Prediction Guard API using Go. The application sends chat completion requests to the API and processes the responses in real-time.

## Table of Contents
- [Introduction](#introduction)
- [Setup and Installation](#setup-and-installation)
- [Running the Project](#running-the-project)
- [Example Output](#example-output)
- [Troubleshooting](#troubleshooting)

## Introduction

This repository showcases an example of integrating the [Prediction Guard API](https://api.predictionguard.com) to generate natural language completions. Specifically, it uses a language model to respond to chat-style prompts, returning structured responses like code snippets.

## Setup and Installation

### Prerequisites

1. **Install Go**: 
   Ensure that you have Go installed. If not, download and install Go from [Go's official site](https://golang.org/doc/install).
   
2. **Obtain an API Key from Prediction Guard**: 
   You'll need an API key to authenticate requests to the Prediction Guard API. Store this key as an environment variable:
   ```bash
   export PGKEY="your_api_key_here"
   ```

### Project Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/satoshinakamoto/llms-example3.git
   cd llms-example3
   ```

2. **Initialize a Go Module**:
   Create a Go module to manage dependencies:
   ```bash
   go mod init github.com/satoshinakamoto/llms-example3
   ```
   Replace `github.com/satoshinakamoto/llms-example3` with the appropriate module path if necessary.

3. **Install Dependencies**:
   Use `go get` to download the required package:
   ```bash
   go get github.com/predictionguard/go-client
   ```

4. **Verify or Downgrade Package Version**:
   If you encounter issues with package compatibility, try downgrading to a specific version:
   ```bash
   go get github.com/predictionguard/go-client@v0.21.0
   ```

### Project Code

In the `main.go` file, the following steps are performed:

- A client is created to interact with the Prediction Guard API.
- A chat completion request is formed with a specific model and prompt.
- The program sends this request to the API and prints the response in real-time.

The most important sections of `main.go` include:
- **Model Selection**: `"Hermes-2-Pro-Llama-3-8B"` is used as the model for generating completions.
- **Request Input**: A system message and a user prompt are sent to the API.
- **Response Handling**: The API's response is printed to the console.

## Running the Project

To run the application and interact with the Prediction Guard API:

```bash
go run main.go
```

This command will execute the program, send a chat completion request, and print the response. Make sure that the environment variable `PGKEY` is set to your valid API key.

## Example Output

Upon running `main.go`, you will see logs indicating the progress of the API request and a response like the following:

```plaintext
2024/09/28 20:56:27 msg: do: rawRequest: started, method: POST, endpoint: https://api.predictionguard.com/chat/completions
2024/09/28 20:56:27 msg: do: rawRequest: completed, status: 200

Here's a simple Go program that prints out random numbers:

package main

import (
        "fmt"
        "math/rand"
        "time"
)

func main() {
        rand.Seed(time.Now().UnixNano())
        for i := 0; i < 10; i++ {
                fmt.Println(rand.Intn(100))
        }
}
```

This response is generated by the `"Hermes-2-Pro-Llama-3-8B"` model, and in this case, it provides a Go code snippet that prints random numbers.

## Troubleshooting

1. **Error: `model name does not exist`**: 
   - Double-check that the model name in `main.go` is correctly spelled and exists within the Prediction Guard API's supported models.

2. **API Key Issues**:
   - Make sure your `PGKEY` environment variable is correctly set and contains a valid API key.

3. **Type Errors**:
   - When setting `MaxTokens` or `Temperature` in `client.ChatSSEInput`, make sure you are passing pointers:
     ```go
     maxTokens := 1000
     temperature := float32(0.3)
     ```

4. **Dependency Conflicts**:
   - If you encounter dependency issues, try downgrading to an earlier version of `go-client` as mentioned in the setup steps.
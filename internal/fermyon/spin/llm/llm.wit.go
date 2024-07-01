// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package llm represents the imported interface "fermyon:spin/llm".
//
// A WASI interface dedicated to performing inferencing for Large Language Models.
package llm

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// InferencingModel represents the string "fermyon:spin/llm#inferencing-model".
//
// A Large Language Model.
//
//	type inferencing-model = string
type InferencingModel string

// InferencingParams represents the record "fermyon:spin/llm#inferencing-params".
//
// Inference request parameters
//
//	record inferencing-params {
//		max-tokens: u32,
//		repeat-penalty: f32,
//		repeat-penalty-last-n-token-count: u32,
//		temperature: f32,
//		top-k: u32,
//		top-p: f32,
//	}
type InferencingParams struct {
	// The maximum tokens that should be inferred.
	//
	// Note: the backing implementation may return less tokens.
	MaxTokens uint32

	// The amount the model should avoid repeating tokens.
	RepeatPenalty float32

	// The number of tokens the model should apply the repeat penalty to.
	RepeatPenaltyLastNTokenCount uint32

	// The randomness with which the next token is selected.
	Temperature float32

	// The number of possible next tokens the model will choose from.
	TopK uint32

	// The probability total of next tokens the model will choose from.
	TopP float32
}

// Error represents the variant "fermyon:spin/llm#error".
//
// The set of errors which may be raised by functions in this interface
//
//	variant error {
//		model-not-supported,
//		runtime-error(string),
//		invalid-input(string),
//	}
type Error cm.Variant[uint8, string, string]

// ErrorModelNotSupported returns a [Error] of case "model-not-supported".
func ErrorModelNotSupported() Error {
	var data struct{}
	return cm.New[Error](0, data)
}

// ModelNotSupported returns true if [Error] represents the variant case "model-not-supported".
func (self *Error) ModelNotSupported() bool {
	return self.Tag() == 0
}

// ErrorRuntimeError returns a [Error] of case "runtime-error".
func ErrorRuntimeError(data string) Error {
	return cm.New[Error](1, data)
}

// RuntimeError returns a non-nil *[string] if [Error] represents the variant case "runtime-error".
func (self *Error) RuntimeError() *string {
	return cm.Case[string](self, 1)
}

// ErrorInvalidInput returns a [Error] of case "invalid-input".
func ErrorInvalidInput(data string) Error {
	return cm.New[Error](2, data)
}

// InvalidInput returns a non-nil *[string] if [Error] represents the variant case "invalid-input".
func (self *Error) InvalidInput() *string {
	return cm.Case[string](self, 2)
}

// InferencingUsage represents the record "fermyon:spin/llm#inferencing-usage".
//
// Usage information related to the inferencing result
//
//	record inferencing-usage {
//		prompt-token-count: u32,
//		generated-token-count: u32,
//	}
type InferencingUsage struct {
	// Number of tokens in the prompt
	PromptTokenCount uint32

	// Number of tokens generated by the inferencing operation
	GeneratedTokenCount uint32
}

// InferencingResult represents the record "fermyon:spin/llm#inferencing-result".
//
// An inferencing result
//
//	record inferencing-result {
//		text: string,
//		usage: inferencing-usage,
//	}
type InferencingResult struct {
	// The text generated by the model
	// TODO: this should be a stream
	Text string

	// Usage information about the inferencing request
	Usage InferencingUsage
}

// EmbeddingModel represents the string "fermyon:spin/llm#embedding-model".
//
// The model used for generating embeddings
//
//	type embedding-model = string
type EmbeddingModel string

// EmbeddingsUsage represents the record "fermyon:spin/llm#embeddings-usage".
//
// Usage related to an embeddings generation request
//
//	record embeddings-usage {
//		prompt-token-count: u32,
//	}
type EmbeddingsUsage struct {
	// Number of tokens in the prompt
	PromptTokenCount uint32
}

// EmbeddingsResult represents the record "fermyon:spin/llm#embeddings-result".
//
// Result of generating embeddings
//
//	record embeddings-result {
//		embeddings: list<list<f32>>,
//		usage: embeddings-usage,
//	}
type EmbeddingsResult struct {
	// The embeddings generated by the request
	Embeddings cm.List[cm.List[float32]]

	// Usage related to the embeddings generation request
	Usage EmbeddingsUsage
}

// Infer represents the imported function "infer".
//
// Perform inferencing using the provided model and prompt with the given optional
// params
//
//	infer: func(model: inferencing-model, prompt: string, params: option<inferencing-params>)
//	-> result<inferencing-result, error>
//
//go:nosplit
func Infer(model InferencingModel, prompt string, params cm.Option[InferencingParams]) (result cm.OKResult[InferencingResult, Error]) {
	model0, model1 := cm.LowerString(model)
	prompt0, prompt1 := cm.LowerString(prompt)
	params0, params1, params2, params3, params4, params5, params6 := lower_OptionInferencingParams(params)
	wasmimport_Infer((*uint8)(model0), (uint32)(model1), (*uint8)(prompt0), (uint32)(prompt1), (uint32)(params0), (uint32)(params1), (float32)(params2), (uint32)(params3), (float32)(params4), (uint32)(params5), (float32)(params6), &result)
	return
}

//go:wasmimport fermyon:spin/llm infer
//go:noescape
func wasmimport_Infer(model0 *uint8, model1 uint32, prompt0 *uint8, prompt1 uint32, params0 uint32, params1 uint32, params2 float32, params3 uint32, params4 float32, params5 uint32, params6 float32, result *cm.OKResult[InferencingResult, Error])

// GenerateEmbeddings represents the imported function "generate-embeddings".
//
// Generate embeddings for the supplied list of text
//
//	generate-embeddings: func(model: embedding-model, text: list<string>) -> result<embeddings-result,
//	error>
//
//go:nosplit
func GenerateEmbeddings(model EmbeddingModel, text cm.List[string]) (result cm.OKResult[EmbeddingsResult, Error]) {
	model0, model1 := cm.LowerString(model)
	text0, text1 := cm.LowerList(text)
	wasmimport_GenerateEmbeddings((*uint8)(model0), (uint32)(model1), (*string)(text0), (uint32)(text1), &result)
	return
}

//go:wasmimport fermyon:spin/llm generate-embeddings
//go:noescape
func wasmimport_GenerateEmbeddings(model0 *uint8, model1 uint32, text0 *string, text1 uint32, result *cm.OKResult[EmbeddingsResult, Error])

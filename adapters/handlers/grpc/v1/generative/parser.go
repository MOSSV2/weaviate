//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package generative

import (
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
	pb "github.com/weaviate/weaviate/grpc/generated/protocol/v1"
	anthropicParams "github.com/weaviate/weaviate/modules/generative-anthropic/parameters"
	anyscaleParams "github.com/weaviate/weaviate/modules/generative-anyscale/parameters"
	awsParams "github.com/weaviate/weaviate/modules/generative-aws/parameters"
	cohereParams "github.com/weaviate/weaviate/modules/generative-cohere/parameters"
	mistralParams "github.com/weaviate/weaviate/modules/generative-mistral/parameters"
	octoaiParams "github.com/weaviate/weaviate/modules/generative-octoai/parameters"
	ollamaParams "github.com/weaviate/weaviate/modules/generative-ollama/parameters"
	openaiParams "github.com/weaviate/weaviate/modules/generative-openai/parameters"
	palmParams "github.com/weaviate/weaviate/modules/generative-palm/parameters"
	"github.com/weaviate/weaviate/usecases/modulecomponents/additional/generate"
)

type Parser struct {
	uses127Api     bool
	providerName   string
	returnMetadata bool
}

func NewParser(uses127Api bool) *Parser {
	return &Parser{
		uses127Api: uses127Api,
	}
}

func (p *Parser) Extract(req *pb.GenerativeSearch, class *models.Class) *generate.Params {
	if req == nil {
		return nil
	}
	if p.uses127Api {
		return p.extract(req, class)
	} else {
		return p.extractDeprecated(req, class)
	}
}

func (p *Parser) ProviderName() string {
	return p.providerName
}

func (p *Parser) ReturnMetadata() bool {
	return p.returnMetadata
}

func (p *Parser) extractDeprecated(req *pb.GenerativeSearch, class *models.Class) *generate.Params {
	generative := generate.Params{}
	if req.SingleResponsePrompt != "" {
		generative.Prompt = &req.SingleResponsePrompt
		singleResultPrompts := generate.ExtractPropsFromPrompt(generative.Prompt)
		generative.PropertiesToExtract = append(generative.PropertiesToExtract, singleResultPrompts...)
	}
	if req.GroupedResponseTask != "" {
		generative.Task = &req.GroupedResponseTask
		if len(req.GroupedProperties) > 0 {
			generative.Properties = req.GroupedProperties
			generative.PropertiesToExtract = append(generative.PropertiesToExtract, generative.Properties...)
		} else {
			// if users do not supply a properties, all properties need to be extracted
			generative.PropertiesToExtract = append(generative.PropertiesToExtract, schema.GetPropertyNamesFromClass(class, false)...)
		}
	}
	return &generative
}

func (p *Parser) extract(req *pb.GenerativeSearch, class *models.Class) *generate.Params {
	generative := generate.Params{}
	if req.Single != nil {
		generative.Prompt = &req.Single.Prompt
		singleResultPrompts := generate.ExtractPropsFromPrompt(generative.Prompt)
		generative.PropertiesToExtract = append(generative.PropertiesToExtract, singleResultPrompts...)
		if len(req.Single.Queries) > 0 {
			var options map[string]any
			var providerName string
			query := req.Single.Queries[0]
			switch query.Kind.(type) {
			case *pb.GenerativeProvider_Anthropic:
				options = p.anthropic(query.GetAnthropic())
				providerName = anthropicParams.Name
			case *pb.GenerativeProvider_Anyscale:
				options = p.anyscale(query.GetAnyscale())
				providerName = anyscaleParams.Name
			case *pb.GenerativeProvider_Aws:
				options = p.aws(query.GetAws())
				providerName = awsParams.Name
			case *pb.GenerativeProvider_Cohere:
				options = p.cohere(query.GetCohere())
				providerName = cohereParams.Name
			case *pb.GenerativeProvider_Mistral:
				options = p.mistral(query.GetMistral())
				providerName = mistralParams.Name
			case *pb.GenerativeProvider_Octoai:
				options = p.octoai(query.GetOctoai())
				providerName = octoaiParams.Name
			case *pb.GenerativeProvider_Ollama:
				options = p.ollama(query.GetOllama())
				providerName = ollamaParams.Name
			case *pb.GenerativeProvider_Openai:
				options = p.openai(query.GetOpenai())
				providerName = openaiParams.Name
			case *pb.GenerativeProvider_Palm:
				options = p.palm(query.GetPalm())
				providerName = palmParams.Name
			default:
				// do nothing
			}
			generative.Options = options
			p.providerName = providerName
			p.returnMetadata = query.ReturnMetadata
		}
	}
	if req.Grouped != nil {
		generative.Task = &req.Grouped.Task
		if req.Grouped.GetProperties() != nil {
			generative.Properties = req.Grouped.Properties.GetValues()
			generative.PropertiesToExtract = append(generative.PropertiesToExtract, generative.Properties...)
		} else {
			// if users do not supply a properties, all properties need to be extracted
			generative.PropertiesToExtract = append(generative.PropertiesToExtract, schema.GetPropertyNamesFromClass(class, false)...)
		}
	}
	return &generative
}

func (p *Parser) anthropic(in *pb.GenerativeAnthropic) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		anthropicParams.Name: anthropicParams.Params{
			BaseURL:       in.GetBaseUrl(),
			Model:         in.GetModel(),
			Temperature:   in.Temperature,
			MaxTokens:     p.int64ToInt(in.MaxTokens),
			StopSequences: in.StopSequences.GetValues(),
			TopP:          in.TopP,
			TopK:          p.int64ToInt(in.TopK),
		},
	}
}

func (p *Parser) anyscale(in *pb.GenerativeAnyscale) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		anyscaleParams.Name: anyscaleParams.Params{
			BaseURL:     in.GetBaseUrl(),
			Model:       in.GetModel(),
			Temperature: in.Temperature,
		},
	}
}

func (p *Parser) aws(in *pb.GenerativeAWS) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		awsParams.Name: awsParams.Params{
			Model:       in.GetModel(),
			Temperature: in.Temperature,
		},
	}
}

func (p *Parser) cohere(in *pb.GenerativeCohere) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		cohereParams.Name: cohereParams.Params{
			BaseURL:          in.GetBaseUrl(),
			Model:            in.GetModel(),
			Temperature:      in.Temperature,
			MaxTokens:        p.int64ToInt(in.MaxTokens),
			K:                p.int64ToInt(in.K),
			P:                in.P,
			StopSequences:    in.StopSequences.GetValues(),
			FrequencyPenalty: in.FrequencyPenalty,
			PresencePenalty:  in.PresencePenalty,
		},
	}
}

func (p *Parser) mistral(in *pb.GenerativeMistral) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		mistralParams.Name: mistralParams.Params{
			BaseURL:     in.GetBaseUrl(),
			MaxTokens:   p.int64ToInt(in.MaxTokens),
			Model:       in.GetModel(),
			Temperature: in.Temperature,
			TopP:        in.TopP,
		},
	}
}

func (p *Parser) octoai(in *pb.GenerativeOctoAI) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		octoaiParams.Name: octoaiParams.Params{
			BaseURL:     in.GetBaseUrl(),
			Model:       in.GetModel(),
			MaxTokens:   p.int64ToInt(in.MaxTokens),
			Temperature: in.Temperature,
			N:           p.int64ToInt(in.N),
			TopP:        in.TopP,
		},
	}
}

func (p *Parser) ollama(in *pb.GenerativeOllama) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		ollamaParams.Name: ollamaParams.Params{
			ApiEndpoint: in.GetApiEndpoint(),
			Model:       in.GetModel(),
			Temperature: in.Temperature,
		},
	}
}

func (p *Parser) openai(in *pb.GenerativeOpenAI) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		openaiParams.Name: openaiParams.Params{
			Model:            in.GetModel(),
			FrequencyPenalty: in.FrequencyPenalty,
			Logprobs:         in.LogProbs,
			TopLogprobs:      p.int64ToInt(in.TopLogProbs),
			MaxTokens:        p.int64ToInt(in.MaxTokens),
			N:                p.int64ToInt(in.N),
			PresencePenalty:  in.PresencePenalty,
			Stop:             in.Stop.GetValues(),
			Temperature:      in.Temperature,
			TopP:             in.TopP,
		},
	}
}

func (p *Parser) palm(in *pb.GenerativePaLM) map[string]any {
	if in == nil {
		return nil
	}
	return map[string]any{
		palmParams.Name: palmParams.Params{
			Model:            in.GetModel(),
			Temperature:      in.Temperature,
			MaxTokens:        p.int64ToInt(in.MaxTokens),
			TopP:             in.TopP,
			TopK:             p.int64ToInt(in.TopK),
			StopSequences:    in.StopSequences.GetValues(),
			PresencePenalty:  in.PresencePenalty,
			FrequencyPenalty: in.FrequencyPenalty,
		},
	}
}

func (p *Parser) int64ToInt(in *int64) *int {
	if in != nil && *in > 0 {
		out := int(*in)
		return &out
	}
	return nil
}

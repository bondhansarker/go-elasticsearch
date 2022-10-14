// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// InferenceProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/ingest/_types/Processors.ts#L237-L242
type InferenceProcessor struct {
	FieldMap        map[Field]interface{} `json:"field_map,omitempty"`
	If              *string               `json:"if,omitempty"`
	IgnoreFailure   *bool                 `json:"ignore_failure,omitempty"`
	InferenceConfig *InferenceConfig      `json:"inference_config,omitempty"`
	ModelId         Id                    `json:"model_id"`
	OnFailure       []ProcessorContainer  `json:"on_failure,omitempty"`
	Tag             *string               `json:"tag,omitempty"`
	TargetField     Field                 `json:"target_field"`
}

// InferenceProcessorBuilder holds InferenceProcessor struct and provides a builder API.
type InferenceProcessorBuilder struct {
	v *InferenceProcessor
}

// NewInferenceProcessor provides a builder for the InferenceProcessor struct.
func NewInferenceProcessorBuilder() *InferenceProcessorBuilder {
	r := InferenceProcessorBuilder{
		&InferenceProcessor{
			FieldMap: make(map[Field]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InferenceProcessor struct
func (rb *InferenceProcessorBuilder) Build() InferenceProcessor {
	return *rb.v
}

func (rb *InferenceProcessorBuilder) FieldMap(value map[Field]interface{}) *InferenceProcessorBuilder {
	rb.v.FieldMap = value
	return rb
}

func (rb *InferenceProcessorBuilder) If_(if_ string) *InferenceProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *InferenceProcessorBuilder) IgnoreFailure(ignorefailure bool) *InferenceProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *InferenceProcessorBuilder) InferenceConfig(inferenceconfig *InferenceConfigBuilder) *InferenceProcessorBuilder {
	v := inferenceconfig.Build()
	rb.v.InferenceConfig = &v
	return rb
}

func (rb *InferenceProcessorBuilder) ModelId(modelid Id) *InferenceProcessorBuilder {
	rb.v.ModelId = modelid
	return rb
}

func (rb *InferenceProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *InferenceProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *InferenceProcessorBuilder) Tag(tag string) *InferenceProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *InferenceProcessorBuilder) TargetField(targetfield Field) *InferenceProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}

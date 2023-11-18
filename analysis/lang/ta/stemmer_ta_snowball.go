//  Copyright (c) 2020 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ta

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/snowballstem/tamil"

	"github.com/blevesearch/snowballstem"
)

const SnowballStemmerName = "stemmer_ta_snowball"

type TamilStemmerFilter struct {
}

func NewTamilStemmerFilter() *TamilStemmerFilter {
	return &TamilStemmerFilter{}
}

func (s *TamilStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	for _, token := range input {
		env := snowballstem.NewEnv(string(token.Term))
		tamil.Stem(env)
		token.Term = []byte(env.Current())
	}
	return input
}

func TamilStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	return NewTamilStemmerFilter(), nil
}

func init() {
	registry.RegisterTokenFilter(SnowballStemmerName, TamilStemmerFilterConstructor)
}

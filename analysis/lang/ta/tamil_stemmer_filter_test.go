//  Copyright (c) 2014 Couchbase, Inc.
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
	"reflect"
	"testing"

	"github.com/blevesearch/bleve/v2/analysis"
)

func TestTamilStemmerFilter(t *testing.T) {
	tests := []struct {
		input  analysis.TokenStream
		output analysis.TokenStream
	}{
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("தருமரிடம்"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("தருமர்"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("தருமரோடு"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("தருமர்"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("நடந்தான்"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("நட"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவன்"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவன்"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவனுடைய"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவன்"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("சென்றுகொண்டிருந்த"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("சென்றுகொண்டிரு"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவனாக"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவனா"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவனாக"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("அவனா"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("வளமாக"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("வளமா"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("நிற்கின்ற"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("நி"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("பறக்கிற"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("பற"),
				},
			},
		},
	}

	tamilStemmerFilter := NewTamilStemmerFilter()
	for _, test := range tests {
		actual := tamilStemmerFilter.Filter(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %#v, got %#v", test.output, actual)
			t.Errorf("expected % x, got % x", test.output[0].Term, actual[0].Term)
		}
	}
}

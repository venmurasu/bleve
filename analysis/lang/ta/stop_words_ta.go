package ta

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const StopName = "stop_ta"

// The stop words list is from https://github.com/AshokR/TamilNLP/blob/master/Resources/TamilStopWords.txt
// under Apache 2.0 license https://github.com/AshokR/TamilNLP/blob/master/LICENSE

var TamilStopWords = []byte(`
ஒரு
என்று
மற்றும்
இந்த
இது
என்ற
கொண்டு
என்பது
பல
ஆகும்
அல்லது
அவர்
நான்
உள்ள
அந்த
இவர்
என
முதல்
என்ன
இருந்து
சில
என்
போன்ற
வேண்டும்
வந்து
இதன்
அது
அவன்
தான்
பலரும்
என்னும்
மேலும்
பின்னர்
கொண்ட
இருக்கும்
தனது
உள்ளது
போது
என்றும்
அதன்
தன்
பிறகு
அவர்கள்
வரை
அவள்
நீ
ஆகிய
இருந்தது
உள்ளன
வந்த
இருந்த
மிகவும்
இங்கு
மீது
ஓர்
இவை
இந்தக்
பற்றி
வரும்
வேறு
இரு
இதில்
போல்
இப்போது
அவரது
மட்டும்
இந்தப்
எனும்
மேல்
பின்
சேர்ந்த
ஆகியோர்
எனக்கு
இன்னும்
அந்தப்
அன்று
ஒரே
மிக
அங்கு
பல்வேறு
விட்டு
பெரும்
அதை
பற்றிய
உன்
அதிக
அந்தக்
பேர்
இதனால்
அவை
அதே
ஏன்
முறை
யார்
என்பதை
எல்லாம்
மட்டுமே
இங்கே
அங்கே
இடம்
இடத்தில்
அதில்
நாம்
அதற்கு
எனவே
பிற
சிறு
மற்ற
விட
எந்த
எனவும்
எனப்படும்
எனினும்
அடுத்த
இதனை
இதை
கொள்ள
இந்தத்
இதற்கு
அதனால்
தவிர
போல
வரையில்
சற்று
எனக்
`)

func TokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	rv := analysis.NewTokenMap()
	err := rv.LoadBytes(TamilStopWords)
	return rv, err
}

func init() {
	registry.RegisterTokenMap(StopName, TokenMapConstructor)
}

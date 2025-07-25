package embeddata

import _ "embed"

//go:embed config/dict.json
var DefaultDict []byte

func GetDefaultDict() []byte {
	return DefaultDict
}

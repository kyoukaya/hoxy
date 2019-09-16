package xdefs

import (
	"encoding/json"
	"hoxy/proxy/defs/types"
)

// ResetBuildCoinData is a custom type for handling ambiguous struct or empty json array
// encoded int data in json
type ResetBuildCoinData struct {
	MaxBuildCoin types.Int
	exists       bool
}

type resetBuildCoinData struct {
	MaxBuildCoin types.Int `json:"max_build_coin"`
}

func (ioe *ResetBuildCoinData) UnmarshalJSON(data []byte) error {
	if data[0] == '[' {
		return nil
	}
	ioe.exists = true

	underlyingStruct := resetBuildCoinData{}
	err := json.Unmarshal(data, &underlyingStruct)
	ioe.MaxBuildCoin = underlyingStruct.MaxBuildCoin
	if err != nil {
		return err
	}
	return nil
}

func (ioe *ResetBuildCoinData) MarshalJSON() ([]byte, error) {
	if !ioe.exists {
		return []byte(`[]`), nil
	}
	underlyingStruct := resetBuildCoinData{ioe.MaxBuildCoin}
	return json.Marshal(&underlyingStruct)
}

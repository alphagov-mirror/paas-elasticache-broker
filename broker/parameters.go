package broker

import (
	"encoding/json"
	"fmt"
)

const ParamRestoreLatestSnapshotOf = "restore_from_latest_snapshot_of"
const ParamMaxMemoryPolicy = "maxmemory_policy"

func ParseProvisionParameters(data []byte) (*ProvisionParameters, error) {
	mapParams := map[string]interface{}{}
	err := json.Unmarshal(data, &mapParams)
	if err != nil {
		return nil, err
	}
	validKeys := []string{ParamRestoreLatestSnapshotOf, ParamMaxMemoryPolicy}
	for key := range mapParams {
		valid := false
		for _, validKey := range validKeys {
			if validKey == key {
				valid = true
				break
			}
		}
		if !valid {
			return nil, fmt.Errorf("unknown parameter: %s", key)
		}
	}
	provisionParameters := &ProvisionParameters{}
	if err := json.Unmarshal(data, provisionParameters); err != nil {
		return nil, err
	}
	return provisionParameters, nil
}

type ProvisionParameters struct {
	RestoreFromLatestSnapshotOf *string `json:"restore_from_latest_snapshot_of"`
	MaxMemoryPolicy             *string `json:"maxmemory_policy"`
}

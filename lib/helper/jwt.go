package helper

import (
	"context"
	"encoding/json"
	context2 "github.com/kodekoding/phastos/v2/go/context"
	"godem/domain/models/user"
)

func GetJWTData(ctx context.Context) *user.JWTData {
	getJWTData := context2.GetJWT(ctx)
	if getJWTData == nil {
		return nil
	}

	jwtData, valid := getJWTData.Data.(map[string]interface{})
	if !valid {
		return nil
	}

	b, _ := json.Marshal(jwtData)
	var result user.JWTData
	_ = json.Unmarshal(b, &result)
	return &result
}

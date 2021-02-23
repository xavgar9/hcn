package helpers

import (
	"encoding/json"
	"hcn/mymodels"
)

// CleanHCN Create the same struct of the HCN without
// the mongo _id
func CleanHCN(hcn mymodels.HCNmongo) mymodels.HCNmongoNoID {
	var newHCN mymodels.HCNmongoNoID
	hcnJSON, _ := json.Marshal(hcn)
	json.Unmarshal([]byte(hcnJSON), &newHCN)
	return newHCN
}

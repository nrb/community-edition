package convert

import (
	"os"

	tkrtypes "github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkr/pkg/types"
)

func ReadTKR(filePath string) (tkrtypes.Bom, error) {
	tkr := tkrtypes.Bom{}
	tkrData, err := os.ReadFile(filePath)

	if err != nil {
		return tkr, err
	}

	tkr, err = tkrtypes.NewBom(tkrData)
	if err != nil {
		return tkrtypes.Bom{}, err
	}

	return tkr, nil
}

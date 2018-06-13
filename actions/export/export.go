package export

import (
	"errors"
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/export"
	exportEnum "github.com/chargebee/chargebee-go/models/export/enum"
	"time"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/exports/%v", id), nil)
}

func RevenueRecognition(params *export.RevenueRecognitionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/revenue_recognition"), params)
}

func DeferredRevenue(params *export.DeferredRevenueRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/deferred_revenue"), params)
}
func WaitForExportCompletion(exp export.Export) (export.Export, error) {
	return WaitForExportCompletionWithEnv(exp, chargebee.DefaultConfig())
}
func WaitForExportCompletionWithEnv(exp export.Export, env chargebee.Environment) (export.Export, error) {
	count := 0
	for exp.Status == exportEnum.StatusInProcess {
		if count > 30 {
			return exp, errors.New("'Export is taking too long'")
		}
		count++
		time.Sleep(chargebee.ExportWaitInSecs)
		result, err := Retrieve(exp.Id).RequestWithEnv(env)
		if err != nil {
			return exp, err
		}
		exp = *result.Export
	}
	return exp, nil
}

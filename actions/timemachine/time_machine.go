package timemachine

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/timemachine"
	timeMachineEnum "github.com/chargebee/chargebee-go/v3/models/timemachine/enum"
	"net/url"
	"time"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/time_machines/%v", url.PathEscape(id)), nil)
}
func StartAfresh(id string, params *timemachine.StartAfreshRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/time_machines/%v/start_afresh", url.PathEscape(id)), params).SetIdempotency(true)
}
func TravelForward(id string, params *timemachine.TravelForwardRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/time_machines/%v/travel_forward", url.PathEscape(id)), params).SetIdempotency(true)
}
func WaitForTimeTravelCompletion(tm timemachine.TimeMachine) (timemachine.TimeMachine, error) {
	return WaitForTimeTravelCompletionWithEnv(tm, chargebee.DefaultConfig())
}
func WaitForTimeTravelCompletionWithEnv(tm timemachine.TimeMachine, env chargebee.Environment) (timemachine.TimeMachine, error) {
	count := 0
	for tm.TimeTravelStatus == timeMachineEnum.TimeTravelStatusInProgress {
		if count > 30 {
			panic(errors.New("time travel is taking too much time"))
		}
		count++
		time.Sleep(chargebee.TimeMachineWaitInSecs)
		result, err := Retrieve(tm.Name).RequestWithEnv(env)
		if err != nil {
			panic(err)
		}
		tm = *result.TimeMachine
	}
	if tm.TimeTravelStatus == timeMachineEnum.TimeTravelStatusFailed {
		error := &chargebee.Error{}
		err := json.Unmarshal([]byte(tm.ErrorJson), &error)
		if err != nil {
			panic(err)
		}
		return tm, error
	}
	if tm.TimeTravelStatus != timeMachineEnum.TimeTravelStatusInProgress && tm.TimeTravelStatus != timeMachineEnum.TimeTravelStatusSucceeded && tm.TimeTravelStatus != timeMachineEnum.TimeTravelStatusFailed {
		return tm, errors.New("time travel state is in wrong state  \"" + string(tm.TimeTravelStatus) + "\"")
	}
	return tm, nil
}

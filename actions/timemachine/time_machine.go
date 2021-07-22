package timemachine

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/timemachine"
	timeMachineEnum "github.com/chargebee/chargebee-go/models/timemachine/enum"
	"time"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/time_machines/%v", id), nil)
}
func StartAfresh(id string, params *timemachine.StartAfreshRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/time_machines/%v/start_afresh", id), params)
}
func TravelForward(id string, params *timemachine.TravelForwardRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/time_machines/%v/travel_forward", id), params)
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

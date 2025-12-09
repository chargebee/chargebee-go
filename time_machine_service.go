package chargebee

import (
	"encoding/json"
	"fmt"

	"errors"
	"time"

	"net/url"
)

type TimeMachineService struct {
	transport *Transport
}

func (s *TimeMachineService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/time_machines/%v", url.PathEscape(id)), nil)
}

func (s *TimeMachineService) StartAfresh(id string, req *StartAfreshRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/time_machines/%v/start_afresh", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *TimeMachineService) TravelForward(id string, req *TravelForwardRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/time_machines/%v/travel_forward", url.PathEscape(id)), req).SetIdempotency(true)
}

func WaitForTimeTravelCompletion(tm TimeMachine) (TimeMachine, error) {
	return WaitForTimeTravelCompletionWithEnv(tm, chargebee.DefaultConfig())
}

func WaitForTimeTravelCompletionWithEnv(tm TimeMachine, env chargebee.Environment) (TimeMachine, error) {
	count := 0
	for tm.TimeTravelStatus == TimeTravelStatusInProgress {
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
	if tm.TimeTravelStatus == TimeTravelStatusFailed {
		error := &chargebee.Error{}
		err := json.Unmarshal([]byte(tm.ErrorJson), &error)
		if err != nil {
			panic(err)
		}
		return tm, error
	}
	if tm.TimeTravelStatus != TimeTravelStatusInProgress && tm.TimeTravelStatus != TimeTravelStatusSucceeded && tm.TimeTravelStatus != TimeTravelStatusFailed {
		return tm, errors.New("time travel state is in wrong state  \"" + string(tm.TimeTravelStatus) + "\"")
	}
	return tm, nil
}

package exestate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SuccessCode      = 0
	ControlledCode   = 1
	UncontrolledCode = 2
)

type StateHandler interface {
	RegisterState(*State)
	GetState(consume bool) *State
}

type State struct {
	Code      int
	UserError string
	err       error
}

func UncontrolledError(UserError string, err error) *State {
	//fmt.Println("User err: " + UserError + " --- errTxt:" + err.Error())
	return &State{UncontrolledCode, "User err: " + UserError + " --- errTxt: " + err.Error(), err}
}

func ControlledError(UserError string) *State {
	//fmt.Println("User err: " + UserError)
	return &State{ControlledCode, UserError, nil}
}

func Ok() *State {
	return &State{SuccessCode, "OK", nil}
}

func (this *State) IsOk() bool {
	return this.Code == SuccessCode
}

func StateEvalJson(w http.ResponseWriter, sth StateHandler) bool {

	if OnError(sth) {
		fmt.Println(sth.GetState(false).UserError)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(sth.GetState(false).UserError)

		return true
	}

	return false
}

func StateEval(sth StateHandler, err error, userError string) bool {

	if OnError(sth) {
		fmt.Println(sth.GetState(false).UserError)

		return true
	}

	return false
}
func OnError(sth StateHandler) bool {
	if sth == nil || sth.GetState(false) == nil || sth.GetState(false).Code == SuccessCode {
		return false
	}

	return true
}

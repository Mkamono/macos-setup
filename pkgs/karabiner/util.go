package main

import (
	"karabiner/core"
)

type requirement struct {
	keyCode    string
	modifiers  []string
	conditions []core.Condition
}

type action struct {
	to           core.To
	toAfterKeyUp core.ToAfterKeyUp
}

func simpleManipulator(requirement requirement, action action) core.Manipulator {
	if action.toAfterKeyUp.SetVariable == nil {
		action.toAfterKeyUp = core.ToAfterKeyUp{}
	}

	var toAfterKeyUp []core.ToAfterKeyUp
	if action.toAfterKeyUp.SetVariable != nil {
		toAfterKeyUp = append(toAfterKeyUp, action.toAfterKeyUp)
	}

	return core.Manipulator{
		From: core.From{
			KeyCode: requirement.keyCode,
			Modifiers: &core.Modifiers{
				Mandatory: requirement.modifiers,
				Optional:  []string{core.ANY},
			},
		},
		Conditions:   requirement.conditions,
		To:           []core.To{action.to},
		ToAfterKeyUp: toAfterKeyUp,
		Type:         core.BASIC,
	}
}

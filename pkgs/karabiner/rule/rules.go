package rule

import (
	"karabiner/core"
	"slices"

	"github.com/samber/lo"
)

type record[T any] struct {
	key   string
	value T
}

type orderedMap[T any] []record[T]

// WASDと矢印キーの対応
func wasdRule() core.Rule {
	trigger := core.LEFT_CONTROL
	manipulators := lo.Map(orderedMap[string]{
		{"w", core.UP_ARROW},
		{"s", core.DOWN_ARROW},
		{"a", core.LEFT_ARROW},
		{"d", core.RIGHT_ARROW},
	}, func(item record[string], index int) core.Manipulator {
		return simpleManipulator(
			requirement{
				keyCode:   item.key,
				modifiers: []string{trigger},
			},
			action{
				to: core.To{KeyCode: item.value},
			})
	})

	return core.Rule{
		Description:  "WASD Arrow",
		Manipulators: manipulators,
	}
}

func escAndDeleteRule() core.Rule {
	trigger := core.LEFT_CONTROL
	manipulators := lo.Map(orderedMap[string]{
		{core.OPEN_BRACKET, core.ESCAPE},
		{"h", core.DELETE_OR_BACKSPACE},
	}, func(item record[string], index int) core.Manipulator {
		return simpleManipulator(
			requirement{
				keyCode:   item.key,
				modifiers: []string{trigger},
			},
			action{
				to: core.To{KeyCode: item.value},
			})
	})

	return core.Rule{
		Description:  "ESC and DELETE",
		Manipulators: manipulators,
	}
}

// マウスのフルエミュレーション
func mouseEmulationRule() core.Rule {
	// マウスポインタの移動
	mouseCondition := []core.Condition{
		{
			Name:  core.MOUSE_KEYS_FULL,
			Type:  core.VARIABLE_IF,
			Value: 1,
		},
	}

	mouseManipulators := lo.Map(orderedMap[core.MouseKey]{
		{"a", core.MouseKey{X: -1536}},
		{"s", core.MouseKey{Y: 1536}},
		{"w", core.MouseKey{Y: -1536}},
		{"d", core.MouseKey{X: 1536}},
	}, func(item record[core.MouseKey],
		index int) core.Manipulator {
		return simpleManipulator(
			requirement{
				keyCode:    item.key,
				modifiers:  []string{core.RIGHT_SHIFT},
				conditions: mouseCondition,
			},
			action{
				to: core.To{MouseKey: &item.value},
			})
	})

	mouseSpeedManipulators := lo.Map(orderedMap[core.MouseKey]{
		{"n", core.MouseKey{SpeedMultiplier: 2.0}},
		{"m", core.MouseKey{SpeedMultiplier: 0.3}},
	}, func(item record[core.MouseKey],
		index int) core.Manipulator {
		return simpleManipulator(
			requirement{
				keyCode:    item.key,
				modifiers:  []string{core.RIGHT_SHIFT},
				conditions: mouseCondition,
			},
			action{
				to: core.To{MouseKey: &item.value},
			})
	})

	// マウスボタン
	mouseButtonManipulators := lo.Map(orderedMap[string]{
		{"j", core.BUTTON1},
		{"k", core.BUTTON3},
		{"l", core.BUTTON2},
	}, func(item record[string],
		index int) core.Manipulator {
		return simpleManipulator(
			requirement{
				keyCode:    item.key,
				modifiers:  []string{core.RIGHT_SHIFT},
				conditions: mouseCondition,
			},
			action{
				to: core.To{PointingButton: item.value},
			})
	})

	// マウススクロール
	scrollConditions := []core.Condition{
		{
			Name:  core.MOUSE_KEYS_FULL,
			Type:  core.VARIABLE_IF,
			Value: 1,
		},
		{
			Name:  core.MOUSE_KEYS_FULL_SCROLL,
			Type:  core.VARIABLE_IF,
			Value: 1,
		},
	}

	scrollManipulators := lo.Map(orderedMap[core.MouseKey]{
		{"a", core.MouseKey{HorizontalWheel: 32}},
		{"s", core.MouseKey{VerticalWheel: 32}},
		{"w", core.MouseKey{VerticalWheel: -32}},
		{"d", core.MouseKey{HorizontalWheel: -32}},
	}, func(item record[core.MouseKey],
		index int) core.Manipulator {
		return simpleManipulator(
			requirement{
				keyCode:    item.key,
				modifiers:  []string{core.RIGHT_SHIFT},
				conditions: scrollConditions,
			},
			action{
				to: core.To{MouseKey: &item.value},
			})
	})

	// 右Shiftキーを押している間だけマウスフルエミュレーションを有効にする
	mouseTriggerManipulator := []core.Manipulator{simpleManipulator(
		requirement{
			keyCode: core.RIGHT_SHIFT,
		},
		action{
			to:           core.To{KeyCode: core.RIGHT_SHIFT},
			toAfterKeyUp: core.ToAfterKeyUp{SetVariable: &core.SetVariable{Name: core.MOUSE_KEYS_FULL, Value: 1}},
		},
	)}

	scrollTriggerManipulator := []core.Manipulator{simpleManipulator(
		requirement{
			keyCode:    core.SEMICOLON,
			modifiers:  []string{core.RIGHT_SHIFT},
			conditions: mouseCondition,
		},
		action{
			to: core.To{SetVariable: &core.SetVariable{Name: core.MOUSE_KEYS_FULL_SCROLL, Value: 1}},
			toAfterKeyUp: core.ToAfterKeyUp{
				SetVariable: &core.SetVariable{
					Name:  core.MOUSE_KEYS_FULL_SCROLL,
					Value: 0,
				},
			},
		},
	)}

	return core.Rule{
		Description: "Mouse full emulation with right command super fast",
		Manipulators: slices.Concat(
			mouseTriggerManipulator,
			scrollManipulators,
			mouseManipulators,
			mouseButtonManipulators,
			mouseSpeedManipulators,
			scrollTriggerManipulator,
		),
	}
}

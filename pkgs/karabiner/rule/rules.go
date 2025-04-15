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

var arrowLiteralKey = struct {
	LEFT_ARROW  string
	RIGHT_ARROW string
	UP_ARROW    string
	DOWN_ARROW  string
}{
	LEFT_ARROW:  "s",
	RIGHT_ARROW: "f",
	UP_ARROW:    "e",
	DOWN_ARROW:  "d",
}

// 矢印キーの設定
func arrowKeyRule() core.Rule {
	trigger := core.LEFT_CONTROL
	manipulators := lo.Map(orderedMap[string]{
		{arrowLiteralKey.UP_ARROW, core.UP_ARROW},
		{arrowLiteralKey.DOWN_ARROW, core.DOWN_ARROW},
		{arrowLiteralKey.LEFT_ARROW, core.LEFT_ARROW},
		{arrowLiteralKey.RIGHT_ARROW, core.RIGHT_ARROW},
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
		Description:  "Arrow Keys",
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
	// 移動量
	pointerMoveAmount := 1536
	wheelMoveAmount := 32
	speedMultiplier := 2.0
	slowMouseSpeed := 0.3

	// マウスポインタの移動
	mouseCondition := []core.Condition{
		{
			Name:  core.MOUSE_KEYS_FULL,
			Type:  core.VARIABLE_IF,
			Value: 1,
		},
	}

	mouseManipulators := lo.Map(orderedMap[core.MouseKey]{
		{arrowLiteralKey.LEFT_ARROW, core.MouseKey{X: -pointerMoveAmount}},
		{arrowLiteralKey.DOWN_ARROW, core.MouseKey{Y: pointerMoveAmount}},
		{arrowLiteralKey.UP_ARROW, core.MouseKey{Y: -pointerMoveAmount}},
		{arrowLiteralKey.RIGHT_ARROW, core.MouseKey{X: pointerMoveAmount}},
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
		{"n", core.MouseKey{SpeedMultiplier: speedMultiplier}},
		{"m", core.MouseKey{SpeedMultiplier: slowMouseSpeed}},
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
		{arrowLiteralKey.LEFT_ARROW, core.MouseKey{HorizontalWheel: wheelMoveAmount}},
		{arrowLiteralKey.DOWN_ARROW, core.MouseKey{VerticalWheel: wheelMoveAmount}},
		{arrowLiteralKey.UP_ARROW, core.MouseKey{VerticalWheel: -wheelMoveAmount}},
		{arrowLiteralKey.RIGHT_ARROW, core.MouseKey{HorizontalWheel: -wheelMoveAmount}},
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

	mouseTriggerManipulator := []core.Manipulator{simpleManipulator(
		requirement{
			keyCode: core.RIGHT_SHIFT,
			// keyCode
		},
		action{
			to:           core.To{KeyCode: core.RIGHT_SHIFT},
			toAfterKeyUp: core.ToAfterKeyUp{SetVariable: &core.SetVariable{Name: core.MOUSE_KEYS_FULL, Value: 1}},
		},
	)}

	scrollTriggerManipulator := []core.Manipulator{simpleManipulator(
		requirement{
			keyCode:    core.SEMICOLON,
			modifiers:  []string{core.RIGHT_SHIFT}, // 後でこれはquoteに変えたい
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

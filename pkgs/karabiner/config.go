package main

import "karabiner/core"

func NewConfig() core.KarabinerConfig {
	return core.KarabinerConfig{
		Profiles: []core.Profile{
			{
				ComplexModifications: core.ComplexModifications{
					Rules: []core.Rule{
						wasdRule(),
						mouseEmulationRule(),
						// 	Description: "Mouse full emulation with right command super fast",
						// 	Manipulators: []core.Manipulator{
						// 		{
						// 			From: core.From{
						// 				KeyCode: core.RIGHT_SHIFT,
						// 				Modifiers: &core.Modifiers{
						// 					Optional: []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{KeyCode: core.RIGHT_SHIFT},
						// 			},
						// 			ToAfterKeyUp: []core.ToAfterKeyUp{
						// 				{
						// 					SetVariable: &core.SetVariable{
						// 						Name:  "mouse_keys_full",
						// 						Value: 1,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 				{
						// 					Name:  "mouse_keys_full_scroll",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "a",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						HorizontalWheel: 32,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 				{
						// 					Name:  "mouse_keys_full_scroll",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "s",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						VerticalWheel: 32,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 				{
						// 					Name:  "mouse_keys_full_scroll",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "w",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						VerticalWheel: -32,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 				{
						// 					Name:  "mouse_keys_full_scroll",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "d",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						HorizontalWheel: -32,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "a",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						X: -1536,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "s",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						Y: 1536,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "w",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						Y: -1536,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "d",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						X: 1536,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "j",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					PointingButton: core.BUTTON1,
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "k",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					PointingButton: core.BUTTON3,
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "l",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					PointingButton: core.BUTTON2,
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "n",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						SpeedMultiplier: 2.0,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: "m",
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					MouseKey: &core.MouseKey{
						// 						SpeedMultiplier: 0.3,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 		{
						// 			Conditions: []core.Condition{
						// 				{
						// 					Name:  "mouse_keys_full",
						// 					Type:  "variable_if",
						// 					Value: 1,
						// 				},
						// 			},
						// 			From: core.From{
						// 				KeyCode: core.SEMICOLON,
						// 				Modifiers: &core.Modifiers{
						// 					Mandatory: []string{core.RIGHT_SHIFT},
						// 					Optional:  []string{core.ANY},
						// 				},
						// 			},
						// 			To: []core.To{
						// 				{
						// 					SetVariable: &core.SetVariable{
						// 						Name:  "mouse_keys_full_scroll",
						// 						Value: 1,
						// 					},
						// 				},
						// 			},
						// 			ToAfterKeyUp: []core.ToAfterKeyUp{
						// 				{
						// 					SetVariable: &core.SetVariable{
						// 						Name:  "mouse_keys_full_scroll",
						// 						Value: 0,
						// 					},
						// 				},
						// 			},
						// 			Type: core.BASIC,
						// 		},
						// 	},
						// },
					},
				},
				Name:     "Default profile",
				Selected: true,
				SimpleModifications: []core.SimpleModification{
					{
						From: core.From{
							KeyCode: core.CAPS_LOCK,
						},
						To: []core.To{
							{KeyCode: core.LEFT_CONTROL},
						},
					},
				},
				VirtualHidKeyboard: core.VirtualHidKeyboard{
					KeyboardType: core.ANSI,
				},
			},
		},
	}
}

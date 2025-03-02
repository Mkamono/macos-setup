package main

import (
	"karabiner/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKarabinerConfig(t *testing.T) {
	assert.Equal(t, core.KarabinerConfig{
		Profiles: []core.Profile{
			{
				ComplexModifications: core.ComplexModifications{
					Rules: []core.Rule{
						{
							Description: "WASD Arrow",
							Manipulators: []core.Manipulator{
								{
									From: core.From{
										KeyCode: "w",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"left_control"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "up_arrow"},
									},
									Type: "basic",
								},
								{
									From: core.From{
										KeyCode: "s",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"left_control"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "down_arrow"},
									},
									Type: "basic",
								},
								{
									From: core.From{
										KeyCode: "a",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"left_control"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "left_arrow"},
									},
									Type: "basic",
								},
								{
									From: core.From{
										KeyCode: "d",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"left_control"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "right_arrow"},
									},
									Type: "basic",
								},
								{
									From: core.From{
										KeyCode: "open_bracket",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"left_control"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "escape"},
									},
									Type: "basic",
								},
								{
									From: core.From{
										KeyCode: "h",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"left_control"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "delete_or_backspace"},
									},
									Type: "basic",
								},
							},
						},
						{
							Description: "Mouse full emulation with right command super fast",
							Manipulators: []core.Manipulator{
								{
									From: core.From{
										KeyCode: "right_shift",
										Modifiers: &core.Modifiers{
											Optional: []string{"any"},
										},
									},
									To: []core.To{
										{KeyCode: "right_shift"},
									},
									ToAfterKeyUp: []core.ToAfterKeyUp{
										{
											SetVariable: &core.SetVariable{
												Name:  "mouse_keys_full",
												Value: 1,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
										{
											Name:  "mouse_keys_full_scroll",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "a",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												HorizontalWheel: 32,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
										{
											Name:  "mouse_keys_full_scroll",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "s",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												VerticalWheel: 32,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
										{
											Name:  "mouse_keys_full_scroll",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "w",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												VerticalWheel: -32,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
										{
											Name:  "mouse_keys_full_scroll",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "d",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												HorizontalWheel: -32,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "a",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												X: -1536,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "s",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												Y: 1536,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "w",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												Y: -1536,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "d",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												X: 1536,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "j",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											PointingButton: "button1",
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "k",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											PointingButton: "button3",
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "l",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											PointingButton: "button2",
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "n",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												SpeedMultiplier: 2.0,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "m",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											MouseKey: &core.MouseKey{
												SpeedMultiplier: 0.3,
											},
										},
									},
									Type: "basic",
								},
								{
									Conditions: []core.Condition{
										{
											Name:  "mouse_keys_full",
											Type:  "variable_if",
											Value: 1,
										},
									},
									From: core.From{
										KeyCode: "semicolon",
										Modifiers: &core.Modifiers{
											Mandatory: []string{"right_shift"},
											Optional:  []string{"any"},
										},
									},
									To: []core.To{
										{
											SetVariable: &core.SetVariable{
												Name:  "mouse_keys_full_scroll",
												Value: 1,
											},
										},
									},
									ToAfterKeyUp: []core.ToAfterKeyUp{
										{
											SetVariable: &core.SetVariable{
												Name:  "mouse_keys_full_scroll",
												Value: 0,
											},
										},
									},
									Type: "basic",
								},
							},
						},
					},
				},
				Name:     "Default profile",
				Selected: true,
				SimpleModifications: []core.SimpleModification{
					{
						From: core.From{
							KeyCode: "caps_lock",
						},
						To: []core.To{
							{KeyCode: "left_control"},
						},
					},
				},
				VirtualHidKeyboard: core.VirtualHidKeyboard{
					KeyboardType: "ansi",
				},
			},
		},
	}, NewConfig())
}

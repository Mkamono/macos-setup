package rule

import "karabiner/core"

func NewConfig() core.KarabinerConfig {
	return core.KarabinerConfig{
		Profiles: []core.Profile{
			{
				ComplexModifications: core.ComplexModifications{
					Rules: []core.Rule{
						arrowKeyRule(),
						escAndDeleteRule(),
						mouseEmulationRule(),
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

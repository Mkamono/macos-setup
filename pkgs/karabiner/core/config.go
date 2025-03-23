package core

const (
	// Modifiers
	LEFT_CONTROL = "left_control"
	RIGHT_SHIFT  = "right_shift"
	ANY          = "any"

	// Special keys
	UP_ARROW            = "up_arrow"
	DOWN_ARROW          = "down_arrow"
	LEFT_ARROW          = "left_arrow"
	RIGHT_ARROW         = "right_arrow"
	ESCAPE              = "escape"
	DELETE_OR_BACKSPACE = "delete_or_backspace"
	CAPS_LOCK           = "caps_lock"
	OPEN_BRACKET        = "open_bracket"
	SEMICOLON           = "semicolon"

	// Mouse buttons
	BUTTON1 = "button1"
	BUTTON2 = "button2"
	BUTTON3 = "button3"

	// Variable names
	MOUSE_KEYS_FULL        = "mouse_keys_full"
	MOUSE_KEYS_FULL_SCROLL = "mouse_keys_full_scroll"

	// Types
	VARIABLE_IF = "variable_if"
	BASIC       = "basic"
	ANSI        = "ansi"
)

type KarabinerConfig struct {
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	ComplexModifications ComplexModifications `json:"complex_modifications"`
	Name                 string               `json:"name"`
	Selected             bool                 `json:"selected"`
	SimpleModifications  []SimpleModification `json:"simple_modifications"`
	VirtualHidKeyboard   VirtualHidKeyboard   `json:"virtual_hid_keyboard"`
}

type ComplexModifications struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Description  string        `json:"description"`
	Manipulators []Manipulator `json:"manipulators"`
}

type Manipulator struct {
	Conditions   []Condition    `json:"conditions,omitempty"`
	From         From           `json:"from"`
	To           []To           `json:"to,omitempty"`
	ToAfterKeyUp []ToAfterKeyUp `json:"to_after_key_up,omitempty"`
	Type         string         `json:"type"`
}

type Condition struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type From struct {
	KeyCode   string     `json:"key_code"`
	Modifiers *Modifiers `json:"modifiers,omitempty"`
}

type Modifiers struct {
	Mandatory []string `json:"mandatory,omitempty"`
	Optional  []string `json:"optional,omitempty"`
}

type To struct {
	KeyCode        string       `json:"key_code,omitempty"`
	MouseKey       *MouseKey    `json:"mouse_key,omitempty"`
	PointingButton string       `json:"pointing_button,omitempty"`
	SetVariable    *SetVariable `json:"set_variable,omitempty"`
}

type ToAfterKeyUp struct {
	SetVariable *SetVariable `json:"set_variable,omitempty"`
}

type SetVariable struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type MouseKey struct {
	X               int     `json:"x,omitempty"`
	Y               int     `json:"y,omitempty"`
	VerticalWheel   int     `json:"vertical_wheel,omitempty"`
	HorizontalWheel int     `json:"horizontal_wheel,omitempty"`
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
}

type SimpleModification struct {
	From From `json:"from"`
	To   []To `json:"to"`
}

type VirtualHidKeyboard struct {
	KeyboardType string `json:"keyboard_type_v2"`
}

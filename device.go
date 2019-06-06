package msapi

// DeviceKind for device kind
type DeviceKind uint8
const (
	// DeviceKindVideoInput for videoinput
	DeviceKindVideoInput DeviceKind = iota
	// DeviceKindAudioInput for audioinput
	DeviceKindAudioInput
	// DeviceKindAudioOutput for audiooutput
	DeviceKindAudioOutput
)

// DeviceInfo for device info
type DeviceInfo struct {
	DeviceID string
	GroupID string
	Kind DeviceKind
	Label string
}
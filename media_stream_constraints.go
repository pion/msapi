package msapi

// MediaStreamConstraints for media stream constraints
type MediaStreamConstraints map[string]interface{}

const (
	// PropertyAudio for audio
	PropertyAudio = "audio"
	// PropertyVideo for video
	PropertyVideo = "video"
)
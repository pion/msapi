package msapi

// MediaTrackConstraints for media track constraints
type MediaTrackConstraints map[string]interface{}

const (
	// Properties of all media tracks

	// PropertyDeviceID for deviceId
	PropertyDeviceID = "deviceId"
	// PropertyGroupID for groupId
	PropertyGroupID = "groupId"

	// Properties of audio tracks

	// PropertyAutoGainControl for autoGainControl
	PropertyAutoGainControl = "autoGainControl"
	// PropertyChannelCount for channelCount
	PropertyChannelCount = "channelCount"
	// PropertyEchoCancellation for echoCancellation
	PropertyEchoCancellation = "echoCancellation"
	// PropertyLatency for latency
	PropertyLatency = "latency"
	// PropertyNoiseSuppression for noiseSuppression
	PropertyNoiseSuppression = "noiseSuppression"
	// PropertySampleRate for sampleRate
	PropertySampleRate = "sampleRate"
	// PropertySampleSize for sampleSize
	PropertySampleSize = "sampleSize"
	// PropertyVolume for volume
	PropertyVolume = "volume"

	
	// Properties of video tracks

	// PropertyAspectRatio for aspectRatio
	PropertyAspectRatio = "aspectRatio"
	// PropertyFacingMode for facingMode
	PropertyFacingMode = "facingMode"
	// PropertyFrameRate for frameRate
	PropertyFrameRate = "frameRate"
	// PropertyHeight for height
	PropertyHeight = "height"
	// PropertyWidth for width
	PropertyWidth = "width"

	// Properties of shared screen tracks

	// PropertyCursor for cursor
	PropertyCursor = "cursor"
	// PropertyDisplaySurface for displaySurface
	PropertyDisplaySurface = "displaySurface"
	// PropertyLogicalSurface for logicalSurface
	PropertyLogicalSurface = "logicalSurface"
)
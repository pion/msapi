package msapi

// DeviceManager provides access to connected media input devices like cameras and microphones, as well as screen sharing
type DeviceManager interface {
	// GetUserMedia produces a MediaStream with tracks containing the requested types of media
	GetUserMedia(MediaStreamConstraints) MediaStream
	EnumerateDevices() []DeviceInfo
	GetDisplayMedia(MediaStreamConstraints) MediaStream
}

// MediaStream is used to group several MediaStreamTrack objects into one unit that can be recorded or rendered in a media element
type MediaStream interface {
	Active() bool
	GetTracks() []MediaStreamTrack
	GetAudioTracks() []MediaStreamTrack
	GetVideoTracks() []MediaStreamTrack
	AddTrack(MediaStreamTrack)
	RemoveTrack(MediaStreamTrack)
	Clone() MediaStream
}

// MediaStreamTrack object represents media of a single type that originates from one media source in the User Agent, e.g. video produced by a web camera
type MediaStreamTrack interface {
	Stop()
	GetConstraints()
	Clone() MediaStreamTrack
}

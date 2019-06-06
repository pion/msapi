package msapi

type mediaStream struct {

}

func (impl *mediaStream) Active() bool {
	return false
}

func (impl *mediaStream) GetTracks() []MediaStreamTrack {
	return nil
}

func (impl *mediaStream) GetAudioTracks() []MediaStreamTrack {
	return nil
}

func (impl *mediaStream) GetVideoTracks() []MediaStreamTrack {
	return nil
}

func (impl *mediaStream) AddTrack(MediaStreamTrack) {
	return
}

func (impl *mediaStream) RemoveTrack(MediaStreamTrack) {

}

func (impl *mediaStream) Clone() MediaStream {
	return nil
}

// NewMediaStream creates a MediaStream
func NewMediaStream() MediaStream {
	return &mediaStream{}
}
package msapi

type mediaStreamTrack struct {

}



func (impl *mediaStreamTrack) Stop() {

}

func (impl *mediaStreamTrack) GetConstraints() {
	
}

func (impl *mediaStreamTrack) Clone() MediaStreamTrack {
	return nil
}


// NewMediaStreamTrack creates a MediaStreamTrack
func NewMediaStreamTrack() MediaStreamTrack {
	return &mediaStreamTrack{}
}
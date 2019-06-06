package msapi

type mediaRecorder struct {
	ms MediaStream
	opts MediaRecorderOptions
}

// BlobEvent for blob event
type BlobEvent struct {
	Blob []byte
	Timecode int64
}

// MediaRecorderOptions for MediaRecorder options
type MediaRecorderOptions struct {
	MimeType string
	VideoBitsPerSecond uint
	AudioBitsPerSecond uint
	BitsPerSecond uint
	OnStart func()
	OnStop func()
	OnDataAvailable func(BlobEvent)
	OnPause func()
	OnResume func()
	OnError func()
}

func (impl *mediaRecorder) Start(timeSlice uint) {
	return
}

func (impl *mediaRecorder) Stop() {
	return
}

func (impl *mediaRecorder) Pause() {
	return
}

func (impl *mediaRecorder) Resume() {
	return
}

func (impl *mediaRecorder) RequestData() {
	return
}

// NewMediaRecorder creates a MediaRecorder
func NewMediaRecorder(ms MediaStream, opts MediaRecorderOptions) MediaRecorder {
	return &mediaRecorder{ms:ms,opts:opts}
}
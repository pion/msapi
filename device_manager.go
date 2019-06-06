package msapi

type deviceManager struct {

}


func (impl *deviceManager) GetUserMedia(MediaStreamConstraints) MediaStream {
	return nil
}

func (impl *deviceManager) EnumerateDevices() []DeviceInfo {
	return nil
}

func (impl *deviceManager) GetDisplayMedia(MediaStreamConstraints) MediaStream {
	return nil
}

// NewDeviceManager creates a DeviceManager
func NewDeviceManager() DeviceManager {
	return &deviceManager{}
}
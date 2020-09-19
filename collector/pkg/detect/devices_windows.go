package detect

func (d *Detect) devicePrefix() string {
	return ""
}

func (d *Detect) Start() ([]models.Device, error) {
	// call the base/common functionality to get a list of devices
	detectedDevices, err := d.smartctlScan()
	if err != nil {
		return nil, err
	}

	//inflate device info for detected devices.
	for ndx, _ := range detectedDevices {
		d.smartCtlInfo(&detectedDevices[ndx]) //ignore errors.
	}

	return detectedDevices, nil
}

//WWN values NVMe and SCSI
func (d *Detect) wwnFallback(detectedDevice *models.Device) {

	//fallback to serial number
	if len(detectedDevice.WWN) == 0 {
		detectedDevice.WWN = detectedDevice.SerialNumber
	}
}

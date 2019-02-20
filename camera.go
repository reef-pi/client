package client

func (c *client) TakePhoto() error {
	return c.post("/api/camera/shoot", nil)
}

func (c *client) CameraConfig() (CameraConfig, error) {
	var conf CameraConfig
	return conf, c.get("/api/camera/config", &conf)
}

func (c *client) UpdateCameraConfig(conf CameraConfig) error {
	return c.post("/api/camera/config", &conf)
}

func (c *client) ListPhotos() ([]ImageItem, error) {
	var images []ImageItem
	return images, c.get("/api/camera/list", &images)
}

func (c *client) LatestPhoto() (map[string]string, error) {
	var d map[string]string
	return d, c.get("/api/camera/latest", &d)
}

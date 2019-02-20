package client

func (c *Client) TakePhoto() error {
	return c.post("/api/camera/shoot", nil)
}

func (c *Client) CameraConfig() (CameraConfig, error) {
	var conf CameraConfig
	return conf, c.get("/api/camera/config", &conf)
}

func (c *Client) UpdateCameraConfig(conf CameraConfig) error {
	return c.post("/api/camera/config", &conf)
}

func (c *Client) ListPhotos() ([]ImageItem, error) {
	var images []ImageItem
	return images, c.get("/api/camera/list", &images)
}

func (c *Client) LatestPhoto() (map[string]string, error) {
	var d map[string]string
	return d, c.get("/api/camera/latest", &d)
}

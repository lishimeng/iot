package module

type GeoData struct {
	Lat float64
	Lng float64
}

// DeviceData 应用层数据结构
type DeviceData struct {
	Application string
	Device      string
	Properties  map[string]interface{}
	Geo         GeoData
}

// NetworkData 网络层数据结构,由各个network bridge收集,提交到上层解析
type NetworkData struct {
	ApplicationId int         `json:"application_id,omitempty"` // 平台应用ID
	DeviceId      int         `json:"device_id,omitempty"`      // 平台设备ID
	Payload       interface{} `json:"payload,omitempty"`        // 数据载体
}

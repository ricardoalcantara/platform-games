package catalog

type VpsCatalogView struct {
	Id               uint    `json:"id"`
	VCpu             int     `json:"vcpu"`
	Memory           int     `json:"memory"`
	DiskSize         int     `json:"disk_size"`
	Price            float64 `json:"price"`
	HourPrice        float64 `json:"hour_price"`
	VpsCatalogTypeId uint8   `json:"vps_catalog_type_id"`
	VpsCatalogType   string  `json:"vps_catalog_type"`
	Ipv4             bool    `json:"ipv4"`
	Ipv6             bool    `json:"ipv6"`
}

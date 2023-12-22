package catalog

import "github.com/ricardoalcantara/platform-games/internal/models"

type CatalogView struct {
	Id          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Tags        []string        `json:"tags"`
	Variations  []VariationView `json:"variations"`
	Images      []ImageView     `json:"images"`
}

type VariationView struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	HourPrice  float64 `json:"hour_price"`
	Active     bool    `json:"active"`
	Memory     int     `json:"memory"`
	DiskSize   int     `json:"diskSize"`
	MaxPlayers int     `json:"max_players"`

	Details []DetailsView `json:"details"`
}

type DetailsView struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ImageView struct {
	Url string `json:"url"`
}

func NewCatalogView(c *models.Catalog, hours int) CatalogView {
	var tags []string
	for _, tag := range c.CatalogTagType {
		tags = append(tags, tag.Name)
	}

	vpsCatalogView := CatalogView{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Tags:        tags,
	}

	for _, v := range c.CatalogVariation {
		hourPrice := v.Price / float64(hours)
		variationView := VariationView{
			Name:       v.Name,
			Price:      v.Price,
			HourPrice:  hourPrice,
			Active:     v.Active,
			Memory:     v.Memory,
			DiskSize:   v.DiskSize,
			MaxPlayers: v.MaxPlayers,
		}

		for _, d := range v.VariationDetail {
			detailsView := DetailsView{
				Name:        d.Name,
				Description: d.Description,
			}
			variationView.Details = append(variationView.Details, detailsView)
		}

		vpsCatalogView.Variations = append(vpsCatalogView.Variations, variationView)
	}

	for _, img := range c.CatalogImage {
		imageView := ImageView{
			Url: img.Url,
		}
		vpsCatalogView.Images = append(vpsCatalogView.Images, imageView)
	}

	return vpsCatalogView
}

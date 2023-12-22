package catalog

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/platform-games/internal/domain"
	"github.com/ricardoalcantara/platform-games/internal/models"
	"github.com/ricardoalcantara/platform-games/internal/utils"
	"github.com/samber/lo"
)

func list(c *gin.Context) {
	vpsCatalogs, err := models.ListCatalog()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	now := time.Now()
	days := utils.GetTotalHoursInAMonth(now.Year(), int(now.Month()))
	result := lo.Map(vpsCatalogs, func(c models.Catalog, _index int) CatalogView {
		return NewCatalogView(&c, days)
	})

	c.JSON(http.StatusOK, domain.ListView[CatalogView]{List: result, Page: 1})
}

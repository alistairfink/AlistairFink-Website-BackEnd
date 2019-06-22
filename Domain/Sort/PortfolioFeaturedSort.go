package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortPortfolioFeaturedBySortOrder(featured *[]DataModels.PortfolioFeaturedDataModel) {
	sortOrder := func(feat1, feat2 DataModels.PortfolioFeaturedDataModel) bool {
		return feat1.SortOrder < feat2.SortOrder
	}

	sortPortfolioFeatured(sortOrder).Sort(featured)
}

type sortPortfolioFeatured func(feat1, feat2 DataModels.PortfolioFeaturedDataModel) bool

func (this sortPortfolioFeatured) Sort(featured *[]DataModels.PortfolioFeaturedDataModel) {
	portfolioFeaturedSorter := &portfolioFeaturedSorter {
		featured: featured,
		by: this,
	}

	sort.Sort(portfolioFeaturedSorter)
}

type portfolioFeaturedSorter struct {
	featured *[]DataModels.PortfolioFeaturedDataModel
	by func(feat1, feat2 DataModels.PortfolioFeaturedDataModel) bool
}

func (this *portfolioFeaturedSorter) Len() int {
	return len(*this.featured)
}

func (this *portfolioFeaturedSorter) Swap(i, j int) {
	(*this.featured)[i], (*this.featured)[j] = (*this.featured)[j], (*this.featured)[i]
}

func (this *portfolioFeaturedSorter) Less(i, j int) bool {
	return this.by((*this.featured)[i], (*this.featured)[j])
}
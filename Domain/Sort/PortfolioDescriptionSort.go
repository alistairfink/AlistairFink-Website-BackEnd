package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortPortfolioDescriptionBySortOrder(description *[]DataModels.PortfolioDescriptionDataModel) {
	sortOrder := func(desc1, desc2 DataModels.PortfolioDescriptionDataModel) bool {
		return desc1.SortOrder < desc2.SortOrder
	}

	sortPortfolioDescription(sortOrder).Sort(description)
}

type sortPortfolioDescription func(desc1, desc2 DataModels.PortfolioDescriptionDataModel) bool

func (this sortPortfolioDescription) Sort(description *[]DataModels.PortfolioDescriptionDataModel) {
	portfolioDescriptionSorter := &portfolioDescriptionSorter {
		description: description,
		by: this,
	}

	sort.Sort(portfolioDescriptionSorter)
}

type portfolioDescriptionSorter struct {
	description *[]DataModels.PortfolioDescriptionDataModel
	by func(desc1, desc2 DataModels.PortfolioDescriptionDataModel) bool
}

func (this *portfolioDescriptionSorter) Len() int {
	return len(*this.description)
}

func (this *portfolioDescriptionSorter) Swap(i, j int) {
	(*this.description)[i], (*this.description)[j] = (*this.description)[j], (*this.description)[i]
}

func (this *portfolioDescriptionSorter) Less(i, j int) bool {
	return this.by((*this.description)[i], (*this.description)[j])
}
package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

func SortPortfolioByYear(portfolio *[]DomainModels.PortfolioDomainModel) {
	sortOrder := func(pItem1, pItem2 DomainModels.PortfolioDomainModel) bool {
		return pItem1.Year.Year() < pItem2.Year.Year()
	}

	sortPortfolio(sortOrder).Sort(portfolio)
}

type sortPortfolio func(pItem1, pItem2 DomainModels.PortfolioDomainModel) bool

func (this sortPortfolio) Sort(portfolio *[]DomainModels.PortfolioDomainModel) {
	portfolioSorter := &portfolioSorter {
		portfolio: portfolio,
		by: this,
	}

	sort.Sort(portfolioSorter)
}

type portfolioSorter struct {
	portfolio *[]DomainModels.PortfolioDomainModel
	by func(pItem1, pItem2 DomainModels.PortfolioDomainModel) bool
}

func (this *portfolioSorter) Len() int {
	return len(*this.portfolio)
}

func (this *portfolioSorter) Swap(i, j int) {
	(*this.portfolio)[i], (*this.portfolio)[j] = (*this.portfolio)[j], (*this.portfolio)[i]
}

func (this *portfolioSorter) Less(i, j int) bool {
	return this.by((*this.portfolio)[i], (*this.portfolio)[j])
}
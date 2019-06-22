package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortPortfolioImageBySortOrder(images *[]DataModels.PortfolioImageDataModel) {
	sortOrder := func(img1, img2 DataModels.PortfolioImageDataModel) bool {
		return img1.SortOrder < img2.SortOrder
	}

	sortPortfolioImage(sortOrder).Sort(images)
}

type sortPortfolioImage func(img1, img2 DataModels.PortfolioImageDataModel) bool

func (this sortPortfolioImage) Sort(images *[]DataModels.PortfolioImageDataModel) {
	portfolioImageSorter := &portfolioImageSorter {
		images: images,
		by: this,
	}

	sort.Sort(portfolioImageSorter)
}

type portfolioImageSorter struct {
	images *[]DataModels.PortfolioImageDataModel
	by func(img1, img2 DataModels.PortfolioImageDataModel) bool
}

func (this *portfolioImageSorter) Len() int {
	return len(*this.images)
}

func (this *portfolioImageSorter) Swap(i, j int) {
	(*this.images)[i], (*this.images)[j] = (*this.images)[j], (*this.images)[i]
}

func (this *portfolioImageSorter) Less(i, j int) bool {
	return this.by((*this.images)[i], (*this.images)[j])
}
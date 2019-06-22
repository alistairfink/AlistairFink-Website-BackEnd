package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortAboutDescriptionBySortOrder(description *[]DataModels.AboutDescriptionDataModel) {
	sortOrder := func(desc1, desc2 DataModels.AboutDescriptionDataModel) bool {
		return desc1.SortOrder < desc2.SortOrder
	}

	sortAboutDescription(sortOrder).Sort(description)
}

type sortAboutDescription func(desc1, desc2 DataModels.AboutDescriptionDataModel) bool

func (this sortAboutDescription) Sort(description *[]DataModels.AboutDescriptionDataModel) {
	aboutDescriptionSorter := &aboutDescriptionSorter {
		description: description,
		by: this,
	}

	sort.Sort(aboutDescriptionSorter)
}

type aboutDescriptionSorter struct {
	description *[]DataModels.AboutDescriptionDataModel
	by func(desc1, desc2 DataModels.AboutDescriptionDataModel) bool
}

func (this *aboutDescriptionSorter) Len() int {
	return len(*this.description)
}

func (this *aboutDescriptionSorter) Swap(i, j int) {
	(*this.description)[i], (*this.description)[j] = (*this.description)[j], (*this.description)[i]
}

func (this *aboutDescriptionSorter) Less(i, j int) bool {
	return this.by((*this.description)[i], (*this.description)[j])
}
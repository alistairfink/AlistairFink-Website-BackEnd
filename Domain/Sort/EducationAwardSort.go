package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortEducationAwardBySortOrder(awards *[]DataModels.EducationAwardDataModel) {
	sortOrder := func(award1, award2 DataModels.EducationAwardDataModel) bool {
		return award1.SortOrder < award2.SortOrder
	}

	sortEducationAward(sortOrder).Sort(awards)
}

type sortEducationAward func(award1, award2 DataModels.EducationAwardDataModel) bool

func (this sortEducationAward) Sort(awards *[]DataModels.EducationAwardDataModel) {
	educationAwardSorter := &educationAwardSorter {
		awards: awards,
		by: this,
	}

	sort.Sort(educationAwardSorter)
}

type educationAwardSorter struct {
	awards *[]DataModels.EducationAwardDataModel
	by func(award1, award2 DataModels.EducationAwardDataModel) bool
}

func (this *educationAwardSorter) Len() int {
	return len(*this.awards)
}

func (this *educationAwardSorter) Swap(i, j int) {
	(*this.awards)[i], (*this.awards)[j] = (*this.awards)[j], (*this.awards)[i]
}

func (this *educationAwardSorter) Less(i, j int) bool {
	return this.by((*this.awards)[i], (*this.awards)[j])
}
package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

func SortEducationByStartDate(education *[]DomainModels.EducationDomainModel) {
	sortOrder := func(ed1, ed2 DomainModels.EducationDomainModel) bool {
		return ed1.StartDate.Unix() < ed2.StartDate.Unix()
	}

	sortEducation(sortOrder).Sort(education)
}

type sortEducation func(ed1, ed2 DomainModels.EducationDomainModel) bool

func (this sortEducation) Sort(education *[]DomainModels.EducationDomainModel) {
	educationSorter := &educationSorter {
		education: education,
		by: this,
	}

	sort.Sort(educationSorter)
}

type educationSorter struct {
	education *[]DomainModels.EducationDomainModel
	by func(ed1, ed2 DomainModels.EducationDomainModel) bool
}

func (this *educationSorter) Len() int {
	return len(*this.education)
}

func (this *educationSorter) Swap(i, j int) {
	(*this.education)[i], (*this.education)[j] = (*this.education)[j], (*this.education)[i]
}

func (this *educationSorter) Less(i, j int) bool {
	return this.by((*this.education)[i], (*this.education)[j])
}
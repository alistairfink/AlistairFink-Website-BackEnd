package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortEducationExtracurricularBySortOrder(extracurriculars *[]DataModels.EducationExtracurricularDataModel) {
	sortOrder := func(extracurricular1, extracurricular2 DataModels.EducationExtracurricularDataModel) bool {
		return extracurricular1.SortOrder < extracurricular2.SortOrder
	}

	sortEducationExtracurricular(sortOrder).Sort(extracurriculars)
}

type sortEducationExtracurricular func(extracurricular1, extracurricular2 DataModels.EducationExtracurricularDataModel) bool

func (this sortEducationExtracurricular) Sort(extracurriculars *[]DataModels.EducationExtracurricularDataModel) {
	educationExtracurricularSorter := &educationExtracurricularSorter {
		extracurriculars: extracurriculars,
		by: this,
	}

	sort.Sort(educationExtracurricularSorter)
}

type educationExtracurricularSorter struct {
	extracurriculars *[]DataModels.EducationExtracurricularDataModel
	by func(extracurricular1, extracurricular2 DataModels.EducationExtracurricularDataModel) bool
}

func (this *educationExtracurricularSorter) Len() int {
	return len(*this.extracurriculars)
}

func (this *educationExtracurricularSorter) Swap(i, j int) {
	(*this.extracurriculars)[i], (*this.extracurriculars)[j] = (*this.extracurriculars)[j], (*this.extracurriculars)[i]
}

func (this *educationExtracurricularSorter) Less(i, j int) bool {
	return this.by((*this.extracurriculars)[i], (*this.extracurriculars)[j])
}
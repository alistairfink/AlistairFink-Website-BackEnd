package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortEducationScholarshipBySortOrder(scholarships *[]DataModels.EducationScholarshipDataModel) {
	sortOrder := func(scholarship1, scholarship2 DataModels.EducationScholarshipDataModel) bool {
		return scholarship1.SortOrder < scholarship2.SortOrder
	}

	sortEducationScholarship(sortOrder).Sort(scholarships)
}

type sortEducationScholarship func(scholarship1, scholarship2 DataModels.EducationScholarshipDataModel) bool

func (this sortEducationScholarship) Sort(scholarships *[]DataModels.EducationScholarshipDataModel) {
	educationScholarshipSorter := &educationScholarshipSorter {
		scholarships: scholarships,
		by: this,
	}

	sort.Sort(educationScholarshipSorter)
}

type educationScholarshipSorter struct {
	scholarships *[]DataModels.EducationScholarshipDataModel
	by func(scholarship1, scholarship2 DataModels.EducationScholarshipDataModel) bool
}

func (this *educationScholarshipSorter) Len() int {
	return len(*this.scholarships)
}

func (this *educationScholarshipSorter) Swap(i, j int) {
	(*this.scholarships)[i], (*this.scholarships)[j] = (*this.scholarships)[j], (*this.scholarships)[i]
}

func (this *educationScholarshipSorter) Less(i, j int) bool {
	return this.by((*this.scholarships)[i], (*this.scholarships)[j])
}
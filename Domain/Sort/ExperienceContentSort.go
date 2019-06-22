package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortExperienceContentBySortOrder(contents *[]DataModels.ExperienceContentDataModel) {
	sortOrder := func(content1, content2 DataModels.ExperienceContentDataModel) bool {
		return content1.SortOrder < content2.SortOrder
	}

	sortExperienceContent(sortOrder).Sort(contents)
}

type sortExperienceContent func(content1, content2 DataModels.ExperienceContentDataModel) bool

func (this sortExperienceContent) Sort(contents *[]DataModels.ExperienceContentDataModel) {
	experienceContentSorter := &experienceContentSorter {
		contents: contents,
		by: this,
	}

	sort.Sort(experienceContentSorter)
}

type experienceContentSorter struct {
	contents *[]DataModels.ExperienceContentDataModel
	by func(content1, content2 DataModels.ExperienceContentDataModel) bool
}

func (this *experienceContentSorter) Len() int {
	return len(*this.contents)
}

func (this *experienceContentSorter) Swap(i, j int) {
	(*this.contents)[i], (*this.contents)[j] = (*this.contents)[j], (*this.contents)[i]
}

func (this *experienceContentSorter) Less(i, j int) bool {
	return this.by((*this.contents)[i], (*this.contents)[j])
}
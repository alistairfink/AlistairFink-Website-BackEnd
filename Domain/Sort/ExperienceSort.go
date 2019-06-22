package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

func SortExperienceByStartDate(experience *[]DomainModels.ExperienceDomainModel) {
	sortOrder := func(exp1, exp2 DomainModels.ExperienceDomainModel) bool {
		return exp1.StartDate.Unix() < exp2.StartDate.Unix()
	}

	sortExperience(sortOrder).Sort(experience)
}

type sortExperience func(exp1, exp2 DomainModels.ExperienceDomainModel) bool

func (this sortExperience) Sort(experience *[]DomainModels.ExperienceDomainModel) {
	experienceSorter := &experienceSorter {
		experience: experience,
		by: this,
	}

	sort.Sort(experienceSorter)
}

type experienceSorter struct {
	experience *[]DomainModels.ExperienceDomainModel
	by func(exp1, exp2 DomainModels.ExperienceDomainModel) bool
}

func (this *experienceSorter) Len() int {
	return len(*this.experience)
}

func (this *experienceSorter) Swap(i, j int) {
	(*this.experience)[i], (*this.experience)[j] = (*this.experience)[j], (*this.experience)[i]
}

func (this *experienceSorter) Less(i, j int) bool {
	return this.by((*this.experience)[i], (*this.experience)[j])
}
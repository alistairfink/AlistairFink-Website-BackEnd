package Sort

import (
	"sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

func SortPortfolioVideoBySortOrder(videos *[]DataModels.PortfolioVideoDataModel) {
	sortOrder := func(vid1, vid2 DataModels.PortfolioVideoDataModel) bool {
		return vid1.SortOrder < vid2.SortOrder
	}

	sortPortfolioVideo(sortOrder).Sort(videos)
}

type sortPortfolioVideo func(vid1, vid2 DataModels.PortfolioVideoDataModel) bool

func (this sortPortfolioVideo) Sort(videos *[]DataModels.PortfolioVideoDataModel) {
	portfolioVideoSorter := &portfolioVideoSorter {
		videos: videos,
		by: this,
	}

	sort.Sort(portfolioVideoSorter)
}

type portfolioVideoSorter struct {
	videos *[]DataModels.PortfolioVideoDataModel
	by func(vid1, vid2 DataModels.PortfolioVideoDataModel) bool
}

func (this *portfolioVideoSorter) Len() int {
	return len(*this.videos)
}

func (this *portfolioVideoSorter) Swap(i, j int) {
	(*this.videos)[i], (*this.videos)[j] = (*this.videos)[j], (*this.videos)[i]
}

func (this *portfolioVideoSorter) Less(i, j int) bool {
	return this.by((*this.videos)[i], (*this.videos)[j])
}
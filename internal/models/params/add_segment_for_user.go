package params

type SegmentParams struct {
	SegmentName []string `json:"segmentName"`
	UserId      int      `json:"userId"`
}

type SegmentsForRandomUsers struct {
	SegmentName []string `json:"segmentName"`
}

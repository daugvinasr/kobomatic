package gen

import (
	"time"

	"github.com/daugvinasr/kobomatic/constants"
	"github.com/daugvinasr/kobomatic/database/kobomatic"
)

type LocationComplete struct {
	Value  string `json:"Value"`
	Type   string `json:"Type"`
	Source string `json:"Source"`
}

type StatusInfoComplete struct {
	LastModified           string `json:"LastModified"`
	Status                 string `json:"Status"`
	TimesStartedReading    int    `json:"TimesStartedReading"`
	LastTimeStartedReading string `json:"LastTimeStartedReading"`
}

type StatisticsComplete struct {
	LastModified         string `json:"LastModified"`
	SpentReadingMinutes  int    `json:"SpentReadingMinutes"`
	RemainingTimeMinutes int    `json:"RemainingTimeMinutes"`
}

type CurrentBookmarkComplete struct {
	LastModified                 string           `json:"LastModified"`
	ProgressPercent              int              `json:"ProgressPercent"`
	ContentSourceProgressPercent int              `json:"ContentSourceProgressPercent"`
	Location                     LocationComplete `json:"Location"`
}

type ReadingStateComplete struct {
	EntitlementID     string                  `json:"EntitlementId"`
	Created           string                  `json:"Created"`
	LastModified      string                  `json:"LastModified"`
	StatusInfo        StatusInfoComplete      `json:"StatusInfo"`
	Statistics        StatisticsComplete      `json:"Statistics"`
	CurrentBookmark   CurrentBookmarkComplete `json:"CurrentBookmark"`
	PriorityTimestamp string                  `json:"PriorityTimestamp"`
}

func CompleteReadingState(readingState kobomatic.ReadingState) (ReadingStateComplete, error) {
	parsedLastModified, err := time.Parse(constants.ISO8601, readingState.LastModified)
	if err != nil {
		return ReadingStateComplete{}, err
	}

	lastModified := parsedLastModified.Format(constants.KoboTime)

	statusInfo := StatusInfoComplete{
		LastModified:           lastModified,
		Status:                 readingState.Status,
		TimesStartedReading:    0,
		LastTimeStartedReading: lastModified,
	}

	statistics := StatisticsComplete{
		LastModified:         lastModified,
		SpentReadingMinutes:  readingState.SpentReadingMinutes,
		RemainingTimeMinutes: readingState.RemainingTimeMinutes,
	}

	location := LocationComplete{
		Value:  readingState.Value,
		Type:   readingState.Type,
		Source: readingState.Source,
	}

	currentBookmark := CurrentBookmarkComplete{
		LastModified:                 lastModified,
		ProgressPercent:              readingState.ProgressPercent,
		ContentSourceProgressPercent: readingState.ContentSourceProgressPercent,
		Location:                     location,
	}

	return ReadingStateComplete{
		EntitlementID:     readingState.EntitlementID,
		Created:           lastModified,
		LastModified:      lastModified,
		StatusInfo:        statusInfo,
		Statistics:        statistics,
		CurrentBookmark:   currentBookmark,
		PriorityTimestamp: lastModified,
	}, nil
}

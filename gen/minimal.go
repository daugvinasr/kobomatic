package gen

import "github.com/daugvinasr/kobomatic/constants"

type StatusInfoMinimal struct {
	LastModified        string `json:"LastModified"`
	Status              string `json:"Status"`
	TimesStartedReading int    `json:"TimesStartedReading"`
}

type StatisticsMinimal struct {
	LastModified string `json:"LastModified"`
}

type CurrentBookmarkMinimal struct {
	LastModified string `json:"LastModified"`
}

type ReadingStateMinimal struct {
	EntitlementID     string                 `json:"EntitlementId"`
	Created           string                 `json:"Created"`
	LastModified      string                 `json:"LastModified"`
	StatusInfo        StatusInfoMinimal      `json:"StatusInfo"`
	Statistics        StatisticsMinimal      `json:"Statistics"`
	CurrentBookmark   CurrentBookmarkMinimal `json:"CurrentBookmark"`
	PriorityTimestamp string                 `json:"PriorityTimestamp"`
}

func MinimalReadingState(entitlementID string) ReadingStateMinimal {
	statusInfo := StatusInfoMinimal{
		LastModified:        constants.KoboOldTimestamp,
		Status:              "ReadyToRead",
		TimesStartedReading: 0,
	}

	statistics := StatisticsMinimal{
		LastModified: constants.KoboOldTimestamp,
	}

	currentBookmark := CurrentBookmarkMinimal{
		LastModified: constants.KoboOldTimestamp,
	}

	return ReadingStateMinimal{
		EntitlementID:     entitlementID,
		Created:           constants.KoboOldTimestamp,
		LastModified:      constants.KoboOldTimestamp,
		StatusInfo:        statusInfo,
		Statistics:        statistics,
		CurrentBookmark:   currentBookmark,
		PriorityTimestamp: constants.KoboOldTimestamp,
	}
}

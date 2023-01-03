package goIMDB

import "time"

/*
Response data models - this is the json of the various repsonses I picked up.

*/
// Title represnet a movie or tvseries title.  The data is extracted as best as possible from the json
// returned by the api
type Title1 struct {
	imdb_id       string
	title         string
	Type          string
	Certification string
	Year          int
	Genres        *[]string
	Writers       *[]People
	Creators      *[]People
	//Credits       *[]People  /removed since hard to get data from json need reflection
	Directors    *[]People
	Stars        *[]People
	Image        Image
	Episodes     int
	Rating_count int
	Releases     *TitleReleases
	Season       int
	Episode      int
	Rating       float64
	Plot_outline string
	Release_date time.Time
	Runtime      int
	// Non exported fields contain the raw data and
	tExtra    *titleExtra
	tTopCrew  *titleTopCrew
	tEpisodes *TitleEpisodes
	tReleases *TitleReleases
	// below fields that are used internally to control errors and actions
	// the client to access the web
	client *IMDBPie
	// title is TvSeries
	isTv bool
	// title recieved ok from API
	hasTitle bool
	// title is a single episode
	isEpisode bool
}

type Image struct {
	Url    string
	Width  int
	Height int
}

type People struct {
	Name       string
	Job        string
	Category   string
	Imdb_id    string
	Characters []string
}

type TitleRelease struct {
	Date   time.Time
	region string
}

type TitleReleases []TitleRelease

//type TitleName struct {
//	Name       string
//	Category   string
//	Imdb_id    string
//	Job        string
//	Characters []People
//}

type PersonName struct {
	Name          string
	Imdb_id       string
	Image         Image
	Birth_place   string
	Gender        string
	Bios          string
	Date_of_birth time.Time
	Filmography   []string
}

type Episode struct {
	Imdb_id string
	Title   string
	Season  int
	Episode int
	Year    int
}

type TitleEpisodes struct {
	Imdb_id  string
	Count    int
	Episodes []Episode
}

// below translated from json

type titleEpisodesDetailed struct {
	Season      int `json:"season"`
	SeriesTitle struct {
		Id    string `json:"id"`
		Image struct {
			Height int    `json:"height"`
			Id     string `json:"id"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image"`
		RunningTimeInMinutes int    `json:"runningTimeInMinutes"`
		NextEpisode          string `json:"nextEpisode"`
		NumberOfEpisodes     int    `json:"numberOfEpisodes"`
		SeriesEndYear        int    `json:"seriesEndYear"`
		SeriesStartYear      int    `json:"seriesStartYear"`
		Title                string `json:"title"`
		TitleType            string `json:"titleType"`
		Year                 int    `json:"year"`
	} `json:"seriesTitle"`
	AllSeasons    []int       `json:"allSeasons"`
	TotalEpisodes int         `json:"totalEpisodes"`
	Start         float64     `json:"start"`
	End           float64     `json:"end"`
	Region        interface{} `json:"region"`
	Episodes      []struct {
		EpisodeNumber int         `json:"episodeNumber"`
		Id            string      `json:"id"`
		Title         string      `json:"title"`
		Year          int         `json:"year"`
		Rating        float64     `json:"rating"`
		UserRating    interface{} `json:"userRating"`
		RatingCount   int         `json:"ratingCount"`
		ReleaseDate   struct {
			Date  interface{} `json:"date"`
			First struct {
				Date    string   `json:"date"`
				Regions []string `json:"regions"`
			} `json:"first"`
		} `json:"releaseDate"`
	} `json:"episodes"`
}

const aux_url = "/title/%s/auxiliary"

// special extra title information
type titleExtra struct {
	Certificate struct {
		Certificate string `json:"certificate"`
		Country     string `json:"country"`
	} `json:"certificate"`
	FilmingLocations []struct {
		Extras           []string `json:"extras,omitempty"`
		Id               string   `json:"id"`
		InterestingVotes struct {
			Up   int `json:"up"`
			Down int `json:"down,omitempty"`
		} `json:"interestingVotes,omitempty"`
		Location   string   `json:"location"`
		Attributes []string `json:"attributes,omitempty"`
	} `json:"filmingLocations"`
	MetacriticInfo struct {
		Type            string `json:"@type"`
		ReviewCount     int    `json:"reviewCount"`
		UserRatingCount int    `json:"userRatingCount"`
	} `json:"metacriticInfo"`
	Plot struct {
		Outline struct {
			Id   string `json:"id"`
			Text string `json:"text"`
		} `json:"outline"`
		TotalSummaries int `json:"totalSummaries"`
	} `json:"plot"`
	Principals []struct {
		Id             string   `json:"id"`
		LegacyNameText string   `json:"legacyNameText"`
		Name           string   `json:"name"`
		Category       string   `json:"category"`
		Characters     []string `json:"characters"`
		EndYear        int      `json:"endYear"`
		EpisodeCount   int      `json:"episodeCount"`
		Job            string   `json:"job"`
		Roles          []struct {
			Character string `json:"character"`
		} `json:"roles"`
		StartYear int `json:"startYear"`
		Image     struct {
			Height int    `json:"height"`
			Id     string `json:"id"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image"`
		Disambiguation string `json:"disambiguation,omitempty"`
	} `json:"principals"`
	Rating        float64 `json:"rating"`
	NumberOfVotes int     `json:"numberOfVotes"`
	CanRate       bool    `json:"canRate"`
	TopRank       struct {
		RankType string `json:"rankType"`
		Rank     int    `json:"rank"`
	} `json:"topRank"`
	UserRating            interface{} `json:"userRating"`
	AlternateTitlesSample []string    `json:"alternateTitlesSample"`
	AlternateTitlesCount  int         `json:"alternateTitlesCount"`
	HasAlternateVersions  bool        `json:"hasAlternateVersions"`
	OriginalTitle         string      `json:"originalTitle"`
	RunningTimes          []struct {
		TimeMinutes int `json:"timeMinutes"`
	} `json:"runningTimes"`
	SpokenLanguages   []string `json:"spokenLanguages"`
	Origins           []string `json:"origins"`
	SimilaritiesCount int      `json:"similaritiesCount"`
	ReleaseDetails    struct {
		Date     string `json:"date"`
		Premiere bool   `json:"premiere"`
		Region   string `json:"region"`
		Wide     bool   `json:"wide"`
	} `json:"releaseDetails"`
	Soundtracks []struct {
		Comment      string `json:"comment"`
		Id           string `json:"id"`
		Name         string `json:"name"`
		RelatedNames []struct {
			Akas  []string `json:"akas,omitempty"`
			Id    string   `json:"id"`
			Image struct {
				Height int    `json:"height"`
				Id     string `json:"id"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image,omitempty"`
			LegacyNameText string `json:"legacyNameText"`
			Name           string `json:"name"`
		} `json:"relatedNames,omitempty"`
	} `json:"soundtracks"`
	Genres        []string `json:"genres"`
	ReviewsTeaser struct {
		Author struct {
			DisplayName string `json:"displayName"`
			UserId      string `json:"userId"`
		} `json:"author"`
		AuthorRating     int     `json:"authorRating"`
		HelpfulnessScore float64 `json:"helpfulnessScore"`
		Id               string  `json:"id"`
		InterestingVotes struct {
			Down int `json:"down"`
			Up   int `json:"up"`
		} `json:"interestingVotes"`
		LanguageCode   string `json:"languageCode"`
		ReviewText     string `json:"reviewText"`
		ReviewTitle    string `json:"reviewTitle"`
		Spoiler        bool   `json:"spoiler"`
		SubmissionDate string `json:"submissionDate"`
		TitleId        string `json:"titleId"`
	} `json:"reviewsTeaser"`
	ReviewsCount       int         `json:"reviewsCount"`
	HasContentGuide    bool        `json:"hasContentGuide"`
	HasSynopsis        bool        `json:"hasSynopsis"`
	HasCriticsReviews  bool        `json:"hasCriticsReviews"`
	CriticsReviewers   []string    `json:"criticsReviewers"`
	CrazyCreditsTeaser interface{} `json:"crazyCreditsTeaser"`
	Awards             struct {
		AwardsSummary struct {
			OtherNominationsCount int `json:"otherNominationsCount"`
			OtherWinsCount        int `json:"otherWinsCount"`
		} `json:"awardsSummary"`
		HighlightedCategory interface{} `json:"highlightedCategory"`
	} `json:"awards"`
	Photos struct {
		Images []struct {
			Caption          string    `json:"caption"`
			CreatedOn        time.Time `json:"createdOn"`
			Height           int       `json:"height"`
			Id               string    `json:"id"`
			Url              string    `json:"url"`
			Width            int       `json:"width"`
			RelatedNamesIds  []string  `json:"relatedNamesIds,omitempty"`
			RelatedTitlesIds []string  `json:"relatedTitlesIds"`
			Source           string    `json:"source"`
			Type             string    `json:"type"`
		} `json:"images"`
		TotalImageCount int `json:"totalImageCount"`
	} `json:"photos"`
	HeroImages []struct {
		Caption     string    `json:"caption"`
		CreatedOn   time.Time `json:"createdOn"`
		Height      int       `json:"height"`
		Id          string    `json:"id"`
		Url         string    `json:"url"`
		Width       int       `json:"width"`
		Source      string    `json:"source"`
		Type        string    `json:"type"`
		Attribution string    `json:"attribution,omitempty"`
		Copyright   string    `json:"copyright,omitempty"`
	} `json:"heroImages"`
	SeasonsInfo []struct {
		Season int `json:"season"`
	} `json:"seasonsInfo"`
	ProductionStatus struct {
		Comment string `json:"comment"`
		Date    string `json:"date"`
		Status  string `json:"status"`
	} `json:"productionStatus"`
	Directors []struct {
		Akas           []string `json:"akas,omitempty"`
		Id             string   `json:"id"`
		LegacyNameText string   `json:"legacyNameText"`
		Name           string   `json:"name"`
		Category       string   `json:"category"`
		EndYear        int      `json:"endYear"`
		EpisodeCount   int      `json:"episodeCount"`
		StartYear      int      `json:"startYear"`
		Disambiguation string   `json:"disambiguation,omitempty"`
		Image          struct {
			Height int    `json:"height"`
			Id     string `json:"id"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image,omitempty"`
	} `json:"directors"`
	Writers []struct {
		Akas  []string `json:"akas,omitempty"`
		Id    string   `json:"id"`
		Image struct {
			Height int    `json:"height"`
			Id     string `json:"id"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image,omitempty"`
		LegacyNameText string   `json:"legacyNameText"`
		Name           string   `json:"name"`
		Attr           []string `json:"attr,omitempty"`
		Category       string   `json:"category"`
		EndYear        int      `json:"endYear"`
		EpisodeCount   int      `json:"episodeCount"`
		Job            string   `json:"job,omitempty"`
		StartYear      int      `json:"startYear"`
		Disambiguation string   `json:"disambiguation,omitempty"`
	} `json:"writers"`
	Videos struct {
		TotalVideoCount int `json:"totalVideoCount"`
		MainTrailer     struct {
			ContentType     string `json:"contentType"`
			Description     string `json:"description"`
			DurationSeconds int    `json:"durationSeconds"`
			Encodings       []struct {
				Definition   string `json:"definition"`
				HeightPixels int    `json:"heightPixels"`
				MimeType     string `json:"mimeType"`
				Play         string `json:"play"`
				VideoCodec   string `json:"videoCodec"`
				WidthPixels  int    `json:"widthPixels"`
			} `json:"encodings"`
			Id    string `json:"id"`
			Image struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			Monetization string `json:"monetization"`
			PrimaryTitle struct {
				Id    string `json:"id"`
				Image struct {
					Height int    `json:"height"`
					Id     string `json:"id"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"primaryTitle"`
			VideoTitle string `json:"videoTitle"`
		} `json:"mainTrailer"`
		HeroVideos []struct {
			ContentType     string `json:"contentType"`
			Description     string `json:"description"`
			DurationSeconds int    `json:"durationSeconds"`
			Encodings       []struct {
				Definition   string `json:"definition"`
				HeightPixels int    `json:"heightPixels"`
				MimeType     string `json:"mimeType"`
				Play         string `json:"play"`
				VideoCodec   string `json:"videoCodec"`
				WidthPixels  int    `json:"widthPixels"`
			} `json:"encodings"`
			Id    string `json:"id"`
			Image struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			Monetization string `json:"monetization"`
			PrimaryTitle struct {
				Id    string `json:"id"`
				Image struct {
					Height int    `json:"height"`
					Id     string `json:"id"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"primaryTitle"`
			VideoTitle string `json:"videoTitle"`
		} `json:"heroVideos"`
		OtherVideos []struct {
			ContentType     string `json:"contentType"`
			Description     string `json:"description"`
			DurationSeconds int    `json:"durationSeconds"`
			Encodings       []struct {
				Definition   string `json:"definition"`
				HeightPixels int    `json:"heightPixels"`
				MimeType     string `json:"mimeType"`
				Play         string `json:"play"`
				VideoCodec   string `json:"videoCodec"`
				WidthPixels  int    `json:"widthPixels"`
			} `json:"encodings"`
			Id    string `json:"id"`
			Image struct {
				Height int    `json:"height"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			Monetization string `json:"monetization"`
			ParentTitle  struct {
				Id    string `json:"id"`
				Image struct {
					Height int    `json:"height"`
					Id     string `json:"id"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"parentTitle"`
			PrimaryTitle struct {
				Episode int    `json:"episode"`
				Id      string `json:"id"`
				Image   struct {
					Height int    `json:"height"`
					Id     string `json:"id"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				Season    int    `json:"season"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"primaryTitle"`
			VideoTitle string `json:"videoTitle"`
		} `json:"otherVideos"`
	} `json:"videos"`
	AdWidgets interface{} `json:"adWidgets"`
	Id        string      `json:"id"`
	Image     struct {
		Height int    `json:"height"`
		Id     string `json:"id"`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"image"`
	RunningTimeInMinutes int    `json:"runningTimeInMinutes"`
	NextEpisode          string `json:"nextEpisode"`
	NumberOfEpisodes     int    `json:"numberOfEpisodes"`
	SeriesEndYear        int    `json:"seriesEndYear"`
	SeriesStartYear      int    `json:"seriesStartYear"`
	Title                string `json:"title"`
	TitleType            string `json:"titleType"`
	Year                 int    `json:"year"`
}

type titleTopCrew struct {
	Directors []struct {
		Akas  []string `json:"akas"`
		Id    string   `json:"id"`
		Image struct {
			Height int    `json:"height"`
			Id     string `json:"id"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image"`
		LegacyNameText     string   `json:"legacyNameText"`
		Name               string   `json:"name"`
		Category           string   `json:"category"`
		FreeTextAttributes []string `json:"freeTextAttributes"`
	} `json:"directors"`
	Writers []struct {
		Akas  []string `json:"akas,omitempty"`
		Id    string   `json:"id"`
		Image struct {
			Height int    `json:"height"`
			Id     string `json:"id"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"image"`
		LegacyNameText        string `json:"legacyNameText"`
		Name                  string `json:"name"`
		Category              string `json:"category"`
		Job                   string `json:"job"`
		WriterCategoryBilling int    `json:"writerCategoryBilling"`
		WriterTeamBilling     int    `json:"writerTeamBilling"`
		Disambiguation        string `json:"disambiguation,omitempty"`
	} `json:"writers"`
}

type T struct {
	Meta struct {
		Operation     string  `json:"operation"`
		RequestId     string  `json:"requestId"`
		ServiceTimeMs float64 `json:"serviceTimeMs"`
	} `json:"@meta"`
	Resource struct {
		Type string `json:"@type"`
		Base struct {
			Episode int    `json:"episode"`
			Id      string `json:"id"`
			Image   struct {
				Height int    `json:"height"`
				Id     string `json:"id"`
				Url    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			RunningTimeInMinutes int    `json:"runningTimeInMinutes"`
			Season               int    `json:"season"`
			NextEpisode          string `json:"nextEpisode"`
			ParentTitle          struct {
				Id    string `json:"id"`
				Image struct {
					Height int    `json:"height"`
					Id     string `json:"id"`
					Url    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"parentTitle"`
			PreviousEpisode string `json:"previousEpisode"`
			SeriesEndYear   int    `json:"seriesEndYear"`
			SeriesStartYear int    `json:"seriesStartYear"`
			Title           string `json:"title"`
			TitleType       string `json:"titleType"`
			Year            int    `json:"year"`
		} `json:"base"`
		FilmingLocations []struct {
			Id       string `json:"id"`
			Location string `json:"location"`
		} `json:"filmingLocations"`
		MetacriticScore struct {
			Type            string `json:"@type"`
			ReviewCount     int    `json:"reviewCount"`
			UserRatingCount int    `json:"userRatingCount"`
		} `json:"metacriticScore"`
		Plot struct {
			Outline struct {
				Id   string `json:"id"`
				Text string `json:"text"`
			} `json:"outline"`
			Summaries []struct {
				Author string `json:"author"`
				Id     string `json:"id"`
				Text   string `json:"text"`
			} `json:"summaries"`
			TotalSummaries int `json:"totalSummaries"`
		} `json:"plot"`
		Ratings struct {
			Episode           int     `json:"episode"`
			Id                string  `json:"id"`
			Season            int     `json:"season"`
			Title             string  `json:"title"`
			TitleType         string  `json:"titleType"`
			Year              int     `json:"year"`
			CanRate           bool    `json:"canRate"`
			Rating            float64 `json:"rating"`
			RatingCount       int     `json:"ratingCount"`
			RatingsHistograms struct {
				FemalesAged3044 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Females Aged 30-44"`
				Aged1829 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Aged 18-29"`
				NonUSUsers struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Non-US users"`
				MalesAged45 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Males Aged 45+"`
				Males struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Males"`
				Females struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Females"`
				MalesAged3044 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Males Aged 30-44"`
				Top1000Voters struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Top 1000 voters"`
				MalesAged1829 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Males Aged 18-29"`
				FemalesAged45 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Females Aged 45+"`
				IMDbUsers struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"IMDb Users"`
				IMDbStaff struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"IMDb Staff"`
				FemalesAgedUnder18 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Females Aged under 18"`
				FemalesAged1829 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Females Aged 18-29"`
				Aged45 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Aged 45+"`
				USUsers struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"US users"`
				MalesAgedUnder18 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Males Aged under 18"`
				AgedUnder18 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Aged under 18"`
				Aged3044 struct {
					AggregateRating float64 `json:"aggregateRating"`
					Demographic     string  `json:"demographic"`
					Histogram       struct {
						Field1  int `json:"1"`
						Field2  int `json:"2"`
						Field3  int `json:"3"`
						Field4  int `json:"4"`
						Field5  int `json:"5"`
						Field6  int `json:"6"`
						Field7  int `json:"7"`
						Field8  int `json:"8"`
						Field9  int `json:"9"`
						Field10 int `json:"10"`
					} `json:"histogram"`
					TotalRatings int `json:"totalRatings"`
				} `json:"Aged 30-44"`
			} `json:"ratingsHistograms"`
		} `json:"ratings"`
	} `json:"resource"`
}

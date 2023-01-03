package goIMDB

import "time"

/*
Holds the various types that are unmarshalled to from the raw API and the Types that are exposed from the package.
Note that some fields from the api response are omitted.
*/

type searchRaw struct {
	Results []struct {
		Image    []interface{} `json:"i"` // as an array!
		ID       string        `json:"id"`
		Title    string        `json:"l"`
		TypeName string        `json:"q"`
		TypeCode string        `json:"qid"`
		TopCast  string        `json:"s"`
		Year     int           `json:"y"`
		Years    string        `json:"yr"`
	} `json:"d"`
	//Q string `json:"q"`
	//V int    `json:"v"`
}

// getTitle - raw with excess removed
type titleRaw struct {
	Resource struct {
		Base struct {
			ID    string `json:"id"`
			Image struct {
				Height int    `json:"height"`
				ID     string `json:"id"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			NextEpisode          string `json:"nextEpisode"`
			NumberOfEpisodes     int    `json:"numberOfEpisodes"`
			RunningTimeInMinutes int    `json:"runningTimeInMinutes"`
			SeriesEndYear        int    `json:"seriesEndYear"`
			SeriesStartYear      int    `json:"seriesStartYear"`
			Title                string `json:"title"`
			TitleType            string `json:"titleType"`
			Year                 int    `json:"year"`
			Episode              int    `json:"episode"`
			ParentTitle          struct {
				ID    string `json:"id"`
				Image struct {
					Height int    `json:"height"`
					ID     string `json:"id"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"parentTitle"`
			PreviousEpisode string `json:"previousEpisode"`
			Season          int    `json:"season"`
		} `json:"base"`
		FilmingLocations []struct {
			Attributes []string `json:"attributes,omitempty"`
			Extras     []string `json:"extras,omitempty"`
			ID         string   `json:"id"`
			Location   string   `json:"location"`
		} `json:"filmingLocations"`
		//MetacriticScore struct {
		//	_Type           string `json:"@type"`
		//	ReviewCount     int    `json:"reviewCount"`
		//	UserRatingCount int    `json:"userRatingCount"`
		//} `json:"metacriticScore"`
		Plot struct {
			Outline struct {
				//ID   string `json:"id"`
				Text string `json:"text"`
			} `json:"outline"`
			Summaries []struct {
				Author string `json:"author,omitempty"`
				//ID     string `json:"id"`
				Text string `json:"text"`
			} `json:"summaries"`
			//TotalSummaries int `json:"totalSummaries"`
		} `json:"plot"`
		Ratings struct {
			Rating      float64 `json:"rating"`
			RatingCount int     `json:"ratingCount"`
		} `json:"ratings"`
		Similarities []struct {
			ID    string `json:"id"`
			Image struct {
				Height int    `json:"height"`
				ID     string `json:"id"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			NextEpisode          string `json:"nextEpisode,omitempty"`
			NumberOfEpisodes     int    `json:"numberOfEpisodes,omitempty"`
			RunningTimeInMinutes int    `json:"runningTimeInMinutes"`
			SeriesEndYear        int    `json:"seriesEndYear,omitempty"`
			SeriesStartYear      int    `json:"seriesStartYear,omitempty"`
			Title                string `json:"title"`
			TitleType            string `json:"titleType"`
			Year                 int    `json:"year"`
		} `json:"similarities"`
		Soundtrack []struct {
			Comment string `json:"comment"`
			//ID           string `json:"id"`
			Name         string `json:"name"`
			RelatedNames []struct {
				Akas           []string `json:"akas,omitempty"`
				Disambiguation string   `json:"disambiguation,omitempty"`
				ID             string   `json:"id"`
				Image          struct {
					Height int    `json:"height"`
					ID     string `json:"id"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"image"`
				LegacyNameText string `json:"legacyNameText"`
				Name           string `json:"name"`
			} `json:"relatedNames"`
		} `json:"soundtrack"`
	} `json:"resource"`
}

// getTitleEpisodesdata
type getEpisodesraw struct {
	Resource struct {
		//ID      string `json:"id"`
		Seasons []struct {
			Season   int `json:"season"`
			Episodes []struct {
				Episode   int    `json:"episode"`
				ID        string `json:"id"`
				Season    int    `json:"season"`
				Title     string `json:"title"`
				TitleType string `json:"titleType"`
				Year      int    `json:"year"`
			} `json:"episodes"`
		} `json:"seasons"`
	} `json:"resource"`
}

// getTitleGenredata
type getGenreRaw struct {
	Resource struct {
		Genres []string `json:"genres"`
	} `json:"resource"`
}

// getTitleVersionsdata
type getVersionsRaw struct {
	Resource struct {
		AlternateTitles []struct {
			Attributes []string `json:"attributes,omitempty"`
			Language   string   `json:"language,omitempty"`
			Region     string   `json:"region,omitempty"`
			Title      string   `json:"title"`
			//Types      []string `json:"types,omitempty"`
		} `json:"alternateTitles"`
		DefaultTitle    string   `json:"defaultTitle"`
		OriginalTitle   string   `json:"originalTitle"`
		Origins         []string `json:"origins"`
		SpokenLanguages []string `json:"spokenLanguages"`
	} `json:"resource"`
}

// get_title_Images
type getImagesRaw struct {
	Resource struct {
		Images []struct {
			Attribution string    `json:"attribution,omitempty"`
			Caption     string    `json:"caption"`
			Copyright   string    `json:"copyright,omitempty"`
			CreatedOn   time.Time `json:"createdOn"`
			Height      int       `json:"height"`
			ID          string    `json:"id"`
			Source      string    `json:"source"`
			Type        string    `json:"type"`
			URL         string    `json:"url"`
			Width       int       `json:"width"`
		} `json:"images"`
		TotalImageCount int `json:"totalImageCount"`
	} `json:"resource"`
}

// raw credits as gotten from API (with many fields removed)
type getCreditsRaw struct {
	Resource struct {
		Credits map[string][]struct {
			Akas           []string `json:"akas,omitempty"`
			Category       string   `json:"category"`
			Disambiguation string   `json:"disambiguation,omitempty"`
			EndYear        int      `json:"endYear"`
			EpisodeCount   int      `json:"episodeCount"`
			ID             string   `json:"id"`
			Image          struct {
				Height int    `json:"height"`
				ID     string `json:"id"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image,omitempty"`
			Job            string `json:"job"`
			LegacyNameText string `json:"legacyNameText"`
			Name           string `json:"name"`
			StartYear      int    `json:"startYear"`
		} `json:"credits"`
	} `json:"resource"`
}

// detailed Name result
type nameDetailedRaw struct {
	Resource struct {
		Base struct {
			Akas              []string `json:"akas"`
			BirthDate         string   `json:"birthDate"`
			BirthPlace        string   `json:"birthPlace"`
			DeathCause        string   `json:"deathCause"`
			DeathDate         string   `json:"deathDate"`
			DeathPlace        string   `json:"deathPlace"`
			Gender            string   `json:"gender"`
			ID                string   `json:"id"`
			HeightCentimeters float64  `json:"heightCentimeters"`

			Image struct {
				Height int    `json:"height"`
				ID     string `json:"id"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"image"`
			IsDead         bool   `json:"isDead"`
			LegacyNameText string `json:"legacyNameText"`
			MiniBios       []struct {
				Author string `json:"author"`
				//ID       string `json:"id"`
				Language string `json:"language"`
				Text     string `json:"text"`
			} `json:"miniBios"`
			Name      string   `json:"name"`
			Nicknames []string `json:"nicknames"`
			RealName  string   `json:"realName"`
			Spouses   []struct {
				Attributes string `json:"attributes"`
				FromDate   string `json:"fromDate"`
				Name       string `json:"name"`
				ToDate     string `json:"toDate"`
			} `json:"spouses"`
		} `json:"base"`
		ID   string   `json:"id"`
		Jobs []string `json:"jobs"`
		//KnownFor []struct {
		//	Cast []struct {
		//		Billing    int      `json:"billing"`
		//		Category   string   `json:"category"`
		//		Characters []string `json:"characters"`
		//		Roles      []struct {
		//			Character   string `json:"character"`
		//			CharacterID string `json:"characterId"`
		//		} `json:"roles"`
		//	} `json:"cast"`
		//	Crew []struct {
		//		Category string `json:"category"`
		//		Job      string `json:"job,omitempty"`
		//	} `json:"crew,omitempty"`
		//	ID      string `json:"id"`
		//	Summary struct {
		//		Category    string   `json:"category"`
		//		Characters  []string `json:"characters"`
		//		DisplayYear string   `json:"displayYear"`
		//	} `json:"summary"`
		//	Title     string `json:"title"`
		//	TitleType string `json:"titleType"`
		//	Year      int    `json:"year"`
		//} `json:"knownFor"`
		//Trivia []struct {
		//	ID               string `json:"id"`
		//	InterestingVotes struct {
		//		Down int `json:"down"`
		//		Up   int `json:"up"`
		//	} `json:"interestingVotes"`
		//	Text string `json:"text"`
		//} `json:"trivia"`
	} `json:"resource"`
}

// holds flags on a title to know if data needs to be called from the API
type titleFlags struct {
	hasFilmLocations bool
	hasExtendedInfo  bool
	isGetTitle       bool // true if this is a true api request or part of a list/parent etc with non full data
	isEpisode        bool
	isSeries         bool
	hasSoundtracks   bool
	hasSimilarities  bool
	hasEpisodes      bool     // true if called get_episodes and all data is here
	hasGenre         bool     // true if genre is called
	hasRelease       bool     // has release information
	hasImages        bool     // true if get_Images called
	hasCredits       bool     // true if get_Credits called
	hasData          bool     // true if there is api data in it
	client           *IMDBPie // client to access web
}

// holds flags on a title to know if data needs to be called from the API

type nameFlags struct {
	hasData bool     // true if there is api data in it
	client  *IMDBPie // client to access web
}

/*
The following types are exposed to the package user, and they are retuned from the various calls
*/

// Size  Image size, used in all images returned by the API
type Size struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

// ImageBasic Basic image data - contain the URL and the size
type ImageBasic struct {
	Url  string `json:"url,omitempty"`
	Size `json:"size"`
}

// ImageDetail Image with more details (as getting it from the image api)
type ImageDetail struct {
	ImageBasic  `json:"imageBasic"`
	Caption     string    `json:"caption,omitempty"`
	CreatedOn   time.Time `json:"createdOn"`
	Copyright   string    `json:"copyright,omitempty"`
	Attribution string    `json:"attribution,omitempty"`
	Source      string    `json:"source,omitempty"`
	Type        string    `json:"type,omitempty"`
}

// Rating basic Rating info
type Rating struct {
	Rating      float64 `json:"rating"`
	RatingCount int     `json:"ratingCount"`
}

// SingleCredit a single credit for a person in a title
type SingleCredit struct {
	ID         string `json:"ID,omitempty"`
	Name       string `json:"name,omitempty"`
	Category   string `json:"category,omitempty"`
	Image      `json:"image"`
	Akas       []string `json:"akas,omitempty"`
	Job        string   `json:"job,omitempty"`
	LegacyName string   `json:"legacyName,omitempty"`
}

// SingleSearchTitleResult a single result when searching for a title
type SingleSearchTitleResult struct {
	Image    ImageBasic `json:"image"`           // imageData
	ID       string     `json:"ID"`              // imdb ID
	Title    string     `json:"title,omitempty"` //title
	TypeName string     `json:"typeName,omitempty"`
	TypeCode string     `json:"typeCode,omitempty"`
	TopCast  []string   `json:"topCast,omitempty"` //top names
	Year     int        `json:"year,omitempty"`    // year out
	Years    string     `json:"years,omitempty"`   // year range if multiple years
}

// SingleSearchNameResult a single result when searching for a name
type SingleSearchNameResult struct {
	Image    ImageBasic `json:"image"`              // imageData
	ID       string     `json:"ID"`                 // imdb ID
	Name     string     `json:"name,omitempty"`     //title
	KnownFor string     `json:"knownFor,omitempty"` //top names
}

// SoundTrack a single SoundTrack result
type SoundTrack struct {
	Comment      string      `json:"comment" `
	Name         string      `json:"name" `
	RelatedNames []NameBasic `json:"relatedNames"`
}

// FilmingLocation a single Film Location
type FilmingLocation struct {
	Attributes []string `json:"attributes,omitempty"` // where (i.e. inside, outside)
	Comment    string   `json:"comment,omitempty"`    // additional info
	Location   string   `json:"location,omitempty"`   // location mame
}

// PlotSummary a single Plot Summary
type PlotSummary struct {
	Author      string `json:"author,omitempty"`
	PlotSummary string `json:"PlotSummary,omitempty"`
}

// NameBasic Basic Name / person information
type NameBasic struct {
	ID         string     `json:"ID,omitempty"`
	Akas       []string   `json:"akas,omitempty"`
	Image      ImageBasic `json:"image"`
	Name       string     `json:"name,omitempty"`
	LegacyName string     `json:"legacyName,omitempty"`
}

// NameDetailed detailed name (I excluded known-for items) seems too much info
type NameDetailed struct {
	NameBasic
	BirthDate         string  `json:"birthDate"`
	BirthPlace        string  `json:"birthPlace"`
	Gender            string  `json:"gender"`
	HeightCentimeters float64 `json:"heightCentimeters"`
	MiniBios          []struct {
		Author string `json:"author"`
		//ID       string `json:"id"`
		Language string `json:"language"`
		Text     string `json:"text"`
	} `json:"miniBios"`
	Nicknames []string `json:"nicknames"`
	RealName  string   `json:"realName"`
	Spouses   []struct {
		Attributes string `json:"attributes"`
		FromDate   string `json:"fromDate"`
		Name       string `json:"name"`
		ToDate     string `json:"toDate"`
	} `json:"spouses"`
	Jobs []string `json:"jobs"`
	nameFlags
}

// SingleEpisode single episode info
type SingleEpisode struct {
	ID        string `json:"id"`
	Season    int    `json:"season"`
	Episode   int    `json:"episode"`
	Title     string `json:"title"`
	TitleType string `json:"titleType"`
	Year      int    `json:"year"`
}

// SingleSeason single season & its episodes
type SingleSeason struct {
	Season   int             `json:"season,omitempty"`
	Episodes []SingleEpisode `json:"episodes,omitempty"`
}

// SingleVersion one Version of a title
type SingleVersion struct {
	Attributes []string `json:"attributes,omitempty"`
	Language   string   `json:"language,omitempty"`
	Region     string   `json:"region,omitempty"`
	Title      string   `json:"title"`
}

// VersionInfo Versions information
type VersionInfo struct {
	Releases        []SingleVersion `json:"releases,omitempty"`
	OriginalTitle   string          `json:"originalTitle,omitempty"`
	DefaultTitle    string          `json:"defaultTitle,omitempty"`
	Origins         []string        `json:"origins,omitempty"`
	SpokenLanguages []string        `json:"spokenLanguages,omitempty"`
}

// TitleMinimal minimal info title (this is returned in similarities etc)
type TitleMinimal struct {
	ID        string     `json:"id"`
	Image     ImageBasic `json:"image"`
	Title     string     `json:"title"`
	TitleType string     `json:"titleType"`
	Year      int        `json:"year"`
	RunTime   int        `json:"runningTimeInMinutes"`
}

// TitleBasic Basic title information
type TitleBasic struct {
	TitleMinimal
	SeriesEndYear    int    `json:"seriesEndYear"`
	SeriesStartYear  int    `json:"seriesStartYear"`
	Episode          int    `json:"episode"`
	Season           int    `json:"season"`
	NextEpisode      string `json:"nextEpisode"`
	PreviousEpisode  string `json:"previousEpisode"`
	NumberOfEpisodes int    `json:"numberOfEpisodes"`
	PlotOutline      string
	ParentTitle      TitleMinimal // holds the id of the parent series

}

// Title type - This type holds the most basic information from the API. Additional calls to its various
// methods will return additional data
type Title struct {
	TitleBasic    `json:"titleBasic"`
	PlotSummaries []PlotSummary `json:"plotSummaries,omitempty"`
	Rating        Rating        `json:"rating" json:"rating"`
	// entries below are only available via Methods, and are not visible directly from outside the package
	filmingLocations []FilmingLocation
	soundTracksBasic []SoundTrack
	similarities     []TitleMinimal
	allEpisodes      []SingleSeason
	credits          map[string][]SingleCredit // map from category (i.e. cast/writes etc) to the persons
	genres           []string
	versions         VersionInfo
	images           []ImageDetail
	titleFlags       // holds all the relevant information we have in this title [interal use only]
}

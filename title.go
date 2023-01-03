package goIMDB

import (
	"errors"
	"fmt"
	"log"
)

/*
Title struct holds all info about a specific title.  Once the GetTitle is called, additional methods on the resulting
object gets further details on that title.  Only a subset of possible api calls were implemented as the rest are less
obviously needed.  However, they are simple to implement.
todo: implement additional api calls and data gathering
*/

// error messages
var ERROR_NO_TITLE_DATA = errors.New("title has no api data, please call get title prior to calling its methods")
var ERROR_NO_EPISODES_FOR_TITLE = errors.New("title is not a tvSeries so can't get episodes")

// GetTitle - Get Basic Info on a title (and film locations, soundtracks) using a single api call
// the call also check for a valid ttxxxxxx id code and that this id is not redirected to a new one
func (m *IMDBPie) GetTitle(imdbid string) (Title, error) {
	if e := validateImdbTitleID(imdbid); e != nil {
		return Title{}, e
	}
	if e := m.redirectionCheck(imdbid); e != nil {
		return Title{}, e
	}
	// get all the data
	aux := titleRaw{}
	e := m.getTitleApiData(routeGetTitle, imdbid, &aux)
	//aux, e := m.getTitleBase(imdbid)
	if e != nil {
		return Title{}, e
	}
	if m.excludeEp && aux.Resource.Base.TitleType == "tvEpisode" {
		return Title{}, ErrEpisodeTitle{msg: fmt.Sprintf("Title %s is an episode title, and exclude episode is set", imdbid)}
	}
	ttl := Title{
		TitleBasic: TitleBasic{
			TitleMinimal: TitleMinimal{
				RunTime:   aux.Resource.Base.RunningTimeInMinutes,
				Title:     aux.Resource.Base.Title,
				TitleType: aux.Resource.Base.TitleType,
				Year:      aux.Resource.Base.Year,
				ID:        getImdbID(aux.Resource.Base.ID),
				Image: ImageBasic{
					Url: aux.Resource.Base.Image.URL,
					Size: Size{
						Height: aux.Resource.Base.Image.Height,
						Width:  aux.Resource.Base.Image.Width,
					},
				},
			},
			NextEpisode:      aux.Resource.Base.NextEpisode,
			PreviousEpisode:  aux.Resource.Base.PreviousEpisode,
			NumberOfEpisodes: aux.Resource.Base.NumberOfEpisodes,
			SeriesEndYear:    aux.Resource.Base.SeriesEndYear,
			SeriesStartYear:  aux.Resource.Base.SeriesStartYear,
			PlotOutline:      aux.Resource.Plot.Outline.Text,
			Season:           aux.Resource.Base.Season,
			Episode:          aux.Resource.Base.Episode,
			ParentTitle: TitleMinimal{
				ID: getImdbID(aux.Resource.Base.ParentTitle.ID),
				Image: ImageBasic{
					Url: aux.Resource.Base.ParentTitle.Image.URL,
					Size: Size{
						Height: aux.Resource.Base.ParentTitle.Image.Height,
						Width:  aux.Resource.Base.ParentTitle.Image.Width,
					},
				},
				Title:     aux.Resource.Base.ParentTitle.Title,
				TitleType: aux.Resource.Base.ParentTitle.TitleType,
				Year:      aux.Resource.Base.ParentTitle.Year,
			},
		},
		Rating: Rating{
			Rating:      aux.Resource.Ratings.Rating,
			RatingCount: aux.Resource.Ratings.RatingCount,
		},
		titleFlags: titleFlags{
			hasFilmLocations: true,
			hasExtendedInfo:  false,
			hasSimilarities:  true,
			hasSoundtracks:   true,
			isGetTitle:       true,
			client:           m,
		},
	}
	ttl.PlotSummaries = []PlotSummary{}
	for _, ps := range aux.Resource.Plot.Summaries {
		ttl.PlotSummaries = append(ttl.PlotSummaries, PlotSummary{
			Author:      ps.Author,
			PlotSummary: ps.Text,
		})
	}
	ttl.filmingLocations = []FilmingLocation{}
	for _, fl := range aux.Resource.FilmingLocations {
		extra := ""
		if len(fl.Extras) > 0 {
			extra = fl.Extras[0]
		}
		if fl.Attributes == nil {
			fl.Attributes = []string{}
		}
		ttl.filmingLocations = append(ttl.filmingLocations, FilmingLocation{
			Attributes: fl.Attributes,
			Comment:    extra,
			Location:   fl.Location,
		})
	}
	ttl.soundTracksBasic = []SoundTrack{}
	for _, st := range aux.Resource.Soundtrack {
		var rnlst []NameBasic
		if len(st.RelatedNames) > 0 {
			for _, rn := range st.RelatedNames {
				if rn.Akas == nil {
					rn.Akas = []string{}
				}
				rnlst = append(rnlst, NameBasic{
					ID:   getImdbID(rn.ID),
					Akas: rn.Akas,
					Image: ImageBasic{
						Url: rn.Image.URL,
						Size: Size{
							Height: rn.Image.Height,
							Width:  rn.Image.Width,
						},
					},
					Name:       rn.Name,
					LegacyName: rn.LegacyNameText,
				})
			}
		}
		ttl.soundTracksBasic = append(ttl.soundTracksBasic, SoundTrack{
			Comment:      st.Comment,
			Name:         st.Name,
			RelatedNames: rnlst,
		})
	}
	ttl.similarities = []TitleMinimal{}
	for _, sim := range aux.Resource.Similarities {
		ttl.similarities = append(ttl.similarities, TitleMinimal{
			ID: getImdbID(sim.ID),
			Image: ImageBasic{
				Url: sim.Image.URL,
				Size: Size{
					Height: sim.Image.Height,
					Width:  sim.Image.Width,
				},
			},
			Title:     sim.Title,
			TitleType: sim.TitleType,
			Year:      sim.Year,
			RunTime:   sim.RunningTimeInMinutes,
		})

	}
	ttl.titleFlags.isSeries = ttl.TitleType == "tvSeries" || ttl.TitleType == "tvMiniSeries"
	ttl.titleFlags.isEpisode = ttl.TitleType == "tvEpisode"
	ttl.titleFlags.hasData = true
	return ttl, nil
}

// GetTitleEpisodes - get all episodes for a series title
func (t *Title) GetTitleEpisodes() ([]SingleSeason, error) {
	if !t.hasData {
		return nil, ERROR_NO_TITLE_DATA
	}
	if !t.isSeries {
		return nil, ERROR_NO_EPISODES_FOR_TITLE
	}
	if t.hasEpisodes {
		return t.allEpisodes, nil
	}
	m := t.titleFlags.client
	if m == nil {
		m = InitIMDBPie()
		log.Println("Title Client is NIL - should not be!")
	}
	aux := getEpisodesraw{}
	//aux, e := m.getTitleEpisodes(t.ID)
	e := m.getTitleApiData(routeGetTitleEpisodes, t.ID, &aux)
	if e != nil {
		return nil, e
	}
	var result []SingleSeason
	for _, sn := range aux.Resource.Seasons {
		var eplist []SingleEpisode
		for _, ep := range sn.Episodes {
			eplist = append(eplist, SingleEpisode{
				ID:        getImdbID(ep.ID),
				Season:    ep.Season,
				Episode:   ep.Episode,
				Title:     ep.Title,
				TitleType: ep.TitleType,
				Year:      ep.Year,
			})
		}
		result = append(result, SingleSeason{
			Season:   sn.Season,
			Episodes: eplist,
		})
	}
	t.allEpisodes = result
	t.titleFlags.hasEpisodes = true
	return result, nil
}

// get the Genres of a title
func (t *Title) GetTitleGenres() ([]string, error) {
	if !t.hasData {
		return nil, ERROR_NO_TITLE_DATA
	}
	if t.hasGenre {
		log.Println("Cache generes met")
		return t.genres, nil
	}
	m := t.titleFlags.client
	if m == nil {
		m = InitIMDBPie()
		log.Println("Title Client is NIL - should not be!")
	}
	var aux getGenreRaw
	//aux, e := m.getTitleGenres(t.ID)
	e := m.getTitleApiData(routeGetTitleGenres, t.ID, &aux)
	if e != nil {
		return nil, e
	}
	t.titleFlags.hasGenre = true
	t.genres = aux.Resource.Genres
	return t.genres, nil
}

// get Versions (i.e. different countires, languages etc) for a title
func (t *Title) GetTitleVersions() (vr VersionInfo, err error) {
	if !t.hasData {
		return t.versions, ERROR_NO_TITLE_DATA
	}
	if t.hasRelease {
		log.Println("Cache generes met")
		return t.versions, nil
	}
	m := t.titleFlags.client
	if m == nil {
		m = InitIMDBPie()
		log.Println("Title Client is NIL - should not be!")
	}
	var aux getVersionsRaw
	//aux, e := m.getTitleVersions(t.ID)
	e := m.getTitleApiData(routeGetTitleVersions, t.ID, &aux)
	if e != nil {
		return t.versions, e
	}
	vr = VersionInfo{
		Releases:        []SingleVersion{},
		OriginalTitle:   aux.Resource.OriginalTitle,
		Origins:         aux.Resource.Origins,
		SpokenLanguages: aux.Resource.SpokenLanguages,
		DefaultTitle:    aux.Resource.DefaultTitle,
	}
	if vr.Origins == nil {
		vr.Origins = []string{}
	}
	for _, r := range aux.Resource.AlternateTitles {
		if r.Attributes == nil {
			r.Attributes = []string{}
		}
		vr.Releases = append(vr.Releases, SingleVersion{
			Attributes: r.Attributes,
			Language:   r.Language,
			Region:     r.Region,
			Title:      r.Title,
		})
	}
	t.titleFlags.hasRelease = true
	t.versions = vr
	return t.versions, nil
}

// get Images  for a title
func (t *Title) GetTitleImages() (imgs []ImageDetail, err error) {
	if !t.hasData {
		return t.images, ERROR_NO_TITLE_DATA
	}
	if t.hasImages {
		log.Println("Cache generes met")
		return t.images, nil
	}
	m := t.titleFlags.client
	if m == nil {
		m = InitIMDBPie()
		log.Println("Title Client is NIL - should not be!")
	}
	var aux getImagesRaw
	e := m.getTitleApiData(routeGetTitleImages, t.ID, &aux)
	//aux, e := m.getTitleImages(t.ID)
	if e != nil {
		return t.images, e
	}
	for _, im := range aux.Resource.Images {
		imgs = append(imgs, ImageDetail{
			ImageBasic: ImageBasic{
				Url: im.URL,
				Size: Size{
					Height: im.Height,
					Width:  im.Width,
				},
			},
			Caption:     im.Caption,
			CreatedOn:   im.CreatedOn,
			Copyright:   im.Copyright,
			Attribution: im.Attribution,
			Source:      im.Source,
			Type:        im.Type,
		})
	}
	log.Println("total of ", len(imgs), "Images for ", t.Title)
	t.titleFlags.hasImages = true
	t.images = imgs
	return t.images, nil
}

// get the credits for a title (all the possible positions and people that were involved in it).
func (t *Title) GetTitleCredits() (crds map[string][]SingleCredit, err error) {
	if !t.hasData {
		return t.credits, ERROR_NO_TITLE_DATA
	}
	if t.hasCredits {
		log.Println("Cache Credits met")
		return t.credits, nil
	}
	m := t.titleFlags.client
	if m == nil {
		m = InitIMDBPie()
		log.Println("Title Client is NIL - should not be!")
	}
	var aux getCreditsRaw
	//aux, e := m.getTitleCredits(t.ID)
	e := m.getTitleApiData(routeGetTitleCredits, t.ID, &aux)
	if e != nil {
		return t.credits, e
	}
	crds = make(map[string][]SingleCredit, 0)
	for k, v := range aux.Resource.Credits {
		for _, cr := range v {
			if cr.Akas == nil {
				cr.Akas = []string{}
			}
			crds[k] = append(crds[k], SingleCredit{
				Akas:     cr.Akas,
				Category: cr.Category,
				ID:       getImdbID(cr.ID),
				Image: Image{
					Url:    cr.Image.URL,
					Width:  cr.Image.Width,
					Height: cr.Image.Height,
				},
				Job:        cr.Job,
				LegacyName: cr.LegacyNameText,
				Name:       cr.Name,
			})
		}
	}
	t.titleFlags.hasCredits = true
	t.credits = crds
	return t.credits, nil
}

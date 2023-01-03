package goIMDB

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

/*
Those are the implemented and not implemented raw api calls as copied from the imdbpy package.
There are title,name,search,special and general route types.
*/

// name routes - calls for nmXXXXXX codes
var apiNameRoutes = map[int8]string{
	getNameImages:      "/name/%s/images",
	getNameVideos:      "/name/%s/videos",
	getName:            "/name/%s/fulldetails",
	getNameFilmography: "/name/%s/filmography",
}

// search api route
const apiSearchRoute = "/suggests/%s/%s.json"

// mapping int to named constants for readability
const (
	getNameImages      = 1
	getNameVideos      = 2
	getName            = 3
	getNameFilmography = 4

	routeGetTitle                  = 1
	routeGetTitleMetaCriticReviews = 2
	routeGetTitleUserReviews       = 3
	routeGetTitleVideos            = 4
	routeGetTitleImages            = 5
	routeGetTitleCompanies         = 6
	routeGetTitleTechnical         = 7
	routeGetTitleTrivia            = 8
	routeGetTitleGoofs             = 9
	routeGetTitleSoundtracks       = 10
	routeGetTitleNews              = 11
	routeGetTitlePlot              = 12
	routeGetTitlePlotSynopsis      = 13
	routeGetTitlePlotTaglines      = 14
	routeGetTitleVersions          = 15
	routeGetTitleReleases          = 16
	routeGetTitleQuotes            = 17
	routeGetTitleConnections       = 18
	routeGetTitleGenres            = 19
	routeGetTitleSimilarities      = 20
	routeGetTitleAwards            = 21
	routeGetTitleRatings           = 22
	routeGetTitleCredits           = 23
	routeGetTitleEpisodes          = 24

	getChartTitles = 1
	getChartTv     = 2
	getChartMovies = 3
)

// Title route map between name and format string
var apiTitleRoutes = map[int8]string{
	routeGetTitle:                  "/title/%s/auxiliary",
	routeGetTitleMetaCriticReviews: "/title/%s/metacritic",
	routeGetTitleUserReviews:       "/title/%s/userreviews",
	routeGetTitleVideos:            "/title/%s/videos",
	routeGetTitleImages:            "/title/%s/images",
	routeGetTitleCompanies:         "/title/%s/companies",
	routeGetTitleTechnical:         "/title/%s/technical",
	routeGetTitleTrivia:            "/title/%s/trivia",
	routeGetTitleGoofs:             "/title/%s/goofs",
	routeGetTitleSoundtracks:       "/title/%s/soundtracks",
	routeGetTitleNews:              "/title/%s/news",
	routeGetTitlePlot:              "/title/%s/plot",
	routeGetTitlePlotSynopsis:      "/title/%s/plotsynopsis",
	routeGetTitlePlotTaglines:      "/title/%s/taglines",
	routeGetTitleVersions:          "/title/%s/versions",
	routeGetTitleReleases:          "/title/%s/releases",
	routeGetTitleQuotes:            "/title/%s/quotes",
	routeGetTitleConnections:       "/title/%s/connections",
	routeGetTitleGenres:            "/title/%s/genres",
	routeGetTitleSimilarities:      "/title/%s/similarities",
	routeGetTitleAwards:            "/title/%s/awards",
	routeGetTitleRatings:           "/title/%s/ratings",
	routeGetTitleCredits:           "/title/%s/fullcredits",
	routeGetTitleEpisodes:          "/title/%s/episodes",
}

// global routes
var apiGlobalRoutes = map[int8]string{
	getChartTitles: "/chart/titlemeter",
	getChartTv:     "/chart/tvmeter",
	getChartMovies: "/chart/tvmeterrt/moviemete",
}

// specialRoutes - routes that needs additional query parameters
type specialRoutes struct {
	route  string            // the non base url
	params map[string]string // the parameters to encode into the request
}

// set the parameters of the route
func (sp *specialRoutes) SetParams(p ...[2]string) {
	for _, pr := range p {
		if pr[0] == "" {
			log.Panic("Bad call to SetParams - no key!", p)
		}
		sp.params[pr[0]] = pr[1]
	}
}

// special routes, user should add the params as needed per each only the fixed parameters are included here
// todo: implement them
var apiSpecialRoutes = map[string]specialRoutes{
	// get addional data on title, need to supply : "tconst":imdbId,"today":time.Now().Format("2006-01-02")
	"routeGetTitleAuxiliary": {
		route: "/template/imdb-ios-writable/title-auxiliary-v31.jstl/render",
		params: map[string]string{
			"inlineBannerAdWeblabOn": "false",
			"minwidth":               "320",
			"osVersion":              "11.3.0",
		},
	},
	// get top crew for a title, need to supply :  "tconst": imdbi_id
	"routeGetTitleTopCrew": {route: "/template/imdb-android-writable/7.3.top-crew.jstl/render"},
	// get specific range of episodes: need to supply the following:
	//'end': limit,  - last episode to show (in that season)(not zero based)
	//'start': offset, - first episode to show (zero based)
	//'season': season - 1, season (zero based)
	//'tconst': imdb_id, series ID
	"routeGetTitleTvEpisodes": {route: "/template/imdb-ios-writable/tv-episodes-v2.jstl/render"},
}

// compose the calling url.
func composeUrl(url string, data string) string {
	p1 := fmt.Sprintf("%s%s", BaseUri, url)
	if data != "" {
		return fmt.Sprintf(p1, data)
	} else {
		return p1
	}
}

// various imdb code validation/extraction

var RegexImdbID = regexp.MustCompile(`^(tt|nm)\d{5,8}$`)
var RegexImdbTitle = regexp.MustCompile(`^tt\d{5,8}$`)
var RegexImdbName = regexp.MustCompile(`^nm\d{5,8}$`)
var RegexImdbExtract = regexp.MustCompile(`((?:tt|nm)\d{5,8})(?:$|\D)`)

// general check on imdb id
func validateImdbId(ID string) error {
	result := RegexImdbID.FindString(ID)
	if result == "" {
		return errors.New("malformed imdb id given: " + ID)
	}
	return nil
}

// check imdb_id is a valid title code
func validateImdbTitleID(imdbID string) error {
	result := RegexImdbTitle.FindString(imdbID)
	if result == "" {
		return errors.New("malformed Title Imdb id given: " + imdbID)
	}
	return nil
}

// check imdb_id is a valid name code
func validateImdbNameID(imdbID string) error {
	result := RegexImdbName.FindString(imdbID)
	if result == "" {
		return errors.New("malformed Name Imdb id given: " + imdbID)
	}
	return nil
}

// extract the imdb code from the returned values in the api (they sometime return additional path parts)
func getImdbID(s string) string {
	result := RegexImdbExtract.FindStringSubmatch(s)
	if len(result) > 0 {
		return result[1]
	} else {
		return ""
	}

}

// check if an imdbid is redirected to a new one and retun an error if so.
func (m *IMDBPie) redirectionCheck(imdbID string) error {
	if isRedirection(imdbID) {
		return ErrRedirect{msg: fmt.Sprintf("title %s is redirected use the new value", imdbID)}
	}
	return nil
}

// isRedirection - check if page is redirected by not allowing redirects and checking the resulting status code
// return true if the page is redirected
func isRedirection(imdbID string) bool {
	pageUrl := fmt.Sprintf("https://www.imdb.com/title/%s/", imdbID)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	// todo: add error handling - in case of internet problems the program fail here
	res, e := client.Get(pageUrl)
	if e != nil {
		return false
	} // assume it's not redirected if get does not work.
	return res.StatusCode == http.StatusMovedPermanently
}

// parseDate - convert date to Time type, possible as a year only or full calendar date
// return time zero if not and a special error
func parseDate(dt string) time.Time {
	date, err := time.Parse("2006-01-02", dt)
	if err == nil {
		return date
	}
	// try year only
	yr, err := strconv.Atoi(dt)
	if err == nil && yr > 1900 && yr < 2100 {
		return time.Date(yr, 1, 1, 0, 0, 0, 0, &time.Location{})
	} else {
		log.Printf("Bad Date in api: %s Ignoring\n", dt)
	}
	return time.Time{}
}

// getApiData call one of the above simple api routes, and json unmarshal the data to the given variable
func (m *IMDBPie) getApiData(endpoint string, imdbID string, jResult interface{}) error {
	url := composeUrl(endpoint, imdbID)
	res, e := m.makeImdbRequest("GET", url, nil, nil)
	if e != nil {
		return e
	}
	resBody, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	e = json.Unmarshal(resBody, jResult)
	return e
}

// get Title API by route code, the result in marshaled in the jResult struct
func (m *IMDBPie) getTitleApiData(routeCode int8, imdbID string, jResult interface{}) error {
	return m.getApiData(apiTitleRoutes[routeCode], imdbID, jResult)
}

// get Name API by route code, the result in marshaled in the jResult struct
func (m *IMDBPie) getNameApiData(routeCode int8, imdbID string, jResult interface{}) error {
	return m.getApiData(apiNameRoutes[routeCode], imdbID, jResult)
}

// get Search Api data and return the marshaled results
func (m *IMDBPie) getSearchApiData(query string, jResult interface{}) error {
	trimmedName := fixQuery(query)
	firstLetter := firstAlphaNum(trimmedName)
	if firstLetter == "" {
		return errors.New("query has no Alphanumeric characters in it")
	}
	sQuery := fmt.Sprintf(apiSearchRoute, firstLetter, query)
	url := fmt.Sprintf("%s%s", SearchBaseUri, sQuery)
	res, e := m.makeImdbRequest("GET", url, nil, nil)
	if e != nil {
		return e
	}
	body, _ := io.ReadAll(res.Body)
	r := cleanJson(&body)
	return json.Unmarshal(*r, jResult)

}

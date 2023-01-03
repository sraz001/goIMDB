package goIMDB

import (
	"fmt"
	"strings"
)

/*
Handles search request for titles and names.
*/

// SearchTitle - search IMDB for the title with an optional year (which is just added as a string to the search).
// note, non title results will not be returned.  (use Name search to search for people)
func (m *IMDBPie) SearchTitle(tName string, tYear int) ([]SingleSearchTitleResult, error) {
	if tYear > 1900 && tYear < 2100 {
		tName = fmt.Sprintf("%s %d", tName, tYear)
	}
	result := []SingleSearchTitleResult{}
	var v searchRaw
	e := m.getSearchApiData(tName, &v)
	if e != nil {
		fmt.Println("Error in searchTitle Api!")
		return result, e
	}
	for _, itm := range v.Results {
		// ignore non title entries
		if !strings.HasPrefix(itm.ID, "tt") {
			continue
		}
		sr := SingleSearchTitleResult{
			ID:       itm.ID,
			Title:    itm.Title,
			TypeName: itm.TypeName,
			TypeCode: itm.TypeCode,
			Year:     itm.Year,
			Years:    itm.Years,
		}
		if itm.TopCast != "" {
			parts := strings.Split(itm.TopCast, ",")
			for i, j := range parts {
				parts[i] = strings.TrimSpace(j)
			}
			sr.TopCast = parts
		}
		if len(itm.Image) > 0 {
			sr.Image = ImageBasic{
				Url:  itm.Image[0].(string),
				Size: Size{Height: int(itm.Image[2].(float64)), Width: int(itm.Image[1].(float64))},
			}
		}
		result = append(result, sr)
	}
	return result, nil
}

// SerachName - search IMDB for a name result.  Only name results will be returned.
func (m *IMDBPie) SearchName(tName string) ([]SingleSearchNameResult, error) {
	var result []SingleSearchNameResult
	var v searchRaw
	e := m.getSearchApiData(tName, &v)
	if e != nil {
		fmt.Println("Error in SearchName api!!!")
		return result, e
	}
	for _, itm := range v.Results {
		// ignore non Name entries
		if !strings.HasPrefix(itm.ID, "nm") {
			continue
		}
		sr := SingleSearchNameResult{
			ID:       itm.ID,
			Name:     itm.Title,
			KnownFor: itm.TopCast,
		}
		if len(itm.Image) > 0 {
			sr.Image = ImageBasic{
				Url:  itm.Image[0].(string),
				Size: Size{Height: int(itm.Image[2].(float64)), Width: int(itm.Image[1].(float64))},
			}
		}
		result = append(result, sr)
	}
	return result, nil
}

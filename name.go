package goIMDB

import "errors"

/*
Handles all calls to Name structure - a person name and its various additional information.
Only the basic call was implemented, additional available api calls were not implemented
todo: implement additional calls
*/

var ERROR_NO_NAME_DATA = errors.New("name has no api data, please call getName prior to calling its methods")

// GetName - Get Info on a Name by its ID.
func (m *IMDBPie) GetName(imdbid string) (nmD NameDetailed, err error) {
	if e := validateImdbNameID(imdbid); e != nil {
		return nmD, e
	}
	if e := m.redirectionCheck(imdbid); e != nil {
		return nmD, e
	}
	// get all the data
	var aux nameDetailedRaw
	e := m.getNameApiData(getName, imdbid, &aux)
	if e != nil {
		return nmD, e
	}
	nmd := NameDetailed{
		NameBasic: NameBasic{
			ID:   getImdbID(aux.Resource.ID),
			Akas: aux.Resource.Base.Akas,
			Image: ImageBasic{
				Url: aux.Resource.Base.Image.URL,
				Size: Size{
					Height: aux.Resource.Base.Image.Height,
					Width:  aux.Resource.Base.Image.Width,
				},
			},
			Name:       aux.Resource.Base.Name,
			LegacyName: aux.Resource.Base.LegacyNameText,
		},
		BirthDate:         aux.Resource.Base.BirthDate,
		BirthPlace:        aux.Resource.Base.BirthPlace,
		Gender:            aux.Resource.Base.Gender,
		HeightCentimeters: aux.Resource.Base.HeightCentimeters,
		MiniBios:          aux.Resource.Base.MiniBios,
		Nicknames:         aux.Resource.Base.Nicknames,
		RealName:          aux.Resource.Base.RealName,
		Spouses:           aux.Resource.Base.Spouses,
		Jobs:              aux.Resource.Jobs,
		nameFlags: nameFlags{
			hasData: true,
			client:  m,
		},
	}
	return nmd, nil
}

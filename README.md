# goIMDB 

A simple and effective library to query the IMDB database using the unofficial mobile app API.

# Description

A frontend to access the unofficial mobile API of IMDB, the famous movie database.  The package allows among others to:

* Search for Titles
* Search for People
* Get Title information, including additional details on cast, genre, versions etc.
* Get People information, including image, jobs, etc.

The implemented calls are only a subset of all available api calls, and in the future, upon need, I will add the remaining
calls.

This package was inspired by the python package [imdb-pie](https://github.com/richardARPANET/imdb-pie) and the unofficial api
calls are copied from that package.

This package handles three parts of the vast IMDB information:
* `Title` - represent a movie/tv title and has a code similar to ```tt123435```
* `Name` - represent a person (actor, director etc.) and has a code similar to ```nm88888```
* `Search` - search results return very basic information on a list of possible matches - and is ordered by their internal, very good, search algo.


# Installation

This package can be installed with the `go get` command:

    go get github.com/sraz001/goIMDB


# API Reference

API documentation can be found [here](http://godoc.org/github.com/sraz001/goIMDB).

# Sample Usage

Below is an example to search for a title, pick the first result (note that the search results are sorted per IMDB sorting algo)
and get more details about the results (like plot and cast), then print some details on the first actor.

```go
package main

import (
  "fmt"
  "github.com/sraz001/goIMDB"
  "log"
)

func main() {
  m := goIMDB.InitIMDBPie()
  // search for The Matrix title:
  result, e := m.SearchTitle("The Matrix", 1999)
  if e != nil {
    log.Fatal(e)
  }
  // print the first result
  fmt.Printf("Got %d Search Results, the first one is %s - %s (%d)\n", len(result), result[0].ID, result[0].Title, result[0].Year)
  // Got 8 Search Results, the first one is tt0133093 - The Matrix (1999)

  // get Basic title information
  title, e := m.GetTitle(result[0].ID)
  if e != nil {
    log.Fatal(e)
  }
  fmt.Printf("Plot of %s is [ %s ]\n", title.ID, title.PlotOutline)
  // Plot of tt0133093 is [ When a beautiful stranger leads computer hacker Neo to a forbidding underworld, he discovers the shocking truth--the life he knows is the elaborate deception of an evil cyber-intelligence. ]

  // get the cast of the title
  credits, _ := title.GetTitleCredits() // skip error check for readability
  actor := credits["cast"][0]           // get first actor
  fmt.Printf("Actor %s is %s with Image url: %s", actor.ID, actor.Name, actor.Image.Url)
  // Actor nm0000206 is Keanu Reeves with Image url: https://m.media-amazon.com/images/M/MV5BNGJmMWEzOGQtMWZkNS00MGNiLTk5NGEtYzg1YzAyZTgzZTZmXkEyXkFqcGdeQXVyMTE1MTYxNDAw._V1_.jpg

}
```


### Features & Future Development

Below is a short table of the implemented api at this stage.

| Example                              | Description                                                   |
|--------------------------------------|---------------------------------------------------------------|
| ```m := goIMDB.InitIMDBPie()```      | return a client to access the Title,Name and Search functions |
| ```t := m.GetTitle(imdbCode)```      | return a Title Struct that holds the title information        |
| ```nm := m.GetName(imdbCode)```      | return a Name Struct that holds the title information         |
| ```t.GetTitleEpisodes()```           | return all episodes for a tv series                           |
| ```t.GetTitleGenres()```             | return all genres for a title                                 |
| ```t.GetTitleImages()```             | return all images for a title                                 |
| ```t.GetTitleCredits()```            | return all credits for a title (large amount of data!)        |
| ```m.SearchTitle(Searchterm,year)``` | lookup Titles with optional year (set to zero to ignore)      |
| ```m.SearchName(Searchterm,)```      | lookup Names with optional year (set to zero to ignore)       |

The entire code is documented and all the available api calls are included (although most are not used).

#### Todo:

*   add disk caching to the program, as the api calls are rather slow.  Caching might be done on a url basis (before encryption)
*   implement additional api calls on Name and Title
*  add tests & examples
*  add better error messages and handling



### License

MIT: http://mattn.mit-license.org/2018

all contributions are welcome.



package goIMDB

import (
	"bytes"
	"encoding/json"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"regexp"
	"strings"
	"unicode"
)

/*
Various used and potential use functions
todo: remove uneeded ones
*/
func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}

func isalphenum(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z')
}

// returns first rune that is alphanumeric as a string
func firstAlphaNum(s string) string {
	for _, r := range s {
		if isalphenum(r) {
			return string(r)
		}
	}
	return ""
}

// replace spaces with _ and remove non alpha
func fixQuery(q string) string {
	re := regexp.MustCompile(`\W+`)
	normalizedName := re.ReplaceAllString(q, "_")
	return strings.Trim(normalizedName, "_")
}

// cleans json that is returned from searches, as its has an invalid start/end
func cleanJson(badJson *[]byte) *[]byte {
	re := regexp.MustCompile(`imdb\$.+\((\{.+\})\)$`)
	matches := re.FindSubmatch(*badJson)
	if len(matches) > 1 {
		//fmt.Printf("Before: \n%s\n------------------------\n\n", string(*badJson))
		//fmt.Printf("after: \n%s\n-----------------------\n\n", string(matches[1]))
		return &matches[1]
	}
	return badJson
}

// pretty print a struct for debugging
// todo: remove this
func pprintStruct(p any) string {
	b1, e := json.Marshal(p)
	if e != nil {
		log.Fatal("Cant Marshal: ", e)
	}
	jIndent := bytes.Buffer{} // the nicely formatted json
	e = json.Indent(&jIndent, b1, "", "\t")
	if e != nil {
		log.Fatal("Cant Indent: ", e)
	}
	return string(string(jIndent.Bytes()))
}

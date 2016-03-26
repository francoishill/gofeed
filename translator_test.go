package gofeed_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/atom"
	"github.com/mmcdole/gofeed/rss"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRSSTranslator_Translate(t *testing.T) {
	files, _ := filepath.Glob("testdata/translator/rss/*.xml")
	for _, f := range files {
		base := filepath.Base(f)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing %s... ", name)

		// Get actual source feed
		ff := fmt.Sprintf("testdata/translator/rss/%s.xml", name)
		f, _ := ioutil.ReadFile(ff)

		// Parse actual feed
		translator := &gofeed.DefaultRSSTranslator{}
		fp := &rss.Parser{}
		rssFeed, _ := fp.ParseFeed(string(f))
		actual := translator.Translate(rssFeed)

		// Get json encoded expected feed result
		ef := fmt.Sprintf("testdata/translator/rss/%s.json", name)
		e, _ := ioutil.ReadFile(ef)

		// Unmarshal expected feed
		expected := &gofeed.Feed{}
		json.Unmarshal(e, &expected)

		if assert.Equal(t, actual, expected, "Feed file %s.xml did not match expected output %s.json", name, name) {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("Failed\n")
		}
	}
}

func TestDefaultAtomTranslator_Translate(t *testing.T) {
	files, _ := filepath.Glob("testdata/translator/atom/*.xml")
	for _, f := range files {
		base := filepath.Base(f)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing %s... ", name)

		// Get actual source feed
		ff := fmt.Sprintf("testdata/translator/atom/%s.xml", name)
		f, _ := ioutil.ReadFile(ff)

		// Parse actual feed
		translator := &gofeed.DefaultAtomTranslator{}
		fp := &atom.Parser{}
		atomFeed, _ := fp.ParseFeed(string(f))
		actual := translator.Translate(atomFeed)

		// Get json encoded expected feed result
		ef := fmt.Sprintf("testdata/translator/atom/%s.json", name)
		e, _ := ioutil.ReadFile(ef)

		// Unmarshal expected feed
		expected := &gofeed.Feed{}
		json.Unmarshal(e, &expected)

		if assert.Equal(t, actual, expected, "Feed file %s.xml did not match expected output %s.json", name, name) {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("Failed\n")
		}
	}
}
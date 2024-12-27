package blog

import (
	"strconv"

	"github.com/cloudfstrife/go-codeforces/client"
)

func View(cli *client.Client, id int64) (*BlogEntry, error) {
	var result = BlogEntry{}
	var err = cli.Call("/blogEntry.view", map[string][]string{
		"blogEntryId": {strconv.FormatInt(id, 10)},
	}, &result)
	return &result, err
}

type BlogEntry struct {
	ID                      int64    `json:"id"`                      // .ID
	OriginalLocale          string   `json:"originalLocale"`          // .Original locale of the blog entry.
	CreationTimeSeconds     int64    `json:"creationTimeSeconds"`     // .Time, when blog entry was created, in unix format.
	AuthorHandle            string   `json:"authorHandle"`            // .Author user handle.
	Title                   string   `json:"title"`                   // .Localized.
	Content                 string   `json:"content"`                 // .Localized. Not included in short version.
	Locale                  string   `json:"locale"`                  //. Locale
	ModificationTimeSeconds int64    `json:"modificationTimeSeconds"` // .Time, when blog entry has been updated, in unix format.
	AllowViewHistory        bool     `json:"allowViewHistory"`        // .If true, you can view any specific revision of the blog entry.
	Tags                    []string `json:"tags"`                    // .tags
	Rating                  int64    `json:"rating"`                  // .rating
}

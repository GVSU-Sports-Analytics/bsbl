package bsbl

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

const BBREFPrefix string = "https://baseball-reference.com"
const BBREFStartLink string = BBREFPrefix + "/leagues/"

// GetTeams -> the main function of this package
func GetTeams(req *Request) (*Team, *Team, error) {

	year_1_link, err := getYearlink(req.Team1Yr)
	if err != nil {
		panic(err)
	}

	fmt.Println(*year_1_link)

	return nil, nil, nil
}

// getDoc -> general scraping function
func getDoc(link string) (*soup.Root, error) {
	r, err := soup.Get(link)
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(r)
	return &doc, nil
}

func getYearlink(yr string) (*string, error) {
	doc, err := getDoc(BBREFStartLink)
	if err != nil {
		return nil, err
	}

	// get all the year links from the active leagues table
	yrLinks := doc.Find(
		"table",
		"id",
		"leagues_active",
	).FindAll("a")

	var yearLink string
	for _, link := range yrLinks {
		if link.Text() != yr {
			continue
		}

		href, hasHref := link.Attrs()["href"]
		if !hasHref {
			return nil, fmt.Errorf("the requested year does not have an href attribute")
		}
		yearLink = BBREFPrefix + href
	}

	return &yearLink, nil
}

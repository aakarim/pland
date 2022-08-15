package plan

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/scanner"
	"time"
)

type PlanFile struct {
	Header            Header
	ArbitrarySections []ArbitrarySection
	Days              []Day // ordered by descending time order
	LastTouched       time.Time
}

type Header struct {
	Contents         string
	tokenStartLoc    scanner.Position
	contentsStartLoc scanner.Position
	sectionEndLoc    scanner.Position
	token            string // the token that is used to designate the header - could be empty.
}

type ArbitrarySection struct {
	Contents         string
	tokenStartLoc    scanner.Position
	contentsStartLoc scanner.Position
	sectionEndLoc    scanner.Position
	token            string
}

type Day struct {
	Date            time.Time
	Contents        string
	tokenStartLoc   scanner.Position
	contentStartLoc scanner.Position
	sectionEndLoc   scanner.Position
}

func Parse(ctx context.Context, planFile string) (*PlanFile, error) {
	p := &PlanFile{}
	rResource, err := regexp.Compile("# ?plan.(.+)?\n")
	if err != nil {
		return nil, err
	}
	fileEndLoc := len(planFile) - 1
	locs := rResource.FindAllStringIndex(planFile, -1)
	for i, loc := range locs {
		locStr := planFile[loc[0]:loc[1]]
		sectionStartLoc := loc[1] + 1
		sectionEndLoc := fileEndLoc + 1
		if i != len(locs)-1 {
			sectionEndLoc = locs[i+1][0] - 1 // character before beginning of next resource
		}
		if strings.Contains(locStr, "plan.header") {
			p.Header.Contents = strings.TrimSpace(planFile[sectionStartLoc:sectionEndLoc])
			continue
		}
		if strings.Contains(locStr, "plan.day") {
			// get date
			rDate, err := regexp.Compile("[\\d+].+-.+")
			if err != nil {
				return nil, err
			}
			dateStr := rDate.FindString(locStr)
			if dateStr == "" {
				return nil, fmt.Errorf("invalid date: %s", locStr)
			}
			t, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return nil, err
			}
			p.Days = append(p.Days, Day{
				Contents: strings.TrimSpace(planFile[sectionStartLoc:sectionEndLoc]),
				Date:     t,
			})
			continue
		}
		// arbitrary token
		p.ArbitrarySections = append(p.ArbitrarySections, ArbitrarySection{
			Contents: strings.TrimSpace(planFile[sectionStartLoc:sectionEndLoc]),
			token:    locStr,
		})
	}

	// sort the dates in descending order
	sort.Slice(p.Days, func(i, j int) bool {
		return p.Days[i].Date.After(p.Days[j].Date)
	})
	return p, nil
}

// func Parse(ctx context.Context, planFile string) (*PlanFile, error) {
// 	// lexer
// 	var s scanner.Scanner
// 	s.Init(strings.NewReader(planFile))
// 	s.Whitespace ^= 1 << '\n' // don't skip new lines

// 	p := &PlanFile{}
// 	var (
// 		lastMagicToken          *scanner.Position
// 		planActive              bool
// 		dotActive               bool
// 		headerTokenStart        *scanner.Position
// 		dayTokenStart           *scanner.Position
// 		daySectionContentsStart *scanner.Position
// 	)
// 	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
// 		txt := s.TokenText()
// 		switch txt {
// 		case ">":
// 			tmp := s.Position
// 			lastMagicToken = &tmp
// 			continue
// 		case "header":
// 			// found the header
// 			if lastMagicToken != nil && dotActive {
// 				headerTokenStart = lastMagicToken
// 				lastMagicToken = nil // reset the start check
// 				continue
// 			}
// 		}
// 		if lastMagicToken != nil {
// 			if txt == "plan" {
// 				planActive = true
// 				continue
// 			}
// 			if planActive && txt == "." {
// 				dotActive = true
// 				continue
// 			}
// 			if dotActive {
// 				switch txt {
// 				case "day":
// 					// close out the previous day
// 					if daySectionContentsStart != nil {
// 						tmp := s.Position
// 						p.Days = append(p.Days, Day{
// 							tokenStartLoc:   *dayTokenStart,
// 							contentStartLoc: *daySectionContentsStart,
// 							sectionEndLoc:   tmp,
// 						})
// 						dayTokenStart = nil
// 						daySectionContentsStart = nil
// 					}
// 					// we've found a day
// 					dayTokenStart = lastMagicToken
// 					lastMagicToken = nil // reset the start check
// 					continue
// 				}
// 			}
// 		}
// 		if dayTokenStart != nil && daySectionContentsStart == nil && tok == '\n' {
// 			tmp := s.Pos()
// 			daySectionContentsStart = &tmp
// 		}
// 	}
// 	// close out last remaining day
// 	if daySectionContentsStart != nil {
// 		p.Days = append(p.Days, Day{
// 			tokenStartLoc:   *dayTokenStart,
// 			contentStartLoc: *daySectionContentsStart,
// 			sectionEndLoc:   s.Position,
// 		})
// 		dayTokenStart = nil
// 		daySectionContentsStart = nil
// 	}

// 	spl := strings.Split(planFile, "\n")

// 	// parser
// 	for i, d := range p.Days {
// 		// get line
// 		lines := spl[d.contentStartLoc.Line-1 : d.sectionEndLoc.Line]
// 		p.Days[i].Contents = strings.TrimSpace(strings.Join(lines, "\n"))
// 	}
// 	return p, nil
// }

func (p PlanFile) String() string {
	var str string
	str += "# plan.header" + "\n\n"
	str += p.Header.Contents + "\n\n"
	for _, a := range p.ArbitrarySections {
		str += a.token + "\n\n"
		str += a.Contents + "\n\n"
	}
	for _, d := range p.Days {
		str += fmt.Sprintf("# plan.day/%s", d.Date.Format("2006-01-02")) + "\n\n"
		str += d.Contents + "\n\n"
	}
	return str
}

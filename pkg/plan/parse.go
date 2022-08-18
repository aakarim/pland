package plan

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cespare/xxhash"
)

type PlanFile struct {
	ParentVersion     int
	Header            Header
	ArbitrarySections []ArbitrarySection
	Days              []Day // ordered by descending time order
	LastTouched       time.Time
	HasConflicts      bool
}

type Header struct {
	Contents string
	// tokenStartLoc    scanner.Position
	// contentsStartLoc scanner.Position
	// sectionEndLoc    scanner.Position
	token string // the token that is used to designate the header - could be empty.
}

type ArbitrarySection struct {
	Contents string
	// tokenStartLoc    scanner.Position
	// contentsStartLoc scanner.Position
	// sectionEndLoc    scanner.Position
	token string
}

type Day struct {
	Date     time.Time
	Contents string
	// tokenStartLoc   scanner.Position
	// contentStartLoc scanner.Position
	// sectionEndLoc   scanner.Position
}

func (p *PlanFile) Digest() string {
	return fmt.Sprintf("%d", xxhash.Sum64String(p.StringExceptVersion()))
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
		locStr := strings.TrimSpace(planFile[loc[0]:loc[1]])
		sectionStartLoc := loc[1] + 1
		sectionEndLoc := fileEndLoc + 1
		if i != len(locs)-1 {
			sectionEndLoc = locs[i+1][0] - 1 // character before beginning of next resource
		}
		if strings.Contains(locStr, "plan.header") {
			headerToken := strings.TrimSpace(planFile[sectionStartLoc:sectionEndLoc])
			p.Header.Contents = headerToken
			spl := strings.Split(locStr, "/")
			// if does not exist then it may be an initial version
			if len(spl) > 1 {
				p.ParentVersion, err = strconv.Atoi(spl[1])
				if err != nil {
					return nil, fmt.Errorf("parsing header: %w", err)
				}
			}
			continue
		}
		if strings.Contains(locStr, "plan.day") {
			// get date
			rDate, err := regexp.Compile(`[\d+]+-[\d+]+-[\d+]+`)
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

// StringExceptHeader returns the string representation of the .plan file without
// the version number. Useful for creating stable digests.
func (p PlanFile) StringExceptVersion() string {
	var str string
	for _, a := range p.ArbitrarySections {
		str += a.token + "\n\n"
		str += a.Contents + "\n\n"
	}
	for i, d := range p.Days {
		str += fmt.Sprintf("# plan.day/%s", d.Date.Format("2006-01-02"))
		if i == 0 {
			str += " 🌱"
		}
		str += "\n\n"
		str += d.Contents + "\n\n"
	}
	return str
}

func (p PlanFile) String() string {
	var str string
	if p.HasConflicts {
		str += "🔀 ⚠️ Conflicts found!\n\n"
		str += "This happens in situations where another machine has\n"
		str += "uploaded your .plan file to the server which has changed the file in ways we\n"
		str += "resolve automatically.\n\n"
		str += "You can resolve this issue by going through each plan section with a 🔀 symbol\n"
		str += "and making it look how you expect it to look.\n\n"
		str += "Remember a .plan file is just a normal text file, there's no magic here. So just\n"
		str += "make it look how you'd expect it to look in the end.\n\n"
	}
	str += "# plan.header" + "/" + strconv.Itoa(p.ParentVersion) + "\n\n"
	str += p.Header.Contents + "\n\n"
	str += p.StringExceptVersion()
	return str
}

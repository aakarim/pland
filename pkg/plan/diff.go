package plan

import (
	"errors"
	"fmt"
	"strings"
)

var ErrConflict = errors.New("conflict")

func makeConflictSection(diff string) string {
	return fmt.Sprintf("🔀 conflict\n\n%s", diff)
}

func prettyDiff(str1, str2 string) string {
	finalStr := []string{}
	spl1 := strings.Split(str1, "\n")
	spl2 := strings.Split(str2, "\n")
	// make them equal sizes by adding empty strings
	lenDiff := len(spl1) - len(spl2)
	if lenDiff > 0 {
		// spl1 > spl2
		for i := lenDiff; i > 0; i-- {
			spl2 = append(spl2, "")
		}
	}
	if lenDiff < 0 {
		// spl2 > spl1
		for i := lenDiff; i < 0; i++ {
			spl1 = append(spl1, "")
		}
	}
	// diff each line
	var hasDiffs bool
	for i := 0; i < len(spl1); i++ {
		if spl1[i] != spl2[i] {
			hasDiffs = true
			finalStr = append(finalStr, "A > "+spl1[i])
			finalStr = append(finalStr, "B > "+spl2[i])
			continue
		}
		finalStr = append(finalStr, spl1[i])
	}
	if !hasDiffs {
		return ""
	}
	return strings.Join(finalStr, "\n")
}

// func prettyDiff(str1, str2 string) string {
// 	finalStr := []string{}
// 	dmp := diffmatchpatch.New()

// 	diffs := dmp.DiffMain(str1, str2, false)
// 	for _, v := range diffs {
// 		switch v.Type {
// 		case diffmatchpatch.DiffEqual:
// 			finalStr = append(finalStr, v.Text)
// 		case diffmatchpatch.DiffDelete:
// 			finalStr = append(finalStr, "B > "+v.Text)
// 		case diffmatchpatch.DiffInsert:
// 			finalStr = append(finalStr, "A > "+v.Text)
// 		}
// 	}
// 	return strings.Join(finalStr, "")
// }

// Diff diffs each known section of the .plan file and outputs the diff in the contents. It has no idea which one is 'correct'
// so it sets no version number.
//
// Relies on the days being pre-sorted by time.
func Diff(p1 *PlanFile, p2 *PlanFile) (*PlanFile, error) {
	if p1.Digest() == p2.Digest() {
		return p1, nil
	}

	var conflict bool
	diffedPlan := &PlanFile{}
	// header diff
	diffedPlan.Header.Contents = p1.Header.Contents
	if diff := prettyDiff(p1.Header.Contents, p2.Header.Contents); diff != "" {
		conflict = true

		diffedPlan.Header.Contents = makeConflictSection(diff)
	}

	// arbitrary sections find the longest and loop over that
	longerSet := p1
	otherSet := p2
	p1Length := len(p1.ArbitrarySections)
	otherSetArbLength := len(p2.ArbitrarySections)
	if p1Length < otherSetArbLength {
		longerSet = p2
		otherSet = p1
		otherSetArbLength = p1Length
	}
	for i, a := range longerSet.ArbitrarySections {
		otherSetContents := ""
		if i <= otherSetArbLength-1 {
			otherSetContents = otherSet.ArbitrarySections[i].Contents
		}
		newArbSection := ArbitrarySection{
			Contents: a.Contents,
		}
		if diff := prettyDiff(a.Contents, otherSetContents); diff != "" {
			conflict = true
			newArbSection.Contents = makeConflictSection(diff)
		}
		diffedPlan.ArbitrarySections = append(diffedPlan.ArbitrarySections, newArbSection)
	}

	// diff days by day
	// what if one set of days is longer... that can only mean it is more recently touched, so we don't have to be worried about going back
	// ASSUMPTION: days are presorted. Otherwise we would not be diffing against the correct one
	mostRecent := p1
	leastRecent := p2
	lenDiff := len(p2.Days) - len(p1.Days)
	if lenDiff > 0 {
		mostRecent = p2
		leastRecent = p1
	}
	// start from the top and add the difference in indexes so we add blank entries at the beginning
	for i := len(mostRecent.Days) - 1; i >= 0; i-- {
		mostRecentDay := mostRecent.Days[i]
		leastRecentContents := ""
		if i-lenDiff > 0 {
			leastRecentContents = leastRecent.Days[i-lenDiff].Contents
		}
		newDay := Day{
			Date:     mostRecentDay.Date,
			Contents: mostRecentDay.Contents,
		}
		if diff := prettyDiff(mostRecentDay.Contents, leastRecentContents); diff != "" {
			conflict = true
			newDay.Contents = makeConflictSection(diff)
		}
		diffedPlan.Days = append(diffedPlan.Days, newDay)
	}
	// reverse the days array
	for i, j := 0, len(diffedPlan.Days)-1; i < j; i, j = i+1, j-1 {
		diffedPlan.Days[i], diffedPlan.Days[j] = diffedPlan.Days[j], diffedPlan.Days[i]
	}
	var err error
	if conflict {
		err = ErrConflict
	}
	return diffedPlan, err
}

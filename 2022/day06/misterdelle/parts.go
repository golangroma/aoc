package main

import (
	"strconv"
)

func PartOne(input []string) string {
	//
	// nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
	// must give 10
	// nznr has duplicate and so is not ok
	//  znrn has duplicate and so is not ok
	//   nrnf has duplicate and so is not ok
	//    rnfr has duplicate and so is not ok
	//     nfrf has duplicate and so is not ok
	//      frfn has duplicate and so is not ok
	//       rfnt has no duplicate and so is OK!!!
	// last character added, t, is at position 10 of the original stream
	//

	count := 0

	//
	// Number of chars representing the signal
	//
	numCharsforSignal := 4
	charStream := input[0]

	//
	// Load the occurences for every first numCharsforSignal (4 for this run) chars
	//
	charsOccurences := make(map[byte]int)
	for i := 0; i < numCharsforSignal; i++ {
		charsOccurences[charStream[i]]++
	}

	for i := numCharsforSignal; i < len(charStream); i++ {
		if len(charsOccurences) == numCharsforSignal && mapHasUniqueValues(charsOccurences) {
			count = i
			break
		}

		//
		// Get first element of the stream
		//
		firstElement := charStream[i-numCharsforSignal]

		//
		// Get the element number of occurences
		//
		occurenceNumber := charsOccurences[firstElement]

		if occurenceNumber > 1 {
			//
			// If the element has more than one occurence let's decrement it
			//
			charsOccurences[firstElement]--
		} else {
			//
			// If the element has only one occurence let's delete it
			//
			delete(charsOccurences, firstElement)
		}

		//
		// Let's add and/or increment current element
		//
		charsOccurences[charStream[i]]++
	}

	return strconv.Itoa(count)
}

func PartTwo(input []string) string {
	//
	// nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
	// must give 29
	// nznrnfrfntjfmv has duplicate and so is not ok
	//  znrnfrfntjfmvf has duplicate and so is not ok
	//   nrnfrfntjfmvfw has duplicate and so is not ok
	//    rnfrfntjfmvfwm has duplicate and so is not ok
	//     nfrfntjfmvfwmz has duplicate and so is not ok
	//      frfntjfmvfwmzd has duplicate and so is not ok
	//       rfntjfmvfwmzdf has duplicate and so is not ok
	//        fntjfmvfwmzdfj has duplicate and so is not ok
	//         ntjfmvfwmzdfjl has duplicate and so is not ok
	//          tjfmvfwmzdfjlv has duplicate and so is not ok
	//           jfmvfwmzdfjlvt has duplicate and so is not ok
	//            fmvfwmzdfjlvtq has duplicate and so is not ok
	//             mvfwmzdfjlvtqn has duplicate and so is not ok
	//              vfwmzdfjlvtqnb has duplicate and so is not ok
	//               fwmzdfjlvtqnbh has duplicate and so is not ok
	//                wmzdfjlvtqnbhc has no duplicate and so is OK!!!
	// last character added, c, is at position 29 of the original stream
	//

	count := 0

	//
	// Number of chars representing the signal
	//
	numCharsforSignal := 14
	charStream := input[0]

	//
	// Load the occurences for every first numCharsforSignal (14 for this run) chars
	//
	charsOccurences := make(map[byte]int)
	for i := 0; i < numCharsforSignal; i++ {
		charsOccurences[charStream[i]]++
	}

	for i := numCharsforSignal; i < len(charStream); i++ {
		if len(charsOccurences) == numCharsforSignal && mapHasUniqueValues(charsOccurences) {
			count = i
			break
		}

		//
		// Get first element of the stream
		//
		firstElement := charStream[i-numCharsforSignal]

		//
		// Get the element number of occurences
		//
		occurenceNumber := charsOccurences[firstElement]

		if occurenceNumber > 1 {
			//
			// If the element has more than one occurence let's decrement it
			//
			charsOccurences[firstElement]--
		} else {
			//
			// If the element has only one occurence let's delete it
			//
			delete(charsOccurences, firstElement)
		}

		//
		// Let's add and/or increment current element
		//
		charsOccurences[charStream[i]]++
	}

	return strconv.Itoa(count)
}

func mapHasUniqueValues(m map[byte]int) bool {
	for _, occurence := range m {
		if occurence > 1 {
			return false
		}
	}

	return true
}

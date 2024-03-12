package main

// input "S" represent "Shot"
// input "R" represent "Revenge"
func bossBaby(s string) string {
	shots, revenges := 0, 0

	for _, action := range s {
		if action == 'S' {
			shots++
		} else if action == 'R' {
			// first condition
			// prevent boss baby to initiated any shooting himself.
			if shots <= 0 {
				return "Bad boy"
			}
			// second condition
			// will add revenge only when shot
			if shots > revenges {
				revenges++
			}
		}
	}

	// Boss Baby has sought revenge for every shot aimed at him at least once
	if shots > revenges {
		return "Bad boy"
	}

	return "Good boy"
}

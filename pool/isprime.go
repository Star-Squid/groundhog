package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i < nb; i++ {
		// if nb is divisible by i, then nb is not prime
		if nb%i == 0 {
			return false
		}
	}

	return true
}

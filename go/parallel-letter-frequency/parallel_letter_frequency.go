package letter

func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap, len(s))
	for _, subs := range s {
		ch <- Frequency(subs)
	}
	close(ch)
	for freqMap :=
	return m
}

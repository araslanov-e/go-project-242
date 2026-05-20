package sizefmt

import "fmt"

const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024
	tb = gb * 1024
	pb = tb * 1024
	eb = pb * 1024
)

func Format(size int64, human bool) string {
	if size < 0 {
		return "0B"
	}
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	sizes := []struct {
		unit  string
		value int64
	}{
		{"EB", eb},
		{"PB", pb},
		{"TB", tb},
		{"GB", gb},
		{"MB", mb},
		{"KB", kb},
	}

	for _, s := range sizes {
		if size >= s.value {
			return fmt.Sprintf("%.1f%s", float64(size)/float64(s.value), s.unit)
		}
	}

	return fmt.Sprintf("%dB", size)
}

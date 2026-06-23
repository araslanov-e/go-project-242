package sizefmt

import "fmt"

// Размеры в байтах для различных единиц измерения
const (
	_  = iota             // пропускаем нулевое значение
	KB = 1 << (10 * iota) // 1024 байт
	MB                    // 1 << (10 * 2) = 1 048 576 байт
	GB                    // 1 << (10 * 3) = 1 073 741 824 байт
	TB                    // 1 << (10 * 4)
	PB                    // 1 << (10 * 5)
	EB                    // 1 << (10 * 6)
)

// units - срез, содержащий единицы измерения и их соответствующие значения в байтах
var units = []struct {
	unit  string
	value int64
}{
	{"EB", EB},
	{"PB", PB},
	{"TB", TB},
	{"GB", GB},
	{"MB", MB},
	{"KB", KB},
}

// defaultUnit - единица измерения по умолчанию для неформатированного вывода
var defaultUnit = "B"

func Format(size int64, human bool) string {
	if size < 0 {
		size = 0
	}

	if human && size > 0 {
		for _, u := range units {
			if size >= u.value {
				uSize := toUnitSize(size, u.value)
				return fmt.Sprintf("%.1f%s", uSize, u.unit)
			}
		}
	}

	return fmt.Sprintf("%d%s", size, defaultUnit)
}

func toUnitSize(size, unitSize int64) float64 {
	return float64(size) / float64(unitSize)
}

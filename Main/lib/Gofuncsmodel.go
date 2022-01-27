package lib

import (
	"encoding/json"
	//"io/ioutil"
	"math"
	"math/cmplx"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type InfoAboutPolynom struct {
	PolynomialCoefficients string
	MultiplicityRoots      map[string]int
	AllRoots               string
}

type StrTable2d struct {
	PointZ   complex128
	AbsPolyZ float64
}

type Table []StrTable2d

func (a Table) Len() int           { return len(a) }
func (a Table) Less(i, j int) bool { return a[i].AbsPolyZ < a[j].AbsPolyZ }
func (a Table) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func generate(r float64) complex128 {
	// Генерация случайного комплексного числа из шара радиуса r
	rand.Seed(time.Now().UnixNano())
	z := complex(rand.Float64()*r, rand.Float64()*r)
	return z
}

func subF(ind complex128, n int) complex128 {
	if n == 0 {
		return 1
	} else {
		return ind * subF(ind-1, n-1)
	}
}

func functionOrderN(z complex128, ar []complex128, n int) complex128 {
	sum, ind := new(complex128), new(complex128)
	for int(real(*ind)) < len(ar) {
		*sum += subF(*ind, n) * ar[int(real(*ind))] * cmplx.Pow(z, *ind-complex(float64(n), 0))
		*ind += 1
	}
	return *sum
}

func metric(z complex128, ar []complex128) float64 {
	return cmplx.Abs(functionOrderN(z, ar, 0))
}

func choice(z0, fi, psi complex128, ar []complex128) (complex128, float64) {
	var n complex128 = 2 * math.Pi / fi
	listComplex := new([]complex128)
	i := new(complex128)
	psi = psi * complex(rand.Float64(), 0)
	for real(*i) < real(n) {
		z := z0 + psi*cmplx.Exp(1i*fi**i)
		*listComplex = append(*listComplex, z)
		*i += 1
	}
	TempleTable := creatTable(*listComplex, ar)
	return TempleTable[0].PointZ, TempleTable[0].AbsPolyZ
}

func creatTable(listComplex, ar []complex128) Table {
	// На вход подается список комплексных чисел и коэффициенты
	// На выходе отсортированная по возрастанию таблица по модулю полинома в точке
	TempleTable := new(Table)
	for _, value := range listComplex {
		*TempleTable = append(*TempleTable, StrTable2d{value, metric(value, ar)})
	}
	sort.Sort(Table(*TempleTable))
	return *TempleTable
}

func unique(Slice []complex128) []complex128 {
	keys := make(map[complex128]bool)
	list := new([]complex128)
	for _, entry := range Slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			*list = append(*list, entry)
		}
	}
	return *list
}

func sliceEq(a, b []complex128) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func conditionUniqueList(listComplex []complex128, psi float64) []complex128 {
	listComplexFinal := new([]complex128)
	for _, value := range listComplex {
		*listComplexFinal = append(*listComplexFinal, conEqv(value, psi))
	}
	return unique(*listComplexFinal)
}

func conEqv(z complex128, psi float64) complex128 {
	return complex(math.Ceil(real(z)/psi), math.Ceil(imag(z)/psi))
}

func conEqvList(listComplex []complex128, psi float64) [][]complex128 {
	// На вход подается массив и точность psi различения корней.
	// На выходе получается массив с массивами близких значений не
	// превышающих порядок точности psi

	tableComplexFinal := new([][]complex128)
	templeComplexList := new([]complex128)
	uniqueList := conditionUniqueList(listComplex, psi)
	for _, uniqueValue := range uniqueList {
		for _, lvalue := range listComplex {
			if cmplx.Abs(uniqueValue-conEqv(lvalue, psi)) <= 9 {
				*templeComplexList = append(*templeComplexList, lvalue)

			}
		}
		sum := 0
		for _, tableValue := range *tableComplexFinal {
			if sliceEq(tableValue, *templeComplexList) {
				sum++
			}
		}
		if sum == 0 {
			*tableComplexFinal = append(*tableComplexFinal, *templeComplexList)
		}
		*templeComplexList = nil

	}
	return *tableComplexFinal
}

func startTemperature(fz0 int) float64 {
	// Возвращает стартовую температуру для метода имитации отжига
	a := strconv.Itoa(fz0)
	return math.Pow(10, float64(len([]rune(a)))+1)
}

func SwarmParticle(z0, h complex128, n int, ar []complex128) []complex128 {
	var fi complex128 = math.Pi / complex(float64(n), 0)
	listComplexFinal := new([]complex128)
	startMetric := metric(z0, ar)
	T := startTemperature(int(startMetric))
	for T > math.Pow(10, -19) {
		T = T * 0.996
		z1, absZ1 := choice(z0, fi, h, ar)
		D := absZ1 - metric(z0, ar)
		if D < 0 {
			z0 = z1
		} else {
			p := math.Exp(-math.Abs(D) / T)
			r := rand.Float64()
			if p < r {
				z0 = z0
			} else {
				z0 = z1
			}
		}
		*listComplexFinal = append(*listComplexFinal, z0)
	}
	return *listComplexFinal
}

func goIP5(inputListZ, ar []complex128, psi float64, it int) []complex128 {
	listComplexFinal := new([]complex128)
	xi, m := new(complex128), new(complex128)
	a := len(ar)
	n := len(inputListZ)
	*m = complex(float64(a), 0) - 1
	*xi = *m - 1
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			p := inputListZ[i]
			for iter := 0; iter < it; iter++ {
				f := functionOrderN(p, ar, 0)
				if f != 0 {
					g := functionOrderN(p, ar, 1)
					h := functionOrderN(p, ar, 2)
					G := g / f
					H := h / f
					Q := cmplx.Sqrt(G*G - (*m*H)/(*m-1))
					p = p - *m/(G+*xi*Q)
					if cmplx.Abs(functionOrderN(p, ar, 0)) <= psi {
						*listComplexFinal = append(*listComplexFinal, p)
					}
				} else {
					*listComplexFinal = append(*listComplexFinal, p)
				}
			}
		}(i)
	}
	wg.Wait()

	return *listComplexFinal
}

func MoteKarlo(ar []complex128, n, iter int, r, psi float64) []complex128 {
	listComplexRand := genListComplex(n, r)
	listComplexFinal := goIP5(listComplexRand, ar, psi, iter)
	return theBestFromClass(ar, conEqvList(listComplexFinal, psi*100))
}

func genListComplex(n int, r float64) []complex128 {
	listComplexFinal := new([]complex128)
	for i := 0; i < n; i++ {
		*listComplexFinal = append(*listComplexFinal, generate(r))
	}
	return *listComplexFinal
}

func theBestFromClass(ar []complex128, listComplexClass [][]complex128) []complex128 {
	listComplexFinal := new([]complex128)
	for _, value := range listComplexClass {
		*listComplexFinal = append(*listComplexFinal, creatTable(value, ar)[0].PointZ)
	}
	return *listComplexFinal
}

func processingBeforeSP(listComplex, ar []complex128) ([]complex128, []complex128) {
	listComplexPerfect, listComplexForSP := new([]complex128), new([]complex128)
	for _, value := range listComplex {
		if metric(value, ar) <= 9e-15 {
			*listComplexPerfect = append(*listComplexPerfect, value)
		} else {
			*listComplexForSP = append(*listComplexForSP, value)
		}
	}
	return *listComplexPerfect, *listComplexForSP
}

func modelAlpha(ar []complex128, psi, r float64, amountSpoint, amountParticle, iter int) []complex128 {
	listComplexFinal := new([][]complex128)
	firstRadix := MoteKarlo(ar, amountSpoint, iter, r, psi)
	listComplexPerfect, listComplexForSP := processingBeforeSP(firstRadix, ar)
	threads := len(listComplexForSP)
	wg := new(sync.WaitGroup)
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func(i int) {
			defer wg.Done()
			applicantsOfRoots := SwarmParticle(listComplexForSP[i], 1e-15, amountParticle, ar)
			if len(applicantsOfRoots) != 0 {
				*listComplexFinal = append(*listComplexFinal, applicantsOfRoots)
			}
		}(i)

	}
	wg.Wait()

	listComplexForSP = theBestFromClass(ar, *listComplexFinal)
	for _, value := range listComplexForSP {
		listComplexPerfect = append(listComplexPerfect, value)
	}
	return FinalProcessing(theBestFromClass(ar, conEqvList(listComplexPerfect, psi*1e+3)), ar)
}

func Multiplicity(ar []complex128, root complex128, j int) int {
	if cmplx.Abs(functionOrderN(root, ar, j)) > 1e-4 {
		return j
	} else {
		return Multiplicity(ar, root, j+1)
	}

}

func MultiplicityOfRoots(ar, roots []complex128) (map[string]int, string) {
	m := make(map[string]int)
	amountRoots := new(int)
	for i := 0; i < len(roots); i++ {
		stringRoot := strconv.FormatComplex(roots[i], 'e', -1, 128)
		amountMultiplicityOfRoot := func() int {
			root := roots[i]

			return Multiplicity(ar, root, 0)
		}
		amount := amountMultiplicityOfRoot()
		if amount >= 1 {
			m[stringRoot] = amount
			*amountRoots += amount
		}

	}
	flag := new(string)
	mustAmountRoots := len(ar) - 1
	if *amountRoots == mustAmountRoots {
		*flag = strconv.FormatBool(true) + "-roots: " + strconv.Itoa(mustAmountRoots) + "|" + strconv.Itoa(mustAmountRoots)
	} else {
		*flag = strconv.FormatBool(false) + "-roots: " + strconv.Itoa(*amountRoots) + "|" + strconv.Itoa(mustAmountRoots)

	}
	return m, *flag
}

func FinalProcessing(listComplex, ar []complex128) []complex128 {
	list := new([]complex128)
	for _, value := range listComplex {
		switch {
		case cmplx.Abs(functionOrderN(value, ar, 0)) >= cmplx.Abs(functionOrderN(complex(math.Round(real(value)), 0), ar, 0)):
			*list = append(*list, complex(math.Round(real(value)), 0))
		case cmplx.Abs(functionOrderN(value, ar, 0)) >= cmplx.Abs(functionOrderN(complex(0, math.Round(imag(value))), ar, 0)):
			*list = append(*list, complex(0, math.Round(imag(value))))
		default:
			*list = append(*list, value)
		}
	}
	return unique(*list)
}

func GlobaloAlphaModel(listComplexParametrs [][]complex128, psi float64, r float64, amounrSpoint, amountParticle, iter int) []byte {
	listInfoPoly := new([]InfoAboutPolynom)
	for _, ar := range listComplexParametrs {
		roots := modelAlpha(ar, psi, r, amounrSpoint, amountParticle, iter)
		multiplicityRoots, flac := MultiplicityOfRoots(ar, roots)
		elemet := InfoAboutPolynom{convListComplexToString(ar), multiplicityRoots, flac}
		*listInfoPoly = append(*listInfoPoly, elemet)

	}

	data, _ := json.MarshalIndent(listInfoPoly, "", "  ")
	//if err := ioutil.WriteFile("./json_files/list_info_polynom.txt", data, 0600); err != nil {
	//}
	return data

}

func convListComplexToString(listComplex []complex128) string {
	finalStr := new(string)
	for _, value := range listComplex {
		convComplex := strconv.FormatComplex(value, 'e', -1, 128)
		*finalStr += convComplex + ", "
	}
	return *finalStr
}

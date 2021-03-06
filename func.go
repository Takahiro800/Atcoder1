import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func min(nums ...int) {
	min := nums[0]
	for _, n := range nums {
		if n <= min {
			min = n
		}
	}
	return min
}

func bitSearch() {
	for bits := 0; bits < (1 << uint64(n)); bits++ {
		// bitsの各要素についてのループ
		for i := 0; i < n; i++ {
			// bitsのi個目の要素の状態がonかどうかチェック
			if (bits>>uint64(i))&1 == 1 {
				// 何かしらの処理
			}
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)

func input() string {
	sc.Scan()
	return sc.Text()
}

func init()        { sc.Buffer(make([]byte, 256), 1e9); sc.Split(bufio.ScanWords) }
func Read() string { sc.Scan(); return sc.Text() }

func ReadInt() int {
	v, e := strconv.Atoi(Read())
	if e != nil {
		panic(e.Error())
	}
	return v
}

func ReadString() string {
	if !sc.Scan() {
		panic(sc.Err())
	}
	return sc.Text()
}

// ReadInstSlice returns an integer slice that han n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

func ReadFloatSlice(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat()
	}
	return b
}

func ReadFloat() float64 {
	v, e := strconv.ParseFloat(Read(), 64)
	if e != nil {
		panic(e.Error())
	}
	return v
}

func judgeInt(x float64) bool {
	if math.Floor(x) == x {
		return true
	}
	return false
}
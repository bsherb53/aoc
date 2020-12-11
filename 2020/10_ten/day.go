package ten

import (
	"fmt"
	"strconv"
	"strings"
)

var shouldLog = false
var max = 0

func Do(log bool) {
	shouldLog = log
	oj := strings.Split(data, "\n")
	ojs := make([]int, 0)
	for i := range oj {
		n, _ := strconv.Atoi(oj[i])
		ojs = append(ojs, n)
		if n > max {
			max = n
		}
	}
	partOne(ojs)
	partTwo(ojs)
}

func partOne(d []int) {
	list := make([]int, 0)
	last := 0
	oneGap, threeGap := 0, 0
	for i := 0; i <= max+3; i++ {
		if contains(d, i) {
			diff := i - last
			if diff == 1 {
				oneGap++
			}
			if diff == 3 {
				threeGap++
			}
			last = i
			list = append(list, i)
		}
	}
	diff := max + 3 - last
	if diff == 1 {
		oneGap++
	}
	if diff == 3 {
		threeGap++
	}
	if shouldLog {
		fmt.Println(fmt.Sprintf("Part One:\n\tOnes (%v) * Threes (%v) = %v", oneGap, threeGap, oneGap*threeGap))
	}
}

func contains(a []int, i int) bool {
	for _, n := range a {
		if i == n {
			return true
		}
	}
	return false
}

func partTwo(d []int) {
	var memo = make(map[int]int, 0)

	o := check(d, 1, memo)
	tw := check(d, 2, memo)
	t := check(d, 3, memo)
	if shouldLog {
		fmt.Println(fmt.Sprintf("Part Two:\n\tTotal Adapter Combinations: %v", o+tw+t))
	}
}

func check(a []int, n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	switch {
	case n == max:
		return 1
	case contains(a, n):
		e := check(a, n+1, memo)
		memo[n+1] = e

		f := check(a, n+2, memo)
		memo[n+2] = f

		g := check(a, n+3, memo)
		memo[n+3] = g
		return e + f + g
	default:
		return 0
	}
}

var example = `16
10
15
5
1
11
7
19
6
12
4`

var exampleTwo = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

var data = `35
111
135
32
150
5
106
154
41
7
27
117
109
63
64
21
138
98
40
71
144
13
66
48
12
55
119
103
54
78
65
112
39
128
53
140
77
34
28
81
151
125
85
124
2
99
131
59
60
6
94
33
42
93
14
141
92
38
104
9
29
100
52
19
147
49
74
70
84
113
120
91
97
17
45
139
90
116
149
129
87
69
20
24
148
18
58
123
76
118
130
132
75
110
105
1
8
86`

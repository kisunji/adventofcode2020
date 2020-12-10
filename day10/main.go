package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	val  int
	j1   *node
	j2   *node
	j3   *node
}

func main() {
	ss := strings.Split(input, "\n")
	var adapters []int
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		adapters = append(adapters, i)
	}
	if adapters == nil {
		panic("nil adapters")
	}

	sort.Ints(adapters)
	fmt.Println(adapters)
	maxAdapter := adapters[len(adapters)-1]

	// make slice with size max adapter + 3 (for final adaptor)
	nodes := make([]*node, maxAdapter+4)
	// outlet is node0
	nodes[0] = &node{
		val: 0,
		j1:  nodes[1],
		j2:  nodes[2],
		j3:  nodes[3],
	}

	for _, adapter := range adapters {
		node := &node{
			val: adapter,
		}
		nodes[adapter] = node
	}
	// last adapter
	last := &node{val: maxAdapter + 3}
	nodes[maxAdapter+3] = last

	for i, node := range nodes {
		if node == nil {
			continue
		}
		if i == maxAdapter+3 {
			break
		}
		node.j1 = nodes[i+1]
		node.j2 = nodes[i+2]
		node.j3 = nodes[i+3]
	}

	outlet := nodes[0]
	dfs(outlet, last)
	fmt.Println(found)
	fmt.Println(cache)
}

var found int
var cache = make(map[int]int)

func dfs(n *node, search *node) {
	if n == nil {
		return
	}
	if cache[n.val] > 0 {
		found += cache[n.val]
		return
	}

	if n == search {
		found++
		return
	}
	dfs(n.j1, search)
	dfs(n.j2, search)
	dfs(n.j3, search)
	cache[n.val] += found
}

func main_1() {
	ss := strings.Split(sample, "\n")
	var adapters []int
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		adapters = append(adapters, i)
	}
	sort.Ints(adapters)

	jolt1, jolt2, jolt3 := 0, 0, 0
	outlet := 0
	for i, adapter := range adapters {
		var diff int
		if i == 0 {
			diff = adapter - outlet
		} else {
			diff = adapter - adapters[i-1]
		}
		switch diff {
		case 1:
			jolt1++
		case 2:
			jolt2++
		case 3:
			jolt3++
		}
	}
	jolt3++ // final adapter
	fmt.Printf("jolt1: %d\njolt2: %d\njolt3: %d\n", jolt1, jolt2, jolt3)

	fmt.Printf("answer: %d", jolt1*jolt3)
}

const sample = `28
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

const input = `71
30
134
33
51
115
122
38
61
103
21
12
44
129
29
89
54
83
96
91
133
102
99
52
144
82
22
68
7
15
93
125
14
92
1
146
67
132
114
59
72
107
34
119
136
60
20
53
8
46
55
26
126
77
65
78
13
108
142
27
75
110
90
35
143
86
116
79
48
113
101
2
123
58
19
76
16
66
135
64
28
9
6
100
124
47
109
23
139
145
5
45
106
41`

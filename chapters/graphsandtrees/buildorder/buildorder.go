package buildorder

import "fmt"

/*
Build Order: You are given a list of projects and a list of dependencies (which is a list of pairs of
projects, where the second project is dependent on the first project). All of a project's dependencies
must be built before the project is. Find a build order that will allow the projects to be built. If there
is no valid build order, return an error.
EXAMPLE
Input:
projects: a, b, c, d, e, f
dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
Output: f, e, a, b, d, c
*/

type Queue struct {
	data []string
}

func (q *Queue) Enqueue(element string) {
	q.data = append(q.data, element)
}

func (q *Queue) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func KahnAlgorithm(projects []string, dependencies [][2]string) ([]string, error) {
	graph := make(map[string][]string)
	inDegree := make(map[string]int)
	order := []string{}
	queue := &Queue{
		data: []string{},
	}

	for _, p := range projects {
		graph[p] = []string{}
		inDegree[p] = 0
	}

	for i := range dependencies {
		prereq := dependencies[i][0]
		dependent := dependencies[i][1]

		if _, ok := inDegree[prereq]; !ok {
			return nil, fmt.Errorf("unknown project %q in dependencies", prereq)
		}

		if _, ok := inDegree[dependent]; !ok {
			return nil, fmt.Errorf("unknown project %q in dependencies", dependent)
		}

		graph[prereq] = append(graph[prereq], dependent)
		inDegree[dependent]++
	}

	for p, deg := range inDegree {
		if deg == 0 {
			queue.Enqueue(p)
		}
	}

	for !queue.IsEmpty() {
		pr := queue.Dequeue()
		order = append(order, pr)
		for _, elem := range graph[pr] {
			inDegree[elem]--
			if inDegree[elem] == 0 {
				queue.Enqueue(elem)
			}
		}
	}

	if len(order) == len(projects) {
		return order, nil
	}
	return nil, fmt.Errorf("cycle detected")
}

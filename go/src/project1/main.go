// This file has been changed
package main

import (
	"fmt"
	"home/cepl/go/src/project1/utility"
	"home/cepl/go/src/project1/utility/packagestruct"
	"sort"
)

var employee = map[string]int{"Mark": 10, "Sandy": 20}

func main() {
	fmt.Println(employee)
	utility.Empmap()
	//utility.Makemap()
	//gpackagestruct.Rectstruct()
	//packagestruct.EmployeeSal()
	packagestruct.Empstruct()
	fmt.Println("nikitagi")
	//struct2.Employeestruct()g

	Employee := make(map[string]int)
	Employee["Gloria"] = 99
	Employee["Darshita"] = 899

	utility.Makemap(Employee)
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	//fmt.Println(unSortedMap)
	keys := make([]string, 0, len(unSortedMap))
	//fmt.Println("map keys", keys)

	for k := range unSortedMap {
		fmt.Println("MAP")
		fmt.Println(unSortedMap)
		fmt.Println(k)
		keys = append(keys, k)
		fmt.Println(keys)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k)
		fmt.Println(k, unSortedMap[k])
	}

}

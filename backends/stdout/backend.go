// Copyright (c) 2014 Eric Paris. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package stdout

import (
	"fmt"
)

type StdoutRouterManager struct {
	lastRouteTable map[string]string
}

func New() (*StdoutRouterManager, error) {
	rm := &StdoutRouterManager{
		lastRouteTable:	make(map[string]string),
	}
	return rm, nil
}

func (rm *StdoutRouterManager) Sync(routeTable map[string]string) error {
	var added map[string]string = make(map[string]string);
	var removed map[string]string = make(map[string]string);

	last := rm.lastRouteTable;
	rm.lastRouteTable = make(map[string]string)

	for ip, subnet := range routeTable {
		rm.lastRouteTable[ip] = subnet
		newsubnet, ok := last[ip]
		if ! ok || newsubnet != subnet {
			added[ip]=subnet
			continue;
		}
	}
	for ip, oldsubnet := range last {
		subnet, ok := routeTable[ip]
		if ! ok || oldsubnet != subnet {
			removed[ip]=oldsubnet
			continue;
		}
	}
	if len(added) == 0 && len(removed) == 0 {
		return nil
	}

	fmt.Printf("*****\n")
	for ip, subnet := range added {
		fmt.Printf("ADD %v via %v\n", subnet, ip)
	}
	for ip, subnet := range removed {
		fmt.Printf("REMOVE %v via %v\n", subnet, ip)
	}
	fmt.Printf("!!!!!\n")
	return nil
}

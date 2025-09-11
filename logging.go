package main

/*
    yrHost - a multi-service home server
    Copyright (C) 2025  PabloMyLove

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import "fmt"

const (
	RED   string = "\033[31m"
	BLUE  string = "\033[34m"
	CYAN  string = "\033[36m"
	PINK  string = "\033[35m"
	GREEN string = "\033[32m"
	RESET string = "\033[0m"
)
const (
	ERROR    int = 0
	ATTEMPT  int = 1
	COMPLETE int = 2
	STEP     int = 3
)

func log(logType int, content string, addSeparator bool) bool {
	switch logType {
	case ERROR:
		fmt.Printf("%s>> %s%s\n", RED, RESET, content)
	case ATTEMPT:
		fmt.Printf("%s>> %s%s\n", BLUE, RESET, content)
	case COMPLETE:
		fmt.Printf("%s>> %s%s\n", GREEN, RESET, content)
	case STEP:
		fmt.Printf("	%s•%s %s\n", CYAN, RESET, content)
	default:
		return false
	}

	if addSeparator {
		fmt.Printf("%s-------------------------------------------------%s\n", PINK, RESET)
	}
	return true
}


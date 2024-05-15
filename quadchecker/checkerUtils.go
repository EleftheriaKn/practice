package quadchecker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func PrintResult(respArr []string) {

	fmt.Println("\nResult : ")
	count := 0
	for _, match := range respArr {
		if match != "" {
			count++
		}
	}

	if count < 1 {
		println("Not a quad function")
	} else {
		var result []string
		for _, match := range respArr {
			if match != "" {
				result = append(result, match)
			}
		}
		var response string
		for i, s := range result {
			if i > 0 && i < len(result) {
				response += " || "
			}
			response += s
		}
		fmt.Println(response)
	}

	println()

}

func GetShape(r io.Reader) []string {

	scanner := bufio.NewScanner(r)

	var lines []string

	fmt.Println(scanner.Text())
	fmt.Println("Received : ")

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		lines = append(lines, scanner.Text())

		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}

	if len(lines) < 1 {
		return nil
	}

	return lines

}

// return matching quads
func Check(lines []string) []string {
	var quads []string

	if len(lines) > 0 {
		quads = append(quads, checkQuadA(lines))
		quads = append(quads, checkQuadB(lines))
		quads = append(quads, checkQuadC(lines))
		quads = append(quads, checkQuadD(lines))
		quads = append(quads, checkQuadE(lines))
	}
	return quads
}

func checkQuadE(lines []string) string {
	resp := ""
	quadBarray := QuadE(len(lines[0]), len(lines))

	flag := true

	for i, s := range lines {
		if s != quadBarray[i] {
			flag = false
		}
	}

	if flag {
		resp = "[QuadE] [" + strconv.Itoa(len(lines[0])) + "] [" + strconv.Itoa(len(lines)) + "]"
		return resp
	}

	return ""
}

func checkQuadD(lines []string) string {
	resp := ""
	quadBarray := QuadD(len(lines[0]), len(lines))

	flag := true

	for i, s := range lines {
		if s != quadBarray[i] {
			flag = false
		}
	}

	if flag {
		resp = "[QuadD] [" + strconv.Itoa(len(lines[0])) + "] [" + strconv.Itoa(len(lines)) + "]"
		return resp
	}

	return ""
}

func checkQuadC(lines []string) string {
	resp := ""
	quadBarray := QuadC(len(lines[0]), len(lines))

	flag := true

	for i, s := range lines {
		if s != quadBarray[i] {
			flag = false
		}
	}

	if flag {
		resp = "[QuadC] [" + strconv.Itoa(len(lines[0])) + "] [" + strconv.Itoa(len(lines)) + "]"
		return resp
	}

	return ""
}

func checkQuadB(lines []string) string {
	resp := ""
	quadBarray := QuadB(len(lines[0]), len(lines))

	flag := true

	for i, s := range lines {
		if s != quadBarray[i] {
			flag = false
		}
	}

	if flag {
		resp = "[QuadB] [" + strconv.Itoa(len(lines[0])) + "] [" + strconv.Itoa(len(lines)) + "]"
		return resp
	}

	return ""
}

func checkQuadA(lines []string) string {
	resp := ""

	quadAarray := QuadA(len(lines[0]), len(lines))

	flag := true

	for i, s := range lines {
		if s != quadAarray[i] {
			flag = false
		}
	}

	if flag {
		resp = "[QuadA] [" + strconv.Itoa(len(lines[0])) + "] [" + strconv.Itoa(len(lines)) + "]"
		return resp
	}

	return ""
}

func CheckForExecutables() {
	filenames := []string{"quadA", "quadB", "quadC", "quadD", "quadE", "quadchecker"}
	buildPaths := []string{"../quadAmain/quada.go", "../quadBmain/quadb.go", "../quadCmain/quadc.go", "../quadDmain/quadd.go", "../quadEmain/quade.go", "../quadCheckerMain/quadchecker.go"}

	creatbin(filenames, buildPaths)
}

func creatbin(filenames []string, paths []string) {
	flag := false
	for i := 0; i < len(filenames); i++ {
		// Check if the file exists
		_, err := os.Stat(filenames[i])
		if err == nil {
			// fmt.Printf("File '%s' exists in the folder.\n", filenames[i])
		} else if os.IsNotExist(err) { // if not exist build the file
			fmt.Println("================== ", filenames[i], "====================")

			fmt.Printf("File '%s' does not exist in the folder.\n", filenames[i])
			cmd := exec.Command("go", "build", "-o", filenames[i], paths[i])
			_, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("file : ", filenames[i], " created ")
			flag = true
		} else {
			fmt.Println("Error:", err)
		}
	}

	if flag {
		fmt.Println("\n\n please try again .. ")
	}
}

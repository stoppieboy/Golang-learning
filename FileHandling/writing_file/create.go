package main
import (
	"os"
	"flag"
	"bufio"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fileName := flag.String("file","file.txt","file name for creating a new file")
	flag.Parse()
	file, err := os.Create(*fileName)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println("Enter the text:")
	text, _ := reader.ReadString('\n')

	file.WriteString(text)
}
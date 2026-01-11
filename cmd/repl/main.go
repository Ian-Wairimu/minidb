package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"minidb/internal/engine"
)

func main() {
	eng, err := engine.NewEngine("data.wal")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("MiniDB REPL (type 'exit' to quit)")

	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')

		if line == "exit\n" {
			break
		}

		res, err := eng.Execute(context.Background(), line)
		if err != nil {
			fmt.Println("ERR:", err)
			continue
		}

		fmt.Println(res)
	}
}

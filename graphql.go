package main

import (
	"io"
	"os"
	"log"
	"fmt"
	"os/exec"
	"encoding/json"
	"gopkg.in/urfave/cli.v1"
)

func checkError(err error) {
    if err != nil {
	log.Fatalf("Error: %s", err)
    }
}

func cleanup() {
	fmt.Println("cleanup")
//	hasura.Process.Kill()
}

//var hasura exec.Cmd;



type Data struct {
    output []byte
    error  error
}


// func runCommand(port string, ch chan<- Data) {
//	hasura := exec.Command(
//		"graphql-engine", "serve",
//		"--server-port", port,
//		"--database-url", "postgres://nas@localhost/spendily_dev",
//	)
//	var outb, errb bytes.Buffer
//	hasura.Stdout = &outb
//	hasura.Stderr = &errb

//	err := hasura.Start()
//	if err != nil {
//		log.Fatal(err)
//	}


//	ch <- Data{
//		error:  err,
//		output: data,
//	}
// }




func startGraphql(c *cli.Context) *exec.Cmd {
	jwtSecret, _ := json.Marshal(
		map[string]string{
			"type": "HS256",
			"key": c.String("session_secret"),
			"claims_namespace": "graphql_claims",
		},
	)

	hasura := exec.Command(
		"graphql-engine", "serve",
		"--server-port", fmt.Sprintf("%d", c.Int("graphql_port")),
		"--database-url", "postgres://nas@localhost/spendily_dev",
	)

	hasura.Env =  append(os.Environ(),
		fmt.Sprintf("HASURA_GRAPHQL_ACCESS_KEY=%s", c.String("graphql_access_key")),
		fmt.Sprintf("HASURA_GRAPHQL_JWT_SECRET=%s", jwtSecret),
	)

	fmt.Printf("STARTED GQL: %s\n", c.String("graphql_access_key"))
	// Create stdout, stderr streams of type io.Reader
	stdout, err := hasura.StdoutPipe()
	checkError(err)
	stderr, err := hasura.StderrPipe()
	checkError(err)

	// Start command
	err = hasura.Start()
	checkError(err)

	// Don't let main() exit before our command has finished running
	//defer hasura.Wait()  // Doesn't block

	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	return hasura;
}

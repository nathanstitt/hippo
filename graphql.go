package hippo

import (
	"io"
	"os"
	"fmt"
	"os/exec"
	"encoding/json"
	"gopkg.in/urfave/cli.v1"
)


func StartGraphql(c *cli.Context) *exec.Cmd {
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
		"--database-url", c.String("db_conn_url"),
	)

	hasura.Env =  append(os.Environ(),
		fmt.Sprintf("HASURA_GRAPHQL_ACCESS_KEY=%s", c.String("graphql_access_key")),
		fmt.Sprintf("HASURA_GRAPHQL_JWT_SECRET=%s", jwtSecret),
	)

	fmt.Printf("STARTED GQL: %s\n", c.String("graphql_access_key"))
	// Create stdout, stderr streams of type io.Reader
	stdout, err := hasura.StdoutPipe()
	CheckError(err)
	stderr, err := hasura.StderrPipe()
	CheckError(err)

	// Start command
	err = hasura.Start()
	CheckError(err)

	// Don't let main() exit before our command has finished running
	//defer hasura.Wait()  // Doesn't block

	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	return hasura;
}

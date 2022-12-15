package pact

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"go.mod/server"
	"testing"
)

func TestProvider(t *testing.T) {
	// It takes a suitable port and run the server
	port, _ := utils.GetFreePort()

	//New server
	svr := server.NewServer()

	//The server run on the given port
	go svr.StartServer(port)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "TodoService",
		Consumer:                 "TodoInterface",
		DisableToolValidityCheck: true,
	}

	// Verified request
	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs: []string{
			"D:\\BURCU\\Code\\Todo App\\frontend\\pact\\pacts\\userinterface-todointerface.json",
		},
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}

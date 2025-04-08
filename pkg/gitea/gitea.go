package gitea

import (
	"crypto/tls"
	"net/http"
	"sync"

	"gitea.com/gitea/gitea-mcp/pkg/flag"
	"gitea.com/gitea/gitea-mcp/pkg/log"

	"code.gitea.io/sdk/gitea"
)

var (
	client     *gitea.Client
	clientOnce sync.Once
)

func Client() *gitea.Client {
	clientOnce.Do(func() {
		var err error
		if client != nil {
			return
		}
		opts := []gitea.ClientOption{
			gitea.SetToken(flag.Token),
		}
		if flag.Insecure {
			httpClient := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
			opts = append(opts, gitea.SetHTTPClient(httpClient))
		}
		client, err = gitea.NewClient(flag.Host, opts...)
		if err != nil {
			log.Fatalf("create gitea client err: %v", err)
		}
	})
	return client
}

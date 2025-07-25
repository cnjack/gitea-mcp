package repo

import (
	"context"
	"fmt"

	"gitea.com/gitea/gitea-mcp/pkg/gitea"
	"gitea.com/gitea/gitea-mcp/pkg/log"
	"gitea.com/gitea/gitea-mcp/pkg/to"

	gitea_sdk "code.gitea.io/sdk/gitea"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	ListRepoCommitsToolName = "list_repo_commits"
)

var ListRepoCommitsTool = mcp.NewTool(
	ListRepoCommitsToolName,
	mcp.WithDescription("List repository commits"),
	mcp.WithString("owner", mcp.Required(), mcp.Description("repository owner")),
	mcp.WithString("repo", mcp.Required(), mcp.Description("repository name")),
	mcp.WithString("sha", mcp.Description("SHA or branch to start listing commits from")),
	mcp.WithString("path", mcp.Description("path indicates that only commits that include the path's file/dir should be returned.")),
	mcp.WithNumber("page", mcp.Required(), mcp.Description("page number"), mcp.DefaultNumber(1), mcp.Min(1)),
	mcp.WithNumber("page_size", mcp.Required(), mcp.Description("page size"), mcp.DefaultNumber(50), mcp.Min(1)),
)

func init() {
	Tool.RegisterRead(server.ServerTool{
		Tool:    ListRepoCommitsTool,
		Handler: ListRepoCommitsFn,
	})
}

func ListRepoCommitsFn(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Debugf("Called ListRepoCommitsFn")
	owner, ok := req.GetArguments()["owner"].(string)
	if !ok {
		return to.ErrorResult(fmt.Errorf("owner is required"))
	}
	repo, ok := req.GetArguments()["repo"].(string)
	if !ok {
		return to.ErrorResult(fmt.Errorf("repo is required"))
	}
	page, ok := req.GetArguments()["page"].(float64)
	if !ok {
		return to.ErrorResult(fmt.Errorf("page is required"))
	}
	pageSize, ok := req.GetArguments()["page_size"].(float64)
	if !ok {
		return to.ErrorResult(fmt.Errorf("page_size is required"))
	}
	sha, _ := req.GetArguments()["sha"].(string)
	path, _ := req.GetArguments()["path"].(string)
	opt := gitea_sdk.ListCommitOptions{
		ListOptions: gitea_sdk.ListOptions{
			Page:     int(page),
			PageSize: int(pageSize),
		},
		SHA:  sha,
		Path: path,
	}
	commits, _, err := gitea.Client().ListRepoCommits(owner, repo, opt)
	if err != nil {
		return to.ErrorResult(fmt.Errorf("list repo commits err: %v", err))
	}
	return to.TextResult(commits)
}

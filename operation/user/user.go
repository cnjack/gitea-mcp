package user

import (
	"context"
	"fmt"

	"gitea.com/gitea/gitea-mcp/pkg/gitea"
	"gitea.com/gitea/gitea-mcp/pkg/log"
	"gitea.com/gitea/gitea-mcp/pkg/to"
	"gitea.com/gitea/gitea-mcp/pkg/tool"

	gitea_sdk "code.gitea.io/sdk/gitea"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	GetMyUserInfoToolName = "get_my_user_info"
	GetUserOrgsToolName   = "get_user_orgs"
)

var Tool = tool.New()

var (
	GetMyUserInfoTool = mcp.NewTool(
		GetMyUserInfoToolName,
		mcp.WithDescription("Get my user info"),
	)

	GetUserOrgsTool = mcp.NewTool(
		GetUserOrgsToolName,
		mcp.WithDescription("Get organizations associated with the authenticated user"),
		mcp.WithNumber("page", mcp.Description("page number"), mcp.DefaultNumber(1)),
		mcp.WithNumber("pageSize", mcp.Description("page size"), mcp.DefaultNumber(100)),
	)
)

func init() {
	Tool.RegisterRead(server.ServerTool{
		Tool:    GetMyUserInfoTool,
		Handler: GetUserInfoFn,
	})

	Tool.RegisterRead(server.ServerTool{
		Tool:    GetUserOrgsTool,
		Handler: GetUserOrgsFn,
	})
}

func GetUserInfoFn(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Debugf("Called GetUserInfoFn")
	user, _, err := gitea.Client().GetMyUserInfo()
	if err != nil {
		return to.ErrorResult(fmt.Errorf("get user info err: %v", err))
	}

	return to.TextResult(user)
}

func GetUserOrgsFn(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Debugf("Called GetUserOrgsFn")
	page, ok := req.GetArguments()["page"].(float64)
	if !ok || page < 1 {
		page = 1
	}
	pageSize, ok := req.GetArguments()["pageSize"].(float64)
	if !ok || pageSize < 1 {
		pageSize = 100
	}
	opt := gitea_sdk.ListOrgsOptions{
		ListOptions: gitea_sdk.ListOptions{
			Page:     int(page),
			PageSize: int(pageSize),
		},
	}
	orgs, _, err := gitea.Client().ListMyOrgs(opt)
	if err != nil {
		return to.ErrorResult(fmt.Errorf("get user orgs err: %v", err))
	}

	return to.TextResult(orgs)
}

# Gitea MCP 伺服器

[English](README.md) | [简体中文](README.zh-cn.md)

**Gitea MCP 伺服器** 是一個整合插件，旨在將 Gitea 與 Model Context Protocol (MCP) 系統連接起來。這允許通過 MCP 兼容的聊天界面無縫執行命令和管理倉庫。

[![在 VS Code 中使用 Docker 安裝](https://img.shields.io/badge/VS_Code-Install_Server-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=gitea&inputs=[{%22id%22:%22gitea_token%22,%22type%22:%22promptString%22,%22description%22:%22Gitea%20Personal%20Access%20Token%22,%22password%22:true}]&config={%22command%22:%22docker%22,%22args%22:[%22run%22,%22-i%22,%22--rm%22,%22-e%22,%22GITEA_ACCESS_TOKEN%22,%22docker.gitea.com/gitea-mcp-server%22],%22env%22:{%22GITEA_ACCESS_TOKEN%22:%22${input:gitea_token}%22}}) [![在 VS Code Insiders 中使用 Docker 安裝](https://img.shields.io/badge/VS_Code_Insiders-Install_Server-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=gitea&inputs=[{%22id%22:%22gitea_token%22,%22type%22:%22promptString%22,%22description%22:%22Gitea%20Personal%20Access%20Token%22,%22password%22:true}]&config={%22command%22:%22docker%22,%22args%22:[%22run%22,%22-i%22,%22--rm%22,%22-e%22,%22GITEA_ACCESS_TOKEN%22,%22docker.gitea.com/gitea-mcp-server%22],%22env%22:{%22GITEA_ACCESS_TOKEN%22:%22${input:gitea_token}%22}}&quality=insiders)

## 什麼是 Gitea？

Gitea 是一個由社群管理的輕量級代碼託管解決方案，使用 Go 語言編寫。它以 MIT 許可證發布。Gitea 提供 Git 託管，包括倉庫查看器、問題追蹤、拉取請求等功能。

## 什麼是 MCP？

Model Context Protocol (MCP) 是一種協議，允許通過聊天界面整合各種工具和系統。它能夠無縫執行命令和管理倉庫、用戶和其他資源。

## 🚧 安裝

### 在 VS Code 中使用

要快速安裝，請使用本 README 頂部的單擊安裝按鈕之一。

要手動安裝，請將以下 JSON 塊添加到 VS Code 的用戶設置 (JSON) 文件中。您可以通過按 `Ctrl + Shift + P` 並輸入 `Preferences: Open User Settings (JSON)` 來完成此操作。

或者，您可以將其添加到工作區中的 `.vscode/mcp.json` 文件中。這將允許您與他人共享配置。

> 請注意，`.vscode/mcp.json` 文件中不需要 `mcp` 鍵。

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "gitea_token",
        "description": "Gitea 個人訪問令牌",
        "password": true
      }
    ],
    "servers": {
      "github": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "-e",
          "GITEA_ACCESS_TOKEN",
          "docker.gitea.com/gitea-mcp-server"
        ],
        "env": {
          "GITEA_ACCESS_TOKEN": "${input:gitea_token}"
        }
      }
    }
  }
}
```

### 📥 下載官方二進制版本

您可以從[這裡](https://gitea.com/gitea/gitea-mcp/releases)下載官方版本。

### 🔧 從源代碼構建

您可以使用 Git 克隆倉庫來下載源代碼：

```bash
git clone https://gitea.com/gitea/gitea-mcp.git
```

在構建之前，請確保您已安裝以下內容：

- make
- Golang (建議使用 Go 1.24 或更高版本)

然後運行：

```bash
make install
```

### 📁 添加到 PATH

安裝後，將二進制文件 gitea-mcp 複製到系統 PATH 中包含的目錄。例如：

```bash
cp gitea-mcp /usr/local/bin/
```

## 🚀 使用

此示例適用於 Cursor，您也可以在 VSCode 中使用插件。
要配置 Gitea 的 MCP 伺服器，請將以下內容添加到您的 MCP 配置文件中：

- **stdio 模式**

```json
{
  "mcpServers": {
    "gitea": {
      "command": "gitea-mcp",
      "args": [
        "-t",
        "stdio",
        "--host",
        "https://gitea.com"
        // "--token", "<your personal access token>"
      ],
      "env": {
        // "GITEA_HOST": "https://gitea.com",
        // "GITEA_INSECURE": "true",
        "GITEA_ACCESS_TOKEN": "<your personal access token>"
      }
    }
  }
}
```

- **sse 模式**

```json
{
  "mcpServers": {
    "gitea": {
      "url": "http://localhost:8080/sse"
    }
  }
}
```

**默認日誌路徑**: `$HOME/.gitea-mcp/gitea-mcp.log`

> [!注意]
> 您可以通過命令行參數或環境變量提供您的 Gitea 主機和訪問令牌。
> 命令行參數具有最高優先級

一切設置完成後，請嘗試在您的 MCP 兼容聊天框中輸入以下內容：

```text
列出我所有的倉庫
```

## ✅ 可用工具

Gitea MCP 伺服器支持以下工具：

|             工具             |   範圍   |             描述             |
| :--------------------------: | :------: | :--------------------------: |
|       get_my_user_info       |   用戶   |     獲取已認證用戶的信息     |
|         create_repo          |   倉庫   |        創建一個新倉庫        |
|          fork_repo           |   倉庫   |         復刻一個倉庫         |
|        list_my_repos         |   倉庫   | 列出已認證用戶擁有的所有倉庫 |
|        create_branch         |   分支   |        創建一個新分支        |
|        delete_branch         |   分支   |         刪除一個分支         |
|        list_branches         |   分支   |     列出倉庫中的所有分支     |
|      list_repo_commits       |   提交   |     列出倉庫中的所有提交     |
|       get_file_content       |   文件   |    獲取文件的內容和元數據    |
|         create_file          |   文件   |        創建一個新文件        |
|         update_file          |   文件   |         更新現有文件         |
|         delete_file          |   文件   |         刪除一個文件         |
|      get_issue_by_index      |   問題   |       根據索引獲取問題       |
|       list_repo_issues       |   問題   |     列出倉庫中的所有問題     |
|         create_issue         |   問題   |        創建一個新問題        |
|     create_issue_comment     |   問題   |       在問題上創建評論       |
|  get_pull_request_by_index   | 拉取請求 |     根據索引獲取拉取請求     |
|   list_repo_pull_requests    | 拉取請求 |   列出倉庫中的所有拉取請求   |
|     create_pull_request      | 拉取請求 |      創建一個新拉取請求      |
|         search_users         |   用戶   |           搜索用戶           |
|       search_org_teams       |   組織   |       搜索組織中的團隊       |
|         search_repos         |   倉庫   |           搜索倉庫           |
| get_gitea_mcp_server_version |  伺服器  | 獲取 Gitea MCP 伺服器的版本  |

## 🐛 調試

要啟用調試模式，請在使用 sse 模式運行 Gitea MCP 伺服器時添加 `-d` 標誌：

```sh
./gitea-mcp -t sse [--port 8080] --token <your personal access token> -d
```

## 🛠 疑難排解

如果您遇到任何問題，以下是一些常見的疑難排解步驟：

1. **檢查您的 PATH**: 確保 `gitea-mcp` 二進制文件位於系統 PATH 中包含的目錄中。
2. **驗證依賴項**: 確保您已安裝所有所需的依賴項，例如 `make` 和 `Golang`。
3. **檢查配置**: 仔細檢查您的 MCP 配置文件是否有任何錯誤或遺漏的信息。
4. **查看日誌**: 檢查日誌中是否有任何錯誤消息或警告，可以提供有關問題的更多信息。

享受通過聊天探索和管理您的 Gitea 倉庫的樂趣！

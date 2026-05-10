# Session 2 Summary — 2026-05-10

## 完成事項

### 1. 清理舊目錄
- 使用者確認後，刪除本機兩個殘留目錄：
  - `~/workspace/GitHubActionsInAction-Part1/`
  - `~/workspace/Globoticket/`
- 兩個目錄原為 fork 而來（非 chenghyang2001 帳號下的 repo），GitHub 端無需操作

### 2. 找到並 Clone《GitHub Actions in Action》（Manning）書本 sample code
- 書名：*GitHub Actions in Action*，作者：Michael Kaufmann、Rob Bos、Marcel de Vries
- 出版社：Manning（非 O'Reilly）
- 主要 repo：`vriesmarcel/github-actions-in-action`（co-author Marcel de Vries 維護）
- Clone 路徑：`~/workspace/github-actions-in-action/`
- 內容：Globoticket 三微服務（catalog / frontend / ordering）+ .github/workflows/

### 3. 發現並 Clone GitHubActionsInAction 組織 repo
- 組織：`https://github.com/GitHubActionsInAction/`
- Clone 了兩個 repo：
  - `GitHubActionsInAction/Part1` → `~/workspace/GitHubActionsInAction-Part1/`（入門概念 workflow）
  - `GitHubActionsInAction/ActionInAction` → `~/workspace/GitHubActionsInAction-ActionInAction/`（Docker container action 動手 lab）

### 4. 整理全書 Workflow 清單（26 個）
- **Part1 repo（4 個）**：MyFirstWorkflow、Main、Workflow in Branch、Labeler
- **Globoticket repo（21 個）**：Chapter 8（compile/test/container/SBOM/security）、Chapter 9（deploy/release/reusable workflow/blue-green）、CodeQL
- **ActionInAction repo（1 個，template）**：讀者自建 Docker container action 的 hands-on lab

## 關鍵技術筆記

### Manning 書本 sample code 結構特點
- 不像 O'Reilly 書按 chapter-X 資料夾分類，Manning 用同一個 Globoticket 應用搭配不同 workflow 示範各章概念
- workflow 名稱內含章節號（如 `"chapter 08: ..."`, `"chapter 09: ..."`）
- Part1 repo 另有 branch 對應不同功能（`matrix`、`new-workflow`、`change-appjs`）

### GitHubActionsInAction 組織 repo 清單
| Repo | 用途 |
|---|---|
| Part1 | 入門 workflow（triggers / context / labeler / masking） |
| Globoticket | 完整 app + Chapter 8-9+ 進階 workflow |
| ActionInAction | Chapter 4 Docker container action template |

### fork repo 處理原則
- Fork 來的 repo 刪本機即可，GitHub 上原始 repo 不是自己的帳號，無需（也不應）刪除

## 產出檔案

| 類型 | 路徑 | 說明 |
|---|---|---|
| Clone | `~/workspace/github-actions-in-action/` | Manning 書主 sample repo |
| Clone | `~/workspace/GitHubActionsInAction-Part1/` | 入門 workflow 範例 |
| Clone | `~/workspace/GitHubActionsInAction-ActionInAction/` | Docker action hands-on lab |
| Summary | `summary-02-sessions/2026-05-10/session2-summary.md` | 本檔 |

---

## HANDOFF（下次 session 優先處理）

### 立即行動
- [ ] 開始跟著書本 Part1 章節練習：從 `GitHubActionsInAction-Part1` 的 `MyFirstWorkflow.yml` 開始，push 到自己的 fork 觸發 workflow
- [ ] 若需要跟著 Chapter 8-9 練習，考慮 fork `vriesmarcel/github-actions-in-action` 到 `chenghyang2001` 帳號並設定 Azure secrets
- [ ] 確認 `ActionInAction` hands-on lab：依 README 步驟建立自己的 `MyActionInAction` repo

### 進行中（需接續）
- 目前已完成三個 repo 的下載與 workflow 清單整理，尚未實際執行任何 workflow
- 書本章節對應關係已確認（Part1 = 入門、Ch8 = Build/Test、Ch9 = Deploy/Release、Ch4 = 自訂 Action）

### 注意事項
- `GitHubActionsInAction/Globoticket` 與 `vriesmarcel/github-actions-in-action` 內容高度重疊（後者是前者的 co-author fork），兩者保留一個即可
- Chapter 9 的 deploy workflow 需要 Azure 帳號 + ACR + AKS，無 Azure 環境無法直接跑
- `create-release.yml` 觸發是 `registry_package: published`，須先有 container push 才會觸發

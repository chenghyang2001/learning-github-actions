# Session 1 Summary

**日期**：2026-05-10
**專案**：learning-github-actions（O'Reilly 書本範例）

---

## 完成事項

### 環境建置
- 安裝 Go（`winget install GoLang.Go`，版本 1.26.3）至 Windows 10 本機
- 驗證 `go run helloworld.go` 本機執行成功

### helloworld.go 修改
- 新增列印當前日期時間（使用 `time.DateTime` 常數，非手寫格式字串）
- 嵌入 `time/tzdata`，強制使用 Asia/Taipei 時區，解決 GitHub Actions Linux runner 無 tzdata 問題
- fallback：載入時區失敗改用 UTC 並印警告

### CLAUDE.md 建立
- 建立專案根目錄 CLAUDE.md，包含：架構說明、章節目錄結構、關鍵設計決策
- 說明章節 YAML 不放 `.github/workflows/` 的原因（避免每次 push 觸發所有示範 workflow）

### Workflow 上傳（`.github/workflows/`）
共將 **22 個 workflow** 從書本章節目錄複製至 `.github/workflows/` 並推上 GitHub：

| 章節 | 數量 | 說明 |
|------|------|------|
| Chapter 4 | 2 | basic.yml、simple-proj.yml（基本 workflow 結構）|
| Chapter 6 | 2 | system-environment-settings.yml、verify-file.yml |
| Chapter 7 | 3 | output-from-a-job.yml、output-from-a-step.yml、create-failure-issue.yml |
| Chapter 8 | 2 | ch8-create-demo-issue-3.yml（2D matrix）、ch8-create-failure-issue.yml |
| Chapter 9 | 3 | evaluate-pr.yml、ch9-forked-example1.yml、ch9-forked-example2.yml |
| Chapter 12 | 6 | ch12-repo-info.yml、ch12-create-demo-issue.yml、ch12-create-repo-issue.yml、ch12-create-demo-issue3.yml、ch12-create-repo-issue3.yml、ch12-get-info.yml |
| Chapter 13 | 4 | ch13-create-issue-via-gh.yml、ch13-create-issues-across-prods.yml、ch13-create-issues-across-prods-and-levels.yml、ch13-create-issues-from-context.yml |

加上原有的 `simple-go-build.yml`，共 **23 個 workflow**。

### GitHub Repo 設定
- 建立 3 個 Configuration Variables（`vars.*`）供 `verify-file.yml` 使用：
  - `EXEC_WF=true`
  - `FILE_TO_CHECK=README.md`
  - `JOB_NAME="Verify my file"`
- 建立 `.github/workflow-templates/` 目錄（Chapter 12 org-level template 範例）

### README 更新
- 改寫 README.md：加入 CI badge、23 個 workflow 完整分類清單、注意事項

---

## 關鍵技術筆記

### Go 時區嵌入
```go
import _ "time/tzdata"
// 解決 GitHub Actions ubuntu-latest runner 預設無 tzdata 套件
// time.LoadLocation("Asia/Taipei") 在 Linux CI 環境不會失敗
```

### Workflow 觸發類型（本專案範例涵蓋）
- `push` / `pull_request` / `pull_request_target`（PR 安全性差異：後者有寫入權限）
- `workflow_dispatch`（手動觸發，支援 inputs）
- `workflow_call`（reusable workflow 被呼叫）
- `repository_dispatch`（外部事件觸發，含 `client_payload`）

### Matrix 演進
1. 1D matrix：`strategy.matrix.prod: [prod1, prod2]`
2. 2D matrix：prod × level，組合爆炸（6 jobs）
3. 動態 matrix：從 `${{ github.event.client_payload.levels }}` 讀取（最進階）

### fork PR 安全性（Chapter 9 重點）
- `pull_request`：fork PR 無 secrets 存取，安全但功能受限
- `pull_request_target`：有 secrets，但 fork PR 可透過 workflow 修改洩漏機密

### Chapter 12/13 外部呼叫限制
- `rndrepos/common` 呼叫預期失敗（需對應 PAT secret）
- 本地 reusable workflow 呼叫（如 `ch12-get-info.yml` → `ch12-repo-info.yml`）可正常運作

---

## 產出檔案

| 檔案 | 動作 | 說明 |
|------|------|------|
| `helloworld.go` | 修改 | 加入 Taipei 時區時間輸出 |
| `CLAUDE.md` | 新增 | 專案架構指引 |
| `README.md` | 修改 | 加入 23 個 workflow 清單與 CI badge |
| `.github/workflows/simple-go-build.yml` | 既有 | CI 主 workflow |
| `.github/workflows/basic.yml` | 新增 | Ch4 |
| `.github/workflows/simple-proj.yml` | 新增 | Ch4 |
| `.github/workflows/system-environment-settings.yml` | 新增 | Ch6 |
| `.github/workflows/verify-file.yml` | 新增 | Ch6 |
| `.github/workflows/output-from-a-job.yml` | 新增 | Ch7 |
| `.github/workflows/output-from-a-step.yml` | 新增 | Ch7 |
| `.github/workflows/create-failure-issue.yml` | 新增 | Ch7 reusable |
| `.github/workflows/ch8-create-demo-issue-3.yml` | 新增 | Ch8 2D matrix |
| `.github/workflows/ch8-create-failure-issue.yml` | 新增 | Ch8 reusable |
| `.github/workflows/evaluate-pr.yml` | 新增 | Ch9 |
| `.github/workflows/ch9-forked-example1.yml` | 新增 | Ch9 |
| `.github/workflows/ch9-forked-example2.yml` | 新增 | Ch9 |
| `.github/workflows/ch12-repo-info.yml` | 新增 | Ch12 reusable |
| `.github/workflows/ch12-create-demo-issue.yml` | 新增 | Ch12 |
| `.github/workflows/ch12-create-repo-issue.yml` | 新增 | Ch12 |
| `.github/workflows/ch12-create-demo-issue3.yml` | 新增 | Ch12 |
| `.github/workflows/ch12-create-repo-issue3.yml` | 新增 | Ch12 |
| `.github/workflows/ch12-get-info.yml` | 新增 | Ch12 |
| `.github/workflows/ch13-create-issue-via-gh.yml` | 新增 | Ch13 gh CLI |
| `.github/workflows/ch13-create-issues-across-prods.yml` | 新增 | Ch13 1D matrix |
| `.github/workflows/ch13-create-issues-across-prods-and-levels.yml` | 新增 | Ch13 2D matrix |
| `.github/workflows/ch13-create-issues-from-context.yml` | 新增 | Ch13 dynamic matrix |
| `test/.gitkeep` | 新增 | Ch6 working-directory 需求 |
| `.github/workflow-templates/rndrepos-info.yml` | 新增 | Ch12 org template |
| `.github/workflow-templates/rndrepos-info.properties.json` | 新增 | Ch12 org template |
| `.github/workflow-templates/check-square.svg` | 新增 | Ch12 org template |
| `summary-02-sessions/2026-05-10/session1-summary.md` | 新增 | 本 session summary |

---

## HANDOFF（下次 session 優先處理）

### 立即行動
- [ ] 書本只到 Chapter 13，若使用者想繼續學習，可考慮自己撰寫自訂 Action（Chapter 11 主題：creating custom actions）
- [ ] 確認 `verify-file.yml` 實際在 GitHub Actions 上執行是否成功（vars 已設定）
- [ ] 若使用者想深入，可練習用 `repository_dispatch` + curl 手動觸發 `ch13-create-issues-from-context.yml`

### 進行中（需接續）
- 無未完成工作，本 session 書本章節 4-13 全數 push 完畢

### 注意事項
- Chapter 12/13 中呼叫外部 `rndrepos/common` 的 workflow 永遠會失敗（設計如此，書中 demo 用）
- `verify-file.yml` 依賴 repo Configuration Variables（`vars.EXEC_WF` 等），未設定時 job 會 skip
- Chapter 9 的 fork PR 示範 workflow 有安全漏洞（故意的），不要在生產 repo 使用
- 本書沒有 chapter-14 目錄，書本範例至 chapter-13 結束

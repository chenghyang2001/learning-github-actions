# learning-github-actions

《Learning GitHub Actions》（O'Reilly，Brent Laster 著）書本範例倉庫。

[![Simple Go Build](https://github.com/chenghyang2001/learning-github-actions/actions/workflows/simple-go-build.yml/badge.svg)](https://github.com/chenghyang2001/learning-github-actions/actions/workflows/simple-go-build.yml)

---

## Workflow 清單

`.github/workflows/` 共 23 個 workflow，依章節分類如下。

### 實際運作的 CI

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `simple-go-build.yml` | push to main | 核心 CI：checkout → setup Go 1.21 → `go run helloworld.go` |

### Chapter 4：基本 Workflow 結構

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `basic.yml` | push/PR to main、手動 | 兩個平行 job：checkout、run echo |
| `simple-proj.yml` | push/PR to main、手動 | demo 基本 workflow 骨架 |

### Chapter 6：環境變數與 Deployment Environments

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `system-environment-settings.yml` | push to main | `defaults.run` 示範，印出環境變數 |
| `verify-file.yml` | push to main（`vars.EXEC_WF` 控制） | 用 Configuration Variables 決定是否執行 |

### Chapter 7：輸出傳遞

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `output-from-a-job.yml` | push to main | job output → 下一個 job 讀取 |
| `output-from-a-step.yml` | push to main | step output → 同 job 下一個 step 讀取 |
| `create-failure-issue.yml` | `workflow_call` | reusable：失敗時用 curl 建立 GitHub Issue |

### Chapter 8：Matrix 與 Workflow Dispatch

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `ch8-create-demo-issue-3.yml` | `workflow_dispatch` | 2D matrix：`[prod1, prod2] × [dev, test, stage]` |
| `ch8-create-failure-issue.yml` | `workflow_call` | reusable：失敗時建立 Issue |

### Chapter 9：PR 安全性

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `evaluate-pr.yml` | `pull_request_target` | PR 評估腳本，具寫入權限 |
| `ch9-forked-example1.yml` | push/PR to main | fork PR 漏洞示範：環境變數洩漏 |
| `ch9-forked-example2.yml` | push/PR to main | fork PR 漏洞示範：`GITHUB_TOKEN` 洩漏 |

### Chapter 12：Reusable Workflows

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `ch12-repo-info.yml` | `workflow_call` | reusable：回傳 repo 資訊（inputs/outputs） |
| `ch12-create-demo-issue.yml` | `workflow_dispatch` | 呼叫外部 reusable workflow 建 issue |
| `ch12-create-repo-issue.yml` | `workflow_dispatch` | 同上，附帶 secrets 傳遞 |
| `ch12-create-demo-issue3.yml` | `workflow_dispatch` | 呼叫外部 reusable workflow（含 outputs） |
| `ch12-create-repo-issue3.yml` | `workflow_dispatch` | 同上，用 outputs 建第二個 issue |
| `ch12-get-info.yml` | `workflow_dispatch` | 呼叫本地 `ch12-repo-info.yml` 取 outputs |

### Chapter 13：GitHub CLI 與進階 Matrix

| 檔案 | 觸發方式 | 說明 |
|------|---------|------|
| `ch13-create-issue-via-gh.yml` | `workflow_call` | reusable：用 `gh` CLI 建 issue |
| `ch13-create-issues-across-prods.yml` | `workflow_dispatch` | 1D matrix `[prod1, prod2]` 建 issue |
| `ch13-create-issues-across-prods-and-levels.yml` | `workflow_dispatch` | 2D matrix `[prod × level]` 建 issue |
| `ch13-create-issues-from-context.yml` | `repository_dispatch` | 動態 matrix 從 `client_payload` 讀取 |

---

## 本機執行

```bash
go run helloworld.go
```

## 注意事項

- 章節 YAML 範例存放於 `chapter-*/` 目錄（不在 `.github/workflows/`），避免每次 push 都觸發所有示範 workflow。
- Chapter 12、13 中呼叫外部 `rndrepos/common` 的 workflow 預期會失敗（需要對應的 PAT secret）。
- `verify-file.yml` 需要在 repo 設定 `vars.EXEC_WF=true`、`vars.FILE_TO_CHECK`、`vars.JOB_NAME` 三個 Configuration Variables。

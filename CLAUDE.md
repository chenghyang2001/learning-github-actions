# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 專案概述

《Learning GitHub Actions》（O'Reilly）書本範例倉庫。包含書中各章節的 YAML workflow 範例，以及一個實際運作在 `.github/workflows/` 的 CI 流程。

## 執行方式

```bash
# 本機執行 Go 程式
go run helloworld.go

# 推送到 main 分支後，GitHub Actions 自動觸發 CI
git push origin main
```

## 架構說明

```
.github/workflows/
  simple-go-build.yml   # 唯一實際觸發的 workflow（push to main）

chapter-2/              # 各章節 workflow 範例（僅供參考，不在 .github/workflows/ 中）
chapter-4/
chapter-5/
chapter-6/              # 環境變數、Deployment Environments
chapter-7/              # job/step outputs、workflow inputs
chapter-8/              # matrix strategy、workflow_dispatch
chapter-9/              # PR validation、forked repo 安全性
chapter-12/             # reusable workflows、inputs/secrets/outputs
chapter-13/             # GitHub CLI、多維度 matrix

helloworld.go           # CI 的最小 Go 目標（go run helloworld.go）
```

## 關鍵設計決策

- **章節 YAML 不在 `.github/workflows/`**：書中範例以原始路徑存放（如 `chapter-2/simple-go-build.yml`），避免每次 push 都觸發所有示範 workflow。
- **實際 CI workflow**（`.github/workflows/simple-go-build.yml`）使用較新的 action 版本（`actions/checkout@v4`、`actions/setup-go@v5`、Go 1.21），而章節原始範例保留書中的舊版本以便對照。
- **reusable workflow 範例**（chapter-12、chapter-13）引用外部 `rndrepos/common` 倉庫，本地無法單獨執行，須在有 PAT secret 的 GitHub 環境下才能跑通。

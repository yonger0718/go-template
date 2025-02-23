# scripts

本目錄可以存放自動化、測試、部署等相關的腳本。

> 注意: 我們應該盡可能地將外部的makefile、shell script等移植到此目錄，並且只保留必要的腳本，也應該盡可能的簡化主路徑下腳本的內容。

## 說明

目前這個目錄是空的，你可以根據需要添加腳本，例如：

*   `build.sh`:  建置專案的腳本。
*   `deploy.sh`:  部署專案的腳本。
*   `test.sh`:  執行測試的腳本。
*   `setup.sh`: 設定開發環境的腳本。

等等。

## example

* [prometheus/scripts](https://github.com/prometheus/prometheus/tree/main/scripts)

* [harness/scripts](https://github.com/harness/harness/tree/main/scripts)
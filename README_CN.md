# Optimus

> 桌面端图片压缩、优化与格式转换工具。

该项目 fork 自 https://github.com/Splode/optimus

## 概览

Optimus 是一款桌面端图片优化应用，支持 WebP、JPEG、PNG 等图片格式之间的转换与压缩。

## 功能特性

- 在 JPEG、PNG、WebP 格式之间相互转换
- 压缩 JPEG、PNG（有损）与 WebP（有损与无损）
- 批量将图片调整为多种尺寸
- 查看本次会话与历史累计的简单统计信息

![Optimus 主编辑界面截图](./.github/optimus_screenshot_editor--1200x742.png)

![Optimus 设置界面截图](./.github/optimus_screenshot_options--1200x936.png)

## 安装

### 下载

请从 [Releases 页面](https://github.com/panyingyun/optimus/releases) 下载最新版本。

Optimus 支持 Windows、macOS 和 Linux。

## 开发

Optimus 基于 [Wails](https://wails.app/) v2.11.0 构建：前端使用 JavaScript，后端使用 Go。

### 环境要求

- Go 1.22.0 或更高版本
- Node.js 与 npm
- Wails v2.11.0

### 初始化

1. 克隆仓库：

```bash
git clone https://github.com/panyingyun/optimus.git
cd optimus
```

2. 安装 Wails：

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

3. 安装系统依赖（Linux）：

Ubuntu 22.04：

```bash
sudo apt install libwebkit2gtk-4.0-dev libgtk-3-dev -y
```

Ubuntu 24.04：

```bash
sudo apt install libwebkit2gtk-4.1-dev libgtk-3-dev -y
```

4. 安装 Go 依赖：

```bash
go mod tidy
```

5. 安装前端依赖：

```bash
cd frontend
npm install
cd ..
```

### 开发模式

启动开发服务：

Ubuntu 22.04：

```bash
make dev_ubuntu2204
# 或
wails dev -tags webkit2_40
```

Ubuntu 24.04：

```bash
make dev_ubuntu2404
# 或
wails dev -tags webkit2_41
```

### 构建

构建应用：

Ubuntu 22.04：

```bash
make build_ubuntu2204
# 或
wails build -tags webkit2_40
```

Ubuntu 24.04：

```bash
make build_ubuntu2404
# 或
wails build -tags webkit2_41
```

Windows：

```bash
make build-windows
# 或
go build -tags desktop,production -ldflags "-w -s -H windowsgui"
```

## 许可证

MIT


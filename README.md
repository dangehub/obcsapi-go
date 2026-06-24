# Obcsapi Go

> 本仓库是 [kkbt/obcsapi-go](https://github.com/kkbt0/obcsapi-go) 的**二次修改版**，保留原项目名称以表对原作者的尊重。
> 如需原版请访问上游仓库。

基于 Obsidian S3 存储、CouchDb、本地存储和 WebDAV 的后端 API。可借助 Obsidian 插件 Remotely-Save 或 Self-hosted LiveSync 保存消息到 Obsidian 库，也支持本地文件夹的文本编辑器。

## 特点

- 前端 Memos / 简答编辑器，支持指令模式、黑暗主题、PWA 应用
- 微信测试号 → Obsidian
- 简悦 SimpRead Webhook 裁剪网页文章
- FV 悬浮球文字图片分享保存
- 静读天下 MoonReader 高亮标注
- 通用 HTTP API
- 使用 Lua & Bash 拓展功能
- WebDAV 服务
- 简易图床，附带命令行上传工具
- SCF 或 Docker 部署

## 与本版相比的改动

详见文档：[分叉说明](https://dangehub.github.io/obcsapi-go/md/go-version/11-分叉说明)

## 文档

文档基于上游文档二次构建：

- GitHub Pages: https://dangehub.github.io/obcsapi-go/
- 上游文档: https://www.ftls.xyz/docs/obcsapi/

## 部署

```bash
# Docker
docker run -d -p 8900:8900 --name obcsapi \
  -v /your/obsidian/vault/:/app/data/webdav/note/ \
  ghcr.io/dangehub/obcsapi-go:latest

# 本地构建
cd server/
go run .
```

## 鸣谢

- [恐咖兵糖 (kkbt)](https://github.com/kkbt0) — 原版作者

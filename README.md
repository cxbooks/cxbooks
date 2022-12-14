<div align=center>
<img src="https://img.shields.io/badge/golang-1.19-blue"/>
<img src="https://img.shields.io/badge/vue-3.2.31-brightgreen"/>
<img src="https://img.shields.io/badge/vuetify-3.0.0-green"/>
</div>

[English](README.en.md) | 简体中文

**项目刚刚开始，预计11月份可以出一个能运行的版本**

# ![cxbooks](docs/images/logo.png)

cxbooks 是使用 Golang 编写的个人电子书管理WEB服务，提供类似 jellyfin 和 navidrome 文档搜刮与管理功能。


> **提醒：中国境内网站，个人是不允许进行在线出版的，维护公开的书籍网站是违法违规的行为！建议仅作为个人使用！**

## 目标与功能支持

- [x] 美观的界面：由于Calibre自带的网页太丑太难用，于是基于Vue，独立编写了新的界面，支持PC访问和手机浏览；
- [ ] 支持多用户：为了网友们更方便使用，开发了多用户功能，支持豆瓣（已废弃）、QQ、微博、Github等社交网站的登录；
- [ ] 支持在线阅读：借助Readium.js 库，支持了网页在线阅读电子书；
- [x] 支持目录书籍搜刮，不需要额外导入过程；
- [ ] 支持邮件推送：可方便推送到Kindle；
- [x] 支持OPDS：可使用KyBooks等APP方便地读书；
- [x] 支持一键安装，网页版初始化配置，轻松启动网站；
- [ ] 支持快捷更新书籍信息：支持从百度百科、豆瓣搜索并导入书籍基础信息；
- [x] 支持文件去重，提供多种去重复机制；
- [ ] 支持上传文件问题
- [ ] 支持其他数据文件阅读支持，诸如 Markdown 文件
- [ ] 提供文件格式转换功能（使用 Calibre ）


## 部署

### docker方式

1. 快速部署默认方式

```bash
docker run -d --name cxbooks -p 8080:80   -e PUID=1000 -e PGID=1000  -v /localdata:/data cxbooks/cxbooks

```

2. 高级参数



## 项目依赖以及三分框架使用说明
- sqlite3: 存储图书目录
- [nutsdb](https://github.com/nutsdb/nutsdb) : 磁盘基本缓存功能（诸如Redis作用）
- Calibre: [可选]提供文档格式转换功能
- 

# 鸣谢

本项目管理界面功能严重参（抄）考（袭）了 [talebook](https://github.com/talebook/talebook) 和 [Calibre-web](https://github.com/janeczku/calibre-web) 界面设计


# ThinkPHP 5.x 远程代码执行漏洞

ThinkPHP 诞生于 2006 年，是一个国产开源的 PHP 开发框架，其借鉴了 Struts 框架的 Action 对象，同时也使用面向对象的开发结构和 MVC 模式。ThinkPHP 可在 Windows 和 Linux 等操作系统运行，支持 MySQL、Sqlite 和 PostgreSQL 等多种数据库以及 PDO 扩展，是一款跨平台、跨版本以及简单易用的 PHP 框架。

在 ThinkPHP 5 中，由于框架对控制器名没有进行足够的检测，会导致在没有开启强制路由的情况下的远程代码执行漏洞。

**影响版本**：ThinkPHP 5.x

**[FOFA](https://fofa.so/result?qbase64=YXBwPSJUaGlua1BIUCI%3D) 查询规则**：app="ThinkPHP"

# Demo

![](thinkphp_5.gif)
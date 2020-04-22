# ThinkPHP 2.1 远程代码执行

ThinkPHP 诞生于 2006 年，是一个国产开源的 PHP 开发框架，其借鉴了 Struts 框架的 Action 对象，同时也使用面向对象的开发结构和 MVC 模式。ThinkPHP 可在 Windows 和 Linux 等操作系统运行，支持 MySql、Sqlite 和 PostgreSQL 等多种数据库以及 PDO 扩展，是一款跨平台，跨版本以及简单易用的 PHP 框架。

ThinkPHP 2.1 版本中，使用 `preg_replace` 的 `/e` 模式匹配路由：

```php
$res = preg_replace('@(\w+)'.$depr.'([^'.$depr.'\/]+)@e', '$var[\'\\1\']="\\2";', implode($depr,$paths));
```

这是个非常危险的参数，如果用了这个参数，`preg_replace` 的第二个参数就会被当做 PHP 代码执行。

**影响版本**：ThinkPHP 2.1

**[FOFA](https://fofa.so/result?qbase64=YXBwPSJUaGlua1BIUCI%3D) 查询规则**：app="ThinkPHP"

# Demo

![](thinkphp_2.1.gif)
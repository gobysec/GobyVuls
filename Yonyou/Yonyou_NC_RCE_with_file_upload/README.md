# Yonyou NC RCE with file upload

用友NC是一款企业级管理软件，在大中型企业广泛使用。实现建模、开发、继承、运行、管理一体化的IT解决方案信息化平台。用友NC 为C/S 架构，使用JAVA编程语言开发，客户端可直接使用UClient，服务端接口为HTTP。用友NC6.5的某个页面，存在任意文件上传漏洞。漏洞成因在于上传文件处未作类型限制，未经身份验证的攻击者可通过向目标系统发送特制数据包来利用此漏洞，成功利用此漏洞的远程攻击者可在目标系统上传任意文件执行命令。

**Affected version**: 用友 NC 6.5

**[FOFA](https://fofa.so/result?q=title%3D%22YONYOU+NC%22&qbase64=dGl0bGU9IllPTllPVSBOQyI%3D&file=&file=) query rule**: title="YONYOU NC"

# Demo

![](Yonyou_NC_RCE_with_file_upload.gif)
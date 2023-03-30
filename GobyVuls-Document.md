# Goby History Update Vulnerability Total Document (Continuously Update) 
The following content is an updated vulnerability from Goby. Some of the vulnerabilities are recorded on the screen for easy viewing.

**Updated document date: March 17, 2023** 

## Cockpit File Upload Vulnerability(CVE-2023-1313)

|   **Vulnerability**  | **Cockpit File Upload Vulnerability(CVE-2023-1313)**  |
| :----:   | :-----|
|  **Chinese name**  | Cockpit 平台 upload 文件上传漏洞（CVE-2023-1313）  |
| **CVSS core**  | 7.2 |
| **FOFA Query**  (click to view the results directly)| [app="cockpit"](https://fofa.info/result?qbase64=YXBwPSJjb2NrcGl0Ig%3D%3D) |
| **Number of assets affected**  | 1643 |
| **Description**  | Cockpit is a self-hosted, flexible and user-friendly headless content platform for creating custom digital experiences.</p><p>Cockpit has a file upload vulnerability, which allows attackers to upload arbitrary files, leading to server control, etc. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

## JeecgBoot Default Password Vulnerability

|   **Vulnerability**  | **JeecgBoot Default Password Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | JeecgBoot 开发平台默认口令漏洞  |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [app="JeecgBoot-企业级低代码平台"](https://fofa.info/result?qbase64=YXBwPSJKZWVjZ0Jvb3Qt5LyB5Lia57qn5L2O5Luj56CB5bmz5Y%2BwIg%3D%3D) |
| **Number of assets affected**  | 3965 |
| **Description**  | JeecgBoot is a low -code development platform based on code generator. |
| **Impact** | Attackers can control the entire platform through default password vulnerabilities and use administrator privileges to operate core functions. |


**CVSS core**: 7.5

**FOFA query** (click to view the results directly):

[app="JeecgBoot-企业级低代码平台"](https://fofa.info/result?qbase64=YXBwPSJKZWVjZ0Jvb3Qt5LyB5Lia57qn5L2O5Luj56CB5bmz5Y%2BwIg%3D%3D)

**Number of assets affected**： 3965

**Description** : JeecgBoot is a low -code development platform based on code generator.

**Impact** : Attackers can control the entire platform through default password vulnerabilities and use administrator privileges to operate core functions.

## Nacos Authentication Bypass Vulnerability
**Chinese name: Nacos 身份认证绕过漏洞**

**Description** : Nacos is a service management platform for building cloud native applications.
The open source service management platform Nacos has a high-risk vulnerability of authentication bypass in versions 0.1.0 to 2.20, which causes attackers to bypass key authentication and enter the background, resulting in system control and other consequences.Nacos is a service management platform for building cloud native applications.

**Impact** : The open source service management platform Nacos has a high-risk vulnerability of authentication bypass in versions 0.1.0 to 2.20, which causes attackers to bypass key authentication and enter the background, resulting in system control and other consequences.

![](https://s3.bmp.ovh/imgs/2023/03/17/98a67053444f1997.gif)

## Weaver e-cology OA browser.jsp keyword SQL Injection Vulnerability

**Chinese name: 泛微-协同办公OA browser.jsp 文件 keyword 参数 SQL注入漏洞**

**Description** : Weaver e-cology OA, also known as Ubiq Collaborative office system, is a high-quality office system built on the principle of simplicity, application and efficiency. The software has more than 20 functional modules including process, portal, knowledge, personnel and communication, and adopts intelligent voice interactive office mode, which can perfectly fit the actual needs of enterprises and open up the whole digital management for enterprises
The browser.jsp file has the s q l injection vulnerability, through which the attacker can obtain sensitive database information.

**Impact** : In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions.

## Atlassian Jira snjFooterNavigationConfig fileName Arbitrary File Read Vulnerability (CVE-2023-26256)、Atlassian Jira snjCustomDesignConfig fileName Arbitrary File Read Vulnerability (CVE-2023-26255)
**Chinese name: Atlassian Jira 缺陷跟踪管理系统 snjFooterNavigationConfig 文件 fileName 参数任意文件读取漏洞（CVE-2023-26256）、Atlassian Jira 缺陷跟踪管理系统 snjCustomDesignConfig 文件 fileName 参数任意文件读取漏洞（CVE-2023-26255）**

**Description** : Atlassian Jira is a set of defect tracking management system of Atlassian company in Australia. The system is mainly used to track and manage various problems and defects in the work.

There is a security vulnerability in Jira plugin STAGIL Navigation before 2.0.52. The vulnerability stems from a path traversal vulnerability, which allows attackers to traverse and read the file system.

**Impact** : There is a security vulnerability in Jira plugin STAGIL Navigation before 2.0.52. The vulnerability stems from a path traversal vulnerability, which allows attackers to traverse and read the file system.

## Fortinet FortiNAC keyUpload.jsp Arbitrary File Upload Vulnerability (CVE-2022-39952)
**Chinese name: Fortinet FortiNAC keyUpload.jsp 任意文件上传漏洞（CVE-2022-39952）**

**Description** : Fortinet FortiNAC is a zero-trust access solution from Fortinet.
Fortinet FortiNAC has a security vulnerability. The attacker uploads a maliciously compressed Trojan horse file through keyUpload.jsp to obtain server permissions.

**Impact** : Fortinet FortiNAC has a security vulnerability. The attacker uploads a maliciously compressed Trojan horse file through keyUpload.jsp to obtain server permissions.

![](https://s3.bmp.ovh/imgs/2023/03/17/381c6d29699d92eb.gif)

## Topsec ACM download.php Any file download
**Chinese name: 天融信 topsec ACM 系统 download.php 任意文件下载**

**Description** : Topsec topsec ACM is a professional product of Topsec company for network behavior management and content auditing for all walks of life with years of experience in security product research and development. The system not only has the functions of preventing illegal information dissemination, sensitive information leakage, real-time monitoring, log tracing, network resource management, but also powerful user management, report statistics and analysis functions.
There is an arbitrary file download vulnerability in the download.php file of TopSec ACM of TopSec ACM, and attackers can use this to read arbitrary files in the system.

**Impact** : There is an arbitrary file download vulnerability in the download.php file of TopSec ACM of TopSec ACM, and attackers can use this to read arbitrary files in the system.

## QNAP Photo Station the filename parameter of the video.php file is read arbitrarily vulnerability
**Chinese name: QNAP NAS网络储存设备video.php文件的filename参数任意文件读取**

**Description** : QNAP NAS is a suite of network storage devices from QNAP Systems. For home, SOHO, and SME users, QNAP Systems Photo Station is a photo management and viewing application that allows users to bring together photos scattered across multiple terminal devices for management, editing, and sharing, with vulnerabilities in the Photo Station and CGI modules.

**Impact** : QNAP NAS is a suite of network storage devices from QNAP Systems. The filename parameter of the /photo/p/api/video.php file of QNAP NAS has an arbitrary file read vulnerability, which is due to the controllable exportFile() parameter, and the identity verification can be bypassed by constructing specific parameters even without authorization, resulting in an arbitrary file read vulnerability, which can allow attackers to view sensitive information and obtain higher privilege access.

## SangFor AD login clsMode Command Execution Vulnerability
**Chinese name: 深信服应用交付管理系统 login 文件 clsMode 参数命令执行漏洞**

**Description** : SangFor AD It provides users with comprehensive solutions, including multi-DC load balancing, multi-link load balancing, and server load balancing. It not only realizes real-time monitoring of the status of each data center, link and server, but also allocates the user's access request to the corresponding data center, link and server according to preset rules, so as to realize the rational distribution of data flow and make full use of all data centers, links and servers. The login interface of version 7.0.8-7.0.8r5 has command execution vulnerabilities. Attackers obtain server permissions through command concatenation

**Impact** : Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.

## QNAP-NAS authLogin.cgi app_token RCE Vulnerability (CVE-2022-27596)
**Chinese name: QNAP-NAS authLogin.cgi 文件 app_token 参数代码执行漏洞（CVE-2022-27596）**

**Description** : QNAP Systems QTS is an operating system used by China's QNAP Systems for entry-level to mid-level QNAP NAS.
There is a security vulnerability in QNAP Systems QTS. The vulnerability stems from the fact that devices running QuTS hero and QTS allow remote attackers to inject malicious code into the app_token parameter field to obtain server permissions.

**Impact** : There is a security vulnerability in QNAP Systems QTS. The vulnerability stems from the fact that devices running QuTS hero and QTS allow remote attackers to inject malicious code into the app_token parameter field to obtain server permissions.

## Nostromo path traversal vulnerability（CVE-2022-48253）
**Chinese name: Nostromo 路径穿越漏洞（CVE-2022-48253）**

**Description** : Nostromo (aka nhttpd) is a simple and fast open source web server.
Nostromo 2.1 was affected by path traversal, which could allow an attacker to do arbitrary file reading and, if run with permissions, execute arbitrary commands on a remote server. (This vulnerability occurs when using the homedirs option)

**Impact** : Nostromo 2.1 was affected by path traversal, which could allow an attacker to do arbitrary file reading and, if run with permissions, execute arbitrary commands on a remote server. (This vulnerability occurs when using the homedirs option)

## Ruckus Wireless Admin Command Execution Vulnerability (CVE-2023-25717)
**Chinese name: Ruckus Wireless Admin 命令执行漏洞（CVE-2023-25717）**

**Description** : Ruckus Wireless Admin is the background management system for multiple routers and hardware devices of ruckuswireless.
A command execution vulnerability exists in Ruckus Wireless Admin version 10.4 and earlier.

**Impact** : Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.

## Aspera Faspex relay_package RCE Vulnerability(CVE-2022-47986)
**Chinese name: Aspera Faspex relay_package 远程代码执行漏洞（CVE-2022-47986）**

**Description** : Aspera Faspex is a set of fast file transfer and streaming solutions based on the IBM FASP protocol built by IBM Corporation of the United States.
There is a security vulnerability in Aspera Faspex. The vulnerability stems from the lack of security check in the /package_relay/relay_package path. Attackers can use this vulnerability to execute arbitrary code to obtain server permissions.

**Impact** : There is a security vulnerability in Aspera Faspex. The vulnerability stems from the lack of security check in the /package_relay/relay_package path. Attackers can use this vulnerability to execute arbitrary code to obtain server permissions.

## SolarView Compact downloader.php RCE (CVE-2023-23333)
**Chinese name: SolarView Compact downloader.php 存在任意命令执行漏洞（CVE-2023-23333）**

**Description** : There is a command injection vulnerability in SolarView Compact through 6.00, attackers can execute commands by bypassing internal restrictions through downloader.php.

**Impact** : Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.

## Smartbi Unauthorized And JDBC Arbitrary Code Execution Vulnerability
**Chinese name: Smartbi 未授权访问及 JDBC 任意代码执行漏洞**

**Description** : Smartbi is a business intelligence BI software launched by Smart Software, which meets the development stage of BI products. Smart software integrates the functional requirements of data analysis and decision support in various industries to meet the big data analysis needs of end users in enterprise-level reports, data visualization analysis, self-service exploration analysis, data mining modeling, AI intelligent analysis and other scenarios.

**Impact** : There is an unauthorized access background interface vulnerability between Smartbi V7 and V10.5.8. Combined with postgresql JDBC, it can write arbitrary files or execute arbitrary code to obtain server permissions.

![](https://s3.bmp.ovh/imgs/2023/03/17/c33297736a9aa870.gif)

## yunucms request_uri method code execution vulnerabilities
**Chinese name: yunucms 城市分站管理系统 request_uri 参数代码执行漏洞**

**Description** : yunucms is a free and open source urban substation management system developed by Yunyou Network Technology Co., Ltd. based on the TP5.0 framework.

**Impact** : There is a code execution vulnerability in the request_uri parameter of the front-end wap/index/index method of the yunucms system v1-2.0.5. Attackers can obtain server permissions through the vulnerability.

## WordPress Plugin BackupBuddy Arbitrary File Read Vulnerability (CVE-2022-31474)
**Chinese name: WordPress BackupBuddy 插件 local-download 参数任意文件读取漏洞（CVE-2022-31474）**

**Description** : WordPress BackupBuddy plugin is a fast and simple plugin for WordPress backup and restore.
WordPress plugin BackupBuddy versions 8.5.8.0 to 8.7.4.1 have an information disclosure vulnerability, which stems from an arbitrary file read and download vulnerability.

**Impact** : Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website.

## TurboMail mail system viewfile file reading vulnerability
**Chinese name: TurboMail 邮件系统 viewfile 文件读取漏洞**

**Description** : TurboMail mail system is an email server system developed for the communication needs of enterprises and institutions. There is a file reading vulnerability in the TurboMail mail system. An attacker can read the configuration file through this vulnerability, and then perform base64 decryption on the password to log in to the background/maintlogin.jsp.

**Impact** : There is a file reading vulnerability in the TurboMail mail system. An attacker can read the configuration file through the /viewfile endpoint, and then decrypt the password to base64 and log in to the background /maintlogin.jsp.

## Apache Guacamole tokens Api Default Credential Vulnerability
**Chinese name: Apache Guacamole tokens 接口默认密码漏洞**

**Description** : Apache Guacamole is a clientless remote desktop gateway. It supports standard protocols like VNC, RDP, and SSH. Apache Guacamole default password may lead information disclosure.

**Impact** : Attackers can control the entire platform through default password vulnerabilities and use administrator privileges to operate core functions.

## Liferay Portal Unauthenticated 7.2.1 RCE (CVE-2020-7961)
**Chinese name: Liferay Portal 7.2.1 版本 invoke 文件远程代码执行漏洞（CVE-2020-7961）**

**Description** : Liferay Portal is a set of J2EE-based portal solutions of American Liferay Company. The program uses EJB and JMS and other technologies, and can be used as Web publishing and sharing workspace, enterprise collaboration platform, social network and so on.

**Impact** : A code issue vulnerability exists in versions prior to Liferay Portal 7.2.1 CE GA2. A remote attacker could exploit this vulnerability to execute arbitrary code using JSON Web services.

## SugarCRM index.php File Upload Vulnerability (CVE-2023-22952)
**Chinese name: SugarCRM index.php 任意文件上传漏洞（CVE-2023-22952）**

**Description** : SugarCRM is a set of open source customer relationship management system (CRM) of American SugarCRM company. The system supports differentiated marketing, management and distribution of sales leads for different customer needs, and realizes information sharing and tracking of sales representatives.

**Impact** : SugarCRM has a security vulnerability. The vulnerability stems from an authorization bypass and PHP local file inclusion vulnerability in the installation component, which allows unauthenticated remote code execution on the configured SugarCRM instance through HTTP requests.

## Oracle E-Business Suite BneViewerXMLService Arbitrary File Upload Vulnerability (CVE-2022-21587)
**Chinese name: Oracle E-Business Suite 软件 BneViewerXMLService 任意文件上传漏洞（CVE-2022-21587）**

**Description** : Oracle E-Business Suite (E-Business Suite) is a set of fully integrated global business management software from Oracle Corporation. The software provides functions such as customer relationship management, service management, and financial management.

**Impact** : A security vulnerability exists in Oracle Web Applications Desktop Integrator versions 12.2.3-12.2.11 of Oracle E-Business Suite. An unauthenticated attacker gains server privileges by uploading a malicious webshell file.

## WordPress Plugin WP Live Chat Support path File Inclusion Vulnerability
**Chinese name: WordPress WP Live Chat Support 插件 path 文件包含漏洞**

**Description** : WP Live Chat Support is a reliable and tested live chat solution for WordPress.
There is a file inclusion vulnerability in WP Live Chat Support <= 9.4.2.Attackers can exploit this vulnerability to obtain sensitive files.

**Impact** : Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website.

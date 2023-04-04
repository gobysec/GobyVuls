# Goby History Update Vulnerability Total Document (Continuously Update) 
The following content is an updated vulnerability from Goby. Some of the vulnerabilities are recorded on the screen for easy viewing.

**Updated document date: April 04, 2023** 

## QNAP-NAS authLogin.cgi app_token RCE Vulnerability (CVE-2022-27596)

|   **Vulnerability**  | **QNAP-NAS authLogin.cgi app_token RCE Vulnerability (CVE-2022-27596)**  |
| :----:   | :-----|
|  **Chinese name**  | QNAP-NAS authLogin.cgi 文件 app_token 参数代码执行漏洞（CVE-2022-27596） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [(((header="http server" \&\& body="redirect_suffix") \|\| body="/css/qnap-default.css" \|\| body="/redirect.html?count=\\\"+Math.random()" \|\| body="/indexnas.cgi?counter=") && body!="Server: couchdb") \|\| (body="qnap_hyperlink" && body="QNAP Systems, Inc.\</a \> All Rights Reserved.")](https://fofa.info/result?qbase64=KCgoaGVhZGVyPSJodHRwIHNlcnZlciIgJiYgYm9keT0icmVkaXJlY3Rfc3VmZml4IikgfHwgYm9keT0iL2Nzcy9xbmFwLWRlZmF1bHQuY3NzIiB8fCBib2R5PSIvcmVkaXJlY3QuaHRtbD9jb3VudD1cXFwiK01hdGgucmFuZG9tKCkiIHx8IGJvZHk9Ii9pbmRleG5hcy5jZ2k%2FY291bnRlcj0iKSAmJiBib2R5IT0iU2VydmVyOiBjb3VjaGRiIikgfHwgKGJvZHk9InFuYXBfaHlwZXJsaW5rIiAmJiBib2R5PSJRTkFQIFN5c3RlbXMsIEluYy48L2EgPiBBbGwgUmlnaHRzIFJlc2VydmVkLiIp) |
| **Number of assets affected**  | 2262781 |
| **Description**  | QNAP Systems QTS is an operating system used by China's QNAP Systems for entry-level to mid-level QNAP NAS. There is a security vulnerability in QNAP Systems QTS. The vulnerability stems from the fact that devices running QuTS hero and QTS allow remote attackers to inject malicious code into the app_token parameter field to obtain server permissions. |
| **Impact** | There is a security vulnerability in QNAP Systems QTS. The vulnerability stems from the fact that devices running QuTS hero and QTS allow remote attackers to inject malicious code into the app_token parameter field to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/04/04/5bfa9b242ae05f6c.gif)

## Zyxel Authentication Bypass Vulnerability (CVE-2022-0342)

|   **Vulnerability**  | **Zyxel Authentication Bypass Vulnerability (CVE-2022-0342)**  |
| :----:   | :-----|
|  **Chinese name**  | Zyxel 认证绕过漏洞 (CVE-2022-0342) |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="/2FA-access.cgi" && body="zyxel zyxel_style1"](https://fofa.info/result?qbase64=Ym9keT0iLzJGQS1hY2Nlc3MuY2dpIiAmJiBib2R5PSJ6eXhlbCB6eXhlbF9zdHlsZTEi) |
| **Number of assets affected**  | 6806 |
| **Description**  | Zyxel USG/ZyWALL is a firewall of China Zyxel Technology (Zyxel). Zyxel USG/ZyWALL 4.20 to 4.70, USG FLEX 4.50 to 5.20, ATP 4.32 to 5.20, VPN 4.30 to 5.20, NSG 1.20 to 1.33 Patch 4 have security vulnerabilities, which can be exploited by attackers to circumvent Authenticate over the web and gain administrative access to the device. |
| **Impact** | Attackers can control the entire system through unauthorized access vulnerabilities, and ultimately lead to an extremely insecure state of the system. |

![](https://s3.bmp.ovh/imgs/2023/04/01/ca2c23fcafe64c05.gif)

## PbootCMS 3.1.2 RCE (CVE-2022-32417)

|   **Vulnerability**  | **PbootCMS 3.1.2 RCE (CVE-2022-32417)**  |
| :----:   | :-----|
|  **Chinese name**  | PbootCMS 3.1.2 远程代码执行漏洞（CVE-2022-32417） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [banner="Set-Cookie: pbootsystem=" \|\| header="Set-Cookie: pbootsystem=" \|\| title="PbootCMS"](https://fofa.info/result?qbase64=CmJhbm5lcj0iU2V0LUNvb2tpZTogcGJvb3RzeXN0ZW09IiB8fCBoZWFkZXI9IlNldC1Db29raWU6IHBib290c3lzdGVtPSIgfHwgdGl0bGU9IlBib290Q01TIg%3D%3D) |
| **Number of assets affected**  | 144504 |
| **Description**  | PbootCMS is an open source enterprise website content management system (CMS) developed by PbootCMS personal developers using PHP language. There is a security vulnerability in PbootCMS version 3.1.2, through which an attacker can cause remote code execution. |
| **Impact** | There is a security vulnerability in PbootCMS version 3.1.2, through which an attacker can cause remote code execution. |

![](https://s3.bmp.ovh/imgs/2023/04/01/47ea94dc539d5cf8.gif)


## Weblogic ForeignOpaqueReference  Remote Code Execution Vulnerability (CVE-2023-21839)

|   **Vulnerability**  | **Weblogic ForeignOpaqueReference  Remote Code Execution Vulnerability (CVE-2023-21839)**  |
| :----:   | :-----|
|  **Chinese name**  | Weblogic ForeignOpaqueReference 反序列化远程代码执行漏洞（CVE-2023-21839） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="Welcome to WebLogic Server") \|\| (title=="Error 404--Not Found") \|\| (((body="\<h1\>BEA WebLogic Server" \|\| server="Weblogic" \|\| body="content=\"WebLogic Server" \|\| body="\<h1\>Welcome to Weblogic Application" \|\| body="\<h1\>BEA WebLogic Server") && header!="couchdb" && header!="boa" && header!="RouterOS" && header!="X-Generator: Drupal") \|\| (banner="Weblogic" && banner!="couchdb" && banner!="drupal" && banner!=" Apache,Tomcat,Jboss" && banner!="ReeCam IP Camera" && banner!="\<h2\>Blog Comments\</h2\>")) \|\| (port="7001" && protocol=="weblogic")](https://fofa.info/result?qbase64=Cihib2R5PSJXZWxjb21lIHRvIFdlYkxvZ2ljIFNlcnZlciIpfHwodGl0bGU9PSJFcnJvciA0MDQtLU5vdCBGb3VuZCIpIHx8ICgoKGJvZHk9IjxoMT5CRUEgV2ViTG9naWMgU2VydmVyIiB8fCBzZXJ2ZXI9IldlYmxvZ2ljIiB8fCBib2R5PSJjb250ZW50PVwiV2ViTG9naWMgU2VydmVyIiB8fCBib2R5PSI8aDE%2BV2VsY29tZSB0byBXZWJsb2dpYyBBcHBsaWNhdGlvbiIgfHwgYm9keT0iPGgxPkJFQSBXZWJMb2dpYyBTZXJ2ZXIiKSAmJiBoZWFkZXIhPSJjb3VjaGRiIiAmJiBoZWFkZXIhPSJib2EiICYmIGhlYWRlciE9IlJvdXRlck9TIiAmJiBoZWFkZXIhPSJYLUdlbmVyYXRvcjogRHJ1cGFsIikgfHwgKGJhbm5lcj0iV2VibG9naWMiICYmIGJhbm5lciE9ImNvdWNoZGIiICYmIGJhbm5lciE9ImRydXBhbCIgJiYgYmFubmVyIT0iIEFwYWNoZSxUb21jYXQsSmJvc3MiICYmIGJhbm5lciE9IlJlZUNhbSBJUCBDYW1lcmEiICYmIGJhbm5lciE9IjxoMj5CbG9nIENvbW1lbnRzPC9oMj4iKSkgfHwgKHBvcnQ9IjcwMDEiICYmIHByb3RvY29sPT0id2VibG9naWMiKQ%3D%3D) |
| **Number of assets affected**  | 128502 |
| **Description**  | WebLogic Server is one of the application server components applicable to cloud and traditional environments. WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution. |
| **Impact** | WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution. |

![](https://s3.bmp.ovh/imgs/2023/04/01/67b39bfe7311567f.gif)

## PHICOMM FIR302B management.cgi RCE (CVE-2022-27373)

|   **Vulnerability**  | **PHICOMM FIR302B management.cgi RCE (CVE-2022-27373)**  |
| :----:   | :-----|
|  **Chinese name**  | 斐讯 FIR302B management.cgi 远程命令执行漏洞 (CVE-2022-27373) |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [title="FIR302B"](https://fofa.info/result?qbase64=dGl0bGU9IkZJUjMwMkIi) |
| **Number of assets affected**  | 14766 |
| **Description**  | phicomm Feixun fir302b is a router of Shanghai Feixun Data Communication Technology Co., Ltd. (phicomm), China. Feixun fir302b has a security vulnerability that stems from the discovery of a Remote Command Execution (RCE) vulnerability through the Ping function. |
| **Impact** | Feixun fir302b has a security vulnerability that stems from the discovery of a Remote Command Execution (RCE) vulnerability through the Ping function. |

![](https://s3.bmp.ovh/imgs/2023/04/01/3a6df3f44e86bfc0.gif)

## Atlassian Confluence Default Login (CVE-2022-26138)

|   **Vulnerability**  | **Atlassian Confluence Default Login (CVE-2022-26138)**  |
| :----:   | :-----|
|  **Chinese name**  | Atlassian Confluence 硬编码用户登陆漏洞 (CVE-2022-26138) |
| **CVSS core**  | 7.0 |
| **FOFA Query**  (click to view the results directly)| [(header="X-Confluence-" && header!="TP-LINK Router UPnP") \|\| (banner="X-Confluence-" && banner!="TP-LINK Router UPnP") \|\| (body="name=\"confluence-base-url\"" && body="id=\"com-atlassian-confluence") \|\| title="Atlassian Confluence" \|\| (title=="Errors" && body="Confluence")](https://fofa.info/result?qbase64=CihoZWFkZXI9IlgtQ29uZmx1ZW5jZS0iICYmIGhlYWRlciE9IlRQLUxJTksgUm91dGVyIFVQblAiKSB8fCAoYmFubmVyPSJYLUNvbmZsdWVuY2UtIiAmJiBiYW5uZXIhPSJUUC1MSU5LIFJvdXRlciBVUG5QIikgfHwgKGJvZHk9Im5hbWU9XCJjb25mbHVlbmNlLWJhc2UtdXJsXCIiICYmIGJvZHk9ImlkPVwiY29tLWF0bGFzc2lhbi1jb25mbHVlbmNlIikgfHwgdGl0bGU9IkF0bGFzc2lhbiBDb25mbHVlbmNlIiB8fCAodGl0bGU9PSJFcnJvcnMiICYmIGJvZHk9IkNvbmZsdWVuY2UiKQ%3D%3D) |
| **Number of assets affected**  | 90658 |
| **Description**  | Atlassian Confluence Server is a server version of Atlassian's collaboration software with enterprise knowledge management functions and support for building enterprise WiKi. A security vulnerability exists in Atlassian Confluence Server, which stems from the use of hard-coded passwords that allow attackers to log in to view sensitive information such as team space members. |
| **Impact** | A security vulnerability exists in Atlassian Confluence Server, which stems from the use of hard-coded passwords that allow attackers to log in to view sensitive information such as team space members. |

![](https://s3.bmp.ovh/imgs/2023/04/01/67b39bfe7311567f.gif)

## Jira Server SSRF (CVE-2022-26135)

|   **Vulnerability**  | **Jira Server SSRF (CVE-2022-26135)**  |
| :----:   | :-----|
|  **Chinese name**  | Jira Server 服务端请求伪造 (CVE-2022-26135) |
| **CVSS core**  | 7.0 |
| **FOFA Query**  (click to view the results directly)| [body="Signup!default.jspa"](https://fofa.info/result?qbase64=Ym9keT0iU2lnbnVwIWRlZmF1bHQuanNwYSI%3D) |
| **Number of assets affected**  | 4586 |
| **Description**  | Atlassian JIRA Server is a server version of a defect tracking management system developed by Atlassian in Australia. The system is mainly used to track and manage various problems and defects in the work. A security vulnerability exists in Atlassian Jira Server. An attacker exploits this vulnerability to perform a server-side request forgery attack via a batch endpoint. |
| **Impact** | A security vulnerability exists in Atlassian Jira Server. An attacker exploits this vulnerability to perform a server-side request forgery attack via a batch endpoint. |

![](https://s3.bmp.ovh/imgs/2023/04/01/492aaf83b98a7363.gif)

## QVIS-NVR Camera Management System RCE (CVE-2021-41419)

|   **Vulnerability**  | **QVIS-NVR Camera Management System RCE (CVE-2021-41419)**  |
| :----:   | :-----|
|  **Chinese name**  | QVIS-NVR Camera Management System JSF 反序列化漏洞（CVE-2021-41419）|
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="qvisBase.js"](https://fofa.info/result?qbase64=Ym9keT0icXZpc0Jhc2UuanMi) |
| **Number of assets affected**  | 1801 |
| **Description**  | QVIS NVR Camera Management System is a monitoring system of QVIS company. A security vulnerability exists in the QVIS NVR Camera Management System due to vulnerability to remote code execution via Java deserialization. |
| **Impact** | A security vulnerability exists in the QVIS NVR Camera Management System due to vulnerability to remote code execution via Java deserialization. |

![](https://s3.bmp.ovh/imgs/2023/04/01/1f12fd865122acb3.gif)


## Apache OFBiz xmlrpc Deserialization Vulnerability (CVE-2020-9496) 

|   **Vulnerability**  | **Apache OFBiz xmlrpc Deserialization Vulnerability (CVE-2020-9496)**  |
| :----:   | :-----|
|  **Chinese name**  | Apache OFBiz xmlrpc 反序列化漏洞 (CVE-2020-9496) |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [cert="Organizational Unit: Apache OFBiz" \|\| (body="www.ofbiz.org" && body="/images/ofbiz_powered.gif")](https://fofa.info/result?qbase64=CmNlcnQ9Ik9yZ2FuaXphdGlvbmFsIFVuaXQ6IEFwYWNoZSBPRkJpeiIgfHwgKGJvZHk9Ind3dy5vZmJpei5vcmciICYmIGJvZHk9Ii9pbWFnZXMvb2ZiaXpfcG93ZXJlZC5naWYiKQ%3D%3D) |
| **Number of assets affected**  | 1226 |
| **Description**  | Apache OFBiz is a suite of business applications flexible enough to be used across any industry. A common architecture allows developers to easily extend or enhance it to create custom features. |
| **Impact** | There is a deserialization vulnerability in the Apache OFBiz SOAPService processing interface. An attacker can obtain server privileges by sending specially constructed deserialized data, executing arbitrary code on the target server, executing system commands, or entering the memory horse. |

![](https://s3.bmp.ovh/imgs/2023/04/01/15ed5f12c7b11753.gif)

## Liferay Portal RCE (CVE-2019-16891)  

|   **Vulnerability**  | **Liferay Portal RCE (CVE-2019-16891)**  |
| :----:   | :-----|
|  **Chinese name**  | Liferay Portal 远程代码执行漏洞（CVE-2019-16891） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="Powered by Liferay Portal" \|\| header="Liferay Portal" \|\| banner="Liferay Portal" \|\| header="guest_language_id=" \|\| banner="guest_language_id=" \|\| body="Liferay.AUI" \|\| body="Liferay.currentURL"](https://fofa.info/result?qbase64=Ym9keT0iUG93ZXJlZCBieSBMaWZlcmF5IFBvcnRhbCIgfHwgaGVhZGVyPSJMaWZlcmF5IFBvcnRhbCIgfHwgYmFubmVyPSJMaWZlcmF5IFBvcnRhbCIgfHwgaGVhZGVyPSJndWVzdF9sYW5ndWFnZV9pZD0iIHx8IGJhbm5lcj0iZ3Vlc3RfbGFuZ3VhZ2VfaWQ9IiB8fCBib2R5PSJMaWZlcmF5LkFVSSIgfHwgYm9keT0iTGlmZXJheS5jdXJyZW50VVJMIg%3D%3D) |
| **Number of assets affected**  | 144504 |
| **Description**  | Liferay Portal is a J2EE-based portal solution developed by American Liferay Company. The solution uses technologies such as EJB and JMS, and can be used as Web publishing and shared workspace, enterprise collaboration platform, social network, etc. A code issue vulnerability exists in Liferay Portal CE version 6.2.5. This vulnerability stems from improper design or implementation problems in the code development process of network systems or products. |
| **Impact** | A code issue vulnerability exists in Liferay Portal CE version 6.2.5. This vulnerability stems from improper design or implementation problems in the code development process of network systems or products. |

![](https://s3.bmp.ovh/imgs/2023/04/01/267b1226161fa8b5.gif)

## WordPress Plugin BackupBuddy Arbitrary File Read Vulnerability (CVE-2022-31474)

|   **Vulnerability**  | **WordPress Plugin BackupBuddy Arbitrary File Read Vulnerability (CVE-2022-31474)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress BackupBuddy 插件 local-download 参数任意文件读取漏洞（CVE-2022-31474） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [header="WordPress" \|\| header="api.w.org" \|\| body="/wp-content/themes/"](https://fofa.info/result?qbase64=aGVhZGVyPSJXb3JkUHJlc3MiIHx8IGhlYWRlcj0iYXBpLncub3JnIiB8fCBib2R5PSIvd3AtY29udGVudC90aGVtZXMvIg%3D%3D) |
| **Number of assets affected**  | 34049801 |
| **Description**  | WordPress BackupBuddy plugin is a fast and simple plugin for WordPress backup and restore. WordPress plugin BackupBuddy versions 8.5.8.0 to 8.7.4.1 have an information disclosure vulnerability, which stems from an arbitrary file read and download vulnerability. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

![](https://s3.bmp.ovh/imgs/2023/04/01/0b9fb8da4ab364e1.gif)

## Zyxel Authentication Bypass Vulnerability (CVE-2022-0342)

|   **Vulnerability**  | **Zyxel Authentication Bypass Vulnerability (CVE-2022-0342)**  |
| :----:   | :-----|
|  **Chinese name**  | Zyxel 认证绕过漏洞 (CVE-2022-0342)  |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="/2FA-access.cgi" && body="zyxel zyxel_style1"](https://fofa.info/result?qbase64=Ym9keT0iLzJGQS1hY2Nlc3MuY2dpIiAmJiBib2R5PSJ6eXhlbCB6eXhlbF9zdHlsZTEi) |
| **Number of assets affected**  | 6806 |
| **Description**  | Zyxel USG/ZyWALL is a firewall of China Zyxel Technology (Zyxel). |
| **Impact** | Attackers can control the entire system through unauthorized access vulnerabilities, and ultimately lead to an extremely insecure state of the system. |

![](https://s3.bmp.ovh/imgs/2023/03/31/ff8f9f2124edc110.gif)

## Zyxel Path Traversal Vulnerability (CVE-2022-2030)

|   **Vulnerability**  | **Zyxel Path Traversal Vulnerability (CVE-2022-2030)**  |
| :----:   | :-----|
|  **Chinese name**  |Zyxel 路径遍历漏洞 (CVE-2022-2030)  |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="/2FA-access.cgi" && body="zyxel zyxel_style1"](https://fofa.info/result?qbase64=Ym9keT0iLzJGQS1hY2Nlc3MuY2dpIiAmJiBib2R5PSJ6eXhlbCB6eXhlbF9zdHlsZTEi) |
| **Number of assets affected**  | 6860 |
| **Description**  | Zyxel USG FLEX is a firewall from China's Zyxel Technology (Zyxel). Offers flexible VPN options (IPsec, SSL or L2TP) to provide flexible and secure remote access for remote work and management.A security vulnerability in Zyxel products stems from a directory traversal vulnerability found in some CGI programs caused by improper handling of specific character sequences in URLs, combined with vulnerability cve-2022-0342 that could allow an unauthenticated attacker to access vulnerable Attack some restricted files on the device. The following products and versions are affected: Zyxel USG FLEX 100(W) firmware version 4.50 to 5.30, USG FLEX 200 firmware version 4.50 to 5.30, USG FLEX 500 firmware version 4.50 to 5.30, USG FLEX 700 firmware version 4.50 to 5.30, USG FLEX 50 (W) firmware version 4.16 to 5.30, USG20(W)-VPN firmware version 4.16 to 5.30, ATP series firmware version 4.32 to 5.30, VPN series firmware version 4.30 to 5.30, USG/ZyWALL series firmware version 4.11 to 4.72. |
| **Impact** | Attackers can control the entire system through unauthorized access vulnerabilities, and ultimately lead to an extremely insecure state of the system. |

![](https://s3.bmp.ovh/imgs/2023/03/31/5f6fb3131fb049d7.gif)

## Smartbi DB2 JDBC Arbitrary Code Execution Vulnerability

|   **Vulnerability**  | **Smartbi DB2 JDBC Arbitrary Code Execution Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  |Smartbi DB2 JDBC 任意代码执行漏洞  |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Smartbi"](https://fofa.info/result?qbase64=KGJvZHk9ImdjZnV0aWwgPSBqc2xvYWRlci5yZXNvbHZlKCdzbWFydGJpLmdjZi5nY2Z1dGlsJykiKSB8fCBib2R5PSJnY2Z1dGlsID0ganNsb2FkZXIucmVzb2x2ZSgnc21hcnRiaS5nY2YuZ2NmdXRpbCcpIg%3D%3D) |
| **Number of assets affected**  | 291 |
| **Description**  | Smartbi is a business intelligence BI software launched by Smart Software, which meets the development stage of BI products. |
| **Impact** | There is an unauthorized access background interface vulnerability between Smartbi V7 and V10.5.8. Combining DB2 JDBC exploitation and bypassing defense checks can lead to JNDI injection vulnerabilities, executing arbitrary code, and obtaining server privileges. |

![](https://s3.bmp.ovh/imgs/2023/03/30/8258465c5b97a719.gif)

## Joomla Web Api Unauthorized Access

|   **Vulnerability**  | **Joomla Web Api Unauthorized Access**  |
| :----:   | :-----|
|  **Chinese name**  |Joomla Web Api接口未授权访问  |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [app="Joomla"](https://fofa.info/result?qbase64=YXBwPSJKb29tbGEi) |
| **Number of assets affected**  | 747187 |
| **Description**  | Attackers can obtain the passwords of MySQL database accounts through unauthorized access vulnerabilities, resulting in sensitive data leakage, and ultimately the system is in an extremely insecure state. |
| **Impact** | Attackers can obtain the passwords of MySQL database accounts through unauthorized access vulnerabilities, resulting in sensitive data leakage, and ultimately the system is in an extremely insecure state. |


## MeterSphere files File Read Vulnerability(CVE-2023-25573)

|   **Vulnerability**  | **MeterSphere files File Read Vulnerability(CVE-2023-25573)**  |
| :----:   | :-----|
|  **Chinese name**  |metersphere 平台 files 文件读取漏洞（CVE-2023-25573）  |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [app="FIT2CLOUD-MeterSphere"](https://fofa.info/result?qbase64=YXBwPSJGSVQyQ0xPVUQtTWV0ZXJTcGhlcmUi) |
| **Number of assets affected**  | 2574 |
| **Description**  | MeterSphere is a one-stop open source continuous testing platform, covering functions such as test tracking, interface testing, UI testing and performance testing, and is fully compatible with mainstream open source standards such as JMeter and Selenium.</p><p>MeterSphere has an unauthorized arbitrary file read vulnerability. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

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

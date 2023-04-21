# Goby History Update Vulnerability Total Document (Continuously Update) 
The following content is an updated vulnerability from Goby. Some of the vulnerabilities are recorded on the screen for easy viewing.

**Updated document date: April 19, 2023** 

## Weblogic ForeignOpaqueReference Remote Code Execution Vulnerability (CVE-2023-21979)

|   **Vulnerability**  | **Weblogic ForeignOpaqueReference Remote Code Execution Vulnerability (CVE-2023-21979)**  |
| :----:   | :-----|
|  **Chinese name**  | Weblogic ForeignOpaqueReference 反序列化远程代码执行漏洞（CVE-2023-21979） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="Welcome to WebLogic Server") \|\| (title=="Error 404--Not Found") \|\| (((body="\<h1>BEA WebLogic Server" \|\| server="Weblogic" \|\| body="content=\"WebLogic Server" \|\| body="\<h1>Welcome to Weblogic Application" \|\| body="\<h1>BEA WebLogic Server") && header!="couchdb" && header!="boa" && header!="RouterOS" && header!="X-Generator: Drupal") \|\| (banner="Weblogic" && banner!="couchdb" && banner!="drupal" && banner!=" Apache,Tomcat,Jboss" && banner!="ReeCam IP Camera" && banner!="\<h2>Blog Comments</h2>")) \|\| (port="7001" && protocol=="weblogic")](https://en.fofa.info/result?qbase64=KGJvZHk9IldlbGNvbWUgdG8gV2ViTG9naWMgU2VydmVyIikgfHwgKHRpdGxlPT0iRXJyb3IgNDA0LS1Ob3QgRm91bmQiKSB8fCAoKChib2R5PSI8aDE%2BQkVBIFdlYkxvZ2ljIFNlcnZlciIgfHwgc2VydmVyPSJXZWJsb2dpYyIgfHwgYm9keT0iY29udGVudD1cIldlYkxvZ2ljIFNlcnZlciIgfHwgYm9keT0iPGgxPldlbGNvbWUgdG8gV2VibG9naWMgQXBwbGljYXRpb24iIHx8IGJvZHk9IjxoMT5CRUEgV2ViTG9naWMgU2VydmVyIikgJiYgaGVhZGVyIT0iY291Y2hkYiIgJiYgaGVhZGVyIT0iYm9hIiAmJiBoZWFkZXIhPSJSb3V0ZXJPUyIgJiYgaGVhZGVyIT0iWC1HZW5lcmF0b3I6IERydXBhbCIpIHx8IChiYW5uZXI9IldlYmxvZ2ljIiAmJiBiYW5uZXIhPSJjb3VjaGRiIiAmJiBiYW5uZXIhPSJkcnVwYWwiICYmIGJhbm5lciE9IiBBcGFjaGUsVG9tY2F0LEpib3NzIiAmJiBiYW5uZXIhPSJSZWVDYW0gSVAgQ2FtZXJhIiAmJiBiYW5uZXIhPSI8aDI%2BQmxvZyBDb21tZW50czwvaDI%2BIikpIHx8IChwb3J0PSI3MDAxIiAmJiBwcm90b2NvbD09IndlYmxvZ2ljIik%3D) |
| **Number of assets affected**  | 126908 |
| **Description**  | WebLogic Server is one of the application server components applicable to cloud and traditional environments. WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution. |
| **Impact** | WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution. |

![](https://s3.bmp.ovh/imgs/2023/04/21/4471db3f7b0147fa.gif)

## Weblogic LinkRef Deserialization Remote Code Execution Vulnerability (CVE-2023-21931)

|   **Vulnerability**  | **Weblogic LinkRef Deserialization Remote Code Execution Vulnerability (CVE-2023-21931)**  |
| :----:   | :-----|
|  **Chinese name**  | Weblogic LinkRef 反序列化远程代码执行漏洞（CVE-2023-21931） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="Welcome to WebLogic Server") \|\| (title=="Error 404--Not Found") \|\| (((body="\<h1>BEA WebLogic Server" \|\| server="Weblogic" \|\| body="content=\"WebLogic Server" \|\| body="\<h1>Welcome to Weblogic Application" \|\| body="\<h1>BEA WebLogic Server") && header!="couchdb" && header!="boa" && header!="RouterOS" && header!="X-Generator: Drupal") \|\| (banner="Weblogic" && banner!="couchdb" && banner!="drupal" && banner!=" Apache,Tomcat,Jboss" && banner!="ReeCam IP Camera" && banner!="\<h2>Blog Comments</h2>")) \|\| (port="7001" && protocol=="weblogic")](https://en.fofa.info/result?qbase64=KGJvZHk9IldlbGNvbWUgdG8gV2ViTG9naWMgU2VydmVyIikgfHwgKHRpdGxlPT0iRXJyb3IgNDA0LS1Ob3QgRm91bmQiKSB8fCAoKChib2R5PSI8aDE%2BQkVBIFdlYkxvZ2ljIFNlcnZlciIgfHwgc2VydmVyPSJXZWJsb2dpYyIgfHwgYm9keT0iY29udGVudD1cIldlYkxvZ2ljIFNlcnZlciIgfHwgYm9keT0iPGgxPldlbGNvbWUgdG8gV2VibG9naWMgQXBwbGljYXRpb24iIHx8IGJvZHk9IjxoMT5CRUEgV2ViTG9naWMgU2VydmVyIikgJiYgaGVhZGVyIT0iY291Y2hkYiIgJiYgaGVhZGVyIT0iYm9hIiAmJiBoZWFkZXIhPSJSb3V0ZXJPUyIgJiYgaGVhZGVyIT0iWC1HZW5lcmF0b3I6IERydXBhbCIpIHx8IChiYW5uZXI9IldlYmxvZ2ljIiAmJiBiYW5uZXIhPSJjb3VjaGRiIiAmJiBiYW5uZXIhPSJkcnVwYWwiICYmIGJhbm5lciE9IiBBcGFjaGUsVG9tY2F0LEpib3NzIiAmJiBiYW5uZXIhPSJSZWVDYW0gSVAgQ2FtZXJhIiAmJiBiYW5uZXIhPSI8aDI%2BQmxvZyBDb21tZW50czwvaDI%2BIikpIHx8IChwb3J0PSI3MDAxIiAmJiBwcm90b2NvbD09IndlYmxvZ2ljIik%3D) |
| **Number of assets affected**  | 127237 |
| **Description**  | WebLogic Server is one of the application server components for cloud and traditional environments.There is a remote code execution vulnerability in WebLogic, which allows an unauthenticated attacker to access and damage the vulnerable WebLogic Server through the IIOP protocol network. Successful exploitation of the vulnerability can lead to WebLogic Server being taken over by the attacker, resulting in remote code execution. |
| **Impact** | There is a remote code execution vulnerability in WebLogic, which allows an unauthenticated attacker to access and damage the vulnerable WebLogic Server through the IIOP protocol network. Successful exploitation of the vulnerability can lead to WebLogic Server being taken over by the attacker, resulting in remote code execution. |

![](https://s3.bmp.ovh/imgs/2023/04/19/564df1e41195fca2.gif)

## Apache CouchDB Unauthenticated Remote Code Execution Vulnerability (CVE-2022-24706)

|   **Vulnerability**  | **Apache CouchDB Unauthenticated Remote Code Execution Vulnerability (CVE-2022-24706)**  |
| :----:   | :-----|
|  **Chinese name**  | Apache CouchDB 未认证远程代码执行漏洞 (CVE-2022-24706) |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [banner=\"name couchdb at\"](https://en.fofa.info/result?qbase64=YmFubmVyPSJuYW1lIGNvdWNoZGIgYXQi) |
| **Number of assets affected**  | 2817 |
| **Description**  | Apache CouchDB is a document-oriented database system developed by the Apache Foundation using Erlang. An access control error vulnerability existed prior to Apache CouchDB 3.2.2 that stemmed from the ability of an attacker to access an incorrect default installation and gain administrator privileges without authenticating. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/04/17/66ab3b4f04224398.gif)

## ThinkPHP Debug Mode Log Information Disclosure Vulnerability

|   **Vulnerability**  | **ThinkPHP Debug Mode Log Information Disclosure Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | ThinkPHP Debug 模式日志信息泄露漏洞 |
| **CVSS core**  | 5.0 |
| **FOFA Query**  (click to view the results directly)| [(((header=\"thinkphp\" \|\| header=\"think_template\") && header!=\"couchdb\" && header!=\"St: upnp:rootdevice\") \|\| body=\"href=\\\"http://www.thinkphp.cn\\\">ThinkPHP\</a \>\<sup\>\" \|\| ((banner=\"thinkphp\" \|\| banner=\"think_template\") && banner!=\"couchdb\" && banner!=\"St: upnp:rootdevice\") \|\| (body=\"ThinkPHP\" && body=\"internal function\"))](https://en.fofa.info/result?qbase64=KCgoaGVhZGVyPSJ0aGlua3BocCIgfHwgaGVhZGVyPSJ0aGlua190ZW1wbGF0ZSIpICYmIGhlYWRlciE9ImNvdWNoZGIiICYmIGhlYWRlciE9IlN0OiB1cG5wOnJvb3RkZXZpY2UiKSB8fCBib2R5PSJocmVmPVwiaHR0cDovL3d3dy50aGlua3BocC5jblwiPlRoaW5rUEhQPC9hID48c3VwPiIgfHwgKChiYW5uZXI9InRoaW5rcGhwIiB8fCBiYW5uZXI9InRoaW5rX3RlbXBsYXRlIikgJiYgYmFubmVyIT0iY291Y2hkYiIgJiYgYmFubmVyIT0iU3Q6IHVwbnA6cm9vdGRldmljZSIpIHx8IChib2R5PSJUaGlua1BIUCIgJiYgYm9keT0iaW50ZXJuYWwgZnVuY3Rpb24iKSk%3D) |
| **Number of assets affected**  | 680923 |
| **Description**  | env configuration leakage: Attacker can fetch env configuration file in laravel framework 5.5.21 and earlier. CVE-2018-15133: In Laravel Framework through 5.5.40 and 5.6.x through 5.6.29, remote code execution might occur as a result of an unserialize call on a potentially untrusted X-XSRF-TOKEN value. This involves the decrypt method in Illuminate/Encryption/Encrypter.php and PendingBroadcast in gadgetchains/Laravel/RCE/3/chain.php in phpggc. The attacker must know the application key, which normally would never occur, but could happen if the attacker previously had privileged access or successfully accomplished a previous attack. When exploit CVE-2018-15133, you need to input a url path that support POST method. |
| **Impact** | Laravel env configuration leakage |

![](https://s3.bmp.ovh/imgs/2023/04/21/a15d8379c113c7b6.gif)


## Laravel env configuration leakage

|   **Vulnerability**  | **Laravel env configuration leakage**  |
| :----:   | :-----|
|  **Chinese name**  | Laravel 环境配置 信息泄露漏洞（CVE-2018-15133） |
| **CVSS core**  | 8.1 |
| **FOFA Query**  (click to view the results directly)| [(header=\"laravel_session\")](https://en.fofa.info/result?qbase64=KGhlYWRlcj0ibGFyYXZlbF9zZXNzaW9uIik%3D) |
| **Number of assets affected**  | 1150688 |
| **Description**  | env configuration leakage: Attacker can fetch env configuration file in laravel framework 5.5.21 and earlier. CVE-2018-15133: In Laravel Framework through 5.5.40 and 5.6.x through 5.6.29, remote code execution might occur as a result of an unserialize call on a potentially untrusted X-XSRF-TOKEN value. This involves the decrypt method in Illuminate/Encryption/Encrypter.php and PendingBroadcast in gadgetchains/Laravel/RCE/3/chain.php in phpggc. The attacker must know the application key, which normally would never occur, but could happen if the attacker previously had privileged access or successfully accomplished a previous attack. When exploit CVE-2018-15133, you need to input a url path that support POST method. |
| **Impact** | Laravel env configuration leakage |

## Nostromo nhttpd RCE (CVE-2019-16278)

|   **Vulnerability**  | **Nostromo nhttpd RCE (CVE-2019-16278)**  |
| :----:   | :-----|
|  **Chinese name**  | Nostromo nhttpd远程代码执行漏洞（CVE-2019-16278） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(header=\"Server: nostromo\" \|\| banner=\"Server: nostromo \")](https://en.fofa.info/result?qbase64=KGhlYWRlcj0iU2VydmVyOiBub3N0cm9tbyIgfHwgYmFubmVyPSJTZXJ2ZXI6IG5vc3Ryb21vICIp) |
| **Number of assets affected**  | 3737 |
| **Description**  | Directory Traversal in the function http_verify in nostromo nhttpd through 1.9.6 allows an attacker to achieve remote code execution via a crafted HTTP request. |
| **Impact** | Nostromo nhttpd RCE (CVE-2019-16278) |

![](https://s3.bmp.ovh/imgs/2023/04/21/5cc3d5eeb458b766.gif)

## Kibana Unauthorized RCE (CVE-2019-7609)

|   **Vulnerability**  | **Kibana Unauthorized RCE (CVE-2019-7609)**  |
| :----:   | :-----|
|  **Chinese name**  | Kibana 远程代码执行漏洞（CVE-2019-7609） |
| **CVSS core**  | 6.0 |
| **FOFA Query**  (click to view the results directly)| [(title=\"Kibana\" \|\| body=\"kbnVersion\" \|\| (header=\"Kbn-Name: kibana\" && header=\"Kbn-Version\") \|\| (body=\"kibana_dashboard_only_user\" && header=\"Kbn-Name\") \|\| (banner=\"Kbn-Name: kibana\" && banner=\"Kbn-Version\")) \|\| (title=\"Kibana\" \|\| body=\"kbnVersion\" \|\| (header=\"Kbn-Name: kibana\" && header=\"Kbn-Version\") \|\| (body=\"kibana_dashboard_only_user\" && header=\"Kbn-Name\") \|\| (banner=\"Kbn-Name: kibana\" && banner=\"Kbn-Version\"))](https://en.fofa.info/result?qbase64=KHRpdGxlPSJLaWJhbmEiIHx8IGJvZHk9ImtiblZlcnNpb24iIHx8IChoZWFkZXI9Iktibi1OYW1lOiBraWJhbmEiICYmIGhlYWRlcj0iS2JuLVZlcnNpb24iKSB8fCAoYm9keT0ia2liYW5hX2Rhc2hib2FyZF9vbmx5X3VzZXIiICYmIGhlYWRlcj0iS2JuLU5hbWUiKSB8fCAoYmFubmVyPSJLYm4tTmFtZToga2liYW5hIiAmJiBiYW5uZXI9Iktibi1WZXJzaW9uIikpIHx8ICh0aXRsZT0iS2liYW5hIiB8fCBib2R5PSJrYm5WZXJzaW9uIiB8fCAoaGVhZGVyPSJLYm4tTmFtZToga2liYW5hIiAmJiBoZWFkZXI9Iktibi1WZXJzaW9uIikgfHwgKGJvZHk9ImtpYmFuYV9kYXNoYm9hcmRfb25seV91c2VyIiAmJiBoZWFkZXI9Iktibi1OYW1lIikgfHwgKGJhbm5lcj0iS2JuLU5hbWU6IGtpYmFuYSIgJiYgYmFubmVyPSJLYm4tVmVyc2lvbiIpKQ%3D%3D) |
| **Number of assets affected**  | 75087 |
| **Description**  | By exploiting the vulnerability, the attacker can initiate relevant requests to kibana through the JavaScript prototype chain pollution attack in the timelion component, so as to take over the server and execute arbitrary commands on the server |
| **Impact** | Kibana Unauthorized RCE (CVE-2019-7609) |


## MCMS Shiro Deserialization Vulnerability (CVE-2022-22928)

|   **Vulnerability**  | **MCMS Shiro Deserialization Vulnerability (CVE-2022-22928)**  |
| :----:   | :-----|
|  **Chinese name**  | 铭飞 MCMS shiro 反序列化漏洞（CVE-2022-22928） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body=\"铭飞Mcms\" \|\| title=\"铭飞Mcms\"](https://en.fofa.info/result?qbase64=Ym9keT0i6ZOt6aOeTWNtcyIgfHwgdGl0bGU9IumTremjnk1jbXMi) |
| **Number of assets affected**  | 295 |
| **Description**  | Mingfei Mcms is a complete open source J2EE system of Mingfei Technology Co., Ltd. Mingfei Mcms V5 2.2 and earlier versions contain a security vulnerability, which stems from the existence of hard coded Shiro key in the software, which allows attackers to exploit the key and execute arbitrary code. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/04/12/6e4ebece1945ba6f.gif)

## GoAnywhere MFT Deserialization Vulnerability (CVE-2023-0669)

|   **Vulnerability**  | **GoAnywhere MFT Deserialization Vulnerability (CVE-2023-0669)**  |
| :----:   | :-----|
|  **Chinese name**  | GoAnywhere MFT 反序列化漏洞（CVE-2023-0669） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [banner=\".goanywhere.com\" \|\| title=\"GoAnywhere\"](https://en.fofa.info/result?qbase64=YmFubmVyPSIuZ29hbnl3aGVyZS5jb20iIHx8IHRpdGxlPSJHb0FueXdoZXJlIg%3D%3D) |
| **Number of assets affected**  | 4399 |
| **Description**  | GoAnywhere MFT is a solution for managing file transfer, which simplifies data exchange between systems, employees, customers and trading partners. It provides centralized control through extensive security settings, detailed audit trails, and helps to process information in files into XML, EDI, CSV, and JSON databases. There is a Java deserialization vulnerability in GoAnywhere MFT. An attacker can use this vulnerability to execute arbitrary code, execute commands on the server, enter memory horses, etc., and obtain server privileges. |
| **Impact** | There is a Java deserialization vulnerability in GoAnywhere MFT. An attacker can use this vulnerability to execute arbitrary code, execute commands on the server, enter memory horses, etc., and obtain server privileges. |

![](https://s3.bmp.ovh/imgs/2023/04/12/a8aaa327c8938de6.gif)

## ZOHO ManageEngine Password Manager Pro RCE (CVE-2022-35405)

|   **Vulnerability**  | **ZOHO ManageEngine Password Manager Pro RCE (CVE-2022-35405)**  |
| :----:   | :-----|
|  **Chinese name**  | ZOHO ManageEngine Password Manager Pro 远程代码执行漏洞（CVE-2022-35405） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [banner=\"Server: PMP\" \|\| header=\"Server: PMP\" \|\| banner=\"Set-Cookie: pmpcc=\" \|\| header=\"Set-Cookie: pmpcc=\" \|\| title=\"ManageEngine Password Manager Pro\"](https://en.fofa.info/result?qbase64=YmFubmVyPSJTZXJ2ZXI6IFBNUCIgfHwgaGVhZGVyPSJTZXJ2ZXI6IFBNUCIgfHwgYmFubmVyPSJTZXQtQ29va2llOiBwbXBjYz0iIHx8IGhlYWRlcj0iU2V0LUNvb2tpZTogcG1wY2M9IiB8fCB0aXRsZT0iTWFuYWdlRW5naW5lIFBhc3N3b3JkIE1hbmFnZXIgUHJvIg%3D%3D) |
| **Number of assets affected**  | 672 |
| **Description**  | ZOHO ManageEngine Password Manager Pro is a password manager from the American company ZOHO. ZOHO ManageEngine Password Manager Pro versions prior to 12101 and PAM360 prior to 5510 have security vulnerabilities, attackers can execute arbitrary commands to gain server privileges. |
| **Impact** | ZOHO ManageEngine Password Manager Pro versions prior to 12101 and PAM360 prior to 5510 have security vulnerabilities, attackers can execute arbitrary commands to gain server privileges. |

![](https://s3.bmp.ovh/imgs/2023/04/12/8c05d22e66e4bd93.gif)

## WordPress plugin Metform forms Information Disclosure (CVE-2022-1442)

|   **Vulnerability**  | **WordPress plugin Metform forms Information Disclosure (CVE-2022-1442)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Metform 插件 forms 文件信息泄露漏洞（CVE-2022-1442） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body=\"wp-content/plugins/metform/\"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL21ldGZvcm0vIg%3D%3D) |
| **Number of assets affected**  | 13517 |
| **Description**  | WordPress plugin Metform is a secure contact form plugin for WordPress. There is a security vulnerability in the WordPress plugin Metform. The vulnerability is caused by improper access control in the ~/core/forms/action.php file, and attackers can obtain various key information of users. |
| **Impact** | There is a security vulnerability in the WordPress plugin Metform. The vulnerability is caused by improper access control in the ~/core/forms/action.php file, and attackers can obtain various key information of users. |

![](https://s3.bmp.ovh/imgs/2023/04/12/d33ddd786b414472.gif)


## CaiMore Gateway formping file Command Execution Vulnerability

|   **Vulnerability**  | **CaiMore Gateway formping file Command Execution Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 厦门才茂通信科技有限公司网关 formping 文件命令执行漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [banner=\"Basic realm=\\\"CaiMore\" \|\| header=\"Basic realm=\\\"CaiMore\"](https://en.fofa.info/result?qbase64=YmFubmVyPSJCYXNpYyByZWFsbT1cXFwiQ2FpTW9yZSIgfHwgaGVhZGVyPSJCYXNpYyByZWFsbT1cXFwiQ2FpTW9yZSI%3D) |
| **Number of assets affected**  | 1265 |
| **Description**  | The gateway of Xiamen Caimao Communication Technology Co., Ltd. is designed with open software architecture. It is a metal shell design, with two Ethernet RJ45 interfaces, and an industrial design wireless gateway using 3G/4G/5G wide area network for Internet communication. There is a command execution vulnerability in the formping file of the gateway of Xiamen Caimao Communication Technology Co., Ltd. An attacker can use this vulnerability to arbitrarily execute code on the server side, write to the back door, obtain server permissions, and then control the entire web server. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/04/12/7d6fdb188ac503b5.gif)

## Hikvision iSecure Center springboot Information disclosure vulnerability

|   **Vulnerability**  | **Hikvision iSecure Center springboot Information disclosure vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 海康综合安防管理平台系统 springboot 信息泄露漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=\"综合安防管理平台\" && body=\"nginxService/v1/download/InstallRootCert.exe\"](https://en.fofa.info/result?qbase64=dGl0bGU9Iue7vOWQiOWuiemYsueuoeeQhuW5s%2BWPsCIgJiYgYm9keT0ibmdpbnhTZXJ2aWNlL3YxL2Rvd25sb2FkL0luc3RhbGxSb290Q2VydC5leGUi) |
| **Number of assets affected**  | 3095 |
| **Description**  | Hikvision iSecure Center is an integrated management platform, which can centrally manage the access video monitoring points to achieve unified deployment, configuration, management and scheduling. the framework it uses has a spring boot information disclosure vulnerability. An attacker can access the exposed route to obtain information such as environment variables, intranet addresses, and user names in the configuration. |
| **Impact** | Hikvision iSecure Center is a spring boot information disclosure vulnerability. An attacker can access and download the heapdump heap to obtain sensitive information such as the intranet account password. |

![](https://s3.bmp.ovh/imgs/2023/04/13/47c0acd2094e7191.gif)


## Ruiyou Tianyi Application Virtualization System Index.php File Remote Code Execution Vulnerability

|   **Vulnerability**  | **Ruiyou Tianyi Application Virtualization System Index.php File Remote Code Execution Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 瑞友天翼应用虚拟化系统 index.php 文件远程代码执行漏洞 |
| **CVSS core**  | 9.3 |
| **FOFA Query**  (click to view the results directly)| [body="瑞友应用虚拟化系统" \|\| body="CASMain.XGI?cmd=" \|\| body="瑞友天翼－应用虚拟化系统" \|\| body="DownLoad.XGI?pram="](https://en.fofa.info/result?qbase64=Ym9keT0i55Ge5Y%2BL5bqU55So6Jma5ouf5YyW57O757ufIiB8fCBib2R5PSJDQVNNYWluLlhHST9jbWQ9IiB8fCBib2R5PSLnkZ7lj4vlpKnnv7zvvI3lupTnlKjomZrmi5%2FljJbns7vnu58iIHx8IGJvZHk9IkRvd25Mb2FkLlhHST9wcmFtPSI%3D) |
| **Number of assets affected**  | 61711 |
| **Description**  | Ruiyou Tianyi Application Virtualization System Remote Code Execution Vulnerability Intelligence (0day) allows attackers to execute arbitrary code through this vulnerability, resulting in the system being attacked and controlled. The Ruiyou Tianyi Application Virtualization System is an application virtualization platform based on server computing architecture. It centrally deploys various user application software to the Ruiyou Tianyi service cluster, and clients can access authorized application software on the server through the WEB, achieving centralized application, remote access, collaborative office, and more. Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/04/11/ad85ca285f103af4.gif)

## Superdata Software V.NET Struts2 Code Execution Vulnerability

|   **Vulnerability**  | **Superdata Software V.NET Struts2 Code Execution Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 速达软件 V.NET home 文件 存在 Struts2 代码执行漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="速达软件技术（广州）有限公司"](https://en.fofa.info/result?qbase64=Ym9keT0i6YCf6L6%2B6L2v5Lu25oqA5pyv77yI5bm%2F5bee77yJ5pyJ6ZmQ5YWs5Y%2B4Ig%3D%3D) |
| **Number of assets affected**  | 16627 |
| **Description**  | Superdata software management system is a complete set of enterprise business management system, which organically integrates enterprise purchase management, sales management, warehousing management and financial management. It is extremely easy to use and practical, and comprehensively improves enterprise management ability and work efficiency. Many products of superdata software technology (Guangzhou) Co., Ltd. have code execution vulnerabilities. The code does not filter the controllable parameters of the user, leading to the direct introduction of execution commands and codes, the execution of maliciously constructed statements, and the execution of arbitrary commands or codes through the vulnerability. Attackers can execute arbitrary commands, read and write files, etc. on the server, which is very harmful. |
| **Impact** | Because the code does not filter the user controllable parameters, it directly leads to the execution of commands and code, and executes maliciously constructed statements and arbitrary commands or code through vulnerabilities. Attackers can execute arbitrary commands, read and write files, etc. on the server, which is very harmful. |

![](https://s3.bmp.ovh/imgs/2023/04/10/df1506737795d6f4.gif)

## OpenCart So Newsletter Custom Popup 4.0 module email parameter SQL injection vulnerability

|   **Vulnerability**  | **OpenCart So Newsletter Custom Popup 4.0 module email parameter SQL injection vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | OpenCart So Newsletter Custom Popup 4.0 模块 email 参数 SQL 注入漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="extension/module/so_newletter_custom_popup/newsletter"](https://en.fofa.info/result?qbase64=Ym9keT0iZXh0ZW5zaW9uL21vZHVsZS9zb19uZXdsZXR0ZXJfY3VzdG9tX3BvcHVwL25ld3NsZXR0ZXIi) |
| **Number of assets affected**  | 4474 |
| **Description**  | The OpenCart Newsletter Custom Popup module is a module for newsletter subscriptions. There is a SQL injection vulnerability in the email parameter of the extension/module/so_newletter_custom_popup/newsletter interface of the Opencart Newsletter Custom Popup 4.0 module due to improper filtering. |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

![](https://s3.bmp.ovh/imgs/2023/04/12/0092879ad5b9054b.gif)

## WordPress plugin AWP Classifieds SQL injection vulnerability (CVE-2022-3254)

|   **Vulnerability**  | **WordPress plugin AWP Classifieds SQL injection vulnerability (CVE-2022-3254)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress AWP Classifieds 插件 admin-ajax.php 文件 type 参数SQL注入漏洞（CVE-2022-3254） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/another-wordpress-classifieds"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2Fub3RoZXItd29yZHByZXNzLWNsYXNzaWZpZWRzIg%3D%3D) |
| **Number of assets affected**  | 3526 |
| **Description**  | WordPress plugin AWP Classifieds is a leading plug-in that quickly and easily adds classified ads sections to your WordPress website in minutes. WordPress plugin AWP Classifieds has an SQL injection vulnerability prior to 4.3, which is caused by the plugin's inability to escape the type parameter correctly. Attackers can exploit the vulnerability to obtain sensitive information such as user names and passwords. |
| **Impact** | WordPress plugin AWP Classifieds has an SQL injection vulnerability prior to 4.3, which is caused by the plugin's inability to escape the type parameter correctly. Attackers can exploit the vulnerability to obtain sensitive information such as user names and passwords. |

![](https://s3.bmp.ovh/imgs/2023/04/13/812af92728d9dd9a.gif)

## GetSimpleCMS theme-edit.php content Arbitrary code execution vulnerability (CVE-2022-41544)

|   **Vulnerability**  | **GetSimpleCMS theme-edit.php content Arbitrary code execution vulnerability (CVE-2022-41544)**  |
| :----:   | :-----|
|  **Chinese name**  | GetSimpleCMS 内容管理系统 theme-edit.php 文件 content 参数任意代码执行漏洞（CVE-2022-41544） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(body=\"content=\\\"GetSimple\" \|\| body=\"Powered by GetSimple\")](https://en.fofa.info/result?qbase64=KGJvZHk9ImNvbnRlbnQ9XFxcIkdldFNpbXBsZSIgfHwgYm9keT0iUG93ZXJlZCBieSBHZXRTaW1wbGUiKQ%3D%3D) |
| **Number of assets affected**  | 2784 |
| **Description**  | GetSimple CMS is a content management system (CMS) written in the PHP language. GetSimple CMS v3.3.16 has a security vulnerability that stems from the discovery of the remote Code execution (RCE) vulnerability through the edited_file parameter in admin/theme-edit.php. |
| **Impact** | GetSimple CMS v3.3.16 has a security vulnerability that stems from the discovery of the remote Code execution (RCE) vulnerability through the edited_file parameter in admin/theme-edit.php. |

![](https://s3.bmp.ovh/imgs/2023/04/13/fc08fe3813440052.gif)


## NUUO NVR __debugging_center_utils___.php Command Execution

|   **Vulnerability**  | **NUUO NVR __debugging_center_utils___.php Command Execution**  |
| :----:   | :-----|
|  **Chinese name**  |   NUUO NVR 摄像机 __debugging_center_utils___.php 命令执行漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [title="Network Video Recorder Login"](https://en.fofa.info/result?qbase64=dGl0bGU9Ik5ldHdvcmsgVmlkZW8gUmVjb3JkZXIgTG9naW4i) |
| **Number of assets affected**  | 10146 |
| **Description**  | The NUUO NVR video storage management device __debugging_center_utils___.php has an unauthorized remote command execution vulnerability. An attacker can execute arbitrary system commands without any permissions, thereby invading the server and obtaining administrator permissions on the server. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.< |

![](https://s3.bmp.ovh/imgs/2023/04/07/c18c326c391b4092.gif)

## Grafana welcome Arbitrary File Reading Vulnerability

|   **Vulnerability**  | **Grafana welcome Arbitrary File Reading Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  |   Grafana 网络应用程序平台 welcome 任意文件读取漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [app="Grafana_Labs-公司产品"](https://en.fofa.info/result?qbase64=YXBwPSJHcmFmYW5hX0xhYnMt5YWs5Y%2B45Lqn5ZOBIg%3D%3D) |
| **Number of assets affected**  | 369673 |
| **Description**  | Grafana is a cross-platform, open source platform for data visualization web applications. After users configure the connected data source, Grafana can display data graphs and warnings in a Web browser. Unauthorized attackers can exploit this vulnerability and gain access to sensitive server files. |
| **Impact** | Grafana can display graphs and warnings in a Web browser. Unauthorized attackers can exploit this vulnerability and gain access to sensitive server files. |

![](https://s3.bmp.ovh/imgs/2023/04/07/ac7eb471dfe138dc.gif)

## MeterSphere File Read Vulnerability(CVE-2023-25814)

|   **Vulnerability**  | **MeterSphere File Read Vulnerability(CVE-2023-25814)**  |
| :----:   | :-----|
|  **Chinese name**  |    MeterSphere 文件读取漏洞（CVE-2023-25814） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [app="FIT2CLOUD-MeterSphere"](https://en.fofa.info/result?qbase64=YXBwPSJGSVQyQ0xPVUQtTWV0ZXJTcGhlcmUi) |
| **Number of assets affected**  | 2552 |
| **Description**  | MeterSphere is a one-stop open source continuous testing platform, covering functions such as test tracking, interface testing, UI testing and performance testing, and is fully compatible with mainstream open source standards such as JMeter and Selenium. MeterSphere has an unauthorized arbitrary file read vulnerability. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

![](https://s3.bmp.ovh/imgs/2023/04/07/4fd8616dc5a1c81c.gif)


## Yonyou NC com.ufsoft.iufo.jiuqi.JiuQiClientReqDispatch Deserialization Command Execution Vulnerability

|   **Vulnerability**  | **Yonyou NC com.ufsoft.iufo.jiuqi.JiuQiClientReqDispatch Deserialization Command Execution Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 		用友NC com.ufsoft.iufo.jiuqi.JiuQiClientReqDispatch 反序列化命令执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Yonyou-UFIDA-NC"](https://en.fofa.info/result?qbase64=YXBwPSJZb255b3UtVUZJREEtTkMi) |
| **Number of assets affected**  | 11642 |
| **Description**  | PlaySMS is a free and open source SMS gateway software. An input validation error vulnerability existed in PlaySMS versions prior to 1.4.3, which was caused by the program not sanitizing malicious strings. An attacker could exploit this vulnerability to execute arbitrary code. |
| **Impact** | An input validation error vulnerability existed in PlaySMS versions prior to 1.4.3, which was caused by the program not sanitizing malicious strings. An attacker could exploit this vulnerability to execute arbitrary code. |

![](https://s3.bmp.ovh/imgs/2023/04/06/05179a798f7fc68a.gif)

## OneThink category method code execution vulnerabilities

|   **Vulnerability**  | **OneThink category method code execution vulnerabilities**  |
| :----:   | :-----|
|  **Chinese name**  | 	OneThink 内容管理框架 category 方法代码执行漏洞 |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [(header="ThinkPHP" && title="Onethink") \|\| body="<a href=\\\"http://www.onethink.cn\\\" target=\\\"_blank\\\">OneThink\</a>" \|\| body="/css/onethink.css"](https://en.fofa.info/result?qbase64=KGhlYWRlcj0iVGhpbmtQSFAiICYmIHRpdGxlPSJPbmV0aGluayIpIHx8IGJvZHk9IjxhIGhyZWY9XFxcImh0dHA6Ly93d3cub25ldGhpbmsuY25cXFwiIHRhcmdldD1cXFwiX2JsYW5rXFxcIj5PbmVUaGluazwvYT4iIHx8IGJvZHk9Ii9jc3Mvb25ldGhpbmsuY3NzIg%3D%3D) |
| **Number of assets affected**  | 2854 |
| **Description**  |  OneThink is an open source content management framework developed by the ThinkPHP team based on ThinkPHP. There is a SQL injection vulnerability in the category parameter of the front-end home/article/index method of the OneThink system v1 version. An attacker can use the vulnerability to jointly query any template, causing the file to be included, and finally obtain server permissions. |
| **Impact** | There is a SQL injection vulnerability in the category parameter of the front-end home/article/index method of the OneThink system v1 version. An attacker can use the vulnerability to jointly query any template, causing the file to be included, and finally obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/04/07/35c9841dac3ce243.gif)

## wavlink mesh.cgi command execution (CVE-2022-2486)

|   **Vulnerability**  | **wavlink mesh.cgi command execution (CVE-2022-2486)**  |
| :----:   | :-----|
|  **Chinese name**  | wavlink mesh.cgi命令执行漏洞（CVE-2022-2486） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="firstFlage"](https://en.fofa.info/result?qbase64=Ym9keT0iZmlyc3RGbGFnZSI%3D) |
| **Number of assets affected**  | 3078 |
| **Description**  | WAVLINK is a router developed by China Ruiyin Technology (WAVLINK) company. The system mesh.cgi file has a command execution vulnerability, and attackers can obtain server privileges through this vulnerability. Including models WN530HG4, WN531G3, WN572HG3, WN535G3, WN575A4, etc. |
| **Impact** | Attackers can use this vulnerability to execute system commands to gain server privileges. |

![](https://s3.bmp.ovh/imgs/2023/04/07/a53061701522ba0d.gif)


## WSO2 API Manager save_artifact_ajaxprocessor.jsp XXE Vulnerability (CVE-2020-24589)

|   **Vulnerability**  | **WSO2 API Manager save_artifact_ajaxprocessor.jsp XXE Vulnerability (CVE-2020-24589)**  |
| :----:   | :-----|
|  **Chinese name**  | WSO2 API Manager 系统 save_artifact_ajaxprocessor.jsp XXE 漏洞（CVE-2020-24589） |
| **CVSS core**  | 9.1 |
| **FOFA Query**  (click to view the results directly)| [title="WSO2" \|\| header="Server: WSO2 Carbon Server" \|\| banner="Server: WSO2 Carbon Server"](https://en.fofa.info/result?qbase64=dGl0bGU9IldTTzIiIHx8IGhlYWRlcj0iU2VydmVyOiBXU08yIENhcmJvbiBTZXJ2ZXIiIHx8IGJhbm5lcj0iU2VydmVyOiBXU08yIENhcmJvbiBTZXJ2ZXIi) |
| **Number of assets affected**  | 15231 |
| **Description**  | WSO2 API Manager is a set of API lifecycle management solutions from WSO2 in the United States. A vulnerability exists in WSO2 API Manager. The following products and versions are affected: WSO2 API Manager from version 3.1.0 and API Microgateway version 2.2.0, the attacker can read arbitrary files and detect intranet information, etc. |
| **Impact** | A vulnerability exists in WSO2 API Manager. The following products and versions are affected: WSO2 API Manager from version 3.1.0 and API Microgateway version 2.2.0, the attacker can read arbitrary files and detect intranet information, etc. |

![](https://s3.bmp.ovh/imgs/2023/04/07/92ab16512332fe0c.gif)


## Liferay Portal Unauthenticated 7.2.1 RCE (CVE-2020-7961)

|   **Vulnerability**  | **Liferay Portal Unauthenticated 7.2.1 RCE (CVE-2020-7961)**  |
| :----:   | :-----|
|  **Chinese name**  | Liferay Portal 7.2.1 版本 invoke 文件远程代码执行漏洞（CVE-2020-7961）） |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [body="Powered by Liferay Portal" \|\| header="Liferay Portal" \|\| banner="Liferay Portal" \|\| header="guest_language_id=" \|\| banner="guest_language_id=" \|\| body="Liferay.AUI" \|\| body="Liferay.currentURL"](https://en.fofa.info/result?qbase64=Ym9keT0iUG93ZXJlZCBieSBMaWZlcmF5IFBvcnRhbCIgfHwgaGVhZGVyPSJMaWZlcmF5IFBvcnRhbCIgfHwgYmFubmVyPSJMaWZlcmF5IFBvcnRhbCIgfHwgaGVhZGVyPSJndWVzdF9sYW5ndWFnZV9pZD0iIHx8IGJhbm5lcj0iZ3Vlc3RfbGFuZ3VhZ2VfaWQ9IiB8fCBib2R5PSJMaWZlcmF5LkFVSSIgfHwgYm9keT0iTGlmZXJheS5jdXJyZW50VVJMIg%3D%3D) |
| **Number of assets affected**  | 59885 |
| **Description**  | Liferay Portal is a set of J2EE-based portal solutions of American Liferay Company. The program uses EJB and JMS and other technologies, and can be used as Web publishing and sharing workspace, enterprise collaboration platform, social network and so on. A code issue vulnerability exists in versions prior to Liferay Portal 7.2.1 CE GA2. A remote attacker could exploit this vulnerability to execute arbitrary code using JSON Web services. |
| **Impact** | A code issue vulnerability exists in versions prior to Liferay Portal 7.2.1 CE GA2. A remote attacker could exploit this vulnerability to execute arbitrary code using JSON Web services. |

![](https://s3.bmp.ovh/imgs/2023/04/07/a16de9eefef6f8a5.gif)


## Apache ShenYu Admin plugin API Unauth Access Vulnerability (CVE-2022-23944)

|   **Vulnerability**  | **Apache ShenYu Admin plugin API Unauth Access Vulnerability (CVE-2022-23944)**  |
| :----:   | :-----|
|  **Chinese name**  | Apache ShenYu Admin plugin 接口未授权访问漏洞（CVE-2022-23944） |
| **CVSS core**  | 9.1 |
| **FOFA Query**  (click to view the results directly)| [body="id=\\\"httpPath\\\" style=\\\"display: none"](https://fofa.info/result?qbase64=Ym9keT0iaWQ9XFxcImh0dHBQYXRoXFxcIiBzdHlsZT1cXFwiZGlzcGxheTogbm9uZSI%3D) |
| **Number of assets affected**  | 74 |
| **Description**  | Apache ShenYu is an asynchronous, high-performance, cross-language, reactive API gateway of the Apache Foundation. Apache ShenYu 2.4.0 and 2.4.1 have an access control error vulnerability that stems from users accessing the /plugin api without authentication. |
| **Impact** | Apache ShenYu 2.4.0 and 2.4.1 have an access control error vulnerability that stems from users accessing the /plugin api without authentication. |

![](https://s3.bmp.ovh/imgs/2023/04/07/7151dc2cc22bed37.gif)


## Microsoft Exchange Server Remote Command Execution Vulnerability (CVE-2021-26857/CVE-2021-26858)

|   **Vulnerability**  | **Microsoft Exchange Server Remote Command Execution Vulnerability (CVE-2021-26857/CVE-2021-26858)**  |
| :----:   | :-----|
|  **Chinese name**  | Microsoft Exchange Server 远程命令执行漏洞（CVE-2021-26857/CVE-2021-26858） |
| **CVSS core**  | 7.8 |
| **FOFA Query**  (click to view the results directly)| [banner="Microsoft ESMTP MAIL Service" \|\| banner="Microsoft Exchange Server" \|\| banner="Microsoft Exchange Internet Mail Service" \|\| banner="Microsoft SMTP MAIL" \|\| banner="Microsoft Exchange" \|\| (banner="owa" && banner="Location" && cert!="Technicolor") \|\| banner="Set-Cookie: OutlookSession" \|\| (((header="owa" && (header="Location" \|\| header="X-Owa-Version" \|\| header="Set-Cookie: OWA-COOKIE")) \|\| (body="href=\\\"/owa/auth/" && (title="Outlook" \|\| title="Exchange " \|\| body="var a_sLgn" \|\| body="aria-label=\\\"Outlook Web App\\\" class=\\\"signInImageHeader"))) && header!="WordPress" && body!="wp-content" && body!="wp-includes") \|\| body="\<!-- owapage = ASP.auth_logon_aspx" \|\| header="x-owa-version" \|\| body="window.location.replace(\\\"/owa/\\\" + window.location.hash);\</script>\</head>\<body>\</body>" \|\| body="\<meta http-equiv=\\\"Refresh\\\" contect=\\\"0;url=/owa\\\">" \|\| body="themes/resources/segoeui-semibold.ttf" \|\| title=="Microsoft Outlook Web Access" \|\| body="aria-label=\\\"Outlook Web App" \|\| title="Outlook Web Access" \|\| header="OutlookSession" \|\| (body=".mouse .owaLogoContainer, .twide .owaLogoContainer" && body="owaLogoContainer") \|\| (body="\<div class=\\\"signInHeader\\\">Outlook\</div>" && body="/owa/") \|\| (body="owapage = ASP.auth_logon_aspx" && body="/owa/" && (body="showPasswordCheck" \|\| body="Outlook")) \|\| (title="Outlook Web App" && body="Microsoft Corporation") \|\| header="realm=\\\"Outlook Web App" \|\| ((body="使用 Outlook Web App " \|\| body=" use Outlook Web App") && body="Microsoft Corporation")](https://fofa.info/result?qbase64=YmFubmVyPSJNaWNyb3NvZnQgRVNNVFAgTUFJTCBTZXJ2aWNlIiB8fCBiYW5uZXI9Ik1pY3Jvc29mdCBFeGNoYW5nZSBTZXJ2ZXIiIHx8IGJhbm5lcj0iTWljcm9zb2Z0IEV4Y2hhbmdlIEludGVybmV0IE1haWwgU2VydmljZSIgfHwgYmFubmVyPSJNaWNyb3NvZnQgU01UUCBNQUlMIiB8fCBiYW5uZXI9Ik1pY3Jvc29mdCBFeGNoYW5nZSIgfHwgKGJhbm5lcj0ib3dhIiAmJiBiYW5uZXI9IkxvY2F0aW9uIiAmJiBjZXJ0IT0iVGVjaG5pY29sb3IiKSB8fCBiYW5uZXI9IlNldC1Db29raWU6IE91dGxvb2tTZXNzaW9uIiB8fCAoKChoZWFkZXI9Im93YSIgJiYgKGhlYWRlcj0iTG9jYXRpb24iIHx8IGhlYWRlcj0iWC1Pd2EtVmVyc2lvbiIgfHwgaGVhZGVyPSJTZXQtQ29va2llOiBPV0EtQ09PS0lFIikpIHx8IChib2R5PSJocmVmPVxcXCIvb3dhL2F1dGgvIiAmJiAodGl0bGU9Ik91dGxvb2siIHx8IHRpdGxlPSJFeGNoYW5nZSAiIHx8IGJvZHk9InZhciBhX3NMZ24iIHx8IGJvZHk9ImFyaWEtbGFiZWw9XFxcIk91dGxvb2sgV2ViIEFwcFxcXCIgY2xhc3M9XFxcInNpZ25JbkltYWdlSGVhZGVyIikpKSAmJiBoZWFkZXIhPSJXb3JkUHJlc3MiICYmIGJvZHkhPSJ3cC1jb250ZW50IiAmJiBib2R5IT0id3AtaW5jbHVkZXMiKSB8fCBib2R5PSI8IS0tIG93YXBhZ2UgPSBBU1AuYXV0aF9sb2dvbl9hc3B4IiB8fCBoZWFkZXI9Ingtb3dhLXZlcnNpb24iIHx8IGJvZHk9IndpbmRvdy5sb2NhdGlvbi5yZXBsYWNlKFxcXCIvb3dhL1xcXCIgKyB3aW5kb3cubG9jYXRpb24uaGFzaCk7PC9zY3JpcHQ%2BPC9oZWFkPjxib2R5PjwvYm9keT4iIHx8IGJvZHk9IjxtZXRhIGh0dHAtZXF1aXY9XFxcIlJlZnJlc2hcXFwiIGNvbnRlY3Q9XFxcIjA7dXJsPS9vd2FcXFwiPiIgfHwgYm9keT0idGhlbWVzL3Jlc291cmNlcy9zZWdvZXVpLXNlbWlib2xkLnR0ZiIgfHwgdGl0bGU9PSJNaWNyb3NvZnQgT3V0bG9vayBXZWIgQWNjZXNzIiB8fCBib2R5PSJhcmlhLWxhYmVsPVxcXCJPdXRsb29rIFdlYiBBcHAiIHx8IHRpdGxlPSJPdXRsb29rIFdlYiBBY2Nlc3MiIHx8IGhlYWRlcj0iT3V0bG9va1Nlc3Npb24iIHx8IChib2R5PSIubW91c2UgLm93YUxvZ29Db250YWluZXIsIC50d2lkZSAub3dhTG9nb0NvbnRhaW5lciIgJiYgYm9keT0ib3dhTG9nb0NvbnRhaW5lciIpIHx8IChib2R5PSI8ZGl2IGNsYXNzPVxcXCJzaWduSW5IZWFkZXJcXFwiPk91dGxvb2s8L2Rpdj4iICYmIGJvZHk9Ii9vd2EvIikgfHwgKGJvZHk9Im93YXBhZ2UgPSBBU1AuYXV0aF9sb2dvbl9hc3B4IiAmJiBib2R5PSIvb3dhLyIgJiYgKGJvZHk9InNob3dQYXNzd29yZENoZWNrIiB8fCBib2R5PSJPdXRsb29rIikpIHx8ICh0aXRsZT0iT3V0bG9vayBXZWIgQXBwIiAmJiBib2R5PSJNaWNyb3NvZnQgQ29ycG9yYXRpb24iKSB8fCBoZWFkZXI9InJlYWxtPVxcXCJPdXRsb29rIFdlYiBBcHAiIHx8ICgoYm9keT0i5L2%2F55SoIE91dGxvb2sgV2ViIEFwcCAiIHx8IGJvZHk9IiB1c2UgT3V0bG9vayBXZWIgQXBwIikgJiYgYm9keT0iTWljcm9zb2Z0IENvcnBvcmF0aW9uIik%3D) |
| **Number of assets affected**  | 2198588 |
| **Description**  | Microsoft Exchange Server is a suite of e-mail services programs from Microsoft Corporation of the United States. It provides mail access, storage, forwarding, voicemail, email filtering and filtering functions. Microsoft Exchange Server has a remote command execution vulnerability. Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | Microsoft Exchange Server has a remote code execution vulnerability. Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/04/07/f5e117459cecb928.gif)


## H2 Database Console login.do Code Execution Vulnerability (CVE-2021-42392)

|   **Vulnerability**  | **H2 Database Console login.do Code Execution Vulnerability (CVE-2021-42392)**  |
| :----:   | :-----|
|  **Chinese name**  | H2 Database 数据库 login.do 文件远程代码执行漏洞 (CVE-2021-42392) |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="login.jsp?jsessionid=" && body="Welcome to H2"](https://fofa.info/result?qbase64=Ym9keT0ibG9naW4uanNwP2pzZXNzaW9uaWQ9IiAmJiBib2R5PSJXZWxjb21lIHRvIEgyIg%3D%3D) |
| **Number of assets affected**  | 488 |
| **Description**  | H2 database is a Java memory database, which is mainly used for unit testing. There is an unauthorized remote code execution vulnerability in the H2 Database Web management page. An attacker can use this vulnerability to arbitrarily execute code on the server side, write to the back door, and obtain server permissions, thereby controlling the entire web server. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/04/07/0a5df04ffd240ed7.gif)

## playSMS 1.4.3 RCE (CVE-2020-8644)

|   **Vulnerability**  | **playSMS 1.4.3 RCE (CVE-2020-8644)**  |
| :----:   | :-----|
|  **Chinese name**  | 	playSMS 1.4.3 远程命令执行漏洞 (CVE-2020-8644) |
| **CVSS core**  | 9.5 |
| **FOFA Query**  (click to view the results directly)| [title=="playSMS"](https://fofa.info/result?qbase64=dGl0bGU9PSJwbGF5U01TIg%3D%3D) |
| **Number of assets affected**  | 722 |
| **Description**  | PlaySMS is a free and open source SMS gateway software. An input validation error vulnerability existed in PlaySMS versions prior to 1.4.3, which was caused by the program not sanitizing malicious strings. An attacker could exploit this vulnerability to execute arbitrary code. |
| **Impact** | An input validation error vulnerability existed in PlaySMS versions prior to 1.4.3, which was caused by the program not sanitizing malicious strings. An attacker could exploit this vulnerability to execute arbitrary code. |

![](https://s3.bmp.ovh/imgs/2023/04/03/70ee3365dd90c1a5.gif)

## YoudianCMS v9.5.0 SQL Injection (CVE-2022-32300)

|   **Vulnerability**  | **YoudianCMS v9.5.0 SQL Injection (CVE-2022-32300)**  |
| :----:   | :-----|
|  **Chinese name**  | YoudianCMS v9.5.0 sql注入（CVE-2022-32300） |
| **CVSS core**  | 8.8 |
| **FOFA Query**  (click to view the results directly)| [body="YoudianCMS"](https://fofa.info/result?qbase64=Ym9keT0iWW91ZGlhbkNNUyI%3D) |
| **Number of assets affected**  | 987 |
| **Description**  | YouDianCMS is a website CMS. YoudianCMS v9.5.0 version exists security holes, the vulnerability stems from a pass/App/Lib/Action/Admin/MailAction class. PHP MailSendID parameters of SQL injection vulnerabilities are found out. |
| **Impact** | Able to read some sensitive files through SQL injection vulnerability. |

![](https://s3.bmp.ovh/imgs/2023/04/06/b6d2916d8bfa2662.gif)

## SolarView Compact downloader.php RCE (CVE-2023-23333)

|   **Vulnerability**  | **SolarView Compact downloader.php RCE (CVE-2023-23333)**  |
| :----:   | :-----|
|  **Chinese name**  | SolarView Compact downloader.php 任意命令执行漏洞（CVE-2023-23333）|
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [body="SolarView Compact"](https://fofa.info/result?qbase64=dGl0bGU9PSJwbGF5U01TIg%3D%3D) |
| **Number of assets affected**  | 5585 |
| **Description**  | There is a command injection vulnerability in SolarView Compact through 6.00, attackers can execute commands by bypassing internal restrictions through downloader.php. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |


![](https://s3.bmp.ovh/imgs/2023/04/03/a32aa1b44858819c.gif)

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

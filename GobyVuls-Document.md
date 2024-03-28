# Goby History Update Vulnerability Total Document (Continuously Update) 
The following content is an updated vulnerability from Goby. Some of the vulnerabilities are recorded on the screen for easy viewing.

**Updated document date: March 28, 2024** 


## 		Adobe ColdFusion /CFIDE/adminapi/_servermanager/servermanager.cfc File Read Vulnerability (CVE-2024-20767)

|   **Vulnerability**  | Adobe ColdFusion /CFIDE/adminapi/_servermanager/servermanager.cfc File Read Vulnerability (CVE-2024-20767) |
| :----:   | :-----|
|  **Chinese name**  |	Adobe ColdFusion /CFIDE/adminapi/_servermanager/servermanager.cfc 文件读取漏洞 (CVE-2024-20767) |
| **CVSS core**  | 8.2 |
| **FOFA Query**  (click to view the results directly)| [app="Adobe-ColdFusion"](https://en.fofa.info/result?qbase64=Ym9keT0iL2NmYWpheC8iIHx8IGhlYWRlcj0iQ0ZUT0tFTiIgfHwgYmFubmVyPSJDRlRPS0VOIiB8fCBib2R5PSJDb2xkRnVzaW9uLkFqYXgiIHx8IGJvZHk9IjxjZnNjcmlwdD4iIHx8IHNlcnZlcj0iQ29sZEZ1c2lvbiIgfHwgdGl0bGU9IkNvbGRGdXNpb24iIHx8IChib2R5PSJjcm9zc2RvbWFpbi54bWwiICYmIGJvZHk9IkNGSURFIikgfHwgKGJvZHk9IiMwMDA4MDgiICYmIGJvZHk9IiNlN2U3ZTciKQ%3D%3D)|
| **Number of assets affected**  | 	504562 |
| **Description**  | Adobe ColdFusion is a commercial application server developed by Adobe for web application development. The vulnerability allows an attacker to read important system files (e.g., database configuration files, system configuration files), database configuration files, etc., resulting in an extremely insecure web site.|
| **Impact** | The vulnerability allows an attacker to read important system files (e.g., database configuration files, system configuration files), database configuration files, etc., resulting in an extremely insecure web site. |

![](https://s3.bmp.ovh/imgs/2024/03/28/d8f5c5bf74cb4017.gif).


## 		JetBrains TeamCity permission bypass vulnerability (CVE-2024-27198 & CVE-2024-27199)

|   **Vulnerability**  | JetBrains TeamCity permission bypass vulnerability (CVE-2024-27198 & CVE-2024-27199)  |
| :----:   | :-----|
|  **Chinese name**  |	JetBrains TeamCity 权限绕过漏洞（CVE-2024-27198 & CVE-2024-27199） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="JET_BRAINS-TeamCity"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJUZWFtY2l0eSIgfHwgYmFubmVyPSJUZWFtY2l0eSIgfHwgdGl0bGU9IlRlYW1DaXR5IiB8fCBib2R5PSJjb250ZW50PVwiVGVhbUNpdHkgKExvZyBpbiB0byBUZWFtQ2l0eSI%3D)|
| **Number of assets affected**  |  141734 |
| **Description**  | JetBrains TeamCity is a continuous integration and continuous delivery (CI/CD) server developed by JetBrains. It provides a powerful platform for automating the building, testing and deployment of software projects. TeamCity aims to simplify team collaboration and software delivery processes, improve development team efficiency and product quality.JetBrains TeamCity has an authentication bypass vulnerability before version 2023.11.4. An attacker can use this vulnerability to bypass the authentication mechanism and directly perform administrator operations. Combined with the background function, the attacker can use this vulnerability to execute arbitrary system commands on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | JetBrains TeamCity has an authentication bypass vulnerability before version 2023.11.4. An attacker can use this vulnerability to bypass the authentication mechanism and directly perform administrator operations. Combined with the background function, the attacker can use this vulnerability to execute arbitrary system commands on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2024/03/21/247110c62504abd3.gif).



## 		Apache Kafka Connect remote code execution vulnerability (CVE-2023-25194)

|   **Vulnerability**  | Apache Kafka Connect remote code execution vulnerability (CVE-2023-25194)  |
| :----:   | :-----|
|  **Chinese name**  |	Apache Druid Kafka Connect 远程代码执行漏洞（CVE-2023-25194） |
| **CVSS core**  | 8.8 |
| **FOFA Query**  (click to view the results directly)| [app="APACHE-Druid"](https://en.fofa.info/result?qbase64=Ym9keT0iQXBhY2hlIERydWlkIGNvbnNvbGUiIHx8IHRpdGxlPSJBcGFjaGUgRHJ1aWQiIHx8IGhlYWRlcj0idW5pZmllZC1jb25zb2xlLmh0bWwiIHx8IGJhbm5lcj0idW5pZmllZC1jb25zb2xlLmh0bWwi)|
| **Number of assets affected**  |2935 |
| **Description**  | 	Apache Druid is an open source distributed data storage and analysis system. It is designed to handle large-scale real-time data and provide fast interactive query and analysis.Apache Druid uses the vulnerable Kafka Connect. An attacker can access the Kafka Connect Worker and create or modify the connector by setting the sasl.jaas.config attribute to a malicious class, which can lead to a JNDI injection vulnerability. This vulnerability can be used Execute code arbitrarily on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | Apache Druid uses the vulnerable Kafka Connect. An attacker can access the Kafka Connect Worker and create or modify the connector by setting the sasl.jaas.config attribute to a malicious class, which can lead to a JNDI injection vulnerability. This vulnerability can be used Execute code arbitrarily on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2024/03/15/7a1efed457cb487c.gif).


## 	ComfyUI follow_symlinks File Read Vulnerability (CVE-2024-23334)

|   **Vulnerability**  | 	ComfyUI follow_symlinks File Read Vulnerability (CVE-2024-23334)  |
| :----:   | :-----|
|  **Chinese name**  | ComfyUI follow_symlinks 文件读取漏洞（CVE-2024-23334） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [app="ComfyUI"](https://en.fofa.info/result?qbase64=dGl0bGU9IkNvbWZ5VUki)|
| **Number of assets affected**  | 1564 |
| **Description**  | 	ComfyUI is a powerful, modular stable diffusion GUI, API, and backend. It provides a graphical/node interface for designing and managing stable diffusion pipelines.ComfyUI uses a low version of aiohttp as a web server and configures static routes with the follow_symlinks option enabled, leading to an arbitrary file read vulnerability. The vulnerability allows an attacker to read leaked source code, database configuration files, etc., resulting in a highly insecure web site. |
| **Impact** | ComfyUI uses a lower version of the aiohttp component as the web server and configures static routing with the follow_symlinks option enabled, resulting in an arbitrary file reading vulnerability. Attackers can use this vulnerability to read leaked source code, database configuration files, etc., causing the website to be in an extremely unsafe state. |

![](https://s3.bmp.ovh/imgs/2024/03/01/2b4bae4fd8ab1926.gif).





## 	WordPress Bricks render_element Remote Code Execution Vulnerability (CVE-2024-25600)

|   **Vulnerability**  | 		WordPress Bricks render_element Remote Code Execution Vulnerability (CVE-2024-25600)  |
| :----:   | :-----|
|  **Chinese name**  | 	WordPress Bricks render_element 远程代码执行漏洞（CVE-2024-25600） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="wordpress-bricks"](https://en.fofa.info/result?qbase64=Ym9keT0iL3dwLWNvbnRlbnQvdGhlbWVzL2JyaWNrcyI%3D)|
| **Number of assets affected**  | 	25433 |
| **Description**  | 		WordPress Bricks is an innovative, community driven, and visual WordPress website builder that allows you to design unique, high-performance, and scalable websites.WordPress Bricks has a remote code execution vulnerability, which allows attackers to execute code arbitrarily on the server side, write backdoors, gain server privileges, and then control the entire web server. |
| **Impact** | WordPress Bricks has a remote code execution vulnerability, which allows attackers to execute code arbitrarily on the server side, write backdoors, gain server privileges, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2024/02/23/70784f1d16f8316e.gif).





## 			Weblogic ForeignOpaqueReference remote code execution vulnerability (CVE-2024-20931)

|   **Vulnerability**  | 		Weblogic ForeignOpaqueReference remote code execution vulnerability (CVE-2024-20931)  |
| :----:   | :-----|
|  **Chinese name**  | 	Weblogic ForeignOpaqueReference 远程代码执行漏洞（CVE-2024-20931） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Weblogic_interface_7001"](https://en.fofa.info/result?qbase64=KGJvZHk9IldlbGNvbWUgdG8gV2ViTG9naWMgU2VydmVyIil8fCh0aXRsZT09IkVycm9yIDQwNC0tTm90IEZvdW5kIikgfHwgKCgoYm9keT0iPGgxPkJFQSBXZWJMb2dpYyBTZXJ2ZXIiIHx8IHNlcnZlcj0iV2VibG9naWMiIHx8IGJvZHk9ImNvbnRlbnQ9XCJXZWJMb2dpYyBTZXJ2ZXIiIHx8IGJvZHk9IjxoMT5XZWxjb21lIHRvIFdlYmxvZ2ljIEFwcGxpY2F0aW9uIiB8fCBib2R5PSI8aDE%2BQkVBIFdlYkxvZ2ljIFNlcnZlciIpICYmIGhlYWRlciE9ImNvdWNoZGIiICYmIGhlYWRlciE9ImJvYSIgJiYgaGVhZGVyIT0iUm91dGVyT1MiICYmIGhlYWRlciE9IlgtR2VuZXJhdG9yOiBEcnVwYWwiKSB8fCAoYmFubmVyPSJXZWJsb2dpYyIgJiYgYmFubmVyIT0iY291Y2hkYiIgJiYgYmFubmVyIT0iZHJ1cGFsIiAmJiBiYW5uZXIhPSIgQXBhY2hlLFRvbWNhdCxKYm9zcyIgJiYgYmFubmVyIT0iUmVlQ2FtIElQIENhbWVyYSIgJiYgYmFubmVyIT0iPGgyPkJsb2cgQ29tbWVudHM8L2gyPiIpKSB8fCAocG9ydD0iNzAwMSIgJiYgcHJvdG9jb2w9PSJ3ZWJsb2dpYyIp)|
| **Number of assets affected**  | 	194125 |
| **Description**  | 		WebLogic Server is one of the application server components suitable for both cloud and traditional environments.WebLogic has a remote code execution vulnerability that allows an unauthenticated attacker to access and destroy a vulnerable WebLogic Server through the IIOP protocol network. Successful exploitation of the vulnerability can cause WebLogic Server to be taken over by an attacker, resulting in remote code execution. |
| **Impact** | There is a remote code execution vulnerability in WebLogic, which allows an unauthenticated attacker to access and damage the vulnerable WebLogic Server through the IIOP protocol network. Successful exploitation of the vulnerability can lead to WebLogic Server being taken over by the attacker, resulting in remote code execution. |

![](https://s3.bmp.ovh/imgs/2024/02/23/70784f1d16f8316e.gif).


## 		Ivanti Connect Secure and Policy Secure saml20.ws server-side request forgery vulnerability (CVE-2024-21893)

|   **Vulnerability**  | 	Ivanti Connect Secure and Policy Secure saml20.ws server-side request forgery vulnerability (CVE-2024-21893)  |
| :----:   | :-----|
|  **Chinese name**  | 		JIvanti Connect Secure 和 Policy Secure saml20.ws 服务端请求伪造漏洞（CVE-2024-21893） |
| **CVSS core**  | 	8.2 |
| **FOFA Query**  (click to view the results directly)| [app="PulseSecure-SSL-VPN"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJEU0Jyb3dzZXJJRCIgfHwgYmFubmVyPSJEU0Jyb3dzZXJJRCIgfHwgYm9keT0iL2RhbmEtbmEvO2V4cGlyZXM9IiB8fCBib2R5PSJkYW5hLWNhY2hlZC9pbWdzL3NwYWNlLmdpZiIgfHwgYm9keT0iL2RhbmEtbmEvaW1ncy9zcGFjZS5naWYiIHx8IGJvZHk9Ii9kYW5hLW5hL2ltZ3MvUHJvZHVjdF9mYXZpY29uLnBuZyIgfHwgYm9keT0iL2RhbmEtbmEvaW1ncy9JdmFudGlfZmF2aWNvbi5wbmciIHx8IGJvZHk9Ii9kYW5hLW5hL2Nzcy9kcy5qcyIgfHwgYm9keT0iZHNfbW9iaWxlX3NhZmFyaS5jc3MiIHx8IGJvZHk9IndlbGNvbWUuY2dpP3A9bG9nbyZzaWduaW5JZD11cmxfZGVmYXVsdCIgfHwgYm9keT0iPGI%2BUHVsc2UgQ29ubmVjdCBTZWN1cmU8L2I%2BIiB8fCB0aXRsZT0iU2VjdXJlIEFjY2VzcyBTU0wgVlBOIiB8fCBiYW5uZXI9Ii9kYW5hLW5hL2F1dGgvdXJsX2RlZmF1bHQiIHx8IGhlYWRlcj0iL2RhbmEtbmEvYXV0aC91cmxfZGVmYXVsdCIgfHwgYm9keT0ic3JjPVwiL2RhbmEtbmEvYXV0aC8i)|
| **Number of assets affected**  | 	399547 |
| **Description**  | 		Ivanti Connect/Policy Secure is a secure remote network connection tool from the American company Ivanti.Ivanti Connect Secure product saml20.ws has a server-side request forgery vulnerability. An attacker can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., causing the website to be in an extremely unsafe state. |
| **Impact** | Ivanti Connect Secure product saml20.ws has a server-side request forgery vulnerability. An attacker can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., causing the website to be in an extremely unsafe state. |

![](https://s3.bmp.ovh/imgs/2024/02/06/b02b7373188c6998.gif).



## 	Jenkins args4j file read vulnerability (CVE-2024-23897)

|   **Vulnerability**  | Jenkins args4j file read vulnerability (CVE-2024-23897)  |
| :----:   | :-----|
|  **Chinese name**  | 		Jenkins args4j 文件读取漏洞（CVE-2024-23897） |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Jenkins"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJYLUplbmtpbnMiIHx8IGJhbm5lcj0iWC1KZW5raW5zIiB8fCBoZWFkZXI9IlgtSHVkc29uIiB8fCBiYW5uZXI9IlgtSHVkc29uIiB8fCBoZWFkZXI9IlgtUmVxdWlyZWQtUGVybWlzc2lvbjogaHVkc29uLm1vZGVsLkh1ZHNvbi5SZWFkIiB8fCBiYW5uZXI9IlgtUmVxdWlyZWQtUGVybWlzc2lvbjogaHVkc29uLm1vZGVsLkh1ZHNvbi5SZWFkIiB8fCBib2R5PSJKZW5raW5zLUFnZW50LVByb3RvY29scyI%3D)|
| **Number of assets affected**  | 		729753 |
| **Description**  | 	CloudBees Jenkins (formerly known as Hudson Labs) is a set of Java-based continuous integration tools developed by American CloudBees Company. It is mainly used to monitor continuous software version release/test projects and some regularly executed tasks.Attackers can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., causing the website to be in an extremely unsafe state. |
| **Impact** | Attackers can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., causing the website to be in an extremely unsafe state. |

![](https://s3.bmp.ovh/imgs/2024/01/26/bb74a2a4f3c0cdbc.gif).




## 	GoAnywhere MFT InitialAccountSetup.xhtml Bypass Vulnerability (CVE-2024-0204)

|   **Vulnerability**  | 	GoAnywhere MFT InitialAccountSetup.xhtml Bypass Vulnerability (CVE-2024-0204)  |
| :----:   | :-----|
|  **Chinese name**  | 		GoAnywhere MFT InitialAccountSetup.xhtml 绕过漏洞（CVE-2024-0204） |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="GoAnywhere-MFT"](https://en.fofa.info/result?qbase64=dGl0bGU9IkdvQW55d2hlcmUiIHx8IGhlYWRlcj0iL2dvYW55d2hlcmUiIHx8IGJhbm5lcj0iL2dvYW55d2hlcmUi)|
| **Number of assets affected**  | 		4468 |
| **Description**  | 	GoAnywhere MFT (Managed File Transfer) is an enterprise-class file transfer solution provided by HelpSystems, designed to meet the needs of organizations for secure, manageable and automated file transfer.Authentication bypass vulnerability in GoAnywhere MFT before Fortra version 7.4.1. Allows unauthorized users to pretend to be administrators through the management portal, create arbitrary management users, and take over the entire system. |
| **Impact** | 	Authentication bypass vulnerability in GoAnywhere MFT before Fortra version 7.4.1. Allows unauthorized users to pretend to be administrators through the management portal, create arbitrary management users, and take over the entire system. |

![](https://s3.bmp.ovh/imgs/2024/01/24/fed22927c8fb9a5f.gif).


## Atlassian Confluence template/aui/text-inline.vm code execution vulnerability (CVE-2023-22527)

|   **Vulnerability**  | 	Atlassian Confluence template/aui/text-inline.vm code execution vulnerability (CVE-2023-22527)  |
| :----:   | :-----|
|  **Chinese name**  | 		Atlassian Confluence template/aui/text-inline.vm 代码执行漏洞（CVE-2023-22527） |
| **CVSS core**  | 	10 |
| **FOFA Query**  (click to view the results directly)| [	app="ATLASSIAN-Confluence"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJYLUNvbmZsdWVuY2UtIiB8fCBiYW5uZXI9IlgtQ29uZmx1ZW5jZS0iIHx8IChib2R5PSJuYW1lPVwiY29uZmx1ZW5jZS1iYXNlLXVybFwiIiAmJiBib2R5PSJpZD1cImNvbS1hdGxhc3NpYW4tY29uZmx1ZW5jZSIpIHx8IHRpdGxlPSJBdGxhc3NpYW4gQ29uZmx1ZW5jZSIgfHwgKHRpdGxlPT0iRXJyb3JzIiAmJiBib2R5PSJDb25mbHVlbmNlIik%3D)|
| **Number of assets affected**  | 			1190821  |
| **Description**  | 	Atlassian Confluence is an enterprise team collaboration and knowledge management software developed by Atlassian that provides a centralized platform for creating, organizing and sharing your team's documents, knowledge base, project plans and collaborative content.Atlassian Confluence has a remote code execution vulnerability in template/aui/text-inline.vm that could allow an unauthorized attacker to execute arbitrary code on an affected version. |
| **Impact** | 	Atlassian Confluence has a remote code execution vulnerability in template/aui/text-inline.vm that could allow an unauthorized attacker to execute arbitrary code on an affected version. |

![](https://s3.bmp.ovh/imgs/2024/01/23/1888fed69ab36f7a.gif).


## Ivanti Connect Secure and Policy Secure keys-status remote command execution vulnerability (CVE-2023-46805/CVE-2024-21887)

|   **Vulnerability**  | 	Ivanti Connect Secure and Policy Secure keys-status remote command execution vulnerability (CVE-2023-46805/CVE-2024-21887)  |
| :----:   | :-----|
|  **Chinese name**  | 	Ivanti Connect Secure 和 Policy Secure keys-status 远程命令执行漏洞（CVE-2023-46805/CVE-2024-21887） |
| **CVSS core**  | 	9.1 |
| **FOFA Query**  (click to view the results directly)| [app="Ivanti Connect Secure"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJEU0Jyb3dzZXJJRCIgfHwgYmFubmVyPSJEU0Jyb3dzZXJJRCIgfHwgYm9keT0iL2RhbmEtbmEvO2V4cGlyZXM9IiB8fCBib2R5PSJkYW5hLWNhY2hlZC9pbWdzL3NwYWNlLmdpZiIgfHwgYm9keT0iL2RhbmEtbmEvaW1ncy9zcGFjZS5naWYiIHx8IGJvZHk9Ii9kYW5hLW5hL2ltZ3MvUHJvZHVjdF9mYXZpY29uLnBuZyIgfHwgYm9keT0iL2RhbmEtbmEvaW1ncy9JdmFudGlfZmF2aWNvbi5wbmciIHx8IGJvZHk9Ii9kYW5hLW5hL2Nzcy9kcy5qcyIgfHwgYm9keT0iZHNfbW9iaWxlX3NhZmFyaS5jc3MiIHx8IGJvZHk9IndlbGNvbWUuY2dpP3A9bG9nbyZzaWduaW5JZD11cmxfZGVmYXVsdCI%3D)|
| **Number of assets affected**  | 		154590 |
| **Description**  | 	Ivanti is a software and information technology services company focused on providing solutions for IT management, security, service management and endpoint management. Ivanti Connect Secure and Ivanti Policy Secure are part of two security solutions from Ivanti for network security and connectivity management.There is a command injection vulnerability in the web component of Ivanti Connect Secure (9.x, 22.x) and Ivanti Policy Secure (9.x, 22.x). The authentication bypass vulnerability and the command injection vulnerability can be used to send specially crafted requests and Execute arbitrary commands on the device. |
| **Impact** | 	which utilizes authentication to bypass the vulnerability and in combination with command injection vulnerabilities, can send customized requests and execute arbitrary commands on the device, gain server privileges, and thus control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/10/10/ec147824884d3597.gif).






## Apache OFBiz webtools/control/ProgramExport remote code execution vulnerability (CVE-2023-51467)

|   **Vulnerability**  | **Apache OFBiz webtools/control/ProgramExport remote code execution vulnerability (CVE-2023-51467)**  |
| :----:   | :-----|
| **Chinese name**  | Apache OFBiz webtools/control/ProgramExport 远程代码执行漏洞（CVE-2023-51467） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Apache_OFBiz"](https://en.fofa.info/result?qbase64=Y2VydD0iT3JnYW5pemF0aW9uYWwgVW5pdDogQXBhY2hlIE9GQml6IiB8fCAoYm9keT0id3d3Lm9mYml6Lm9yZyIgJiYgYm9keT0iL2ltYWdlcy9vZmJpel9wb3dlcmVkLmdpZiIpIHx8IGhlYWRlcj0iU2V0LUNvb2tpZTogT0ZCaXouVmlzaXRvciIgfHwgYmFubmVyPSJTZXQtQ29va2llOiBPRkJpei5WaXNpdG9yIg%3D%3D) |
| **Number of assets affected**  | 	5912 |
| **Description**  |Apache OFBiz is an open source enterprise resource planning (ERP) system that provides a variety of business functions and modules.Apache OFBiz has a code execution vulnerability in webtools/control/ProgramExport. An attacker can use this vulnerability to execute arbitrary code on the server side, write a backdoor, obtain server permissions, and then control the entire web server. |
| **Impact** | 	Apache OFBiz has a code execution vulnerability in webtools/control/ProgramExport. An attacker can use this vulnerability to execute arbitrary code on the server side, write a backdoor, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/12/28/f81ea49af8383f1b.gif).






## Apache OFBiz webtools/control/xmlrpc Remote Code Execution Vulnerability (CVE-2023-49070)

|   **Vulnerability**  | 	Apache OFBiz webtools/control/xmlrpc Remote Code Execution Vulnerability (CVE-2023-49070)  |
| :----:   | :-----|
|  **Chinese name**  | Apache OFBiz webtools/control/xmlrpc 远程代码执行漏洞（CVE-2023-49070） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Apache_OFBiz"](https://en.fofa.info/result?qbase64=Y2VydD0iT3JnYW5pemF0aW9uYWwgVW5pdDogQXBhY2hlIE9GQml6IiB8fCAoYm9keT0id3d3Lm9mYml6Lm9yZyIgJiYgYm9keT0iL2ltYWdlcy9vZmJpel9wb3dlcmVkLmdpZiIpIHx8IGhlYWRlcj0iU2V0LUNvb2tpZTogT0ZCaXouVmlzaXRvciIgfHwgYmFubmVyPSJTZXQtQ29va2llOiBPRkJpei5WaXNpdG9yIg%3D%3D) |
| **Number of assets affected**  | 	5883 |
| **Description**  | Apache OFBiz is an open source enterprise resource planning (ERP) system that provides a variety of business functions and modules.Apache OFBiz has a deserialization code execution vulnerability in webtools/control/xmlrpc. An attacker can use this vulnerability to execute arbitrary code on the server side, write a backdoor, obtain server permissions, and then control the entire web server. |
| **Impact** | Apache OFBiz has a deserialization code execution vulnerability in webtools/control/xmlrpc. An attacker can use this vulnerability to execute arbitrary code on the server side, write a backdoor, obtain server permissions, and then control the entire web server. |
.
![](https://s3.bmp.ovh/imgs/2023/12/07/a73e0ef58369cbd8.gif).



## CrushFTP as2-to Authentication Permission bypass Vulnerability (CVE-2023-43177)

|   **Vulnerability**  | 	CrushFTP as2-to Authentication Permission bypass Vulnerability (CVE-2023-43177)  |
| :----:   | :-----|
|  **Chinese name**  | CrushFTP as2-to 认证权限绕过漏洞（CVE-2023-43177） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [app="crushftp-WebInterface"](https://en.fofa.info/result?qbase64=c2VydmVyPSJDcnVzaEZUUCIgfHwgaGVhZGVyPSIvV2ViSW50ZXJmYWNlL2xvZ2luLmh0bWwiIHx8IGJhbm5lcj0iL1dlYkludGVyZmFjZS9sb2dpbi5odG1sIiB8fCBoZWFkZXI9Ii9XZWJJbnRlcmZhY2UvdzNjL3AzcC54bWwiIHx8IGJhbm5lcj0iL1dlYkludGVyZmFjZS93M2MvcDNwLnhtbCIgfHwgdGl0bGU9IkNydXNoRlRQIg%3D%3D) |
| **Number of assets affected**  | 	38695 |
| **Description**  | CrushFTP is a powerful file transfer server suitable for secure and efficient file transfer and management for individual or enterprise users.CrashFTP has a permission bypass vulnerability, where attackers can bypass system permission control by constructing malicious as2 to request authentication, achieving arbitrary execution of malicious operations such as file read and delete. |
| **Impact** | CrashFTP has a permission bypass vulnerability, where attackers can bypass system permission control by constructing malicious as2 to request authentication, achieving arbitrary execution of malicious operations such as file read and delete. |

![](https://s3.bmp.ovh/imgs/2023/11/30/8d49b65293d87b3a.gif)


## Splunk Enterprise XSLT Command Execute Vulnerability (CVE-2023-46214)

|   **Vulnerability**  | Splunk Enterprise XSLT Command Execute Vulnerability (CVE-2023-46214)  |
| :----:   | :-----|
|  **Chinese name**  | Splunk Enterprise XSLT 命令执行漏洞（CVE-2023-46214）	 |
| **CVSS core**  | 	8.0 |
| **FOFA Query**  (click to view the results directly)| [app="splunk-Enterprise"](https://en.fofa.info/result?qbase64=Ym9keT0iX19zcGx1bmtkX3BhcnRpYWxzX18iICB8fCAoaGVhZGVyPSJTZXQtQ29va2llOiBzcGx1bmt3ZWJfdWlkPSIgJiYgYm9keT0iZW50ZXJwcmlzZSIp)|
| **Number of assets affected**  | 134567 |
| **Description**  | Splunk Enterprise is a data collection and analysis software developed by Splunk Corporation in the United States. This software is mainly used for collecting, indexing, and analyzing the data it generates, including data generated by all IT systems and infrastructure (physical, virtual machines, and cloud).Splunk Enterprise has a command execution vulnerability that does not securely clean up user provided Extensible Stylesheet Language Transformations (XSLTs). Attackers can exploit this vulnerability to upload malicious XSLTs and remotely execute commands on Splunk Enterprise instances. |
| **Impact** | 	Splunk Enterprise has a command execution vulnerability that does not securely clean up user provided Extensible Stylesheet Language Transformations (XSLTs). Attackers can exploit this vulnerability to upload malicious XSLTs and remotely execute commands on Splunk Enterprise instances. |

![](https://s3.bmp.ovh/imgs/2023/11/27/4f67b965b611da68.gif)



## SysAid userentry file upload vulnerability (CVE-2023-47246)

|   **Vulnerability**  | SysAid userentry file upload vulnerability (CVE-2023-47246)  |
| :----:   | :-----|
|  **Chinese name**  | SysAid userentry 文件上传漏洞（CVE-2023-47246）	 |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [	app="SysAid-Help-Desk"](https://en.fofa.info/result?qbase64=Ym9keT0ic3lzYWlkLWxvZ28tZGFyay1ncmVlbi5wbmciIHx8IHRpdGxlPSJTeXNBaWQgSGVscCBEZXNrIFNvZnR3YXJlIiB8fCBib2R5PSJIZWxwIERlc2sgc29mdHdhcmUgPGEgaHJlZj1cImh0dHA6Ly93d3cuc3lzYWlkLmNvbVwiPmJ5IFN5c0FpZDwvYT4i)|
| **Number of assets affected**  | 1819 |
| **Description**  | 	SysAid is an information technology (IT) service management and help desk solution designed to help organizations more effectively manage their IT infrastructure, help desk support and user needs. SysAid provides a series of functions, including fault reporting, asset management, problem management, change management, knowledge base, automated workflow, etc., to help enterprises improve the efficiency and quality of IT services.SysAid has a file upload vulnerability in userentry. An attacker can use the file upload vulnerability to execute malicious code, write backdoors, and read sensitive files, which may cause the server to be attacked and controlled. |
| **Impact** | 	SysAid has a file upload vulnerability in userentry. An attacker can use the file upload vulnerability to execute malicious code, write backdoors, and read sensitive files, which may cause the server to be attacked and controlled. |

![](https://s3.bmp.ovh/imgs/2023/11/15/105d49cb7220f659.gif)



## Honeywell PM43 loadfile.lp file command execution vulnerability (CVE-2023-3710)

|   **Vulnerability**  | Honeywell PM43 loadfile.lp file command execution vulnerability (CVE-2023-3710)  |
| :----:   | :-----|
|  **Chinese name**  | Honeywell PM43 loadfile.lp 文件命令执行漏洞（CVE-2023-3710）	 |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="Honeywell PM43 "](https://en.fofa.info/result?qbase64=aGVhZGVyPSJQTTQzIiB8fCBiYW5uZXI9IlBNNDMiIHx8IHRpdGxlPSJQTTQzIiB8fCBib2R5PSIvbWFpbi9sb2dpbi5sdWE%2FcGFnZWlkPUNvbmZpZ3VyZSI%3D)|
| **Number of assets affected**  | 96 |
| **Description**  | 	The Honeywell PM43 is a printer product of the American company Honeywell.Honeywell PM43P10.19.050004 and earlier versions of the input verification error vulnerability, attackers can arbitrarily execute code on the server side, write a backdoor, obtain server permissions, and then control the entire web server. |
| **Impact** | 	Honeywell PM43P10.19.050004 and earlier versions of the input verification error vulnerability, attackers can arbitrarily execute code on the server side, write a backdoor, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/10/17/ff602decce69f83b.gif)




## JetBrains TeamCity remote command execution vulnerability (CVE-2023-42793)

|   **Vulnerability**  | 	JetBrains TeamCity remote command execution vulnerability (CVE-2023-42793)  |
| :----:   | :-----|
|  **Chinese name**  | 	JetBrains TeamCity 远程命令执行漏洞（CVE-2023-42793） |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="JET_BRAINS-TeamCity"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJUZWFtY2l0eSIgfHwgYmFubmVyPSJUZWFtY2l0eSIgfHwgdGl0bGU9IlRlYW1DaXR5IiB8fCBib2R5PSJjb250ZW50PVwiVGVhbUNpdHkgKExvZyBpbiB0byBUZWFtQ2l0eSI%3D)|
| **Number of assets affected**  | 26963 |
| **Description**  | 	JetBrains TeamCity is a general CI/CD software platform developed by JetBrains.JetBrains TeamCity can obtain the valid token of the corresponding id user by accessing the /app/rest/users/{{id}}/tokens/RPC2 endpoint. Accessing the restricted endpoint with the admin token will cause remote command execution or the creation of a background administrator user. |
| **Impact** | 	JetBrains TeamCity can obtain the valid token of the corresponding id user by accessing the /app/rest/users/{{id}}/tokens/RPC2 endpoint. Accessing the restricted endpoint with the admin token will cause remote command execution or the creation of a background administrator user. |

![](https://s3.bmp.ovh/imgs/2023/10/13/d46d67aca09ab4ae.gif)



## JeeSpringCloud uploadFile.jsp file upload vulnerability

|   **Vulnerability**  | 	JeeSpringCloud uploadFile.jsp file upload vulnerability  |
| :----:   | :-----|
|  **Chinese name**  | 	JeeSpringCloud uploadFile.jsp 文件上传漏洞 |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="JeeSpringCloud"](https://en.fofa.info/result?qbase64=Ym9keT0iL2plZVNwcmluZ1N0YXRpYy9wbHVncy9qcXVlcnkvanF1ZXJ5IiB8fCBoZWFkZXI9ImNvbS5qZWVzcHJpbmcuc2Vzc2lvbi5pZCIgfHwgaGVhZGVyPSJjb20uamVlc3ByaW5nLnNlc3Npb24uaWQi)|
| **Number of assets affected**  | 	282 |
| **Description**  | 	JeeSpringCloud is a free and open source Java Internet cloud rapid development platform.JeeSpringCloud can upload any file by accessing /static/uploadify/uploadFile.jsp and specify the file upload path through the ?uploadPath parameter, causing the server to be controlled. |
| **Impact** | 	An attacker can use this vulnerability to write a backdoor on the server side, execute code, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/10/12/800e0ccfe95d6cbb.gif)

## Atlassian Confluence permission bypass vulnerability (CVE-2023-22515)

|   **Vulnerability**  | 	Atlassian Confluence permission bypass vulnerability (CVE-2023-22515)  |
| :----:   | :-----|
|  **Chinese name**  | 	Atlassian Confluence 权限绕过漏洞（CVE-2023-22515） |
| **CVSS core**  | 	10.0 |
| **FOFA Query**  (click to view the results directly)| [app="ATLASSIAN-Confluence"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJDb25mbHVlbmNlIiB8fCBiYW5uZXI9IkNvbmZsdWVuY2UiIHx8IGJvZHk9ImNvbmZsdWVuY2UtYmFzZS11cmwiIHx8IGJvZHk9ImNvbS1hdGxhc3NpYW4tY29uZmx1ZW5jZSIgfHwgIHRpdGxlPSJBdGxhc3NpYW4gQ29uZmx1ZW5jZSIgfHwgKHRpdGxlPT0iRXJyb3JzIiAmJiBib2R5PSJDb25mbHVlbmNlIik%3D)|
| **Number of assets affected**  | 	97667 |
| **Description**  | 	Atlassian Confluence is a software developed by Atlassian based on the online enterprise wiki (collaboration software).A vulnerability exists in the Atlassian Confluence data center and server. The /server-info.action endpoint is used to pass the bootstrapStatusProvider.applicationConfig.setupComplete parameter, leaving the server in an incomplete state to access restricted endpoints and create unauthorized Confluence administrator accounts. Log in to the Confluence instance backend. |
| **Impact** | 	A vulnerability exists in the Atlassian Confluence data center and server. The /server-info.action endpoint is used to pass the bootstrapStatusProvider.applicationConfig.setupComplete parameter, leaving the server in an incomplete state to access restricted endpoints and create unauthorized Confluence administrator accounts. Log in to the Confluence instance backend. |

![](https://s3.bmp.ovh/imgs/2023/10/11/c0c440512d0c5ee2.gif)


## Junos webauth_operation.php PHPRC Code Execution Vulnerability (CVE-2023-36845/CVE-2023-36846)

|   **Vulnerability**  | 	Junos webauth_operation.php PHPRC Code Execution Vulnerability (CVE-2023-36845/CVE-2023-36846)  |
| :----:   | :-----|
|  **Chinese name**  | 	Junos webauth_operation.php PHPRC 代码执行漏洞（CVE-2023-36845/CVE-2023-36846） |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="JUNIPer-Web-Device-Manager"](https://en.fofa.info/result?qbase64=IHRpdGxlPSJKdW5pcGVyIFdlYiBEZXZpY2UgTWFuYWdlciIgfHwgYmFubmVyPSJqdW5pcGVyIiB8fCBoZWFkZXI9Imp1bmlwZXIiIHx8IGJvZHk9InN2ZzRldmVyeWJvZHkvc3ZnNGV2ZXJ5Ym9keS5qcyIgfHwgYm9keT0ianVuaXBlci5uZXQvdXMvZW4vbGVnYWwtbm90aWNlcyIgfHwgYm9keT0ibmF0aXZlbG9naW5fbG9naW5fY3JlZGVudGlhbHMi)|
| **Number of assets affected**  | 	43627 |
| **Description**  | 	Junos is a reliable, high-performance network operating system from Juniper Networks.An attacker can use the J-Web service of the Junos operating system to pass in the PHPRC environment variable, turn on the allow_url_include setting, run the incoming encoded PHP code, and gain control of the entire web server. |
| **Impact** | 	Attackers can use this vulnerability to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/10/11/3e4434d6602a32a6.gif)



## Apache Superset Cookie Permission Bypass Vulnerability (CVE-2023-30776)

|   **Vulnerability**  | 	Apache Superset Cookie Permission Bypass Vulnerability (CVE-2023-30776)  |
| :----:   | :-----|
|  **Chinese name**  | 	Apache Superset Cookie 权限绕过漏洞（CVE-2023-27524） |
| **CVSS core**  | 	9.8 |
| **FOFA Query**  (click to view the results directly)| [app="APACHE-Superset"](https://en.fofa.info/result?qbase64=KHRpdGxlPSJTdXBlcnNldCIgJiYgKGJvZHk9ImFwcGJ1aWxkZXIiIHx8IGJvZHk9IjxpbWcgc3JjPVwiaHR0cHM6Ly9qb2luc3VwZXJzZXQuY29tL2ltZy9zdXBlcnNldGxvZ292ZWN0b3Iuc3ZnIikpIHx8IGJvZHk9IjxhIGhyZWY9XCJodHRwczovL21hbmFnZS5hcHAtc2R4LnByZXNldC5pb1wiIGNsYXNzPVwiYnV0dG9uXCI%2BQmFjayB0byB3b3Jrc3BhY2VzPC9hPjwvc2VjdGlvbj4iIHx8IChib2R5PSIvc3RhdGljL2Fzc2V0cy9kaXN0L2NvbW1vbi42NDRhZTdhZTk3M2IwMGFiYzE0Yi5lbnRyeS5qcyIgfHwgKGJvZHk9Ii9zdGF0aWMvYXNzZXRzL2ltYWdlcy9mYXZpY29uLnBuZyIgJiYgYm9keT0iL3N0YXRpYy9hcHBidWlsZGVyL2pzL2pxdWVyeS1sYXRlc3QuanMiKSAmJiBib2R5PSJTdXBlcnNldCIpIHx8IGhlYWRlcj0iL3N1cGVyc2V0L3dlbGNvbWUvIiB8fCB0aXRsZT0iNTAwOiBJbnRlcm5hbCBzZXJ2ZXIgZXJyb3IgfCBTdXBlcnNldCIgfHwgdGl0bGU9IjQwNDogTm90IGZvdW5kIHwgU3VwZXJzZXQiIHx8IGJhbm5lcj0iL3N1cGVyc2V0L3dlbGNvbWUvIiB8fCBiYW5uZXI9Ii9zdXBlcnNldC9kYXNoYm9hcmQvIg%3D%3D)|
| **Number of assets affected**  | 		56089 |
| **Description**  | 	Apache Superset is an open source modern data exploration and visualization platform.Apache Superset Cookie has a permission bypass vulnerability that allows an attacker to control the entire system, ultimately leaving the system in an extremely unsafe state. |
| **Impact** | 	Apache Superset Cookie has a permission bypass vulnerability that allows an attacker to control the entire system, ultimately leaving the system in an extremely unsafe state. |

![](https://s3.bmp.ovh/imgs/2023/10/10/ec147824884d3597.gif)


## Cockpit assetsmanager/upload file upload vulnerability (CVE-2023-1313)

|   **Vulnerability**  | 	Cockpit assetsmanager/upload file upload vulnerability (CVE-2023-1313)  |
| :----:   | :-----|
|  **Chinese name**  | 	Cockpit assetsmanager/upload 文件上传漏洞（CVE-2023-1313） |
| **CVSS core**  | 	7.2 |
| **FOFA Query**  (click to view the results directly)| [title="Authenticate Please!" \|\| body="password:this.refs.password.value"  \|\| body="UIkit.components.formPassword.prototype.defaults.lblShow"  \|\| body="App.request('/auth/check'"](https://en.fofa.info/result?qbase64=dGl0bGU9IkF1dGhlbnRpY2F0ZSBQbGVhc2UhIiB8fCBib2R5PSJwYXNzd29yZDp0aGlzLnJlZnMucGFzc3dvcmQudmFsdWUiIHx8IGJvZHk9IlVJa2l0LmNvbXBvbmVudHMuZm9ybVBhc3N3b3JkLnByb3RvdHlwZS5kZWZhdWx0cy5sYmxTaG93IiB8fCBib2R5PSJBcHAucmVxdWVzdCgnL2F1dGgvY2hlY2snIg%3D%3D)|
| **Number of assets affected**  | 	3185 |
| **Description**  | 	Cockpit is a self-hosted, flexible and user-friendly headless content platform for creating custom digital experiences.Cockpit has a file upload vulnerability, which allows attackers to upload arbitrary files, leading to server control, etc. |
| **Impact** | 	Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/09/29/7fa28d36837b1e06.gif)

## Revive Adserver adxmlrpc.php Remote Code Execution Vulnerability (CVE-2019-5434)

|   **Vulnerability**  | 	Revive Adserver adxmlrpc.php Remote Code Execution Vulnerability (CVE-2019-5434)  |
| :----:   | :-----|
|  **Chinese name**  | 	Revive Adserver 广告管理系统 adxmlrpc.php 文件远程代码执行漏洞（CVE-2019-5434） |
| **CVSS core**  | 	9.0 |
| **FOFA Query**  (click to view the results directly)| [title="Revive Adserver" \|\| body="strPasswordMinLength" \|\| body="Welcome to Revive Adserver"](https://en.fofa.info/result?qbase64=dGl0bGU9IlJldml2ZSBBZHNlcnZlciIgfHwgYm9keT0ic3RyUGFzc3dvcmRNaW5MZW5ndGgiIHx8IGJvZHk9IldlbGNvbWUgdG8gUmV2aXZlIEFkc2VydmVyIg%3D%3D)|
| **Number of assets affected**  | 	5667 |
| **Description**  | 	Revive Adserver is an open source advertising management system developed by the Revive Adserver team. The system provides functions such as advertising placement, advertising space management, and data statistics.The delivery XML-RPC script in versions prior to Revive Adserver 4.2.0 has a code problem vulnerability, and an attacker can execute arbitrary code to obtain server permissions. |
| **Impact** | 	The delivery XML-RPC script in versions prior to Revive Adserver 4.2.0 has a code problem vulnerability, and an attacker can execute arbitrary code to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/09/25/74874bad3ffe8fc6.gif)




## Revive Adserver adxmlrpc.php Remote Code Execution Vulnerability (CVE-2019-5434)

|   **Vulnerability**  | 	Revive Adserver adxmlrpc.php Remote Code Execution Vulnerability (CVE-2019-5434)  |
| :----:   | :-----|
|  **Chinese name**  | 	Revive Adserver 广告管理系统 adxmlrpc.php 文件远程代码执行漏洞（CVE-2019-5434） |
| **CVSS core**  | 	9.0 |
| **FOFA Query**  (click to view the results directly)| [title="Revive Adserver" \|\| body="strPasswordMinLength" \|\| body="Welcome to Revive Adserver"](https://en.fofa.info/result?qbase64=dGl0bGU9IlJldml2ZSBBZHNlcnZlciIgfHwgYm9keT0ic3RyUGFzc3dvcmRNaW5MZW5ndGgiIHx8IGJvZHk9IldlbGNvbWUgdG8gUmV2aXZlIEFkc2VydmVyIg%3D%3D)|
| **Number of assets affected**  | 	5667 |
| **Description**  | 	Revive Adserver is an open source advertising management system developed by the Revive Adserver team. The system provides functions such as advertising placement, advertising space management, and data statistics.The delivery XML-RPC script in versions prior to Revive Adserver 4.2.0 has a code problem vulnerability, and an attacker can execute arbitrary code to obtain server permissions. |
| **Impact** | 	The delivery XML-RPC script in versions prior to Revive Adserver 4.2.0 has a code problem vulnerability, and an attacker can execute arbitrary code to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/09/25/74874bad3ffe8fc6.gif)




## Weaver E-office flow_xml.php file SORT_ID parameter SQL injection vulnerability

|   **Vulnerability**  | 	Weaver E-office flow_xml.php file SORT_ID parameter SQL injection vulnerability  |
| :----:   | :-----|
|  **Chinese name**  | 	泛微 E-office flow_xml.php 文件 SORT_ID 参数 SQL 注入漏洞 |
| **CVSS core**  | 	7.8 |
| **FOFA Query**  (click to view the results directly)| [body="href=\"/eoffice" \|\| body="/eoffice10/client" \|\| body="eoffice_loading_tip" \|\| body="eoffice_init" \|\| header="general/login/index.php" \|\| banner="general/login/index.php" \|\| body="/general/login/view//images/updateLoad.gif" \|\| (body="szFeatures" && body="eoffice") \|\| header="eOffice" \|\| banner="eOffice"](https://en.fofa.info/result?qbase64=Ym9keT0iaHJlZj1cIi9lb2ZmaWNlIiB8fCBib2R5PSIvZW9mZmljZTEwL2NsaWVudCIgfHwgYm9keT0iZW9mZmljZV9sb2FkaW5nX3RpcCIgfHwgYm9keT0iZW9mZmljZV9pbml0IiB8fCBoZWFkZXI9ImdlbmVyYWwvbG9naW4vaW5kZXgucGhwIiB8fCBiYW5uZXI9ImdlbmVyYWwvbG9naW4vaW5kZXgucGhwIiB8fCBib2R5PSIvZ2VuZXJhbC9sb2dpbi92aWV3Ly9pbWFnZXMvdXBkYXRlTG9hZC5naWYiIHx8IChib2R5PSJzekZlYXR1cmVzIiAmJiBib2R5PSJlb2ZmaWNlIikgfHwgaGVhZGVyPSJlT2ZmaWNlIiB8fCBiYW5uZXI9ImVPZmZpY2Ui)|
| **Number of assets affected**  | 	21632 |
| **Description**  | 	Weaver&nbsp;e-office is an OA product for small and medium-sized organizations, developed by Weaver Network Technology Co., LTD.There is an SQL injection vulnerability in flow_xml.php, which can be used by attackers to obtain information in the database (for example, administrator background password, site user personal information). |
| **Impact** | 	An attacker can exploit the SQL injection vulnerability to obtain information from the database (for example, administrator background passwords, site user personal information). |

![](https://s3.bmp.ovh/imgs/2023/09/25/6a416f12923360a7.gif)




## GJP SelectImage.aspx file upload vulnerability

|   **Vulnerability**  | 	GJP SelectImage.aspx file upload vulnerability  |
| :----:   | :-----|
|  **Chinese name**  | 	管家婆订货易在线商城 SelectImage.aspx 文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="订货易"](https://en.fofa.info/result?qbase64=dGl0bGU9Iuiuoui0p%2BaYkyI%3D) |
| **Number of assets affected**  | 	2617 |
| **Description**  | 	Renwoxing took the lead in launching the Guanjiapo purchase, sales, inventory and financial integration software for small and medium-sized enterprises.There is a SelectImage.aspx arbitrary file upload vulnerability in the Guanjiapo Ordering Online Mall. An attacker can use this vulnerability to control the entire system, ultimately causing the system to be in an extremely unsafe state. |
| **Impact** | An attacker can take control of the entire system through this vulnerability, ultimately leaving the system in an extremely unsafe state. |

![](https://s3.bmp.ovh/imgs/2023/09/15/68d54cb7ae080127.gif)


## Junos webauth_operation.php File Upload Vulnerability (CVE-2023-36844)

|   **Vulnerability**  | 	Junos webauth_operation.php File Upload Vulnerability (CVE-2023-36844)  |
| :----:   | :-----|
|  **Chinese name**  | Junos webauth_operation.php 文件上传漏洞（CVE-2023-36844） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="Juniper Web Device Manager" \|\| banner="juniper" \|\| header="juniper" \|\| body="svg4everybody/svg4everybody.js" \|\| body="juniper.net/us/en/legal-notices" \|\| body="nativelogin_login_credentials"](https://en.fofa.info/result?qbase64=dGl0bGU9Ikp1bmlwZXIgV2ViIERldmljZSBNYW5hZ2VyIiB8fCBiYW5uZXI9Imp1bmlwZXIiIHx8IGhlYWRlcj0ianVuaXBlciIgfHwgYm9keT0ic3ZnNGV2ZXJ5Ym9keS9zdmc0ZXZlcnlib2R5LmpzIiB8fCBib2R5PSJqdW5pcGVyLm5ldC91cy9lbi9sZWdhbC1ub3RpY2VzIiB8fCBib2R5PSJuYXRpdmVsb2dpbl9sb2dpbl9jcmVkZW50aWFscyI%3D) |
| **Number of assets affected**  | 	47518 |
| **Description**  | Junos is a reliable, high-performance network operating system from Juniper Networks.An attacker can use the J-Web service /webauth_operation.php route of the Junos operating system to upload a php webshell, include it through the ?PHPRC parameter, and gain control of the entire web server. |
| **Impact** | Attackers can use this vulnerability to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/09/14/58f32076aab47bc2.gif)


## Weaver ecology XmlRpcServlet Path File Read Vulnerability

|   **Vulnerability**  | **Weaver ecology XmlRpcServlet Path File Read Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 泛微 e-cology XmlRpcServlet 接口文件读取漏洞 |
| **CVSS core**  | 7.8 |
| **FOFA Query**  (click to view the results directly)| [((body="szFeatures" && body="redirectUrl") \|\| (body="rndData" && body="isdx") \|\| (body="typeof poppedWindow" && body="client/jquery.client_wev8.js") \|\| body="/theme/ecology8/jquery/js/zDialog_wev8.js" \|\| body="ecology8/lang/weaver_lang_7_wev8.js" \|\| body="src=\"/js/jquery/jquery_wev8.js" \|\| (header="Server: WVS" && (title!="404 Not Found" && header!="404 Not Found"))) && header!="testBanCookie" && header!="Couchdb" && header!="JoomlaWor" && body!="<title>28ZE</title>"](https://en.fofa.info/result?qbase64=KChib2R5PSJzekZlYXR1cmVzIiAmJiBib2R5PSJyZWRpcmVjdFVybCIpIHx8IChib2R5PSJybmREYXRhIiAmJiBib2R5PSJpc2R4IikgfHwgKGJvZHk9InR5cGVvZiBwb3BwZWRXaW5kb3ciICYmIGJvZHk9ImNsaWVudC9qcXVlcnkuY2xpZW50X3dldjguanMiKSB8fCBib2R5PSIvdGhlbWUvZWNvbG9neTgvanF1ZXJ5L2pzL3pEaWFsb2dfd2V2OC5qcyIgfHwgYm9keT0iZWNvbG9neTgvbGFuZy93ZWF2ZXJfbGFuZ183X3dldjguanMiIHx8IGJvZHk9InNyYz1cIi9qcy9qcXVlcnkvanF1ZXJ5X3dldjguanMiIHx8IChoZWFkZXI9IlNlcnZlcjogV1ZTIiAmJiAodGl0bGUhPSI0MDQgTm90IEZvdW5kIiAmJiBoZWFkZXIhPSI0MDQgTm90IEZvdW5kIikpKSAmJiBoZWFkZXIhPSJ0ZXN0QmFuQ29va2llIiAmJiBoZWFkZXIhPSJDb3VjaGRiIiAmJiBoZWFkZXIhPSJKb29tbGFXb3IiICYmIGJvZHkhPSI8dGl0bGU%2BMjhaRTwvdGl0bGU%2BIg%3D%3D) |
| **Number of assets affected**  | 111321 |
| **Description**  | Weaver e-cology is an OA office system specifically designed for large and medium-sized enterprises, supporting simultaneous work on PC, mobile, and WeChat platforms.Attackers can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., resulting in an extremely insecure state of the website. |
| **Impact** | Attackers can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., resulting in an extremely insecure state of the website. |

![](https://s3.bmp.ovh/imgs/2023/09/07/af06beef9479a1c1.gif)



## UF U8 Cloud upload.jsp file upload vulnerability

|   **Vulnerability**  | **UF U8 Cloud upload.jsp file upload vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 用友 U8 Cloud upload.jsp 文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="开启U8 cloud云端之旅"](https://en.fofa.info/result?qbase64=Ym9keT0i5byA5ZCvVTggY2xvdWTkupHnq6%2FkuYvml4Ui) |
| **Number of assets affected**  | 13473 |
| **Description**  | yonyou U8 cloud is a cloud ERP developed by yonyou.There is a file upload vulnerability in yonyou U8 upload.jsp, which can be exploited by attackers to gain server privileges. |
| **Impact** | Attackers can use this vulnerability to upload file, execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/09/06/407cf745d8210300.gif)


## Ruijie EWEB Network Management System flwo.control.php type Arbitrary Command Execution Vulnerability

|   **Vulnerability**  | **Ruijie EWEB Network Management System flwo.control.php type Arbitrary Command Execution Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | Ruijie-EWEB 网管系统 flwo.control.php 文件 type 参数任意命令执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(body="<span class=\"resource\" mark=\"login.copyRight\">锐捷网络</span>" && body="login.getDeviceInfo") \|\| title="锐捷网络-EWEB网管系统"]((https://en.fofa.info/result?qbase64=KGJvZHk9IjxzcGFuIGNsYXNzPVwicmVzb3VyY2VcIiBtYXJrPVwibG9naW4uY29weVJpZ2h0XCI%2B6ZSQ5o23572R57ucPC9zcGFuPiIgJiYgYm9keT0ibG9naW4uZ2V0RGV2aWNlSW5mbyIpIHx8IHRpdGxlPSLplJDmjbfnvZHnu5wtRVdFQue9keeuoeezu%2Be7nyI%3D)) |
| **Number of assets affected**  | 11544 |
| **Description**  | Ruijie Network Management System is a new generation of cloud based network management software developed by Beijing Ruijie Data Era Technology Co., Ltd. With the slogan of "Innovative Network Management and Information Security in the Data Age", it is positioned as a unified solution for terminal security, IT operations, and enterprise service-oriented management.Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/09/06/3fa09e24d0bce4cb.gif)



## Adobe ColdFusion WDDX JGroups remote code execution vulnerability

|   **Vulnerability**  | **Adobe ColdFusion WDDX JGroups remote code execution vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | Adobe ColdFusion WDDX JGroups 远程代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="/cfajax/" \|\| header="CFTOKEN" \|\| banner="CFTOKEN" \|\| body="ColdFusion.Ajax" \|\| body="<cfscript>" \|\| server="ColdFusion" \|\| title="ColdFusion" \|\| (body="crossdomain.xml" && body="CFIDE") \|\| (body="#000808" && body="#e7e7e7")](https://en.fofa.info/result?qbase64=Ym9keT0iL2NmYWpheC8iIHx8IGhlYWRlcj0iQ0ZUT0tFTiIgfHwgYmFubmVyPSJDRlRPS0VOIiB8fCBib2R5PSJDb2xkRnVzaW9uLkFqYXgiIHx8IGJvZHk9IjxjZnNjcmlwdD4iIHx8IHNlcnZlcj0iQ29sZEZ1c2lvbiIgfHwgdGl0bGU9IkNvbGRGdXNpb24iIHx8IChib2R5PSJjcm9zc2RvbWFpbi54bWwiICYmIGJvZHk9IkNGSURFIikgfHwgKGJvZHk9IiMwMDA4MDgiICYmIGJvZHk9IiNlN2U3ZTciKQ%3D%3D) |
| **Number of assets affected**  | 567468 |
| **Description**  | Adobe Coldfusion is a commercial application server developed by Adobe for web applications.The attacker can send unbelievable serialized data and trigger derivativeization to the Coldfusion server, thereby executing any code. |
| **Impact** |The attacker can execute the code at the server through this vulnerability, obtain the server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/09/05/459741d98c251494.gif)



## Adobe ColdFusion WDDX C3P0 remote code execution vulnerability

|   **Vulnerability**  | **Adobe ColdFusion WDDX C3P0 remote code execution vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | Adobe ColdFusion WDDX C3P0 远程代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="/cfajax/" \|\| header="CFTOKEN" \|\| banner="CFTOKEN" \|\| body="ColdFusion.Ajax" \|\| body="<cfscript>" \|\| server="ColdFusion" \|\| title="ColdFusion" \|\| (body="crossdomain.xml" && body="CFIDE") \|\| (body="#000808" && body="#e7e7e7")](https://en.fofa.info/result?qbase64=Ym9keT0iL2NmYWpheC8iIHx8IGhlYWRlcj0iQ0ZUT0tFTiIgfHwgYmFubmVyPSJDRlRPS0VOIiB8fCBib2R5PSJDb2xkRnVzaW9uLkFqYXgiIHx8IGJvZHk9IjxjZnNjcmlwdD4iIHx8IHNlcnZlcj0iQ29sZEZ1c2lvbiIgfHwgdGl0bGU9IkNvbGRGdXNpb24iIHx8IChib2R5PSJjcm9zc2RvbWFpbi54bWwiICYmIGJvZHk9IkNGSURFIikgfHwgKGJvZHk9IiMwMDA4MDgiICYmIGJvZHk9IiNlN2U3ZTciKQ%3D%3D) |
| **Number of assets affected**  | 567468 |
| **Description**  | Adobe Coldfusion is a commercial application server developed by Adobe for web applications.The attacker can send unbelievable serialized data and trigger derivativeization to the Coldfusion server, thereby executing any code. |
| **Impact** | The attacker can execute the code at the server through this vulnerability, obtain the server permissions, and then control the entire web server. |




## Dahua Smart Park Integrated Management Platform searchJson SQL injection vulnerability

|   **Vulnerability**  | **Dahua Smart Park Integrated Management Platform searchJson SQL injection vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 大华智慧园区综合管理平台 searchJson SQL 注入漏洞 |
| **CVSS core**  | 8.2 |
| **FOFA Query**  (click to view the results directly)| [body="src=\\"/WPMS/asset/common/js/jsencrypt.min.js\'"](https://en.fofa.info/result?qbase64=Ym9keT0ic3JjPVwiL1dQTVMvYXNzZXQvY29tbW9uL2pzL2pzZW5jcnlwdC5taW4uanMnIg%3D%3D) |
| **Number of assets affected**  | 5415 |
| **Description**  | Dahua Smart Park Integrated Management Platform is a comprehensive management platform for smart parks built to provide security and efficient management of general public buildings. By integrating Dahua’s experience and cutting-edge technologies in the field of security and intelligence, it integrates video, access control, alarm, parking lot, attendance, visitor, video intercom, information release and other business subsystems to provide customers with a set of integrated, efficient, open, flexible and scalable platform software products, forming a comprehensive security solution for the six major areas of "public management, infrastructure, economic development, ecological protection, security, and social services". There is a sql injection vulnerability in the comprehensive management platform of Dahua Smart Park. In addition to using the SQL injection vulnerability to obtain information in the database (for example, administrator background passwords, personal information of site users), attackers can even write Trojan horses into the server under high-privilege conditions to further obtain server system permissions. |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

## Dahua Smart Park Integrated Management Platform source/publishing/publishing/material/file/video File Upload Vulnerability

|   **Vulnerability**  | **Dahua Smart Park Integrated Management Platform source/publishing/publishing/material/file/video File Upload Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 大华智慧园区综合管理平台 source/publishing/publishing/material/file/video 文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="src=\\"/WPMS/asset/common/js/jsencrypt.min.js\"](https://en.fofa.info/result?qbase64=Ym9keT0ic3JjPVwiL1dQTVMvYXNzZXQvY29tbW9uL2pzL2pzZW5jcnlwdC5taW4uanMi) |
| **Number of assets affected**  | 5420 |
| **Description**  | Dahua Smart Park Comprehensive Management Platform The smart park comprehensive management platform is built to provide safe and efficient management for common public buildings. By integrating Dahua’s experience and cutting-edge technologies in the field of security and intelligence, it integrates video, access control, alarm, and parking Field, attendance, visitor, video intercom, information release and other business subsystems, providing customers with a set of integrated, efficient, open, flexible and scalable platform software products, forming a "public management, infrastructure, economic development Comprehensive security solutions in the six major areas of , ecological protection, security, and social services. There is a file upload vulnerability in the Dahua Smart Park system /publishing/, which leads to the server being controlled. |
| **Impact** | There is a file upload vulnerability in the comprehensive management platform of the smart park of Zhejiang Dahua Technology Co., Ltd. An attacker can use this vulnerability to obtain server permissions by uploading a specific configuration file. |

## SolarView Compact downloader.php RCE (CVE-2023-23333)

|   **Vulnerability**  | **SolarView Compact downloader.php RCE (CVE-2023-23333)**  |
| :----:   | :-----|
|  **Chinese name**  | SolarView Compact downloader.php 任意命令执行漏洞（CVE-2023-23333） |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [body="SolarView Compact"](https://en.fofa.info/result?qbase64=Ym9keT0iU29sYXJWaWV3IENvbXBhY3Qi) |
| **Number of assets affected**  | 4941 |
| **Description**  | There is a command injection vulnerability in SolarView Compact through 6.00, attackers can execute commands by bypassing internal restrictions through downloader.php. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/07/31/9e95a65910332af8.gif)

## EMERSON-XWEB-EVO upload.cgi path Directory Traversal Vulnerability (CVE-2021-45427)

|   **Vulnerability**  | **EMERSON-XWEB-EVO upload.cgi path Directory Traversal Vulnerability (CVE-2021-45427)**  |
| :----:   | :-----|
|  **Chinese name**  | EMERSON-XWEB-EVO upload.cgi 文件 path 参数目录遍历漏洞（CVE-2021-45427） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="src=\\"img/xweb-logo.png" \|\| body="src=\\"/css/images/Logo_XWEB_alpha.png"](https://en.fofa.info/result?qbase64=Ym9keT0ic3JjPVwiaW1nL3h3ZWItbG9nby5wbmciIHx8IGJvZHk9InNyYz1cIi9jc3MvaW1hZ2VzL0xvZ29fWFdFQl9hbHBoYS5wbmci) |
| **Number of assets affected**  | 15849 |
| **Description**  | Emerson XWEB 300D EVO is an energy-saving air conditioner of Emerson Company in the United States. Emerson XWEB 300D EVO 3.0.7 -- 3ee403 has a directory traversal vulnerability (CVE-2021-45427). An attacker may access some secret files including configuration files, logs, source codes, etc. by browsing the directory structure. With the comprehensive utilization of other vulnerabilities, the attacker can easily obtain higher permissions. |
| **Impact** | Emerson XWEB 300D EVO is an energy-saving air conditioner of Emerson Company in the United States. Emerson XWEB 300D EVO 3.0.7 -- 3ee403 has a directory traversal vulnerability (CVE-2021-45427). An attacker may access some secret files including configuration files, logs, source codes, etc. by browsing the directory structure. With the comprehensive utilization of other vulnerabilities, the attacker can easily obtain higher permissions. |

## Acmailer init_ctl.cgi sendmail_path Remote Command Execution Vulnerability (CVE-2021-20617)

|   **Vulnerability**  | **Acmailer init_ctl.cgi sendmail_path Remote Command Execution Vulnerability (CVE-2021-20617)**  |
| :----:   | :-----|
|  **Chinese name**  | Acmailer 邮件系统 init_ctl.cgi 文件 sendmail_path 参数远程命令执行漏洞（CVE-2021-20617） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="CGI acmailer"](https://en.fofa.info/result?qbase64=Ym9keT0iQ0dJIGFjbWFpbGVyIg%3D%3D) |
| **Number of assets affected**  | 557 |
| **Description**  | Acmailer is a CGI software used to support mail services. Acmailer 4.0.2 and earlier versions have a security vulnerability. The vulnerability is due to the fact that init_ctl.cgi does not strictly verify input parameters, and attackers can execute arbitrary commands to obtain server permissions. |
| **Impact** | Acmailer 4.0.2 and earlier versions have a security vulnerability. The vulnerability is due to the fact that init_ctl.cgi does not strictly verify input parameters, and attackers can execute arbitrary commands to obtain server permissions. |

## Acmailer enq_form.cgi Authentication Bypass Vulnerability (CVE-2021-20618)

|   **Vulnerability**  | **Acmailer enq_form.cgi Authentication Bypass Vulnerability (CVE-2021-20618)**  |
| :----:   | :-----|
|  **Chinese name**  | Acmailer 邮件系统 enq_form.cgi 认证绕过漏洞（CVE-2021-20618） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="CGI acmailer"](https://en.fofa.info/result?qbase64=Ym9keT0iQ0dJIGFjbWFpbGVyIg%3D%3D) |
| **Number of assets affected**  | 552 |
| **Description**  | Acmailer is a CGI software used to support mail services.Acmailer 4.0.2 and earlier versions have security vulnerabilities, which allow remote attackers to bypass authentication and gain administrative privileges. |
| **Impact** | Acmailer 4.0.2 and earlier versions have security vulnerabilities, which allow remote attackers to bypass authentication and gain administrative privileges. |

## Frappe Framework frappe.core.doctype.data_import.data_import.get_preview_from_template import_file Arbitrary File Read Vulnerability (CVE-2022-41712)

|   **Vulnerability**  | **Frappe Framework frappe.core.doctype.data_import.data_import.get_preview_from_template import_file Arbitrary File Read Vulnerability (CVE-2022-41712)**  |
| :----:   | :-----|
|  **Chinese name**  | Frappe-Framework 框架 frappe.core.doctype.data_import.data_import.get_preview_from_template 文件 import_file 参数任意文件读取漏洞（CVE-2022-41712） |
| **CVSS core**  | 6.5 |
| **FOFA Query**  (click to view the results directly)| [body="<meta name=\\"generator\\" content=\\"frappe" \|\| body="frappe.ready_events.push (fn):" \|\| header= "Link: </assets/frappe/js/lib/jquery/query.min.js" \|\| header="</assets/frap pe/dist/js/frappe-web.bundle.7XJQJMPF.js"](https://en.fofa.info/result?qbase64=Ym9keT0iPG1ldGEgbmFtZT1cImdlbmVyYXRvclwiIGNvbnRlbnQ9XCJmcmFwcGUiIHx8IGJvZHk9ImZyYXBwZS5yZWFkeV9ldmVudHMucHVzaCAoZm4pOiIgfHwgaGVhZGVyPSAiTGluazogPC9hc3NldHMvZnJhcHBlL2pzL2xpYi9qcXVlcnkvcXVlcnkubWluLmpzIiB8fCBoZWFkZXI9IjwvYXNzZXRzL2ZyYXAgcGUvZGlzdC9qcy9mcmFwcGUtd2ViLmJ1bmRsZS43WEpRSk1QRi5qcyI%3D) |
| **Number of assets affected**  | 48857 |
| **Description**  | An attacker can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in the extremely insecure state of the website. |
| **Impact** | At present, the manufacturer has issued an upgrade patch to fix the vulnerability. The patch access link is:https://github.com/frappe/frappe/releases/tag/v14.12.0 |

## Netgear Devices boardDataWW.php Unauthenticated Remote Command Execution

|   **Vulnerability**  | **Netgear Devices boardDataWW.php Unauthenticated Remote Command Execution**  |
| :----:   | :-----|
|  **Chinese name**  | Netgear 多款设备 boardDataWW.php 文件命令执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title=="Netgear"](https://en.fofa.info/result?qbase64=dGl0bGU9PSJOZXRnZWFyIg%3D%3D) |
| **Number of assets affected**  | 556 |
| **Description**  | (1) boardData102.php, (2) boardData103.php, (3) boardDataJP.php, (4) boardDataNA.php, and (5) boardDataWW.php in Netgear WN604 before 3.3.3 and WN802Tv2, WNAP210v2, WNAP320, WNDAP350, WNDAP360, and WNDAP660 before 3.5.5.0 allow remote attackers to execute arbitrary commands. |
| **Impact** | (1) boardData102.php, (2) boardData103.php, (3) boardDataJP.php, (4) boardDataNA.php, and (5) boardDataWW.php in Netgear WN604 before 3.3.3 and WN802Tv2, WNAP210v2, WNAP320, WNDAP350, WNDAP360, and WNDAP660 before 3.5.5.0 allow remote attackers to execute arbitrary commands. |

![](https://s3.bmp.ovh/imgs/2023/07/31/4509202e858461b5.gif)

## Metabase JDBC Remote Code Execution Vulnerability (CVE-2023-38646)

|   **Vulnerability**  | **Metabase JDBC Remote Code Execution Vulnerability (CVE-2023-38646)**  |
| :----:   | :-----|
| **Chinese name**  | Metabase JDBC 远程代码执行漏洞（CVE-2023-38646） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title=="Metabase" \|\| ((body="<script type=\"application/json\" id=\"_metabaseBootstrap\">" \|\| body="window.MetabaseLocalization = JSON.parse(document.getElementById(\"_metabaseLocalization\").textContent);") && body="window.MetabaseRoot = actualRoot;")](https://en.fofa.info/result?qbase64=dGl0bGU9PSJNZXRhYmFzZSIgfHwgKChib2R5PSI8c2NyaXB0IHR5cGU9XCJhcHBsaWNhdGlvbi9qc29uXCIgaWQ9XCJfbWV0YWJhc2VCb290c3RyYXBcIj4iIHx8IGJvZHk9IndpbmRvdy5NZXRhYmFzZUxvY2FsaXphdGlvbiA9IEpTT04ucGFyc2UoZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJfbWV0YWJhc2VMb2NhbGl6YXRpb25cIikudGV4dENvbnRlbnQpOyIpICYmIGJvZHk9IndpbmRvdy5NZXRhYmFzZVJvb3QgPSBhY3R1YWxSb290OyIp) |
| **Number of assets affected**  | 66604 |
| **Description**  | Metabase is an open source data analysis and visualization tool that helps users easily connect to various data sources, including databases, cloud services, and APIs, and then use an intuitive interface for data query, analysis, and visualization.A remote code execution vulnerability exists in Metabase that could allow an attacker to execute arbitrary code on a server running with Metabase server privileges. |
| **Impact** | A remote code execution vulnerability exists in Metabase that could allow an attacker to execute arbitrary code on a server running with Metabase server privileges. |

![](https://s3.bmp.ovh/imgs/2023/07/28/4a0b2c90aaf1b387.gif)

## Adobe Coldfusion remote code execution vulnerability (CVE-2023-38203)

|   **Vulnerability**  | **Adobe Coldfusion remote code execution vulnerability (CVE-2023-38203)**  |
| :----:   | :-----|
| **Chinese name**  | Adobe ColdFusion 远程代码执行漏洞（CVE-2023-38203） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(body="crossdomain.xml" && body="CFIDE") \|\| (body="#000808" && body="#e7e7e7")](https://en.fofa.info/result?qbase64=IChib2R5PSJjcm9zc2RvbWFpbi54bWwiICYmIGJvZHk9IkNGSURFIikgfHwgKGJvZHk9IiMwMDA4MDgiICYmIGJvZHk9IiNlN2U3ZTciKQ%3D%3D) |
| **Number of assets affected**  | 3740 |
| **Description**  | Adobe Coldfusion is a commercial application server developed by Adobe for web applications.The attacker can send unbelievable serialized data and trigger derivativeization to the Coldfusion server, thereby executing any code. |
| **Impact** | The attacker can execute the code at the server through this vulnerability, obtain the server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/07/24/514d4dd7f7e3b52c.gif)

## SANGFOR-IOMS catjs.php File Read Vulnerability

|   **Vulnerability**  | **SANGFOR-IOMS catjs.php File Read Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 深信服上网优化管理系统 catjs.php 文件读取漏洞 |
| **CVSS core**  | 6.0 |
| **FOFA Query**  (click to view the results directly)| [title="SANGFOR上网优化管理"](https://en.fofa.info/result?qbase64=CXRpdGxlPSJTQU5HRk9S5LiK572R5LyY5YyW566h55CGIg%3D%3D) |
| **Number of assets affected**  | 97 |
| **Description**  | Convinced by the Internet optimization management system deployment does not need to be adjusted, and transparent bridging mode is supported in organizational networks. At the same time, Intranet users can directly access the Internet regardless of any changes and maintain the original Internet access habits. This enables all data centers, links, and servers to be fully utilized.catjs.php file has any file reading vulnerability, through which an attacker can download any file in the server and leak sensitive information of the server. |
| **Impact** | Attackers can use this vulnerability to read important server files, such as system configuration files, database configuration files, and so on, causing the website to be in an extremely insecure state. |

## Command Execution Vulnerability in Hikvision Operations Management Center

|   **Vulnerability**  | **Command Execution Vulnerability in Hikvision Operations Management Center**  |
| :----:   | :-----|
| **Chinese name**  | 海康运行管理中心命令执行漏洞 |
| **CVSS core**  | 9.6 |
| **FOFA Query**  (click to view the results directly)| [header="X-Content-Type-Options: nosniff" && body="\<h1>Welcome to OpenResty!\</h1>" && header="X-Xss-Protection: 1; mode=block"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJYLUNvbnRlbnQtVHlwZS1PcHRpb25zOiBub3NuaWZmIiAmJiBib2R5PSI8aDE%2BV2VsY29tZSB0byBPcGVuUmVzdHkhPC9oMT4iICYmIGhlYWRlcj0iWC1Yc3MtUHJvdGVjdGlvbjogMTsgbW9kZT1ibG9jayI%3D) |
| **Number of assets affected**  | 5905 |
| **Description**  | Hikvision is a video-centric provider of intelligent IoT solutions and big data services. A command execution vulnerability exists in the operation and management center system of Hangzhou Hikvision Digital Technology Co. An attacker could use the vulnerability to gain server privileges. |
| **Impact** | The latest version has fixed the vulnerability, upgrade the system version to the latest version :https://www.hikvision.com/cn/19th-asian-games/isecure-center/?q=%E6%B5%B7%E5%BA%B7%E5%9F%9F%E8%A7%81%E7%BB%BC%E5%90%88%E5%AE%89%E9%98%B2%E7%AE%A1%E7%90%86%E5%B9%B3%E5%8F%B0 |

![](https://s3.bmp.ovh/imgs/2023/07/21/9e8983187d4816a8.gif)

## Netgod SecGate 3600 Firewall obj_area_import_save File Upload Vulnerability

|   **Vulnerability**  | **Netgod SecGate 3600 Firewall obj_area_import_save File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 网神 SecGate 3600 防火墙 obj_area_import_save 文件上传漏洞 |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [title="网神SecGate 3600防火墙"](https://en.fofa.info/result?qbase64=CXRpdGxlPSLnvZHnpZ5TZWNHYXRlIDM2MDDpmLLngavlopki) |
| **Number of assets affected**  | 725 |
| **Description**  | Netgod SecGate 3600 firewall is a composite hardware firewall based on status detection packet filtering and application level agents. It is a new generation of professional firewall equipment specially developed for large and medium-sized enterprises, governments, military, universities and other users. It supports external attack prevention, internal network security, network access control, network traffic monitoring and bandwidth management, dynamic routing, web content filtering, email content filtering, IP conflict detection and other functions, It can effectively ensure the security of the network; The product provides flexible network routing/bridging capabilities, supports policy routing and multi outlet link aggregation; It provides a variety of intelligent analysis and management methods, supports email alarm, supports log audit, provides comprehensive network management monitoring, and assists network administrators in completing network security management.There is a file upload vulnerability in SecGate 3600 firewall, which allows attackers to gain server control permissions. |
| **Impact** | There is a file upload vulnerability in SecGate 3600 firewall, which allows attackers to gain server control permissions. |

![](https://s3.bmp.ovh/imgs/2023/07/21/cfe7cdf43c6e18a4.gif)

## Netgod SecGate 3600 Firewall obj_area_import_save File Upload Vulnerability

|   **Vulnerability**  | **Netgod SecGate 3600 Firewall app_av_import_save File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 网神 SecGate 3600 防火墙 app_av_import_save 文件上传漏洞 |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [title="网神SecGate 3600防火墙"](https://en.fofa.info/result?qbase64=CXRpdGxlPSLnvZHnpZ5TZWNHYXRlIDM2MDDpmLLngavlopki) |
| **Number of assets affected**  | 725 |
| **Description**  | Netgod SecGate 3600 firewall is a composite hardware firewall based on status detection packet filtering and application level agents. It is a new generation of professional firewall equipment specially developed for large and medium-sized enterprises, governments, military, universities and other users. It supports external attack prevention, internal network security, network access control, network traffic monitoring and bandwidth management, dynamic routing, web content filtering, email content filtering, IP conflict detection and other functions, It can effectively ensure the security of the network; The product provides flexible network routing/bridging capabilities, supports policy routing and multi outlet link aggregation; It provides a variety of intelligent analysis and management methods, supports email alarm, supports log audit, provides comprehensive network management monitoring, and assists network administrators in completing network security management.There is a file upload vulnerability in SecGate 3600 firewall, which allows attackers to gain server control permissions. |
| **Impact** | There is a file upload vulnerability in SecGate 3600 firewall, which allows attackers to gain server control permissions. |

![](https://s3.bmp.ovh/imgs/2023/07/21/66573e8f0df89d78.gif)

## Kingdee Apusic Application Server deployApp Arbitrary File Upload Vulnerability

|   **Vulnerability**  | **Kingdee Apusic Application Server deployApp Arbitrary File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | Apusic应用服务器 deployApp 任意文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="Apusic应用服务器"](https://en.fofa.info/result?qbase64=dGl0bGU9IkFwdXNpY%2BW6lOeUqOacjeWKoeWZqCI%3D) |
| **Number of assets affected**  | 232 |
| **Description**  | Kingdee Apusic application server is the first pure Java application server in China with its own intellectual property rights following the J2EE standard.There is an arbitrary file upload vulnerability in the deployApp interface of the Kingdee Apusic application server. Attackers can use double slashes to bypass authentication and upload malicious compressed packages to take over server permissions. |
| **Impact** | There is an arbitrary file upload vulnerability in the deployApp interface of the Kingdee Apusic application server. Attackers can use double slashes to bypass authentication and upload malicious compressed packages to take over server permissions. |

## DOCBOX dynamiccontent.properties.xhtml Remote Code Execution Vulnerability

|   **Vulnerability**  | **DOCBOX dynamiccontent.properties.xhtml Remote Code Execution Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | DOCBOX dynamiccontent.properties.xhtml 文件 cmd 参数远程代码执行漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="docbox.webapp"](https://en.fofa.info/result?qbase64=Ym9keT0iZG9jYm94LndlYmFwcCI%3D) |
| **Number of assets affected**  | 657 |
| **Description**  | DOCBOX is a solution that can improve healthcare, is easy to use, and is based on a secure, open system.There is a code execution vulnerability in the javax.faces.resource of the DOCBOX system, and an attacker can execute arbitrary code to obtain server permissions. |
| **Impact** | There is a code execution vulnerability in the javax.faces.resource of the DOCBOX system, and an attacker can execute arbitrary code to obtain server permissions. |

## Kingdee-EAS easWebClient Arbitrary File Download Vulnerability

|   **Vulnerability**  | **Kingdee-EAS easWebClient Arbitrary File Download Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 金蝶-EAS easWebClient 任意文件下载漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="easSessionId" \|\| header="easportal" \|\| header="eassso/login" \|\| banner="eassso/login" \|\| body="/eassso/common" \|\| (title="EAS系统登录" && body="金蝶")](https://en.fofa.info/result?qbase64=Ym9keT0iZWFzU2Vzc2lvbklkIiB8fCBoZWFkZXI9ImVhc3BvcnRhbCIgfHwgaGVhZGVyPSJlYXNzc28vbG9naW4iIHx8IGJhbm5lcj0iZWFzc3NvL2xvZ2luIiB8fCBib2R5PSIvZWFzc3NvL2NvbW1vbiIgfHwgKHRpdGxlPSJFQVPns7vnu5%2FnmbvlvZUiICYmIGJvZHk9IumHkeidtiIp) |
| **Number of assets affected**  | 255 |
| **Description**  | Kingdee-EAS is a leading enterprise management system, which helps enterprises to build an integrated platform for industry, treasury, tax and invoice files, covering human resource management, tax management, financial sharing, procurement management, inventory management, production and manufacturing, etc.There is an arbitrary file reading vulnerability in Kingdee-EAS easWebClient, and attackers can read sensitive configuration file information such as config.jar. |
| **Impact** | There is an arbitrary file reading vulnerability in Kingdee-EAS easWebClient, and attackers can read sensitive configuration file information such as config.jar. |

## seeyon M1 Server userTokenService Code Execution Vulnerability

|   **Vulnerability**  | **seeyon M1 Server userTokenService Code Execution Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 致远 M1 移动端 userTokenService 代码执行漏洞 |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [title=="M1-Server 已启动"](https://en.fofa.info/result?qbase64=dGl0bGU9PSJNMS1TZXJ2ZXIg5bey5ZCv5YqoIg%3D%3D) |
| **Number of assets affected**  | 7050 |
| **Description**  | Seeyon M1 Server is a mobile device.Seeyon M1 Server userTokenService code execution vulnerability, attackers can arbitrarily execute code on the server side, write back door, obtain server permissions, and then control the entire web server. |
| **Impact** | Seeyon M1 Server userTokenService code execution vulnerability, attackers can arbitrarily execute code on the server side, write back door, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/07/14/d3a8bec6a5065f70.gif)

## Yonyou KSOA QueryService SQL Injection vulnerability

|   **Vulnerability**  | **Yonyou KSOA QueryService SQL Injection vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 用友时空 KSOA QueryService 处 content 参数 SQL 注入漏洞 |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [body="onmouseout=\"this.classname='btn btnOff'\""](https://en.fofa.info/result?qbase64=Ym9keT0ib25tb3VzZW91dD1cInRoaXMuY2xhc3NuYW1lPSdidG4gYnRuT2ZmJ1wiIg%3D%3D) |
| **Number of assets affected**  | 3995 |
| **Description**  | Yonyou KSOA spacetime is based on the KSOA concept under the guidance of research and development of a new generation of products, is according to the forefront of circulation enterprises IT requirements to launch the unification of the IT infrastructure, IT can make circulation enterprises established between IT systems in different historical periods, relaxed conversation with each other, help circulation enterprises to protect the existing IT investments, simplify IT management, enhance competition ability, Ensure that the overall strategic objectives and innovation activities of the enterprise are achieved. SQL injection vulnerability exists in some function of Yonyou spatio-temporal KSOA, which can be used by attackers to obtain database sensitive information. |
| **Impact** | In addition to using SQL injection vulnerability to obtain information in the database (for example, administrator background password, site user personal information), the attacker can even write Trojan horse to the server in the case of high permission to further obtain server system permission. |

## Qi An Xin Tianqing Terminal Security Management System information disclosure vulnerability

|   **Vulnerability**  | **Qi An Xin Tianqing Terminal Security Management System information disclosure vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 奇安信天擎终端安全管理系统信息泄露漏洞 |
| **CVSS core**  | 5.6 |
| **FOFA Query**  (click to view the results directly)| [title="新天擎"](https://en.fofa.info/result?qbase64=dGl0bGU9IuaWsOWkqeaTjiI%3D) |
| **Number of assets affected**  | 574 |
| **Description**  | Tianqing Terminal Security Management System is an integrated terminal security product solution for government and enterprise units.Tianqing Terminal Security Management System has an information disclosure vulnerability,the attacker reads the sensitive information of the system by constructing a special URL address. |
| **Impact** | Tianqing Terminal Security Management System has an information disclosure vulnerability,the attacker reads the sensitive information of the system by constructing a special URL address. |

## Tianqing terminal security management system YII_CSRF_TOKEN remote code execution vulnerability

|   **Vulnerability**  | **Tianqing terminal security management system YII_CSRF_TOKEN remote code execution vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 天擎终端安全管理系统 YII_CSRF_TOKEN 远程代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="360新天擎" \|\| body="appid\":\"skylar6" \|\| body="/task/index/detail?id={item.id}" \|\| body="已过期或者未授权，购买请联系4008-136-360" \|\| title="360天擎" \|\| title="360天擎终端安全管理系统"](https://en.fofa.info/result?qbase64=dGl0bGU9IjM2MOaWsOWkqeaTjiIgfHwgYm9keT0iYXBwaWRcIjpcInNreWxhcjYiIHx8IGJvZHk9Ii90YXNrL2luZGV4L2RldGFpbD9pZD17aXRlbS5pZH0iIHx8IGJvZHk9IuW3sui%2Fh%2Bacn%2BaIluiAheacquaOiOadg%2B%2B8jOi0reS5sOivt%2BiBlOezuzQwMDgtMTM2LTM2MCIgfHwgdGl0bGU9IjM2MOWkqeaTjiIgfHwgdGl0bGU9IjM2MOWkqeaTjue7iOerr%2BWuieWFqOeuoeeQhuezu%2Be7nyI%3D) |
| **Number of assets affected**  | 875 |
| **Description**  | Qi Anxin Tianqing is a terminal security management system (referred to as "Tianqing") product of Qi Anxin Group dedicated to integrated terminal security solutions.The web part of Qi'an Xintianqing terminal security management system uses the yii framework. This version of the framework has its own deserialization entry point, and the attacker can execute arbitrary code to obtain server permissions. |
| **Impact** | The web part of Qi'an Xintianqing terminal security management system uses the yii framework. This version of the framework has its own deserialization entry point, and the attacker can execute arbitrary code to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/07/14/fdc6987a22268e3b.gif)

## 91skzy Enterprise process control system login File Read vulnerability

|   **Vulnerability**  | **91skzy Enterprise process control system login File Read vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 时空智友企业流程化管控系统 login 文件读取漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="企业流程化管控系统" && body="密码(Password):"](https://en.fofa.info/result?qbase64=Ym9keT0i5LyB5Lia5rWB56iL5YyW566h5o6n57O757ufIiAmJiBib2R5PSLlr4bnoIEoUGFzc3dvcmQpOiI%3D) |
| **Number of assets affected**  | 1467 |
| **Description**  | Spatiotemporal Intelligent Friend enterprise process management and control system is a system that uses JAVA development to provide process management and control for enterprises.Spatiotemporal Zhiyou enterprise process control system login file read vulnerability, attackers can use the vulnerability to obtain sensitive information of the system. |
| **Impact** | Spatiotemporal Zhiyou enterprise process control system login file read vulnerability, attackers can use the vulnerability to obtain sensitive information of the system. |

## 91skzy Enterprise process control system formservice File Upload vulnerability

|   **Vulnerability**  | **91skzy Enterprise process control system formservice File Upload vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 时空智友企业流程化管控系统 formservice 文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="企业流程化管控系统" && body="密码(Password):"](https://en.fofa.info/result?qbase64=Ym9keT0i5LyB5Lia5rWB56iL5YyW566h5o6n57O757ufIiAmJiBib2R5PSLlr4bnoIEoUGFzc3dvcmQpOiI%3D) |
| **Number of assets affected**  | 1467 |
| **Description**  | Spatiotemporal Intelligent Friend enterprise process management and control system is a system that uses JAVA development to provide process management and control for enterprises.Spatiotemporal Zhiyou enterprise process control system formservice file upload vulnerability, attackers can use the vulnerability to obtain system permissions. |
| **Impact** | Spatiotemporal Zhiyou enterprise process control system formservice file upload vulnerability, attackers can use the vulnerability to obtain system permissions. |

![](https://s3.bmp.ovh/imgs/2023/07/14/5b49872a996cbd1e.gif)

## Glodon-Linkworks GetUserByEmployeeCode employeeCode SQL Injection Vulnerability

|   **Vulnerability**  | **Glodon-Linkworks GetUserByEmployeeCode employeeCode SQL Injection Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 广联达-Linkworks 协同办公管理平台 GetUserByEmployeeCode 文件 employeeCode 参数 SQL注入漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="Services/Identification/login.ashx" \|\| header="Services/Identification/login.ashx" \|\| banner="Services/Identification/login.ashx"](https://en.fofa.info/result?qbase64=Ym9keT0iU2VydmljZXMvSWRlbnRpZmljYXRpb24vbG9naW4uYXNoeCIgfHwgaGVhZGVyPSJTZXJ2aWNlcy9JZGVudGlmaWNhdGlvbi9sb2dpbi5hc2h4IiB8fCBiYW5uZXI9IlNlcnZpY2VzL0lkZW50aWZpY2F0aW9uL2xvZ2luLmFzaHgi) |
| **Number of assets affected**  | 27341 |
| **Description**  | Glodon-Linkworks collaborative office management platform is a management system that focuses on the entire life cycle of engineering projects and provides customers with digital software and hardware products and solutions.Glodon-Linkworks collaborative office management platform GetUserByEmployeeCode has a SQL injection vulnerability, and attackers can obtain sensitive information such as usernames and passwords. |
| **Impact** | Glodon-Linkworks collaborative office management platform GetUserByEmployeeCode has a SQL injection vulnerability, and attackers can obtain sensitive information such as usernames and passwords. |

![](https://s3.bmp.ovh/imgs/2023/07/14/30159400f31ca801.gif)

## Huatian-OA8000 MyHttpServlet reportFile Arbitrary File Upload Vulnerability

|   **Vulnerability**  | **Huatian-OA8000 MyHttpServlet reportFile Arbitrary File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 华天动力-OA8000 MyHttpServlet 文件 reportFile 参数文件上传漏洞 |
| **CVSS core**  | 8.6 |
| **FOFA Query**  (click to view the results directly)| [body="/OAapp/WebObjects/OAapp.woa"](https://en.fofa.info/result?qbase64=Ym9keT0iL09BYXBwL1dlYk9iamVjdHMvT0FhcHAud29hIg%3D%3D) |
| **Number of assets affected**  | 2226 |
| **Description**  | Huatian-OA8000 is a combination of advanced management ideas, management models, software technology and network technology, providing users with a low-cost, high-efficiency collaborative office and management platform.There is an arbitrary file upload vulnerability in Huatian Power OA MyHttpServlet. Attackers can upload malicious raq files and execute arbitrary sql statements in the raq files to obtain sensitive information such as user account passwords. |
| **Impact** | 	There is an arbitrary file upload vulnerability in Huatian Power OA MyHttpServlet. Attackers can upload malicious raq files and execute arbitrary sql statements in the raq files to obtain sensitive information such as user account passwords. |

![](https://s3.bmp.ovh/imgs/2023/07/07/ee0dff7305687815.gif)

## Ruijie WEB Management System EXCU_SHELL Information Disclosure Vulnerability

|   **Vulnerability**  | **Ruijie WEB Management System EXCU_SHELL Information Disclosure Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 锐捷交换机 WEB 管理系统 EXCU_SHELL 信息泄露漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="img/free_login_ge.gif" && body="./img/login_bg.gif"](https://en.fofa.info/result?qbase64=Ym9keT0iaW1nL2ZyZWVfbG9naW5fZ2UuZ2lmIiAmJiBib2R5PSIuL2ltZy9sb2dpbl9iZy5naWYi) |
| **Number of assets affected**  | 912 |
| **Description**  | Ruijie WEB management system is a switch device widely used in government, education, finance, medical and health care, and enterprises.Ruijie WEB management system EXCU_SHELL has an information leakage vulnerability, and attackers can obtain sensitive information such as system passwords to further control the system. |
| **Impact** | 	Ruijie WEB management system EXCU_SHELL has an information leakage vulnerability, and attackers can obtain sensitive information such as system passwords to further control the system. |

![](https://s3.bmp.ovh/imgs/2023/07/07/7bc1e0fd82aa53eb.gif)

## RSeeyou-OA wpsAssistServlet templateUrl Arbitrary File Read Vulnerability

|   **Vulnerability**  | **Seeyou-OA wpsAssistServlet templateUrl Arbitrary File Read Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 致远互联-OA wpsAssistServlet 文件 templateUrl 参数任意文件读取漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="/seeyon/USER-DATA/IMAGES/LOGIN/login.gif" \|\| title="用友致远A" \|\| (body="/yyoa/" && body!="本站内容均采集于") \|\| header="path=/yyoa" \|\| server=="SY8044" \|\| (body="A6-V5企业版" && body="seeyon" && body="seeyonProductId") \|\| (body="/seeyon/common/" && body="var _ctxpath = '/seeyon'") \|\| (body="A8-V5企业版" && body="/seeyon/") \|\| banner="Server: SY8044"](https://en.fofa.info/result?qbase64=Ym9keT0iL3NlZXlvbi9VU0VSLURBVEEvSU1BR0VTL0xPR0lOL2xvZ2luLmdpZiIgfHwgdGl0bGU9IueUqOWPi%2BiHtOi%2FnEEiIHx8IChib2R5PSIveXlvYS8iICYmIGJvZHkhPSLmnKznq5nlhoXlrrnlnYfph4fpm4bkuo4iKSB8fCBoZWFkZXI9InBhdGg9L3l5b2EiIHx8IHNlcnZlcj09IlNZODA0NCIgfHwgKGJvZHk9IkE2LVY15LyB5Lia54mIIiAmJiBib2R5PSJzZWV5b24iICYmIGJvZHk9InNlZXlvblByb2R1Y3RJZCIpIHx8IChib2R5PSIvc2VleW9uL2NvbW1vbi8iICYmIGJvZHk9InZhciBfY3R4cGF0aCA9ICcvc2VleW9uJyIpIHx8IChib2R5PSJBOC1WNeS8geS4mueJiCIgJiYgYm9keT0iL3NlZXlvbi8iKSB8fCBiYW5uZXI9IlNlcnZlcjogU1k4MDQ0Ig%3D%3D) |
| **Number of assets affected**  | 	53406 |
| **Description**  | Seeyou-OA is a collaborative office software that digitally builds the digital collaborative operation platform of enterprises and provides one-stop big data analysis solutions for various business scenarios of enterprises.Seeyou-OA wpsAssistServlet has arbitrary file reading vulnerabilities, and attackers can read sensitive information such as system passwords to further control the system. |
| **Impact** | 	Seeyou-OA wpsAssistServlet has arbitrary file reading vulnerabilities, and attackers can read sensitive information such as system passwords to further control the system. |

![](https://s3.bmp.ovh/imgs/2023/07/14/033b7613462dfe6a.gif)

## koronsoft AIO management system UtilServlet fileName File Read vulnerability

|   **Vulnerability**  | **koronsoft AIO management system UtilServlet fileName File Read vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 科荣 AIO 管理系统 UtilServlet 文件 fileName 参数文件读取漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="changeAccount('8000')"](https://en.fofa.info/result?qbase64=Ym9keT0iY2hhbmdlQWNjb3VudCgnODAwMCcpIg%3D%3D) |
| **Number of assets affected**  | 1976 |
| **Description**  | KoronsoftAIO management system is a very excellent enterprise management tool.The UtilServlet file reading vulnerability of koronsoftAIO management system can be used to obtain sensitive information of the system. |
| **Impact** | 	The UtilServlet file reading vulnerability ofkoronsoftAIO management system can be used to obtain sensitive information of the system. |

![](https://s3.bmp.ovh/imgs/2023/07/07/57f1d03fa857d0ea.gif)

## 91skzy Enterprise process control system wc.db Information Disclosure vulnerability

|   **Vulnerability**  | **91skzy Enterprise process control system wc.db Information Disclosure vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 时空智友企业流程化管控系统 wc.db 文件信息泄露漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="企业流程化管控系统" && body="密码(Password):"](https://en.fofa.info/result?qbase64=Ym9keT0i5LyB5Lia5rWB56iL5YyW566h5o6n57O757ufIiAmJiBib2R5PSLlr4bnoIEoUGFzc3dvcmQpOiI%3D) |
| **Number of assets affected**  | 1213 |
| **Description**  | Spatiotemporal Intelligent Friend enterprise process management and control system is a system that uses JAVA development to provide process management and control for enterprises.Spatiotemporal Wisdom enterprise process control system wc.db information leakage vulnerability, attackers can use the vulnerability to obtain sensitive information of the system. |
| **Impact** | 	Spatiotemporal Wisdom enterprise process control system wc.db information leakage vulnerability, attackers can use the vulnerability to obtain sensitive information of the system. |

![](https://s3.bmp.ovh/imgs/2023/07/07/894f6e466736766e.gif)

## Kingdee Cloud Starry Sky-Management Center Kingdee.BOS.ServiceFacade.ServicesStub.DevReportService.GetBusinessObjectData.common.kdsvc Arbitrary Code Execution Vulnerability

|   **Vulnerability**  | **Kingdee Cloud Starry Sky-Management Center Kingdee.BOS.ServiceFacade.ServicesStub.DevReportService.GetBusinessObjectData.common.kdsvc Arbitrary Code Execution Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 金蝶云星空 Kingdee.BOS.ServiceFacade.ServicesStub.DevReportService.GetBusinessObjectData.common.kdsvc 任意代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="金蝶云星空"](https://en.fofa.info/result?qbase64=CXRpdGxlPSLph5HonbbkupHmmJ%2Fnqboi) |
| **Number of assets affected**  | 6729 |
| **Description**  | Kingdee Cloud Starry Sky-Management Center is based on a leading assembleable low-code PaaS platform, which comprehensively serves customers' transformation in R&amp;D, production, marketing, supply chain, finance and other fields.There is a deserialization vulnerability in the Kingdee.BOS.ServiceFacade.ServicesStub.DevReportService.GetBusinessObjectData.common.kdsvc interface of Kingdee Cloud Star-Management Center, and an attacker can execute arbitrary commands to obtain server permissions. |
| **Impact** | There is a deserialization vulnerability in the Kingdee.BOS.ServiceFacade.ServicesStub.DevReportService.GetBusinessObjectData.common.kdsvc interface of Kingdee Cloud Star-Management Center, and an attacker can execute arbitrary commands to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/07/01/de00b732452eea00.gif)

## Chanjet T+ DownloadProxy.aspx Path File Read Vulnerability

|   **Vulnerability**  | **Chanjet T+ DownloadProxy.aspx Path File Read Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 畅捷通T+ DownloadProxy.aspx 文件 Path 参数文件读取漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="><script>location='/tplus/';\</script>\</body>" \|\| title=="畅捷通 T+"](https://en.fofa.info/result?qbase64=Ym9keT0iPjxzY3JpcHQ%2BbG9jYXRpb249Jy90cGx1cy8nOzwvc2NyaXB0PjwvYm9keT4iIHx8IHRpdGxlPT0i55WF5o236YCaIFQrIg%3D%3D) |
| **Number of assets affected**  | 112547 |
| **Description**  | Chanjet T+ is a smart, flexible and stylish enterprise management software based on the Internet era.Chanjet T+ DownloadProxy.aspx has arbitrary file reading vulnerabilities, and attackers can read sensitive information such as web.config to further control server permissions. |
| **Impact** | Chanjet T+ DownloadProxy.aspx has arbitrary file reading vulnerabilities, and attackers can read sensitive information such as web.config to further control server permissions. |

![](https://s3.bmp.ovh/imgs/2023/07/07/191a72363fbc96e6.gif)

## WordPress Plugin js-support-ticket File Upload Vulnerability

|   **Vulnerability**  | **WordPress Plugin js-support-ticket File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | WordPress js-support-ticket 插件 saveconfiguration 功能文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/js-support-ticket"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2pzLXN1cHBvcnQtdGlja2V0Ig%3D%3D) |
| **Number of assets affected**  | 1115 |
| **Description**  | JS Help Desk is a professional, simple, easy to use and complete customer support system. JS Help Desk comes packed with lot features than most of the expensive(and complex) support ticket system on market.JS Help Desk &lt;= 2.7.1 Unauthenticated Arbitrary File Upload. |
| **Impact** | An attacker can use the uploaded malicious script file to control the whole website or even control the server. This malicious script file, also known as WebShell, can also be referred to as a kind of web backdoor. WebShell scripts have very powerful functions, such as viewing server directories, files in the server, executing system commands, etc. |

## 91skzy Enterprise process control system formservice SQL Injection vulnerability

|   **Vulnerability**  | **91skzy Enterprise process control system formservice SQL Injection vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 时空智友企业流程化管控系统 formservice SQL 注入漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="企业流程化管控系统"](https://en.fofa.info/result?qbase64=Ym9keT0i5LyB5Lia5rWB56iL5YyW566h5o6n57O757ufIg%3D%3D) |
| **Number of assets affected**  | 1461 |
| **Description**  | Spatiotemporal Intelligent Friend enterprise process management and control system is a system that uses JAVA development to provide process management and control for enterprises.Spatiotemporal Wisdom enterprise process management and control system formservice SQL injection vulnerability, attackers can use the vulnerability to obtain sensitive database information. |
| **Impact** | Spatiotemporal Wisdom enterprise process management and control system formservice SQL injection vulnerability, attackers can use the vulnerability to obtain sensitive database information. |

## Hongfan OA zyy_AttFile.asmx File SQL Injection Vulnerability

|   **Vulnerability**  | **Hongfan OA zyy_AttFile.asmx File SQL Injection Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 红帆 OA zyy_AttFile.asmx 文件 fileName 参数 SQL 注入漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="iOffice.net" \|\| body="/iOffice/js" \|\| (body="iOffice.net" && header!="couchdb" && header!="drupal") \|\| body="iOfficeOcxSetup.exe" \|\| body="Hongfan. All Rights Reserved"](https://en.fofa.info/result?qbase64=dGl0bGU9ImlPZmZpY2UubmV0IiB8fCBib2R5PSIvaU9mZmljZS9qcyIgfHwgKGJvZHk9ImlPZmZpY2UubmV0IiAmJiBoZWFkZXIhPSJjb3VjaGRiIiAmJiBoZWFkZXIhPSJkcnVwYWwiKSB8fCBib2R5PSJpT2ZmaWNlT2N4U2V0dXAuZXhlIiB8fCBib2R5PSJIb25nZmFuLiBBbGwgUmlnaHRzIFJlc2VydmVkIg%3D%3D) |
| **Number of assets affected**  | 261 |
| **Description**  | Hongfan OA is an information management platform developed by Hongfan Technology based on the latest technology of Microsoft. NET. The Hongfan OA system provides the hospital with oA functions and completes administrative office services such as information release, process approval, document management, schedule management, work arrangement, document transfer, online communication, etc. Hongfan collaborative office system is the most professional and successful hospital OA in China. |
| **Impact** | There is a SQL injection vulnerability in Hongfan iOffice Hospital Edition, which can be used by attackers to obtain sensitive database information. |

## iOffice OA iorepsavexml.aspx Arbitrary File Upload Vulnerability

|   **Vulnerability**  | **iOffice OA iorepsavexml.aspx Arbitrary File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 红帆-ioffice iorepsavexml.aspx 任意文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="iOffice.net" \|\| body="/iOffice/js" \|\| (body="iOffice.net" && header!="couchdb" && header!="drupal") \|\| body="iOfficeOcxSetup.exe" \|\| body="Hongfan. All Rights Reserved"](https://en.fofa.info/result?qbase64=dGl0bGU9ImlPZmZpY2UubmV0IiB8fCBib2R5PSIvaU9mZmljZS9qcyIgfHwgKGJvZHk9ImlPZmZpY2UubmV0IiAmJiBoZWFkZXIhPSJjb3VjaGRiIiAmJiBoZWFkZXIhPSJkcnVwYWwiKSB8fCBib2R5PSJpT2ZmaWNlT2N4U2V0dXAuZXhlIiB8fCBib2R5PSJIb25nZmFuLiBBbGwgUmlnaHRzIFJlc2VydmVkIg%3D%3D) |
| **Number of assets affected**  | 261 |
| **Description**  | Hongfan OA is an oA function that provides hospitals with information release, process approval, document management, schedule management, work arrangement, file delivery, online communication and other administrative office services.There is an arbitrary file upload vulnerability in the Hongfan OA iorepsavexml.aspx file. Attackers can upload malicious Trojan horses to obtain server permissions. |
| **Impact** | There is an arbitrary file upload vulnerability in the Hongfan OA iorepsavexml.aspx file. Attackers can upload malicious Trojan horses to obtain server permissions. |

## WordPress plugin perfect survey admin-ajax.php question_id SQL Injection Vulnerability (CVE-2021-24762)

|   **Vulnerability**  | **WordPress plugin perfect survey admin-ajax.php question_id SQL Injection Vulnerability (CVE-2021-24762)**  |
| :----:   | :-----|
| **Chinese name**  | WordPress perfect survey 插件 admin-ajax.php 文件 question_id 参数 SQL注入漏洞（CVE-2021-24762） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="/wp-content/plugins/perfect-survey"](https://en.fofa.info/result?qbase64=Ym9keT0iL3dwLWNvbnRlbnQvcGx1Z2lucy9wZXJmZWN0LXN1cnZleSI%3D) |
| **Number of assets affected**  | 628 |
| **Description**  | WordPress plugin perfect survey is a plugin for surveying user feedback issues.WordPress plugin perfect survey version before 1.5.2 has a SQL injection vulnerability, the vulnerability stems from the lack of validation of externally input SQL statements in database-based applications. Attackers can exploit this vulnerability to execute illegal SQL commands to obtain sensitive information such as user passwords. |
| **Impact** | WordPress plugin perfect survey version before 1.5.2 has a SQL injection vulnerability, the vulnerability stems from the lack of validation of externally input SQL statements in database-based applications. Attackers can exploit this vulnerability to execute illegal SQL commands to obtain sensitive information such as user passwords. |

![](https://s3.bmp.ovh/imgs/2023/07/13/7a534dcb3646e5c6.gif)

## WordPress Plugin QuadMenu admin-ajax.php output File Upload Vulnerability

|   **Vulnerability**  | **WordPress Plugin QuadMenu admin-ajax.php output File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | WordPress QuadMenu 插件 admin-ajax.php 文件 output 参数任意文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/quadmenu"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3F1YWRtZW51Ig%3D%3D) |
| **Number of assets affected**  | 7573 |
| **Description**  | WordPress plugin perfect survey is a plugin for surveying user feedback issues.WordPress plugin perfect survey version before 1.5.2 has a SQL injection vulnerability, the vulnerability stems from the lack of validation of externally input SQL statements in database-based applications. Attackers can exploit this vulnerability to execute illegal SQL commands to obtain sensitive information such as user passwords. |
| **Impact** | WordPress plugin perfect survey version before 1.5.2 has a SQL injection vulnerability, the vulnerability stems from the lack of validation of externally input SQL statements in database-based applications. Attackers can exploit this vulnerability to execute illegal SQL commands to obtain sensitive information such as user passwords. |

## WordPress Plugin SuperStoreFinder-wp import.php File Upload Vulnerability

|   **Vulnerability**  | **WordPress Plugin SuperStoreFinder-wp import.php File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | WordPress SuperStoreFinder-wp 插件 import.php 任意文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/superstorefinder-wp"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3N1cGVyc3RvcmVmaW5kZXItd3Ai) |
| **Number of assets affected**  | 2363 |
| **Description**  | WordPress Plugin SuperStoreFinder-wp is a plugin with precise geolocation built in to let customers route and reach your store outlets in the easiest way.The WordPress Plugin SuperStoreFinder-wp plugin does not properly check file uploads. An attacker can set the Content-Type header to text/csv and use double extensions to bypass the existing checks. An attacker can upload malicious files to gain server permissions. |
| **Impact** | The WordPress Plugin SuperStoreFinder-wp plugin does not properly check file uploads. An attacker can set the Content-Type header to text/csv and use double extensions to bypass the existing checks. An attacker can upload malicious files to gain server permissions. |

## WordPress Plugin WP Hotel Booking thimpress_hotel_booking_1 RCE Vulnerability (CVE-2020-29047)

|   **Vulnerability**  | **WordPress Plugin WP Hotel Booking thimpress_hotel_booking_1 RCE Vulnerability (CVE-2020-29047)**  |
| :----:   | :-----|
| **Chinese name**  | WordPress 插件 WP Hotel Booking thimpress_hotel_booking_1 参数远程代码执行漏洞（CVE-2020-29047） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/wp-hotel-booking"](https://en.fofa.info/result?qbase64=CWJvZHk9IndwLWNvbnRlbnQvcGx1Z2lucy93cC1ob3RlbC1ib29raW5nIg%3D%3D) |
| **Number of assets affected**  | 1940 |
| **Description**  | WordPress Plugin WP Hotel Booking is a complete hotel booking plugin.WordPress Plugin WP Hotel Booking version 1.10.2 has a code execution vulnerability, and attackers can execute malicious code to control the server. |
| **Impact** | WordPress Plugin WP Hotel Booking version 1.10.2 has a code execution vulnerability, and attackers can execute malicious code to control the server. |

## EasyCVR default password

|   **Vulnerability**  | **EasyCVR default password**  |
| :----:   | :-----|
| **Chinese name**  | EasyCVR智能边缘网关默认口令漏洞 |
| **CVSS core**  | 8.6 |
| **FOFA Query**  (click to view the results directly)| [body="EasyGBS" \|\| body="EasyDarwin.Body" \|\| body="EasyCVR"](https://en.fofa.info/result?qbase64=Ym9keT0iRWFzeUdCUyIgfHwgYm9keT0iRWFzeURhcndpbi5Cb2R5IiB8fCBib2R5PSJFYXN5Q1ZSIg%3D%3D) |
| **Number of assets affected**  | 25111 |
| **Description**  | EasyCVR intelligent edge gateway is a product of TSINGSEE's software and hardware integration, which can provide multi-protocol (RTSP/RTMP/GB28181/Haikang Ehome/Dahua, Haikang SDK, etc.) device video access, Collection, AI intelligent detection, processing, distribution and other services.The EasyCVR video management platform has a default password vulnerability, where attackers can exploit the default credentials "easycvr/easycvr" to log in to the administration interface. |
| **Impact** | The EasyCVR video management platform has a default password vulnerability, where attackers can exploit the default credentials "easycvr/easycvr" to log in to the administration interface. |

## Apache Solr Velocity Template Injection Vulnerability (CVE-2019-17558)

|   **Vulnerability**  | **Apache Solr Velocity Template Injection Vulnerability (CVE-2019-17558)**  |
| :----:   | :-----|
| **Chinese name**  | Apache Solr Velocity 模版注入漏洞（CVE-2019-17558） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title="Solr Admin" \|\| body="SolrCore Initialization Failures" \|\| body="app_config.solr_path" \|\| (banner="/solr/" && banner="Location" && banner!="couchdb" && banner!="drupal")](https://en.fofa.info/result?qbase64=dGl0bGU9IlNvbHIgQWRtaW4iIHx8IGJvZHk9IlNvbHJDb3JlIEluaXRpYWxpemF0aW9uIEZhaWx1cmVzIiB8fCBib2R5PSJhcHBfY29uZmlnLnNvbHJfcGF0aCIgfHwgKGJhbm5lcj0iL3NvbHIvIiAmJiBiYW5uZXI9IkxvY2F0aW9uIiAmJiBiYW5uZXIhPSJjb3VjaGRiIiAmJiBiYW5uZXIhPSJkcnVwYWwiKQ%3D%3D) |
| **Number of assets affected**  | 1128540 |
| **Description**  | Apache Solr is a search server based on Lucene, developed by the Apache Software Foundation. The software supports features such as faceted search, vertical search, and highlighting of search results.A vulnerability has been discovered in Apache Solr versions 5.0.0 to 8.3.1 that allows injection attacks. Attackers can exploit this vulnerability using Velocity templates to execute arbitrary code on the system. |
| **Impact** | A vulnerability has been discovered in Apache Solr versions 5.0.0 to 8.3.1 that allows injection attacks. Attackers can exploit this vulnerability using Velocity templates to execute arbitrary code on the system. |

## 红帆-ioffice ioAssistance2.asmx 文件 SQL 注入漏洞 

|   **Vulnerability**  | **红帆-ioffice ioAssistance2.asmx 文件 SQL 注入漏洞 **  |
| :----:   | :-----|
| **Chinese name**  | 红帆-ioffice ioAssistance2.asmx 文件 SQL 注入漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="iOffice.net" \|\| body="/iOffice/js" \|\| (body="iOffice.net" && header!="couchdb" && header!="drupal") \|\| body="iOfficeOcxSetup.exe" \|\| body="Hongfan. All Rights Reserved"](https://en.fofa.info/result?qbase64=dGl0bGU9IlNvbHIgQWRtaW4iIHx8IGJvZHk9IlNvbHJDb3JlIEluaXRpYWxpemF0aW9uIEZhaWx1cmVzIiB8fCBib2R5PSJhcHBfY29uZmlnLnNvbHJfcGF0aCIgfHwgKGJhbm5lcj0iL3NvbHIvIiAmJiBiYW5uZXI9IkxvY2F0aW9uIiAmJiBiYW5uZXIhPSJjb3VjaGRiIiAmJiBiYW5uZXIhPSJkcnVwYWwiKQ%3D%3D) |
| **Number of assets affected**  | 261 |
| **Description**  | Hongfan OA is an oA function that provides hospitals with information release, process approval, document management, schedule management, work arrangement, file delivery, online communication and other administrative office services.There is a SQL injection vulnerability in the Hongfan OA ioAssistance2.asmx file. An attacker can obtain information such as database passwords and execute commands to obtain server permissions. |
| **Impact** | There is a SQL injection vulnerability in the Hongfan OA ioAssistance2.asmx file. An attacker can obtain information such as database passwords and execute commands to obtain server permissions. |

## VNC remote desktop system default password vulnerability

|   **Vulnerability**  | **VNC remote desktop system default password vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | VNC 远程桌面系统默认口令漏洞 |
| **CVSS core**  | 7.3 |
| **FOFA Query**  (click to view the results directly)| [(protocol="vnc" \|\| body="<APPLET code=VncViewer.class archive=VncViewer.jar")](https://en.fofa.info/result?qbase64=KHByb3RvY29sPSJ2bmMiIHx8IGJvZHk9IjxBUFBMRVQgY29kZT1WbmNWaWV3ZXIuY2xhc3MgYXJjaGl2ZT1WbmNWaWV3ZXIuamFyIik%3D) |
| **Number of assets affected**  | 2611849 |
| **Description**  | VNC is a screen sharing and remote operation software using RFB protocol. This software can send keyboard and mouse movements and real-time screen images through the network. VNC has nothing to do with the operating system, so it can be used across platforms, for example, you can use Windows to connect to a Linux computer, and vice versa. Even in a computer without a client program installed, as long as there is a browser that supports JAVA, it can be used. The product has weak passwords, and attackers can use the 123456 password to enter the system, view system information, and modify system configuration, which affects the use of users. |
| **Impact** | The product has weak passwords, and attackers can use the 123456 password to enter the system, view system information, and modify system configuration, which affects the use of users. |

## WordPress plugin Membership wps_membership_csv_file_upload File Upload Vulnerability (CVE-2022-4395)

|   **Vulnerability**  | **WordPress plugin Membership wps_membership_csv_file_upload File Upload Vulnerability (CVE-2022-4395)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Membership 插件 wps_membership_csv_file_upload 文件上传漏洞（CVE-2022-4395） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/Membership"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL01lbWJlcnNoaXAi) |
| **Number of assets affected**  | 1151 |
| **Description**  | WordPress plugin Membership is a plugin that helps attract customers using membership programs, offers users membership programs as subscriptions to limit access to your eCommerce store. WordPress plugin Membership versions before 2.1.7 have a code problem vulnerability, which is caused by not verifying uploaded files. Attackers exploit this vulnerability to upload arbitrary files, such as malicious PHP code, and execute code remotely. |
| **Impact** | WordPress plugin Membership versions before 2.1.7 have a code problem vulnerability, which is caused by not verifying uploaded files. Attackers exploit this vulnerability to upload arbitrary files, such as malicious PHP code, and execute code remotely. |

## WordPress Theme Motor File Inclusion Vulnerability(CVE-2021-24375)

|   **Vulnerability**  | **WordPress Theme Motor File Inclusion Vulnerability(CVE-2021-24375)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Motor 主题 admin-ajax.php 文件包含漏洞（CVE-2021-24375） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/themes/motor"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC90aGVtZXMvbW90b3Ii) |
| **Number of assets affected**  | 711 |
| **Description**  | Motor is a professional WordPress WooCommerce Theme for dealers, retailers, shops and mechanics. WordPress Motor  Theme < 3.1.0 is vulnerable to Local File Inclusion. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

![](https://s3.bmp.ovh/imgs/2023/06/28/56a22a12996a7f9c.gif)

## Arbitrary file reading vulnerability of edusoho classroom-course-statistics（CNVD-2023-03903）

|   **Vulnerability**  | **Arbitrary file reading vulnerability of edusoho classroom-course-statistics（CNVD-2023-03903）**  |
| :----:   | :-----|
|  **Chinese name**  | edusoho 教培系统 classroom-course-statistics 任意文件读取漏洞（CNVD-2023-03903） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [title="Powered By EduSoho" \|\| body="Powered by \<a href=\"http://www.edusoho.com/\" target=\"_blank\">EduSoho" \|\| (body="Powered By EduSoho" && body="var app")](https://en.fofa.info/result?qbase64=dGl0bGU9IlBvd2VyZWQgQnkgRWR1U29obyIgfHwgYm9keT0iUG93ZXJlZCBieSA8YSBocmVmPVwiaHR0cDovL3d3dy5lZHVzb2hvLmNvbS9cIiB0YXJnZXQ9XCJfYmxhbmtcIj5FZHVTb2hvIiB8fCAoYm9keT0iUG93ZXJlZCBCeSBFZHVTb2hvIiAmJiBib2R5PSJ2YXIgYXBwIik%3D) |
| **Number of assets affected**  | 6957 |
| **Description**  | The edusoho education and training system <v22.4.7 has unauthorized file reading vulnerability. Through this vulnerability, an attacker can read the contents of the config/parameters.yml file and obtain sensitive information such as the secret value saved in the file and database account password. After the secret value is obtained, an attacker can implement RCE with symfony _fragment routing |
| **Impact** | EduSoho Education and training system is an open source network school system developed by Hangzhou Kozhi Network Technology Company. The education and training system <v22.4.7 has unauthorized arbitrary file reading vulnerability, through which an attacker can read the contents of the config/parameters.yml file. Get the secret value saved in the file, database account password and other sensitive information. After the secret value is obtained, an attacker can implement RCE with symfony _fragment routing |

![](https://s3.bmp.ovh/imgs/2023/06/28/2639f5440b5697a4.gif)

## dst-admin sendBroadcast message RCE Vulnerability (CVE-2023-0649)

|   **Vulnerability**  | **dst-admin sendBroadcast message RCE Vulnerability (CVE-2023-0649)**  |
| :----:   | :-----|
|  **Chinese name**  | dst-admin 饥荒管理后台 sendBroadcast 文件 message 参数远程命令执行漏洞（CVE-2023-0649） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=="饥荒管理后台"](https://en.fofa.info/result?qbase64=dGl0bGU9PSLppaXojZLnrqHnkIblkI7lj7Ai) |
| **Number of assets affected**  | 1999 |
| **Description**  | dst-admin is a web program written in Java language by the individual developer of qinming99. There is a command injection vulnerability in dst-admin version 1.5.0. The vulnerability comes from the unknown function of the file home/sendBroadcast. Command injection can be performed through the parameter message, and the attacker can obtain server privileges. |
| **Impact** | There is a command injection vulnerability in dst-admin version 1.5.0. The vulnerability comes from the unknown function of the file home/sendBroadcast. Command injection can be performed through the parameter message, and the attacker can obtain server privileges. |

## WordPress plugin Build App Online admin-ajax.php vendor SQL Vulnerability (CVE-2022-3241)

|   **Vulnerability**  | **WordPress plugin Build App Online admin-ajax.php vendor SQL Vulnerability (CVE-2022-3241)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Build App Online 插件 admin-ajax.php 文件 vendor 参数 SQL注入漏洞（CVE-2022-3241） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/build-app-online"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2J1aWxkLWFwcC1vbmxpbmUi) |
| **Number of assets affected**  | 327 |
| **Description**  | WordPress plugin Build App Online is a plugin that helps you create and run mobile apps for woocommerce. WordPress plugin Build App Online version before 1.0.19 has a SQL injection vulnerability. The vulnerability stems from the fact that some parameters are not properly cleaned and escaped before the SQL statement uses them, resulting in SQL injection, and the attacker can obtain sensitive information such as account passwords. information. |
| **Impact** | WordPress plugin Build App Online version before 1.0.19 has a SQL injection vulnerability. The vulnerability stems from the fact that some parameters are not properly cleaned and escaped before the SQL statement uses them, resulting in SQL injection, and the attacker can obtain sensitive information such as account passwords. information. |

## WordPress plugin Nirweb support admin-ajax.php id_form SQL Injection Vulnerability (CVE-2022-0781)

|   **Vulnerability**  | **WordPress plugin Nirweb support admin-ajax.php id_form SQL Injection Vulnerability (CVE-2022-0781)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Nirweb support 插件 admin-ajax.php 文件 id_form 参数SQL注入漏洞（CVE-2022-0781） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/nirweb-support"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL25pcndlYi1zdXBwb3J0Ig%3D%3D) |
| **Number of assets affected**  | 408 |
| **Description**  | WordPress plugin Nirweb support is a plugin with sales, support, management and other functions. There is a SQL injection vulnerability in versions before WordPress plugin Nirweb support 2.8.2. The vulnerability stems from the failure to clean and escape parameters. Attackers exploiting this vulnerability can lead to SQL injection attacks. |
| **Impact** | There is a SQL injection vulnerability in versions before WordPress plugin Nirweb support 2.8.2. The vulnerability stems from the failure to clean and escape parameters. Attackers exploiting this vulnerability can lead to SQL injection attacks. |

## WordPress theme Listingo listingo_temp_uploader File Upload Vulnerability (CVE-2022-3921)

|   **Vulnerability**  | **WordPress theme Listingo listingo_temp_uploader File Upload Vulnerability (CVE-2022-3921)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress 主题 Listingo listingo_temp_uploader 功能任意文件上传漏洞（CVE-2022-3921） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/themes/listingo"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC90aGVtZXMvbGlzdGluZ28i) |
| **Number of assets affected**  | 209 |
| **Description**  | WordPress theme Listingo is a WordPress theme for displaying business listings and directories. WordPress theme Listingo version before 3.2.7 has a code problem vulnerability. The vulnerability is caused by not verifying the files to be uploaded through AJAX operations. Attackers can upload malicious webshells to obtain server permissions. |
| **Impact** | WordPress theme Listingo version before 3.2.7 has a code problem vulnerability. The vulnerability is caused by not verifying the files to be uploaded through AJAX operations. Attackers can upload malicious webshells to obtain server permissions. |

## dst-admin cavesConsole RCE Vulnerability (CVE-2023-0646)

|   **Vulnerability**  | **dst-admin cavesConsole RCE Vulnerability (CVE-2023-0646)**  |
| :----:   | :-----|
|  **Chinese name**  | dst-admin cavesConsole 远程命令执行漏洞（CVE-2023-0646） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=="饥荒管理后台"](https://en.fofa.info/result?qbase64=dGl0bGU9PSLppaXojZLnrqHnkIblkI7lj7Ai) |
| **Number of assets affected**  | 1980 |
| **Description**  | dst-admin is a web program written in Java language by the individual developer of qinming99.There is a command injection vulnerability in dst-admin version 1.5.0. The vulnerability comes from the unknown function of the file home/cavesConsole. Command injection can be performed through the parameter command, and the attacker can obtain server privileges. |
| **Impact** | There is a command injection vulnerability in dst-admin version 1.5.0. The vulnerability comes from the unknown function of the file home/cavesConsole. Command injection can be performed through the parameter command, and the attacker can obtain server privileges.|

## dst-admin kickPlayer RCE Vulnerability (CVE-2023-0647)

|   **Vulnerability**  | **dst-admin kickPlayer RCE Vulnerability (CVE-2023-0647)**  |
| :----:   | :-----|
|  **Chinese name**  | dst-admin kickPlayer 远程命令执行漏洞（CVE-2023-0647） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=="饥荒管理后台"](https://en.fofa.info/result?qbase64=dGl0bGU9PSLppaXojZLnrqHnkIblkI7lj7Ai) |
| **Number of assets affected**  | 1980 |
| **Description**  | dst-admin is a web program written in Java language by the individual developer of qinming99.There is a command injection vulnerability in dst-admin version 1.5.0. The vulnerability comes from the unknown function of the file home/cavesConsole. Command injection can be performed through the parameter command, and the attacker can obtain server privileges. |
| **Impact** | There is a command injection vulnerability in dst-admin version 1.5.0. The vulnerability comes from the unknown function of the file home/cavesConsole. Command injection can be performed through the parameter command, and the attacker can obtain server privileges.|

## 	VNC remote desktop system week password vulnerability

|   **Vulnerability**  | **VNC remote desktop system week password vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | VNC 远程桌面系统弱口令漏洞 |
| **CVSS core**  | 7.3 |
| **FOFA Query**  (click to view the results directly)| [(protocol="vnc" \|\| body="<APPLET code=VncViewer.class archive=VncViewer.jar")](https://en.fofa.info/result?qbase64=KHByb3RvY29sPSJ2bmMiIHx8IGJvZHk9IjxBUFBMRVQgY29kZT1WbmNWaWV3ZXIuY2xhc3MgYXJjaGl2ZT1WbmNWaWV3ZXIuamFyIik%3D) |
| **Number of assets affected**  | 2611849 |
| **Description**  | VNC is a screen sharing and remote operation software using RFB protocol. This software can send keyboard and mouse movements and real-time screen images through the network. VNC has nothing to do with the operating system, so it can be used across platforms, for example, you can use Windows to connect to a Linux computer, and vice versa. Even in a computer without a client program installed, as long as there is a browser that supports JAVA, it can be used. The product has weak passwords, and attackers can use the 123456 password to enter the system, view system information, and modify system configuration, which affects the use of users. |
| **Impact** | The product has weak passwords, and attackers can use the 123456 password to enter the system, view system information, and modify system configuration, which affects the use of users. |

## VMware VRealize Network Insight resttosaasservlet Remote Command Execution Vulnerability (CVE-2023-20887)

|   **Vulnerability**  | **VMware VRealize Network Insight resttosaasservlet Remote Command Execution Vulnerability (CVE-2023-20887)**  |
| :----:   | :-----|
| **Chinese name**  | VMware VRealize Network Insight resttosaasservlet 远程命令执行漏洞（CVE-2023-20887） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="VMware vRealize Network Insight" \|\| body="vneraapp/assets/fonts/bootstrap/glyphicons-halflings-regular" \|\| title="Operations for Networks"](https://en.fofa.info/result?qbase64=dGl0bGU9IlZNd2FyZSB2UmVhbGl6ZSBOZXR3b3JrIEluc2lnaHQiIHx8IGJvZHk9InZuZXJhYXBwL2Fzc2V0cy9mb250cy9ib290c3RyYXAvZ2x5cGhpY29ucy1oYWxmbGluZ3MtcmVndWxhciIgfHwgdGl0bGU9Ik9wZXJhdGlvbnMgZm9yIE5ldHdvcmtzIg%3D%3D) |
| **Number of assets affected**  | 32 |
| **Description**  | VMware Aria Operations is a unified, AI-driven autonomous IT operations management platform from VMware Inc. It is designed for private cloud, hybrid cloud, and multi-cloud environments.A security vulnerability exists in the /saas./resttosaasservlet component of VMware Aria Operations Networks 6.x series versions, which allows attackers to execute command injection attacks and subsequently result in remote code execution. |
| **Impact** | A security vulnerability exists in the saasresttosaasservlet component of VMware Aria Operations Networks 6.x series versions, which allows attackers to execute command injection attacks and subsequently result in remote code execution. |

![](https://github.com/jwy456829/Goby/blob/master/CVE-2023-20887.gif)

## VMware VRealize Network Insight saasresttosaasservlet Remote Command Execution Vulnerability (CVE-2022-31702)

|   **Vulnerability**  | **VMware VRealize Network Insight saasresttosaasservlet Remote Command Execution Vulnerability (CVE-2022-31702)**  |
| :----:   | :-----|
| **Chinese name**  | VMware VRealize Network Insight saasresttosaasservlet 远程命令执行漏洞（CVE-2022-31702） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="VMware vRealize Network Insight" \|\| body="vneraapp/assets/fonts/bootstrap/glyphicons-halflings-regular" \|\| title="Operations for Networks"](https://en.fofa.info/result?qbase64=dGl0bGU9IlZNd2FyZSB2UmVhbGl6ZSBOZXR3b3JrIEluc2lnaHQiIHx8IGJvZHk9InZuZXJhYXBwL2Fzc2V0cy9mb250cy9ib290c3RyYXAvZ2x5cGhpY29ucy1oYWxmbGluZ3MtcmVndWxhciIgfHwgdGl0bGU9Ik9wZXJhdGlvbnMgZm9yIE5ldHdvcmtzIg%3D%3D) |
| **Number of assets affected**  | 32 |
| **Description**  | VMware Aria Operations is a unified, AI-driven autonomous IT operations management platform from VMware Inc. It is designed for private cloud, hybrid cloud, and multi-cloud environments.A security vulnerability exists in the saasresttosaasservlet component of VMware Aria Operations Networks 6.x series versions, which allows attackers to execute command injection attacks and subsequently result in remote code execution. |
| **Impact** | A security vulnerability exists in the saasresttosaasservlet component of VMware Aria Operations Networks 6.x series versions, which allows attackers to execute command injection attacks and subsequently result in remote code execution. |

![](https://s3.bmp.ovh/imgs/2023/06/16/ab6c2f05e446d56a.gif)

## Ignite Realtime Openfire Permission Bypass Vulnerability (CVE-2023-32315)

|   **Vulnerability**  | **Ignite Realtime Openfire Permission Bypass Vulnerability (CVE-2023-32315)**  |
| :----:   | :-----|
| **Chinese name**  | Ignite Realtime Openfire 权限绕过漏洞（CVE-2023-32315） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="background: transparent url(images/login_logo.gif) no-repeat" && body="Openfire") \|\| (body="class=\"row justify-content-center\"" && body="\<title>Openfire 管理界面\</title>") \|\| title="Openfire Admin Console" \|\| title="Openfire HTTP Binding Service" \|\| (body="align=\"right\" id=\"jive-loginVersion" && body="Openfire") \|\| title="Открытый огонь Консоль Администрации" \|\| title=="Openfire 管理界面"](https://en.fofa.info/result?qbase64=KGJvZHk9ImJhY2tncm91bmQ6IHRyYW5zcGFyZW50IHVybChpbWFnZXMvbG9naW5fbG9nby5naWYpIG5vLXJlcGVhdCIgJiYgYm9keT0iT3BlbmZpcmUiKSB8fCAoYm9keT0iY2xhc3M9XCJyb3cganVzdGlmeS1jb250ZW50LWNlbnRlclwiIiAmJiBib2R5PSI8dGl0bGU%2BT3BlbmZpcmUg566h55CG55WM6Z2iPC90aXRsZT4iKSB8fCB0aXRsZT0iT3BlbmZpcmUgQWRtaW4gQ29uc29sZSIgfHwgdGl0bGU9Ik9wZW5maXJlIEhUVFAgQmluZGluZyBTZXJ2aWNlIiB8fCAoYm9keT0iYWxpZ249XCJyaWdodFwiIGlkPVwiaml2ZS1sb2dpblZlcnNpb24iICYmIGJvZHk9Ik9wZW5maXJlIikgfHwgdGl0bGU9ItCe0YLQutGA0YvRgtGL0Lkg0L7Qs9C%2B0L3RjCDQmtC%2B0L3RgdC%2B0LvRjCDQkNC00LzQuNC90LjRgdGC0YDQsNGG0LjQuCIgfHwgdGl0bGU9PSJPcGVuZmlyZSDnrqHnkIbnlYzpnaIi) |
| **Number of assets affected**  | 49936 |
| **Description**  | Apache RocketMQ is a lightweight data processing platform and messaging engine developed by the Apache Software Foundation in the United States. There is a code injection vulnerability in Apache RocketMQ 5.1.0 and earlier versions, which originates from a remote command execution vulnerability. Attackers can exploit this vulnerability to execute commands with system user privileges using the update configuration function. |
| **Impact** | There is a code injection vulnerability in Apache RocketMQ 5.1.0 and earlier versions, which originates from a remote command execution vulnerability. Attackers can exploit this vulnerability to execute commands with system user privileges using the update configuration function. |

![](https://s3.bmp.ovh/imgs/2023/06/16/ab6c2f05e446d56a.gif)

## RocketMQ Broker rocketmqHome Config Remote Command Execution Vulnerability (CVE-2023-33246)

|   **Vulnerability**  | **RocketMQ Broker rocketmqHome Config Remote Command Execution Vulnerability (CVE-2023-33246)**  |
| :----:   | :-----|
| **Chinese name**  | RocketMQ Broker rocketmqHome Config 远程命令执行漏洞（CVE-2023-33246） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [protocol="rocketmq-broker"](https://en.fofa.info/result?qbase64=cHJvdG9jb2w9InJvY2tldG1xLWJyb2tlciI%3D) |
| **Number of assets affected**  | 10087 |
| **Description**  | Apache RocketMQ is a lightweight data processing platform and messaging engine developed by the Apache Software Foundation in the United States. There is a code injection vulnerability in Apache RocketMQ 5.1.0 and earlier versions, which originates from a remote command execution vulnerability. Attackers can exploit this vulnerability to execute commands with system user privileges using the update configuration function. |
| **Impact** | There is a code injection vulnerability in Apache RocketMQ 5.1.0 and earlier versions, which originates from a remote command execution vulnerability. Attackers can exploit this vulnerability to execute commands with system user privileges using the update configuration function. |
  
![](https://github.com/gobysec/GobyVuls/blob/master/CVE-2023-33246.gif)

## Telecom system /manager/teletext/material/upload.php fileupload vulnerability

|   **Vulnerability**  | **Telecom system /manager/teletext/material/upload.php fileupload vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 电信网关配置管理系统后台 /manager/teletext/material/upload.php 文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="img/login_bg3.png" && body="系统登录"](https://en.fofa.info/result?qbase64=Ym9keT0iaW1nL2xvZ2luX2JnMy5wbmciICYmIGJvZHk9Iuezu%2Be7n%2BeZu%2BW9lSI%3D) |
| **Number of assets affected**  | 856 |
| **Description**  | China Telecom Group Co., Ltd. (English name \"China Telecom\", referred to as \"China Telecom\") was established in September 2000. It is a large state-owned telecommunications company in China and a global partner of the Shanghai World Expo. There is a file upload vulnerability in the background of the telecom gateway configuration management system. An attacker can exploit this vulnerability to obtain a device shell.|
| **Impact** | China Telecom Group Co., Ltd. (English name \"China Telecom\", referred to as \"China Telecom\") was established in September 2000. It is a large state-owned telecommunications company in China and a global partner of the Shanghai World Expo. There is a file upload vulnerability in the background of the telecom gateway configuration management system. An attacker can exploit this vulnerability to obtain a device shell. |

![](https://s3.bmp.ovh/imgs/2023/06/21/dd91a5dac2dab90d.gif)

## Weaver E-office Office Automation System download.php filename file read vulnerability

|   **Vulnerability**  | **Weaver E-office Office Automation System download.php filename file read vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 泛微E-office协同办公系统 download.php 文件 filename 参数文件读取漏洞 |
| **CVSS core**  | 5.0 |
| **FOFA Query**  (click to view the results directly)| [((header="general/login/index.php" \|\| body="/general/login/view//images/updateLoad.gif" \|\| (body="szFeatures" && body="eoffice") \|\| header="Server: eOffice") && body!="Server: couchdb") \|\| banner="general/login/index.php"](https://en.fofa.info/result?qbase64=KChoZWFkZXI9ImdlbmVyYWwvbG9naW4vaW5kZXgucGhwIiB8fCBib2R5PSIvZ2VuZXJhbC9sb2dpbi92aWV3Ly9pbWFnZXMvdXBkYXRlTG9hZC5naWYiIHx8IChib2R5PSJzekZlYXR1cmVzIiAmJiBib2R5PSJlb2ZmaWNlIikgfHwgaGVhZGVyPSJTZXJ2ZXI6IGVPZmZpY2UiKSAmJiBib2R5IT0iU2VydmVyOiBjb3VjaGRiIikgfHwgYmFubmVyPSJnZW5lcmFsL2xvZ2luL2luZGV4LnBocCI%3D) |
| **Number of assets affected**  | 4971 |
| **Description**  | Weaver E-office Office Automation System is a professional office software, is for small business or team work platform. There are file reading vulnerabilities in the Weaver E-office Office Automation System.Attackers can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., resulting in an extremely insecure state of the website. |
| **Impact** | There are file reading vulnerabilities in the Weaver E-office Office Automation System.Attackers can use this vulnerability to read important system files (such as database configuration files, system configuration files), database configuration files, etc., resulting in an extremely insecure state of the website. |

![](https://s3.bmp.ovh/imgs/2023/06/21/008f7969a69cf000.gif)

## WordPress Plugin Extensive VC Addons File Inclusion Vulnerability

|   **Vulnerability**  | **WordPress Plugin Extensive VC Addons File Inclusion Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | WordPress Extensive VC Addons 插件 options[template] 文件包含漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/extensive-vc-addon"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2V4dGVuc2l2ZS12Yy1hZGRvbiI%3D) |
| **Number of assets affected**  | 2583 |
| **Description**  | Extensive VC is a powerful WordPress tool which allows you to add unique, flexible and fully responsive shortcode elements on your site. xtensive VC Addons &lt; 1.9.1 is vulnerable to Local File Inclusion. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

![](https://s3.bmp.ovh/imgs/2023/06/21/10e1c4496ca1db1b.gif)

## Avaya Aura Device Services PhoneBackup File Upload Vulnerability

|   **Vulnerability**  | **Avaya Aura Device Services PhoneBackup File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | Avaya Aura Device Services r软件 PhoneBackup 任意文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [((body="vmsTitle\">Avaya Aura&#8482;&nbsp;Utility Server" \|\| body="/webhelp/Base/Utility_toc.htm" \|\| (body="Avaya Aura&reg;&nbsp;Utility Services" && body="Avaya Inc. All Rights Reserved")) && body!="Server: couchdb")](https://en.fofa.info/result?qbase64=KChib2R5PSJ2bXNUaXRsZVwiPkF2YXlhIEF1cmEmIzg0ODI7Jm5ic3A7VXRpbGl0eSBTZXJ2ZXIiIHx8IGJvZHk9Ii93ZWJoZWxwL0Jhc2UvVXRpbGl0eV90b2MuaHRtIiB8fCAoYm9keT0iQXZheWEgQXVyYSZyZWc7Jm5ic3A7VXRpbGl0eSBTZXJ2aWNlcyIgJiYgYm9keT0iQXZheWEgSW5jLiBBbGwgUmlnaHRzIFJlc2VydmVkIikpICYmIGJvZHkhPSJTZXJ2ZXI6IGNvdWNoZGIiKQ%3D%3D) |
| **Number of assets affected**  | 565 |
| **Description**  | Avaya Aura Device Services is an application software of Avaya Corporation in the United States. Provides a function to manage Avaya endpoints. Avaya Aura Device Services versions 7.0 to 8.1.4.0 have security vulnerabilities. Attackers can bypass authentication and upload arbitrary files to obtain server permissions. |
| **Impact** | Avaya Aura Device Services versions 7.0 to 8.1.4.0 have security vulnerabilities. Attackers can bypass authentication and upload arbitrary files to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/06/21/06ca1ac2bfc684e0.gif)

## Weaver OA PluginViewServlet Authentication Bypass Vulnerability

|   **Vulnerability**  | **Weaver OA PluginViewServlet Authentication Bypass Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 泛微OA办公系统 PluginViewServlet 认证绕过漏洞 |
| **CVSS core**  | 8.0 |
| **FOFA Query**  (click to view the results directly)| [(header="testBanCookie" \|\| banner="testBanCookie" \|\| body="/wui/common/css/w7OVFont.css" \|\| (body="typeof poppedWindow" && body="client/jquery.client_wev8.js") \|\| body="/theme/ecology8/jquery/js/zDialog_wev8.js" \|\| body="ecology8/lang/weaver_lang_7_wev8.js")](https://en.fofa.info/result?qbase64=KGhlYWRlcj0idGVzdEJhbkNvb2tpZSIgfHwgYmFubmVyPSJ0ZXN0QmFuQ29va2llIiB8fCBib2R5PSIvd3VpL2NvbW1vbi9jc3MvdzdPVkZvbnQuY3NzIiB8fCAoYm9keT0idHlwZW9mIHBvcHBlZFdpbmRvdyIgJiYgYm9keT0iY2xpZW50L2pxdWVyeS5jbGllbnRfd2V2OC5qcyIpIHx8IGJvZHk9Ii90aGVtZS9lY29sb2d5OC9qcXVlcnkvanMvekRpYWxvZ193ZXY4LmpzIiB8fCBib2R5PSJlY29sb2d5OC9sYW5nL3dlYXZlcl9sYW5nXzdfd2V2OC5qcyIp) |
| **Number of assets affected**  | 45034 |
| **Description**  | Weaver OA is a professional and powerful multi-functional office management software that supports mobile approval, attendance, query, sharing and other functions, effectively improving the user's office efficiency. There is an authentication bypass vulnerability in Panwei OA weaver.mobile.plugin.ecology.service.PluginViewServlet, and attackers can log in arbitrarily to obtain administrator privileges. |
| **Impact** | There is an authentication bypass vulnerability in Panwei OA weaver.mobile.plugin.ecology.service.PluginViewServlet, and attackers can log in arbitrarily to obtain administrator privileges. |

![](https://s3.bmp.ovh/imgs/2023/06/21/38cfb28426294995.gif)

## Netgod SecGate 3600 Firewall sys_export_conf_local_save File Read Vulnerability

|   **Vulnerability**  | **Netgod SecGate 3600 Firewall sys_export_conf_local_save File Read Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | 网神SecGate 3600防火墙 sys_export_conf_local_save 文件读取漏洞 |
| **CVSS core**  | 8.0 |
| **FOFA Query**  (click to view the results directly)| [title="网神SecGate 3600防火墙"](https://en.fofa.info/result?qbase64=dGl0bGU9Iue9keelnlNlY0dhdGUgMzYwMOmYsueBq%2BWimSI%3D) |
| **Number of assets affected**  | 738 |
| **Description**  | Netgod SecGate 3600 firewall is a composite hardware firewall based on status detection packet filtering and application level agents. It is a new generation of professional firewall equipment specially developed for large and medium-sized enterprises, governments, military, universities and other users. It supports external attack prevention, internal network security, network access control, network traffic monitoring and bandwidth management, dynamic routing, web content filtering, email content filtering, IP conflict detection and other functions, It can effectively ensure the security of the network; The product provides flexible network routing/bridging capabilities, supports policy routing and multi outlet link aggregation; It provides a variety of intelligent analysis and management methods, supports email alarm, supports log audit, provides comprehensive network management monitoring, and assists network administrators in completing network security management. There is a file reading vulnerability in the Netgod SecGate 3600 firewall, which allows attackers to obtain sensitive information from the server. |
| **Impact** | There is a file reading vulnerability in the Netgod SecGate 3600 firewall, which allows attackers to obtain sensitive information from the server. |

![](https://s3.bmp.ovh/imgs/2023/06/21/cffc80cace17c408.gif)

## PandoraFMS upload_head_image.php Arbitrary File Upload Vulnerability

|   **Vulnerability**  | **PandoraFMS upload_head_image.php Arbitrary File Upload Vulnerability**  |
| :----:   | :-----|
| **Chinese name**  | PandoraFMS 软件 upload_head_image.php 任意文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="pandora_console/"](https://en.fofa.info/result?qbase64=Ym9keT0icGFuZG9yYV9jb25zb2xlLyI%3D) |
| **Number of assets affected**  | 768 |
| **Description**  | PandoraFMS is an application software of American PandoraFMS. Provides a monitoring function.<br></p><p>There is an unauthorized file upload vulnerability in PandoraFMS upload_head_image.php. Attackers can upload malicious Trojan horses to obtain server permissions. |
| **Impact** | There is an unauthorized file upload vulnerability in PandoraFMS upload_head_image.php. Attackers can upload malicious Trojan horses to obtain server permissions. |

![](https://s3.bmp.ovh/imgs/2023/06/21/c89f35d6d3a29740.gif)

## WordPress plugins User Post Gallery upg_datatable RCE Vulnerability (CVE-2022-4060)

|   **Vulnerability**  | **WordPress plugins User Post Gallery upg_datatable RCE Vulnerability (CVE-2022-4060)**  |
| :----:   | :-----|
| **Chinese name**  | WordPress User Post Gallery 插件 upg_datatable 远程代码执行漏洞（CVE-2022-4060） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/wp-upg"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3dwLXVwZyI%3D) |
| **Number of assets affected**  | 383 |
| **Description**  | WordPress plugins User Post Gallery is a plugin that allows users to select albums, generate tags, upload pictures and videos from the front end. There is a code injection vulnerability in WordPress plugin User Post Gallery 2.19 and earlier versions. The vulnerability stems from the fact that the callback function allows any user to call it. Attackers can use this vulnerability to run code on its site. |
| **Impact** | There is a code injection vulnerability in WordPress plugin User Post Gallery 2.19 and earlier versions. The vulnerability stems from the fact that the callback function allows any user to call it. Attackers can use this vulnerability to run code on its site. |

![](https://s3.bmp.ovh/imgs/2023/06/21/f82bbe1d63c40785.gif)

## WordPress plugins User Verification Authentication Bypass Vulnerability (CVE-2022-4693)

|   **Vulnerability**  | **WordPress plugins User Verification Authentication Bypass Vulnerability (CVE-2022-4693)**  |
| :----:   | :-----|
| **Chinese name**  | WordPress User Verification 插件 user_verification_send_otp 页面认证绕过漏洞（CVE-2022-4693） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/user-verification"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3VzZXItdmVyaWZpY2F0aW9uIg%3D%3D) |
| **Number of assets affected**  | 707 |
| **Description**  | WordPress plugins User Verification is a plugin to protect your website from spam users and block instant access by using spam email addresses. There is an authorization problem vulnerability in WordPress plugins User Verification before version 1.0.94. The vulnerability stems from the fact that login verification can be bypassed. |
| **Impact** | There is an authorization problem vulnerability in WordPress plugins User Verification before version 1.0.94. The vulnerability stems from the fact that login verification can be bypassed. |

![](https://s3.bmp.ovh/imgs/2023/06/21/5e1c0207e077f467.gif)

## WordPress plugin Wholesale Market ced_cwsm_csv_import_export_module_download_error_log File Read Vulnerability (CVE-2022-4298)

|   **Vulnerability**  | **WordPress plugin Wholesale Market ced_cwsm_csv_import_export_module_download_error_log File Read Vulnerability (CVE-2022-4298)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Wholesale Market 插件 ced_cwsm_csv_import_export_module_download_error_log 任意文件读取漏洞（CVE-2022-4298） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/wholesale-market"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3dob2xlc2FsZS1tYXJrZXQi) |
| **Number of assets affected**  | 120 |
| **Description**  | The WordPress plugin Wholesale Market is a woocommerce extension plugin that enables your store to create wholesale users and set wholesale prices for products by. The WordPress plugin Wholesale Market version prior to 2.2.1 has a path traversal vulnerability, which is caused by not performing authorization checks and not validating user input. Attackers exploit this vulnerability to download arbitrary files from the server. |
| **Impact** | The WordPress plugin Wholesale Market version prior to 2.2.1 has a path traversal vulnerability, which is caused by not performing authorization checks and not validating user input. Attackers exploit this vulnerability to download arbitrary files from the server. |

![](https://s3.bmp.ovh/imgs/2023/06/21/f9ab68386ba634ba.gif)

## yongyou GRP-U8 U8AppProxy Arbitrary file upload vulnerability

|   **Vulnerability**  | **yongyou GRP-U8 U8AppProxy Arbitrary file upload vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 用友GRP-U8 软件 U8AppProxy 任意文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="window.location.replace(\"login.jsp?up=1\")" \|\| body="GRP-U8"](https://en.fofa.info/result?qbase64=Ym9keT0id2luZG93LmxvY2F0aW9uLnJlcGxhY2UoXCJsb2dpbi5qc3A%2FdXA9MVwiKSIgfHwgYm9keT0iR1JQLVU4Ig%3D%3D) |
| **Number of assets affected**  | 1308 |
| **Description**  | Yonyou GRP-U8 management software is a new generation of products launched by UFIDA focusing on national e-government affairs and based on cloud computing technology. It is the most professional government financial management software in the field of administrative affairs and finance in my country. UFIDA GRP-U8 management software U8AppProxy has an arbitrary file upload vulnerability, an attacker can upload a webshell to obtain server permissions.|
| **Impact** | UFIDA GRP-U8 management software U8AppProxy has an arbitrary file upload vulnerability, an attacker can upload a webshell to obtain server permissions. |
  
![](https://s3.bmp.ovh/imgs/2023/06/08/5cccc970d4c3d964.gif)

## WordPress Plugin Events Made Easy SQL Injection Vulnerability(CVE-2022-10271)

|   **Vulnerability**  | **WordPress Plugin Events Made Easy SQL Injection Vulnerability(CVE-2022-10271)**  |
| :----:   | :-----|
|  **Chinese name**  | WebLogic CoordinatorPortType 远程代码执行漏洞（CVE-2017-10271） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="Welcome to WebLogic Server") \|\| (title=="Error 404--Not Found") \|\| (((body="\<h1>BEA WebLogic Server" \|\| server="Weblogic" \|\| body="content=\"WebLogic Server" \|\| body="\<h1>Welcome to Weblogic Application" \|\| body="\<h1>BEA WebLogic Server") && header!="couchdb" && header!="boa" && header!="RouterOS" && header!="X-Generator: Drupal") \|\| (banner="Weblogic" && banner!="couchdb" && banner!="drupal" && banner!=" Apache,Tomcat,Jboss" && banner!="ReeCam IP Camera" && banner!="\<h2>Blog Comments\</h2>")) \|\| (port="7001" && protocol=="weblogic")](https://en.fofa.info/result?qbase64=KGJvZHk9IldlbGNvbWUgdG8gV2ViTG9naWMgU2VydmVyIikgfHwgKHRpdGxlPT0iRXJyb3IgNDA0LS1Ob3QgRm91bmQiKSB8fCAoKChib2R5PSI8aDE%2BQkVBIFdlYkxvZ2ljIFNlcnZlciIgfHwgc2VydmVyPSJXZWJsb2dpYyIgfHwgYm9keT0iY29udGVudD1cIldlYkxvZ2ljIFNlcnZlciIgfHwgYm9keT0iPGgxPldlbGNvbWUgdG8gV2VibG9naWMgQXBwbGljYXRpb24iIHx8IGJvZHk9IjxoMT5CRUEgV2ViTG9naWMgU2VydmVyIikgJiYgaGVhZGVyIT0iY291Y2hkYiIgJiYgaGVhZGVyIT0iYm9hIiAmJiBoZWFkZXIhPSJSb3V0ZXJPUyIgJiYgaGVhZGVyIT0iWC1HZW5lcmF0b3I6IERydXBhbCIpIHx8IChiYW5uZXI9IldlYmxvZ2ljIiAmJiBiYW5uZXIhPSJjb3VjaGRiIiAmJiBiYW5uZXIhPSJkcnVwYWwiICYmIGJhbm5lciE9IiBBcGFjaGUsVG9tY2F0LEpib3NzIiAmJiBiYW5uZXIhPSJSZWVDYW0gSVAgQ2FtZXJhIiAmJiBiYW5uZXIhPSI8aDI%2BQmxvZyBDb21tZW50czwvaDI%2BIikpIHx8IChwb3J0PSI3MDAxIiAmJiBwcm90b2NvbD09IndlYmxvZ2ljIik%3D) |
| **Number of assets affected**  | 127705 |
| **Description**  | WebLogic Server is one of the application server components suitable for both cloud and traditional environments. Due to the default activation of the WLS WebService component during the deployment process, WebLogic utilizes XMLDecoder to parse serialized data. Attackers can exploit this by constructing malicious XML files to achieve remote command execution, potentially allowing them to execute arbitrary code on the server and gain control over the entire web server. |
| **Impact** | Since WebLogic enables the WLS WebService component by default during the deployment process, this component uses XMLDecoder to parse the serialized data. An attacker can implement remote command execution by constructing a malicious XML file, which may cause the attacker to execute arbitrary code on the server side. And then control the entire web server. |
  
## WordPress Plugin Events Made Easy SQL Injection Vulnerability(CVE-2022-1905)

|   **Vulnerability**  | **WordPress Plugin Events Made Easy SQL Injection Vulnerability(CVE-2022-1905)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Events Made Easy 插件 admin-ajax.php 文件 lang 参数SQL注入漏洞（CVE-2022-1905） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/events-made-easy"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2V2ZW50cy1tYWRlLWVhc3ki) |
| **Number of assets affected**  | 4021 |
| **Description**  | Events Made Easy is a full-featured event and membership management solution for WordPress. Events Made Easy 2.2.81 has an unauthorized SQL injection vulnerability. |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

## Bifrost X-Requested-With Authentication Bypass Vulnerability (CVE-2022-39267)

|   **Vulnerability**  | **Bifrost X-Requested-With Authentication Bypass Vulnerability (CVE-2022-39267)**  |
| :----:   | :-----|
|  **Chinese name**  | Bifrost 中间件 X-Requested-With 系统身份认证绕过漏洞（CVE-2022-39267） |
| **CVSS core**  | 8.8 |
| **FOFA Query**  (click to view the results directly)| [body="/dologin" && body="Bifrost"](https://en.fofa.info/result?qbase64=Ym9keT0iL2RvbG9naW4iICYmIGJvZHk9IkJpZnJvc3Qi) |
| **Number of assets affected**  | 14 |
| **Description**  | Bifrost is a heterogeneous middleware that synchronizes MySQL, MariaDB to Redis, MongoDB, ClickHouse, MySQL and other services for production environments. Versions prior to 1.8.8-release are subject to authentication bypass in the admin and monitor user groups by deleting the X-Requested-With: XMLHttpRequest field in the request header. This issue has been patched in 1.8.8-release. There are no known workarounds. |
| **Impact** | Bifrost is a heterogeneous middleware that synchronizes MySQL, MariaDB and Kafka to Redis, MongoDB, ClickHouse and other services for production environments. It can bypass identity authentication by deleting request headers and obtain passwords for various database accounts configured in the environment. |

![](https://s3.bmp.ovh/imgs/2023/06/09/5d975955f9fd76d9.gif)

## MDT KNX manager panel default credentials vulnerability

|   **Vulnerability**  | **MDT KNX manager panel default credentials vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | MDT KNX 管理面板默认口令 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title="MDT Technologies GmbH" && server="DEFAULT IP PLATFORM"](https://en.fofa.info/result?qbase64=dGl0bGU9Ik1EVCBUZWNobm9sb2dpZXMgR21iSCIgJiYgc2VydmVyPSJERUZBVUxUIElQIFBMQVRGT1JNIg%3D%3D) |
| **Number of assets affected**  | 1135 |
| **Description**  | MDT Technologies is an intelligent building automation service provider based on KNX technology for product manufacturing. Its KNX-IP Interface/ Knx-ip Object Server panel is used to access every bus device in the KNX bus system. These panels have default passwords and malicious attackers can take over the target panel system. Default passwords exist on the KNX-IP Interface and KNX-IP Object Server management panel of MDT Technologies. Malicious attackers can use these passwords to take over the target web system. |
| **Impact** | Default passwords exist on the KNX-IP Interface and KNX-IP Object Server management panel of MDT Technologies. Malicious attackers can use these passwords to take over the target web system. |

![](https://s3.bmp.ovh/imgs/2023/06/09/c82a201acb398b4b.gif)

## WordPress Plugin LearnPress archive-course File Inclusion Vulnerability (CVE-2022-47615)

|   **Vulnerability**  | **WordPress Plugin LearnPress archive-course File Inclusion Vulnerability (CVE-2022-47615)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Plugin LearnPress archive-course 文件包含漏洞（CVE-2022-47615） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/learnpress"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2xlYXJucHJlc3Mi) |
| **Number of assets affected**  | 48623 |
| **Description**  | LearnPress is a comprehensive WordPress LMS Plugin for WordPress. This is one of the best WordPress LMS Plugins which can be used to easily create &amp; sell courses online.</p><p>WordPress LearnPress Plugin &lt;= 4.1.7.3.2 is vulnerable to Local File Inclusion. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

## maxView Storage Manager dynamiccontent.properties.xhtml RCE

|   **Vulnerability**  | **maxView Storage Manager dynamiccontent.properties.xhtml RCE**  |
| :----:   | :-----|
|  **Chinese name**  | maxView Storage Manager 系统 dynamiccontent.properties.xhtml 远程代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title=="maxView Storage Manager - Login"](https://en.fofa.info/result?qbase64=dGl0bGU9PSJtYXhWaWV3IFN0b3JhZ2UgTWFuYWdlciAtIExvZ2luIg%3D%3D) |
| **Number of assets affected**  | 1465 |
| **Description**  | maxView Storage Manager is a management system for enterprise storage and communication solutions.<br></p><p>There is a code execution vulnerability in maxView Storage Manager, an attacker can execute arbitrary code through dynamiccontent.properties.xhtml to gain server privileges. |
| **Impact** | There is a code execution vulnerability in maxView Storage Manager, an attacker can execute arbitrary code through dynamiccontent.properties.xhtml to gain server privileges. |
  
![](https://s3.bmp.ovh/imgs/2023/06/09/b3b919992e1c09f1.gif)

## SXETUX dynamiccontent.properties.xhtml RCE

|   **Vulnerability**  | **XETUX dynamiccontent.properties.xhtml RCE**  |
| :----:   | :-----|
|  **Chinese name**  | XETUX 软件 dynamiccontent.properties.xhtml 远程代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title="@XETUX" && title="XPOS" && body = "BackEnd"](https://en.fofa.info/result?qbase64=dGl0bGU9IkBYRVRVWCIgJiYgdGl0bGU9IlhQT1MiICYmIGJvZHkgPSAiQmFja0VuZCI%3D) |
| **Number of assets affected**  | 2002 |
| **Description**  | XETUX is a comprehensive solution comprising a set of safe, powerful and monitorable software programs, designed and developed for automatic control of restaurants and retail. There is a code execution vulnerability in XETUX, an attacker can execute arbitrary code through dynamiccontent.properties.xhtml to gain server privileges. |
| **Impact** | There is a code execution vulnerability in XETUX, an attacker can execute arbitrary code through dynamiccontent.properties.xhtml to gain server privileges. |


## Some Hikvision iVMS file upload vulnerabilities

|   **Vulnerability**  | **Some Hikvision iVMS file upload vulnerabilities**  |
| :----:   | :-----|
|  **Chinese name**  | 海康威视部分iVMS系统存在文件上传漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(body="class=\"enname\">iVMS-4200" && body="laRemPassword") \|\| (body="home/locationIndex.action?time=" && body="result.data.indexUrl;") \|\| (body="//caoshiyan modify 2015-06-30 中转页面" && body="/home/locationIndex.action?time=" \|\| body="home/licenseUpload.action") \|\| (body="class=\"out\">\<a href=\"download/iVMS-") \|\| ((body="tab-border code-iivms\">" \|\| body="login?service=" \|\| body="/eop/common/css/reset.css" \|\| header="/cms/web/gateway/"\|\| body="/cms/web/gateway/" \|\| header="/login?service=" \|\| title="iVMS") && header="Server: If you want know, you can ask me" && header!="404 Not Found") \|\| (body="var uuid = \"2b73083e-9b29-4005-a123-1d4ec47a36d5\"; // 用于检测VMS是否超时, chenliangyf1") \|\| (body="/cas/login" && body="js/login/login.service.js") \|\| (body="daysOflicenseDatedWarn" && body="/cas/login") \|\| (body="/ivms-ui/default/css/login.css") \|\| (server="Apache-Coyote/1.1" && body="/baseui/js/plugins/ui/jquery.placeholder.js") \|\| (body="/cas/static/js/jquery.placeholder.js") \|\| (body="IVMS.files/logo.gif") \|\| (body="license!getExpireDateOfDays.action" && body=" window.document.location = '/license!getExpireDateOfDays.action';") \|\| (body="iVMS-A100" && title="登录") \|\| (body="/error/browser.do" && body="/portal" && body="settings.skinStyle" && (body="src=\"/portal/common/js/commonVar.js" \|\| body="nginxService/v1/download/InstallRootCert.exe"))](https://en.fofa.info/result?qbase64=KGJvZHk9ImNsYXNzPVwiZW5uYW1lXCI%2BaVZNUy00MjAwIiAmJiBib2R5PSJsYVJlbVBhc3N3b3JkIikgfHwgKGJvZHk9ImhvbWUvbG9jYXRpb25JbmRleC5hY3Rpb24%2FdGltZT0iICYmIGJvZHk9InJlc3VsdC5kYXRhLmluZGV4VXJsOyIpIHx8IChib2R5PSIvL2Nhb3NoaXlhbiBtb2RpZnkgMjAxNS0wNi0zMCDkuK3ovazpobXpnaIiICYmIGJvZHk9Ii9ob21lL2xvY2F0aW9uSW5kZXguYWN0aW9uP3RpbWU9IiB8fCBib2R5PSJob21lL2xpY2Vuc2VVcGxvYWQuYWN0aW9uIikgfHwgKGJvZHk9ImNsYXNzPVwib3V0XCI%2BPGEgaHJlZj1cImRvd25sb2FkL2lWTVMtIikgfHwgKChib2R5PSJ0YWItYm9yZGVyIGNvZGUtaWl2bXNcIj4iIHx8IGJvZHk9ImxvZ2luP3NlcnZpY2U9IiB8fCBib2R5PSIvZW9wL2NvbW1vbi9jc3MvcmVzZXQuY3NzIiB8fCBoZWFkZXI9Ii9jbXMvd2ViL2dhdGV3YXkvInx8IGJvZHk9Ii9jbXMvd2ViL2dhdGV3YXkvIiB8fCBoZWFkZXI9Ii9sb2dpbj9zZXJ2aWNlPSIgfHwgdGl0bGU9ImlWTVMiKSAmJiBoZWFkZXI9IlNlcnZlcjogSWYgeW91IHdhbnQga25vdywgeW91IGNhbiBhc2sgbWUiICYmIGhlYWRlciE9IjQwNCBOb3QgRm91bmQiKSB8fCAoYm9keT0idmFyIHV1aWQgPSBcIjJiNzMwODNlLTliMjktNDAwNS1hMTIzLTFkNGVjNDdhMzZkNVwiOyAvLyDnlKjkuo7mo4DmtYtWTVPmmK%2FlkKbotoXml7YsIGNoZW5saWFuZ3lmMSIpIHx8IChib2R5PSIvY2FzL2xvZ2luIiAmJiBib2R5PSJqcy9sb2dpbi9sb2dpbi5zZXJ2aWNlLmpzIikgfHwgKGJvZHk9ImRheXNPZmxpY2Vuc2VEYXRlZFdhcm4iICYmIGJvZHk9Ii9jYXMvbG9naW4iKSB8fCAoYm9keT0iL2l2bXMtdWkvZGVmYXVsdC9jc3MvbG9naW4uY3NzIikgfHwgKHNlcnZlcj0iQXBhY2hlLUNveW90ZS8xLjEiICYmIGJvZHk9Ii9iYXNldWkvanMvcGx1Z2lucy91aS9qcXVlcnkucGxhY2Vob2xkZXIuanMiKSB8fCAoYm9keT0iL2Nhcy9zdGF0aWMvanMvanF1ZXJ5LnBsYWNlaG9sZGVyLmpzIikgfHwgKGJvZHk9IklWTVMuZmlsZXMvbG9nby5naWYiKSB8fCAoYm9keT0ibGljZW5zZSFnZXRFeHBpcmVEYXRlT2ZEYXlzLmFjdGlvbiIgJiYgYm9keT0iIHdpbmRvdy5kb2N1bWVudC5sb2NhdGlvbiA9ICcvbGljZW5zZSFnZXRFeHBpcmVEYXRlT2ZEYXlzLmFjdGlvbic7IikgfHwgKGJvZHk9ImlWTVMtQTEwMCIgJiYgdGl0bGU9IueZu%2BW9lSIpIHx8IChib2R5PSIvZXJyb3IvYnJvd3Nlci5kbyIgJiYgYm9keT0iL3BvcnRhbCIgJiYgYm9keT0ic2V0dGluZ3Muc2tpblN0eWxlIiAmJiAoYm9keT0ic3JjPVwiL3BvcnRhbC9jb21tb24vanMvY29tbW9uVmFyLmpzIiB8fCBib2R5PSJuZ2lueFNlcnZpY2UvdjEvZG93bmxvYWQvSW5zdGFsbFJvb3RDZXJ0LmV4ZSIpKQ%3D%3D) |
| **Number of assets affected**  | 15294 |
| **Description**  | Hikvision-iVMS comprehensive security management platform is an \"integrated\", \"digital\" and \"intelligent\" platform, including video, alarm, access control, visitor, elevator control, inspection, attendance, consumption, parking lot, Video intercom and other subsystems. The attacker constructs a token arbitrarily by obtaining the key, and requests an interface to upload files arbitrarily, resulting in obtaining the webshell permission of the server and executing malicious code remotely. |
| **Impact** | Hikvision-iVMS comprehensive security management platform is an \"integrated\", \"digital\" and \"intelligent\" platform, including video, alarm, access control, visitor, elevator control, inspection, attendance, consumption, parking lot, Video intercom and other subsystems. The attacker constructs a token arbitrarily by obtaining the key, and requests an interface to upload files arbitrarily, resulting in obtaining the webshell permission of the server and executing malicious code remotely. |

![](https://s3.bmp.ovh/imgs/2023/06/02/48cbd695f8499d33.gif)

## WordPress Plugin IWS SQL Injection Vulnerability (CVE-2022-4117)

|   **Vulnerability**  | **WordPress Plugin IWS SQL Injection Vulnerability (CVE-2022-4117)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Plugin IWS SQL注入漏洞（CVE-2022-4117） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/iws-geo-form-fields"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2l3cy1nZW8tZm9ybS1maWVsZHMi) |
| **Number of assets affected**  | 2186 |
| **Description**  | iws-geo-form-fields is a easy to use WordPress plugin, It uses Ajax to dynamically populate Select fields in your form,It can add Country - State - City select field in your WordPress website. Has an unauthorized SQL injection vulnerability. |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

![](https://s3.bmp.ovh/imgs/2023/06/01/a686613b77d15f5c.gif)

## Netgod SecGate 3600 Firewall File Upload Vulnerability

|   **Vulnerability**  | **Netgod SecGate 3600 Firewall File Upload Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 网神SecGate 3600防火墙 文件上传漏洞 |
| **CVSS core**  | 10.0 |
| **FOFA Query**  (click to view the results directly)| [title="网神SecGate 3600防火墙"](https://en.fofa.info/result?qbase64=dGl0bGU9Iue9keelnlNlY0dhdGUgMzYwMOmYsueBq%2BWimSI%3D) |
| **Number of assets affected**  | 747 |
| **Description**  | Netgod SecGate 3600 firewall is a composite hardware firewall based on status detection packet filtering and application level agents. It is a new generation of professional firewall equipment specially developed for large and medium-sized enterprises, governments, military, universities and other users. It supports external attack prevention, internal network security, network access control, network traffic monitoring and bandwidth management, dynamic routing, web content filtering, email content filtering, IP conflict detection and other functions, It can effectively ensure the security of the network; The product provides flexible network routing/bridging capabilities, supports policy routing and multi outlet link aggregation; It provides a variety of intelligent analysis and management methods, supports email alarm, supports log audit, provides comprehensive network management monitoring, and assists network administrators in completing network security management. There is a file upload vulnerability in SecGate 3600 firewall, which allows attackers to gain server control permission |
| **Impact** | There is a file upload vulnerability in SecGate 3600 firewall, which allows attackers to gain server control permissions. |

![](https://s3.bmp.ovh/imgs/2023/06/01/25197f6a68da33df.gif)

## Hangzhou new Zhongda NetcallServer management console default password

|   **Vulnerability**  | **Hangzhou new Zhongda NetcallServer management console default password**  |
| :----:   | :-----|
|  **Chinese name**  | 杭州新中大 NetcallServer 管理控制台默认口令 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=="netcallServer 管理控制台"](https://en.fofa.info/result?qbase64=dGl0bGU9PSJuZXRjYWxsU2VydmVyIOeuoeeQhuaOp%2BWItuWPsCI%3D) |
| **Number of assets affected**  | 567 |
| **Description**  | Hangzhou New Zhongda NetcallServer Management console is an instant messaging software of Hangzhou New Zhongda Technology Co., LTD. There is a default password in the NetcallServer management console of Hangzhou New CUHK, which can be exploited by attackers to obtain sensitive information. |
| **Impact** | The attacker can control the whole platform through the default password vulnerability and operate the core functions with the administrator rights. Cause sensitive information to leak. |

![](https://s3.bmp.ovh/imgs/2023/06/01/334aa827a0cbf890.gif)

## D-Link DCS-960L HNAP LoginPassword Authentication Bypass Vulnerability

|   **Vulnerability**  | **D-Link DCS-960L HNAP LoginPassword Authentication Bypass Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | D-Link DCS-960L HNAP LoginPassword 认证绕过漏洞 |
| **CVSS core**  | 8.8 |
| **FOFA Query**  (click to view the results directly)| [header="DCS-960L" \|\| banner="DCS-960L"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJEQ1MtOTYwTCIgfHwgYmFubmVyPSJEQ1MtOTYwTCI%3D) |
| **Number of assets affected**  | 16014 |
| **Description**  | D-Link DCS-960L is a network camera product of China Taiwan D-Link Company.<br></p><p>When D-Link DCS-960L processes the HNAP login request, the processing logic of the parameter LoginPassword is wrong, and the attacker can construct a special login request to bypass the login verification. |
| **Impact** | When D-Link DCS-960L processes the HNAP login request, the processing logic of the parameter LoginPassword is wrong, and the attacker can construct a special login request to bypass the login verification. |
  
![](https://s3.bmp.ovh/imgs/2023/06/02/78c7b86ddb23a225.gif)
  
## Array Networks AG/vxAG RCE (CVE-2022-42897)

|   **Vulnerability**  | **Array Networks AG/vxAG RCE (CVE-2022-42897)**  |
| :----:   | :-----|
|  **Chinese name**  | Array Networks AG/vxAG 远程代码执行漏洞（CVE-2022-42897） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [banner="/prx/000/http" \|\| header="/prx/000/http" \|\| body="an_util.js"](https://en.fofa.info/result?qbase64=YmFubmVyPSIvcHJ4LzAwMC9odHRwIiB8fCBoZWFkZXI9Ii9wcngvMDAwL2h0dHAiIHx8IGJvZHk9ImFuX3V0aWwuanMi) |
| **Number of assets affected**  | 10117 |
| **Description**  | Array Networks AG/vxAG is an Array SSL-VPN gateway product of Array Networks in the United States.<br></p><p>Array Networks AG/vxAG with ArrayOS AG prior to 9.4.0.469 has a security vulnerability that allows an unauthenticated attacker to achieve command injection, resulting in privilege escalation and control over the system. |
| **Impact** | Array Networks AG/vxAG with ArrayOS AG prior to 9.4.0.469 has a security vulnerability that allows an unauthenticated attacker to achieve command injection, resulting in privilege escalation and control over the system. |
  
![](https://s3.bmp.ovh/imgs/2023/06/02/15f9b0a6d6522815.gif)

## ASUS RT-AX56U Sensitive Information Disclosure Vulnerability

|   **Vulnerability**  | **ASUS RT-AX56U Sensitive Information Disclosure Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | ASUS RT-AX56U 敏感信息泄漏漏洞 |
| **CVSS core**  | 5.0 |
| **FOFA Query**  (click to view the results directly)| [banner="ASUS RT-AX56U" \|\| (body="RT-AX56U" && title=="ASUS Login")](https://fofa.info/result?qbase64=YmFubmVyPSJBU1VTIFJULUFYNTZVIiB8fCAoYm9keT0iUlQtQVg1NlUiICYmIHRpdGxlPT0iQVNVUyBMb2dpbiIp) |
| **Number of assets affected**  | 291164 |
| **Description**  | The ASUS RT-AX56U is a WiFi6 dual band 1800M E-sports route that supports the WiFi6 (802.11ax) standard and 80MHz bandwidth to provide better network performance and efficiency. With Trend Micro ™ The supported AiProtection commercial level security protection function provides network security protection for all connected intelligent devices.</p><p>After the construction request is sent to the vulnerable device, the passwd or shadow file in the system can be read, causing the password information disclosure problem of the administrator user. |
| **Impact** | After the construction request is sent to the vulnerable device, the passwd or shadow file in the system can be read, causing the password information disclosure problem of the administrator user. |
  
![](https://s3.bmp.ovh/imgs/2023/06/02/ef698a4fd96bc55f.gif)
  
## kkFileView onlinePreview Arbitrary File Read

|   **Vulnerability**  | **kkFileView onlinePreview Arbitrary File Read**  |
| :----:   | :-----|
|  **Chinese name**  | kkFileView onlinePreview 任意文件读取漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="/onlinePreview?url"](https://en.fofa.info/result?qbase64=Ym9keT0iL29ubGluZVByZXZpZXc%2FdXJsIg%3D%3D) |
| **Number of assets affected**  | 2360 |
| **Description**  | Keking kkFileView is a Spring-Boot online preview project for creating file documents of Keking Technology Co., Ltd. in China. here is a security vulnerability in Keking kkFileview, which stems from reading arbitrary files through directory traversal vulnerabilities, which may lead to the leakage of sensitive files on related hosts. |
| **Impact** | There is a security vulnerability in Keking kkFileview, which stems from reading arbitrary files through directory traversal vulnerabilities, which may lead to the leakage of sensitive files on related hosts. |
  
## Ruijie NBR Router webgl.data information

|   **Vulnerability**  | **Ruijie NBR Router webgl.data information**  |
| :----:   | :-----|
|  **Chinese name**  | 锐捷网络 NBR路由器 webgl.data 信息泄露漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="Ruijie - NBR" \|\| (body="support.ruijie.com.cn" && body="\<p>系统负荷过高，导致网络拥塞，建议降低系统负荷或重启路由器") \|\| body="class=\"line resource\" id=\"nbr_1\"" \|\| title="锐捷网络 --NBR路由器--登录界面" \|\| title=="锐捷网络") && body!="Server: couchdb"](https://en.fofa.info/result?qbase64=KGJvZHk9IlJ1aWppZSAtIE5CUiIgfHwgKGJvZHk9InN1cHBvcnQucnVpamllLmNvbS5jbiIgJiYgYm9keT0iPHA%2B57O757uf6LSf6I236L%2BH6auY77yM5a%2B86Ie0572R57uc5oul5aGe77yM5bu66K6u6ZmN5L2O57O757uf6LSf6I235oiW6YeN5ZCv6Lev55Sx5ZmoIikgfHwgYm9keT0iY2xhc3M9XCJsaW5lIHJlc291cmNlXCIgaWQ9XCJuYnJfMVwiIiB8fCB0aXRsZT0i6ZSQ5o23572R57ucIC0tTkJS6Lev55Sx5ZmoLS3nmbvlvZXnlYzpnaIiIHx8IHRpdGxlPT0i6ZSQ5o23572R57ucIikgJiYgYm9keSE9IlNlcnZlcjogY291Y2hkYiI%3D) |
| **Number of assets affected**  | 204290 |
| **Description**  | Ruijie Network NBR700G router is a wireless routing equipment of Ruijie Network Co., LTD. The NBR700G router of Ruijie Network has information vulnerability, which can be used by attackers to obtain sensitive information. |
| **Impact** | Attackers can use this vulnerability to obtain Ruijie network NBR700G router account and password, resulting in sensitive information leakage.|

![](https://s3.bmp.ovh/imgs/2023/05/26/f90351322821464c.gif)

## Gemtek Modem Configuration Interface Default password vulnerability

|   **Vulnerability**  | **Gemtek Modem Configuration Interface Default password vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 中保無限Modem Configuration Interface 默认口令漏洞 |
| **CVSS core**  | 5.0 |
| **FOFA Query**  (click to view the results directly)| [(title="Modem configuration interface" && body="status_device_status.asp" && body!="Huawei") && header!="Couchdb" && header!="JoomlaWor"](https://en.fofa.info/result?qbase64=KHRpdGxlPSJNb2RlbSBjb25maWd1cmF0aW9uIGludGVyZmFjZSIgJiYgYm9keT0ic3RhdHVzX2RldmljZV9zdGF0dXMuYXNwIiAmJiBib2R5IT0iSHVhd2VpIikgJiYgaGVhZGVyIT0iQ291Y2hkYiIgJiYgaGVhZGVyIT0iSm9vbWxhV29yIg%3D%3D) |
| **Number of assets affected**  | 4521 |
| **Description**  | Modem Configuration Interface is an unlimited router management system of China Insurance Corporation. There is a default password in the system. An attacker can control the entire platform through the default password (sigmu/secom) and operate the core functions with administrator privileges. |
| **Impact** | attackers can control the entire platform through the default password(sigmu/secom) vulnerability, and use administrator privileges to operate core functions. |
  
![](https://s3.bmp.ovh/imgs/2023/05/26/fc6fd428ce123e79.gif)

## Weaver e-cology CheckServer.jsp file sql injection vulnerability

|   **Vulnerability**  | **Weaver e-cology CheckServer.jsp file sql injection vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 泛微 e-cology CheckServer.jsp 文件 SQL 注入漏洞 |
| **CVSS core**  | 7.8 |
| **FOFA Query**  (click to view the results directly)| [((body="szFeatures" && body="redirectUrl") \|\| (body="rndData" && body="isdx") \|\| (body="typeof poppedWindow" && body="client/jquery.client_wev8.js") \|\| body="/theme/ecology8/jquery/js/zDialog_wev8.js" \|\| body="ecology8/lang/weaver_lang_7_wev8.js" \|\| body="src=\"/js/jquery/jquery_wev8.js" \|\| (header="Server: WVS" && (title!="404 Not Found" && header!="404 Not Found"))) && header!="testBanCookie" && header!="Couchdb" && header!="JoomlaWor" && body!="\<title>28ZE\</title>"](https://en.fofa.info/result?qbase64=KChib2R5PSJzekZlYXR1cmVzIiAmJiBib2R5PSJyZWRpcmVjdFVybCIpIHx8IChib2R5PSJybmREYXRhIiAmJiBib2R5PSJpc2R4IikgfHwgKGJvZHk9InR5cGVvZiBwb3BwZWRXaW5kb3ciICYmIGJvZHk9ImNsaWVudC9qcXVlcnkuY2xpZW50X3dldjguanMiKSB8fCBib2R5PSIvdGhlbWUvZWNvbG9neTgvanF1ZXJ5L2pzL3pEaWFsb2dfd2V2OC5qcyIgfHwgYm9keT0iZWNvbG9neTgvbGFuZy93ZWF2ZXJfbGFuZ183X3dldjguanMiIHx8IGJvZHk9InNyYz1cIi9qcy9qcXVlcnkvanF1ZXJ5X3dldjguanMiIHx8IChoZWFkZXI9IlNlcnZlcjogV1ZTIiAmJiAodGl0bGUhPSI0MDQgTm90IEZvdW5kIiAmJiBoZWFkZXIhPSI0MDQgTm90IEZvdW5kIikpKSAmJiBoZWFkZXIhPSJ0ZXN0QmFuQ29va2llIiAmJiBoZWFkZXIhPSJDb3VjaGRiIiAmJiBoZWFkZXIhPSJKb29tbGFXb3IiICYmIGJvZHkhPSI8dGl0bGU%2BMjhaRTwvdGl0bGU%2BIg%3D%3D) |
| **Number of assets affected**  | 		105760 |
| **Description**  | Weaver e-cology OA is a high-quality OA office system built on the principles of simplicity, applicability, and efficiency. The software is equipped with over 20 functional modules for processes, portals, knowledge, personnel, and communication, and adopts an intelligent voice interaction office mode. It can perfectly meet the actual needs of enterprises and provide them with full digital management.Weaver e-cology OA has an SQL injection vulnerability, which allows attackers to not only obtain information from the database (such as administrator background passwords, user personal information of the site) through SQL injection vulnerabilities, but also write Trojan horses to the server under high privileges to further gain server system privileges. |
| **Impact** | Weavere-cology OA has an SQL injection vulnerability, which allows attackers to not only obtain information from the database (such as administrator background passwords, user personal information of the site) through SQL injection vulnerabilities, but also write Trojan horses to the server under high privileges to further gain server system privileges.|

## Command Execution in Multiple TP-LINK Routers (CVE-2020-9374)

|   **Vulnerability**  | **Command Execution in Multiple TP-LINK Routers (CVE-2020-9374)**  |
| :----:   | :-----|
|  **Chinese name**  | TP_LINK 多款路由器命令执行（CVE-2020-9374） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="tplinkwifi.net" && body="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="](https://en.fofa.info/result?qbase64=Ym9keT0idHBsaW5rd2lmaS5uZXQiICYmIGJvZHk9IkFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaYWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXowMTIzNDU2Nzg5Ky89Ig%3D%3D) |
| **Number of assets affected**  | 	378381 |
| **Description**  | Multiple models of TP-Link routers from TP-Link Technologies Co., Ltd., including TL-WR841N, TL-WR840N, Archer C20, TL-WR849N, Archer C55, Archer C50, TL-WA801ND, TL-WR841HP, TL-WR845N, Archer C20i, Archer C2, are vulnerable to a command execution flaw. Attackers can exploit this vulnerability to execute arbitrary code, inject backdoors, gain server privileges, and ultimately take control of the entire web server. |
| **Impact** | An attacker can exploit this vulnerability by sending shell metacharacters through the routing trace feature of the dashboard, allowing them to execute arbitrary commands, inject backdoors, gain server privileges, and ultimately take control of the entire web server.|

## WebLogic JtaTransactionManager Remote Code Execution Vulnerability (CVE-2020-2551)

|   **Vulnerability**  | **WebLogic JtaTransactionManager Remote Code Execution Vulnerability (CVE-2020-2551)**  |
| :----:   | :-----|
|  **Chinese name**  | Weblogic JtaTransactionManager 反序列化远程代码执行漏洞（CVE-2020-2551） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(body="Welcome to WebLogic Server")\|\|(title=="Error 404--Not Found") \|\| (((body="\<h1>BEA WebLogic Server" \|\| server="Weblogic" \|\| body="content=\"WebLogic Server" \|\| body="\<h1>Welcome to Weblogic Application" \|\| body="\<h1>BEA WebLogic Server") && header!="couchdb" && header!="boa" && header!="RouterOS" && header!="X-Generator: Drupal") \|\| (banner="Weblogic" && banner!="couchdb" && banner!="drupal" && banner!=" Apache,Tomcat,Jboss" && banner!="ReeCam IP Camera" && banner!="\<h2>Blog Comments\</h2>")) \|\| (port="7001" && protocol=="weblogic")](https://en.fofa.info/result?qbase64=KGJvZHk9IldlbGNvbWUgdG8gV2ViTG9naWMgU2VydmVyIil8fCh0aXRsZT09IkVycm9yIDQwNC0tTm90IEZvdW5kIikgfHwgKCgoYm9keT0iPGgxPkJFQSBXZWJMb2dpYyBTZXJ2ZXIiIHx8IHNlcnZlcj0iV2VibG9naWMiIHx8IGJvZHk9ImNvbnRlbnQ9XCJXZWJMb2dpYyBTZXJ2ZXIiIHx8IGJvZHk9IjxoMT5XZWxjb21lIHRvIFdlYmxvZ2ljIEFwcGxpY2F0aW9uIiB8fCBib2R5PSI8aDE%2BQkVBIFdlYkxvZ2ljIFNlcnZlciIpICYmIGhlYWRlciE9ImNvdWNoZGIiICYmIGhlYWRlciE9ImJvYSIgJiYgaGVhZGVyIT0iUm91dGVyT1MiICYmIGhlYWRlciE9IlgtR2VuZXJhdG9yOiBEcnVwYWwiKSB8fCAoYmFubmVyPSJXZWJsb2dpYyIgJiYgYmFubmVyIT0iY291Y2hkYiIgJiYgYmFubmVyIT0iZHJ1cGFsIiAmJiBiYW5uZXIhPSIgQXBhY2hlLFRvbWNhdCxKYm9zcyIgJiYgYmFubmVyIT0iUmVlQ2FtIElQIENhbWVyYSIgJiYgYmFubmVyIT0iPGgyPkJsb2cgQ29tbWVudHM8L2gyPiIpKSB8fCAocG9ydD0iNzAwMSIgJiYgcHJvdG9jb2w9PSJ3ZWJsb2dpYyIp) |
| **Number of assets affected**  | 	127541 |
| **Description**  | WebLogic Server is one of the application server components applicable to cloud and traditional environments. WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution. |
| **Impact** | WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution. |

## ActiveMQ Arbitrary File Write Vulnerability (CVE-2016-3088)

|   **Vulnerability**  | **ActiveMQ Arbitrary File Write Vulnerability (CVE-2016-3088)**  |
| :----:   | :-----|
|  **Chinese name**  | ActiveMQ 消息代理系统 fileserver 文件上传漏洞（CVE-2016-3088） |
| **CVSS core**  | 9.6 |
| **FOFA Query**  (click to view the results directly)| [((((title="Apache ActiveMQ" \|\| (port="8161" && header="Server: Jetty") \|\| header="realm=\"ActiveMQRealm") && header!="couchdb" && header!="drupal" && body!="Server: couchdb") \|\| (banner="server:ActiveMQ" \|\| banner="Magic:ActiveMQ" \|\| banner="realm=\"ActiveMQRealm") \|\| banner="Apache ActiveMQ") \|\| (((title="Apache ActiveMQ" \|\| (port="8161" && header="Server: Jetty") \|\| header="realm=\"ActiveMQRealm") && header!="couchdb" && header!="drupal" && body!="Server: couchdb") \|\| (banner="server:ActiveMQ" \|\| banner="Magic:ActiveMQ" \|\| banner="realm=\"ActiveMQRealm") \|\| banner="Apache ActiveMQ")) && protocol!="activemq" && protocol!="stomp"](https://en.fofa.info/result?qbase64=KCgoKHRpdGxlPSJBcGFjaGUgQWN0aXZlTVEiIHx8IChwb3J0PSI4MTYxIiAmJiBoZWFkZXI9IlNlcnZlcjogSmV0dHkiKSB8fCBoZWFkZXI9InJlYWxtPVwiQWN0aXZlTVFSZWFsbSIpICYmIGhlYWRlciE9ImNvdWNoZGIiICYmIGhlYWRlciE9ImRydXBhbCIgJiYgYm9keSE9IlNlcnZlcjogY291Y2hkYiIpIHx8IChiYW5uZXI9InNlcnZlcjpBY3RpdmVNUSIgfHwgYmFubmVyPSJNYWdpYzpBY3RpdmVNUSIgfHwgYmFubmVyPSJyZWFsbT1cIkFjdGl2ZU1RUmVhbG0iKSB8fCBiYW5uZXI9IkFwYWNoZSBBY3RpdmVNUSIpIHx8ICgoKHRpdGxlPSJBcGFjaGUgQWN0aXZlTVEiIHx8IChwb3J0PSI4MTYxIiAmJiBoZWFkZXI9IlNlcnZlcjogSmV0dHkiKSB8fCBoZWFkZXI9InJlYWxtPVwiQWN0aXZlTVFSZWFsbSIpICYmIGhlYWRlciE9ImNvdWNoZGIiICYmIGhlYWRlciE9ImRydXBhbCIgJiYgYm9keSE9IlNlcnZlcjogY291Y2hkYiIpIHx8IChiYW5uZXI9InNlcnZlcjpBY3RpdmVNUSIgfHwgYmFubmVyPSJNYWdpYzpBY3RpdmVNUSIgfHwgYmFubmVyPSJyZWFsbT1cIkFjdGl2ZU1RUmVhbG0iKSB8fCBiYW5uZXI9IkFwYWNoZSBBY3RpdmVNUSIpKSAmJiBwcm90b2NvbCE9ImFjdGl2ZW1xIiAmJiBwcm90b2NvbCE9InN0b21wIg%3D%3D) |
| **Number of assets affected**  | 42641 |
| **Description**  | Apache ActiveMQ is the most popular open source, multi-protocol, Java-based message broker.</p><p>The Fileserver web application in Apache ActiveMQ 5.x before 5.14.0 allows remote attackers to upload and execute arbitrary files via an HTTP PUT followed by an HTTP MOVE request. |
| **Impact** | The Fileserver web application in Apache ActiveMQ 5.x before 5.14.0 allows remote attackers to upload and execute arbitrary files via an HTTP PUT followed by an HTTP MOVE request. |

## Apache Superset Permission Bypass Vulnerability (CVE-2023-27524)

|   **Vulnerability**  | **Apache Superset Permission Bypass Vulnerability (CVE-2023-27524)**  |
| :----:   | :-----|
|  **Chinese name**  | Apache Superset 权限绕过漏洞（CVE-2023-27524） |
| **CVSS core**  | 8.9 |
| **FOFA Query**  (click to view the results directly)| [(title="Superset" && (body="appbuilder" \|\| body="\<img src=\"https://joinsuperset.com/img/supersetlogovector.svg")) \|\| body="\<a href=\"https://manage.app-sdx.preset.io\" class=\"button\">Back to workspaces\</a>\</section>" \|\| (body="/static/assets/dist/common.644ae7ae973b00abc14b.entry.js" \|\| (body="/static/assets/images/favicon.png" && body="/static/appbuilder/js/jquery-latest.js") && body="Superset") \|\| header="/superset/welcome/" \|\|  title="500: Internal server error | Superset" \|\| title="404: Not found | Superset" \|\| banner="/superset/welcome/" \|\| banner="/superset/dashboard/"](https://en.fofa.info/result?qbase64=KHRpdGxlPSJTdXBlcnNldCIgJiYgKGJvZHk9ImFwcGJ1aWxkZXIiIHx8IGJvZHk9IjxpbWcgc3JjPVwiaHR0cHM6Ly9qb2luc3VwZXJzZXQuY29tL2ltZy9zdXBlcnNldGxvZ292ZWN0b3Iuc3ZnIikpIHx8IGJvZHk9IjxhIGhyZWY9XCJodHRwczovL21hbmFnZS5hcHAtc2R4LnByZXNldC5pb1wiIGNsYXNzPVwiYnV0dG9uXCI%2BQmFjayB0byB3b3Jrc3BhY2VzPC9hPjwvc2VjdGlvbj4iIHx8IChib2R5PSIvc3RhdGljL2Fzc2V0cy9kaXN0L2NvbW1vbi42NDRhZTdhZTk3M2IwMGFiYzE0Yi5lbnRyeS5qcyIgfHwgKGJvZHk9Ii9zdGF0aWMvYXNzZXRzL2ltYWdlcy9mYXZpY29uLnBuZyIgJiYgYm9keT0iL3N0YXRpYy9hcHBidWlsZGVyL2pzL2pxdWVyeS1sYXRlc3QuanMiKSAmJiBib2R5PSJTdXBlcnNldCIpIHx8IGhlYWRlcj0iL3N1cGVyc2V0L3dlbGNvbWUvIiB8fCAgdGl0bGU9IjUwMDogSW50ZXJuYWwgc2VydmVyIGVycm9yIHwgU3VwZXJzZXQiIHx8IHRpdGxlPSI0MDQ6IE5vdCBmb3VuZCB8IFN1cGVyc2V0IiB8fCBiYW5uZXI9Ii9zdXBlcnNldC93ZWxjb21lLyIgfHwgYmFubmVyPSIvc3VwZXJzZXQvZGFzaGJvYXJkLyI%3D) |
| **Number of assets affected**  | 43325 |
| **Description**  | Apache Superset is a data visualization and data exploration platform of the Apache Foundation. Apache Superset versions 2.0.1 and earlier have security vulnerabilities. Attackers exploit this vulnerability to verify and access unauthorized resources. |
| **Impact** | Attackers can exploit this vulnerability to verify and access unauthorized resources. |

![](https://s3.bmp.ovh/imgs/2023/05/22/46c693629791a204.gif)

## Apache Archiva RepositoryServlet internal Arbitrary File Read (CVE-2022-40308)

|   **Vulnerability**  | **Apache Archiva RepositoryServlet internal Arbitrary File Read (CVE-2022-40308)**  |
| :----:   | :-----|
|  **Chinese name**  | Apache Archiva RepositoryServlet 代理功能 internal 文件任意文件读取漏洞（CVE-2022-40308） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title="Apache Archiva" \|\| body="/archiva.js" \|\| body="/archiva.css"](https://en.fofa.info/result?qbase64=dGl0bGU9IkFwYWNoZSBBcmNoaXZhIiB8fCBib2R5PSIvYXJjaGl2YS5qcyIgfHwgYm9keT0iL2FyY2hpdmEuY3NzIg%3D%3D) |
| **Number of assets affected**  | 910 |
| **Description**  | Apache Archiva is a set of software used by the Apache Foundation of the United States to manage one or more remote storages. The software provides features such as remote Repository agents, secure role-based access management, and usage reporting. Versions prior to Apache Archiva 2.2.9 have a security vulnerability, which stems from the ability to read database files directly without logging in if anonymous reading is enabled. |
| **Impact** | Versions prior to Apache Archiva 2.2.9 have a security vulnerability, which stems from the ability to read database files directly without logging in if anonymous reading is enabled. |

![](https://s3.bmp.ovh/imgs/2023/05/19/cd6e4b533cae79ff.gif)

## WordPress Plugin InPost Gallery popup_shortcode_attributes File Inclusion Vulnerability(CVE-2022-4063)

|   **Vulnerability**  | **WordPress Plugin InPost Gallery popup_shortcode_attributes File Inclusion Vulnerability(CVE-2022-4063)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress InPost Gallery 插件 popup_shortcode_attributes 参数文件包含漏洞（CVE-2022-4063） |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/inpost-gallery"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2lucG9zdC1nYWxsZXJ5Ig%3D%3D) |
| **Number of assets affected**  | 556 |
| **Description**  | InPost Gallery is a powerful and very pleasing photo gallery plugin for working with images in WordPress.There is a file inclusion vulnerability in InPost Gallery < 2.1.4.1. Attackers can exploit this vulnerability to obtain sensitive files. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

![](https://s3.bmp.ovh/imgs/2023/05/19/e7c2e191cff78bbd.gif)

## SOPHOS-Netgenie Default Password

|   **Vulnerability**  | **SOPHOS-Netgenie  Default Password**  |
| :----:   | :-----|
|  **Chinese name**  | SOPHOS-Netgenie 默认口令 |
| **CVSS core**  | 5.0 |
| **FOFA Query**  (click to view the results directly)| [header="Server: Netgenie" \|\| banner="Server: Netgenie"](https://en.fofa.info/result?qbase64=aGVhZGVyPSJTZXJ2ZXI6IE5ldGdlbmllIiB8fCBiYW5uZXI9IlNlcnZlcjogTmV0Z2VuaWUi) |
| **Number of assets affected**  | 1566 |
| **Description**  | With NetGenie, get support for all types of Internet connectivity, viz. VDSL2, ADSL2+, Cable Internet and 3G connection, along with excellent wireless range, high performance, Gigabit port and threat-free Wi-Fi over multiple devices. Get Internet activity reports of children at home along with security reports of your home network.The command center of this series of printers has admin/admin default password. |
| **Impact** | SOPHOS-Netgenie  have default passwords. Attackers can use the default password admin/admin to log in to the system background without authorization, perform other sensitive operations, and obtain more sensitive information. |

![](https://s3.bmp.ovh/imgs/2023/05/23/a7c7faa7524301b1.gif)

## WAVLINK WN535 G3 router live_ Check.shtml file information disclosure vulnerability (CVE-2022-31845)

|   **Vulnerability**  | **WAVLINK WN535 G3 router live_ Check.shtml file information disclosure vulnerability (CVE-2022-31845)**  |
| :----:   | :-----|
|  **Chinese name**  | WAVLINK WN535 G3 路由器 live_check.shtml 文件信息泄露漏洞（CVE-2022-31845） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="firstFlage"](https://en.fofa.info/result?qbase64=Ym9keT0iZmlyc3RGbGFnZSI%3D) |
| **Number of assets affected**  | 3001 |
| **Description**  | WAVLINK WN535 is a dual band 4G LTE intelligent router. There is a security vulnerability in WAVLINK WN535 G3 M35G3R.V5030.180927, which originates in live_ There is a vulnerability in check.shtml. Attackers can use this vulnerability to obtain sensitive router information by executing exec cmd functions. |
| **Impact** | There is a security vulnerability in WAVLINK WN535 G3 M35G3R.V5030.180927, which originates in live_ There is a vulnerability in check.shtml. Attackers can use this vulnerability to obtain sensitive router information by executing exec cmd functions. |

![](https://s3.bmp.ovh/imgs/2023/05/23/5ac982d1b5b0c5b4.gif)

## I3Geo codemirror.php file pagina parameter file read vulnerability (CVE-2022-32409)

|   **Vulnerability**  | **I3Geo codemirror.php file pagina parameter file read vulnerability (CVE-2022-32409)**  |
| :----:   | :-----|
|  **Chinese name**  | i3Geo codemirror.php 文件 pagina 参数文件读取漏洞（CVE-2022-32409） |
| **CVSS core**  | 7.6 |
| **FOFA Query**  (click to view the results directly)| [body="i3geo"](https://en.fofa.info/result?qbase64=Ym9keT0iaTNnZW8i) |
| **Number of assets affected**  | 88 |
| **Description**  | I3geo is an open source application of salade situacao for developing interactive network maps. I3Geo has a file reading vulnerability, through which an attacker can read important system files (such as database configuration files, system configuration files), database configuration files, etc., causing the website to be in an extremely insecure state. |
| **Impact** | I3Geo has a file reading vulnerability, through which an attacker can read important system files (such as database configuration files, system configuration files), database configuration files, etc., causing the website to be in an extremely insecure state. |

![](https://s3.bmp.ovh/imgs/2023/05/23/17685ef4dce000de.gif)

## Telos Alliance Omnia MPX Node downloadMainLog fnameFile Reading Vulnerability(CVE-2022-36642)

|   **Vulnerability**  | **Telos Alliance Omnia MPX Node downloadMainLog fnameFile Reading Vulnerability(CVE-2022-36642)**  |
| :----:   | :-----|
|  **Chinese name**  | Telos Alliance Omnia MPX Node 硬件编解码器 downloadMainLog 文件 fname 参数文件读取漏洞（CVE-2022-36642） |
| **CVSS core**  | 7.6 |
| **FOFA Query**  (click to view the results directly)| [body="Omnia MPX"](https://en.fofa.info/result?qbase64=Ym9keT0iT21uaWEgTVBYIg%3D%3D) |
| **Number of assets affected**  | 49 |
| **Description**  | Telos Alliance Omnia MPX Node is a special hardware codec of Telos Alliance of the United States. Ability to leverage Omnia μ The MPXTM algorithm sends or receives complete FM signals at data rates as low as 320 kbps, making it ideal for networks with limited capacity, including IP radios. There is a security vulnerability in Telos Alliance Omnia MPX Node 1.5.0+r1 and earlier versions, which originates from the local file disclosure vulnerability in/appConfig/userDB.json. An attacker uses this vulnerability to elevate privileges to root and execute arbitrary commands. |
| **Impact** | There is a security vulnerability in Telos Alliance Omnia MPX Node 1.5.0+r1 and earlier versions, which originates from the local file disclosure vulnerability in/appConfig/userDB.json. An attacker uses this vulnerability to elevate privileges to root and execute arbitrary commands. |

![](https://s3.bmp.ovh/imgs/2023/05/23/e024d90bde2b5088.gif)

## TamronOS IPTV backup file down vulnerability

|   **Vulnerability**  | **TamronOS IPTV backup file down vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | TamronOS IPTV 系统 backup 任意文件下载漏洞 |
| **CVSS core**  | 5.0 |
| **FOFA Query**  (click to view the results directly)| [title="TamronOS IPTV系统"](https://en.fofa.info/result?qbase64=dGl0bGU9IlRhbXJvbk9TIElQVFbns7vnu58i) |
| **Number of assets affected**  | 472 |
| **Description**  | TamronOS IPTV system is an intelligent TV management system. The system has an arbitrary file download vulnerability, through which an attacker can read system files and obtain sensitive information. |
| **Impact** | an attacker can read system files and obtain sensitive information. |

![](https://s3.bmp.ovh/imgs/2023/05/23/701178b7dafee00c.gif)

## Hikvision NCG Networking Gateway login.php Directory traversal Vulnerability

|   **Vulnerability**  | **Hikvision NCG Networking Gateway login.php Directory traversal Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 海康威视 NCG 联网网关 login.php 文件目录遍历漏洞 |
| **CVSS core**  | 7.8 |
| **FOFA Query**  (click to view the results directly)| [body="data/login.php"](https://en.fofa.info/result?qbase64=Ym9keT0iZGF0YS9sb2dpbi5waHAi) |
| **Number of assets affected**  | 735 |
| **Description**  | The Hikvision NCG Networking Gateway  of Hikvision is a carrier level network gateway device integrating signaling gateway service, media gateway service, security authentication, authority management, log management and network management functions. An attacker can read important system files (such as database configuration files, system configuration files), database configuration files, etc. through this vulnerability, causing the website to be in an extremely insecure state. |
| **Impact** | An attacker can read important system files (such as database configuration files, system configuration files), database configuration files, etc. through this vulnerability, causing the website to be in an extremely insecure state. |

![](https://s3.bmp.ovh/imgs/2023/05/23/a4a531adf15cb89f.gif)

## Wordpress wpjobboard plugin wpjobboard directory traversal vulnerability (CVE-2022-2544)

|   **Vulnerability**  | **Wordpress wpjobboard plugin wpjobboard directory traversal vulnerability (CVE-2022-2544)**  |
| :----:   | :-----|
|  **Chinese name**  | Wordpress wpjobboard 插件 wpjobboard 页面目录遍历漏洞（CVE-2022-2544） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body="wp-content/plugins/wpjobboard"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3dwam9iYm9hcmQi) |
| **Number of assets affected**  | 1201 |
| **Description**  | Wpjobboard is a plugin of Wordpress. The Wpjobboard plug-in allows website owners to embed payment forms and make payments via Visa, American Express, Discover and Mastercard through their Click&amp;Lead merchant accounts.The Wpjobboard plug-in has a directory traversal vulnerability, through which an attacker can view sensitive directories and files in the server, control the entire system, and finally cause the system to be in an extremely insecure state. |
| **Impact** | 	The Wpjobboard plug-in has a directory traversal vulnerability, through which an attacker can view sensitive directories and files in the server, control the entire system, and finally cause the system to be in an extremely insecure state. |

![](https://s3.bmp.ovh/imgs/2023/05/23/4ddb35a567453502.gif)

## Telrad-WLTMS-110 Default Password

|   **Vulnerability**  | **Telrad-WLTMS-110 Default Password**  |
| :----:   | :-----|
|  **Chinese name**  | Telrad-WLTMS-110 默认口令 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body="WLTMS-110 Telrad" && body="frameRtoLControl.js") \|\| body="var multipleParameters = \" WLTMS-110"](https://en.fofa.info/result?qbase64=KGJvZHk9IldMVE1TLTExMCBUZWxyYWQiICYmIGJvZHk9ImZyYW1lUnRvTENvbnRyb2wuanMiKSB8fCBib2R5PSJ2YXIgbXVsdGlwbGVQYXJhbWV0ZXJzID0gXCIgV0xUTVMtMTEwIg%3D%3D) |
| **Number of assets affected**  | 1201 |
| **Description**  | The Telrad-WLTMS-110 offers deployment flexibility. The high throughput and transmit power of the CPEs combine with the small tower footprint and high capacity of our flagship BreezeCOMPACT base stations - reducing the density of base stations in a network and enabling faster, more affordable LTE deployments.The command center of this series of printers has admin/admin default password. |
| **Impact** | 	Telrad-WLTMS-110&nbsp; have default passwords. Attackers can use the default password admin/admin to log in to the system background without authorization, perform other sensitive operations, and obtain more sensitive information. |

## Weaver e-cology ofsLogin.jsp User Login Bypass Vulnerability

|   **Vulnerability**  | **Weaver e-cology ofsLogin.jsp User Login Bypass Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | Weaver e-cology ofsLogin.jsp 用户登陆绕过漏洞 |
| **CVSS core**  | 9.3 |
| **FOFA Query**  (click to view the results directly)| [body="/wui/common/" \|\| body="/wui/index.html"](https://en.fofa.info/result?qbase64=Ym9keT0iL3d1aS9jb21tb24vInx8Ym9keT0iL3d1aS9pbmRleC5odG1sIg%3D%3D) |
| **Number of assets affected**  | 92980 |
| **Description**  | The Weaver management application platform (e-cology) is a comprehensive enterprise management platform. It has diversified functions, including enterprise information portal, knowledge document management, work process management, human resource management, customer relationship management, project management, financial management, asset management, supply chain management and data center. This platform helps enterprises integrate various resources, including management, marketing, sales, research and development, personnel, and administrative fields. Through e-cology, these resources can be integrated on a unified platform and provide users with a unified interface for easy operation and information retrieval&nbsp;. The Weaver management application platform (e-cology) has a privilege bypass vulnerability, which allows attackers to bypass system privileges and log in to the system to perform malicious operations |
| **Impact** | The Weaver management application platform (e-cology) has a privilege bypass vulnerability, which allows attackers to bypass system privileges and log in to the system to perform malicious operations. |

![](https://s3.bmp.ovh/imgs/2023/05/17/b2f501929dc9ba6c.gif)

## Weblogic Commons Collections serialization code execution vulnerability (CVE-2015-4852)

|   **Vulnerability**  | **Weblogic Commons Collections serialization code execution vulnerability (CVE-2015-4852)**  |
| :----:   | :-----|
|  **Chinese name**  | Weblogic Commons Collections 序列化代码执行漏洞（CVE-2015-4852） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [(body=\"Welcome to WebLogic Server\") \|\| (title==\"Error 404--Not Found\") \|\| (((body=\"\<h1\>BEA WebLogic Server\" \|\| server=\"Weblogic\" \|\| body=\"content=\\\"WebLogic Server\" \|\| body=\"\<h1\>Welcome to Weblogic Application\" \|\| body=\"\<h1\>BEA WebLogic Server\") && header!=\"couchdb\" && header!=\"boa\" && header!=\"RouterOS\" && header!=\"X-Generator: Drupal\") \|\| (banner=\"Weblogic\" && banner!=\"couchdb\" && banner!=\"drupal\" && banner!=\" Apache,Tomcat,Jboss\" && banner!=\"ReeCam IP Camera\" && banner!=\"\<h2\>Blog Comments\</h2\>\")) \|\| (port=\"7001\" && protocol==\"weblogic\")](https://en.fofa.info/result?qbase64=KGJvZHk9IldlbGNvbWUgdG8gV2ViTG9naWMgU2VydmVyIikgfHwgKHRpdGxlPT0iRXJyb3IgNDA0LS1Ob3QgRm91bmQiKSB8fCAoKChib2R5PSI8aDE%2BQkVBIFdlYkxvZ2ljIFNlcnZlciIgfHwgc2VydmVyPSJXZWJsb2dpYyIgfHwgYm9keT0iY29udGVudD1cIldlYkxvZ2ljIFNlcnZlciIgfHwgYm9keT0iPGgxPldlbGNvbWUgdG8gV2VibG9naWMgQXBwbGljYXRpb24iIHx8IGJvZHk9IjxoMT5CRUEgV2ViTG9naWMgU2VydmVyIikgJiYgaGVhZGVyIT0iY291Y2hkYiIgJiYgaGVhZGVyIT0iYm9hIiAmJiBoZWFkZXIhPSJSb3V0ZXJPUyIgJiYgaGVhZGVyIT0iWC1HZW5lcmF0b3I6IERydXBhbCIpIHx8IChiYW5uZXI9IldlYmxvZ2ljIiAmJiBiYW5uZXIhPSJjb3VjaGRiIiAmJiBiYW5uZXIhPSJkcnVwYWwiICYmIGJhbm5lciE9IiBBcGFjaGUsVG9tY2F0LEpib3NzIiAmJiBiYW5uZXIhPSJSZWVDYW0gSVAgQ2FtZXJhIiAmJiBiYW5uZXIhPSI8aDI%2BQmxvZyBDb21tZW50czwvaDI%2BIikpIHx8IChwb3J0PSI3MDAxIiAmJiBwcm90b2NvbD09IndlYmxvZ2ljIik%3D) |
| **Number of assets affected**  | 127703 |
| **Description**  | WebLogic Server is an application server component suitable for both cloud and traditional environments. The WebLogic Commons Collections component has a remote code execution vulnerability that allows unauthenticated attackers to access vulnerable WebLogic Servers through the IIOP protocol and compromise them. Successful exploitation of the vulnerability can lead to the attacker taking over the WebLogic Server, resulting in remote code execution. |
| **Impact** | The WebLogic Commons Collections component has a remote code execution vulnerability that allows unauthenticated attackers to access vulnerable WebLogic Servers through the IIOP protocol and compromise them. Successful exploitation of the vulnerability can lead to the attacker taking over the WebLogic Server, resulting in remote code execution. |

![](https://s3.bmp.ovh/imgs/2023/05/13/a5c75ddf8c43c69e.gif)

## PowerJob /job/list api unauthorized access vulnerability

|   **Vulnerability**  | **PowerJob /job/list api unauthorized access vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | PowerJob /job/list 接口未授权访问漏洞 |
| **CVSS core**  | 7.3 |
| **FOFA Query**  (click to view the results directly)| [(title=\"PowerJob\" && body=\"We're sorry but oms-console\") \|\| (banner=\"Content-Length: 1222\" \|\| banner=\"Content-Length: 1260\") && banner=\"Vary: Origin\" && banner=\"Vary: Access-Control-Request-Headers\" && banner!=\"X-Content-Type-Options: nosniff\"](https://en.fofa.info/result?qbase64=KHRpdGxlPSJQb3dlckpvYiIgJiYgYm9keT0iV2UncmUgc29ycnkgYnV0IG9tcy1jb25zb2xlIikgfHwgKGJhbm5lcj0iQ29udGVudC1MZW5ndGg6IDEyMjIiIHx8IGJhbm5lcj0iQ29udGVudC1MZW5ndGg6IDEyNjAiKSAmJiBiYW5uZXI9IlZhcnk6IE9yaWdpbiIgJiYgYmFubmVyPSJWYXJ5OiBBY2Nlc3MtQ29udHJvbC1SZXF1ZXN0LUhlYWRlcnMiICYmIGJhbm5lciE9IlgtQ29udGVudC1UeXBlLU9wdGlvbnM6IG5vc25pZmYi) |
| **Number of assets affected**  | 656 |
| **Description**  | PowerJob (formerly OhMyScheduler) is a new generation of distributed scheduling and computing framework that allows you to easily complete job scheduling and distributed computing of complex tasks. Attackers can control the entire system through unauthorized access vulnerabilities, and ultimately lead to an extremely insecure state of the system. |
| **Impact** | Attackers can exploit an unauthorized access vulnerability in /job/list to obtain task information for the entire system, which could ultimately result in the system being in an extremely insecure state. |

![](https://s3.bmp.ovh/imgs/2023/05/23/0eba0b14510d1610.gif)

## nginxWebUI runCmd file remote command execution vulnerability

|   **Vulnerability**  | **nginxWebUI runCmd file remote command execution vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | nginxWebUI runCmd 文件命令执行漏洞 |
| **CVSS core**  | 9.2 |
| **FOFA Query**  (click to view the results directly)| [title=\"nginxWebUI\" && body=\"refreshCode('codeImg')\"](https://en.fofa.info/result?qbase64=dGl0bGU9Im5naW54V2ViVUkiICYmIGJvZHk9InJlZnJlc2hDb2RlKCdjb2RlSW1nJyki) |
| **Number of assets affected**  | 5856 |
| **Description**  | NginxWebUI is a tool for graphical management of nginx configuration. You can use web pages to quickly configure various functions of nginx, including http protocol forwarding, tcp protocol forwarding, reverse proxy, load balancing, static html server, automatic application, renewal and configuration of ssl certificates. After configuration, you can create nginx. conf file, and control nginx to use this file to start and reload, completing the graphical control loop of nginx. Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | Attackers can use this vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |

![](https://s3.bmp.ovh/imgs/2023/05/23/a77433ad12687464.gif)

## Yun-Box authService fastjson serialization code execution vulnerability

|   **Vulnerability**  | **Yun-Box authService fastjson serialization code execution vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 云匣子 authService fastjson 序列化代码执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [(body=\"id=mTokenPlugin width=0 height=0 style=\\\"position: absolute;LEFT: 0px; TOP: 0px\\\"\" && body=\"type=application/x-xtx-axhost\") && (cert=\"Domain Control Validated\" \|\| cert=\"云匣子\")](https://en.fofa.info/result?qbase64=KGJvZHk9ImlkPW1Ub2tlblBsdWdpbiB3aWR0aD0wIGhlaWdodD0wIHN0eWxlPVwicG9zaXRpb246IGFic29sdXRlO0xFRlQ6IDBweDsgVE9QOiAwcHhcIiIgJiYgYm9keT0idHlwZT1hcHBsaWNhdGlvbi94LXh0eC1heGhvc3QiKSAmJiAoY2VydD0iRG9tYWluIENvbnRyb2wgVmFsaWRhdGVkIiB8fCBjZXJ0PSLkupHljKPlrZAiKQ%3D%3D) |
| **Number of assets affected**  | 620 |
| **Description**  | Yun-Box is a secure management tool developed by Yunanbao for tenants to connect to cloud resources, which can help cloud tenants manage virtual machines, databases, and other resources on the cloud in a more secure and precise manner. With years of experience in operations and security, Yun-Box combines operations and security on the cloud to achieve pre-planned operations, in-process control, and post-audit. Additionally, Yun-Box integrates features such as automated operations, asset topology discovery, and account security to provide comprehensive and reliable cloud security management services. |
| **Impact** | Yun-Box uses the vulnerable fastjson component, and hackers can launch attacks on Yun-Box by exploiting the fastjson serialization vulnerability to gain server privileges. |

![](https://s3.bmp.ovh/imgs/2023/05/23/7853202174123e25.gif)

## Realor Tianyi AVS ConsoleExternalApi.XGI file SQL Injection vulnerability

|   **Vulnerability**  | **Realor Tianyi AVS ConsoleExternalApi.XGI file SQL Injection vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 瑞友天翼应用虚拟化系统 ConsoleExternalApi.XGI 文件 iDisplayStart 参数 SQL 注入漏洞 |
| **CVSS core**  | 9.2 |
| **FOFA Query**  (click to view the results directly)| [title=\"瑞友天翼－应用虚拟化系统\" \|\| title=\"瑞友应用虚拟化系统\"](https://en.fofa.info/result?qbase64=dGl0bGU9IueRnuWPi%2BWkqee%2FvO%2B8jeW6lOeUqOiZmuaLn%2BWMluezu%2Be7nyIgfHwgdGl0bGU9IueRnuWPi%2BW6lOeUqOiZmuaLn%2BWMluezu%2Be7nyI%3D) |
| **Number of assets affected**  | 55178 |
| **Description**  | Realor Tianyi Application Virtualization System is an application virtualization platform based on server computing architecture. It centrally deploys various user application software to the Ruiyou Tianyi service cluster, and clients can access authorized application software on the server through the WEB, achieving centralized application, remote access, collaborative office, and more. In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

## Realor Tianyi AVS ConsoleExternalApi.XGI file account param sql injection vulnerability

|   **Vulnerability**  | **Realor Tianyi AVS ConsoleExternalApi.XGI file account param sql injection vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 瑞友天翼应用虚拟化系统 ConsoleExternalApi.XGI account 参数 SQL 注入漏洞 |
| **CVSS core**  | 9.2 |
| **FOFA Query**  (click to view the results directly)| [title=\"瑞友天翼－应用虚拟化系统\" \|\| title=\"瑞友应用虚拟化系统\"](https://en.fofa.info/result?qbase64=dGl0bGU9IueRnuWPi%2BWkqee%2FvO%2B8jeW6lOeUqOiZmuaLn%2BWMluezu%2Be7nyIgfHwgdGl0bGU9IueRnuWPi%2BW6lOeUqOiZmuaLn%2BWMluezu%2Be7nyI%3D) |
| **Number of assets affected**  | 55178 |
| **Description**  | Realor Tianyi Application Virtualization System is an application virtualization platform based on server computing architecture. It centrally deploys various user application software to the Ruiyou Tianyi service cluster, and clients can access authorized application software on the server through the WEB, achieving centralized application, remote access, collaborative office, and more. Attackers can use this sql injection vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |
| **Impact** | Attackers can use this sql injection vulnerability to arbitrarily execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server. |


## Telesquare TLR-2005Ksh setSyncTimeHost RCE

|   **Vulnerability**  | **Telesquare TLR-2005Ksh setSyncTimeHost RCE**  |
| :----:   | :-----|
|  **Chinese name**  | Telesquare TLR-2005Ksh 路由器 setSyncTimeHost 命令执行漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [title=\"TLR-2005KSH\" \|\| banner=\"TLR-2005KSH login:\"](https://en.fofa.info/result?qbase64=dGl0bGU9IlRMUi0yMDA1S1NIIiB8fCBiYW5uZXI9IlRMUi0yMDA1S1NIIGxvZ2luOiI%3D) |
| **Number of assets affected**  | 25826 |
| **Description**  | Telesquare Tlr-2005Ksh is a Sk Telecom LTE router produced by Telesquare Korea. There is a security vulnerability in Telesquare TLR-2005Ksh, attackers can execute arbitrary commands through setSyncTimeHost to obtain server privileges. |
| **Impact** | There is a security vulnerability in Telesquare TLR-2005Ksh, attackers can execute arbitrary commands through setSyncTimeHost to obtain server privileges. |

![](https://s3.bmp.ovh/imgs/2023/05/11/99ee97cd77ca8767.gif)

## Telesquare TLR-2005Ksh getUsernamePassword Information Disclosure

|   **Vulnerability**  | **Telesquare TLR-2005Ksh getUsernamePassword Information Disclosure**  |
| :----:   | :-----|
|  **Chinese name**  | Telesquare TLR-2005Ksh 路由器 getUsernamePassword 信息泄露漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [title=\"TLR-2005KSH\" \|\| banner=\"TLR-2005KSH login:\"](https://en.fofa.info/result?qbase64=dGl0bGU9IlRMUi0yMDA1S1NIIiB8fCBiYW5uZXI9IlRMUi0yMDA1S1NIIGxvZ2luOiI%3D) |
| **Number of assets affected**  | 25826 |
| **Description**  | Telesquare Tlr-2005Ksh is a Sk Telecom LTE router produced by Telesquare Korea. There is a security hole in Telesquare TLR-2005Ksh. Attackers can obtain sensitive information such as username and password through getUsernamePassword. |
| **Impact** | There is a security hole in Telesquare TLR-2005Ksh. Attackers can obtain sensitive information such as username and password through getUsernamePassword. |

![](https://s3.bmp.ovh/imgs/2023/05/12/de09634d3b12cc18.gif)

## Telesquare TLR-2005Ksh ExportSettings.sh file download (CVE-2021-46423)

|   **Vulnerability**  | **Telesquare TLR-2005Ksh ExportSettings.sh file download (CVE-2021-46423)**  |
| :----:   | :-----|
|  **Chinese name**  | Telesquare TLR-2005Ksh 路由器 ExportSettings.sh 文件下载漏洞（CVE-2021-46423） |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=\"TLR-2005KSH\" \|\| banner=\"TLR-2005KSH login:\"](https://en.fofa.info/result?qbase64=dGl0bGU9IlRMUi0yMDA1S1NIIiB8fCBiYW5uZXI9IlRMUi0yMDA1S1NIIGxvZ2luOiI%3D) |
| **Number of assets affected**  | 25826 |
| **Description**  | Telesquare Tlr-2005K and so on are the Sk Telecom Lte routers of Korea Telesquare Company. There are security vulnerabilities in Telesquare TLR-2005Ksh, etc., which originate from unauthenticated file downloads. A remote attacker could exploit this vulnerability to download a complete configuration file. |
| **Impact** | There are security vulnerabilities in Telesquare TLR-2005Ksh, etc., which originate from unauthenticated file downloads. A remote attacker could exploit this vulnerability to download a complete configuration file. |

![](https://s3.bmp.ovh/imgs/2023/05/12/076acaa0dba4f960.gif)

## Sinovision Cloud CDN live default passwd

|   **Vulnerability**  | **Sinovision Cloud CDN live default passwd**  |
| :----:   | :-----|
|  **Chinese name**  | 华视私云-CDN直播加速服务器默认口令漏洞 |
| **CVSS core**  | 6.5 |
| **FOFA Query**  (click to view the results directly)| [body=\"src=\\\"img/dl.gif\\\"\" && title=\"系统登录\" && body=\"华视美达\"](https://en.fofa.info/result?qbase64=Ym9keT0ic3JjPVwiaW1nL2RsLmdpZlwiIiAmJiB0aXRsZT0i57O757uf55m75b2VIiAmJiBib2R5PSLljY7op4bnvo7ovr4i) |
| **Number of assets affected**  | 737 |
| **Description**  | CDN Live Broadcast Acceleration Server is a server for CDN live broadcast acceleration. The weak password vulnerability exists in the CDN Live broadcast acceleration server. The attacker can use the default password admin/admin to log in to the system background and obtain the background administrator permission. |
| **Impact** | attackers can control the entire platform through default password vulnerabilities and use administrator privileges to operate core functions. |

![](https://s3.bmp.ovh/imgs/2023/05/12/2d290c42299026fa.gif)

## WordPress plugin Welcart e-Commerce content-log.php logfile File Read Vulnerability

|   **Vulnerability**  | **WordPress plugin Welcart e-Commerce content-log.php logfile File Read Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress Welcart e-Commerce 插件 content-log.php 文件 logfile 参数文件读取漏洞 |
| **CVSS core**  | 9.8 |
| **FOFA Query**  (click to view the results directly)| [body=\"wp-content/plugins/usc-e-shop\"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL3VzYy1lLXNob3Ai) |
| **Number of assets affected**  | 5453 |
| **Description**  | Welcart is a free e-commerce plugin for WordPress with top market share in Japan.An arbitrary file read vulnerability exists in Welcart e-Commerce < 2.8.5, and attackers can exploit this vulnerability to obtain sensitive files. |
| **Impact** | Attackers can use this vulnerability to read the leaked source code, database configuration files, etc., resulting in an extremely insecure website. |

![](https://s3.bmp.ovh/imgs/2023/05/12/2474ac119a44c003.gif)

## secnet Intelligent Router actpt_5g.data Infoleakage

|   **Vulnerability**  | **secnet Intelligent Router actpt_5g.data Infoleakage**  |
| :----:   | :-----|
|  **Chinese name**  | secnet-智能路由系统 actpt_5g.data 信息泄露 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [title=\"安网-智能路由系统\" \|\| title==\"智能路由系统\" \|\| title=\"安网科技-智能路由系统\" \|\| banner=\"HTTPD_ac 1.0\" \|\| header=\"HTTPD_ac 1.0\"](https://en.fofa.info/result?qbase64=dGl0bGU9IuWuiee9kS3mmbrog73ot6%2FnlLHns7vnu58iIHx8IHRpdGxlPT0i5pm66IO96Lev55Sx57O757ufIiB8fCB0aXRsZT0i5a6J572R56eR5oqALeaZuuiDvei3r%2BeUseezu%2Be7nyIgfHwgYmFubmVyPSJIVFRQRF9hYyAxLjAiIHx8IGhlYWRlcj0iSFRUUERfYWMgMS4wIg%3D%3D) |
| **Number of assets affected**  | 71768 |
| **Description**  | secnet Intelligent AC management system is the wireless AP management system of Guangzhou Secure Network Communication Technology Co., LTD. (" Secure Network Communication "for short). The secnet intelligent AC management system has information vulnerabilities, which can be used by attackers to obtain sensitive information. |
| **Impact** | An attacker can use this vulnerability to obtain the WEB login account and password of the AC intelligent routing system and obtain the WEB administrator permission. As a result, sensitive information is leaked. |

![](https://s3.bmp.ovh/imgs/2023/05/12/f8eebf0ce38f975b.gif)

## SQL injection exists on Lotus ERP DictionaryEdit.aspx pag

|   **Vulnerability**  | **SQL injection exists on Lotus ERP DictionaryEdit.aspx pag**  |
| :----:   | :-----|
|  **Chinese name**  | 商混ERP系统 DictionaryEdit.aspx 页面存在SQL注入 |
| **CVSS core**  | 8.5 |
| **FOFA Query**  (click to view the results directly)| [title="商混ERP系统"](https://en.fofa.info/result?qbase64=dGl0bGU9IuWVhua3t0VSUOezu%2Be7nyI%3D) |
| **Number of assets affected**  | 616 |
| **Description**  | Hangzhou Lotus Software Co., Ltd. developed the commercial ERP system. This system mainly deals with the management of the mixing station of the construction company or various projects, including the sales module, production management module, laboratory module, personnel management, etc. The company's commercial concrete ERP system/Sys/DictionaryEdit dict at aspx_ SQL error injection vulnerability exists in the key parameter, which allows attackers to obtain database permissions. |
| **Impact** | In addition to taking advantage of SQL injection vulnerabilities to obtain information in the database (for example, administrator background password, site user personal information), attackers can even write Trojan horses to the server under high permissions to further obtain server system permissions. |

![](https://s3.bmp.ovh/imgs/2023/05/12/7fe36b3b6ee2d967.gif)

## V2Board admin.php Permission Bypass Vulnerability

|   **Vulnerability**  | **V2Board admin.php Permission Bypass Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | V2Board admin.php 越权访问漏洞 |
| **CVSS core**  | 8.0 |
| **FOFA Query**  (click to view the results directly)| [body="/theme/v2board/assets/umi.js"](https://en.fofa.info/result?qbase64=Ym9keT0iL3RoZW1lL3YyYm9hcmQvYXNzZXRzL3VtaS5qcyI%3D) |
| **Number of assets affected**  | 13299 |
| **Description**  | V2Board is a stable, simple, fast and easy to use multi-agent protocol management system.bV2Board v1.6.1 has an unauthorized access vulnerability. The authentication method is changed to obtain the cache from Redis to determine whether there is an interface that can be called. As a result, any user can call the interface with administrator privileges to obtain background privileges.  |
| **Impact** | Due to the lack of strict checks and restrictions on the user's access to the role, the current account can perform related operations on other accounts, such as viewing and modifying. |

## ZyXEL routers Export_Log arbitrary file read

|   **Vulnerability**  | **ZyXEL routers Export_Log arbitrary file read**  |
| :----:   | :-----|
|  **Chinese name**  | ZyXEL 路由器 Export_Log 任意文件读取 |
| **CVSS core**  | 8.0 |
| **FOFA Query**  (click to view the results directly)| [(title=".:: Welcome to the Web-Based Configuration::." && body="ZyXEL") \|\| (title="Welcome to the Web-Based Configurator" && (body="/zycss.css" \|\| body="zyxel")) \|\| title="do Router ZyXEL" \|\| title="Welcome to ZyROUTER" \|\| title="ZyXEL Router" \|\| body="\<friendlyName>ZyXEL Router\</friendlyName>" \|\| banner="ZyXEL-router"](https://en.fofa.info/result?qbase64=KHRpdGxlPSIuOjogV2VsY29tZSB0byB0aGUgV2ViLUJhc2VkIENvbmZpZ3VyYXRpb246Oi4iICYmIGJvZHk9Ilp5WEVMIikgfHwgKHRpdGxlPSJXZWxjb21lIHRvIHRoZSBXZWItQmFzZWQgQ29uZmlndXJhdG9yIiAmJiAoYm9keT0iL3p5Y3NzLmNzcyIgfHwgYm9keT0ienl4ZWwiKSkgfHwgdGl0bGU9ImRvIFJvdXRlciBaeVhFTCIgfHwgdGl0bGU9IldlbGNvbWUgdG8gWnlST1VURVIiIHx8IHRpdGxlPSJaeVhFTCBSb3V0ZXIiIHx8IGJvZHk9IjxmcmllbmRseU5hbWU%2BWnlYRUwgUm91dGVyPC9mcmllbmRseU5hbWU%2BIiB8fCBiYW5uZXI9Ilp5WEVMLXJvdXRlciIK) |
| **Number of assets affected**  | 733803 |
| **Description**  | ZyXEL routers are various router products of ZyXEL company. Several ZyXEL routers have an arbitrary file read vulnerability in /Export_Log. |
| **Impact** | Several ZyXEL routers have an arbitrary file read vulnerability in /Export_Log. |

![](https://s3.bmp.ovh/imgs/2023/05/12/3909e1c09af8eb25.gif)

## ZXHN H108NS Router tools_admin.asp Permission Bypass Vulnerability

|   **Vulnerability**  | **ZXHN H108NS Router tools_admin.asp Permission Bypass Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 中兴 H108NS 路由器 tools_admin.asp 文件权限绕过漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [banner=\"Basic realm=\\\"H108NS\\\"\" \|\| header=\"Basic realm=\\\"H108NS\\\"\"](https://en.fofa.info/result?qbase64=YmFubmVyPSJCYXNpYyByZWFsbT1cIkgxMDhOU1wiIiB8fCBoZWFkZXI9IkJhc2ljIHJlYWxtPVwiSDEwOE5TXCIi) |
| **Number of assets affected**  | 8245 |
| **Description**  | ZTE H108NS router is a router product that integrates WiFi management, route allocation, dynamic access to Internet connections and other functions. The ZTE H108NS router has an identity authentication bypass vulnerability. An attacker can use this vulnerability to bypass identity authentication and allow access to the router's management panel to modify the administrator password to obtain sensitive user information. |
| **Impact** | An attacker can use this vulnerability to bypass identity authentication and allow access to the management panel of the router to modify the administrator password and obtain sensitive information of the user. |

![](https://s3.bmp.ovh/imgs/2023/04/27/79bd32a6c7ab12b2.gif)

## ezOFFICE OA OfficeServer.jsp Arbitrarily File Upload Vulnerability

|   **Vulnerability**  | **ezOFFICE OA OfficeServer.jsp Arbitrarily File Upload Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 万户 OA OfficeServer.jsp 任意文件上传漏洞 |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [(banner=\"OASESSIONID\" && banner=\"/defaultroot/\") \|\| (header=\"OASESSIONID\" && header=\"/defaultroot/\")\|\|body=\"/defaultroot/themes/common/common.css\"\|\|body=\"ezofficeDomainAccount\"\|\|title=\"Wanhu ezOFFICE\" \|\| title=\"万户ezOFFICE\"](https://en.fofa.info/result?qbase64=KGJhbm5lcj0iT0FTRVNTSU9OSUQiICYmIGJhbm5lcj0iL2RlZmF1bHRyb290LyIpIHx8IChoZWFkZXI9Ik9BU0VTU0lPTklEIiAmJiBoZWFkZXI9Ii9kZWZhdWx0cm9vdC8iKXx8Ym9keT0iL2RlZmF1bHRyb290L3RoZW1lcy9jb21tb24vY29tbW9uLmNzcyJ8fGJvZHk9ImV6b2ZmaWNlRG9tYWluQWNjb3VudCJ8fHRpdGxlPSJXYW5odSBlek9GRklDRSIgfHwgdGl0bGU9IuS4h%2BaIt2V6T0ZGSUNFIg%3D%3D) |
| **Number of assets affected**  | 4715 |
| **Description**  | ezOFFICE OA is a FlexOffice independent security cooperative office platform for government organizations, enterprises and institutions. ezOFFICE OA OfficeServer There is an arbitrary file upload vulnerability in jsp, through which an attacker can upload arbitrary files to control the entire server. |
| **Impact** | File upload vulnerabilities are usually caused by the lax filtering of files uploaded by the file upload function in the code or the unrepaired parsing vulnerabilities related to the web server. Attackers can upload arbitrary files through the file upload point, including the website backdoor file (webshell), to control the entire website. |

![](https://s3.bmp.ovh/imgs/2023/04/27/4287e0695068b5fe.gif)

## seaflysoft ERP getylist_login.do SQL Injection

|   **Vulnerability**  | **seaflysoft ERP getylist_login.do SQL Injection**  |
| :----:   | :-----|
|  **Chinese name**  | 海翔云平台 getylist_login.do SQL 注入漏洞 |
| **CVSS core**  | 8.0 |
| **FOFA Query**  (click to view the results directly)| [body=\"checkMacWaitingSecond\"](https://en.fofa.info/result?qbase64=Ym9keT0iY2hlY2tNYWNXYWl0aW5nU2Vjb25kIg%3D%3D) |
| **Number of assets affected**  | 773 |
| **Description**  | seaflysoft cloud platform one-stop overall solution provider, business covers wholesale, chain, retail industry ERP solutions, wms warehousing solutions, e-commerce, field work, mobile terminal (PDA, APP, small program) solutions. There is a SQL injection vulnerability in the system getylist_login.do, through which an attacker can obtain database permissions |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

![](https://s3.bmp.ovh/imgs/2023/05/12/416d432cd922426c.gif)

## MCMS list Interface sqlWhere Sql Injection Vulnerability

|   **Vulnerability**  | **MCMS list Interface sqlWhere Sql Injection Vulnerability**  |
| :----:   | :-----|
|  **Chinese name**  | 铭飞 CMS list 接口 sqlWhere 参数 sql 注入漏洞 |
| **CVSS core**  | 7.5 |
| **FOFA Query**  (click to view the results directly)| [body=\"铭飞MCMS\" \|\| body=\"/mdiy/formData/save.do\" \|\| body=\"static/plugins/ms/1.0.0/ms.js\"](https://en.fofa.info/result?qbase64=Ym9keT0i6ZOt6aOeTUNNUyIgfHwgYm9keT0iL21kaXkvZm9ybURhdGEvc2F2ZS5kbyIgfHwgYm9keT0ic3RhdGljL3BsdWdpbnMvbXMvMS4wLjAvbXMuanMi) |
| **Number of assets affected**  | 3091 |
| **Description**  | MCMS is a set of lightweight open source content management system developed based on java. It is simple, safe, open source and free. It can run on Linux, Windows, MacOSX, Solaris and other platforms. The system has an sql injection vulnerability before the 5.2.10 version. You can use this vulnerability to obtain sensitive information |
| **Impact** | In addition to using SQL injection vulnerabilities to obtain information in the database (for example, the administrator's back-end password, the user's personal information of the site), an attacker can write a Trojan horse to the server even in a high-privileged situation to further obtain server system permissions. |

![](https://s3.bmp.ovh/imgs/2023/05/04/9119224cdf0a37f4.gif)

## WordPress booking-calendar admin-ajax.php File Upload (CVE-2022-3982)

|   **Vulnerability**  | **WordPress booking-calendar admin-ajax.php File Upload (CVE-2022-3982)**  |
| :----:   | :-----|
|  **Chinese name**  | WordPress booking-calendar 插件 admin-ajax.php 任意文件上传漏洞（CVE-2022-3982） |
| **CVSS core**  | 9.0 |
| **FOFA Query**  (click to view the results directly)| [body=\"wp-content/plugins/booking-calendar/\"](https://en.fofa.info/result?qbase64=Ym9keT0id3AtY29udGVudC9wbHVnaW5zL2Jvb2tpbmctY2FsZW5kYXIvIg%3D%3D) |
| **Number of assets affected**  | 1074 |
| **Description**  | WordPress booking-calendar is a plugin for creating booking system scheduling calendars for WordPress sites. WordPress Plugin Booking Calendar versions before 3.2.2 have a code problem vulnerability. The vulnerability stems from the fact that the plugin does not verify uploaded files and allows unauthenticated users to upload arbitrary files. Attackers can exploit this vulnerability to achieve RCE. |
| **Impact** | WordPress Plugin Booking Calendar versions before 3.2.2 have a code problem vulnerability. The vulnerability stems from the fact that the plugin does not verify uploaded files and allows unauthenticated users to upload arbitrary files. Attackers can exploit this vulnerability to achieve RCE. |

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
](https://fofa.info/result?qbase64=KHRpdGxlPSIuOjogV2VsY29tZSB0byB0aGUgV2ViLUJhc2VkIENvbmZpZ3VyYXRpb246Oi4iICYmIGJvZHk9Ilp5WEVMIikgfHwgKHRpdGxlPSJXZWxjb21lIHRvIHRoZSBXZWItQmFzZWQgQ29uZmlndXJhdG9yIiAmJiAoYm9keT0iL3p5Y3NzLmNzcyIgfHwgYm9keT0ienl4ZWwiKSkgfHwgdGl0bGU9ImRvIFJvdXRlciBaeVhFTCIgfHwgdGl0bGU9IldlbGNvbWUgdG8gWnlST1VURVIiIHx8IHRpdGxlPSJaeVhFTCBSb3V0ZXIiIHx8IGJvZHk9IjxmcmllbmRseU5hbWU%2BWnlYRUwgUm91dGVyPC9mcmllbmRseU5hbWU%2BIiB8fCBiYW5uZXI9Ilp5WEVMLXJvdXRlciIK)

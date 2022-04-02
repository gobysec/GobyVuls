
# Spring Cloud Function SPEL Vulnerability

Spring cloud function is a function calculation framework based on spring boot. By abstracting the transmission details and infrastructure, it retains familiar development tools and development processes for developers, so that developers can focus on realizing business logic, so as to improve development efficiency.There is spring in the HTTP request header for accessing spring cloud function&nbsp;cloud.&nbsp;function.&nbsp;Routing expression parameter, whose spel expression can be injected and executed through StandardeValuationContext parsing.&nbsp;Eventually, an attacker can perform remote command execution through this vulnerability.

FOFA **query rule**: [((header="Server: Netty@SpringBoot" || (body="Whitelabel Error Page" && body="There was an unexpected error")) && body!="couchdb") || title="SpringBootAdmin-Server" || body="SpringBoot"](https://fofa.info/result?qbase64=KChoZWFkZXI9IlNlcnZlcjogTmV0dHlAU3ByaW5nQm9vdCIgfHwgKGJvZHk9IldoaXRlbGFiZWwgRXJyb3IgUGFnZSIgJiYgYm9keT0iVGhlcmUgd2FzIGFuIHVuZXhwZWN0ZWQgZXJyb3IiKSkgJiYgYm9keSE9ImNvdWNoZGIiKSB8fCB0aXRsZT0iU3ByaW5nQm9vdEFkbWluLVNlcnZlciIgfHwgYm9keT0iU3ByaW5nQm9vdCI%3D)

# Demo

![Spring_Cloud_Function_SPEL_Vulnerability](Spring_Cloud_Function_SPEL_Vulnerability.gif)

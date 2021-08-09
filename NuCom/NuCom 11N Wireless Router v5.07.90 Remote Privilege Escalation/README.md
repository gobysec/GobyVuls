# NuCom 11N Wireless Router v5.07.90 Remote Privilege Escalation

The application suffers from a privilege escalation vulnerability. The non-privileged default user (user:user) can elevate his privileges by sending a HTTP GET request to the configuration backup endpoint and disclose the http super password (admin credentials) in Base64 encoded value. Once authenticated as admin, an attacker will be granted access to the additional and privileged pages.

FOFA **query rule**: [title="NuCom 11N Wireless Router"||body="NuCom 11N Wireless Router"](https://fofa.so/result?qbase64=dGl0bGU9Ik51Q29tIDExTiBXaXJlbGVzcyBSb3V0ZXIifHxib2R5PSJOdUNvbSAxMU4gV2lyZWxlc3MgUm91dGVyIg%3D%3D)

# Demo

![](NuCom_11N_Wireless_Router_V5_07_Remote_Privilege_Escalation.gif)


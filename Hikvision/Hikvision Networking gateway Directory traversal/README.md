# Hikvision Networking gateway Directory traversal

Hikvision networking gateway has a directory traversal vulnerability, use ::data to read sensitive information.

FOFA **query rule**: [body="通用系统" && body="data/login.php"](https://fofa.so/result?qbase64=Ym9keT0i6YCa55So57O757ufIiAmJiBib2R5PSJkYXRhL2xvZ2luLnBocCI%3D)

# Demo

![Hikvision_Networking_gateway_Directory_traversal.gif](Hikvision_Networking_gateway_Directory_traversal.gif)

# Ricon Industrial Cellular Router apply.cgi RCE

The router suffers from an authenticated OS command injection vulnerability, This can be exploited to inject and execute arbitrary shell commands as the admin user via the ping_server_ip POST parameter. Also vulnerable to Heartbleed.

FOFA **query rule**: [body="Industrial Cellular" && server="WEB-ROUTER"](https://fofa.so/result?qbase64=Ym9keT0iSW5kdXN0cmlhbCBDZWxsdWxhciIgJiYgc2VydmVyPSJXRUItUk9VVEVSIg%3D%3D)

# Demo

![](Ricon_Industrial_Cellular_Router_apply_cgi_RCE.gif)


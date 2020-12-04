# Apache NiFi Api RCE

This module uses the NiFi API to create an ExecuteProcess processor that will execute OS commands. The API must be unsecured (or credentials provided) and the ExecuteProcess processor must be available. An ExecuteProcessor processor is created then is configured with the payload and started. The processor is then stopped and deleted.

**Affected version**: Apache NiFi

**[FOFA](https://fofa.so/result?q=title%3D%22NiFi%22&qbase64=dGl0bGU9Ik5pRmki&file=&file=) query rule**: title="NiFi"

# Demo

![](Apache_Nifi_Api_RCE.gif)

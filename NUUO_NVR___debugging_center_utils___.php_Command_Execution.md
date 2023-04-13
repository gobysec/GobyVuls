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

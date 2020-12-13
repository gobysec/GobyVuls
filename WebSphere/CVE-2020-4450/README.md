# CVE-2020-24571 NexusDB path traversal

NexusQA NexusDB before 4.50.23 allows the reading of files via ../ directory traversal.

**Affected version**: nexusdb < 4.50.23

**[FOFA](https://fofa.so/result?q=header%3D%22Server%3A+NexusDB%22&qbase64=aGVhZGVyPSJTZXJ2ZXI6IE5leHVzREIi&file=&file=) query rule**: header="Server: NexusDB"

# Demo

![](CVE-2020-24571.gif)

# pfSense Arbitrary File Write to RCE

diag_routes.php in pfSense 2.5.2 allows sed data injection. The data is retrieved by executing the netstat utility, and then its output is parsed via the sed utility.

FOFA **query rule**: [app="pfSense"](https://fofa.info/result?qbase64=YXBwPSJwZlNlbnNlIg%3D%3D)

# Demo

![pfSense_Arbitrary_File_Write_to_RCE](pfSense_Arbitrary_File_Write_to_RCE.gif)

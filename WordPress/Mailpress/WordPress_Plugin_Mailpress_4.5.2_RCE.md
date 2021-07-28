# WordPress Plugin Mailpress 4.5.2 RCE

In the WordPress Mailpress Plugin, the subject parameter in the iview function in the mailpress/mp-includes/class/MP_Actions.class.php file is not filtered, and pass to do_eval function, leading to remote code execution.

**Affected version**: WordPress Plugin Mailpress <= 4.5.2

**FOFA query rule**: [app="WordPress"](https://fofa.so/result?qbase64=YXBwPSJXb3JkUHJlc3Mi)

# Demo

![](WordPress_Plugin_Mailpress_4.5.2_RCE.gif)
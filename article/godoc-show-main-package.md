# godoc 显示main包的文档


```diff
diff --git a/godoc/server.go b/godoc/server.go
index bbc3409..631d12b 100644
--- a/godoc/server.go
+++ b/godoc/server.go
@@ -174,7 +174,8 @@ func (h *handlerServer) GetPageInfo(abspath, relpath string, mode PageInfoMode)
                        }
                        info.PAst = files
                }
-               info.IsMain = pkgname == "main"
+               // info.IsMain = pkgname == "main"
+               info.IsMain = false && pkgname == "main"
        }

        // get directory information, if any
(END)

```

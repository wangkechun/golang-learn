# 修改go编译器允许代码存在未使用的变量和包

先直接看修改之后的结果，这样一份代码
```go
package main

import "fmt"
import "log"

func main() {
	a := 1
	log.Println("aaaaa")
}
```
编译结果如下：
```bash
➜  test  go run w.go
# command-line-arguments
./w.go:3: imported and not used2: "fmt"
./w.go:7: a declared and not used2
2015/09/21 17:36:34 aaaaa
```
可以看到，虽然提示了代码的问题，但是还是编译并运行成功了。

接下来就告诉大家如何修改。
首先确认你的golang版本和我的一样，
```bash
➜  test  go version
go version go1.5.1 darwin/amd64
```
然后进入$GOROOT目录(golang安装目录)， 按照下面的diff修改即可。

```diff
diff --git a/src/cmd/compile/internal/gc/lex.go b/src/cmd/compile/internal/gc/lex.go
index 606298b..ef9b59a 100644
--- a/src/cmd/compile/internal/gc/lex.go
+++ b/src/cmd/compile/internal/gc/lex.go
@@ -2584,9 +2584,9 @@ func pkgnotused(lineno int, path string, name string) {
                elem = elem[i+1:]
        }
        if name == "" || elem == name {
-               yyerrorl(int(lineno), "imported and not used: %q", path)
+               adderr(int(lineno), "imported and not used2: %q", path)
        } else {
-               yyerrorl(int(lineno), "imported and not used: %q as %s", path, name)
+               adderr(int(lineno), "imported and not used3: %q as %s", path, name)
        }
 }

diff --git a/src/cmd/compile/internal/gc/subr.go b/src/cmd/compile/internal/gc/subr.go
index 866d8e1..56d3875 100644
--- a/src/cmd/compile/internal/gc/subr.go
+++ b/src/cmd/compile/internal/gc/subr.go
@@ -348,7 +348,7 @@ func importdot(opkg *Pkg, pack *Node) {

        if n == 0 {
                // can't possibly be used - there were no symbols
-               yyerrorl(int(pack.Lineno), "imported and not used: %q", opkg.Path)
+               adderr(int(pack.Lineno), "imported and not used1: %q", opkg.Path)
        }
 }

diff --git a/src/cmd/compile/internal/gc/walk.go b/src/cmd/compile/internal/gc/walk.go
index af3e1cc..f9c35b0 100644
--- a/src/cmd/compile/internal/gc/walk.go
+++ b/src/cmd/compile/internal/gc/walk.go
@@ -51,11 +51,11 @@ func walk(fn *Node) {
                                continue
                        }
                        lineno = defn.Left.Lineno
-                       Yyerror("%v declared and not used", l.N.Sym)
+                       adderr(parserline(), "%v declared and not used", l.N.Sym)
                        defn.Left.Used = true // suppress repeats
                } else {
                        lineno = l.N.Lineno
-                       Yyerror("%v declared and not used", l.N.Sym)
+                       adderr(parserline(), "%v declared and not used2", l.N.Sym)
                }
        }

(END)
```
然后
```
➜  test  go build -x cmd/compile
WORK=/var/folders/5x/bg9pbd6x5dv70kt_mq5ys16w0000gn/T/go-build774122134
mkdir -p $WORK/cmd/internal/gcprog/_obj/
mkdir -p $WORK/cmd/internal/
cd /usr/local/go/src/cmd/internal/gcprog
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/cmd/internal/gcprog.a -trimpath $WORK -p cmd/internal/gcprog -complete -buildid f02cd5f66694d70425b46f6ebd53b46222f5086a -D _/usr/local/go/src/cmd/internal/gcprog -I $WORK -pack ./gcprog.go
mkdir -p $WORK/cmd/compile/internal/big/_obj/
mkdir -p $WORK/cmd/compile/internal/
cd /usr/local/go/src/cmd/compile/internal/big
......
```
这样会在当前目录生成compile。然后上面输出可以找到 `/usr/local/go/pkg/tool/darwin_amd64/compile `这个二进制文件，
用当前目录生成的compile覆盖之前的就ok了。


文章地址 https://github.com/wangkechun/golang-learn/blob/master/article/go-compile-allowed-unused-packages-and-vars.md

# 修改go编译器允许没有使用的变量和包

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

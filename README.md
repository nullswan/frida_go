```Go
file, _ := ioutil.ReadFile(__file__)
session, _ := frida_go.Attach(ProcessID)
script, _ := session.CerateScriptSync(name, string(file))
_ := script.Load()
```
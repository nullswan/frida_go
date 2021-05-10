```Go
file, _ := ioutil.ReadFile(__file__)
session, _ := frida_go.Attach(ProcessID)
script, _ := session.CreateScriptSync(name, string(file))
_ := script.Load()
```
currentpath = createobject("Scripting.FileSystemObject").GetFile(Wscript.ScriptFullName).ParentFolder.Path
set ws=WScript.CreateObject("WScript.Shell")
ws.Run currentpath &"\keyhook.bat",0
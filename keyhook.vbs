currentpath = createobject("Scripting.FileSystemObject").GetFile(Wscript.ScriptFullName).ParentFolder.Path &"keyhook.bat"
CreateObject("WScript.Shell").Run currentpath,0 
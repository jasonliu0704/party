package MetadataStore

import "bytes"
import "data"


type File struct{
  version int
  fileCache map[string]data.FileInfo
}

type MetadataStore struct{
  blockStore BlockStore
  fileCollection map[string]data.FileInfo
}

/*
Read and return the requested file
return FileInfo object to the client
*/
(MetadataStore s) readFile(fname string){

  // get file blocks
  if val, ok := s.fileCollection[fname]; ok {
    return val
  }

  // error handling
}

/*
client upload a file, server find out missing block and return the missing block
for client to upload
*/
(MetadataStore s) modifyFile(file data.FileInfo){
  upload_version := file.version
  cur_version := 0
  fname := file.filename

  if(f,ok := s.fileCollection[fname]; ok) cur_version := f.version

  // the request file version must be one bigger than the current version
  if(((cur_version != 0) && (version != (cur_version + 1))) || (version < 0)){
      response = builder.setResult(WriteResult.Result.OLD_VERSION).build();
      responseObserver.onNext(response);
      responseObserver.onCompleted();
      return;
  }
}

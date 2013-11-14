	//- \subsection{The File System Interface}
	//--------------------------------------------------

	//- Conventional programs expect a conventional file system interface.  
	//- The /texttt{msfsfs/msfsfiles.go} file implements the file system functionality of \texttt{msfs}:

// msfsfiles (in msfsfiles.go) implements the msfs fuse interface
package msfsfiles
//=

	//- In addition to the standard \texttt{go} library, \texttt{fs.go} relies on functionality provided 
	//- by the \texttt{fuse} package:

// dependencies

import (
	"fmt"
	"os"

	"rputbl.com/msfs/msfshashes"
	"rputbl.com/msfs/msfscas"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)
//=

	//- \subsubsection{File System Structures and Methods}
	//--------------------------------------------------
 
	//- The \texttt{FS} structure holds the reference to the root directory and global information
	//-  about the state of the file system.

// FS encapsulates whole-file-system features
type FS struct{
	storeversion, execversion,serverloc string
	Hi *msfshashes.HSC
	Cas	*msfscas.CASC
	root *Dir
	Own *Dir
}
//=

	//- The file system as a whole is initialized when the system starts.

//    Prepare() sets up the file system for use.
func (msfs *FS) Prepare( hi *msfshashes.HSC, cas *msfscas.CASC) fuse.Error {
	msfs.execversion = "msfs.0.0.1"
	msfs.storeversion = "n/a"
	msfs.serverloc = "some server"
	msfs.Hi = hi
	msfs.Cas = cas

	msfs.root = new(Dir)
	msfs.Own = new(Dir)

	msfs.root.Prepare( []fuse.Dirent{ 
		{Inode: 2, Name: ".own", Type: fuse.DT_Dir}, 
		{Inode: 3, Name: ".status", Type: fuse.DT_File}, 
		{Inode: 8, Name: "testfile", Type: fuse.DT_File}, },msfs)

	msfs.Own.Prepare( []fuse.Dirent{ 
		{Inode: 4, Name: "privateKey", Type: fuse.DT_File}, 
		{Inode: 5, Name: "publicKey", Type: fuse.DT_File}, },msfs)



	fmt.Fprintf(os.Stderr, "Status: %s\n",msfs.Cas.Status)

	return  nil
}
//=

	//- Root returns the root directory to the FUSE interface.

//    
func (msfs *FS) Root() (fs.Node, fuse.Error) {
	return msfs.root, nil
}
//=


	//- \subsubsection{Directory Structures and Methods}
	//--------------------------------------------------

	//- \texttt{msfs} maps file system directory operations on hashes of the names of the contents of the directories.


// Dir implements both Node and Handle for the root directory
type Dir struct{
	DirList  []fuse.Dirent
	Fs *FS
}

func (msd *Dir) Prepare(contents []fuse.Dirent, Fs *FS)  fuse.Error {
	msd.DirList = contents
	msd.Fs = Fs
	return nil
}


func (Dir) Attr() fuse.Attr {
	return fuse.Attr{Inode: 1, Mode: os.ModeDir | 0555}
}

func (msd *Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	return msd.DirList, nil
}

func (msd *Dir) Lookup(name string, intr fs.Intr) (fs.Node, fuse.Error) {
	if name == ".own" {
		return msd.Fs.Own, nil
	}else{
		for _,v:= range msd.DirList{
			if name == v.Name {
				f := new(File)
				f.Prepare( v.Inode, v.Name, msd )
				return f, nil
			}
		}
	}

	return nil, fuse.ENOENT
}


//=

	//- \subsubsection{File Structures and Methods}
	//--------------------------------------------------

// File implements both Node and Handle for the hello file
type File struct{
	Inode uint64
	Name string
	Msd *Dir
}

func (fi *File) Prepare(i uint64, n string, Msd *Dir)  fuse.Error {
	fi.Inode = i
	fi.Name = n
	fi.Msd = Msd
	return nil
}

func (File) Attr() fuse.Attr {
	return fuse.Attr{Mode: 0444}
}

func (fi File) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	if fi.Inode == 3{
		return []byte(fmt.Sprintf("CAS Client State: %s\nCAS Client Mode: %s\n",fi.Msd.Fs.Cas.Status,fi.Msd.Fs.Cas.Mode)), nil
	}
	if fi.Inode == 5{
		return []byte(fi.Msd.Fs.Hi.PubCert), nil
	}
	return []byte(fmt.Sprintf("Test File Contents for file %s!\n",fi.Name)), nil

}
//=


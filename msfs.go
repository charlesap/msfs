//bin/echo "Creating latex documents from go source"  >&2 ;F () { echo " processing $1" >&2;\
//bin/cat $1 | sed '	/\/\/----------/d'|\
//bin/sed 's/\/\/ \(.*\)/\\begin{lstlisting}[title={{\\color{blue}\1}}]/g' |\
//bin/sed 's/\/\/-\(.*\)/ \1/g' |\
//bin/sed 's/\/\/=\(.*\)/\\end{lstlisting}/g' |\
//usr/bin/tr '~' '\\' > $1.tex; \
//bin/true ; } ;\
//bin/cat $0 | sed 1,11d | sed 38d > $0.trimmed ; F $0.trimmed ;\
//bin/true; for d in */; do  for f in $d*.go ; do \
//bin/true; F $f ; done ; done; mv $0.trimmed.tex $0.latex; rm $0.trimmed; exit
/*

~documentclass[11pt]{article}
~setlength{~topmargin}{-.5in}
~setlength{~textheight}{9in}
~setlength{~oddsidemargin}{.125in}
~setlength{~textwidth}{6.25in}
~setlength{~parskip=4pt}{~baselineskip=0pt}%
~setcounter{secnumdepth}{3}
~usepackage{tikz}
~usepackage{datetime}
~usepackage{listings}
~usepackage{courier}
~usepackage{color}
~usepackage{changepage}
~lstset{
%language=C,
basicstyle=~small~sffamily,
numbers=none,
basicstyle=\footnotesize\ttfamily,
columns=fullflexible,showstringspaces=false,
frame=single,                    % adds a frame around the code
keepspaces=true,                 % keeps spaces in text, useful for keeping indentation of code (possibly needs %columns=flexible,
%captionstyle=\color{blue},
keywordstyle=\color{blue},       % keyword style
%language=c,                 % the language of the code
morekeywords={import,var,func,type,package,return,struct}, 
}
~begin{document}
~title{A Mutually Suspicious File System}
~author{Charles Perkins~~
Chief Cypherpunk and Brewmaster, Kuracali Sake and Beer Brewery}
~renewcommand{~today}{
~shortmonthname{} ~twodigit{~day} 2013} 
~maketitle
~part*{}


//- */

	//- Traditional computing systems delegate the protection of stored data to a trusted, priveleged party
	//- (e.g. the kernel, the supervisor, the administrator, the sysadmin, etc.) This trust is not always warranted.
	//- The Mutually Suspicious File System
	//- demonstrates a method by which the user-level process may retain the authority and responsibility for access
	//- control of stored\footnote{Protecting data and computations in memory from the kernel and from other users
	//- is also critical but is beyond the scope of this paper.} data.
	//- 
	//- This paper presents an implementation of \texttt{msfs} in the \texttt{go}\footnote{
	//- Team, Go. The Go programming language specification. Technical Report. http://golang.org/ref/spec  (retrieved Oct 2013.)} 
	//-  programming language. The program source code is open source and freely redistributable under the MIT license:
	//- 
	//- \begin{adjustwidth}{.2in}{.5in}
	//- {\scriptsize \texttt{Copyright (c) 2013 Charles Perkins}
	//- 
	//- \texttt{Permission is hereby granted, free of charge, to any person
	//- obtaining a copy of this software and associated documentation
	//- files (the "Software"), to deal in the Software without
	//- restriction, including without limitation the rights to use,
	//- copy, modify, merge, publish, distribute, sublicense, and/or sell
	//- copies of the Software, and to permit persons to whom the
	//- Software is furnished to do so, subject to the following
	//- conditions:}
	//- 
	//- \texttt{The above copyright notice and this permission notice shall be
	//- included in all copies or substantial portions of the Software.}
	//- 
	//- \texttt{THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	//- EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
	//- OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	//- NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	//- HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
	//- WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
	//- FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
	//- OTHER DEALINGS IN THE SOFTWARE.}}
	//- \end{adjustwidth}

	//- The program text is presented in the body of this paper like so:


	//- \definecolor{Color1}{rgb}{.9,.9,.55}
	//- \lstset{backgroundcolor=\color{Color1}}


// about this code
//- source code statements 
//=


	//- This  {\LaTeX} type-set paper {\it [A Mutually Suspicous...]} is drawn from the comments in the \texttt{go} source code via 
	//- a \texttt{bash}-executable but \texttt{go}- and \texttt{godoc}-invisible preamble, unifying the 
	//- program and the documentation
	//- in one text and thereby forming a minimally capable literate programming
	//- environment.\footnote{Donald E. Knuth, Literate Programming, Stanford, California: 
	//- 1992, CLSI Lecture Notes, no. 27.}


	//- \setcounter{section}{1}

	//--------------------------------------------------
	//- \section{Paper Overview}
	//--------------------------------------------------

	//- A successful program design may be approached in stages with each stage buidling on concepts and capabilities 
	//- introduced in an earlier stage. An approcahable paper likewise may sequentially introduce and expand upon  
	//- concepts and assertions. This paper is structured in several sections:  

	//- \begin{enumerate}\itemsep1pt
	//- \item Introduction
	//- \item Paper Overview
	//- \item Theory of Operation
	//- \item Subsystems
	//- \item Program Flow
	//- \item File System Usage
	//- \item Literate Coding Conventions
	//- \end{enumerate}

	//- By the end of this paper the reader should have a good understanding of how a mutually suspicious file system
	//- may be implemented.


	//--------------------------------------------------
	//- \section{Theory of Operation}
	//--------------------------------------------------

	//- \texttt{msfs} is based on the observation that public key cryptography, 
	//- symmetric key ciphers and hash algorithms may combine to allow mutually suspicous use of a shared resource:

	//- \begin{itemize}\itemsep1pt
	//- \item RSA/DSA certificates provide user identity and introduction
	//- \item Symmetric-key ciphers provide efficiency and security
	//- \item Hash algorithms provide object reference and retrieval
	//- \item Content-Addressable Storage serializes and persists data to block storage
	//- \end{itemize}

	//- A content-addressable hash store is the shared resource in the case of the \texttt{msfs} file system. 
	//-  The hash store is not a privileged agent--it is not given the ability to decrypt stored values unless the 
	//- values are intended to be publically accessible. The hash store serializes
	//- the access requests of multiple readers and writers and persists the values to block storage.  

	//- The theory of operations can be further broken down as follows:

	//- \begin{enumerate}\itemsep1pt
	//- \item File System Operations
	//- \item \texttt{msfs} Notional Structures 
	//- \item Reading
	//- \item Writing 
	//- \item Access Control
	//- \item Notional to Block Mapping
	//- \item Serializing Access
	//- \end{enumerate}


	//- \subsection{File System Operations}
	//--------------------------------------------------

	//- \texttt{msfs} should appear to be a conventional file system from the point of view of the user and the user's 
	//- executing programs. FUSE (File System in User Space) is a protocol specification and operating system interface 
	//- supported on Linux and OSX that allows user-level programs to appear to other user-level programs as a file system.
	//- The \texttt{msfs} program implements this interface in order to implement file-system functionality for other programs.


	//- To implement the FUSE interface the program responds to the following requests:
	//-


	//- \begin{itemize}\itemsep1pt
	//- \item (FS) Root() -- return the root directory
	//- \item (Dir) Attr() -- return directory attributes
	//- \item (Dir) Lookup( ) -- return a file entry
	//- \item (Dir) ReadDir( ) -- return an array of directory entries
	//- \item (File) Attr() -- return the attributes of a file
	//- \item (File) ReadAll( ) -- return the contents of a file
	//- \end{itemize}


	//- \subsection{Notional Structures}
	//--------------------------------------------------

	//- The \texttt{msfs} user-level client will maintain several notional structures that will cache information
	//- and assist in the translation of file i/o operations to block-level operations.

	//- \subsection{Reading}
	//--------------------------------------------------

	//- Reading from a file or a directory involves several steps.

	//- \subsection{Writing}
	//--------------------------------------------------

	//- Creating a file and writing to it involves several steps.

	//- \subsection{Access Control}
	//--------------------------------------------------

	//- Changes in access control are made by writing certificates and distributing keys.

	//- \subsection{Notional to Block Mapping}
	//--------------------------------------------------

	//- The notional structures maintained by the client are mapped to blocks-in-storage.

	//- \subsection{Serializing Access}
	//--------------------------------------------------

	//- Block-level hash storage is mediated by a goroutine and accessed over TCP/IP.

	//--------------------------------------------------
	//- \section{Subsystems}
	//--------------------------------------------------

	//- The \texttt{msfs} program includes several packages, each of which ecapsulates a portion
	//- of the program's functionaltiy, each in a separate file:


	//- \begin{enumerate}\itemsep1pt
	//- \item The File System Interface (in \texttt{msfsfiles.go})
	//- \item Hashing and Encryption (in \texttt{msfshashes.go})
	//- \item Block Storage (in \texttt{msfscas.go})
	//- \end{enumerate}

	//- The above files are imported by the main program executable, described in Section 5, Program Flow:

	//- \begin{enumerate}\itemsep1pt
	//- \setcounter{enumi}{3}
	//- \item Program Sequencing, UI and Internetworking (in \texttt{msfs.go})
	//- \end{enumerate}

	//- \definecolor{Color2}{cmyk}{.07,.01,0.1,.01}
	//- \lstset{backgroundcolor=\color{Color2}}

//- \input{msfsfiles/msfsfiles.go}

	//- \definecolor{Color3}{rgb}{.77,.9,.9}
	//- \lstset{backgroundcolor=\color{Color3}}


//- \input{msfshashes/msfshashes.go}

	//- \definecolor{Color4}{rgb}{.9,.88,.9}
	//- \lstset{backgroundcolor=\color{Color4}}


//- \input{msfscas/msfscas.go}

	//- \section{Program Flow}
	//--------------------------------------------------

	//- \definecolor{Color5}{rgb}{.8,.88,1}
	//- \lstset{backgroundcolor=\color{Color5}}


	//- The main program begins here::

// msfs (in msfs.go) implements a mutually suspicious file system
package main
//=

	//- In addition to the standard \texttt{go} library, \texttt{msfs.go} relies on functionality provided 
	//- by the \texttt{fuse} and \texttt{directio} packages and includes the packages defined earlier in Section 4:

// dependencies

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net/http"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

	"rputbl.com/msfs/msfsfiles"
	"rputbl.com/msfs/msfshashes"
	"rputbl.com/msfs/msfscas"
)
//=

	//- The \texttt{msfs} main executable has several sections:  

	//- \begin{itemize}\itemsep1pt
	//- \item Startup and Shutdown
	//- \item The File System Service Loop
	//- \item The Networking Service Loop
	//- \item The Block Storage Service Loop
	//- \end{itemize}

	//- \subsection{Startup and Shutdown}
	//--------------------------------------------------



	//-  In startup, the \texttt{msfs} configuration is assisted by the \texttt{flag} package, which needs a Usage function
	//- to instruct the user on options and parameters:

// Usage prints invocation options
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	a := "mountpoint  blockdevice host port rsafile"
	fmt.Fprintf(os.Stderr, " %s %s\n", os.Args[0],a)
	flag.PrintDefaults()
}
//=

	//- The \texttt{msfs} executable checks on startup to see if it has the right number of parameters, initializes the subsystems,
	//- then passes execution in parallel to the file system, hash services, communications, and block storage subsystems. Upon return from 
	//- any of these systems the rest are cleaned up and the program terminates.

// main checks parameters and initializes services
func main() {
	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() != 5{
			Usage()
			os.Exit(2)
	}
	mountpoint := flag.Arg(0)
	blockdevice := flag.Arg(1)
	host := flag.Arg(2)
	port := flag.Arg(3)
	rsafile := flag.Arg(4)

	casreq, casans:= make(chan *msfscas.CasReq), make(chan *msfscas.CasAns)

	blockstoreserver := new(msfscas.CASS)
	if err:=blockstoreserver.Prepare(blockdevice,host,port,casreq,casans); err != nil {
		log.Fatal(err)
	}

	blockstoreclient := new(msfscas.CASC)
	if err:=blockstoreclient.Prepare(*blockstoreserver); err != nil {
		log.Fatal(err)
	}

	hashinterface := new(msfshashes.HSC)
	if err:=hashinterface.Prepare( rsafile ); err != nil {
		log.Fatal(err)
	}

	filesystem := new(msfsfiles.FS)
	if err:=filesystem.Prepare(hashinterface, blockstoreclient ); err != nil {
		log.Fatal(err)
	}

	go serveFS(mountpoint,filesystem)

	go serveNET(blockdevice,host,port)

	blockstoreserver.Serve()

}

//=

	//- \subsection{The File System Service Loop}
	//--------------------------------------------------

// serveFS is a process for each user that converts file system operations into hash operations
func serveFS(mountpoint string, filesystem *msfsfiles.FS ) {


	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Fatal(err)
	}

	fs.Serve(c, filesystem)
}
//=


	//- \subsection{The Networking Service Loop}
	//--------------------------------------------------

// serveNET extends the programs services over the network
func serveNET(blockdevice, host, port string) {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
//=



	//--------------------------------------------------
	//- \section{File System Usage}
	//--------------------------------------------------

	//- In ordinary usage a user will orient the \texttt{msfs} file system in the \texttt{sec} subdirectory of their home directory,
	//- e.g. \texttt{/home/chuck/sec}.  
	//-
	//- To unmount the filesystem: \texttt{fusermount -u mountdir}

	//--------------------------------------------------
	//- \section{Literate Coding Conventions}
	//--------------------------------------------------

	//- \subsection{Obtaining the \texttt{msfs} source code}
	//--------------------------------------------------

	//- \subsection{Creating the \texttt{msfs} executable}
	//--------------------------------------------------

	//- \subsection{Generating the \texttt{msfs} document}
	//--------------------------------------------------

	//- \subsection{Reporting \texttt{msfs} bugs}
	//--------------------------------------------------

	//- \subsection{Suggested exercises in extending \texttt{msfs}}
	//--------------------------------------------------



	//-\end{document}


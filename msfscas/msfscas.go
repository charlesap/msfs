	//- \subsection{Content Aware Storage}
	//--------------------------------------------------

	//- The file system structure maintains information needed to map file system operations (exposed by FUSE to the \texttt{msfs} executable)
	//- onto  assertions and queries made via TCP/IP to a local or remote content-aware hash store.

	//- The /texttt{msfsfs/msfsblocks.go} file implements the content aware storage functionality of \texttt{msfs}:

// msfscas (in msfscas.go) implements msfs content-aware persistent storage
package msfscas
//=

	//- In addition to the standard \texttt{go} library, \texttt{fs.go} relies on functionality provided 
	//- by the \texttt{fuse} package:

// dependencies

import (
	"fmt"
	"os"
	"io"
	"syscall"

	"github.com/ncw/directio"
)
//=

	//- \subsubsection{CAS Client Structures and Methods}
	//--------------------------------------------------


// CASC implements the block store client
type CASC struct{
	Cass CASS
	Status string
	Mode string
}
//=


	//- Status may be disconnected, unintroduced, or connected, mode may be direct, local, or remote.

// Prepare() initializes the Content Aware Storage Client
func (casc *CASC) Prepare(cass CASS)  error {
	casc.Cass=cass
	casc.Status="disconnected"
	casc.Mode="direct"
	return nil
}

type CasReq struct {
	Request      string
}
type CasAns struct {
	Answer      string
}

//=

	//- \subsubsection{CAS Server Structures and Methods}
	//--------------------------------------------------

// CASS implements the block store server
type CASS struct{
	Requests chan *CasReq
	Answers chan *CasAns
	BlockDevice string
	BlockHost string
	BlockPort string
	BlockFormat string
	BlockSize int64
	BlockMagic string
	BlockStat syscall.Stat_t
	Status string 
}
//=

	//-  

// Prepare must be called before using a Content Aware Storage Server
func (cass *CASS) Prepare(bd, bh, bp string, casreq chan *CasReq, casans chan *CasAns)  error {

	var  err error

	cass.Requests = casreq
	cass.Answers = casans
	cass.BlockDevice=bd
	cass.BlockHost=bh
	cass.BlockPort=bp
	cass.BlockFormat = "Unknown"
	cass.Status="disconnected"

	xfi, err := os.Stat(cass.BlockDevice)
	if err != nil {
		return err
	}
	cass.BlockSize = xfi.Size()

	in, err := directio.OpenFile(cass.BlockDevice, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	block := directio.AlignedBlock(BlockSize)
        _, err = io.ReadFull(in, block)
	if err != nil {
		return err
	}

	cass.BlockMagic = fmt.Sprintf("%s ",block[0:10])

	err = syscall.Stat(cass.BlockDevice,&cass.BlockStat)
	if err != nil {
		return err
	}
	fmt.Printf("Size: %d Blocksize: %d\n",cass.BlockStat.Size,cass.BlockStat.Blksize)

	fmt.Printf("OK\n")




	return err
}

func (cass *CASS) Serve() {

	for s:= range cass.Requests{
		fmt.Printf("Block request: %s\n",s.Request)
	}

}

const (
    AlignSize = 512
    BlockSize = 512
)

//=



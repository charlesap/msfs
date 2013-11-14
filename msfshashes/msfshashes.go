	//- \subsection{Hashing and Encryption}
	//--------------------------------------------------

	//- The file system structure maintains information needed to map file system operations (exposed by FUSE to the \texttt{msfs} executable)
	//- onto  assertions and queries made via TCP/IP to a local or remote  hash store.

	//- The /texttt{msfsfs/msfshashes.go} file implements the hashing functionality of \texttt{msfs}:

// msfshashes (in msfshashes.go) implements msfs hash functionality
package msfshashes
//=

	//- In addition to the standard \texttt{go} library, \texttt{fs.go} relies on functionality provided 
	//- by the \texttt{fuse} package:

// dependencies

import (
	"bazil.org/fuse"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)
//=

	//- \subsubsection{Hash Store Client Structures and Methods}
	//--------------------------------------------------


// HSC implements the hash store client
type HSC struct{
	PrivKey	*rsa.PrivateKey
	PubCert	[]byte
}

func (hsc *HSC) Prepare(rsaName string)  error {
	var err error
	hsc.PrivKey, hsc.PubCert, err  = GetPKI( rsaName )
	return err
}

//=

	//- \subsubsection{Hash Store Server Structures and Methods}
	//--------------------------------------------------

// HSS implements the hash store server
type HSS struct{}

func (hss *HSS) Prepare(destination string)  fuse.Error {
	return nil
}
//=


	//- \subsubsection{Miscellaneous Structures and Methods}
	//--------------------------------------------------



// Sha224base64 performs a sha224 hash on a byte array and then perfroms a base64 encoding on the result.
func Sha224base64(item []byte) (string, []byte) {

	phash := sha256.New224()
	io.WriteString(phash, string(item))
	phashbytes := phash.Sum(nil)
	return base64.StdEncoding.EncodeToString(phashbytes), phashbytes
}
//=

// Un64 decodes a base-64 encoded hash string and returns a byte array or an error.
func Un64( hash64val string) ([]byte, error){
	return base64.StdEncoding.DecodeString(hash64val)
}
//=

// Sign64 signs a byte array with a private key.
func Sign64(rsakey *rsa.PrivateKey, item []byte) (string, []byte) {

	hashFunc := crypto.SHA1
	h := hashFunc.New()
	h.Write(item)
	digest := h.Sum(nil)
	signresult, _ := rsa.SignPKCS1v15(rand.Reader, rsakey, hashFunc, digest)
	return base64.StdEncoding.EncodeToString(signresult), signresult
}
//=

// GetPKI retrieves 'rputn' RSA public and private key values
func GetPKI( rsaName string ) (*rsa.PrivateKey, []byte, error) {

	rsa_file := fmt.Sprintf("%s/.ssh/%s", os.Getenv("HOME"),rsaName)
	rsapub_file := fmt.Sprintf("%s/.ssh/%s.pub", os.Getenv("HOME"),rsaName)

	_, err := os.Stat(rsa_file)
	if err == nil {
		_, err = os.Stat(rsapub_file)
	}
	if err != nil {
		estr:="Please generate a reputation public/private key pair\n" +
		"(ssh-keygen -t rsa -C \"<user>@<host>\" -f ~/.ssh/<filename>)\n"+
		"rsa_file: "+rsa_file+"\n"+
		"rsapub_file: "+rsapub_file+"\n"
		return nil, nil, errors.New(estr)
	}

	rputn_rsa, _ := ioutil.ReadFile(rsa_file)
	rputn_rsa_pub, _ := ioutil.ReadFile(rsapub_file)
	block, _ := pem.Decode(rputn_rsa)
	rsakey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return rsakey, rputn_rsa_pub, nil
}
//=



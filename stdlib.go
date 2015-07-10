package main

var stdlibImports = `

import (
 "archive/tar"
 "archive/zip"
 "bufio"
 "bytes"
 "compress/bzip2"
 "compress/flate"
 "compress/gzip"
 "compress/lzw"
 "compress/zlib"
 "container/heap"
 "container/list"
 "container/ring"
 "crypto"
 "crypto/aes"
 "crypto/cipher"
 "crypto/des"
 "crypto/dsa"
 "crypto/ecdsa"
 "crypto/elliptic"
 "crypto/hmac"
 "crypto/md5"
 cryptorand "crypto/rand"
 "crypto/rc4"
 "crypto/rsa"
 "crypto/sha1"
 "crypto/sha256"
 "crypto/sha512"
 "crypto/subtle"
 "crypto/tls"
 "crypto/x509"
 "database/sql"
 "debug/dwarf"
 "debug/elf"
 "debug/gosym"
 "debug/macho"
 "debug/pe"
 "debug/plan9obj"
 "encoding"
 "encoding/ascii85"
 "encoding/asn1"
 "encoding/base32"
 "encoding/base64"
 "encoding/binary"
 "encoding/csv"
 "encoding/gob"
 "encoding/hex"
 "encoding/json"
 "encoding/pem"
 "encoding/xml"
 "errors"
 "expvar"
 "flag"
 "fmt"
 "go/ast"
 "go/build"
 "go/constant"
 "go/doc"
 "go/format"
 "go/importer"
 "go/parser"
 "go/printer"
 "go/scanner"
 "go/token"
 "go/types"
 "hash"
 "hash/adler32"
 "hash/crc32"
 "hash/crc64"
 "hash/fnv"
 "html"
 "html/template"
 "image"
 "image/color"
 "image/draw"
 "image/gif"
 "image/jpeg"
 "image/png"
 "index/suffixarray"
 "io"
 "io/ioutil"
 "log"
 "log/syslog"
 "math"
 "math/big"
 "math/cmplx"
 "math/rand"
 "mime"
 "mime/multipart"
 "mime/quotedprintable"
 "net"
 "net/http"
 "net/mail"
 "net/rpc"
 "net/smtp"
 "net/textproto"
 "net/url"
 "os"
 "os/exec"
 "os/signal"
 "os/user"
 "path"
 "path/filepath"
 "reflect"
 "regexp"
 "regexp/syntax"
 "runtime"
 "runtime/cgo"
 "runtime/debug"
 "runtime/pprof"
 "runtime/race"
 "sort"
 "strconv"
 "strings"
 "sync"
 "sync/atomic"
 "syscall"
 "testing"
 "testing/iotest"
 "testing/quick"
 textscanner "text/scanner"
 "text/tabwriter"
 texttemplate "text/template"
 "time"
 "unicode"
 "unicode/utf16"
 "unicode/utf8"
 "unsafe"
)

 // for scripting; avoid compile errors if unused libs.

 var _ = tar.TypeReg
 var _ = zip.Store
 var _ = bufio.MaxScanTokenSize
 var _ = bytes.MinRead
 var _ = bzip2.NewReader
 var _ = flate.NewReader
 var _ = gzip.NewReader
 var _ = lzw.NewReader
 var _ = zlib.NewReader
 var _ = heap.Fix
 var _ = list.New
 var _ = ring.New
 var _ = crypto.RegisterHash
 var _ = aes.BlockSize
 var _ = cipher.NewGCM
 var _ = des.BlockSize
 var _ = dsa.GenerateKey
 var _ = ecdsa.Sign
 var _ = elliptic.GenerateKey
 var _ = hmac.New
 var _ = md5.New
 var _ = cryptorand.Read
 var _ = rc4.NewCipher
 var _ = rsa.PSSSaltLengthAuto
 var _ = sha1.New
 var _ = sha256.New
 var _ = sha512.New
 var _ = subtle.ConstantTimeByteEq
 var _ = tls.Listen
 var _ = x509.CreateCertificate
 var _ = sql.Drivers
 var _ = dwarf.AttrSibling
 var _ = elf.EI_CLASS
 var _ = gosym.NewLineTable
 var _ = macho.Magic32
 var _ = pe.IMAGE_FILE_MACHINE_UNKNOWN
 var _ = plan9obj.NewFile
 var _ = encoding.BinaryMarshaler(nil)
 var _ = ascii85.Decode
 var _ = asn1.Marshal
 var _ = base32.NewDecoder
 var _ = base64.NewDecoder
 var _ = binary.MaxVarintLen16
 var _ = csv.NewReader
 var _ = gob.Register
 var _ = hex.Decode
 var _ = json.NewDecoder
 var _ = pem.Encode
 var _ = xml.Unmarshal
 var _ = errors.New
 var _ = expvar.Do
 var _ = flag.Arg
 var _ = fmt.Printf
 var _ = ast.Walk
 var _ = constant.New
 var _ = doc.New
 var _ = format.New
 var _ = importer.New
 var _ = parser.New
 var _ = printer.New
 var _ = scanner.PrintError
 var _ = token.New
 var _ = types.New
 var _ = hash.New
 var _ = adler32.New
 var _ = crc32.New
 var _ = crc64.New
 var _ = fnv.New
 var _ = html.New
 var _ = template.New
 var _ = image.New
 var _ = color.New
 var _ = draw.New
 var _ = gif.New
 var _ = jpeg.New
 var _ = png.New
 var _ = suffixarray.New
 var _ = io.Copy
 var _ = ioutil.ReadAll
 var _ = log.New
 var _ = syslog.New
 var _ = math.New
 var _ = big.New
 var _ = cmplx.New
 var _ = rand.New
 var _ = mime.New
 var _ = multipart.New
 var _ = quotedprintable.New
 var _ = net.New
 var _ = http.New
 var _ = mail.New
 var _ = rpc.New
 var _ = smtp.New
 var _ = textproto.New
 var _ = url.New
 var _ = os.Chdir
 var _ = exec.Command
 var _ = signal.New
 var _ = user.New
 var _ = path.Base
 var _ = filepath.Abs
 var _ = reflect.New
 var _ = regexp.New
 var _ = syntax.New
 var _ = runtime.New
 var _ = cgo.New
 var _ = debug.New
 var _ = pprof.New
 var _ = race.New
 var _ = sort.New
 var _ = strconv.New
 var _ = strings.New
 var _ = sync.New
 var _ = atomic.New
 var _ = syscall.Accept
 var _ = testing.New
 var _ = iotest.New
 var _ = quick.New
 var _ = textscanner.ScanIdents
 var _ = tabwriter.New
 var _ = texttemplate.New
 var _ = time.Now
 var _ = unicode.New
 var _ = utf16.New
 var _ = utf8.New
 var _ = unsafe.Pointer(nil)
 // above for scripting. Remove _ assigns in production code.New

` // end of var stdlibImports

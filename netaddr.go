package salticidae

// #include <stdlib.h>
// #include "salticidae/netaddr.h"
import "C"
import "runtime"

// The C pointer type for a NetAddr object
type CNetAddr = *C.netaddr_t
type netAddr struct { inner CNetAddr }
// Network address object.
type NetAddr = *netAddr

// Convert an existing C pointer into a go object. Notice that when the go
// object does *not* own the resource of the C pointer, so it is only valid to
// the extent in which the given C pointer is valid. The C memory will not be
// deallocated when the go object is finalized by GC. This applies to all other
// "FromC" functions.
func NetAddrFromC(ptr CNetAddr) NetAddr {
    return &netAddr{ inner: ptr }
}

type netAddrArray struct {
    inner *C.netaddr_array_t
}

// An array of network address.
type NetAddrArray = *netAddrArray

// Create NetAddr from a TCP socket format string (e.g. 127.0.0.1:8888).
func NewAddrFromIPPortString(addr string) (res NetAddr) {
    c_str := C.CString(addr)
    res = &netAddr{ inner: C.netaddr_new_from_sipport(c_str) }
    C.free(rawptr_t(c_str))
    runtime.SetFinalizer(res, func(self NetAddr) { self.free() })
    return
}

// Convert a Go slice of net addresses to NetAddrArray.
func NewAddrArrayFromAddrs(arr []NetAddr) (res NetAddrArray) {
    size := len(arr)
    if size > 0 {
        // FIXME: here we assume struct of a single pointer has the same memory
        // footprint the pointer
        base := (**C.netaddr_t)(rawptr_t(&arr[0]))
        res = &netAddrArray{ inner: C.netaddr_array_new_from_addrs(base, C.size_t(size)) }
    } else {
        res = &netAddrArray{ inner: C.netaddr_array_new() }
    }
    runtime.SetFinalizer(res, func(self NetAddrArray) { self.free() })
    return
}

func (self NetAddr) free() { C.netaddr_free(self.inner) }

// Check if two addresses are the same.
func (self NetAddr) IsEq(other NetAddr) bool { return bool(C.netaddr_is_eq(self.inner, other.inner)) }

func (self NetAddr) IsNull() bool { return bool(C.netaddr_is_null(self.inner)) }

// Get the 32-bit IP representation.
func (self NetAddr) GetIP() uint32 { return uint32(C.netaddr_get_ip(self.inner)) }

// Get the 16-bit port number (in UNIX network byte order, so need to apply
// ntohs(), for example, to convert the returned integer to the local endianness).
func (self NetAddr) GetPort() uint16 { return uint16(C.netaddr_get_port(self.inner)) }

// Make a copy of the object. This is required if you want to keep the NetAddr
// returned (or passed as a callback parameter) by other salticidae objects
// (such like MsgNetwork/PeerNetwork), unless the function returns a moved object.
func (self NetAddr) Copy() NetAddr {
    res := &netAddr{ inner: C.netaddr_copy(self.inner) }
    runtime.SetFinalizer(res, func(self NetAddr) { self.free() })
    return res
}

func (self NetAddrArray) free() { C.netaddr_array_free(self.inner) }

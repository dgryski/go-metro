//go:build !noasm && amd64 && !gccgo && !tinygo
// +build !noasm,amd64,!gccgo,!tinygo

package metro

//go:generate python -m peachpy.x86_64 metro.py -S -o metro_amd64.s -mabi=goasm
//go:noescape

func Hash64(buffer []byte, seed uint64) uint64
func Hash64Str(buffer string, seed uint64) uint64

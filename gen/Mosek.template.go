/* MOSEK package.

   Thin wrapper for MOSEK solver API.
*/
package mosek

// /*<comment>*/
// // SEE https://github.com/golang/go/wiki/cgo#function-variables
//
// #include <stdlib.h>
// #include <stdint.h>
//
// typedef void * MSKuserhandle_t;
// typedef const int8_t * string_t;
// typedef void * MSKtask_t;
// typedef void * MSKenv_t;
//
//<extern>
import "C"

import (
	"bytes"
)

//<consts>

//<methods>

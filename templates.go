package main

var sources = map[string]string{
	"c": `

/* %s: XXX implementation. */


#include <sys/types.h>


`,
	"h": `

/* %s: XXX definitions. */


#ifndef GUARD
#define GUARD




#endif /* GUARD */
`,
	"fth": `
\ %s: definitions for XXX.

`,
	"cc": `

// %s: XXX implementation.


namespace NAMESPACE {




} // namespace NAMESPACE


`,
	"hh": `

// %s: XXX definitions.


#ifndef GUARD
#define GUARD


namespace NAMESPACE {



} // namespace NAMESPACE


#endif // GUARD
`,
}

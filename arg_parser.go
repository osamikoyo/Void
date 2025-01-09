package void

import "strings"

type ArgRouter struct {
	args      []string
	flags     map[string]string
	boolFlags map[string]bool
}

// NewArgRouter creates a new argument router
func NewArgRouter(args []string) *ArgRouter {
	router := &ArgRouter{
		args:      make([]string, 0),
		flags:     make(map[string]string),
		boolFlags: make(map[string]bool),
	}
	
	// Parse arguments and flags
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			// Handle long flags
			name := strings.TrimPrefix(arg, "--")
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				router.flags[name] = args[i+1]
				i++
			} else {
				router.boolFlags[name] = true
			}
		} else if strings.HasPrefix(arg, "-") {
			// Handle short flags
			name := strings.TrimPrefix(arg, "-")
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				router.flags[name] = args[i+1]
				i++
			} else {
				router.boolFlags[name] = true
			}
		} else {
			// Regular argument
			router.args = append(router.args, arg)
		}
	}
	
	return router
}
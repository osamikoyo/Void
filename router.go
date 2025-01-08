package void

type ArgRouter struct{
	Parameters map[string]string
	ArgList []string
}

func NewArgRouter(args []string) *ArgRouter {
	var router ArgRouter
	for i, arg := range args{
		if arg[0] == '-' {
			router.Parameters[arg] = args[i + 1]
		}
	}

	return &router
}
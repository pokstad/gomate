package gomate

import "context"

// Dep abstracts away a specific gomate dependency
type Dep interface {
	// Satisfied tells us if a dependency is satisfied
	Satisfied(Env) (bool, error)
	// Install will attempt to install the dependency
	Install(context.Context, Env) error
	// DepName is a unique identifier for a dependency
	DepName() string
}

// depReg is a global dependency registry for all tools needed by various
// gomate pkg's (yay, global state!)
var depReg = map[string]Dep{}

// MustRegisterDep raises a panic when a dependency ID is registered twice. This
// function should only be called in a package's init function.
func MustRegisterDep(d Dep) {
	_, ok := depReg[d.DepName()]

	if ok {
		panic(E("dependency already exists with name " + d.DepName()))
	}

	depReg[d.DepName()] = d
}

// UninstalledDeps returns a list of any uninstalled dependencies from the list
// provided. If no list is provided, all uninstalled deps will be returned.
func UninstalledDeps(deps []string, env Env) ([]Dep, error) {
	var uninstalled []Dep

	if len(deps) == 0 {
		for _, d := range depReg {
			deps = append(deps, d.DepName())
		}
	}

	for _, d := range deps {
		dep, ok := depReg[d]
		if !ok {
			return nil, E("dependency does not exist in registry: " + d)
		}

		satisfied, err := dep.Satisfied(env)
		if err != nil {
			return nil, PushE(err, "cannot determine if dep is satisfied")
		}

		if !satisfied {
			uninstalled = append(uninstalled, dep)
		}
	}

	return uninstalled, nil
}

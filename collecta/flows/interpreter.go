package flows

import (
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

var memoizedEvaluation map[string]string
var memoizedCounter map[string]uint64

const thresholdAlloc = 40
const thresholdUses = 5

type input struct {
	state          string
	externalInputs []string
}

func sanitizeInput(in *input) {
	in.state = strings.TrimSpace(in.state)

	for i, s := range in.externalInputs {
		in.externalInputs[i] = strings.TrimSpace(s)
	}
}

func cleanMemoized() {
	for k := range memoizedEvaluation {
		if memoizedCounter[k] < thresholdUses {
			delete(memoizedEvaluation, k)
			delete(memoizedCounter, k)
		}
	}
}

func memo(input input) string {
	if memoizedEvaluation == nil {
		memoizedEvaluation = map[string]string{}
	}

	if memoizedCounter == nil {
		memoizedCounter = map[string]uint64{}
	}

	in := []string{input.state}
	in = append(in, input.externalInputs...)

	chain := strings.Join(in, ".")

	val, exist := memoizedEvaluation[chain]
	if !exist {
		return ""
	}

	if memoizedCounter[chain] += 1; len(memoizedEvaluation) > thresholdAlloc {
		cleanMemoized()
	}

	return val
}

func evalProgram(ctx context.Context, program string, input input) (string, error) {
	sanitizeInput(&input) // TODO: Improve this

	// Memoized
	if mem := memo(input); mem != "" {
		return mem, nil
	}

	scr := tengo.NewScript([]byte(program))

	if err := scr.Add("state", input.state); err != nil {
		return "", errors.Wrap(err, "error at add state param to tengo script")
	}

	externalInputs := make([]tengo.Object, 0)
	for _, exIn := range input.externalInputs {
		obj, _ := tengo.FromInterface(exIn)
		externalInputs = append(externalInputs, obj)
	}

	if err := scr.Add("external", externalInputs); err != nil {
		return "", errors.Wrap(err, "error at add externals param to tengo script")
	}

	compiled, err := scr.RunContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "error at compile and run tengo script")
	}

	if !compiled.IsDefined("res") {
		return "", errors.New("your script not response with a standard variable: 'res'")
	}

	newState := compiled.Get("res").String()
	if _, err = uuid.Parse(newState); err != nil {
		return "", errors.Wrap(err, "invalid response, it must be an uuid")
	}

	return newState, nil
}

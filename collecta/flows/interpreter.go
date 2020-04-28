package flows

import (
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

var memoizedEvaluation map[string]flowResponse
var memoizedCounter map[string]uint64

const thresholdAlloc = 40
const thresholdUses = 5

type flowResponse struct {
	next string
	last string
}

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

func memo(input input) flowResponse {
	if memoizedEvaluation == nil {
		memoizedEvaluation = map[string]flowResponse{}
	}

	if memoizedCounter == nil {
		memoizedCounter = map[string]uint64{}
	}

	in := []string{input.state}
	in = append(in, input.externalInputs...)

	chain := strings.Join(in, ".")

	val, exist := memoizedEvaluation[chain]
	if !exist {
		return flowResponse{}
	}

	if memoizedCounter[chain] += 1; len(memoizedEvaluation) > thresholdAlloc {
		cleanMemoized()
	}

	return val
}


func evalProgram(ctx context.Context, program string, input input) (flowResponse, error) {
	sanitizeInput(&input) // TODO: Improve this

	// Memoized
	if flowResponse := memo(input); flowResponse.last != "" && flowResponse.next != "" {
		return flowResponse, nil
	}

	scr := tengo.NewScript([]byte(program))

	if err := scr.Add("state", input.state); err != nil {
		return flowResponse{}, errors.Wrap(err, "error at add state param to tengo script")
	}

	externalInputs := make([]tengo.Object, 0)
	for _, exIn := range input.externalInputs {
		obj, _ := tengo.FromInterface(exIn)
		externalInputs = append(externalInputs, obj)
	}

	if err := scr.Add("input", externalInputs); err != nil {
		return flowResponse{}, errors.Wrap(err, "error at add externals param to tengo script")
	}

	compiled, err := scr.RunContext(ctx)
	if err != nil {
		return flowResponse{}, errors.Wrap(err, "error at compile and run tengo script")
	}

	if !compiled.IsDefined("next") {
		return flowResponse{next: input.state, last: input.state}, nil
	}

	if !compiled.IsDefined("last") {
		return flowResponse{next: input.state, last: input.state}, nil
	}

	nextState := compiled.Get("next").String()
	if _, err = uuid.Parse(nextState); err != nil {
		return flowResponse{}, errors.Wrap(err, "invalid response, it must be an uuid")
	}

	lastState := compiled.Get("last").String()
	if _, err = uuid.Parse(lastState); err != nil {
		return flowResponse{}, errors.Wrap(err, "invalid response, it must be an uuid")
	}

	return flowResponse{
		next: nextState,
		last: lastState,
	}, nil
}

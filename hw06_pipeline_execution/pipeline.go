package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = executeStage(in, done, stage)
	}
	return in
}

func executeStage(in In, done In, stage Stage) Out {
	outChan := make(chan interface{})
	go func() {
		defer close(outChan)
		stageOut := stage(in)
		for {
			select {
			case <-done:
				return
			case data, ok := <-stageOut:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case outChan <- data:
				}
			}
		}
	}()
	return outChan
}

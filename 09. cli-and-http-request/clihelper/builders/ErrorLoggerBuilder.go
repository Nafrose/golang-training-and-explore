package builders

import "github.com/nafrose/exploring/clirunner/clihelper/interfaces"

type ErrorLoggerBuilder struct {
	*WritersCollectionBuilder
}

func (elb *ErrorLoggerBuilder) Add(w *clihelper.Writer) *ErrorLoggerBuilder {
	elb.WritersCollectionBuilder.WritersCollection.ErrorOutputLoggers =
		append(elb.WritersCollectionBuilder.WritersCollection.ErrorOutputLoggers, *w)

	return elb
}

func (elb *ErrorLoggerBuilder) AddMany(writers []*clihelper.Writer) *ErrorLoggerBuilder {
	for _, writer := range writers {
		elb.Add(writer)
	}

	return elb
}
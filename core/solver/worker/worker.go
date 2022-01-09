package worker

import "log"

// ErrorHandler Handles the errors sent on an error channel
// when the error channel is closed it sends the total number
// of errors handled on the given totalErrors channel
func ErrorHandler(errors chan error) {
	// process the errors on the channel
	for {
		// read from the errors channel
		err, ok := <-errors

		// test if the channel has been closed
		if !ok {
			return
		}

		log.Print(err)
	}
}

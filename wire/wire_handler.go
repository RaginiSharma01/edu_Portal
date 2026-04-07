package wire

import "smp/handler"

type Handlers struct {
	UserHandler      *handler.UserHandler
	ClassroomHandler *handler.ClassroomHandler
}

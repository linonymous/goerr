package log

import (
	"github.com/pkg/errors"
	"log"
)

func Log(err error){
	if err != nil{
		log.Println("%+v", errors.Wrap(err, err.Error()))
	}
}

func Fatal(err error)  {
	if err != nil{
		panic(errors.Wrap(err, err.Error()))
	}
}

func Recover(errs ...*error) {
	var e *error
	for _, err := range errs {
		e = err
		break
	}
	//handle panic
	if r := recover(); r != nil {
		var errmsg error
		//Preserve error which might have happened before panic/recover
		if e != nil && *e != nil {
			errmsg = errors.Wrap(*e, r.(error).Error())
		} else {
			//No error occurred just add a stacktrace
			errmsg = errors.Wrap(r.(error), "")
		}
		//If error can't bubble up -> Log it
		if e != nil {
			*e = errmsg
		} else {
			//log.Printf("%+v", errmsg)
		}
	}
}


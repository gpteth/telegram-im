package databus_util

import (
	"sync"

	"github.com/pkg/errors"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/pkg/log"
)

func NewDatabusHandler() *DatabusHandler {
	return &DatabusHandler{
		closeCh:    make(chan struct{}),
		databusMap: make(map[*databus.Databus]int, 10),
	}
}

type DatabusHandler struct {
	closeCh    chan struct{}
	wg         sync.WaitGroup
	databusMap map[*databus.Databus]int
	lock       sync.Mutex
}

func (s *DatabusHandler) Close() {
	// close all databus
	s.lock.Lock()
	for k, v := range s.databusMap {
		log.Info("closing databus, %v, routine=%d", k, v)
		k.Close()
	}
	s.lock.Unlock()
	close(s.closeCh)
	s.wg.Wait()
}

func (s *DatabusHandler) incWatch(bus *databus.Databus) {
	s.lock.Lock()
	s.databusMap[bus] = s.databusMap[bus] + 1
	s.lock.Unlock()
}

func (s *DatabusHandler) decWatch(bus *databus.Databus) {
	s.lock.Lock()
	s.databusMap[bus] = s.databusMap[bus] - 1
	s.lock.Unlock()
}

func (s *DatabusHandler) Watch(bus *databus.Databus, handler func(msg *databus.Message) error) {
	defer func() {
		s.wg.Done()
		s.decWatch(bus)
		if r := recover(); r != nil {
			r = errors.WithStack(r.(error))
			log.Error("Runtime error caught, try recover: %+v", r)
			s.wg.Add(1)
			go s.Watch(bus, handler)
		}
	}()
	var msgs = bus.Messages()
	s.incWatch(bus)
	log.Info("start watch databus, %+v", bus)
	for {
		var (
			msg *databus.Message
			ok  bool
			err error
		)
		select {
		case msg, ok = <-msgs:
			if !ok {
				log.Error("s.archiveNotifyT.messages closed")
				return
			}
		case <-s.closeCh:
			log.Info("server close")
			return
		}
		msg.Commit()
		if err = handler(msg); err != nil {
			log.Error("handle databus error topic(%s) key(%s) value(%s) err(%v)", msg.Topic, msg.Key, msg.Value, err)
			continue
		}
	}
}

func (s *DatabusHandler) GoWatch(bus *databus.Databus, handler func(msg *databus.Message) error) {
	go s.Watch(bus, handler)
}

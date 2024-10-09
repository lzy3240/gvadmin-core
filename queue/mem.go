package queue

import "time"

//TODO
// 实现消息异步处理
// 内存消息队列: channel

type memClient struct {
	qmap map[string]*mqueue
}

type mqueue struct {
	name string
	ch   chan string
}

func newMemClient() *memClient {
	return &memClient{
		qmap: make(map[string]*mqueue),
	}
}

func (m *memClient) RegisterTopic(topic string) error {
	m.qmap[topic] = &mqueue{
		name: topic,
		ch:   make(chan string, 1024),
	}
	return nil
}

func (m *memClient) Publish(topic string, message string) error {
	m.qmap[topic].ch <- message
	return nil
}

func (m *memClient) Subscribe(topic string, f func(param string)) {
	//协程死锁
	//for value := range m.qmap[topic].ch {
	//	f(value)
	//}

	for {
		select {
		case value := <-m.qmap[topic].ch:
			f(value)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

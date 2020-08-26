/*
TODO : change kafka module confluent to sarama
*/

package confluent

import (
	"github.com/seheee/PDK/logic-core/domain/model"
	"github.com/seheee/PDK/logic-core/setting"
)

var kafkaConsumer *group

type group struct {
	c   *consumer
	out chan model.KafkaData
}

func NewKafkaConsumer() *group {
	if kafkaConsumer != nil {
		return kafkaConsumer
	}

	outBufSize := setting.KafkaSetting.ChanBufSize

	kafkaConsumer = &group{
		out: make(chan model.KafkaData, outBufSize),
	}

	kafkaConsumer.c = NewConsumer()
	go kafkaConsumer.c.run(kafkaConsumer.out)

	return kafkaConsumer
}

func (g *group) GetOutput() <-chan model.KafkaData {
	if g != nil {
		return g.out
	}
	return nil
}

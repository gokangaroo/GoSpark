package message

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

type KafkaServer struct {
	Host string
	Port int
}

var (
	wg sync.WaitGroup
)

/**
 * 同步生产者
 */
func (s *KafkaServer) ProducerSync() error {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型 返回一个分区器  该分区器随机选取一个分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待失败成功后的响应
	config.Producer.Return.Successes = true


	msg := &sarama.ProducerMessage{}
	msg.Topic = "geek-kafka-topic"
	msg.Value = sarama.StringEncoder("this is a test message")
	//使用代理地址和配置创建一个同步生产者
	producer,err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err != nil{
		fmt.Println("producer close err:",err)
		return err
	}
	defer producer.Close()

	pid,offset,err := producer.SendMessage(msg)
	if err != nil{
		fmt.Println("send message failed: ",err)
		return err
	}
	fmt.Printf("pid:%v offset:%\n",pid,offset)
	return err
}

/**
 * 消费者
 */
func (s *KafkaServer) Consumer() error {
	//根据代理地址和配置创建一个消费者
	consumer,err := sarama.NewConsumer([]string{"127.0.0.1:9092"},nil)
	if err != nil{
		return err
	}
	//Partitions(topic) 返回topic下的所有分区id
	partitionList,err := consumer.Partitions("geek-kafka-topic")
	if err != nil{
		return err
	}

	for partition := range partitionList{
		//ConsumePartition 根据topic 分区和给定的偏移量创建相应的分区消费者
		//如果该分区消费者已经消费了该消息 则返回error
		//sarama.offsetNew 表明了为最新的消息
		pc,err := consumer.ConsumePartition("geek-kafka-topic",int32(partition),sarama.OffsetNewest)
		if err != nil{
			return err
		}
		defer pc.AsyncClose()

		wg.Add(1)
		go func(sarama.PartitionConsumer){
			defer wg.Done()
			//Messages() 该方法返回一个消息类型的只读通道 由代理产生
			for msg := range pc.Messages(){
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic,msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}

	wg.Wait()
	_ = consumer.Close()
	return err
}
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var buildVersion string

func usage() {
	fmt.Println(usageText)
	os.Exit(0)
}

var usageText = `kinesis-bulk-loader [options...]

Options:
-stream string       (Required) Kinesis stream name
-region string       Region for Kinesis stream
                     By default "ap-northeast-1"
-k string            (Required) Partition key
-m string            (Required) Message payload to put into the stream
-c connections       Number of parallel simultaneous Kinesis session
                     By default 1; Must be more than 0
-n num-calls         Run for exactly this number of calls by each Kinesis session
                     By default 1; Must be more than 0
-r retry-num         Number fo Retry in each message send
                     By default 1; Must be more than 0
-endpoint-url string The URL to send the API request to
                     By default "", which mean the AWS SDK automatically determines the URL
-version             Prints out build version information
-verbose             Verbose option
-h                   help message
`

type KinesisDataStreamProducer struct {
	StreamName   string
	Region       string
	EndpointUrl  string
	PartitionKey string
	Message      string
	Connections  int
	NumCalls     int
	RetryNum     int
	Verbose      bool
}

func randomStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// https://stackoverflow.com/questions/47606761/repeat-code-if-an-error-occured
func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)
		fmt.Printf("retrying after error:%s\n", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func (c *KinesisDataStreamProducer) Run() {
	successCount := uint32(0)
	errorCount := uint32(0)
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= c.Connections; i++ {
		wg.Add(1)
		go c.startWorker(i, &wg, &successCount, &errorCount)
	}
	wg.Wait()

	duration := time.Since(startTime).Seconds()
	duration_ms := time.Since(startTime).Milliseconds()
	average_ms := duration_ms / (int64(successCount) + int64(errorCount))

	fmt.Println("-----------------------")
	fmt.Println("Kinesis Bulk Loader Summary")
	fmt.Println("-----------------------")
	fmt.Printf("Sent messages: %v\n", successCount)
	fmt.Printf("Errors: %v\n", errorCount)
	fmt.Printf("Duration (sec): %v\n", duration)
	fmt.Printf("Average (ms): %v\n", average_ms)
}

func (c *KinesisDataStreamProducer) startWorker(id int, wg *sync.WaitGroup, successCount *uint32, errorCount *uint32) {
	defer wg.Done()

	kc := getKinesisSession(c.Region, c.EndpointUrl)
	randomString := randomStr(10)
	message := c.Message + randomString
	for i := 1; i <= c.NumCalls; i++ {
		if c.Verbose {
			fmt.Printf("[Verbose] Mssage: PartitionKey %s Data %s\n", c.PartitionKey, message)
		}
		err := retry(c.RetryNum, 2*time.Second, func() (err error) {
			pout, kcerr := kc.PutRecord(
				&kinesis.PutRecordInput{
					Data:         []byte(message),
					StreamName:   &c.StreamName,
					PartitionKey: &c.PartitionKey,
				})
			if c.Verbose {
				fmt.Printf("[Verbose] PutRecord Result: PartitionKey %s SequenceNumber %s ShardId %s\n",
					c.PartitionKey, *pout.SequenceNumber, *pout.ShardId)
			}
			return kcerr
		})

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			atomic.AddUint32(errorCount, 1)
			continue
		}

		atomic.AddUint32(successCount, 1)
	}
}

func getKinesisSession(region string, endpointUrl string) *kinesis.Kinesis {
	var sess *session.Session
	if endpointUrl != "" {
		sess = session.New(&aws.Config{Region: aws.String(region), Endpoint: aws.String(endpointUrl)})
	} else {
		sess = session.New(&aws.Config{Region: aws.String(region)})
	}
	return kinesis.New(sess)
}

func main() {

	var (
		streamName   string
		region       string
		endpointUrl  string
		partitionKey string
		message      string
		connections  int
		numCalls     int
		retryNum     int
		version      bool
		verbose      bool
	)

	flag.StringVar(&streamName, "stream", "", "(Required) Kinesis stream name")
	flag.StringVar(&region, "region", "ap-northeast-1", "Region for Kinesis stream")
	flag.StringVar(&endpointUrl, "endpoint-url", "", "The URL to send the API request to")
	flag.StringVar(&partitionKey, "k", "", "(Required) Partition Key")
	flag.StringVar(&message, "m", "", "(Required) Message payload to put into the stream")
	flag.IntVar(&connections, "c", 1, "Number of parallel simultaneous Kinesis session")
	flag.IntVar(&numCalls, "n", 1, "Run for exactly this number of calls by each Kinesis session")
	flag.IntVar(&retryNum, "r", 1, "Number fo Retry in each message send")
	flag.BoolVar(&version, "version", false, "Build version")
	flag.BoolVar(&verbose, "verbose", false, "Verbose option")
	flag.Usage = usage
	flag.Parse()

	if version {
		fmt.Printf("version: %s\n", buildVersion)
		os.Exit(0)
	}

	if streamName == "" || partitionKey == "" || message == "" {
		fmt.Println("[ERROR] Invalid Command Options! Minimum required options are \"-stream\", \"-k\" and \"-m\"")
		usage()
	}

	s := KinesisDataStreamProducer{
		StreamName:   streamName,
		Region:       region,
		EndpointUrl:  endpointUrl,
		PartitionKey: partitionKey,
		Message:      message,
		Connections:  connections,
		NumCalls:     numCalls,
		RetryNum:     retryNum,
		Verbose:      verbose,
	}

	s.Run()
}

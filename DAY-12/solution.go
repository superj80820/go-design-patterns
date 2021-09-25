package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SendInfo struct {
	NewsName   string
	FinishTime time.Time
}

type NewsSender struct {
	inputsChan  chan string
	jobsChan    chan string
	resultsChan chan SendInfo
	wg          *sync.WaitGroup
}

type NewsCompleteLogger struct {
	completeNews <-chan SendInfo
	done         chan bool
}

func (c *NewsSender) StartConsume(ctx context.Context) {
	for {
		select {
		case in := <-c.inputsChan:
			if ctx.Err() != nil {
				close(c.jobsChan)
				return
			}
			c.jobsChan <- in
		case <-ctx.Done():
			close(c.jobsChan)
			return
		}
	}
}

func (c *NewsSender) StartWorker(ctx context.Context, id int) {
	defer c.wg.Done()
	fmt.Printf("create worker %d\n", id)
	time.Sleep(time.Duration(3 * time.Second)) //模擬與某service建立socket的時間
	for {
		select {
		case job, ok := <-c.jobsChan:
			if !ok {
				fmt.Printf("close worker %d\n", id)
				return
			}
			time.Sleep(time.Duration(3 * time.Second)) //模擬推播運行的時間
			fmt.Printf("<<worker %d finish %s>>\n", id, job)
			c.resultsChan <- SendInfo{job, time.Now()}
		case <-ctx.Done():
			fmt.Printf("close worker %d\n", id)
			return
		}
	}
}

func (c *NewsSender) CreateWorkerPool(ctx context.Context, poolSize int) {
	c.wg = &sync.WaitGroup{}
	c.wg.Add(poolSize)
	for w := 0; w < poolSize; w++ {
		go c.StartWorker(ctx, w)
	}
}

func (c *NewsSender) StopWait(ctx context.Context) {
	c.wg.Wait()
	close(c.resultsChan)
}

func CreateNewsSender(ctx context.Context) *NewsSender {
	newsSender := NewsSender{
		inputsChan:  make(chan string),
		jobsChan:    make(chan string),
		resultsChan: make(chan SendInfo),
	}
	return &newsSender
}

func ProduceToNewsSender(allNews []string, inputsChan chan<- string) {
	for _, news := range allNews {
		fmt.Printf("<<producer send %s>>\n", news)
		inputsChan <- news
	}
}

func CreateNewsCompleteLogger(ctx context.Context, completeNews <-chan SendInfo) *NewsCompleteLogger {
	newsCompleteLogger := NewsCompleteLogger{
		completeNews: completeNews,
		done:         make(chan bool),
	}
	return &newsCompleteLogger
}

func (n *NewsCompleteLogger) StartLog(ctx context.Context) {
	for result := range n.completeNews {
		fmt.Printf("<<fan in news>> news %s is sent at %s\n", result.NewsName, result.FinishTime)
	}
	close(n.done)
}

func (n *NewsCompleteLogger) StopWait(ctx context.Context) {
	<-n.done
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	newsSender := CreateNewsSender(ctx)
	newsCompleteLogger := CreateNewsCompleteLogger(ctx, newsSender.resultsChan)
	newsSender.CreateWorkerPool(ctx, 3)

	go ProduceToNewsSender([]string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
		"也不要吃太撐",
	}, newsSender.inputsChan)
	go ProduceToNewsSender([]string{
		"床前明月光",
		"疑似地上霜",
		"舉頭望明月",
		"低頭思故鄉",
	}, newsSender.inputsChan)
	go ProduceToNewsSender([]string{
		"渭城朝雨邑輕塵",
		"客舍青青柳色新",
		"勸君更盡一杯酒",
		"西出陽關無故人",
	}, newsSender.inputsChan)

	go newsSender.StartConsume(ctx)
	go newsCompleteLogger.StartLog(ctx)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	fmt.Println("================\nctrl+c event\n================")
	cancel()
	newsSender.StopWait(ctx)
	newsCompleteLogger.StopWait(ctx)
}

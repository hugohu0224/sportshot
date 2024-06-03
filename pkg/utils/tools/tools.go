package tools

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net"
	"sportshot/internal/crawler"
	"time"
)

func GetLocalHost() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// make sure IPv4
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not IPv4
			}

			return ip.String(), nil
		}
	}
	return "", errors.New("connected network is not found")
}

func TimeOutCtx(s int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(s)*time.Second)
}

func CrawlerTicker(crawler crawler.Crawler, timeInterval int) {
	zap.S().Infof("Start to crawl")
	events := crawler.Crawl()
	crawler.SaveToMongo(events)

	// make interval
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			events = crawler.Crawl()
			crawler.SaveToMongo(events)
		}
	}
}

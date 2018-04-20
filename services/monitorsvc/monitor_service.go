package monitorsvc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/WayneShenHH/toolsgo/repository"
)

type MonitorService struct {
	repository.Repository
}

// New instence MonitorService
func New(ctx repository.Repository, machineName, emojiName, channelName string, min int) *MonitorService {
	env = os.Getenv("GO_ENV")
	minute, machine, emoji, channel = min, machineName, emojiName, channelName
	return &MonitorService{
		Repository: ctx,
	}
}

var (
	minute  int
	machine string
	emoji   string
	channel string
	env     string
)

const (
	offerSvc      = "libgo_comparer,libgo_matchstat,libgo_offer,libgo_variant"
	envProduction = "production"
	slackURL      = "https://hooks.slack.com/services/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"
)

// Start 開始process check
func (service *MonitorService) Start() {
	workers := getWorker()
	alertChannel("start watching "+env+" services", "#libgo-offer")
	for {
		log.Println(fmt.Sprintf("Sleep %v minute ...", minute))
		time.Sleep(time.Duration(minute) * time.Minute)
		log.Println("Start checking worker")

		for _, w := range workers {
			checkProcess(w)
		}
	}
}

func checkProcess(worker string) {
	if serviceIsFailed(worker) {
		log.Println(fmt.Sprintf("%v failed", worker))
		msg := serviceStatus(worker)
		serviceRestart(worker)
		msg = msg + fmt.Sprintf("Restart %v\n", worker)
		alert(msg)
		if env == envProduction {
			go productionRestart(worker, 3)
		}
	} else {
		log.Println(fmt.Sprintf("%v active", worker))
		return
	}
}
func productionRestart(svc string, countdown int) {
	if strings.Index(svc, offerSvc) < 0 {
		return
	}
	time.Sleep(time.Second * 5)
	if !serviceIsFailed(svc) {
		return
	}
	countdown--
	if countdown < 0 {
		alertChannel(svc+" failed more than 3 times, close http:server.", "#libgo-offer")
		shutDownAPIServer()
		return
	}
	productionRestart(svc, countdown)
}
func serviceRestart(svc string) {
	cmdS := exec.Command("sudo", "systemctl", "restart", svc+".service")
	_, err := cmdS.Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(cmdS)
}
func serviceIsFailed(svc string) bool {
	c := exec.Command("sudo", "systemctl", "is-failed", svc+".service")
	o, _ := c.Output()
	state := strings.TrimSpace(string(o))
	return state == "failed"
}
func serviceStatus(svc string) string {
	cmdS := exec.Command("sudo", "systemctl", "status", svc+".service")
	o, _ := cmdS.Output()
	msg := fmt.Sprintf("%v is now not running.\n%v\n", svc, string(o))
	return msg
}
func shutDownAPIServer() {
	cmd := exec.Command("sudo", "systemctl", "stop", "libgo_httpapi.service")
	_, e := cmd.Output()
	if e != nil || !serviceIsFailed("libgo_httpapi") {
		shutDownAPIServer()
	}
}
func alert(msg string) {
	alertChannel(msg, channel)
}
func alertChannel(msg string, ch string) {
	payload := SlackPayload{
		Text:      msg,
		Username:  machine,
		Channel:   ch,
		IconEmoji: fmt.Sprintf(":%v:", emoji)}
	SendSlack(payload)
}
func getWorker() (s []string) {
	cmmd := exec.Command("sudo", "systemctl", "list-unit-files")
	out, err := cmmd.Output()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return convertSlice(out)
}

func convertSlice(out []byte) (s []string) {
	str := string(out)
	uasSlice := strings.Split(str, "\n")
	for _, uas := range uasSlice {
		if !valid(uas) {
			continue
		}
		svc := strings.Split(uas, ".service")[0]
		svc = strings.Trim(svc, " \t")
		s = append(s, svc)
	}
	return
}
func valid(uas string) bool {
	enable, _ := regexp.MatchString(".service\\s+enabled", uas)                                 // 跳過非enabled
	skip, _ := regexp.MatchString("(process_checker|sshd|syslog|service_monitor).service", uas) // 跳過自己 & 系統程序
	return enable && !skip
}
func SendSlack(payload SlackPayload) {
	jsonStr, _ := json.Marshal(payload)
	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("[response]", string(body))
}

type SlackPayload struct {
	Text      string `json:"text"`
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

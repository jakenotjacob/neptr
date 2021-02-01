package plugin

import (
	"log"
	"net/http"
	"os/exec"
	"sync"
)

type Plugin struct {
	Name  string                                       `json:"name"`
	Fn    func(w http.ResponseWriter, r *http.Request) `json:"fn"`
	State State
	mu    sync.Mutex
}

type State string

const (
	OK   State = "OK"
	ERR  State = "ERR"
	STOP State = "STOPPED"
	RUN  State = "RUNNING"
)

// A callback script must implement this interface
// to bind to a pattern
type Runnable interface {
	Status()
	Run()
}

func New(name string, fn func(w http.ResponseWriter, r *http.Request)) *Plugin {
	return &Plugin{
		Name:  name,
		Fn:    fn,
		State: STOP,
	}
}

func (p *Plugin) Register() {
	http.HandleFunc(p.Name, p.Fn)
}

func (p *Plugin) Status() State {
	return p.State
}

func (p *Plugin) Exec(args string) {
	p.mu.Lock()
	p.Run()
	log.Printf("Executing '%v'\n", args)

	cmd := exec.Command("sleep", args)
	cmd.Run()

	p.Stop()
	p.mu.Unlock()
}

func (p *Plugin) Run() {
	log.Printf("Running plugin: %v", p.Name)
	p.State = RUN
}

func (p *Plugin) Stop() {
	log.Printf("Running plugin: %v", p.Name)
	p.State = STOP
}

func init() {
}

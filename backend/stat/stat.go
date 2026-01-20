package stat

import (
	"context"
	"encoding/json"

	"optimus/backend/localstore"
)

const filename = "stats.json"

// Stat represents application statistics.
type Stat struct {
	ByteCount  int64 `json:"byteCount"`
	ImageCount int   `json:"imageCount"`
	TimeCount  int64 `json:"timeCount"`

	ctx        context.Context
	localStore *localstore.LocalStore
}

// NewStat returns a new Stat instance.
func NewStat() *Stat {
	s := &Stat{
		localStore: localstore.NewLocalStore(),
	}

	d, _ := s.localStore.Load(filename)
	_ = json.Unmarshal(d, &s)
	return s
}

// OnStartup is called when the app starts.
func (s *Stat) OnStartup(ctx context.Context) {
	s.ctx = ctx
}

// GetStats returns the application stats.
func (s *Stat) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"byteCount":  s.ByteCount,
		"imageCount": s.ImageCount,
		"timeCount":  s.TimeCount,
	}
}

// SetByteCount adds and persists the given byte count to the app stats.
func (s *Stat) SetByteCount(b int64) {
	if b <= 0 {
		return
	}
	s.ByteCount += b
	_ = s.store()
}

// SetImageCount adds and persists the given image count to the app stats.
func (s *Stat) SetImageCount(i int) {
	if i <= 0 {
		return
	}
	s.ImageCount += i
	_ = s.store()
}

// SetTimeCount adds and persists the given time count to the app stats.
func (s *Stat) SetTimeCount(t int64) {
	if t < 0 {
		return
	}
	s.TimeCount += t
	_ = s.store()
}

// store stores the app stats to the file system.
func (s *Stat) store() error {
	js, err := json.Marshal(s.GetStats())
	if err != nil {
		return err
	}
	if err = s.localStore.Store(js, filename); err != nil {
		return err
	}
	return nil
}

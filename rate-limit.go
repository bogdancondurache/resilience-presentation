perRequest := config.per / time.Duration(rate)
l := &atomicLimiter{
	perRequest: perRequest,
	maxSlack:   -1 * time.Duration(config.slack) * perRequest,
	clock:      config.clock,
}

initialState := state{
	last:     time.Time{},
	sleepFor: 0,
}

func (t *mutexLimiter) Take() time.Time {
	t.Lock()
	defer t.Unlock()
	now := t.clock.Now()
	if t.last.IsZero() {
		t.last = now
		return t.last
	}
	t.sleepFor += t.perRequest - now.Sub(t.last)
	if t.sleepFor < t.maxSlack {
		t.sleepFor = t.maxSlack
	}
	if t.sleepFor > 0 {
		t.clock.Sleep(t.sleepFor)
		t.last = now.Add(t.sleepFor)
		t.sleepFor = 0
	} else {
		t.last = now
	}
	return t.last
}


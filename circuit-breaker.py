def call(self, func, *args, **kwargs):
    if self.opened:
        raise CircuitBreakerError(self)
    try:
        result = func(*args, **kwargs)
    except self._expected_exceptions as e:
        self._call_failed()
        raise

    self._call_succeeded()
    return result

def _call_succeeded(self):
    self._state = self.STATE_CLOSED
    self._failure_count = 0

def _call_failed(self):
    self._failure_count += 1
    if self._failure_count >= self._failure_threshold:
        self._state = self.STATE_OPEN
        self._opened = datetime.utcnow()

@property
def state(self):
    if self._state == self.STATE_OPEN and self.open_remaining <= 0:
        return self.STATE_HALF_OPEN

    return self._state
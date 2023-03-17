package tcp

// per tcp_states.h

type State uint

const (
	TCP_UNKNOWN State = iota
	TCP_ESTABLISHED
	TCP_SYN_SENT
	TCP_SYN_RECV
	TCP_FIN_WAIT1
	TCP_FIN_WAIT2
	TCP_TIME_WAIT
	TCP_CLOSE
	TCP_CLOSE_WAIT
	TCP_LAST_ACK
	TCP_LISTEN
	TCP_CLOSING /* Now a valid state */
	TCP_NEW_SYN_RECV

	TCP_MAX_STATES /* Leave at the end! */
)

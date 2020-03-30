package irc

import (
	"net"
	"time"

	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

/*
Client contains information for our IRC bridge, including the girc Client
and the IRCSettings that were passed into NewClient
*/
type Client struct {
	*girc.Client
	Settings internal.IRCSettings
	logger   internal.DebugLogger
	sendToTg func(string)
}

/*
NewClient returns a new IRCClient based on the provided settings
*/
func NewClient(settings internal.IRCSettings, logger internal.DebugLogger) Client {
	logger.LogInfo("Creating new IRC bot client...")
	client := girc.New(girc.Config{
		Server: settings.Server,
		Port:   settings.Port,
		Nick:   settings.BotName,
		User:   settings.BotName,
	})
	if settings.NickServPassword != "" {
		client.Config.SASL = &girc.SASLPlain{
			User: settings.BotName,
			Pass: settings.NickServPassword,
		}
	}
	return Client{client, settings, logger, nil}
}

/*
StartBot adds necessary handlers to the client and then connects,
returns any errors that occur
*/
func (c Client) StartBot(errChan chan<- error, sendMessage func(string)) {
	c.logger.LogInfo("Starting up IRC bot...")
	c.sendToTg = sendMessage
	c.addHandlers()
	// 10 second timeout for connection
	if err := c.ConnectDialer(&net.Dialer{Timeout: 10 * time.Second}); err != nil {
		errChan <- err
		c.logger.LogError(err)
	} else {
		errChan <- nil
	}
}

func (c Client) AddHandler(eventType string, cb func(*girc.Client, girc.Event)) {
	c.Handlers.Add(eventType, cb)
}

func (c Client) ConnectDialer(dialer girc.Dialer) error {
	return c.DialerConnect(dialer)
}

func (c Client) Message(channel string, msg string) {
	c.Cmd.Message(channel, msg)
}

/*
SendMessage sends a message to the IRC channel specified in the
settings
*/
func (c Client) SendMessage(msg string) {
	c.Message(c.Settings.Channel, msg)
}

/*
addHandlers adds handlers for the client struct based on the settings
that were passed in to NewClient
*/
func (c Client) addHandlers() {
	for eventType, handler := range getHandlerMapping() {
		c.logger.LogDebug("Adding IRC event handler:", eventType)
		c.AddHandler(eventType, handler(c))
	}
}

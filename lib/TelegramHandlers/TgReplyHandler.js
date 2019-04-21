const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user sends a message containing a reply
 * to the Telegram Group.
 */
class TgReplyHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {*} ircConfig - IRC Configuration that specifies what the 
     *                        username's prefix and suffix shall be, and the 
     *                        maximum message length that can be sent through IRC.
     * @param {booleasn} enabled - Is this handler enabled?
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(ircConfig, enabled, action) {
        this._ircConfig = ircConfig;
        this.Enabled = enabled;
        this._action = action;
    }
    // ---------------- Functions ----------------

    /**
     * 
     * @param {*} from - Object that contains the information about the username.
     * @param {*} userMessage - The message the user sent that we want to relay.
     *                          Contains a Message object that user replied to.
     *                          
     */
    RelayMessage(from, userMessage) {
        const self = this;
        if (!self.Enabled) {
            return;
        }

        // bold username for IRC users
        let username = '\x02' + TgHelpers.ResolveUserName(from) + '\x02';
        let replyMessage = _GetReplyMessage(username, userMessage);
    }

    /**
     * @param {*} from - Object that contains information about the username
     * @param {*} userMessage - Message in which to get reply Message object from.
     * 
     */
    _GetReplyMessage(from, userMessage) {
        // get reply message object from userMessage
    }
}

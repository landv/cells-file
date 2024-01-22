/*
 * Copyright 2007-2017 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */
import React from 'react'
import PropTypes from 'prop-types'

const MessagesConsumerMixin = {
    contextTypes: {
        messages:PropTypes.object,
        getMessage:PropTypes.func
    }
};

const MessagesProviderMixin = {

    childContextTypes: {
        messages:PropTypes.object,
        getMessage:PropTypes.func
    },

    getChildContext: function() {
        var messages = this.props.pydio.MessageHash;
        return {
            messages: messages,
            getMessage: function(messageId, namespace='ajxp_admin'){
                try{
                    return messages[namespace + (namespace?".":"") + messageId] || messageId;
                }catch(e){
                    return messageId;
                }
            }
        };
    }

};

const PydioConsumerMixin = {
    contextTypes:{
        pydio:PropTypes.instanceOf(Pydio)
    }
};

const PydioProviderMixin = {
    childContextTypes:{
        pydio:PropTypes.instanceOf(Pydio)
    },

    getChildContext: function(){
        return {
            pydio: this.props.pydio
        };
    }
};

export {MessagesConsumerMixin, MessagesProviderMixin, PydioConsumerMixin, PydioProviderMixin};
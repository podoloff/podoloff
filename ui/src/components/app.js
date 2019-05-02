import { h, Component } from 'preact';
import { Router } from 'preact-router';

import Login from './login'
import Dashboard from './dashboard';
export default class App extends Component {

	/** Gets fired when the route changes.
	 *	@param {Object} event		"change" event from [preact-router](http://git.io/preact-router)
	 *	@param {string} event.url	The newly routed URL
	 */
    handleRoute = e => {
        this.currentUrl = e.url;
    };

    render() {
        return (
            <div id="app">
                {false ? // TODO: replace with login check
                    <Router onChange={this.handleRoute}>
                        <Dashboard path="/" />
                    </Router>
                    : <Login></Login>
                }
            </div>
        );
    }
}
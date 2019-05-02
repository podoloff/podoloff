import { h, Component } from 'preact';
import linkstate from 'linkstate'


class Login extends Component {
    render({ }, { email, password }) {
        return (
            <div>
                <h3>Login</h3>
                <hr />
                <div>
                    <h1>Dashboard</h1>
                    <div>
                        <h3>Email</h3>
                        <input type="text" value={email} onInput={linkstate(this, 'email')} />
                    </div>
                    <div>
                        <h3>Password</h3>
                        <input type="text" value={password} onInput={linkstate(this, 'password')} />
                    </div>

                    <div>
                        <br />
                        <input type="button" value="Submit" />
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;
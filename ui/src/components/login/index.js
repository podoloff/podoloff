import { h, Component } from 'preact';
import linkstate from 'linkstate'

import { login, isAuthed } from '../../util/auth'

class Login extends Component {

    state = { res: "" };

    Login() {
        login(this.state.email, this.state.password).then((res) => {
            this.setState({ res });
        });
    }

    Auth() {
        isAuthed().then((res) => {
            this.setState({ res });
        });
    }

    render({ }, { email, password, res }) {
        return (
            <div>
                <h3>Login</h3>
                <hr />
                {res}
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
                        <input type="button" value="Submit" onClick={() => this.Login()} />
                        <br />
                        <input type="button" value="Check" onClick={() => this.Auth()} />
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;
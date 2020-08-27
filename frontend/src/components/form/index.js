import React from 'react'

export class SignInForm extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            error_message: "",
            email: "",
            password: ""
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    render() {

        let { error_message, email, password } = this.state

        let message = null;
        if (error_message.length !== 0) {
            message = <h5 className="mb-4 text-danger">{error_message}</h5>;
        }

        return (
            <div>
                {message}
                <form onSubmit={this.handleSubmit}>
                    <h5 className="mb-4">Basic Info</h5>
                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input value={email} name="email" type="email" className="form-control" id="email" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="passwrord">Password:</label>
                        <input value={password} name="password" type="password" className="form-control" id="passwrord" onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-success btn-large">Sign In</button>
                        </div>
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-link text-info" onClick={() => this.props.handleNewUser()}> New User? Register</button>
                        </div>
                    </div>
                </form>
            </div>
        )
    }

    handleChange(event) {
        const name = event.target.name
        const value = event.target.value
        this.setState(
            {
                [name]: value
            }
        )
    }

    handleSubmit(event) {
        event.preventDefault()
        console.log(this.state)
    }
}

export class RegistrationForm extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            error_message: "",
            username: "",
            email: "",
            password: "",
            confirm_assword: "",
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    render() {

        const { username, email, password, confirm_assword, error_message } = this.state

        let message = null;
        if (error_message.length !== 0) {
            message = <h5 className="mb-4 text-danger">{error_message}</h5>;
        }

        return (
            <div>
                {message}
                <form onSubmit={this.handleSubmit}>
                    <h5 className="mb-4">Registeration</h5>
                    <div className="form-group">
                        <label htmlFor="username">User Name:</label>
                        <input value={username} id="username" name='username' className="form-control" placeholder='John Doe' type='text' onChange={this.handleChange} />
                    </div>

                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input value={email} type="email" name='email' className="form-control" id="email" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Password:</label>
                        <input value={password} type="password" name='password' className="form-control" id="password" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Confirm password:</label>
                        <input value={confirm_assword} type="password" name='confirm_password' className="form-control" id="confirm-password" onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-success btn-large">Register</button>
                        </div>
                    </div>
                </form>
            </div>
        )
    }

    handleChange(event) {
        const name = event.target.name
        const value = event.target.value
        this.setState(
            {
                [name]: value
            }
        )
    }

    handleSubmit(event) {
        event.preventDefault()
        console.log(this.state)
    }
}
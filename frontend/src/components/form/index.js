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
        fetch("/users/signin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: this.state.email,
                password: this.state.password
            })
        }).then(resp => resp.json()).
        then(result => console.log(result))

    }
}

export class RegistrationForm extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            error_message: "",
            firstname: "",
            lastname: "",
            email: "",
            password: "",
            confirm_password: "",
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    render() {

        const { firstname, lastname, email, password, confirm_password, error_message } = this.state

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
                        <label htmlFor="firstname">First Name:</label>
                        <input value={firstname} id="firstname" name='firstname' className="form-control" placeholder='John' type='text' onChange={this.handleChange} />
                    </div>

                    <div className="form-group">
                        <label htmlFor="lastname">Last Name:</label>
                        <input value={lastname} id="lastname" name='lastname' className="form-control" placeholder='Doe' type='text' onChange={this.handleChange} />
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
                        <input value={confirm_password} type="password" name='confirm_password' className="form-control" id="confirm-password" onChange={this.handleChange} />
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
        fetch("/users", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(this.state)
        }).then(resp => resp.json()).
        then(result => console.log(result))
    }
}
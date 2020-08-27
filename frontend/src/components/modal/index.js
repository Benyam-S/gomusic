import React from 'react'
import { Modal, ModalHeader, ModalBody } from 'reactstrap';
import CreditCardInformation from '../cerditcard'
import { SignInForm, RegistrationForm } from '../form'

export function BuyModalWindow(props) {
    return (
        <Modal id="buy" tabIndex="-1" role="dialog" isOpen={props.showModal} toggle={props.toggle}>
            <div role="document">
                <ModalHeader toggle={props.toggle} className="bg-success text-white">
                    Buy Item
                  </ModalHeader>
                <ModalBody>
                    {<CreditCardInformation show={true} operation="Charge" toggle={props.toggle} />}
                </ModalBody>
            </div>
        </Modal>
    );
}

export class RegistrationModalWindow extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            title: "Sign In",
        }

        this.handleNewUser = this.handleNewUser.bind(this)
    }

    render() {

        const { title } = this.state;

        return (
            <Modal id="buy" tabIndex="-1" role="dialog" isOpen={this.props.showModal} toggle={() => {
                this.setState({ title: "Sign In" })
                this.props.toggle()
            }} >
                <div role="document">
                    <ModalHeader toggle={() => {
                        this.setState({ title: "Sign In" })
                        this.props.toggle()
                    }} className="bg-success text-white">
                        {title}
                    </ModalHeader>
                    <ModalBody>
                        {title === "Sign In" ? <SignInForm handleNewUser={this.handleNewUser} /> : <RegistrationForm />}
                    </ModalBody>
                </div>
            </Modal>
        )
    }

    handleNewUser() {

        this.setState(
            { title: "Register" }
        )

    }
}
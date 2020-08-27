import React from 'react';
import { CardContainer, OrderContainer } from './components/container'
import { BuyModalWindow, RegistrationModalWindow } from './components/modal'
import { Navigation } from './components/navigation'
import { BrowserRouter as Router, Route } from "react-router-dom";
import About from './components/about'

class App extends React.Component {

  constructor(props) {
    super(props)

    this.state = {
      user: {
        loggedIn: false,
        name: ""
      },
      buyModalState: false,
      registrationModalState: false
    }

    this.showBuyModal = this.showBuyModal.bind(this)
    this.showModalWindow = this.showModalWindow.bind(this)
    this.toggleBuyModal = this.toggleBuyModal.bind(this)
    this.toggleRegistrationModal = this.toggleRegistrationModal.bind(this)

  }

  render() {

    const { buyModalState, registrationModalState, user } = this.state
    return (
      <div className="App container row">
        <Router>

          <Navigation user={user} showModalWindow={this.showModalWindow} />

          <div className="container pt-5 mt-4">

            <Route exact path="/">
              <CardContainer location="cards.json" promo={false} showBuyModal={this.showBuyModal} />
            </Route>

            <Route path="/promos">
              <CardContainer location="promos.json" promo={true} showBuyModal={this.showBuyModal} />
            </Route>

            {
              user.loggedIn ? <Route path="/myorders"> < OrderContainer /> </Route> : null
            }

            <Route path="/about">
              <About />
            </Route>

          </div>

        </Router>

        <BuyModalWindow showModal={buyModalState} toggle={this.toggleBuyModal} />
        <RegistrationModalWindow showModal={registrationModalState} toggle={this.toggleRegistrationModal} />

      </div>
    )
  }

  showBuyModal() {
    this.setState({
      buyModalState: true
    })
  }

  showModalWindow() {
    this.setState({
      registrationModalState: true
    })
  }

  toggleBuyModal() {
    let { buyModalState } = this.state
    this.setState({ buyModalState: !buyModalState })
  }

  toggleRegistrationModal() {
    let { registrationModalState } = this.state
    this.setState({ registrationModalState: !registrationModalState })
  }

}

export default App;

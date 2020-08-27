import React from "react";
import Card from "../card"
import Order from '../order'

export class CardContainer extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            items: []
        }
    }

    render() {

        const { items } = this.state

        return (
            <div className="row">

                {
                    items.map(
                        (item) => {
                            return <Card key={item.id} {...item} promo={this.props.promo} showBuyModal={this.props.showBuyModal} />
                        }
                    )
                }

            </div>
        )

    }

    componentDidMount() {
        fetch(this.props.location).
            then(resp => resp.json()).
            then(result => {
                this.setState({ items: result })
            })
    }
}

export class OrderContainer extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            orders: []
        };
    }

    render() {

        return (
            <div>
                {
                    this.state.orders.map(
                        order => {
                            return <Order key={order.id} {...order} />
                        }
                    )
                }
            </div>
        )
    }
}
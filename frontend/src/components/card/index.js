import React from 'react';

export default  function Card(props) {
    const priceColor = (props.promo) ? 'text-danger' : 'text-dark'
    return (<div key={props.id} className="col-md-6 col-lg-4 d-flex align-items-stretch">
        <div className="card mb-3">
            <img className="card-img-top" src={props.imgSrc} alt={props.imgAlt} />
            <div className="card-body">
                <h4 className="card-title">{props.productName}</h4>
            Price: <strong className={priceColor}>{props.price}</strong>
                <p className="card-text">{props.desc}</p>
                <a href="#" className="btn btn-primary" onClick={props.showBuyModal}>Buy</a>
            </div>
        </div>
    </div>)
}